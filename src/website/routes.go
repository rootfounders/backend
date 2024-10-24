package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/microcosm-cc/bluemonday"
	"github.com/rootfounders/rootfounders/database"
)

type ProjectForm struct {
	Owner           string
	ShortName       string
	Description     string
	DetailsLocation string
}

var policy = bluemonday.StrictPolicy()

func NewProjectForm(values *url.Values) *ProjectForm {
	return &ProjectForm{
		Owner:       values.Get("owner"),
		ShortName:   values.Get("shortName"),
		Description: values.Get("description"),
	}
}

func (p *ProjectForm) Validate() (errs map[string]string) {
	errs = make(map[string]string)
	// if !common.IsHexAddress(p.Owner) {
	// 	errs["owner"] = "invalid address"
	// }
	if len(p.ShortName) < 3 {
		errs["shortName"] = "value too short"
	}
	if len(p.ShortName) > 160 {
		errs["shortName"] = "value too long"
	}
	if len(p.Description) < 10 {
		errs["description"] = "empty or too short"
	}
	if len(errs) > 0 {
		return
	} else {
		return nil
	}
}

func InitRoutes(r *chi.Mux) {
	RegisterTemplateRoute(r, "/", "main")
	RegisterTemplateRoute(r, "/projects/new", "projects_new")

	r.Get("/projects", func(w http.ResponseWriter, r *http.Request) {
		projects, err := dbConn.ListProjects(context.TODO())
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			log.Println("listing projects:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		renderTemplate(w, "/projects", "projects", projects)
	})

	r.Post("/projects/new", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println("/projects/draft form parse error:", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		form := NewProjectForm(&r.Form)
		if errs := form.Validate(); errs != nil {
			log.Println("form errors", errs)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(map[string]any{"validation": errs}); err != nil {
				log.Println("encoding json:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		var details bytes.Buffer
		if _, err := details.WriteString(policy.Sanitize(form.Description)); err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		cid, err := Pin(&details)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("pinning", cid)

		if err := json.NewEncoder(w).Encode(map[string]any{"cid": cid}); err != nil {
			log.Println("encoding json:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	r.Get("/projects/{projectId}", func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")

		pid, _ := big.NewInt(0).SetString(projectId, 10)
		project, err := dbConn.GetProject(context.TODO(), pid)
		if err != nil {
			if err == pgx.ErrNoRows {
				data := struct {
					ProjectId     string
					DataAvailable bool
				}{
					ProjectId:     projectId,
					DataAvailable: false,
				}

				renderTemplate(w, "/project/{projectId}", "project", data)
				return
			}
			log.Println("getting project data from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// TODO: implement getting project data from database
		description, err := ReadFromIpfs(project.DetailsLocation)
		if err != nil {
			log.Println("reading from IPFS error", err, pid, project.DetailsLocation)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tipsTotal, _ := big.NewInt(0).SetString("1247932781501679000", 10)
		comments, err := dbConn.GetComments(context.TODO(), pid)
		if err != nil {
			log.Println("getting comments from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		commentsJson, _ := json.Marshal(comments)

		updates, err := dbConn.GetUpdates(context.TODO(), pid)
		if err != nil {
			log.Println("getting updates from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		updatesJson, _ := json.Marshal(updates)

		team, err := dbConn.GetTeam(context.TODO(), pid)
		if err != nil && err != pgx.ErrNoRows {
			log.Println("getting team from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			ProjectId     string
			DataAvailable bool
			ShortName     string
			Owner         string
			Description   string
			TipJar        string
			TipsTotal     *big.Int
			Comments      string
			Updates       string
			Team          []string
		}{
			// TOOD
			ProjectId:     projectId,
			DataAvailable: true,
			ShortName:     project.ShortName,
			Owner:         project.Owner,
			// description should really be JSON
			Description: string(description),
			TipJar:      project.TipJar,
			TipsTotal:   tipsTotal,
			Comments:    string(commentsJson),
			Updates:     string(updatesJson),
			Team:        team,
		}

		renderTemplate(w, "/project/{projectId}", "project", data)
	})

	r.Get("/projects/{projectId}/apply", func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")

		pid, _ := big.NewInt(0).SetString(projectId, 10)
		project, err := dbConn.GetProject(context.TODO(), pid)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Redirect(w, r, "/projects", http.StatusSeeOther)
				return
			}
			log.Println("getting project data from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			ProjectId string
			ShortName string
		}{
			// TOOD
			ProjectId: projectId,
			ShortName: project.ShortName,
		}

		renderTemplate(w, "/project/{projectId}/apply", "project_apply", data)
	})

	r.Post("/projects/{projectId}/apply", func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")

		pid, _ := big.NewInt(0).SetString(projectId, 10)
		_, err := dbConn.GetProject(context.TODO(), pid)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Redirect(w, r, "/projects", http.StatusSeeOther)
				return
			}
			log.Println("getting project data from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		payload := struct {
			Signature string `json:"signature"`
			Address   string `json:"address"`
			Value     string `json:"value"`
		}{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payload); err != nil {
			log.Println("post apply json error:", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		verified, err := verifySignature(payload.Address, payload.Signature, payload.Value)
		if err != nil {
			log.Println("invalid signatue", err, payload.Address, payload.Signature, payload.Value)
			http.Error(w, "Invalid signature", http.StatusBadRequest)
			return
		}
		if !verified {
			http.Error(w, "Cannot verify signature", http.StatusBadRequest)
			return
		}

		proposalValue := struct {
			ProjectId string `json:"projectId"`
			// teamApplicationProposal
			Value string `json:"value"`
		}{}

		if err := json.Unmarshal([]byte(payload.Value), &proposalValue); err != nil {
			log.Println("invalid proposal JSON", err)
			http.Error(w, "Invalid proposal", http.StatusBadRequest)
			return
		}

		if projectId != proposalValue.ProjectId {
			http.Error(w, "Invalid project", http.StatusBadRequest)
			return
		}

		err = dbConn.CreateProposal(context.TODO(), &database.ProposalRecord{
			ProjectId: projectId,
			Address:   payload.Address,
			Message:   proposalValue.Value,
		})
		if err != nil {
			log.Println("saving proposal to database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("added new proposal", projectId, payload.Address)
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/projects/{projectId}/proposals", func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")

		pid, _ := big.NewInt(0).SetString(projectId, 10)
		project, err := dbConn.GetProject(context.TODO(), pid)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Redirect(w, r, "/projects", http.StatusSeeOther)
				return
			}
			log.Println("getting project data from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			ProjectId string
			ShortName string
		}{
			ProjectId: projectId,
			ShortName: project.ShortName,
		}

		renderTemplate(w, "/project/{projectId}/apply", "project_proposals", data)
	})

	r.Get("/projects/{projectId}/proposals/list", func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")

		pid, _ := big.NewInt(0).SetString(projectId, 10)
		project, err := dbConn.GetProject(context.TODO(), pid)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Redirect(w, r, "/projects", http.StatusSeeOther)
				return
			}
			log.Println("getting project data from the database:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		signature := r.URL.Query().Get("signature")

		verified, err := verifySignature(project.Owner, signature, "connect wallet")
		if err != nil {
			log.Println("invalid signatue")
			http.Error(w, "Invalid signature", http.StatusBadRequest)
			return
		}
		if !verified {
			http.Error(w, "Cannot verify signature", http.StatusBadRequest)
			return
		}

		proposals, err := dbConn.GetProposals(context.TODO(), pid)
		if err != nil {
			log.Println("loading proposals:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(proposals); err != nil {
			log.Println("encoding proposals:", err, pid)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

}

func RegisterTemplateRoute(r *chi.Mux, route string, name string) {
	r.Get(route, func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, route, name, map[string]any{})
	})
}

func renderTemplate(w http.ResponseWriter, route string, name string, data any) {
	template, err := GetTemplate(name)
	if err != nil {
		log.Printf("getting template for route %s: %s\n", route, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := template.Execute(w, data); err != nil {
		log.Printf("template error for route %s: %s\n", route, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

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
)

type ProjectForm struct {
	Owner           string
	ShortName       string
	Description     string
	DetailsLocation string
}

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
		if _, err := details.WriteString(form.Description); err != nil {
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

		data := struct {
			ProjectId     string
			DataAvailable bool
			ShortName     string
			Owner         string
			Description   string
			TipsTotal     *big.Int
			Comments      string
			Updates       string
		}{
			// TOOD
			ProjectId:     projectId,
			DataAvailable: true,
			ShortName:     project.ShortName,
			Owner:         project.Owner,
			// description should really be JSON
			Description: string(description),
			TipsTotal:   tipsTotal,
			Comments:    string(commentsJson),
			Updates:     string(updatesJson),
		}

		renderTemplate(w, "/project/{projectId}", "project", data)
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

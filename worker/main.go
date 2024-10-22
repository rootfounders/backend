package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5"
	"github.com/rootfounders/backend/config"
	rootFounders "github.com/rootfounders/backend/contracts"
	database "github.com/rootfounders/backend/database"
)

var conn *database.Connection

func main() {
	// flags
	configPath := flag.String("config", "./rootfounders.toml", "configuration file")
	flag.Parse()

	config, err := config.Load(*configPath)
	if err != nil {
		log.Fatalln("could not load config file", err)
	}
	if len(config.MainContract) == 0 {
		log.Fatalln("main contract address is missing")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := ethclient.DialContext(ctx, config.Rpc)
	if err != nil {
		log.Fatalln(err)
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if config.Chain > 0 {
		if fmt.Sprint(config.Chain) != chainId.String() {
			log.Fatalf("chain: wanted %d, but RPC returns %s\n", config.Chain, chainId.String())
		}
	}
	log.Println("Watching chain", chainId.String())

	conn = &database.Connection{
		Chain:    uint(chainId.Uint64()),
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.Database,
	}
	if err := conn.Connect(ctx); err != nil {
		log.Println("db connect err:", err)
		return
	}
	defer conn.Conn().Close(ctx)

	log.Println("watching contract at", config.MainContract)

	lastBlock, err := conn.Progress(ctx)
	if err != nil && err != pgx.ErrNoRows {
		log.Println(err)
		return
	}
	opts := &bind.WatchOpts{}
	if lastBlock > 0 {
		opts.Start = &lastBlock
		log.Println("continuing watching from block", lastBlock)
	}

	mainContract, err := rootFounders.NewRootFounders(config.MainContract, client)
	if err != nil {
		log.Println("abi init error:", err)
		return
	}

	projects := make(chan *rootFounders.RootFoundersProjectCreated, 100)
	projectCreatedSub, err := mainContract.WatchProjectCreated(opts, projects, nil, nil)
	if err != nil {
		log.Println("projects sub:", err)
		return
	}

	comments := make(chan *rootFounders.RootFoundersCommented, 100)
	commentedSub, err := mainContract.WatchCommented(opts, comments, nil, nil)
	if err != nil {
		log.Println("comments sub:", err)
		return
	}

	applications := make(chan *rootFounders.RootFoundersApplied, 100)
	appliedSub, err := mainContract.WatchApplied(opts, applications, nil, nil)
	if err != nil {
		log.Println("applications sub:", err)
		return
	}

	teamJoins := make(chan *rootFounders.RootFoundersJoinedTeam, 100)
	teamJoinsSub, err := mainContract.WatchJoinedTeam(opts, teamJoins, nil, nil)
	if err != nil {
		log.Println("team joins:", err)
		return
	}

	teamLeaves := make(chan *rootFounders.RootFoundersLeftTeam, 100)
	teamLeavesSub, err := mainContract.WatchLeftTeam(opts, teamLeaves, nil, nil)
	if err != nil {
		log.Println("team leaves:", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Canceling")
			return

		case err := <-projectCreatedSub.Err():
			log.Fatal("project subscription error", err)
		case projectCreated := <-projects:
			log.Println("new project:", projectCreated.Project)
			if err := conn.CreateProject(ctx, &projectCreated.Project); err != nil {
				logDbError("project", projectCreated.Project, err)
				return
			}
			if err := conn.SaveProgress(ctx, projectCreated.Raw.BlockNumber); err != nil {
				logDbError("progress", projectCreated.Raw.BlockNumber, err)
			}

		case err := <-commentedSub.Err():
			log.Fatalln("comment subscription error", err)
		case comment := <-comments:
			if err := conn.CreateComment(ctx, comment); err != nil {
				logDbError("comment", comment, err)
				return
			}
			if err := conn.SaveProgress(ctx, comment.Raw.BlockNumber); err != nil {
				logDbError("progress", comment.Raw.BlockNumber, err)
			}

		case err := <-appliedSub.Err():
			log.Fatalln("applied subscription error", err)
		case application := <-applications:
			if err := conn.CreateTeamMember(ctx, application); err != nil {
				logDbError("application", application, err)
				return
			}
			if err := conn.SaveProgress(ctx, application.Raw.BlockNumber); err != nil {
				logDbError("progress", application.Raw.BlockNumber, err)
			}

		case err := <-teamJoinsSub.Err():
			log.Fatalln("joined team subscription error", err)
		case teamJoin := <-teamJoins:
			if err := conn.AcceptTeamMember(ctx, teamJoin.ProjectId, &teamJoin.Who); err != nil {
				logDbError("joined team", teamJoin, err)
				return
			}
			if err := conn.SaveProgress(ctx, teamJoin.Raw.BlockNumber); err != nil {
				logDbError("progress", teamJoin.Raw.BlockNumber, err)
			}

		case err := <-teamLeavesSub.Err():
			log.Fatalln("left team subscription error", err)
		case teamLeave := <-teamLeaves:
			if err := conn.MarkTeamMemberAsRemoved(ctx, teamLeave.ProjectId, &teamLeave.Who); err != nil {
				logDbError("left team", teamLeave, err)
				return
			}
			if err := conn.SaveProgress(ctx, teamLeave.Raw.BlockNumber); err != nil {
				logDbError("progress", teamLeave.Raw.BlockNumber, err)
			}
		}
	}
}

func logDbError(recordType string, recordData any, err error) {
	log.Printf("database error for %s: %s\n", recordType, err)
	log.Printf("record data: %+v\n", recordData)
}

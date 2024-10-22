package database

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/microcosm-cc/bluemonday"
	rootFounders "github.com/rootfounders/backend/contracts"
)

var policy = bluemonday.StrictPolicy()

func (c *Connection) Progress(ctx context.Context) (blockNumber uint64, err error) {
	err = c.conn.QueryRow(
		ctx,
		`SELECT lastBlock FROM progress WHERE chain = $1`,
		c.Chain,
	).Scan(&blockNumber)

	return
}

func (c *Connection) SaveProgress(ctx context.Context, blockNumber uint64) (err error) {
	_, err = c.conn.Exec(
		ctx,
		`INSERT INTO progress(chain, lastBlock) VALUES($1, $2) ON CONFLICT (chain) DO UPDATE SET lastBlock = EXCLUDED.lastBlock`,
		c.Chain,
		blockNumber,
	)

	return
}

func (c *Connection) CreateProject(ctx context.Context, project *rootFounders.Project) (err error) {
	_, err = c.conn.Exec(
		ctx,
		`INSERT INTO projects(id, chain, owner, detailsLocationType, detailsLocation, shortName, tipJarAddress) VALUES($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING`,
		project.Id,
		c.Chain,
		project.Owner.Hex(),
		project.DetailsLocationType,
		project.DetailsLocation,
		project.ShortName,
		project.TipJar.Hex(),
	)

	return
}

// This type exists only as a workaround for bigint conversion problems between PSQL and Go
type ProjectRecord struct {
	Id                  string
	Owner               string
	DetailsLocationType uint8
	DetailsLocation     string
	ShortName           string
	TipJar              string
}

func (c *Connection) GetProject(ctx context.Context, projectId *big.Int) (project *ProjectRecord, err error) {
	project = &ProjectRecord{}
	rows, err := c.conn.Query(
		ctx,
		`SELECT id::text, owner, detailsLocationType, detailsLocation, shortName, tipJarAddress as tipJar FROM projects WHERE id = $1`,
		projectId,
	)
	if err != nil {
		return
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[ProjectRecord])
	if err != nil {
		return
	}
	project = &r
	return
}

func (c *Connection) FeaturedProject(ctx context.Context, count uint) (projects []ProjectRecord, err error) {
	rows, err := c.conn.Query(
		ctx,
		`SELECT id::text, owner, detailsLocationType, detailsLocation, shortName, tipJarAddress as tipJar FROM projects ORDER BY dbAdded DESC LIMIT $1`,
		count,
	)
	if err != nil {
		return
	}

	projects, err = pgx.CollectRows(rows, pgx.RowToStructByName[ProjectRecord])
	return
}

func (c *Connection) CreateComment(ctx context.Context, commentEvent *rootFounders.RootFoundersCommented) (err error) {
	safeText := policy.Sanitize(commentEvent.Comment)
	_, err = c.conn.Exec(
		ctx,
		`INSERT INTO comments(tx, projectId, chain, author, content) VALUES($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`,
		commentEvent.Raw.TxHash.Hex(),
		commentEvent.ProjectId,
		c.Chain,
		commentEvent.From.Hex(),
		safeText,
	)

	return
}

type Comment struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (c *Connection) GetComments(ctx context.Context, projectId *big.Int) (comments []Comment, err error) {
	rows, err := c.conn.Query(
		ctx,
		`SELECT author, content FROM comments WHERE projectId = $1`,
		projectId,
	)
	if err != nil {
		return
	}

	return pgx.CollectRows[Comment](rows, pgx.RowToStructByName)
}

func (c *Connection) CreateTeamMember(ctx context.Context, appliedEvent *rootFounders.RootFoundersApplied) (err error) {
	_, err = c.conn.Exec(
		ctx,
		`INSERT INTO team_members(tx, projectId, chain, address, accepted) VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING`,
		appliedEvent.Raw.TxHash.Hex(),
		appliedEvent.ProjectId,
		c.Chain,
		appliedEvent.Who.Hex(),
	)

	return
}

func (c *Connection) AcceptTeamMember(ctx context.Context, projectId *big.Int, address *common.Address) (err error) {
	_, err = c.conn.Exec(
		ctx,
		`UPDATE team_members SET accepted = TRUE WHERE projectId = $2 AND address = $3`,
		projectId,
		address.Hex(),
	)

	return
}

func (c *Connection) MarkTeamMemberAsRemoved(ctx context.Context, projectId *big.Int, address *common.Address) (err error) {
	_, err = c.conn.Exec(
		ctx,
		`UPDATE team_members SET left = TRUE WHERE projectId = $1 AND address = $2`,
		projectId,
		address.Hex(),
	)

	return
}

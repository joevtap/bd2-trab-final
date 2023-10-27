// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	Agent           *agent
	Event           *event
	EventOccurrence *eventOccurrence
	Project         *project
	Space           *space
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Agent = &Q.Agent
	Event = &Q.Event
	EventOccurrence = &Q.EventOccurrence
	Project = &Q.Project
	Space = &Q.Space
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		Agent:           newAgent(db, opts...),
		Event:           newEvent(db, opts...),
		EventOccurrence: newEventOccurrence(db, opts...),
		Project:         newProject(db, opts...),
		Space:           newSpace(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Agent           agent
	Event           event
	EventOccurrence eventOccurrence
	Project         project
	Space           space
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Agent:           q.Agent.clone(db),
		Event:           q.Event.clone(db),
		EventOccurrence: q.EventOccurrence.clone(db),
		Project:         q.Project.clone(db),
		Space:           q.Space.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Agent:           q.Agent.replaceDB(db),
		Event:           q.Event.replaceDB(db),
		EventOccurrence: q.EventOccurrence.replaceDB(db),
		Project:         q.Project.replaceDB(db),
		Space:           q.Space.replaceDB(db),
	}
}

type queryCtx struct {
	Agent           IAgentDo
	Event           IEventDo
	EventOccurrence IEventOccurrenceDo
	Project         IProjectDo
	Space           ISpaceDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Agent:           q.Agent.WithContext(ctx),
		Event:           q.Event.WithContext(ctx),
		EventOccurrence: q.EventOccurrence.WithContext(ctx),
		Project:         q.Project.WithContext(ctx),
		Space:           q.Space.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}

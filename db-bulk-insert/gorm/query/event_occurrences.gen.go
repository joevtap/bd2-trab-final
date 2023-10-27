// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func newEventOccurrence(db *gorm.DB, opts ...gen.DOOption) eventOccurrence {
	_eventOccurrence := eventOccurrence{}

	_eventOccurrence.eventOccurrenceDo.UseDB(db, opts...)
	_eventOccurrence.eventOccurrenceDo.UseModel(&model.EventOccurrence{})

	tableName := _eventOccurrence.eventOccurrenceDo.TableName()
	_eventOccurrence.ALL = field.NewAsterisk(tableName)
	_eventOccurrence.ID = field.NewInt32(tableName, "id")
	_eventOccurrence.Startson = field.NewTime(tableName, "_startson")
	_eventOccurrence.Endson = field.NewTime(tableName, "_endson")
	_eventOccurrence.Startsat = field.NewTime(tableName, "_startsat")
	_eventOccurrence.Endsat = field.NewTime(tableName, "_endsat")
	_eventOccurrence.Frequency = field.NewString(tableName, "frequency")
	_eventOccurrence.Separation = field.NewInt32(tableName, "separation")
	_eventOccurrence.Count_ = field.NewInt32(tableName, "_count")
	_eventOccurrence.Until = field.NewTime(tableName, "_until")
	_eventOccurrence.Timezonename = field.NewString(tableName, "timezonename")
	_eventOccurrence.Event = field.NewInt32(tableName, "event")
	_eventOccurrence.Eventid = field.NewInt32(tableName, "eventid")
	_eventOccurrence.Space = field.NewInt32(tableName, "space")
	_eventOccurrence.Spaceid = field.NewInt32(tableName, "spaceid")
	_eventOccurrence.Rule = field.NewString(tableName, "rule")
	_eventOccurrence.Status = field.NewInt32(tableName, "status")

	_eventOccurrence.fillFieldMap()

	return _eventOccurrence
}

type eventOccurrence struct {
	eventOccurrenceDo

	ALL          field.Asterisk
	ID           field.Int32
	Startson     field.Time
	Endson       field.Time
	Startsat     field.Time
	Endsat       field.Time
	Frequency    field.String
	Separation   field.Int32
	Count_       field.Int32
	Until        field.Time
	Timezonename field.String
	Event        field.Int32
	Eventid      field.Int32
	Space        field.Int32
	Spaceid      field.Int32
	Rule         field.String
	Status       field.Int32

	fieldMap map[string]field.Expr
}

func (e eventOccurrence) Table(newTableName string) *eventOccurrence {
	e.eventOccurrenceDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e eventOccurrence) As(alias string) *eventOccurrence {
	e.eventOccurrenceDo.DO = *(e.eventOccurrenceDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *eventOccurrence) updateTableName(table string) *eventOccurrence {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Startson = field.NewTime(table, "_startson")
	e.Endson = field.NewTime(table, "_endson")
	e.Startsat = field.NewTime(table, "_startsat")
	e.Endsat = field.NewTime(table, "_endsat")
	e.Frequency = field.NewString(table, "frequency")
	e.Separation = field.NewInt32(table, "separation")
	e.Count_ = field.NewInt32(table, "_count")
	e.Until = field.NewTime(table, "_until")
	e.Timezonename = field.NewString(table, "timezonename")
	e.Event = field.NewInt32(table, "event")
	e.Eventid = field.NewInt32(table, "eventid")
	e.Space = field.NewInt32(table, "space")
	e.Spaceid = field.NewInt32(table, "spaceid")
	e.Rule = field.NewString(table, "rule")
	e.Status = field.NewInt32(table, "status")

	e.fillFieldMap()

	return e
}

func (e *eventOccurrence) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *eventOccurrence) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 16)
	e.fieldMap["id"] = e.ID
	e.fieldMap["_startson"] = e.Startson
	e.fieldMap["_endson"] = e.Endson
	e.fieldMap["_startsat"] = e.Startsat
	e.fieldMap["_endsat"] = e.Endsat
	e.fieldMap["frequency"] = e.Frequency
	e.fieldMap["separation"] = e.Separation
	e.fieldMap["_count"] = e.Count_
	e.fieldMap["_until"] = e.Until
	e.fieldMap["timezonename"] = e.Timezonename
	e.fieldMap["event"] = e.Event
	e.fieldMap["eventid"] = e.Eventid
	e.fieldMap["space"] = e.Space
	e.fieldMap["spaceid"] = e.Spaceid
	e.fieldMap["rule"] = e.Rule
	e.fieldMap["status"] = e.Status
}

func (e eventOccurrence) clone(db *gorm.DB) eventOccurrence {
	e.eventOccurrenceDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e eventOccurrence) replaceDB(db *gorm.DB) eventOccurrence {
	e.eventOccurrenceDo.ReplaceDB(db)
	return e
}

type eventOccurrenceDo struct{ gen.DO }

type IEventOccurrenceDo interface {
	gen.SubQuery
	Debug() IEventOccurrenceDo
	WithContext(ctx context.Context) IEventOccurrenceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEventOccurrenceDo
	WriteDB() IEventOccurrenceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEventOccurrenceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEventOccurrenceDo
	Not(conds ...gen.Condition) IEventOccurrenceDo
	Or(conds ...gen.Condition) IEventOccurrenceDo
	Select(conds ...field.Expr) IEventOccurrenceDo
	Where(conds ...gen.Condition) IEventOccurrenceDo
	Order(conds ...field.Expr) IEventOccurrenceDo
	Distinct(cols ...field.Expr) IEventOccurrenceDo
	Omit(cols ...field.Expr) IEventOccurrenceDo
	Join(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo
	Group(cols ...field.Expr) IEventOccurrenceDo
	Having(conds ...gen.Condition) IEventOccurrenceDo
	Limit(limit int) IEventOccurrenceDo
	Offset(offset int) IEventOccurrenceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEventOccurrenceDo
	Unscoped() IEventOccurrenceDo
	Create(values ...*model.EventOccurrence) error
	CreateInBatches(values []*model.EventOccurrence, batchSize int) error
	Save(values ...*model.EventOccurrence) error
	First() (*model.EventOccurrence, error)
	Take() (*model.EventOccurrence, error)
	Last() (*model.EventOccurrence, error)
	Find() ([]*model.EventOccurrence, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EventOccurrence, err error)
	FindInBatches(result *[]*model.EventOccurrence, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EventOccurrence) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEventOccurrenceDo
	Assign(attrs ...field.AssignExpr) IEventOccurrenceDo
	Joins(fields ...field.RelationField) IEventOccurrenceDo
	Preload(fields ...field.RelationField) IEventOccurrenceDo
	FirstOrInit() (*model.EventOccurrence, error)
	FirstOrCreate() (*model.EventOccurrence, error)
	FindByPage(offset int, limit int) (result []*model.EventOccurrence, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEventOccurrenceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e eventOccurrenceDo) Debug() IEventOccurrenceDo {
	return e.withDO(e.DO.Debug())
}

func (e eventOccurrenceDo) WithContext(ctx context.Context) IEventOccurrenceDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e eventOccurrenceDo) ReadDB() IEventOccurrenceDo {
	return e.Clauses(dbresolver.Read)
}

func (e eventOccurrenceDo) WriteDB() IEventOccurrenceDo {
	return e.Clauses(dbresolver.Write)
}

func (e eventOccurrenceDo) Session(config *gorm.Session) IEventOccurrenceDo {
	return e.withDO(e.DO.Session(config))
}

func (e eventOccurrenceDo) Clauses(conds ...clause.Expression) IEventOccurrenceDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e eventOccurrenceDo) Returning(value interface{}, columns ...string) IEventOccurrenceDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e eventOccurrenceDo) Not(conds ...gen.Condition) IEventOccurrenceDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e eventOccurrenceDo) Or(conds ...gen.Condition) IEventOccurrenceDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e eventOccurrenceDo) Select(conds ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e eventOccurrenceDo) Where(conds ...gen.Condition) IEventOccurrenceDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e eventOccurrenceDo) Order(conds ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e eventOccurrenceDo) Distinct(cols ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e eventOccurrenceDo) Omit(cols ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e eventOccurrenceDo) Join(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e eventOccurrenceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e eventOccurrenceDo) RightJoin(table schema.Tabler, on ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e eventOccurrenceDo) Group(cols ...field.Expr) IEventOccurrenceDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e eventOccurrenceDo) Having(conds ...gen.Condition) IEventOccurrenceDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e eventOccurrenceDo) Limit(limit int) IEventOccurrenceDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e eventOccurrenceDo) Offset(offset int) IEventOccurrenceDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e eventOccurrenceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEventOccurrenceDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e eventOccurrenceDo) Unscoped() IEventOccurrenceDo {
	return e.withDO(e.DO.Unscoped())
}

func (e eventOccurrenceDo) Create(values ...*model.EventOccurrence) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e eventOccurrenceDo) CreateInBatches(values []*model.EventOccurrence, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e eventOccurrenceDo) Save(values ...*model.EventOccurrence) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e eventOccurrenceDo) First() (*model.EventOccurrence, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EventOccurrence), nil
	}
}

func (e eventOccurrenceDo) Take() (*model.EventOccurrence, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EventOccurrence), nil
	}
}

func (e eventOccurrenceDo) Last() (*model.EventOccurrence, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EventOccurrence), nil
	}
}

func (e eventOccurrenceDo) Find() ([]*model.EventOccurrence, error) {
	result, err := e.DO.Find()
	return result.([]*model.EventOccurrence), err
}

func (e eventOccurrenceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EventOccurrence, err error) {
	buf := make([]*model.EventOccurrence, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e eventOccurrenceDo) FindInBatches(result *[]*model.EventOccurrence, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e eventOccurrenceDo) Attrs(attrs ...field.AssignExpr) IEventOccurrenceDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e eventOccurrenceDo) Assign(attrs ...field.AssignExpr) IEventOccurrenceDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e eventOccurrenceDo) Joins(fields ...field.RelationField) IEventOccurrenceDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e eventOccurrenceDo) Preload(fields ...field.RelationField) IEventOccurrenceDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e eventOccurrenceDo) FirstOrInit() (*model.EventOccurrence, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EventOccurrence), nil
	}
}

func (e eventOccurrenceDo) FirstOrCreate() (*model.EventOccurrence, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EventOccurrence), nil
	}
}

func (e eventOccurrenceDo) FindByPage(offset int, limit int) (result []*model.EventOccurrence, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e eventOccurrenceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e eventOccurrenceDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e eventOccurrenceDo) Delete(models ...*model.EventOccurrence) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *eventOccurrenceDo) withDO(do gen.Dao) *eventOccurrenceDo {
	e.DO = *do.(*gen.DO)
	return e
}

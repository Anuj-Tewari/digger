// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/ee/drift/model"
)

func newDiggerCiJob(db *gorm.DB, opts ...gen.DOOption) diggerCiJob {
	_diggerCiJob := diggerCiJob{}

	_diggerCiJob.diggerCiJobDo.UseDB(db, opts...)
	_diggerCiJob.diggerCiJobDo.UseModel(&model.DiggerCiJob{})

	tableName := _diggerCiJob.diggerCiJobDo.TableName()
	_diggerCiJob.ALL = field.NewAsterisk(tableName)
	_diggerCiJob.ID = field.NewString(tableName, "id")
	_diggerCiJob.Spectype = field.NewString(tableName, "spectype")
	_diggerCiJob.Commentid = field.NewString(tableName, "commentid")
	_diggerCiJob.Runname = field.NewString(tableName, "runname")
	_diggerCiJob.Jobspec = field.NewField(tableName, "jobspec")
	_diggerCiJob.Reporterspec = field.NewField(tableName, "reporterspec")
	_diggerCiJob.Commentupdaterspec = field.NewField(tableName, "commentupdaterspec")
	_diggerCiJob.Lockspec = field.NewField(tableName, "lockspec")
	_diggerCiJob.Backendspec = field.NewField(tableName, "backendspec")
	_diggerCiJob.Vcsspec = field.NewField(tableName, "vcsspec")
	_diggerCiJob.Policyspec = field.NewField(tableName, "policyspec")
	_diggerCiJob.Variablesspec = field.NewField(tableName, "variablesspec")
	_diggerCiJob.CreatedAt = field.NewTime(tableName, "created_at")
	_diggerCiJob.UpdatedAt = field.NewTime(tableName, "updated_at")
	_diggerCiJob.DeletedAt = field.NewField(tableName, "deleted_at")
	_diggerCiJob.WorkflowFile = field.NewString(tableName, "workflow_file")
	_diggerCiJob.WorkflowURL = field.NewString(tableName, "workflow_url")
	_diggerCiJob.Status = field.NewString(tableName, "status")
	_diggerCiJob.ResourcesCreated = field.NewInt32(tableName, "resources_created")
	_diggerCiJob.ResourcesUpdated = field.NewInt32(tableName, "resources_updated")
	_diggerCiJob.ResourcesDeleted = field.NewInt32(tableName, "resources_deleted")
	_diggerCiJob.ProjectID = field.NewString(tableName, "project_id")
	_diggerCiJob.DiggerJobID = field.NewString(tableName, "digger_job_id")
	_diggerCiJob.TerraformOutput = field.NewString(tableName, "terraform_output")

	_diggerCiJob.fillFieldMap()

	return _diggerCiJob
}

type diggerCiJob struct {
	diggerCiJobDo

	ALL                field.Asterisk
	ID                 field.String
	Spectype           field.String
	Commentid          field.String
	Runname            field.String
	Jobspec            field.Field
	Reporterspec       field.Field
	Commentupdaterspec field.Field
	Lockspec           field.Field
	Backendspec        field.Field
	Vcsspec            field.Field
	Policyspec         field.Field
	Variablesspec      field.Field
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field
	WorkflowFile       field.String
	WorkflowURL        field.String
	Status             field.String
	ResourcesCreated   field.Int32
	ResourcesUpdated   field.Int32
	ResourcesDeleted   field.Int32
	ProjectID          field.String
	DiggerJobID        field.String
	TerraformOutput    field.String

	fieldMap map[string]field.Expr
}

func (d diggerCiJob) Table(newTableName string) *diggerCiJob {
	d.diggerCiJobDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d diggerCiJob) As(alias string) *diggerCiJob {
	d.diggerCiJobDo.DO = *(d.diggerCiJobDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *diggerCiJob) updateTableName(table string) *diggerCiJob {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Spectype = field.NewString(table, "spectype")
	d.Commentid = field.NewString(table, "commentid")
	d.Runname = field.NewString(table, "runname")
	d.Jobspec = field.NewField(table, "jobspec")
	d.Reporterspec = field.NewField(table, "reporterspec")
	d.Commentupdaterspec = field.NewField(table, "commentupdaterspec")
	d.Lockspec = field.NewField(table, "lockspec")
	d.Backendspec = field.NewField(table, "backendspec")
	d.Vcsspec = field.NewField(table, "vcsspec")
	d.Policyspec = field.NewField(table, "policyspec")
	d.Variablesspec = field.NewField(table, "variablesspec")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.WorkflowFile = field.NewString(table, "workflow_file")
	d.WorkflowURL = field.NewString(table, "workflow_url")
	d.Status = field.NewString(table, "status")
	d.ResourcesCreated = field.NewInt32(table, "resources_created")
	d.ResourcesUpdated = field.NewInt32(table, "resources_updated")
	d.ResourcesDeleted = field.NewInt32(table, "resources_deleted")
	d.ProjectID = field.NewString(table, "project_id")
	d.DiggerJobID = field.NewString(table, "digger_job_id")
	d.TerraformOutput = field.NewString(table, "terraform_output")

	d.fillFieldMap()

	return d
}

func (d *diggerCiJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *diggerCiJob) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 24)
	d.fieldMap["id"] = d.ID
	d.fieldMap["spectype"] = d.Spectype
	d.fieldMap["commentid"] = d.Commentid
	d.fieldMap["runname"] = d.Runname
	d.fieldMap["jobspec"] = d.Jobspec
	d.fieldMap["reporterspec"] = d.Reporterspec
	d.fieldMap["commentupdaterspec"] = d.Commentupdaterspec
	d.fieldMap["lockspec"] = d.Lockspec
	d.fieldMap["backendspec"] = d.Backendspec
	d.fieldMap["vcsspec"] = d.Vcsspec
	d.fieldMap["policyspec"] = d.Policyspec
	d.fieldMap["variablesspec"] = d.Variablesspec
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["workflow_file"] = d.WorkflowFile
	d.fieldMap["workflow_url"] = d.WorkflowURL
	d.fieldMap["status"] = d.Status
	d.fieldMap["resources_created"] = d.ResourcesCreated
	d.fieldMap["resources_updated"] = d.ResourcesUpdated
	d.fieldMap["resources_deleted"] = d.ResourcesDeleted
	d.fieldMap["project_id"] = d.ProjectID
	d.fieldMap["digger_job_id"] = d.DiggerJobID
	d.fieldMap["terraform_output"] = d.TerraformOutput
}

func (d diggerCiJob) clone(db *gorm.DB) diggerCiJob {
	d.diggerCiJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d diggerCiJob) replaceDB(db *gorm.DB) diggerCiJob {
	d.diggerCiJobDo.ReplaceDB(db)
	return d
}

type diggerCiJobDo struct{ gen.DO }

type IDiggerCiJobDo interface {
	gen.SubQuery
	Debug() IDiggerCiJobDo
	WithContext(ctx context.Context) IDiggerCiJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDiggerCiJobDo
	WriteDB() IDiggerCiJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDiggerCiJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDiggerCiJobDo
	Not(conds ...gen.Condition) IDiggerCiJobDo
	Or(conds ...gen.Condition) IDiggerCiJobDo
	Select(conds ...field.Expr) IDiggerCiJobDo
	Where(conds ...gen.Condition) IDiggerCiJobDo
	Order(conds ...field.Expr) IDiggerCiJobDo
	Distinct(cols ...field.Expr) IDiggerCiJobDo
	Omit(cols ...field.Expr) IDiggerCiJobDo
	Join(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo
	Group(cols ...field.Expr) IDiggerCiJobDo
	Having(conds ...gen.Condition) IDiggerCiJobDo
	Limit(limit int) IDiggerCiJobDo
	Offset(offset int) IDiggerCiJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDiggerCiJobDo
	Unscoped() IDiggerCiJobDo
	Create(values ...*model.DiggerCiJob) error
	CreateInBatches(values []*model.DiggerCiJob, batchSize int) error
	Save(values ...*model.DiggerCiJob) error
	First() (*model.DiggerCiJob, error)
	Take() (*model.DiggerCiJob, error)
	Last() (*model.DiggerCiJob, error)
	Find() ([]*model.DiggerCiJob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DiggerCiJob, err error)
	FindInBatches(result *[]*model.DiggerCiJob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DiggerCiJob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDiggerCiJobDo
	Assign(attrs ...field.AssignExpr) IDiggerCiJobDo
	Joins(fields ...field.RelationField) IDiggerCiJobDo
	Preload(fields ...field.RelationField) IDiggerCiJobDo
	FirstOrInit() (*model.DiggerCiJob, error)
	FirstOrCreate() (*model.DiggerCiJob, error)
	FindByPage(offset int, limit int) (result []*model.DiggerCiJob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDiggerCiJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d diggerCiJobDo) Debug() IDiggerCiJobDo {
	return d.withDO(d.DO.Debug())
}

func (d diggerCiJobDo) WithContext(ctx context.Context) IDiggerCiJobDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d diggerCiJobDo) ReadDB() IDiggerCiJobDo {
	return d.Clauses(dbresolver.Read)
}

func (d diggerCiJobDo) WriteDB() IDiggerCiJobDo {
	return d.Clauses(dbresolver.Write)
}

func (d diggerCiJobDo) Session(config *gorm.Session) IDiggerCiJobDo {
	return d.withDO(d.DO.Session(config))
}

func (d diggerCiJobDo) Clauses(conds ...clause.Expression) IDiggerCiJobDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d diggerCiJobDo) Returning(value interface{}, columns ...string) IDiggerCiJobDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d diggerCiJobDo) Not(conds ...gen.Condition) IDiggerCiJobDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d diggerCiJobDo) Or(conds ...gen.Condition) IDiggerCiJobDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d diggerCiJobDo) Select(conds ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d diggerCiJobDo) Where(conds ...gen.Condition) IDiggerCiJobDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d diggerCiJobDo) Order(conds ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d diggerCiJobDo) Distinct(cols ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d diggerCiJobDo) Omit(cols ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d diggerCiJobDo) Join(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d diggerCiJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d diggerCiJobDo) RightJoin(table schema.Tabler, on ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d diggerCiJobDo) Group(cols ...field.Expr) IDiggerCiJobDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d diggerCiJobDo) Having(conds ...gen.Condition) IDiggerCiJobDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d diggerCiJobDo) Limit(limit int) IDiggerCiJobDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d diggerCiJobDo) Offset(offset int) IDiggerCiJobDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d diggerCiJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDiggerCiJobDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d diggerCiJobDo) Unscoped() IDiggerCiJobDo {
	return d.withDO(d.DO.Unscoped())
}

func (d diggerCiJobDo) Create(values ...*model.DiggerCiJob) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d diggerCiJobDo) CreateInBatches(values []*model.DiggerCiJob, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d diggerCiJobDo) Save(values ...*model.DiggerCiJob) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d diggerCiJobDo) First() (*model.DiggerCiJob, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerCiJob), nil
	}
}

func (d diggerCiJobDo) Take() (*model.DiggerCiJob, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerCiJob), nil
	}
}

func (d diggerCiJobDo) Last() (*model.DiggerCiJob, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerCiJob), nil
	}
}

func (d diggerCiJobDo) Find() ([]*model.DiggerCiJob, error) {
	result, err := d.DO.Find()
	return result.([]*model.DiggerCiJob), err
}

func (d diggerCiJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DiggerCiJob, err error) {
	buf := make([]*model.DiggerCiJob, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d diggerCiJobDo) FindInBatches(result *[]*model.DiggerCiJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d diggerCiJobDo) Attrs(attrs ...field.AssignExpr) IDiggerCiJobDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d diggerCiJobDo) Assign(attrs ...field.AssignExpr) IDiggerCiJobDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d diggerCiJobDo) Joins(fields ...field.RelationField) IDiggerCiJobDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d diggerCiJobDo) Preload(fields ...field.RelationField) IDiggerCiJobDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d diggerCiJobDo) FirstOrInit() (*model.DiggerCiJob, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerCiJob), nil
	}
}

func (d diggerCiJobDo) FirstOrCreate() (*model.DiggerCiJob, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerCiJob), nil
	}
}

func (d diggerCiJobDo) FindByPage(offset int, limit int) (result []*model.DiggerCiJob, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d diggerCiJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d diggerCiJobDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d diggerCiJobDo) Delete(models ...*model.DiggerCiJob) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *diggerCiJobDo) withDO(do gen.Dao) *diggerCiJobDo {
	d.DO = *do.(*gen.DO)
	return d
}

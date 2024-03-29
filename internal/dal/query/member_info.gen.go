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

	"github.com/Godyu97/vegeGoWebTemplate/internal/dal/model"
)

func newMember(db *gorm.DB, opts ...gen.DOOption) member {
	_member := member{}

	_member.memberDo.UseDB(db, opts...)
	_member.memberDo.UseModel(&model.Member{})

	tableName := _member.memberDo.TableName()
	_member.ALL = field.NewAsterisk(tableName)
	_member.ID = field.NewInt64(tableName, "id")
	_member.Code = field.NewString(tableName, "code")
	_member.FollowInviter = field.NewString(tableName, "follow_inviter")
	_member.PaymentInviter = field.NewString(tableName, "payment_inviter")
	_member.Openid = field.NewString(tableName, "openid")
	_member.Unionid = field.NewString(tableName, "unionid")
	_member.Nickname = field.NewString(tableName, "nickname")
	_member.Avatar = field.NewString(tableName, "avatar")
	_member.Account = field.NewString(tableName, "account")
	_member.Password = field.NewString(tableName, "password")
	_member.Email = field.NewString(tableName, "email")
	_member.Mobile = field.NewString(tableName, "mobile")
	_member.Username = field.NewString(tableName, "username")
	_member.Company = field.NewString(tableName, "company")
	_member.Province = field.NewInt64(tableName, "province")
	_member.City = field.NewInt64(tableName, "city")
	_member.Address = field.NewString(tableName, "address")
	_member.Level = field.NewInt64(tableName, "level")
	_member.Expires = field.NewTime(tableName, "expires")
	_member.Readable = field.NewInt64(tableName, "readable")
	_member.Status = field.NewInt64(tableName, "status")
	_member.UpdateTime = field.NewTime(tableName, "update_time")
	_member.CreateTime = field.NewTime(tableName, "create_time")
	_member.RegTyp = field.NewString(tableName, "reg_typ")

	_member.fillFieldMap()

	return _member
}

type member struct {
	memberDo memberDo

	ALL            field.Asterisk
	ID             field.Int64
	Code           field.String // 用户编码
	FollowInviter  field.String // 关注邀请人
	PaymentInviter field.String // 支付邀请人
	Openid         field.String // openid
	Unionid        field.String // unionid
	Nickname       field.String // 昵称
	Avatar         field.String // 头像
	Account        field.String // 账号
	Password       field.String // 密码
	Email          field.String // 邮箱
	Mobile         field.String // 手机号
	Username       field.String // 用户名
	Company        field.String // 公司名
	Province       field.Int64  // 省份
	City           field.Int64  // 市
	Address        field.String // 地址
	/*
		会员等级
		0: 非会员
		10: VIP会员
		11: 高级会员
		12: 钻石会员
		17: 试用会员
		18: 测试会员
		19: 内部会员
		21: 原1 VIP会员
		22: 原2 高级会员
		23: 原3 钻石会员
		24: 原4 试用会员
		25: 原5 内部会员
		26: 原6 其他会员
	*/
	Level    field.Int64
	Expires  field.Time  // 会员到期时间
	Readable field.Int64 // 剩余数量
	/*
		会员状态
		0: 默认状态
		-1: 公众号取关
	*/
	Status     field.Int64
	UpdateTime field.Time   // 创建时间
	CreateTime field.Time   // 更新时间
	RegTyp     field.String // 注册渠道

	fieldMap map[string]field.Expr
}

func (m member) Table(newTableName string) *member {
	m.memberDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m member) As(alias string) *member {
	m.memberDo.DO = *(m.memberDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *member) updateTableName(table string) *member {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.Code = field.NewString(table, "code")
	m.FollowInviter = field.NewString(table, "follow_inviter")
	m.PaymentInviter = field.NewString(table, "payment_inviter")
	m.Openid = field.NewString(table, "openid")
	m.Unionid = field.NewString(table, "unionid")
	m.Nickname = field.NewString(table, "nickname")
	m.Avatar = field.NewString(table, "avatar")
	m.Account = field.NewString(table, "account")
	m.Password = field.NewString(table, "password")
	m.Email = field.NewString(table, "email")
	m.Mobile = field.NewString(table, "mobile")
	m.Username = field.NewString(table, "username")
	m.Company = field.NewString(table, "company")
	m.Province = field.NewInt64(table, "province")
	m.City = field.NewInt64(table, "city")
	m.Address = field.NewString(table, "address")
	m.Level = field.NewInt64(table, "level")
	m.Expires = field.NewTime(table, "expires")
	m.Readable = field.NewInt64(table, "readable")
	m.Status = field.NewInt64(table, "status")
	m.UpdateTime = field.NewTime(table, "update_time")
	m.CreateTime = field.NewTime(table, "create_time")
	m.RegTyp = field.NewString(table, "reg_typ")

	m.fillFieldMap()

	return m
}

func (m *member) WithContext(ctx context.Context) *memberDo { return m.memberDo.WithContext(ctx) }

func (m member) TableName() string { return m.memberDo.TableName() }

func (m member) Alias() string { return m.memberDo.Alias() }

func (m member) Columns(cols ...field.Expr) gen.Columns { return m.memberDo.Columns(cols...) }

func (m *member) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *member) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 24)
	m.fieldMap["id"] = m.ID
	m.fieldMap["code"] = m.Code
	m.fieldMap["follow_inviter"] = m.FollowInviter
	m.fieldMap["payment_inviter"] = m.PaymentInviter
	m.fieldMap["openid"] = m.Openid
	m.fieldMap["unionid"] = m.Unionid
	m.fieldMap["nickname"] = m.Nickname
	m.fieldMap["avatar"] = m.Avatar
	m.fieldMap["account"] = m.Account
	m.fieldMap["password"] = m.Password
	m.fieldMap["email"] = m.Email
	m.fieldMap["mobile"] = m.Mobile
	m.fieldMap["username"] = m.Username
	m.fieldMap["company"] = m.Company
	m.fieldMap["province"] = m.Province
	m.fieldMap["city"] = m.City
	m.fieldMap["address"] = m.Address
	m.fieldMap["level"] = m.Level
	m.fieldMap["expires"] = m.Expires
	m.fieldMap["readable"] = m.Readable
	m.fieldMap["status"] = m.Status
	m.fieldMap["update_time"] = m.UpdateTime
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["reg_typ"] = m.RegTyp
}

func (m member) clone(db *gorm.DB) member {
	m.memberDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m member) replaceDB(db *gorm.DB) member {
	m.memberDo.ReplaceDB(db)
	return m
}

type memberDo struct{ gen.DO }

func (m memberDo) Debug() *memberDo {
	return m.withDO(m.DO.Debug())
}

func (m memberDo) WithContext(ctx context.Context) *memberDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberDo) ReadDB() *memberDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberDo) WriteDB() *memberDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberDo) Session(config *gorm.Session) *memberDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberDo) Clauses(conds ...clause.Expression) *memberDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberDo) Returning(value interface{}, columns ...string) *memberDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberDo) Not(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberDo) Or(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberDo) Select(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberDo) Where(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberDo) Order(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberDo) Distinct(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberDo) Omit(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberDo) Join(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberDo) LeftJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberDo) RightJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberDo) Group(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberDo) Having(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberDo) Limit(limit int) *memberDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberDo) Offset(offset int) *memberDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *memberDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberDo) Unscoped() *memberDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberDo) Create(values ...*model.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberDo) CreateInBatches(values []*model.Member, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberDo) Save(values ...*model.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberDo) First() (*model.Member, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Take() (*model.Member, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Last() (*model.Member, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Find() ([]*model.Member, error) {
	result, err := m.DO.Find()
	return result.([]*model.Member), err
}

func (m memberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Member, err error) {
	buf := make([]*model.Member, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberDo) FindInBatches(result *[]*model.Member, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberDo) Attrs(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberDo) Assign(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberDo) Joins(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberDo) Preload(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberDo) FirstOrInit() (*model.Member, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) FirstOrCreate() (*model.Member, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) FindByPage(offset int, limit int) (result []*model.Member, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberDo) Delete(models ...*model.Member) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberDo) withDO(do gen.Dao) *memberDo {
	m.DO = *do.(*gen.DO)
	return m
}

// Code generated by gola 0.0.2; DO NOT EDIT.

package blogs

import (
	"database/sql"
	"strings"

	"github.com/olachat/gola/coredb"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

// Blog represents blogs table
type Blog struct {
	//  int
	Id
	// User Id int
	UserId
	// Slug varchar(255)
	Slug
	// Title varchar(255)
	Title
	// Category Id int
	CategoryId
	// Is pinned to top tinyint
	IsPinned
	// Is VIP reader only tinyint
	IsVip
	// Country of the blog user varchar(255)
	Country
	// Created Timestamp int unsigned
	CreatedAt
	// Updated Timestamp int unsigned
	UpdatedAt
}

type BlogTable struct{}

func (*BlogTable) GetTableName() string {
	return "blogs"
}

var table *BlogTable

type WithPK interface {
	GetId() int
}

// FetchBlogByPKs returns a row from blogs table with given primary key value
func FetchBlogByPK(val int) *Blog {
	return coredb.FetchByPK[Blog](_db, []string{"id"}, val)
}

// FetchByPKs returns a row with selected fields from blogs table with given primary key value
func FetchByPK[T any](val int) *T {
	return coredb.FetchByPK[T](_db, []string{"id"}, val)
}

// FetchBlogByPKs returns rows with from blogs table with given primary key values
func FetchBlogByPKs(vals ...int) []*Blog {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Blog](pks, "id", _db)
}

// FetchByPKs returns rows with selected fields from blogs table with given primary key values
func FetchByPKs[T any](vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](pks, "id", _db)
}

// FindOneBlog returns a row from blogs table with arbitary where query
// whereSQL must start with "where ..."
func FindOneBlog(whereSQL string, params ...any) *Blog {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Blog](w, _db)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM blogs "+whereSQL, _db, params...)
}

// FindOne returns a row with selected fields from blogs table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, _db)
}

// FindBlog returns rows from blogs table with arbitary where query
// whereSQL must start with "where ..."
func FindBlog(whereSQL string, params ...any) ([]*Blog, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Blog](w, _db)
}

// Find returns rows with selected fields from blogs table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, _db)
}

// Column types

// Id field
//
type Id struct {
	_updated bool
	val      int
}

func (c *Id) GetId() int {
	return c.val
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) IsUpdated() bool {
	return c._updated
}

func (c *Id) resetUpdated() {
	c._updated = false
}

func (c *Id) IsPrimaryKey() bool {
	return true
}

func (c *Id) GetValPointer() any {
	return &c.val
}

func (c *Id) GetTableType() coredb.TableType {
	return table
}

// UserId field
// User Id
type UserId struct {
	_updated bool
	val      int
}

func (c *UserId) GetUserId() int {
	return c.val
}

func (c *UserId) SetUserId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) IsUpdated() bool {
	return c._updated
}

func (c *UserId) resetUpdated() {
	c._updated = false
}

func (c *UserId) IsPrimaryKey() bool {
	return false
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) GetTableType() coredb.TableType {
	return table
}

// Slug field
// Slug
type Slug struct {
	_updated bool
	val      string
}

func (c *Slug) GetSlug() string {
	return c.val
}

func (c *Slug) SetSlug(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Slug) GetColumnName() string {
	return "slug"
}

func (c *Slug) IsUpdated() bool {
	return c._updated
}

func (c *Slug) resetUpdated() {
	c._updated = false
}

func (c *Slug) IsPrimaryKey() bool {
	return false
}

func (c *Slug) GetValPointer() any {
	return &c.val
}

func (c *Slug) GetTableType() coredb.TableType {
	return table
}

// Title field
// Title
type Title struct {
	_updated bool
	val      string
}

func (c *Title) GetTitle() string {
	return c.val
}

func (c *Title) SetTitle(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Title) GetColumnName() string {
	return "title"
}

func (c *Title) IsUpdated() bool {
	return c._updated
}

func (c *Title) resetUpdated() {
	c._updated = false
}

func (c *Title) IsPrimaryKey() bool {
	return false
}

func (c *Title) GetValPointer() any {
	return &c.val
}

func (c *Title) GetTableType() coredb.TableType {
	return table
}

// CategoryId field
// Category Id
type CategoryId struct {
	_updated bool
	val      int
}

func (c *CategoryId) GetCategoryId() int {
	return c.val
}

func (c *CategoryId) SetCategoryId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CategoryId) GetColumnName() string {
	return "category_id"
}

func (c *CategoryId) IsUpdated() bool {
	return c._updated
}

func (c *CategoryId) resetUpdated() {
	c._updated = false
}

func (c *CategoryId) IsPrimaryKey() bool {
	return false
}

func (c *CategoryId) GetValPointer() any {
	return &c.val
}

func (c *CategoryId) GetTableType() coredb.TableType {
	return table
}

// IsPinned field
// Is pinned to top
type IsPinned struct {
	_updated bool
	val      int8
}

func (c *IsPinned) GetIsPinned() int8 {
	return c.val
}

func (c *IsPinned) SetIsPinned(val int8) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsPinned) GetColumnName() string {
	return "is_pinned"
}

func (c *IsPinned) IsUpdated() bool {
	return c._updated
}

func (c *IsPinned) resetUpdated() {
	c._updated = false
}

func (c *IsPinned) IsPrimaryKey() bool {
	return false
}

func (c *IsPinned) GetValPointer() any {
	return &c.val
}

func (c *IsPinned) GetTableType() coredb.TableType {
	return table
}

// IsVip field
// Is VIP reader only
type IsVip struct {
	_updated bool
	val      int8
}

func (c *IsVip) GetIsVip() int8 {
	return c.val
}

func (c *IsVip) SetIsVip(val int8) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsVip) GetColumnName() string {
	return "is_vip"
}

func (c *IsVip) IsUpdated() bool {
	return c._updated
}

func (c *IsVip) resetUpdated() {
	c._updated = false
}

func (c *IsVip) IsPrimaryKey() bool {
	return false
}

func (c *IsVip) GetValPointer() any {
	return &c.val
}

func (c *IsVip) GetTableType() coredb.TableType {
	return table
}

// Country field
// Country of the blog user
type Country struct {
	_updated bool
	val      string
}

func (c *Country) GetCountry() string {
	return c.val
}

func (c *Country) SetCountry(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Country) GetColumnName() string {
	return "country"
}

func (c *Country) IsUpdated() bool {
	return c._updated
}

func (c *Country) resetUpdated() {
	c._updated = false
}

func (c *Country) IsPrimaryKey() bool {
	return false
}

func (c *Country) GetValPointer() any {
	return &c.val
}

func (c *Country) GetTableType() coredb.TableType {
	return table
}

// CreatedAt field
// Created Timestamp
type CreatedAt struct {
	_updated bool
	val      uint
}

func (c *CreatedAt) GetCreatedAt() uint {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CreatedAt) GetColumnName() string {
	return "created_at"
}

func (c *CreatedAt) IsUpdated() bool {
	return c._updated
}

func (c *CreatedAt) resetUpdated() {
	c._updated = false
}

func (c *CreatedAt) IsPrimaryKey() bool {
	return false
}

func (c *CreatedAt) GetValPointer() any {
	return &c.val
}

func (c *CreatedAt) GetTableType() coredb.TableType {
	return table
}

// UpdatedAt field
// Updated Timestamp
type UpdatedAt struct {
	_updated bool
	val      uint
}

func (c *UpdatedAt) GetUpdatedAt() uint {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdatedAt) GetColumnName() string {
	return "updated_at"
}

func (c *UpdatedAt) IsUpdated() bool {
	return c._updated
}

func (c *UpdatedAt) resetUpdated() {
	c._updated = false
}

func (c *UpdatedAt) IsPrimaryKey() bool {
	return false
}

func (c *UpdatedAt) GetValPointer() any {
	return &c.val
}

func (c *UpdatedAt) GetTableType() coredb.TableType {
	return table
}

func NewBlog() *Blog {
	return &Blog{
		Id{},
		UserId{val: int(0)},
		Slug{val: ""},
		Title{val: ""},
		CategoryId{val: int(0)},
		IsPinned{val: int8(0)},
		IsVip{val: int8(0)},
		Country{val: ""},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
	}
}

func NewBlogWithPK(val int) *Blog {
	c := &Blog{
		Id{},
		UserId{val: int(0)},
		Slug{val: ""},
		Title{val: ""},
		CategoryId{val: int(0)},
		IsPinned{val: int8(0)},
		IsVip{val: int8(0)},
		Country{val: ""},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
	}
	c.Id.val = val
	return c
}

func (c *Blog) Insert() error {
	sql := `INSERT INTO blogs (user_id, slug, title, category_id, is_pinned, is_vip, country, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := coredb.Exec(sql, _db, c.GetUserId(), c.GetSlug(), c.GetTitle(), c.GetCategoryId(), c.GetIsPinned(), c.GetIsVip(), c.GetCountry(), c.GetCreatedAt(), c.GetUpdatedAt())

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.Id.val = int(id)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return coredb.ErrAvoidInsert
	}

	c.resetUpdated()
	return nil
}

func (c *Blog) resetUpdated() {
	c.Id.resetUpdated()
	c.UserId.resetUpdated()
	c.Slug.resetUpdated()
	c.Title.resetUpdated()
	c.CategoryId.resetUpdated()
	c.IsPinned.resetUpdated()
	c.IsVip.resetUpdated()
	c.Country.resetUpdated()
	c.CreatedAt.resetUpdated()
	c.UpdatedAt.resetUpdated()
}

func (obj *Blog) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Id.IsUpdated() {
		return false, coredb.ErrPKChanged
	}
	if obj.UserId.IsUpdated() {
		updatedFields = append(updatedFields, "user_id = ?")
		params = append(params, obj.GetUserId())
	}
	if obj.Slug.IsUpdated() {
		updatedFields = append(updatedFields, "slug = ?")
		params = append(params, obj.GetSlug())
	}
	if obj.Title.IsUpdated() {
		updatedFields = append(updatedFields, "title = ?")
		params = append(params, obj.GetTitle())
	}
	if obj.CategoryId.IsUpdated() {
		updatedFields = append(updatedFields, "category_id = ?")
		params = append(params, obj.GetCategoryId())
	}
	if obj.IsPinned.IsUpdated() {
		updatedFields = append(updatedFields, "is_pinned = ?")
		params = append(params, obj.GetIsPinned())
	}
	if obj.IsVip.IsUpdated() {
		updatedFields = append(updatedFields, "is_vip = ?")
		params = append(params, obj.GetIsVip())
	}
	if obj.Country.IsUpdated() {
		updatedFields = append(updatedFields, "country = ?")
		params = append(params, obj.GetCountry())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "created_at = ?")
		params = append(params, obj.GetCreatedAt())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "updated_at = ?")
		params = append(params, obj.GetUpdatedAt())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE blogs SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE id = ?"
	params = append(params, obj.GetId())

	result, err := coredb.Exec(sql, _db, params...)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, coredb.ErrAvoidUpdate
	}
	if affectedRows > 1 {
		return false, coredb.ErrMultipleUpdate
	}

	obj.resetUpdated()
	return true, nil
}

func Update[T any](obj *T) (bool, error) {
	return coredb.Update(obj, _db)
}

func (obj *Blog) Delete() error {
	sql := `DELETE FROM blogs WHERE id = ?`

	_, err := coredb.Exec(sql, _db, obj.GetId())
	return err
}

func Delete(obj WithPK) error {
	sql := `DELETE FROM blogs WHERE id = ?`

	_, err := coredb.Exec(sql, _db, obj.GetId())
	return err
}

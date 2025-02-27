// Code generated by gola 0.0.6; DO NOT EDIT.

package blogs

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

const DBName string = "testdata"
const TableName string = "blogs"

// Blog represents `blogs` table
type Blog struct {
	//  int
	Id `json:"id"`
	// User Id int
	UserId `json:"user_id"`
	// Slug varchar
	Slug `json:"slug"`
	// Title varchar
	Title `json:"title"`
	// Category Id int
	CategoryId `json:"category_id"`
	// Is pinned to top tinyint
	IsPinned `json:"is_pinned"`
	// Is VIP reader only tinyint
	IsVip `json:"is_vip"`
	// Country of the blog user varchar
	Country `json:"country"`
	// Created Timestamp int unsigned
	CreatedAt `json:"created_at"`
	// Updated Timestamp int unsigned
	UpdatedAt `json:"updated_at"`
}

type withPK interface {
	GetId() int
}

// FetchByPK returns a row from `blogs` table with given primary key value
func FetchByPK(val int) *Blog {
	return coredb.FetchByPK[Blog](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from blogs table with given primary key value
func FetchFieldsByPK[T any](val int) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `blogs` table with given primary key values
func FetchByPKs(vals ...int) []*Blog {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Blog](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `blogs` table with given primary key values
func FetchFieldsByPKs[T any](vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *Blog {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Blog](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*Blog, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Blog](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `blogs` "+whereSQL, params...)
}

// Column types

// Id field
//
type Id struct {
	isAssigned bool
	val        int
}

func (c *Id) GetId() int {
	return c.val
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) GetValPointer() any {
	return &c.val
}

func (c *Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Id) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *UserId) IsUpdated() bool {
	return c._updated
}

func (c *UserId) resetUpdated() {
	c._updated = false
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UserId) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Slug) IsUpdated() bool {
	return c._updated
}

func (c *Slug) resetUpdated() {
	c._updated = false
}

func (c *Slug) GetColumnName() string {
	return "slug"
}

func (c *Slug) GetValPointer() any {
	return &c.val
}

func (c *Slug) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Slug) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Title) IsUpdated() bool {
	return c._updated
}

func (c *Title) resetUpdated() {
	c._updated = false
}

func (c *Title) GetColumnName() string {
	return "title"
}

func (c *Title) GetValPointer() any {
	return &c.val
}

func (c *Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Title) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *CategoryId) IsUpdated() bool {
	return c._updated
}

func (c *CategoryId) resetUpdated() {
	c._updated = false
}

func (c *CategoryId) GetColumnName() string {
	return "category_id"
}

func (c *CategoryId) GetValPointer() any {
	return &c.val
}

func (c *CategoryId) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *CategoryId) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// IsPinned field
// Is pinned to top
type IsPinned struct {
	_updated bool
	val      bool
}

func (c *IsPinned) GetIsPinned() bool {
	return c.val
}

func (c *IsPinned) SetIsPinned(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsPinned) IsUpdated() bool {
	return c._updated
}

func (c *IsPinned) resetUpdated() {
	c._updated = false
}

func (c *IsPinned) GetColumnName() string {
	return "is_pinned"
}

func (c *IsPinned) GetValPointer() any {
	return &c.val
}

func (c *IsPinned) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *IsPinned) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// IsVip field
// Is VIP reader only
type IsVip struct {
	_updated bool
	val      bool
}

func (c *IsVip) GetIsVip() bool {
	return c.val
}

func (c *IsVip) SetIsVip(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsVip) IsUpdated() bool {
	return c._updated
}

func (c *IsVip) resetUpdated() {
	c._updated = false
}

func (c *IsVip) GetColumnName() string {
	return "is_vip"
}

func (c *IsVip) GetValPointer() any {
	return &c.val
}

func (c *IsVip) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *IsVip) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Country) IsUpdated() bool {
	return c._updated
}

func (c *Country) resetUpdated() {
	c._updated = false
}

func (c *Country) GetColumnName() string {
	return "country"
}

func (c *Country) GetValPointer() any {
	return &c.val
}

func (c *Country) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Country) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *CreatedAt) IsUpdated() bool {
	return c._updated
}

func (c *CreatedAt) resetUpdated() {
	c._updated = false
}

func (c *CreatedAt) GetColumnName() string {
	return "created_at"
}

func (c *CreatedAt) GetValPointer() any {
	return &c.val
}

func (c *CreatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *CreatedAt) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *UpdatedAt) IsUpdated() bool {
	return c._updated
}

func (c *UpdatedAt) resetUpdated() {
	c._updated = false
}

func (c *UpdatedAt) GetColumnName() string {
	return "updated_at"
}

func (c *UpdatedAt) GetValPointer() any {
	return &c.val
}

func (c *UpdatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UpdatedAt) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// New return new *Blog with default values
func New() *Blog {
	return &Blog{
		Id{},
		UserId{val: int(0)},
		Slug{},
		Title{},
		CategoryId{val: int(0)},
		IsPinned{val: false},
		IsVip{val: false},
		Country{},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
	}
}

// NewWithPK takes "id"
// and returns new *Blog with given PK
func NewWithPK(val int) *Blog {
	c := &Blog{
		Id{},
		UserId{val: int(0)},
		Slug{},
		Title{},
		CategoryId{val: int(0)},
		IsPinned{val: false},
		IsVip{val: false},
		Country{},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
	}
	c.Id.val = val
	c.Id.isAssigned = true
	return c
}

const insertWithoutPK string = "INSERT IGNORE INTO `blogs` (`user_id`, `slug`, `title`, `category_id`, `is_pinned`, `is_vip`, `country`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
const insertWithPK string = "INSERT IGNORE INTO `blogs` (`id`, `user_id`, `slug`, `title`, `category_id`, `is_pinned`, `is_vip`, `country`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

// Insert Blog struct to `blogs` table
func (c *Blog) Insert() error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.Exec(DBName, insertWithPK, c.GetId(), c.GetUserId(), c.GetSlug(), c.GetTitle(), c.GetCategoryId(), c.GetIsPinned(), c.GetIsVip(), c.GetCountry(), c.GetCreatedAt(), c.GetUpdatedAt())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.Exec(DBName, insertWithoutPK, c.GetUserId(), c.GetSlug(), c.GetTitle(), c.GetCategoryId(), c.GetIsPinned(), c.GetIsVip(), c.GetCountry(), c.GetCreatedAt(), c.GetUpdatedAt())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = int(id)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return coredb.ErrAvoidInsert
	}

	c.resetUpdated()
	return nil
}

func (c *Blog) resetUpdated() {
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

// Update Blog struct in `blogs` table
func (obj *Blog) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.UserId.IsUpdated() {
		updatedFields = append(updatedFields, "`user_id` = ?")
		params = append(params, obj.GetUserId())
	}
	if obj.Slug.IsUpdated() {
		updatedFields = append(updatedFields, "`slug` = ?")
		params = append(params, obj.GetSlug())
	}
	if obj.Title.IsUpdated() {
		updatedFields = append(updatedFields, "`title` = ?")
		params = append(params, obj.GetTitle())
	}
	if obj.CategoryId.IsUpdated() {
		updatedFields = append(updatedFields, "`category_id` = ?")
		params = append(params, obj.GetCategoryId())
	}
	if obj.IsPinned.IsUpdated() {
		updatedFields = append(updatedFields, "`is_pinned` = ?")
		params = append(params, obj.GetIsPinned())
	}
	if obj.IsVip.IsUpdated() {
		updatedFields = append(updatedFields, "`is_vip` = ?")
		params = append(params, obj.GetIsVip())
	}
	if obj.Country.IsUpdated() {
		updatedFields = append(updatedFields, "`country` = ?")
		params = append(params, obj.GetCountry())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.GetCreatedAt())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.GetUpdatedAt())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.Exec(DBName, sql, params...)
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

	obj.resetUpdated()
	return true, nil
}

// Update Blog struct with given fields in `blogs` table
func Update(obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *UserId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`user_id` = ?")
				params = append(params, c.GetUserId())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Slug:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`slug` = ?")
				params = append(params, c.GetSlug())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Title:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`title` = ?")
				params = append(params, c.GetTitle())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CategoryId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`category_id` = ?")
				params = append(params, c.GetCategoryId())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsPinned:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_pinned` = ?")
				params = append(params, c.GetIsPinned())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsVip:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_vip` = ?")
				params = append(params, c.GetIsVip())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Country:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`country` = ?")
				params = append(params, c.GetCountry())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`created_at` = ?")
				params = append(params, c.GetCreatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`updated_at` = ?")
				params = append(params, c.GetUpdatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.Exec(DBName, sql, params...)
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

	for _, f := range resetFuncs {
		f()
	}
	return true, nil
}

const deleteSql string = "DELETE FROM `blogs` WHERE `id` = ?"

// DeleteByPK delete a row from blogs table with given primary key value
func DeleteByPK(val int) error {
	_, err := coredb.Exec(DBName, deleteSql, val)
	return err
}

// Code generated by gola 0.0.2; DO NOT EDIT.

package song_user_favourites

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"

	"time"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

// SongUserFavourite represents song_user_favourites table
type SongUserFavourite struct {
	// User ID int unsigned
	UserId
	// Song ID int unsigned
	SongId
	// favourite remark varchar
	Remark
	// Is favourite tinyint
	IsFavourite
	// Create Time timestamp
	CreatedAt
	// Last Update Time timestamp
	UpdatedAt
}

type SongUserFavouriteTable struct{}

func (*SongUserFavouriteTable) GetTableName() string {
	return "song_user_favourites"
}

var table *SongUserFavouriteTable

type PK struct {
	UserId uint
	SongId uint
}

type withPK interface {
	GetUserId() uint
	GetSongId() uint
}

// FetchSongUserFavouriteByPKs returns a row from song_user_favourites table with given primary key value
func FetchSongUserFavouriteByPK(val PK) *SongUserFavourite {
	return coredb.FetchByPK[SongUserFavourite](_db, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FetchByPKs returns a row with selected fields from song_user_favourites table with given primary key value
func FetchByPK[T any](val PK) *T {
	return coredb.FetchByPK[T](_db, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FindOneSongUserFavourite returns a row from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindOneSongUserFavourite(whereSQL string, params ...any) *SongUserFavourite {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[SongUserFavourite](w, _db)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM song_user_favourites "+whereSQL, _db, params...)
}

// FindOne returns a row with selected fields from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, _db)
}

// FindSongUserFavourite returns rows from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindSongUserFavourite(whereSQL string, params ...any) ([]*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[SongUserFavourite](w, _db)
}

// Find returns rows with selected fields from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, _db)
}

// Column types

// UserId field
// User ID
type UserId struct {
	_updated bool
	val      uint
}

func (c *UserId) GetUserId() uint {
	return c.val
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
	return true
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) GetTableType() coredb.TableType {
	return table
}

// SongId field
// Song ID
type SongId struct {
	_updated bool
	val      uint
}

func (c *SongId) GetSongId() uint {
	return c.val
}

func (c *SongId) GetColumnName() string {
	return "song_id"
}

func (c *SongId) IsUpdated() bool {
	return c._updated
}

func (c *SongId) resetUpdated() {
	c._updated = false
}

func (c *SongId) IsPrimaryKey() bool {
	return true
}

func (c *SongId) GetValPointer() any {
	return &c.val
}

func (c *SongId) GetTableType() coredb.TableType {
	return table
}

// Remark field
// favourite remark
type Remark struct {
	_updated bool
	val      string
}

func (c *Remark) GetRemark() string {
	return c.val
}

func (c *Remark) SetRemark(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Remark) GetColumnName() string {
	return "remark"
}

func (c *Remark) IsUpdated() bool {
	return c._updated
}

func (c *Remark) resetUpdated() {
	c._updated = false
}

func (c *Remark) IsPrimaryKey() bool {
	return false
}

func (c *Remark) GetValPointer() any {
	return &c.val
}

func (c *Remark) GetTableType() coredb.TableType {
	return table
}

// IsFavourite field
// Is favourite
type IsFavourite struct {
	_updated bool
	val      bool
}

func (c *IsFavourite) GetIsFavourite() bool {
	return c.val
}

func (c *IsFavourite) SetIsFavourite(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsFavourite) GetColumnName() string {
	return "is_favourite"
}

func (c *IsFavourite) IsUpdated() bool {
	return c._updated
}

func (c *IsFavourite) resetUpdated() {
	c._updated = false
}

func (c *IsFavourite) IsPrimaryKey() bool {
	return false
}

func (c *IsFavourite) GetValPointer() any {
	return &c.val
}

func (c *IsFavourite) GetTableType() coredb.TableType {
	return table
}

// CreatedAt field
// Create Time
type CreatedAt struct {
	_updated bool
	val      time.Time
}

func (c *CreatedAt) GetCreatedAt() time.Time {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val time.Time) bool {
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
// Last Update Time
type UpdatedAt struct {
	_updated bool
	val      time.Time
}

func (c *UpdatedAt) GetUpdatedAt() time.Time {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val time.Time) bool {
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

func NewSongUserFavouriteWithPK(val PK) *SongUserFavourite {
	c := &SongUserFavourite{
		UserId{},
		SongId{},
		Remark{},
		IsFavourite{val: true},
		CreatedAt{val: time.Now()},
		UpdatedAt{val: time.Now()},
	}
	c.UserId.val = val.UserId
	c.SongId.val = val.SongId
	return c
}

func (c *SongUserFavourite) Insert() error {
	sql := `INSERT INTO song_user_favourites (user_id, song_id, remark, is_favourite, created_at, updated_at) values (?, ?, ?, ?, ?, ?)`

	result, err := coredb.Exec(sql, _db, c.GetUserId(), c.GetSongId(), c.GetRemark(), c.GetIsFavourite(), c.GetCreatedAt(), c.GetUpdatedAt())

	if err != nil {
		return err
	}

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

func (c *SongUserFavourite) resetUpdated() {
	c.UserId.resetUpdated()
	c.SongId.resetUpdated()
	c.Remark.resetUpdated()
	c.IsFavourite.resetUpdated()
	c.CreatedAt.resetUpdated()
	c.UpdatedAt.resetUpdated()
}

func (obj *SongUserFavourite) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.UserId.IsUpdated() {
		return false, coredb.ErrPKChanged
	}
	if obj.SongId.IsUpdated() {
		return false, coredb.ErrPKChanged
	}
	if obj.Remark.IsUpdated() {
		updatedFields = append(updatedFields, "remark = ?")
		params = append(params, obj.GetRemark())
	}
	if obj.IsFavourite.IsUpdated() {
		updatedFields = append(updatedFields, "is_favourite = ?")
		params = append(params, obj.GetIsFavourite())
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

	sql := "UPDATE song_user_favourites SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE user_id = ? and song_id = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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
		case *Remark:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "remark = ?")
				params = append(params, c.GetRemark())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsFavourite:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "is_favourite = ?")
				params = append(params, c.GetIsFavourite())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "created_at = ?")
				params = append(params, c.GetCreatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "updated_at = ?")
				params = append(params, c.GetUpdatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE song_user_favourites SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE user_id = ? and song_id = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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

	for _, f := range resetFuncs {
		f()
	}
	return true, nil
}

func (obj *SongUserFavourite) Delete() error {
	sql := `DELETE FROM song_user_favourites WHERE user_id = ? and song_id = ?`

	_, err := coredb.Exec(sql, _db, obj.GetUserId(), obj.GetSongId())
	return err
}

func Delete(obj withPK) error {
	sql := `DELETE FROM song_user_favourites WHERE user_id = ? and song_id = ?`

	_, err := coredb.Exec(sql, _db, obj.GetUserId(), obj.GetSongId())
	return err
}

// Code generated by gola 0.0.3; DO NOT EDIT.

package songs

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

// Song represents songs table
type Song struct {
	//  int unsigned
	Id
	// Song title varchar
	Title
	// Song Ranking mediumint
	Rank
	//  enum('','101','1+9','%1')
	Type
	// Song file hash checksum varchar
	Hash
}

type SongTable struct{}

func (*SongTable) GetTableName() string {
	return "songs"
}

var table *SongTable

type withPK interface {
	GetId() uint
}

// FetchSongByPKs returns a row from songs table with given primary key value
func FetchSongByPK(val uint) *Song {
	return coredb.FetchByPK[Song](_db, []string{"id"}, val)
}

// FetchByPKs returns a row with selected fields from songs table with given primary key value
func FetchByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](_db, []string{"id"}, val)
}

// FetchSongByPKs returns rows with from songs table with given primary key values
func FetchSongByPKs(vals ...uint) []*Song {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Song](pks, "`id`", _db)
}

// FetchByPKs returns rows with selected fields from songs table with given primary key values
func FetchByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](pks, "`id`", _db)
}

// FindOneSong returns a row from songs table with arbitary where query
// whereSQL must start with "where ..."
func FindOneSong(whereSQL string, params ...any) *Song {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Song](w, _db)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM `songs` "+whereSQL, _db, params...)
}

// FindOne returns a row with selected fields from songs table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, _db)
}

// FindSong returns rows from songs table with arbitary where query
// whereSQL must start with "where ..."
func FindSong(whereSQL string, params ...any) ([]*Song, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Song](w, _db)
}

// Find returns rows with selected fields from songs table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, _db)
}

// Column types
type SongType string

const (
	SongTypeEmpty SongType = ""
	SongType101   SongType = "101"
	SongType1x2B9 SongType = "1+9"
	SongTypex251  SongType = "%1"
)

// Id field
//
type Id struct {
	val uint
}

func (c *Id) GetId() uint {
	return c.val
}

func (c *Id) GetColumnName() string {
	return "id"
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

// Title field
// Song title
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

func (c *Title) IsPrimaryKey() bool {
	return false
}

func (c *Title) GetValPointer() any {
	return &c.val
}

func (c *Title) GetTableType() coredb.TableType {
	return table
}

// Rank field
// Song Ranking
type Rank struct {
	_updated bool
	val      int
}

func (c *Rank) GetRank() int {
	return c.val
}

func (c *Rank) SetRank(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Rank) IsUpdated() bool {
	return c._updated
}

func (c *Rank) resetUpdated() {
	c._updated = false
}

func (c *Rank) GetColumnName() string {
	return "rank"
}

func (c *Rank) IsPrimaryKey() bool {
	return false
}

func (c *Rank) GetValPointer() any {
	return &c.val
}

func (c *Rank) GetTableType() coredb.TableType {
	return table
}

// Type field
//
type Type struct {
	_updated bool
	val      SongType
}

func (c *Type) GetType() SongType {
	return c.val
}

func (c *Type) SetType(val SongType) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Type) IsUpdated() bool {
	return c._updated
}

func (c *Type) resetUpdated() {
	c._updated = false
}

func (c *Type) GetColumnName() string {
	return "type"
}

func (c *Type) IsPrimaryKey() bool {
	return false
}

func (c *Type) GetValPointer() any {
	return &c.val
}

func (c *Type) GetTableType() coredb.TableType {
	return table
}

// Hash field
// Song file hash checksum
type Hash struct {
	_updated bool
	val      string
}

func (c *Hash) GetHash() string {
	return c.val
}

func (c *Hash) SetHash(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Hash) IsUpdated() bool {
	return c._updated
}

func (c *Hash) resetUpdated() {
	c._updated = false
}

func (c *Hash) GetColumnName() string {
	return "hash"
}

func (c *Hash) IsPrimaryKey() bool {
	return false
}

func (c *Hash) GetValPointer() any {
	return &c.val
}

func (c *Hash) GetTableType() coredb.TableType {
	return table
}

func NewSong() *Song {
	return &Song{
		Id{},
		Title{},
		Rank{val: int(0)},
		Type{},
		Hash{},
	}
}

func NewSongWithPK(val uint) *Song {
	c := &Song{
		Id{},
		Title{},
		Rank{val: int(0)},
		Type{},
		Hash{},
	}
	c.Id.val = val
	return c
}

func (c *Song) Insert() error {
	sql := "INSERT IGNORE INTO `songs` (`title`, `rank`, `type`, `hash`) values (?, ?, ?, ?)"

	result, err := coredb.Exec(sql, _db, c.GetTitle(), c.GetRank(), c.GetType(), c.GetHash())

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.Id.val = uint(id)

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

func (c *Song) resetUpdated() {
	c.Title.resetUpdated()
	c.Rank.resetUpdated()
	c.Type.resetUpdated()
	c.Hash.resetUpdated()
}

func (obj *Song) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Title.IsUpdated() {
		updatedFields = append(updatedFields, "`title` = ?")
		params = append(params, obj.GetTitle())
	}
	if obj.Rank.IsUpdated() {
		updatedFields = append(updatedFields, "`rank` = ?")
		params = append(params, obj.GetRank())
	}
	if obj.Type.IsUpdated() {
		updatedFields = append(updatedFields, "`type` = ?")
		params = append(params, obj.GetType())
	}
	if obj.Hash.IsUpdated() {
		updatedFields = append(updatedFields, "`hash` = ?")
		params = append(params, obj.GetHash())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `songs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
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
		case *Title:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`title` = ?")
				params = append(params, c.GetTitle())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Rank:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`rank` = ?")
				params = append(params, c.GetRank())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Type:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`type` = ?")
				params = append(params, c.GetType())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Hash:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`hash` = ?")
				params = append(params, c.GetHash())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `songs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
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

	for _, f := range resetFuncs {
		f()
	}
	return true, nil
}

func (obj *Song) Delete() error {
	sql := "DELETE FROM `songs` WHERE `id` = ?"

	_, err := coredb.Exec(sql, _db, obj.GetId())
	return err
}

func Delete(obj withPK) error {
	sql := "DELETE FROM `songs` WHERE `id` = ?"

	_, err := coredb.Exec(sql, _db, obj.GetId())
	return err
}

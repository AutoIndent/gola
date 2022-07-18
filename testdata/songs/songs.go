// Code generated by gola 0.0.2; DO NOT EDIT.

package songs

import (
	"database/sql"
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
	// Song title varchar(100)
	Title
	// Song file hash checksum varchar(128)
	Hash
}

type SongTable struct{}

func (*SongTable) GetTableName() string {
	return "songs"
}

var table *SongTable

// FetchSongByPKs returns a row from songs table with given primary key value
func FetchSongByPK(val uint) *Song {
	return coredb.FetchByPK[Song](val, "id", _db)
}

// FetchByPKs returns a row with selected fields from songs table with given primary key value
func FetchByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](val, "id", _db)
}

// FetchSongByPKs returns rows with from songs table with given primary key values
func FetchSongByPKs(vals ...uint) []*Song {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Song](pks, "id", _db)
}

// FetchByPKs returns rows with selected fields from songs table with given primary key values
func FetchByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](pks, "id", _db)
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
	return coredb.QueryInt("SELECT COUNT(*) FROM songs "+whereSQL, _db, params...)
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

// Id field
//
type Id struct {
	_updated bool
	val      uint
}

func (c *Id) GetId() uint {
	return c.val
}

func (c *Id) SetId(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
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

func (c *Hash) GetColumnName() string {
	return "hash"
}

func (c *Hash) IsUpdated() bool {
	return c._updated
}

func (c *Hash) resetUpdated() {
	c._updated = false
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
		Hash{val: ""},
	}
}

func (c *Song) Insert() error {
	sql := `INSERT INTO songs (title, hash) values (?, ?)`

	result, err := coredb.Exec(sql, _db, c.GetTitle(), c.GetHash())

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.SetId(uint(id))

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

func (c *Song) resetUpdated() {
	c.Id.resetUpdated()
	c.Title.resetUpdated()
	c.Hash.resetUpdated()
}

func (c *Song) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if c.Title.IsUpdated() {
		updatedFields = append(updatedFields, "title = ?")
		params = append(params, c.GetTitle())
	}
	if c.Hash.IsUpdated() {
		updatedFields = append(updatedFields, "hash = ?")
		params = append(params, c.GetHash())
	}

	sql := `UPDATE songs SET `

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql = sql + strings.Join(updatedFields, ",") + " WHERE id = ?"
	params = append(params, c.GetId())

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

	c.resetUpdated()
	return true, nil
}

func (c *Song) Delete() error {
	sql := `DELETE FROM songs WHERE id = ?`

	_, err := coredb.Exec(sql, _db, c.GetId())
	return err
}

func Update[T any](obj *T) (bool, error) {
	return coredb.Update(obj, _db)
}

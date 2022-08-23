// Code generated by gola 0.0.6; DO NOT EDIT.

package songs

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

const DBName string = "testdata"
const TableName string = "songs"

// Song represents `songs` table
type Song struct {
	//  int unsigned
	Id `json:"id"`
	// Song title varchar
	Title `json:"title"`
	// Song Ranking mediumint
	Rank `json:"rank"`
	//  enum('','101','1+9','%1','0.9')
	Type `json:"type"`
	// Song file hash checksum varchar
	Hash `json:"hash"`
}

type withPK interface {
	GetId() uint
}

// FetchByPK returns a row from `songs` table with given primary key value
func FetchByPK(val uint) *Song {
	return coredb.FetchByPK[Song](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from songs table with given primary key value
func FetchFieldsByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `songs` table with given primary key values
func FetchByPKs(vals ...uint) []*Song {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Song](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `songs` table with given primary key values
func FetchFieldsByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `songs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *Song {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Song](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `songs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `songs` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*Song, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Song](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `songs` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `songs` "+whereSQL, params...)
}

// Column types
type SongType string

const (
	SongTypeEmpty SongType = ""
	SongType101   SongType = "101"
	SongType1x2B9 SongType = "1+9"
	SongTypex251  SongType = "%1"
	SongType0dot9 SongType = "0.9"
)

// Id field
//
type Id struct {
	isAssigned bool
	val        uint
}

func (c *Id) GetId() uint {
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

func (c *Rank) GetValPointer() any {
	return &c.val
}

func (c *Rank) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Rank) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Type) GetValPointer() any {
	return &c.val
}

func (c *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Type) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Hash) GetValPointer() any {
	return &c.val
}

func (c *Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Hash) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// New return new *Song with default values
func New() *Song {
	return &Song{
		Id{},
		Title{},
		Rank{val: int(0)},
		Type{},
		Hash{},
	}
}

// NewWithPK takes "id"
// and returns new *Song with given PK
func NewWithPK(val uint) *Song {
	c := &Song{
		Id{},
		Title{},
		Rank{val: int(0)},
		Type{},
		Hash{},
	}
	c.Id.val = val
	c.Id.isAssigned = true
	return c
}

const insertWithoutPK string = "INSERT IGNORE INTO `songs` (`title`, `rank`, `type`, `hash`) values (?, ?, ?, ?)"
const insertWithPK string = "INSERT IGNORE INTO `songs` (`id`, `title`, `rank`, `type`, `hash`) values (?, ?, ?, ?, ?)"

// Insert Song struct to `songs` table
func (c *Song) Insert() error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.Exec(DBName, insertWithPK, c.GetId(), c.GetTitle(), c.GetRank(), c.GetType(), c.GetHash())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.Exec(DBName, insertWithoutPK, c.GetTitle(), c.GetRank(), c.GetType(), c.GetHash())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = uint(id)
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

func (c *Song) resetUpdated() {
	c.Title.resetUpdated()
	c.Rank.resetUpdated()
	c.Type.resetUpdated()
	c.Hash.resetUpdated()
}

// Update Song struct in `songs` table
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

// Update Song struct with given fields in `songs` table
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

const deleteSql string = "DELETE FROM `songs` WHERE `id` = ?"

// DeleteByPK delete a row from songs table with given primary key value
func DeleteByPK(val uint) error {
	_, err := coredb.Exec(DBName, deleteSql, val)
	return err
}

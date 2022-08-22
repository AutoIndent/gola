// Code generated by gola 0.0.4; DO NOT EDIT.

package users

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

const DBName string = "testdata"
const TableName string = "users"

// User represents `users` table
type User struct {
	//  int
	Id
	// Name varchar
	Name
	// Email address varchar
	Email
	// Created Timestamp int unsigned
	CreatedAt
	// Updated Timestamp int unsigned
	UpdatedAt
	// float float
	FloatType
	// double double
	DoubleType
	// user hobby enum('swimming','running','singing')
	Hobby
	// user hobby enum('swimming','running','singing')
	HobbyNoDefault
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	Sports
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	Sports2
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	SportsNoDefault
}

type withPK interface {
	GetId() int
}

// FetchByPK returns a row from `users` table with given primary key value
func FetchByPK(val int) *User {
	return coredb.FetchByPK[User](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from users table with given primary key value
func FetchFieldsByPK[T any](val int) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `users` table with given primary key values
func FetchByPKs(vals ...int) []*User {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[User](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `users` table with given primary key values
func FetchFieldsByPKs[T any](vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *User {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[User](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `users` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[User](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `users` "+whereSQL, params...)
}

// Column types
type UserHobby string

const (
	UserHobbySwimming UserHobby = "swimming"
	UserHobbyRunning  UserHobby = "running"
	UserHobbySinging  UserHobby = "singing"
)

type UserHobbyNoDefault string

const (
	UserHobbyNoDefaultSwimming UserHobbyNoDefault = "swimming"
	UserHobbyNoDefaultRunning  UserHobbyNoDefault = "running"
	UserHobbyNoDefaultSinging  UserHobbyNoDefault = "singing"
)

type UserSports string

const (
	UserSportsSwim       UserSports = "swim"
	UserSportsTennis     UserSports = "tennis"
	UserSportsBasketball UserSports = "basketball"
	UserSportsFootball   UserSports = "football"
	UserSportsSquash     UserSports = "squash"
	UserSportsBadminton  UserSports = "badminton"
)

type UserSports2 string

const (
	UserSports2Swim       UserSports2 = "swim"
	UserSports2Tennis     UserSports2 = "tennis"
	UserSports2Basketball UserSports2 = "basketball"
	UserSports2Football   UserSports2 = "football"
	UserSports2Squash     UserSports2 = "squash"
	UserSports2Badminton  UserSports2 = "badminton"
)

type UserSportsNoDefault string

const (
	UserSportsNoDefaultSwim       UserSportsNoDefault = "swim"
	UserSportsNoDefaultTennis     UserSportsNoDefault = "tennis"
	UserSportsNoDefaultBasketball UserSportsNoDefault = "basketball"
	UserSportsNoDefaultFootball   UserSportsNoDefault = "football"
	UserSportsNoDefaultSquash     UserSportsNoDefault = "squash"
	UserSportsNoDefaultBadminton  UserSportsNoDefault = "badminton"
)

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

// Name field
// Name
type Name struct {
	_updated bool
	val      string
}

func (c *Name) GetName() string {
	return c.val
}

func (c *Name) SetName(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Name) IsUpdated() bool {
	return c._updated
}

func (c *Name) resetUpdated() {
	c._updated = false
}

func (c *Name) GetColumnName() string {
	return "name"
}

func (c *Name) GetValPointer() any {
	return &c.val
}

// Email field
// Email address
type Email struct {
	_updated bool
	val      string
}

func (c *Email) GetEmail() string {
	return c.val
}

func (c *Email) SetEmail(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Email) IsUpdated() bool {
	return c._updated
}

func (c *Email) resetUpdated() {
	c._updated = false
}

func (c *Email) GetColumnName() string {
	return "email"
}

func (c *Email) GetValPointer() any {
	return &c.val
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

// FloatType field
// float
type FloatType struct {
	_updated bool
	val      float32
}

func (c *FloatType) GetFloatType() float32 {
	return c.val
}

func (c *FloatType) SetFloatType(val float32) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *FloatType) IsUpdated() bool {
	return c._updated
}

func (c *FloatType) resetUpdated() {
	c._updated = false
}

func (c *FloatType) GetColumnName() string {
	return "float_type"
}

func (c *FloatType) GetValPointer() any {
	return &c.val
}

// DoubleType field
// double
type DoubleType struct {
	_updated bool
	val      float64
}

func (c *DoubleType) GetDoubleType() float64 {
	return c.val
}

func (c *DoubleType) SetDoubleType(val float64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *DoubleType) IsUpdated() bool {
	return c._updated
}

func (c *DoubleType) resetUpdated() {
	c._updated = false
}

func (c *DoubleType) GetColumnName() string {
	return "double_type"
}

func (c *DoubleType) GetValPointer() any {
	return &c.val
}

// Hobby field
// user hobby
type Hobby struct {
	_updated bool
	val      UserHobby
}

func (c *Hobby) GetHobby() UserHobby {
	return c.val
}

func (c *Hobby) SetHobby(val UserHobby) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Hobby) IsUpdated() bool {
	return c._updated
}

func (c *Hobby) resetUpdated() {
	c._updated = false
}

func (c *Hobby) GetColumnName() string {
	return "hobby"
}

func (c *Hobby) GetValPointer() any {
	return &c.val
}

// HobbyNoDefault field
// user hobby
type HobbyNoDefault struct {
	_updated bool
	val      UserHobbyNoDefault
}

func (c *HobbyNoDefault) GetHobbyNoDefault() UserHobbyNoDefault {
	return c.val
}

func (c *HobbyNoDefault) SetHobbyNoDefault(val UserHobbyNoDefault) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *HobbyNoDefault) IsUpdated() bool {
	return c._updated
}

func (c *HobbyNoDefault) resetUpdated() {
	c._updated = false
}

func (c *HobbyNoDefault) GetColumnName() string {
	return "hobby_no_default"
}

func (c *HobbyNoDefault) GetValPointer() any {
	return &c.val
}

// Sports field
// user sports
type Sports struct {
	_updated bool
	val      string
}

func (c *Sports) GetSports() []UserSports {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSports, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSports(strings.ToLower(s)))
	}
	return valSlice
}

func (c *Sports) SetSports(val []UserSports) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *Sports) IsUpdated() bool {
	return c._updated
}

func (c *Sports) resetUpdated() {
	c._updated = false
}

func (c *Sports) GetColumnName() string {
	return "sports"
}

func (c *Sports) GetValPointer() any {
	return &c.val
}

// Sports2 field
// user sports
type Sports2 struct {
	_updated bool
	val      string
}

func (c *Sports2) GetSports2() []UserSports2 {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSports2, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSports2(strings.ToLower(s)))
	}
	return valSlice
}

func (c *Sports2) SetSports2(val []UserSports2) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *Sports2) IsUpdated() bool {
	return c._updated
}

func (c *Sports2) resetUpdated() {
	c._updated = false
}

func (c *Sports2) GetColumnName() string {
	return "sports2"
}

func (c *Sports2) GetValPointer() any {
	return &c.val
}

// SportsNoDefault field
// user sports
type SportsNoDefault struct {
	_updated bool
	val      string
}

func (c *SportsNoDefault) GetSportsNoDefault() []UserSportsNoDefault {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSportsNoDefault, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSportsNoDefault(strings.ToLower(s)))
	}
	return valSlice
}

func (c *SportsNoDefault) SetSportsNoDefault(val []UserSportsNoDefault) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *SportsNoDefault) IsUpdated() bool {
	return c._updated
}

func (c *SportsNoDefault) resetUpdated() {
	c._updated = false
}

func (c *SportsNoDefault) GetColumnName() string {
	return "sports_no_default"
}

func (c *SportsNoDefault) GetValPointer() any {
	return &c.val
}

// New return new *User with default values
func New() *User {
	return &User{
		Id{},
		Name{},
		Email{},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
		FloatType{val: float32(0)},
		DoubleType{val: float64(0)},
		Hobby{val: "swimming"},
		HobbyNoDefault{},
		Sports{val: "swim, football"},
		Sports2{val: "swim,football"},
		SportsNoDefault{},
	}
}

// NewWithPK takes "id"
// and returns new *User with given PK
func NewWithPK(val int) *User {
	c := &User{
		Id{},
		Name{},
		Email{},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
		FloatType{val: float32(0)},
		DoubleType{val: float64(0)},
		Hobby{val: "swimming"},
		HobbyNoDefault{},
		Sports{val: "swim, football"},
		Sports2{val: "swim,football"},
		SportsNoDefault{},
	}
	c.Id.val = val
	c.Id.isAssigned = true
	return c
}

const insertWithoutPK string = "INSERT IGNORE INTO `users` (`name`, `email`, `created_at`, `updated_at`, `float_type`, `double_type`, `hobby`, `hobby_no_default`, `sports`, `sports2`, `sports_no_default`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
const insertWithPK string = "INSERT IGNORE INTO `users` (`id`, `name`, `email`, `created_at`, `updated_at`, `float_type`, `double_type`, `hobby`, `hobby_no_default`, `sports`, `sports2`, `sports_no_default`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

// Insert User struct to `users` table
func (c *User) Insert() error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.Exec(DBName, insertWithPK, c.GetId(), c.GetName(), c.GetEmail(), c.GetCreatedAt(), c.GetUpdatedAt(), c.GetFloatType(), c.GetDoubleType(), c.GetHobby(), c.GetHobbyNoDefault(), c.GetSports(), c.GetSports2(), c.GetSportsNoDefault())
	} else {
		result, err = coredb.Exec(DBName, insertWithoutPK, c.GetName(), c.GetEmail(), c.GetCreatedAt(), c.GetUpdatedAt(), c.GetFloatType(), c.GetDoubleType(), c.GetHobby(), c.GetHobbyNoDefault(), c.GetSports(), c.GetSports2(), c.GetSportsNoDefault())
	}

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
	if affectedRows == 0 {
		return coredb.ErrAvoidInsert
	}

	c.resetUpdated()
	return nil
}

func (c *User) resetUpdated() {
	c.Name.resetUpdated()
	c.Email.resetUpdated()
	c.CreatedAt.resetUpdated()
	c.UpdatedAt.resetUpdated()
	c.FloatType.resetUpdated()
	c.DoubleType.resetUpdated()
	c.Hobby.resetUpdated()
	c.HobbyNoDefault.resetUpdated()
	c.Sports.resetUpdated()
	c.Sports2.resetUpdated()
	c.SportsNoDefault.resetUpdated()
}

// Update User struct in `users` table
func (obj *User) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Name.IsUpdated() {
		updatedFields = append(updatedFields, "`name` = ?")
		params = append(params, obj.GetName())
	}
	if obj.Email.IsUpdated() {
		updatedFields = append(updatedFields, "`email` = ?")
		params = append(params, obj.GetEmail())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.GetCreatedAt())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.GetUpdatedAt())
	}
	if obj.FloatType.IsUpdated() {
		updatedFields = append(updatedFields, "`float_type` = ?")
		params = append(params, obj.GetFloatType())
	}
	if obj.DoubleType.IsUpdated() {
		updatedFields = append(updatedFields, "`double_type` = ?")
		params = append(params, obj.GetDoubleType())
	}
	if obj.Hobby.IsUpdated() {
		updatedFields = append(updatedFields, "`hobby` = ?")
		params = append(params, obj.GetHobby())
	}
	if obj.HobbyNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "`hobby_no_default` = ?")
		params = append(params, obj.GetHobbyNoDefault())
	}
	if obj.Sports.IsUpdated() {
		updatedFields = append(updatedFields, "`sports` = ?")
		params = append(params, obj.GetSports())
	}
	if obj.Sports2.IsUpdated() {
		updatedFields = append(updatedFields, "`sports2` = ?")
		params = append(params, obj.GetSports2())
	}
	if obj.SportsNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "`sports_no_default` = ?")
		params = append(params, obj.GetSportsNoDefault())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `users` SET "
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

// Update User struct with given fields in `users` table
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
		case *Name:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`name` = ?")
				params = append(params, c.GetName())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Email:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`email` = ?")
				params = append(params, c.GetEmail())
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
		case *FloatType:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`float_type` = ?")
				params = append(params, c.GetFloatType())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *DoubleType:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`double_type` = ?")
				params = append(params, c.GetDoubleType())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Hobby:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`hobby` = ?")
				params = append(params, c.GetHobby())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *HobbyNoDefault:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`hobby_no_default` = ?")
				params = append(params, c.GetHobbyNoDefault())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Sports:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports` = ?")
				params = append(params, c.GetSports())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Sports2:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports2` = ?")
				params = append(params, c.GetSports2())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *SportsNoDefault:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports_no_default` = ?")
				params = append(params, c.GetSportsNoDefault())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `users` SET "
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

const deleteSql string = "DELETE FROM `users` WHERE `id` = ?"

// DeleteByPK delete a row from users table with given primary key value
func DeleteByPK(val int) error {
	_, err := coredb.Exec(DBName, deleteSql, val)
	return err
}

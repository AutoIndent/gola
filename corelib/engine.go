package corelib

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var _db *sql.DB

// Setup the default db instance
func Setup(db *sql.DB) {
	_db = db
}

var typeColumnNames = make(map[reflect.Type]string)
var typeTableNames = make(map[reflect.Type]string)

func getDB(db *sql.DB) *sql.DB {
	if db != nil {
		return db
	}
	if _db != nil {
		return _db
	}
	panic("No db instance available")
}

// FetchByPK returns a row of T type with given primary key value
func FetchByPK[T any](val interface{}, pkName string, db *sql.DB) *T {
	u := new(T)
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	data := StrutForScan(u)

	query := fmt.Sprintf("SELECT %s from %s where %s=?", columnsNames, tableName, pkName)

	mydb := getDB(db)
	err2 := mydb.QueryRow(query, val).Scan(data...)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err2)
	}

	return u
}

// FetchByPKs returns rows of T type with given primary key values
func FetchByPKs[T any](vals []interface{}, pkName string, db *sql.DB) []*T {
	tableName, columnsNames := GetTableAndColumnsNames[T]()

	// idstr := JoinInts(ids, ",")
	query := fmt.Sprintf("SELECT %s from %s where %s in(%s)",
		columnsNames, tableName, pkName, GetParamPlaceHolder(len(vals)))

	return Query[T](query, db, vals...)
}

// Exec given query with given db instances or default
func Exec(query string, db *sql.DB, params ...interface{}) (sql.Result, error) {
	mydb := getDB(db)
	return mydb.Exec(query, params...)
}

// FindOne returns a row from given table type with where query
func FindOne[T any](where WhereQuery, db *sql.DB) *T {
	u := new(T)
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s from %s where %s", columnsNames,
		tableName, whereSQL)

	mydb := getDB(db)
	err2 := mydb.QueryRow(query, params...).Scan(data...)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err2)
	}

	return u
}

// Find returns rows from given table type with where query
func Find[T any](where WhereQuery, db *sql.DB) []*T {
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s from %s %s", columnsNames,
		tableName, whereSQL)

	return Query[T](query, db, params...)
}

// Query rows from given table type with where query & params
func Query[T any](query string, db *sql.DB, params ...interface{}) []*T {
	var result []*T
	var u *T

	mydb := getDB(db)
	rows, err2 := mydb.Query(query, params...)

	if err2 != nil {
		log.Fatal(err2)
	}

	for rows.Next() {
		u = new(T)
		data := StrutForScan(u)
		rows.Scan(data...)
		result = append(result, u)
	}

	return result
}

// Update given obj changes with given db instances or default
func Update[T any](obj *T, db *sql.DB) (bool, error) {
	return true, nil
}

// GetTableAndColumnsNames returns tablesName & column names joined by , of given type
func GetTableAndColumnsNames[T any]() (tableName string, joinedColumnNames string) {
	var o *T
	t := reflect.TypeOf(o)
	joinedColumnNames, ok := typeColumnNames[t]
	if ok {
		tableName = typeTableNames[t]
		return
	}

	o = new(T)
	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			columnNames = append(columnNames, f.GetColumnName())
			if tableName == "" {
				tableName = f.GetTableType().GetTableName()
			}
		}
	}

	joinedColumnNames = strings.Join(columnNames, ",")
	typeTableNames[t] = tableName
	typeColumnNames[t] = joinedColumnNames

	return
}

// StrutForScan returns value pointers of given obj
func StrutForScan[T any](u *T) (pointers []interface{}) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]interface{}, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			pointers = append(pointers, f.GetValPointer())
		}
	}
	return
}

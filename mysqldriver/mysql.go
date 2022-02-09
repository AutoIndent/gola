package mysqldriver

// modified from: https://github.com/volatiletech/sqlboiler/blob/v4.6.0/drivers/sqlboiler-mysql/driver/mysql.go

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/structs"
	"github.com/pkg/errors"
)

// Assemble is more useful for calling into the library so you don't
// have to instantiate an empty type.
func Assemble(config Config) (dbinfo *structs.DBInfo, err error) {
	driver := MySQLDriver{}
	return driver.Assemble(config)
}

// MySQLDriver holds the database connection string and a handle
// to the database connection.
type MySQLDriver struct {
	connStr string
	conn    *sql.DB

	tinyIntAsInt bool
}

// Assemble all the information we need to provide back to the driver
func (m *MySQLDriver) Assemble(config Config) (dbinfo *structs.DBInfo, err error) {
	defer func() {
		if r := recover(); r != nil && err == nil {
			dbinfo = nil
			err = r.(error)
		}
	}()

	user := config.MustString(structs.ConfigUser)
	pass, _ := config.String(structs.ConfigPass)
	dbname := config.MustString(structs.ConfigDBName)
	host := config.MustString(structs.ConfigHost)
	port := config.DefaultInt(structs.ConfigPort, 3306)
	sslmode := config.DefaultString(structs.ConfigSSLMode, "true")

	schema := dbname
	whitelist, _ := config.StringSlice(structs.ConfigWhitelist)
	blacklist, _ := config.StringSlice(structs.ConfigBlacklist)

	tinyIntAsIntIntf, ok := config["tinyint_as_int"]
	if ok {
		if b, ok := tinyIntAsIntIntf.(bool); ok {
			m.tinyIntAsInt = b
		}
	}

	m.connStr = MySQLBuildQueryString(user, pass, dbname, host, port, sslmode)
	m.conn, err = sql.Open("mysql", m.connStr)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler-mysql failed to connect to database")
	}

	defer func() {
		if e := m.conn.Close(); e != nil {
			dbinfo = nil
			err = e
		}
	}()

	dbinfo = &structs.DBInfo{}

	dbinfo.Schema = schema
	dbinfo.Tables, err = structs.Tables(m, schema, whitelist, blacklist)
	if err != nil {
		return nil, err
	}

	return dbinfo, err
}

// MySQLBuildQueryString builds a query string for MySQL.
func MySQLBuildQueryString(user, pass, dbname, host string, port int, sslmode string) string {
	config := mysql.NewConfig()

	config.User = user
	if len(pass) != 0 {
		config.Passwd = pass
	}
	config.DBName = dbname
	config.Net = "tcp"
	config.Addr = host
	if port == 0 {
		port = 3306
	}
	config.Addr += ":" + strconv.Itoa(port)
	config.TLSConfig = sslmode

	// MySQL is a bad, and by default reads date/datetime into a []byte
	// instead of a time.Time. Tell it to stop being a bad.
	config.ParseTime = true

	return config.FormatDSN()
}

// TableNames connects to the mysql database and
// retrieves all table names from the information_schema where the
// table schema is public.
func (m *MySQLDriver) TableNames(schema string, whitelist, blacklist []string) ([]string, error) {
	var names []string

	query := `select table_name from information_schema.tables where table_schema = ? and table_type = 'BASE TABLE'`
	args := []interface{}{schema}
	if len(whitelist) > 0 {
		tables := TablesFromList(whitelist)
		if len(tables) > 0 {
			query += fmt.Sprintf(" and table_name in (%s)", strings.Repeat(",?", len(tables))[1:])
			for _, w := range tables {
				args = append(args, w)
			}
		}
	} else if len(blacklist) > 0 {
		tables := TablesFromList(blacklist)
		if len(tables) > 0 {
			query += fmt.Sprintf(" and table_name not in (%s)", strings.Repeat(",?", len(tables))[1:])
			for _, b := range tables {
				args = append(args, b)
			}
		}
	}

	query += ` order by table_name;`

	rows, err := m.conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}

// Columns takes a table name and attempts to retrieve the table information
// from the database information_schema.columns. It retrieves the column names
// and column types and returns those as a []Column after TranslateColumnType()
// converts the SQL types to Go types, for example: "varchar" to "string"
func (m *MySQLDriver) Columns(schema string, table *structs.Table, tableName string, whitelist, blacklist []string) ([]structs.Column, error) {
	var columns []structs.Column
	args := []interface{}{tableName, schema}

	query := `
	select
	c.column_name,
	c.column_type,
	c.column_comment,
	if(c.data_type = 'enum', c.column_type, c.data_type),
	if(extra = 'auto_increment','auto_increment',
		if(version() like '%MariaDB%' and c.column_default = 'NULL', '',
		if(version() like '%MariaDB%' and c.data_type in ('varchar','char','binary','date','datetime','time'),
			replace(substring(c.column_default,2,length(c.column_default)-2),'\'\'','\''),
				c.column_default))),
	c.is_nullable = 'YES',
		0 as is_unique
	from information_schema.columns as c
	where table_name = ? and table_schema = ? and c.extra not like '%VIRTUAL%'`

	if len(whitelist) > 0 {
		cols := ColumnsFromList(whitelist, tableName)
		if len(cols) > 0 {
			query += fmt.Sprintf(" and c.column_name in (%s)", strings.Repeat(",?", len(cols))[1:])
			for _, w := range cols {
				args = append(args, w)
			}
		}
	} else if len(blacklist) > 0 {
		cols := ColumnsFromList(blacklist, tableName)
		if len(cols) > 0 {
			query += fmt.Sprintf(" and c.column_name not in (%s)", strings.Repeat(",?", len(cols))[1:])
			for _, w := range cols {
				args = append(args, w)
			}
		}
	}

	query += ` order by c.ordinal_position;`

	rows, err := m.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var colName, colFullType, colComment, colType string
		var nullable, unique bool
		var defaultValue *string
		if err := rows.Scan(&colName, &colFullType, &colComment, &colType, &defaultValue, &nullable, &unique); err != nil {
			return nil, errors.Wrapf(err, "unable to scan for table %s", tableName)
		}

		column := structs.Column{
			Name:       colName,
			Comment:    colComment,
			FullDBType: colFullType, // example: tinyint(1) instead of tinyint
			DBType:     colType,
			Nullable:   nullable,
			Unique:     unique,
		}

		if defaultValue != nil && *defaultValue != "NULL" {
			column.Default = *defaultValue
		}

		column.Comment = strings.ReplaceAll(column.Comment, "\r\n", " ")
		column.Comment = strings.ReplaceAll(column.Comment, "\n", " ")
		column.Comment = strings.ReplaceAll(column.Comment, "\"", "'")

		column.Table = table

		columns = append(columns, column)
	}

	return columns, nil
}

// PrimaryKeyInfo looks up the primary key for a table.
func (m *MySQLDriver) PrimaryKeyInfo(schema, tableName string) (*structs.PrimaryKey, error) {
	pkey := &structs.PrimaryKey{}
	var err error

	query := `
	select tc.constraint_name
	from information_schema.table_constraints as tc
	where tc.table_name = ? and tc.constraint_type = 'PRIMARY KEY' and tc.table_schema = ?;`

	row := m.conn.QueryRow(query, tableName, schema)
	if err = row.Scan(&pkey.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	queryColumns := `
	select kcu.column_name
	from   information_schema.key_column_usage as kcu
	where  table_name = ? and constraint_name = ? and table_schema = ?
	order by kcu.ordinal_position;`

	var rows *sql.Rows
	if rows, err = m.conn.Query(queryColumns, tableName, pkey.Name, schema); err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var column string

		err = rows.Scan(&column)
		if err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	pkey.Columns = columns

	return pkey, nil
}

// ForeignKeyInfo retrieves the foreign keys for a given table name.
func (m *MySQLDriver) ForeignKeyInfo(schema, tableName string) ([]structs.ForeignKey, error) {
	var fkeys []structs.ForeignKey

	query := `
	select constraint_name, table_name, column_name, referenced_table_name, referenced_column_name
	from information_schema.key_column_usage
	where table_schema = ? and referenced_table_schema = ? and table_name = ?
	order by constraint_name, table_name, column_name, referenced_table_name, referenced_column_name
	`

	var rows *sql.Rows
	var err error
	if rows, err = m.conn.Query(query, schema, schema, tableName); err != nil {
		return nil, err
	}

	for rows.Next() {
		var fkey structs.ForeignKey
		var sourceTable string

		fkey.Table = tableName
		err = rows.Scan(&fkey.Name, &sourceTable, &fkey.Column, &fkey.ForeignTable, &fkey.ForeignColumn)
		if err != nil {
			return nil, err
		}

		fkeys = append(fkeys, fkey)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fkeys, nil
}

// TranslateColumnType converts mysql database types to Go types, for example
// "varchar" to "string" and "bigint" to "int64". It returns this parsed data
// as a Column object.
func (m *MySQLDriver) TranslateColumnType(c structs.Column) structs.Column {
	unsigned := strings.Contains(c.FullDBType, "unsigned")
	if c.Nullable {
		switch c.DBType {
		case "tinyint":
			// map tinyint(1) to bool if TinyintAsBool is true
			if !m.tinyIntAsInt && c.FullDBType == "tinyint(1)" {
				c.Type = "null.Bool"
			} else if unsigned {
				c.Type = "null.Uint8"
			} else {
				c.Type = "null.Int8"
			}
		case "smallint":
			if unsigned {
				c.Type = "null.Uint16"
			} else {
				c.Type = "null.Int16"
			}
		case "mediumint":
			if unsigned {
				c.Type = "null.Uint32"
			} else {
				c.Type = "null.Int32"
			}
		case "int", "integer":
			if unsigned {
				c.Type = "null.Uint"
			} else {
				c.Type = "null.Int"
			}
		case "bigint":
			if unsigned {
				c.Type = "null.Uint64"
			} else {
				c.Type = "null.Int64"
			}
		case "float":
			c.Type = "null.Float32"
		case "double", "double precision", "real":
			c.Type = "null.Float64"
		case "boolean", "bool":
			c.Type = "null.Bool"
		case "date", "datetime", "timestamp":
			c.Type = "null.Time"
		case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
			c.Type = "null.Bytes"
		case "numeric", "decimal", "dec", "fixed":
			c.Type = "types.NullDecimal"
		case "json":
			c.Type = "null.JSON"
		default:
			c.Type = "null.String"
		}
	} else {
		switch c.DBType {
		case "tinyint":
			// map tinyint(1) to bool if TinyintAsBool is true
			if !m.tinyIntAsInt && c.FullDBType == "tinyint(1)" {
				c.Type = "bool"
			} else if unsigned {
				c.Type = "uint8"
			} else {
				c.Type = "int8"
			}
		case "smallint":
			if unsigned {
				c.Type = "uint16"
			} else {
				c.Type = "int16"
			}
		case "mediumint":
			if unsigned {
				c.Type = "uint32"
			} else {
				c.Type = "int32"
			}
		case "int", "integer":
			if unsigned {
				c.Type = "uint"
			} else {
				c.Type = "int"
			}
		case "bigint":
			if unsigned {
				c.Type = "uint64"
			} else {
				c.Type = "int64"
			}
		case "float":
			c.Type = "float32"
		case "double", "double precision", "real":
			c.Type = "float64"
		case "boolean", "bool":
			c.Type = "bool"
		case "date", "datetime", "timestamp":
			c.Type = "time.Time"
		case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
			c.Type = "[]byte"
		case "numeric", "decimal", "dec", "fixed":
			c.Type = "types.Decimal"
		case "json":
			c.Type = "types.JSON"
		default:
			c.Type = "string"
		}
	}

	return c
}

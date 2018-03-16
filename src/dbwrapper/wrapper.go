package dbwrapper

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	Host       string
	User       string
	Password   string
	Dbname     string
	connection *sql.DB
}

//Connect to data base using User, Password, Dbname fields.
//Return error or nill if connection was successfull.
func (con *DBConnection) Connect() error {
	connpath := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", con.User, con.Password, con.Host, con.Dbname)
	connection, err := sql.Open("postgres", connpath)
	if err != nil {
		con.connection = nil
		return err
	}
	con.connection = connection
	return nil
}

//Select all fields data in specified table
func (conn *DBConnection) Select(table string, fields ...string) (*sql.Rows, error) {
	flds := ""
	if len(fields) > 0 {
		flds = strings.Join(fields, ", ")
	} else {
		flds = "*"
	}
	query := fmt.Sprintf("SELECT %s FROM %s;", flds, table)
	return conn.connection.Query(query)
}

//Select all fields data in specified table
func (conn *DBConnection) SelectBy(table string, keys map[string]interface{}, fields ...string) (*sql.Rows, error) {
	flds := ""
	filterExpression := joinKeys(keys)
	if len(filterExpression) > 0 {
		filterExpression = "WHERE " + filterExpression
	}
	if len(fields) > 0 {
		flds = strings.Join(fields, ", ")
	} else {
		flds = "*"
	}
	query := fmt.Sprintf("SELECT %s FROM %s %s;", flds, table, filterExpression)
	return conn.connection.Query(query)
}

//Update table with specified key=>value map
func (conn *DBConnection) Update(table string, keys map[string]interface{}, args map[string]interface{}) error {
	filterExpression := joinKeys(keys)
	if len(filterExpression) > 0 {
		filterExpression = "WHERE " + filterExpression
	}
	joinedArgs := joinArgs(args)
	query := fmt.Sprintf("UPDATE %s SET %s %s;", table, joinedArgs, filterExpression)
	fmt.Println(query)
	_, err := conn.connection.Exec(query)
	return err
}

func (conn *DBConnection) Insert(table string, data map[string]interface{}) error {
	fields := ""
	values := ""
	for key, val := range data {
		fields += key + ","
		values += fmt.Sprintf("'%s',", val)
	}
	fields = fields[:len(fields)-1]
	values = values[:len(values)-1]
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s);", table, fields, values)
	_, err := conn.connection.Exec(query)
	return err
}

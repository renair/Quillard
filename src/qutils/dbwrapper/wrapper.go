package dbwrapper

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	Host                string
	User                string
	Password            string
	Dbname              string
	connection          *sql.DB
	transactionRequired bool
	transactionDepth    int16
}

//Connect to data base using User, Password, Dbname fields.
//Return error or nill if connection was successfull.
func (conn *DBConnection) Connect() error {
	connpath := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conn.User, conn.Password, conn.Host, conn.Dbname)
	connection, err := sql.Open("postgres", connpath)
	if err != nil {
		conn.connection = nil
		return err
	}
	conn.connection = connection
	conn.transactionRequired = false
	conn.transactionDepth = 0
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
	_, err := conn.connection.Exec(query)
	return err
}

func (conn *DBConnection) Insert(table string, data map[string]interface{}) error {
	fields := ""
	values := ""
	for key, val := range data {
		fields += key + ","
		values += fmt.Sprintf("'%v',", val)
	}
	fields = fields[:len(fields)-1]
	values = values[:len(values)-1]
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s);", table, fields, values)
	_, err := conn.connection.Exec(query)
	return err
}

func (conn *DBConnection) MultipleInsert(table string, columns []string, data [][]interface{}) error {
	fields := strings.Join(columns, ",")
	valuesBuilder := strings.Builder{}
	for i, values := range data {
		valuesBuilder.WriteRune('(')
		for q := 0; q < len(values); q++ {
			stringifiedVal := fmt.Sprintf("'%v'", values[q])
			if q == len(values)-1 {
				valuesBuilder.WriteString(stringifiedVal)
			} else {
				valuesBuilder.WriteString(stringifiedVal + ",")
			}
		}
		if i == len(data)-1 {
			valuesBuilder.WriteRune(')')
		} else {
			valuesBuilder.WriteString("),")
		}
	}
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s;", table, fields, valuesBuilder.String())
	_, err := conn.connection.Exec(query)
	return err
}

func (conn *DBConnection) ManualQuery(query string) (*sql.Rows, error) {
	return conn.connection.Query(query)
}

func (conn *DBConnection) BeginTransaction() {
	conn.transactionDepth++
	if conn.transactionDepth == 1 {
		conn.ManualQuery("START TRANSACTION;")
		conn.transactionRequired = true
	}
}

func (conn *DBConnection) CommitTransaction() {
	conn.transactionDepth--
	if conn.transactionRequired && conn.transactionDepth == 0 {
		conn.ManualQuery("COMMIT TRANSACTION;")
		conn.transactionRequired = false
	}
}

func (conn *DBConnection) RollbackTransaction() {
	conn.transactionDepth--
	if conn.transactionRequired {
		conn.ManualQuery("ROLLBACK TRANSACTION;")
		conn.transactionRequired = false
	}
}

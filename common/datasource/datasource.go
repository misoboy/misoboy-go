package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
	"fmt"
)

func init(){
	/*info := ConnectionInfo{
		"host" : "localhost",
		"port" : 3306,
		"uname" : "root",
		"passwd" : "thd2243",
		"dbname" : "go",
	}*/
}

/*type DataSource interface {
	Select (query string, params ...interface{})

}*/

type DataSource interface {
	SelectQuery(query string, params ...interface{}) []map[string]interface{}
	SelectOneQuery(query string, params ...interface{}) map[string]interface{}
	UpdateQuery(query string, params ...interface{}) int64
	DeleteQuery(query string, params ...interface{}) int64
}

type dataSource struct {
	db *sql.DB
}

type Any interface{}
type ConnectionInfo map[string] Any

func NewDataSource() DataSource {
	return &dataSource{ db : getOpenConnection() }
}

func getOpenConnection () *sql.DB {
	db, err := sql.Open("mysql", "go:go@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

func (r * dataSource) SelectQuery(query string, params ...interface{}) []map[string]interface{} {
	defer r.db.Close()
	var tempRows *sql.Rows

	if params != nil && len(params) > 0 {
		rows, err := r.db.Query(query, params...)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		tempRows = rows
	} else {
		rows, _ := r.db.Query(query)
		tempRows = rows
	}
	columns, err := tempRows.Columns()
	if err != nil {
		panic(err.Error())
	}

	defer tempRows.Close()

	dataMap := make([]map[string]interface{}, 0)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for tempRows.Next() {
		tempRows.Scan(scanArgs...)

		var value string
		var rowDataMap = make(map[string]interface{})
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			rowDataMap[columns[i]] = value
		}
		dataMap = append(dataMap, rowDataMap)
	}
	return dataMap
}

func (r * dataSource) SelectOneQuery(query string, params ...interface{}) map[string]interface{} {
	defer r.db.Close()
	var tempRows *sql.Rows

	if params != nil && len(params) > 0 {
		rows, err := r.db.Query(query, params...)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		tempRows = rows
	} else {
		rows, _ := r.db.Query(query)
		tempRows = rows
	}
	columns, err := tempRows.Columns()
	if err != nil {
		panic(err.Error())
	}

	defer tempRows.Close()

	dataMap := make(map[string]interface{}, 0)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if tempRows.Next() {
		tempRows.Scan(scanArgs...)

		var value string
		var rowDataMap = make(map[string]interface{})
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			rowDataMap[columns[i]] = value
		}
		dataMap = rowDataMap
	}

	return dataMap
}

func (r * dataSource) UpdateQuery(query string, params ...interface{}) int64 {
	defer r.db.Close()
	result, err := r.db.Exec(query, params...)
	if err != nil {
		panic(err.Error())
		return 0
	}
	num, _ := result.RowsAffected()

	return num
}

func (r * dataSource) DeleteQuery(query string, params ...interface{}) int64 {
	defer r.db.Close()
	result, err := r.db.Exec(query, params...)
	if err != nil {
		panic(err.Error())
		return 0
	}
	num, _ := result.RowsAffected()

	return num
}
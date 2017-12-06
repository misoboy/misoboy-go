package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
	"fmt"
	"regexp"
	"reflect"
	"strings"
	"misoboy_web/common/pagination"
	"strconv"
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
	SelectQuery(query string, param interface{}) []map[string]interface{}
	SelectOneQuery(query string, param interface{}) map[string]interface{}
	UpdateQuery(query string, param interface{}) int64
	DeleteQuery(query string, param interface{}) int64
}

type dataSource struct {
	db *sql.DB
}

type Any interface{}
type ConnectionInfo map[string] Any

func NewDataSource() DataSource {
	return &dataSource{ }
}

func getOpenConnection () *sql.DB {
	db, err := sql.Open("mysql", "go:go@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

func (r * dataSource) SelectQuery(query string, param interface{}) []map[string]interface{} {
	r.db = getOpenConnection ()

	defer r.db.Close()
	var tempRows *sql.Rows

	if param != nil {
		query = queryBindParameters(query, param)
		query = makePagination(query, param)

		rows, err := r.db.Query(query)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		tempRows = rows
	} else {
		rows, _ := r.db.Query(query)
		tempRows = rows
	}

	fmt.Println(query)

	dataMap := make([]map[string]interface{}, 0)
	if tempRows != nil {
		columns, err := tempRows.Columns()
		if err != nil {
			panic(err.Error())
		}

		defer tempRows.Close()

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
	}

	return dataMap
}

func (r * dataSource) SelectOneQuery(query string, param interface{}) map[string]interface{} {
	r.db = getOpenConnection ()

	defer r.db.Close()
	var tempRows *sql.Rows

	if param != nil {
		query = queryBindParameters(query, param)
		rows, err := r.db.Query(query)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		tempRows = rows
	} else {
		rows, _ := r.db.Query(query)
		tempRows = rows
	}

	fmt.Println(query)

	dataMap := make(map[string]interface{}, 0)
	if tempRows != nil {
		columns, err := tempRows.Columns()
		if err != nil {
			panic(err.Error())
		}

		defer tempRows.Close()

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
	}

	return dataMap
}

func (r * dataSource) UpdateQuery(query string, param interface{}) int64 {
	r.db = getOpenConnection ()

	defer r.db.Close()
	query = queryBindParameters(query, param)
	result, err := r.db.Exec(query)
	fmt.Println(query)
	if err != nil {
		panic(err.Error())
		return 0
	}
	num, _ := result.RowsAffected()

	return num
}

func (r * dataSource) DeleteQuery(query string, param interface{}) int64 {
	r.db = getOpenConnection ()

	defer r.db.Close()
	query = queryBindParameters(query, param)
	result, err := r.db.Exec(query)

	fmt.Println(query)
	if err != nil {
		panic(err.Error())
		return 0
	}
	num, _ := result.RowsAffected()

	return num
}


func makePagination(query string, param interface {}) string {
	if param != nil {
		_, ptr := inspectStructV(reflect.ValueOf(param), "Pagination")
		if ptr != nil {
			paginationVo := ptr.(pagination.Pagination)

			if paginationVo.PaginationEnable == pagination.PAGINE_ENABLE_ON {
				pageIndex := paginationVo.CurrentPageNo
				recordCountPerPage := paginationVo.RecordCountPerPageNo
				condOrder := paginationVo.CondOrder
				condAlign := paginationVo.CondAlign

				paginationConfig := pagination.Pagination{CurrentPageNo: pageIndex, RecordCountPerPageNo: recordCountPerPage}

				prefix := " SELECT * FROM (SELECT RESULT_LIST.*, @NO := @NO + 1 AS RNUM FROM( "
				suffix := fmt.Sprintf(" ) RESULT_LIST, ( SELECT @NO := 0 ) RESULT_NO ) RESULT WHERE RESULT.RNUM <= %s + %s AND RESULT.RNUM > %s", paginationConfig.FirstRecordIndex(), paginationConfig.RecordCountPerPageNo, paginationConfig.FirstRecordIndex())
				order := fmt.Sprintf(" ORDER BY %s %s", condOrder, condAlign)

				if len(condOrder) > 0 && len(condAlign) > 0 {
					return prefix + query + suffix + order
				} else {
					return prefix + query + suffix
				}
			}
		}
	}

	return query
}


func queryBindParameters(query string, param interface{}) string {

	regexStr := regexp.MustCompile("(\\#|\\$){.*?}")
	regexVar := regexp.MustCompile("([^(\\#|\\$)\\{]).*[^}]")
	findStrs := regexStr.FindAllString(query, -1)
	for i, v := range findStrs {
		fieldName := strings.Title(regexVar.FindString(v))
		isQuot := strings.HasPrefix(v, "#")
		resultParam, _ := inspectStructV(reflect.ValueOf(param), fieldName)

		if isQuot {
			// parameter : ''
			resultParam = "'" + resultParam + "'"
		} else {
			// parameter : non ''
		}

		query = strings.Replace(query, v, resultParam, i + 1)
	}

	return query
}

func inspectStructV(val reflect.Value, fieldName string) (string, interface {}) {

	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//var address uintptr

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()

		}
		if valueField.CanAddr() {
			//address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
			//address = valueField.Addr().Pointer()
		}

		if !valueField.IsValid() {
			continue
		}

		/*fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n", typeField.Name,
			valueField.Interface(), address, typeField.Type, valueField.Kind())*/

		if fieldName == typeField.Name {

			var v string
			switch valueField.Kind() {
				case reflect.String:
					v = valueField.String()
				case reflect.Int:
					v = strconv.FormatInt(valueField.Int(), 10)
				case reflect.Int32:
					v = strconv.FormatInt(valueField.Int(), 10)
				case reflect.Int64:
					v = strconv.FormatInt(valueField.Int(), 10)
				default:
					fmt.Println("Not support type of struct")
					v = valueField.String()
			}

			return v, valueField.Interface()
		}

		if valueField.Kind() == reflect.Struct {
			value, ptr := inspectStructV(valueField, fieldName)
			if value != "" {
				return value, ptr
			}
		}
	}

	return "", nil
}
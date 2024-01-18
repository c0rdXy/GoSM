package output

import "fmt"

func GoMSInit() {
	fmt.Println("  ________           _________   _____   ")
	fmt.Println(" /  _____/   ____   /   _____/  /     \\  ")
	fmt.Println("/   \\  ___  /  _ \\  \\_____  \\  /  \\ /  \\ ")
	fmt.Println("\\    \\_\\  \\(  <_> ) /        \\/    Y    \\")
	fmt.Println(" \\______  / \\____/ /_______  /\\____|__  /")
	fmt.Println("        \\/                 \\/         \\/ ")

	fmt.Println("GoSM - 一个简单的Go语言SQL注入工具")
	fmt.Println("作者: xxx")
}

// PrintCurrentDatabase 输出当前数据库信息
func PrintCurrentDatabase(database string) {
	fmt.Printf("[*] Current Database: %s\n", database)
}

// PrintDatabase 输出数据库信息
func PrintDatabase(database string) {
	fmt.Printf("\t[*] Database: %s\n", database)
}

// PrintTable 输出表信息
func PrintTable(table string) {
	fmt.Printf("\t\t[*] Table: %s\n", table)
}

// PrintColumn 输出字段信息
func PrintColumn(columns string) {
	fmt.Printf("\t\t\t[+] Columns: %v\n", columns)
}

// PrintCurrentDatabases 输出当前数据库信息
func PrintCurrentDatabases(database []string) {
	fmt.Printf("[*] Current Database: %s\n", database)
}

// PrintDatabase 输出数据库信息
func PrintDatabases(database []string) {
	fmt.Printf("\t[*] Database: %s\n", database)
}

// PrintTable 输出表信息
func PrintTables(table []string) {
	fmt.Printf("\t\t[*] Tables: %s\n", table)
}

// PrintColumn 输出字段信息
func PrintColumns(columns []string) {
	fmt.Printf("\t\t\t[+] Columns: %v\n", columns)
}

// PrintColumnValues 输出字段值信息
func PrintColumnValues(column string, value []string, index int) {
	fmt.Printf("\t\t\t\t[>] Column: %s\n", column)
	fmt.Printf("\t\t\t\t\t[%d] %s\n", index+1, value)
}

// NewLine 打印空行
func NewLine() {
	fmt.Println()
}

// PrintCurrentDatabaseArray 输出当前数据库信息
func PrintCurrentDatabaseArray(databases []string) {
	for _, database := range databases {
		fmt.Printf("[*] Current Database: %s\n", database)
	}
}

// PrintDatabaseArray 输出数据库信息
func PrintDatabaseArray(databases []string) {
	for _, database := range databases {
		fmt.Printf("\t[*] Database: %s\n", database)
	}
}

// PrintTableArray 输出表信息
func PrintTableArray(tables []string) {
	for _, table := range tables {
		fmt.Printf("\t\t[*] Table: %s\n", table)
	}
}

// PrintColumnsArray 输出字段信息
func PrintColumnsArray(columns []string) {
	for _, column := range columns {
		fmt.Printf("\t\t\t[+] Column: %s\n", column)
	}
}

// PrintColumnValuesArray 输出字段值信息
func PrintColumnValuesArray(column string, values []string) {
	fmt.Printf("\t\t\t\t[>] Column: %s\n", column)
	for i, value := range values {
		fmt.Printf("\t\t\t\t\t[%d] %s\n", i+1, value)
	}
}

// 输出URL和payload
func PrintURLAndPayload(url string, payload string) {
	fmt.Printf("[+] URL: %s\n", url)
	fmt.Printf("[+] Payload: %s\n", payload)
}

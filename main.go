package main

import (
	"GoSM/bool-blind2"
	"GoSM/error"
	"GoSM/myUtils/output"
	"flag"
	"fmt"
	"log"
)

var (
	TargetURL string
	Method    string
	Database  string
	Table     string
	Column    string
	Dump      bool
)

func init() {
	output.GoMSInit()

	flag.StringVar(&TargetURL, "url", "", "目标URL")
	flag.StringVar(&Method, "method", "", "注入方法（目前支持：E 报错注入，B 布尔盲注）")
	flag.StringVar(&Database, "database", "", "数据库名")
	flag.StringVar(&Table, "table", "", "表名")
	flag.StringVar(&Column, "column", "", "字段名")
	flag.BoolVar(&Dump, "dump", false, "自动注入得到所有信息")

	flag.Parse()

	// 验证必需参数
	if TargetURL == "" {
		log.Fatal("目标URL是必需的。用法: go run main.go -url \"http://example.com\" -method E -database shop -table users -column username")
	}

	// 如果输入 -h，则显示帮助信息
	if flag.Lookup("h") != nil {
		flag.Usage()
	}
}

func main() {
	// 执行SQL注入操作
	performInjection()
}

// 执行SQL注入操作
func performInjection() {
	// 在这里添加你的注入逻辑
	// 根据 method、database、table、column 等参数执行相应的注入操作
	// 输出结果使用 output 包中的函数

	switch Method {
	case "U":
		performUnionBasedInjection()
	case "B":
		performBlindInjection()
	case "E":
		performErrorBasedInjection()
	// 可以根据需要添加其他注入方式的处理逻辑
	default:
		log.Fatalf("不支持的注入方式: %s", Method)
	}

}

// 执行基于Union的注入操作
func performUnionBasedInjection() {
	// 在这里添加执行Union-Based注入的具体逻辑
	// 使用相应的库或工具执行 SQL 查询，或者调用相应的 ORM 方法
	// 最后将结果输出到文件或打印到控制台，根据需要进行后续处理
	fmt.Println("执行基于Union的注入操作...")
}

// 执行盲注（二分法）的注入操作
func performBlindInjection() {
	fmt.Println("执行盲注（二分法）的注入操作...")

	if Dump {
		// 输出所有信息
		bool_blind2.GetAllInfo(TargetURL)
	} else {
		// 根据命令行参数输出相应的内容
		if Database == "" && Table == "" && Column == "" {
			// 输出所有数据库名
			bool_blind2.GetCurrentDatabase(TargetURL)
			bool_blind2.GetAllDatabases(TargetURL)
		} else if Database != "" && Table == "" && Column == "" {
			output.PrintCurrentDatabase(Database)
			// 输出某个数据库的所有表名
			bool_blind2.GetTablesInDatabase(TargetURL, Database)
		} else if Database != "" && Table != "" && Column == "" {
			output.PrintCurrentDatabase(Database)
			output.PrintTable(Table)
			// 输出某个表的所有字段名
			bool_blind2.GetColumnsInTable(TargetURL, Database, Table)
		} else if Database != "" && Table != "" && Column != "" {
			output.PrintCurrentDatabase(Database)
			output.PrintTable(Table)
			output.PrintColumn(Column)
			// 输出某个字段的所有值
			bool_blind2.GetColumnsAndValues(TargetURL, Database, Table, Column)
		}
	}
}

// 执行报错注入操作
func performErrorBasedInjection() {
	fmt.Println("执行报错注入操作...")

	if Dump {
		// 输出所有信息
		error.GetAllInfo(TargetURL)
	} else {
		// 根据命令行参数输出相应的内容
		if Database == "" && Table == "" && Column == "" {
			// 输出所有数据库名
			error.GetCurrentDatabase(TargetURL)
			error.GetAllDatabases(TargetURL)
		} else if Database != "" && Table == "" && Column == "" {
			// 输出某个数据库的所有表名
			error.GetTablesInDatabase(TargetURL, Database)
		} else if Database != "" && Table != "" && Column == "" {
			// 输出某个表的所有字段名
			error.GetColumnsAndValuesInTable(TargetURL, Table)
		} else if Database != "" && Table != "" && Column != "" {
			// 输出某个字段的所有值
			error.GetColumnsAndValuesInTable(TargetURL, Column)
		}
	}

}

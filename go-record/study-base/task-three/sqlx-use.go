package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	//init database
	dsn := "user:pass@tcp(127.0.0.1:3306)/company?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("connect db:", err)
	}
	defer db.Close()
}

type Employee struct {
	ID         int64   `db:"id"`
	Name       string  `db:"Name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

/*
使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 Name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

func queryData() {
	sqlStr := "select id,Name,department,salary from employees where department = ?"
	var employees = make([]Employee, 0)
	err := db.Select(&employees, sqlStr, "技术部")
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	fmt.Printf("employees = %v\n", employees)

	var maxSalary Employee
	sqlStr = "select id,Name,department,salary from employees order by salary desc limit 1"
	err = db.Get(&maxSalary, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("maxSalary = %v\n", maxSalary.Salary)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func queryData2() {
	sqlStr := "select id,Name,department,salary from books where price > ?"
	var books = make([]Book, 0)
	err := db.Select(&books, sqlStr, 50)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	fmt.Printf("books = %v\n", books)
}

func main() {
	queryData()
	queryData2()
}

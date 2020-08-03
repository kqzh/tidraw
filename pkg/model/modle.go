package model

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	PictureHeight      = 128
	DropTableSql       = "drop table pix%02x;"
	SelectTableSql     = "select * from pix%02x limit %v;"
	CreateTableSql     = `CREATE TABLE pix%02x (id INT NOT NULL AUTO_INCREMENT, pix_data VARCHAR(45) DEFAULT "hello, pingCap", PRIMARY KEY (id));`
	InsertTableDataSql = "insert into pix%02x values()"
)

var (
	DB *sql.DB
	wg sync.WaitGroup
)

func InitDB(host, port string) (err error) {
	dsn := fmt.Sprintf("root:@tcp(%v:%v)/test", host, port)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect tidb :%v", err)
	}
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(50)
	begin := time.Now()
	initTables()
	log.Println("Init Tables Run time: ", time.Since(begin))
	return
}

func initTables() {
	stmt := ""
	for i := 0; i < PictureHeight; i++ {
		stmt += fmt.Sprintf(CreateTableSql, i)
	}
	_, err := DB.Exec(stmt)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < PictureHeight; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			stmt := fmt.Sprintf(InsertTableDataSql, index)
			for j := 0; j < 256; j++ {
				stmt += ",()"
			}
			_, err := DB.Exec(stmt)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(i)
	}
	wg.Wait()
}

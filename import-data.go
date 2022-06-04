package main

import (
	"fmt"
	"sync"

	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vtuber struct {
	gorm.Model
	Id   string
	Name string
}

func importData(w *sync.WaitGroup, db *gorm.DB, data Vtuber) {
	defer w.Done()
	db.Create(&data)
}

func main() {

	var wg sync.WaitGroup

	dsn := "root:root@tcp(127.0.0.1:3306)/goimport?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	// migrate
	db.AutoMigrate(&Vtuber{})

	// read file
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		wg.Add(1)
		go importData(&wg, db, Vtuber{Id: row[0], Name: row[1]})
	}

	wg.Wait()
	fmt.Println("import done")

}

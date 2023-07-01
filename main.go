package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Start    time.Time
	Stop     time.Time
	Duration int
	Project  string
}

func main() {
	//connect database
	dsn := "root@tcp(127.0.0.1:3306)/timemanagementapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	if db != nil {
		fmt.Println("a ...any")
	} else {
		fmt.Println(err)
	}
	// create record
	record1 := Record{Start: time.Now(), Stop: time.Now(), Duration: 10, Project: "meinv"}
	result := db.Create(&record1)
	fmt.Println(result.RowsAffected)

	// read record
	var record Record
	db.First(&record)
	//In the context of gorm.DB, the First() method does not directly return the retrieved record. Instead, it modifies the provided record2 variable and returns the gorm.DB object itself.
	//So when you write result2 := db.First(&record2), the result2 variable is of type *gorm.DB, not the record2 variable. That's why result2.ID is undefined because the *gorm.DB type does not have an ID field.
	//To access the retrieved record's fields, you need to use the record2 variable, which was passed as a reference to the First() method. This method modifies the record2 variable with the fetched record.

	fmt.Println("ID:", record.ID)
	fmt.Println("Create Time:", record.CreatedAt)
	fmt.Println("Project:", record.Project)

}

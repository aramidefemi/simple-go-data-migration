package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "log"
    "os" 
	// "strings"
	// "reflect"

    _ "github.com/lib/pq"
)

func main() {
    // PostgreSQL connection URL
    psqlURL := "postgres://your_postgres_user:your_postgres_password@your_postgres_host:5432/your_postgres_db?sslmode=disable"

	type data_row struct {
		email string
		name string
		profession string
		country string 
	}
	
	dataRows := []data_row{}

	emailsStore := make(map[string]int)

    // Connect to PostgreSQL
    db, err := sql.Open("postgres", psqlURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Open CSV file
    file, err := os.Open("data.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse CSV file
    csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    // Data clean up and processing 
    for _, record := range records {
		// data clean up
		 
		
		var newRows data_row;

		newRows.email = record[4];
		newRows.name = record[2] + " " + record[3];
		newRows.profession= record[5];
		newRows.country= record[6]; 

		if(emailsStore[newRows.email] == 1){ // avoid dubs from csv files
			continue;
		}
 		
		var count int
        err := db.QueryRow("SELECT COUNT(*) FROM wait_list WHERE email = $1", record[4]).Scan(&count)
        if err != nil {
            log.Fatal(err)
        }


				fmt.Printf("count: %v  ", count)
				if(count == 0){
				emailsStore[newRows.email] = 1
				dataRows = append(dataRows, newRows)
				}
				
    }
	
    // Iterate over CSV records and insert into PostgreSQL
    for index, row := range dataRows {
        // Customize this part to map CSV columns to PostgreSQL table columns
		fmt.Printf("\n========== inserting : %v , index : %v total jobs: %v \n\n", row.email, index, len(dataRows))
        _, err := db.Exec("INSERT INTO wait_list (email, name, profession, country) VALUES ($1, $2, $3, $4)", row.email, row.name, row.profession, row.country)
        if err != nil {
            log.Fatal(err)
        }
    }


	fmt.Printf("\nData migration complete total jobs: %v of %v \n\n",   len(dataRows),  len(records))
 
}

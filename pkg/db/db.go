package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"zzucturum-app/pkg/models"
)

var log = logrus.New()

func createConnection() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("DB_PARAMS"))
	fmt.Printf(os.Getenv("DB_PARAMS"))
	if err != nil {
		log.Fatalf("Can't open connection. %v", err)
		panic(err)
	}

	if err != nil {
		log.Fatalf("No connection to posgres server select! %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("No connection to posgres server! %v", err)
		panic(err)
	}

	log.Info("Successfully connected!")
	return db
}

func InsertCounter(counter models.Counter) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO social_media (url, number_requests, created_at) VALUES 
                    ($1, $2, to_timestamp($3)) RETURNING id`

	var id int64
	err := db.QueryRow(sqlStatement, counter.Domain, counter.NumberOfRequests, counter.Timestamp).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Info("Inserted a single record ", id)

	return id
}

func GetCounterStatsByMinutes() ([]models.Stats, error) {

	db := createConnection()

	defer db.Close()
	var stats []models.Stats

	sqlStatement := `SELECT domain_name, sum(number_requests) as nor
						FROM counter_data
						WHERE  created_at >= date_trunc('minute', NOW() - INTERVAL '1 minute')
						group by domain_name 
						order by nor desc 
						limit 10`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stat models.Stats

		err = rows.Scan(&stat.Domain, &stat.NumberOfRequests)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		stats = append(stats, stat)

	}
	return stats, err
}

func GetCounterStatsByHour() ([]models.Stats, error) {
	db := createConnection()

	defer db.Close()

	var stats []models.Stats

	sqlStatement := `SELECT domain_name, sum(number_requests) as nor
						FROM counter_data
						WHERE  created_at >= date_trunc('minute', NOW() - INTERVAL '1 hour')
						group by domain_name 
						order by nor desc 
						limit 10`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stat models.Stats

		err = rows.Scan(&stat.Domain, &stat.NumberOfRequests)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		stats = append(stats, stat)

	}
	return stats, err
}

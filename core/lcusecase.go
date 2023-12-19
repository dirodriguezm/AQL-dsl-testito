package core

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Detection struct {
	Candid int64   `json:"candid"`
	Oid    string  `json:"oid"`
	Mjd    float32 `json:"mjd"`
	Mag    float32 `json:"mag"`
	Fid    int8    `json:"fid"`
}

type NonDetection struct {
	Aid string `json:"aid"`
}

type LightcurveUseCase struct{}

func getDbConn() *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func makeDetectionQuery(queryArgs *QueryArguments) string {
	query := ""
	if queryArgs.Oid != "" {
		query = "SELECT candid, oid, mjd, magpsf, fid from detection where oid=$1"
	}
	return query
}

func (uc *LightcurveUseCase) GetDetection(queryArgs *QueryArguments) []Detection {
	sqlQuery := makeDetectionQuery(queryArgs)
	conn := getDbConn()
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), sqlQuery, queryArgs.Oid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var detections []Detection
	for rows.Next() {
		var detection Detection
		err := rows.Scan(&detection.Candid, &detection.Oid, &detection.Mjd, &detection.Mag, &detection.Fid)
		if err != nil {
			panic(err)
		}
		detections = append(detections, detection)
	}
	return detections
}

func (uc *LightcurveUseCase) GetNonDetection(queryArgs *QueryArguments) []NonDetection {
	sqlQuery := makeDetectionQuery(queryArgs)
	conn := getDbConn()
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), sqlQuery, queryArgs.Oid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var nondetections []NonDetection
	for rows.Next() {
		var nondetection NonDetection
		err := rows.Scan(&nondetection.Aid)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", nondetection)
		nondetections = append(nondetections, nondetection)
	}
	return nondetections
}

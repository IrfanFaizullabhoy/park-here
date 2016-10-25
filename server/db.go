package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToPG(dbName string) *sql.DB {
	db, err := sql.Open("postgres", "postgres://finch:password@"+os.Getenv("DB_PORT_5432_TCP_ADDR")+"/nest?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func SetupTable(db *sql.DB, tableName string) {
	var tables = InitTables()
	for _, element := range tables {
		_, err := db.Exec(element)
		if err != nil {
			panic(err)
			return
		}
	}
}

func GetListingsDB(db *sql.DB) []Listing {
	var listings []Listing
	rows, err := db.Query("SELECT * FROM listings")
	check(err)
	for rows.Next() {
		var ListingID int
		var Address string
		var ListingName string
		var CompanyID int
		var CompanyName string
		var Type int
		var Price int
		var SquareFeet int
		var NumBed int
		var NumBath float32
		var MinsWalk float32
		var Description string
		var CoverURL string
		var ThumbnailURL string
		err = rows.Scan(&ListingID, &Address, &ListingName, &CompanyID, &CompanyName, &Type, &Price, &SquareFeet, &NumBed, &NumBath, &MinsWalk, &Description, &CoverURL, &ThumbnailURL)
		check(err)
		listing := Listing{
			ListingID:    ListingID,
			Address:      Address,
			ListingName:  ListingName,
			CompanyID:    CompanyID,
			CompanyName:  CompanyName,
			Type:         Type,
			Price:        Price,
			SquareFeet:   SquareFeet,
			NumBed:       NumBed,
			NumBath:      NumBath,
			MinsWalk:     MinsWalk,
			Description:  Description,
			CoverURL:     CoverURL,
			ThumbnailURL: ThumbnailURL}

		listings = append(listings, listing)
	}
	return listings
}

func RegisterListingToDb(db *sql.DB, listing *Listing) {
	_, err := db.Exec(" INSERT INTO listings (listing_id, address, listing_name , company_id , company_name , type , price , square_feet , num_bed, num_bath, mins_walk, description, cover_url, thumbnail_url ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		listing.ListingID, listing.Address, listing.ListingName, listing.CompanyID, listing.CompanyName, listing.Type, listing.Price, listing.SquareFeet, listing.NumBed, listing.NumBath, listing.MinsWalk, listing.Description, listing.CoverURL, listing.ThumbnailURL)
	check(err)
}

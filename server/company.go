package main

type Company struct {
	Name     string    `json"company_name"`
	Rating   float32   `json"company_rating"`
	Listings []Listing `json"company_listings"`
}

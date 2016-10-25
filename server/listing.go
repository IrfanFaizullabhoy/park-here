package main

type Listing struct {
	ListingID    int     `json:"listing_id,string"`
	Address      string  `json:"address"`
	ListingName  string  `json:"listing_name"`
	CompanyID    int     `json:"company_id,string"`
	CompanyName  string  `json:"company_name"`
	Type         int     `json:"type,string"`
	Price        int     `json:"price,string"`
	SquareFeet   int     `json:"square_feet,string"`
	NumBed       int     `json:"num_bed,string"`
	NumBath      float32 `json:"num_bath,string"`
	MinsWalk     float32 `json:"mins_walk,string"`
	Description  string  `json:"description"`
	CoverURL     string  `json:"cover_url"`
	ThumbnailURL string  `json:"thumbnail_url"`
}

type Listings struct {
	Listings []Listing `json:"listings"`
}

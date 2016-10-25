/* Schema.go
 *
 */
package main

import ()

/*
 ***********************
 *                     *
 *       Tables        *
 *                     *
 ***********************
 */

/* Table: Listings
 *
 *	Columns:
 *
 * listing_id 		(int)
 *	address			  (text)
 *	listing_name	(text)
 *	company_id		(text)
 *	company_name	(text)
 *	type			    (int)
 *	price			    (int)
 * square_feet		(int)
 *	num_bed			  (int)
 * num_bath		  (float32)
 * mins_walk		 (decimal)
 */

var tables = make([]string, 10)

func InitTables() []string {
	tables = append(tables, "CREATE TABLE listings (listing_id integer PRIMARY KEY, address text, listing_name text, company_id text, company_name text, type integer, price integer, square_feet integer, num_bed integer, num_bath real, mins_walk decimal, description text, cover_url text, thumbnail_url text)")

	return tables
}

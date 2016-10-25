// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "database/sql"
)

var dat map[string]interface{}  
var listings Listings

func ReadSeed(db *sql.DB) {
    file, err := ioutil.ReadFile("data.txt"); 
    check(err)
    fmt.Println("In here")

    err = json.Unmarshal(file, &listings);
    check(err)

    for _, listing := range listings.Listings {
        RegisterListingToDb(db, &listing)
    }
}

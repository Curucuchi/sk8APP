package dbb

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Skater struct {
	SkaterID  int    `db:"skater_id"`
	Firstname string `db:"first_name"`
	Lastname  string `db:"last_name"`
	Stance    string `db:"stance"`
}

type Tricks struct {
	ID              int    `db:"trick_id"`
	Trickname       string `db:"trick_name"`
	Trickdifficulty int    `db:"trick_difficulty"`
}

type Parks struct {
	ID          int    `db:"park_id"`
	Parkname    string `db:"park_name"`
	Parkaddress string `db:"park_addr"`
}

const DSN = "root:@tcp(localhost:3636)/sk8opia"

func getSkaters() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println("OPEN error")
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM Skater")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var s Skater
		err := results.Scan(&s.SkaterID, &s.Firstname, &s.Lastname, &s.Stance)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Skater name:", s.Firstname, s.Lastname, "\n", "Stance:", s.Stance)
	}
}

func GetTricks() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT COUNT(trick_id) FROM Tricks")
	if err != nil {
		panic(err)
	}

	for results.Next() {
		var tID Tricks
		err := results.Scan(&tID.ID)
		if err != nil {
			panic(err.Error())
		}

		// time.now function helps us get a new random number each time, without it, we will get same number everytime
		rand.Seed(time.Now().UnixNano())
		randTrick := rand.Intn(tID.ID)

		Trick, err := db.Query(`SELECT trick_name, trick_difficulty FROM Tricks WHERE trick_id =` + strconv.Itoa(randTrick))
		if err != nil {
			panic(err)
		}
		for Trick.Next() {
			var t Tricks
			err := Trick.Scan(&t.Trickname, &t.Trickdifficulty)
			if err != nil {
				panic(err)
			}
			fmt.Println("Trick name:", t.Trickname, "\n", "Trick difficulty:", t.Trickdifficulty, "\n")

		}

	}

}

func GetParks() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	parks, err := db.Query(`SELECT COUNT(park_id) FROM Parks`)
	if err != nil {
		panic(err.Error())
	}

	for parks.Next() {
		var pID Parks
		err := parks.Scan(&pID.ID)
		if err != nil {
			panic(err.Error())
		}

		rand.Seed(time.Now().UnixNano())
		randPark := rand.Intn(pID.ID)

		Park, err := db.Query(`SELECT park_name, park_addr FROM Parks WHERE park_id =` + strconv.Itoa(randPark))
		for Park.Next() {
			var p Parks
			err := Park.Scan(&p.Parkname, &p.Parkaddress)
			if err != nil {
				panic(err.Error)
			}
			fmt.Println("Park name:", p.Parkname, "\n", "Park address:", p.Parkaddress, "\n")
		}
	}

}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	dbURL := "postgres://postgres:test1234@localhost:5432/seminar?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can not connect to database;", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Create **just one** simple API to get the the list of visitor list base on/filter by seminar event.
	app.Get("/visitor", func(c *fiber.Ctx) error {
		eventID := c.Query("event_id")
		eventName := c.Query("event_name")

		type Res struct {
			EventID      int    `json:"event_id"`
			EventName    string `json:"event_name"`
			VisitorName  string `json:"visitor_name"`
			VisitorEmail string `json:"visitor_email"`
			InviteCode   string `json:"invite_code"`
			IsAccept     bool   `json:"is_accept"`
		}

		var (
			res      []Res
			err      error
			queryStr string
			rows     *sql.Rows
		)
		queryStr = `select e.id, e.name, v.name, v.email, v.invite_code, v.is_accept 
					from visitor v
					left join events e on v.event_id = e.id`

		if eventID != "" {
			log.Println("event_id: ", eventID)
			eID, err := strconv.Atoi(eventID)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return c.JSON(fiber.Map{
					"message": err,
				})
			}
			queryStr = queryStr + ` where e.id = $1`
			rows, err = db.Query(queryStr, eID)
		} else if eventName != "" {
			log.Println("event_name: ", eventName)
			queryStr = queryStr + ` where e.name = $1`
			rows, err = db.Query(queryStr, eventName)
		} else {
			rows, err = db.Query(queryStr)
		}

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		defer rows.Close()

		for rows.Next() {
			var r Res
			err = rows.Scan(&r.EventID, &r.EventName, &r.VisitorName, &r.VisitorEmail, &r.InviteCode, &r.IsAccept)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": err,
				})
			}
			res = append(res, r)
		}

		return c.JSON(res)
	})

	app.Listen(":3000")
}

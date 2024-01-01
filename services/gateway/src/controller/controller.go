package controller

import (
	"encoding/json"
	"io"
	"os"

	// "fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type message struct {
	Values []string `json:"values"`
}

func Login(c *fiber.Ctx) error {
	auth := c.Get("Authorization")

	if auth == "" || auth[:5] != "Basic " {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing Auth",
		})
	}

	encodedCreds := auth[6:]

	url := os.Getenv("AUTH_SERVICE_URL")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	// send the encodedCreds( base64 ) as Authorization header

	client := &http.Client{}
	req.Header.Set("Authorization", encodedCreds)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	token := string(body)

	return c.Status(fiber.StatusOK).SendString(token)
}

// LEARN
func login0(w http.ResponseWriter, r *http.Request) {
	// parse the request
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := r.FormValue("data")
	verde := r.FormValue("verde")
	if data == "" {
		http.Error(w, "Not Data Provided", http.StatusBadRequest)
		return
	}

	resonse := message{Values: []string{data, verde}}
	responseJSON, _ := json.Marshal(resonse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

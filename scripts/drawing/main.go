package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
  "time"
	"math/rand"
	"net/http"

	"github.com/cpietsch/c3place/backend/pixel"
)

func main() {
	server := "http://localhost:4000"
	noiseLine(server, 255, 0, 0, 0, 100)
	noiseLine(server, 0, 255, 0, 0, 200)
	noiseLine(server, 0, 0, 255, 0, 300)
}

func PostPixel(server string, body pixel.PostPixel) error {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", server+"/pixel", bytes.NewBuffer(bodyJSON))
	req.Header.Set("UserAgent", "pixel-drawer")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(respBody))

	return nil
}

func noiseLine(server string, r, g, b uint8, x, y int) {
	rand.Seed(1337)

	for i := 0; i < 1000; i += 3 {
		random := rand.Intn(10)
		body := pixel.PostPixel{r, g, b, x + i, y + random}
		err := PostPixel(server, body)
		if err != nil {
			log.Println("ERROR", err)
		}
    time.Sleep(100*time.Millisecond)
	}
}

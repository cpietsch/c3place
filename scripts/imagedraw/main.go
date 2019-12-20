package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cpietsch/c3place/backend/pixel"
)

func main() {
	inputFilepath := flag.String("i", "", "input image")
	apiServer := flag.String("a", "http://localhost:4000", "api server")
	xPos := flag.Int("x", 0, "x pos")
	yPos := flag.Int("y", 0, "y pos")
	sleep := flag.Int("s", 100, "sleep between the api requests")
	flag.Parse()

	// load image
	log.Println("read image", *inputFilepath)
	inputData, err := os.Open(*inputFilepath)
	if err != nil {
		panic(err)
	}
	defer inputData.Close()

	img, err := png.Decode(inputData)
	if err != nil {
		panic(err)
	}

	log.Println(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {

			c := img.At(x, y)
			r, g, b, _ := c.RGBA()

			p := pixel.PostPixel{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				X: *xPos + x,
				Y: *yPos + y,
			}
			PostPixel(*apiServer, p)

			time.Sleep(time.Duration(*sleep) * time.Millisecond)
		}
	}
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

	log.Println("Response Status:", resp.Status, "Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response Body:", string(respBody))

	return nil
}

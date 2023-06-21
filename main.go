package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Product struct {
	ID                 int      `json:"ID"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float32  `json:"discountPercentage"`
	Rating             float32  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

type Response struct {
	Products []Product `json:"products"`
}

const (
	HOST             = "https://dummyjson.com"
	PRODUCT_ENDPOINT = "/products"
)

func GetURL(host, endpoint string) string {
	return host + endpoint
}

func main() {
	productURL := GetURL(HOST, PRODUCT_ENDPOINT)
	log.Printf("=== Product URL: %s", productURL)

	response, err := http.Get(productURL)
	if err != nil {
		log.Fatalf("Failed to make the HTTP request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status code: %d", response.StatusCode)
	}

	var responseProducts Response
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	err = json.Unmarshal(responseBody, &responseProducts)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	var products []Product

	for _, product := range responseProducts.Products {
		var buffer bytes.Buffer
		err := PrettyEncode(product, &buffer)
		if err != nil {
			log.Println(err)
		}
		log.Println(buffer.String())
		products = append(products, product)
	}
}

func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

package main

import (
	"github.com/DiasOrazbaev/ads-service/internal/ads"
	"github.com/oschwald/geoip2-golang"
	"log"
)

// TODO: Make filter by array of country and browser
// TODO: Make browser and country enum
// TODO: Refactoring code
func main() {
	geoip, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatalln(err)
	}
	s := ads.NewServer(geoip)

	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}

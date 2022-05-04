package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	postgres "github.com/kadirgonen/Go-Bootcamp/lesson_4/location-service/common/db"
	"github.com/kadirgonen/Go-Bootcamp/lesson_4/location-service/domain/city"
	"github.com/kadirgonen/Go-Bootcamp/lesson_4/location-service/domain/country"
)

func main() {
	//Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	// Repositories
	cityRepo := city.NewCityRepository(db)
	cityRepo.Migrations()
	cityRepo.InsertSampleData()

	//fmt.Println(len(cityRepo.FindAll()))
	//fmt.Println(cityRepo.FindByCountryCode("TR"))

	countryRepo := country.NewCountryRepository(db)
	countryRepo.Migration()
	countryRepo.InsertSampleData()

	fmt.Println(countryRepo.GetAllCountriesWithCityInformation())

}

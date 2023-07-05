package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andreluizmicro/go-soccer/internal/country/domain/entity"
	_ "github.com/go-sql-driver/mysql"
)

type CountryRepository struct {
}

func NewCountryRepository() *CountryRepository {
	return &CountryRepository{}
}

func (repository *CountryRepository) Create(country *entity.Country) error {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/go_soccer")
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO countries(uuid, name, capital, area, language, timezone, continent, official_color, population) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	_, err = stmt.Exec(country.ID, country.Name, country.Capital, country.Area, country.Language, country.Timezone, country.Continent, country.OfficalColor, country.Population)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (repository *CountryRepository) FindAll(page, limit int, sort string) ([]entity.Country, error) {
	var countries []entity.Country
	return countries, nil

}

func (repository *CountryRepository) FindByID(id string) (*entity.Country, error) {
	var country entity.Country
	return &country, nil
}

func (repository *CountryRepository) Update(country *entity.Country) error {
	return nil
}

func (repository *CountryRepository) Delete(id string) error {
	return nil
}

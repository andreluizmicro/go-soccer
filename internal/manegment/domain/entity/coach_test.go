package entity

import (
	"testing"

	country "github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/stretchr/testify/assert"
)

const city = "São Paulo"
const club = "São Paulo Futebol Clube"

func TestNewCoach(t *testing.T) {
	stadium, err := NewStadium("Morumbi", city, 65000)
	assert.Nil(t, err)

	player, err := NewPlayer("Rafael", "Brasileiro", 10, 28, 1)
	assert.Nil(t, err)
	player2, err := NewPlayer("Luciano", "Brasileiro", 10, 30, 10)
	assert.Nil(t, err)
	player3, err := NewPlayer("Luciano", "Brasileiro", 10, 30, 10)
	assert.Nil(t, err)

	players := []Player{*player, *player2, *player3}

	coach, err := NewCoach("Dorival Júnior", "Brasileiro", 60)
	assert.Nil(t, err)

	params := country.Params{
		Name:         "Brasil",
		Capital:      "Brasília",
		Area:         "8.510.000 km²",
		Languages:    "Portuguese",
		Timezone:     "UTC",
		Continent:    "Americano",
		OfficalColor: "verde e amarelo",
		Population:   100000,
	}

	country, err := country.NewCountry(params)
	assert.Nil(t, err)

	props := Props{"São Paulo Futebol Clube", "http://saopaulo.jpeg", *country, *coach, players, *stadium, 1919, 100}

	team, err := NewTeam(props)
	assert.Nil(t, err)
	assert.Equal(t, club, team.Name)
}

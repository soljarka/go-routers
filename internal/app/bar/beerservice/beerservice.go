package beerservice

var beers = []string{"aldaris", "labietis", "malduguns"}

func GetAll() []string {
	return beers
}

func GetById(id int) string {
	return beers[id]
}

func Add(beer string) {
	beers = append(beers, beer)
}

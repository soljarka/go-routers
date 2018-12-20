package snackservice

var snacks = []string{"cheese", "garlic bread", "squid rings"}

func GetAll() []string {
	return snacks
}

func GetById(id int) string {
	return snacks[id]
}

func Add(snack string) {
	snacks = append(snacks, snack)
}

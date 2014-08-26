package benchfiles

type User struct {
	FirstName string
	LastName  string
	Hobbies   []string
	Age       int
}

var user = &User{
	FirstName: "Joe",
	LastName:  "Seabee",
	Hobbies:   []string{"running", "scuba", "rock climbing"},
	Age:       21,
}

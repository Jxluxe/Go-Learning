package main


/*
 defines the template for a Todo item with a struct giving us a way to store and manipulate Todo Items.
 */
type Todo struct {
	ID	int
	Title	string
	Completed	bool
}

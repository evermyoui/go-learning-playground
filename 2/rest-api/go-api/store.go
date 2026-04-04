package main

type UserStorer interface {
	Create(name, email string) (User, error)
	GetByID(id int) (User, error)
	Update(id int, name, email string) (User, error)
	// Delete(id int) error
	List() ([]User, error)
}

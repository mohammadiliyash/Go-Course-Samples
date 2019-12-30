package models

//User struct to store user data for user creation
type User struct{

	ID string `json:"id, omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
}


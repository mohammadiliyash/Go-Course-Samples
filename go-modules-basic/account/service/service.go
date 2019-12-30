package account

import "context"

//Service Interface
//These are the methods you want to be exposed to transport
type Service interface{
	//Create a user based on email and pwd
	CreateUser(c context.Context, email string, pwd string)(string, error)
	//Get a user based on the user id
	GetUser()c context.Context, id string)(string, error)
}



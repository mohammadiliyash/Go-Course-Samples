package respository

import(
	"go-kit-create-user/account/models"
	"context"
)

//UserRepository => Repository for user 
type UserRepository interface{

	CreateUser(ctx context.context, user models.User)
	GetUser(ctx context.Context, id string)

}
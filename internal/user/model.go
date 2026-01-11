package user

type User struct {
	ID    int    	`json:"id"`					
	Name  string 	`json:"name"`
	Email string 	`json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
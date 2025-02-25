package schemas

// UserSignupSchemaIn represents the input data for user sign up
type UserSignupSchemaIn struct {
	FirstName string `json:"first_name" binding:"required,min=1,max=100" example:"John"`
	LastName  string `json:"last_name" binding:"required,min=1,max=100" example:"Doe"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=100"`
}

// UserLoginSchemaIn represents the input data for user log in
type UserLoginSchemaIn struct {
	Email    string `json:"email" binding:"required,email" example:"string@string.com"`
	Password string `json:"password" binding:"required,min=8,max=100" example:"stringstring"`
}

type UserUpdateSchemaIn struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserSchemaOut struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

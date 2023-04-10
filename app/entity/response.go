package entity

//ResponseFailed represents the failed response
type ResponseFailed struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
}

//ResponseSuccess represents the failed response
type ResponseSuccess struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
}

//ResponseLogin represents the response after login
type ResponseLogin struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
	Data    struct {
		Token string `json:"token" example:"eyJhbGciOiJI...."`
	} `json:"data"`
}

//ResponseRegister represents the response after register
type ResponseRegister struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
	Data    struct {
		ID    uint   `json:"id" example:"1"`
		Email string `json:"email" example:"user@mail.com"`
		Uname string `json:"username" example:"user"`
		Age   int    `json:"age" example:"18"`
	} `json:"data"`
}

//Response represents the all the response
type Response struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"created"`
	Data    interface{} `json:"data"`
}

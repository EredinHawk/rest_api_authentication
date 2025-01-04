package auth

// Credentails - структура тела запроса на регистрацию нового пользователя
type Credentails struct {
	Login    string `json:"login"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// FakeDataBase - имитация базы данных пользователей
var FakeDataBase = []Credentails{
	{Login: "login_1", Password: "password_1"},
	{Login: "login_2", Password: "password_2"},
	{Login: "login_3", Password: "password_3"},
	{Login: "login_4", Password: "password_4"},
}

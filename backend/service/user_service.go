package services

import (
	userClient "Arqui_Soft_I/backend/clients/user"
	"Arqui_Soft_I/backend/dto"
	"Arqui_Soft_I/backend/model"
	e "Arqui_Soft_I/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// Creo la estructura
type userService struct{}

// Creo interfaz de la estructura para darle funcionalidad
// Serian los metodos de la estructura
type userServiceInterface interface {
	UserAuth(userDto dto.UserDto) (bool, string, int)
	GetUsers() (dto.Users, e.ApiError)
	RegisterUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	GetUserById(id int) (dto.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

// Implemento las funciones que tenfo en la interface
// Esta funcion verifica si el usuario esta en la bd y si la passwoerd
// es correcta. La implementacion del login se hace en el front
func (s *userService) UserAuth(userDto dto.UserDto) (bool, string, int) {
	//el operador := declara e inicializa!!
	user := userClient.GetUserByUsername(userDto.Username) //Accede a la bd
	//busca en la bd al usuario por su nombre de usuario(Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		return false, "", -1
	}

	return true, user.Rol, user.ID

}

func (s *userService) GetUsers() (dto.Users, e.ApiError) {

	var users model.Users = userClient.GetUsers()
	var usersDto dto.Users

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Username = user.Username
		userDto.ID = user.ID
		userDto.Rol = user.Rol
		userDto.Nombre = user.Nombre
		userDto.Apellido = user.Apellido

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) RegisterUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		// Handle error
	}
	var user model.User

	user.Nombre = userDto.Nombre
	user.Apellido = userDto.Apellido
	user.Password = string(hashedPassword)
	user.Username = userDto.Username
	user.Rol = userDto.Rol

	user = userClient.RegisterUser(user)

	userDto.ID = user.ID

	return userDto, nil
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.Nombre = user.Nombre
	userDto.Apellido = user.Apellido
	userDto.Username = user.Username
	userDto.Rol = user.Rol
	userDto.ID = user.ID

	return userDto, nil
}

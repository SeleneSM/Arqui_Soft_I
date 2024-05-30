package services

import (
	userClient "Arqui_Soft_I/clients/user"
	"Arqui_Soft_I/dto"

	"golang.org/x/crypto/bcrypt"
)

// Creo la estructura
type userService struct{}

// Creo interfaz de la estructura para darle funcionalidad
// Serian los metodos de la estructura
type userServiceInterface interface {
	UserAuth(userDto dto.UserDto) (bool, string, int)
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

	/*if user.Password == userDto.Password {
		return true, user.Rol, user.ID
	}

	return false, "", -1 */
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		return false, "", -1
	}

	return true, user.Rol, user.ID

}

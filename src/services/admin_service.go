package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/samuelbeaulieu1/gimlet/logger"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
)

const (
	tokenDurationMinutes = 15
)

type AdminService struct{}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (service *AdminService) verifyNewPassword(req *classes.UpdateAdminRequest) responses.Error {
	if err := service.Authenticate(req.Password); err != nil {
		return err
	}
	if req.NewPassword != req.NewPasswordRepeat {
		return responses.NewFieldsError([]string{"Les mots de passe ne sont pas identiques"}, []string{"new_password", "new_password_repeat"})
	}
	if req.NewPassword == "" {
		return responses.NewFieldsError([]string{"Le mot de passe est obligatoire"}, []string{"new_password", "new_password_repeat"})
	}

	return nil
}

func (service *AdminService) UpdatePassword(req *classes.UpdateAdminRequest) responses.Error {
	if err := service.verifyNewPassword(req); err != nil {
		return err
	}
	if err := entities.NewAdmin().Update(req.NewPassword); err != nil {
		return responses.NewError("Une erreur inattendue est survenue en sauvegardant le nouveau mot de passe")
	}

	return nil
}

func (service *AdminService) Authenticate(password string) responses.Error {
	adminEntity := entities.NewAdmin()
	admin, err := adminEntity.Get()
	if err != nil {
		return responses.NewError("Une erreur inattendue est survenue en vérifiant le compte administrateur")
	}

	if ok := adminEntity.VerifyPassword(password, admin.Password); !ok {
		return responses.NewFieldsError([]string{"Mot de passe invalide"}, []string{"password"})
	}

	return nil
}

func tokenValidator(token *jwt.Token) (interface{}, error) {
	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		return nil, responses.NewError("Une erreur inattendue est survenue en vérifiant l'authentification")
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, responses.NewError("Authentification invalide")
	}

	return []byte(secret), nil
}

func (service *AdminService) ValidateToken(tokenStr string) responses.Error {
	token, err := jwt.Parse(tokenStr, tokenValidator)
	if err != nil {
		return err
	}
	if !token.Valid {
		return responses.NewError("Authentification invalide")
	}

	return nil
}

func (service *AdminService) CreateSession(req *classes.AdminAuthRequest) (*dto.AuthResponseDTO, responses.Error) {
	if err := service.Authenticate(req.Password); err != nil {
		return nil, err
	}
	expiration := time.Now().Add(time.Minute * tokenDurationMinutes)
	token, err := service.createToken(expiration)

	return &dto.AuthResponseDTO{
		Token:     token,
		ExpiresAt: expiration,
	}, err
}

func (service *AdminService) createToken(expiration time.Time) (string, responses.Error) {
	secret := os.Getenv("TOKEN_SECRET")
	if secret == "" {
		logger.PrintError("No token secret is configured for admin sessions")
		return "", responses.NewError("Une erreur inattendue est survenue en vérifiant l'authentification")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expiration.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", responses.NewError("Une erreur inattendue est survenue en créant la clé d'authentification")
	}

	return tokenString, nil
}

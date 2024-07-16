package mappings

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/mapper"
	registeringuserdtosv1 "github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/features/registering_user/v1/dtos"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/models"
)

func ConfigureMappings() error {
	err := mapper.CreateMap[*models.User, *registeringuserdtosv1.RegisterUserResponseDto]()
	if err != nil {
		return err
	}
	return err
}

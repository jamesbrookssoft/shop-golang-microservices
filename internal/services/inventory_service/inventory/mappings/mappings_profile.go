package mappings

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/mapper"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/inventory_service/inventory/consumers/events"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/inventory_service/inventory/dtos"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/inventory_service/inventory/models"
)

func ConfigureMappings() error {
	err := mapper.CreateMap[*models.Inventory, *dtos.InventoryDto]()
	if err != nil {
		return err
	}

	err = mapper.CreateMap[*models.ProductItem, *events.InventoryUpdated]()
	if err != nil {
		return err
	}

	return nil
}

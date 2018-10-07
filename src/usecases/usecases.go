package usecases

import (
	"fmt"

	"github.com/bkyoung/inventory/src/domain"
)

// type UserRepository interface {
// 	Store(user User)
// 	FindById(id int) User
// }

// type User struct {
// 	Id       int
// 	IsAdmin  bool
// 	Customer domain.Customer
// }

// Logger allows us to log without worrying about implementation details
type Logger interface {
	Log(args ...interface{})
}

// InventoryInteractor allows us to interact with Inventory
type InventoryInteractor struct {
	// UserRepository  UserRepository
	InventoryRepository domain.InventoryRepository
	GroupRepository     domain.GroupRepository
	HostRespoitory      domain.HostRepository
	VariableRepository  domain.VariableRepository
	Logger              Logger
}

// FindById gets the specified Inventory
func (interactor *InventoryInteractor) FindById(inventoryId int) (domain.Inventory, error) {
	inventory := interactor.InventoryRepository.FindById(inventoryId)
	return inventory, nil
}

// AddGroupToInventory allows us to add groups to inventory
func (interactor *InventoryInteractor) AddGroupToInventory(inventoryId, groupId int) error {
	inventory := interactor.InventoryRepository.FindById(inventoryId)
	group := interactor.GroupRepository.FindById(groupId)
	if err := inventory.AddGroupToInventory(group); err != nil {
		e := fmt.Errorf("Unable to add group %d to inventory %d: %s", group.Id, inventory.Id, err.Error())
		interactor.Logger.Log(e.Error())
		return e
	}
	interactor.GroupRepository.AddToInventory(inventory, group)
	interactor.Logger.Log(fmt.Sprintf("Added group '%s' to inventory '%s' (id %d)", group.Name, inventory.Name, inventory.Id))
	return nil
}

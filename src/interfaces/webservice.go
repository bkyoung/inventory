package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bkyoung/inventory/src/domain"
)

// InventoryInteractor allows us to inject this interface 
// into inner layers, to honor the dependency rule
type InventoryInteractor interface {
	FindById(inventoryId int) (domain.Inventory, error)
	AddGroupToInventory(inventoryId, groupId int) error
}

// WebserviceHandler wraps our interactor to map it to various handler functions
type WebserviceHandler struct {
	InventoryInteractor InventoryInteractor
}

// ShowInventory sends a complete inventory to requestor
func (handler WebserviceHandler) ShowInventory(res http.ResponseWriter, req *http.Request) {
	inventoryId, _ := strconv.Atoi(req.FormValue("inventoryId"))
	inventory, _ := handler.InventoryInteractor.FindById(inventoryId)
	io.WriteString(res, fmt.Sprintf("inventory: %v\n", inventory))
}

package main

import (
	"github.com/bkyoung/inventory/src/usecases"
	"github.com/bkyoung/inventory/src/interfaces"
	"github.com/bkyoung/inventory/src/infrastructure"
	"net/http"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("test.sqlite")

	handlers := make(map[string] interfaces.DbHandler)
	handlers["DbInventoryRepo"] = dbHandler
	handlers["DbGroupRepo"] = dbHandler
	handlers["DbHostRepo"] = dbHandler
	handlers["DbVariableRepo"] = dbHandler

	inventoryInteractor := new(usecases.InventoryInteractor)
	//inventoryInteractor.VariableRepository = interfaces.NewDbVariableRepo(handlers)
	//inventoryInteractor.HostRepository = interfaces.NewDbHostRepo(handlers)
	inventoryInteractor.GroupRepository = interfaces.NewDbGroupRepo(handlers)
	inventoryInteractor.InventoryRepository = interfaces.NewDbInventoryRepo(handlers)
	inventoryInteractor.Logger = new(infrastructure.Logger)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.InventoryInteractor = inventoryInteractor

	http.HandleFunc("/inventory", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowInventory(res, req)
	})
	http.ListenAndServe(":8080", nil)
}

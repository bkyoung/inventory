package domain

// VariableRepository is our interface definition for interacting with inventory/group/host variables
type VariableRepository interface {
	Store(variable Variable)
	FindById(id int) Variable
}

// HostRepository is our interface for interacting with Host objects
type HostRepository interface {
	Store(host Host)
	FindById(id int) Host
}

// GroupRepository is our interface for interacting with Group objects
type GroupRepository interface {
	Store(group Group)
	FindById(id int) Group
	AddToInventory(inventory Inventory, group Group)
}

// InventoryRepository is our interface for interacting with Inventory objects
type InventoryRepository interface {
	Store(inventory Inventory)
	FindById(id int) Inventory
	AddGroupToInventory(inventory Inventory, group Group)
}

// Variable allows us to represent inventory/group/host vars
type Variable struct {
	Id    int
	Name  string
	Value interface{}
}

// Host represents a host in inventory
type Host struct {
	Id   int
	Name string
}

// Group represents a group in inventory
type Group struct {
	Id   int
	Name string
}

// Inventory is a collection of Groups, Hosts, and Variables
type Inventory struct {
	Id        int
	Name      string
	Groups    []Group
	Hosts     []Host
	Variables []Variable
}

func (inventory *Inventory) AddGroupToInventory(group Group) error {
	inventory.Groups = append(inventory.Groups, group)
	return nil
}

func (inventory *Inventory) AddHostToInventory(host Host) error {
	inventory.Hosts = append(inventory.Hosts, host)
	return nil
}

func (inventory *Inventory) AddVariableToInventory(variable Variable) error {
	inventory.Variables = append(inventory.Variables, variable)
	return nil
}
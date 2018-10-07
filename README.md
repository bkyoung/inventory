# Inventory
A configuration management inventory service.  Initially it is targeting Ansible, but the database is general enough to support other configuration management systems in the future without a great deal of structural modification.

The structure of the project follows the principles of [clean architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).

## How to use
Very early state of development.  Not at all usable for real work.

Be sure you have the go-sqlite3 package installed:
```
$ go get -u github.com/mattn/go-sqlite3
```

Create a test.sqlite database in the top level of this repo:
```
$ sqlite3 test.sqlite < inventory.sql
$ sqlite3 test.sqlite < inventory_test.sql
```

Run the service:
```
$ go run main.go
```

Make an http request for a basic inventory:
```
$ curl 'http://localhost:8080/inventory?inventoryId=1'
```

## Design(ish)
A host is modeled as an entity that contains:
* host-level variables (i.e. host_vars in Ansible)

A group is modeled as an entity that contains:
* group-level variables (i.e. group_vars in Ansible)
* groups
* hosts

An inventory is modeled as a entity that contains:
* inventory-level variables (i.e. [all:vars] in Ansible)
* groups
* hosts

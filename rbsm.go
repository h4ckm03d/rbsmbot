package rbsm

type Database interface {
	PromoteAdmin(username string) error
	DemoteAdmin(username string) error
}

type rbsm struct {
	db          Database
	superAdmins []string
	debugMode   bool
}

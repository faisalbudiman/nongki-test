package factory

import (
	"nongki-test/database"
	"nongki-test/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db                *gorm.DB
	UserRepository    repository.User
	AddressRepository repository.Address
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("V0")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
	f.AddressRepository = repository.NewAddress(f.Db)
}

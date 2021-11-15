package dao

type GenericDB interface {
	Connect()
	Insert(entity interface{}) (interface{}, error)
	Update(entity interface{}) (interface{}, error)
	Get(id string, entity interface{}, preload ...string) (bool, error)
	GetAll(entity interface{}) error
	Delete(entity interface{}, id string) (interface{}, error)
}

type DBConnection struct {
	GenericDB
}

func NewDBConnection(db GenericDB) DBConnection {
	return DBConnection{db}
}

func (c DBConnection) DBConnect() {
	c.Connect()
}

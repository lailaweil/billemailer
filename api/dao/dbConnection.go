package dao

type GenericDB interface {
	Connect()
	Insert(entity interface{}) (interface{}, error)
	Get(id string, entity interface{}) error
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

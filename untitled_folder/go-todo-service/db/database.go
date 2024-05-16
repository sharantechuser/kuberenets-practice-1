package db

var Handle dbHandle

type dbHandle interface {
	Create(i interface{}) error
	GetAll(coll_name string, results interface{}) (interface{}, error)
	Update(_id string, v interface{}) (int64, error)
	Delete(_id string, v interface{}) (int64, error)
	IsValid() bool
}

func Initialize() bool {
	Handle = NewMongo()

	return Handle.IsValid()
}

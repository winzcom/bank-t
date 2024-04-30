package db

type Storage struct{}

var tkl string

func (s Storage) Save(t interface{}) {

}

func GetStorage() Storage {
	return Storage{}
}

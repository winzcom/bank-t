package db

var store []interface{}

func Save(t interface{}) {
	store = append(store, t)
}

package biz

type ICollection interface {
	Add(docs ...interface{}) error
}

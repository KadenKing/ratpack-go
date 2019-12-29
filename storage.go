package main

type storage interface {
	Save(key string, value interface{}) error
}

type postgres struct {
	db string
}

func newPostgres() *postgres {
	return &postgres{}
}

func (p *postgres) Save(key string, value interface{}) error {
	return nil
}

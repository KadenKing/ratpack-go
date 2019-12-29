package main

type storage interface {
	pointIncrementer
}

type pointIncrementer interface {
	IncrementPoints(user string, points int64) error
}

type postgres struct {
	db string
}

func newPostgres() *postgres {
	return &postgres{}
}

func (p *postgres) IncrementPoints(user string, points int64) error {
	return nil
}

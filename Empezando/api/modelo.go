package api

type inter interface {
	ahre()
}

type Modelo struct {
	id   int
	name string
	date bool
}

func (m Modelo) ahre(id int, name string, date bool) {
	m.id = id
	m.name = name
	m.date = date
}

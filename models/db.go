package models

type DB interface {
	Put(p *Post) error
	Get(p *Post) (string, error)
	Count(p *Post) (string, error)
}

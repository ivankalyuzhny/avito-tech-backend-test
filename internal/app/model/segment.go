package model

type Segment struct {
	ID   int64  `db:"id"`
	Slug string `db:"slug"`
}

func NewSegment(id int64, slug string) *Segment {
	return &Segment{
		ID:   id,
		Slug: slug,
	}
}

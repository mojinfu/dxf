package entity

type Insert struct {
	*entity
}

func (t *Insert) IsEntity() bool {
	return true
}

// NewText creates a new Insert.
func NewInsert() *Insert {
	t := &Insert{
		entity: NewEntity(INSERT),
	}
	return t
}
func (p *Insert) BBox() ([]float64, []float64) {
	return []float64{}, []float64{}
}

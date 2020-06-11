package entity

type Mtext struct {
	*entity
}

func (t *Mtext) IsEntity() bool {
	return true
}

// NewText creates a new Mtext.
func NewMtext() *Mtext {
	t := &Mtext{
		entity: NewEntity(MTEXT),
	}
	return t
}
func (p *Mtext) BBox() ([]float64, []float64) {
	return []float64{}, []float64{}
}

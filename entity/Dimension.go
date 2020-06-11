package entity

type Dimension struct {
	*entity
}

func (t *Dimension) IsEntity() bool {
	return true
}

// NewText creates a new Dimension.
func NewDimension() *Dimension {
	t := &Dimension{
		entity: NewEntity(DIMENSION),
	}
	return t
}
func (p *Dimension) BBox() ([]float64, []float64) {
	return []float64{}, []float64{}
}

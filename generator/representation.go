package generator

type Representation struct {
    name string
    base int
}

func (representation Representation) Name() string {
    return representation.name
}

func (representation Representation) Base() int {
    return representation.base
}



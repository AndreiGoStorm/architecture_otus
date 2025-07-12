package hw5

type SpaceObject interface {
	GetProperty(key string) interface{}
	SetProperty(key string, value interface{})
}

type Spaceship struct {
	properties map[string]interface{}
}

func NewSpaceship() SpaceObject {
	return &Spaceship{make(map[string]interface{})}
}

func (s *Spaceship) GetProperty(key string) interface{} {
	if item, ok := s.properties[key]; ok {
		return item
	}
	return nil
}

func (s *Spaceship) SetProperty(key string, value interface{}) {
	s.properties[key] = value
}

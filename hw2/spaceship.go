package hw2

type SpaceObject interface {
	getProperty(key string) interface{}
	setProperty(key string, value interface{})
}

type Spaceship struct {
	properties map[string]interface{}
}

func NewSpaceship() SpaceObject {
	return &Spaceship{make(map[string]interface{})}
}

func (s *Spaceship) getProperty(key string) interface{} {
	if item, ok := s.properties[key]; ok {
		return item
	}
	return nil
}

func (s *Spaceship) setProperty(key string, value interface{}) {
	s.properties[key] = value
}

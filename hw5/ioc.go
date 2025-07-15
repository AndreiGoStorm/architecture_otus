package hw5

import (
	"fmt"
	"sync"
)

type IoC struct {
	registrations map[string]func(...interface{}) interface{}
	singletons    map[string]interface{}
	mutex         sync.Mutex
}

func NewIoC() *IoC {
	return &IoC{
		registrations: make(map[string]func(...interface{}) interface{}),
		singletons:    make(map[string]interface{}),
	}
}

func (ioc *IoC) register(args ...interface{}) {
	if len(args) < 2 {
		panic(ErrNotEnoughArgs)
	}
	name, ok := args[0].(string)
	if !ok {
		panic(ErrNoKeyArgument)
	}
	factory, ok := args[1].(func(...interface{}) interface{})
	if !ok {
		panic(ErrNoFactoryArgument)
	}

	if ioc.isSingleton(args...) {
		ioc.registrations[name] = func(args ...interface{}) interface{} {
			ioc.mutex.Lock()
			defer ioc.mutex.Unlock()
			if instance, ok := ioc.singletons[name]; ok {
				return instance
			}
			instance := factory(args...)
			ioc.singletons[name] = instance
			return instance
		}
	} else {
		ioc.mutex.Lock()
		ioc.registrations[name] = factory
		ioc.mutex.Unlock()
	}
}

func (ioc *IoC) isSingleton(args ...interface{}) bool {
	isSingleton := false
	if len(args) >= 3 {
		if flag, ok := args[2].(bool); ok {
			isSingleton = flag
		}
	}
	return isSingleton
}

func (ioc *IoC) Resolve(key string, args ...interface{}) interface{} {
	if key == "Ioc.Register" {
		ioc.register(args...)
		return nil
	}

	ioc.mutex.Lock()
	factory, ok := ioc.registrations[key]
	ioc.mutex.Unlock()
	if !ok {
		panic(fmt.Errorf("%w%s", ErrNoRegistration, key))
	}

	return factory(args...)
}

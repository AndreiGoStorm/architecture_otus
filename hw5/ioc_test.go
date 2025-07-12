package hw5

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIocContainer(t *testing.T) {
	t.Parallel()
	t.Run("register spaceship object", func(t *testing.T) {
		t.Parallel()
		// Arrange
		container := NewIoC()

		// Act
		container.Resolve("Ioc.Register", "Spaceship",
			func(...interface{}) interface{} {
				return NewSpaceship()
			},
		)

		// Assert
		spaceShip := container.Resolve("Spaceship").(SpaceObject)
		assert.NotNil(t, spaceShip)
		assert.IsType(t, &Spaceship{}, spaceShip)
	})

	t.Run("register rotate adapter", func(t *testing.T) {
		t.Parallel()
		// Arrange
		container := NewIoC()
		container.Resolve("Ioc.Register", "Spaceship",
			func(...interface{}) interface{} {
				return NewSpaceship()
			},
		)

		// Act
		container.Resolve("Ioc.Register", "RotatableObjectAdapter",
			func(args ...interface{}) interface{} {
				return NewRotatableObjectAdapter(args[0].(SpaceObject))
			},
		)

		// Assert
		spaceShip := container.Resolve("Spaceship").(SpaceObject)
		adapter := container.Resolve("RotatableObjectAdapter", spaceShip).(Rotatable)
		assert.IsType(t, &RotatableObjectAdapter{}, adapter)
	})

	t.Run("register move command", func(t *testing.T) {
		t.Parallel()
		// Arrange
		container := NewIoC()
		container.Resolve("Ioc.Register", "Spaceship",
			func(...interface{}) interface{} {
				return NewSpaceship()
			},
		)
		container.Resolve("Ioc.Register", "RotatableObjectAdapter",
			func(args ...interface{}) interface{} {
				return NewRotatableObjectAdapter(args[0].(SpaceObject))
			},
		)

		// Act
		container.Resolve("Ioc.Register", "RotateCommand",
			func(args ...interface{}) interface{} {
				return NewRotate(
					container.Resolve("RotatableObjectAdapter", args[0].(SpaceObject)).(Rotatable),
				)
			},
		)

		// Assert
		spaceShip := container.Resolve("Spaceship").(SpaceObject)
		cmd := container.Resolve("RotateCommand", spaceShip).(Command)
		assert.IsType(t, &Rotate{}, cmd)
	})
}

func TestIocContainerSingleton(t *testing.T) {
	t.Parallel()
	t.Run("register spaceship object singleton", func(t *testing.T) {
		t.Parallel()
		// Arrange
		container := NewIoC()

		// Act
		container.Resolve("Ioc.Register", "Spaceship",
			func(...interface{}) interface{} {
				return NewSpaceship()
			},
			true,
		)

		// Assert
		spaceShip1 := container.Resolve("Spaceship").(SpaceObject)
		spaceShip2 := container.Resolve("Spaceship").(SpaceObject)
		if spaceShip1 != spaceShip2 {
			t.Error("objects should be the same")
		}
	})

	t.Run("register spaceship object without singleton", func(t *testing.T) {
		t.Parallel()
		// Arrange
		container := NewIoC()

		// Act
		container.Resolve("Ioc.Register", "Spaceship",
			func(...interface{}) interface{} {
				return NewSpaceship()
			},
		)

		// Assert
		spaceShip1 := container.Resolve("Spaceship").(SpaceObject)
		spaceShip2 := container.Resolve("Spaceship").(SpaceObject)
		if spaceShip1 == spaceShip2 {
			t.Error("objects should not be the same")
		}
	})
}

func TestIocContainerConcurrency(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	const workers = 100

	t.Run("concurrency test for resolves", func(t *testing.T) {
		// Arrange
		container := NewIoC()
		var wg sync.WaitGroup

		// Act
		for i := 0; i < workers; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				container.Resolve("Ioc.Register", "Spaceship",
					func(...interface{}) interface{} {
						return NewSpaceship()
					},
				)
				container.Resolve("Ioc.Register", "RotatableObjectAdapter",
					func(args ...interface{}) interface{} {
						return NewRotatableObjectAdapter(args[0].(SpaceObject))
					},
				)
				container.Resolve("Ioc.Register", "RotateCommand",
					func(args ...interface{}) interface{} {
						return NewRotate(
							container.Resolve("RotatableObjectAdapter", args[0].(SpaceObject)).(Rotatable),
						)
					},
				)
				spaceShip := container.Resolve("Spaceship").(SpaceObject)
				cmd := container.Resolve("RotateCommand", spaceShip).(Command)

				// Assert
				assert.NotNil(t, cmd)
			}()
		}

		wg.Wait()
	})
}

func TestIocContainerForPanic(t *testing.T) {
	t.Run("panic: not enough arguments", func(t *testing.T) {
		// Arrange
		container := NewIoC()

		// Assert
		assert.PanicsWithError(t, ErrNotEnoughArgs.Error(), func() {
			container.Resolve("Ioc.Register")
		})
	})

	t.Run("panic: first argument must be string", func(t *testing.T) {
		// Arrange
		container := NewIoC()

		// Assert
		assert.PanicsWithError(t, ErrNoKeyArgument.Error(), func() {
			container.Resolve("Ioc.Register", 12, true)
		})
	})

	t.Run("panic: second argument must be function", func(t *testing.T) {
		// Arrange
		container := NewIoC()

		// Assert
		assert.PanicsWithError(t, ErrNoFactoryArgument.Error(), func() {
			container.Resolve("Ioc.Register", "MoveCommand", true)
		})
	})

	t.Run("panic: no registration for key", func(t *testing.T) {
		// Arrange
		container := NewIoC()
		key := "MoveCommand"

		// Assert
		assert.PanicsWithError(t, fmt.Sprintf("%v%s", ErrNoRegistration, key), func() {
			container.Resolve(key)
		})
	})
}

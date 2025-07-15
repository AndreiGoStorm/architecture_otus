package hw5

import "errors"

var ErrNotEnoughArgs = errors.New("not enough arguments")
var ErrNoKeyArgument = errors.New("first argument must be string")
var ErrNoFactoryArgument = errors.New("second argument must be function")
var ErrNoRegistration = errors.New("no registration for key: ")
var ErrNoAngle = errors.New("cannot read angle from properties")
var ErrNoRotationAngle = errors.New("cannot read rotation angle from properties")

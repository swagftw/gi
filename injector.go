package gi

import (
	"errors"
	"reflect"
	"sync"
)

var (
	i = New()
)

// Injector is a dependency injection container
type Injector struct {
	mu        sync.Mutex
	container map[any]any
}

// New returns a new Injector
func New() *Injector {
	return &Injector{
		mu:        sync.Mutex{},
		container: make(map[any]any),
	}
}

// Inject injects a dependency into the container
func Inject(o any) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	val := reflect.ValueOf(o)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		t := reflect.TypeOf(o)
		return &ErrInvalidArg{Type: t}
	}

	i.container[val.Type().String()] = o

	return nil
}

// ErrInvalidArg error returned when the argument to Inject is not a pointer
// (The argument to Inject must be a non-nil pointer.)
type ErrInvalidArg struct {
	Type reflect.Type
}

func (e *ErrInvalidArg) Error() string {
	if e.Type == nil {
		return "gi: Inject(passed nil value)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "gi: Inject(non-pointer " + e.Type.String() + ")"
	}

	return "gi: Inject(nil " + e.Type.String() + ")"
}

func Invoke[T any]() (t T, err error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// get passed type
	pt := empty[T]()

	typ := reflect.TypeOf(pt)
	if typ == nil {
		return empty[T](), ErrNilType
	}

	obj := i.container[typ.String()]

	if _, ok := obj.(T); ok {
		return obj.(T), nil
	}

	return empty[T](), ErrServiceNotFound
}

func empty[T any]() (t T) {
	return
}

var (
	ErrServiceNotFound error = errors.New("service for provide type not found")
	ErrNilType         error = errors.New("nil type not allowed")
)

package gi

import (
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Random struct {
	Number int
}

func TestInject(t *testing.T) {
	tests := []struct {
		name    string
		obj     any
		wantErr error
	}{
		{
			name:    "Injecting a nil value",
			wantErr: &ErrInvalidArg{Type: reflect.TypeOf(nil)},
			obj:     nil,
		},
		{
			name:    "Injecting a non-pointer value",
			wantErr: &ErrInvalidArg{Type: reflect.TypeOf(Random{Number: 1})},
			obj:     Random{Number: 1},
		},
		{
			name:    "Injecting a pointer",
			wantErr: nil,
			obj:     &Random{Number: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Inject(test.obj)
			assert.Equal(t, err, test.wantErr)
		})
	}
}

func TestInvoke(t *testing.T) {
	tests := []struct {
		name      string
		injectObj any
		wantSame  bool
		wantErr   bool
	}{
		{
			name:      "getting a injected value",
			wantSame:  true,
			wantErr:   false,
			injectObj: &Random{Number: 1},
		},
		{
			name:      "getting wrong value",
			wantSame:  false,
			wantErr:   false,
			injectObj: &Random{Number: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Inject(test.injectObj)
			assert.Equal(t, err, nil)

			var val any

			val, err = Invoke[*Random]()

			assert.Equal(t, err != nil, test.wantErr)

			if test.wantSame {
				assert.Equal(t, val, test.injectObj)
			} else {
				assert.NotEqual(t, val, nil)
			}
		})
	}
}

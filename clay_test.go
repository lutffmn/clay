package clay

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	right := &App{
		name:        "app",
		description: "description",
		usage:       "usage",
		version:     "v1",
		feature:     make(Feature),
	}

	left := New("app", "description", "usage", "v1")

	if !reflect.DeepEqual(left, right) {
		t.Errorf("left: %v\tright: %v", left, right)
	}
}

func TestShortFlag(t *testing.T) {
	right := make(Feature)
	right["-d"] = Action{
		handler: dummyHandler,
	}

	left := New("app", "description", "usage", "v1")
	left.Short("d", dummyHandler)

	for key, rightAction := range right {
		leftAction, ok := left.feature[key]
		if !ok {
			t.Errorf("left: %v\tright: %v\n", leftAction, rightAction)
		}

		leftResult := reflect.ValueOf(leftAction.handler)
		rightResult := reflect.ValueOf(rightAction.handler)

		if !reflect.DeepEqual(leftResult.Pointer(), rightResult.Pointer()) {
			t.Errorf("left: %v\tright: %v\n", leftAction, rightAction)
		}
	}
}

func TestLongFlag(t *testing.T) {
	right := make(Feature)
	right["--dummy"] = Action{
		handler: dummyHandler,
	}

	left := New("app", "description", "usage", "v1")
	left.Long("dummy", dummyHandler)

	for key, rightAction := range right {
		leftAction, ok := left.feature[key]
		if !ok {
			t.Errorf("left: %v\tright: %v\n", leftAction, rightAction)
		}

		leftResult := reflect.ValueOf(leftAction.handler)
		rightResult := reflect.ValueOf(rightAction.handler)

		if !reflect.DeepEqual(leftResult.Pointer(), rightResult.Pointer()) {
			t.Errorf("left: %v\tright: %v\n", leftAction, rightAction)
		}
	}
}

func TestParse(t *testing.T) {
	// TODO: test parse

	app := New("app", "description", "usage", "v1")
	app.Short("d", func() {
		t.Log("short flag\n")
	})
	app.Long("dummy", func(arg string) {
		t.Logf("long flag with argument: %s\n", arg)
	})

	os.Args = []string{"app", "-d", "--dummy", "argument"}

	app.Parse()
}

func dummyHandler() {
}

package tests

type Test interface {
	// Name returns the name of the test.
	Name() string
	// Run runs the test.
	Run() error
}

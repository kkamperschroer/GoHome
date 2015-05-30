package core

type Rule interface {
	Event() Event
	Actions() []Action
	Conditions() []Condition
	Name() string
}

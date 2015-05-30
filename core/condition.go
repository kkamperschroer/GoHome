package core

type Condition interface {
	IsAnd() bool
	IsOrg() bool
	IsNot() bool
	Description() string
}

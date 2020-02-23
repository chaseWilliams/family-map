package models

/*
Model is an interface for all persistent data resources
*/
type Model interface {
	get(ids ...int) []Model
}

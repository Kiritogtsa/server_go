package services

type services interface {
	validacao() bool
	verificacao() bool
	Create() bool
	Update() bool
	Delete() bool
}

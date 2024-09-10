package repostory

type Crud[T any] interface {
	save(*T) error
	update(*T) error
	Persti(*T) error
	Delete(int) error
}

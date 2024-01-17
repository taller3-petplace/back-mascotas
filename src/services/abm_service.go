package services

type ABMService[T any] interface {
	New(T) (T, error)
	Get(int) (T, error)
	Edit(int, T) (T, error)
	Delete(int)
}

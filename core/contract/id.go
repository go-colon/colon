package contract

const IDKey = "colon:id"

type IDService interface {
	NewID() string
}

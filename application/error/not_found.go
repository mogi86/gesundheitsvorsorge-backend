package error

type NotFoundErr struct{}

func (e *NotFoundErr) Error() string {
	return "Not Found Entity"
}

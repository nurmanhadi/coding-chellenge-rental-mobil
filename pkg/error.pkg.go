package pkg

type BadRequest struct {
	Message string
}
type NotFound struct {
	Message string
}

func (v *BadRequest) Error() string {
	return v.Message
}
func (v *NotFound) Error() string {
	return v.Message
}

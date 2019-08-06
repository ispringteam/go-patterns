package errors

type causer interface {
	Cause() error
}

func getCause(err error) error {
	if c, ok := err.(causer); ok {
		return c.Cause()
	}
	return nil
}

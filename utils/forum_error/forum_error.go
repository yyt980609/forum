package forum_error

import "fmt"

type ForumError struct {
	Code int
}

func New(code int) error {
	return &ForumError{Code: code}
}

func (e ForumError) Error() string {
	return fmt.Sprintf("code:%d", e.Code)
}

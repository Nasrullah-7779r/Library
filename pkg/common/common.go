package common

import "time"

type Student struct {
	//ID       primitive.ObjectID `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentOut struct {
	//ID       primitive.ObjectID `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginCred struct {
	Name     string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type RequestStatus string

const (
	PENDING    RequestStatus = "pending"
	IN_PROCESS RequestStatus = "in_process"
	COMPLETED  RequestStatus = "completed"
)

type BorrowRequest struct {
	RequestID    string        `json:"request-id"`
	BookTitle    string        `json:"book_title" binding:"required" validate:"required"`
	BookAuthor   string        `json:"book_author" binding:"required" validate:"required"`
	BorrowerName string        `json:"borrower_name" binding:"required" validate:"required"`
	Status       RequestStatus `json:"status"`
	Time         time.Time     `json:"time"`
}

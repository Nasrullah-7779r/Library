package common

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
	Name     string `form:"username" binding:"required" validate:"required"`
	Password string `form:"password" binding:"required" validate:"required"`
}

type RequestStatus string

const (
	PENDING   RequestStatus = "pending"
	ISSUED    RequestStatus = "issued"
	COMPLETED RequestStatus = "completed"
)

type BorrowRequest struct {
	RequestID    string        `json:"request_id"`
	BookID       uint          `json:"book_id"`
	BookTitle    string        `json:"book_title" binding:"required" validate:"required"`
	BookAuthor   string        `json:"book_author" binding:"required" validate:"required"`
	BorrowerName string        `json:"borrower_name"  validate:"required"`
	Status       RequestStatus `json:"status"`
	CreatedAt    string        `json:"created_at"`
}

package librarian

type librarian struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type book struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	IsBorrowed  bool   `json:"is_borrowed" default:"false"`
}

type bookStatus struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

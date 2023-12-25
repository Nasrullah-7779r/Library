package student

type Student struct {
	//ID       primitive.ObjectID `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

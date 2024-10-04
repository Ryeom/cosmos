package user

type User struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	StaffType   string `json:"staff_type" bson:"staff_type"`     // 사장/매니저/직원 권한 다름
	JobFunction string `json:"job_function" bson:"job_function"` // 주방/홀/보안 등 역할군을 정의
}

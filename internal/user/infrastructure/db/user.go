package db

type User struct {
	tableName struct{} `pg:"users"`

	ID       int64  `pg:"id,pk"`
	Email    string `pg:"email,unique"`
	Password string `pg:"password"`
	Role     string `pg:"role"`
	CreateAt int64  `pg:"create_at"`
}

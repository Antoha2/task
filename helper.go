package task

type ContextKey string

const (
	USER_ID   ContextKey = "userId"
	USER_ROLE ContextKey = "userRole"
)

type User struct {
	UserId    int      `json:"-" db:"user_id" gorm:"primaryKey"`  //; index:,unique"` //gorm:"primaryKey column:UserId"` //; many2many:userlist_to_roles; index:,unique; ; foreignKey:UserId;joinForeignKey:RoleIdID;References:RoleId;joinReferences:RolesUserId"`
	FirstName string   `json:"firstname" gorm:"column:firstname"` // many2many:userlist_to_roles;"`
	LastName  string   `json:"lastname" gorm:"column:lastname"`
	Username  string   `json:"username" gorm:"column:username"`
	Password  string   `json:"password"`
	Roles     []*Roles `gorm:"many2many:userlist_to_roles"`
}

func (User) TableName() string {
	return "userlist"
}

type UsersRoles struct {
	Roles []string `json:"roles"`
}

type Roles struct {
	RoleId    int     `gorm:"primaryKey"` //  index:,unique;"`
	RolesName string  `gorm:"column:rolesname"`
	Users     []*User `gorm:"many2many:userlist_to_roles;"`
}

func (Roles) TableName() string {
	return "roles"
}

type UserlistToRoles struct {
	UserId int `gorm:"primaryKey"`
	RoleId int `gorm:"primaryKey"`
}

func (UserlistToRoles) TableName() string {
	return "userlist_to_roles"
}

type Task struct {
	Id     int    `json:"task_id" gorm:"primaryKey; column:task_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone" gorm:"column:isdone"`
}

func (Task) TableName() string {
	return "todolist"
}

package store

type User struct {
	ID       uint
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Session struct {
	ID        uint   `bson:"objectId"`
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	User      User
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type SessionStore interface {
	CreateSessios(session *Session) (*Session, error)
	GetUserSession(sessionID string, userID string) (*User, error)
}

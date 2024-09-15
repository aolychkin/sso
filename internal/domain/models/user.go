package models

type UserID struct {
	ID       int64
	Email    string
	PassHash []byte
}

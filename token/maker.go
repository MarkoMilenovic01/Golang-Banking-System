package token

import "time"

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}

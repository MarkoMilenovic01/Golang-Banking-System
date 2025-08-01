package test

import (
	"github.com/dgrijalva/jwt-go"
	token2 "github.com/marko/simplebank/token"
	"github.com/marko/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker, err := token2.NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

//func TestExpiredToken(t *testing.T) {
//	maker, err := NewJWTMaker(util.RandomString(32))
//	require.NoError(t, err)
//
//	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
//	require.NoError(t, err)
//	require.NotEmpty(t, token)
//
//	payload, err := maker.VerifyToken(token)
//	require.Error(t, err)
//	require.EqualError(t, err, "Token is expired")
//	require.Nil(t, payload)
//}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := token2.NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	maker, err := token2.NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, token2.ErrInvalidToken.Error())
	require.Nil(t, payload)

}

package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bernie-pham/cloneSimpleBank/ultilities"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, username, fullname string) User {

	newUser := CreateUserParams{
		Username:       username,
		FullName:       fullname,
		Email:          fmt.Sprintf("%s@gmail.com", username),
		HashedPassword: "123456",
	}
	user, err := testQueries.CreateUser(context.Background(), newUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	require.Equal(t, fmt.Sprintf("%s@gmail.com", username), user.Email)
	require.Equal(t, fullname, user.FullName)
	return user
}

func TestCreateUser(t *testing.T) {
	username := ultilities.RandomString(6)
	fullname := ultilities.RandomFullName()
	createRandomUser(t, username, fullname)
}

func TestGetUser(t *testing.T) {
	username := ultilities.RandomString(6)
	fullname := ultilities.RandomFullName()
	user1 := createRandomUser(t, username, fullname)
	user2, err := testQueries.GetUser(context.Background(), username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)

}

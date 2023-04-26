package db

import (
	"context"
	"testing"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedpassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)
	arg := CreateUsersParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedpassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	// require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUsers(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user1, err := testQueries.GetUsers(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Email, user1.Email)

	require.WithinDuration(t, user.PasswordChangedAt, user1.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, user1.CreatedAt, time.Second)

}

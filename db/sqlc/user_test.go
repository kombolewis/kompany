package db

import (
	"context"
	"testing"
	"time"

	"github.com/kombolewis/kompani/utils"
	"github.com/stretchr/testify/require"
)

func createNewUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Name:     utils.RandomName(),
		Email:    utils.RandomEmail(),
		Password: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createNewUser(t)
}
func TestGetUser(t *testing.T) {
	user1 := createNewUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

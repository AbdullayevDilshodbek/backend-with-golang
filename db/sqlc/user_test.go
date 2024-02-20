package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AbdullayevDilshodbek/backend-with-golang/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Role:     util.RandomRule(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}


func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID: user1.ID,
		Username: user1.Username,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
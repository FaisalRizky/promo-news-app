package db

import (
	"context"
	"github/promo-news-app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUsersParams{
		Email:             util.RandomEmail(),
		Name:              util.RandomString(10),
		Username:          util.RandomString(8),
		Password:          util.RandomString(100),
		PasswordChangedAt: util.RandomTime(),
		PhoneNumber:       util.RandomInt(800000000, 899999999),
		DeviceToken:       util.RandomString(100),
		Lang:              "id",
		Avatar:            "null",
		UserLevel:         "9",
		IsActive:          util.RandomBoolean(),
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.DeviceToken, user.DeviceToken)
	require.Equal(t, arg.Lang, user.Lang)
	require.Equal(t, arg.Avatar, user.Avatar)
	require.Equal(t, arg.UserLevel, user.UserLevel)
	require.Equal(t, arg.IsActive, user.IsActive)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateRandomUser(t *testing.T) {
	createRandomUser(t)
}

func TestToogleUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := ToogleActiveUsersParams{
		ID:       user1.ID,
		IsActive: !user1.IsActive,
	}
	user1, err := testQueries.ToogleActiveUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user1.IsActive, arg.IsActive)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUsers(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	require.Equal(t, user1.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, user1.DeviceToken, user2.DeviceToken)
	require.Equal(t, user1.Lang, user2.Lang)
	require.Equal(t, user1.Avatar, user2.Avatar)
	require.Equal(t, user1.UserLevel, user2.UserLevel)
	require.Equal(t, user1.IsActive, user2.IsActive)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUsersParams{
		Email:             user1.Email,
		Name:              user1.Name,
		Username:          util.RandomString(8),
		Password:          util.RandomString(100),
		PasswordChangedAt: util.RandomTime(),
		PhoneNumber:       util.RandomInt(800000000, 899999999),
		DeviceToken:       util.RandomString(100),
		Lang:              "id",
		Avatar:            "null",
		UserLevel:         "9",
		IsActive:          util.RandomBoolean(),
	}

	user2, err := testQueries.UpdateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	user, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, user, 5)

	for _, user := range user {
		require.NotEmpty(t, user)
	}
}

package db

import (
	"context"
	"github/promo-news-app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStore(t *testing.T) Store {
	operationalTime := createRandomOperationalTime(t)
	arg := CreateStoresParams{
		Name:          util.RandomString(10),
		Address:       util.RandomString(25),
		Description:   util.RandomString(150),
		PhoneNumber:   util.RandomInt(800000000, 899999999),
		OperationalID: operationalTime.ID,
		IsActive:      util.RandomBoolean(),
	}

	store, err := testQueries.CreateStores(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, arg.Name, store.Name)
	require.Equal(t, arg.Address, store.Address)
	require.Equal(t, arg.Description, store.Description)
	require.Equal(t, arg.PhoneNumber, store.PhoneNumber)
	require.Equal(t, arg.OperationalID, store.OperationalID)
	require.Equal(t, arg.IsActive, store.IsActive)

	require.NotZero(t, store.ID)
	require.NotZero(t, store.CreatedAt)

	return store
}

func TestCreateRandomStore(t *testing.T) {
	createRandomStore(t)
}

func TestGetStore(t *testing.T) {
	store1 := createRandomStore(t)
	store2, err := testQueries.GetStores(context.Background(), store1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, store2)

	require.Equal(t, store1.Name, store2.Name)
	require.Equal(t, store1.Address, store2.Address)
	require.Equal(t, store1.Description, store2.Description)
	require.Equal(t, store1.PhoneNumber, store2.PhoneNumber)
	require.Equal(t, store1.OperationalID, store2.OperationalID)
	require.Equal(t, store1.IsActive, store2.IsActive)
	require.WithinDuration(t, store1.CreatedAt, store2.CreatedAt, time.Second)
}

func TestUpdateStore(t *testing.T) {
	operationalTime := createRandomOperationalTime(t)
	store1 := createRandomStore(t)
	arg := UpdateStoresParams{
		ID:            store1.ID,
		Name:          store1.Name,
		Address:       store1.Address,
		Description:   store1.Description,
		PhoneNumber:   util.RandomInt(81313131, 899999999),
		OperationalID: operationalTime.ID,
		IsActive:      store1.IsActive,
	}

	store2, err := testQueries.UpdateStores(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store2)

	require.Equal(t, arg.Name, store2.Name)
	require.Equal(t, store1.Address, store2.Address)
	require.Equal(t, store1.Description, store2.Description)
	require.Equal(t, store1.IsActive, store2.IsActive)
	require.WithinDuration(t, store1.CreatedAt, store2.CreatedAt, time.Second)
}
func TestToogleStores(t *testing.T) {
	user1 := createRandomStore(t)
	arg := ToogleActiveStoresParams{
		ID:       user1.ID,
		IsActive: !user1.IsActive,
	}
	user1, err := testQueries.ToogleActiveStores(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user1.IsActive, arg.IsActive)
}

func TestListStore(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomStore(t)
	}

	arg := ListStoresParams{
		Limit:  5,
		Offset: 5,
	}

	store, err := testQueries.ListStores(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, store, 5)

	for _, store := range store {
		require.NotEmpty(t, store)
	}
}

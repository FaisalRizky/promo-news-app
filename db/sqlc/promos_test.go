package db

// import (
// 	"context"
// 	"github/promo-news-app/util"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// func createRandomPromo(t *testing.T) Promo {
// 	arg := CreatePromosParams{
// 		PromoName:        util.RandomString(15),
// 		StoreID:          1,
// 		PromoCode:        util.RandomString(5),
// 		PromoDescription: util.RandomString(150),
// 		Quantity:         util.RandomInt(10, 1000),
// 		StartAt:          util.RandomUnixTime(),
// 		ExpiredAt:        util.RandomUnixTime(),
// 		IsActive:         util.RandomBoolean(),
// 		CreatedBy:        1,
// 	}

// 	promo, err := testQueries.CreatePromos(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, promo)

// 	require.Equal(t, arg.PromoName, promo.PromoName)
// 	require.Equal(t, arg.PromoCode, promo.PromoCode)
// 	require.Equal(t, arg.StoreID, promo.StoreID)
// 	require.Equal(t, arg.PromoDescription, promo.PromoDescription)
// 	require.Equal(t, arg.Quantity, promo.Quantity)
// 	require.Equal(t, arg.StartAt, promo.StartAt)
// 	require.Equal(t, arg.ExpiredAt, promo.ExpiredAt)
// 	require.Equal(t, arg.CreatedBy, promo.CreatedBy)
// 	require.Equal(t, arg.IsActive, promo.IsActive)

// 	require.NotZero(t, promo.ID)
// 	require.NotZero(t, promo.CreatedAt)

// 	return promo
// }

// func TestCreateRandomPromo(t *testing.T) {
// 	createRandomPromo(t)
// }

// func TestGetPromo(t *testing.T) {
// 	promo1 := createRandomPromo(t)
// 	promo2, err := testQueries.GetPromos(context.Background(), promo1.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, promo2)

// 	require.Equal(t, promo1.PromoName, promo2.PromoName)
// 	require.Equal(t, promo1.PromoCode, promo2.PromoCode)
// 	require.Equal(t, promo1.StoreID, promo2.StoreID)
// 	require.Equal(t, promo1.PromoDescription, promo2.PromoDescription)
// 	require.Equal(t, promo1.Quantity, promo2.Quantity)
// 	require.Equal(t, promo1.StartAt, promo2.StartAt)
// 	require.Equal(t, promo1.ExpiredAt, promo2.ExpiredAt)
// 	require.Equal(t, promo1.CreatedBy, promo2.CreatedBy)
// 	require.Equal(t, promo1.IsActive, promo2.IsActive)
// 	require.WithinDuration(t, promo1.CreatedAt, promo2.CreatedAt, time.Second)
// }

// func TestUpdatePromo(t *testing.T) {
// 	promo1 := createRandomPromo(t)
// 	arg := UpdatePromosParams{
// 		ID:               promo1.ID,
// 		StoreID:          1,
// 		PromoCode:        promo1.PromoCode,
// 		PromoDescription: util.RandomString(150),
// 		Quantity:         promo1.Quantity,
// 		StartAt:          promo1.StartAt,
// 		ExpiredAt:        promo1.ExpiredAt,
// 		IsActive:         promo1.IsActive,
// 		CreatedBy:        1,
// 	}
// 	promo2, err := testQueries.UpdatePromos(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, promo2)

// 	require.Equal(t, arg.PromoName, promo2.PromoName)
// 	require.Equal(t, promo1.PromoCode, promo2.PromoCode)
// 	require.Equal(t, promo1.StoreID, promo2.StoreID)
// 	require.Equal(t, promo1.Quantity, promo2.Quantity)
// 	require.Equal(t, promo1.StartAt, promo2.StartAt)
// 	require.Equal(t, promo1.ExpiredAt, promo2.ExpiredAt)
// 	require.Equal(t, promo1.CreatedBy, promo2.CreatedBy)
// 	require.Equal(t, promo1.IsActive, promo2.IsActive)
// 	require.WithinDuration(t, promo1.CreatedAt, promo2.CreatedAt, time.Second)
// }

// func TestListPromo(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomPromo(t)
// 	}

// 	arg := ListPromosParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	promo, err := testQueries.ListPromos(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, promo, 5)

// 	for _, promo := range promo {
// 		require.NotEmpty(t, promo)
// 	}
// }

// func TestTooglePromos(t *testing.T) {
// 	user1 := createRandomPromo(t)
// 	arg := ToogleActivePromosParams{
// 		ID:       user1.ID,
// 		IsActive: !user1.IsActive,
// 	}
// 	user1, err := testQueries.ToogleActivePromos(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Equal(t, user1.IsActive, arg.IsActive)
// }

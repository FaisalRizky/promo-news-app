package db

import (
	"context"
	"github/promo-news-app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomOperationalTime(t *testing.T) OperationalTime {
	arg := CreateOperationalTimeParams{
		OpeningTime:     util.RandomAmHours(),
		ClosingTime:     util.RandomPmHours(),
		OperationalDays: util.RandomDay(),
		OffDays:         "12:00-13:00",
		IsActive:        true,
	}

	operationalTime, err := testQueries.CreateOperationalTime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, operationalTime)

	require.Equal(t, arg.OpeningTime, operationalTime.OpeningTime)
	require.Equal(t, arg.ClosingTime, operationalTime.ClosingTime)
	require.Equal(t, arg.OperationalDays, operationalTime.OperationalDays)
	require.Equal(t, arg.IsActive, operationalTime.IsActive)

	require.NotZero(t, operationalTime.ID)
	require.NotZero(t, operationalTime.CreatedAt)

	return operationalTime
}

func TestCreateRandomOperationalTime(t *testing.T) {
	createRandomOperationalTime(t)
}

func TestGetOperationalTime(t *testing.T) {
	operationalTime1 := createRandomOperationalTime(t)
	operationalTime2, err := testQueries.GetOperationalTime(context.Background(), operationalTime1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, operationalTime2)

	require.Equal(t, operationalTime1.OpeningTime, operationalTime2.OpeningTime)
	require.Equal(t, operationalTime1.ClosingTime, operationalTime2.ClosingTime)
	require.Equal(t, operationalTime1.OperationalDays, operationalTime2.OperationalDays)
	require.Equal(t, operationalTime1.OffDays, operationalTime2.OffDays)
	require.Equal(t, operationalTime1.IsActive, operationalTime2.IsActive)
	require.WithinDuration(t, operationalTime1.CreatedAt, operationalTime2.CreatedAt, time.Second)
}

func TestUpdateOperationalTime(t *testing.T) {
	operationalTime1 := createRandomOperationalTime(t)
	arg := UpdateOperationalTimeParams{
		ID:              operationalTime1.ID,
		OpeningTime:     util.RandomAmHours(),
		ClosingTime:     operationalTime1.ClosingTime,
		OperationalDays: operationalTime1.OperationalDays,
		OffDays:         operationalTime1.OffDays,
		IsActive:        operationalTime1.IsActive,
	}
	operationalTime2, err := testQueries.UpdateOperationalTime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, operationalTime2)

	require.Equal(t, arg.OpeningTime, operationalTime2.OpeningTime)
	require.Equal(t, operationalTime1.ClosingTime, operationalTime2.ClosingTime)
	require.Equal(t, operationalTime1.OperationalDays, operationalTime2.OperationalDays)
	require.Equal(t, operationalTime1.OffDays, operationalTime2.OffDays)
	require.Equal(t, operationalTime1.IsActive, operationalTime2.IsActive)
	require.WithinDuration(t, operationalTime1.CreatedAt, operationalTime2.CreatedAt, time.Second)
}

func TestOperationalTime(t *testing.T) {
	operationalTime1 := createRandomOperationalTime(t)
	arg := ToogleActiveOperationalTimeParams{
		ID:       operationalTime1.ID,
		IsActive: !operationalTime1.IsActive,
	}
	operationalTime1, err := testQueries.ToogleActiveOperationalTime(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, operationalTime1.IsActive, arg.IsActive)
}

func TestListOperationalTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOperationalTime(t)
	}

	arg := ListOperationalTimeParams{
		Limit:  5,
		Offset: 5,
	}

	operationalTime, err := testQueries.ListOperationalTime(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, operationalTime, 5)

	for _, account := range operationalTime {
		require.NotEmpty(t, account)
	}
}

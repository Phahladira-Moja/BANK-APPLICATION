package db

import (
	"context"
	"github.com/phahladira-moja/simple-bank-application/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, account Account, account2 Account) Transfer {

	args := CreateTransferParams{
		FromAccountID: account.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.Amount, transfer.Amount)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account, account2)
}

func TestGetTransfer(t *testing.T) {
	account := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer2.ID)
	require.Equal(t, transfer2.Amount, transfer2.Amount)
	require.Equal(t, transfer2.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer2.FromAccountID, transfer2.FromAccountID)
	require.WithinDuration(t, transfer2.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account, account2)
	}

	args := ListTransfersParams{
		FromAccountID: account.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
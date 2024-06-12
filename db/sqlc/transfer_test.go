package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/danielHieu/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1, account2 Account) Transfer{
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	CreateRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := CreateRandomTransfer(t, account1, account2) 
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)	

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt,time.Second)

}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
  for i:= 0; i < 10; i++ {
		CreateRandomTransfer(t,account1,account2)
    CreateRandomTransfer(t,account2,account1)
  }

  arg := ListTransfersParams{
  	FromAccountID: account1.ID,
  	ToAccountID:   account1.ID,
  	Limit:         5,
  	Offset:        5,
  }

  transfers, err := testQueries.ListTransfers(context.Background(), arg)
  require.NoError(t, err)
  require.Len(t, transfers, 5)

  for _, transfer := range transfers {
    require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
  }
  
}

func TestUpdateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
  transfer1 := CreateRandomTransfer(t, account1, account2)

  arg := UpdateTransferParams{
  	ID:     transfer1.ID,
  	Amount: util.RandomMoney(),
  }

  transfer2, err := testQueries.UpdateTransfer(context.Background(), arg )
  require.NoError(t, err)
  require.NotEmpty(t, transfer2)

  require.Equal(t, transfer1.ID, transfer2.ID)
  require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
  require.Equal(t, arg.Amount, transfer2.Amount)
  require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)

  require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
  transfer1 := CreateRandomTransfer(t, account1, account2)

  err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
  require.NoError(t, err)

  transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
  require.Error(t, err)
  require.EqualError(t, err, sql.ErrNoRows.Error()) 
  require.Empty(t, transfer2)
}
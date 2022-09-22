package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	fmt.Printf(">> Before fromAccount = %v, toAccount = %v\n", fromAccount.Balance, toAccount.Balance)
	// run a concurrent goroutine 
	n := 5
	errChannel := make(chan error)
	resultChannel := make(chan TransferTxResult)
	amount := int64(3)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})

			errChannel <- err
			resultChannel <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errChannel
		require.NoError(t, err)

		res := <-resultChannel
		require.NotEmpty(t, res)

		// Check Transfer
		transfer := res.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, fromAccount.ID)
		require.Equal(t, transfer.ToAccountID, toAccount.ID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.ID)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check FromEntry
		fromEntry := res.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.Amount, -amount)
		require.NotEmpty(t, fromEntry.CreatedAt)
		require.NotEmpty(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		// Check ToEntry
		toEntry := res.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.Amount, amount)
		require.NotEmpty(t, toEntry.CreatedAt)
		require.NotEmpty(t, toEntry.ID)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		account1 := res.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		account2 := res.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		fmt.Printf(">> After fromAccount = %v, toAccount = %v\n", account1.Balance, account2.Balance)
		// Check account's balance
		diff1 := fromAccount.Balance - account1.Balance
		diff2 := account2.Balance - toAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err := store.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)
	updatedAccount2, err := store.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	require.Equal(t, fromAccount.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, toAccount.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxForDeadLock(t *testing.T) {
	store := NewStore(testDB)
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	fmt.Printf(">> Before fromAccount = %v, toAccount = %v\n", fromAccount.Balance, toAccount.Balance)
	// run a concurrent goroutine
	n := 10
	errChannel := make(chan error)
	// resultChannel := make(chan TransferTxResult)
	amount := int64(3)

	for i := 0; i < n; i++ {
		var toAID, fromAID int64
		if i%2 == 0 {
			toAID = toAccount.ID
			fromAID = fromAccount.ID
		} else {
			toAID = fromAccount.ID
			fromAID = toAccount.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAID,
				ToAccountID:   toAID,
				Amount:        amount,
			})

			errChannel <- err
			// resultChannel <- result
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errChannel
		require.NoError(t, err)

		// res := <-resultChannel
	}

	updatedAccount1, err := store.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)
	updatedAccount2, err := store.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	require.Equal(t, fromAccount.Balance, updatedAccount1.Balance)
	require.Equal(t, toAccount.Balance, updatedAccount2.Balance)
}

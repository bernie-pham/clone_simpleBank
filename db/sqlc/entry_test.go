package db

import (
	"context"
	"testing"
	"time"

	"github.com/bernie-pham/cloneSimpleBank/ultilities"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	entryArg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    ultilities.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), entryArg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entryArg.Amount, entry.Amount)
	require.Equal(t, entryArg.AccountID, entry.AccountID)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	var lastEntry Entry
	account := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		lastEntry = createRandomEntry(t, account)
	}

	require.NotEmpty(t, lastEntry)
	listEntryArg := ListEntryParams{
		AccountID: account.ID,
		Offset:    0,
		Limit:     10,
	}
	entries, err := testQueries.ListEntry(context.Background(), listEntryArg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}

}

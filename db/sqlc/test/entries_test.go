package test

import (
	"context"
	"database/sql"
	"github.com/marko/simplebank/db/sqlc"
	"github.com/marko/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T, account db.Account) db.Entry {
	arg := db.CreateEntryParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
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
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	newAmount := util.RandomMoney()

	arg := db.UpdateEntryParams{
		ID:     entry1.ID,
		Amount: newAmount,
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, newAmount, entry2.Amount)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	limit := 5
	offset := 0

	entries, err := testQueries.ListEntries(context.Background(), db.ListEntriesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	require.NoError(t, err)
	require.Len(t, entries, limit)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

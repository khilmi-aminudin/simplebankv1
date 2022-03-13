package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotZero(t, account.ID)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccounts(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)

	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	n := 5
	for i := 0; i < n*2; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  int32(n),
		Offset: int32(n),
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, n)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: 100,
	}

	account1, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, arg.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, arg.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account1)

}

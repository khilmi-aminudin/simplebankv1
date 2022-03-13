package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, account)

	assert.NotZero(t, account.ID)
	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Currency, account.Currency)

	assert.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccounts(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, account1)

	assert.Equal(t, account.ID, account1.ID)
	assert.Equal(t, account.Owner, account1.Owner)
	assert.Equal(t, account.Balance, account1.Balance)
	assert.Equal(t, account.Currency, account1.Currency)

	assert.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
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
	assert.NoError(t, err)
	assert.Len(t, accounts, n)

	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: 100,
	}

	account1, err := testQueries.UpdateAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, account1)

	assert.Equal(t, arg.ID, account1.ID)
	assert.Equal(t, account.Owner, account1.Owner)
	assert.Equal(t, arg.Balance, account1.Balance)
	assert.Equal(t, account.Currency, account1.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	assert.NoError(t, err)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, account1)

}

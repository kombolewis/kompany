package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/kombolewis/kompani/utils"
	"github.com/stretchr/testify/require"
)

func createNewCompany(t *testing.T) Company {
	arg := CreateCompanyParams{
		Name:        utils.RandomName(),
		Description: sql.NullString{String: "Industrial Cleaning Manufacturing Company", Valid: true},
		Amount:      utils.RandomAmount(),
		Registered:  true,
		Type:        utils.RandomType(),
	}

	company, err := testQueries.CreateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, arg.Name, company.Name)
	// require.Equal(t, arg.Description, company.Description)
	require.Equal(t, arg.Amount, company.Amount)
	require.Equal(t, arg.Registered, company.Registered)
	require.Equal(t, arg.Type, company.Type)

	require.NotZero(t, company.ID)
	require.NotZero(t, company.CreatedAt)

	return company
}

func TestCreateCompany(t *testing.T) {
	createNewCompany(t)
}
func TestGetCompany(t *testing.T) {
	company1 := createNewCompany(t)
	company2, err := testQueries.GetCompany(context.Background(), company1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, company2)

	require.Equal(t, company1.Name, company2.Name)
	require.Equal(t, company1.Description, company2.Description)
	require.Equal(t, company1.Amount, company2.Amount)
	require.Equal(t, company1.Registered, company2.Registered)
	require.Equal(t, company1.Type, company2.Type)
	require.WithinDuration(t, company1.CreatedAt, company2.CreatedAt, time.Second)

}

func TestUpdateCompany(t *testing.T) {
	company := createNewCompany(t)

	arg := UpdateCompanyParams{
		ID:          company.ID,
		Name:        utils.RandomName(),
		Description: company.Description,
		Amount:      utils.RandomAmount(),
		Registered:  company.Registered,
		Type:        utils.RandomType(),
	}

	updatedCompany, err := testQueries.UpdateCompany(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedCompany)

	require.Equal(t, arg.Name, updatedCompany.Name)
	require.Equal(t, company.Description, company.Description)
	require.Equal(t, arg.Amount, updatedCompany.Amount)
	require.Equal(t, company.Registered, updatedCompany.Registered)
	require.Equal(t, arg.Type, updatedCompany.Type)
	require.WithinDuration(t, company.CreatedAt, updatedCompany.CreatedAt, time.Second)

}

func TestDeleteCompany(t *testing.T) {
	company := createNewCompany(t)
	err := testQueries.DeleteCompany(context.Background(), company.ID)
	require.NoError(t, err)

	deletedCompany, err := testQueries.GetCompany(context.Background(), company.ID)
	require.Empty(t, deletedCompany)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListCompanies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createNewCompany(t)
	}

	arg := ListCompaniesParams{
		Limit:  5,
		Offset: 5,
	}

	companies, err := testQueries.ListCompanies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 5)

	for _, company := range companies {
		require.NotEmpty(t, company)
	}
}

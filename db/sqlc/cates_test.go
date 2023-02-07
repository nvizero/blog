package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCate(t *testing.T) Cate {
	arg := CreateCateParams{
		Name: "aa11",
	}
	cate, err := testQueries.CreateCate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cate)

	require.Equal(t, arg.Name, cate.Name)

	require.NotZero(t, cate.ID)
	require.NotZero(t, cate.CreatedAt)
	return cate
}

func TestCreateCate(t *testing.T) {
	createRandomCate(t)
}

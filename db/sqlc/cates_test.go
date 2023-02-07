package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCate(t *testing.T) Cate {
	//arg := CreateCateParams{
	//	Name: "aaqq1",
	//}
	cate, err := testQueries.CreateCate(context.Background(), "asd")
	require.NotEmpty(t, cate)
	require.NoError(t, err)
	//require.Equal(t, arg.Name, cate.Name)

	return cate
}

func TestCreateCate(t *testing.T) {
	createRandomCate(t)
}

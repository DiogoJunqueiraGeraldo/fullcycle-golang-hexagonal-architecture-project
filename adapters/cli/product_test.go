package cli_test

import (
	"fmt"
	"testing"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/adapters/cli"
	mock_application "github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Product Test"
	productPrice := 30.0
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := fmt.Sprintf(
		`Product %s { "id": "%s", "name": "%s", "price": %f, "status": "%s" }`,
		"create",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf(
		`Product %s { "id": "%s", "name": "%s", "price": %f, "status": "%s" }`,
		"enable",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "enable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf(
		`Product %s { "id": "%s", "name": "%s", "price": %f, "status": "%s" }`,
		"disable",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "disable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf(
		`Product %s { "id": "%s", "name": "%s", "price": %f, "status": "%s" }`,
		"",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}

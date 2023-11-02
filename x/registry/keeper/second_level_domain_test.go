package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSecondLevelDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecondLevelDomain {
	items := make([]types.SecondLevelDomain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)
		items[i].Parent = strconv.Itoa(i)

		keeper.SetSecondLevelDomain(ctx, items[i])
	}
	return items
}

func createNSecondLevelDomainResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecondLevelDomainResponse {
	items := createNSecondLevelDomain(keeper, ctx, n)
	responses := make([]types.SecondLevelDomainResponse, n)
	for i := range responses {
		responses[i].Name = items[i].Name
		responses[i].Parent = items[i].Parent
		responses[i].ExpirationDate = items[i].ExpirationDate
	}
	return responses
}

func TestSecondLevelDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSecondLevelDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		_, found := keeper.GetSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.False(t, found)
	}
}

func TestDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecondLevelDomain(ctx)),
	)
}

func (suite *KeeperTestSuite) TestGetValidSecondLevelDomain() {
	testCases := []struct {
		secondLevelDomain types.SecondLevelDomain
		expErr            error
	}{
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "cel",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, 20).UnixNano(),
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "cel",
				ExpirationDate: 0,
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, -20).UnixNano(),
			},
			expErr: errorsmod.Wrapf(types.ErrDomainNotFound, "test"),
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Set domain
			suite.app.RegistryKeeper.SetSecondLevelDomain(suite.ctx, tc.secondLevelDomain)

			// Get valid domain
			secondLevelDomain, err := suite.app.RegistryKeeper.GetValidSecondLevelDomain(suite.ctx, tc.secondLevelDomain.Name, tc.secondLevelDomain.Parent)
			if tc.expErr == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.secondLevelDomain, secondLevelDomain)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.expErr.Error(), err.Error())
			}
		})
	}

}

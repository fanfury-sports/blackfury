package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	blacktypes "github.com/blackfury-zone/blackfury/types"
	"github.com/blackfury-zone/blackfury/x/gauge/types"
)

func (suite *KeeperTestSuite) TestBribe_Deposit() {
	suite.SetupTest()
	k := suite.app.GaugeKeeper
	depoistDenom := blacktypes.BaseDenom
	k.SetGauge(suite.ctx, depoistDenom)
	bribe := k.Bribe(suite.ctx, depoistDenom)
	veID := uint64(100)
	amount := sdk.NewInt(100)
	bribe.Deposit(suite.ctx, veID, amount)

	suite.Require().Equal(amount, bribe.GetTotalDepositedAmount(suite.ctx))
	suite.Require().Equal(amount, bribe.GetDepositedAmountByUser(suite.ctx, veID))
}

func (suite *KeeperTestSuite) TestBribe_Withdraw() {
	suite.SetupTest()
	k := suite.app.GaugeKeeper
	depoistDenom := blacktypes.BaseDenom
	k.SetGauge(suite.ctx, depoistDenom)
	bribe := k.Bribe(suite.ctx, depoistDenom)
	veID := uint64(100)
	amount := sdk.NewInt(100)
	err := bribe.Withdraw(suite.ctx, veID, amount)
	suite.Require().Error(err, types.ErrTooLargeAmount)

	bribe.Deposit(suite.ctx, veID, amount)
	err = bribe.Withdraw(suite.ctx, veID, sdk.NewInt(50))
	suite.Require().Nil(err)

	suite.Require().Equal(sdk.NewInt(50), bribe.GetTotalDepositedAmount(suite.ctx))
	suite.Require().Equal(sdk.NewInt(50), bribe.GetDepositedAmountByUser(suite.ctx, veID))

	err = bribe.Withdraw(suite.ctx, veID, sdk.NewInt(50))
	suite.Require().Nil(err)

	suite.Require().Equal(sdk.ZeroInt(), bribe.GetTotalDepositedAmount(suite.ctx))
	suite.Require().Equal(sdk.ZeroInt(), bribe.GetDepositedAmountByUser(suite.ctx, veID))
}

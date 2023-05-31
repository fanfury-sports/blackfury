package keeper_test

import (
	"fmt"

	blacktypes "github.com/blackfury-zone/blackfury/types"
	"github.com/blackfury-zone/blackfury/x/gauge/keeper"
	"github.com/blackfury-zone/blackfury/x/gauge/types"
)

func (suite *KeeperTestSuite) TestBase_PoolDenom() {
	suite.SetupTest()
	k := suite.app.GaugeKeeper
	depoistDenom := blacktypes.BaseDenom
	base := keeper.NewBase(k, depoistDenom, types.GaugeKey(depoistDenom), true)
	suite.Require().Equal(depoistDenom, base.PoolDenom())
}

func (suite *KeeperTestSuite) TestBase_PoolName() {
	suite.SetupTest()
	k := suite.app.GaugeKeeper
	depoistDenom := blacktypes.BaseDenom
	testCases := []struct {
		name    string
		gauge   bool
		pooName string
	}{
		{"bribe pool name", false, fmt.Sprintf("bribe_%s", depoistDenom)},
		{"gauge pool name", true, fmt.Sprintf("gauge_%s", depoistDenom)},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			base := keeper.NewBase(k, depoistDenom, types.GaugeKey(depoistDenom), tc.gauge)
			suite.Require().Equal(tc.pooName, base.PoolName())
		})
	}
}

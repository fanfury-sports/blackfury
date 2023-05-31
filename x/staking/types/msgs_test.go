package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blackfury-zone/blackfury/app"
	blacktypes "github.com/blackfury-zone/blackfury/types"
	"github.com/blackfury-zone/blackfury/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestMsgVeDelegate_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc             string
		delegatorAddress string
		validatorAddress string
		veID             string
		amount           sdk.Coin
		valid            bool
	}{
		{
			desc:             "ErrEmptyDelegatorAddr",
			delegatorAddress: "",
		},
		{
			desc:             "ErrEmptyValidatorAddr",
			delegatorAddress: "black1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5",
			validatorAddress: "",
		},
		{
			desc:             "invalid ve id",
			delegatorAddress: "black1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5",
			validatorAddress: "blackvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pctrjr3",
			veID:             "",
		},
		{
			desc:             "invalid delegation amount",
			delegatorAddress: "black1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5",
			validatorAddress: "blackvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pctrjr3",
			veID:             "ve-100",
			amount:           sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(0)),
		},
		{
			desc:             "valid",
			delegatorAddress: "black1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5",
			validatorAddress: "blackvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pctrjr3",
			veID:             "ve-100",
			amount:           sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
			valid:            true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgVeDelegate{
				DelegatorAddress: tc.delegatorAddress,
				ValidatorAddress: tc.validatorAddress,
				VeId:             tc.veID,
				Amount:           tc.amount,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgVeDelegate_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgVeDelegate{
		DelegatorAddress: "black1mnfm9c7cdgqnkk66sganp78m0ydmcr4ppeaeg5",
		ValidatorAddress: "blackvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pctrjr3",
		Amount:           sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
		VeId:             "ve-100",
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

package types

import (
	"testing"

	blackfury "github.com/fanfury-sports/blackfury/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	params := DefaultParams()
	require.Equal(t, blackfury.BaseDenom, params.LockDenom)
}

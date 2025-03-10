package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// AddressHandler is an interface that defines the methods to handle addresses
type AddressHandler interface {

	// GetSheAddressFromString parses an address string and returns the corresponding sdk.AccAddress.
	// Address string does not have to be a bech32 address. It could be a 0x prefixed (EVM) address, etc.
	GetSheAddressFromString(ctx sdk.Context, address string) (sdk.AccAddress, error)
}

type SheAddressHandler struct{}

// GetSheAddressFromString parses a bech32 address formatted string and returns the corresponding sdk.AccAddress
func (h SheAddressHandler) GetSheAddressFromString(_ sdk.Context, address string) (sdk.AccAddress, error) {
	parsedAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, err
	}
	return parsedAddress, nil
}

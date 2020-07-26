package asserter

import (
	"errors"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Coin returns an error if the provided *types.Coin is invalid.
func Coin(coin *types.Coin) error {
	if coin == nil {
		return errors.New("coin cannot be nil")
	}

	if err := CoinIdentifier(coin.CoinIdentifier); err != nil {
		return fmt.Errorf("%w: coin identifier is invalid", err)
	}

	if err := Amount(coin.Amount); err != nil {
		return fmt.Errorf("%w: coin amount is invalid", err)
	}

	return nil
}

// Coins returns an error if the provided
// []*types.Coin is invalid. If there are any
// duplicate identifiers, this function
// will also return an error.
func Coins(coins []*types.Coin) error {
	ids := map[string]struct{}{}
	for _, coin := range coins {
		if err := Coin(coin); err != nil {
			return fmt.Errorf("%w: coin is invalid", err)
		}

		if _, exists := ids[coin.CoinIdentifier.Identifier]; exists {
			return fmt.Errorf("duplicate coin identifier detected: %s", coin.CoinIdentifier.Identifier)
		}

		ids[coin.CoinIdentifier.Identifier] = struct{}{}
	}

	return nil
}

// CoinIdentifier returns an error if the provided *types.CoinIdentifier
// is invalid.
func CoinIdentifier(coinIdentifier *types.CoinIdentifier) error {
	if coinIdentifier == nil {
		return errors.New("coin identifier cannot be nil")
	}

	if len(coinIdentifier.Identifier) == 0 {
		return errors.New("coin identifier cannot be empty")
	}

	return nil
}

// CoinChange returns an error if the provided *types.CoinChange
// is invalid.
func CoinChange(change *types.CoinChange) error {
	if change == nil {
		return errors.New("coin change cannot be nil")
	}

	if err := CoinIdentifier(change.CoinIdentifier); err != nil {
		return fmt.Errorf("%w: coin identifier is invalid", err)
	}

	if err := CoinAction(change.CoinAction); err != nil {
		return fmt.Errorf("%w: coin action is invalid", err)
	}

	return nil
}

// CoinAction returns an error if the provided types.CoinAction
// is invalid.
func CoinAction(action types.CoinAction) error {
	switch action {
	case types.CoinCreated, types.CoinSpent:
	default:
		return fmt.Errorf("%s is not a valid CoinAction", action)
	}

	return nil
}

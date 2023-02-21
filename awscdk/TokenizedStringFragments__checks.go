//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (t *jsiiProxy_TokenizedStringFragments) validateAddIntrinsicParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateAddLiteralParameters(lit interface{}) error {
	if lit == nil {
		return fmt.Errorf("parameter lit is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateAddTokenParameters(token IResolvable) error {
	if token == nil {
		return fmt.Errorf("parameter token is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateJoinParameters(concat IFragmentConcatenator) error {
	if concat == nil {
		return fmt.Errorf("parameter concat is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TokenizedStringFragments) validateMapTokensParameters(mapper ITokenMapper) error {
	if mapper == nil {
		return fmt.Errorf("parameter mapper is required, but nil was provided")
	}

	return nil
}


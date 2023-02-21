//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateToken_AsAnyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateToken_AsListParameters(value interface{}, options *EncodingOptions) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateToken_AsNumberParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateToken_AsStringParameters(value interface{}, options *EncodingOptions) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateToken_CompareStringsParameters(possibleToken1 *string, possibleToken2 *string) error {
	if possibleToken1 == nil {
		return fmt.Errorf("parameter possibleToken1 is required, but nil was provided")
	}

	if possibleToken2 == nil {
		return fmt.Errorf("parameter possibleToken2 is required, but nil was provided")
	}

	return nil
}

func validateToken_IsUnresolvedParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}


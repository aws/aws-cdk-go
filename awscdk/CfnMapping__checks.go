//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CfnMapping) validateFindInMapParameters(key1 *string, key2 *string) error {
	if key1 == nil {
		return fmt.Errorf("parameter key1 is required, but nil was provided")
	}

	if key2 == nil {
		return fmt.Errorf("parameter key2 is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnMapping) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnMapping) validateSetValueParameters(key1 *string, key2 *string, value interface{}) error {
	if key1 == nil {
		return fmt.Errorf("parameter key1 is required, but nil was provided")
	}

	if key2 == nil {
		return fmt.Errorf("parameter key2 is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCfnMapping_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnMapping_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCfnMappingParameters(scope constructs.Construct, id *string, props *CfnMappingProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


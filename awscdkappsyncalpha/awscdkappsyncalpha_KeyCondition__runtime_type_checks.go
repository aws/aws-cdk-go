//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"fmt"
)

func (k *jsiiProxy_KeyCondition) validateAndParameters(keyCond KeyCondition) error {
	if keyCond == nil {
		return fmt.Errorf("parameter keyCond is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_BeginsWithParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_BetweenParameters(keyName *string, arg1 *string, arg2 *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg1 == nil {
		return fmt.Errorf("parameter arg1 is required, but nil was provided")
	}

	if arg2 == nil {
		return fmt.Errorf("parameter arg2 is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_EqParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_GeParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_GtParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_LeParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

func validateKeyCondition_LtParameters(keyName *string, arg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}


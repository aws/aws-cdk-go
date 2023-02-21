//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FeatureFlags) validateIsEnabledParameters(featureFlag *string) error {
	if featureFlag == nil {
		return fmt.Errorf("parameter featureFlag is required, but nil was provided")
	}

	return nil
}

func validateFeatureFlags_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}


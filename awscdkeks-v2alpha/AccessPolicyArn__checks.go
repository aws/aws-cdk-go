//go:build !no_runtime_type_checking

package awscdkeks-v2alpha

import (
	"fmt"
)

func validateAccessPolicyArn_OfParameters(policyName *string) error {
	if policyName == nil {
		return fmt.Errorf("parameter policyName is required, but nil was provided")
	}

	return nil
}

func validateNewAccessPolicyArnParameters(policyName *string) error {
	if policyName == nil {
		return fmt.Errorf("parameter policyName is required, but nil was provided")
	}

	return nil
}


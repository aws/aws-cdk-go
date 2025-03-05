//go:build !no_runtime_type_checking

package awscdkeks-v2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAccessPolicy_FromAccessPolicyNameParameters(policyName *string, options *AccessPolicyNameOptions) error {
	if policyName == nil {
		return fmt.Errorf("parameter policyName is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewAccessPolicyParameters(props *AccessPolicyProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


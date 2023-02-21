//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (g *jsiiProxy_Grant) validateCombineParameters(rhs Grant) error {
	if rhs == nil {
		return fmt.Errorf("parameter rhs is required, but nil was provided")
	}

	return nil
}

func validateGrant_AddToPrincipalParameters(options *GrantOnPrincipalOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGrant_AddToPrincipalAndResourceParameters(options *GrantOnPrincipalAndResourceOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGrant_AddToPrincipalOrResourceParameters(options *GrantWithResourceOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGrant_DropParameters(grantee IGrantable, _intent *string) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if _intent == nil {
		return fmt.Errorf("parameter _intent is required, but nil was provided")
	}

	return nil
}


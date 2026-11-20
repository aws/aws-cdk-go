//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_CfnTemplateMetadataContext) validateSetParameters(context *CfnTemplateMetadataContextProps) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}

func validateCfnTemplateMetadataContext_OfParameters(stack Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}


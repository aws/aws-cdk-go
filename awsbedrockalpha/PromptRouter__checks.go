//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (p *jsiiProxy_PromptRouter) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validatePromptRouter_FromDefaultIdParameters(defaultRouter DefaultPromptRouterIdentifier, region *string) error {
	if defaultRouter == nil {
		return fmt.Errorf("parameter defaultRouter is required, but nil was provided")
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}

func validateNewPromptRouterParameters(props *PromptRouterProps, region *string) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	return nil
}


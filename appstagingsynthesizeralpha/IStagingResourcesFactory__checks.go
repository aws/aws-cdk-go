//go:build !no_runtime_type_checking

package appstagingsynthesizeralpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_IStagingResourcesFactory) validateObtainStagingResourcesParameters(stack awscdk.Stack, context *ObtainStagingResourcesContext) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}


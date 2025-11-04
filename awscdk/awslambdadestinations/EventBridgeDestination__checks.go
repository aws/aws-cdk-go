//go:build !no_runtime_type_checking

package awslambdadestinations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EventBridgeDestination) validateBindParameters(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if fn == nil {
		return fmt.Errorf("parameter fn is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


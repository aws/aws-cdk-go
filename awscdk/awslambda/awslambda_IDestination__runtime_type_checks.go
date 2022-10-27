//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_IDestination) validateBindParameters(scope awscdk.Construct, fn IFunction, options *DestinationOptions) error {
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


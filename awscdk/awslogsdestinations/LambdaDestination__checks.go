//go:build !no_runtime_type_checking

package awslogsdestinations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaDestination) validateBindParameters(scope constructs.Construct, sourceLogGroup awslogs.ILogGroup) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if sourceLogGroup == nil {
		return fmt.Errorf("parameter sourceLogGroup is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaDestinationParameters(fn awslambda.IFunction, options *LambdaDestinationOptions) error {
	if fn == nil {
		return fmt.Errorf("parameter fn is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


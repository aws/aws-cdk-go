//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (l *jsiiProxy_LambdaIntegration) validateBindParameters(method Method) error {
	if method == nil {
		return fmt.Errorf("parameter method is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaIntegrationParameters(handler awslambda.IFunction, options *LambdaIntegrationOptions) error {
	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


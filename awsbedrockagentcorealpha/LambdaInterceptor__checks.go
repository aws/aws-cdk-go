//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaInterceptor) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}

func validateLambdaInterceptor_ForRequestParameters(lambdaFunction awslambda.IFunction, options *InterceptorOptions) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateLambdaInterceptor_ForResponseParameters(lambdaFunction awslambda.IFunction, options *InterceptorOptions) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


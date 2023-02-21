//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaFunctionProcessor) validateBindParameters(_scope constructs.Construct, options *DataProcessorBindOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewLambdaFunctionProcessorParameters(lambdaFunction awslambda.IFunction, props *DataProcessorProps) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


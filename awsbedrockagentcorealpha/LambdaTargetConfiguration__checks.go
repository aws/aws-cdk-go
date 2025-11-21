//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaTargetConfiguration) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}

func validateLambdaTargetConfiguration_CreateParameters(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if toolSchema == nil {
		return fmt.Errorf("parameter toolSchema is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaTargetConfigurationParameters(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if toolSchema == nil {
		return fmt.Errorf("parameter toolSchema is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaAction) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaActionParameters(lambdaFunction interface{}) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}
	switch lambdaFunction.(type) {
	case awslambda.IFunction:
		// ok
	case awslambda.IVersion:
		// ok
	case awslambda.IAlias:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(lambdaFunction) {
			return fmt.Errorf("parameter lambdaFunction must be one of the allowed types: awslambda.IFunction, awslambda.IVersion, awslambda.IAlias; received %#v (a %T)", lambdaFunction, lambdaFunction)
		}
	}

	return nil
}


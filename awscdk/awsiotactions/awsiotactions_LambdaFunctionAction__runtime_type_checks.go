//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

func (l *jsiiProxy_LambdaFunctionAction) validateBindParameters(topicRule awsiot.ITopicRule) error {
	if topicRule == nil {
		return fmt.Errorf("parameter topicRule is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaFunctionActionParameters(func_ awslambda.IFunction) error {
	if func_ == nil {
		return fmt.Errorf("parameter func_ is required, but nil was provided")
	}

	return nil
}


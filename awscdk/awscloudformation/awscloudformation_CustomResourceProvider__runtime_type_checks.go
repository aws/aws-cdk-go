//go:build !no_runtime_type_checking

package awscloudformation

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

func (c *jsiiProxy_CustomResourceProvider) validateBindParameters(_arg awscdk.Construct) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func validateCustomResourceProvider_FromLambdaParameters(handler awslambda.IFunction) error {
	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	return nil
}

func validateCustomResourceProvider_FromTopicParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

func validateCustomResourceProvider_LambdaParameters(handler awslambda.IFunction) error {
	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	return nil
}

func validateCustomResourceProvider_TopicParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}


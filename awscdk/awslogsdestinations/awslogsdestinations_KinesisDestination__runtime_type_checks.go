//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awslogsdestinations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

func (k *jsiiProxy_KinesisDestination) validateBindParameters(scope awscdk.Construct, _sourceLogGroup awslogs.ILogGroup) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _sourceLogGroup == nil {
		return fmt.Errorf("parameter _sourceLogGroup is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisDestinationParameters(stream awskinesis.IStream, props *KinesisDestinationProps) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


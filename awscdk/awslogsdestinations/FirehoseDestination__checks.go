//go:build !no_runtime_type_checking

package awslogsdestinations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FirehoseDestination) validateBindParameters(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _sourceLogGroup == nil {
		return fmt.Errorf("parameter _sourceLogGroup is required, but nil was provided")
	}

	return nil
}

func validateNewFirehoseDestinationParameters(stream awskinesisfirehose.IDeliveryStream, props *FirehoseDestinationProps) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


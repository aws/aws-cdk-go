//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(stage IStageRef) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}

func validateNewFirehoseLogDestinationParameters(stream awskinesisfirehose.CfnDeliveryStream) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}


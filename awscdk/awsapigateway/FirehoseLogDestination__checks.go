//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(_stage IStage) error {
	if _stage == nil {
		return fmt.Errorf("parameter _stage is required, but nil was provided")
	}

	return nil
}

func validateNewFirehoseLogDestinationParameters(stream awskinesisfirehose.CfnDeliveryStream) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}


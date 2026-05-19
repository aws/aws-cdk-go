//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateLoggingDestination_CloudWatchLogsParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

func validateLoggingDestination_FirehoseParameters(stream awskinesisfirehose.IDeliveryStream) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}

func validateLoggingDestination_S3Parameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


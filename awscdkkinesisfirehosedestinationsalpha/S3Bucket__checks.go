//go:build !no_runtime_type_checking

package awscdkkinesisfirehosedestinationsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_S3Bucket) validateBindParameters(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateNewS3BucketParameters(bucket awss3.IBucket, props *S3BucketProps) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


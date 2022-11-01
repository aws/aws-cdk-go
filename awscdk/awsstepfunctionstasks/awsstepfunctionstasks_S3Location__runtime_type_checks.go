//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

func (s *jsiiProxy_S3Location) validateBindParameters(task ISageMakerTask, opts *S3LocationBindOptions) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateS3Location_FromBucketParameters(bucket awss3.IBucket, keyPrefix *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if keyPrefix == nil {
		return fmt.Errorf("parameter keyPrefix is required, but nil was provided")
	}

	return nil
}

func validateS3Location_FromJsonExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}


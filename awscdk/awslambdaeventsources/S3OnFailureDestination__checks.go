//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func (s *jsiiProxy_S3OnFailureDestination) validateBindParameters(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	if _target == nil {
		return fmt.Errorf("parameter _target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

func validateNewS3OnFailureDestinationParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateS3OutputDestination_ToBucketParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

func validateS3OutputDestination_UrlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}


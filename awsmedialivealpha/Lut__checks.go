//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateLut_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateLut_UrlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_IBucketNotificationDestination) validateBindParameters(scope awscdk.Construct, bucket IBucket) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


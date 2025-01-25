//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func (b *jsiiProxy_BucketWebsiteTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewBucketWebsiteTargetParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

func (b *jsiiProxy_BucketWebsiteTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateNewBucketWebsiteTargetParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


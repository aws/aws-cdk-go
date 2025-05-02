//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

func validateEndpoint_FromKinesisStreamParameters(stream awskinesis.IStream) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}


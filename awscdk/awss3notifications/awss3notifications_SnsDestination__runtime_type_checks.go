//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awss3notifications

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SnsDestination) validateBindParameters(_scope constructs.Construct, bucket awss3.IBucket) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

func validateNewSnsDestinationParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}


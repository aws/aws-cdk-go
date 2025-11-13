//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IBucketNotificationDestination) validateBindParameters(scope constructs.Construct, bucket interfacesawss3.IBucketRef) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


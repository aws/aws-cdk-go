//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FirehoseLogDestination) validateGrantPushParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewFirehoseLogDestinationParameters(deliveryStream awskinesisfirehose.IDeliveryStream) error {
	if deliveryStream == nil {
		return fmt.Errorf("parameter deliveryStream is required, but nil was provided")
	}

	return nil
}


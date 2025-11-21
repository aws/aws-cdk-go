//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
)

func (d *jsiiProxy_DeliveryStreamGrants) validatePutRecordsParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateDeliveryStreamGrants_FromDeliveryStreamParameters(resource interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}


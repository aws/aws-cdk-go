//go:build !no_runtime_type_checking

package previewawsm2mixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (c *jsiiProxy_CfnApplicationConsoleLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnApplicationConsoleLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	if deliveryStream == nil {
		return fmt.Errorf("parameter deliveryStream is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnApplicationConsoleLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnApplicationConsoleLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


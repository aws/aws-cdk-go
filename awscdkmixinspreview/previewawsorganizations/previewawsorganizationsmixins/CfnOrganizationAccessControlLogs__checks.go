//go:build !no_runtime_type_checking

package previewawsorganizationsmixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (c *jsiiProxy_CfnOrganizationAccessControlLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	if deliveryStream == nil {
		return fmt.Errorf("parameter deliveryStream is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnOrganizationAccessControlLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnOrganizationAccessControlLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}


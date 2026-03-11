//go:build !no_runtime_type_checking

package previewawseksmixins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps) error {
	if deliveryStream == nil {
		return fmt.Errorf("parameter deliveryStream is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


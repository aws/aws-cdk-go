package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnEventBusLogsMixin to generate INFO_LOGS for CfnEventBus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusInfoLogs := awscdkmixinspreview.Mixins.NewCfnEventBusInfoLogs()
//
type CfnEventBusInfoLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusInfoLogsDestProps) CfnEventBusLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusInfoLogsFirehoseProps) CfnEventBusLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusInfoLogsLogGroupProps) CfnEventBusLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnEventBusInfoLogsS3Props) CfnEventBusLogsMixin
}

// The jsii proxy struct for CfnEventBusInfoLogs
type jsiiProxy_CfnEventBusInfoLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnEventBusInfoLogs() CfnEventBusInfoLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusInfoLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEventBusInfoLogs_Override(c CfnEventBusInfoLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnEventBusInfoLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusInfoLogsDestProps) CfnEventBusLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnEventBusLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventBusInfoLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusInfoLogsFirehoseProps) CfnEventBusLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnEventBusLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventBusInfoLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusInfoLogsLogGroupProps) CfnEventBusLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnEventBusLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventBusInfoLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnEventBusInfoLogsS3Props) CfnEventBusLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnEventBusLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


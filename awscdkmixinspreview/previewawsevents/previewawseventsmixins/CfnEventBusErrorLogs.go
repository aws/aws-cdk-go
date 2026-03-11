package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnEventBusLogsMixin to generate ERROR_LOGS for CfnEventBus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusErrorLogs := awscdkmixinspreview.Mixins.NewCfnEventBusErrorLogs()
//
type CfnEventBusErrorLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusErrorLogsDestProps) CfnEventBusLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusErrorLogsFirehoseProps) CfnEventBusLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusErrorLogsLogGroupProps) CfnEventBusLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnEventBusErrorLogsS3Props) CfnEventBusLogsMixin
}

// The jsii proxy struct for CfnEventBusErrorLogs
type jsiiProxy_CfnEventBusErrorLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnEventBusErrorLogs() CfnEventBusErrorLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusErrorLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEventBusErrorLogs_Override(c CfnEventBusErrorLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnEventBusErrorLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusErrorLogsDestProps) CfnEventBusLogsMixin {
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

func (c *jsiiProxy_CfnEventBusErrorLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusErrorLogsFirehoseProps) CfnEventBusLogsMixin {
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

func (c *jsiiProxy_CfnEventBusErrorLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusErrorLogsLogGroupProps) CfnEventBusLogsMixin {
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

func (c *jsiiProxy_CfnEventBusErrorLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnEventBusErrorLogsS3Props) CfnEventBusLogsMixin {
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


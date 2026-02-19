package previewawsmediapackagev2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnChannelGroupLogsMixin to generate EGRESS_ACCESS_LOGS for CfnChannelGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupEgressAccessLogs := awscdkmixinspreview.Mixins.NewCfnChannelGroupEgressAccessLogs()
//
type CfnChannelGroupEgressAccessLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnChannelGroupLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnChannelGroupLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnChannelGroupLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnChannelGroupLogsMixin
}

// The jsii proxy struct for CfnChannelGroupEgressAccessLogs
type jsiiProxy_CfnChannelGroupEgressAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnChannelGroupEgressAccessLogs() CfnChannelGroupEgressAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnChannelGroupEgressAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnChannelGroupEgressAccessLogs_Override(c CfnChannelGroupEgressAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnChannelGroupLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnChannelGroupLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnChannelGroupLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnChannelGroupLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnChannelGroupLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnChannelGroupLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnChannelGroupLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnChannelGroupLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


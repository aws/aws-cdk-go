package previewawsmediapackagev2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnChannelGroupLogsMixin to generate INGRESS_ACCESS_LOGS for CfnChannelGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupIngressAccessLogs := awscdkmixinspreview.Mixins.NewCfnChannelGroupIngressAccessLogs()
//
type CfnChannelGroupIngressAccessLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnChannelGroupLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnChannelGroupLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnChannelGroupLogsMixin
}

// The jsii proxy struct for CfnChannelGroupIngressAccessLogs
type jsiiProxy_CfnChannelGroupIngressAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnChannelGroupIngressAccessLogs() CfnChannelGroupIngressAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnChannelGroupIngressAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnChannelGroupIngressAccessLogs_Override(c CfnChannelGroupIngressAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnChannelGroupLogsMixin {
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

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnChannelGroupLogsMixin {
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

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnChannelGroupLogsMixin {
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


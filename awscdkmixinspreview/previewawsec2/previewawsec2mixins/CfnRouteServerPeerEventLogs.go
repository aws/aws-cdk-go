package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnRouteServerPeerLogsMixin to generate EVENT_LOGS for CfnRouteServerPeer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPeerEventLogs := awscdkmixinspreview.Mixins.NewCfnRouteServerPeerEventLogs()
//
type CfnRouteServerPeerEventLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnRouteServerPeerLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnRouteServerPeerLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnRouteServerPeerLogsMixin
}

// The jsii proxy struct for CfnRouteServerPeerEventLogs
type jsiiProxy_CfnRouteServerPeerEventLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnRouteServerPeerEventLogs() CfnRouteServerPeerEventLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnRouteServerPeerEventLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRouteServerPeerEventLogs_Override(c CfnRouteServerPeerEventLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnRouteServerPeerEventLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnRouteServerPeerLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnRouteServerPeerLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRouteServerPeerEventLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnRouteServerPeerLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnRouteServerPeerLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRouteServerPeerEventLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnRouteServerPeerLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnRouteServerPeerLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


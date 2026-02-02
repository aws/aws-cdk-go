package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnVPNConnectionLogsMixin to generate CONNECTION_LOGS for CfnVPNConnection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionConnectionLogs := awscdkmixinspreview.Mixins.NewCfnVPNConnectionConnectionLogs()
//
type CfnVPNConnectionConnectionLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVPNConnectionLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVPNConnectionLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnVPNConnectionLogsMixin
}

// The jsii proxy struct for CfnVPNConnectionConnectionLogs
type jsiiProxy_CfnVPNConnectionConnectionLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnVPNConnectionConnectionLogs() CfnVPNConnectionConnectionLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnVPNConnectionConnectionLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVPNConnectionConnectionLogs_Override(c CfnVPNConnectionConnectionLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnVPNConnectionConnectionLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVPNConnectionLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnVPNConnectionLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnectionConnectionLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVPNConnectionLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnVPNConnectionLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPNConnectionConnectionLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnVPNConnectionLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnVPNConnectionLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


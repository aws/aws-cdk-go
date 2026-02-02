package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnInstanceLogsMixin to generate AMAZON_CONNECT_FLOW_LOGS for CfnInstance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAmazonConnectFlowLogs := awscdkmixinspreview.Mixins.NewCfnInstanceAmazonConnectFlowLogs()
//
type CfnInstanceAmazonConnectFlowLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnInstanceLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnInstanceLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnInstanceLogsMixin
}

// The jsii proxy struct for CfnInstanceAmazonConnectFlowLogs
type jsiiProxy_CfnInstanceAmazonConnectFlowLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnInstanceAmazonConnectFlowLogs() CfnInstanceAmazonConnectFlowLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceAmazonConnectFlowLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnInstanceAmazonConnectFlowLogs_Override(c CfnInstanceAmazonConnectFlowLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnInstanceLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnInstanceLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnInstanceLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


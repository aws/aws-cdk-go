package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
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
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceAmazonConnectFlowLogsDestProps) CfnInstanceLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceAmazonConnectFlowLogsFirehoseProps) CfnInstanceLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceAmazonConnectFlowLogsLogGroupProps) CfnInstanceLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnInstanceAmazonConnectFlowLogsS3Props) CfnInstanceLogsMixin
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

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceAmazonConnectFlowLogsDestProps) CfnInstanceLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceAmazonConnectFlowLogsFirehoseProps) CfnInstanceLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceAmazonConnectFlowLogsLogGroupProps) CfnInstanceLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnInstanceAmazonConnectFlowLogsS3Props) CfnInstanceLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
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


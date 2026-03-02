package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnFlowLogsMixin to generate APPLICATION_LOGS for CfnFlow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowApplicationLogs := awscdkmixinspreview.Mixins.NewCfnFlowApplicationLogs()
//
type CfnFlowApplicationLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFlowApplicationLogsDestProps) CfnFlowLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFlowApplicationLogsFirehoseProps) CfnFlowLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnFlowApplicationLogsLogGroupProps) CfnFlowLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnFlowApplicationLogsS3Props) CfnFlowLogsMixin
}

// The jsii proxy struct for CfnFlowApplicationLogs
type jsiiProxy_CfnFlowApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnFlowApplicationLogs() CfnFlowApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnFlowApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFlowApplicationLogs_Override(c CfnFlowApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnFlowApplicationLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFlowApplicationLogsDestProps) CfnFlowLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnFlowLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFlowApplicationLogsFirehoseProps) CfnFlowLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnFlowLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnFlowApplicationLogsLogGroupProps) CfnFlowLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnFlowLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnFlowApplicationLogsS3Props) CfnFlowLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnFlowLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnRestApiLogsMixin to generate EXECUTION_LOGS for CfnRestApi.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiExecutionLogs := awscdkmixinspreview.Mixins.NewCfnRestApiExecutionLogs()
//
type CfnRestApiExecutionLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRestApiExecutionLogsDestProps) CfnRestApiLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRestApiExecutionLogsFirehoseProps) CfnRestApiLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnRestApiExecutionLogsLogGroupProps) CfnRestApiLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnRestApiExecutionLogsS3Props) CfnRestApiLogsMixin
}

// The jsii proxy struct for CfnRestApiExecutionLogs
type jsiiProxy_CfnRestApiExecutionLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnRestApiExecutionLogs() CfnRestApiExecutionLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnRestApiExecutionLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRestApiExecutionLogs_Override(c CfnRestApiExecutionLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRestApiExecutionLogsDestProps) CfnRestApiLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnRestApiLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRestApiExecutionLogsFirehoseProps) CfnRestApiLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnRestApiLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnRestApiExecutionLogsLogGroupProps) CfnRestApiLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnRestApiLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnRestApiExecutionLogsS3Props) CfnRestApiLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnRestApiLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


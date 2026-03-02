package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnRuntimeLogsMixin to generate APPLICATION_LOGS for CfnRuntime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeApplicationLogs := awscdkmixinspreview.Mixins.NewCfnRuntimeApplicationLogs()
//
type CfnRuntimeApplicationLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeApplicationLogsDestProps) CfnRuntimeLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRuntimeApplicationLogsFirehoseProps) CfnRuntimeLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnRuntimeApplicationLogsLogGroupProps) CfnRuntimeLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnRuntimeApplicationLogsS3Props) CfnRuntimeLogsMixin
}

// The jsii proxy struct for CfnRuntimeApplicationLogs
type jsiiProxy_CfnRuntimeApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnRuntimeApplicationLogs() CfnRuntimeApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnRuntimeApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRuntimeApplicationLogs_Override(c CfnRuntimeApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeApplicationLogsDestProps) CfnRuntimeLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRuntimeApplicationLogsFirehoseProps) CfnRuntimeLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnRuntimeApplicationLogsLogGroupProps) CfnRuntimeLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnRuntimeApplicationLogsS3Props) CfnRuntimeLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


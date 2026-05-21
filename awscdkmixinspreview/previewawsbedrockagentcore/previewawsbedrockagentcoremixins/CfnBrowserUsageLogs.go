package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnBrowserLogsMixin to generate USAGE_LOGS for CfnBrowser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserUsageLogs := awscdkmixinspreview.Mixins.NewCfnBrowserUsageLogs()
//
type CfnBrowserUsageLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserUsageLogsDestProps) CfnBrowserLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserUsageLogsFirehoseProps) CfnBrowserLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserUsageLogsLogGroupProps) CfnBrowserLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnBrowserUsageLogsS3Props) CfnBrowserLogsMixin
}

// The jsii proxy struct for CfnBrowserUsageLogs
type jsiiProxy_CfnBrowserUsageLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserUsageLogs() CfnBrowserUsageLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserUsageLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserUsageLogs_Override(c CfnBrowserUsageLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnBrowserUsageLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserUsageLogsDestProps) CfnBrowserLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserUsageLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserUsageLogsFirehoseProps) CfnBrowserLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserUsageLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserUsageLogsLogGroupProps) CfnBrowserLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserUsageLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnBrowserUsageLogsS3Props) CfnBrowserLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


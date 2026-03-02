package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnBrowserCustomLogsMixin to generate USAGE_LOGS for CfnBrowserCustom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogs := awscdkmixinspreview.Mixins.NewCfnBrowserCustomUsageLogs()
//
type CfnBrowserCustomUsageLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomUsageLogsDestProps) CfnBrowserCustomLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserCustomUsageLogsFirehoseProps) CfnBrowserCustomLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserCustomUsageLogsLogGroupProps) CfnBrowserCustomLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnBrowserCustomUsageLogsS3Props) CfnBrowserCustomLogsMixin
}

// The jsii proxy struct for CfnBrowserCustomUsageLogs
type jsiiProxy_CfnBrowserCustomUsageLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserCustomUsageLogs() CfnBrowserCustomUsageLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserCustomUsageLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserCustomUsageLogs_Override(c CfnBrowserCustomUsageLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomUsageLogsDestProps) CfnBrowserCustomLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserCustomUsageLogsFirehoseProps) CfnBrowserCustomLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserCustomUsageLogsLogGroupProps) CfnBrowserCustomLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnBrowserCustomUsageLogsS3Props) CfnBrowserCustomLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


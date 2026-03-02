package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCodeInterpreterCustomLogsMixin to generate USAGE_LOGS for CfnCodeInterpreterCustom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomUsageLogs := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterCustomUsageLogs()
//
type CfnCodeInterpreterCustomUsageLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterCustomUsageLogsDestProps) CfnCodeInterpreterCustomLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterCustomUsageLogsFirehoseProps) CfnCodeInterpreterCustomLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterCustomUsageLogsLogGroupProps) CfnCodeInterpreterCustomLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterCustomUsageLogsS3Props) CfnCodeInterpreterCustomLogsMixin
}

// The jsii proxy struct for CfnCodeInterpreterCustomUsageLogs
type jsiiProxy_CfnCodeInterpreterCustomUsageLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterCustomUsageLogs() CfnCodeInterpreterCustomUsageLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterCustomUsageLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterCustomUsageLogs_Override(c CfnCodeInterpreterCustomUsageLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterCustomUsageLogsDestProps) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterCustomUsageLogsFirehoseProps) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterCustomUsageLogsLogGroupProps) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterCustomUsageLogsS3Props) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


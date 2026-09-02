package previewawsrtbfabricmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnLinkLogsMixin to generate APPLICATION_LOGS for CfnLink.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkApplicationLogs := awscdkmixinspreview.Mixins.NewCfnLinkApplicationLogs()
//
type CfnLinkApplicationLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLinkApplicationLogsDestProps) CfnLinkLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLinkApplicationLogsFirehoseProps) CfnLinkLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLinkApplicationLogsLogGroupProps) CfnLinkLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnLinkApplicationLogsS3Props) CfnLinkLogsMixin
}

// The jsii proxy struct for CfnLinkApplicationLogs
type jsiiProxy_CfnLinkApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnLinkApplicationLogs() CfnLinkApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnLinkApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLinkApplicationLogs_Override(c CfnLinkApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnLinkApplicationLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLinkApplicationLogsDestProps) CfnLinkLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnLinkLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLinkApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLinkApplicationLogsFirehoseProps) CfnLinkLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnLinkLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLinkApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLinkApplicationLogsLogGroupProps) CfnLinkLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnLinkLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLinkApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnLinkApplicationLogsS3Props) CfnLinkLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnLinkLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


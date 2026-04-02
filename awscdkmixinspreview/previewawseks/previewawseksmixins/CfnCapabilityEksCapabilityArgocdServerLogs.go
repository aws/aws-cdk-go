package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCapabilityLogsMixin to generate EKS_CAPABILITY_ARGOCD_SERVER_LOGS for CfnCapability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdServerLogs := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdServerLogs()
//
type CfnCapabilityEksCapabilityArgocdServerLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdServerLogsDestProps) CfnCapabilityLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps) CfnCapabilityLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdServerLogsLogGroupProps) CfnCapabilityLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdServerLogsS3Props) CfnCapabilityLogsMixin
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdServerLogs
type jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdServerLogs() CfnCapabilityEksCapabilityArgocdServerLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdServerLogs_Override(c CfnCapabilityEksCapabilityArgocdServerLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdServerLogsDestProps) CfnCapabilityLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnCapabilityLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps) CfnCapabilityLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnCapabilityLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdServerLogsLogGroupProps) CfnCapabilityLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnCapabilityLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdServerLogsS3Props) CfnCapabilityLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnCapabilityLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


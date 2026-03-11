package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCapabilityLogsMixin to generate EKS_CAPABILITY_ARGOCD_APPLICATION_LOGS for CfnCapability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationLogs := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdApplicationLogs()
//
type CfnCapabilityEksCapabilityArgocdApplicationLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps) CfnCapabilityLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps) CfnCapabilityLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps) CfnCapabilityLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props) CfnCapabilityLogsMixin
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdApplicationLogs
type jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationLogs() CfnCapabilityEksCapabilityArgocdApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationLogs_Override(c CfnCapabilityEksCapabilityArgocdApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props) CfnCapabilityLogsMixin {
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


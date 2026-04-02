package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCapabilityLogsMixin to generate EKS_CAPABILITY_ARGOCD_APPLICATIONSET_LOGS for CfnCapability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationsetLogs := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdApplicationsetLogs()
//
type CfnCapabilityEksCapabilityArgocdApplicationsetLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps) CfnCapabilityLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsFirehoseProps) CfnCapabilityLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsLogGroupProps) CfnCapabilityLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsS3Props) CfnCapabilityLogsMixin
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdApplicationsetLogs
type jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationsetLogs() CfnCapabilityEksCapabilityArgocdApplicationsetLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationsetLogs_Override(c CfnCapabilityEksCapabilityArgocdApplicationsetLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsFirehoseProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsLogGroupProps) CfnCapabilityLogsMixin {
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

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationsetLogsS3Props) CfnCapabilityLogsMixin {
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


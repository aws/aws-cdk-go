package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnMailManagerIngressPointLogsMixin to generate TRAFFIC_POLICY_DEBUG_LOGS for CfnMailManagerIngressPoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointTrafficPolicyDebugLogs := awscdkmixinspreview.Mixins.NewCfnMailManagerIngressPointTrafficPolicyDebugLogs()
//
type CfnMailManagerIngressPointTrafficPolicyDebugLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps) CfnMailManagerIngressPointLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps) CfnMailManagerIngressPointLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps) CfnMailManagerIngressPointLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props) CfnMailManagerIngressPointLogsMixin
}

// The jsii proxy struct for CfnMailManagerIngressPointTrafficPolicyDebugLogs
type jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerIngressPointTrafficPolicyDebugLogs() CfnMailManagerIngressPointTrafficPolicyDebugLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerIngressPointTrafficPolicyDebugLogs_Override(c CfnMailManagerIngressPointTrafficPolicyDebugLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnLoadBalancerLogsMixin to generate ALB_ACCESS_LOGS for CfnLoadBalancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbAccessLogs := awscdkmixinspreview.Mixins.NewCfnLoadBalancerAlbAccessLogs()
//
type CfnLoadBalancerAlbAccessLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbAccessLogsDestProps) CfnLoadBalancerLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbAccessLogsFirehoseProps) CfnLoadBalancerLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbAccessLogsLogGroupProps) CfnLoadBalancerLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbAccessLogsS3Props) CfnLoadBalancerLogsMixin
}

// The jsii proxy struct for CfnLoadBalancerAlbAccessLogs
type jsiiProxy_CfnLoadBalancerAlbAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerAlbAccessLogs() CfnLoadBalancerAlbAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerAlbAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerAlbAccessLogs_Override(c CfnLoadBalancerAlbAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbAccessLogsDestProps) CfnLoadBalancerLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnLoadBalancerLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbAccessLogsFirehoseProps) CfnLoadBalancerLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnLoadBalancerLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbAccessLogsLogGroupProps) CfnLoadBalancerLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnLoadBalancerLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbAccessLogsS3Props) CfnLoadBalancerLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnLoadBalancerLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnLoadBalancerLogsMixin to generate ALB_HEALTH_CHECK_LOGS for CfnLoadBalancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbHealthCheckLogs := awscdkmixinspreview.Mixins.NewCfnLoadBalancerAlbHealthCheckLogs()
//
type CfnLoadBalancerAlbHealthCheckLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbHealthCheckLogsDestProps) CfnLoadBalancerLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbHealthCheckLogsFirehoseProps) CfnLoadBalancerLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbHealthCheckLogsLogGroupProps) CfnLoadBalancerLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbHealthCheckLogsS3Props) CfnLoadBalancerLogsMixin
}

// The jsii proxy struct for CfnLoadBalancerAlbHealthCheckLogs
type jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerAlbHealthCheckLogs() CfnLoadBalancerAlbHealthCheckLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerAlbHealthCheckLogs_Override(c CfnLoadBalancerAlbHealthCheckLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbHealthCheckLogsDestProps) CfnLoadBalancerLogsMixin {
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

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbHealthCheckLogsFirehoseProps) CfnLoadBalancerLogsMixin {
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

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbHealthCheckLogsLogGroupProps) CfnLoadBalancerLogsMixin {
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

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbHealthCheckLogsS3Props) CfnLoadBalancerLogsMixin {
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


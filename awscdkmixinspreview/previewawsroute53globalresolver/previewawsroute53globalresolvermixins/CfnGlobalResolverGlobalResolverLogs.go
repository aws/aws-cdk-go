package previewawsroute53globalresolvermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnGlobalResolverLogsMixin to generate GLOBAL_RESOLVER_LOGS for CfnGlobalResolver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalResolverGlobalResolverLogs := awscdkmixinspreview.Mixins.NewCfnGlobalResolverGlobalResolverLogs()
//
type CfnGlobalResolverGlobalResolverLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGlobalResolverGlobalResolverLogsDestProps) CfnGlobalResolverLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnGlobalResolverGlobalResolverLogsFirehoseProps) CfnGlobalResolverLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnGlobalResolverGlobalResolverLogsLogGroupProps) CfnGlobalResolverLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnGlobalResolverGlobalResolverLogsS3Props) CfnGlobalResolverLogsMixin
}

// The jsii proxy struct for CfnGlobalResolverGlobalResolverLogs
type jsiiProxy_CfnGlobalResolverGlobalResolverLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnGlobalResolverGlobalResolverLogs() CfnGlobalResolverGlobalResolverLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnGlobalResolverGlobalResolverLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGlobalResolverGlobalResolverLogs_Override(c CfnGlobalResolverGlobalResolverLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGlobalResolverGlobalResolverLogsDestProps) CfnGlobalResolverLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnGlobalResolverLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnGlobalResolverGlobalResolverLogsFirehoseProps) CfnGlobalResolverLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnGlobalResolverLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnGlobalResolverGlobalResolverLogsLogGroupProps) CfnGlobalResolverLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnGlobalResolverLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnGlobalResolverGlobalResolverLogsS3Props) CfnGlobalResolverLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnGlobalResolverLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


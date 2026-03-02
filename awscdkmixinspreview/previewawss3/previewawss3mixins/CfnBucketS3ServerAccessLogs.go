package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnBucketLogsMixin to generate S3_SERVER_ACCESS_LOGS for CfnBucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBucketS3ServerAccessLogs := awscdkmixinspreview.Mixins.NewCfnBucketS3ServerAccessLogs()
//
type CfnBucketS3ServerAccessLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBucketS3ServerAccessLogsDestProps) CfnBucketLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBucketS3ServerAccessLogsFirehoseProps) CfnBucketLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBucketS3ServerAccessLogsLogGroupProps) CfnBucketLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnBucketS3ServerAccessLogsS3Props) CfnBucketLogsMixin
}

// The jsii proxy struct for CfnBucketS3ServerAccessLogs
type jsiiProxy_CfnBucketS3ServerAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnBucketS3ServerAccessLogs() CfnBucketS3ServerAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnBucketS3ServerAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBucketS3ServerAccessLogs_Override(c CfnBucketS3ServerAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBucketS3ServerAccessLogsDestProps) CfnBucketLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnBucketLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBucketS3ServerAccessLogsFirehoseProps) CfnBucketLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnBucketLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnBucketS3ServerAccessLogsLogGroupProps) CfnBucketLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnBucketLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnBucketS3ServerAccessLogsS3Props) CfnBucketLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnBucketLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnPlaybackConfigurationLogsMixin to generate MANIFEST_SERVICE_LOGS for CfnPlaybackConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationManifestServiceLogs := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationManifestServiceLogs()
//
type CfnPlaybackConfigurationManifestServiceLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnPlaybackConfigurationManifestServiceLogsDestProps) CfnPlaybackConfigurationLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPlaybackConfigurationManifestServiceLogsFirehoseProps) CfnPlaybackConfigurationLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnPlaybackConfigurationManifestServiceLogsLogGroupProps) CfnPlaybackConfigurationLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnPlaybackConfigurationManifestServiceLogsS3Props) CfnPlaybackConfigurationLogsMixin
}

// The jsii proxy struct for CfnPlaybackConfigurationManifestServiceLogs
type jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnPlaybackConfigurationManifestServiceLogs() CfnPlaybackConfigurationManifestServiceLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPlaybackConfigurationManifestServiceLogs_Override(c CfnPlaybackConfigurationManifestServiceLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnPlaybackConfigurationManifestServiceLogsDestProps) CfnPlaybackConfigurationLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnPlaybackConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPlaybackConfigurationManifestServiceLogsFirehoseProps) CfnPlaybackConfigurationLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnPlaybackConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnPlaybackConfigurationManifestServiceLogsLogGroupProps) CfnPlaybackConfigurationLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnPlaybackConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnPlaybackConfigurationManifestServiceLogsS3Props) CfnPlaybackConfigurationLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnPlaybackConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnInstanceLogsMixin to generate CONSOLE_LOGS for CfnInstance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceConsoleLogs := awscdkmixinspreview.Mixins.NewCfnInstanceConsoleLogs()
//
type CfnInstanceConsoleLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceConsoleLogsDestProps) CfnInstanceLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceConsoleLogsFirehoseProps) CfnInstanceLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceConsoleLogsLogGroupProps) CfnInstanceLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnInstanceConsoleLogsS3Props) CfnInstanceLogsMixin
}

// The jsii proxy struct for CfnInstanceConsoleLogs
type jsiiProxy_CfnInstanceConsoleLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnInstanceConsoleLogs() CfnInstanceConsoleLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceConsoleLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnInstanceConsoleLogs_Override(c CfnInstanceConsoleLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceConsoleLogsDestProps) CfnInstanceLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceConsoleLogsFirehoseProps) CfnInstanceLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceConsoleLogsLogGroupProps) CfnInstanceLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnInstanceConsoleLogsS3Props) CfnInstanceLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
		panic(err)
	}
	var returns CfnInstanceLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


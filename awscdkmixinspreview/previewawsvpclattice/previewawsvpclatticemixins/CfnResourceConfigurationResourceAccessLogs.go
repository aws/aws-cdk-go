package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnResourceConfigurationLogsMixin to generate RESOURCE_ACCESS_LOGS for CfnResourceConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceConfigurationResourceAccessLogs := awscdkmixinspreview.Mixins.NewCfnResourceConfigurationResourceAccessLogs()
//
type CfnResourceConfigurationResourceAccessLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnResourceConfigurationLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnResourceConfigurationLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnResourceConfigurationLogsMixin
}

// The jsii proxy struct for CfnResourceConfigurationResourceAccessLogs
type jsiiProxy_CfnResourceConfigurationResourceAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnResourceConfigurationResourceAccessLogs() CfnResourceConfigurationResourceAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceConfigurationResourceAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnResourceConfigurationResourceAccessLogs_Override(c CfnResourceConfigurationResourceAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnResourceConfigurationLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnResourceConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnResourceConfigurationLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnResourceConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnResourceConfigurationLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnResourceConfigurationLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


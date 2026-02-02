package previewawsshieldmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnProtectionLogsMixin to generate FLOW_LOGS for CfnProtection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProtectionFlowLogs := awscdkmixinspreview.Mixins.NewCfnProtectionFlowLogs()
//
type CfnProtectionFlowLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnProtectionLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnProtectionLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnProtectionLogsMixin
}

// The jsii proxy struct for CfnProtectionFlowLogs
type jsiiProxy_CfnProtectionFlowLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnProtectionFlowLogs() CfnProtectionFlowLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnProtectionFlowLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnProtectionFlowLogs_Override(c CfnProtectionFlowLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnProtectionFlowLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnProtectionLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnProtectionLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProtectionFlowLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnProtectionLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnProtectionLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProtectionFlowLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnProtectionLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnProtectionLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


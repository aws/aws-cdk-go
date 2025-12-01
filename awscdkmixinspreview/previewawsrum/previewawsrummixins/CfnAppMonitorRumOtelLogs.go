package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnAppMonitorLogsMixin to generate RUM_OTEL_LOGS for CfnAppMonitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelLogs := awscdkmixinspreview.Mixins.NewCfnAppMonitorRumOtelLogs()
//
type CfnAppMonitorRumOtelLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAppMonitorLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAppMonitorLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnAppMonitorLogsMixin
}

// The jsii proxy struct for CfnAppMonitorRumOtelLogs
type jsiiProxy_CfnAppMonitorRumOtelLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnAppMonitorRumOtelLogs() CfnAppMonitorRumOtelLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnAppMonitorRumOtelLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAppMonitorRumOtelLogs_Override(c CfnAppMonitorRumOtelLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnAppMonitorRumOtelLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAppMonitorLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnAppMonitorLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitorRumOtelLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAppMonitorLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnAppMonitorLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitorRumOtelLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnAppMonitorLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnAppMonitorLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


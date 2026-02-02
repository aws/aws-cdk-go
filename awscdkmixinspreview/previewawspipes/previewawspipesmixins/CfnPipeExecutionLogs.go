package previewawspipesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnPipeLogsMixin to generate EXECUTION_LOGS for CfnPipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipeExecutionLogs := awscdkmixinspreview.Mixins.NewCfnPipeExecutionLogs()
//
type CfnPipeExecutionLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnPipeLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnPipeLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnPipeLogsMixin
}

// The jsii proxy struct for CfnPipeExecutionLogs
type jsiiProxy_CfnPipeExecutionLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnPipeExecutionLogs() CfnPipeExecutionLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeExecutionLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPipeExecutionLogs_Override(c CfnPipeExecutionLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnPipeExecutionLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnPipeLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnPipeLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeExecutionLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnPipeLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnPipeLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeExecutionLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnPipeLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnPipeLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawsosismixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnPipelineLogsMixin to generate PIPELINE_LOGS for CfnPipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipelinePipelineLogs := awscdkmixinspreview.Mixins.NewCfnPipelinePipelineLogs()
//
type CfnPipelinePipelineLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnPipelineLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnPipelineLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnPipelineLogsMixin
}

// The jsii proxy struct for CfnPipelinePipelineLogs
type jsiiProxy_CfnPipelinePipelineLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnPipelinePipelineLogs() CfnPipelinePipelineLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnPipelinePipelineLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPipelinePipelineLogs_Override(c CfnPipelinePipelineLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnPipelineLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnPipelineLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnPipelineLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnPipelineLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnPipelineLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnPipelineLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


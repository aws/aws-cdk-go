package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnKnowledgeBaseLogsMixin to generate RUNTIME_LOGS for CfnKnowledgeBase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseRuntimeLogs := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseRuntimeLogs()
//
type CfnKnowledgeBaseRuntimeLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnKnowledgeBaseLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnKnowledgeBaseLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnKnowledgeBaseLogsMixin
}

// The jsii proxy struct for CfnKnowledgeBaseRuntimeLogs
type jsiiProxy_CfnKnowledgeBaseRuntimeLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnKnowledgeBaseRuntimeLogs() CfnKnowledgeBaseRuntimeLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBaseRuntimeLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnKnowledgeBaseRuntimeLogs_Override(c CfnKnowledgeBaseRuntimeLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnKnowledgeBaseLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnKnowledgeBaseLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnKnowledgeBaseLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnKnowledgeBaseLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnKnowledgeBaseLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnKnowledgeBaseLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


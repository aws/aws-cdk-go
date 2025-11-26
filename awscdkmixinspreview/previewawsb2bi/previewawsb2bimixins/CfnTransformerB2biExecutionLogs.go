package previewawsb2bimixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnTransformerLogsMixin to generate B2BI_EXECUTION_LOGS for CfnTransformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerB2biExecutionLogs := awscdkmixinspreview.Mixins.NewCfnTransformerB2biExecutionLogs()
//
type CfnTransformerB2biExecutionLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnTransformerLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnTransformerLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnTransformerLogsMixin
}

// The jsii proxy struct for CfnTransformerB2biExecutionLogs
type jsiiProxy_CfnTransformerB2biExecutionLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnTransformerB2biExecutionLogs() CfnTransformerB2biExecutionLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnTransformerB2biExecutionLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnTransformerB2biExecutionLogs_Override(c CfnTransformerB2biExecutionLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnTransformerLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnTransformerLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnTransformerLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnTransformerLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnTransformerLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnTransformerLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnAgentLogsMixin to generate APPLICATION_LOGS for CfnAgent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentApplicationLogs := awscdkmixinspreview.Mixins.NewCfnAgentApplicationLogs()
//
type CfnAgentApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAgentLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAgentLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnAgentLogsMixin
}

// The jsii proxy struct for CfnAgentApplicationLogs
type jsiiProxy_CfnAgentApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnAgentApplicationLogs() CfnAgentApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnAgentApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAgentApplicationLogs_Override(c CfnAgentApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnAgentApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAgentLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnAgentLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgentApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAgentLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnAgentLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgentApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnAgentLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnAgentLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnAssistantLogsMixin to generate EVENT_LOGS for CfnAssistant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantEventLogs := awscdkmixinspreview.Mixins.NewCfnAssistantEventLogs()
//
type CfnAssistantEventLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAssistantLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAssistantLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnAssistantLogsMixin
}

// The jsii proxy struct for CfnAssistantEventLogs
type jsiiProxy_CfnAssistantEventLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnAssistantEventLogs() CfnAssistantEventLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnAssistantEventLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAssistantEventLogs_Override(c CfnAssistantEventLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnAssistantEventLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnAssistantLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnAssistantLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssistantEventLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnAssistantLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnAssistantLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssistantEventLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnAssistantLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnAssistantLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


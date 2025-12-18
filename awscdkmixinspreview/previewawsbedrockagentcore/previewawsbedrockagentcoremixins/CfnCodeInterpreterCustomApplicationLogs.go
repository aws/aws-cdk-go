package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCodeInterpreterCustomLogsMixin to generate APPLICATION_LOGS for CfnCodeInterpreterCustom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomApplicationLogs := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterCustomApplicationLogs()
//
type CfnCodeInterpreterCustomApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCodeInterpreterCustomLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCodeInterpreterCustomLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnCodeInterpreterCustomLogsMixin
}

// The jsii proxy struct for CfnCodeInterpreterCustomApplicationLogs
type jsiiProxy_CfnCodeInterpreterCustomApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterCustomApplicationLogs() CfnCodeInterpreterCustomApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterCustomApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterCustomApplicationLogs_Override(c CfnCodeInterpreterCustomApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnCodeInterpreterCustomLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnCodeInterpreterCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


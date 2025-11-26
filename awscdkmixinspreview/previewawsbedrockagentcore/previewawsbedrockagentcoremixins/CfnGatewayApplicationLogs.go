package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnGatewayLogsMixin to generate APPLICATION_LOGS for CfnGateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayApplicationLogs := awscdkmixinspreview.Mixins.NewCfnGatewayApplicationLogs()
//
type CfnGatewayApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnGatewayLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnGatewayLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnGatewayLogsMixin
}

// The jsii proxy struct for CfnGatewayApplicationLogs
type jsiiProxy_CfnGatewayApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnGatewayApplicationLogs() CfnGatewayApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGatewayApplicationLogs_Override(c CfnGatewayApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnGatewayLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnGatewayLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnGatewayLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnGatewayLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnGatewayLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnGatewayLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


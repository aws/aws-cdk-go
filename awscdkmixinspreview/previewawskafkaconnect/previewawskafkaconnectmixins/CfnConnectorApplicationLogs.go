package previewawskafkaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnConnectorLogsMixin to generate APPLICATION_LOGS for CfnConnector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorApplicationLogs := awscdkmixinspreview.Mixins.NewCfnConnectorApplicationLogs()
//
type CfnConnectorApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnConnectorLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnConnectorLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnConnectorLogsMixin
}

// The jsii proxy struct for CfnConnectorApplicationLogs
type jsiiProxy_CfnConnectorApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnConnectorApplicationLogs() CfnConnectorApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnConnectorApplicationLogs_Override(c CfnConnectorApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnConnectorLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnConnectorLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnConnectorLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnConnectorLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnConnectorLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnConnectorLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


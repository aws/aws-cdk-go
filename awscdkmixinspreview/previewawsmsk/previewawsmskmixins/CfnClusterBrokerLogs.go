package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnClusterLogsMixin to generate BROKER_LOGS for CfnCluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterBrokerLogs := awscdkmixinspreview.Mixins.NewCfnClusterBrokerLogs()
//
type CfnClusterBrokerLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnClusterLogsMixin
}

// The jsii proxy struct for CfnClusterBrokerLogs
type jsiiProxy_CfnClusterBrokerLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterBrokerLogs() CfnClusterBrokerLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterBrokerLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterBrokerLogs_Override(c CfnClusterBrokerLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnClusterBrokerLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterBrokerLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterBrokerLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnClusterLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


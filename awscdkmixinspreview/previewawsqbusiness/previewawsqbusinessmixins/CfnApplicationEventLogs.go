package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnApplicationLogsMixin to generate EVENT_LOGS for CfnApplication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEventLogs := awscdkmixinspreview.Mixins.NewCfnApplicationEventLogs()
//
type CfnApplicationEventLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnApplicationLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnApplicationLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnApplicationLogsMixin
}

// The jsii proxy struct for CfnApplicationEventLogs
type jsiiProxy_CfnApplicationEventLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationEventLogs() CfnApplicationEventLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationEventLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationEventLogs_Override(c CfnApplicationEventLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnApplicationEventLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnApplicationLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationEventLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnApplicationLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationEventLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnApplicationLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


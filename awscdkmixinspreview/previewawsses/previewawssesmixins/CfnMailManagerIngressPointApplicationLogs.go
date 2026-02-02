package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnMailManagerIngressPointLogsMixin to generate APPLICATION_LOGS for CfnMailManagerIngressPoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointApplicationLogs := awscdkmixinspreview.Mixins.NewCfnMailManagerIngressPointApplicationLogs()
//
type CfnMailManagerIngressPointApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMailManagerIngressPointLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMailManagerIngressPointLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMailManagerIngressPointLogsMixin
}

// The jsii proxy struct for CfnMailManagerIngressPointApplicationLogs
type jsiiProxy_CfnMailManagerIngressPointApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerIngressPointApplicationLogs() CfnMailManagerIngressPointApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerIngressPointApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerIngressPointApplicationLogs_Override(c CfnMailManagerIngressPointApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMailManagerIngressPointLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMailManagerIngressPointLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


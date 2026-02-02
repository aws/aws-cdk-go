package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnMailManagerRuleSetLogsMixin to generate APPLICATION_LOGS for CfnMailManagerRuleSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerRuleSetApplicationLogs := awscdkmixinspreview.Mixins.NewCfnMailManagerRuleSetApplicationLogs()
//
type CfnMailManagerRuleSetApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMailManagerRuleSetLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMailManagerRuleSetLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMailManagerRuleSetLogsMixin
}

// The jsii proxy struct for CfnMailManagerRuleSetApplicationLogs
type jsiiProxy_CfnMailManagerRuleSetApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerRuleSetApplicationLogs() CfnMailManagerRuleSetApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerRuleSetApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerRuleSetApplicationLogs_Override(c CfnMailManagerRuleSetApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMailManagerRuleSetLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnMailManagerRuleSetLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMailManagerRuleSetLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnMailManagerRuleSetLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMailManagerRuleSetLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMailManagerRuleSetLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


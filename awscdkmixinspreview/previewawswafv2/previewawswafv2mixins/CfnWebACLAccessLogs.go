package previewawswafv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnWebACLLogsMixin to generate ACCESS_LOGS for CfnWebACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLAccessLogs := awscdkmixinspreview.Mixins.NewCfnWebACLAccessLogs()
//
type CfnWebACLAccessLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnWebACLLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnWebACLLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnWebACLLogsMixin
}

// The jsii proxy struct for CfnWebACLAccessLogs
type jsiiProxy_CfnWebACLAccessLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnWebACLAccessLogs() CfnWebACLAccessLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLAccessLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWebACLAccessLogs_Override(c CfnWebACLAccessLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnWebACLAccessLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnWebACLLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnWebACLLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAccessLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnWebACLLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnWebACLLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAccessLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnWebACLLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnWebACLLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


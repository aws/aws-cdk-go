package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnBrowserCustomLogsMixin to generate USAGE_LOGS for CfnBrowserCustom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogs := awscdkmixinspreview.Mixins.NewCfnBrowserCustomUsageLogs()
//
type CfnBrowserCustomUsageLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnBrowserCustomLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnBrowserCustomLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnBrowserCustomLogsMixin
}

// The jsii proxy struct for CfnBrowserCustomUsageLogs
type jsiiProxy_CfnBrowserCustomUsageLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserCustomUsageLogs() CfnBrowserCustomUsageLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserCustomUsageLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserCustomUsageLogs_Override(c CfnBrowserCustomUsageLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnBrowserCustomLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnBrowserCustomLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnBrowserCustomLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


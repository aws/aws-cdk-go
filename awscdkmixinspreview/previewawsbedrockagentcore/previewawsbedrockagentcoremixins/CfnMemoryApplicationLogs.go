package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnMemoryLogsMixin to generate APPLICATION_LOGS for CfnMemory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryApplicationLogs := awscdkmixinspreview.Mixins.NewCfnMemoryApplicationLogs()
//
type CfnMemoryApplicationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMemoryLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMemoryLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnMemoryLogsMixin
}

// The jsii proxy struct for CfnMemoryApplicationLogs
type jsiiProxy_CfnMemoryApplicationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMemoryApplicationLogs() CfnMemoryApplicationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMemoryApplicationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMemoryApplicationLogs_Override(c CfnMemoryApplicationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMemoryLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMemoryLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnMemoryLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


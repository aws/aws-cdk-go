package previewawstransfermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnServerLogsMixin to generate TRANSFER_LOGS for CfnServer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServerTransferLogs := awscdkmixinspreview.Mixins.NewCfnServerTransferLogs()
//
type CfnServerTransferLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnServerLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnServerLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnServerLogsMixin
}

// The jsii proxy struct for CfnServerTransferLogs
type jsiiProxy_CfnServerTransferLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnServerTransferLogs() CfnServerTransferLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnServerTransferLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnServerTransferLogs_Override(c CfnServerTransferLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerTransferLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnServerTransferLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnServerLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnServerLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerTransferLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnServerLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnServerLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerTransferLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnServerLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnServerLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


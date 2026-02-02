package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnClusterLogsMixin to generate AUTO_MODE_COMPUTE_LOGS for CfnCluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeComputeLogs := awscdkmixinspreview.Mixins.NewCfnClusterAutoModeComputeLogs()
//
type CfnClusterAutoModeComputeLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnClusterLogsMixin
}

// The jsii proxy struct for CfnClusterAutoModeComputeLogs
type jsiiProxy_CfnClusterAutoModeComputeLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAutoModeComputeLogs() CfnClusterAutoModeComputeLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAutoModeComputeLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAutoModeComputeLogs_Override(c CfnClusterAutoModeComputeLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin {
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

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin {
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

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnClusterLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


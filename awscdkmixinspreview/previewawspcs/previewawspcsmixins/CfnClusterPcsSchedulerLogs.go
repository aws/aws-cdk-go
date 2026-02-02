package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnClusterLogsMixin to generate PCS_SCHEDULER_LOGS for CfnCluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerLogs := awscdkmixinspreview.Mixins.NewCfnClusterPcsSchedulerLogs()
//
type CfnClusterPcsSchedulerLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnClusterLogsMixin
}

// The jsii proxy struct for CfnClusterPcsSchedulerLogs
type jsiiProxy_CfnClusterPcsSchedulerLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterPcsSchedulerLogs() CfnClusterPcsSchedulerLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterPcsSchedulerLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterPcsSchedulerLogs_Override(c CfnClusterPcsSchedulerLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnClusterLogsMixin {
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

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnClusterLogsMixin {
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

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnClusterLogsMixin {
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


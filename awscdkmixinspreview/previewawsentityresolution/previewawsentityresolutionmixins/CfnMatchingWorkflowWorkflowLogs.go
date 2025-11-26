package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnMatchingWorkflowLogsMixin to generate WORKFLOW_LOGS for CfnMatchingWorkflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowWorkflowLogs := awscdkmixinspreview.Mixins.NewCfnMatchingWorkflowWorkflowLogs()
//
type CfnMatchingWorkflowWorkflowLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMatchingWorkflowLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMatchingWorkflowLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnMatchingWorkflowLogsMixin
}

// The jsii proxy struct for CfnMatchingWorkflowWorkflowLogs
type jsiiProxy_CfnMatchingWorkflowWorkflowLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMatchingWorkflowWorkflowLogs() CfnMatchingWorkflowWorkflowLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchingWorkflowWorkflowLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMatchingWorkflowWorkflowLogs_Override(c CfnMatchingWorkflowWorkflowLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMatchingWorkflowLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnMatchingWorkflowLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMatchingWorkflowLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnMatchingWorkflowLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnMatchingWorkflowLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMatchingWorkflowLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
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
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnMatchingWorkflowLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMatchingWorkflowLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMatchingWorkflowLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMatchingWorkflowLogsMixin
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

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnMatchingWorkflowLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnMatchingWorkflowLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnMatchingWorkflowLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMatchingWorkflowLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


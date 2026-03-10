package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnClusterLogsMixin to generate PCS_SCHEDULER_AUDIT_LOGS for CfnCluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerAuditLogs := awscdkmixinspreview.Mixins.NewCfnClusterPcsSchedulerAuditLogs()
//
type CfnClusterPcsSchedulerAuditLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are S3, CWL, FH
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsSchedulerAuditLogsDestProps) CfnClusterLogsMixin
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsSchedulerAuditLogsFirehoseProps) CfnClusterLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsSchedulerAuditLogsLogGroupProps) CfnClusterLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsSchedulerAuditLogsS3Props) CfnClusterLogsMixin
}

// The jsii proxy struct for CfnClusterPcsSchedulerAuditLogs
type jsiiProxy_CfnClusterPcsSchedulerAuditLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterPcsSchedulerAuditLogs() CfnClusterPcsSchedulerAuditLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterPcsSchedulerAuditLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterPcsSchedulerAuditLogs_Override(c CfnClusterPcsSchedulerAuditLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsSchedulerAuditLogsDestProps) CfnClusterLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsSchedulerAuditLogsFirehoseProps) CfnClusterLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsSchedulerAuditLogsLogGroupProps) CfnClusterLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) ToS3(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsSchedulerAuditLogsS3Props) CfnClusterLogsMixin {
	if err := c.validateToS3Parameters(bucket, props); err != nil {
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


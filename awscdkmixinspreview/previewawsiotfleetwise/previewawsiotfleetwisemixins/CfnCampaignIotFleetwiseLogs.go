package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnCampaignLogsMixin to generate IOT_FLEETWISE_LOGS for CfnCampaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCampaignIotFleetwiseLogs := awscdkmixinspreview.Mixins.NewCfnCampaignIotFleetwiseLogs()
//
type CfnCampaignIotFleetwiseLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCampaignLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCampaignLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnCampaignLogsMixin
}

// The jsii proxy struct for CfnCampaignIotFleetwiseLogs
type jsiiProxy_CfnCampaignIotFleetwiseLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCampaignIotFleetwiseLogs() CfnCampaignIotFleetwiseLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCampaignIotFleetwiseLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCampaignIotFleetwiseLogs_Override(c CfnCampaignIotFleetwiseLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCampaignLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnCampaignLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCampaignLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnCampaignLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnCampaignLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnCampaignLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnVehicleLogsMixin to generate IOT_FLEETWISE_LOGS for CfnVehicle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVehicleIotFleetwiseLogs := awscdkmixinspreview.Mixins.NewCfnVehicleIotFleetwiseLogs()
//
type CfnVehicleIotFleetwiseLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVehicleLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVehicleLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnVehicleLogsMixin
}

// The jsii proxy struct for CfnVehicleIotFleetwiseLogs
type jsiiProxy_CfnVehicleIotFleetwiseLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnVehicleIotFleetwiseLogs() CfnVehicleIotFleetwiseLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnVehicleIotFleetwiseLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVehicleIotFleetwiseLogs_Override(c CfnVehicleIotFleetwiseLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVehicleLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnVehicleLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVehicleLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnVehicleLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnVehicleLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnVehicleLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


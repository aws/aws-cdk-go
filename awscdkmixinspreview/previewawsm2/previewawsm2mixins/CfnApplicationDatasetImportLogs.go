package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnApplicationLogsMixin to generate DATASET_IMPORT_LOGS for CfnApplication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationDatasetImportLogs := awscdkmixinspreview.Mixins.NewCfnApplicationDatasetImportLogs()
//
type CfnApplicationDatasetImportLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnApplicationLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnApplicationLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnApplicationLogsMixin
}

// The jsii proxy struct for CfnApplicationDatasetImportLogs
type jsiiProxy_CfnApplicationDatasetImportLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationDatasetImportLogs() CfnApplicationDatasetImportLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationDatasetImportLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationDatasetImportLogs_Override(c CfnApplicationDatasetImportLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnApplicationLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnApplicationLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnApplicationLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnApplicationLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


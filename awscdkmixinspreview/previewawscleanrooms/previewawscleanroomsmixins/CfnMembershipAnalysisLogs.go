package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnMembershipLogsMixin to generate ANALYSIS_LOGS for CfnMembership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMembershipAnalysisLogs := awscdkmixinspreview.Mixins.NewCfnMembershipAnalysisLogs()
//
type CfnMembershipAnalysisLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMembershipLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMembershipLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnMembershipLogsMixin
}

// The jsii proxy struct for CfnMembershipAnalysisLogs
type jsiiProxy_CfnMembershipAnalysisLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnMembershipAnalysisLogs() CfnMembershipAnalysisLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnMembershipAnalysisLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMembershipAnalysisLogs_Override(c CfnMembershipAnalysisLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnMembershipLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnMembershipLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnMembershipLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnMembershipLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnMembershipLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnMembershipLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


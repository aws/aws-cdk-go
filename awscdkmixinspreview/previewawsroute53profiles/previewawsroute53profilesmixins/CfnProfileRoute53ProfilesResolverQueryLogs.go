package previewawsroute53profilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnProfileLogsMixin to generate ROUTE53_PROFILES_RESOLVER_QUERY_LOGS for CfnProfile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfileRoute53ProfilesResolverQueryLogs := awscdkmixinspreview.Mixins.NewCfnProfileRoute53ProfilesResolverQueryLogs()
//
type CfnProfileRoute53ProfilesResolverQueryLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnProfileLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnProfileLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnProfileLogsMixin
}

// The jsii proxy struct for CfnProfileRoute53ProfilesResolverQueryLogs
type jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnProfileRoute53ProfilesResolverQueryLogs() CfnProfileRoute53ProfilesResolverQueryLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnProfileRoute53ProfilesResolverQueryLogs_Override(c CfnProfileRoute53ProfilesResolverQueryLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnProfileLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnProfileLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnProfileLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnProfileLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnProfileLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnProfileLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


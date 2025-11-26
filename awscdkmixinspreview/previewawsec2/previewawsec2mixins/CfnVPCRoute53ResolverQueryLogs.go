package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnVPCLogsMixin to generate ROUTE53_RESOLVER_QUERY_LOGS for CfnVPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCRoute53ResolverQueryLogs := awscdkmixinspreview.Mixins.NewCfnVPCRoute53ResolverQueryLogs()
//
type CfnVPCRoute53ResolverQueryLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVPCLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVPCLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnVPCLogsMixin
}

// The jsii proxy struct for CfnVPCRoute53ResolverQueryLogs
type jsiiProxy_CfnVPCRoute53ResolverQueryLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnVPCRoute53ResolverQueryLogs() CfnVPCRoute53ResolverQueryLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnVPCRoute53ResolverQueryLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVPCRoute53ResolverQueryLogs_Override(c CfnVPCRoute53ResolverQueryLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnVPCLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnVPCLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnVPCLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnVPCLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnVPCLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnVPCLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


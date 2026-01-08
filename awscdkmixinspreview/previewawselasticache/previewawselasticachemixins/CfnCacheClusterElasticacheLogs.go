package previewawselasticachemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnCacheClusterLogsMixin to generate ELASTICACHE_LOGS for CfnCacheCluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCacheClusterElasticacheLogs := awscdkmixinspreview.Mixins.NewCfnCacheClusterElasticacheLogs()
//
type CfnCacheClusterElasticacheLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCacheClusterLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCacheClusterLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnCacheClusterLogsMixin
}

// The jsii proxy struct for CfnCacheClusterElasticacheLogs
type jsiiProxy_CfnCacheClusterElasticacheLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnCacheClusterElasticacheLogs() CfnCacheClusterElasticacheLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnCacheClusterElasticacheLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterElasticacheLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCacheClusterElasticacheLogs_Override(c CfnCacheClusterElasticacheLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterElasticacheLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnCacheClusterLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnCacheClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnCacheClusterLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnCacheClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnCacheClusterLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnCacheClusterLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}


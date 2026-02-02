package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
)

// Builder for CfnWorkteamLogsMixin to generate ACTIVITY_LOGS for CfnWorkteam.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkteamActivityLogs := awscdkmixinspreview.Mixins.NewCfnWorkteamActivityLogs()
//
type CfnWorkteamActivityLogs interface {
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnWorkteamLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnWorkteamLogsMixin
}

// The jsii proxy struct for CfnWorkteamActivityLogs
type jsiiProxy_CfnWorkteamActivityLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnWorkteamActivityLogs() CfnWorkteamActivityLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkteamActivityLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWorkteamActivityLogs_Override(c CfnWorkteamActivityLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnWorkteamActivityLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnWorkteamLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnWorkteamLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteamActivityLogs) ToS3(bucket interfacesawss3.IBucketRef, props previewawslogs.IS3LogsDestinationProps) CfnWorkteamLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnWorkteamLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}


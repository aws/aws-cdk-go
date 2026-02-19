package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//   import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var bucket Bucket
//
//   distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(bucket),
//   	},
//   })
//
//   // Create log destination
//   logGroup := logs.NewLogGroup(scope, jsii.String("DeliveryLogGroup"))
//
//   // Configure log delivery using the mixin
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToLogGroup(logGroup))
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html
//
type CfnDistributionLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDistributionLogsMixin
type jsiiProxy_CfnDistributionLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDistributionLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::CloudFront::Distribution`.
func NewCfnDistributionLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnDistributionLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnDistributionLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDistributionLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::CloudFront::Distribution`.
func NewCfnDistributionLogsMixin_Override(c CfnDistributionLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDistributionLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDistributionLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionLogsMixin_ACCESS_LOGS() CfnDistributionAccessLogs {
	_init_.Initialize()
	var returns CfnDistributionAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		"ACCESS_LOGS",
		&returns,
	)
	return returns
}

func CfnDistributionLogsMixin_CONNECTION_LOGS() CfnDistributionConnectionLogs {
	_init_.Initialize()
	var returns CfnDistributionConnectionLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		"CONNECTION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDistributionLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


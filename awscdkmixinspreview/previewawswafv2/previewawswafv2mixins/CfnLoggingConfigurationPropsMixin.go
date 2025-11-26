package previewawswafv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswafv2/previewawswafv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an association between logging destinations and a web ACL resource, for logging from AWS WAF .
//
// As part of the association, you can specify parts of the standard logging fields to keep out of the logs and you can specify filters so that you log only a subset of the logging records.
//
// If you configure data protection for the web ACL, the protection applies to the data that AWS WAF sends to the logs.
//
// > You can define one logging destination per web ACL.
//
// You can access information about the traffic that AWS WAF inspects using the following steps:
//
// - Create your logging destination. You can use an Amazon CloudWatch Logs log group, an Amazon Simple Storage Service (Amazon S3) bucket, or an Amazon Kinesis Data Firehose.
//
// The name that you give the destination must start with `aws-waf-logs-` . Depending on the type of destination, you might need to configure additional settings or permissions.
//
// For configuration requirements and pricing information for each destination type, see [Logging web ACL traffic](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the *AWS WAF Developer Guide* .
// - Associate your logging destination to your web ACL using a `PutLoggingConfiguration` request.
//
// When you successfully enable logging using a `PutLoggingConfiguration` request, AWS WAF creates an additional role or policy that is required to write logs to the logging destination. For an Amazon CloudWatch Logs log group, AWS WAF creates a resource policy on the log group. For an Amazon S3 bucket, AWS WAF creates a bucket policy. For an Amazon Kinesis Data Firehose, AWS WAF creates a service-linked role.
//
// For additional information about web ACL logging, see [Logging web ACL traffic information](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the *AWS WAF Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var jsonBody interface{}
//   var loggingFilter interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   cfnLoggingConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnLoggingConfigurationPropsMixin(&CfnLoggingConfigurationMixinProps{
//   	LogDestinationConfigs: []*string{
//   		jsii.String("logDestinationConfigs"),
//   	},
//   	LoggingFilter: loggingFilter,
//   	RedactedFields: []interface{}{
//   		&FieldToMatchProperty{
//   			JsonBody: jsonBody,
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			UriPath: uriPath,
//   		},
//   	},
//   	ResourceArn: jsii.String("resourceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html
//
type CfnLoggingConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoggingConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoggingConfigurationPropsMixin
type jsiiProxy_CfnLoggingConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoggingConfigurationPropsMixin) Props() *CfnLoggingConfigurationMixinProps {
	var returns *CfnLoggingConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFv2::LoggingConfiguration`.
func NewCfnLoggingConfigurationPropsMixin(props *CfnLoggingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoggingConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoggingConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnLoggingConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFv2::LoggingConfiguration`.
func NewCfnLoggingConfigurationPropsMixin_Override(c CfnLoggingConfigurationPropsMixin, props *CfnLoggingConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnLoggingConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoggingConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoggingConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnLoggingConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggingConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnLoggingConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggingConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


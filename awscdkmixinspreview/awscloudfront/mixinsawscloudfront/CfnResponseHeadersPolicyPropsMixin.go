package mixinsawscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscloudfront/mixinsawscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A response headers policy.
//
// A response headers policy contains information about a set of HTTP response headers.
//
// After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
//
// For more information, see [Adding or removing HTTP headers in CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponseHeadersPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnResponseHeadersPolicyPropsMixin(&CfnResponseHeadersPolicyMixinProps{
//   	ResponseHeadersPolicyConfig: &ResponseHeadersPolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		CorsConfig: &CorsConfigProperty{
//   			AccessControlAllowCredentials: jsii.Boolean(false),
//   			AccessControlAllowHeaders: &AccessControlAllowHeadersProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlAllowMethods: &AccessControlAllowMethodsProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlAllowOrigins: &AccessControlAllowOriginsProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlExposeHeaders: &AccessControlExposeHeadersProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlMaxAgeSec: jsii.Number(123),
//   			OriginOverride: jsii.Boolean(false),
//   		},
//   		CustomHeadersConfig: &CustomHeadersConfigProperty{
//   			Items: []interface{}{
//   				&CustomHeaderProperty{
//   					Header: jsii.String("header"),
//   					Override: jsii.Boolean(false),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		RemoveHeadersConfig: &RemoveHeadersConfigProperty{
//   			Items: []interface{}{
//   				&RemoveHeaderProperty{
//   					Header: jsii.String("header"),
//   				},
//   			},
//   		},
//   		SecurityHeadersConfig: &SecurityHeadersConfigProperty{
//   			ContentSecurityPolicy: &ContentSecurityPolicyProperty{
//   				ContentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   				Override: jsii.Boolean(false),
//   			},
//   			ContentTypeOptions: &ContentTypeOptionsProperty{
//   				Override: jsii.Boolean(false),
//   			},
//   			FrameOptions: &FrameOptionsProperty{
//   				FrameOption: jsii.String("frameOption"),
//   				Override: jsii.Boolean(false),
//   			},
//   			ReferrerPolicy: &ReferrerPolicyProperty{
//   				Override: jsii.Boolean(false),
//   				ReferrerPolicy: jsii.String("referrerPolicy"),
//   			},
//   			StrictTransportSecurity: &StrictTransportSecurityProperty{
//   				AccessControlMaxAgeSec: jsii.Number(123),
//   				IncludeSubdomains: jsii.Boolean(false),
//   				Override: jsii.Boolean(false),
//   				Preload: jsii.Boolean(false),
//   			},
//   			XssProtection: &XSSProtectionProperty{
//   				ModeBlock: jsii.Boolean(false),
//   				Override: jsii.Boolean(false),
//   				Protection: jsii.Boolean(false),
//   				ReportUri: jsii.String("reportUri"),
//   			},
//   		},
//   		ServerTimingHeadersConfig: &ServerTimingHeadersConfigProperty{
//   			Enabled: jsii.Boolean(false),
//   			SamplingRate: jsii.Number(123),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-responseheaderspolicy.html
//
type CfnResponseHeadersPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResponseHeadersPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResponseHeadersPolicyPropsMixin
type jsiiProxy_CfnResponseHeadersPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResponseHeadersPolicyPropsMixin) Props() *CfnResponseHeadersPolicyMixinProps {
	var returns *CfnResponseHeadersPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::ResponseHeadersPolicy`.
func NewCfnResponseHeadersPolicyPropsMixin(props *CfnResponseHeadersPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResponseHeadersPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResponseHeadersPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResponseHeadersPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::ResponseHeadersPolicy`.
func NewCfnResponseHeadersPolicyPropsMixin_Override(c CfnResponseHeadersPolicyPropsMixin, props *CfnResponseHeadersPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResponseHeadersPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResponseHeadersPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResponseHeadersPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResponseHeadersPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResponseHeadersPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


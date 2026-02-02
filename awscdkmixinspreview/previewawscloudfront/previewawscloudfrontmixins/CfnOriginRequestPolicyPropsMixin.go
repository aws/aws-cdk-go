package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An origin request policy.
//
// When it's attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
//
// - The request body and the URL path (without the domain name) from the viewer request.
// - The headers that CloudFront automatically includes in every origin request, including `Host` , `User-Agent` , and `X-Amz-Cf-Id` .
// - All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
//
// CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use `CachePolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOriginRequestPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnOriginRequestPolicyPropsMixin(&CfnOriginRequestPolicyMixinProps{
//   	OriginRequestPolicyConfig: &OriginRequestPolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		CookiesConfig: &CookiesConfigProperty{
//   			CookieBehavior: jsii.String("cookieBehavior"),
//   			Cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		HeadersConfig: &HeadersConfigProperty{
//   			HeaderBehavior: jsii.String("headerBehavior"),
//   			Headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		QueryStringsConfig: &QueryStringsConfigProperty{
//   			QueryStringBehavior: jsii.String("queryStringBehavior"),
//   			QueryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html
//
type CfnOriginRequestPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOriginRequestPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOriginRequestPolicyPropsMixin
type jsiiProxy_CfnOriginRequestPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOriginRequestPolicyPropsMixin) Props() *CfnOriginRequestPolicyMixinProps {
	var returns *CfnOriginRequestPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::OriginRequestPolicy`.
func NewCfnOriginRequestPolicyPropsMixin(props *CfnOriginRequestPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOriginRequestPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOriginRequestPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginRequestPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::OriginRequestPolicy`.
func NewCfnOriginRequestPolicyPropsMixin_Override(c CfnOriginRequestPolicyPropsMixin, props *CfnOriginRequestPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOriginRequestPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginRequestPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginRequestPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginRequestPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOriginRequestPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


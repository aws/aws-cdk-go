package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A cache policy.
//
// When it's attached to a cache behavior, the cache policy determines the following:
//
// - The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
// - The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
// The headers, cookies, and query strings that are included in the cache key are also included in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a valid object in its cache that matches the request's cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCachePolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnCachePolicyPropsMixin(&CfnCachePolicyMixinProps{
//   	CachePolicyConfig: &CachePolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		DefaultTtl: jsii.Number(123),
//   		MaxTtl: jsii.Number(123),
//   		MinTtl: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		ParametersInCacheKeyAndForwardedToOrigin: &ParametersInCacheKeyAndForwardedToOriginProperty{
//   			CookiesConfig: &CookiesConfigProperty{
//   				CookieBehavior: jsii.String("cookieBehavior"),
//   				Cookies: []*string{
//   					jsii.String("cookies"),
//   				},
//   			},
//   			EnableAcceptEncodingBrotli: jsii.Boolean(false),
//   			EnableAcceptEncodingGzip: jsii.Boolean(false),
//   			HeadersConfig: &HeadersConfigProperty{
//   				HeaderBehavior: jsii.String("headerBehavior"),
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   			},
//   			QueryStringsConfig: &QueryStringsConfigProperty{
//   				QueryStringBehavior: jsii.String("queryStringBehavior"),
//   				QueryStrings: []*string{
//   					jsii.String("queryStrings"),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cachepolicy.html
//
type CfnCachePolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCachePolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCachePolicyPropsMixin
type jsiiProxy_CfnCachePolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCachePolicyPropsMixin) Props() *CfnCachePolicyMixinProps {
	var returns *CfnCachePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::CachePolicy`.
func NewCfnCachePolicyPropsMixin(props *CfnCachePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCachePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCachePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCachePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::CachePolicy`.
func NewCfnCachePolicyPropsMixin_Override(c CfnCachePolicyPropsMixin, props *CfnCachePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCachePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCachePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCachePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCachePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCachePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


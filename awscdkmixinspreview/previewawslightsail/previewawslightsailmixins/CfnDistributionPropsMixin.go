package previewawslightsailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslightsail/previewawslightsailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::Distribution` resource specifies a content delivery network (CDN) distribution.
//
// You can create distributions only in the `us-east-1` AWS Region.
//
// A distribution is a globally distributed network of caching servers that improve the performance of your website or web application hosted on a Lightsail instance, static content hosted on a Lightsail bucket, or through a Lightsail load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionPropsMixin := awscdkmixinspreview.Mixins.NewCfnDistributionPropsMixin(&CfnDistributionMixinProps{
//   	BundleId: jsii.String("bundleId"),
//   	CacheBehaviors: []interface{}{
//   		&CacheBehaviorPerPathProperty{
//   			Behavior: jsii.String("behavior"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	CacheBehaviorSettings: &CacheSettingsProperty{
//   		AllowedHttpMethods: jsii.String("allowedHttpMethods"),
//   		CachedHttpMethods: jsii.String("cachedHttpMethods"),
//   		DefaultTtl: jsii.Number(123),
//   		ForwardedCookies: &CookieObjectProperty{
//   			CookiesAllowList: []*string{
//   				jsii.String("cookiesAllowList"),
//   			},
//   			Option: jsii.String("option"),
//   		},
//   		ForwardedHeaders: &HeaderObjectProperty{
//   			HeadersAllowList: []*string{
//   				jsii.String("headersAllowList"),
//   			},
//   			Option: jsii.String("option"),
//   		},
//   		ForwardedQueryStrings: &QueryStringObjectProperty{
//   			Option: jsii.Boolean(false),
//   			QueryStringsAllowList: []*string{
//   				jsii.String("queryStringsAllowList"),
//   			},
//   		},
//   		MaximumTtl: jsii.Number(123),
//   		MinimumTtl: jsii.Number(123),
//   	},
//   	CertificateName: jsii.String("certificateName"),
//   	DefaultCacheBehavior: &CacheBehaviorProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   	DistributionName: jsii.String("distributionName"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	IsEnabled: jsii.Boolean(false),
//   	Origin: &InputOriginProperty{
//   		Name: jsii.String("name"),
//   		ProtocolPolicy: jsii.String("protocolPolicy"),
//   		RegionName: jsii.String("regionName"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html
//
type CfnDistributionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDistributionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDistributionPropsMixin
type jsiiProxy_CfnDistributionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDistributionPropsMixin) Props() *CfnDistributionMixinProps {
	var returns *CfnDistributionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::Distribution`.
func NewCfnDistributionPropsMixin(props *CfnDistributionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDistributionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDistributionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDistributionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::Distribution`.
func NewCfnDistributionPropsMixin_Override(c CfnDistributionPropsMixin, props *CfnDistributionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDistributionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDistributionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDistributionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


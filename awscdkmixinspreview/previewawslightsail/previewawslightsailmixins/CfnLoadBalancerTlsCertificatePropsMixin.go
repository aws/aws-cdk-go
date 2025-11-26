package previewawslightsailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslightsail/previewawslightsailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::LoadBalancerTlsCertificate` resource specifies a TLS certificate that can be used with a Lightsail load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerTlsCertificatePropsMixin := awscdkmixinspreview.Mixins.NewCfnLoadBalancerTlsCertificatePropsMixin(&CfnLoadBalancerTlsCertificateMixinProps{
//   	CertificateAlternativeNames: []*string{
//   		jsii.String("certificateAlternativeNames"),
//   	},
//   	CertificateDomainName: jsii.String("certificateDomainName"),
//   	CertificateName: jsii.String("certificateName"),
//   	HttpsRedirectionEnabled: jsii.Boolean(false),
//   	IsAttached: jsii.Boolean(false),
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html
//
type CfnLoadBalancerTlsCertificatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoadBalancerTlsCertificateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoadBalancerTlsCertificatePropsMixin
type jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin) Props() *CfnLoadBalancerTlsCertificateMixinProps {
	var returns *CfnLoadBalancerTlsCertificateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::LoadBalancerTlsCertificate`.
func NewCfnLoadBalancerTlsCertificatePropsMixin(props *CfnLoadBalancerTlsCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoadBalancerTlsCertificatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoadBalancerTlsCertificatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::LoadBalancerTlsCertificate`.
func NewCfnLoadBalancerTlsCertificatePropsMixin_Override(c CfnLoadBalancerTlsCertificatePropsMixin, props *CfnLoadBalancerTlsCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoadBalancerTlsCertificatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoadBalancerTlsCertificatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancerTlsCertificatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawselasticloadbalancingmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticloadbalancing/previewawselasticloadbalancingmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Classic Load Balancer.
//
// If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the VPC-gateway attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributes interface{}
//
//   cfnLoadBalancerPropsMixin := awscdkmixinspreview.Mixins.NewCfnLoadBalancerPropsMixin(&CfnLoadBalancerMixinProps{
//   	AccessLoggingPolicy: &AccessLoggingPolicyProperty{
//   		EmitInterval: jsii.Number(123),
//   		Enabled: jsii.Boolean(false),
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	AppCookieStickinessPolicy: []interface{}{
//   		&AppCookieStickinessPolicyProperty{
//   			CookieName: jsii.String("cookieName"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ConnectionDrainingPolicy: &ConnectionDrainingPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//   		Timeout: jsii.Number(123),
//   	},
//   	ConnectionSettings: &ConnectionSettingsProperty{
//   		IdleTimeout: jsii.Number(123),
//   	},
//   	CrossZone: jsii.Boolean(false),
//   	HealthCheck: &HealthCheckProperty{
//   		HealthyThreshold: jsii.String("healthyThreshold"),
//   		Interval: jsii.String("interval"),
//   		Target: jsii.String("target"),
//   		Timeout: jsii.String("timeout"),
//   		UnhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	Instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	LbCookieStickinessPolicy: []interface{}{
//   		&LBCookieStickinessPolicyProperty{
//   			CookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	Listeners: []interface{}{
//   		&ListenersProperty{
//   			InstancePort: jsii.String("instancePort"),
//   			InstanceProtocol: jsii.String("instanceProtocol"),
//   			LoadBalancerPort: jsii.String("loadBalancerPort"),
//   			PolicyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			Protocol: jsii.String("protocol"),
//   			SslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	Policies: []interface{}{
//   		&PoliciesProperty{
//   			Attributes: []interface{}{
//   				attributes,
//   			},
//   			InstancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			LoadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   			PolicyName: jsii.String("policyName"),
//   			PolicyType: jsii.String("policyType"),
//   		},
//   	},
//   	Scheme: jsii.String("scheme"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html
//
type CfnLoadBalancerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoadBalancerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoadBalancerPropsMixin
type jsiiProxy_CfnLoadBalancerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoadBalancerPropsMixin) Props() *CfnLoadBalancerMixinProps {
	var returns *CfnLoadBalancerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticLoadBalancing::LoadBalancer`.
func NewCfnLoadBalancerPropsMixin(props *CfnLoadBalancerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoadBalancerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoadBalancerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoadBalancerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticLoadBalancing::LoadBalancer`.
func NewCfnLoadBalancerPropsMixin_Override(c CfnLoadBalancerPropsMixin, props *CfnLoadBalancerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoadBalancerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoadBalancerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoadBalancerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLoadBalancerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


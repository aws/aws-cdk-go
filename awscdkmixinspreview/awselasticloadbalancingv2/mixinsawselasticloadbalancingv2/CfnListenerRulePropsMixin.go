package mixinsawselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awselasticloadbalancingv2/mixinsawselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a listener rule.
//
// The listener must be associated with an Application Load Balancer. Each rule consists of a priority, one or more actions, and one or more conditions.
//
// For more information, see [Quotas for your Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html) in the *User Guide for Application Load Balancers* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnListenerRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnListenerRulePropsMixin(&CfnListenerRuleMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   				UserPoolArn: jsii.String("userPoolArn"),
//   				UserPoolClientId: jsii.String("userPoolClientId"),
//   				UserPoolDomain: jsii.String("userPoolDomain"),
//   			},
//   			AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				Issuer: jsii.String("issuer"),
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   				UseExistingClientSecret: jsii.Boolean(false),
//   				UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   			},
//   			FixedResponseConfig: &FixedResponseConfigProperty{
//   				ContentType: jsii.String("contentType"),
//   				MessageBody: jsii.String("messageBody"),
//   				StatusCode: jsii.String("statusCode"),
//   			},
//   			ForwardConfig: &ForwardConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupTupleProperty{
//   						TargetGroupArn: jsii.String("targetGroupArn"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   				TargetGroupStickinessConfig: &TargetGroupStickinessConfigProperty{
//   					DurationSeconds: jsii.Number(123),
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   			JwtValidationConfig: &JwtValidationConfigProperty{
//   				AdditionalClaims: []interface{}{
//   					&JwtValidationActionAdditionalClaimProperty{
//   						Format: jsii.String("format"),
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				Issuer: jsii.String("issuer"),
//   				JwksEndpoint: jsii.String("jwksEndpoint"),
//   			},
//   			Order: jsii.Number(123),
//   			RedirectConfig: &RedirectConfigProperty{
//   				Host: jsii.String("host"),
//   				Path: jsii.String("path"),
//   				Port: jsii.String("port"),
//   				Protocol: jsii.String("protocol"),
//   				Query: jsii.String("query"),
//   				StatusCode: jsii.String("statusCode"),
//   			},
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&RuleConditionProperty{
//   			Field: jsii.String("field"),
//   			HostHeaderConfig: &HostHeaderConfigProperty{
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			HttpHeaderConfig: &HttpHeaderConfigProperty{
//   				HttpHeaderName: jsii.String("httpHeaderName"),
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			HttpRequestMethodConfig: &HttpRequestMethodConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PathPatternConfig: &PathPatternConfigProperty{
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			QueryStringConfig: &QueryStringConfigProperty{
//   				Values: []interface{}{
//   					&QueryStringKeyValueProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			RegexValues: []*string{
//   				jsii.String("regexValues"),
//   			},
//   			SourceIpConfig: &SourceIpConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	ListenerArn: jsii.String("listenerArn"),
//   	Priority: jsii.Number(123),
//   	Transforms: []interface{}{
//   		&TransformProperty{
//   			HostHeaderRewriteConfig: &RewriteConfigObjectProperty{
//   				Rewrites: []interface{}{
//   					&RewriteConfigProperty{
//   						Regex: jsii.String("regex"),
//   						Replace: jsii.String("replace"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   			UrlRewriteConfig: &RewriteConfigObjectProperty{
//   				Rewrites: []interface{}{
//   					&RewriteConfigProperty{
//   						Regex: jsii.String("regex"),
//   						Replace: jsii.String("replace"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
//
type CfnListenerRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnListenerRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnListenerRulePropsMixin
type jsiiProxy_CfnListenerRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnListenerRulePropsMixin) Props() *CfnListenerRuleMixinProps {
	var returns *CfnListenerRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::ListenerRule`.
func NewCfnListenerRulePropsMixin(props *CfnListenerRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnListenerRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnListenerRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnListenerRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::ListenerRule`.
func NewCfnListenerRulePropsMixin_Override(c CfnListenerRulePropsMixin, props *CfnListenerRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnListenerRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnListenerRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListenerRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnListenerRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnListenerRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


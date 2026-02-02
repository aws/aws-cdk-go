package previewawswafv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswafv2/previewawswafv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019.
//
// For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF developer guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `WebACL` to define a collection of rules to use to inspect and control web requests. Each rule in a web ACL has a statement that defines what to look for in web requests and an action that AWS WAF applies to requests that match the statement. In the web ACL, you assign a default action to take (allow, block) for any request that doesn't match any of the rules.
//
// The rules in a web ACL can be a combination of explicitly defined rules and rule groups that you reference from the web ACL. The rule groups can be rule groups that you manage or rule groups that are managed by others.
//
// You can associate a web ACL with one or more AWS resources to protect. The resources can be an Amazon CloudFront distribution, an  REST API, an Application Load Balancer , an AWS AppSync GraphQL API , an Amazon Cognito user pool, an AWS App Runner service, an AWS Amplify application, or an AWS Verified Access instance.
//
// For more information, see [Web access control lists (web ACLs)](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl.html) in the *AWS WAF developer guide* .
//
// *Web ACLs used in AWS Shield Advanced automatic application layer DDoS mitigation*
//
// If you use Shield Advanced automatic application layer DDoS mitigation, the web ACLs that you use with automatic mitigation have a rule group rule whose name starts with `ShieldMitigationRuleGroup` . This rule is used for automatic mitigations and it's managed for you in the web ACL by Shield Advanced and AWS WAF . You'll see the rule listed among the web ACL rules when you view the web ACL through the AWS WAF interfaces.
//
// When you manage the web ACL through CloudFormation interfaces, you won't see the Shield Advanced rule. CloudFormation doesn't include this type of rule in the stack drift status between the actual configuration of the web ACL and your web ACL template.
//
// Don't add the Shield Advanced rule group rule to your web ACL template. The rule shouldn't be in your template. When you update the web ACL template in a stack, the Shield Advanced rule is maintained for you by AWS WAF in the resulting web ACL.
//
// For more information, see [Shield Advanced automatic application layer DDoS mitigation](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-automatic-app-layer-response.html) in the *AWS Shield Advanced developer guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html
//
type CfnWebACLPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWebACLMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWebACLPropsMixin
type jsiiProxy_CfnWebACLPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWebACLPropsMixin) Props() *CfnWebACLMixinProps {
	var returns *CfnWebACLMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WAFv2::WebACL`.
func NewCfnWebACLPropsMixin(props *CfnWebACLMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWebACLPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWebACLPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWebACLPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WAFv2::WebACL`.
func NewCfnWebACLPropsMixin_Override(c CfnWebACLPropsMixin, props *CfnWebACLMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWebACLPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWebACLPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACLPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebACLPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWebACLPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


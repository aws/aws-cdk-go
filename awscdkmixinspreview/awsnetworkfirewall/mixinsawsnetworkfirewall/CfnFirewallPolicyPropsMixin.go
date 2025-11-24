package mixinsawsnetworkfirewall

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsnetworkfirewall/mixinsawsnetworkfirewall/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the firewall policy to define the stateless and stateful network traffic filtering behavior for your firewall.
//
// You can use one firewall policy for multiple firewalls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnFirewallPolicyPropsMixin(&CfnFirewallPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	FirewallPolicy: &FirewallPolicyProperty{
//   		EnableTlsSessionHolding: jsii.Boolean(false),
//   		PolicyVariables: &PolicyVariablesProperty{
//   			RuleVariables: map[string]interface{}{
//   				"ruleVariablesKey": map[string][]*string{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   		},
//   		StatefulDefaultActions: []*string{
//   			jsii.String("statefulDefaultActions"),
//   		},
//   		StatefulEngineOptions: &StatefulEngineOptionsProperty{
//   			FlowTimeouts: &FlowTimeoutsProperty{
//   				TcpIdleTimeoutSeconds: jsii.Number(123),
//   			},
//   			RuleOrder: jsii.String("ruleOrder"),
//   			StreamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   		},
//   		StatefulRuleGroupReferences: []interface{}{
//   			&StatefulRuleGroupReferenceProperty{
//   				DeepThreatInspection: jsii.Boolean(false),
//   				Override: &StatefulRuleGroupOverrideProperty{
//   					Action: jsii.String("action"),
//   				},
//   				Priority: jsii.Number(123),
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   		StatelessCustomActions: []interface{}{
//   			&CustomActionProperty{
//   				ActionDefinition: &ActionDefinitionProperty{
//   					PublishMetricAction: &PublishMetricActionProperty{
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				ActionName: jsii.String("actionName"),
//   			},
//   		},
//   		StatelessDefaultActions: []*string{
//   			jsii.String("statelessDefaultActions"),
//   		},
//   		StatelessFragmentDefaultActions: []*string{
//   			jsii.String("statelessFragmentDefaultActions"),
//   		},
//   		StatelessRuleGroupReferences: []interface{}{
//   			&StatelessRuleGroupReferenceProperty{
//   				Priority: jsii.Number(123),
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   		TlsInspectionConfigurationArn: jsii.String("tlsInspectionConfigurationArn"),
//   	},
//   	FirewallPolicyName: jsii.String("firewallPolicyName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html
//
type CfnFirewallPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFirewallPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFirewallPolicyPropsMixin
type jsiiProxy_CfnFirewallPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFirewallPolicyPropsMixin) Props() *CfnFirewallPolicyMixinProps {
	var returns *CfnFirewallPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::FirewallPolicy`.
func NewCfnFirewallPolicyPropsMixin(props *CfnFirewallPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFirewallPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFirewallPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewallPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::FirewallPolicy`.
func NewCfnFirewallPolicyPropsMixin_Override(c CfnFirewallPolicyPropsMixin, props *CfnFirewallPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFirewallPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewallPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFirewallPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkfirewall/previewawsnetworkfirewallmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the [](https://docs.aws.amazon.com/RuleGroup) to define a reusable collection of stateless or stateful network traffic filtering rules. You use rule groups in an firewall policy to specify the filtering behavior of an firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuleGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnRuleGroupPropsMixin(&CfnRuleGroupMixinProps{
//   	Capacity: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	RuleGroup: &RuleGroupProperty{
//   		ReferenceSets: &ReferenceSetsProperty{
//   			IpSetReferences: map[string]interface{}{
//   				"ipSetReferencesKey": map[string]*string{
//   					"referenceArn": jsii.String("referenceArn"),
//   				},
//   			},
//   		},
//   		RulesSource: &RulesSourceProperty{
//   			RulesSourceList: &RulesSourceListProperty{
//   				GeneratedRulesType: jsii.String("generatedRulesType"),
//   				Targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				TargetTypes: []*string{
//   					jsii.String("targetTypes"),
//   				},
//   			},
//   			RulesString: jsii.String("rulesString"),
//   			StatefulRules: []interface{}{
//   				&StatefulRuleProperty{
//   					Action: jsii.String("action"),
//   					Header: &HeaderProperty{
//   						Destination: jsii.String("destination"),
//   						DestinationPort: jsii.String("destinationPort"),
//   						Direction: jsii.String("direction"),
//   						Protocol: jsii.String("protocol"),
//   						Source: jsii.String("source"),
//   						SourcePort: jsii.String("sourcePort"),
//   					},
//   					RuleOptions: []interface{}{
//   						&RuleOptionProperty{
//   							Keyword: jsii.String("keyword"),
//   							Settings: []*string{
//   								jsii.String("settings"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			StatelessRulesAndCustomActions: &StatelessRulesAndCustomActionsProperty{
//   				CustomActions: []interface{}{
//   					&CustomActionProperty{
//   						ActionDefinition: &ActionDefinitionProperty{
//   							PublishMetricAction: &PublishMetricActionProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						ActionName: jsii.String("actionName"),
//   					},
//   				},
//   				StatelessRules: []interface{}{
//   					&StatelessRuleProperty{
//   						Priority: jsii.Number(123),
//   						RuleDefinition: &RuleDefinitionProperty{
//   							Actions: []*string{
//   								jsii.String("actions"),
//   							},
//   							MatchAttributes: &MatchAttributesProperty{
//   								DestinationPorts: []interface{}{
//   									&PortRangeProperty{
//   										FromPort: jsii.Number(123),
//   										ToPort: jsii.Number(123),
//   									},
//   								},
//   								Destinations: []interface{}{
//   									&AddressProperty{
//   										AddressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								Protocols: []interface{}{
//   									jsii.Number(123),
//   								},
//   								SourcePorts: []interface{}{
//   									&PortRangeProperty{
//   										FromPort: jsii.Number(123),
//   										ToPort: jsii.Number(123),
//   									},
//   								},
//   								Sources: []interface{}{
//   									&AddressProperty{
//   										AddressDefinition: jsii.String("addressDefinition"),
//   									},
//   								},
//   								TcpFlags: []interface{}{
//   									&TCPFlagFieldProperty{
//   										Flags: []*string{
//   											jsii.String("flags"),
//   										},
//   										Masks: []*string{
//   											jsii.String("masks"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RuleVariables: &RuleVariablesProperty{
//   			IpSets: map[string]interface{}{
//   				"ipSetsKey": map[string][]*string{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   			PortSets: map[string]interface{}{
//   				"portSetsKey": &PortSetProperty{
//   					"definition": []*string{
//   						jsii.String("definition"),
//   					},
//   				},
//   			},
//   		},
//   		StatefulRuleOptions: &StatefulRuleOptionsProperty{
//   			RuleOrder: jsii.String("ruleOrder"),
//   		},
//   	},
//   	RuleGroupName: jsii.String("ruleGroupName"),
//   	SummaryConfiguration: &SummaryConfigurationProperty{
//   		RuleOptions: []*string{
//   			jsii.String("ruleOptions"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-rulegroup.html
//
type CfnRuleGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRuleGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRuleGroupPropsMixin
type jsiiProxy_CfnRuleGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRuleGroupPropsMixin) Props() *CfnRuleGroupMixinProps {
	var returns *CfnRuleGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::RuleGroup`.
func NewCfnRuleGroupPropsMixin(props *CfnRuleGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRuleGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRuleGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRuleGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::RuleGroup`.
func NewCfnRuleGroupPropsMixin_Override(c CfnRuleGroupPropsMixin, props *CfnRuleGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRuleGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRuleGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuleGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRuleGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


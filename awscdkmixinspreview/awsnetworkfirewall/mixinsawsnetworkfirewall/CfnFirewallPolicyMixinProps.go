package mixinsawsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFirewallPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallPolicyMixinProps := &CfnFirewallPolicyMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html
//
type CfnFirewallPolicyMixinProps struct {
	// A description of the firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy
	//
	FirewallPolicy interface{} `field:"optional" json:"firewallPolicy" yaml:"firewallPolicy"`
	// The descriptive name of the firewall policy.
	//
	// You can't change the name of a firewall policy after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicyname
	//
	FirewallPolicyName *string `field:"optional" json:"firewallPolicyName" yaml:"firewallPolicyName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFirewallPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallPolicyProps := &CfnFirewallPolicyProps{
//   	FirewallPolicy: &FirewallPolicyProperty{
//   		StatelessDefaultActions: []*string{
//   			jsii.String("statelessDefaultActions"),
//   		},
//   		StatelessFragmentDefaultActions: []*string{
//   			jsii.String("statelessFragmentDefaultActions"),
//   		},
//
//   		// the properties below are optional
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
//   			RuleOrder: jsii.String("ruleOrder"),
//   			StreamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   		},
//   		StatefulRuleGroupReferences: []interface{}{
//   			&StatefulRuleGroupReferenceProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//
//   				// the properties below are optional
//   				Override: &StatefulRuleGroupOverrideProperty{
//   					Action: jsii.String("action"),
//   				},
//   				Priority: jsii.Number(123),
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
//   		StatelessRuleGroupReferences: []interface{}{
//   			&StatelessRuleGroupReferenceProperty{
//   				Priority: jsii.Number(123),
//   				ResourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   	},
//   	FirewallPolicyName: jsii.String("firewallPolicyName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html
//
type CfnFirewallPolicyProps struct {
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy
	//
	FirewallPolicy interface{} `field:"required" json:"firewallPolicy" yaml:"firewallPolicy"`
	// The descriptive name of the firewall policy.
	//
	// You can't change the name of a firewall policy after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicyname
	//
	FirewallPolicyName *string `field:"required" json:"firewallPolicyName" yaml:"firewallPolicyName"`
	// A description of the firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


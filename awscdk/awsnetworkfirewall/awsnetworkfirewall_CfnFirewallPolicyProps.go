package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFirewallPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallPolicyProps := &cfnFirewallPolicyProps{
//   	firewallPolicy: &firewallPolicyProperty{
//   		statelessDefaultActions: []*string{
//   			jsii.String("statelessDefaultActions"),
//   		},
//   		statelessFragmentDefaultActions: []*string{
//   			jsii.String("statelessFragmentDefaultActions"),
//   		},
//
//   		// the properties below are optional
//   		statefulDefaultActions: []*string{
//   			jsii.String("statefulDefaultActions"),
//   		},
//   		statefulEngineOptions: &statefulEngineOptionsProperty{
//   			ruleOrder: jsii.String("ruleOrder"),
//   			streamExceptionPolicy: jsii.String("streamExceptionPolicy"),
//   		},
//   		statefulRuleGroupReferences: []interface{}{
//   			&statefulRuleGroupReferenceProperty{
//   				resourceArn: jsii.String("resourceArn"),
//
//   				// the properties below are optional
//   				override: &statefulRuleGroupOverrideProperty{
//   					action: jsii.String("action"),
//   				},
//   				priority: jsii.Number(123),
//   			},
//   		},
//   		statelessCustomActions: []interface{}{
//   			&customActionProperty{
//   				actionDefinition: &actionDefinitionProperty{
//   					publishMetricAction: &publishMetricActionProperty{
//   						dimensions: []interface{}{
//   							&dimensionProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				actionName: jsii.String("actionName"),
//   			},
//   		},
//   		statelessRuleGroupReferences: []interface{}{
//   			&statelessRuleGroupReferenceProperty{
//   				priority: jsii.Number(123),
//   				resourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   	},
//   	firewallPolicyName: jsii.String("firewallPolicyName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallPolicyProps struct {
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	FirewallPolicy interface{} `field:"required" json:"firewallPolicy" yaml:"firewallPolicy"`
	// The descriptive name of the firewall policy.
	//
	// You can't change the name of a firewall policy after you create it.
	FirewallPolicyName *string `field:"required" json:"firewallPolicyName" yaml:"firewallPolicyName"`
	// A description of the firewall policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


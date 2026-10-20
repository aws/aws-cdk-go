package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPolicyMixinProps := &CfnPolicyMixinProps{
//   	AssociatedTemplateAndRuleList: []interface{}{
//   		&AssociatedTemplateAndRuleProperty{
//   			RuleArn: jsii.String("ruleArn"),
//   			TemplateArn: jsii.String("templateArn"),
//   		},
//   	},
//   	FirewallType: jsii.String("firewallType"),
//   	PolicyConfiguration: &PolicyConfigurationProperty{
//   		RemediationEnabled: jsii.Boolean(false),
//   		ResourcesCleanUp: jsii.Boolean(false),
//   		WafConfig: &WafConfigProperty{
//   			ConflictResolution: jsii.String("conflictResolution"),
//   			ExistingCustomerWebAclResolution: jsii.String("existingCustomerWebAclResolution"),
//   		},
//   	},
//   	PolicyDescription: jsii.String("policyDescription"),
//   	PolicyName: jsii.String("policyName"),
//   	Priority: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html
//
type CfnPolicyMixinProps struct {
	// List of templates and rules associated with this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-associatedtemplateandrulelist
	//
	AssociatedTemplateAndRuleList interface{} `field:"optional" json:"associatedTemplateAndRuleList" yaml:"associatedTemplateAndRuleList"`
	// The type of firewall.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-firewalltype
	//
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// Configuration settings for policy behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policyconfiguration
	//
	PolicyConfiguration interface{} `field:"optional" json:"policyConfiguration" yaml:"policyConfiguration"`
	// A description of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policydescription
	//
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The priority of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The tags associated with the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


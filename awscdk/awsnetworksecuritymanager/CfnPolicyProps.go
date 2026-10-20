package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	FirewallType: jsii.String("firewallType"),
//   	PolicyConfiguration: &PolicyConfigurationProperty{
//   		RemediationEnabled: jsii.Boolean(false),
//   		ResourcesCleanUp: jsii.Boolean(false),
//   		WafConfig: &WafConfigProperty{
//   			ConflictResolution: jsii.String("conflictResolution"),
//   			ExistingCustomerWebAclResolution: jsii.String("existingCustomerWebAclResolution"),
//   		},
//   	},
//   	PolicyName: jsii.String("policyName"),
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	AssociatedTemplateAndRuleList: []interface{}{
//   		&AssociatedTemplateAndRuleProperty{
//   			RuleArn: jsii.String("ruleArn"),
//   			TemplateArn: jsii.String("templateArn"),
//   		},
//   	},
//   	PolicyDescription: jsii.String("policyDescription"),
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
type CfnPolicyProps struct {
	// The type of firewall.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-firewalltype
	//
	FirewallType *string `field:"required" json:"firewallType" yaml:"firewallType"`
	// Configuration settings for policy behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policyconfiguration
	//
	PolicyConfiguration interface{} `field:"required" json:"policyConfiguration" yaml:"policyConfiguration"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The priority of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// List of templates and rules associated with this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-associatedtemplateandrulelist
	//
	AssociatedTemplateAndRuleList interface{} `field:"optional" json:"associatedTemplateAndRuleList" yaml:"associatedTemplateAndRuleList"`
	// A description of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-policydescription
	//
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// The tags associated with the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-policy.html#cfn-networksecuritymanager-policy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


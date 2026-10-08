package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTemplateMixinProps := &CfnTemplateMixinProps{
//   	AssociatedRuleList: []interface{}{
//   		&AssociatedRuleProperty{
//   			RuleArn: jsii.String("ruleArn"),
//   		},
//   	},
//   	FirewallType: jsii.String("firewallType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html
//
type CfnTemplateMixinProps struct {
	// List of rules associated with this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html#cfn-networksecuritymanager-template-associatedrulelist
	//
	AssociatedRuleList interface{} `field:"optional" json:"associatedRuleList" yaml:"associatedRuleList"`
	// The type of firewall.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html#cfn-networksecuritymanager-template-firewalltype
	//
	FirewallType *string `field:"optional" json:"firewallType" yaml:"firewallType"`
	// The tags associated with the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html#cfn-networksecuritymanager-template-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html#cfn-networksecuritymanager-template-templatedescription
	//
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
	// The name of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-template.html#cfn-networksecuritymanager-template-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}


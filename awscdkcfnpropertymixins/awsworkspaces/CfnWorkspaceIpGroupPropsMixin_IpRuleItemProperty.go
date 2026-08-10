package awsworkspaces


// Describes a rule for an IP access control group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ipRuleItemProperty := &IpRuleItemProperty{
//   	IpRule: jsii.String("ipRule"),
//   	RuleDesc: jsii.String("ruleDesc"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspaceipgroup-ipruleitem.html
//
type CfnWorkspaceIpGroupPropsMixin_IpRuleItemProperty struct {
	// The IP address range, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspaceipgroup-ipruleitem.html#cfn-workspaces-workspaceipgroup-ipruleitem-iprule
	//
	IpRule *string `field:"optional" json:"ipRule" yaml:"ipRule"`
	// The description of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspaceipgroup-ipruleitem.html#cfn-workspaces-workspaceipgroup-ipruleitem-ruledesc
	//
	RuleDesc *string `field:"optional" json:"ruleDesc" yaml:"ruleDesc"`
}


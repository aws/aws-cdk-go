package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkspaceIpGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkspaceIpGroupMixinProps := &CfnWorkspaceIpGroupMixinProps{
//   	GroupDesc: jsii.String("groupDesc"),
//   	GroupName: jsii.String("groupName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserRules: []interface{}{
//   		&IpRuleItemProperty{
//   			IpRule: jsii.String("ipRule"),
//   			RuleDesc: jsii.String("ruleDesc"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspaceipgroup.html
//
type CfnWorkspaceIpGroupMixinProps struct {
	// The description of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspaceipgroup.html#cfn-workspaces-workspaceipgroup-groupdesc
	//
	GroupDesc *string `field:"optional" json:"groupDesc" yaml:"groupDesc"`
	// The name of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspaceipgroup.html#cfn-workspaces-workspaceipgroup-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The tags for the IP access control group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspaceipgroup.html#cfn-workspaces-workspaceipgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The rules for the IP access control group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspaceipgroup.html#cfn-workspaces-workspaceipgroup-userrules
	//
	UserRules interface{} `field:"optional" json:"userRules" yaml:"userRules"`
}


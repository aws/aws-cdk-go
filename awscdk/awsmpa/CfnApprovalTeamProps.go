package awsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApprovalTeam`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApprovalTeamProps := &CfnApprovalTeamProps{
//   	ApprovalStrategy: &ApprovalStrategyProperty{
//   		MofN: &MofNApprovalStrategyProperty{
//   			MinApprovalsRequired: jsii.Number(123),
//   		},
//   	},
//   	Approvers: []interface{}{
//   		&ApproverProperty{
//   			PrimaryIdentityId: jsii.String("primaryIdentityId"),
//   			PrimaryIdentitySourceArn: jsii.String("primaryIdentitySourceArn"),
//
//   			// the properties below are optional
//   			ApproverId: jsii.String("approverId"),
//   			PrimaryIdentityStatus: jsii.String("primaryIdentityStatus"),
//   			ResponseTime: jsii.String("responseTime"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policies: []interface{}{
//   		&PolicyProperty{
//   			PolicyArn: jsii.String("policyArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html
//
type CfnApprovalTeamProps struct {
	// Contains details for how an approval team grants approval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-approvalstrategy
	//
	ApprovalStrategy interface{} `field:"required" json:"approvalStrategy" yaml:"approvalStrategy"`
	// Contains details for an approver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-approvers
	//
	Approvers interface{} `field:"required" json:"approvers" yaml:"approvers"`
	// Description for the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains details for a policy.
	//
	// Policies define what operations a team that define the permissions for team resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-policies
	//
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// Tags that you have added to the specified resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


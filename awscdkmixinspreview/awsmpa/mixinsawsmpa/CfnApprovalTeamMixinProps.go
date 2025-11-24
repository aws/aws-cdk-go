package mixinsawsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApprovalTeamPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApprovalTeamMixinProps := &CfnApprovalTeamMixinProps{
//   	ApprovalStrategy: &ApprovalStrategyProperty{
//   		MofN: &MofNApprovalStrategyProperty{
//   			MinApprovalsRequired: jsii.Number(123),
//   		},
//   	},
//   	Approvers: []interface{}{
//   		&ApproverProperty{
//   			ApproverId: jsii.String("approverId"),
//   			PrimaryIdentityId: jsii.String("primaryIdentityId"),
//   			PrimaryIdentitySourceArn: jsii.String("primaryIdentitySourceArn"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html
//
type CfnApprovalTeamMixinProps struct {
	// Contains details for how an approval team grants approval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-approvalstrategy
	//
	ApprovalStrategy interface{} `field:"optional" json:"approvalStrategy" yaml:"approvalStrategy"`
	// Contains details for an approver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-approvers
	//
	Approvers interface{} `field:"optional" json:"approvers" yaml:"approvers"`
	// Description for the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains details for a policy.
	//
	// Policies define what operations a team that define the permissions for team resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// Tags that you have added to the specified resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html#cfn-mpa-approvalteam-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


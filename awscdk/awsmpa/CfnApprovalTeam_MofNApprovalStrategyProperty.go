package awsmpa


// Strategy for how an approval team grants approval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mofNApprovalStrategyProperty := &MofNApprovalStrategyProperty{
//   	MinApprovalsRequired: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-mofnapprovalstrategy.html
//
type CfnApprovalTeam_MofNApprovalStrategyProperty struct {
	// Minimum number of approvals (M) required for a total number of approvers (N).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-mofnapprovalstrategy.html#cfn-mpa-approvalteam-mofnapprovalstrategy-minapprovalsrequired
	//
	MinApprovalsRequired *float64 `field:"required" json:"minApprovalsRequired" yaml:"minApprovalsRequired"`
}


package previewawsmpamixins


// Strategy for how an approval team grants approval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   approvalStrategyProperty := &ApprovalStrategyProperty{
//   	MofN: &MofNApprovalStrategyProperty{
//   		MinApprovalsRequired: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approvalstrategy.html
//
type CfnApprovalTeamPropsMixin_ApprovalStrategyProperty struct {
	// Minimum number of approvals (M) required for a total number of approvers (N).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approvalstrategy.html#cfn-mpa-approvalteam-approvalstrategy-mofn
	//
	MofN interface{} `field:"optional" json:"mofN" yaml:"mofN"`
}


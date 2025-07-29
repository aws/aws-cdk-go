package awsmpa


// Contains details for an approver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approverProperty := &ApproverProperty{
//   	PrimaryIdentityId: jsii.String("primaryIdentityId"),
//   	PrimaryIdentitySourceArn: jsii.String("primaryIdentitySourceArn"),
//
//   	// the properties below are optional
//   	ApproverId: jsii.String("approverId"),
//   	PrimaryIdentityStatus: jsii.String("primaryIdentityStatus"),
//   	ResponseTime: jsii.String("responseTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html
//
type CfnApprovalTeam_ApproverProperty struct {
	// ID for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html#cfn-mpa-approvalteam-approver-primaryidentityid
	//
	PrimaryIdentityId *string `field:"required" json:"primaryIdentityId" yaml:"primaryIdentityId"`
	// Amazon Resource Name (ARN) for the identity source.
	//
	// The identity source manages the user authentication for approvers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html#cfn-mpa-approvalteam-approver-primaryidentitysourcearn
	//
	PrimaryIdentitySourceArn *string `field:"required" json:"primaryIdentitySourceArn" yaml:"primaryIdentitySourceArn"`
	// ID for the approver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html#cfn-mpa-approvalteam-approver-approverid
	//
	ApproverId *string `field:"optional" json:"approverId" yaml:"approverId"`
	// Status for the identity source.
	//
	// For example, if an approver has accepted a team invitation with a user authentication method managed by the identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html#cfn-mpa-approvalteam-approver-primaryidentitystatus
	//
	PrimaryIdentityStatus *string `field:"optional" json:"primaryIdentityStatus" yaml:"primaryIdentityStatus"`
	// Timestamp when the approver responded to an approval team invitation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-approvalteam-approver.html#cfn-mpa-approvalteam-approver-responsetime
	//
	ResponseTime *string `field:"optional" json:"responseTime" yaml:"responseTime"`
}


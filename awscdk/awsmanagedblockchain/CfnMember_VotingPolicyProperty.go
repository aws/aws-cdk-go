package awsmanagedblockchain


// The voting rules for the network to decide if a proposal is accepted.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   votingPolicyProperty := &VotingPolicyProperty{
//   	ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   		ProposalDurationInHours: jsii.Number(123),
//   		ThresholdComparator: jsii.String("thresholdComparator"),
//   		ThresholdPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-votingpolicy.html
//
type CfnMember_VotingPolicyProperty struct {
	// Defines the rules for the network for voting on proposals, such as the percentage of `YES` votes required for the proposal to be approved and the duration of the proposal.
	//
	// The policy applies to all proposals and is specified when the network is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-votingpolicy.html#cfn-managedblockchain-member-votingpolicy-approvalthresholdpolicy
	//
	ApprovalThresholdPolicy interface{} `field:"optional" json:"approvalThresholdPolicy" yaml:"approvalThresholdPolicy"`
}


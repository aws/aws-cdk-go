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
//   votingPolicyProperty := &votingPolicyProperty{
//   	approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   		proposalDurationInHours: jsii.Number(123),
//   		thresholdComparator: jsii.String("thresholdComparator"),
//   		thresholdPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnMember_VotingPolicyProperty struct {
	// Defines the rules for the network for voting on proposals, such as the percentage of `YES` votes required for the proposal to be approved and the duration of the proposal.
	//
	// The policy applies to all proposals and is specified when the network is created.
	ApprovalThresholdPolicy interface{} `field:"optional" json:"approvalThresholdPolicy" yaml:"approvalThresholdPolicy"`
}


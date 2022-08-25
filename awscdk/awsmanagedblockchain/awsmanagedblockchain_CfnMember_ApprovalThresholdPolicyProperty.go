package awsmanagedblockchain


// A policy type that defines the voting rules for the network.
//
// The rules decide if a proposal is approved. Approval may be based on criteria such as the percentage of `YES` votes and the duration of the proposal. The policy applies to all proposals and is specified when the network is created.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvalThresholdPolicyProperty := &approvalThresholdPolicyProperty{
//   	proposalDurationInHours: jsii.Number(123),
//   	thresholdComparator: jsii.String("thresholdComparator"),
//   	thresholdPercentage: jsii.Number(123),
//   }
//
type CfnMember_ApprovalThresholdPolicyProperty struct {
	// The duration from the time that a proposal is created until it expires.
	//
	// If members cast neither the required number of `YES` votes to approve the proposal nor the number of `NO` votes required to reject it before the duration expires, the proposal is `EXPIRED` and `ProposalActions` are not carried out.
	ProposalDurationInHours *float64 `field:"optional" json:"proposalDurationInHours" yaml:"proposalDurationInHours"`
	// Determines whether the vote percentage must be greater than the `ThresholdPercentage` or must be greater than or equal to the `ThreholdPercentage` to be approved.
	ThresholdComparator *string `field:"optional" json:"thresholdComparator" yaml:"thresholdComparator"`
	// The percentage of votes among all members that must be `YES` for a proposal to be approved.
	//
	// For example, a `ThresholdPercentage` value of `50` indicates 50%. The `ThresholdComparator` determines the precise comparison. If a `ThresholdPercentage` value of `50` is specified on a network with 10 members, along with a `ThresholdComparator` value of `GREATER_THAN` , this indicates that 6 `YES` votes are required for the proposal to be approved.
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}


package awsmanagedblockchain


// Configuration properties of the network to which the member belongs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	Framework: jsii.String("framework"),
//   	FrameworkVersion: jsii.String("frameworkVersion"),
//   	Name: jsii.String("name"),
//   	VotingPolicy: &VotingPolicyProperty{
//   		ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   			ProposalDurationInHours: jsii.Number(123),
//   			ThresholdComparator: jsii.String("thresholdComparator"),
//   			ThresholdPercentage: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	NetworkFrameworkConfiguration: &NetworkFrameworkConfigurationProperty{
//   		NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   			Edition: jsii.String("edition"),
//   		},
//   	},
//   }
//
type CfnMember_NetworkConfigurationProperty struct {
	// The blockchain framework that the network uses.
	Framework *string `field:"required" json:"framework" yaml:"framework"`
	// The version of the blockchain framework that the network uses.
	FrameworkVersion *string `field:"required" json:"frameworkVersion" yaml:"frameworkVersion"`
	// The name of the network.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The voting rules that the network uses to decide if a proposal is accepted.
	VotingPolicy interface{} `field:"required" json:"votingPolicy" yaml:"votingPolicy"`
	// Attributes of the blockchain framework for the network.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties relevant to the network for the blockchain framework that the network uses.
	NetworkFrameworkConfiguration interface{} `field:"optional" json:"networkFrameworkConfiguration" yaml:"networkFrameworkConfiguration"`
}


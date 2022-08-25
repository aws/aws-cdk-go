package awsmanagedblockchain


// Configuration properties of the network to which the member belongs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	framework: jsii.String("framework"),
//   	frameworkVersion: jsii.String("frameworkVersion"),
//   	name: jsii.String("name"),
//   	votingPolicy: &votingPolicyProperty{
//   		approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   			proposalDurationInHours: jsii.Number(123),
//   			thresholdComparator: jsii.String("thresholdComparator"),
//   			thresholdPercentage: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	networkFrameworkConfiguration: &networkFrameworkConfigurationProperty{
//   		networkFabricConfiguration: &networkFabricConfigurationProperty{
//   			edition: jsii.String("edition"),
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
	// The voting rules for the network to decide if a proposal is accepted.
	VotingPolicy interface{} `field:"required" json:"votingPolicy" yaml:"votingPolicy"`
	// Attributes of the blockchain framework for the network.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties relevant to the network for the blockchain framework that the network uses.
	NetworkFrameworkConfiguration interface{} `field:"optional" json:"networkFrameworkConfiguration" yaml:"networkFrameworkConfiguration"`
}


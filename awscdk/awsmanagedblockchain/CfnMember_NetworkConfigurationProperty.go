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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html
//
type CfnMember_NetworkConfigurationProperty struct {
	// The blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-framework
	//
	Framework *string `field:"required" json:"framework" yaml:"framework"`
	// The version of the blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-frameworkversion
	//
	FrameworkVersion *string `field:"required" json:"frameworkVersion" yaml:"frameworkVersion"`
	// The name of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The voting rules that the network uses to decide if a proposal is accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-votingpolicy
	//
	VotingPolicy interface{} `field:"required" json:"votingPolicy" yaml:"votingPolicy"`
	// Attributes of the blockchain framework for the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties relevant to the network for the blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-networkframeworkconfiguration
	//
	NetworkFrameworkConfiguration interface{} `field:"optional" json:"networkFrameworkConfiguration" yaml:"networkFrameworkConfiguration"`
}


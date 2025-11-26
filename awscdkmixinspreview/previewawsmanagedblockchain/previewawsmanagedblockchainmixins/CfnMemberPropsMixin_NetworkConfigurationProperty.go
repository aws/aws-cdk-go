package previewawsmanagedblockchainmixins


// Configuration properties of the network to which the member belongs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	Description: jsii.String("description"),
//   	Framework: jsii.String("framework"),
//   	FrameworkVersion: jsii.String("frameworkVersion"),
//   	Name: jsii.String("name"),
//   	NetworkFrameworkConfiguration: &NetworkFrameworkConfigurationProperty{
//   		NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   			Edition: jsii.String("edition"),
//   		},
//   	},
//   	VotingPolicy: &VotingPolicyProperty{
//   		ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   			ProposalDurationInHours: jsii.Number(123),
//   			ThresholdComparator: jsii.String("thresholdComparator"),
//   			ThresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html
//
type CfnMemberPropsMixin_NetworkConfigurationProperty struct {
	// Attributes of the blockchain framework for the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-framework
	//
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// The version of the blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-frameworkversion
	//
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// The name of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configuration properties relevant to the network for the blockchain framework that the network uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-networkframeworkconfiguration
	//
	NetworkFrameworkConfiguration interface{} `field:"optional" json:"networkFrameworkConfiguration" yaml:"networkFrameworkConfiguration"`
	// The voting rules that the network uses to decide if a proposal is accepted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkconfiguration.html#cfn-managedblockchain-member-networkconfiguration-votingpolicy
	//
	VotingPolicy interface{} `field:"optional" json:"votingPolicy" yaml:"votingPolicy"`
}


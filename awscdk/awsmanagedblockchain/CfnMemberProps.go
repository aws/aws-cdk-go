package awsmanagedblockchain


// Properties for defining a `CfnMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberProps := &CfnMemberProps{
//   	MemberConfiguration: &MemberConfigurationProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		MemberFrameworkConfiguration: &MemberFrameworkConfigurationProperty{
//   			MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   				AdminPassword: jsii.String("adminPassword"),
//   				AdminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	InvitationId: jsii.String("invitationId"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		Framework: jsii.String("framework"),
//   		FrameworkVersion: jsii.String("frameworkVersion"),
//   		Name: jsii.String("name"),
//   		VotingPolicy: &VotingPolicyProperty{
//   			ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   				ProposalDurationInHours: jsii.Number(123),
//   				ThresholdComparator: jsii.String("thresholdComparator"),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		NetworkFrameworkConfiguration: &NetworkFrameworkConfigurationProperty{
//   			NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   				Edition: jsii.String("edition"),
//   			},
//   		},
//   	},
//   	NetworkId: jsii.String("networkId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html
//
type CfnMemberProps struct {
	// Configuration properties of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-memberconfiguration
	//
	MemberConfiguration interface{} `field:"required" json:"memberConfiguration" yaml:"memberConfiguration"`
	// The unique identifier of the invitation to join the network sent to the account that creates the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-invitationid
	//
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Configuration properties of the network to which the member belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The unique identifier of the network to which the member belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-networkid
	//
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}


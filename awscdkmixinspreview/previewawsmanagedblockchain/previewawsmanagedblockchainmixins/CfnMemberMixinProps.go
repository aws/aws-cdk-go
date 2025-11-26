package previewawsmanagedblockchainmixins


// Properties for CfnMemberPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemberMixinProps := &CfnMemberMixinProps{
//   	InvitationId: jsii.String("invitationId"),
//   	MemberConfiguration: &MemberConfigurationProperty{
//   		Description: jsii.String("description"),
//   		MemberFrameworkConfiguration: &MemberFrameworkConfigurationProperty{
//   			MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   				AdminPassword: jsii.String("adminPassword"),
//   				AdminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		Description: jsii.String("description"),
//   		Framework: jsii.String("framework"),
//   		FrameworkVersion: jsii.String("frameworkVersion"),
//   		Name: jsii.String("name"),
//   		NetworkFrameworkConfiguration: &NetworkFrameworkConfigurationProperty{
//   			NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   				Edition: jsii.String("edition"),
//   			},
//   		},
//   		VotingPolicy: &VotingPolicyProperty{
//   			ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   				ProposalDurationInHours: jsii.Number(123),
//   				ThresholdComparator: jsii.String("thresholdComparator"),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	NetworkId: jsii.String("networkId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html
//
type CfnMemberMixinProps struct {
	// The unique identifier of the invitation to join the network sent to the account that creates the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-invitationid
	//
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Configuration properties of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-memberconfiguration
	//
	MemberConfiguration interface{} `field:"optional" json:"memberConfiguration" yaml:"memberConfiguration"`
	// Configuration properties of the network to which the member belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The unique identifier of the network to which the member belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html#cfn-managedblockchain-member-networkid
	//
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}


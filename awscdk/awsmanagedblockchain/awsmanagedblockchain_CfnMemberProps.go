package awsmanagedblockchain


// Properties for defining a `CfnMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberProps := &cfnMemberProps{
//   	memberConfiguration: &memberConfigurationProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		memberFrameworkConfiguration: &memberFrameworkConfigurationProperty{
//   			memberFabricConfiguration: &memberFabricConfigurationProperty{
//   				adminPassword: jsii.String("adminPassword"),
//   				adminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	invitationId: jsii.String("invitationId"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		framework: jsii.String("framework"),
//   		frameworkVersion: jsii.String("frameworkVersion"),
//   		name: jsii.String("name"),
//   		votingPolicy: &votingPolicyProperty{
//   			approvalThresholdPolicy: &approvalThresholdPolicyProperty{
//   				proposalDurationInHours: jsii.Number(123),
//   				thresholdComparator: jsii.String("thresholdComparator"),
//   				thresholdPercentage: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		networkFrameworkConfiguration: &networkFrameworkConfigurationProperty{
//   			networkFabricConfiguration: &networkFabricConfigurationProperty{
//   				edition: jsii.String("edition"),
//   			},
//   		},
//   	},
//   	networkId: jsii.String("networkId"),
//   }
//
type CfnMemberProps struct {
	// Configuration properties of the member.
	MemberConfiguration interface{} `field:"required" json:"memberConfiguration" yaml:"memberConfiguration"`
	// The unique identifier of the invitation to join the network sent to the account that creates the member.
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// Configuration properties of the network to which the member belongs.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The unique identifier of the network to which the member belongs.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}


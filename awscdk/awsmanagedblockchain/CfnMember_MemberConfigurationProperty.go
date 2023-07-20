package awsmanagedblockchain


// Configuration properties of the member.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberConfigurationProperty := &MemberConfigurationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MemberFrameworkConfiguration: &MemberFrameworkConfigurationProperty{
//   		MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   			AdminPassword: jsii.String("adminPassword"),
//   			AdminUsername: jsii.String("adminUsername"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html
//
type CfnMember_MemberConfigurationProperty struct {
	// The name of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration properties of the blockchain framework relevant to the member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberconfiguration.html#cfn-managedblockchain-member-memberconfiguration-memberframeworkconfiguration
	//
	MemberFrameworkConfiguration interface{} `field:"optional" json:"memberFrameworkConfiguration" yaml:"memberFrameworkConfiguration"`
}


package awsmanagedblockchain


// Configuration properties relevant to a member for the blockchain framework that the Managed Blockchain network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberFrameworkConfigurationProperty := &MemberFrameworkConfigurationProperty{
//   	MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   		AdminPassword: jsii.String("adminPassword"),
//   		AdminUsername: jsii.String("adminUsername"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberframeworkconfiguration.html
//
type CfnMember_MemberFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-memberframeworkconfiguration.html#cfn-managedblockchain-member-memberframeworkconfiguration-memberfabricconfiguration
	//
	MemberFabricConfiguration interface{} `field:"optional" json:"memberFabricConfiguration" yaml:"memberFabricConfiguration"`
}


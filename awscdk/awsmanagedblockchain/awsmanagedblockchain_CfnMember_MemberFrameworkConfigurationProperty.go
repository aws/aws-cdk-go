package awsmanagedblockchain


// Configuration properties relevant to a member for the blockchain framework that the Managed Blockchain network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberFrameworkConfigurationProperty := &memberFrameworkConfigurationProperty{
//   	memberFabricConfiguration: &memberFabricConfigurationProperty{
//   		adminPassword: jsii.String("adminPassword"),
//   		adminUsername: jsii.String("adminUsername"),
//   	},
//   }
//
type CfnMember_MemberFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric.
	MemberFabricConfiguration interface{} `field:"optional" json:"memberFabricConfiguration" yaml:"memberFabricConfiguration"`
}


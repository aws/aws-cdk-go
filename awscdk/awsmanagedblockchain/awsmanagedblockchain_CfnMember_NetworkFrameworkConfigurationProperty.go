package awsmanagedblockchain


// Configuration properties relevant to the network for the blockchain framework that the network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkFrameworkConfigurationProperty := &networkFrameworkConfigurationProperty{
//   	networkFabricConfiguration: &networkFabricConfigurationProperty{
//   		edition: jsii.String("edition"),
//   	},
//   }
//
type CfnMember_NetworkFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric for a member in a Managed Blockchain network using the Hyperledger Fabric framework.
	NetworkFabricConfiguration interface{} `field:"optional" json:"networkFabricConfiguration" yaml:"networkFabricConfiguration"`
}


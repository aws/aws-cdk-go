package awsmanagedblockchain


// Hyperledger Fabric configuration properties for the network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkFabricConfigurationProperty := &networkFabricConfigurationProperty{
//   	edition: jsii.String("edition"),
//   }
//
type CfnMember_NetworkFabricConfigurationProperty struct {
	// The edition of Amazon Managed Blockchain that the network uses.
	//
	// Valid values are `standard` and `starter` . For more information, see
	Edition *string `field:"required" json:"edition" yaml:"edition"`
}


package mixinsawsmanagedblockchain


// Hyperledger Fabric configuration properties for the network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkFabricConfigurationProperty := &NetworkFabricConfigurationProperty{
//   	Edition: jsii.String("edition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html
//
type CfnMemberPropsMixin_NetworkFabricConfigurationProperty struct {
	// The edition of Amazon Managed Blockchain that the network uses.
	//
	// Valid values are `standard` and `starter` . For more information, see [Amazon Managed Blockchain Pricing](https://docs.aws.amazon.com/managed-blockchain/pricing/)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html#cfn-managedblockchain-member-networkfabricconfiguration-edition
	//
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
}


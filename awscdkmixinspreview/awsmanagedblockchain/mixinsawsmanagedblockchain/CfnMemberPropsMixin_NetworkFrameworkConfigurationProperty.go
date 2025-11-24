package mixinsawsmanagedblockchain


// Configuration properties relevant to the network for the blockchain framework that the network uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkFrameworkConfigurationProperty := &NetworkFrameworkConfigurationProperty{
//   	NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   		Edition: jsii.String("edition"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkframeworkconfiguration.html
//
type CfnMemberPropsMixin_NetworkFrameworkConfigurationProperty struct {
	// Configuration properties for Hyperledger Fabric for a member in a Managed Blockchain network that is using the Hyperledger Fabric framework.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkframeworkconfiguration.html#cfn-managedblockchain-member-networkframeworkconfiguration-networkfabricconfiguration
	//
	NetworkFabricConfiguration interface{} `field:"optional" json:"networkFabricConfiguration" yaml:"networkFabricConfiguration"`
}


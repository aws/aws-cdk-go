package awsmanagedblockchain


// Properties for defining a `CfnNode`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNodeProps := &CfnNodeProps{
//   	NetworkId: jsii.String("networkId"),
//   	NodeConfiguration: &NodeConfigurationProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		InstanceType: jsii.String("instanceType"),
//   	},
//
//   	// the properties below are optional
//   	MemberId: jsii.String("memberId"),
//   }
//
type CfnNodeProps struct {
	// The unique identifier of the network for the node.
	//
	// Ethereum public networks have the following `NetworkId` s:
	//
	// - `n-ethereum-mainnet`
	// - `n-ethereum-goerli`
	// - `n-ethereum-rinkeby`.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Configuration properties of a peer node.
	NodeConfiguration interface{} `field:"required" json:"nodeConfiguration" yaml:"nodeConfiguration"`
	// The unique identifier of the member to which the node belongs.
	//
	// Applies only to Hyperledger Fabric.
	MemberId *string `field:"optional" json:"memberId" yaml:"memberId"`
}


package previewawsmanagedblockchainmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAccessorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessorMixinProps := &CfnAccessorMixinProps{
//   	AccessorType: jsii.String("accessorType"),
//   	NetworkType: jsii.String("networkType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-accessor.html
//
type CfnAccessorMixinProps struct {
	// The type of the accessor.
	//
	// > Currently, accessor type is restricted to `BILLING_TOKEN` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-accessor.html#cfn-managedblockchain-accessor-accessortype
	//
	AccessorType *string `field:"optional" json:"accessorType" yaml:"accessorType"`
	// The blockchain network that the `Accessor` token is created for.
	//
	// > We recommend using the appropriate `networkType` value for the blockchain network that you are creating the `Accessor` token for. You cannot use the value `ETHEREUM_MAINNET_AND_GOERLI` to specify a `networkType` for your Accessor token.
	// >
	// > The default value of `ETHEREUM_MAINNET_AND_GOERLI` is only applied:
	// >
	// > - when the `CreateAccessor` action does not set a `networkType` .
	// > - to all existing `Accessor` tokens that were created before the `networkType` property was introduced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-accessor.html#cfn-managedblockchain-accessor-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The tags assigned to the Accessor.
	//
	// For more information about tags, see [Tagging Resources](https://docs.aws.amazon.com/managed-blockchain/latest/ethereum-dev/tagging-resources.html) in the *Amazon Managed Blockchain Ethereum Developer Guide* , or [Tagging Resources](https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html) in the *Amazon Managed Blockchain Hyperledger Fabric Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-accessor.html#cfn-managedblockchain-accessor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


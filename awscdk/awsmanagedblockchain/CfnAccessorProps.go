package awsmanagedblockchain

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAccessor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessorProps := &CfnAccessorProps{
//   	AccessorType: jsii.String("accessorType"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAccessorProps struct {
	// The type of the accessor.
	//
	// > Currently, accessor type is restricted to `BILLING_TOKEN` .
	AccessorType *string `field:"required" json:"accessorType" yaml:"accessorType"`
	// The tags assigned to the Accessor.
	//
	// For more information about tags, see [Tagging Resources](https://docs.aws.amazon.com/managed-blockchain/latest/ethereum-dev/tagging-resources.html) in the *Amazon Managed Blockchain Ethereum Developer Guide* , or [Tagging Resources](https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html) in the *Amazon Managed Blockchain Hyperledger Fabric Developer Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


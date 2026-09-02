package awswickr


// Properties for CfnNetworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNetworkMixinProps := &CfnNetworkMixinProps{
//   	AccessLevel: jsii.String("accessLevel"),
//   	NetworkName: jsii.String("networkName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wickr-network.html
//
type CfnNetworkMixinProps struct {
	// The access level of the network, which determines available features and capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wickr-network.html#cfn-wickr-network-accesslevel
	//
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The name of the network.
	//
	// Must be between 1 and 20 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wickr-network.html#cfn-wickr-network-networkname
	//
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
}


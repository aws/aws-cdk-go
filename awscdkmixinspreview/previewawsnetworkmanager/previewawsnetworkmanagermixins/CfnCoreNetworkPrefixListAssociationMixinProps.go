package previewawsnetworkmanagermixins


// Properties for CfnCoreNetworkPrefixListAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCoreNetworkPrefixListAssociationMixinProps := &CfnCoreNetworkPrefixListAssociationMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	PrefixListAlias: jsii.String("prefixListAlias"),
//   	PrefixListArn: jsii.String("prefixListArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetworkprefixlistassociation.html
//
type CfnCoreNetworkPrefixListAssociationMixinProps struct {
	// The ID of the core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetworkprefixlistassociation.html#cfn-networkmanager-corenetworkprefixlistassociation-corenetworkid
	//
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The alias of the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetworkprefixlistassociation.html#cfn-networkmanager-corenetworkprefixlistassociation-prefixlistalias
	//
	PrefixListAlias *string `field:"optional" json:"prefixListAlias" yaml:"prefixListAlias"`
	// The Amazon Resource Name (ARN) of the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetworkprefixlistassociation.html#cfn-networkmanager-corenetworkprefixlistassociation-prefixlistarn
	//
	PrefixListArn *string `field:"optional" json:"prefixListArn" yaml:"prefixListArn"`
}


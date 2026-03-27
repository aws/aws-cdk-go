package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIPAMPrefixListResolverTargetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnIPAMPrefixListResolverTargetMixinProps := &CfnIPAMPrefixListResolverTargetMixinProps{
//   	DesiredVersion: jsii.Number(123),
//   	IpamPrefixListResolverId: jsii.String("ipamPrefixListResolverId"),
//   	PrefixListId: jsii.String("prefixListId"),
//   	PrefixListRegion: jsii.String("prefixListRegion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackLatestVersion: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html
//
type CfnIPAMPrefixListResolverTargetMixinProps struct {
	// The desired version of the Prefix List Resolver that this Target should synchronize with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-desiredversion
	//
	DesiredVersion *float64 `field:"optional" json:"desiredVersion" yaml:"desiredVersion"`
	// The Id of the IPAM Prefix List Resolver associated with this Target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-ipamprefixlistresolverid
	//
	IpamPrefixListResolverId *string `field:"optional" json:"ipamPrefixListResolverId" yaml:"ipamPrefixListResolverId"`
	// The Id of the Managed Prefix List.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-prefixlistid
	//
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// The region that the Managed Prefix List is located in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-prefixlistregion
	//
	PrefixListRegion *string `field:"optional" json:"prefixListRegion" yaml:"prefixListRegion"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether this Target automatically tracks the latest version of the Prefix List Resolver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolvertarget.html#cfn-ec2-ipamprefixlistresolvertarget-tracklatestversion
	//
	TrackLatestVersion interface{} `field:"optional" json:"trackLatestVersion" yaml:"trackLatestVersion"`
}


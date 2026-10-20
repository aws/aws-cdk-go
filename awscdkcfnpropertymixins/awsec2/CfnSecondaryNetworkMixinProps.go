package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSecondaryNetworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSecondaryNetworkMixinProps := &CfnSecondaryNetworkMixinProps{
//   	Ipv4CidrBlock: jsii.String("ipv4CidrBlock"),
//   	NetworkType: jsii.String("networkType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-secondarynetwork.html
//
type CfnSecondaryNetworkMixinProps struct {
	// The IPv4 CIDR block for the secondary network.
	//
	// The CIDR block size must be between /12 and /28.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-secondarynetwork.html#cfn-ec2-secondarynetwork-ipv4cidrblock
	//
	Ipv4CidrBlock *string `field:"optional" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The type of secondary network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-secondarynetwork.html#cfn-ec2-secondarynetwork-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The tags for the secondary network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-secondarynetwork.html#cfn-ec2-secondarynetwork-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


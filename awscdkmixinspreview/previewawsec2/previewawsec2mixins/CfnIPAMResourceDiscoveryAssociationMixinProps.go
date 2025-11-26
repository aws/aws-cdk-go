package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIPAMResourceDiscoveryAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIPAMResourceDiscoveryAssociationMixinProps := &CfnIPAMResourceDiscoveryAssociationMixinProps{
//   	IpamId: jsii.String("ipamId"),
//   	IpamResourceDiscoveryId: jsii.String("ipamResourceDiscoveryId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscoveryassociation.html
//
type CfnIPAMResourceDiscoveryAssociationMixinProps struct {
	// The IPAM ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscoveryassociation.html#cfn-ec2-ipamresourcediscoveryassociation-ipamid
	//
	IpamId *string `field:"optional" json:"ipamId" yaml:"ipamId"`
	// The resource discovery ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscoveryassociation.html#cfn-ec2-ipamresourcediscoveryassociation-ipamresourcediscoveryid
	//
	IpamResourceDiscoveryId *string `field:"optional" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
	// A tag is a label that you assign to an AWS resource.
	//
	// Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscoveryassociation.html#cfn-ec2-ipamresourcediscoveryassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


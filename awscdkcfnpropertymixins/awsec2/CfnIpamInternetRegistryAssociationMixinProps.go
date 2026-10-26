package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIpamInternetRegistryAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnIpamInternetRegistryAssociationMixinProps := &CfnIpamInternetRegistryAssociationMixinProps{
//   	Description: jsii.String("description"),
//   	IpamId: jsii.String("ipamId"),
//   	OrganizationHandle: jsii.String("organizationHandle"),
//   	Rir: jsii.String("rir"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html
//
type CfnIpamInternetRegistryAssociationMixinProps struct {
	// A description for the internet registry association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html#cfn-ec2-ipaminternetregistryassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the IPAM to associate with the internet registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html#cfn-ec2-ipaminternetregistryassociation-ipamid
	//
	IpamId *string `field:"optional" json:"ipamId" yaml:"ipamId"`
	// The organization handle at the internet registry (for example, a RIPE NCC organization ID or ARIN Org ID).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html#cfn-ec2-ipaminternetregistryassociation-organizationhandle
	//
	OrganizationHandle *string `field:"optional" json:"organizationHandle" yaml:"organizationHandle"`
	// The Regional Internet Registry to associate with.
	//
	// The service returns these values in upper case, so the enum uses the round-tripping form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html#cfn-ec2-ipaminternetregistryassociation-rir
	//
	Rir *string `field:"optional" json:"rir" yaml:"rir"`
	// The tags assigned to the internet registry association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipaminternetregistryassociation.html#cfn-ec2-ipaminternetregistryassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


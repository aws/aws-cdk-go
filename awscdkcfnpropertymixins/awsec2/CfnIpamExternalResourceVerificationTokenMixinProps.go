package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIpamExternalResourceVerificationTokenPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnIpamExternalResourceVerificationTokenMixinProps := &CfnIpamExternalResourceVerificationTokenMixinProps{
//   	IpamId: jsii.String("ipamId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamexternalresourceverificationtoken.html
//
type CfnIpamExternalResourceVerificationTokenMixinProps struct {
	// The ID of the IPAM that will create the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamexternalresourceverificationtoken.html#cfn-ec2-ipamexternalresourceverificationtoken-ipamid
	//
	IpamId *string `field:"optional" json:"ipamId" yaml:"ipamId"`
	// The tags for the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamexternalresourceverificationtoken.html#cfn-ec2-ipamexternalresourceverificationtoken-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


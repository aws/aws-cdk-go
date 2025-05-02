package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnectionGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionGroupProps := &CfnConnectionGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AnycastIpListId: jsii.String("anycastIpListId"),
//   	Enabled: jsii.Boolean(false),
//   	Ipv6Enabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html
//
type CfnConnectionGroupProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-anycastiplistid
	//
	AnycastIpListId *string `field:"optional" json:"anycastIpListId" yaml:"anycastIpListId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-ipv6enabled
	//
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


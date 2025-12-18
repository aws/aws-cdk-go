package awscloudfront


// Properties for defining a `CfnAnycastIpList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnycastIpListProps := &CfnAnycastIpListProps{
//   	IpCount: jsii.Number(123),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	IpAddressType: jsii.String("ipAddressType"),
//   	IpamCidrConfigs: []interface{}{
//   		&IpamCidrConfigProperty{
//   			Cidr: jsii.String("cidr"),
//   			IpamPoolArn: jsii.String("ipamPoolArn"),
//   		},
//   	},
//   	Tags: &TagsProperty{
//   		Items: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html
//
type CfnAnycastIpListProps struct {
	// The number of IP addresses in the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html#cfn-cloudfront-anycastiplist-ipcount
	//
	IpCount *float64 `field:"required" json:"ipCount" yaml:"ipCount"`
	// The name of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html#cfn-cloudfront-anycastiplist-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IP address type for the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html#cfn-cloudfront-anycastiplist-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html#cfn-cloudfront-anycastiplist-ipamcidrconfigs
	//
	IpamCidrConfigs interface{} `field:"optional" json:"ipamCidrConfigs" yaml:"ipamCidrConfigs"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html#cfn-cloudfront-anycastiplist-tags
	//
	Tags *CfnAnycastIpList_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}


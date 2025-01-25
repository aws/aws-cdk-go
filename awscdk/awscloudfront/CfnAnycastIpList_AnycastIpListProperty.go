package awscloudfront


// An Anycast static IP list.
//
// For more information, see [Request Anycast static IPs to use for allowlisting](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/request-static-ips.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anycastIpListProperty := &AnycastIpListProperty{
//   	AnycastIps: []*string{
//   		jsii.String("anycastIps"),
//   	},
//   	Arn: jsii.String("arn"),
//   	Id: jsii.String("id"),
//   	IpCount: jsii.Number(123),
//   	LastModifiedTime: jsii.String("lastModifiedTime"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html
//
type CfnAnycastIpList_AnycastIpListProperty struct {
	// The static IP addresses that are allocated to the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-anycastips
	//
	AnycastIps *[]*string `field:"required" json:"anycastIps" yaml:"anycastIps"`
	// The Amazon Resource Name (ARN) of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The ID of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The number of IP addresses in the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-ipcount
	//
	IpCount *float64 `field:"required" json:"ipCount" yaml:"ipCount"`
	// The last time the Anycast static IP list was modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-lastmodifiedtime
	//
	LastModifiedTime *string `field:"required" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// The name of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The status of the Anycast static IP list.
	//
	// Valid values: `Deployed` , `Deploying` , or `Failed` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-anycastiplist.html#cfn-cloudfront-anycastiplist-anycastiplist-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}


package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamCidrConfigResultProperty := &IpamCidrConfigResultProperty{
//   	AnycastIp: jsii.String("anycastIp"),
//   	Cidr: jsii.String("cidr"),
//   	IpamPoolArn: jsii.String("ipamPoolArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html
//
type CfnAnycastIpList_IpamCidrConfigResultProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-anycastip
	//
	AnycastIp *string `field:"optional" json:"anycastIp" yaml:"anycastIp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-ipampoolarn
	//
	IpamPoolArn *string `field:"optional" json:"ipamPoolArn" yaml:"ipamPoolArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}


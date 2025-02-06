package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOriginConfigProperty := &VpcOriginConfigProperty{
//   	VpcOriginId: jsii.String("vpcOriginId"),
//
//   	// the properties below are optional
//   	OriginKeepaliveTimeout: jsii.Number(123),
//   	OriginReadTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html
//
type CfnDistribution_VpcOriginConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-vpcoriginid
	//
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-originkeepalivetimeout
	//
	// Default: - 5.
	//
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-originreadtimeout
	//
	// Default: - 30.
	//
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
}


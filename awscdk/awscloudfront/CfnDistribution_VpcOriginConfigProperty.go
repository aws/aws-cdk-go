package awscloudfront


// An Amazon CloudFront VPC origin configuration.
//
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
//   	OwnerAccountId: jsii.String("ownerAccountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html
//
type CfnDistribution_VpcOriginConfigProperty struct {
	// The VPC origin ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-vpcoriginid
	//
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is 5 seconds.
	//
	// For more information, see [Keep-alive timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-originkeepalivetimeout
	//
	// Default: - 5.
	//
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin.
	//
	// This is also known as the *origin response timeout* . The minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is 30 seconds.
	//
	// For more information, see [Response timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-originreadtimeout
	//
	// Default: - 30.
	//
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// The account ID of the AWS account that owns the VPC origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-vpcoriginconfig.html#cfn-cloudfront-distribution-vpcoriginconfig-owneraccountid
	//
	OwnerAccountId *string `field:"optional" json:"ownerAccountId" yaml:"ownerAccountId"`
}


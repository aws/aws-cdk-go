package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheTagConfigProperty := &CacheTagConfigProperty{
//   	HeaderName: jsii.String("headerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-cachetagconfig.html
//
type CfnDistribution_CacheTagConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-cachetagconfig.html#cfn-cloudfront-distribution-cachetagconfig-headername
	//
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}


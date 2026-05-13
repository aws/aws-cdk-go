package awscloudfront


// A complex type that specifies the HTTP header name from which CloudFront extracts cache tags from origin responses.
//
// When you add ``CacheTagConfig`` to a distribution, CloudFront reads the specified header from origin responses, parses the comma-separated tag values, and stores them with the cached object. You can then invalidate cached objects by tag using the ``CreateInvalidation`` API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cacheTagConfigProperty := &CacheTagConfigProperty{
//   	HeaderName: jsii.String("headerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-cachetagconfig.html
//
type CfnDistributionPropsMixin_CacheTagConfigProperty struct {
	// The name of the HTTP header that your origin includes in responses.
	//
	// CloudFront uses this header to extract cache tags. The header value must contain comma-separated tag values (for example, ``product:electronics, category:tv, brand:example``).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-cachetagconfig.html#cfn-cloudfront-distribution-cachetagconfig-headername
	//
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
}


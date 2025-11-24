package mixinsawscloudfront


// Properties for CfnCachePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCachePolicyMixinProps := &CfnCachePolicyMixinProps{
//   	CachePolicyConfig: &CachePolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		DefaultTtl: jsii.Number(123),
//   		MaxTtl: jsii.Number(123),
//   		MinTtl: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		ParametersInCacheKeyAndForwardedToOrigin: &ParametersInCacheKeyAndForwardedToOriginProperty{
//   			CookiesConfig: &CookiesConfigProperty{
//   				CookieBehavior: jsii.String("cookieBehavior"),
//   				Cookies: []*string{
//   					jsii.String("cookies"),
//   				},
//   			},
//   			EnableAcceptEncodingBrotli: jsii.Boolean(false),
//   			EnableAcceptEncodingGzip: jsii.Boolean(false),
//   			HeadersConfig: &HeadersConfigProperty{
//   				HeaderBehavior: jsii.String("headerBehavior"),
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   			},
//   			QueryStringsConfig: &QueryStringsConfigProperty{
//   				QueryStringBehavior: jsii.String("queryStringBehavior"),
//   				QueryStrings: []*string{
//   					jsii.String("queryStrings"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cachepolicy.html
//
type CfnCachePolicyMixinProps struct {
	// The cache policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cachepolicy.html#cfn-cloudfront-cachepolicy-cachepolicyconfig
	//
	CachePolicyConfig interface{} `field:"optional" json:"cachePolicyConfig" yaml:"cachePolicyConfig"`
}


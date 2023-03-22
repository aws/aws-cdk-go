package awscloudfront


// Properties for defining a `CfnCachePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCachePolicyProps := &CfnCachePolicyProps{
//   	CachePolicyConfig: &CachePolicyConfigProperty{
//   		DefaultTtl: jsii.Number(123),
//   		MaxTtl: jsii.Number(123),
//   		MinTtl: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		ParametersInCacheKeyAndForwardedToOrigin: &ParametersInCacheKeyAndForwardedToOriginProperty{
//   			CookiesConfig: &CookiesConfigProperty{
//   				CookieBehavior: jsii.String("cookieBehavior"),
//
//   				// the properties below are optional
//   				Cookies: []*string{
//   					jsii.String("cookies"),
//   				},
//   			},
//   			EnableAcceptEncodingGzip: jsii.Boolean(false),
//   			HeadersConfig: &HeadersConfigProperty{
//   				HeaderBehavior: jsii.String("headerBehavior"),
//
//   				// the properties below are optional
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   			},
//   			QueryStringsConfig: &QueryStringsConfigProperty{
//   				QueryStringBehavior: jsii.String("queryStringBehavior"),
//
//   				// the properties below are optional
//   				QueryStrings: []*string{
//   					jsii.String("queryStrings"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			EnableAcceptEncodingBrotli: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
type CfnCachePolicyProps struct {
	// The cache policy configuration.
	CachePolicyConfig interface{} `field:"required" json:"cachePolicyConfig" yaml:"cachePolicyConfig"`
}


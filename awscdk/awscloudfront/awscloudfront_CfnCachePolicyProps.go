package awscloudfront


// Properties for defining a `CfnCachePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCachePolicyProps := &cfnCachePolicyProps{
//   	cachePolicyConfig: &cachePolicyConfigProperty{
//   		defaultTtl: jsii.Number(123),
//   		maxTtl: jsii.Number(123),
//   		minTtl: jsii.Number(123),
//   		name: jsii.String("name"),
//   		parametersInCacheKeyAndForwardedToOrigin: &parametersInCacheKeyAndForwardedToOriginProperty{
//   			cookiesConfig: &cookiesConfigProperty{
//   				cookieBehavior: jsii.String("cookieBehavior"),
//
//   				// the properties below are optional
//   				cookies: []*string{
//   					jsii.String("cookies"),
//   				},
//   			},
//   			enableAcceptEncodingGzip: jsii.Boolean(false),
//   			headersConfig: &headersConfigProperty{
//   				headerBehavior: jsii.String("headerBehavior"),
//
//   				// the properties below are optional
//   				headers: []*string{
//   					jsii.String("headers"),
//   				},
//   			},
//   			queryStringsConfig: &queryStringsConfigProperty{
//   				queryStringBehavior: jsii.String("queryStringBehavior"),
//
//   				// the properties below are optional
//   				queryStrings: []*string{
//   					jsii.String("queryStrings"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			enableAcceptEncodingBrotli: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		comment: jsii.String("comment"),
//   	},
//   }
//
type CfnCachePolicyProps struct {
	// The cache policy configuration.
	CachePolicyConfig interface{} `field:"required" json:"cachePolicyConfig" yaml:"cachePolicyConfig"`
}


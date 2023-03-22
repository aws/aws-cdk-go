package awscloudfront


// Properties for defining a `CfnOriginRequestPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginRequestPolicyProps := &CfnOriginRequestPolicyProps{
//   	OriginRequestPolicyConfig: &OriginRequestPolicyConfigProperty{
//   		CookiesConfig: &CookiesConfigProperty{
//   			CookieBehavior: jsii.String("cookieBehavior"),
//
//   			// the properties below are optional
//   			Cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		HeadersConfig: &HeadersConfigProperty{
//   			HeaderBehavior: jsii.String("headerBehavior"),
//
//   			// the properties below are optional
//   			Headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		QueryStringsConfig: &QueryStringsConfigProperty{
//   			QueryStringBehavior: jsii.String("queryStringBehavior"),
//
//   			// the properties below are optional
//   			QueryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
type CfnOriginRequestPolicyProps struct {
	// The origin request policy configuration.
	OriginRequestPolicyConfig interface{} `field:"required" json:"originRequestPolicyConfig" yaml:"originRequestPolicyConfig"`
}


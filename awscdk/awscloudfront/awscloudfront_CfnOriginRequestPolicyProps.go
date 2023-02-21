package awscloudfront


// Properties for defining a `CfnOriginRequestPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginRequestPolicyProps := &cfnOriginRequestPolicyProps{
//   	originRequestPolicyConfig: &originRequestPolicyConfigProperty{
//   		cookiesConfig: &cookiesConfigProperty{
//   			cookieBehavior: jsii.String("cookieBehavior"),
//
//   			// the properties below are optional
//   			cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		headersConfig: &headersConfigProperty{
//   			headerBehavior: jsii.String("headerBehavior"),
//
//   			// the properties below are optional
//   			headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		name: jsii.String("name"),
//   		queryStringsConfig: &queryStringsConfigProperty{
//   			queryStringBehavior: jsii.String("queryStringBehavior"),
//
//   			// the properties below are optional
//   			queryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		comment: jsii.String("comment"),
//   	},
//   }
//
type CfnOriginRequestPolicyProps struct {
	// The origin request policy configuration.
	OriginRequestPolicyConfig interface{} `field:"required" json:"originRequestPolicyConfig" yaml:"originRequestPolicyConfig"`
}


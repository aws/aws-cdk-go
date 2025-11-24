package mixinsawscloudfront


// Properties for CfnOriginRequestPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOriginRequestPolicyMixinProps := &CfnOriginRequestPolicyMixinProps{
//   	OriginRequestPolicyConfig: &OriginRequestPolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		CookiesConfig: &CookiesConfigProperty{
//   			CookieBehavior: jsii.String("cookieBehavior"),
//   			Cookies: []*string{
//   				jsii.String("cookies"),
//   			},
//   		},
//   		HeadersConfig: &HeadersConfigProperty{
//   			HeaderBehavior: jsii.String("headerBehavior"),
//   			Headers: []*string{
//   				jsii.String("headers"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		QueryStringsConfig: &QueryStringsConfigProperty{
//   			QueryStringBehavior: jsii.String("queryStringBehavior"),
//   			QueryStrings: []*string{
//   				jsii.String("queryStrings"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html
//
type CfnOriginRequestPolicyMixinProps struct {
	// The origin request policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig
	//
	OriginRequestPolicyConfig interface{} `field:"optional" json:"originRequestPolicyConfig" yaml:"originRequestPolicyConfig"`
}


package awss3


// Rule that define when a redirect is applied and the redirect behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var replaceKey replaceKey
//
//   routingRule := &RoutingRule{
//   	Condition: &RoutingRuleCondition{
//   		HttpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   		KeyPrefixEquals: jsii.String("keyPrefixEquals"),
//   	},
//   	HostName: jsii.String("hostName"),
//   	HttpRedirectCode: jsii.String("httpRedirectCode"),
//   	Protocol: awscdk.Aws_s3.RedirectProtocol_HTTP,
//   	ReplaceKey: replaceKey,
//   }
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html
//
type RoutingRule struct {
	// Specifies a condition that must be met for the specified redirect to apply.
	Condition *RoutingRuleCondition `field:"optional" json:"condition" yaml:"condition"`
	// The host name to use in the redirect request.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response.
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests.
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Specifies the object key prefix to use in the redirect request.
	ReplaceKey ReplaceKey `field:"optional" json:"replaceKey" yaml:"replaceKey"`
}


package awss3


// Specifies the redirect behavior and when a redirect is applied.
//
// For more information about routing rules, see [Configuring advanced conditional redirects](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingRuleProperty := &routingRuleProperty{
//   	redirectRule: &redirectRuleProperty{
//   		hostName: jsii.String("hostName"),
//   		httpRedirectCode: jsii.String("httpRedirectCode"),
//   		protocol: jsii.String("protocol"),
//   		replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   		replaceKeyWith: jsii.String("replaceKeyWith"),
//   	},
//
//   	// the properties below are optional
//   	routingRuleCondition: &routingRuleConditionProperty{
//   		httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   		keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   	},
//   }
//
type CfnBucket_RoutingRuleProperty struct {
	// Container for redirect information.
	//
	// You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
	RedirectRule interface{} `field:"required" json:"redirectRule" yaml:"redirectRule"`
	// A container for describing a condition that must be met for the specified redirect to apply.
	//
	// For example, 1. If request is for pages in the `/docs` folder, redirect to the `/documents` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
	RoutingRuleCondition interface{} `field:"optional" json:"routingRuleCondition" yaml:"routingRuleCondition"`
}


package awss3


// Specifies website configuration parameters for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   websiteConfigurationProperty := &websiteConfigurationProperty{
//   	errorDocument: jsii.String("errorDocument"),
//   	indexDocument: jsii.String("indexDocument"),
//   	redirectAllRequestsTo: &redirectAllRequestsToProperty{
//   		hostName: jsii.String("hostName"),
//
//   		// the properties below are optional
//   		protocol: jsii.String("protocol"),
//   	},
//   	routingRules: []interface{}{
//   		&routingRuleProperty{
//   			redirectRule: &redirectRuleProperty{
//   				hostName: jsii.String("hostName"),
//   				httpRedirectCode: jsii.String("httpRedirectCode"),
//   				protocol: jsii.String("protocol"),
//   				replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   				replaceKeyWith: jsii.String("replaceKeyWith"),
//   			},
//
//   			// the properties below are optional
//   			routingRuleCondition: &routingRuleConditionProperty{
//   				httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   				keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_WebsiteConfigurationProperty struct {
	// The name of the error document for the website.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// The name of the index document for the website.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// The redirect behavior for every request to this bucket's website endpoint.
	//
	// > If you specify this property, you can't specify any other property.
	RedirectAllRequestsTo interface{} `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Rules that define when a redirect is applied and the redirect behavior.
	RoutingRules interface{} `field:"optional" json:"routingRules" yaml:"routingRules"`
}


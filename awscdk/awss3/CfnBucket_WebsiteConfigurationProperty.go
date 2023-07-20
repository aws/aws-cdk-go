package awss3


// Specifies website configuration parameters for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   websiteConfigurationProperty := &WebsiteConfigurationProperty{
//   	ErrorDocument: jsii.String("errorDocument"),
//   	IndexDocument: jsii.String("indexDocument"),
//   	RedirectAllRequestsTo: &RedirectAllRequestsToProperty{
//   		HostName: jsii.String("hostName"),
//
//   		// the properties below are optional
//   		Protocol: jsii.String("protocol"),
//   	},
//   	RoutingRules: []interface{}{
//   		&RoutingRuleProperty{
//   			RedirectRule: &RedirectRuleProperty{
//   				HostName: jsii.String("hostName"),
//   				HttpRedirectCode: jsii.String("httpRedirectCode"),
//   				Protocol: jsii.String("protocol"),
//   				ReplaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   				ReplaceKeyWith: jsii.String("replaceKeyWith"),
//   			},
//
//   			// the properties below are optional
//   			RoutingRuleCondition: &RoutingRuleConditionProperty{
//   				HttpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   				KeyPrefixEquals: jsii.String("keyPrefixEquals"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-websiteconfiguration.html
//
type CfnBucket_WebsiteConfigurationProperty struct {
	// The name of the error document for the website.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-websiteconfiguration.html#cfn-s3-bucket-websiteconfiguration-errordocument
	//
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// The name of the index document for the website.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-websiteconfiguration.html#cfn-s3-bucket-websiteconfiguration-indexdocument
	//
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// The redirect behavior for every request to this bucket's website endpoint.
	//
	// > If you specify this property, you can't specify any other property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-websiteconfiguration.html#cfn-s3-bucket-websiteconfiguration-redirectallrequeststo
	//
	RedirectAllRequestsTo interface{} `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Rules that define when a redirect is applied and the redirect behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-websiteconfiguration.html#cfn-s3-bucket-websiteconfiguration-routingrules
	//
	RoutingRules interface{} `field:"optional" json:"routingRules" yaml:"routingRules"`
}


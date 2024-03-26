package awswafv2


// Specifies custom configurations for the associations between the web ACL and protected resources.
//
// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
//
// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
//
// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationConfigProperty := &AssociationConfigProperty{
//   	RequestBody: map[string]interface{}{
//   		"requestBodyKey": &RequestBodyAssociatedResourceTypeConfigProperty{
//   			"defaultSizeInspectionLimit": jsii.String("defaultSizeInspectionLimit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-associationconfig.html
//
type CfnWebACL_AssociationConfigProperty struct {
	// Customizes the maximum size of the request body that your protected CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access resources forward to AWS WAF for inspection.
	//
	// The default size is 16 KB (16,384 bytes). You can change the setting for any of the available resource types.
	//
	// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
	//
	// Example JSON: `{ "API_GATEWAY": "KB_48", "APP_RUNNER_SERVICE": "KB_32" }`
	//
	// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-associationconfig.html#cfn-wafv2-webacl-associationconfig-requestbody
	//
	RequestBody interface{} `field:"optional" json:"requestBody" yaml:"requestBody"`
}


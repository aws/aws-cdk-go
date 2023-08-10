package awswafv2


// Specifies custom configurations for the associations between the web ACL and protected resources.
//
// Use this to customize the maximum size of the request body that your protected CloudFront distributions forward to AWS WAF for inspection. The default is 16 KB (16,384 bytes).
//
// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
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
	// Customizes the maximum size of the request body that your protected CloudFront distributions forward to AWS WAF for inspection.
	//
	// The default size is 16 KB (16,384 bytes).
	//
	// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-associationconfig.html#cfn-wafv2-webacl-associationconfig-requestbody
	//
	RequestBody interface{} `field:"optional" json:"requestBody" yaml:"requestBody"`
}


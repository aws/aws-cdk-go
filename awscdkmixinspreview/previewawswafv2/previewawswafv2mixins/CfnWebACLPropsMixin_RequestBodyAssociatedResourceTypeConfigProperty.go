package previewawswafv2mixins


// Customizes the maximum size of the request body that your protected CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access resources forward to AWS WAF for inspection.
//
// The default size is 16 KB (16,384 bytes). You can change the setting for any of the available resource types.
//
// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
//
// Example JSON: `{ "API_GATEWAY": "KB_48", "APP_RUNNER_SERVICE": "KB_32" }`
//
// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
//
// This is used in the `AssociationConfig` of the web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestBodyAssociatedResourceTypeConfigProperty := &RequestBodyAssociatedResourceTypeConfigProperty{
//   	DefaultSizeInspectionLimit: jsii.String("defaultSizeInspectionLimit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestbodyassociatedresourcetypeconfig.html
//
type CfnWebACLPropsMixin_RequestBodyAssociatedResourceTypeConfigProperty struct {
	// Specifies the maximum size of the web request body component that an associated CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resource should send to AWS WAF for inspection.
	//
	// This applies to statements in the web ACL that inspect the body or JSON body.
	//
	// Default: `16 KB (16,384 bytes)`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestbodyassociatedresourcetypeconfig.html#cfn-wafv2-webacl-requestbodyassociatedresourcetypeconfig-defaultsizeinspectionlimit
	//
	DefaultSizeInspectionLimit *string `field:"optional" json:"defaultSizeInspectionLimit" yaml:"defaultSizeInspectionLimit"`
}


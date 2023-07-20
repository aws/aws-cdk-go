package awswafv2


// Customizes the maximum size of the request body that your protected CloudFront distributions forward to AWS WAF for inspection.
//
// The default size is 16 KB (16,384 kilobytes).
//
// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
//
// This is used in the `AssociationConfig` of the web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestBodyAssociatedResourceTypeConfigProperty := &RequestBodyAssociatedResourceTypeConfigProperty{
//   	DefaultSizeInspectionLimit: jsii.String("defaultSizeInspectionLimit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestbodyassociatedresourcetypeconfig.html
//
type CfnWebACL_RequestBodyAssociatedResourceTypeConfigProperty struct {
	// Specifies the maximum size of the web request body component that an associated CloudFront distribution should send to AWS WAF for inspection.
	//
	// This applies to statements in the web ACL that inspect the body or JSON body.
	//
	// Default: `16 KB (16,384 kilobytes)`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestbodyassociatedresourcetypeconfig.html#cfn-wafv2-webacl-requestbodyassociatedresourcetypeconfig-defaultsizeinspectionlimit
	//
	DefaultSizeInspectionLimit *string `field:"required" json:"defaultSizeInspectionLimit" yaml:"defaultSizeInspectionLimit"`
}


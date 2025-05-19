package awscloudfront


// The AWS WAF web ACL customization specified for the distribution tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webAclCustomizationProperty := &WebAclCustomizationProperty{
//   	Action: jsii.String("action"),
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-webaclcustomization.html
//
type CfnDistributionTenant_WebAclCustomizationProperty struct {
	// The action for the AWS WAF web ACL customization.
	//
	// You can specify `override` to specify a separate AWS WAF web ACL for the distribution tenant. If you specify `disable` , the distribution tenant won't have AWS WAF web ACL protections and won't inherit from the multi-tenant distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-webaclcustomization.html#cfn-cloudfront-distributiontenant-webaclcustomization-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The Amazon Resource Name (ARN) of the AWS WAF web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-webaclcustomization.html#cfn-cloudfront-distributiontenant-webaclcustomization-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}


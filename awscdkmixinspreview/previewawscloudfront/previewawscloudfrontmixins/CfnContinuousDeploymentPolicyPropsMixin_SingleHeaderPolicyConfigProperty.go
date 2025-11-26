package previewawscloudfrontmixins


// Defines a single header policy for a CloudFront distribution.
//
// > This property is legacy. We recommend that you use [TrafficConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html) and specify the [SingleHeaderConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleheaderconfig) property instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singleHeaderPolicyConfigProperty := &SingleHeaderPolicyConfigProperty{
//   	Header: jsii.String("header"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig.html
//
type CfnContinuousDeploymentPolicyPropsMixin_SingleHeaderPolicyConfigProperty struct {
	// The name of the HTTP header that CloudFront uses to configure for the single header policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig-header
	//
	Header *string `field:"optional" json:"header" yaml:"header"`
	// Specifies the value to assign to the header for a single header policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


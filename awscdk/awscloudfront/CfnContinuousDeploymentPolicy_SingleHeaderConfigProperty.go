package awscloudfront


// Determines which HTTP requests are sent to the staging distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleHeaderConfigProperty := &SingleHeaderConfigProperty{
//   	Header: jsii.String("header"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html
//
type CfnContinuousDeploymentPolicy_SingleHeaderConfigProperty struct {
	// The request header name that you want CloudFront to send to your staging distribution.
	//
	// The header must contain the prefix `aws-cf-cd-` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-header
	//
	Header *string `field:"required" json:"header" yaml:"header"`
	// The request header value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


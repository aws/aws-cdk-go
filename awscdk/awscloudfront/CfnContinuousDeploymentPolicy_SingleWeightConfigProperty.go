package awscloudfront


// This configuration determines the percentage of HTTP requests that are sent to the staging distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleWeightConfigProperty := &SingleWeightConfigProperty{
//   	Weight: jsii.Number(123),
//
//   	// the properties below are optional
//   	SessionStickinessConfig: &SessionStickinessConfigProperty{
//   		IdleTtl: jsii.Number(123),
//   		MaximumTtl: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightconfig.html
//
type CfnContinuousDeploymentPolicy_SingleWeightConfigProperty struct {
	// The percentage of traffic to send to a staging distribution, expressed as a decimal number between 0 and .15.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightconfig-weight
	//
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Session stickiness provides the ability to define multiple requests from a single viewer as a single session.
	//
	// This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightconfig-sessionstickinessconfig
	//
	SessionStickinessConfig interface{} `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
}


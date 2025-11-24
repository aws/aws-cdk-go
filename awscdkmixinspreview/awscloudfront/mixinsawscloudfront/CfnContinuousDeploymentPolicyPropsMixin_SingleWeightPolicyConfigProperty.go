package mixinsawscloudfront


// Configure a policy that CloudFront uses to route requests to different origins or use different cache settings, based on the weight assigned to each option.
//
// > This property is legacy. We recommend that you use [TrafficConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html) and specify the [SingleWeightConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleweightconfig) property instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singleWeightPolicyConfigProperty := &SingleWeightPolicyConfigProperty{
//   	SessionStickinessConfig: &SessionStickinessConfigProperty{
//   		IdleTtl: jsii.Number(123),
//   		MaximumTtl: jsii.Number(123),
//   	},
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html
//
type CfnContinuousDeploymentPolicyPropsMixin_SingleWeightPolicyConfigProperty struct {
	// Enable session stickiness for the associated origin or cache settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig-sessionstickinessconfig
	//
	SessionStickinessConfig interface{} `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
	// The percentage of requests that CloudFront will use to send to an associated origin or cache settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


package mixinsawscloudfront


// The traffic configuration of your continuous deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trafficConfigProperty := &TrafficConfigProperty{
//   	SingleHeaderConfig: &SingleHeaderConfigProperty{
//   		Header: jsii.String("header"),
//   		Value: jsii.String("value"),
//   	},
//   	SingleWeightConfig: &SingleWeightConfigProperty{
//   		SessionStickinessConfig: &SessionStickinessConfigProperty{
//   			IdleTtl: jsii.Number(123),
//   			MaximumTtl: jsii.Number(123),
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html
//
type CfnContinuousDeploymentPolicyPropsMixin_TrafficConfigProperty struct {
	// Determines which HTTP requests are sent to the staging distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleheaderconfig
	//
	SingleHeaderConfig interface{} `field:"optional" json:"singleHeaderConfig" yaml:"singleHeaderConfig"`
	// Contains the percentage of traffic to send to the staging distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleweightconfig
	//
	SingleWeightConfig interface{} `field:"optional" json:"singleWeightConfig" yaml:"singleWeightConfig"`
	// The type of traffic configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


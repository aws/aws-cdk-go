package mixinsawscloudfront


// Properties for CfnContinuousDeploymentPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContinuousDeploymentPolicyMixinProps := &CfnContinuousDeploymentPolicyMixinProps{
//   	ContinuousDeploymentPolicyConfig: &ContinuousDeploymentPolicyConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		SingleHeaderPolicyConfig: &SingleHeaderPolicyConfigProperty{
//   			Header: jsii.String("header"),
//   			Value: jsii.String("value"),
//   		},
//   		SingleWeightPolicyConfig: &SingleWeightPolicyConfigProperty{
//   			SessionStickinessConfig: &SessionStickinessConfigProperty{
//   				IdleTtl: jsii.Number(123),
//   				MaximumTtl: jsii.Number(123),
//   			},
//   			Weight: jsii.Number(123),
//   		},
//   		StagingDistributionDnsNames: []*string{
//   			jsii.String("stagingDistributionDnsNames"),
//   		},
//   		TrafficConfig: &TrafficConfigProperty{
//   			SingleHeaderConfig: &SingleHeaderConfigProperty{
//   				Header: jsii.String("header"),
//   				Value: jsii.String("value"),
//   			},
//   			SingleWeightConfig: &SingleWeightConfigProperty{
//   				SessionStickinessConfig: &SessionStickinessConfigProperty{
//   					IdleTtl: jsii.Number(123),
//   					MaximumTtl: jsii.Number(123),
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-continuousdeploymentpolicy.html
//
type CfnContinuousDeploymentPolicyMixinProps struct {
	// Contains the configuration for a continuous deployment policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-continuousdeploymentpolicy.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig
	//
	ContinuousDeploymentPolicyConfig interface{} `field:"optional" json:"continuousDeploymentPolicyConfig" yaml:"continuousDeploymentPolicyConfig"`
}


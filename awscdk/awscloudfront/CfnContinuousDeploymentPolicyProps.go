package awscloudfront


// Properties for defining a `CfnContinuousDeploymentPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContinuousDeploymentPolicyProps := &CfnContinuousDeploymentPolicyProps{
//   	ContinuousDeploymentPolicyConfig: &ContinuousDeploymentPolicyConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		StagingDistributionDnsNames: []*string{
//   			jsii.String("stagingDistributionDnsNames"),
//   		},
//
//   		// the properties below are optional
//   		TrafficConfig: &TrafficConfigProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			SingleHeaderConfig: &SingleHeaderConfigProperty{
//   				Header: jsii.String("header"),
//   				Value: jsii.String("value"),
//   			},
//   			SingleWeightConfig: &SingleWeightConfigProperty{
//   				Weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				SessionStickinessConfig: &SessionStickinessConfigProperty{
//   					IdleTtl: jsii.Number(123),
//   					MaximumTtl: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnContinuousDeploymentPolicyProps struct {
	// Contains the configuration for a continuous deployment policy.
	ContinuousDeploymentPolicyConfig interface{} `field:"required" json:"continuousDeploymentPolicyConfig" yaml:"continuousDeploymentPolicyConfig"`
}


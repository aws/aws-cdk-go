package awscloudfront


// Properties for defining a `CfnContinuousDeploymentPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContinuousDeploymentPolicyProps := &cfnContinuousDeploymentPolicyProps{
//   	continuousDeploymentPolicyConfig: &continuousDeploymentPolicyConfigProperty{
//   		enabled: jsii.Boolean(false),
//   		stagingDistributionDnsNames: []*string{
//   			jsii.String("stagingDistributionDnsNames"),
//   		},
//
//   		// the properties below are optional
//   		trafficConfig: &trafficConfigProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			singleHeaderConfig: &singleHeaderConfigProperty{
//   				header: jsii.String("header"),
//   				value: jsii.String("value"),
//   			},
//   			singleWeightConfig: &singleWeightConfigProperty{
//   				weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				sessionStickinessConfig: &sessionStickinessConfigProperty{
//   					idleTtl: jsii.Number(123),
//   					maximumTtl: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnContinuousDeploymentPolicyProps struct {
	// `AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig`.
	ContinuousDeploymentPolicyConfig interface{} `field:"required" json:"continuousDeploymentPolicyConfig" yaml:"continuousDeploymentPolicyConfig"`
}


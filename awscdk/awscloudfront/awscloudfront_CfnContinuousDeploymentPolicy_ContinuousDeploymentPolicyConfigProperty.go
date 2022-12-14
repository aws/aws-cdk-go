package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousDeploymentPolicyConfigProperty := &continuousDeploymentPolicyConfigProperty{
//   	enabled: jsii.Boolean(false),
//   	stagingDistributionDnsNames: []*string{
//   		jsii.String("stagingDistributionDnsNames"),
//   	},
//
//   	// the properties below are optional
//   	trafficConfig: &trafficConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		singleHeaderConfig: &singleHeaderConfigProperty{
//   			header: jsii.String("header"),
//   			value: jsii.String("value"),
//   		},
//   		singleWeightConfig: &singleWeightConfigProperty{
//   			weight: jsii.Number(123),
//
//   			// the properties below are optional
//   			sessionStickinessConfig: &sessionStickinessConfigProperty{
//   				idleTtl: jsii.Number(123),
//   				maximumTtl: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfigProperty struct {
	// `CfnContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfigProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfigProperty.StagingDistributionDnsNames`.
	StagingDistributionDnsNames *[]*string `field:"required" json:"stagingDistributionDnsNames" yaml:"stagingDistributionDnsNames"`
	// `CfnContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfigProperty.TrafficConfig`.
	TrafficConfig interface{} `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
}


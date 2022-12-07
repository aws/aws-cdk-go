package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleWeightConfigProperty := &singleWeightConfigProperty{
//   	weight: jsii.Number(123),
//
//   	// the properties below are optional
//   	sessionStickinessConfig: &sessionStickinessConfigProperty{
//   		idleTtl: jsii.Number(123),
//   		maximumTtl: jsii.Number(123),
//   	},
//   }
//
type CfnContinuousDeploymentPolicy_SingleWeightConfigProperty struct {
	// `CfnContinuousDeploymentPolicy.SingleWeightConfigProperty.Weight`.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// `CfnContinuousDeploymentPolicy.SingleWeightConfigProperty.SessionStickinessConfig`.
	SessionStickinessConfig interface{} `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
}


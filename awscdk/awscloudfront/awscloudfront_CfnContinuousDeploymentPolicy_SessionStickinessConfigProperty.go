package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionStickinessConfigProperty := &sessionStickinessConfigProperty{
//   	idleTtl: jsii.Number(123),
//   	maximumTtl: jsii.Number(123),
//   }
//
type CfnContinuousDeploymentPolicy_SessionStickinessConfigProperty struct {
	// `CfnContinuousDeploymentPolicy.SessionStickinessConfigProperty.IdleTTL`.
	IdleTtl *float64 `field:"required" json:"idleTtl" yaml:"idleTtl"`
	// `CfnContinuousDeploymentPolicy.SessionStickinessConfigProperty.MaximumTTL`.
	MaximumTtl *float64 `field:"required" json:"maximumTtl" yaml:"maximumTtl"`
}


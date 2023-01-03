// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for enabling SageMaker Endpoint utilization tracking.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var model model
//
//
//   variantName := "my-variant"
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &endpointConfigProps{
//   	instanceProductionVariants: []instanceProductionVariantProps{
//   		&instanceProductionVariantProps{
//   			model: model,
//   			variantName: variantName,
//   		},
//   	},
//   })
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &endpointProps{
//   	endpointConfig: endpointConfig,
//   })
//   productionVariant := endpoint.findInstanceProductionVariant(variantName)
//   instanceCount := productionVariant.autoScaleInstanceCount(&enableScalingProps{
//   	maxCapacity: jsii.Number(3),
//   })
//   instanceCount.scaleOnInvocations(jsii.String("LimitRPS"), &invocationsScalingProps{
//   	maxRequestsPerSecond: jsii.Number(30),
//   })
//
// Experimental.
type InvocationsScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// Max RPS per instance used for calculating the target SageMaker variant invocation per instance.
	//
	// More documentation available here: https://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-scaling-loadtest.html
	// Experimental.
	MaxRequestsPerSecond *float64 `field:"required" json:"maxRequestsPerSecond" yaml:"maxRequestsPerSecond"`
	// Safty factor for calculating the target SageMaker variant invocation per instance.
	//
	// More documentation available here: https://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-scaling-loadtest.html
	// Experimental.
	SafetyFactor *float64 `field:"optional" json:"safetyFactor" yaml:"safetyFactor"`
}


package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for enabling SageMaker Endpoint utilization tracking.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var model model
//
//
//   variantName := "my-variant"
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &EndpointConfigProps{
//   	InstanceProductionVariants: []instanceProductionVariantProps{
//   		&instanceProductionVariantProps{
//   			Model: model,
//   			VariantName: variantName,
//   		},
//   	},
//   })
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &EndpointProps{
//   	EndpointConfig: EndpointConfig,
//   })
//   productionVariant := endpoint.FindInstanceProductionVariant(variantName)
//   instanceCount := productionVariant.AutoScaleInstanceCount(&EnableScalingProps{
//   	MaxCapacity: jsii.Number(3),
//   })
//   instanceCount.ScaleOnInvocations(jsii.String("LimitRPS"), &InvocationsScalingProps{
//   	MaxRequestsPerSecond: jsii.Number(30),
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
	// Default: false.
	//
	// Experimental.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Default: - Automatically generated name.
	//
	// Experimental.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Default: Duration.seconds(300) for the following scalable targets: ECS services,
	// Spot Fleet requests, EMR clusters, AppStream 2.0 fleets, Aurora DB clusters,
	// Amazon SageMaker endpoint variants, Custom resources. For all other scalable
	// targets, the default value is Duration.seconds(0): DynamoDB tables, DynamoDB
	// global secondary indexes, Amazon Comprehend document classification endpoints,
	// Lambda provisioned concurrency.
	//
	// Experimental.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Default: Duration.seconds(300) for the following scalable targets: ECS services,
	// Spot Fleet requests, EMR clusters, AppStream 2.0 fleets, Aurora DB clusters,
	// Amazon SageMaker endpoint variants, Custom resources. For all other scalable
	// targets, the default value is Duration.seconds(0): DynamoDB tables, DynamoDB
	// global secondary indexes, Amazon Comprehend document classification endpoints,
	// Lambda provisioned concurrency.
	//
	// Experimental.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// Max RPS per instance used for calculating the target SageMaker variant invocation per instance.
	//
	// More documentation available here: https://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-scaling-loadtest.html
	// Experimental.
	MaxRequestsPerSecond *float64 `field:"required" json:"maxRequestsPerSecond" yaml:"maxRequestsPerSecond"`
	// Safety factor for calculating the target SageMaker variant invocation per instance.
	//
	// More documentation available here: https://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-scaling-loadtest.html
	// Default: 0.5
	//
	// Experimental.
	SafetyFactor *float64 `field:"optional" json:"safetyFactor" yaml:"safetyFactor"`
}


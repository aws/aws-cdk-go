package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for enabling DynamoDB utilization tracking.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//
//   readCapacity := table.AutoScaleReadCapacity(&EnableScalingProps{
//   	MinCapacity: jsii.Number(10),
//   	MaxCapacity: jsii.Number(1000),
//   })
//   readCapacity.ScaleOnUtilization(&UtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(60),
//   })
//
type UtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Default: false.
	//
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Default: - Automatically generated name.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Default: Duration.seconds(300) for the following scalable targets: ECS services,
	// Spot Fleet requests, EMR clusters, AppStream 2.0 fleets, Aurora DB clusters,
	// Amazon SageMaker endpoint variants, Custom resources. For all other scalable
	// targets, the default value is Duration.seconds(0): DynamoDB tables, DynamoDB
	// global secondary indexes, Amazon Comprehend document classification endpoints,
	// Lambda provisioned concurrency.
	//
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Default: Duration.seconds(300) for the following scalable targets: ECS services,
	// Spot Fleet requests, EMR clusters, AppStream 2.0 fleets, Aurora DB clusters,
	// Amazon SageMaker endpoint variants, Custom resources. For all other scalable
	// targets, the default value is Duration.seconds(0): DynamoDB tables, DynamoDB
	// global secondary indexes, Amazon Comprehend document classification endpoints,
	// Lambda provisioned concurrency.
	//
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// Target utilization percentage for the attribute.
	TargetUtilizationPercent *float64 `field:"required" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}


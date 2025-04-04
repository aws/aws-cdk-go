package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for enabling Lambda utilization tracking.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   alias := fn.AddAlias(jsii.String("prod"))
//
//   // Create AutoScaling target
//   as := alias.AddAutoScaling(&AutoScalingOptions{
//   	MaxCapacity: jsii.Number(50),
//   })
//
//   // Configure Target Tracking
//   as.ScaleOnUtilization(&UtilizationScalingOptions{
//   	UtilizationTarget: jsii.Number(0.5),
//   })
//
//   // Configure Scheduled Scaling
//   as.ScaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &ScalingSchedule{
//   	Schedule: autoscaling.Schedule_Cron(&CronOptions{
//   		Hour: jsii.String("8"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(20),
//   })
//
type UtilizationScalingOptions struct {
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
	// Utilization target for the attribute.
	//
	// For example, .5 indicates that 50 percent of allocated provisioned concurrency is in use.
	UtilizationTarget *float64 `field:"required" json:"utilizationTarget" yaml:"utilizationTarget"`
}


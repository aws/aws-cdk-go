package awsapplicationautoscaling

import (
	"time"
)

// A scheduled scaling action.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.Service.AutoScaleTaskCount(&EnableScalingProps{
//   	MinCapacity: jsii.Number(5),
//   	MaxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.ScaleOnSchedule(jsii.String("DaytimeScaleDown"), &ScalingSchedule{
//   	Schedule: appscaling.Schedule_Cron(&CronOptions{
//   		Hour: jsii.String("8"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(1),
//   })
//
//   scalableTarget.ScaleOnSchedule(jsii.String("EveningRushScaleUp"), &ScalingSchedule{
//   	Schedule: appscaling.Schedule_*Cron(&CronOptions{
//   		Hour: jsii.String("20"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(10),
//   })
//
// Experimental.
type ScalingSchedule struct {
	// When to perform this action.
	// Experimental.
	Schedule Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// When this scheduled action expires.
	// Experimental.
	EndTime *time.Time `field:"optional" json:"endTime" yaml:"endTime"`
	// The new maximum capacity.
	//
	// During the scheduled time, the current capacity is above the maximum
	// capacity, Application Auto Scaling scales in to the maximum capacity.
	//
	// At least one of maxCapacity and minCapacity must be supplied.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The new minimum capacity.
	//
	// During the scheduled time, if the current capacity is below the minimum
	// capacity, Application Auto Scaling scales out to the minimum capacity.
	//
	// At least one of maxCapacity and minCapacity must be supplied.
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// When this scheduled action becomes active.
	// Experimental.
	StartTime *time.Time `field:"optional" json:"startTime" yaml:"startTime"`
}


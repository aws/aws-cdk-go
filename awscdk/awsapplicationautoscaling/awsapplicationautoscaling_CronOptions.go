package awsapplicationautoscaling


// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(5),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("DaytimeScaleDown"), &scalingSchedule{
//   	schedule: appscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(1),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("EveningRushScaleUp"), &scalingSchedule{
//   	schedule: appscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(10),
//   })
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
//
type CronOptions struct {
	// The day of the month to run this rule at.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	Month *string `field:"optional" json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	Year *string `field:"optional" json:"year" yaml:"year"`
}


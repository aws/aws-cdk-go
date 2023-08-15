package awsapplicationautoscaling


// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
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
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
//
type CronOptions struct {
	// The day of the month to run this rule at.
	// Default: - Every day of the month.
	//
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Default: - Every hour.
	//
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Default: - Every minute.
	//
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Default: - Every month.
	//
	Month *string `field:"optional" json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	Year *string `field:"optional" json:"year" yaml:"year"`
}


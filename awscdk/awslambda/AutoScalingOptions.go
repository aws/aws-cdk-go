package awslambda


// Properties for enabling Lambda autoscaling.
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
type AutoScalingOptions struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Default: 1.
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}


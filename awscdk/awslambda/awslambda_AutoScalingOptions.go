package awslambda


// Properties for enabling Lambda autoscaling.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   alias := fn.addAlias(jsii.String("prod"))
//
//   // Create AutoScaling target
//   as := alias.addAutoScaling(&autoScalingOptions{
//   	maxCapacity: jsii.Number(50),
//   })
//
//   // Configure Target Tracking
//   as.scaleOnUtilization(&utilizationScalingOptions{
//   	utilizationTarget: jsii.Number(0.5),
//   })
//
//   // Configure Scheduled Scaling
//   as.scaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &scalingSchedule{
//   	schedule: autoscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
type AutoScalingOptions struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}


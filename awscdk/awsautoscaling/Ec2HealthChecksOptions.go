package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// EC2 Heath checks options.
//
// Example:
//   var vpc vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	HealthChecks: autoscaling.HealthChecks_Ec2(&Ec2HealthChecksOptions{
//   		GracePeriod: awscdk.Duration_Seconds(jsii.Number(100)),
//   	}),
//   })
//
type Ec2HealthChecksOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html
	//
	// Default: Duration.seconds(0)
	//
	GracePeriod awscdk.Duration `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}


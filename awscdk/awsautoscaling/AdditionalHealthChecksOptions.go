package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Additional Heath checks options.
//
// Example:
//   var vpc Vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	HealthChecks: autoscaling.HealthChecks_WithAdditionalChecks(&AdditionalHealthChecksOptions{
//   		GracePeriod: awscdk.Duration_Seconds(jsii.Number(100)),
//   		AdditionalTypes: []AdditionalHealthCheckType{
//   			autoscaling.AdditionalHealthCheckType_EBS,
//   			autoscaling.AdditionalHealthCheckType_ELB,
//   			autoscaling.AdditionalHealthCheckType_VPC_LATTICE,
//   		},
//   	}),
//   })
//
type AdditionalHealthChecksOptions struct {
	// One or more health check types other than EC2.
	AdditionalTypes *[]AdditionalHealthCheckType `field:"required" json:"additionalTypes" yaml:"additionalTypes"`
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html
	//
	// Default: Duration.seconds(0)
	//
	GracePeriod awscdk.Duration `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}


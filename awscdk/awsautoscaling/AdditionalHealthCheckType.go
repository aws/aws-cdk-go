package awsautoscaling


// Additional Health Check Type.
//
// Example:
//   var vpc vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	HealthChecks: autoscaling.HealthChecks_WithAdditionalChecks(&AdditionalHealthChecksOptions{
//   		GracePeriod: awscdk.Duration_Seconds(jsii.Number(100)),
//   		AdditionalTypes: []additionalHealthCheckType{
//   			autoscaling.*additionalHealthCheckType_EBS,
//   			autoscaling.*additionalHealthCheckType_ELB,
//   			autoscaling.*additionalHealthCheckType_VPC_LATTICE,
//   		},
//   	}),
//   })
//
type AdditionalHealthCheckType string

const (
	// ELB Health Check.
	AdditionalHealthCheckType_ELB AdditionalHealthCheckType = "ELB"
	// EBS Health Check.
	AdditionalHealthCheckType_EBS AdditionalHealthCheckType = "EBS"
	// VPC LATTICE Health Check.
	AdditionalHealthCheckType_VPC_LATTICE AdditionalHealthCheckType = "VPC_LATTICE"
)


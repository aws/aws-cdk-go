package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// MixedInstancesPolicy allows you to configure a group that diversifies across On-Demand Instances and Spot Instances of multiple instance types.
//
// For more information, see Auto Scaling groups with
// multiple instance types and purchase options in the Amazon EC2 Auto Scaling User Guide:
//
// https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html
//
// Example:
//   var vpc vpc
//   var launchTemplate1 launchTemplate
//   var launchTemplate2 launchTemplate
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	MixedInstancesPolicy: &MixedInstancesPolicy{
//   		InstancesDistribution: &InstancesDistribution{
//   			OnDemandPercentageAboveBaseCapacity: jsii.Number(50),
//   		},
//   		LaunchTemplate: launchTemplate1,
//   		LaunchTemplateOverrides: []launchTemplateOverrides{
//   			&launchTemplateOverrides{
//   				InstanceType: ec2.NewInstanceType(jsii.String("t3.micro")),
//   			},
//   			&launchTemplateOverrides{
//   				InstanceType: ec2.NewInstanceType(jsii.String("t3a.micro")),
//   			},
//   			&launchTemplateOverrides{
//   				InstanceType: ec2.NewInstanceType(jsii.String("t4g.micro")),
//   				LaunchTemplate: launchTemplate2,
//   			},
//   		},
//   	},
//   })
//
type MixedInstancesPolicy struct {
	// Launch template to use.
	LaunchTemplate awsec2.ILaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// InstancesDistribution to use.
	// Default: - The value for each property in it uses a default value.
	//
	InstancesDistribution *InstancesDistribution `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
	// Launch template overrides.
	//
	// The maximum number of instance types that can be associated with an Auto Scaling group is 40.
	//
	// The maximum number of distinct launch templates you can define for an Auto Scaling group is 20.
	// Default: - Do not provide any overrides.
	//
	LaunchTemplateOverrides *[]*LaunchTemplateOverrides `field:"optional" json:"launchTemplateOverrides" yaml:"launchTemplateOverrides"`
}


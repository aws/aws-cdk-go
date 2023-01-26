package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// LaunchTemplateOverrides is a subproperty of LaunchTemplate that describes an override for a launch template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceType instanceType
//   var launchTemplate launchTemplate
//
//   launchTemplateOverrides := &launchTemplateOverrides{
//   	instanceType: instanceType,
//
//   	// the properties below are optional
//   	launchTemplate: launchTemplate,
//   	weightedCapacity: jsii.Number(123),
//   }
//
type LaunchTemplateOverrides struct {
	// The instance type, such as m3.xlarge. You must use an instance type that is supported in your requested Region and Availability Zones.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Provides the launch template to be used when launching the instance type.
	//
	// For example, some instance types might
	// require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template
	// that's defined for your mixed instances policy.
	LaunchTemplate awsec2.ILaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The number of capacity units provided by the specified instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is provisioned, the
	// capacity units count toward the desired capacity. Amazon EC2 Auto Scaling provisions instances until the desired
	// capacity is totally fulfilled, even if this results in an overage. Value must be in the range of 1 to 999.
	//
	// For example, If there are 2 units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only provision
	// an instance with a WeightedCapacity of 5 units, the instance is provisioned, and the desired capacity is exceeded
	// by 3 units.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html
	//
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


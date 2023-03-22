package awsec2


// Specifies the number of units to request for an EC2 Fleet.
//
// You can choose to set the target capacity in terms of instances or a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. If the request type is `maintain` , you can specify a target capacity of `0` and add capacity later.
//
// `TargetCapacitySpecificationRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetCapacitySpecificationRequestProperty := &TargetCapacitySpecificationRequestProperty{
//   	TotalTargetCapacity: jsii.Number(123),
//
//   	// the properties below are optional
//   	DefaultTargetCapacityType: jsii.String("defaultTargetCapacityType"),
//   	OnDemandTargetCapacity: jsii.Number(123),
//   	SpotTargetCapacity: jsii.Number(123),
//   	TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   }
//
type CfnEC2Fleet_TargetCapacitySpecificationRequestProperty struct {
	// The number of units to request, filled using `DefaultTargetCapacityType` .
	TotalTargetCapacity *float64 `field:"required" json:"totalTargetCapacity" yaml:"totalTargetCapacity"`
	// The default `TotalTargetCapacity` , which is either `Spot` or `On-Demand` .
	DefaultTargetCapacityType *string `field:"optional" json:"defaultTargetCapacityType" yaml:"defaultTargetCapacityType"`
	// The number of On-Demand units to request.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// The number of Spot units to request.
	SpotTargetCapacity *float64 `field:"optional" json:"spotTargetCapacity" yaml:"spotTargetCapacity"`
	// The unit for the target capacity. `TargetCapacityUnitType` can only be specified when `InstanceRequirements` is specified.
	//
	// Default: `units` (translates to number of instances).
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
}


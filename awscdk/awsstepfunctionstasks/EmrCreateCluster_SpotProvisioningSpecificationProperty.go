package awsstepfunctionstasks


// The launch specification for Spot instances in the instance fleet, which determines the defined duration and provisioning timeout behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotProvisioningSpecificationProperty := &SpotProvisioningSpecificationProperty{
//   	TimeoutAction: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.SpotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   	TimeoutDurationMinutes: jsii.Number(123),
//
//   	// the properties below are optional
//   	AllocationStrategy: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   	BlockDurationMinutes: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
type EmrCreateCluster_SpotProvisioningSpecificationProperty struct {
	// The action to take when TargetSpotCapacity has not been fulfilled when the TimeoutDurationMinutes has expired.
	TimeoutAction EmrCreateCluster_SpotTimeoutAction `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// The spot provisioning timeout period in minutes.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Specifies the strategy to use in launching Spot Instance fleets.
	AllocationStrategy EmrCreateCluster_SpotAllocationStrategy `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The defined duration for Spot instances (also known as Spot blocks) in minutes.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}


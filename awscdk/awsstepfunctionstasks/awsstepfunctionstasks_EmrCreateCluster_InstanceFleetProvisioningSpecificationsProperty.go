package awsstepfunctionstasks


// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &InstanceFleetProvisioningSpecificationsProperty{
//   	SpotSpecification: &SpotProvisioningSpecificationProperty{
//   		TimeoutAction: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.SpotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   		TimeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		AllocationStrategy: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   		BlockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetProvisioningSpecifications.html
//
// Experimental.
type EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
	// Experimental.
	SpotSpecification *EmrCreateCluster_SpotProvisioningSpecificationProperty `field:"required" json:"spotSpecification" yaml:"spotSpecification"`
}


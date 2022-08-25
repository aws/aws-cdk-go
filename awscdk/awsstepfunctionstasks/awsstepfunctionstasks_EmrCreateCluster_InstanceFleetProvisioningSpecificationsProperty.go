package awsstepfunctionstasks


// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &instanceFleetProvisioningSpecificationsProperty{
//   	spotSpecification: &spotProvisioningSpecificationProperty{
//   		timeoutAction: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.spotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   		timeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotAllocationStrategy_CAPACITY_OPTIMIZED,
//   		blockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetProvisioningSpecifications.html
//
type EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
	SpotSpecification *EmrCreateCluster_SpotProvisioningSpecificationProperty `field:"required" json:"spotSpecification" yaml:"spotSpecification"`
}


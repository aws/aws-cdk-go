package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The launch specification for Spot instances in the instance fleet, which determines the defined duration and provisioning timeout behavior.
//
// Example:
//   tasks.NewEmrCreateCluster(this, jsii.String("OnDemandSpecification"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []InstanceFleetConfigProperty{
//   			&InstanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   					OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   						AllocationStrategy: tasks.EmrCreateCluster.OnDemandAllocationStrategy_LOWEST_PRICE,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("OnDemandCluster"),
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   })
//
//   tasks.NewEmrCreateCluster(this, jsii.String("SpotSpecification"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []InstanceFleetConfigProperty{
//   			&InstanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   					SpotSpecification: &SpotProvisioningSpecificationProperty{
//   						AllocationStrategy: tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   						TimeoutAction: tasks.EmrCreateCluster.SpotTimeoutAction_TERMINATE_CLUSTER,
//   						Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("SpotCluster"),
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   })
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
type EmrCreateCluster_SpotProvisioningSpecificationProperty struct {
	// The action to take when TargetSpotCapacity has not been fulfilled when the TimeoutDurationMinutes has expired.
	TimeoutAction EmrCreateCluster_SpotTimeoutAction `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// Specifies the strategy to use in launching Spot Instance fleets.
	// Default: - No allocation strategy, i.e. spot instance type will be chosen based on current price only
	//
	AllocationStrategy EmrCreateCluster_SpotAllocationStrategy `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The defined duration for Spot instances (also known as Spot blocks) in minutes.
	// Default: - No blockDurationMinutes.
	//
	// Deprecated: - Spot Instances with a defined duration (also known as Spot blocks) are no longer available to new customers from July 1, 2021.
	// For customers who have previously used the feature, we will continue to support Spot Instances with a defined duration until December 31, 2022.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// The spot provisioning timeout period in minutes.
	//
	// The value must be between 5 and 1440 minutes.
	//
	// You must specify one of `timeout` and `timeoutDurationMinutes`.
	// Default: - The value in `timeoutDurationMinutes` is used.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The spot provisioning timeout period in minutes.
	//
	// The value must be between 5 and 1440 minutes.
	//
	// You must specify one of `timeout` and `timeoutDurationMinutes`.
	// Default: - The value in `timeout` is used.
	//
	// Deprecated: - Use `timeout`.
	TimeoutDurationMinutes *float64 `field:"optional" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
}


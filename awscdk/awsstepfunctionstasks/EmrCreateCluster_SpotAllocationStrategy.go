package awsstepfunctionstasks


// Spot Allocation Strategies.
//
// Specifies the strategy to use in launching Spot Instance fleets. For example, "capacity-optimized" launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
//
// Example:
//   tasks.NewEmrCreateCluster(this, jsii.String("OnDemandSpecification"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []instanceFleetConfigProperty{
//   			&instanceFleetConfigProperty{
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
//   		InstanceFleets: []*instanceFleetConfigProperty{
//   			&instanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   					SpotSpecification: &SpotProvisioningSpecificationProperty{
//   						AllocationStrategy: tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   						TimeoutAction: tasks.EmrCreateCluster.SpotTimeoutAction_TERMINATE_CLUSTER,
//   						TimeoutDurationMinutes: jsii.Number(60),
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
type EmrCreateCluster_SpotAllocationStrategy string

const (
	// Capacity-optimized, which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
)


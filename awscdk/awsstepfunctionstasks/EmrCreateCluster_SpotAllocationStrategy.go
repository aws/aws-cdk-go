package awsstepfunctionstasks


// Spot Allocation Strategies.
//
// Specifies the strategy to use in launching Spot Instance fleets. For example, "capacity-optimized" launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
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
type EmrCreateCluster_SpotAllocationStrategy string

const (
	// Capacity-optimized, which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
	// Price-capacity-optimized, which launches instances from Spot Instance pools with the highest capacity availability for the number of instances that are launching.
	//
	// Recommended.
	EmrCreateCluster_SpotAllocationStrategy_PRICE_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "PRICE_CAPACITY_OPTIMIZED"
	// Lowest-price, which launches instances from the lowest priced pool that has available capacity.
	EmrCreateCluster_SpotAllocationStrategy_LOWEST_PRICE EmrCreateCluster_SpotAllocationStrategy = "LOWEST_PRICE"
	// Diversified, which launches instances across all Spot capacity pools.
	EmrCreateCluster_SpotAllocationStrategy_DIVERSIFIED EmrCreateCluster_SpotAllocationStrategy = "DIVERSIFIED"
)


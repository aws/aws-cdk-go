package awsstepfunctionstasks


// Spot Timeout Actions.
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
type EmrCreateCluster_SpotTimeoutAction string

const (
	// SWITCH_TO_ON_DEMAND.
	EmrCreateCluster_SpotTimeoutAction_SWITCH_TO_ON_DEMAND EmrCreateCluster_SpotTimeoutAction = "SWITCH_TO_ON_DEMAND"
	// TERMINATE_CLUSTER.
	EmrCreateCluster_SpotTimeoutAction_TERMINATE_CLUSTER EmrCreateCluster_SpotTimeoutAction = "TERMINATE_CLUSTER"
)


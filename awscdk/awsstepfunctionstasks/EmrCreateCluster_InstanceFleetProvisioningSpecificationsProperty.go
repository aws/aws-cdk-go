package awsstepfunctionstasks


// The launch specification for On-Demand and Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior, and allocation strategy.
//
// The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions.
// On-Demand and Spot instance allocation strategies are available in Amazon EMR releases 5.12.1 and later.
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
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetProvisioningSpecifications.html
//
type EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
	//
	// The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions.
	// On-Demand Instances allocation strategy is available in Amazon EMR releases 5.12.1 and later.
	// Default: - no on-demand specification.
	//
	OnDemandSpecification *EmrCreateCluster_OnDemandProvisioningSpecificationProperty `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
	// Default: - no spot specification.
	//
	SpotSpecification *EmrCreateCluster_SpotProvisioningSpecificationProperty `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}


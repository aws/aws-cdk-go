package awsstepfunctionstasks


// The EC2 unit limits for a managed scaling policy.
//
// Example:
//   tasks.NewEmrCreateCluster(this, jsii.String("CreateCluster"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []instanceFleetConfigProperty{
//   			&instanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_CORE,
//   				InstanceTypeConfigs: []instanceTypeConfigProperty{
//   					&instanceTypeConfigProperty{
//   						InstanceType: jsii.String("m5.xlarge"),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(1),
//   			},
//   			&instanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				InstanceTypeConfigs: []*instanceTypeConfigProperty{
//   					&instanceTypeConfigProperty{
//   						InstanceType: jsii.String("m5.xlarge"),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(1),
//   			},
//   		},
//   	},
//   	Name: jsii.String("ClusterName"),
//   	ReleaseLabel: jsii.String("emr-7.9.0"),
//   	ManagedScalingPolicy: &ManagedScalingPolicyProperty{
//   		ComputeLimits: &ManagedScalingComputeLimitsProperty{
//   			UnitType: tasks.EmrCreateCluster.ComputeLimitsUnitType_INSTANCE_FLEET_UNITS,
//   			MaximumCapacityUnits: jsii.Number(4),
//   			MinimumCapacityUnits: jsii.Number(1),
//   			MaximumOnDemandCapacityUnits: jsii.Number(4),
//   			MaximumCoreCapacityUnits: jsii.Number(2),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ComputeLimits.html
//
type EmrCreateCluster_ManagedScalingComputeLimitsProperty struct {
	// The upper boundary of Amazon EC2 units.
	MaximumCapacityUnits *float64 `field:"required" json:"maximumCapacityUnits" yaml:"maximumCapacityUnits"`
	// The lower boundary of Amazon EC2 units.
	MinimumCapacityUnits *float64 `field:"required" json:"minimumCapacityUnits" yaml:"minimumCapacityUnits"`
	// The unit type used for specifying a managed scaling policy.
	UnitType EmrCreateCluster_ComputeLimitsUnitType `field:"required" json:"unitType" yaml:"unitType"`
	// The upper boundary of Amazon EC2 units for core node type in a cluster.
	// Default: - None.
	//
	MaximumCoreCapacityUnits *float64 `field:"optional" json:"maximumCoreCapacityUnits" yaml:"maximumCoreCapacityUnits"`
	// The upper boundary of On-Demand Amazon EC2 units.
	// Default: - None.
	//
	MaximumOnDemandCapacityUnits *float64 `field:"optional" json:"maximumOnDemandCapacityUnits" yaml:"maximumOnDemandCapacityUnits"`
}


package awsstepfunctionstasks


// Instance Role Types.
//
// Example:
//   tasks.NewEmrCreateCluster(this, jsii.String("CreateCluster"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []InstanceFleetConfigProperty{
//   			&InstanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_CORE,
//   				InstanceTypeConfigs: []InstanceTypeConfigProperty{
//   					&InstanceTypeConfigProperty{
//   						InstanceType: jsii.String("m5.xlarge"),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(1),
//   			},
//   			&InstanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				InstanceTypeConfigs: []InstanceTypeConfigProperty{
//   					&InstanceTypeConfigProperty{
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
type EmrCreateCluster_InstanceRoleType string

const (
	// Master Node.
	EmrCreateCluster_InstanceRoleType_MASTER EmrCreateCluster_InstanceRoleType = "MASTER"
	// Core Node.
	EmrCreateCluster_InstanceRoleType_CORE EmrCreateCluster_InstanceRoleType = "CORE"
	// Task Node.
	EmrCreateCluster_InstanceRoleType_TASK EmrCreateCluster_InstanceRoleType = "TASK"
)


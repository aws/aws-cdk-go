package awsstepfunctionstasks


// The unit type for managed scaling policy compute limits.
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
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ComputeLimits.html
//
type EmrCreateCluster_ComputeLimitsUnitType string

const (
	// InstanceFleetUnits.
	EmrCreateCluster_ComputeLimitsUnitType_INSTANCE_FLEET_UNITS EmrCreateCluster_ComputeLimitsUnitType = "INSTANCE_FLEET_UNITS"
	// Instances.
	EmrCreateCluster_ComputeLimitsUnitType_INSTANCES EmrCreateCluster_ComputeLimitsUnitType = "INSTANCES"
	// VCPU.
	EmrCreateCluster_ComputeLimitsUnitType_VCPU EmrCreateCluster_ComputeLimitsUnitType = "VCPU"
)


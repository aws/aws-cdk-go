package awsstepfunctionstasks


// The managed scaling policy for an Amazon EMR cluster.
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
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ManagedScalingPolicy.html
//
type EmrCreateCluster_ManagedScalingPolicyProperty struct {
	// The Amazon EC2 unit limits for a managed scaling policy.
	// Default: - None.
	//
	ComputeLimits *EmrCreateCluster_ManagedScalingComputeLimitsProperty `field:"optional" json:"computeLimits" yaml:"computeLimits"`
}


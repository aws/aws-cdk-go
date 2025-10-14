package awsstepfunctionstasks


// A specification of the number and type of Amazon EC2 instances.
//
// See the RunJobFlow API for complete documentation on input parameters.
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
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_JobFlowInstancesConfig.html
//
type EmrCreateCluster_InstancesConfigProperty struct {
	// A list of additional Amazon EC2 security group IDs for the master node.
	// Default: - None.
	//
	AdditionalMasterSecurityGroups *[]*string `field:"optional" json:"additionalMasterSecurityGroups" yaml:"additionalMasterSecurityGroups"`
	// A list of additional Amazon EC2 security group IDs for the core and task nodes.
	// Default: - None.
	//
	AdditionalSlaveSecurityGroups *[]*string `field:"optional" json:"additionalSlaveSecurityGroups" yaml:"additionalSlaveSecurityGroups"`
	// The name of the EC2 key pair that can be used to ssh to the master node as the user called "hadoop.".
	// Default: - None.
	//
	Ec2KeyName *string `field:"optional" json:"ec2KeyName" yaml:"ec2KeyName"`
	// Applies to clusters that use the uniform instance group configuration.
	//
	// To launch the cluster in Amazon Virtual Private Cloud (Amazon VPC),
	// set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch.
	// Default: EMR selected default.
	//
	Ec2SubnetId *string `field:"optional" json:"ec2SubnetId" yaml:"ec2SubnetId"`
	// Applies to clusters that use the instance fleet configuration.
	//
	// When multiple EC2 subnet IDs are specified, Amazon EMR evaluates them and
	// launches instances in the optimal subnet.
	// Default: EMR selected default.
	//
	Ec2SubnetIds *[]*string `field:"optional" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The identifier of the Amazon EC2 security group for the master node.
	// Default: - None.
	//
	EmrManagedMasterSecurityGroup *string `field:"optional" json:"emrManagedMasterSecurityGroup" yaml:"emrManagedMasterSecurityGroup"`
	// The identifier of the Amazon EC2 security group for the core and task nodes.
	// Default: - None.
	//
	EmrManagedSlaveSecurityGroup *string `field:"optional" json:"emrManagedSlaveSecurityGroup" yaml:"emrManagedSlaveSecurityGroup"`
	// Applies only to Amazon EMR release versions earlier than 4.0. The Hadoop version for the cluster.
	// Default: - 0.18 if the AmiVersion parameter is not set. If AmiVersion is set, the version of Hadoop for that AMI version is used.
	//
	HadoopVersion *string `field:"optional" json:"hadoopVersion" yaml:"hadoopVersion"`
	// The number of EC2 instances in the cluster.
	// Default: 0.
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Describes the EC2 instances and instance configurations for clusters that use the instance fleet configuration.
	//
	// The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	// Default: - None.
	//
	InstanceFleets *[]*EmrCreateCluster_InstanceFleetConfigProperty `field:"optional" json:"instanceFleets" yaml:"instanceFleets"`
	// Configuration for the instance groups in a cluster.
	// Default: - None.
	//
	InstanceGroups *[]*EmrCreateCluster_InstanceGroupConfigProperty `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
	// The EC2 instance type of the master node.
	// Default: - None.
	//
	MasterInstanceType *string `field:"optional" json:"masterInstanceType" yaml:"masterInstanceType"`
	// The Availability Zone in which the cluster runs.
	// Default: - EMR selected default.
	//
	Placement *EmrCreateCluster_PlacementTypeProperty `field:"optional" json:"placement" yaml:"placement"`
	// The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.
	// Default: - None.
	//
	ServiceAccessSecurityGroup *string `field:"optional" json:"serviceAccessSecurityGroup" yaml:"serviceAccessSecurityGroup"`
	// The EC2 instance type of the core and task nodes.
	// Default: - None.
	//
	SlaveInstanceType *string `field:"optional" json:"slaveInstanceType" yaml:"slaveInstanceType"`
	// Specifies whether to lock the cluster to prevent the Amazon EC2 instances from being terminated by API call, user intervention, or in the event of a job-flow error.
	// Default: false.
	//
	TerminationProtected *bool `field:"optional" json:"terminationProtected" yaml:"terminationProtected"`
}


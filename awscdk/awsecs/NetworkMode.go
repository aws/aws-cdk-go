package awsecs


// The networking mode to use for the containers in the task.
//
// Example:
//   var vpc Vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   	Description: jsii.String("Security group for managed instances"),
//   })
//
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		VCpuCountMin: jsii.Number(1),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(2)),
//   	},
//   })
//
//   // Optionally configure security group rules using IConnectable interface
//   miCapacityProvider.Connections.AllowFrom(ec2.Peer_Ipv4(vpc.VpcCidrBlock), ec2.Port_Tcp(jsii.Number(80)))
//
//   // Add the capacity provider to the cluster
//   cluster.AddManagedInstancesCapacityProvider(miCapacityProvider)
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &TaskDefinitionProps{
//   	MemoryMiB: jsii.String("512"),
//   	Cpu: jsii.String("256"),
//   	NetworkMode: ecs.NetworkMode_AWS_VPC,
//   	Compatibility: ecs.Compatibility_MANAGED_INSTANCES,
//   })
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	CapacityProviderStrategies: []CapacityProviderStrategy{
//   		&CapacityProviderStrategy{
//   			CapacityProvider: miCapacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type NetworkMode string

const (
	// The task's containers do not have external connectivity and port mappings can't be specified in the container definition.
	NetworkMode_NONE NetworkMode = "NONE"
	// The task utilizes Docker's built-in virtual network which runs inside each container instance.
	NetworkMode_BRIDGE NetworkMode = "BRIDGE"
	// The task is allocated an elastic network interface.
	NetworkMode_AWS_VPC NetworkMode = "AWS_VPC"
	// The task bypasses Docker's built-in virtual network and maps container ports directly to the EC2 instance's network interface directly.
	//
	// In this mode, you can't run multiple instantiations of the same task on a
	// single container instance when port mappings are used.
	NetworkMode_HOST NetworkMode = "HOST"
	// The task utilizes Docker's built-in virtual network which runs inside each Windows container instance.
	NetworkMode_NAT NetworkMode = "NAT"
)


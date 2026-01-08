package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Options to run an ECS task on EC2 in StepFunctions and ECS.
//
// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroup"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
//   	Compatibility: ecs.Compatibility_EC2,
//   })
//
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
//   	MemoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchTarget: tasks.NewEcsEc2LaunchTarget(&EcsEc2LaunchTargetOptions{
//   		PlacementStrategies: []PlacementStrategy{
//   			ecs.PlacementStrategy_SpreadAcrossInstances(),
//   			ecs.PlacementStrategy_PackedByCpu(),
//   			ecs.PlacementStrategy_Randomly(),
//   		},
//   		PlacementConstraints: []PlacementConstraint{
//   			ecs.PlacementConstraint_MemberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   	PropagatedTagSource: ecs.PropagatedTagSource_TASK_DEFINITION,
//   })
//
type EcsEc2LaunchTargetOptions struct {
	// The capacity provider options to use for the task.
	//
	// This property allows you to set the capacity provider strategy for the task.
	//
	// If you want to set the capacity provider strategy for the task, specify
	// `CapacityProviderOptions.custom()`.
	//
	// If you want to use the cluster's default capacity provider strategy, specify
	// `CapacityProviderOptions.default()`.
	// Default: - 'EC2' LaunchType running tasks on Amazon EC2 instances registered to
	// your cluster is used without the capacity provider strategy.
	//
	CapacityProviderOptions CapacityProviderOptions `field:"optional" json:"capacityProviderOptions" yaml:"capacityProviderOptions"`
	// Placement constraints.
	// Default: - None.
	//
	PlacementConstraints *[]awsecs.PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// Placement strategies.
	// Default: - None.
	//
	PlacementStrategies *[]awsecs.PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
}


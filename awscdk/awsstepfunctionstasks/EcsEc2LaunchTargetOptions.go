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
//   		PlacementStrategies: []placementStrategy{
//   			ecs.*placementStrategy_SpreadAcrossInstances(),
//   			ecs.*placementStrategy_PackedByCpu(),
//   			ecs.*placementStrategy_Randomly(),
//   		},
//   		PlacementConstraints: []placementConstraint{
//   			ecs.*placementConstraint_MemberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   	PropagatedTagSource: ecs.PropagatedTagSource_TASK_DEFINITION,
//   })
//
type EcsEc2LaunchTargetOptions struct {
	// Placement constraints.
	PlacementConstraints *[]awsecs.PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// Placement strategies.
	PlacementStrategies *[]awsecs.PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
}


package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The placement strategies to use for tasks in the service. For more information, see [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
//
// Tasks will preferentially be placed on instances that match these rules.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroup"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	compatibility: ecs.compatibility_EC2,
//   })
//
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	launchTarget: tasks.NewEcsEc2LaunchTarget(&ecsEc2LaunchTargetOptions{
//   		placementStrategies: []placementStrategy{
//   			ecs.*placementStrategy.spreadAcrossInstances(),
//   			ecs.*placementStrategy.packedByCpu(),
//   			ecs.*placementStrategy.randomly(),
//   		},
//   		placementConstraints: []placementConstraint{
//   			ecs.*placementConstraint.memberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   })
//
type PlacementStrategy interface {
	// Return the placement JSON.
	ToJson() *[]*CfnService_PlacementStrategyProperty
}

// The jsii proxy struct for PlacementStrategy
type jsiiProxy_PlacementStrategy struct {
	_ byte // padding
}

// Places tasks on the container instances with the least available capacity of the specified resource.
func PlacementStrategy_PackedBy(resource BinPackResource) PlacementStrategy {
	_init_.Initialize()

	if err := validatePlacementStrategy_PackedByParameters(resource); err != nil {
		panic(err)
	}
	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"packedBy",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Places tasks on container instances with the least available amount of CPU capacity.
//
// This minimizes the number of instances in use.
func PlacementStrategy_PackedByCpu() PlacementStrategy {
	_init_.Initialize()

	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"packedByCpu",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Places tasks on container instances with the least available amount of memory capacity.
//
// This minimizes the number of instances in use.
func PlacementStrategy_PackedByMemory() PlacementStrategy {
	_init_.Initialize()

	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"packedByMemory",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Places tasks randomly.
func PlacementStrategy_Randomly() PlacementStrategy {
	_init_.Initialize()

	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"randomly",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Places tasks evenly based on the specified value.
//
// You can use one of the built-in attributes found on `BuiltInAttributes`
// or supply your own custom instance attributes. If more than one attribute
// is supplied, spreading is done in order.
func PlacementStrategy_SpreadAcross(fields ...*string) PlacementStrategy {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"spreadAcross",
		args,
		&returns,
	)

	return returns
}

// Places tasks evenly across all container instances in the cluster.
func PlacementStrategy_SpreadAcrossInstances() PlacementStrategy {
	_init_.Initialize()

	var returns PlacementStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementStrategy",
		"spreadAcrossInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PlacementStrategy) ToJson() *[]*CfnService_PlacementStrategyProperty {
	var returns *[]*CfnService_PlacementStrategyProperty

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}


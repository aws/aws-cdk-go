package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The placement constraints to use for tasks in the service. For more information, see [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
//
// Tasks will only be placed on instances that match these rules.
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
type PlacementConstraint interface {
	// Return the placement JSON.
	ToJson() *[]*CfnService_PlacementConstraintProperty
}

// The jsii proxy struct for PlacementConstraint
type jsiiProxy_PlacementConstraint struct {
	_ byte // padding
}

// Use distinctInstance to ensure that each task in a particular group is running on a different container instance.
func PlacementConstraint_DistinctInstances() PlacementConstraint {
	_init_.Initialize()

	var returns PlacementConstraint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementConstraint",
		"distinctInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use memberOf to restrict the selection to a group of valid candidates specified by a query expression.
//
// Multiple expressions can be specified. For more information, see
// [Cluster Query Language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
//
// You can specify multiple expressions in one call. The tasks will only be placed on instances matching all expressions.
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html
//
func PlacementConstraint_MemberOf(expressions ...*string) PlacementConstraint {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range expressions {
		args = append(args, a)
	}

	var returns PlacementConstraint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.PlacementConstraint",
		"memberOf",
		args,
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PlacementConstraint) ToJson() *[]*CfnService_PlacementConstraintProperty {
	var returns *[]*CfnService_PlacementConstraintProperty

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}


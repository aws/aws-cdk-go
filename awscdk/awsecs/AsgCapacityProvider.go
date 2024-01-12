package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Auto Scaling Group Capacity Provider.
//
// This allows an ECS cluster to target
// a specific EC2 Auto Scaling Group for the placement of tasks. Optionally (and
// recommended), ECS can manage the number of instances in the ASG to fit the
// tasks, and can ensure that instances are not prematurely terminated while
// there are still tasks running on them.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Either add default capacity
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   // Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux(),
//   	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
//   	// machineImage: EcsOptimizedImage.amazonLinux2(),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
//   	AutoScalingGroup: AutoScalingGroup,
//   })
//   cluster.AddAsgCapacityProvider(capacityProvider)
//
type AsgCapacityProvider interface {
	constructs.Construct
	// Auto Scaling Group.
	AutoScalingGroup() awsautoscaling.AutoScalingGroup
	// Specifies whether the containers can access the container instance role.
	// Default: false.
	//
	CanContainersAccessInstanceRole() *bool
	// Capacity provider name.
	// Default: Chosen by CloudFormation.
	//
	CapacityProviderName() *string
	// Whether managed draining is enabled.
	EnableManagedDraining() *bool
	// Whether managed termination protection is enabled.
	EnableManagedTerminationProtection() *bool
	// Auto Scaling Group machineImageType.
	MachineImageType() MachineImageType
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AsgCapacityProvider
type jsiiProxy_AsgCapacityProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AsgCapacityProvider) AutoScalingGroup() awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup
	_jsii_.Get(
		j,
		"autoScalingGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) CanContainersAccessInstanceRole() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canContainersAccessInstanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) CapacityProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) EnableManagedDraining() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableManagedDraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) EnableManagedTerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableManagedTerminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) MachineImageType() MachineImageType {
	var returns MachineImageType
	_jsii_.Get(
		j,
		"machineImageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsgCapacityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAsgCapacityProvider(scope constructs.Construct, id *string, props *AsgCapacityProviderProps) AsgCapacityProvider {
	_init_.Initialize()

	if err := validateNewAsgCapacityProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsgCapacityProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AsgCapacityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAsgCapacityProvider_Override(a AsgCapacityProvider, scope constructs.Construct, id *string, props *AsgCapacityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AsgCapacityProvider",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AsgCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsgCapacityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AsgCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsgCapacityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


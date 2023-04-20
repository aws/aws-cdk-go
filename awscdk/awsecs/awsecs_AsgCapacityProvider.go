package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
//   	MinCapacity: jsii.Number(0),
//   	MaxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
//   	AutoScalingGroup: AutoScalingGroup,
//   })
//   cluster.AddAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: capacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type AsgCapacityProvider interface {
	awscdk.Construct
	// Auto Scaling Group.
	// Experimental.
	AutoScalingGroup() awsautoscaling.AutoScalingGroup
	// Specifies whether the containers can access the container instance role.
	// Experimental.
	CanContainersAccessInstanceRole() *bool
	// Capacity provider name.
	// Experimental.
	CapacityProviderName() *string
	// Whether managed termination protection is enabled.
	// Experimental.
	EnableManagedTerminationProtection() *bool
	// Auto Scaling Group machineImageType.
	// Experimental.
	MachineImageType() MachineImageType
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for AsgCapacityProvider
type jsiiProxy_AsgCapacityProvider struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_AsgCapacityProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAsgCapacityProvider(scope constructs.Construct, id *string, props *AsgCapacityProviderProps) AsgCapacityProvider {
	_init_.Initialize()

	if err := validateNewAsgCapacityProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsgCapacityProvider{}

	_jsii_.Create(
		"monocdk.aws_ecs.AsgCapacityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAsgCapacityProvider_Override(a AsgCapacityProvider, scope constructs.Construct, id *string, props *AsgCapacityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.AsgCapacityProvider",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func AsgCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsgCapacityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AsgCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsgCapacityProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsgCapacityProvider) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AsgCapacityProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsgCapacityProvider) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsgCapacityProvider) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_AsgCapacityProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


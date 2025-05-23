package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Fleet represents a managed set of EC2 instances.
//
// The Fleet models a number of AutoScalingGroups, a launch configuration, a
// security group and an instance role.
//
// It allows adding arbitrary commands to the startup scripts of the instances
// in the fleet.
//
// The ASG spans the availability zones specified by vpcSubnets, falling back to
// the Vpc default strategy if not specified.
//
// Example:
//   var vpc vpc
//
//
//   mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   	SecurityGroup: mySecurityGroup,
//   })
//
type AutoScalingGroup interface {
	awscdk.Resource
	IAutoScalingGroup
	awsec2.IConnectable
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	AlbTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	SetAlbTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup)
	// Arn of the AutoScalingGroup.
	AutoScalingGroupArn() *string
	// Name of the AutoScalingGroup.
	AutoScalingGroupName() *string
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The principal to grant permissions to.
	GrantPrincipal() awsiam.IPrincipal
	HasCalledScaleOnRequestCount() *bool
	SetHasCalledScaleOnRequestCount(val *bool)
	// The maximum amount of time that an instance can be in service.
	MaxInstanceLifetime() awscdk.Duration
	NewInstancesProtectedFromScaleIn() *bool
	SetNewInstancesProtectedFromScaleIn(val *bool)
	// The tree node.
	Node() constructs.Node
	// The type of OS instances of this fleet are running.
	OsType() awsec2.OperatingSystemType
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The IAM Role in the instance profile.
	Role() awsiam.IRole
	// The maximum spot price configured for the autoscaling group.
	//
	// `undefined`
	// indicates that this group uses on-demand capacity.
	SpotPrice() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The Base64-encoded user data to make available to the launched EC2 instances.
	UserData() awsec2.UserData
	// Send a message to either an SQS queue or SNS topic when instances launch or terminate.
	AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook
	// Add the security group to all instances via the launch template security groups array.
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	// Adds a statement to the IAM role assumed by instances of this fleet.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Add command to the startup script of fleet instances.
	//
	// The command must be in the scripting language supported by the fleet's OS (i.e. Linux/Windows).
	// Does nothing for imported ASGs.
	AddUserData(commands ...*string)
	// Add a pool of pre-initialized EC2 instances that sits alongside an Auto Scaling group.
	AddWarmPool(options *WarmPoolOptions) WarmPool
	// Use a CloudFormation Init configuration at instance startup.
	//
	// This does the following:
	//
	// - Attaches the CloudFormation Init metadata to the AutoScalingGroup resource.
	// - Add commands to the UserData to run `cfn-init` and `cfn-signal`.
	// - Update the instance's CreationPolicy to wait for `cfn-init` to finish
	//   before reporting success.
	ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns `true` if newly-launched instances are protected from scale-in.
	AreNewInstancesProtectedFromScaleIn() *bool
	// Attach to ELBv2 Application Target Group.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Attach to a classic load balancer.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// Attach to ELBv2 Application Target Group.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Ensures newly-launched instances are protected from scale-in.
	ProtectNewInstancesFromScaleIn()
	// Scale out or in to achieve a target CPU utilization.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target network ingress rate.
	ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in, in response to a metric.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in to achieve a target network egress rate.
	ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target request handling rate.
	//
	// The AutoScalingGroup must have been attached to an Application Load Balancer
	// in order to be able to call this.
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy
	// Scale out or in based on time.
	ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction
	// Scale out or in in order to keep a metric around a target value.
	ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AutoScalingGroup
type jsiiProxy_AutoScalingGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IAutoScalingGroup
	internal.Type__awsec2IConnectable
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

func (j *jsiiProxy_AutoScalingGroup) AlbTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"albTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) HasCalledScaleOnRequestCount() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasCalledScaleOnRequestCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) MaxInstanceLifetime() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) NewInstancesProtectedFromScaleIn() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) OsType() awsec2.OperatingSystemType {
	var returns awsec2.OperatingSystemType
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) UserData() awsec2.UserData {
	var returns awsec2.UserData
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


func NewAutoScalingGroup(scope constructs.Construct, id *string, props *AutoScalingGroupProps) AutoScalingGroup {
	_init_.Initialize()

	if err := validateNewAutoScalingGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoScalingGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAutoScalingGroup_Override(a AutoScalingGroup, scope constructs.Construct, id *string, props *AutoScalingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AutoScalingGroup)SetAlbTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"albTargetGroup",
		val,
	)
}

func (j *jsiiProxy_AutoScalingGroup)SetHasCalledScaleOnRequestCount(val *bool) {
	if err := j.validateSetHasCalledScaleOnRequestCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasCalledScaleOnRequestCount",
		val,
	)
}

func (j *jsiiProxy_AutoScalingGroup)SetNewInstancesProtectedFromScaleIn(val *bool) {
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

func AutoScalingGroup_FromAutoScalingGroupName(scope constructs.Construct, id *string, autoScalingGroupName *string) IAutoScalingGroup {
	_init_.Initialize()

	if err := validateAutoScalingGroup_FromAutoScalingGroupNameParameters(scope, id, autoScalingGroupName); err != nil {
		panic(err)
	}
	var returns IAutoScalingGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"fromAutoScalingGroupName",
		[]interface{}{scope, id, autoScalingGroupName},
		&returns,
	)

	return returns
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
func AutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoScalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func AutoScalingGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAutoScalingGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AutoScalingGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAutoScalingGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func AutoScalingGroup_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook {
	if err := a.validateAddLifecycleHookParameters(id, props); err != nil {
		panic(err)
	}
	var returns LifecycleHook

	_jsii_.Invoke(
		a,
		"addLifecycleHook",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	if err := a.validateAddSecurityGroupParameters(securityGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := a.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddUserData(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addUserData",
		args,
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddWarmPool(options *WarmPoolOptions) WarmPool {
	if err := a.validateAddWarmPoolParameters(options); err != nil {
		panic(err)
	}
	var returns WarmPool

	_jsii_.Invoke(
		a,
		"addWarmPool",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions) {
	if err := a.validateApplyCloudFormationInitParameters(init, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyCloudFormationInit",
		[]interface{}{init, options},
	)
}

func (a *jsiiProxy_AutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AreNewInstancesProtectedFromScaleIn() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"areNewInstancesProtectedFromScaleIn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	if err := a.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ProtectNewInstancesFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"protectNewInstancesFromScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnCpuUtilizationParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnIncomingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnIncomingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	if err := a.validateScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns StepScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnOutgoingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnOutgoingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnRequestCountParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnRequestCount",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction {
	if err := a.validateScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	var returns ScheduledAction

	_jsii_.Invoke(
		a,
		"scaleOnSchedule",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


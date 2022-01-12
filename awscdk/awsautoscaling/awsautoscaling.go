package awsautoscaling

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// An adjustment.
//
// TODO: EXAMPLE
//
type AdjustmentTier struct {
	// What number to adjust the capacity with.
	//
	// The number is interpeted as an added capacity, a new fixed capacity or an
	// added percentage depending on the AdjustmentType value of the
	// StepScalingPolicy.
	//
	// Can be positive or negative.
	Adjustment *float64 `json:"adjustment"`
	// Lower bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is higher than this value.
	LowerBound *float64 `json:"lowerBound"`
	// Upper bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is lower than this value.
	UpperBound *float64 `json:"upperBound"`
}

// How adjustment numbers are interpreted.
//
// TODO: EXAMPLE
//
type AdjustmentType string

const (
	AdjustmentType_CHANGE_IN_CAPACITY AdjustmentType = "CHANGE_IN_CAPACITY"
	AdjustmentType_PERCENT_CHANGE_IN_CAPACITY AdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	AdjustmentType_EXACT_CAPACITY AdjustmentType = "EXACT_CAPACITY"
)

// Options for applying CloudFormation init to an instance or instance group.
//
// TODO: EXAMPLE
//
type ApplyCloudFormationInitOptions struct {
	// ConfigSet to activate.
	ConfigSets *[]*string `json:"configSets"`
	// Force instance replacement by embedding a config fingerprint.
	//
	// If `true` (the default), a hash of the config will be embedded into the
	// UserData, so that if the config changes, the UserData changes and
	// instances will be replaced (given an UpdatePolicy has been configured on
	// the AutoScalingGroup).
	//
	// If `false`, no such hash will be embedded, and if the CloudFormation Init
	// config changes nothing will happen to the running instances. If a
	// config update introduces errors, you will not notice until after the
	// CloudFormation deployment successfully finishes and the next instance
	// fails to launch.
	EmbedFingerprint *bool `json:"embedFingerprint"`
	// Don't fail the instance creation when cfn-init fails.
	//
	// You can use this to prevent CloudFormation from rolling back when
	// instances fail to start up, to help in debugging.
	IgnoreFailures *bool `json:"ignoreFailures"`
	// Include --role argument when running cfn-init and cfn-signal commands.
	//
	// This will be the IAM instance profile attached to the EC2 instance
	IncludeRole *bool `json:"includeRole"`
	// Include --url argument when running cfn-init and cfn-signal commands.
	//
	// This will be the cloudformation endpoint in the deployed region
	// e.g. https://cloudformation.us-east-1.amazonaws.com
	IncludeUrl *bool `json:"includeUrl"`
	// Print the results of running cfn-init to the Instance System Log.
	//
	// By default, the output of running cfn-init is written to a log file
	// on the instance. Set this to `true` to print it to the System Log
	// (visible from the EC2 Console), `false` to not print it.
	//
	// (Be aware that the system log is refreshed at certain points in
	// time of the instance life cycle, and successful execution may
	// not always show up).
	PrintLog *bool `json:"printLog"`
}

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
// TODO: EXAMPLE
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
	AutoScalingGroupArn() *string
	AutoScalingGroupName() *string
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	MaxInstanceLifetime() awscdk.Duration
	NewInstancesProtectedFromScaleIn() *bool
	SetNewInstancesProtectedFromScaleIn(val *bool)
	Node() constructs.Node
	OsType() awsec2.OperatingSystemType
	PhysicalName() *string
	Role() awsiam.IRole
	SpotPrice() *string
	Stack() awscdk.Stack
	UserData() awsec2.UserData
	AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	AddUserData(commands ...*string)
	ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AreNewInstancesProtectedFromScaleIn() *bool
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ProtectNewInstancesFromScaleIn()
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy
	ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy
	ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction
	ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy
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

func (j *jsiiProxy_AutoScalingGroup) SetAlbTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"albTargetGroup",
		val,
	)
}

func (j *jsiiProxy_AutoScalingGroup) SetNewInstancesProtectedFromScaleIn(val *bool) {
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

func AutoScalingGroup_FromAutoScalingGroupName(scope constructs.Construct, id *string, autoScalingGroupName *string) IAutoScalingGroup {
	_init_.Initialize()

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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AutoScalingGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Send a message to either an SQS queue or SNS topic when instances launch or terminate.
func (a *jsiiProxy_AutoScalingGroup) AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook {
	var returns LifecycleHook

	_jsii_.Invoke(
		a,
		"addLifecycleHook",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Add the security group to all instances via the launch configuration security groups array.
func (a *jsiiProxy_AutoScalingGroup) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

// Adds a statement to the IAM role assumed by instances of this fleet.
func (a *jsiiProxy_AutoScalingGroup) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

// Add command to the startup script of fleet instances.
//
// The command must be in the scripting language supported by the fleet's OS (i.e. Linux/Windows).
// Does nothing for imported ASGs.
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

// Use a CloudFormation Init configuration at instance startup.
//
// This does the following:
//
// - Attaches the CloudFormation Init metadata to the AutoScalingGroup resource.
// - Add commands to the UserData to run `cfn-init` and `cfn-signal`.
// - Update the instance's CreationPolicy to wait for `cfn-init` to finish
//    before reporting success.
func (a *jsiiProxy_AutoScalingGroup) ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions) {
	_jsii_.InvokeVoid(
		a,
		"applyCloudFormationInit",
		[]interface{}{init, options},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_AutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns `true` if newly-launched instances are protected from scale-in.
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

// Attach to ELBv2 Application Target Group.
func (a *jsiiProxy_AutoScalingGroup) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Attach to a classic load balancer.
func (a *jsiiProxy_AutoScalingGroup) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		a,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// Attach to ELBv2 Application Target Group.
func (a *jsiiProxy_AutoScalingGroup) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_AutoScalingGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_AutoScalingGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Ensures newly-launched instances are protected from scale-in.
func (a *jsiiProxy_AutoScalingGroup) ProtectNewInstancesFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"protectNewInstancesFromScaleIn",
		nil, // no parameters
	)
}

// Scale out or in to achieve a target CPU utilization.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in to achieve a target network ingress rate.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnIncomingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in, in response to a metric.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	var returns StepScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in to achieve a target network egress rate.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnOutgoingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in to achieve a target request handling rate.
//
// The AutoScalingGroup must have been attached to an Application Load Balancer
// in order to be able to call this.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnRequestCount",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in based on time.
func (a *jsiiProxy_AutoScalingGroup) ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction {
	var returns ScheduledAction

	_jsii_.Invoke(
		a,
		"scaleOnSchedule",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Scale out or in in order to keep a metric around a target value.
func (a *jsiiProxy_AutoScalingGroup) ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
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

// Properties of a Fleet.
//
// TODO: EXAMPLE
//
type AutoScalingGroupProps struct {
	// Whether the instances can initiate connections to anywhere by default.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	AssociatePublicIpAddress *bool `json:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*BlockDevice `json:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	GroupMetrics *[]GroupMetrics `json:"groupMetrics"`
	// Configuration for health checks.
	HealthCheck HealthCheck `json:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	IgnoreUnmodifiedSizeProperties *bool `json:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	InstanceMonitoring Monitoring `json:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	KeyName *string `json:"keyName"`
	// Maximum number of instances in the fleet.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The maximum amount of time that an instance can be in service.
	//
	// The maximum duration applies
	// to all current and future instances in the group. As an instance approaches its maximum duration,
	// it is terminated and replaced, and cannot be used again.
	//
	// You must specify a value of at least 604,800 seconds (7 days). To clear a previously set value,
	// leave this property undefined.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html
	//
	MaxInstanceLifetime awscdk.Duration `json:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
	MinCapacity *float64 `json:"minCapacity"`
	// Whether newly-launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// By default, Auto Scaling can terminate an instance at any time after launch
	// when scaling in an Auto Scaling Group, subject to the group's termination
	// policy. However, you may wish to protect newly-launched instances from
	// being scaled in if they are going to run critical applications that should
	// not be prematurely terminated.
	//
	// This flag must be enabled if the Auto Scaling Group will be associated with
	// an ECS Capacity Provider with managed termination protection.
	NewInstancesProtectedFromScaleIn *bool `json:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	Notifications *[]*NotificationConfiguration `json:"notifications"`
	// Configure waiting for signals during deployment.
	//
	// Use this to pause the CloudFormation deployment to wait for the instances
	// in the AutoScalingGroup to report successful startup during
	// creation and updates. The UserData script needs to invoke `cfn-signal`
	// with a success or failure code after it is done setting up the instance.
	//
	// Without waiting for signals, the CloudFormation deployment will proceed as
	// soon as the AutoScalingGroup has been created or updated but before the
	// instances in the group have been started.
	//
	// For example, to have instances wait for an Elastic Load Balancing health check before
	// they signal success, add a health-check verification by using the
	// cfn-init helper script. For an example, see the verify_instance_health
	// command in the Auto Scaling rolling updates sample template:
	//
	// https://github.com/awslabs/aws-cloudformation-templates/blob/master/aws/services/AutoScaling/AutoScalingRollingUpdates.yaml
	Signals Signals `json:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	SpotPrice *string `json:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	TerminationPolicies *[]TerminationPolicy `json:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	UpdatePolicy UpdatePolicy `json:"updatePolicy"`
	// Where to place instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// Type of instance to launch.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// AMI to launch.
	MachineImage awsec2.IMachineImage `json:"machineImage"`
	// VPC to launch these instances in.
	Vpc awsec2.IVpc `json:"vpc"`
	// Apply the given CloudFormation Init configuration to the instances in the AutoScalingGroup at startup.
	//
	// If you specify `init`, you must also specify `signals` to configure
	// the number of instances to wait for and the timeout for waiting for the
	// init process.
	Init awsec2.CloudFormationInit `json:"init"`
	// Use the given options for applying CloudFormation Init.
	//
	// Describes the configsets to use and the timeout to wait
	InitOptions *ApplyCloudFormationInitOptions `json:"initOptions"`
	// Whether IMDSv2 should be required on launched instances.
	RequireImdsv2 *bool `json:"requireImdsv2"`
	// An IAM role to associate with the instance profile assigned to this Auto Scaling Group.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// TODO: EXAMPLE
	//
	Role awsiam.IRole `json:"role"`
	// Security group to launch the instances in.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// Specific UserData to use.
	//
	// The UserData may still be mutated after creation.
	UserData awsec2.UserData `json:"userData"`
}

// Aspect that makes IMDSv2 required on instances deployed by AutoScalingGroups.
//
// TODO: EXAMPLE
//
type AutoScalingGroupRequireImdsv2Aspect interface {
	awscdk.IAspect
	Visit(node constructs.IConstruct)
	Warn(node constructs.IConstruct, message *string)
}

// The jsii proxy struct for AutoScalingGroupRequireImdsv2Aspect
type jsiiProxy_AutoScalingGroupRequireImdsv2Aspect struct {
	internal.Type__awscdkIAspect
}

func NewAutoScalingGroupRequireImdsv2Aspect() AutoScalingGroupRequireImdsv2Aspect {
	_init_.Initialize()

	j := jsiiProxy_AutoScalingGroupRequireImdsv2Aspect{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroupRequireImdsv2Aspect",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAutoScalingGroupRequireImdsv2Aspect_Override(a AutoScalingGroupRequireImdsv2Aspect) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.AutoScalingGroupRequireImdsv2Aspect",
		nil, // no parameters
		a,
	)
}

// All aspects can visit an IConstruct.
func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		a,
		"visit",
		[]interface{}{node},
	)
}

// Adds a warning annotation to a node.
func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) Warn(node constructs.IConstruct, message *string) {
	_jsii_.InvokeVoid(
		a,
		"warn",
		[]interface{}{node, message},
	)
}

// Base interface for target tracking props.
//
// Contains the attributes that are common to target tracking policies,
// except the ones relating to the metric and to the scalable target.
//
// This interface is reused by more specific target tracking props objects.
//
// TODO: EXAMPLE
//
type BaseTargetTrackingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
}

// Basic properties for a lifecycle hook.
//
// TODO: EXAMPLE
//
type BasicLifecycleHookProps struct {
	// The state of the Amazon EC2 instance to which you want to attach the lifecycle hook.
	LifecycleTransition LifecycleTransition `json:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	DefaultResult DefaultResult `json:"defaultResult"`
	// Maximum time between calls to RecordLifecycleActionHeartbeat for the hook.
	//
	// If the lifecycle hook times out, perform the action in DefaultResult.
	HeartbeatTimeout awscdk.Duration `json:"heartbeatTimeout"`
	// Name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName"`
	// Additional data to pass to the lifecycle hook target.
	NotificationMetadata *string `json:"notificationMetadata"`
	// The target of the lifecycle hook.
	NotificationTarget ILifecycleHookTarget `json:"notificationTarget"`
	// The role that allows publishing to the notification target.
	Role awsiam.IRole `json:"role"`
}

// Properties for a scheduled scaling action.
//
// TODO: EXAMPLE
//
type BasicScheduledActionProps struct {
	// When to perform this action.
	//
	// Supports cron expressions.
	//
	// For more information about cron expressions, see https://en.wikipedia.org/wiki/Cron.
	Schedule Schedule `json:"schedule"`
	// The new desired capacity.
	//
	// At the scheduled time, set the desired capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// When this scheduled action expires.
	EndTime *time.Time `json:"endTime"`
	// The new maximum capacity.
	//
	// At the scheduled time, set the maximum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The new minimum capacity.
	//
	// At the scheduled time, set the minimum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	MinCapacity *float64 `json:"minCapacity"`
	// When this scheduled action becomes active.
	StartTime *time.Time `json:"startTime"`
	// Specifies the time zone for a cron expression.
	//
	// If a time zone is not provided, UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the IANA Time Zone Database (such as Etc/GMT+9 or Pacific/Tahiti).
	//
	// For more information, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	TimeZone *string `json:"timeZone"`
}

// TODO: EXAMPLE
//
type BasicStepScalingPolicyProps struct {
	// Metric to scale on.
	Metric awscloudwatch.IMetric `json:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	ScalingSteps *[]*ScalingInterval `json:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	AdjustmentType AdjustmentType `json:"adjustmentType"`
	// Grace period after scaling activity.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	EvaluationPeriods *float64 `json:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude"`
}

// Properties for a Target Tracking policy that include the metric but exclude the target.
//
// TODO: EXAMPLE
//
type BasicTargetTrackingScalingPolicyProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// The target value for the metric.
	TargetValue *float64 `json:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	CustomMetric awscloudwatch.IMetric `json:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	PredefinedMetric PredefinedMetric `json:"predefinedMetric"`
	// The resource label associated with the predefined metric.
	//
	// Should be supplied if the predefined metric is ALBRequestCountPerTarget, and the
	// format should be:
	//
	// app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>
	ResourceLabel *string `json:"resourceLabel"`
}

// Options needed to bind a target to a lifecycle hook.
//
// [disable-awslint:ref-via-interface] The lifecycle hook to attach to and an IRole to use
//
// TODO: EXAMPLE
//
type BindHookTargetOptions struct {
	// The lifecycle hook to attach to.
	//
	// [disable-awslint:ref-via-interface]
	LifecycleHook LifecycleHook `json:"lifecycleHook"`
	// The role to use when attaching to the lifecycle hook.
	//
	// [disable-awslint:ref-via-interface]
	Role awsiam.IRole `json:"role"`
}

// Block device.
//
// TODO: EXAMPLE
//
type BlockDevice struct {
	// The device name exposed to the EC2 instance.
	//
	// Supply a value like `/dev/sdh`, `xvdh`.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html
	//
	DeviceName *string `json:"deviceName"`
	// Defines the block device volume, to be either an Amazon EBS volume or an ephemeral instance store volume.
	//
	// Supply a value like `BlockDeviceVolume.ebs(15)`, `BlockDeviceVolume.ephemeral(0)`.
	Volume BlockDeviceVolume `json:"volume"`
}

// Describes a block device mapping for an EC2 instance or Auto Scaling group.
//
// TODO: EXAMPLE
//
type BlockDeviceVolume interface {
	EbsDevice() *EbsDeviceProps
	VirtualName() *string
}

// The jsii proxy struct for BlockDeviceVolume
type jsiiProxy_BlockDeviceVolume struct {
	_ byte // padding
}

func (j *jsiiProxy_BlockDeviceVolume) EbsDevice() *EbsDeviceProps {
	var returns *EbsDeviceProps
	_jsii_.Get(
		j,
		"ebsDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockDeviceVolume) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}


func NewBlockDeviceVolume(ebsDevice *EbsDeviceProps, virtualName *string) BlockDeviceVolume {
	_init_.Initialize()

	j := jsiiProxy_BlockDeviceVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		&j,
	)

	return &j
}

func NewBlockDeviceVolume_Override(b BlockDeviceVolume, ebsDevice *EbsDeviceProps, virtualName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		b,
	)
}

// Creates a new Elastic Block Storage device.
func BlockDeviceVolume_Ebs(volumeSize *float64, options *EbsDeviceOptions) BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ebs",
		[]interface{}{volumeSize, options},
		&returns,
	)

	return returns
}

// Creates a new Elastic Block Storage device from an existing snapshot.
func BlockDeviceVolume_EbsFromSnapshot(snapshotId *string, options *EbsDeviceSnapshotOptions) BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ebsFromSnapshot",
		[]interface{}{snapshotId, options},
		&returns,
	)

	return returns
}

// Creates a virtual, ephemeral device.
//
// The name will be in the form ephemeral{volumeIndex}.
func BlockDeviceVolume_Ephemeral(volumeIndex *float64) BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ephemeral",
		[]interface{}{volumeIndex},
		&returns,
	)

	return returns
}

// Supresses a volume mapping.
func BlockDeviceVolume_NoDevice() BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"noDevice",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::AutoScaling::AutoScalingGroup`.
//
// The `AWS::AutoScaling::AutoScalingGroup` resource defines an Amazon EC2 Auto Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical grouping for the purposes of automatic scaling and management.
//
// > Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a launch template or a launch configuration. We recommend that you use a launch template to make sure that you can use the latest features of Amazon EC2, such as Dedicated Hosts and T2 Unlimited instances. For more information, see [Creating a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) . You can find sample launch templates in [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) .
//
// For more information, see [CreateAutoScalingGroup](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html) and [UpdateAutoScalingGroup](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_UpdateAutoScalingGroup.html) in the *Amazon EC2 Auto Scaling API Reference* . For more information about Amazon EC2 Auto Scaling, see the [Amazon EC2 Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html) .
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	CapacityRebalance() interface{}
	SetCapacityRebalance(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Context() *string
	SetContext(val *string)
	Cooldown() *string
	SetCooldown(val *string)
	CreationStack() *[]*string
	DesiredCapacity() *string
	SetDesiredCapacity(val *string)
	DesiredCapacityType() *string
	SetDesiredCapacityType(val *string)
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	InstanceId() *string
	SetInstanceId(val *string)
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	LifecycleHookSpecificationList() interface{}
	SetLifecycleHookSpecificationList(val interface{})
	LoadBalancerNames() *[]*string
	SetLoadBalancerNames(val *[]*string)
	LogicalId() *string
	MaxInstanceLifetime() *float64
	SetMaxInstanceLifetime(val *float64)
	MaxSize() *string
	SetMaxSize(val *string)
	MetricsCollection() interface{}
	SetMetricsCollection(val interface{})
	MinSize() *string
	SetMinSize(val *string)
	MixedInstancesPolicy() interface{}
	SetMixedInstancesPolicy(val interface{})
	NewInstancesProtectedFromScaleIn() interface{}
	SetNewInstancesProtectedFromScaleIn(val interface{})
	Node() constructs.Node
	NotificationConfigurations() interface{}
	SetNotificationConfigurations(val interface{})
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	Ref() *string
	ServiceLinkedRoleArn() *string
	SetServiceLinkedRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TargetGroupArns() *[]*string
	SetTargetGroupArns(val *[]*string)
	TerminationPolicies() *[]*string
	SetTerminationPolicies(val *[]*string)
	UpdatedProperites() *map[string]interface{}
	VpcZoneIdentifier() *[]*string
	SetVpcZoneIdentifier(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAutoScalingGroup
type jsiiProxy_CfnAutoScalingGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) DesiredCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LaunchTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LifecycleHookSpecificationList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHookSpecificationList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LoadBalancerNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MetricsCollection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MinSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MixedInstancesPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) NewInstancesProtectedFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) NotificationConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::AutoScalingGroup`.
func NewCfnAutoScalingGroup(scope constructs.Construct, id *string, props *CfnAutoScalingGroupProps) CfnAutoScalingGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnAutoScalingGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::AutoScalingGroup`.
func NewCfnAutoScalingGroup_Override(c CfnAutoScalingGroup, scope constructs.Construct, id *string, props *CfnAutoScalingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetCapacityRebalance(val interface{}) {
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetCooldown(val *string) {
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetDesiredCapacity(val *string) {
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetDesiredCapacityType(val *string) {
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetHealthCheckGracePeriod(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetHealthCheckType(val *string) {
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetLaunchConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetLaunchTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"launchTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetLifecycleHookSpecificationList(val interface{}) {
	_jsii_.Set(
		j,
		"lifecycleHookSpecificationList",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetLoadBalancerNames(val *[]*string) {
	_jsii_.Set(
		j,
		"loadBalancerNames",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetMaxInstanceLifetime(val *float64) {
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetMaxSize(val *string) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetMetricsCollection(val interface{}) {
	_jsii_.Set(
		j,
		"metricsCollection",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetMinSize(val *string) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetMixedInstancesPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"mixedInstancesPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetNewInstancesProtectedFromScaleIn(val interface{}) {
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetNotificationConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"notificationConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetPlacementGroup(val *string) {
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetServiceLinkedRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetTargetGroupArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetTerminationPolicies(val *[]*string) {
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup) SetVpcZoneIdentifier(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAutoScalingGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAutoScalingGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnAutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutoScalingGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAutoScalingGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnAutoScalingGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAutoScalingGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnAutoScalingGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAutoScalingGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAutoScalingGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnAutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnAutoScalingGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAutoScalingGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnAutoScalingGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAutoScalingGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnAutoScalingGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnAutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AcceleratorCountRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum number of accelerators for an instance type.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_AcceleratorCountRequestProperty struct {
	// The maximum value.
	Max *float64 `json:"max"`
	// The minimum value.
	Min *float64 `json:"min"`
}

// `AcceleratorTotalMemoryMiBRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum total memory size for the accelerators for an instance type, in MiB.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_AcceleratorTotalMemoryMiBRequestProperty struct {
	// The memory maximum in MiB.
	Max *float64 `json:"max"`
	// The memory minimum in MiB.
	Min *float64 `json:"min"`
}

// `BaselineEbsBandwidthMbpsRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum baseline bandwidth performance for an instance type, in Mbps.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_BaselineEbsBandwidthMbpsRequestProperty struct {
	// The maximum value in Mbps.
	Max *float64 `json:"max"`
	// The minimum value in Mbps.
	Min *float64 `json:"min"`
}

// `InstanceRequirements` specifies a set of requirements for the types of instances that can be launched by an `AWS::AutoScaling::AutoScalingGroup` resource.
//
// `InstanceRequirements` is a property of the `LaunchTemplateOverrides` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) property type.
//
// You must specify `VCpuCount` and `MemoryMiB` , but all other properties are optional. Any unspecified optional property is set to its default.
//
// When you specify multiple properties, you get instance types that satisfy all of the specified properties. If you specify multiple values for a property, you get instance types that satisfy any of the specified values.
//
// For more information, see [Creating an Auto Scaling group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_InstanceRequirementsProperty struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.
	//
	// To exclude accelerator-enabled instance types, set `Max` to `0` .
	//
	// Default: No minimum or maximum
	AcceleratorCount interface{} `json:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// - For instance types with NVIDIA devices, specify `nvidia` .
	// - For instance types with AMD devices, specify `amd` .
	// - For instance types with AWS devices, specify `amazon-web-services` .
	// - For instance types with Xilinx devices, specify `xilinx` .
	//
	// Default: Any manufacturer
	AcceleratorManufacturers *[]*string `json:"acceleratorManufacturers"`
	// Lists the accelerators that must be on an instance type.
	//
	// - For instance types with NVIDIA A100 GPUs, specify `a100` .
	// - For instance types with NVIDIA V100 GPUs, specify `v100` .
	// - For instance types with NVIDIA K80 GPUs, specify `k80` .
	// - For instance types with NVIDIA T4 GPUs, specify `t4` .
	// - For instance types with NVIDIA M60 GPUs, specify `m60` .
	// - For instance types with AMD Radeon Pro V520 GPUs, specify `radeon-pro-v520` .
	// - For instance types with Xilinx VU9P FPGAs, specify `vu9p` .
	//
	// Default: Any accelerator
	AcceleratorNames *[]*string `json:"acceleratorNames"`
	// The minimum and maximum total memory size for the accelerators on an instance type, in MiB.
	//
	// Default: No minimum or maximum
	AcceleratorTotalMemoryMiB interface{} `json:"acceleratorTotalMemoryMiB"`
	// Lists the accelerator types that must be on an instance type.
	//
	// - For instance types with GPU accelerators, specify `gpu` .
	// - For instance types with FPGA accelerators, specify `fpga` .
	// - For instance types with inference accelerators, specify `inference` .
	//
	// Default: Any accelerator type
	AcceleratorTypes *[]*string `json:"acceleratorTypes"`
	// Indicates whether bare metal instance types are included, excluded, or required.
	//
	// Default: `excluded`
	BareMetal *string `json:"bareMetal"`
	// The minimum and maximum baseline bandwidth performance for an instance type, in Mbps.
	//
	// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: No minimum or maximum
	BaselineEbsBandwidthMbps interface{} `json:"baselineEbsBandwidthMbps"`
	// Indicates whether burstable performance instance types are included, excluded, or required.
	//
	// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: `excluded`
	BurstablePerformance *string `json:"burstablePerformance"`
	// Lists which specific CPU manufacturers to include.
	//
	// - For instance types with Intel CPUs, specify `intel` .
	// - For instance types with AMD CPUs, specify `amd` .
	// - For instance types with AWS CPUs, specify `amazon-web-services` .
	//
	// > Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	//
	// Default: Any manufacturer
	CpuManufacturers *[]*string `json:"cpuManufacturers"`
	// Lists which instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ). The following are examples: `c5*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` , you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , you are excluding all the M5a instance types, but not the M5n instance types.
	//
	// Default: No excluded instance types
	ExcludedInstanceTypes *[]*string `json:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// - For current generation instance types, specify `current` . The current generation includes EC2 instance types currently recommended for use. This typically includes the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// - For previous generation instance types, specify `previous` .
	//
	// Default: Any current or previous generation
	InstanceGenerations *[]*string `json:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, see [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: `included`
	LocalStorage *string `json:"localStorage"`
	// Indicates the type of local storage that is required.
	//
	// - For instance types with hard disk drive (HDD) storage, specify `hdd` .
	// - For instance types with solid state drive (SSD) storage, specify `sdd` .
	//
	// Default: Any local storage type
	LocalStorageTypes *[]*string `json:"localStorageTypes"`
	// The minimum and maximum amount of memory per vCPU for an instance type, in GiB.
	//
	// Default: No minimum or maximum
	MemoryGiBPerVCpu interface{} `json:"memoryGiBPerVCpu"`
	// The minimum and maximum instance memory size for an instance type, in MiB.
	MemoryMiB interface{} `json:"memoryMiB"`
	// The minimum and maximum number of network interfaces for an instance type.
	//
	// Default: No minimum or maximum
	NetworkInterfaceCount interface{} `json:"networkInterfaceCount"`
	// The price protection threshold for On-Demand Instances.
	//
	// This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as `999999` .
	//
	// Default: `20`
	OnDemandMaxPricePercentageOverLowestPrice *float64 `json:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must provide On-Demand Instance hibernation support.
	//
	// Default: `false`
	RequireHibernateSupport interface{} `json:"requireHibernateSupport"`
	// The price protection threshold for Spot Instances.
	//
	// This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as `999999` .
	//
	// Default: `100`
	SpotMaxPricePercentageOverLowestPrice *float64 `json:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum total local storage size for an instance type, in GB.
	//
	// Default: No minimum or maximum
	TotalLocalStorageGb interface{} `json:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs for an instance type.
	VCpuCount interface{} `json:"vCpuCount"`
}

// `InstancesDistribution` is a property of the [AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html) property type that describes an instances distribution for an Auto Scaling group. All properties have a default value, which is the value that is used or assumed when the property is not specified.
//
// The instances distribution specifies the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances, and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities.
//
// For more information and example configurations, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_InstancesDistributionProperty struct {
	// The order of the launch template overrides to use in fulfilling On-Demand capacity.
	//
	// If you specify `lowest-price` , Amazon EC2 Auto Scaling uses price to determine the order, launching the lowest price first.
	//
	// If you specify `prioritized` , Amazon EC2 Auto Scaling uses the priority that you assigned to each launch template override, launching the highest priority first. If all your On-Demand capacity cannot be fulfilled using your highest priority instance, then Amazon EC2 Auto Scaling launches the remaining capacity using the second priority instance type, and so on.
	//
	// Default: `lowest-price` for Auto Scaling groups that specify the `InstanceRequirements` property in the overrides and `prioritized` for Auto Scaling groups that don't.
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy"`
	// The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances.
	//
	// This base portion is launched first as your group scales.
	//
	// If you specify weights for the instance types in the overrides, the base capacity is measured in the same unit of measurement as the instance types. If you specify the `InstanceRequirements` property in the overrides, the base capacity is measured in the same unit of measurement as your group's desired capacity.
	//
	// Default: `0`
	//
	// > An update to this setting means a gradual replacement of instances to adjust the current On-Demand Instance levels. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
	OnDemandBaseCapacity *float64 `json:"onDemandBaseCapacity"`
	// Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond `OnDemandBaseCapacity` .
	//
	// Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). If set to 100, only On-Demand Instances are used.
	//
	// Default: `100`
	//
	// > An update to this setting means a gradual replacement of instances to adjust the current On-Demand and Spot Instance levels for your additional capacity higher than the base capacity. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
	OnDemandPercentageAboveBaseCapacity *float64 `json:"onDemandPercentageAboveBaseCapacity"`
	// If the allocation strategy is `lowest-price` , the Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify.
	//
	// If the allocation strategy is `capacity-optimized` (recommended), the Auto Scaling group launches instances using Spot pools that are optimally chosen based on the available Spot capacity. Alternatively, you can use `capacity-optimized-prioritized` and set the order of instance types in the list of launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling honors the instance type priorities on a best-effort basis but optimizes for capacity first.
	//
	// Default: `lowest-price`
	//
	// Valid values: `lowest-price` | `capacity-optimized` | `capacity-optimized-prioritized`
	SpotAllocationStrategy *string `json:"spotAllocationStrategy"`
	// The number of Spot Instance pools to use to allocate your Spot capacity.
	//
	// The Spot pools are determined from the different instance types in the overrides. Valid only when the Spot allocation strategy is `lowest-price` . Value must be in the range of 1–20.
	//
	// Default: `2`
	SpotInstancePools *float64 `json:"spotInstancePools"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// If you leave the value at its default (empty), Amazon EC2 Auto Scaling uses the On-Demand price as the maximum Spot price. To remove a value that you previously set, include the property but specify an empty string ("") for the value.
	//
	// > If your maximum price is lower than the Spot price for the instance types that you selected, your Spot Instances are not launched.
	//
	// Valid Range: Minimum value of 0.001
	SpotMaxPrice *string `json:"spotMaxPrice"`
}

// `LaunchTemplateOverrides` is a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) property type that describes an override for a launch template.
//
// If you supply your own instance types, the maximum number of instance types that can be associated with an Auto Scaling group is 40. The maximum number of distinct launch templates you can define for an Auto Scaling group is 20.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_LaunchTemplateOverridesProperty struct {
	// The instance requirements.
	//
	// When you specify instance requirements, Amazon EC2 Auto Scaling finds instance types that satisfy your requirements, and then uses your On-Demand and Spot allocation strategies to launch instances from these instance types, in the same way as when you specify a list of specific instance types.
	//
	// > `InstanceRequirements` are incompatible with the `InstanceType` property. If you specify both of these properties, Amazon EC2 Auto Scaling will return a `ValidationException` exception.
	InstanceRequirements interface{} `json:"instanceRequirements"`
	// The instance type, such as `m3.xlarge` . You must use an instance type that is supported in your requested Region and Availability Zones. For more information, see [Available instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the *Amazon EC2 User Guide for Linux Instances.*.
	InstanceType *string `json:"instanceType"`
	// Provides the launch template to be used when launching the instance type specified in `InstanceType` .
	//
	// For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's defined for your mixed instances policy. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide* .
	LaunchTemplateSpecification interface{} `json:"launchTemplateSpecification"`
	// The number of capacity units provided by the instance type specified in `InstanceType` in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is provisioned, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling provisions instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the desired capacity is exceeded by 3 units. For more information, see [Instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide* . Value must be in the range of 1-999.
	//
	// > Every Auto Scaling group has three size parameters ( `DesiredCapacity` , `MaxSize` , and `MinSize` ). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.
	WeightedCapacity *string `json:"weightedCapacity"`
}

// `LaunchTemplate` is a property of the [AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html) property type that describes a launch template and overrides. The overrides are used to override the instance type specified by the launch template with multiple instance types that can be used to launch On-Demand Instances and Spot Instances.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_LaunchTemplateProperty struct {
	// The launch template to use.
	LaunchTemplateSpecification interface{} `json:"launchTemplateSpecification"`
	// Any properties that you specify override the same properties in the launch template.
	//
	// If not provided, Amazon EC2 Auto Scaling uses the instance type or instance requirements specified in the launch template when it launches an instance.
	//
	// The overrides can include either one or more instance types or a set of instance requirements, but not both.
	Overrides interface{} `json:"overrides"`
}

// `LaunchTemplateSpecification` specifies a launch template and version for the `LaunchTemplate` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource. It is also a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) and [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property types.
//
// The launch template that is specified must be configured for use with an Auto Scaling group. For information about creating a launch template, see [Creating a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// You can find a sample template snippets in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` documentation and in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#aws-resource-ec2-launchtemplate--examples) section of the `AWS::EC2::LaunchTemplate` documentation.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_LaunchTemplateSpecificationProperty struct {
	// The version number.
	//
	// CloudFormation does not support specifying $Latest, or $Default for the template version number. However, you can specify `LatestVersionNumber` or `DefaultVersionNumber` using the `Fn::GetAtt` function.
	//
	// > For an example of using the `Fn::GetAtt` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` documentation.
	Version *string `json:"version"`
	// The ID of the [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) . You must specify either a `LaunchTemplateName` or a `LaunchTemplateId` .
	LaunchTemplateId *string `json:"launchTemplateId"`
	// The name of the [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) . You must specify either a `LaunchTemplateName` or a `LaunchTemplateId` .
	LaunchTemplateName *string `json:"launchTemplateName"`
}

// `LifecycleHookSpecification` specifies a lifecycle hook for the `LifecycleHookSpecificationList` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource. A lifecycle hook specifies actions to perform when Amazon EC2 Auto Scaling launches or terminates instances.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#aws-resource-as-lifecyclehook--examples) section of the `AWS::AutoScaling::LifecycleHook` documentation.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_LifecycleHookSpecificationProperty struct {
	// The name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName"`
	// The state of the EC2 instance to attach the lifecycle hook to. The valid values are:.
	//
	// - autoscaling:EC2_INSTANCE_LAUNCHING
	// - autoscaling:EC2_INSTANCE_TERMINATING
	LifecycleTransition *string `json:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	//
	// The valid values are `CONTINUE` and `ABANDON` (default).
	//
	// For more information, see [Adding lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/adding-lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* .
	DefaultResult *string `json:"defaultResult"`
	// The maximum time, in seconds, that can elapse before the lifecycle hook times out.
	//
	// If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the default action.
	HeartbeatTimeout *float64 `json:"heartbeatTimeout"`
	// Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata"`
	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook.
	//
	// You can specify an Amazon SQS queue or an Amazon SNS topic.
	NotificationTargetArn *string `json:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
	//
	// For information about creating this role, see [Configuring a notification target for a lifecycle hook](https://docs.aws.amazon.com/autoscaling/ec2/userguide/configuring-lifecycle-hook-notifications.html) in the *Amazon EC2 Auto Scaling User Guide* .
	RoleArn *string `json:"roleArn"`
}

// `MemoryGiBPerVCpuRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum amount of memory per vCPU for an instance type, in GiB.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_MemoryGiBPerVCpuRequestProperty struct {
	// The memory maximum in GiB.
	Max *float64 `json:"max"`
	// The memory minimum in GiB.
	Min *float64 `json:"min"`
}

// `MemoryMiBRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum instance memory size for an instance type, in MiB.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_MemoryMiBRequestProperty struct {
	// The memory maximum in MiB.
	Max *float64 `json:"max"`
	// The memory minimum in MiB.
	Min *float64 `json:"min"`
}

// `MetricsCollection` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource that describes the group metrics that an Amazon EC2 Auto Scaling group sends to Amazon CloudWatch. These metrics describe the group rather than any of its instances.
//
// For more information, see [Monitoring CloudWatch metrics for your Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` documentation.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_MetricsCollectionProperty struct {
	// The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch.
	//
	// *Allowed Values* : `1Minute`
	Granularity *string `json:"granularity"`
	// Specifies which group-level metrics to start collecting.
	//
	// *Allowed Values* :
	//
	// - `GroupMinSize`
	// - `GroupMaxSize`
	// - `GroupDesiredCapacity`
	// - `GroupInServiceInstances`
	// - `GroupPendingInstances`
	// - `GroupStandbyInstances`
	// - `GroupTerminatingInstances`
	// - `GroupTotalInstances`
	// - `GroupInServiceCapacity`
	// - `GroupPendingCapacity`
	// - `GroupStandbyCapacity`
	// - `GroupTerminatingCapacity`
	// - `GroupTotalCapacity`
	// - `WarmPoolDesiredCapacity`
	// - `WarmPoolWarmedCapacity`
	// - `WarmPoolPendingCapacity`
	// - `WarmPoolTerminatingCapacity`
	// - `WarmPoolTotalCapacity`
	// - `GroupAndWarmPoolDesiredCapacity`
	// - `GroupAndWarmPoolTotalCapacity`
	//
	// If you specify `Granularity` and don't specify any metrics, all metrics are enabled.
	Metrics *[]*string `json:"metrics"`
}

// `MixedInstancesPolicy` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource. It allows you to configure a group that diversifies across On-Demand Instances and Spot Instances of multiple instance types. For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// You can create a mixed instances policy for a new Auto Scaling group, or you can create it for an existing group by updating the group to specify `MixedInstancesPolicy` as the top-level property instead of a launch template or launch configuration. If you specify a `MixedInstancesPolicy` , you must specify a launch template as a property of the policy. You cannot specify a launch configuration for the policy.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_MixedInstancesPolicyProperty struct {
	// Specifies the launch template to use and optionally the instance types (overrides) that are used to provision EC2 instances to fulfill On-Demand and Spot capacities.
	LaunchTemplate interface{} `json:"launchTemplate"`
	// The instances distribution to use.
	//
	// If you leave this property unspecified, the value for each property in `InstancesDistribution` uses a default value.
	InstancesDistribution interface{} `json:"instancesDistribution"`
}

// `NetworkInterfaceCountRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum number of network interfaces for an instance type.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_NetworkInterfaceCountRequestProperty struct {
	// The maximum number of network interfaces.
	Max *float64 `json:"max"`
	// The minimum number of network interfaces.
	Min *float64 `json:"min"`
}

// `NotificationConfiguration` specifies a notification configuration for the `NotificationConfigurations` property of [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) . `NotificationConfiguration` specifies the events that the Amazon EC2 Auto Scaling group sends notifications for.
//
// For example snippets, see [Declaring an Auto Scaling group with a launch template and notifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html#scenario-as-notification) .
//
// For more information, see [Getting Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_NotificationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic.
	TopicArn *string `json:"topicArn"`
	// A list of event types that trigger a notification. Event types can include any of the following types.
	//
	// *Allowed Values* :
	//
	// - `autoscaling:EC2_INSTANCE_LAUNCH`
	// - `autoscaling:EC2_INSTANCE_LAUNCH_ERROR`
	// - `autoscaling:EC2_INSTANCE_TERMINATE`
	// - `autoscaling:EC2_INSTANCE_TERMINATE_ERROR`
	// - `autoscaling:TEST_NOTIFICATION`
	NotificationTypes *[]*string `json:"notificationTypes"`
}

// `TagProperty` specifies a tag for the `Tags` property of [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) . `TagProperty` adds tags to all associated instances in an Auto Scaling group.
//
// For more information, see [Tagging Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` documentation.
//
// CloudFormation adds the following tags to all Auto Scaling groups and associated instances:
//
// - aws:cloudformation:stack-name
// - aws:cloudformation:stack-id
// - aws:cloudformation:logical-id
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_TagPropertyProperty struct {
	// The tag key.
	Key *string `json:"key"`
	// Set to `true` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group.
	//
	// Set to `false` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.
	PropagateAtLaunch interface{} `json:"propagateAtLaunch"`
	// The tag value.
	Value *string `json:"value"`
}

// `TotalLocalStorageGBRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum total local storage size for an instance type, in GB.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_TotalLocalStorageGBRequestProperty struct {
	// The storage maximum in GB.
	Max *float64 `json:"max"`
	// The storage minimum in GB.
	Min *float64 `json:"min"`
}

// `VCpuCountRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum number of vCPUs for an instance type.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroup_VCpuCountRequestProperty struct {
	// The maximum number of vCPUs.
	Max *float64 `json:"max"`
	// The minimum number of vCPUs.
	Min *float64 `json:"min"`
}

// Properties for defining a `CfnAutoScalingGroup`.
//
// TODO: EXAMPLE
//
type CfnAutoScalingGroupProps struct {
	// The maximum size of the group.
	//
	// > With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to go above `MaxSize` to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will never go above `MaxSize` by more than your largest instance weight (weights that define how many units each instance contributes to the desired capacity of the group).
	MaxSize *string `json:"maxSize"`
	// The minimum size of the group.
	MinSize *string `json:"minSize"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// A list of Availability Zones where instances in the Auto Scaling group can be created.
	//
	// You must specify one of the following properties: `VPCZoneIdentifier` or `AvailabilityZones` . If your account supports EC2-Classic and VPC, this property is required to create an Auto Scaling group that launches instances into EC2-Classic.
	AvailabilityZones *[]*string `json:"availabilityZones"`
	// Indicates whether Capacity Rebalancing is enabled.
	//
	// For more information, see [Amazon EC2 Auto Scaling Capacity Rebalancing](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the *Amazon EC2 Auto Scaling User Guide* .
	CapacityRebalance interface{} `json:"capacityRebalance"`
	// Reserved.
	Context *string `json:"context"`
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	//
	// The default value is `300` . This setting applies when using simple scaling policies, but not when using other scaling policies or scheduled scaling. For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	Cooldown *string `json:"cooldown"`
	// The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain.
	//
	// It can scale beyond this capacity if you configure automatic scaling.
	//
	// The number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity when creating the stack, the default is the minimum size of the group.
	//
	// CloudFormation marks the Auto Scaling group as successful (by setting its status to CREATE_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set in the launch template or launch configuration that you specified, then desired capacity is not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance capacity and your maximum price.
	DesiredCapacity *string `json:"desiredCapacity"`
	// The unit of measurement for the value specified for desired capacity.
	//
	// Amazon EC2 Auto Scaling supports `DesiredCapacityType` for attribute-based instance type selection only. For more information, see [Creating an Auto Scaling group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// By default, Amazon EC2 Auto Scaling specifies `units` , which translates into number of instances.
	//
	// Valid values: `units` | `vcpu` | `memory-mib`
	DesiredCapacityType *string `json:"desiredCapacityType"`
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	//
	// The default value is `0` . For more information, see [Health checks for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you are adding an `ELB` health check, you must specify this property.
	HealthCheckGracePeriod *float64 `json:"healthCheckGracePeriod"`
	// The service to use for the health checks.
	//
	// The valid values are `EC2` (default) and `ELB` . If you configure an Auto Scaling group to use load balancer (ELB) health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks. For more information, see [Health checks for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the *Amazon EC2 Auto Scaling User Guide* .
	HealthCheckType *string `json:"healthCheckType"`
	// The ID of the instance used to base the launch configuration on.
	//
	// If specified, Amazon EC2 Auto Scaling uses the configuration values from the specified instance to create a new launch configuration. For more information, see [Creating an Auto Scaling group using an EC2 instance](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// To get the instance ID, use the EC2 [DescribeInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html) API operation.
	//
	// If you specify `LaunchTemplate` , `MixedInstancesPolicy` , or `LaunchConfigurationName` , don't specify `InstanceId` .
	InstanceId *string `json:"instanceId"`
	// The name of the [launch configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html) to use to launch instances.
	//
	// If you specify `LaunchTemplate` , `MixedInstancesPolicy` , or `InstanceId` , don't specify `LaunchConfigurationName` .
	LaunchConfigurationName *string `json:"launchConfigurationName"`
	// Properties used to specify the [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) and version to use to launch instances. You can alternatively associate a launch template to the Auto Scaling group by specifying a `MixedInstancesPolicy` .
	//
	// If you omit this property, you must specify `MixedInstancesPolicy` , `LaunchConfigurationName` , or `InstanceId` .
	LaunchTemplate interface{} `json:"launchTemplate"`
	// One or more lifecycle hooks for the group, which specify actions to perform when Amazon EC2 Auto Scaling launches or terminates instances.
	LifecycleHookSpecificationList interface{} `json:"lifecycleHookSpecificationList"`
	// A list of Classic Load Balancers associated with this Auto Scaling group.
	//
	// For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers, specify the `TargetGroupARNs` property instead.
	LoadBalancerNames *[]*string `json:"loadBalancerNames"`
	// The maximum amount of time, in seconds, that an instance can be in service.
	//
	// The default is null. If specified, the value must be either 0 or a number equal to or greater than 86,400 seconds (1 day). For more information, see [Replacing Auto Scaling instances based on maximum instance lifetime](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html) in the *Amazon EC2 Auto Scaling User Guide* .
	MaxInstanceLifetime *float64 `json:"maxInstanceLifetime"`
	// Enables the monitoring of group metrics of an Auto Scaling group.
	//
	// By default, these metrics are disabled.
	MetricsCollection interface{} `json:"metricsCollection"`
	// An embedded object that specifies a mixed instances policy.
	//
	// The policy includes properties that not only define the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but also the properties that specify the instance configuration information—the launch template and instance types. The policy can also include a weight for each instance type and different launch templates for individual instance types.
	//
	// For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify `LaunchTemplate` , `InstanceId` , or `LaunchConfigurationName` , don't specify `MixedInstancesPolicy` .
	MixedInstancesPolicy interface{} `json:"mixedInstancesPolicy"`
	// Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// For more information about preventing instances from terminating on scale in, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *Amazon EC2 Auto Scaling User Guide* .
	NewInstancesProtectedFromScaleIn interface{} `json:"newInstancesProtectedFromScaleIn"`
	// Configures an Auto Scaling group to send notifications when specified events take place.
	NotificationConfigurations interface{} `json:"notificationConfigurations"`
	// The name of the placement group into which you want to launch your instances.
	//
	// A placement group is a logical grouping of instances within a single Availability Zone. For more information, see [Placement Groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances* .
	PlacementGroup *string `json:"placementGroup"`
	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
	//
	// By default, Amazon EC2 Auto Scaling uses a service-linked role named `AWSServiceRoleForAutoScaling` , which it creates if it does not exist. For more information, see [Service-linked roles for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	ServiceLinkedRoleArn *string `json:"serviceLinkedRoleArn"`
	// One or more tags.
	//
	// You can tag your Auto Scaling group and propagate the tags to the Amazon EC2 instances it launches. For more information, see [Tagging Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide* .
	Tags *[]*CfnAutoScalingGroup_TagPropertyProperty `json:"tags"`
	// One or more Amazon Resource Names (ARN) of load balancer target groups to associate with the Auto Scaling group.
	//
	// Instances are registered as targets in a target group, and traffic is routed to the target group. For more information, see [Elastic Load Balancing and Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html) in the *Amazon EC2 Auto Scaling User Guide* .
	TargetGroupArns *[]*string `json:"targetGroupArns"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them. The termination policies supported by Amazon EC2 Auto Scaling: `OldestInstance` , `OldestLaunchConfiguration` , `NewestInstance` , `ClosestToNextInstanceHour` , `Default` , `OldestLaunchTemplate` , and `AllocationStrategy` . For more information, see [Controlling which Auto Scaling instances terminate during scale in](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html) in the *Amazon EC2 Auto Scaling User Guide* .
	TerminationPolicies *[]*string `json:"terminationPolicies"`
	// A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.
	//
	// If you specify `VPCZoneIdentifier` with `AvailabilityZones` , the subnets that you specify for this property must reside in those Availability Zones.
	//
	// If this resource specifies public subnets and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html) .
	//
	// > When you update `VPCZoneIdentifier` , this retains the same Auto Scaling group and replaces old instances with new ones, according to the specified subnets. You can optionally specify how CloudFormation handles these updates by using an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) .
	VpcZoneIdentifier *[]*string `json:"vpcZoneIdentifier"`
}

// A CloudFormation `AWS::AutoScaling::LaunchConfiguration`.
//
// The `AWS::AutoScaling::LaunchConfiguration` resource specifies the launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.
//
// When you update the launch configuration for an Auto Scaling group, CloudFormation deletes that resource and creates a new launch configuration with the updated properties and a new name. Existing instances are not affected. To update existing instances when you update the `AWS::AutoScaling::LaunchConfiguration` resource, you can specify an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) for the group. You can find sample update policies for rolling updates in [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) .
//
// For more information, see [CreateLaunchConfiguration](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateLaunchConfiguration.html) in the *Amazon EC2 Auto Scaling API Reference* and [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchConfiguration.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// > To configure Amazon EC2 instances launched as part of the Auto Scaling group, you can specify a launch template or a launch configuration. We recommend that you use a [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) to make sure that you can use the latest features of Amazon EC2, such as Dedicated Hosts and T2 Unlimited instances. For more information, see [Creating a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) .
//
// TODO: EXAMPLE
//
type CfnLaunchConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClassicLinkVpcId() *string
	SetClassicLinkVpcId(val *string)
	ClassicLinkVpcSecurityGroups() *[]*string
	SetClassicLinkVpcSecurityGroups(val *[]*string)
	CreationStack() *[]*string
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	ImageId() *string
	SetImageId(val *string)
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceMonitoring() interface{}
	SetInstanceMonitoring(val interface{})
	InstanceType() *string
	SetInstanceType(val *string)
	KernelId() *string
	SetKernelId(val *string)
	KeyName() *string
	SetKeyName(val *string)
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
	LogicalId() *string
	MetadataOptions() interface{}
	SetMetadataOptions(val interface{})
	Node() constructs.Node
	PlacementTenancy() *string
	SetPlacementTenancy(val *string)
	RamDiskId() *string
	SetRamDiskId(val *string)
	Ref() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SpotPrice() *string
	SetSpotPrice(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	UserData() *string
	SetUserData(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLaunchConfiguration
type jsiiProxy_CfnLaunchConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchConfiguration) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) BlockDeviceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ClassicLinkVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicLinkVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ClassicLinkVpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classicLinkVpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) MetadataOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::LaunchConfiguration`.
func NewCfnLaunchConfiguration(scope constructs.Construct, id *string, props *CfnLaunchConfigurationProps) CfnLaunchConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunchConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::LaunchConfiguration`.
func NewCfnLaunchConfiguration_Override(c CfnLaunchConfiguration, scope constructs.Construct, id *string, props *CfnLaunchConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetAssociatePublicIpAddress(val interface{}) {
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetBlockDeviceMappings(val interface{}) {
	_jsii_.Set(
		j,
		"blockDeviceMappings",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetClassicLinkVpcId(val *string) {
	_jsii_.Set(
		j,
		"classicLinkVpcId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetClassicLinkVpcSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"classicLinkVpcSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetEbsOptimized(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetIamInstanceProfile(val *string) {
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetInstanceMonitoring(val interface{}) {
	_jsii_.Set(
		j,
		"instanceMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetKernelId(val *string) {
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetKeyName(val *string) {
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetLaunchConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetMetadataOptions(val interface{}) {
	_jsii_.Set(
		j,
		"metadataOptions",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetPlacementTenancy(val *string) {
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetRamDiskId(val *string) {
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetSpotPrice(val *string) {
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration) SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunchConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunchConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnLaunchConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnLaunchConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLaunchConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnLaunchConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLaunchConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnLaunchConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLaunchConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLaunchConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnLaunchConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnLaunchConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLaunchConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnLaunchConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLaunchConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnLaunchConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnLaunchConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `BlockDeviceMapping` specifies a block device mapping for the `BlockDeviceMappings` property of the [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html) resource.
//
// Each instance that is launched has an associated root device volume, either an Amazon EBS volume or an instance store volume. You can use block device mappings to specify additional EBS volumes or instance store volumes to attach to an instance when it is launched.
//
// For more information, see [Example block device mapping](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html#block-device-mapping-ex) in the *Amazon EC2 User Guide for Linux Instances* .
//
// TODO: EXAMPLE
//
type CfnLaunchConfiguration_BlockDeviceMappingProperty struct {
	// The device name exposed to the EC2 instance (for example, `/dev/sdh` or `xvdh` ).
	//
	// For more information, see [Device naming on Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html) in the *Amazon EC2 User Guide for Linux Instances* .
	DeviceName *string `json:"deviceName"`
	// Parameters used to automatically set up EBS volumes when an instance is launched.
	//
	// You can specify either `VirtualName` or `Ebs` , but not both.
	Ebs interface{} `json:"ebs"`
	// Setting this value to `true` suppresses the specified device included in the block device mapping of the AMI.
	//
	// If `NoDevice` is `true` for the root device, instances might fail the EC2 health check. In that case, Amazon EC2 Auto Scaling launches replacement instances.
	//
	// If you specify `NoDevice` , you cannot specify `Ebs` .
	NoDevice interface{} `json:"noDevice"`
	// The name of the virtual device.
	//
	// The name must be in the form ephemeral *X* where *X* is a number starting from zero (0), for example, `ephemeral0` .
	//
	// You can specify either `VirtualName` or `Ebs` , but not both.
	VirtualName *string `json:"virtualName"`
}

// `BlockDevice` is a property of the `EBS` property of the [AWS::AutoScaling::LaunchConfiguration BlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html) property type that describes an Amazon EBS volume.
//
// TODO: EXAMPLE
//
type CfnLaunchConfiguration_BlockDeviceProperty struct {
	// Indicates whether the volume is deleted on instance termination.
	//
	// For Amazon EC2 Auto Scaling, the default value is `true` .
	DeleteOnTermination interface{} `json:"deleteOnTermination"`
	// Specifies whether the volume should be encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances) . If your AMI uses encrypted volumes, you can also only launch it on supported instance types.
	//
	// > If you are creating a volume from a snapshot, you cannot create an unencrypted volume from an encrypted snapshot. Also, you cannot specify a KMS key ID when using a launch configuration.
	// >
	// > If you enable encryption by default, the EBS volumes that you create are always encrypted, either using the AWS managed KMS key or a customer-managed KMS key, regardless of whether the snapshot was encrypted.
	// >
	// > For more information, see [Using AWS KMS keys to encrypt Amazon EBS volumes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-data-protection.html#encryption) in the *Amazon EC2 Auto Scaling User Guide* .
	Encrypted interface{} `json:"encrypted"`
	// The number of input/output (I/O) operations per second (IOPS) to provision for the volume.
	//
	// For `gp3` and `io1` volumes, this represents the number of IOPS that are provisioned for the volume. For `gp2` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//
	// The following are the supported values for each volume type:
	//
	// - `gp3` : 3,000-16,000 IOPS
	// - `io1` : 100-64,000 IOPS
	//
	// For `io1` volumes, we guarantee 64,000 IOPS only for [Instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) . Other instance families guarantee performance up to 32,000 IOPS.
	//
	// `Iops` is supported when the volume type is `gp3` or `io1` and required only when the volume type is `io1` . (Not used with `standard` , `gp2` , `st1` , or `sc1` volumes.)
	Iops *float64 `json:"iops"`
	// The snapshot ID of the volume to use.
	//
	// You must specify either a `VolumeSize` or a `SnapshotId` .
	SnapshotId *string `json:"snapshotId"`
	// The throughput (MiBps) to provision for a `gp3` volume.
	Throughput *float64 `json:"throughput"`
	// The volume size, in GiBs. The following are the supported volumes sizes for each volume type:.
	//
	// - `gp2` and `gp3` : 1-16,384
	// - `io1` : 4-16,384
	// - `st1` and `sc1` : 125-16,384
	// - `standard` : 1-1,024
	//
	// You must specify either a `SnapshotId` or a `VolumeSize` . If you specify both `SnapshotId` and `VolumeSize` , the volume size must be equal or greater than the size of the snapshot.
	VolumeSize *float64 `json:"volumeSize"`
	// The volume type.
	//
	// For more information, see [Amazon EBS Volume Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Valid Values: `standard` | `io1` | `gp2` | `st1` | `sc1` | `gp3`
	VolumeType *string `json:"volumeType"`
}

// `MetadataOptions` is a property of [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html) that describes metadata options for the instances.
//
// For more information, see [Configuring the instance metadata options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnLaunchConfiguration_MetadataOptionsProperty struct {
	// This parameter enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is `enabled` .
	//
	// > If you specify a value of `disabled` , you will not be able to access your instance metadata.
	HttpEndpoint *string `json:"httpEndpoint"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//
	// Default: 1
	HttpPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit"`
	// The state of token usage for your instance metadata requests.
	//
	// If the parameter is not specified in the request, the default state is `optional` .
	//
	// If the state is `optional` , you can choose to retrieve instance metadata with or without a signed token header on your request. If you retrieve the IAM role credentials without a token, the version 1.0 role credentials are returned. If you retrieve the IAM role credentials using a valid signed token, the version 2.0 role credentials are returned.
	//
	// If the state is `required` , you must send a signed token header with any instance metadata retrieval requests. In this state, retrieving the IAM role credentials always returns the version 2.0 credentials; the version 1.0 credentials are not available.
	HttpTokens *string `json:"httpTokens"`
}

// Properties for defining a `CfnLaunchConfiguration`.
//
// TODO: EXAMPLE
//
type CfnLaunchConfigurationProps struct {
	// Provides the unique ID of the Amazon Machine Image (AMI) that was assigned during registration.
	//
	// For more information, see [Finding a Linux AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the *Amazon EC2 User Guide for Linux Instances* .
	ImageId *string `json:"imageId"`
	// Specifies the instance type of the EC2 instance.
	//
	// For information about available instance types, see [Available instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the *Amazon EC2 User Guide for Linux Instances* .
	InstanceType *string `json:"instanceType"`
	// For Auto Scaling groups that are running in a virtual private cloud (VPC), specifies whether to assign a public IP address to the group's instances.
	//
	// If you specify `true` , each instance in the Auto Scaling group receives a unique public IP address. For more information, see [Launching Auto Scaling instances in a VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If an instance receives a public IP address and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html) .
	//
	// > If the instance is launched into a default subnet, the default is to assign a public IP address, unless you disabled the option to assign a public IP address on the subnet. If the instance is launched into a nondefault subnet, the default is not to assign a public IP address, unless you enabled the option to assign a public IP address on the subnet.
	AssociatePublicIpAddress interface{} `json:"associatePublicIpAddress"`
	// Specifies how block devices are exposed to the instance.
	//
	// You can specify virtual devices and EBS volumes.
	BlockDeviceMappings interface{} `json:"blockDeviceMappings"`
	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to.
	//
	// For more information, see [ClassicLink](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in the *Amazon EC2 User Guide for Linux Instances* and [Linking EC2-Classic instances to a VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// This property can only be used if you are launching EC2-Classic instances.
	ClassicLinkVpcId *string `json:"classicLinkVpcId"`
	// The IDs of one or more security groups for the VPC that you specified in the `ClassicLinkVPCId` property.
	//
	// If you specify the `ClassicLinkVPCId` property, you must specify this property.
	ClassicLinkVpcSecurityGroups *[]*string `json:"classicLinkVpcSecurityGroups"`
	// Specifies whether the launch configuration is optimized for EBS I/O ( `true` ) or not ( `false` ).
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal EBS I/O performance. Additional fees are incurred when you enable EBS optimization for an instance type that is not EBS-optimized by default. For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// The default value is `false` .
	EbsOptimized interface{} `json:"ebsOptimized"`
	// Provides the name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.
	//
	// The instance profile contains the IAM role.
	//
	// For more information, see [IAM role for applications that run on Amazon EC2 instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	IamInstanceProfile *string `json:"iamInstanceProfile"`
	// The ID of the Amazon EC2 instance you want to use to create the launch configuration.
	//
	// Use this property if you want the launch configuration to use settings from an existing Amazon EC2 instance. When you use an instance to create a launch configuration, all properties are derived from the instance with the exception of `BlockDeviceMapping` and `AssociatePublicIpAddress` . You can override any properties from the instance by specifying them in the launch configuration.
	InstanceId *string `json:"instanceId"`
	// Controls whether instances in this group are launched with detailed ( `true` ) or basic ( `false` ) monitoring.
	//
	// The default value is `true` (enabled).
	//
	// > When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes. For more information, see [Configure monitoring for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics) in the *Amazon EC2 Auto Scaling User Guide* .
	InstanceMonitoring interface{} `json:"instanceMonitoring"`
	// Provides the ID of the kernel associated with the EC2 AMI.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	KernelId *string `json:"kernelId"`
	// Provides the name of the EC2 key pair.
	//
	// > If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in. For information on creating a key pair, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon EC2 User Guide for Linux Instances* .
	KeyName *string `json:"keyName"`
	// The name of the launch configuration.
	//
	// This name must be unique per Region per account.
	LaunchConfigurationName *string `json:"launchConfigurationName"`
	// The metadata options for the instances.
	//
	// For more information, see [Configuring the Instance Metadata Options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the *Amazon EC2 Auto Scaling User Guide* .
	MetadataOptions interface{} `json:"metadataOptions"`
	// The tenancy of the instance, either `default` or `dedicated` .
	//
	// An instance with `dedicated` tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC. You must set the value of this property to `dedicated` if want to launch dedicated instances in a shared tenancy VPC (a VPC with the instance placement tenancy attribute set to default).
	//
	// If you specify this property, you must specify at least one subnet in the `VPCZoneIdentifier` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource.
	//
	// For more information, see [Configuring instance tenancy with Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	PlacementTenancy *string `json:"placementTenancy"`
	// The ID of the RAM disk to select.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	RamDiskId *string `json:"ramDiskId"`
	// A list that contains the security groups to assign to the instances in the Auto Scaling group.
	//
	// The list can contain both the IDs of existing security groups and references to [SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
	//
	// For more information, see [Security groups for your VPC](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	SecurityGroups *[]*string `json:"securityGroups"`
	// The maximum hourly price you are willing to pay for any Spot Instances launched to fulfill the request.
	//
	// Spot Instances are launched when the price you specify exceeds the current Spot price. For more information, see [Requesting Spot Instances for fault-tolerant and flexible applications](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configuration-requesting-spot-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// > When you change your maximum price by creating a new launch configuration, running instances will continue to run as long as the maximum price for those running instances is higher than the current Spot price.
	//
	// Valid Range: Minimum value of 0.001
	SpotPrice *string `json:"spotPrice"`
	// The Base64-encoded user data to make available to the launched EC2 instances.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon EC2 User Guide for Linux Instances* .
	UserData *string `json:"userData"`
}

// A CloudFormation `AWS::AutoScaling::LifecycleHook`.
//
// The `AWS::AutoScaling::LifecycleHook` resource specifies lifecycle hooks for an Auto Scaling group. These hooks enable an Auto Scaling group to be aware of events in the Auto Scaling instance lifecycle, and then perform a custom action when the corresponding lifecycle event occurs. A lifecycle hook provides a specified amount of time (one hour by default) to complete the lifecycle action before the instance transitions to the next state.
//
// There are two types of lifecycle hooks that can be implemented: launch lifecycle hooks and termination lifecycle hooks. Use a launch lifecycle hook to prepare instances for use or to delay instances from registering behind the load balancer before their configuration has been applied completely. Use a termination lifecycle hook to prepare running instances to be shut down.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* and [PutLifecycleHook](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PutLifecycleHook.html) in the *Amazon EC2 Auto Scaling API Reference* .
//
// TODO: EXAMPLE
//
type CfnLifecycleHook interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultResult() *string
	SetDefaultResult(val *string)
	HeartbeatTimeout() *float64
	SetHeartbeatTimeout(val *float64)
	LifecycleHookName() *string
	SetLifecycleHookName(val *string)
	LifecycleTransition() *string
	SetLifecycleTransition(val *string)
	LogicalId() *string
	Node() constructs.Node
	NotificationMetadata() *string
	SetNotificationMetadata(val *string)
	NotificationTargetArn() *string
	SetNotificationTargetArn(val *string)
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLifecycleHook
type jsiiProxy_CfnLifecycleHook struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLifecycleHook) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) DefaultResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) HeartbeatTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) LifecycleHookName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleHookName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) LifecycleTransition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) NotificationMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) NotificationTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHook) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::LifecycleHook`.
func NewCfnLifecycleHook(scope constructs.Construct, id *string, props *CfnLifecycleHookProps) CfnLifecycleHook {
	_init_.Initialize()

	j := jsiiProxy_CfnLifecycleHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::LifecycleHook`.
func NewCfnLifecycleHook_Override(c CfnLifecycleHook, scope constructs.Construct, id *string, props *CfnLifecycleHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetDefaultResult(val *string) {
	_jsii_.Set(
		j,
		"defaultResult",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetHeartbeatTimeout(val *float64) {
	_jsii_.Set(
		j,
		"heartbeatTimeout",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetLifecycleHookName(val *string) {
	_jsii_.Set(
		j,
		"lifecycleHookName",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetLifecycleTransition(val *string) {
	_jsii_.Set(
		j,
		"lifecycleTransition",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetNotificationMetadata(val *string) {
	_jsii_.Set(
		j,
		"notificationMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetNotificationTargetArn(val *string) {
	_jsii_.Set(
		j,
		"notificationTargetArn",
		val,
	)
}

func (j *jsiiProxy_CfnLifecycleHook) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLifecycleHook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLifecycleHook_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnLifecycleHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecycleHook_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnLifecycleHook",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLifecycleHook) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnLifecycleHook) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLifecycleHook) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnLifecycleHook) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLifecycleHook) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLifecycleHook) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnLifecycleHook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnLifecycleHook) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLifecycleHook) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnLifecycleHook) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLifecycleHook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLifecycleHook) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnLifecycleHook) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnLifecycleHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecycleHook) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLifecycleHook`.
//
// TODO: EXAMPLE
//
type CfnLifecycleHookProps struct {
	// The name of the Auto Scaling group for the lifecycle hook.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// The instance state to which you want to attach the lifecycle hook. The valid values are:.
	//
	// - autoscaling:EC2_INSTANCE_LAUNCHING
	// - autoscaling:EC2_INSTANCE_TERMINATING
	LifecycleTransition *string `json:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	//
	// The valid values are `CONTINUE` and `ABANDON` (default).
	//
	// For more information, see [Adding lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/adding-lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* .
	DefaultResult *string `json:"defaultResult"`
	// The maximum time, in seconds, that can elapse before the lifecycle hook times out.
	//
	// The range is from `30` to `7200` seconds. The default value is `3600` seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the `DefaultResult` property.
	HeartbeatTimeout *float64 `json:"heartbeatTimeout"`
	// The name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName"`
	// Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata"`
	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook.
	//
	// You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
	NotificationTargetArn *string `json:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
	//
	// For information about creating this role, see [Configuring a notification target for a lifecycle hook](https://docs.aws.amazon.com/autoscaling/ec2/userguide/configuring-lifecycle-hook-notifications.html) in the *Amazon EC2 Auto Scaling User Guide* .
	RoleArn *string `json:"roleArn"`
}

// A CloudFormation `AWS::AutoScaling::ScalingPolicy`.
//
// The `AWS::AutoScaling::ScalingPolicy` resource specifies an Amazon EC2 Auto Scaling scaling policy so that the Auto Scaling group can scale the number of instances available for your application.
//
// For more information about using scaling policies to scale your Auto Scaling group automatically, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) and [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AdjustmentType() *string
	SetAdjustmentType(val *string)
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cooldown() *string
	SetCooldown(val *string)
	CreationStack() *[]*string
	EstimatedInstanceWarmup() *float64
	SetEstimatedInstanceWarmup(val *float64)
	LogicalId() *string
	MetricAggregationType() *string
	SetMetricAggregationType(val *string)
	MinAdjustmentMagnitude() *float64
	SetMinAdjustmentMagnitude(val *float64)
	Node() constructs.Node
	PolicyType() *string
	SetPolicyType(val *string)
	PredictiveScalingConfiguration() interface{}
	SetPredictiveScalingConfiguration(val interface{})
	Ref() *string
	ScalingAdjustment() *float64
	SetScalingAdjustment(val *float64)
	Stack() awscdk.Stack
	StepAdjustments() interface{}
	SetStepAdjustments(val interface{})
	TargetTrackingConfiguration() interface{}
	SetTargetTrackingConfiguration(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalingPolicy
type jsiiProxy_CfnScalingPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalingPolicy) AdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) EstimatedInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) MetricAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) MinAdjustmentMagnitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) PredictiveScalingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predictiveScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) ScalingAdjustment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) StepAdjustments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) TargetTrackingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::ScalingPolicy`.
func NewCfnScalingPolicy(scope constructs.Construct, id *string, props *CfnScalingPolicyProps) CfnScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::ScalingPolicy`.
func NewCfnScalingPolicy_Override(c CfnScalingPolicy, scope constructs.Construct, id *string, props *CfnScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetAdjustmentType(val *string) {
	_jsii_.Set(
		j,
		"adjustmentType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetCooldown(val *string) {
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetEstimatedInstanceWarmup(val *float64) {
	_jsii_.Set(
		j,
		"estimatedInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetMetricAggregationType(val *string) {
	_jsii_.Set(
		j,
		"metricAggregationType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetMinAdjustmentMagnitude(val *float64) {
	_jsii_.Set(
		j,
		"minAdjustmentMagnitude",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetPolicyType(val *string) {
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetPredictiveScalingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"predictiveScalingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetScalingAdjustment(val *float64) {
	_jsii_.Set(
		j,
		"scalingAdjustment",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetStepAdjustments(val interface{}) {
	_jsii_.Set(
		j,
		"stepAdjustments",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetTargetTrackingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"targetTrackingConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScalingPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScalingPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnScalingPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnScalingPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScalingPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnScalingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnScalingPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnScalingPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnScalingPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnScalingPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScalingPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnScalingPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnScalingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnScalingPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains customized metric specification information for a target tracking scaling policy for Amazon EC2 Auto Scaling.
//
// To create your customized metric specification:
//
// - Add values for each required property from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see [Publish Custom Metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the *Amazon CloudWatch User Guide* .
// - Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases.
//
// For more information about CloudWatch, see [Amazon CloudWatch Concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) .
//
// `CustomizedMetricSpecification` is a property of the [AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html) property type.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_CustomizedMetricSpecificationProperty struct {
	// The name of the metric.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html) .
	MetricName *string `json:"metricName"`
	// The namespace of the metric.
	Namespace *string `json:"namespace"`
	// The statistic of the metric.
	Statistic *string `json:"statistic"`
	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.
	Dimensions interface{} `json:"dimensions"`
	// The unit of the metric.
	//
	// For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference* .
	Unit *string `json:"unit"`
}

// `MetricDimension` specifies a name/value pair that is part of the identity of a CloudWatch metric for the `Dimensions` property of the [AWS::AutoScaling::ScalingPolicy CustomizedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html) property type. Duplicate dimensions are not allowed.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_MetricDimensionProperty struct {
	// The name of the dimension.
	Name *string `json:"name"`
	// The value of the dimension.
	Value *string `json:"value"`
}

// Contains predefined metric specification information for a target tracking scaling policy for Amazon EC2 Auto Scaling.
//
// `PredefinedMetricSpecification` is a property of the [AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html) property type.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredefinedMetricSpecificationProperty struct {
	// The metric type. The following predefined metrics are available.
	//
	// - `ASGAverageCPUUtilization` - Average CPU utilization of the Auto Scaling group.
	// - `ASGAverageNetworkIn` - Average number of bytes received on all network interfaces by the Auto Scaling group.
	// - `ASGAverageNetworkOut` - Average number of bytes sent out on all network interfaces by the Auto Scaling group.
	// - `ALBRequestCountPerTarget` - Number of requests completed per target in an Application Load Balancer target group.
	PredefinedMetricType *string `json:"predefinedMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is `ALBRequestCountPerTarget` and there is a target group attached to the Auto Scaling group.
	//
	// The format is `app/ *load-balancer-name* / *load-balancer-id* /targetgroup/ *target-group-name* / *target-group-id*` , where
	//
	// - `app/ *load-balancer-name* / *load-balancer-id*` is the final portion of the load balancer ARN, and
	// - `targetgroup/ *target-group-name* / *target-group-id*` is the final portion of the target group ARN.
	ResourceLabel *string `json:"resourceLabel"`
}

// `PredictiveScalingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html) resource that specifies a predictive scaling policy for Amazon EC2 Auto Scaling.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PutScalingPolicy.html) in the *Amazon EC2 Auto Scaling API Reference* and [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredictiveScalingConfigurationProperty struct {
	// An array that contains information about the metrics and target utilization to use for predictive scaling.
	//
	// > Adding more than one predictive scaling metric specification to the array is currently not supported.
	MetricSpecifications interface{} `json:"metricSpecifications"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity of the Auto Scaling group.
	//
	// Defaults to `HonorMaxCapacity` if not specified.
	//
	// The following are possible values:
	//
	// - `HonorMaxCapacity` - Amazon EC2 Auto Scaling cannot scale out capacity higher than the maximum capacity. The maximum capacity is enforced as a hard limit.
	// - `IncreaseMaxCapacity` - Amazon EC2 Auto Scaling can scale out capacity higher than the maximum capacity when the forecast capacity is close to or exceeds the maximum capacity. The upper limit is determined by the forecasted capacity and the value for `MaxCapacityBuffer` .
	MaxCapacityBreachBehavior *string `json:"maxCapacityBreachBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer, such that if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// If set to 0, Amazon EC2 Auto Scaling may scale capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	//
	// Required if the `MaxCapacityBreachBehavior` property is set to `IncreaseMaxCapacity` , and cannot be used otherwise.
	MaxCapacityBuffer *float64 `json:"maxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// Defaults to `ForecastOnly` if not specified.
	Mode *string `json:"mode"`
	// The amount of time, in seconds, by which the instance launch time can be advanced.
	//
	// For example, the forecast says to add capacity at 10:00 AM, and you choose to pre-launch instances by 5 minutes. In that case, the instances will be launched at 9:55 AM. The intention is to give resources time to be provisioned. It can take a few minutes to launch an EC2 instance. The actual amount of time required depends on several factors, such as the size of the instance and whether there are startup scripts to complete.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). Defaults to 300 seconds if not specified.
	SchedulingBufferTime *float64 `json:"schedulingBufferTime"`
}

// A structure that specifies a metric specification for the `MetricSpecifications` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html) property type.
//
// For more information, see [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredictiveScalingMetricSpecificationProperty struct {
	// Specifies the target utilization.
	//
	// > Some metrics are based on a count instead of a percentage, such as the request count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy specifies one of these metrics, specify the target utilization as the optimal average request or message count per instance during any one-minute interval.
	TargetValue *float64 `json:"targetValue"`
	// The load metric specification.
	//
	// If you specify `PredefinedMetricPairSpecification` , don't specify this property.
	PredefinedLoadMetricSpecification interface{} `json:"predefinedLoadMetricSpecification"`
	// The metric pair specification from which Amazon EC2 Auto Scaling determines the appropriate scaling metric and load metric to use.
	//
	// > With predictive scaling, you must specify either a metric pair, or a load metric and a scaling metric individually. Specifying a metric pair instead of individual metrics provides a simpler way to configure metrics for a scaling policy. You choose the metric pair, and the policy automatically knows the correct sum and average statistics to use for the load metric and the scaling metric.
	PredefinedMetricPairSpecification interface{} `json:"predefinedMetricPairSpecification"`
	// The scaling metric specification.
	//
	// If you specify `PredefinedMetricPairSpecification` , don't specify this property.
	PredefinedScalingMetricSpecification interface{} `json:"predefinedScalingMetricSpecification"`
}

// Contains load metric information for the `PredefinedLoadMetricSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html) property type.
//
// > Does not apply to policies that use a *metric pair* for the metric specification.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredictiveScalingPredefinedLoadMetricProperty struct {
	// The metric type.
	PredefinedMetricType *string `json:"predefinedMetricType"`
	// A label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
	//
	// You can't specify a resource label unless the target group is attached to the Auto Scaling group.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//
	// `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` .
	//
	// Where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel"`
}

// Contains metric pair information for the `PredefinedMetricPairSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html) property type.
//
// For more information, see [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredictiveScalingPredefinedMetricPairProperty struct {
	// Indicates which metrics to use.
	//
	// There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric. For example, if the metric type is `ASGCPUUtilization` , the Auto Scaling group's total CPU metric is used as the load metric, and the average CPU metric is used for the scaling metric.
	PredefinedMetricType *string `json:"predefinedMetricType"`
	// A label that uniquely identifies a specific Application Load Balancer target group from which to determine the total and average request count served by your Auto Scaling group.
	//
	// You can't specify a resource label unless the target group is attached to the Auto Scaling group.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//
	// `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` .
	//
	// Where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel"`
}

// Contains scaling metric information for the `PredefinedScalingMetricSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.html) property type.
//
// > Does not apply to policies that use a *metric pair* for the metric specification.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_PredictiveScalingPredefinedScalingMetricProperty struct {
	// The metric type.
	PredefinedMetricType *string `json:"predefinedMetricType"`
	// A label that uniquely identifies a specific Application Load Balancer target group from which to determine the average request count served by your Auto Scaling group.
	//
	// You can't specify a resource label unless the target group is attached to the Auto Scaling group.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//
	// `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` .
	//
	// Where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel"`
}

// `StepAdjustment` specifies a step adjustment for the `StepAdjustments` property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html) resource.
//
// For the following examples, suppose that you have an alarm with a breach threshold of 50:
//
// - To trigger a step adjustment when the metric is greater than or equal to 50 and less than 60, specify a lower bound of 0 and an upper bound of 10.
// - To trigger a step adjustment when the metric is greater than 40 and less than or equal to 50, specify a lower bound of -10 and an upper bound of 0.
//
// There are a few rules for the step adjustments for your step policy:
//
// - The ranges of your step adjustments can't overlap or have a gap.
// - At most one step adjustment can have a null lower bound. If one step adjustment has a negative lower bound, then there must be a step adjustment with a null lower bound.
// - At most one step adjustment can have a null upper bound. If one step adjustment has a positive upper bound, then there must be a step adjustment with a null upper bound.
// - The upper and lower bound can't be null in the same step adjustment.
//
// For more information, see [Step adjustments](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-steps) in the *Amazon EC2 Auto Scaling User Guide* .
//
// You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#aws-properties-as-policy--examples) section of the `AWS::AutoScaling::ScalingPolicy` documentation.
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_StepAdjustmentProperty struct {
	// The amount by which to scale.
	//
	// The adjustment is based on the value that you specified in the `AdjustmentType` property (either an absolute number or a percentage). A positive value adds to the current capacity and a negative number subtracts from the current capacity.
	ScalingAdjustment *float64 `json:"scalingAdjustment"`
	// The lower bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.
	MetricIntervalLowerBound *float64 `json:"metricIntervalLowerBound"`
	// The upper bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.
	MetricIntervalUpperBound *float64 `json:"metricIntervalUpperBound"`
}

// `TargetTrackingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html) resource that specifies a target tracking scaling policy configuration for Amazon EC2 Auto Scaling.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PutScalingPolicy.html) in the *Amazon EC2 Auto Scaling API Reference* . For more information about scaling policies, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPolicy_TargetTrackingConfigurationProperty struct {
	// The target value for the metric.
	TargetValue *float64 `json:"targetValue"`
	// A customized metric.
	//
	// You must specify either a predefined metric or a customized metric.
	CustomizedMetricSpecification interface{} `json:"customizedMetricSpecification"`
	// Indicates whether scaling in by the target tracking scaling policy is disabled.
	//
	// If scaling in is disabled, the target tracking scaling policy doesn't remove instances from the Auto Scaling group. Otherwise, the target tracking scaling policy can remove instances from the Auto Scaling group. The default is `false` .
	DisableScaleIn interface{} `json:"disableScaleIn"`
	// A predefined metric.
	//
	// You must specify either a predefined metric or a customized metric.
	PredefinedMetricSpecification interface{} `json:"predefinedMetricSpecification"`
}

// Properties for defining a `CfnScalingPolicy`.
//
// TODO: EXAMPLE
//
type CfnScalingPolicyProps struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// Specifies how the scaling adjustment is interpreted. The valid values are `ChangeInCapacity` , `ExactCapacity` , and `PercentChangeInCapacity` .
	//
	// Required if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	AdjustmentType *string `json:"adjustmentType"`
	// The duration of the policy's cooldown period, in seconds.
	//
	// When a cooldown period is specified here, it overrides the default cooldown period defined for the Auto Scaling group.
	//
	// Valid only if the policy type is `SimpleScaling` . For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	Cooldown *string `json:"cooldown"`
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	//
	// If not provided, the default is to use the value from the default cooldown period for the Auto Scaling group.
	//
	// Valid only if the policy type is `TargetTrackingScaling` or `StepScaling` .
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	//
	// The valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	//
	// Valid only if the policy type is `StepScaling` .
	MetricAggregationType *string `json:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances.
	//
	// Valid only if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// > Some Auto Scaling groups use instance weights. In this case, set the `MinAdjustmentMagnitude` to a value that is at least as large as your largest instance weight.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude"`
	// One of the following policy types:.
	//
	// - `TargetTrackingScaling`
	// - `StepScaling`
	// - `SimpleScaling` (default)
	// - `PredictiveScaling`
	//
	// For more information, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html) and [Step and simple scaling policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html) in the *Amazon EC2 Auto Scaling User Guide* .
	PolicyType *string `json:"policyType"`
	// A predictive scaling policy.
	//
	// Includes support for predefined metrics only.
	PredictiveScalingConfiguration interface{} `json:"predictiveScalingConfiguration"`
	// The amount by which to scale, based on the specified adjustment type.
	//
	// A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value.
	//
	// Required if the policy type is `SimpleScaling` . (Not used with any other policy type.)
	ScalingAdjustment *float64 `json:"scalingAdjustment"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// Required if the policy type is `StepScaling` . (Not used with any other policy type.)
	StepAdjustments interface{} `json:"stepAdjustments"`
	// A target tracking scaling policy. Includes support for predefined or customized metrics.
	//
	// The following predefined metrics are available:
	//
	// - `ASGAverageCPUUtilization`
	// - `ASGAverageNetworkIn`
	// - `ASGAverageNetworkOut`
	// - `ALBRequestCountPerTarget`
	//
	// If you specify `ALBRequestCountPerTarget` for the metric, you must specify the `ResourceLabel` property with the `PredefinedMetricSpecification` .
	TargetTrackingConfiguration interface{} `json:"targetTrackingConfiguration"`
}

// A CloudFormation `AWS::AutoScaling::ScheduledAction`.
//
// The `AWS::AutoScaling::ScheduledAction` resource specifies an Amazon EC2 Auto Scaling scheduled action so that the Auto Scaling group can change the number of instances available for your application in response to predictable load changes.
//
// When you update a stack with an Auto Scaling group and scheduled action, CloudFormation always sets the min size, max size, and desired capacity properties of your group to the values that are defined in the `AWS::AutoScaling::AutoScalingGroup` section of your template. However, you might not want CloudFormation to do that when you have a scheduled action in effect. You can use an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) to prevent CloudFormation from changing the min size, max size, or desired capacity property values during a stack update unless you modified the individual values in your template. If you have rolling updates enabled, before you can update the Auto Scaling group, you must suspend scheduled actions by specifying an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) for the Auto Scaling group. You can find a sample update policy for rolling updates in [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) .
//
// For more information, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html) and [Suspending and resuming scaling processes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScheduledAction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	EndTime() *string
	SetEndTime(val *string)
	LogicalId() *string
	MaxSize() *float64
	SetMaxSize(val *float64)
	MinSize() *float64
	SetMinSize(val *float64)
	Node() constructs.Node
	Recurrence() *string
	SetRecurrence(val *string)
	Ref() *string
	Stack() awscdk.Stack
	StartTime() *string
	SetStartTime(val *string)
	TimeZone() *string
	SetTimeZone(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScheduledAction
type jsiiProxy_CfnScheduledAction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScheduledAction) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Recurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::ScheduledAction`.
func NewCfnScheduledAction(scope constructs.Construct, id *string, props *CfnScheduledActionProps) CfnScheduledAction {
	_init_.Initialize()

	j := jsiiProxy_CfnScheduledAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::ScheduledAction`.
func NewCfnScheduledAction_Override(c CfnScheduledAction, scope constructs.Construct, id *string, props *CfnScheduledActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetDesiredCapacity(val *float64) {
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetRecurrence(val *string) {
	_jsii_.Set(
		j,
		"recurrence",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScheduledAction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScheduledAction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnScheduledAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledAction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnScheduledAction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnScheduledAction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnScheduledAction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScheduledAction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnScheduledAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnScheduledAction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnScheduledAction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnScheduledAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnScheduledAction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScheduledAction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnScheduledAction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnScheduledAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScheduledAction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnScheduledAction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnScheduledAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnScheduledAction`.
//
// TODO: EXAMPLE
//
type CfnScheduledActionProps struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
	//
	// It can scale beyond this capacity if you add more scaling conditions.
	//
	// You must specify at least one of the following properties: `MaxSize` , `MinSize` , or `DesiredCapacity` .
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// The date and time for the recurring schedule to end, in UTC.
	//
	// For example, `"2021-06-01T00:00:00Z"` .
	EndTime *string `json:"endTime"`
	// The maximum size of the Auto Scaling group.
	//
	// You must specify at least one of the following properties: `MaxSize` , `MinSize` , or `DesiredCapacity` .
	MaxSize *float64 `json:"maxSize"`
	// The minimum size of the Auto Scaling group.
	//
	// You must specify at least one of the following properties: `MaxSize` , `MinSize` , or `DesiredCapacity` .
	MinSize *float64 `json:"minSize"`
	// The recurring schedule for this action.
	//
	// This format consists of five fields separated by white spaces: [Minute] [Hour] [Day_of_Month] [Month_of_Year] [Day_of_Week]. For more information about this format, see [Crontab](https://docs.aws.amazon.com/http://crontab.org) .
	//
	// When `StartTime` and `EndTime` are specified with `Recurrence` , they form the boundaries of when the recurring action starts and stops.
	//
	// Cron expressions use Universal Coordinated Time (UTC) by default.
	Recurrence *string `json:"recurrence"`
	// The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in UTC/GMT only. For example, `"2021-06-01T00:00:00Z"` .
	//
	// If you specify `Recurrence` and `StartTime` , Amazon EC2 Auto Scaling performs the action at this time, and then performs the action based on the specified recurrence.
	StartTime *string `json:"startTime"`
	// Specifies the time zone for a cron expression.
	//
	// If a time zone is not provided, UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the IANA Time Zone Database (such as `Etc/GMT+9` or `Pacific/Tahiti` ). For more information, see [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) .
	TimeZone *string `json:"timeZone"`
}

// A CloudFormation `AWS::AutoScaling::WarmPool`.
//
// The `AWS::AutoScaling::WarmPool` resource creates a pool of pre-initialized EC2 instances that sits alongside the Auto Scaling group. Whenever your application needs to scale out, the Auto Scaling group can draw on the warm pool to meet its new desired capacity.
//
// When you create a warm pool, you can define a minimum size. When your Auto Scaling group scales out and the size of the warm pool shrinks, Amazon EC2 Auto Scaling launches new instances into the warm pool to maintain its minimum size.
//
// For more information, see [Warm pools for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// > CloudFormation supports the `UpdatePolicy` attribute for Auto Scaling groups. During an update, if `UpdatePolicy` is set to `AutoScalingRollingUpdate` , CloudFormation replaces `InService` instances only. Instances in the warm pool are not replaced. The difference in which instances are replaced can potentially result in different instance configurations after the stack update completes. If `UpdatePolicy` is set to `AutoScalingReplacingUpdate` , you do not encounter this issue because CloudFormation replaces both the Auto Scaling group and the warm pool.
//
// TODO: EXAMPLE
//
type CfnWarmPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MaxGroupPreparedCapacity() *float64
	SetMaxGroupPreparedCapacity(val *float64)
	MinSize() *float64
	SetMinSize(val *float64)
	Node() constructs.Node
	PoolState() *string
	SetPoolState(val *string)
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWarmPool
type jsiiProxy_CfnWarmPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWarmPool) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) MaxGroupPreparedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) PoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWarmPool) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::WarmPool`.
func NewCfnWarmPool(scope constructs.Construct, id *string, props *CfnWarmPoolProps) CfnWarmPool {
	_init_.Initialize()

	j := jsiiProxy_CfnWarmPool{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::WarmPool`.
func NewCfnWarmPool_Override(c CfnWarmPool, scope constructs.Construct, id *string, props *CfnWarmPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWarmPool) SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnWarmPool) SetMaxGroupPreparedCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxGroupPreparedCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnWarmPool) SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnWarmPool) SetPoolState(val *string) {
	_jsii_.Set(
		j,
		"poolState",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWarmPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWarmPool_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnWarmPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWarmPool_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnWarmPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnWarmPool) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnWarmPool) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnWarmPool) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnWarmPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnWarmPool) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnWarmPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnWarmPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnWarmPool) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnWarmPool) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnWarmPool) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnWarmPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWarmPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnWarmPool) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnWarmPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWarmPool) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWarmPool`.
//
// TODO: EXAMPLE
//
type CfnWarmPoolProps struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// Specifies the maximum number of instances that are allowed to be in the warm pool or in any state except `Terminated` for the Auto Scaling group.
	//
	// This is an optional property. Specify it only if you do not want the warm pool size to be determined by the difference between the group's maximum capacity and its desired capacity.
	//
	// > If a value for `MaxGroupPreparedCapacity` is not specified, Amazon EC2 Auto Scaling launches and maintains the difference between the group's maximum capacity and its desired capacity. If you specify a value for `MaxGroupPreparedCapacity` , Amazon EC2 Auto Scaling uses the difference between the `MaxGroupPreparedCapacity` and the desired capacity instead.
	// >
	// > The size of the warm pool is dynamic. Only when `MaxGroupPreparedCapacity` and `MinSize` are set to the same value does the warm pool have an absolute size.
	//
	// If the desired capacity of the Auto Scaling group is higher than the `MaxGroupPreparedCapacity` , the capacity of the warm pool is 0, unless you specify a value for `MinSize` . To remove a value that you previously set, include the property but specify -1 for the value.
	MaxGroupPreparedCapacity *float64 `json:"maxGroupPreparedCapacity"`
	// Specifies the minimum number of instances to maintain in the warm pool.
	//
	// This helps you to ensure that there is always a certain number of warmed instances available to handle traffic spikes. Defaults to 0 if not specified.
	MinSize *float64 `json:"minSize"`
	// Sets the instance state to transition to after the lifecycle actions are complete.
	//
	// Default is `Stopped` .
	PoolState *string `json:"poolState"`
}

// Basic properties of an AutoScalingGroup, except the exact machines to run and where they should run.
//
// Constructs that want to create AutoScalingGroups can inherit
// this interface and specialize the essential parts in various ways.
//
// TODO: EXAMPLE
//
type CommonAutoScalingGroupProps struct {
	// Whether the instances can initiate connections to anywhere by default.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	AssociatePublicIpAddress *bool `json:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `json:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*BlockDevice `json:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	GroupMetrics *[]GroupMetrics `json:"groupMetrics"`
	// Configuration for health checks.
	HealthCheck HealthCheck `json:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	IgnoreUnmodifiedSizeProperties *bool `json:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	InstanceMonitoring Monitoring `json:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	KeyName *string `json:"keyName"`
	// Maximum number of instances in the fleet.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The maximum amount of time that an instance can be in service.
	//
	// The maximum duration applies
	// to all current and future instances in the group. As an instance approaches its maximum duration,
	// it is terminated and replaced, and cannot be used again.
	//
	// You must specify a value of at least 604,800 seconds (7 days). To clear a previously set value,
	// leave this property undefined.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html
	//
	MaxInstanceLifetime awscdk.Duration `json:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
	MinCapacity *float64 `json:"minCapacity"`
	// Whether newly-launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// By default, Auto Scaling can terminate an instance at any time after launch
	// when scaling in an Auto Scaling Group, subject to the group's termination
	// policy. However, you may wish to protect newly-launched instances from
	// being scaled in if they are going to run critical applications that should
	// not be prematurely terminated.
	//
	// This flag must be enabled if the Auto Scaling Group will be associated with
	// an ECS Capacity Provider with managed termination protection.
	NewInstancesProtectedFromScaleIn *bool `json:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	Notifications *[]*NotificationConfiguration `json:"notifications"`
	// Configure waiting for signals during deployment.
	//
	// Use this to pause the CloudFormation deployment to wait for the instances
	// in the AutoScalingGroup to report successful startup during
	// creation and updates. The UserData script needs to invoke `cfn-signal`
	// with a success or failure code after it is done setting up the instance.
	//
	// Without waiting for signals, the CloudFormation deployment will proceed as
	// soon as the AutoScalingGroup has been created or updated but before the
	// instances in the group have been started.
	//
	// For example, to have instances wait for an Elastic Load Balancing health check before
	// they signal success, add a health-check verification by using the
	// cfn-init helper script. For an example, see the verify_instance_health
	// command in the Auto Scaling rolling updates sample template:
	//
	// https://github.com/awslabs/aws-cloudformation-templates/blob/master/aws/services/AutoScaling/AutoScalingRollingUpdates.yaml
	Signals Signals `json:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	SpotPrice *string `json:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	TerminationPolicies *[]TerminationPolicy `json:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	UpdatePolicy UpdatePolicy `json:"updatePolicy"`
	// Where to place instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Properties for enabling scaling based on CPU utilization.
//
// TODO: EXAMPLE
//
type CpuUtilizationScalingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// Target average CPU utilization across the task.
	TargetUtilizationPercent *float64 `json:"targetUtilizationPercent"`
}

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// TODO: EXAMPLE
//
// See: http://crontab.org/
//
type CronOptions struct {
	// The day of the month to run this rule at.
	Day *string `json:"day"`
	// The hour to run this rule at.
	Hour *string `json:"hour"`
	// The minute to run this rule at.
	Minute *string `json:"minute"`
	// The month to run this rule at.
	Month *string `json:"month"`
	// The day of the week to run this rule at.
	WeekDay *string `json:"weekDay"`
}

type DefaultResult string

const (
	DefaultResult_CONTINUE DefaultResult = "CONTINUE"
	DefaultResult_ABANDON DefaultResult = "ABANDON"
)

// Block device options for an EBS volume.
//
// TODO: EXAMPLE
//
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `json:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `json:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	Encrypted *bool `json:"encrypted"`
}

// Base block device options for an EBS volume.
//
// TODO: EXAMPLE
//
type EbsDeviceOptionsBase struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `json:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `json:"volumeType"`
}

// Properties of an EBS block device.
//
// TODO: EXAMPLE
//
type EbsDeviceProps struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `json:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `json:"volumeType"`
	// The volume size, in Gibibytes (GiB).
	//
	// If you specify volumeSize, it must be equal or greater than the size of the snapshot.
	VolumeSize *float64 `json:"volumeSize"`
	// The snapshot ID of the volume to use.
	SnapshotId *string `json:"snapshotId"`
}

// Block device options for an EBS volume created from a snapshot.
//
// TODO: EXAMPLE
//
type EbsDeviceSnapshotOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `json:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `json:"volumeType"`
	// The volume size, in Gibibytes (GiB).
	//
	// If you specify volumeSize, it must be equal or greater than the size of the snapshot.
	VolumeSize *float64 `json:"volumeSize"`
}

// Supported EBS volume types for blockDevices.
type EbsDeviceVolumeType string

const (
	EbsDeviceVolumeType_STANDARD EbsDeviceVolumeType = "STANDARD"
	EbsDeviceVolumeType_IO1 EbsDeviceVolumeType = "IO1"
	EbsDeviceVolumeType_GP2 EbsDeviceVolumeType = "GP2"
	EbsDeviceVolumeType_GP3 EbsDeviceVolumeType = "GP3"
	EbsDeviceVolumeType_ST1 EbsDeviceVolumeType = "ST1"
	EbsDeviceVolumeType_SC1 EbsDeviceVolumeType = "SC1"
)

// EC2 Heath check options.
//
// TODO: EXAMPLE
//
type Ec2HealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	Grace awscdk.Duration `json:"grace"`
}

// ELB Heath check options.
//
// TODO: EXAMPLE
//
type ElbHealthCheckOptions struct {
	// Specified the time Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	//
	// This option is required for ELB health checks.
	Grace awscdk.Duration `json:"grace"`
}

// Group metrics that an Auto Scaling group sends to Amazon CloudWatch.
//
// TODO: EXAMPLE
//
type GroupMetric interface {
	Name() *string
}

// The jsii proxy struct for GroupMetric
type jsiiProxy_GroupMetric struct {
	_ byte // padding
}

func (j *jsiiProxy_GroupMetric) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewGroupMetric(name *string) GroupMetric {
	_init_.Initialize()

	j := jsiiProxy_GroupMetric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewGroupMetric_Override(g GroupMetric, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		[]interface{}{name},
		g,
	)
}

func GroupMetric_DESIRED_CAPACITY() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"DESIRED_CAPACITY",
		&returns,
	)
	return returns
}

func GroupMetric_IN_SERVICE_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"IN_SERVICE_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_MAX_SIZE() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"MAX_SIZE",
		&returns,
	)
	return returns
}

func GroupMetric_MIN_SIZE() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"MIN_SIZE",
		&returns,
	)
	return returns
}

func GroupMetric_PENDING_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"PENDING_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_STANDBY_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"STANDBY_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_TERMINATING_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"TERMINATING_INSTANCES",
		&returns,
	)
	return returns
}

func GroupMetric_TOTAL_INSTANCES() GroupMetric {
	_init_.Initialize()
	var returns GroupMetric
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.GroupMetric",
		"TOTAL_INSTANCES",
		&returns,
	)
	return returns
}

// A set of group metrics.
//
// TODO: EXAMPLE
//
type GroupMetrics interface {
}

// The jsii proxy struct for GroupMetrics
type jsiiProxy_GroupMetrics struct {
	_ byte // padding
}

func NewGroupMetrics(metrics ...GroupMetric) GroupMetrics {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	j := jsiiProxy_GroupMetrics{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		args,
		&j,
	)

	return &j
}

func NewGroupMetrics_Override(g GroupMetrics, metrics ...GroupMetric) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range metrics {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		args,
		g,
	)
}

// Report all group metrics.
func GroupMetrics_All() GroupMetrics {
	_init_.Initialize()

	var returns GroupMetrics

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.GroupMetrics",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Health check settings.
//
// TODO: EXAMPLE
//
type HealthCheck interface {
	GracePeriod() awscdk.Duration
	Type() *string
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

func (j *jsiiProxy_HealthCheck) GracePeriod() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Use EC2 for health checks.
func HealthCheck_Ec2(options *Ec2HealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthCheck",
		"ec2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use ELB for health checks.
//
// It considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks.
func HealthCheck_Elb(options *ElbHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	var returns HealthCheck

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.HealthCheck",
		"elb",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// An AutoScalingGroup.
type IAutoScalingGroup interface {
	awsiam.IGrantable
	awscdk.IResource
	// Send a message to either an SQS queue or SNS topic when instances launch or terminate.
	AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook
	// Add command to the startup script of fleet instances.
	//
	// The command must be in the scripting language supported by the fleet's OS (i.e. Linux/Windows).
	// Does nothing for imported ASGs.
	AddUserData(commands ...*string)
	// Scale out or in to achieve a target CPU utilization.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target network ingress rate.
	ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in, in response to a metric.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in to achieve a target network egress rate.
	ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in based on time.
	ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction
	// Scale out or in in order to keep a metric around a target value.
	ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy
	// The arn of the AutoScalingGroup.
	AutoScalingGroupArn() *string
	// The name of the AutoScalingGroup.
	AutoScalingGroupName() *string
	// The operating system family that the instances in this auto-scaling group belong to.
	//
	// Is 'UNKNOWN' for imported ASGs.
	OsType() awsec2.OperatingSystemType
}

// The jsii proxy for IAutoScalingGroup
type jsiiProxy_IAutoScalingGroup struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAutoScalingGroup) AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook {
	var returns LifecycleHook

	_jsii_.Invoke(
		i,
		"addLifecycleHook",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) AddUserData(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addUserData",
		args,
	)
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnIncomingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	var returns StepScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnOutgoingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction {
	var returns ScheduledAction

	_jsii_.Invoke(
		i,
		"scaleOnSchedule",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAutoScalingGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) OsType() awsec2.OperatingSystemType {
	var returns awsec2.OperatingSystemType
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A basic lifecycle hook object.
type ILifecycleHook interface {
	awscdk.IResource
	// The role for the lifecycle hook to execute.
	Role() awsiam.IRole
}

// The jsii proxy for ILifecycleHook
type jsiiProxy_ILifecycleHook struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILifecycleHook) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

// Interface for autoscaling lifecycle hook targets.
type ILifecycleHookTarget interface {
	// Called when this object is used as the target of a lifecycle hook.
	Bind(scope constructs.Construct, options *BindHookTargetOptions) *LifecycleHookTargetConfig
}

// The jsii proxy for ILifecycleHookTarget
type jsiiProxy_ILifecycleHookTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ILifecycleHookTarget) Bind(scope constructs.Construct, options *BindHookTargetOptions) *LifecycleHookTargetConfig {
	var returns *LifecycleHookTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Define a life cycle hook.
//
// TODO: EXAMPLE
//
type LifecycleHook interface {
	awscdk.Resource
	ILifecycleHook
	Env() *awscdk.ResourceEnvironment
	LifecycleHookName() *string
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for LifecycleHook
type jsiiProxy_LifecycleHook struct {
	internal.Type__awscdkResource
	jsiiProxy_ILifecycleHook
}

func (j *jsiiProxy_LifecycleHook) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LifecycleHook) LifecycleHookName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleHookName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LifecycleHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LifecycleHook) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LifecycleHook) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LifecycleHook) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLifecycleHook(scope constructs.Construct, id *string, props *LifecycleHookProps) LifecycleHook {
	_init_.Initialize()

	j := jsiiProxy_LifecycleHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.LifecycleHook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLifecycleHook_Override(l LifecycleHook, scope constructs.Construct, id *string, props *LifecycleHookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.LifecycleHook",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LifecycleHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.LifecycleHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LifecycleHook_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.LifecycleHook",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (l *jsiiProxy_LifecycleHook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LifecycleHook) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (l *jsiiProxy_LifecycleHook) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (l *jsiiProxy_LifecycleHook) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LifecycleHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Lifecycle hook.
//
// TODO: EXAMPLE
//
type LifecycleHookProps struct {
	// The state of the Amazon EC2 instance to which you want to attach the lifecycle hook.
	LifecycleTransition LifecycleTransition `json:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	DefaultResult DefaultResult `json:"defaultResult"`
	// Maximum time between calls to RecordLifecycleActionHeartbeat for the hook.
	//
	// If the lifecycle hook times out, perform the action in DefaultResult.
	HeartbeatTimeout awscdk.Duration `json:"heartbeatTimeout"`
	// Name of the lifecycle hook.
	LifecycleHookName *string `json:"lifecycleHookName"`
	// Additional data to pass to the lifecycle hook target.
	NotificationMetadata *string `json:"notificationMetadata"`
	// The target of the lifecycle hook.
	NotificationTarget ILifecycleHookTarget `json:"notificationTarget"`
	// The role that allows publishing to the notification target.
	Role awsiam.IRole `json:"role"`
	// The AutoScalingGroup to add the lifecycle hook to.
	AutoScalingGroup IAutoScalingGroup `json:"autoScalingGroup"`
}

// Result of binding a lifecycle hook to a target.
//
// TODO: EXAMPLE
//
type LifecycleHookTargetConfig struct {
	// The IRole that was used to bind the lifecycle hook to the target.
	CreatedRole awsiam.IRole `json:"createdRole"`
	// The targetArn that the lifecycle hook was bound to.
	NotificationTargetArn *string `json:"notificationTargetArn"`
}

// What instance transition to attach the hook to.
type LifecycleTransition string

const (
	LifecycleTransition_INSTANCE_LAUNCHING LifecycleTransition = "INSTANCE_LAUNCHING"
	LifecycleTransition_INSTANCE_TERMINATING LifecycleTransition = "INSTANCE_TERMINATING"
)

// How the scaling metric is going to be aggregated.
type MetricAggregationType string

const (
	MetricAggregationType_AVERAGE MetricAggregationType = "AVERAGE"
	MetricAggregationType_MINIMUM MetricAggregationType = "MINIMUM"
	MetricAggregationType_MAXIMUM MetricAggregationType = "MAXIMUM"
)

// Properties for enabling tracking of an arbitrary metric.
//
// TODO: EXAMPLE
//
type MetricTargetTrackingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// Metric to track.
	//
	// The metric must represent a utilization, so that if it's higher than the
	// target value, your ASG should scale out, and if it's lower it should
	// scale in.
	Metric awscloudwatch.IMetric `json:"metric"`
	// Value to keep the metric around.
	TargetValue *float64 `json:"targetValue"`
}

// The monitoring mode for instances launched in an autoscaling group.
type Monitoring string

const (
	Monitoring_BASIC Monitoring = "BASIC"
	Monitoring_DETAILED Monitoring = "DETAILED"
)

// Properties for enabling scaling based on network utilization.
//
// TODO: EXAMPLE
//
type NetworkUtilizationScalingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// Target average bytes/seconds on each instance.
	TargetBytesPerSecond *float64 `json:"targetBytesPerSecond"`
}

// AutoScalingGroup fleet change notifications configurations.
//
// You can configure AutoScaling to send an SNS notification whenever your Auto Scaling group scales.
//
// TODO: EXAMPLE
//
type NotificationConfiguration struct {
	// SNS topic to send notifications about fleet scaling events.
	Topic awssns.ITopic `json:"topic"`
	// Which fleet scaling events triggers a notification.
	ScalingEvents ScalingEvents `json:"scalingEvents"`
}

// One of the predefined autoscaling metrics.
type PredefinedMetric string

const (
	PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION PredefinedMetric = "ASG_AVERAGE_CPU_UTILIZATION"
	PredefinedMetric_ASG_AVERAGE_NETWORK_IN PredefinedMetric = "ASG_AVERAGE_NETWORK_IN"
	PredefinedMetric_ASG_AVERAGE_NETWORK_OUT PredefinedMetric = "ASG_AVERAGE_NETWORK_OUT"
	PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET PredefinedMetric = "ALB_REQUEST_COUNT_PER_TARGET"
)

// Input for Signals.renderCreationPolicy.
//
// TODO: EXAMPLE
//
type RenderSignalsOptions struct {
	// The desiredCapacity of the ASG.
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// The minSize of the ASG.
	MinCapacity *float64 `json:"minCapacity"`
}

// Properties for enabling scaling based on request/second.
//
// TODO: EXAMPLE
//
type RequestCountScalingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// Target average requests/minute on each instance.
	TargetRequestsPerMinute *float64 `json:"targetRequestsPerMinute"`
}

// Options for customizing the rolling update.
//
// TODO: EXAMPLE
//
type RollingUpdateOptions struct {
	// The maximum number of instances that AWS CloudFormation updates at once.
	//
	// This number affects the speed of the replacement.
	MaxBatchSize *float64 `json:"maxBatchSize"`
	// The minimum number of instances that must be in service before more instances are replaced.
	//
	// This number affects the speed of the replacement.
	MinInstancesInService *float64 `json:"minInstancesInService"`
	// The percentage of instances that must signal success for the update to succeed.
	MinSuccessPercentage *float64 `json:"minSuccessPercentage"`
	// The pause time after making a change to a batch of instances.
	PauseTime awscdk.Duration `json:"pauseTime"`
	// Specifies the Auto Scaling processes to suspend during a stack update.
	//
	// Suspending processes prevents Auto Scaling from interfering with a stack
	// update.
	SuspendProcesses *[]ScalingProcess `json:"suspendProcesses"`
	// Specifies whether the Auto Scaling group waits on signals from new instances during an update.
	WaitOnResourceSignals *bool `json:"waitOnResourceSignals"`
}

// Fleet scaling events.
type ScalingEvent string

const (
	ScalingEvent_INSTANCE_LAUNCH ScalingEvent = "INSTANCE_LAUNCH"
	ScalingEvent_INSTANCE_TERMINATE ScalingEvent = "INSTANCE_TERMINATE"
	ScalingEvent_INSTANCE_TERMINATE_ERROR ScalingEvent = "INSTANCE_TERMINATE_ERROR"
	ScalingEvent_INSTANCE_LAUNCH_ERROR ScalingEvent = "INSTANCE_LAUNCH_ERROR"
	ScalingEvent_TEST_NOTIFICATION ScalingEvent = "TEST_NOTIFICATION"
)

// A list of ScalingEvents, you can use one of the predefined lists, such as ScalingEvents.ERRORS or create a custom group by instantiating a `NotificationTypes` object, e.g: `new NotificationTypes(`NotificationType.INSTANCE_LAUNCH`)`.
//
// TODO: EXAMPLE
//
type ScalingEvents interface {
}

// The jsii proxy struct for ScalingEvents
type jsiiProxy_ScalingEvents struct {
	_ byte // padding
}

func NewScalingEvents(types ...ScalingEvent) ScalingEvents {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range types {
		args = append(args, a)
	}

	j := jsiiProxy_ScalingEvents{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		args,
		&j,
	)

	return &j
}

func NewScalingEvents_Override(s ScalingEvents, types ...ScalingEvent) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range types {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		args,
		s,
	)
}

func ScalingEvents_ALL() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		"ALL",
		&returns,
	)
	return returns
}

func ScalingEvents_ERRORS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		"ERRORS",
		&returns,
	)
	return returns
}

func ScalingEvents_LAUNCH_EVENTS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		"LAUNCH_EVENTS",
		&returns,
	)
	return returns
}

func ScalingEvents_TERMINATION_EVENTS() ScalingEvents {
	_init_.Initialize()
	var returns ScalingEvents
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.ScalingEvents",
		"TERMINATION_EVENTS",
		&returns,
	)
	return returns
}

// A range of metric values in which to apply a certain scaling operation.
//
// TODO: EXAMPLE
//
type ScalingInterval struct {
	// The capacity adjustment to apply in this interval.
	//
	// The number is interpreted differently based on AdjustmentType:
	//
	// - ChangeInCapacity: add the adjustment to the current capacity.
	//   The number can be positive or negative.
	// - PercentChangeInCapacity: add or remove the given percentage of the current
	//    capacity to itself. The number can be in the range [-100..100].
	// - ExactCapacity: set the capacity to this number. The number must
	//    be positive.
	Change *float64 `json:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	Lower *float64 `json:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	Upper *float64 `json:"upper"`
}

type ScalingProcess string

const (
	ScalingProcess_LAUNCH ScalingProcess = "LAUNCH"
	ScalingProcess_TERMINATE ScalingProcess = "TERMINATE"
	ScalingProcess_HEALTH_CHECK ScalingProcess = "HEALTH_CHECK"
	ScalingProcess_REPLACE_UNHEALTHY ScalingProcess = "REPLACE_UNHEALTHY"
	ScalingProcess_AZ_REBALANCE ScalingProcess = "AZ_REBALANCE"
	ScalingProcess_ALARM_NOTIFICATION ScalingProcess = "ALARM_NOTIFICATION"
	ScalingProcess_SCHEDULED_ACTIONS ScalingProcess = "SCHEDULED_ACTIONS"
	ScalingProcess_ADD_TO_LOAD_BALANCER ScalingProcess = "ADD_TO_LOAD_BALANCER"
)

// Schedule for scheduled scaling actions.
//
// TODO: EXAMPLE
//
type Schedule interface {
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Create a schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
// See: http://crontab.org/
//
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Define a scheduled scaling action.
//
// TODO: EXAMPLE
//
type ScheduledAction interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ScheduledAction
type jsiiProxy_ScheduledAction struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ScheduledAction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledAction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduledAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewScheduledAction(scope constructs.Construct, id *string, props *ScheduledActionProps) ScheduledAction {
	_init_.Initialize()

	j := jsiiProxy_ScheduledAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.ScheduledAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewScheduledAction_Override(s ScheduledAction, scope constructs.Construct, id *string, props *ScheduledActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.ScheduledAction",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ScheduledAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.ScheduledAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ScheduledAction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.ScheduledAction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (s *jsiiProxy_ScheduledAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_ScheduledAction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (s *jsiiProxy_ScheduledAction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (s *jsiiProxy_ScheduledAction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ScheduledAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a scheduled action on an AutoScalingGroup.
//
// TODO: EXAMPLE
//
type ScheduledActionProps struct {
	// When to perform this action.
	//
	// Supports cron expressions.
	//
	// For more information about cron expressions, see https://en.wikipedia.org/wiki/Cron.
	Schedule Schedule `json:"schedule"`
	// The new desired capacity.
	//
	// At the scheduled time, set the desired capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	DesiredCapacity *float64 `json:"desiredCapacity"`
	// When this scheduled action expires.
	EndTime *time.Time `json:"endTime"`
	// The new maximum capacity.
	//
	// At the scheduled time, set the maximum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The new minimum capacity.
	//
	// At the scheduled time, set the minimum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	MinCapacity *float64 `json:"minCapacity"`
	// When this scheduled action becomes active.
	StartTime *time.Time `json:"startTime"`
	// Specifies the time zone for a cron expression.
	//
	// If a time zone is not provided, UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the IANA Time Zone Database (such as Etc/GMT+9 or Pacific/Tahiti).
	//
	// For more information, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	TimeZone *string `json:"timeZone"`
	// The AutoScalingGroup to apply the scheduled actions to.
	AutoScalingGroup IAutoScalingGroup `json:"autoScalingGroup"`
}

// Configure whether the AutoScalingGroup waits for signals.
//
// If you do configure waiting for signals, you should make sure the instances
// invoke `cfn-signal` somewhere in their UserData to signal that they have
// started up (either successfully or unsuccessfully).
//
// Signals are used both during intial creation and subsequent updates.
//
// TODO: EXAMPLE
//
type Signals interface {
	DoRender(options *SignalsOptions, count *float64) *awscdk.CfnCreationPolicy
	RenderCreationPolicy(renderOptions *RenderSignalsOptions) *awscdk.CfnCreationPolicy
}

// The jsii proxy struct for Signals
type jsiiProxy_Signals struct {
	_ byte // padding
}

func NewSignals_Override(s Signals) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.Signals",
		nil, // no parameters
		s,
	)
}

// Wait for the desiredCapacity of the AutoScalingGroup amount of signals to have been received.
//
// If no desiredCapacity has been configured, wait for minCapacity signals intead.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForAll(options *SignalsOptions) Signals {
	_init_.Initialize()

	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForAll",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Wait for a specific amount of signals to have been received.
//
// You should send one signal per instance, so this represents the number of
// instances to wait for.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForCount(count *float64, options *SignalsOptions) Signals {
	_init_.Initialize()

	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForCount",
		[]interface{}{count, options},
		&returns,
	)

	return returns
}

// Wait for the minCapacity of the AutoScalingGroup amount of signals to have been received.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForMinCapacity(options *SignalsOptions) Signals {
	_init_.Initialize()

	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForMinCapacity",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Helper to render the actual creation policy, as the logic between them is quite similar.
func (s *jsiiProxy_Signals) DoRender(options *SignalsOptions, count *float64) *awscdk.CfnCreationPolicy {
	var returns *awscdk.CfnCreationPolicy

	_jsii_.Invoke(
		s,
		"doRender",
		[]interface{}{options, count},
		&returns,
	)

	return returns
}

// Render the ASG's CreationPolicy.
func (s *jsiiProxy_Signals) RenderCreationPolicy(renderOptions *RenderSignalsOptions) *awscdk.CfnCreationPolicy {
	var returns *awscdk.CfnCreationPolicy

	_jsii_.Invoke(
		s,
		"renderCreationPolicy",
		[]interface{}{renderOptions},
		&returns,
	)

	return returns
}

// Customization options for Signal handling.
//
// TODO: EXAMPLE
//
type SignalsOptions struct {
	// The percentage of signals that need to be successful.
	//
	// If this number is less than 100, a percentage of signals may be failure
	// signals while still succeeding the creation or update in CloudFormation.
	MinSuccessPercentage *float64 `json:"minSuccessPercentage"`
	// How long to wait for the signals to be sent.
	//
	// This should reflect how long it takes your instances to start up
	// (including instance start time and instance initialization time).
	Timeout awscdk.Duration `json:"timeout"`
}

// Define a step scaling action.
//
// This kind of scaling policy adjusts the target capacity in configurable
// steps. The size of the step is configurable based on the metric's distance
// to its alarm threshold.
//
// This Action must be used as the target of a CloudWatch alarm to take effect.
//
// TODO: EXAMPLE
//
type StepScalingAction interface {
	constructs.Construct
	Node() constructs.Node
	ScalingPolicyArn() *string
	AddAdjustment(adjustment *AdjustmentTier)
	ToString() *string
}

// The jsii proxy struct for StepScalingAction
type jsiiProxy_StepScalingAction struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StepScalingAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingAction) ScalingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyArn",
		&returns,
	)
	return returns
}


func NewStepScalingAction(scope constructs.Construct, id *string, props *StepScalingActionProps) StepScalingAction {
	_init_.Initialize()

	j := jsiiProxy_StepScalingAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStepScalingAction_Override(s StepScalingAction, scope constructs.Construct, id *string, props *StepScalingActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingAction",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StepScalingAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.StepScalingAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add an adjusment interval to the ScalingAction.
func (s *jsiiProxy_StepScalingAction) AddAdjustment(adjustment *AdjustmentTier) {
	_jsii_.InvokeVoid(
		s,
		"addAdjustment",
		[]interface{}{adjustment},
	)
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StepScalingAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a scaling policy.
//
// TODO: EXAMPLE
//
type StepScalingActionProps struct {
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `json:"autoScalingGroup"`
	// How the adjustment numbers are interpreted.
	AdjustmentType AdjustmentType `json:"adjustmentType"`
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude"`
}

// Define a acaling strategy which scales depending on absolute values of some metric.
//
// You can specify the scaling behavior for various values of the metric.
//
// Implemented using one or more CloudWatch alarms and Step Scaling Policies.
//
// TODO: EXAMPLE
//
type StepScalingPolicy interface {
	constructs.Construct
	LowerAction() StepScalingAction
	LowerAlarm() awscloudwatch.Alarm
	Node() constructs.Node
	UpperAction() StepScalingAction
	UpperAlarm() awscloudwatch.Alarm
	ToString() *string
}

// The jsii proxy struct for StepScalingPolicy
type jsiiProxy_StepScalingPolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StepScalingPolicy) LowerAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"lowerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) LowerAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"lowerAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"upperAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"upperAlarm",
		&returns,
	)
	return returns
}


func NewStepScalingPolicy(scope constructs.Construct, id *string, props *StepScalingPolicyProps) StepScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_StepScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStepScalingPolicy_Override(s StepScalingPolicy, scope constructs.Construct, id *string, props *StepScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StepScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StepScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type StepScalingPolicyProps struct {
	// Metric to scale on.
	Metric awscloudwatch.IMetric `json:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	ScalingSteps *[]*ScalingInterval `json:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	AdjustmentType AdjustmentType `json:"adjustmentType"`
	// Grace period after scaling activity.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	EvaluationPeriods *float64 `json:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude"`
	// The auto scaling group.
	AutoScalingGroup IAutoScalingGroup `json:"autoScalingGroup"`
}

// TODO: EXAMPLE
//
type TargetTrackingScalingPolicy interface {
	constructs.Construct
	Node() constructs.Node
	ScalingPolicyArn() *string
	ToString() *string
}

// The jsii proxy struct for TargetTrackingScalingPolicy
type jsiiProxy_TargetTrackingScalingPolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) ScalingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyArn",
		&returns,
	)
	return returns
}


func NewTargetTrackingScalingPolicy(scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_TargetTrackingScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTargetTrackingScalingPolicy_Override(t TargetTrackingScalingPolicy, scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TargetTrackingScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TargetTrackingScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a concrete TargetTrackingPolicy.
//
// Adds the scalingTarget.
//
// TODO: EXAMPLE
//
type TargetTrackingScalingPolicyProps struct {
	// Period after a scaling completes before another scaling activity can start.
	Cooldown awscdk.Duration `json:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `json:"estimatedInstanceWarmup"`
	// The target value for the metric.
	TargetValue *float64 `json:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	CustomMetric awscloudwatch.IMetric `json:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	PredefinedMetric PredefinedMetric `json:"predefinedMetric"`
	// The resource label associated with the predefined metric.
	//
	// Should be supplied if the predefined metric is ALBRequestCountPerTarget, and the
	// format should be:
	//
	// app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>
	ResourceLabel *string `json:"resourceLabel"`
	AutoScalingGroup IAutoScalingGroup `json:"autoScalingGroup"`
}

// Specifies the termination criteria to apply before Amazon EC2 Auto Scaling chooses an instance for termination.
//
// TODO: EXAMPLE
//
type TerminationPolicy string

const (
	TerminationPolicy_ALLOCATION_STRATEGY TerminationPolicy = "ALLOCATION_STRATEGY"
	TerminationPolicy_CLOSEST_TO_NEXT_INSTANCE_HOUR TerminationPolicy = "CLOSEST_TO_NEXT_INSTANCE_HOUR"
	TerminationPolicy_DEFAULT TerminationPolicy = "DEFAULT"
	TerminationPolicy_NEWEST_INSTANCE TerminationPolicy = "NEWEST_INSTANCE"
	TerminationPolicy_OLDEST_INSTANCE TerminationPolicy = "OLDEST_INSTANCE"
	TerminationPolicy_OLDEST_LAUNCH_CONFIGURATION TerminationPolicy = "OLDEST_LAUNCH_CONFIGURATION"
	TerminationPolicy_OLDEST_LAUNCH_TEMPLATE TerminationPolicy = "OLDEST_LAUNCH_TEMPLATE"
)

// How existing instances should be updated.
//
// TODO: EXAMPLE
//
type UpdatePolicy interface {
}

// The jsii proxy struct for UpdatePolicy
type jsiiProxy_UpdatePolicy struct {
	_ byte // padding
}

func NewUpdatePolicy_Override(u UpdatePolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		nil, // no parameters
		u,
	)
}

// Create a new AutoScalingGroup and switch over to it.
func UpdatePolicy_ReplacingUpdate() UpdatePolicy {
	_init_.Initialize()

	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		"replacingUpdate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Replace the instances in the AutoScalingGroup one by one, or in batches.
func UpdatePolicy_RollingUpdate(options *RollingUpdateOptions) UpdatePolicy {
	_init_.Initialize()

	var returns UpdatePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.UpdatePolicy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}


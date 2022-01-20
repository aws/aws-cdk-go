package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

// Construct for installing the AWS ALB Contoller on EKS clusters.
//
// Use the factory functions `get` and `getOrCreate` to obtain/create instances of this controller.
//
// TODO: EXAMPLE
//
// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
//
// Experimental.
type AlbController interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for AlbController
type jsiiProxy_AlbController struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_AlbController) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlbController(scope constructs.Construct, id *string, props *AlbControllerProps) AlbController {
	_init_.Initialize()

	j := jsiiProxy_AlbController{}

	_jsii_.Create(
		"monocdk.aws_eks.AlbController",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlbController_Override(a AlbController, scope constructs.Construct, id *string, props *AlbControllerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.AlbController",
		[]interface{}{scope, id, props},
		a,
	)
}

// Create the controller construct associated with this cluster and scope.
//
// Singleton per stack/cluster.
// Experimental.
func AlbController_Create(scope constructs.Construct, props *AlbControllerProps) AlbController {
	_init_.Initialize()

	var returns AlbController

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.AlbController",
		"create",
		[]interface{}{scope, props},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AlbController_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.AlbController",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AlbController) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AlbController) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AlbController) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AlbController) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AlbController) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AlbController) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AlbController) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `AlbController`.
//
// TODO: EXAMPLE
//
// Experimental.
type AlbControllerOptions struct {
	// Version of the controller.
	// Experimental.
	Version AlbControllerVersion `json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Experimental.
	Policy interface{} `json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
}

// Properties for `AlbController`.
//
// TODO: EXAMPLE
//
// Experimental.
type AlbControllerProps struct {
	// Version of the controller.
	// Experimental.
	Version AlbControllerVersion `json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Experimental.
	Policy interface{} `json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// [disable-awslint:ref-via-interface] Cluster to install the controller onto.
	// Experimental.
	Cluster Cluster `json:"cluster" yaml:"cluster"`
}

// Controller version.
//
// Corresponds to the image tag of 'amazon/aws-load-balancer-controller' image.
//
// TODO: EXAMPLE
//
// Experimental.
type AlbControllerVersion interface {
	Custom() *bool
	Version() *string
}

// The jsii proxy struct for AlbControllerVersion
type jsiiProxy_AlbControllerVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AlbControllerVersion) Custom() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbControllerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Specify a custom version.
//
// Use this if the version you need is not available in one of the predefined versions.
// Note that in this case, you will also need to provide an IAM policy in the controller options.
// Experimental.
func AlbControllerVersion_Of(version *string) AlbControllerVersion {
	_init_.Initialize()

	var returns AlbControllerVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.AlbControllerVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func AlbControllerVersion_V2_0_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_0_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_0_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_0_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_1_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_1_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_1_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_1_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_2_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_2_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_2_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_4() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_2_4",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_3_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

// ALB Scheme.
// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/guide/ingress/annotations/#scheme
//
// Experimental.
type AlbScheme string

const (
	AlbScheme_INTERNAL AlbScheme = "INTERNAL"
	AlbScheme_INTERNET_FACING AlbScheme = "INTERNET_FACING"
)

// Options for adding worker nodes.
//
// TODO: EXAMPLE
//
// Experimental.
type AutoScalingGroupCapacityOptions struct {
	// Whether the instances can initiate connections to anywhere by default.
	// Experimental.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	// Experimental.
	AssociatePublicIpAddress *bool `json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	// Experimental.
	AutoScalingGroupName *string `json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	// Experimental.
	BlockDevices *[]*awsautoscaling.BlockDevice `json:"blockDevices" yaml:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	// Experimental.
	Cooldown awscdk.Duration `json:"cooldown" yaml:"cooldown"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	// Experimental.
	DesiredCapacity *float64 `json:"desiredCapacity" yaml:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	// Experimental.
	GroupMetrics *[]awsautoscaling.GroupMetrics `json:"groupMetrics" yaml:"groupMetrics"`
	// Configuration for health checks.
	// Experimental.
	HealthCheck awsautoscaling.HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	// Experimental.
	IgnoreUnmodifiedSizeProperties *bool `json:"ignoreUnmodifiedSizeProperties" yaml:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	// Experimental.
	InstanceMonitoring awsautoscaling.Monitoring `json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	// Experimental.
	KeyName *string `json:"keyName" yaml:"keyName"`
	// Maximum number of instances in the fleet.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
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
	// Experimental.
	MaxInstanceLifetime awscdk.Duration `json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
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
	// Experimental.
	NewInstancesProtectedFromScaleIn *bool `json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	// Experimental.
	Notifications *[]*awsautoscaling.NotificationConfiguration `json:"notifications" yaml:"notifications"`
	// SNS topic to send notifications about fleet changes.
	// Deprecated: use `notifications`
	NotificationsTopic awssns.ITopic `json:"notificationsTopic" yaml:"notificationsTopic"`
	// Configuration for replacing updates.
	//
	// Only used if updateType == UpdateType.ReplacingUpdate. Specifies how
	// many instances must signal success for the update to succeed.
	// Deprecated: Use `signals` instead
	ReplacingUpdateMinSuccessfulInstancesPercent *float64 `json:"replacingUpdateMinSuccessfulInstancesPercent" yaml:"replacingUpdateMinSuccessfulInstancesPercent"`
	// How many ResourceSignal calls CloudFormation expects before the resource is considered created.
	// Deprecated: Use `signals` instead.
	ResourceSignalCount *float64 `json:"resourceSignalCount" yaml:"resourceSignalCount"`
	// The length of time to wait for the resourceSignalCount.
	//
	// The maximum value is 43200 (12 hours).
	// Deprecated: Use `signals` instead.
	ResourceSignalTimeout awscdk.Duration `json:"resourceSignalTimeout" yaml:"resourceSignalTimeout"`
	// Configuration for rolling updates.
	//
	// Only used if updateType == UpdateType.RollingUpdate.
	// Deprecated: Use `updatePolicy` instead
	RollingUpdateConfiguration *awsautoscaling.RollingUpdateConfiguration `json:"rollingUpdateConfiguration" yaml:"rollingUpdateConfiguration"`
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
	// Experimental.
	Signals awsautoscaling.Signals `json:"signals" yaml:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	// Experimental.
	SpotPrice *string `json:"spotPrice" yaml:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	// Experimental.
	TerminationPolicies *[]awsautoscaling.TerminationPolicy `json:"terminationPolicies" yaml:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	// Experimental.
	UpdatePolicy awsautoscaling.UpdatePolicy `json:"updatePolicy" yaml:"updatePolicy"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	// Deprecated: Use `updatePolicy` instead
	UpdateType awsautoscaling.UpdateType `json:"updateType" yaml:"updateType"`
	// Where to place instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Instance type of the instances to start.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Experimental.
	BootstrapEnabled *bool `json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// EKS node bootstrapping options.
	// Experimental.
	BootstrapOptions *BootstrapOptions `json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Machine image type.
	// Experimental.
	MachineImageType MachineImageType `json:"machineImageType" yaml:"machineImageType"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Experimental.
	MapRole *bool `json:"mapRole" yaml:"mapRole"`
	// Installs the AWS spot instance interrupt handler on the cluster if it's not already added.
	//
	// Only relevant if `spotPrice` is used.
	// Experimental.
	SpotInterruptHandler *bool `json:"spotInterruptHandler" yaml:"spotInterruptHandler"`
}

// Options for adding an AutoScalingGroup as capacity.
//
// TODO: EXAMPLE
//
// Experimental.
type AutoScalingGroupOptions struct {
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Experimental.
	BootstrapEnabled *bool `json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// Allows options for node bootstrapping through EC2 user data.
	// Experimental.
	BootstrapOptions *BootstrapOptions `json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Allow options to specify different machine image type.
	// Experimental.
	MachineImageType MachineImageType `json:"machineImageType" yaml:"machineImageType"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Experimental.
	MapRole *bool `json:"mapRole" yaml:"mapRole"`
	// Installs the AWS spot instance interrupt handler on the cluster if it's not already added.
	//
	// Only relevant if `spotPrice` is configured on the auto-scaling group.
	// Experimental.
	SpotInterruptHandler *bool `json:"spotInterruptHandler" yaml:"spotInterruptHandler"`
}

// Manages mapping between IAM users and roles to Kubernetes RBAC configuration.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/en_us/eks/latest/userguide/add-user-role.html
//
// Experimental.
type AwsAuth interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	AddAccount(accountId *string)
	AddMastersRole(role awsiam.IRole, username *string)
	AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping)
	AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for AwsAuth
type jsiiProxy_AwsAuth struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_AwsAuth) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAuth(scope constructs.Construct, id *string, props *AwsAuthProps) AwsAuth {
	_init_.Initialize()

	j := jsiiProxy_AwsAuth{}

	_jsii_.Create(
		"monocdk.aws_eks.AwsAuth",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAuth_Override(a AwsAuth, scope constructs.Construct, id *string, props *AwsAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.AwsAuth",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func AwsAuth_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.AwsAuth",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Additional AWS account to add to the aws-auth configmap.
// Experimental.
func (a *jsiiProxy_AwsAuth) AddAccount(accountId *string) {
	_jsii_.InvokeVoid(
		a,
		"addAccount",
		[]interface{}{accountId},
	)
}

// Adds the specified IAM role to the `system:masters` RBAC group, which means that anyone that can assume it will be able to administer this Kubernetes system.
// Experimental.
func (a *jsiiProxy_AwsAuth) AddMastersRole(role awsiam.IRole, username *string) {
	_jsii_.InvokeVoid(
		a,
		"addMastersRole",
		[]interface{}{role, username},
	)
}

// Adds a mapping between an IAM role to a Kubernetes user and groups.
// Experimental.
func (a *jsiiProxy_AwsAuth) AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping) {
	_jsii_.InvokeVoid(
		a,
		"addRoleMapping",
		[]interface{}{role, mapping},
	)
}

// Adds a mapping between an IAM user to a Kubernetes user and groups.
// Experimental.
func (a *jsiiProxy_AwsAuth) AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping) {
	_jsii_.InvokeVoid(
		a,
		"addUserMapping",
		[]interface{}{user, mapping},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AwsAuth) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AwsAuth) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AwsAuth) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AwsAuth) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AwsAuth) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AwsAuth) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AwsAuth) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AwsAuth mapping.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsAuthMapping struct {
	// A list of groups within Kubernetes to which the role is mapped.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	Groups *[]*string `json:"groups" yaml:"groups"`
	// The user name within Kubernetes to map to the IAM role.
	// Experimental.
	Username *string `json:"username" yaml:"username"`
}

// Configuration props for the AwsAuth construct.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsAuthProps struct {
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster Cluster `json:"cluster" yaml:"cluster"`
}

// EKS node bootstrapping options.
//
// TODO: EXAMPLE
//
// Experimental.
type BootstrapOptions struct {
	// Additional command line arguments to pass to the `/etc/eks/bootstrap.sh` command.
	// See: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh
	//
	// Experimental.
	AdditionalArgs *string `json:"additionalArgs" yaml:"additionalArgs"`
	// Number of retry attempts for AWS API call (DescribeCluster).
	// Experimental.
	AwsApiRetryAttempts *float64 `json:"awsApiRetryAttempts" yaml:"awsApiRetryAttempts"`
	// Overrides the IP address to use for DNS queries within the cluster.
	// Experimental.
	DnsClusterIp *string `json:"dnsClusterIp" yaml:"dnsClusterIp"`
	// The contents of the `/etc/docker/daemon.json` file. Useful if you want a custom config differing from the default one in the EKS AMI.
	// Experimental.
	DockerConfigJson *string `json:"dockerConfigJson" yaml:"dockerConfigJson"`
	// Restores the docker default bridge network.
	// Experimental.
	EnableDockerBridge *bool `json:"enableDockerBridge" yaml:"enableDockerBridge"`
	// Extra arguments to add to the kubelet. Useful for adding labels or taints.
	//
	// For example, `--node-labels foo=bar,goo=far`.
	// Experimental.
	KubeletExtraArgs *string `json:"kubeletExtraArgs" yaml:"kubeletExtraArgs"`
	// Sets `--max-pods` for the kubelet based on the capacity of the EC2 instance.
	// Experimental.
	UseMaxPods *bool `json:"useMaxPods" yaml:"useMaxPods"`
}

// Capacity type of the managed node group.
//
// TODO: EXAMPLE
//
// Experimental.
type CapacityType string

const (
	CapacityType_SPOT CapacityType = "SPOT"
	CapacityType_ON_DEMAND CapacityType = "ON_DEMAND"
)

// A CloudFormation `AWS::EKS::Addon`.
//
// Creates an Amazon EKS add-on.
//
// Amazon EKS add-ons help to automate the provisioning and lifecycle management of common operational software for Amazon EKS clusters. Amazon EKS add-ons require clusters running version 1.18 or later because Amazon EKS add-ons rely on the Server-side Apply Kubernetes feature, which is only available in Kubernetes 1.18 and later. For more information, see [Amazon EKS add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the *Amazon EKS User Guide* .
//
// TODO: EXAMPLE
//
type CfnAddon interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AddonName() *string
	SetAddonName(val *string)
	AddonVersion() *string
	SetAddonVersion(val *string)
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClusterName() *string
	SetClusterName(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	ResolveConflicts() *string
	SetResolveConflicts(val *string)
	ServiceAccountRoleArn() *string
	SetServiceAccountRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAddon
type jsiiProxy_CfnAddon struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) AddonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) ResolveConflicts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveConflicts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) ServiceAccountRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAddon) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EKS::Addon`.
func NewCfnAddon(scope awscdk.Construct, id *string, props *CfnAddonProps) CfnAddon {
	_init_.Initialize()

	j := jsiiProxy_CfnAddon{}

	_jsii_.Create(
		"monocdk.aws_eks.CfnAddon",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EKS::Addon`.
func NewCfnAddon_Override(c CfnAddon, scope awscdk.Construct, id *string, props *CfnAddonProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.CfnAddon",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAddon) SetAddonName(val *string) {
	_jsii_.Set(
		j,
		"addonName",
		val,
	)
}

func (j *jsiiProxy_CfnAddon) SetAddonVersion(val *string) {
	_jsii_.Set(
		j,
		"addonVersion",
		val,
	)
}

func (j *jsiiProxy_CfnAddon) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnAddon) SetResolveConflicts(val *string) {
	_jsii_.Set(
		j,
		"resolveConflicts",
		val,
	)
}

func (j *jsiiProxy_CfnAddon) SetServiceAccountRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnAddon_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnAddon",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAddon_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnAddon",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAddon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnAddon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAddon_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.CfnAddon",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAddon) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAddon) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAddon) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAddon) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAddon) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAddon) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAddon) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAddon) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAddon) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnAddon) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAddon) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnAddon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAddon) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAddon) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAddon`.
//
// TODO: EXAMPLE
//
type CfnAddonProps struct {
	// The name of the add-on.
	AddonName *string `json:"addonName" yaml:"addonName"`
	// The name of the cluster.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The version of the add-on.
	AddonVersion *string `json:"addonVersion" yaml:"addonVersion"`
	// How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.
	ResolveConflicts *string `json:"resolveConflicts" yaml:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	ServiceAccountRoleArn *string `json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::EKS::Cluster`.
//
// Creates an Amazon EKS control plane.
//
// The Amazon EKS control plane consists of control plane instances that run the Kubernetes software, such as `etcd` and the API server. The control plane runs in an account managed by AWS , and the Kubernetes API is exposed by the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is single tenant and unique. It runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones and fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS also provisions elastic network interfaces in your VPC subnets to provide connectivity from the control plane instances to the nodes (for example, to support `kubectl exec` , `logs` , and `proxy` data flows).
//
// Amazon EKS nodes run in your AWS account and connect to your cluster's control plane over the Kubernetes API server endpoint and a certificate file that is created for your cluster.
//
// In most cases, it takes several minutes to create a cluster. After you create an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate with the API server and launch nodes into your cluster. For more information, see [Managing Cluster Authentication](https://docs.aws.amazon.com/eks/latest/userguide/managing-auth.html) and [Launching Amazon EKS nodes](https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html) in the *Amazon EKS User Guide* .
//
// TODO: EXAMPLE
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCertificateAuthorityData() *string
	AttrClusterSecurityGroupId() *string
	AttrEncryptionConfigKeyArn() *string
	AttrEndpoint() *string
	AttrKubernetesNetworkConfigServiceIpv6Cidr() *string
	AttrOpenIdConnectIssuerUrl() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EncryptionConfig() interface{}
	SetEncryptionConfig(val interface{})
	KubernetesNetworkConfig() interface{}
	SetKubernetesNetworkConfig(val interface{})
	Logging() interface{}
	SetLogging(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	ResourcesVpcConfig() interface{}
	SetResourcesVpcConfig(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrKubernetesNetworkConfigServiceIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKubernetesNetworkConfigServiceIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EncryptionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KubernetesNetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ResourcesVpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::EKS::Cluster`.
func NewCfnCluster(scope awscdk.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"monocdk.aws_eks.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EKS::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope awscdk.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetEncryptionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKubernetesNetworkConfig(val interface{}) {
	_jsii_.Set(
		j,
		"kubernetesNetworkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLogging(val interface{}) {
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetResourcesVpcConfig(val interface{}) {
	_jsii_.Set(
		j,
		"resourcesVpcConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The cluster control plane logging configuration for your cluster.
//
// > When updating a resource, you must include this `ClusterLogging` property if the previous CloudFormation template of the resource had it.
//
// TODO: EXAMPLE
//
type CfnCluster_ClusterLoggingProperty struct {
	// The enabled control plane logs for your cluster. All log types are disabled if the array is empty.
	//
	// > When updating a resource, you must include this `EnabledTypes` property if the previous CloudFormation template of the resource had it.
	EnabledTypes interface{} `json:"enabledTypes" yaml:"enabledTypes"`
}

// The encryption configuration for the cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionConfigProperty struct {
	// The encryption provider for the cluster.
	Provider interface{} `json:"provider" yaml:"provider"`
	// Specifies the resources to be encrypted.
	//
	// The only supported value is "secrets".
	Resources *[]*string `json:"resources" yaml:"resources"`
}

// The Kubernetes network configuration for the cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_KubernetesNetworkConfigProperty struct {
	// Specify which IP family is used to assign Kubernetes Pod and Service IP addresses.
	//
	// If you don't specify a value, `ipv4` is used by default. You can only specify an IP family when you create a cluster and can't change this value once the cluster is created. If you specify `ipv6` , the VPC and subnets that you specify for cluster creation must have both IPv4 and IPv6 CIDR blocks assigned to them.
	//
	// You can only specify `ipv6` for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on. If you specify `ipv6` , then ensure that your VPC meets the requirements listed in the considerations listed in [Assigning IPv6 addresses to Pods and Services](https://docs.aws.amazon.com/eks/latest/userguide/cni-ipv6.html) in the Amazon EKS User Guide. Kubernetes assigns Services IPv6 addresses from the unique local address range (fc00::/7). You can't specify a custom IPv6 CIDR block. Pod addresses are assigned from the subnet's IPv6 CIDR.
	IpFamily *string `json:"ipFamily" yaml:"ipFamily"`
	// Don't specify a value if you select `ipv6` for *ipFamily* .
	//
	// The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. The block must meet the following requirements:
	//
	// - Within one of the following private IP address blocks: 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16.
	// - Doesn't overlap with any CIDR block assigned to the VPC that you selected for VPC.
	// - Between /24 and /12.
	//
	// > You can only specify a custom CIDR block when you create a cluster and can't change this value once the cluster is created.
	ServiceIpv4Cidr *string `json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// The CIDR block that Kubernetes Pod and Service IP addresses are assigned from if you created a 1.21 or later cluster with version 1.10.1 or later of the Amazon VPC CNI add-on and specified `ipv6` for *ipFamily* when you created the cluster. Kubernetes assigns Service addresses from the unique local address range ( `fc00::/7` ) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	ServiceIpv6Cidr *string `json:"serviceIpv6Cidr" yaml:"serviceIpv6Cidr"`
}

// Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs.
//
// By default, cluster control plane logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS Cluster control plane logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the **Amazon EKS User Guide** .
//
// > When updating a resource, you must include this `Logging` property if the previous CloudFormation template of the resource had it. > CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported control plane logs. For more information, see [CloudWatch Pricing](https://docs.aws.amazon.com/cloudwatch/pricing/) .
//
// TODO: EXAMPLE
//
type CfnCluster_LoggingProperty struct {
	// The cluster control plane logging configuration for your cluster.
	ClusterLogging interface{} `json:"clusterLogging" yaml:"clusterLogging"`
}

// The enabled logging type.
//
// For a list of the valid logging types, see the [`types` property of `LogSetup`](https://docs.aws.amazon.com/eks/latest/APIReference/API_LogSetup.html#AmazonEKS-Type-LogSetup-types) in the *Amazon EKS API Reference* .
//
// TODO: EXAMPLE
//
type CfnCluster_LoggingTypeConfigProperty struct {
	// The name of the log type.
	Type *string `json:"type" yaml:"type"`
}

// An object representing the VPC configuration to use for an Amazon EKS cluster.
//
// > When updating a resource, you must include these properties if the previous CloudFormation template of the resource had them:
// >
// > - `EndpointPublicAccess`
// > - `EndpointPrivateAccess`
// > - `PublicAccessCidrs`
//
// TODO: EXAMPLE
//
type CfnCluster_ResourcesVpcConfigProperty struct {
	// Specify subnets for your Amazon EKS nodes.
	//
	// Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Set this value to `true` to enable private access for your cluster's Kubernetes API server endpoint.
	//
	// If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is `false` , which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that `publicAccessCidrs` includes the necessary CIDR blocks for communication with the nodes or Fargate pods. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	EndpointPrivateAccess interface{} `json:"endpointPrivateAccess" yaml:"endpointPrivateAccess"`
	// Set this value to `false` to disable public access to your cluster's Kubernetes API server endpoint.
	//
	// If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is `true` , which enables public access for your Kubernetes API server. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	EndpointPublicAccess interface{} `json:"endpointPublicAccess" yaml:"endpointPublicAccess"`
	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.
	//
	// Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is `0.0.0.0/0` . If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	PublicAccessCidrs *[]*string `json:"publicAccessCidrs" yaml:"publicAccessCidrs"`
	// Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use that allow communication between your nodes and the Kubernetes control plane.
	//
	// If you don't specify any security groups, then familiarize yourself with the difference between Amazon EKS defaults for clusters deployed with Kubernetes:
	//
	// - 1.14 Amazon EKS platform version `eks.2` and earlier
	// - 1.14 Amazon EKS platform version `eks.3` and later
	//
	// For more information, see [Amazon EKS security group considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the **Amazon EKS User Guide** .
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
}

// Properties for defining a `CfnCluster`.
//
// TODO: EXAMPLE
//
type CfnClusterProps struct {
	// The VPC configuration that's used by the cluster control plane.
	//
	// Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
	//
	// > Updates require replacement of the `SecurityGroupIds` and `SubnetIds` sub-properties.
	ResourcesVpcConfig interface{} `json:"resourcesVpcConfig" yaml:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// For more information, see [Amazon EKS Service IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the **Amazon EKS User Guide** .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The encryption configuration for the cluster.
	EncryptionConfig interface{} `json:"encryptionConfig" yaml:"encryptionConfig"`
	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig interface{} `json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// The logging configuration for your cluster.
	Logging interface{} `json:"logging" yaml:"logging"`
	// The unique name to give to your cluster.
	Name *string `json:"name" yaml:"name"`
	// The metadata that you apply to the cluster to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Cluster tags don't propagate to any other resources associated with the cluster.
	//
	// > You must have the `eks:TagResource` and `eks:UntagResource` permissions in your IAM user or IAM role used to manage the CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version *string `json:"version" yaml:"version"`
}

// A CloudFormation `AWS::EKS::FargateProfile`.
//
// Creates an AWS Fargate profile for your Amazon EKS cluster. You must have at least one Fargate profile in a cluster to be able to run pods on Fargate.
//
// The Fargate profile allows an administrator to declare which pods run on Fargate and specify which pods run on which Fargate profile. This declaration is done through the profile’s selectors. Each profile can have up to five selectors that contain a namespace and labels. A namespace is required for every selector. The label field consists of multiple optional key-value pairs. Pods that match the selectors are scheduled on Fargate. If a to-be-scheduled pod matches any of the selectors in the Fargate profile, then that pod is run on Fargate.
//
// When you create a Fargate profile, you must specify a pod execution role to use with the pods that are scheduled with the profile. This role is added to the cluster's Kubernetes [Role Based Access Control](https://docs.aws.amazon.com/https://kubernetes.io/docs/admin/authorization/rbac/) (RBAC) for authorization so that the `kubelet` that is running on the Fargate infrastructure can register with your Amazon EKS cluster so that it can appear in your cluster as a node. The pod execution role also provides IAM permissions to the Fargate infrastructure to allow read access to Amazon ECR image repositories. For more information, see [Pod Execution Role](https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the *Amazon EKS User Guide* .
//
// Fargate profiles are immutable. However, you can create a new updated profile to replace an existing profile and then delete the original after the updated profile has finished creating.
//
// If any Fargate profiles in a cluster are in the `DELETING` status, you must wait for that Fargate profile to finish deleting before you can create any other profiles in that cluster.
//
// For more information, see [AWS Fargate Profile](https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html) in the *Amazon EKS User Guide* .
//
// TODO: EXAMPLE
//
type CfnFargateProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClusterName() *string
	SetClusterName(val *string)
	CreationStack() *[]*string
	FargateProfileName() *string
	SetFargateProfileName(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	PodExecutionRoleArn() *string
	SetPodExecutionRoleArn(val *string)
	Ref() *string
	Selectors() interface{}
	SetSelectors(val interface{})
	Stack() awscdk.Stack
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	Tags() awscdk.TagManager
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFargateProfile
type jsiiProxy_CfnFargateProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFargateProfile) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) FargateProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fargateProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) PodExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Selectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFargateProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EKS::FargateProfile`.
func NewCfnFargateProfile(scope awscdk.Construct, id *string, props *CfnFargateProfileProps) CfnFargateProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnFargateProfile{}

	_jsii_.Create(
		"monocdk.aws_eks.CfnFargateProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EKS::FargateProfile`.
func NewCfnFargateProfile_Override(c CfnFargateProfile, scope awscdk.Construct, id *string, props *CfnFargateProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.CfnFargateProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFargateProfile) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnFargateProfile) SetFargateProfileName(val *string) {
	_jsii_.Set(
		j,
		"fargateProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnFargateProfile) SetPodExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"podExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFargateProfile) SetSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"selectors",
		val,
	)
}

func (j *jsiiProxy_CfnFargateProfile) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnFargateProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnFargateProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFargateProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnFargateProfile",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFargateProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnFargateProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFargateProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.CfnFargateProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFargateProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFargateProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFargateProfile) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnFargateProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A key-value pair.
//
// TODO: EXAMPLE
//
type CfnFargateProfile_LabelProperty struct {
	// Enter a key.
	Key *string `json:"key" yaml:"key"`
	// Enter a value.
	Value *string `json:"value" yaml:"value"`
}

// An object representing an AWS Fargate profile selector.
//
// TODO: EXAMPLE
//
type CfnFargateProfile_SelectorProperty struct {
	// The Kubernetes namespace that the selector should match.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain all of the labels that are specified in the selector for it to be considered a match.
	Labels interface{} `json:"labels" yaml:"labels"`
}

// Properties for defining a `CfnFargateProfile`.
//
// TODO: EXAMPLE
//
type CfnFargateProfileProps struct {
	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see [Pod Execution Role](https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the *Amazon EKS User Guide* .
	PodExecutionRoleArn *string `json:"podExecutionRoleArn" yaml:"podExecutionRoleArn"`
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors interface{} `json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	FargateProfileName *string `json:"fargateProfileName" yaml:"fargateProfileName"`
	// The IDs of subnets to launch your pods into.
	//
	// At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// The metadata to apply to the Fargate profile to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::EKS::Nodegroup`.
//
// Creates a managed node group for an Amazon EKS cluster. You can only create a node group for your cluster that is equal to the current Kubernetes version for the cluster. All node groups are created with the latest AMI release version for the respective minor Kubernetes version of the cluster, unless you deploy a custom AMI using a launch template. For more information about using launch templates, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) .
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and associated Amazon EC2 instances that are managed by AWS for an Amazon EKS cluster. Each node group uses a version of the Amazon EKS optimized Amazon Linux 2 AMI. For more information, see [Managed Node Groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in the *Amazon EKS User Guide* .
//
// TODO: EXAMPLE
//
type CfnNodegroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AmiType() *string
	SetAmiType(val *string)
	AttrArn() *string
	AttrClusterName() *string
	AttrNodegroupName() *string
	CapacityType() *string
	SetCapacityType(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClusterName() *string
	SetClusterName(val *string)
	CreationStack() *[]*string
	DiskSize() *float64
	SetDiskSize(val *float64)
	ForceUpdateEnabled() interface{}
	SetForceUpdateEnabled(val interface{})
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	Labels() interface{}
	SetLabels(val interface{})
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	NodegroupName() *string
	SetNodegroupName(val *string)
	NodeRole() *string
	SetNodeRole(val *string)
	Ref() *string
	ReleaseVersion() *string
	SetReleaseVersion(val *string)
	RemoteAccess() interface{}
	SetRemoteAccess(val interface{})
	ScalingConfig() interface{}
	SetScalingConfig(val interface{})
	Stack() awscdk.Stack
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	Tags() awscdk.TagManager
	Taints() interface{}
	SetTaints(val interface{})
	UpdateConfig() interface{}
	SetUpdateConfig(val interface{})
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnNodegroup
type jsiiProxy_CfnNodegroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnNodegroup) AmiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) AttrClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) AttrNodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) CapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) ForceUpdateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Labels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) LaunchTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) NodeRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) RemoteAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) ScalingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Taints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) UpdateConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodegroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::EKS::Nodegroup`.
func NewCfnNodegroup(scope awscdk.Construct, id *string, props *CfnNodegroupProps) CfnNodegroup {
	_init_.Initialize()

	j := jsiiProxy_CfnNodegroup{}

	_jsii_.Create(
		"monocdk.aws_eks.CfnNodegroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EKS::Nodegroup`.
func NewCfnNodegroup_Override(c CfnNodegroup, scope awscdk.Construct, id *string, props *CfnNodegroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.CfnNodegroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetAmiType(val *string) {
	_jsii_.Set(
		j,
		"amiType",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetCapacityType(val *string) {
	_jsii_.Set(
		j,
		"capacityType",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetForceUpdateEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"forceUpdateEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetInstanceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetLabels(val interface{}) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetLaunchTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"launchTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetNodegroupName(val *string) {
	_jsii_.Set(
		j,
		"nodegroupName",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetNodeRole(val *string) {
	_jsii_.Set(
		j,
		"nodeRole",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetReleaseVersion(val *string) {
	_jsii_.Set(
		j,
		"releaseVersion",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetRemoteAccess(val interface{}) {
	_jsii_.Set(
		j,
		"remoteAccess",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetScalingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"scalingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetTaints(val interface{}) {
	_jsii_.Set(
		j,
		"taints",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetUpdateConfig(val interface{}) {
	_jsii_.Set(
		j,
		"updateConfig",
		val,
	)
}

func (j *jsiiProxy_CfnNodegroup) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnNodegroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnNodegroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnNodegroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnNodegroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnNodegroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnNodegroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNodegroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.CfnNodegroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnNodegroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNodegroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnNodegroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnNodegroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnNodegroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnNodegroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object representing a node group launch template specification.
//
// The launch template cannot include [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) , [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) , [`RequestSpotInstances`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotInstances.html) , [`HibernationOptions`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_HibernationOptionsRequest.html) , or [`TerminateInstances`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TerminateInstances.html) , or the node group deployment or update will fail. For more information about launch templates, see [`CreateLaunchTemplate`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateLaunchTemplate.html) in the Amazon EC2 API Reference. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
//
// Specify either `name` or `id` , but not both.
//
// TODO: EXAMPLE
//
type CfnNodegroup_LaunchTemplateSpecificationProperty struct {
	// The ID of the launch template.
	Id *string `json:"id" yaml:"id"`
	// The name of the launch template.
	Name *string `json:"name" yaml:"name"`
	// The version of the launch template to use.
	//
	// If no version is specified, then the template's default version is used.
	Version *string `json:"version" yaml:"version"`
}

// An object representing the remote access configuration for the managed node group.
//
// TODO: EXAMPLE
//
type CfnNodegroup_RemoteAccessProperty struct {
	// The Amazon EC2 SSH key that provides access for SSH communication with the nodes in the managed node group.
	//
	// For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Linux Instances* .
	Ec2SshKey *string `json:"ec2SshKey" yaml:"ec2SshKey"`
	// The security groups that are allowed SSH access (port 22) to the nodes.
	//
	// If you specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group, then port 22 on the nodes is opened to the internet (0.0.0.0/0). For more information, see [Security Groups for Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	SourceSecurityGroups *[]*string `json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

// An object representing the scaling configuration details for the Auto Scaling group that is associated with your node group.
//
// When creating a node group, you must specify all or none of the properties. When updating a node group, you can specify any or none of the properties.
//
// TODO: EXAMPLE
//
type CfnNodegroup_ScalingConfigProperty struct {
	// The current number of nodes that the managed node group should maintain.
	//
	// > If you use Cluster Autoscaler, you shouldn't change the desiredSize value directly, as this can cause the Cluster Autoscaler to suddenly scale up or scale down.
	//
	// Whenever this parameter changes, the number of worker nodes in the node group is updated to the specified size. If this parameter is given a value that is smaller than the current number of running worker nodes, the necessary number of worker nodes are terminated to match the given value. When using CloudFormation, no action occurs if you remove this parameter from your CFN template.
	//
	// This parameter can be different from minSize in some cases, such as when starting with extra hosts for testing. This parameter can also be different when you want to start with an estimated number of needed hosts, but let Cluster Autoscaler reduce the number if there are too many. When Cluster Autoscaler is used, the desiredSize parameter is altered by Cluster Autoscaler (but can be out-of-date for short periods of time). Cluster Autoscaler doesn't scale a managed node group lower than minSize or higher than maxSize.
	DesiredSize *float64 `json:"desiredSize" yaml:"desiredSize"`
	// The maximum number of nodes that the managed node group can scale out to.
	//
	// For information about the maximum number that you can specify, see [Amazon EKS service quotas](https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the *Amazon EKS User Guide* .
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The minimum number of nodes that the managed node group can scale in to.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
}

// A property that allows a node to repel a set of pods.
//
// TODO: EXAMPLE
//
type CfnNodegroup_TaintProperty struct {
	// The effect of the taint.
	Effect *string `json:"effect" yaml:"effect"`
	// The key of the taint.
	Key *string `json:"key" yaml:"key"`
	// The value of the taint.
	Value *string `json:"value" yaml:"value"`
}

// The update configuration for the node group.
//
// TODO: EXAMPLE
//
type CfnNodegroup_UpdateConfigProperty struct {
	// The maximum number of nodes unavailable at once during a version update.
	//
	// Nodes will be updated in parallel. This value or `maxUnavailablePercentage` is required to have a value.The maximum number is 100.
	MaxUnavailable *float64 `json:"maxUnavailable" yaml:"maxUnavailable"`
	// The maximum percentage of nodes unavailable during a version update.
	//
	// This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or `maxUnavailable` is required to have a value.
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
}

// Properties for defining a `CfnNodegroup`.
//
// TODO: EXAMPLE
//
type CfnNodegroupProps struct {
	// The name of the cluster to create the node group in.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// The Amazon EKS worker node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/worker_node_IAM_role.html) in the **Amazon EKS User Guide** . If you specify `launchTemplate` , then don't specify [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	NodeRole *string `json:"nodeRole" yaml:"nodeRole"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// If you specify `launchTemplate` , then don't specify [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// The AMI type for your node group.
	//
	// GPU instance types should use the `AL2_x86_64_GPU` AMI type. Non-GPU instances should use the `AL2_x86_64` AMI type. Arm instances should use the `AL2_ARM_64` AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `amiType` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	AmiType *string `json:"amiType" yaml:"amiType"`
	// The capacity type of your managed node group.
	CapacityType *string `json:"capacityType" yaml:"capacityType"`
	// The root device disk size (in GiB) for your node group instances.
	//
	// The default disk size is 20 GiB. If you specify `launchTemplate` , then don't specify `diskSize` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	DiskSize *float64 `json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old node whether or not any pods are running on the node.
	ForceUpdateEnabled interface{} `json:"forceUpdateEnabled" yaml:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	//
	// If you specify a GPU instance type, be sure to specify `AL2_x86_64_GPU` with the `amiType` parameter. If you specify `launchTemplate` , then you can specify zero or one instance type in your launch template *or* you can specify 0-20 instance types for `instanceTypes` . If however, you specify an instance type in your launch template *and* specify any `instanceTypes` , the node group deployment will fail. If you don't specify an instance type in a launch template or for `instanceTypes` , then `t3.medium` is used, by default. If you specify `Spot` for `capacityType` , then we recommend specifying multiple values for `instanceTypes` . For more information, see [Managed node group capacity types](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	InstanceTypes *[]*string `json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels interface{} `json:"labels" yaml:"labels"`
	// An object representing a node group's launch template specification.
	//
	// If specified, then do not specify `instanceTypes` , `diskSize` , or `remoteAccess` and make sure that the launch template meets the requirements in `launchTemplateSpecification` .
	LaunchTemplate interface{} `json:"launchTemplate" yaml:"launchTemplate"`
	// The unique name to give your node group.
	NodegroupName *string `json:"nodegroupName" yaml:"nodegroupName"`
	// The AMI version of the Amazon EKS optimized AMI to use with your node group (for example, `1.14.7- *YYYYMMDD*` ). By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see [Amazon EKS optimized Linux AMI Versions](https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the *Amazon EKS User Guide* .
	//
	// > Changing this value triggers an update of the node group if one is available. However, only the latest available AMI release version is valid as an input. You cannot roll back to a previous AMI release version.
	ReleaseVersion *string `json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// If you specify `launchTemplate` , then don't specify `remoteAccess` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	RemoteAccess interface{} `json:"remoteAccess" yaml:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig interface{} `json:"scalingConfig" yaml:"scalingConfig"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.
	Tags interface{} `json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Effect is one of `No_Schedule` , `Prefer_No_Schedule` , or `No_Execute` . Kubernetes taints can be used together with tolerations to control how workloads are scheduled to your nodes.
	Taints interface{} `json:"taints" yaml:"taints"`
	// The node group update configuration.
	UpdateConfig interface{} `json:"updateConfig" yaml:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	//
	// By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `version` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Version *string `json:"version" yaml:"version"`
}

// A Cluster represents a managed Kubernetes Service (EKS).
//
// This is a fully managed cluster of API Servers (control-plane)
// The user is still required to create the worker nodes.
//
// TODO: EXAMPLE
//
// Experimental.
type Cluster interface {
	awscdk.Resource
	ICluster
	AdminRole() awsiam.Role
	AlbController() AlbController
	AwsAuth() AwsAuth
	ClusterArn() *string
	ClusterCertificateAuthorityData() *string
	ClusterEncryptionConfigKeyArn() *string
	ClusterEndpoint() *string
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	ClusterName() *string
	ClusterOpenIdConnectIssuer() *string
	ClusterOpenIdConnectIssuerUrl() *string
	ClusterSecurityGroup() awsec2.ISecurityGroup
	ClusterSecurityGroupId() *string
	Connections() awsec2.Connections
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	DefaultNodegroup() Nodegroup
	Env() *awscdk.ResourceEnvironment
	KubectlEnvironment() *map[string]*string
	KubectlLambdaRole() awsiam.IRole
	KubectlLayer() awslambda.ILayerVersion
	KubectlMemory() awscdk.Size
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	KubectlRole() awsiam.IRole
	KubectlSecurityGroup() awsec2.ISecurityGroup
	Node() awscdk.ConstructNode
	OnEventLayer() awslambda.ILayerVersion
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	PhysicalName() *string
	Prune() *bool
	Role() awsiam.IRole
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	GeneratePhysicalName() *string
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) AdminRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"adminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AlbController() AlbController {
	var returns AlbController
	_jsii_.Get(
		j,
		"albController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AwsAuth() AwsAuth {
	var returns AwsAuth
	_jsii_.Get(
		j,
		"awsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterHandlerSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterHandlerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterOpenIdConnectIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultCapacity() awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup
	_jsii_.Get(
		j,
		"defaultCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultNodegroup() Nodegroup {
	var returns Nodegroup
	_jsii_.Get(
		j,
		"defaultNodegroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlEnvironment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kubectlEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlLambdaRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlLambdaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"kubectlLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlMemory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"kubectlMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlPrivateSubnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"kubectlPrivateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KubectlSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"kubectlSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) OnEventLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"onEventLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Initiates an EKS Cluster with the supplied arguments.
// Experimental.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"monocdk.aws_eks.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Initiates an EKS Cluster with the supplied arguments.
// Experimental.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing cluster.
// Experimental.
func Cluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Cluster",
		"fromClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Cluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add nodes to this EKS cluster.
//
// The nodes will automatically be configured with the right VPC and AMI
// for the instance type and Kubernetes version.
//
// Note that if you specify `updateType: RollingUpdate` or `updateType: ReplacingUpdate`, your nodes might be replaced at deploy
// time without notice in case the recommended AMI for your machine image type has been updated by AWS.
// The default behavior for `updateType` is `None`, which means only new instances will be launched using the new AMI.
//
// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
// In addition, the [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
// daemon will be installed on all spot instances to handle
// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
// Experimental.
func (c *jsiiProxy_Cluster) AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		c,
		"addAutoScalingGroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CDK8s chart in this cluster.
//
// Returns: a `KubernetesManifest` construct representing the chart.
// Experimental.
func (c *jsiiProxy_Cluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	var returns KubernetesManifest

	_jsii_.Invoke(
		c,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

// Adds a Fargate profile to this cluster.
// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
//
// Experimental.
func (c *jsiiProxy_Cluster) AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile {
	var returns FargateProfile

	_jsii_.Invoke(
		c,
		"addFargateProfile",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a Helm chart in this cluster.
//
// Returns: a `HelmChart` construct
// Experimental.
func (c *jsiiProxy_Cluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	var returns HelmChart

	_jsii_.Invoke(
		c,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a Kubernetes resource in this cluster.
//
// The manifest will be applied/deleted using kubectl as needed.
//
// Returns: a `KubernetesResource` object.
// Experimental.
func (c *jsiiProxy_Cluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		c,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

// Add managed nodegroup to this Amazon EKS cluster.
//
// This method will create a new managed nodegroup and add into the capacity.
// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
//
// Experimental.
func (c *jsiiProxy_Cluster) AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup {
	var returns Nodegroup

	_jsii_.Invoke(
		c,
		"addNodegroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Creates a new service account with corresponding IAM Role (IRSA).
// Experimental.
func (c *jsiiProxy_Cluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	var returns ServiceAccount

	_jsii_.Invoke(
		c,
		"addServiceAccount",
		[]interface{}{id, options},
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
// Experimental.
func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Connect capacity in the form of an existing AutoScalingGroup to the EKS cluster.
//
// The AutoScalingGroup must be running an EKS-optimized AMI containing the
// /etc/eks/bootstrap.sh script. This method will configure Security Groups,
// add the right policies to the instance role, apply the right tags, and add
// the required user data to the instance's launch configuration.
//
// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
// If kubectl is enabled, the
// [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
// daemon will be installed on all spot instances to handle
// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
//
// Prefer to use `addAutoScalingGroupCapacity` if possible.
// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
//
// Experimental.
func (c *jsiiProxy_Cluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		c,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

// Experimental.
func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Fetch the load balancer address of an ingress backed by a load balancer.
// Experimental.
func (c *jsiiProxy_Cluster) GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getIngressLoadBalancerAddress",
		[]interface{}{ingressName, options},
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
// Experimental.
func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
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
// Experimental.
func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Fetch the load balancer address of a service of type 'LoadBalancer'.
// Experimental.
func (c *jsiiProxy_Cluster) GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getServiceLoadBalancerAddress",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_Cluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_Cluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_Cluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_Cluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_Cluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_Cluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for EKS clusters.
//
// TODO: EXAMPLE
//
// Experimental.
type ClusterAttributes struct {
	// The physical name of the Cluster.
	// Experimental.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData *string `json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn *string `json:"clusterEncryptionConfigKeyArn" yaml:"clusterEncryptionConfigKeyArn"`
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint *string `json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// A security group id to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	// Experimental.
	ClusterHandlerSecurityGroupId *string `json:"clusterHandlerSecurityGroupId" yaml:"clusterHandlerSecurityGroupId"`
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId *string `json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
	// Environment variables to use when running `kubectl` against this cluster.
	// Experimental.
	KubectlEnvironment *map[string]*string `json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands
	// to the cluster.
	// Experimental.
	KubectlLambdaRole awsiam.IRole `json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// This layer
	// is used by the kubectl handler to apply manifests and install helm charts.
	//
	// The handler expects the layer to include the following executables:
	//
	//     helm/helm
	//     kubectl/kubectl
	//     awscli/aws
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `json:"kubectlMemory" yaml:"kubectlMemory"`
	// Subnets to host the `kubectl` compute resources.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlPrivateSubnetIds *[]*string `json:"kubectlPrivateSubnetIds" yaml:"kubectlPrivateSubnetIds"`
	// KubectlProvider for issuing kubectl commands.
	// Experimental.
	KubectlProvider IKubectlProvider `json:"kubectlProvider" yaml:"kubectlProvider"`
	// An IAM role with cluster administrator and "system:masters" permissions.
	// Experimental.
	KubectlRoleArn *string `json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
	// A security group to use for `kubectl` execution.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlSecurityGroupId *string `json:"kubectlSecurityGroupId" yaml:"kubectlSecurityGroupId"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// The handler expects the layer to include the following node_modules:
	//
	//     proxy-agent
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `json:"onEventLayer" yaml:"onEventLayer"`
	// An Open ID Connect provider for this cluster that can be used to configure service accounts.
	//
	// You can either import an existing provider using `iam.OpenIdConnectProvider.fromProviderArn`,
	// or create a new provider using `new eks.OpenIdConnectProvider`
	// Experimental.
	OpenIdConnectProvider awsiam.IOpenIdConnectProvider `json:"openIdConnectProvider" yaml:"openIdConnectProvider"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// Additional security groups associated with this cluster.
	// Experimental.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Options for EKS clusters.
//
// TODO: EXAMPLE
//
// Experimental.
type ClusterOptions struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-kubectl" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally, visit
	// https://github.com/aws-samples/aws-lambda-layer-kubectl/blob/master/cdk/README.md
	// for instructions on how to prepare the .zip file and then define it in your
	// app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'kubectl-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.PROVIDED],
	// });
	// ```
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `json:"mastersRole" yaml:"mastersRole"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-node-proxy-agent" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally define it in your app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'proxy-agent-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.NODEJS_12_X],
	// });
	// ```
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

// Common configuration props for EKS clusters.
//
// TODO: EXAMPLE
//
// Experimental.
type ClusterProps struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-kubectl" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally, visit
	// https://github.com/aws-samples/aws-lambda-layer-kubectl/blob/master/cdk/README.md
	// for instructions on how to prepare the .zip file and then define it in your
	// app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'kubectl-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.PROVIDED],
	// });
	// ```
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `json:"mastersRole" yaml:"mastersRole"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-node-proxy-agent" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally define it in your app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'proxy-agent-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.NODEJS_12_X],
	// });
	// ```
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// Number of instances to allocate as an initial capacity for this cluster.
	//
	// Instance type can be configured through `defaultCapacityInstanceType`,
	// which defaults to `m5.large`.
	//
	// Use `cluster.addAutoScalingGroupCapacity` to add additional customized capacity. Set this
	// to `0` is you wish to avoid the initial capacity allocation.
	// Experimental.
	DefaultCapacity *float64 `json:"defaultCapacity" yaml:"defaultCapacity"`
	// The instance type to use for the default capacity.
	//
	// This will only be taken
	// into account if `defaultCapacity` is > 0.
	// Experimental.
	DefaultCapacityInstance awsec2.InstanceType `json:"defaultCapacityInstance" yaml:"defaultCapacityInstance"`
	// The default capacity type for the cluster.
	// Experimental.
	DefaultCapacityType DefaultCapacityType `json:"defaultCapacityType" yaml:"defaultCapacityType"`
	// The IAM role to pass to the Kubectl Lambda Handler.
	// Experimental.
	KubectlLambdaRole awsiam.IRole `json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
	// The tags assigned to the EKS cluster.
	// Experimental.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Options for configuring an EKS cluster.
//
// TODO: EXAMPLE
//
// Experimental.
type CommonClusterOptions struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// The type of compute resources to use for CoreDNS.
// Experimental.
type CoreDnsComputeType string

const (
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

// CPU architecture.
// Experimental.
type CpuArch string

const (
	CpuArch_ARM_64 CpuArch = "ARM_64"
	CpuArch_X86_64 CpuArch = "X86_64"
)

// The default capacity type for the cluster.
//
// TODO: EXAMPLE
//
// Experimental.
type DefaultCapacityType string

const (
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

// Construct an Amazon Linux 2 image from the latest EKS Optimized AMI published in SSM.
//
// TODO: EXAMPLE
//
// Experimental.
type EksOptimizedImage interface {
	awsec2.IMachineImage
	GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EksOptimizedImage
type jsiiProxy_EksOptimizedImage struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Experimental.
func NewEksOptimizedImage(props *EksOptimizedImageProps) EksOptimizedImage {
	_init_.Initialize()

	j := jsiiProxy_EksOptimizedImage{}

	_jsii_.Create(
		"monocdk.aws_eks.EksOptimizedImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Experimental.
func NewEksOptimizedImage_Override(e EksOptimizedImage, props *EksOptimizedImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.EksOptimizedImage",
		[]interface{}{props},
		e,
	)
}

// Return the correct image.
// Experimental.
func (e *jsiiProxy_EksOptimizedImage) GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig {
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		e,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for EksOptimizedImage.
//
// TODO: EXAMPLE
//
// Experimental.
type EksOptimizedImageProps struct {
	// What cpu architecture to retrieve the image for (arm64 or x86_64).
	// Experimental.
	CpuArch CpuArch `json:"cpuArch" yaml:"cpuArch"`
	// The Kubernetes version to use.
	// Experimental.
	KubernetesVersion *string `json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// What instance type to retrieve the image for (standard or GPU-optimized).
	// Experimental.
	NodeType NodeType `json:"nodeType" yaml:"nodeType"`
}

// Endpoint access characteristics.
//
// TODO: EXAMPLE
//
// Experimental.
type EndpointAccess interface {
	OnlyFrom(cidr ...*string) EndpointAccess
}

// The jsii proxy struct for EndpointAccess
type jsiiProxy_EndpointAccess struct {
	_ byte // padding
}

func EndpointAccess_PRIVATE() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"monocdk.aws_eks.EndpointAccess",
		"PRIVATE",
		&returns,
	)
	return returns
}

func EndpointAccess_PUBLIC() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"monocdk.aws_eks.EndpointAccess",
		"PUBLIC",
		&returns,
	)
	return returns
}

func EndpointAccess_PUBLIC_AND_PRIVATE() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"monocdk.aws_eks.EndpointAccess",
		"PUBLIC_AND_PRIVATE",
		&returns,
	)
	return returns
}

// Restrict public access to specific CIDR blocks.
//
// If public access is disabled, this method will result in an error.
// Experimental.
func (e *jsiiProxy_EndpointAccess) OnlyFrom(cidr ...*string) EndpointAccess {
	args := []interface{}{}
	for _, a := range cidr {
		args = append(args, a)
	}

	var returns EndpointAccess

	_jsii_.Invoke(
		e,
		"onlyFrom",
		args,
		&returns,
	)

	return returns
}

// Defines an EKS cluster that runs entirely on AWS Fargate.
//
// The cluster is created with a default Fargate Profile that matches the
// "default" and "kube-system" namespaces. You can add additional profiles using
// `addFargateProfile`.
//
// TODO: EXAMPLE
//
// Experimental.
type FargateCluster interface {
	Cluster
	AdminRole() awsiam.Role
	AlbController() AlbController
	AwsAuth() AwsAuth
	ClusterArn() *string
	ClusterCertificateAuthorityData() *string
	ClusterEncryptionConfigKeyArn() *string
	ClusterEndpoint() *string
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	ClusterName() *string
	ClusterOpenIdConnectIssuer() *string
	ClusterOpenIdConnectIssuerUrl() *string
	ClusterSecurityGroup() awsec2.ISecurityGroup
	ClusterSecurityGroupId() *string
	Connections() awsec2.Connections
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	DefaultNodegroup() Nodegroup
	DefaultProfile() FargateProfile
	Env() *awscdk.ResourceEnvironment
	KubectlEnvironment() *map[string]*string
	KubectlLambdaRole() awsiam.IRole
	KubectlLayer() awslambda.ILayerVersion
	KubectlMemory() awscdk.Size
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	KubectlRole() awsiam.IRole
	KubectlSecurityGroup() awsec2.ISecurityGroup
	Node() awscdk.ConstructNode
	OnEventLayer() awslambda.ILayerVersion
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	PhysicalName() *string
	Prune() *bool
	Role() awsiam.IRole
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	GeneratePhysicalName() *string
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for FargateCluster
type jsiiProxy_FargateCluster struct {
	jsiiProxy_Cluster
}

func (j *jsiiProxy_FargateCluster) AdminRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"adminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) AlbController() AlbController {
	var returns AlbController
	_jsii_.Get(
		j,
		"albController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) AwsAuth() AwsAuth {
	var returns AwsAuth
	_jsii_.Get(
		j,
		"awsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterHandlerSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterHandlerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterOpenIdConnectIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterOpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOpenIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultCapacity() awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup
	_jsii_.Get(
		j,
		"defaultCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultNodegroup() Nodegroup {
	var returns Nodegroup
	_jsii_.Get(
		j,
		"defaultNodegroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) DefaultProfile() FargateProfile {
	var returns FargateProfile
	_jsii_.Get(
		j,
		"defaultProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlEnvironment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kubectlEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlLambdaRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlLambdaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"kubectlLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlMemory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"kubectlMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlPrivateSubnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"kubectlPrivateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) KubectlSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"kubectlSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) OnEventLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"onEventLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateCluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateCluster(scope constructs.Construct, id *string, props *FargateClusterProps) FargateCluster {
	_init_.Initialize()

	j := jsiiProxy_FargateCluster{}

	_jsii_.Create(
		"monocdk.aws_eks.FargateCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateCluster_Override(f FargateCluster, scope constructs.Construct, id *string, props *FargateClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.FargateCluster",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing cluster.
// Experimental.
func FargateCluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.FargateCluster",
		"fromClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func FargateCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.FargateCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FargateCluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.FargateCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add nodes to this EKS cluster.
//
// The nodes will automatically be configured with the right VPC and AMI
// for the instance type and Kubernetes version.
//
// Note that if you specify `updateType: RollingUpdate` or `updateType: ReplacingUpdate`, your nodes might be replaced at deploy
// time without notice in case the recommended AMI for your machine image type has been updated by AWS.
// The default behavior for `updateType` is `None`, which means only new instances will be launched using the new AMI.
//
// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
// In addition, the [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
// daemon will be installed on all spot instances to handle
// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
// Experimental.
func (f *jsiiProxy_FargateCluster) AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		f,
		"addAutoScalingGroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CDK8s chart in this cluster.
//
// Returns: a `KubernetesManifest` construct representing the chart.
// Experimental.
func (f *jsiiProxy_FargateCluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	var returns KubernetesManifest

	_jsii_.Invoke(
		f,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

// Adds a Fargate profile to this cluster.
// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
//
// Experimental.
func (f *jsiiProxy_FargateCluster) AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile {
	var returns FargateProfile

	_jsii_.Invoke(
		f,
		"addFargateProfile",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a Helm chart in this cluster.
//
// Returns: a `HelmChart` construct
// Experimental.
func (f *jsiiProxy_FargateCluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	var returns HelmChart

	_jsii_.Invoke(
		f,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a Kubernetes resource in this cluster.
//
// The manifest will be applied/deleted using kubectl as needed.
//
// Returns: a `KubernetesResource` object.
// Experimental.
func (f *jsiiProxy_FargateCluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		f,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

// Add managed nodegroup to this Amazon EKS cluster.
//
// This method will create a new managed nodegroup and add into the capacity.
// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
//
// Experimental.
func (f *jsiiProxy_FargateCluster) AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup {
	var returns Nodegroup

	_jsii_.Invoke(
		f,
		"addNodegroupCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Creates a new service account with corresponding IAM Role (IRSA).
// Experimental.
func (f *jsiiProxy_FargateCluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	var returns ServiceAccount

	_jsii_.Invoke(
		f,
		"addServiceAccount",
		[]interface{}{id, options},
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
// Experimental.
func (f *jsiiProxy_FargateCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Connect capacity in the form of an existing AutoScalingGroup to the EKS cluster.
//
// The AutoScalingGroup must be running an EKS-optimized AMI containing the
// /etc/eks/bootstrap.sh script. This method will configure Security Groups,
// add the right policies to the instance role, apply the right tags, and add
// the required user data to the instance's launch configuration.
//
// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
// If kubectl is enabled, the
// [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
// daemon will be installed on all spot instances to handle
// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
//
// Prefer to use `addAutoScalingGroupCapacity` if possible.
// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
//
// Experimental.
func (f *jsiiProxy_FargateCluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		f,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

// Experimental.
func (f *jsiiProxy_FargateCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Fetch the load balancer address of an ingress backed by a load balancer.
// Experimental.
func (f *jsiiProxy_FargateCluster) GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getIngressLoadBalancerAddress",
		[]interface{}{ingressName, options},
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
// Experimental.
func (f *jsiiProxy_FargateCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
// Experimental.
func (f *jsiiProxy_FargateCluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Fetch the load balancer address of a service of type 'LoadBalancer'.
// Experimental.
func (f *jsiiProxy_FargateCluster) GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getServiceLoadBalancerAddress",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FargateCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FargateCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (f *jsiiProxy_FargateCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FargateCluster) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FargateCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (f *jsiiProxy_FargateCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (f *jsiiProxy_FargateCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configuration props for EKS Fargate.
//
// TODO: EXAMPLE
//
// Experimental.
type FargateClusterProps struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-kubectl" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally, visit
	// https://github.com/aws-samples/aws-lambda-layer-kubectl/blob/master/cdk/README.md
	// for instructions on how to prepare the .zip file and then define it in your
	// app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'kubectl-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.PROVIDED],
	// });
	// ```
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `json:"mastersRole" yaml:"mastersRole"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-node-proxy-agent" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally define it in your app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'proxy-agent-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.NODEJS_12_X],
	// });
	// ```
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// Fargate Profile to create along with the cluster.
	// Experimental.
	DefaultProfile *FargateProfileOptions `json:"defaultProfile" yaml:"defaultProfile"`
}

// Fargate profiles allows an administrator to declare which pods run on Fargate.
//
// This declaration is done through the profile’s selectors. Each
// profile can have up to five selectors that contain a namespace and optional
// labels. You must define a namespace for every selector. The label field
// consists of multiple optional key-value pairs. Pods that match a selector (by
// matching a namespace for the selector and all of the labels specified in the
// selector) are scheduled on Fargate. If a namespace selector is defined
// without any labels, Amazon EKS will attempt to schedule all pods that run in
// that namespace onto Fargate using the profile. If a to-be-scheduled pod
// matches any of the selectors in the Fargate profile, then that pod is
// scheduled on Fargate.
//
// If a pod matches multiple Fargate profiles, Amazon EKS picks one of the
// matches at random. In this case, you can specify which profile a pod should
// use by adding the following Kubernetes label to the pod specification:
// eks.amazonaws.com/fargate-profile: profile_name. However, the pod must still
// match a selector in that profile in order to be scheduled onto Fargate.
//
// TODO: EXAMPLE
//
// Experimental.
type FargateProfile interface {
	awscdk.Construct
	awscdk.ITaggable
	FargateProfileArn() *string
	FargateProfileName() *string
	Node() awscdk.ConstructNode
	PodExecutionRole() awsiam.IRole
	Tags() awscdk.TagManager
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for FargateProfile
type jsiiProxy_FargateProfile struct {
	internal.Type__awscdkConstruct
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_FargateProfile) FargateProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fargateProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) FargateProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fargateProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) PodExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"podExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateProfile(scope constructs.Construct, id *string, props *FargateProfileProps) FargateProfile {
	_init_.Initialize()

	j := jsiiProxy_FargateProfile{}

	_jsii_.Create(
		"monocdk.aws_eks.FargateProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateProfile_Override(f FargateProfile, scope constructs.Construct, id *string, props *FargateProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.FargateProfile",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FargateProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.FargateProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FargateProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FargateProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (f *jsiiProxy_FargateProfile) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (f *jsiiProxy_FargateProfile) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (f *jsiiProxy_FargateProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (f *jsiiProxy_FargateProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (f *jsiiProxy_FargateProfile) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for defining EKS Fargate Profiles.
//
// TODO: EXAMPLE
//
// Experimental.
type FargateProfileOptions struct {
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace.
	//
	// At least one selector is required and you may specify up to five selectors.
	// Experimental.
	Selectors *[]*Selector `json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName *string `json:"fargateProfileName" yaml:"fargateProfileName"`
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
	//
	// Experimental.
	PodExecutionRole awsiam.IRole `json:"podExecutionRole" yaml:"podExecutionRole"`
	// Select which subnets to launch your pods into.
	//
	// At this time, pods running
	// on Fargate are not assigned public IP addresses, so only private subnets
	// (with no direct route to an Internet Gateway) are allowed.
	//
	// You must specify the VPC to customize the subnet selection
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC from which to select subnets to launch your pods into.
	//
	// By default, all private subnets are selected. You can customize this using
	// `subnetSelection`.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Configuration props for EKS Fargate Profiles.
//
// TODO: EXAMPLE
//
// Experimental.
type FargateProfileProps struct {
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace.
	//
	// At least one selector is required and you may specify up to five selectors.
	// Experimental.
	Selectors *[]*Selector `json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName *string `json:"fargateProfileName" yaml:"fargateProfileName"`
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
	//
	// Experimental.
	PodExecutionRole awsiam.IRole `json:"podExecutionRole" yaml:"podExecutionRole"`
	// Select which subnets to launch your pods into.
	//
	// At this time, pods running
	// on Fargate are not assigned public IP addresses, so only private subnets
	// (with no direct route to an Internet Gateway) are allowed.
	//
	// You must specify the VPC to customize the subnet selection
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC from which to select subnets to launch your pods into.
	//
	// By default, all private subnets are selected. You can customize this using
	// `subnetSelection`.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The EKS cluster to apply the Fargate profile to.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster Cluster `json:"cluster" yaml:"cluster"`
}

// Represents a helm chart within the Kubernetes system.
//
// Applies/deletes the resources using `kubectl` in sync with the resource.
//
// TODO: EXAMPLE
//
// Experimental.
type HelmChart interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for HelmChart
type jsiiProxy_HelmChart struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_HelmChart) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewHelmChart(scope constructs.Construct, id *string, props *HelmChartProps) HelmChart {
	_init_.Initialize()

	j := jsiiProxy_HelmChart{}

	_jsii_.Create(
		"monocdk.aws_eks.HelmChart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHelmChart_Override(h HelmChart, scope constructs.Construct, id *string, props *HelmChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.HelmChart",
		[]interface{}{scope, id, props},
		h,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func HelmChart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.HelmChart",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HelmChart_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.HelmChart",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (h *jsiiProxy_HelmChart) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HelmChart) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (h *jsiiProxy_HelmChart) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (h *jsiiProxy_HelmChart) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HelmChart) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HelmChart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (h *jsiiProxy_HelmChart) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Helm Chart options.
//
// TODO: EXAMPLE
//
// Experimental.
type HelmChartOptions struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Experimental.
	Chart *string `json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Experimental.
	ChartAsset awss3assets.Asset `json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Experimental.
	CreateNamespace *bool `json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Experimental.
	Release *string `json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://kubernetes-charts.storage.googleapis.com/
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Experimental.
	Wait *bool `json:"wait" yaml:"wait"`
}

// Helm Chart properties.
//
// TODO: EXAMPLE
//
// Experimental.
type HelmChartProps struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Experimental.
	Chart *string `json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Experimental.
	ChartAsset awss3assets.Asset `json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Experimental.
	CreateNamespace *bool `json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Experimental.
	Release *string `json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://kubernetes-charts.storage.googleapis.com/
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Experimental.
	Wait *bool `json:"wait" yaml:"wait"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
}

// An EKS cluster.
// Experimental.
type ICluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	// Experimental.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct
	// Experimental.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesManifest` object.
	// Experimental.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Creates a new service account with corresponding IAM Role (IRSA).
	// Experimental.
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
	// Connect capacity in the form of an existing AutoScalingGroup to the EKS cluster.
	//
	// The AutoScalingGroup must be running an EKS-optimized AMI containing the
	// /etc/eks/bootstrap.sh script. This method will configure Security Groups,
	// add the right policies to the instance role, apply the right tags, and add
	// the required user data to the instance's launch configuration.
	//
	// Spot instances will be labeled `lifecycle=Ec2Spot` and tainted with `PreferNoSchedule`.
	// If kubectl is enabled, the
	// [spot interrupt handler](https://github.com/awslabs/ec2-spot-labs/tree/master/ec2-spot-eks-solution/spot-termination-handler)
	// daemon will be installed on all spot instances to handle
	// [EC2 Spot Instance Termination Notices](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/).
	//
	// Prefer to use `addAutoScalingGroupCapacity` if possible.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
	//
	// Experimental.
	ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions)
	// The unique ARN assigned to the service by AWS in the form of arn:aws:eks:.
	// Experimental.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn() *string
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint() *string
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	// The physical name of the Cluster.
	// Experimental.
	ClusterName() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId() *string
	// Custom environment variables when running `kubectl` against this cluster.
	// Experimental.
	KubectlEnvironment() *map[string]*string
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands to the cluster.
	// Experimental.
	KubectlLambdaRole() awsiam.IRole
	// An AWS Lambda layer that includes `kubectl`, `helm` and the `aws` CLI.
	//
	// If not defined, a default layer will be used.
	// Experimental.
	KubectlLayer() awslambda.ILayerVersion
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory() awscdk.Size
	// Subnets to host the `kubectl` compute resources.
	//
	// If this is undefined, the k8s endpoint is expected to be accessible
	// publicly.
	// Experimental.
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	// Kubectl Provider for issuing kubectl commands against it.
	//
	// If not defined, a default provider will be used
	// Experimental.
	KubectlProvider() IKubectlProvider
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	// Experimental.
	KubectlRole() awsiam.IRole
	// A security group to use for `kubectl` execution.
	//
	// If this is undefined, the k8s endpoint is expected to be accessible
	// publicly.
	// Experimental.
	KubectlSecurityGroup() awsec2.ISecurityGroup
	// An AWS Lambda layer that includes the NPM dependency `proxy-agent`.
	//
	// If not defined, a default layer will be used.
	// Experimental.
	OnEventLayer() awslambda.ILayerVersion
	// The Open ID Connect Provider of the cluster used to configure Service Accounts.
	// Experimental.
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Indicates whether Kubernetes resources can be automatically pruned.
	//
	// When
	// this is enabled (default), prune labels will be allocated and injected to
	// each resource. These labels will then be used when issuing the `kubectl
	// apply` operation with the `--prune` switch.
	// Experimental.
	Prune() *bool
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest {
	var returns KubernetesManifest

	_jsii_.Invoke(
		i,
		"addCdk8sChart",
		[]interface{}{id, chart, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddHelmChart(id *string, options *HelmChartOptions) HelmChart {
	var returns HelmChart

	_jsii_.Invoke(
		i,
		"addHelmChart",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest {
	args := []interface{}{id}
	for _, a := range manifest {
		args = append(args, a)
	}

	var returns KubernetesManifest

	_jsii_.Invoke(
		i,
		"addManifest",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount {
	var returns ServiceAccount

	_jsii_.Invoke(
		i,
		"addServiceAccount",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		i,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterCertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCertificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterEncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEncryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterHandlerSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterHandlerSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"clusterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlEnvironment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kubectlEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlLambdaRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlLambdaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"kubectlLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlMemory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"kubectlMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlPrivateSubnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"kubectlPrivateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlProvider() IKubectlProvider {
	var returns IKubectlProvider
	_jsii_.Get(
		j,
		"kubectlProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"kubectlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) KubectlSecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"kubectlSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) OnEventLayer() awslambda.ILayerVersion {
	var returns awslambda.ILayerVersion
	_jsii_.Get(
		j,
		"onEventLayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) OpenIdConnectProvider() awsiam.IOpenIdConnectProvider {
	var returns awsiam.IOpenIdConnectProvider
	_jsii_.Get(
		j,
		"openIdConnectProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Prune() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"prune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Imported KubectlProvider that can be used in place of the default one created by CDK.
// Experimental.
type IKubectlProvider interface {
	awscdk.IConstruct
	// The IAM execution role of the handler.
	// Experimental.
	HandlerRole() awsiam.IRole
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	RoleArn() *string
	// The custom resource provider's service token.
	// Experimental.
	ServiceToken() *string
}

// The jsii proxy for IKubectlProvider
type jsiiProxy_IKubectlProvider struct {
	internal.Type__awscdkIConstruct
}

func (j *jsiiProxy_IKubectlProvider) HandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"handlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKubectlProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKubectlProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

// NodeGroup interface.
// Experimental.
type INodegroup interface {
	awscdk.IResource
	// Name of the nodegroup.
	// Experimental.
	NodegroupName() *string
}

// The jsii proxy for INodegroup
type jsiiProxy_INodegroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

// Options for fetching an IngressLoadBalancerAddress.
//
// TODO: EXAMPLE
//
// Experimental.
type IngressLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
}

// Implementation of Kubectl Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type KubectlProvider interface {
	awscdk.NestedStack
	IKubectlProvider
	Account() *string
	ArtifactId() *string
	AvailabilityZones() *[]*string
	Dependencies() *[]awscdk.Stack
	Environment() *string
	HandlerRole() awsiam.IRole
	Nested() *bool
	NestedStackParent() awscdk.Stack
	NestedStackResource() awscdk.CfnResource
	Node() awscdk.ConstructNode
	NotificationArns() *[]*string
	ParentStack() awscdk.Stack
	Partition() *string
	Region() *string
	RoleArn() *string
	ServiceToken() *string
	StackId() *string
	StackName() *string
	Synthesizer() awscdk.IStackSynthesizer
	Tags() awscdk.TagManager
	TemplateFile() *string
	TemplateOptions() awscdk.ITemplateOptions
	TerminationProtection() *bool
	UrlSuffix() *string
	AddDependency(target awscdk.Stack, reason *string)
	AddDockerImageAsset(asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation
	AddFileAsset(asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation
	AddTransform(transform *string)
	AllocateLogicalId(cfnElement awscdk.CfnElement) *string
	ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string
	FormatArn(components *awscdk.ArnComponents) *string
	GetLogicalId(element awscdk.CfnElement) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ParseArn(arn *string, sepIfToken *string, hasName *bool) *awscdk.ArnComponents
	Prepare()
	PrepareCrossReference(_sourceStack awscdk.Stack, reference awscdk.Reference) awscdk.IResolvable
	RegionalFact(factName *string, defaultValue *string) *string
	RenameLogicalId(oldId *string, newId *string)
	ReportMissingContext(report *cxapi.MissingContext)
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	Resolve(obj interface{}) interface{}
	SetParameter(name *string, value *string)
	SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents
	Synthesize(session awscdk.ISynthesisSession)
	ToJsonString(obj interface{}, space *float64) *string
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for KubectlProvider
type jsiiProxy_KubectlProvider struct {
	internal.Type__awscdkNestedStack
	jsiiProxy_IKubectlProvider
}

func (j *jsiiProxy_KubectlProvider) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Dependencies() *[]awscdk.Stack {
	var returns *[]awscdk.Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) HandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"handlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) NestedStackParent() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) NestedStackResource() awscdk.CfnResource {
	var returns awscdk.CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) ParentStack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"parentStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Synthesizer() awscdk.IStackSynthesizer {
	var returns awscdk.IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) TemplateOptions() awscdk.ITemplateOptions {
	var returns awscdk.ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubectlProvider(scope constructs.Construct, id *string, props *KubectlProviderProps) KubectlProvider {
	_init_.Initialize()

	j := jsiiProxy_KubectlProvider{}

	_jsii_.Create(
		"monocdk.aws_eks.KubectlProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubectlProvider_Override(k KubectlProvider, scope constructs.Construct, id *string, props *KubectlProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.KubectlProvider",
		[]interface{}{scope, id, props},
		k,
	)
}

// Import an existing provider.
// Experimental.
func KubectlProvider_FromKubectlProviderAttributes(scope constructs.Construct, id *string, attrs *KubectlProviderAttributes) IKubectlProvider {
	_init_.Initialize()

	var returns IKubectlProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"fromKubectlProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Take existing provider or create new based on cluster.
// Experimental.
func KubectlProvider_GetOrCreate(scope constructs.Construct, cluster ICluster) IKubectlProvider {
	_init_.Initialize()

	var returns IKubectlProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"getOrCreate",
		[]interface{}{scope, cluster},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func KubectlProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is an object of type `NestedStack`.
// Experimental.
func KubectlProvider_IsNestedStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"isNestedStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
// Experimental.
func KubectlProvider_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
// Experimental.
func KubectlProvider_Of(construct constructs.IConstruct) awscdk.Stack {
	_init_.Initialize()

	var returns awscdk.Stack

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubectlProvider",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a dependency between this stack and another stack.
//
// This can be used to define dependencies between any two stacks within an
// app, and also supports nested stacks.
// Experimental.
func (k *jsiiProxy_KubectlProvider) AddDependency(target awscdk.Stack, reason *string) {
	_jsii_.InvokeVoid(
		k,
		"addDependency",
		[]interface{}{target, reason},
	)
}

// Register a docker image asset on this Stack.
// Deprecated: Use `stack.synthesizer.addDockerImageAsset()` if you are calling,
// and a different `IStackSynthesizer` class if you are implementing.
func (k *jsiiProxy_KubectlProvider) AddDockerImageAsset(asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation {
	var returns *awscdk.DockerImageAssetLocation

	_jsii_.Invoke(
		k,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Register a file asset on this Stack.
// Deprecated: Use `stack.synthesizer.addFileAsset()` if you are calling,
// and a different IStackSynthesizer class if you are implementing.
func (k *jsiiProxy_KubectlProvider) AddFileAsset(asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation {
	var returns *awscdk.FileAssetLocation

	_jsii_.Invoke(
		k,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
//
// Duplicate values are removed when stack is synthesized.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
//
// Experimental.
func (k *jsiiProxy_KubectlProvider) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		k,
		"addTransform",
		[]interface{}{transform},
	)
}

// Returns the naming scheme used to allocate logical IDs.
//
// By default, uses
// the `HashedAddressingScheme` but this method can be overridden to customize
// this behavior.
//
// In order to make sure logical IDs are unique and stable, we hash the resource
// construct tree path (i.e. toplevel/secondlevel/.../myresource) and add it as
// a suffix to the path components joined without a separator (CloudFormation
// IDs only allow alphanumeric characters).
//
// The result will be:
//
//    <path.join('')><md5(path.join('/')>
//      "human"      "hash"
//
// If the "human" part of the ID exceeds 240 characters, we simply trim it so
// the total ID doesn't exceed CloudFormation's 255 character limit.
//
// We only take 8 characters from the md5 hash (0.000005 chance of collision).
//
// Special cases:
//
// - If the path only contains a single component (i.e. it's a top-level
//    resource), we won't add the hash to it. The hash is not needed for
//    disamiguation and also, it allows for a more straightforward migration an
//    existing CloudFormation template to a CDK stack without logical ID changes
//    (or renames).
// - For aesthetic reasons, if the last components of the path are the same
//    (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
//    resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
//    instead of `L1L2PipelinePipeline<HASH>`
// - If a component is named "Default" it will be omitted from the path. This
//    allows refactoring higher level abstractions around constructs without affecting
//    the IDs of already deployed resources.
// - If a component is named "Resource" it will be omitted from the user-visible
//    path, but included in the hash. This reduces visual noise in the human readable
//    part of the identifier.
// Experimental.
func (k *jsiiProxy_KubectlProvider) AllocateLogicalId(cfnElement awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

// Create a CloudFormation Export for a value.
//
// Returns a string representing the corresponding `Fn.importValue()`
// expression for this Export. You can control the name for the export by
// passing the `name` option.
//
// If you don't supply a value for `name`, the value you're exporting must be
// a Resource attribute (for example: `bucket.bucketName`) and it will be
// given the same name as the automatic cross-stack reference that would be created
// if you used the attribute in another Stack.
//
// One of the uses for this method is to *remove* the relationship between
// two Stacks established by automatic cross-stack references. It will
// temporarily ensure that the CloudFormation Export still exists while you
// remove the reference from the consuming stack. After that, you can remove
// the resource and the manual export.
//
// ## Example
//
// Here is how the process works. Let's say there are two stacks,
// `producerStack` and `consumerStack`, and `producerStack` has a bucket
// called `bucket`, which is referenced by `consumerStack` (perhaps because
// an AWS Lambda Function writes into it, or something like that).
//
// It is not safe to remove `producerStack.bucket` because as the bucket is being
// deleted, `consumerStack` might still be using it.
//
// Instead, the process takes two deployments:
//
// ### Deployment 1: break the relationship
//
// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
//    stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
//    remove the Lambda Function altogether).
// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
//    will make sure the CloudFormation Export continues to exist while the relationship
//    between the two stacks is being broken.
// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
//
// ### Deployment 2: remove the bucket resource
//
// - You are now free to remove the `bucket` resource from `producerStack`.
// - Don't forget to remove the `exportValue()` call as well.
// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
// Experimental.
func (k *jsiiProxy_KubectlProvider) ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

// Creates an ARN from components.
//
// If `partition`, `region` or `account` are not specified, the stack's
// partition, region and account will be used.
//
// If any component is the empty string, an empty string will be inserted
// into the generated ARN at the location that component corresponds to.
//
// The ARN will be formatted as follows:
//
//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}}{resource-name}
//
// The required ARN pieces that are omitted will be taken from the stack that
// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
// can be 'undefined'.
// Experimental.
func (k *jsiiProxy_KubectlProvider) FormatArn(components *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
//
// This method is called when a `CfnElement` is created and used to render the
// initial logical identity of resources. Logical ID renames are applied at
// this stage.
//
// This method uses the protected method `allocateLogicalId` to render the
// logical ID for an element. To modify the naming scheme, extend the `Stack`
// class and override this method.
// Experimental.
func (k *jsiiProxy_KubectlProvider) GetLogicalId(element awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubectlProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubectlProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubectlProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Given an ARN, parses it and returns components.
//
// IF THE ARN IS A CONCRETE STRING...
//
// ...it will be parsed and validated. The separator (`sep`) will be set to '/'
// if the 6th component includes a '/', in which case, `resource` will be set
// to the value before the '/' and `resourceName` will be the rest. In case
// there is no '/', `resource` will be set to the 6th components and
// `resourceName` will be set to the rest of the string.
//
// IF THE ARN IS A TOKEN...
//
// ...it cannot be validated, since we don't have the actual value yet at the
// time of this function call. You will have to supply `sepIfToken` and
// whether or not ARNs of the expected format usually have resource names
// in order to parse it properly. The resulting `ArnComponents` object will
// contain tokens for the subexpressions of the ARN, not string literals.
//
// If the resource name could possibly contain the separator char, the actual
// resource name cannot be properly parsed. This only occurs if the separator
// char is '/', and happens for example for S3 object ARNs, IAM Role ARNs,
// IAM OIDC Provider ARNs, etc. To properly extract the resource name from a
// Tokenized ARN, you must know the resource type and call
// `Arn.extractResourceName`.
//
// Returns: an ArnComponents object which allows access to the various
// components of the ARN.
// Deprecated: use splitArn instead
func (k *jsiiProxy_KubectlProvider) ParseArn(arn *string, sepIfToken *string, hasName *bool) *awscdk.ArnComponents {
	var returns *awscdk.ArnComponents

	_jsii_.Invoke(
		k,
		"parseArn",
		[]interface{}{arn, sepIfToken, hasName},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubectlProvider) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

// Deprecated.
//
// Returns: reference itself without any change
// See: https://github.com/aws/aws-cdk/pull/7187
//
// Deprecated: cross reference handling has been moved to `App.prepare()`.
func (k *jsiiProxy_KubectlProvider) PrepareCrossReference(_sourceStack awscdk.Stack, reference awscdk.Reference) awscdk.IResolvable {
	var returns awscdk.IResolvable

	_jsii_.Invoke(
		k,
		"prepareCrossReference",
		[]interface{}{_sourceStack, reference},
		&returns,
	)

	return returns
}

// Look up a fact value for the given fact for the region of this stack.
//
// Will return a definite value only if the region of the current stack is resolved.
// If not, a lookup map will be added to the stack and the lookup will be done at
// CDK deployment time.
//
// What regions will be included in the lookup map is controlled by the
// `@aws-cdk/core:target-partitions` context value: it must be set to a list
// of partitions, and only regions from the given partitions will be included.
// If no such context key is set, all regions will be included.
//
// This function is intended to be used by construct library authors. Application
// builders can rely on the abstractions offered by construct libraries and do
// not have to worry about regional facts.
//
// If `defaultValue` is not given, it is an error if the fact is unknown for
// the given region.
// Experimental.
func (k *jsiiProxy_KubectlProvider) RegionalFact(factName *string, defaultValue *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"regionalFact",
		[]interface{}{factName, defaultValue},
		&returns,
	)

	return returns
}

// Rename a generated logical identities.
//
// To modify the naming scheme strategy, extend the `Stack` class and
// override the `allocateLogicalId` method.
// Experimental.
func (k *jsiiProxy_KubectlProvider) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		k,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

// DEPRECATED.
// Deprecated: use `reportMissingContextKey()`
func (k *jsiiProxy_KubectlProvider) ReportMissingContext(report *cxapi.MissingContext) {
	_jsii_.InvokeVoid(
		k,
		"reportMissingContext",
		[]interface{}{report},
	)
}

// Indicate that a context key was expected.
//
// Contains instructions which will be emitted into the cloud assembly on how
// the key should be supplied.
// Experimental.
func (k *jsiiProxy_KubectlProvider) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		k,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

// Resolve a tokenized value in the context of the current stack.
// Experimental.
func (k *jsiiProxy_KubectlProvider) Resolve(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Assign a value to one of the nested stack parameters.
// Experimental.
func (k *jsiiProxy_KubectlProvider) SetParameter(name *string, value *string) {
	_jsii_.InvokeVoid(
		k,
		"setParameter",
		[]interface{}{name, value},
	)
}

// Splits the provided ARN into its components.
//
// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
// and a Token representing a dynamic CloudFormation expression
// (in which case the returned components will also be dynamic CloudFormation expressions,
// encoded as Tokens).
// Experimental.
func (k *jsiiProxy_KubectlProvider) SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents {
	var returns *awscdk.ArnComponents

	_jsii_.Invoke(
		k,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubectlProvider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

// Convert an object, potentially containing tokens, to a JSON string.
// Experimental.
func (k *jsiiProxy_KubectlProvider) ToJsonString(obj interface{}, space *float64) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (k *jsiiProxy_KubectlProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubectlProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Kubectl Provider Attributes.
//
// TODO: EXAMPLE
//
// Experimental.
type KubectlProviderAttributes struct {
	// The kubectl provider lambda arn.
	// Experimental.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
	// The IAM execution role of the handler.
	//
	// This role must be able to assume kubectlRoleArn
	// Experimental.
	HandlerRole awsiam.IRole `json:"handlerRole" yaml:"handlerRole"`
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	KubectlRoleArn *string `json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
}

// Kubectl Provider Properties.
//
// TODO: EXAMPLE
//
// Experimental.
type KubectlProviderProps struct {
	// The cluster to control.
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
}

// Represents a manifest within the Kubernetes system.
//
// Alternatively, you can use `cluster.addManifest(resource[, resource, ...])`
// to define resources on this cluster.
//
// Applies/deletes the manifest using `kubectl`.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesManifest interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for KubernetesManifest
type jsiiProxy_KubernetesManifest struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_KubernetesManifest) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesManifest(scope constructs.Construct, id *string, props *KubernetesManifestProps) KubernetesManifest {
	_init_.Initialize()

	j := jsiiProxy_KubernetesManifest{}

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesManifest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesManifest_Override(k KubernetesManifest, scope constructs.Construct, id *string, props *KubernetesManifestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesManifest",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func KubernetesManifest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubernetesManifest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesManifest_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesManifest",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesManifest) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `KubernetesManifest`.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesManifestOptions struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	// Experimental.
	IngressAlb *bool `json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	// Experimental.
	IngressAlbScheme AlbScheme `json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
	// When a resource is removed from a Kubernetes manifest, it no longer appears in the manifest, and there is no way to know that this resource needs to be deleted.
	//
	// To address this, `kubectl apply` has a `--prune` option which will
	// query the cluster for all resources with a specific label and will remove
	// all the labeld resources that are not part of the applied manifest. If this
	// option is disabled and a resource is removed, it will become "orphaned" and
	// will not be deleted from the cluster.
	//
	// When this option is enabled (default), the construct will inject a label to
	// all Kubernetes resources included in this manifest which will be used to
	// prune resources when the manifest changes via `kubectl apply --prune`.
	//
	// The label name will be `aws.cdk.eks/prune-<ADDR>` where `<ADDR>` is the
	// 42-char unique address of this construct in the construct tree. Value is
	// empty.
	// See: https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#alternative-kubectl-apply-f-directory-prune-l-your-label
	//
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	// Experimental.
	SkipValidation *bool `json:"skipValidation" yaml:"skipValidation"`
}

// Properties for KubernetesManifest.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesManifestProps struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	// Experimental.
	IngressAlb *bool `json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	// Experimental.
	IngressAlbScheme AlbScheme `json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
	// When a resource is removed from a Kubernetes manifest, it no longer appears in the manifest, and there is no way to know that this resource needs to be deleted.
	//
	// To address this, `kubectl apply` has a `--prune` option which will
	// query the cluster for all resources with a specific label and will remove
	// all the labeld resources that are not part of the applied manifest. If this
	// option is disabled and a resource is removed, it will become "orphaned" and
	// will not be deleted from the cluster.
	//
	// When this option is enabled (default), the construct will inject a label to
	// all Kubernetes resources included in this manifest which will be used to
	// prune resources when the manifest changes via `kubectl apply --prune`.
	//
	// The label name will be `aws.cdk.eks/prune-<ADDR>` where `<ADDR>` is the
	// 42-char unique address of this construct in the construct tree. Value is
	// empty.
	// See: https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#alternative-kubectl-apply-f-directory-prune-l-your-label
	//
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	// Experimental.
	SkipValidation *bool `json:"skipValidation" yaml:"skipValidation"`
	// The EKS cluster to apply this manifest to.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
	// The manifest to apply.
	//
	// Consists of any number of child resources.
	//
	// When the resources are created/updated, this manifest will be applied to the
	// cluster through `kubectl apply` and when the resources or the stack is
	// deleted, the resources in the manifest will be deleted through `kubectl delete`.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Manifest *[]*map[string]interface{} `json:"manifest" yaml:"manifest"`
	// Overwrite any existing resources.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the resource is created. Otherwise, if there is already a resource
	// in the cluster with the same name, the operation will fail.
	// Experimental.
	Overwrite *bool `json:"overwrite" yaml:"overwrite"`
}

// Represents a value of a specific object deployed in the cluster.
//
// Use this to fetch any information available by the `kubectl get` command.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesObjectValue interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	Value() *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for KubernetesObjectValue
type jsiiProxy_KubernetesObjectValue struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_KubernetesObjectValue) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesObjectValue) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesObjectValue(scope constructs.Construct, id *string, props *KubernetesObjectValueProps) KubernetesObjectValue {
	_init_.Initialize()

	j := jsiiProxy_KubernetesObjectValue{}

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesObjectValue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesObjectValue_Override(k KubernetesObjectValue, scope constructs.Construct, id *string, props *KubernetesObjectValueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesObjectValue",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func KubernetesObjectValue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubernetesObjectValue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesObjectValue_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesObjectValue",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesObjectValue) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for KubernetesObjectValue.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesObjectValueProps struct {
	// The EKS cluster to fetch attributes from.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
	// JSONPath to the specific value.
	// See: https://kubernetes.io/docs/reference/kubectl/jsonpath/
	//
	// Experimental.
	JsonPath *string `json:"jsonPath" yaml:"jsonPath"`
	// The name of the object to query.
	// Experimental.
	ObjectName *string `json:"objectName" yaml:"objectName"`
	// The object type to query.
	//
	// (e.g 'service', 'pod'...)
	// Experimental.
	ObjectType *string `json:"objectType" yaml:"objectType"`
	// The namespace the object belongs to.
	// Experimental.
	ObjectNamespace *string `json:"objectNamespace" yaml:"objectNamespace"`
	// Timeout for waiting on a value.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
}

// A CloudFormation resource which applies/restores a JSON patch into a Kubernetes resource.
//
// TODO: EXAMPLE
//
// See: https://kubernetes.io/docs/tasks/run-application/update-api-object-kubectl-patch/
//
// Experimental.
type KubernetesPatch interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for KubernetesPatch
type jsiiProxy_KubernetesPatch struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_KubernetesPatch) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesPatch(scope constructs.Construct, id *string, props *KubernetesPatchProps) KubernetesPatch {
	_init_.Initialize()

	j := jsiiProxy_KubernetesPatch{}

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesPatch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesPatch_Override(k KubernetesPatch, scope constructs.Construct, id *string, props *KubernetesPatchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesPatch",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func KubernetesPatch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubernetesPatch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (k *jsiiProxy_KubernetesPatch) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for KubernetesPatch.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesPatchProps struct {
	// The JSON object to pass to `kubectl patch` when the resource is created/updated.
	// Experimental.
	ApplyPatch *map[string]interface{} `json:"applyPatch" yaml:"applyPatch"`
	// The cluster to apply the patch to.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
	// The full name of the resource to patch (e.g. `deployment/coredns`).
	// Experimental.
	ResourceName *string `json:"resourceName" yaml:"resourceName"`
	// The JSON object to pass to `kubectl patch` when the resource is removed.
	// Experimental.
	RestorePatch *map[string]interface{} `json:"restorePatch" yaml:"restorePatch"`
	// The patch type to pass to `kubectl patch`.
	//
	// The default type used by `kubectl patch` is "strategic".
	// Experimental.
	PatchType PatchType `json:"patchType" yaml:"patchType"`
	// The kubernetes API namespace.
	// Experimental.
	ResourceNamespace *string `json:"resourceNamespace" yaml:"resourceNamespace"`
}

// Kubernetes cluster version.
//
// TODO: EXAMPLE
//
// Experimental.
type KubernetesVersion interface {
	Version() *string
}

// The jsii proxy struct for KubernetesVersion
type jsiiProxy_KubernetesVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KubernetesVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
// Experimental.
func KubernetesVersion_Of(version *string) KubernetesVersion {
	_init_.Initialize()

	var returns KubernetesVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubernetesVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KubernetesVersion_V1_14() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_14",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_15() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_15",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_16() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_16",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_17() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_17",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_18() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_18",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_19() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_19",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_20() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_20",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_21() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.KubernetesVersion",
		"V1_21",
		&returns,
	)
	return returns
}

// Launch template property specification.
//
// TODO: EXAMPLE
//
// Experimental.
type LaunchTemplateSpec struct {
	// The Launch template ID.
	// Experimental.
	Id *string `json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// The machine image type.
//
// TODO: EXAMPLE
//
// Experimental.
type MachineImageType string

const (
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

// Whether the worker nodes should support GPU or just standard instances.
// Experimental.
type NodeType string

const (
	NodeType_STANDARD NodeType = "STANDARD"
	NodeType_GPU NodeType = "GPU"
	NodeType_INFERENTIA NodeType = "INFERENTIA"
)

// The Nodegroup resource class.
//
// TODO: EXAMPLE
//
// Experimental.
type Nodegroup interface {
	awscdk.Resource
	INodegroup
	Cluster() ICluster
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	NodegroupArn() *string
	NodegroupName() *string
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Nodegroup
type jsiiProxy_Nodegroup struct {
	internal.Type__awscdkResource
	jsiiProxy_INodegroup
}

func (j *jsiiProxy_Nodegroup) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) NodegroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodegroup(scope constructs.Construct, id *string, props *NodegroupProps) Nodegroup {
	_init_.Initialize()

	j := jsiiProxy_Nodegroup{}

	_jsii_.Create(
		"monocdk.aws_eks.Nodegroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNodegroup_Override(n Nodegroup, scope constructs.Construct, id *string, props *NodegroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.Nodegroup",
		[]interface{}{scope, id, props},
		n,
	)
}

// Import the Nodegroup from attributes.
// Experimental.
func Nodegroup_FromNodegroupName(scope constructs.Construct, id *string, nodegroupName *string) INodegroup {
	_init_.Initialize()

	var returns INodegroup

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Nodegroup",
		"fromNodegroupName",
		[]interface{}{scope, id, nodegroupName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Nodegroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Nodegroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Nodegroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.Nodegroup",
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
// Experimental.
func (n *jsiiProxy_Nodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (n *jsiiProxy_Nodegroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
// Experimental.
func (n *jsiiProxy_Nodegroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
// Experimental.
func (n *jsiiProxy_Nodegroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (n *jsiiProxy_Nodegroup) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (n *jsiiProxy_Nodegroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (n *jsiiProxy_Nodegroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (n *jsiiProxy_Nodegroup) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (n *jsiiProxy_Nodegroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_Nodegroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (n *jsiiProxy_Nodegroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The AMI type for your node group.
//
// GPU instance types should use the `AL2_x86_64_GPU` AMI type, which uses the
// Amazon EKS-optimized Linux AMI with GPU support. Non-GPU instances should use the `AL2_x86_64` AMI type, which
// uses the Amazon EKS-optimized Linux AMI.
//
// TODO: EXAMPLE
//
// Experimental.
type NodegroupAmiType string

const (
	NodegroupAmiType_AL2_X86_64 NodegroupAmiType = "AL2_X86_64"
	NodegroupAmiType_AL2_X86_64_GPU NodegroupAmiType = "AL2_X86_64_GPU"
	NodegroupAmiType_AL2_ARM_64 NodegroupAmiType = "AL2_ARM_64"
	NodegroupAmiType_BOTTLEROCKET_ARM_64 NodegroupAmiType = "BOTTLEROCKET_ARM_64"
	NodegroupAmiType_BOTTLEROCKET_X86_64 NodegroupAmiType = "BOTTLEROCKET_X86_64"
)

// The Nodegroup Options for addNodeGroup() method.
//
// TODO: EXAMPLE
//
// Experimental.
type NodegroupOptions struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Experimental.
	AmiType NodegroupAmiType `json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Experimental.
	CapacityType CapacityType `json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Experimental.
	DesiredSize *float64 `json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Experimental.
	DiskSize *float64 `json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Experimental.
	ForceUpdate *bool `json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU` with the amiType parameter.
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Experimental.
	Labels *map[string]*string `json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Experimental.
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Experimental.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Experimental.
	NodegroupName *string `json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Experimental.
	NodeRole awsiam.IRole `json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Experimental.
	ReleaseVersion *string `json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Experimental.
	Subnets *awsec2.SubnetSelection `json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Experimental.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Experimental.
	Taints *[]*TaintSpec `json:"taints" yaml:"taints"`
}

// NodeGroup properties interface.
//
// TODO: EXAMPLE
//
// Experimental.
type NodegroupProps struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Experimental.
	AmiType NodegroupAmiType `json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Experimental.
	CapacityType CapacityType `json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Experimental.
	DesiredSize *float64 `json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Experimental.
	DiskSize *float64 `json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Experimental.
	ForceUpdate *bool `json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU` with the amiType parameter.
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Experimental.
	Labels *map[string]*string `json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Experimental.
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Experimental.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Experimental.
	NodegroupName *string `json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Experimental.
	NodeRole awsiam.IRole `json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Experimental.
	ReleaseVersion *string `json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Experimental.
	Subnets *awsec2.SubnetSelection `json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Experimental.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Experimental.
	Taints *[]*TaintSpec `json:"taints" yaml:"taints"`
	// Cluster resource.
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
}

// The remote access (SSH) configuration to use with your node group.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-remoteaccess.html
//
// Experimental.
type NodegroupRemoteAccess struct {
	// The Amazon EC2 SSH key that provides access for SSH communication with the worker nodes in the managed node group.
	// Experimental.
	SshKeyName *string `json:"sshKeyName" yaml:"sshKeyName"`
	// The security groups that are allowed SSH access (port 22) to the worker nodes.
	//
	// If you specify an Amazon EC2 SSH
	// key but do not specify a source security group when you create a managed node group, then port 22 on the worker
	// nodes is opened to the internet (0.0.0.0/0).
	// Experimental.
	SourceSecurityGroups *[]awsec2.ISecurityGroup `json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

// IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce.
//
// You use an IAM OIDC identity provider
// when you want to establish trust between an OIDC-compatible IdP and your AWS
// account.
//
// This implementation has default values for thumbprints and clientIds props
// that will be compatible with the eks cluster
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
//
// Experimental.
type OpenIdConnectProvider interface {
	awsiam.OpenIdConnectProvider
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	OpenIdConnectProviderArn() *string
	OpenIdConnectProviderIssuer() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for OpenIdConnectProvider
type jsiiProxy_OpenIdConnectProvider struct {
	internal.Type__awsiamOpenIdConnectProvider
}

func (j *jsiiProxy_OpenIdConnectProvider) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) OpenIdConnectProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) OpenIdConnectProviderIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Defines an OpenID Connect provider.
// Experimental.
func NewOpenIdConnectProvider(scope constructs.Construct, id *string, props *OpenIdConnectProviderProps) OpenIdConnectProvider {
	_init_.Initialize()

	j := jsiiProxy_OpenIdConnectProvider{}

	_jsii_.Create(
		"monocdk.aws_eks.OpenIdConnectProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an OpenID Connect provider.
// Experimental.
func NewOpenIdConnectProvider_Override(o OpenIdConnectProvider, scope constructs.Construct, id *string, props *OpenIdConnectProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.OpenIdConnectProvider",
		[]interface{}{scope, id, props},
		o,
	)
}

// Imports an Open ID connect provider from an ARN.
// Experimental.
func OpenIdConnectProvider_FromOpenIdConnectProviderArn(scope constructs.Construct, id *string, openIdConnectProviderArn *string) awsiam.IOpenIdConnectProvider {
	_init_.Initialize()

	var returns awsiam.IOpenIdConnectProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.OpenIdConnectProvider",
		"fromOpenIdConnectProviderArn",
		[]interface{}{scope, id, openIdConnectProviderArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func OpenIdConnectProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.OpenIdConnectProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func OpenIdConnectProvider_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.OpenIdConnectProvider",
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
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
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
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		o,
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
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		o,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		o,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) Prepare() {
	_jsii_.InvokeVoid(
		o,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		o,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (o *jsiiProxy_OpenIdConnectProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization properties for `OpenIdConnectProvider`.
//
// TODO: EXAMPLE
//
// Experimental.
type OpenIdConnectProviderProps struct {
	// The URL of the identity provider.
	//
	// The URL must begin with https:// and
	// should correspond to the iss claim in the provider's OpenID Connect ID
	// tokens. Per the OIDC standard, path components are allowed but query
	// parameters are not. Typically the URL consists of only a hostname, like
	// https://server.example.org or https://example.com.
	//
	// You can find your OIDC Issuer URL by:
	// aws eks describe-cluster --name %cluster_name% --query "cluster.identity.oidc.issuer" --output text
	// Experimental.
	Url *string `json:"url" yaml:"url"`
}

// Values for `kubectl patch` --type argument.
// Experimental.
type PatchType string

const (
	PatchType_JSON PatchType = "JSON"
	PatchType_MERGE PatchType = "MERGE"
	PatchType_STRATEGIC PatchType = "STRATEGIC"
)

// Fargate profile selector.
//
// TODO: EXAMPLE
//
// Experimental.
type Selector struct {
	// The Kubernetes namespace that the selector should match.
	//
	// You must specify a namespace for a selector. The selector only matches pods
	// that are created in this namespace, but you can create multiple selectors
	// to target multiple namespaces.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain
	// all of the labels that are specified in the selector for it to be
	// considered a match.
	// Experimental.
	Labels *map[string]*string `json:"labels" yaml:"labels"`
}

// Service Account.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceAccount interface {
	awscdk.Construct
	awsiam.IPrincipal
	AssumeRoleAction() *string
	GrantPrincipal() awsiam.IPrincipal
	Node() awscdk.ConstructNode
	PolicyFragment() awsiam.PrincipalPolicyFragment
	Role() awsiam.IRole
	ServiceAccountName() *string
	ServiceAccountNamespace() *string
	AddToPolicy(statement awsiam.PolicyStatement) *bool
	AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for ServiceAccount
type jsiiProxy_ServiceAccount struct {
	internal.Type__awscdkConstruct
	internal.Type__awsiamIPrincipal
}

func (j *jsiiProxy_ServiceAccount) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) PolicyFragment() awsiam.PrincipalPolicyFragment {
	var returns awsiam.PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ServiceAccountNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNamespace",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceAccount(scope constructs.Construct, id *string, props *ServiceAccountProps) ServiceAccount {
	_init_.Initialize()

	j := jsiiProxy_ServiceAccount{}

	_jsii_.Create(
		"monocdk.aws_eks.ServiceAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceAccount_Override(s ServiceAccount, scope constructs.Construct, id *string, props *ServiceAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.ServiceAccount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ServiceAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.ServiceAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add to the policy of this principal.
// Deprecated: use `addToPrincipalPolicy()`
func (s *jsiiProxy_ServiceAccount) AddToPolicy(statement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Add to the policy of this principal.
// Experimental.
func (s *jsiiProxy_ServiceAccount) AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult {
	var returns *awsiam.AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_ServiceAccount) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServiceAccount) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_ServiceAccount) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_ServiceAccount) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_ServiceAccount) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_ServiceAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_ServiceAccount) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `ServiceAccount`.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceAccountOptions struct {
	// The name of the service account.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The namespace of the service account.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
}

// Properties for defining service accounts.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceAccountProps struct {
	// The name of the service account.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The namespace of the service account.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The cluster to apply the patch to.
	// Experimental.
	Cluster ICluster `json:"cluster" yaml:"cluster"`
}

// Options for fetching a ServiceLoadBalancerAddress.
//
// TODO: EXAMPLE
//
// Experimental.
type ServiceLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
}

// Effect types of kubernetes node taint.
//
// TODO: EXAMPLE
//
// Experimental.
type TaintEffect string

const (
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

// Taint interface.
//
// TODO: EXAMPLE
//
// Experimental.
type TaintSpec struct {
	// Effect type.
	// Experimental.
	Effect TaintEffect `json:"effect" yaml:"effect"`
	// Taint key.
	// Experimental.
	Key *string `json:"key" yaml:"key"`
	// Taint value.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}


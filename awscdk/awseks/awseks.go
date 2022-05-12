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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion albControllerVersion
//   var cluster cluster
//   var policy interface{}
//
//   albController := awscdk.Aws_eks.NewAlbController(this, jsii.String("MyAlbController"), &albControllerProps{
//   	cluster: cluster,
//   	version: albControllerVersion,
//
//   	// the properties below are optional
//   	policy: policy,
//   	repository: jsii.String("repository"),
//   })
//
// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
//
// Experimental.
type AlbController interface {
	awscdk.Construct
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

func (a *jsiiProxy_AlbController) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbController) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (a *jsiiProxy_AlbController) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbController) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	albController: &albControllerOptions{
//   		version: eks.albControllerVersion_V2_4_1(),
//   	},
//   })
//
// Experimental.
type AlbControllerOptions struct {
	// Version of the controller.
	// Experimental.
	Version AlbControllerVersion `field:"required" json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Experimental.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

// Properties for `AlbController`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion albControllerVersion
//   var cluster cluster
//   var policy interface{}
//
//   albControllerProps := &albControllerProps{
//   	cluster: cluster,
//   	version: albControllerVersion,
//
//   	// the properties below are optional
//   	policy: policy,
//   	repository: jsii.String("repository"),
//   }
//
// Experimental.
type AlbControllerProps struct {
	// Version of the controller.
	// Experimental.
	Version AlbControllerVersion `field:"required" json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Experimental.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// [disable-awslint:ref-via-interface] Cluster to install the controller onto.
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

// Controller version.
//
// Corresponds to the image tag of 'amazon/aws-load-balancer-controller' image.
//
// Example:
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	albController: &albControllerOptions{
//   		version: eks.albControllerVersion_V2_4_1(),
//   	},
//   })
//
// Experimental.
type AlbControllerVersion interface {
	// Whether or not its a custom version.
	// Experimental.
	Custom() *bool
	// The version string.
	// Experimental.
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

func AlbControllerVersion_V2_4_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"monocdk.aws_eks.AlbControllerVersion",
		"V2_4_1",
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
	// The nodes of an internal load balancer have only private IP addresses.
	//
	// The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes.
	// Therefore, internal load balancers can only route requests from clients with access to the VPC for the load balancer.
	// Experimental.
	AlbScheme_INTERNAL AlbScheme = "INTERNAL"
	// An internet-facing load balancer has a publicly resolvable DNS name, so it can route requests from clients over the internet to the EC2 instances that are registered with the load balancer.
	// Experimental.
	AlbScheme_INTERNET_FACING AlbScheme = "INTERNET_FACING"
)

// Options for adding worker nodes.
//
// Example:
//   var cluster cluster
//
//   cluster.addAutoScalingGroupCapacity(jsii.String("BottlerocketNodes"), &autoScalingGroupCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   	minCapacity: jsii.Number(2),
//   	machineImageType: eks.machineImageType_BOTTLEROCKET,
//   })
//
// Experimental.
type AutoScalingGroupCapacityOptions struct {
	// Whether the instances can initiate connections to anywhere by default.
	// Experimental.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	// Experimental.
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	// Experimental.
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	// Experimental.
	BlockDevices *[]*awsautoscaling.BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	// Experimental.
	GroupMetrics *[]awsautoscaling.GroupMetrics `field:"optional" json:"groupMetrics" yaml:"groupMetrics"`
	// Configuration for health checks.
	// Experimental.
	HealthCheck awsautoscaling.HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	// Experimental.
	IgnoreUnmodifiedSizeProperties *bool `field:"optional" json:"ignoreUnmodifiedSizeProperties" yaml:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	// Experimental.
	InstanceMonitoring awsautoscaling.Monitoring `field:"optional" json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Maximum number of instances in the fleet.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
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
	MaxInstanceLifetime awscdk.Duration `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
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
	NewInstancesProtectedFromScaleIn *bool `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	// Experimental.
	Notifications *[]*awsautoscaling.NotificationConfiguration `field:"optional" json:"notifications" yaml:"notifications"`
	// SNS topic to send notifications about fleet changes.
	// Deprecated: use `notifications`.
	NotificationsTopic awssns.ITopic `field:"optional" json:"notificationsTopic" yaml:"notificationsTopic"`
	// Configuration for replacing updates.
	//
	// Only used if updateType == UpdateType.ReplacingUpdate. Specifies how
	// many instances must signal success for the update to succeed.
	// Deprecated: Use `signals` instead.
	ReplacingUpdateMinSuccessfulInstancesPercent *float64 `field:"optional" json:"replacingUpdateMinSuccessfulInstancesPercent" yaml:"replacingUpdateMinSuccessfulInstancesPercent"`
	// How many ResourceSignal calls CloudFormation expects before the resource is considered created.
	// Deprecated: Use `signals` instead.
	ResourceSignalCount *float64 `field:"optional" json:"resourceSignalCount" yaml:"resourceSignalCount"`
	// The length of time to wait for the resourceSignalCount.
	//
	// The maximum value is 43200 (12 hours).
	// Deprecated: Use `signals` instead.
	ResourceSignalTimeout awscdk.Duration `field:"optional" json:"resourceSignalTimeout" yaml:"resourceSignalTimeout"`
	// Configuration for rolling updates.
	//
	// Only used if updateType == UpdateType.RollingUpdate.
	// Deprecated: Use `updatePolicy` instead.
	RollingUpdateConfiguration *awsautoscaling.RollingUpdateConfiguration `field:"optional" json:"rollingUpdateConfiguration" yaml:"rollingUpdateConfiguration"`
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
	Signals awsautoscaling.Signals `field:"optional" json:"signals" yaml:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	// Experimental.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	// Experimental.
	TerminationPolicies *[]awsautoscaling.TerminationPolicy `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	// Experimental.
	UpdatePolicy awsautoscaling.UpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	// Deprecated: Use `updatePolicy` instead.
	UpdateType awsautoscaling.UpdateType `field:"optional" json:"updateType" yaml:"updateType"`
	// Where to place instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Instance type of the instances to start.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Experimental.
	BootstrapEnabled *bool `field:"optional" json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// EKS node bootstrapping options.
	// Experimental.
	BootstrapOptions *BootstrapOptions `field:"optional" json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Machine image type.
	// Experimental.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Experimental.
	MapRole *bool `field:"optional" json:"mapRole" yaml:"mapRole"`
	// Installs the AWS spot instance interrupt handler on the cluster if it's not already added.
	//
	// Only relevant if `spotPrice` is used.
	// Experimental.
	SpotInterruptHandler *bool `field:"optional" json:"spotInterruptHandler" yaml:"spotInterruptHandler"`
}

// Options for adding an AutoScalingGroup as capacity.
//
// Example:
//   var cluster cluster
//   var asg autoScalingGroup
//
//   cluster.connectAutoScalingGroupCapacity(asg, &autoScalingGroupOptions{
//   })
//
// Experimental.
type AutoScalingGroupOptions struct {
	// Configures the EC2 user-data script for instances in this autoscaling group to bootstrap the node (invoke `/etc/eks/bootstrap.sh`) and associate it with the EKS cluster.
	//
	// If you wish to provide a custom user data script, set this to `false` and
	// manually invoke `autoscalingGroup.addUserData()`.
	// Experimental.
	BootstrapEnabled *bool `field:"optional" json:"bootstrapEnabled" yaml:"bootstrapEnabled"`
	// Allows options for node bootstrapping through EC2 user data.
	// Experimental.
	BootstrapOptions *BootstrapOptions `field:"optional" json:"bootstrapOptions" yaml:"bootstrapOptions"`
	// Allow options to specify different machine image type.
	// Experimental.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Will automatically update the aws-auth ConfigMap to map the IAM instance role to RBAC.
	//
	// This cannot be explicitly set to `true` if the cluster has kubectl disabled.
	// Experimental.
	MapRole *bool `field:"optional" json:"mapRole" yaml:"mapRole"`
	// Installs the AWS spot instance interrupt handler on the cluster if it's not already added.
	//
	// Only relevant if `spotPrice` is configured on the auto-scaling group.
	// Experimental.
	SpotInterruptHandler *bool `field:"optional" json:"spotInterruptHandler" yaml:"spotInterruptHandler"`
}

// Manages mapping between IAM users and roles to Kubernetes RBAC configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   awsAuth := awscdk.Aws_eks.NewAwsAuth(this, jsii.String("MyAwsAuth"), &awsAuthProps{
//   	cluster: cluster,
//   })
//
// See: https://docs.aws.amazon.com/en_us/eks/latest/userguide/add-user-role.html
//
// Experimental.
type AwsAuth interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Additional AWS account to add to the aws-auth configmap.
	// Experimental.
	AddAccount(accountId *string)
	// Adds the specified IAM role to the `system:masters` RBAC group, which means that anyone that can assume it will be able to administer this Kubernetes system.
	// Experimental.
	AddMastersRole(role awsiam.IRole, username *string)
	// Adds a mapping between an IAM role to a Kubernetes user and groups.
	// Experimental.
	AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping)
	// Adds a mapping between an IAM user to a Kubernetes user and groups.
	// Experimental.
	AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping)
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

func (a *jsiiProxy_AwsAuth) AddAccount(accountId *string) {
	_jsii_.InvokeVoid(
		a,
		"addAccount",
		[]interface{}{accountId},
	)
}

func (a *jsiiProxy_AwsAuth) AddMastersRole(role awsiam.IRole, username *string) {
	_jsii_.InvokeVoid(
		a,
		"addMastersRole",
		[]interface{}{role, username},
	)
}

func (a *jsiiProxy_AwsAuth) AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping) {
	_jsii_.InvokeVoid(
		a,
		"addRoleMapping",
		[]interface{}{role, mapping},
	)
}

func (a *jsiiProxy_AwsAuth) AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping) {
	_jsii_.InvokeVoid(
		a,
		"addUserMapping",
		[]interface{}{user, mapping},
	)
}

func (a *jsiiProxy_AwsAuth) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuth) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (a *jsiiProxy_AwsAuth) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuth) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//
//   adminUser := iam.NewUser(this, jsii.String("Admin"))
//   cluster.awsAuth.addUserMapping(adminUser, &awsAuthMapping{
//   	groups: []*string{
//   		jsii.String("system:masters"),
//   	},
//   })
//
// Experimental.
type AwsAuthMapping struct {
	// A list of groups within Kubernetes to which the role is mapped.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	Groups *[]*string `field:"required" json:"groups" yaml:"groups"`
	// The user name within Kubernetes to map to the IAM role.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Configuration props for the AwsAuth construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   awsAuthProps := &awsAuthProps{
//   	cluster: cluster,
//   }
//
// Experimental.
type AwsAuthProps struct {
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

// EKS node bootstrapping options.
//
// Example:
//   var cluster cluster
//
//   cluster.addAutoScalingGroupCapacity(jsii.String("spot"), &autoScalingGroupCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   	minCapacity: jsii.Number(2),
//   	bootstrapOptions: &bootstrapOptions{
//   		kubeletExtraArgs: jsii.String("--node-labels foo=bar,goo=far"),
//   		awsApiRetryAttempts: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type BootstrapOptions struct {
	// Additional command line arguments to pass to the `/etc/eks/bootstrap.sh` command.
	// See: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh
	//
	// Experimental.
	AdditionalArgs *string `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// Number of retry attempts for AWS API call (DescribeCluster).
	// Experimental.
	AwsApiRetryAttempts *float64 `field:"optional" json:"awsApiRetryAttempts" yaml:"awsApiRetryAttempts"`
	// Overrides the IP address to use for DNS queries within the cluster.
	// Experimental.
	DnsClusterIp *string `field:"optional" json:"dnsClusterIp" yaml:"dnsClusterIp"`
	// The contents of the `/etc/docker/daemon.json` file. Useful if you want a custom config differing from the default one in the EKS AMI.
	// Experimental.
	DockerConfigJson *string `field:"optional" json:"dockerConfigJson" yaml:"dockerConfigJson"`
	// Restores the docker default bridge network.
	// Experimental.
	EnableDockerBridge *bool `field:"optional" json:"enableDockerBridge" yaml:"enableDockerBridge"`
	// Extra arguments to add to the kubelet. Useful for adding labels or taints.
	//
	// For example, `--node-labels foo=bar,goo=far`.
	// Experimental.
	KubeletExtraArgs *string `field:"optional" json:"kubeletExtraArgs" yaml:"kubeletExtraArgs"`
	// Sets `--max-pods` for the kubelet based on the capacity of the EC2 instance.
	// Experimental.
	UseMaxPods *bool `field:"optional" json:"useMaxPods" yaml:"useMaxPods"`
}

// Capacity type of the managed node group.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng-spot"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("c5.large")),
//   		ec2.NewInstanceType(jsii.String("c5a.large")),
//   		ec2.NewInstanceType(jsii.String("c5d.large")),
//   	},
//   	minSize: jsii.Number(3),
//   	capacityType: eks.capacityType_SPOT,
//   })
//
// Experimental.
type CapacityType string

const (
	// spot instances.
	// Experimental.
	CapacityType_SPOT CapacityType = "SPOT"
	// on-demand instances.
	// Experimental.
	CapacityType_ON_DEMAND CapacityType = "ON_DEMAND"
)

// A CloudFormation `AWS::EKS::Addon`.
//
// Creates an Amazon EKS add-on.
//
// Amazon EKS add-ons help to automate the provisioning and lifecycle management of common operational software for Amazon EKS clusters. Amazon EKS add-ons require clusters running version 1.18 or later because Amazon EKS add-ons rely on the Server-side Apply Kubernetes feature, which is only available in Kubernetes 1.18 and later. For more information, see [Amazon EKS add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAddon := awscdk.Aws_eks.NewCfnAddon(this, jsii.String("MyCfnAddon"), &cfnAddonProps{
//   	addonName: jsii.String("addonName"),
//   	clusterName: jsii.String("clusterName"),
//
//   	// the properties below are optional
//   	addonVersion: jsii.String("addonVersion"),
//   	resolveConflicts: jsii.String("resolveConflicts"),
//   	serviceAccountRoleArn: jsii.String("serviceAccountRoleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAddon interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the add-on.
	AddonName() *string
	SetAddonName(val *string)
	// The version of the add-on.
	AddonVersion() *string
	SetAddonVersion(val *string)
	// The ARN of the add-on, such as `arn:aws:eks:us-west-2:111122223333:addon/1-19/vpc-cni/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The name of the cluster.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.
	ResolveConflicts() *string
	SetResolveConflicts(val *string)
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	ServiceAccountRoleArn() *string
	SetServiceAccountRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
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

func (c *jsiiProxy_CfnAddon) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAddon) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAddon) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAddon) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAddon) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAddon) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnAddon) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAddon) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAddon) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnAddon) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnAddon) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnAddon) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAddon`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAddonProps := &cfnAddonProps{
//   	addonName: jsii.String("addonName"),
//   	clusterName: jsii.String("clusterName"),
//
//   	// the properties below are optional
//   	addonVersion: jsii.String("addonVersion"),
//   	resolveConflicts: jsii.String("resolveConflicts"),
//   	serviceAccountRoleArn: jsii.String("serviceAccountRoleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAddonProps struct {
	// The name of the add-on.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of the add-on.
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.
	ResolveConflicts *string `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_eks.NewCfnCluster(this, jsii.String("MyCfnCluster"), &cfnClusterProps{
//   	resourcesVpcConfig: &resourcesVpcConfigProperty{
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		endpointPrivateAccess: jsii.Boolean(false),
//   		endpointPublicAccess: jsii.Boolean(false),
//   		publicAccessCidrs: []*string{
//   			jsii.String("publicAccessCidrs"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	encryptionConfig: []interface{}{
//   		&encryptionConfigProperty{
//   			provider: &providerProperty{
//   				keyArn: jsii.String("keyArn"),
//   			},
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   	},
//   	kubernetesNetworkConfig: &kubernetesNetworkConfigProperty{
//   		ipFamily: jsii.String("ipFamily"),
//   		serviceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   		serviceIpv6Cidr: jsii.String("serviceIpv6Cidr"),
//   	},
//   	logging: &loggingProperty{
//   		clusterLogging: &clusterLoggingProperty{
//   			enabledTypes: []interface{}{
//   				&loggingTypeConfigProperty{
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	version: jsii.String("version"),
//   })
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the cluster, such as `arn:aws:eks:us-west-2:666666666666:cluster/prod` .
	AttrArn() *string
	// The `certificate-authority-data` for your cluster.
	AttrCertificateAuthorityData() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	//
	// Managed node groups use this security group for control plane to data plane communication.
	//
	// This parameter is only returned by Amazon EKS clusters that support managed node groups. For more information, see [Managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in the *Amazon EKS User Guide* .
	AttrClusterSecurityGroupId() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	AttrEncryptionConfigKeyArn() *string
	// The endpoint for your Kubernetes API server, such as `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com` .
	AttrEndpoint() *string
	// The CIDR block that Kubernetes Service IP addresses are assigned from if you created a 1.21 or later cluster with version 1.10.1 or later of the Amazon VPC CNI add-on and specified `ipv6` for *ipFamily* when you created the cluster. Kubernetes assigns Service addresses from the unique local address range ( `fc00::/7` ) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	AttrKubernetesNetworkConfigServiceIpv6Cidr() *string
	// The issuer URL for the OIDC identity provider.
	AttrOpenIdConnectIssuerUrl() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The encryption configuration for the cluster.
	EncryptionConfig() interface{}
	SetEncryptionConfig(val interface{})
	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig() interface{}
	SetKubernetesNetworkConfig(val interface{})
	// The logging configuration for your cluster.
	Logging() interface{}
	SetLogging(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The unique name to give to your cluster.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The VPC configuration that's used by the cluster control plane.
	//
	// Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
	//
	// > Updates require replacement of the `SecurityGroupIds` and `SubnetIds` sub-properties.
	ResourcesVpcConfig() interface{}
	SetResourcesVpcConfig(val interface{})
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// For more information, see [Amazon EKS Service IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the **Amazon EKS User Guide** .
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The metadata that you apply to the cluster to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Cluster tags don't propagate to any other resources associated with the cluster.
	//
	// > You must have the `eks:TagResource` and `eks:UntagResource` permissions in your IAM user or IAM role used to manage the CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version() *string
	SetVersion(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
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

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterLoggingProperty := &clusterLoggingProperty{
//   	enabledTypes: []interface{}{
//   		&loggingTypeConfigProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnCluster_ClusterLoggingProperty struct {
	// The enabled control plane logs for your cluster. All log types are disabled if the array is empty.
	//
	// > When updating a resource, you must include this `EnabledTypes` property if the previous CloudFormation template of the resource had it.
	EnabledTypes interface{} `field:"optional" json:"enabledTypes" yaml:"enabledTypes"`
}

// The encryption configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &encryptionConfigProperty{
//   	provider: &providerProperty{
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
type CfnCluster_EncryptionConfigProperty struct {
	// The encryption provider for the cluster.
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
	// Specifies the resources to be encrypted.
	//
	// The only supported value is "secrets".
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

// The Kubernetes network configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubernetesNetworkConfigProperty := &kubernetesNetworkConfigProperty{
//   	ipFamily: jsii.String("ipFamily"),
//   	serviceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   	serviceIpv6Cidr: jsii.String("serviceIpv6Cidr"),
//   }
//
type CfnCluster_KubernetesNetworkConfigProperty struct {
	// Specify which IP family is used to assign Kubernetes pod and service IP addresses.
	//
	// If you don't specify a value, `ipv4` is used by default. You can only specify an IP family when you create a cluster and can't change this value once the cluster is created. If you specify `ipv6` , the VPC and subnets that you specify for cluster creation must have both IPv4 and IPv6 CIDR blocks assigned to them. You can't specify `ipv6` for clusters in China Regions.
	//
	// You can only specify `ipv6` for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on. If you specify `ipv6` , then ensure that your VPC meets the requirements listed in the considerations listed in [Assigning IPv6 addresses to pods and services](https://docs.aws.amazon.com/eks/latest/userguide/cni-ipv6.html) in the Amazon EKS User Guide. Kubernetes assigns services IPv6 addresses from the unique local address range (fc00::/7). You can't specify a custom IPv6 CIDR block. Pod addresses are assigned from the subnet's IPv6 CIDR.
	IpFamily *string `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Don't specify a value if you select `ipv6` for *ipFamily* .
	//
	// The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. The block must meet the following requirements:
	//
	// - Within one of the following private IP address blocks: 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16.
	// - Doesn't overlap with any CIDR block assigned to the VPC that you selected for VPC.
	// - Between /24 and /12.
	//
	// > You can only specify a custom CIDR block when you create a cluster and can't change this value once the cluster is created.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// The CIDR block that Kubernetes pod and service IP addresses are assigned from if you created a 1.21 or later cluster with version 1.10.1 or later of the Amazon VPC CNI add-on and specified `ipv6` for *ipFamily* when you created the cluster. Kubernetes assigns service addresses from the unique local address range ( `fc00::/7` ) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	ServiceIpv6Cidr *string `field:"optional" json:"serviceIpv6Cidr" yaml:"serviceIpv6Cidr"`
}

// Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs.
//
// By default, cluster control plane logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS Cluster control plane logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the **Amazon EKS User Guide** .
//
// > When updating a resource, you must include this `Logging` property if the previous CloudFormation template of the resource had it. > CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported control plane logs. For more information, see [CloudWatch Pricing](https://docs.aws.amazon.com/cloudwatch/pricing/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	clusterLogging: &clusterLoggingProperty{
//   		enabledTypes: []interface{}{
//   			&loggingTypeConfigProperty{
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnCluster_LoggingProperty struct {
	// The cluster control plane logging configuration for your cluster.
	ClusterLogging interface{} `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
}

// The enabled logging type.
//
// For a list of the valid logging types, see the [`types` property of `LogSetup`](https://docs.aws.amazon.com/eks/latest/APIReference/API_LogSetup.html#AmazonEKS-Type-LogSetup-types) in the *Amazon EKS API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingTypeConfigProperty := &loggingTypeConfigProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnCluster_LoggingTypeConfigProperty struct {
	// The name of the log type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Identifies the AWS Key Management Service ( AWS KMS ) key used to encrypt the secrets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &providerProperty{
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnCluster_ProviderProperty struct {
	// Amazon Resource Name (ARN) or alias of the KMS key.
	//
	// The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key. For more information, see [Allowing Users in Other Accounts to Use a KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html) in the *AWS Key Management Service Developer Guide* .
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

// An object representing the VPC configuration to use for an Amazon EKS cluster.
//
// > When updating a resource, you must include these properties if the previous CloudFormation template of the resource had them:
// >
// > - `EndpointPublicAccess`
// > - `EndpointPrivateAccess`
// > - `PublicAccessCidrs`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcesVpcConfigProperty := &resourcesVpcConfigProperty{
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	endpointPrivateAccess: jsii.Boolean(false),
//   	endpointPublicAccess: jsii.Boolean(false),
//   	publicAccessCidrs: []*string{
//   		jsii.String("publicAccessCidrs"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type CfnCluster_ResourcesVpcConfigProperty struct {
	// Specify subnets for your Amazon EKS nodes.
	//
	// Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Set this value to `true` to enable private access for your cluster's Kubernetes API server endpoint.
	//
	// If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is `false` , which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that `publicAccessCidrs` includes the necessary CIDR blocks for communication with the nodes or Fargate pods. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	EndpointPrivateAccess interface{} `field:"optional" json:"endpointPrivateAccess" yaml:"endpointPrivateAccess"`
	// Set this value to `false` to disable public access to your cluster's Kubernetes API server endpoint.
	//
	// If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is `true` , which enables public access for your Kubernetes API server. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	EndpointPublicAccess interface{} `field:"optional" json:"endpointPublicAccess" yaml:"endpointPublicAccess"`
	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.
	//
	// Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is `0.0.0.0/0` . If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks. For more information, see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the **Amazon EKS User Guide** .
	PublicAccessCidrs *[]*string `field:"optional" json:"publicAccessCidrs" yaml:"publicAccessCidrs"`
	// Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use that allow communication between your nodes and the Kubernetes control plane.
	//
	// If you don't specify any security groups, then familiarize yourself with the difference between Amazon EKS defaults for clusters deployed with Kubernetes:
	//
	// - 1.14 Amazon EKS platform version `eks.2` and earlier
	// - 1.14 Amazon EKS platform version `eks.3` and later
	//
	// For more information, see [Amazon EKS security group considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the **Amazon EKS User Guide** .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	resourcesVpcConfig: &resourcesVpcConfigProperty{
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		endpointPrivateAccess: jsii.Boolean(false),
//   		endpointPublicAccess: jsii.Boolean(false),
//   		publicAccessCidrs: []*string{
//   			jsii.String("publicAccessCidrs"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	encryptionConfig: []interface{}{
//   		&encryptionConfigProperty{
//   			provider: &providerProperty{
//   				keyArn: jsii.String("keyArn"),
//   			},
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   	},
//   	kubernetesNetworkConfig: &kubernetesNetworkConfigProperty{
//   		ipFamily: jsii.String("ipFamily"),
//   		serviceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   		serviceIpv6Cidr: jsii.String("serviceIpv6Cidr"),
//   	},
//   	logging: &loggingProperty{
//   		clusterLogging: &clusterLoggingProperty{
//   			enabledTypes: []interface{}{
//   				&loggingTypeConfigProperty{
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnClusterProps struct {
	// The VPC configuration that's used by the cluster control plane.
	//
	// Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
	//
	// > Updates require replacement of the `SecurityGroupIds` and `SubnetIds` sub-properties.
	ResourcesVpcConfig interface{} `field:"required" json:"resourcesVpcConfig" yaml:"resourcesVpcConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	//
	// For more information, see [Amazon EKS Service IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the **Amazon EKS User Guide** .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The encryption configuration for the cluster.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig interface{} `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// The logging configuration for your cluster.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The unique name to give to your cluster.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The metadata that you apply to the cluster to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Cluster tags don't propagate to any other resources associated with the cluster.
	//
	// > You must have the `eks:TagResource` and `eks:UntagResource` permissions in your IAM user or IAM role used to manage the CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The desired Kubernetes version for your cluster.
	//
	// If you don't specify a value here, the latest version available in Amazon EKS is used.
	Version *string `field:"optional" json:"version" yaml:"version"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFargateProfile := awscdk.Aws_eks.NewCfnFargateProfile(this, jsii.String("MyCfnFargateProfile"), &cfnFargateProfileProps{
//   	clusterName: jsii.String("clusterName"),
//   	podExecutionRoleArn: jsii.String("podExecutionRoleArn"),
//   	selectors: []interface{}{
//   		&selectorProperty{
//   			namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			labels: []interface{}{
//   				&labelProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	fargateProfileName: jsii.String("fargateProfileName"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFargateProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the cluster, such as `arn:aws:eks:us-west-2:666666666666:fargateprofile/myCluster/myFargateProfile/1cb1a11a-1dc1-1d11-cf11-1111f11fa111` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The name of the Fargate profile.
	FargateProfileName() *string
	SetFargateProfileName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see [Pod Execution Role](https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the *Amazon EKS User Guide* .
	PodExecutionRoleArn() *string
	SetPodExecutionRoleArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors() interface{}
	SetSelectors(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The IDs of subnets to launch your pods into.
	//
	// At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// The metadata to apply to the Fargate profile to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
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

func (c *jsiiProxy_CfnFargateProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFargateProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFargateProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFargateProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFargateProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFargateProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFargateProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnFargateProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFargateProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFargateProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnFargateProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnFargateProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnFargateProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A key-value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &labelProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFargateProfile_LabelProperty struct {
	// Enter a key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Enter a value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// An object representing an AWS Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectorProperty := &selectorProperty{
//   	namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	labels: []interface{}{
//   		&labelProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFargateProfile_SelectorProperty struct {
	// The Kubernetes namespace that the selector should match.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain all of the labels that are specified in the selector for it to be considered a match.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

// Properties for defining a `CfnFargateProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFargateProfileProps := &cfnFargateProfileProps{
//   	clusterName: jsii.String("clusterName"),
//   	podExecutionRoleArn: jsii.String("podExecutionRoleArn"),
//   	selectors: []interface{}{
//   		&selectorProperty{
//   			namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			labels: []interface{}{
//   				&labelProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	fargateProfileName: jsii.String("fargateProfileName"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFargateProfileProps struct {
	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see [Pod Execution Role](https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the *Amazon EKS User Guide* .
	PodExecutionRoleArn *string `field:"required" json:"podExecutionRoleArn" yaml:"podExecutionRoleArn"`
	// The selectors to match for pods to use this Fargate profile.
	//
	// Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	FargateProfileName *string `field:"optional" json:"fargateProfileName" yaml:"fargateProfileName"`
	// The IDs of subnets to launch your pods into.
	//
	// At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the Fargate profile to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::EKS::IdentityProviderConfig`.
//
// Associate an identity provider configuration to a cluster.
//
// If you want to authenticate identities using an identity provider, you can create an identity provider configuration and associate it to your cluster. After configuring authentication to your cluster you can create Kubernetes `roles` and `clusterroles` to assign permissions to the roles, and then bind the roles to the identities using Kubernetes `rolebindings` and `clusterrolebindings` . For more information see [Using RBAC Authorization](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/access-authn-authz/rbac/) in the Kubernetes documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProviderConfig := awscdk.Aws_eks.NewCfnIdentityProviderConfig(this, jsii.String("MyCfnIdentityProviderConfig"), &cfnIdentityProviderConfigProps{
//   	clusterName: jsii.String("clusterName"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	identityProviderConfigName: jsii.String("identityProviderConfigName"),
//   	oidc: &oidcIdentityProviderConfigProperty{
//   		clientId: jsii.String("clientId"),
//   		issuerUrl: jsii.String("issuerUrl"),
//
//   		// the properties below are optional
//   		groupsClaim: jsii.String("groupsClaim"),
//   		groupsPrefix: jsii.String("groupsPrefix"),
//   		requiredClaims: []interface{}{
//   			&requiredClaimProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		usernameClaim: jsii.String("usernameClaim"),
//   		usernamePrefix: jsii.String("usernamePrefix"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnIdentityProviderConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) associated with the identity provider config.
	AttrIdentityProviderConfigArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The cluster that the configuration is associated to.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The name of the configuration.
	IdentityProviderConfigName() *string
	SetIdentityProviderConfigName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// An object that represents an OpenID Connect (OIDC) identity provider configuration.
	Oidc() interface{}
	SetOidc(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The metadata to apply to the provider configuration to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both.
	Tags() awscdk.TagManager
	// The type of the identity provider configuration.
	//
	// The only type available is `oidc` .
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnIdentityProviderConfig
type jsiiProxy_CfnIdentityProviderConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIdentityProviderConfig) AttrIdentityProviderConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityProviderConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) IdentityProviderConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Oidc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EKS::IdentityProviderConfig`.
func NewCfnIdentityProviderConfig(scope awscdk.Construct, id *string, props *CfnIdentityProviderConfigProps) CfnIdentityProviderConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnIdentityProviderConfig{}

	_jsii_.Create(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EKS::IdentityProviderConfig`.
func NewCfnIdentityProviderConfig_Override(c CfnIdentityProviderConfig, scope awscdk.Construct, id *string, props *CfnIdentityProviderConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIdentityProviderConfig) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityProviderConfig) SetIdentityProviderConfigName(val *string) {
	_jsii_.Set(
		j,
		"identityProviderConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityProviderConfig) SetOidc(val interface{}) {
	_jsii_.Set(
		j,
		"oidc",
		val,
	)
}

func (j *jsiiProxy_CfnIdentityProviderConfig) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnIdentityProviderConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIdentityProviderConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnIdentityProviderConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityProviderConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.CfnIdentityProviderConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIdentityProviderConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that represents the configuration for an OpenID Connect (OIDC) identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcIdentityProviderConfigProperty := &oidcIdentityProviderConfigProperty{
//   	clientId: jsii.String("clientId"),
//   	issuerUrl: jsii.String("issuerUrl"),
//
//   	// the properties below are optional
//   	groupsClaim: jsii.String("groupsClaim"),
//   	groupsPrefix: jsii.String("groupsPrefix"),
//   	requiredClaims: []interface{}{
//   		&requiredClaimProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	usernameClaim: jsii.String("usernameClaim"),
//   	usernamePrefix: jsii.String("usernamePrefix"),
//   }
//
type CfnIdentityProviderConfig_OidcIdentityProviderConfigProperty struct {
	// This is also known as *audience* .
	//
	// The ID of the client application that makes authentication requests to the OIDC identity provider.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The URL of the OIDC identity provider that allows the API server to discover public signing keys for verifying tokens.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The JSON web token (JWT) claim that the provider uses to return your groups.
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// The prefix that is prepended to group claims to prevent clashes with existing names (such as `system:` groups).
	//
	// For example, the value `oidc:` creates group names like `oidc:engineering` and `oidc:infra` . The prefix can't contain `system:`
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// The key-value pairs that describe required claims in the identity token.
	//
	// If set, each claim is verified to be present in the token with a matching value.
	RequiredClaims interface{} `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// The JSON Web token (JWT) claim that is used as the username.
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// The prefix that is prepended to username claims to prevent clashes with existing names.
	//
	// The prefix can't contain `system:`.
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

// A key-value pair that describes a required claim in the identity token.
//
// If set, each claim is verified to be present in the token with a matching value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requiredClaimProperty := &requiredClaimProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnIdentityProviderConfig_RequiredClaimProperty struct {
	// The key to match from the token.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the key from the token.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Properties for defining a `CfnIdentityProviderConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProviderConfigProps := &cfnIdentityProviderConfigProps{
//   	clusterName: jsii.String("clusterName"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	identityProviderConfigName: jsii.String("identityProviderConfigName"),
//   	oidc: &oidcIdentityProviderConfigProperty{
//   		clientId: jsii.String("clientId"),
//   		issuerUrl: jsii.String("issuerUrl"),
//
//   		// the properties below are optional
//   		groupsClaim: jsii.String("groupsClaim"),
//   		groupsPrefix: jsii.String("groupsPrefix"),
//   		requiredClaims: []interface{}{
//   			&requiredClaimProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		usernameClaim: jsii.String("usernameClaim"),
//   		usernamePrefix: jsii.String("usernamePrefix"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityProviderConfigProps struct {
	// The cluster that the configuration is associated to.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The type of the identity provider configuration.
	//
	// The only type available is `oidc` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the configuration.
	IdentityProviderConfigName *string `field:"optional" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// An object that represents an OpenID Connect (OIDC) identity provider configuration.
	Oidc interface{} `field:"optional" json:"oidc" yaml:"oidc"`
	// The metadata to apply to the provider configuration to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::EKS::Nodegroup`.
//
// Creates a managed node group for an Amazon EKS cluster. You can only create a node group for your cluster that is equal to the current Kubernetes version for the cluster. All node groups are created with the latest AMI release version for the respective minor Kubernetes version of the cluster, unless you deploy a custom AMI using a launch template. For more information about using launch templates, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) .
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and associated Amazon EC2 instances that are managed by AWS for an Amazon EKS cluster. Each node group uses a version of the Amazon EKS optimized Amazon Linux 2 AMI. For more information, see [Managed Node Groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
//   var tags interface{}
//
//   cfnNodegroup := awscdk.Aws_eks.NewCfnNodegroup(this, jsii.String("MyCfnNodegroup"), &cfnNodegroupProps{
//   	clusterName: jsii.String("clusterName"),
//   	nodeRole: jsii.String("nodeRole"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	amiType: jsii.String("amiType"),
//   	capacityType: jsii.String("capacityType"),
//   	diskSize: jsii.Number(123),
//   	forceUpdateEnabled: jsii.Boolean(false),
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	labels: labels,
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	nodegroupName: jsii.String("nodegroupName"),
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &remoteAccessProperty{
//   		ec2SshKey: jsii.String("ec2SshKey"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []*string{
//   			jsii.String("sourceSecurityGroups"),
//   		},
//   	},
//   	scalingConfig: &scalingConfigProperty{
//   		desiredSize: jsii.Number(123),
//   		maxSize: jsii.Number(123),
//   		minSize: jsii.Number(123),
//   	},
//   	tags: tags,
//   	taints: []interface{}{
//   		&taintProperty{
//   			effect: jsii.String("effect"),
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	updateConfig: &updateConfigProperty{
//   		maxUnavailable: jsii.Number(123),
//   		maxUnavailablePercentage: jsii.Number(123),
//   	},
//   	version: jsii.String("version"),
//   })
//
type CfnNodegroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The AMI type for your node group.
	//
	// GPU instance types should use the `AL2_x86_64_GPU` AMI type. Non-GPU instances should use the `AL2_x86_64` AMI type. Arm instances should use the `AL2_ARM_64` AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `amiType` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	AmiType() *string
	SetAmiType(val *string)
	// The Amazon Resource Name (ARN) associated with the managed node group.
	AttrArn() *string
	// The name of the cluster that the managed node group resides in.
	AttrClusterName() *string
	AttrId() *string
	// The name associated with an Amazon EKS managed node group.
	AttrNodegroupName() *string
	// The capacity type of your managed node group.
	CapacityType() *string
	SetCapacityType(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The name of the cluster to create the node group in.
	ClusterName() *string
	SetClusterName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The root device disk size (in GiB) for your node group instances.
	//
	// The default disk size is 20 GiB. If you specify `launchTemplate` , then don't specify `diskSize` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	DiskSize() *float64
	SetDiskSize(val *float64)
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old node whether or not any pods are running on the node.
	ForceUpdateEnabled() interface{}
	SetForceUpdateEnabled(val interface{})
	// Specify the instance types for a node group.
	//
	// If you specify a GPU instance type, be sure to specify `AL2_x86_64_GPU` with the `amiType` parameter. If you specify `launchTemplate` , then you can specify zero or one instance type in your launch template *or* you can specify 0-20 instance types for `instanceTypes` . If however, you specify an instance type in your launch template *and* specify any `instanceTypes` , the node group deployment will fail. If you don't specify an instance type in a launch template or for `instanceTypes` , then `t3.medium` is used, by default. If you specify `Spot` for `capacityType` , then we recommend specifying multiple values for `instanceTypes` . For more information, see [Managed node group capacity types](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels() interface{}
	SetLabels(val interface{})
	// An object representing a node group's launch template specification.
	//
	// If specified, then do not specify `instanceTypes` , `diskSize` , or `remoteAccess` and make sure that the launch template meets the requirements in `launchTemplateSpecification` .
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The unique name to give your node group.
	NodegroupName() *string
	SetNodegroupName(val *string)
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// The Amazon EKS worker node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the **Amazon EKS User Guide** . If you specify `launchTemplate` , then don't specify [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	NodeRole() *string
	SetNodeRole(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The AMI version of the Amazon EKS optimized AMI to use with your node group (for example, `1.14.7- *YYYYMMDD*` ). By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see [Amazon EKS optimized Linux AMI Versions](https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the *Amazon EKS User Guide* .
	//
	// > Changing this value triggers an update of the node group if one is available. However, only the latest available AMI release version is valid as an input. You cannot roll back to a previous AMI release version.
	ReleaseVersion() *string
	SetReleaseVersion(val *string)
	// The remote access (SSH) configuration to use with your node group.
	//
	// If you specify `launchTemplate` , then don't specify `remoteAccess` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	RemoteAccess() interface{}
	SetRemoteAccess(val interface{})
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig() interface{}
	SetScalingConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// If you specify `launchTemplate` , then don't specify [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.
	Tags() awscdk.TagManager
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Effect is one of `No_Schedule` , `Prefer_No_Schedule` , or `No_Execute` . Kubernetes taints can be used together with tolerations to control how workloads are scheduled to your nodes. For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
	Taints() interface{}
	SetTaints(val interface{})
	// The node group update configuration.
	UpdateConfig() interface{}
	SetUpdateConfig(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The Kubernetes version to use for your managed nodes.
	//
	// By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `version` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Version() *string
	SetVersion(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
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

func (j *jsiiProxy_CfnNodegroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (c *jsiiProxy_CfnNodegroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNodegroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNodegroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNodegroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNodegroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNodegroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnNodegroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNodegroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNodegroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnNodegroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnNodegroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &launchTemplateSpecificationProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnNodegroup_LaunchTemplateSpecificationProperty struct {
	// The ID of the launch template.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the launch template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the launch template to use.
	//
	// If no version is specified, then the template's default version is used.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// An object representing the remote access configuration for the managed node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteAccessProperty := &remoteAccessProperty{
//   	ec2SshKey: jsii.String("ec2SshKey"),
//
//   	// the properties below are optional
//   	sourceSecurityGroups: []*string{
//   		jsii.String("sourceSecurityGroups"),
//   	},
//   }
//
type CfnNodegroup_RemoteAccessProperty struct {
	// The Amazon EC2 SSH key that provides access for SSH communication with the nodes in the managed node group.
	//
	// For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Linux Instances* .
	Ec2SshKey *string `field:"required" json:"ec2SshKey" yaml:"ec2SshKey"`
	// The security groups that are allowed SSH access (port 22) to the nodes.
	//
	// If you specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group, then port 22 on the nodes is opened to the internet (0.0.0.0/0). For more information, see [Security Groups for Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	SourceSecurityGroups *[]*string `field:"optional" json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

// An object representing the scaling configuration details for the Auto Scaling group that is associated with your node group.
//
// When creating a node group, you must specify all or none of the properties. When updating a node group, you can specify any or none of the properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigProperty := &scalingConfigProperty{
//   	desiredSize: jsii.Number(123),
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   }
//
type CfnNodegroup_ScalingConfigProperty struct {
	// The current number of nodes that the managed node group should maintain.
	//
	// > If you use Cluster Autoscaler, you shouldn't change the desiredSize value directly, as this can cause the Cluster Autoscaler to suddenly scale up or scale down.
	//
	// Whenever this parameter changes, the number of worker nodes in the node group is updated to the specified size. If this parameter is given a value that is smaller than the current number of running worker nodes, the necessary number of worker nodes are terminated to match the given value. When using CloudFormation, no action occurs if you remove this parameter from your CFN template.
	//
	// This parameter can be different from minSize in some cases, such as when starting with extra hosts for testing. This parameter can also be different when you want to start with an estimated number of needed hosts, but let Cluster Autoscaler reduce the number if there are too many. When Cluster Autoscaler is used, the desiredSize parameter is altered by Cluster Autoscaler (but can be out-of-date for short periods of time). Cluster Autoscaler doesn't scale a managed node group lower than minSize or higher than maxSize.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The maximum number of nodes that the managed node group can scale out to.
	//
	// For information about the maximum number that you can specify, see [Amazon EKS service quotas](https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the *Amazon EKS User Guide* .
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of nodes that the managed node group can scale in to.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

// A property that allows a node to repel a set of pods.
//
// For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taintProperty := &taintProperty{
//   	effect: jsii.String("effect"),
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnNodegroup_TaintProperty struct {
	// The effect of the taint.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The key of the taint.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the taint.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The update configuration for the node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateConfigProperty := &updateConfigProperty{
//   	maxUnavailable: jsii.Number(123),
//   	maxUnavailablePercentage: jsii.Number(123),
//   }
//
type CfnNodegroup_UpdateConfigProperty struct {
	// The maximum number of nodes unavailable at once during a version update.
	//
	// Nodes will be updated in parallel. This value or `maxUnavailablePercentage` is required to have a value.The maximum number is 100.
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The maximum percentage of nodes unavailable during a version update.
	//
	// This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or `maxUnavailable` is required to have a value.
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
}

// Properties for defining a `CfnNodegroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
//   var tags interface{}
//
//   cfnNodegroupProps := &cfnNodegroupProps{
//   	clusterName: jsii.String("clusterName"),
//   	nodeRole: jsii.String("nodeRole"),
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	amiType: jsii.String("amiType"),
//   	capacityType: jsii.String("capacityType"),
//   	diskSize: jsii.Number(123),
//   	forceUpdateEnabled: jsii.Boolean(false),
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	labels: labels,
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		id: jsii.String("id"),
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	nodegroupName: jsii.String("nodegroupName"),
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &remoteAccessProperty{
//   		ec2SshKey: jsii.String("ec2SshKey"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []*string{
//   			jsii.String("sourceSecurityGroups"),
//   		},
//   	},
//   	scalingConfig: &scalingConfigProperty{
//   		desiredSize: jsii.Number(123),
//   		maxSize: jsii.Number(123),
//   		minSize: jsii.Number(123),
//   	},
//   	tags: tags,
//   	taints: []interface{}{
//   		&taintProperty{
//   			effect: jsii.String("effect"),
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	updateConfig: &updateConfigProperty{
//   		maxUnavailable: jsii.Number(123),
//   		maxUnavailablePercentage: jsii.Number(123),
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnNodegroupProps struct {
	// The name of the cluster to create the node group in.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// The Amazon EKS worker node `kubelet` daemon makes calls to AWS APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the **Amazon EKS User Guide** . If you specify `launchTemplate` , then don't specify [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	NodeRole *string `field:"required" json:"nodeRole" yaml:"nodeRole"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// If you specify `launchTemplate` , then don't specify [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The AMI type for your node group.
	//
	// GPU instance types should use the `AL2_x86_64_GPU` AMI type. Non-GPU instances should use the `AL2_x86_64` AMI type. Arm instances should use the `AL2_ARM_64` AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `amiType` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of your managed node group.
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The root device disk size (in GiB) for your node group instances.
	//
	// The default disk size is 20 GiB. If you specify `launchTemplate` , then don't specify `diskSize` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old node whether or not any pods are running on the node.
	ForceUpdateEnabled interface{} `field:"optional" json:"forceUpdateEnabled" yaml:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	//
	// If you specify a GPU instance type, be sure to specify `AL2_x86_64_GPU` with the `amiType` parameter. If you specify `launchTemplate` , then you can specify zero or one instance type in your launch template *or* you can specify 0-20 instance types for `instanceTypes` . If however, you specify an instance type in your launch template *and* specify any `instanceTypes` , the node group deployment will fail. If you don't specify an instance type in a launch template or for `instanceTypes` , then `t3.medium` is used, by default. If you specify `Spot` for `capacityType` , then we recommend specifying multiple values for `instanceTypes` . For more information, see [Managed node group capacity types](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// An object representing a node group's launch template specification.
	//
	// If specified, then do not specify `instanceTypes` , `diskSize` , or `remoteAccess` and make sure that the launch template meets the requirements in `launchTemplateSpecification` .
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The unique name to give your node group.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The AMI version of the Amazon EKS optimized AMI to use with your node group (for example, `1.14.7- *YYYYMMDD*` ). By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see [Amazon EKS optimized Linux AMI Versions](https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the *Amazon EKS User Guide* .
	//
	// > Changing this value triggers an update of the node group if one is available. However, only the latest available AMI release version is valid as an input. You cannot roll back to a previous AMI release version.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// If you specify `launchTemplate` , then don't specify `remoteAccess` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	RemoteAccess interface{} `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Effect is one of `No_Schedule` , `Prefer_No_Schedule` , or `No_Execute` . Kubernetes taints can be used together with tolerations to control how workloads are scheduled to your nodes. For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The node group update configuration.
	UpdateConfig interface{} `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	//
	// By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify `launchTemplate` , and your launch template uses a custom AMI, then don't specify `version` , or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// A Cluster represents a managed Kubernetes Service (EKS).
//
// This is a fully managed cluster of API Servers (control-plane)
// The user is still required to create the worker nodes.
//
// Example:
//   var vpc vpc
//
//
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	vpc: vpc,
//   	vpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE,
//   		},
//   	},
//   })
//
// Experimental.
type Cluster interface {
	awscdk.Resource
	ICluster
	// An IAM role with administrative permissions to create or update the cluster.
	//
	// This role also has `systems:master` permissions.
	// Experimental.
	AdminRole() awsiam.Role
	// The ALB Controller construct defined for this cluster.
	//
	// Will be undefined if `albController` wasn't configured.
	// Experimental.
	AlbController() AlbController
	// Lazily creates the AwsAuth resource, which manages AWS authentication mapping.
	// Experimental.
	AwsAuth() AwsAuth
	// The AWS generated ARN for the Cluster resource.
	//
	// For example, `arn:aws:eks:us-west-2:666666666666:cluster/prod`.
	// Experimental.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn() *string
	// The endpoint URL for the Cluster.
	//
	// This is the URL inside the kubeconfig file to use with kubectl
	//
	// For example, `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com`
	// Experimental.
	ClusterEndpoint() *string
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	// The Name of the created EKS Cluster.
	// Experimental.
	ClusterName() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	// Experimental.
	ClusterOpenIdConnectIssuer() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer url.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	// Experimental.
	ClusterOpenIdConnectIssuerUrl() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId() *string
	// Manages connection rules (Security Group Rules) for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The auto scaling group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is not `EC2` or
	// `defaultCapacityType` is `EC2` but default capacity is set to 0.
	// Experimental.
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	// The node group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is `EC2` or
	// `defaultCapacityType` is `NODEGROUP` but default capacity is set to 0.
	// Experimental.
	DefaultNodegroup() Nodegroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
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
	// The AWS Lambda layer that contains `kubectl`, `helm` and the AWS CLI.
	//
	// If
	// undefined, a SAR app that contains this layer will be used.
	// Experimental.
	KubectlLayer() awslambda.ILayerVersion
	// The amount of memory allocated to the kubectl provider's lambda function.
	// Experimental.
	KubectlMemory() awscdk.Size
	// Subnets to host the `kubectl` compute resources.
	// Experimental.
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	// Experimental.
	KubectlRole() awsiam.IRole
	// A security group to use for `kubectl` execution.
	// Experimental.
	KubectlSecurityGroup() awsec2.ISecurityGroup
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The AWS Lambda layer that contains the NPM dependency `proxy-agent`.
	//
	// If
	// undefined, a SAR app that contains this layer will be used.
	// Experimental.
	OnEventLayer() awslambda.ILayerVersion
	// An `OpenIdConnectProvider` resource associated with this cluster, and which can be used to link this cluster to AWS IAM.
	//
	// A provider will only be defined if this property is accessed (lazy initialization).
	// Experimental.
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Determines if Kubernetes resources can be pruned automatically.
	// Experimental.
	Prune() *bool
	// IAM role assumed by the EKS Control Plane.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc() awsec2.IVpc
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
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	// Experimental.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Adds a Fargate profile to this cluster.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
	//
	// Experimental.
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct.
	// Experimental.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesResource` object.
	// Experimental.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Add managed nodegroup to this Amazon EKS cluster.
	//
	// This method will create a new managed nodegroup and add into the capacity.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
	//
	// Experimental.
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	// Creates a new service account with corresponding IAM Role (IRSA).
	// Experimental.
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
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
	// Experimental.
	GeneratePhysicalName() *string
	// Fetch the load balancer address of an ingress backed by a load balancer.
	// Experimental.
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Fetch the load balancer address of a service of type 'LoadBalancer'.
	// Experimental.
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
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

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Cluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		c,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

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

func (c *jsiiProxy_Cluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_Cluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//   var asg autoScalingGroup
//
//   importedCluster := eks.cluster.fromClusterAttributes(this, jsii.String("ImportedCluster"), &clusterAttributes{
//   	clusterName: cluster.clusterName,
//   	clusterSecurityGroupId: cluster.clusterSecurityGroupId,
//   })
//
//   importedCluster.connectAutoScalingGroupCapacity(asg, &autoScalingGroupOptions{
//   })
//
// Experimental.
type ClusterAttributes struct {
	// The physical name of the Cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData *string `field:"optional" json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn *string `field:"optional" json:"clusterEncryptionConfigKeyArn" yaml:"clusterEncryptionConfigKeyArn"`
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// A security group id to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	// Experimental.
	ClusterHandlerSecurityGroupId *string `field:"optional" json:"clusterHandlerSecurityGroupId" yaml:"clusterHandlerSecurityGroupId"`
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId *string `field:"optional" json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
	// Environment variables to use when running `kubectl` against this cluster.
	// Experimental.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands
	// to the cluster.
	// Experimental.
	KubectlLambdaRole awsiam.IRole `field:"optional" json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// This layer
	// is used by the kubectl handler to apply manifests and install helm charts.
	//
	// The handler expects the layer to include the following executables:
	//
	//     helm/helm
	//     kubectl/kubectl
	// awscli/aws.
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// Subnets to host the `kubectl` compute resources.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlPrivateSubnetIds *[]*string `field:"optional" json:"kubectlPrivateSubnetIds" yaml:"kubectlPrivateSubnetIds"`
	// KubectlProvider for issuing kubectl commands.
	// Experimental.
	KubectlProvider IKubectlProvider `field:"optional" json:"kubectlProvider" yaml:"kubectlProvider"`
	// An IAM role with cluster administrator and "system:masters" permissions.
	// Experimental.
	KubectlRoleArn *string `field:"optional" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
	// A security group to use for `kubectl` execution.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlSecurityGroupId *string `field:"optional" json:"kubectlSecurityGroupId" yaml:"kubectlSecurityGroupId"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// The handler expects the layer to include the following node_modules:
	//
	// proxy-agent.
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// An Open ID Connect provider for this cluster that can be used to configure service accounts.
	//
	// You can either import an existing provider using `iam.OpenIdConnectProvider.fromProviderArn`,
	// or create a new provider using `new eks.OpenIdConnectProvider`
	// Experimental.
	OpenIdConnectProvider awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProvider" yaml:"openIdConnectProvider"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Additional security groups associated with this cluster.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// EKS cluster logging types.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	// ...
//   	version: eks.kubernetesVersion_V1_21(),
//   	clusterLogging: []clusterLoggingTypes{
//   		eks.*clusterLoggingTypes_API,
//   		eks.*clusterLoggingTypes_AUTHENTICATOR,
//   		eks.*clusterLoggingTypes_SCHEDULER,
//   	},
//   })
//
// Experimental.
type ClusterLoggingTypes string

const (
	// Logs pertaining to API requests to the cluster.
	// Experimental.
	ClusterLoggingTypes_API ClusterLoggingTypes = "API"
	// Logs pertaining to cluster access via the Kubernetes API.
	// Experimental.
	ClusterLoggingTypes_AUDIT ClusterLoggingTypes = "AUDIT"
	// Logs pertaining to authentication requests into the cluster.
	// Experimental.
	ClusterLoggingTypes_AUTHENTICATOR ClusterLoggingTypes = "AUTHENTICATOR"
	// Logs pertaining to state of cluster controllers.
	// Experimental.
	ClusterLoggingTypes_CONTROLLER_MANAGER ClusterLoggingTypes = "CONTROLLER_MANAGER"
	// Logs pertaining to scheduling decisions.
	// Experimental.
	ClusterLoggingTypes_SCHEDULER ClusterLoggingTypes = "SCHEDULER"
)

// Options for EKS clusters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion albControllerVersion
//   var endpointAccess endpointAccess
//   var key key
//   var kubernetesVersion kubernetesVersion
//   var layerVersion layerVersion
//   var policy interface{}
//   var role role
//   var securityGroup securityGroup
//   var size size
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   clusterOptions := &clusterOptions{
//   	version: kubernetesVersion,
//
//   	// the properties below are optional
//   	albController: &albControllerOptions{
//   		version: albControllerVersion,
//
//   		// the properties below are optional
//   		policy: policy,
//   		repository: jsii.String("repository"),
//   	},
//   	clusterHandlerEnvironment: map[string]*string{
//   		"clusterHandlerEnvironmentKey": jsii.String("clusterHandlerEnvironment"),
//   	},
//   	clusterHandlerSecurityGroup: securityGroup,
//   	clusterName: jsii.String("clusterName"),
//   	coreDnsComputeType: awscdk.Aws_eks.coreDnsComputeType_EC2,
//   	endpointAccess: endpointAccess,
//   	kubectlEnvironment: map[string]*string{
//   		"kubectlEnvironmentKey": jsii.String("kubectlEnvironment"),
//   	},
//   	kubectlLayer: layerVersion,
//   	kubectlMemory: size,
//   	mastersRole: role,
//   	onEventLayer: layerVersion,
//   	outputClusterName: jsii.Boolean(false),
//   	outputConfigCommand: jsii.Boolean(false),
//   	outputMastersRoleArn: jsii.Boolean(false),
//   	placeClusterHandlerInVpc: jsii.Boolean(false),
//   	prune: jsii.Boolean(false),
//   	role: role,
//   	secretsEncryptionKey: key,
//   	securityGroup: securityGroup,
//   	serviceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   	vpc: vpc,
//   	vpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnetName: jsii.String("subnetName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   		},
//   	},
//   }
//
// Experimental.
type ClusterOptions struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `field:"optional" json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
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
	// ```.
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
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
	// ```.
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `field:"optional" json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

// Common configuration props for EKS clusters.
//
// Example:
//   var vpc vpc
//
//
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	vpc: vpc,
//   	vpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE,
//   		},
//   	},
//   })
//
// Experimental.
type ClusterProps struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `field:"optional" json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
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
	// ```.
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
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
	// ```.
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `field:"optional" json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// The cluster log types which you want to enable.
	// Experimental.
	ClusterLogging *[]ClusterLoggingTypes `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
	// Number of instances to allocate as an initial capacity for this cluster.
	//
	// Instance type can be configured through `defaultCapacityInstanceType`,
	// which defaults to `m5.large`.
	//
	// Use `cluster.addAutoScalingGroupCapacity` to add additional customized capacity. Set this
	// to `0` is you wish to avoid the initial capacity allocation.
	// Experimental.
	DefaultCapacity *float64 `field:"optional" json:"defaultCapacity" yaml:"defaultCapacity"`
	// The instance type to use for the default capacity.
	//
	// This will only be taken
	// into account if `defaultCapacity` is > 0.
	// Experimental.
	DefaultCapacityInstance awsec2.InstanceType `field:"optional" json:"defaultCapacityInstance" yaml:"defaultCapacityInstance"`
	// The default capacity type for the cluster.
	// Experimental.
	DefaultCapacityType DefaultCapacityType `field:"optional" json:"defaultCapacityType" yaml:"defaultCapacityType"`
	// The IAM role to pass to the Kubectl Lambda Handler.
	// Experimental.
	KubectlLambdaRole awsiam.IRole `field:"optional" json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
	// The tags assigned to the EKS cluster.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// Options for configuring an EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var kubernetesVersion kubernetesVersion
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   commonClusterOptions := &commonClusterOptions{
//   	version: kubernetesVersion,
//
//   	// the properties below are optional
//   	clusterName: jsii.String("clusterName"),
//   	outputClusterName: jsii.Boolean(false),
//   	outputConfigCommand: jsii.Boolean(false),
//   	role: role,
//   	securityGroup: securityGroup,
//   	vpc: vpc,
//   	vpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnetName: jsii.String("subnetName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   		},
//   	},
//   }
//
// Experimental.
type CommonClusterOptions struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// The type of compute resources to use for CoreDNS.
// Experimental.
type CoreDnsComputeType string

const (
	// Deploy CoreDNS on EC2 instances.
	// Experimental.
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	// Deploy CoreDNS on Fargate-managed instances.
	// Experimental.
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

// CPU architecture.
// Experimental.
type CpuArch string

const (
	// arm64 CPU type.
	// Experimental.
	CpuArch_ARM_64 CpuArch = "ARM_64"
	// x86_64 CPU type.
	// Experimental.
	CpuArch_X86_64 CpuArch = "X86_64"
)

// The default capacity type for the cluster.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacityType: eks.defaultCapacityType_EC2,
//   })
//
// Experimental.
type DefaultCapacityType string

const (
	// managed node group.
	// Experimental.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	// Experimental.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

// Construct an Amazon Linux 2 image from the latest EKS Optimized AMI published in SSM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImage := awscdk.Aws_eks.NewEksOptimizedImage(&eksOptimizedImageProps{
//   	cpuArch: awscdk.*Aws_eks.cpuArch_ARM_64,
//   	kubernetesVersion: jsii.String("kubernetesVersion"),
//   	nodeType: awscdk.*Aws_eks.nodeType_STANDARD,
//   })
//
// Experimental.
type EksOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Experimental.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImageProps := &eksOptimizedImageProps{
//   	cpuArch: awscdk.Aws_eks.cpuArch_ARM_64,
//   	kubernetesVersion: jsii.String("kubernetesVersion"),
//   	nodeType: awscdk.*Aws_eks.nodeType_STANDARD,
//   }
//
// Experimental.
type EksOptimizedImageProps struct {
	// What cpu architecture to retrieve the image for (arm64 or x86_64).
	// Experimental.
	CpuArch CpuArch `field:"optional" json:"cpuArch" yaml:"cpuArch"`
	// The Kubernetes version to use.
	// Experimental.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// What instance type to retrieve the image for (standard or GPU-optimized).
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
}

// Endpoint access characteristics.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	endpointAccess: eks.endpointAccess_PRIVATE(),
//   })
//
// Experimental.
type EndpointAccess interface {
	// Restrict public access to specific CIDR blocks.
	//
	// If public access is disabled, this method will result in an error.
	// Experimental.
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
// Example:
//   cluster := eks.NewFargateCluster(this, jsii.String("MyCluster"), &fargateClusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   })
//
// Experimental.
type FargateCluster interface {
	Cluster
	// An IAM role with administrative permissions to create or update the cluster.
	//
	// This role also has `systems:master` permissions.
	// Experimental.
	AdminRole() awsiam.Role
	// The ALB Controller construct defined for this cluster.
	//
	// Will be undefined if `albController` wasn't configured.
	// Experimental.
	AlbController() AlbController
	// Lazily creates the AwsAuth resource, which manages AWS authentication mapping.
	// Experimental.
	AwsAuth() AwsAuth
	// The AWS generated ARN for the Cluster resource.
	//
	// For example, `arn:aws:eks:us-west-2:666666666666:cluster/prod`.
	// Experimental.
	ClusterArn() *string
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData() *string
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn() *string
	// The endpoint URL for the Cluster.
	//
	// This is the URL inside the kubeconfig file to use with kubectl
	//
	// For example, `https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com`
	// Experimental.
	ClusterEndpoint() *string
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup() awsec2.ISecurityGroup
	// The Name of the created EKS Cluster.
	// Experimental.
	ClusterName() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	// Experimental.
	ClusterOpenIdConnectIssuer() *string
	// If this cluster is kubectl-enabled, returns the OpenID Connect issuer url.
	//
	// This is because the values is only be retrieved by the API and not exposed
	// by CloudFormation. If this cluster is not kubectl-enabled (i.e. uses the
	// stock `CfnCluster`), this is `undefined`.
	// Experimental.
	ClusterOpenIdConnectIssuerUrl() *string
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroup() awsec2.ISecurityGroup
	// The id of the cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId() *string
	// Manages connection rules (Security Group Rules) for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The auto scaling group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is not `EC2` or
	// `defaultCapacityType` is `EC2` but default capacity is set to 0.
	// Experimental.
	DefaultCapacity() awsautoscaling.AutoScalingGroup
	// The node group that hosts the default capacity for this cluster.
	//
	// This will be `undefined` if the `defaultCapacityType` is `EC2` or
	// `defaultCapacityType` is `NODEGROUP` but default capacity is set to 0.
	// Experimental.
	DefaultNodegroup() Nodegroup
	// Fargate Profile that was created with the cluster.
	// Experimental.
	DefaultProfile() FargateProfile
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
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
	// The AWS Lambda layer that contains `kubectl`, `helm` and the AWS CLI.
	//
	// If
	// undefined, a SAR app that contains this layer will be used.
	// Experimental.
	KubectlLayer() awslambda.ILayerVersion
	// The amount of memory allocated to the kubectl provider's lambda function.
	// Experimental.
	KubectlMemory() awscdk.Size
	// Subnets to host the `kubectl` compute resources.
	// Experimental.
	KubectlPrivateSubnets() *[]awsec2.ISubnet
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	// Experimental.
	KubectlRole() awsiam.IRole
	// A security group to use for `kubectl` execution.
	// Experimental.
	KubectlSecurityGroup() awsec2.ISecurityGroup
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The AWS Lambda layer that contains the NPM dependency `proxy-agent`.
	//
	// If
	// undefined, a SAR app that contains this layer will be used.
	// Experimental.
	OnEventLayer() awslambda.ILayerVersion
	// An `OpenIdConnectProvider` resource associated with this cluster, and which can be used to link this cluster to AWS IAM.
	//
	// A provider will only be defined if this property is accessed (lazy initialization).
	// Experimental.
	OpenIdConnectProvider() awsiam.IOpenIdConnectProvider
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Determines if Kubernetes resources can be pruned automatically.
	// Experimental.
	Prune() *bool
	// IAM role assumed by the EKS Control Plane.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc() awsec2.IVpc
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
	AddAutoScalingGroupCapacity(id *string, options *AutoScalingGroupCapacityOptions) awsautoscaling.AutoScalingGroup
	// Defines a CDK8s chart in this cluster.
	//
	// Returns: a `KubernetesManifest` construct representing the chart.
	// Experimental.
	AddCdk8sChart(id *string, chart constructs.Construct, options *KubernetesManifestOptions) KubernetesManifest
	// Adds a Fargate profile to this cluster.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
	//
	// Experimental.
	AddFargateProfile(id *string, options *FargateProfileOptions) FargateProfile
	// Defines a Helm chart in this cluster.
	//
	// Returns: a `HelmChart` construct.
	// Experimental.
	AddHelmChart(id *string, options *HelmChartOptions) HelmChart
	// Defines a Kubernetes resource in this cluster.
	//
	// The manifest will be applied/deleted using kubectl as needed.
	//
	// Returns: a `KubernetesResource` object.
	// Experimental.
	AddManifest(id *string, manifest ...*map[string]interface{}) KubernetesManifest
	// Add managed nodegroup to this Amazon EKS cluster.
	//
	// This method will create a new managed nodegroup and add into the capacity.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
	//
	// Experimental.
	AddNodegroupCapacity(id *string, options *NodegroupOptions) Nodegroup
	// Creates a new service account with corresponding IAM Role (IRSA).
	// Experimental.
	AddServiceAccount(id *string, options *ServiceAccountOptions) ServiceAccount
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
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
	// Experimental.
	GeneratePhysicalName() *string
	// Fetch the load balancer address of an ingress backed by a load balancer.
	// Experimental.
	GetIngressLoadBalancerAddress(ingressName *string, options *IngressLoadBalancerAddressOptions) *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Fetch the load balancer address of a service of type 'LoadBalancer'.
	// Experimental.
	GetServiceLoadBalancerAddress(serviceName *string, options *ServiceLoadBalancerAddressOptions) *string
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

func (f *jsiiProxy_FargateCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FargateCluster) ConnectAutoScalingGroupCapacity(autoScalingGroup awsautoscaling.AutoScalingGroup, options *AutoScalingGroupOptions) {
	_jsii_.InvokeVoid(
		f,
		"connectAutoScalingGroupCapacity",
		[]interface{}{autoScalingGroup, options},
	)
}

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

func (f *jsiiProxy_FargateCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (f *jsiiProxy_FargateCluster) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   cluster := eks.NewFargateCluster(this, jsii.String("MyCluster"), &fargateClusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   })
//
// Experimental.
type FargateClusterProps struct {
	// The Kubernetes version to run in the cluster.
	// Experimental.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	// Experimental.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	// Experimental.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE }]`
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	// Experimental.
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	// Experimental.
	ClusterHandlerEnvironment *map[string]*string `field:"optional" json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	// Experimental.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	// Experimental.
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	// Experimental.
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	// Experimental.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
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
	// ```.
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
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
	// ```.
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	// Experimental.
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	// Experimental.
	PlaceClusterHandlerInVpc *bool `field:"optional" json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	// Experimental.
	SecretsEncryptionKey awskms.IKey `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	// Experimental.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// Fargate Profile to create along with the cluster.
	// Experimental.
	DefaultProfile *FargateProfileOptions `field:"optional" json:"defaultProfile" yaml:"defaultProfile"`
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
// Example:
//   var cluster cluster
//
//   eks.NewFargateProfile(this, jsii.String("MyProfile"), &fargateProfileProps{
//   	cluster: cluster,
//   	selectors: []selector{
//   		&selector{
//   			namespace: jsii.String("default"),
//   		},
//   	},
//   })
//
// Experimental.
type FargateProfile interface {
	awscdk.Construct
	awscdk.ITaggable
	// The full Amazon Resource Name (ARN) of the Fargate profile.
	// Experimental.
	FargateProfileArn() *string
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// Experimental.
	PodExecutionRole() awsiam.IRole
	// Resource tags.
	// Experimental.
	Tags() awscdk.TagManager
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

func (f *jsiiProxy_FargateProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (f *jsiiProxy_FargateProfile) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//
//   cluster.addFargateProfile(jsii.String("MyProfile"), &fargateProfileOptions{
//   	selectors: []selector{
//   		&selector{
//   			namespace: jsii.String("default"),
//   		},
//   	},
//   })
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
	Selectors *[]*Selector `field:"required" json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName *string `field:"optional" json:"fargateProfileName" yaml:"fargateProfileName"`
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
	//
	// Experimental.
	PodExecutionRole awsiam.IRole `field:"optional" json:"podExecutionRole" yaml:"podExecutionRole"`
	// Select which subnets to launch your pods into.
	//
	// At this time, pods running
	// on Fargate are not assigned public IP addresses, so only private subnets
	// (with no direct route to an Internet Gateway) are allowed.
	//
	// You must specify the VPC to customize the subnet selection.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC from which to select subnets to launch your pods into.
	//
	// By default, all private subnets are selected. You can customize this using
	// `subnetSelection`.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Configuration props for EKS Fargate Profiles.
//
// Example:
//   var cluster cluster
//
//   eks.NewFargateProfile(this, jsii.String("MyProfile"), &fargateProfileProps{
//   	cluster: cluster,
//   	selectors: []selector{
//   		&selector{
//   			namespace: jsii.String("default"),
//   		},
//   	},
//   })
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
	Selectors *[]*Selector `field:"required" json:"selectors" yaml:"selectors"`
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName *string `field:"optional" json:"fargateProfileName" yaml:"fargateProfileName"`
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
	//
	// Experimental.
	PodExecutionRole awsiam.IRole `field:"optional" json:"podExecutionRole" yaml:"podExecutionRole"`
	// Select which subnets to launch your pods into.
	//
	// At this time, pods running
	// on Fargate are not assigned public IP addresses, so only private subnets
	// (with no direct route to an Internet Gateway) are allowed.
	//
	// You must specify the VPC to customize the subnet selection.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC from which to select subnets to launch your pods into.
	//
	// By default, all private subnets are selected. You can customize this using
	// `subnetSelection`.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The EKS cluster to apply the Fargate profile to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

// Represents a helm chart within the Kubernetes system.
//
// Applies/deletes the resources using `kubectl` in sync with the resource.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("NginxIngress"), &helmChartProps{
//   	cluster: cluster,
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
//   // or, option2: use `addHelmChart`
//   cluster.addHelmChart(jsii.String("NginxIngress"), &helmChartOptions{
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
// Experimental.
type HelmChart interface {
	awscdk.Construct
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

func (h *jsiiProxy_HelmChart) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmChart) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (h *jsiiProxy_HelmChart) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmChart) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   import s3Assets "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   chartAsset := s3Assets.NewAsset(this, jsii.String("ChartAsset"), &assetProps{
//   	path: jsii.String("/path/to/asset"),
//   })
//
//   cluster.addHelmChart(jsii.String("test-chart"), &helmChartOptions{
//   	chartAsset: chartAsset,
//   })
//
// Experimental.
type HelmChartOptions struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Experimental.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Experimental.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Experimental.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Experimental.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://kubernetes-charts.storage.googleapis.com/
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Experimental.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
}

// Helm Chart properties.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("NginxIngress"), &helmChartProps{
//   	cluster: cluster,
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
//   // or, option2: use `addHelmChart`
//   cluster.addHelmChart(jsii.String("NginxIngress"), &helmChartOptions{
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
// Experimental.
type HelmChartProps struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Experimental.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Experimental.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Experimental.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Experimental.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://kubernetes-charts.storage.googleapis.com/
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Experimental.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
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
	// Returns: a `HelmChart` construct.
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
	// If not defined, a default provider will be used.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   ingressLoadBalancerAddressOptions := &ingressLoadBalancerAddressOptions{
//   	namespace: jsii.String("namespace"),
//   	timeout: duration,
//   }
//
// Experimental.
type IngressLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// Implementation of Kubectl Lambda.
//
// Example:
//   handlerRole := iam.role.fromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   kubectlProvider := eks.kubectlProvider.fromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &kubectlProviderAttributes{
//   	functionArn: jsii.String("arn:aws:lambda:us-east-2:123456789012:function:my-function:1"),
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456789012:role/kubectl-role"),
//   	handlerRole: handlerRole,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("Cluster"), &clusterAttributes{
//   	clusterName: jsii.String("cluster"),
//   	kubectlProvider: kubectlProvider,
//   })
//
// Experimental.
type KubectlProvider interface {
	awscdk.NestedStack
	IKubectlProvider
	// The AWS account into which this stack will be deployed.
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.account` when the stack is defined. This can
	//     either be a concerete account (e.g. `585695031111`) or the
	//     `Aws.accountId` token.
	// 3. `Aws.accountId`, which represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::AccountId" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.account)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **account-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	// Experimental.
	Account() *string
	// The ID of the cloud assembly artifact for this stack.
	// Experimental.
	ArtifactId() *string
	// Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack.
	//
	// If the stack is environment-agnostic (either account and/or region are
	// tokens), this property will return an array with 2 tokens that will resolve
	// at deploy-time to the first two availability zones returned from CloudFormation's
	// `Fn::GetAZs` intrinsic function.
	//
	// If they are not available in the context, returns a set of dummy values and
	// reports them as missing, and let the CLI resolve them by calling EC2
	// `DescribeAvailabilityZones` on the target environment.
	//
	// To specify a different strategy for selecting availability zones override this method.
	// Experimental.
	AvailabilityZones() *[]*string
	// Indicates whether the stack requires bundling or not.
	// Experimental.
	BundlingRequired() *bool
	// Return the stacks this stack depends on.
	// Experimental.
	Dependencies() *[]awscdk.Stack
	// The environment coordinates in which this stack is deployed.
	//
	// In the form
	// `aws://account/region`. Use `stack.account` and `stack.region` to obtain
	// the specific values, no need to parse.
	//
	// You can use this value to determine if two stacks are targeting the same
	// environment.
	//
	// If either `stack.account` or `stack.region` are not concrete values (e.g.
	// `Aws.account` or `Aws.region`) the special strings `unknown-account` and/or
	// `unknown-region` will be used respectively to indicate this stack is
	// region/account-agnostic.
	// Experimental.
	Environment() *string
	// The IAM execution role of the handler.
	// Experimental.
	HandlerRole() awsiam.IRole
	// Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.
	// Experimental.
	Nested() *bool
	// If this is a nested stack, returns it's parent stack.
	// Experimental.
	NestedStackParent() awscdk.Stack
	// If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.
	//
	// `undefined` for top-level (non-nested) stacks.
	// Experimental.
	NestedStackResource() awscdk.CfnResource
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns the list of notification Amazon Resource Names (ARNs) for the current stack.
	// Experimental.
	NotificationArns() *[]*string
	// Returns the parent of a nested stack.
	// Deprecated: use `nestedStackParent`.
	ParentStack() awscdk.Stack
	// The partition in which this stack is defined.
	// Experimental.
	Partition() *string
	// The AWS region into which this stack will be deployed (e.g. `us-west-2`).
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.region` when the stack is defined. This can
	//     either be a concerete region (e.g. `us-west-2`) or the `Aws.region`
	//     token.
	// 3. `Aws.region`, which is represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::Region" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.region)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **region-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	// Experimental.
	Region() *string
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	RoleArn() *string
	// The custom resource provider's service token.
	// Experimental.
	ServiceToken() *string
	// An attribute that represents the ID of the stack.
	//
	// This is a context aware attribute:
	// - If this is referenced from the parent stack, it will return `{ "Ref": "LogicalIdOfNestedStackResource" }`.
	// - If this is referenced from the context of the nested stack, it will return `{ "Ref": "AWS::StackId" }`
	//
	// Example value: `arn:aws:cloudformation:us-east-2:123456789012:stack/mystack-mynestedstack-sggfrhxhum7w/f449b250-b969-11e0-a185-5081d0136786`.
	// Experimental.
	StackId() *string
	// An attribute that represents the name of the nested stack.
	//
	// This is a context aware attribute:
	// - If this is referenced from the parent stack, it will return a token that parses the name from the stack ID.
	// - If this is referenced from the context of the nested stack, it will return `{ "Ref": "AWS::StackName" }`
	//
	// Example value: `mystack-mynestedstack-sggfrhxhum7w`.
	// Experimental.
	StackName() *string
	// Synthesis method for this stack.
	// Experimental.
	Synthesizer() awscdk.IStackSynthesizer
	// Tags to be applied to the stack.
	// Experimental.
	Tags() awscdk.TagManager
	// The name of the CloudFormation template file emitted to the output directory during synthesis.
	//
	// Example value: `MyStack.template.json`
	// Experimental.
	TemplateFile() *string
	// Options for CloudFormation template (like version, transform, description).
	// Experimental.
	TemplateOptions() awscdk.ITemplateOptions
	// Whether termination protection is enabled for this stack.
	// Experimental.
	TerminationProtection() *bool
	// The Amazon domain suffix for the region in which this stack is defined.
	// Experimental.
	UrlSuffix() *string
	// Add a dependency between this stack and another stack.
	//
	// This can be used to define dependencies between any two stacks within an
	// app, and also supports nested stacks.
	// Experimental.
	AddDependency(target awscdk.Stack, reason *string)
	// Register a docker image asset on this Stack.
	// Deprecated: Use `stack.synthesizer.addDockerImageAsset()` if you are calling,
	// and a different `IStackSynthesizer` class if you are implementing.
	AddDockerImageAsset(asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation
	// Register a file asset on this Stack.
	// Deprecated: Use `stack.synthesizer.addFileAsset()` if you are calling,
	// and a different IStackSynthesizer class if you are implementing.
	AddFileAsset(asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation
	// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
	//
	// Duplicate values are removed when stack is synthesized.
	//
	// Example:
	//   declare const stack: Stack;
	//
	//   stack.addTransform('AWS::Serverless-2016-10-31')
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
	//
	// Experimental.
	AddTransform(transform *string)
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
	AllocateLogicalId(cfnElement awscdk.CfnElement) *string
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
	ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string
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
	FormatArn(components *awscdk.ArnComponents) *string
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
	GetLogicalId(element awscdk.CfnElement) *string
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
	// Deprecated: use splitArn instead.
	ParseArn(arn *string, sepIfToken *string, hasName *bool) *awscdk.ArnComponents
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
	// Deprecated.
	//
	// Returns: reference itself without any change.
	// See: https://github.com/aws/aws-cdk/pull/7187
	//
	// Deprecated: cross reference handling has been moved to `App.prepare()`.
	PrepareCrossReference(_sourceStack awscdk.Stack, reference awscdk.Reference) awscdk.IResolvable
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
	RegionalFact(factName *string, defaultValue *string) *string
	// Rename a generated logical identities.
	//
	// To modify the naming scheme strategy, extend the `Stack` class and
	// override the `allocateLogicalId` method.
	// Experimental.
	RenameLogicalId(oldId *string, newId *string)
	// DEPRECATED.
	// Deprecated: use `reportMissingContextKey()`.
	ReportMissingContext(report *cxapi.MissingContext)
	// Indicate that a context key was expected.
	//
	// Contains instructions which will be emitted into the cloud assembly on how
	// the key should be supplied.
	// Experimental.
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	// Resolve a tokenized value in the context of the current stack.
	// Experimental.
	Resolve(obj interface{}) interface{}
	// Assign a value to one of the nested stack parameters.
	// Experimental.
	SetParameter(name *string, value *string)
	// Splits the provided ARN into its components.
	//
	// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
	// and a Token representing a dynamic CloudFormation expression
	// (in which case the returned components will also be dynamic CloudFormation expressions,
	// encoded as Tokens).
	// Experimental.
	SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Convert an object, potentially containing tokens, to a JSON string.
	// Experimental.
	ToJsonString(obj interface{}, space *float64) *string
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

func (j *jsiiProxy_KubectlProvider) BundlingRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bundlingRequired",
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

func (k *jsiiProxy_KubectlProvider) AddDependency(target awscdk.Stack, reason *string) {
	_jsii_.InvokeVoid(
		k,
		"addDependency",
		[]interface{}{target, reason},
	)
}

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

func (k *jsiiProxy_KubectlProvider) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		k,
		"addTransform",
		[]interface{}{transform},
	)
}

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

func (k *jsiiProxy_KubectlProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubectlProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (k *jsiiProxy_KubectlProvider) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

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

func (k *jsiiProxy_KubectlProvider) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		k,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

func (k *jsiiProxy_KubectlProvider) ReportMissingContext(report *cxapi.MissingContext) {
	_jsii_.InvokeVoid(
		k,
		"reportMissingContext",
		[]interface{}{report},
	)
}

func (k *jsiiProxy_KubectlProvider) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		k,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

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

func (k *jsiiProxy_KubectlProvider) SetParameter(name *string, value *string) {
	_jsii_.InvokeVoid(
		k,
		"setParameter",
		[]interface{}{name, value},
	)
}

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

func (k *jsiiProxy_KubectlProvider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   handlerRole := iam.role.fromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   kubectlProvider := eks.kubectlProvider.fromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &kubectlProviderAttributes{
//   	functionArn: jsii.String("arn:aws:lambda:us-east-2:123456789012:function:my-function:1"),
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456789012:role/kubectl-role"),
//   	handlerRole: handlerRole,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("Cluster"), &clusterAttributes{
//   	clusterName: jsii.String("cluster"),
//   	kubectlProvider: kubectlProvider,
//   })
//
// Experimental.
type KubectlProviderAttributes struct {
	// The kubectl provider lambda arn.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The IAM execution role of the handler.
	//
	// This role must be able to assume kubectlRoleArn.
	// Experimental.
	HandlerRole awsiam.IRole `field:"required" json:"handlerRole" yaml:"handlerRole"`
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	KubectlRoleArn *string `field:"required" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
}

// Kubectl Provider Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   kubectlProviderProps := &kubectlProviderProps{
//   	cluster: cluster,
//   }
//
// Experimental.
type KubectlProviderProps struct {
	// The cluster to control.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

// Represents a manifest within the Kubernetes system.
//
// Alternatively, you can use `cluster.addManifest(resource[, resource, ...])`
// to define resources on this cluster.
//
// Applies/deletes the manifest using `kubectl`.
//
// Example:
//   var cluster cluster
//
//   namespace := cluster.addManifest(jsii.String("my-namespace"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Namespace"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("my-app"),
//   	},
//   })
//
//   service := cluster.addManifest(jsii.String("my-service"), map[string]interface{}{
//   	"metadata": map[string]*string{
//   		"name": jsii.String("myservice"),
//   		"namespace": jsii.String("my-app"),
//   	},
//   	"spec": map[string]interface{}{
//   	},
//   })
//
//   service.node.addDependency(namespace)
//
// Experimental.
type KubernetesManifest interface {
	awscdk.Construct
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

func (k *jsiiProxy_KubernetesManifest) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesManifest) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (k *jsiiProxy_KubernetesManifest) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesManifest) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubernetesManifestOptions := &kubernetesManifestOptions{
//   	ingressAlb: jsii.Boolean(false),
//   	ingressAlbScheme: awscdk.Aws_eks.albScheme_INTERNAL,
//   	prune: jsii.Boolean(false),
//   	skipValidation: jsii.Boolean(false),
//   }
//
// Experimental.
type KubernetesManifestOptions struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	// Experimental.
	IngressAlb *bool `field:"optional" json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	// Experimental.
	IngressAlbScheme AlbScheme `field:"optional" json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
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
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

// Properties for KubernetesManifest.
//
// Example:
//   var cluster cluster
//
//   appLabel := map[string]*string{
//   	"app": jsii.String("hello-kubernetes"),
//   }
//
//   deployment := map[string]interface{}{
//   	"apiVersion": jsii.String("apps/v1"),
//   	"kind": jsii.String("Deployment"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"replicas": jsii.Number(3),
//   		"selector": map[string]map[string]*string{
//   			"matchLabels": appLabel,
//   		},
//   		"template": map[string]map[string]map[string]*string{
//   			"metadata": map[string]map[string]*string{
//   				"labels": appLabel,
//   			},
//   			"spec": map[string][]map[string]interface{}{
//   				"containers": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("hello-kubernetes"),
//   						"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   						"ports": []map[string]*f64{
//   							map[string]*f64{
//   								"containerPort": jsii.Number(8080),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
//   service := map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Service"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"type": jsii.String("LoadBalancer"),
//   		"ports": []map[string]*f64{
//   			map[string]*f64{
//   				"port": jsii.Number(80),
//   				"targetPort": jsii.Number(8080),
//   			},
//   		},
//   		"selector": appLabel,
//   	},
//   }
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewKubernetesManifest(this, jsii.String("hello-kub"), &kubernetesManifestProps{
//   	cluster: cluster,
//   	manifest: []map[string]interface{}{
//   		deployment,
//   		service,
//   	},
//   })
//
//   // or, option2: use `addManifest`
//   cluster.addManifest(jsii.String("hello-kub"), service, deployment)
//
// Experimental.
type KubernetesManifestProps struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	// Experimental.
	IngressAlb *bool `field:"optional" json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	// Experimental.
	IngressAlbScheme AlbScheme `field:"optional" json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
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
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// The EKS cluster to apply this manifest to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The manifest to apply.
	//
	// Consists of any number of child resources.
	//
	// When the resources are created/updated, this manifest will be applied to the
	// cluster through `kubectl apply` and when the resources or the stack is
	// deleted, the resources in the manifest will be deleted through `kubectl delete`.
	//
	// Example:
	//   []map[string]interface{}{
	//   	map[string]interface{}{
	//   		"apiVersion": jsii.String("v1"),
	//   		"kind": jsii.String("Pod"),
	//   		"metadata": map[string]*string{
	//   			"name": jsii.String("mypod"),
	//   		},
	//   		"spec": map[string][]map[string]interface{}{
	//   			"containers": []map[string]interface{}{
	//   				map[string]interface{}{
	//   					"name": jsii.String("hello"),
	//   					"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
	//   					"ports": []map[string]*f64{
	//   						map[string]*f64{
	//   							"containerPort": jsii.Number(8080),
	//   						},
	//   					},
	//   				},
	//   			},
	//   		},
	//   	},
	//   }
	//
	// Experimental.
	Manifest *[]*map[string]interface{} `field:"required" json:"manifest" yaml:"manifest"`
	// Overwrite any existing resources.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the resource is created. Otherwise, if there is already a resource
	// in the cluster with the same name, the operation will fail.
	// Experimental.
	Overwrite *bool `field:"optional" json:"overwrite" yaml:"overwrite"`
}

// Represents a value of a specific object deployed in the cluster.
//
// Use this to fetch any information available by the `kubectl get` command.
//
// Example:
//   var cluster cluster
//
//   // query the load balancer address
//   myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &kubernetesObjectValueProps{
//   	cluster: cluster,
//   	objectType: jsii.String("service"),
//   	objectName: jsii.String("my-service"),
//   	jsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
//   })
//
//   // pass the address to a lambda function
//   proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("my-code")),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	environment: map[string]*string{
//   		"myServiceAddress": myServiceAddress.value,
//   	},
//   })
//
// Experimental.
type KubernetesObjectValue interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The value as a string token.
	// Experimental.
	Value() *string
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

func (k *jsiiProxy_KubernetesObjectValue) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesObjectValue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (k *jsiiProxy_KubernetesObjectValue) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesObjectValue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//
//   // query the load balancer address
//   myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &kubernetesObjectValueProps{
//   	cluster: cluster,
//   	objectType: jsii.String("service"),
//   	objectName: jsii.String("my-service"),
//   	jsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
//   })
//
//   // pass the address to a lambda function
//   proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("my-code")),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	environment: map[string]*string{
//   		"myServiceAddress": myServiceAddress.value,
//   	},
//   })
//
// Experimental.
type KubernetesObjectValueProps struct {
	// The EKS cluster to fetch attributes from.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// JSONPath to the specific value.
	// See: https://kubernetes.io/docs/reference/kubectl/jsonpath/
	//
	// Experimental.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The name of the object to query.
	// Experimental.
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// The object type to query.
	//
	// (e.g 'service', 'pod'...)
	// Experimental.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// The namespace the object belongs to.
	// Experimental.
	ObjectNamespace *string `field:"optional" json:"objectNamespace" yaml:"objectNamespace"`
	// Timeout for waiting on a value.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// A CloudFormation resource which applies/restores a JSON patch into a Kubernetes resource.
//
// Example:
//   var cluster cluster
//
//   eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &kubernetesPatchProps{
//   	cluster: cluster,
//   	resourceName: jsii.String("deployment/hello-kubernetes"),
//   	applyPatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(5),
//   		},
//   	},
//   	restorePatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(3),
//   		},
//   	},
//   })
//
// See: https://kubernetes.io/docs/tasks/run-application/update-api-object-kubectl-patch/
//
// Experimental.
type KubernetesPatch interface {
	awscdk.Construct
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

func (k *jsiiProxy_KubernetesPatch) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesPatch) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (k *jsiiProxy_KubernetesPatch) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesPatch) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//
//   eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &kubernetesPatchProps{
//   	cluster: cluster,
//   	resourceName: jsii.String("deployment/hello-kubernetes"),
//   	applyPatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(5),
//   		},
//   	},
//   	restorePatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(3),
//   		},
//   	},
//   })
//
// Experimental.
type KubernetesPatchProps struct {
	// The JSON object to pass to `kubectl patch` when the resource is created/updated.
	// Experimental.
	ApplyPatch *map[string]interface{} `field:"required" json:"applyPatch" yaml:"applyPatch"`
	// The cluster to apply the patch to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The full name of the resource to patch (e.g. `deployment/coredns`).
	// Experimental.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The JSON object to pass to `kubectl patch` when the resource is removed.
	// Experimental.
	RestorePatch *map[string]interface{} `field:"required" json:"restorePatch" yaml:"restorePatch"`
	// The patch type to pass to `kubectl patch`.
	//
	// The default type used by `kubectl patch` is "strategic".
	// Experimental.
	PatchType PatchType `field:"optional" json:"patchType" yaml:"patchType"`
	// The kubernetes API namespace.
	// Experimental.
	ResourceNamespace *string `field:"optional" json:"resourceNamespace" yaml:"resourceNamespace"`
}

// Kubernetes cluster version.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacityType: eks.defaultCapacityType_EC2,
//   })
//
// Experimental.
type KubernetesVersion interface {
	// cluster version number.
	// Experimental.
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
// Example:
//   var cluster cluster
//
//
//   userData := "MIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"==MYBOUNDARY==\"\n\n--==MYBOUNDARY==\nContent-Type: text/x-shellscript; charset=\"us-ascii\"\n\n#!/bin/bash\necho \"Running custom user data script\"\n\n--==MYBOUNDARY==--\\\n"
//   lt := ec2.NewCfnLaunchTemplate(this, jsii.String("LaunchTemplate"), &cfnLaunchTemplateProps{
//   	launchTemplateData: &launchTemplateDataProperty{
//   		instanceType: jsii.String("t3.small"),
//   		userData: awscdk.Fn.base64(userData),
//   	},
//   })
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng"), &nodegroupOptions{
//   	launchTemplateSpec: &launchTemplateSpec{
//   		id: lt.ref,
//   		version: lt.attrLatestVersionNumber,
//   	},
//   })
//
// Experimental.
type LaunchTemplateSpec struct {
	// The Launch template ID.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// The machine image type.
//
// Example:
//   var cluster cluster
//
//   cluster.addAutoScalingGroupCapacity(jsii.String("BottlerocketNodes"), &autoScalingGroupCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   	minCapacity: jsii.Number(2),
//   	machineImageType: eks.machineImageType_BOTTLEROCKET,
//   })
//
// Experimental.
type MachineImageType string

const (
	// Amazon EKS-optimized Linux AMI.
	// Experimental.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	// Experimental.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

// Whether the worker nodes should support GPU or just standard instances.
// Experimental.
type NodeType string

const (
	// Standard instances.
	// Experimental.
	NodeType_STANDARD NodeType = "STANDARD"
	// GPU instances.
	// Experimental.
	NodeType_GPU NodeType = "GPU"
	// Inferentia instances.
	// Experimental.
	NodeType_INFERENTIA NodeType = "INFERENTIA"
)

// The Nodegroup resource class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   nodegroup := awscdk.Aws_eks.NewNodegroup(this, jsii.String("MyNodegroup"), &nodegroupProps{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	amiType: awscdk.*Aws_eks.nodegroupAmiType_AL2_X86_64,
//   	capacityType: awscdk.*Aws_eks.capacityType_SPOT,
//   	desiredSize: jsii.Number(123),
//   	diskSize: jsii.Number(123),
//   	forceUpdate: jsii.Boolean(false),
//   	instanceType: instanceType,
//   	instanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	launchTemplateSpec: &launchTemplateSpec{
//   		id: jsii.String("id"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	nodegroupName: jsii.String("nodegroupName"),
//   	nodeRole: role,
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &nodegroupRemoteAccess{
//   		sshKeyName: jsii.String("sshKeyName"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   	},
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	taints: []taintSpec{
//   		&taintSpec{
//   			effect: awscdk.*Aws_eks.taintEffect_NO_SCHEDULE,
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
// Experimental.
type Nodegroup interface {
	awscdk.Resource
	INodegroup
	// the Amazon EKS cluster resource.
	// Experimental.
	Cluster() ICluster
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// ARN of the nodegroup.
	// Experimental.
	NodegroupArn() *string
	// Nodegroup name.
	// Experimental.
	NodegroupName() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// IAM role of the instance profile for the nodegroup.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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

func (n *jsiiProxy_Nodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (n *jsiiProxy_Nodegroup) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Nodegroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (n *jsiiProxy_Nodegroup) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Nodegroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	defaultCapacity: jsii.Number(0),
//   })
//
//   cluster.addNodegroupCapacity(jsii.String("custom-node-group"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	minSize: jsii.Number(4),
//   	diskSize: jsii.Number(100),
//   	amiType: eks.nodegroupAmiType_AL2_X86_64_GPU,
//   })
//
// Experimental.
type NodegroupAmiType string

const (
	// Amazon Linux 2 (x86-64).
	// Experimental.
	NodegroupAmiType_AL2_X86_64 NodegroupAmiType = "AL2_X86_64"
	// Amazon Linux 2 with GPU support.
	// Experimental.
	NodegroupAmiType_AL2_X86_64_GPU NodegroupAmiType = "AL2_X86_64_GPU"
	// Amazon Linux 2 (ARM-64).
	// Experimental.
	NodegroupAmiType_AL2_ARM_64 NodegroupAmiType = "AL2_ARM_64"
	// Bottlerocket Linux(ARM-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_ARM_64 NodegroupAmiType = "BOTTLEROCKET_ARM_64"
	// Bottlerocket(x86-64).
	// Experimental.
	NodegroupAmiType_BOTTLEROCKET_X86_64 NodegroupAmiType = "BOTTLEROCKET_X86_64"
)

// The Nodegroup Options for addNodeGroup() method.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng-spot"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("c5.large")),
//   		ec2.NewInstanceType(jsii.String("c5a.large")),
//   		ec2.NewInstanceType(jsii.String("c5d.large")),
//   	},
//   	minSize: jsii.Number(3),
//   	capacityType: eks.capacityType_SPOT,
//   })
//
// Experimental.
type NodegroupOptions struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Experimental.
	AmiType NodegroupAmiType `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Experimental.
	CapacityType CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Experimental.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Experimental.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Experimental.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU` with the amiType parameter.
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Experimental.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Experimental.
	NodeRole awsiam.IRole `field:"optional" json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Experimental.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Experimental.
	Taints *[]*TaintSpec `field:"optional" json:"taints" yaml:"taints"`
}

// NodeGroup properties interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   nodegroupProps := &nodegroupProps{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	amiType: awscdk.Aws_eks.nodegroupAmiType_AL2_X86_64,
//   	capacityType: awscdk.*Aws_eks.capacityType_SPOT,
//   	desiredSize: jsii.Number(123),
//   	diskSize: jsii.Number(123),
//   	forceUpdate: jsii.Boolean(false),
//   	instanceType: instanceType,
//   	instanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	launchTemplateSpec: &launchTemplateSpec{
//   		id: jsii.String("id"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	nodegroupName: jsii.String("nodegroupName"),
//   	nodeRole: role,
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &nodegroupRemoteAccess{
//   		sshKeyName: jsii.String("sshKeyName"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   	},
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	taints: []taintSpec{
//   		&taintSpec{
//   			effect: awscdk.*Aws_eks.taintEffect_NO_SCHEDULE,
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type NodegroupProps struct {
	// The AMI type for your node group.
	//
	// If you explicitly specify the launchTemplate with custom AMI, do not specify this property, or
	// the node group deployment will fail. In other cases, you will need to specify correct amiType for the nodegroup.
	// Experimental.
	AmiType NodegroupAmiType `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of the nodegroup.
	// Experimental.
	CapacityType CapacityType `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The current number of worker nodes that the managed node group should maintain.
	//
	// If not specified,
	// the nodewgroup will initially create `minSize` instances.
	// Experimental.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// The root device disk size (in GiB) for your node group instances.
	// Experimental.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// If an update fails because pods could not be drained, you can force the update after it fails to terminate the old
	// node whether or not any pods are
	// running on the node.
	// Experimental.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// The instance type to use for your node group.
	//
	// Currently, you can specify a single instance type for a node group.
	// The default value for this parameter is `t3.medium`. If you choose a GPU instance type, be sure to specify the
	// `AL2_x86_64_GPU` with the amiType parameter.
	// Deprecated: Use `instanceTypes` instead.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The instance types to use for your node group.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Launch template specification used for the nodegroup.
	// See: - https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
	//
	// Experimental.
	LaunchTemplateSpec *LaunchTemplateSpec `field:"optional" json:"launchTemplateSpec" yaml:"launchTemplateSpec"`
	// The maximum number of worker nodes that the managed node group can scale out to.
	//
	// Managed node groups can support up to 100 nodes by default.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of worker nodes that the managed node group can scale in to.
	//
	// This number must be greater than or equal to zero.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Name of the Nodegroup.
	// Experimental.
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The IAM role to associate with your node group.
	//
	// The Amazon EKS worker node kubelet daemon
	// makes calls to AWS APIs on your behalf. Worker nodes receive permissions for these API calls through
	// an IAM instance profile and associated policies. Before you can launch worker nodes and register them
	// into a cluster, you must create an IAM role for those worker nodes to use when they are launched.
	// Experimental.
	NodeRole awsiam.IRole `field:"optional" json:"nodeRole" yaml:"nodeRole"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group (for example, `1.14.7-YYYYMMDD`).
	// Experimental.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Disabled by default, however, if you
	// specify an Amazon EC2 SSH key but do not specify a source security group when you create a managed node group,
	// then port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	// Experimental.
	RemoteAccess *NodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// By specifying the
	// SubnetSelection, the selected subnets will automatically apply required tags i.e.
	// `kubernetes.io/cluster/CLUSTER_NAME` with a value of `shared`, where `CLUSTER_NAME` is replaced with
	// the name of your cluster.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The metadata to apply to the node group to assist with categorization and organization.
	//
	// Each tag consists of
	// a key and an optional value, both of which you define. Node group tags do not propagate to any other resources
	// associated with the node group, such as the Amazon EC2 instances or subnets.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	// Experimental.
	Taints *[]*TaintSpec `field:"optional" json:"taints" yaml:"taints"`
	// Cluster resource.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

// The remote access (SSH) configuration to use with your node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   nodegroupRemoteAccess := &nodegroupRemoteAccess{
//   	sshKeyName: jsii.String("sshKeyName"),
//
//   	// the properties below are optional
//   	sourceSecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-remoteaccess.html
//
// Experimental.
type NodegroupRemoteAccess struct {
	// The Amazon EC2 SSH key that provides access for SSH communication with the worker nodes in the managed node group.
	// Experimental.
	SshKeyName *string `field:"required" json:"sshKeyName" yaml:"sshKeyName"`
	// The security groups that are allowed SSH access (port 22) to the worker nodes.
	//
	// If you specify an Amazon EC2 SSH
	// key but do not specify a source security group when you create a managed node group, then port 22 on the worker
	// nodes is opened to the internet (0.0.0.0/0).
	// Experimental.
	SourceSecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

// IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce.
//
// You use an IAM OIDC identity provider
// when you want to establish trust between an OIDC-compatible IdP and your AWS
// account.
//
// This implementation has default values for thumbprints and clientIds props
// that will be compatible with the eks cluster.
//
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.openIdConnectProvider.fromOpenIdConnectProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &openIdConnectProviderProps{
//   	url: issuerUrl,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("MyCluster"), &clusterAttributes{
//   	clusterName: jsii.String("Cluster"),
//   	openIdConnectProvider: provider,
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
//
// Experimental.
type OpenIdConnectProvider interface {
	awsiam.OpenIdConnectProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	// Experimental.
	OpenIdConnectProviderArn() *string
	// The issuer for OIDC Provider.
	// Experimental.
	OpenIdConnectProviderIssuer() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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

func (o *jsiiProxy_OpenIdConnectProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

func (o *jsiiProxy_OpenIdConnectProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		o,
		"onPrepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpenIdConnectProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		o,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (o *jsiiProxy_OpenIdConnectProvider) Prepare() {
	_jsii_.InvokeVoid(
		o,
		"prepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpenIdConnectProvider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		o,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.openIdConnectProvider.fromOpenIdConnectProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &openIdConnectProviderProps{
//   	url: issuerUrl,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("MyCluster"), &clusterAttributes{
//   	clusterName: jsii.String("Cluster"),
//   	openIdConnectProvider: provider,
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
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
	Url *string `field:"required" json:"url" yaml:"url"`
}

// Values for `kubectl patch` --type argument.
// Experimental.
type PatchType string

const (
	// JSON Patch, RFC 6902.
	// Experimental.
	PatchType_JSON PatchType = "JSON"
	// JSON Merge patch.
	// Experimental.
	PatchType_MERGE PatchType = "MERGE"
	// Strategic merge patch.
	// Experimental.
	PatchType_STRATEGIC PatchType = "STRATEGIC"
)

// Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selector := &selector{
//   	namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   }
//
// Experimental.
type Selector struct {
	// The Kubernetes namespace that the selector should match.
	//
	// You must specify a namespace for a selector. The selector only matches pods
	// that are created in this namespace, but you can create multiple selectors
	// to target multiple namespaces.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain
	// all of the labels that are specified in the selector for it to be
	// considered a match.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

// Service Account.
//
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.openIdConnectProvider.fromOpenIdConnectProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &openIdConnectProviderProps{
//   	url: issuerUrl,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("MyCluster"), &clusterAttributes{
//   	clusterName: jsii.String("Cluster"),
//   	openIdConnectProvider: provider,
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
//
// Experimental.
type ServiceAccount interface {
	awscdk.Construct
	awsiam.IPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return the policy fragment that identifies this principal in a Policy.
	// Experimental.
	PolicyFragment() awsiam.PrincipalPolicyFragment
	// The role which is linked to the service account.
	// Experimental.
	Role() awsiam.IRole
	// The name of the service account.
	// Experimental.
	ServiceAccountName() *string
	// The namespace where the service account is located in.
	// Experimental.
	ServiceAccountNamespace() *string
	// Add to the policy of this principal.
	// Deprecated: use `addToPrincipalPolicy()`.
	AddToPolicy(statement awsiam.PolicyStatement) *bool
	// Add to the policy of this principal.
	// Experimental.
	AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
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

func (s *jsiiProxy_ServiceAccount) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceAccount) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (s *jsiiProxy_ServiceAccount) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceAccount) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

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
// Example:
//   var cluster cluster
//
//   // add service account with annotations and labels
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"), &serviceAccountOptions{
//   	annotations: map[string]*string{
//   		"eks.amazonaws.com/sts-regional-endpoints": jsii.String("false"),
//   	},
//   	labels: map[string]*string{
//   		"some-label": jsii.String("with-some-value"),
//   	},
//   })
//
// Experimental.
type ServiceAccountOptions struct {
	// Additional annotations of the service account.
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Additional labels of the service account.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

// Properties for defining service accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   serviceAccountProps := &serviceAccountProps{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	annotations: map[string]*string{
//   		"annotationsKey": jsii.String("annotations"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	name: jsii.String("name"),
//   	namespace: jsii.String("namespace"),
//   }
//
// Experimental.
type ServiceAccountProps struct {
	// Additional annotations of the service account.
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Additional labels of the service account.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The cluster to apply the patch to.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

// Options for fetching a ServiceLoadBalancerAddress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   serviceLoadBalancerAddressOptions := &serviceLoadBalancerAddressOptions{
//   	namespace: jsii.String("namespace"),
//   	timeout: duration,
//   }
//
// Experimental.
type ServiceLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// Effect types of kubernetes node taint.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("custom-node-group"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	taints: []taintSpec{
//   		&taintSpec{
//   			effect: eks.taintEffect_NO_SCHEDULE,
//   			key: jsii.String("foo"),
//   			value: jsii.String("bar"),
//   		},
//   	},
//   })
//
// Experimental.
type TaintEffect string

const (
	// NoSchedule.
	// Experimental.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	// Experimental.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// NoExecute.
	// Experimental.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

// Taint interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taintSpec := &taintSpec{
//   	effect: awscdk.Aws_eks.taintEffect_NO_SCHEDULE,
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// Experimental.
type TaintSpec struct {
	// Effect type.
	// Experimental.
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// Taint key.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Taint value.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


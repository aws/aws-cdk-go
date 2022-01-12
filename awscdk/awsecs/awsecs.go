package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

// The properties for adding an AutoScalingGroup.
//
// TODO: EXAMPLE
//
type AddAutoScalingGroupCapacityOptions struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `json:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `json:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `json:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `json:"topicEncryptionKey"`
}

// The properties for adding instance capacity to an AutoScalingGroup.
//
// TODO: EXAMPLE
//
type AddCapacityOptions struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `json:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `json:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `json:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `json:"topicEncryptionKey"`
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
	BlockDevices *[]*awsautoscaling.BlockDevice `json:"blockDevices"`
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
	GroupMetrics *[]awsautoscaling.GroupMetrics `json:"groupMetrics"`
	// Configuration for health checks.
	HealthCheck awsautoscaling.HealthCheck `json:"healthCheck"`
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
	InstanceMonitoring awsautoscaling.Monitoring `json:"instanceMonitoring"`
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
	Notifications *[]*awsautoscaling.NotificationConfiguration `json:"notifications"`
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
	Signals awsautoscaling.Signals `json:"signals"`
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
	TerminationPolicies *[]awsautoscaling.TerminationPolicy `json:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	UpdatePolicy awsautoscaling.UpdatePolicy `json:"updatePolicy"`
	// Where to place instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The EC2 instance type to use when launching instances into the AutoScalingGroup.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The ECS-optimized AMI variant to use.
	//
	// The default is to use an ECS-optimized AMI of Amazon Linux 2 which is
	// automatically updated to the latest version on every deployment. This will
	// replace the instances in the AutoScalingGroup. Make sure you have not disabled
	// task draining, to avoid downtime when the AMI updates.
	//
	// To use an image that does not update on every deployment, pass:
	//
	// ```ts
	// const machineImage = ecs.EcsOptimizedImage.amazonLinux2(ecs.AmiHardwareType.STANDARD, {
	//    cachedInContext: true,
	// });
	// ```
	//
	// For more information, see [Amazon ECS-optimized
	// AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
	//
	// You must define either `machineImage` or `machineImageType`, not both.
	MachineImage awsec2.IMachineImage `json:"machineImage"`
}

// The ECS-optimized AMI variant to use.
//
// For more information, see
// [Amazon ECS-optimized AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
//
// TODO: EXAMPLE
//
type AmiHardwareType string

const (
	AmiHardwareType_STANDARD AmiHardwareType = "STANDARD"
	AmiHardwareType_GPU AmiHardwareType = "GPU"
	AmiHardwareType_ARM AmiHardwareType = "ARM"
)

// The class for App Mesh proxy configurations.
//
// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent and at least version
// 1.26.0-1 of the ecs-init package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized
// AMI version 20190301 or later, then they contain the required versions of the container agent and ecs-init.
// For more information, see [Amazon ECS-optimized AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
//
// For tasks using the Fargate launch type, the task or service requires platform version 1.3.0 or later.
//
// TODO: EXAMPLE
//
type AppMeshProxyConfiguration interface {
	ProxyConfiguration
	Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for AppMeshProxyConfiguration
type jsiiProxy_AppMeshProxyConfiguration struct {
	jsiiProxy_ProxyConfiguration
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
func NewAppMeshProxyConfiguration(props *AppMeshProxyConfigurationConfigProps) AppMeshProxyConfiguration {
	_init_.Initialize()

	j := jsiiProxy_AppMeshProxyConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
func NewAppMeshProxyConfiguration_Override(a AppMeshProxyConfiguration, props *AppMeshProxyConfigurationConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		a,
	)
}

// Called when the proxy configuration is configured on a task definition.
func (a *jsiiProxy_AppMeshProxyConfiguration) Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

// The configuration to use when setting an App Mesh proxy configuration.
//
// TODO: EXAMPLE
//
type AppMeshProxyConfigurationConfigProps struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `json:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin.
	Properties *AppMeshProxyConfigurationProps `json:"properties"`
}

// Interface for setting the properties of proxy configuration.
//
// TODO: EXAMPLE
//
type AppMeshProxyConfigurationProps struct {
	// The list of ports that the application uses.
	//
	// Network traffic to these ports is forwarded to the ProxyIngressPort and ProxyEgressPort.
	AppPorts *[]*float64 `json:"appPorts"`
	// Specifies the port that outgoing traffic from the AppPorts is directed to.
	ProxyEgressPort *float64 `json:"proxyEgressPort"`
	// Specifies the port that incoming traffic to the AppPorts is directed to.
	ProxyIngressPort *float64 `json:"proxyIngressPort"`
	// The egress traffic going to these specified IP addresses is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredIPs *[]*string `json:"egressIgnoredIPs"`
	// The egress traffic going to these specified ports is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredPorts *[]*float64 `json:"egressIgnoredPorts"`
	// The group ID (GID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredUID is specified, this field can be empty.
	IgnoredGID *float64 `json:"ignoredGID"`
	// The user ID (UID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredGID is specified, this field can be empty.
	IgnoredUID *float64 `json:"ignoredUID"`
}

// An Auto Scaling Group Capacity Provider.
//
// This allows an ECS cluster to target
// a specific EC2 Auto Scaling Group for the placement of tasks. Optionally (and
// recommended), ECS can manage the number of instances in the ASG to fit the
// tasks, and can ensure that instances are not prematurely terminated while
// there are still tasks running on them.
//
// TODO: EXAMPLE
//
type AsgCapacityProvider interface {
	constructs.Construct
	AutoScalingGroup() awsautoscaling.AutoScalingGroup
	CapacityProviderName() *string
	EnableManagedTerminationProtection() *bool
	MachineImageType() MachineImageType
	Node() constructs.Node
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AsgCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AsgCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
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

// The options for creating an Auto Scaling Group Capacity Provider.
//
// TODO: EXAMPLE
//
type AsgCapacityProviderProps struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `json:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `json:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `json:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `json:"topicEncryptionKey"`
	// The autoscaling group to add as a Capacity Provider.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `json:"autoScalingGroup"`
	// The name of the capacity provider.
	//
	// If a name is specified,
	// it cannot start with `aws`, `ecs`, or `fargate`. If no name is specified,
	// a default name in the CFNStackName-CFNResourceName-RandomString format is used.
	CapacityProviderName *string `json:"capacityProviderName"`
	// Whether to enable managed scaling.
	EnableManagedScaling *bool `json:"enableManagedScaling"`
	// Whether to enable managed termination protection.
	EnableManagedTerminationProtection *bool `json:"enableManagedTerminationProtection"`
	// Maximum scaling step size.
	//
	// In most cases this should be left alone.
	MaximumScalingStepSize *float64 `json:"maximumScalingStepSize"`
	// Minimum scaling step size.
	//
	// In most cases this should be left alone.
	MinimumScalingStepSize *float64 `json:"minimumScalingStepSize"`
	// Target capacity percent.
	//
	// In most cases this should be left alone.
	TargetCapacityPercent *float64 `json:"targetCapacityPercent"`
}

// Environment file from a local directory.
//
// TODO: EXAMPLE
//
type AssetEnvironmentFile interface {
	EnvironmentFile
	Path() *string
	Bind(scope constructs.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for AssetEnvironmentFile
type jsiiProxy_AssetEnvironmentFile struct {
	jsiiProxy_EnvironmentFile
}

func (j *jsiiProxy_AssetEnvironmentFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


func NewAssetEnvironmentFile(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	j := jsiiProxy_AssetEnvironmentFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetEnvironmentFile_Override(a AssetEnvironmentFile, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		a,
	)
}

// Loads the environment file from a local disk path.
func AssetEnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func AssetEnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Called when the container is initialized to allow this object to bind to the stack.
func (a *jsiiProxy_AssetEnvironmentFile) Bind(scope constructs.Construct) *EnvironmentFileConfig {
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// An image that will be built from a local directory with a Dockerfile.
//
// TODO: EXAMPLE
//
type AssetImage interface {
	ContainerImage
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for AssetImage
type jsiiProxy_AssetImage struct {
	jsiiProxy_ContainerImage
}

// Constructs a new instance of the AssetImage class.
func NewAssetImage(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	j := jsiiProxy_AssetImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AssetImage class.
func NewAssetImage_Override(a AssetImage, directory *string, props *AssetImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		a,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func AssetImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func AssetImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func AssetImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func AssetImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func AssetImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

// Called when the image is used by a ContainerDefinition.
func (a *jsiiProxy_AssetImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// The properties for building an AssetImage.
//
// TODO: EXAMPLE
//
type AssetImageProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *awsecrassets.DockerImageAssetInvalidationOptions `json:"invalidation"`
	// Docker target to build to.
	Target *string `json:"target"`
}

// The options for using a cloudmap service.
//
// TODO: EXAMPLE
//
type AssociateCloudMapServiceOptions struct {
	// The cloudmap service to register with.
	Service awsservicediscovery.IService `json:"service"`
	// The container to point to for a SRV record.
	Container ContainerDefinition `json:"container"`
	// The port to point to for a SRV record.
	ContainerPort *float64 `json:"containerPort"`
}

// The authorization configuration details for the Amazon EFS file system.
//
// TODO: EXAMPLE
//
type AuthorizationConfig struct {
	// The access point ID to use.
	//
	// If an access point is specified, the root directory value will be
	// relative to the directory set for the access point.
	// If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	AccessPointId *string `json:"accessPointId"`
	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the EFSVolumeConfiguration.
	//
	// Valid values: ENABLED | DISABLED
	Iam *string `json:"iam"`
}

// A log driver that sends log information to CloudWatch Logs.
//
// TODO: EXAMPLE
//
type AwsLogDriver interface {
	LogDriver
	LogGroup() awslogs.ILogGroup
	SetLogGroup(val awslogs.ILogGroup)
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for AwsLogDriver
type jsiiProxy_AwsLogDriver struct {
	jsiiProxy_LogDriver
}

func (j *jsiiProxy_AwsLogDriver) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}


// Constructs a new instance of the AwsLogDriver class.
func NewAwsLogDriver(props *AwsLogDriverProps) AwsLogDriver {
	_init_.Initialize()

	j := jsiiProxy_AwsLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AwsLogDriver class.
func NewAwsLogDriver_Override(a AwsLogDriver, props *AwsLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AwsLogDriver) SetLogGroup(val awslogs.ILogGroup) {
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func AwsLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (a *jsiiProxy_AwsLogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// awslogs provides two modes for delivering messages from the container to the log driver.
//
// TODO: EXAMPLE
//
type AwsLogDriverMode string

const (
	AwsLogDriverMode_BLOCKING AwsLogDriverMode = "BLOCKING"
	AwsLogDriverMode_NON_BLOCKING AwsLogDriverMode = "NON_BLOCKING"
)

// Specifies the awslogs log driver configuration options.
//
// TODO: EXAMPLE
//
type AwsLogDriverProps struct {
	// Prefix for the log streams.
	//
	// The awslogs-stream-prefix option allows you to associate a log stream
	// with the specified prefix, the container name, and the ID of the Amazon
	// ECS task to which the container belongs. If you specify a prefix with
	// this option, then the log stream takes the following format:
	//
	//      prefix-name/container-name/ecs-task-id
	StreamPrefix *string `json:"streamPrefix"`
	// This option defines a multiline start pattern in Python strftime format.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	DatetimeFormat *string `json:"datetimeFormat"`
	// The log group to log to.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs when the log group is automatically created by this construct.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// The delivery mode of log messages from the container to awslogs.
	Mode AwsLogDriverMode `json:"mode"`
	// This option defines a multiline start pattern using a regular expression.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	//
	// This option is ignored if datetimeFormat is also configured.
	MultilinePattern *string `json:"multilinePattern"`
}

// TODO: EXAMPLE
//
type BaseLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
}

// The base class for Ec2Service and FargateService services.
type BaseService interface {
	awscdk.Resource
	IBaseService
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	CloudMapService() awsservicediscovery.IService
	Cluster() ICluster
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	Node() constructs.Node
	PhysicalName() *string
	ServiceArn() *string
	ServiceName() *string
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	Stack() awscdk.Stack
	TaskDefinition() TaskDefinition
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	ToString() *string
}

// The jsii proxy struct for BaseService
type jsiiProxy_BaseService struct {
	internal.Type__awscdkResource
	jsiiProxy_IBaseService
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

func (j *jsiiProxy_BaseService) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the BaseService class.
func NewBaseService_Override(b BaseService, scope constructs.Construct, id *string, props *BaseServiceProps, additionalProps interface{}, taskDefinition TaskDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BaseService",
		[]interface{}{scope, id, props, additionalProps, taskDefinition},
		b,
	)
}

func (j *jsiiProxy_BaseService) SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_BaseService) SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_BaseService) SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_BaseService) SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BaseService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BaseService_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
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
func (b *jsiiProxy_BaseService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associates this service with a CloudMap service.
func (b *jsiiProxy_BaseService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		b,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

// This method is called to attach this service to an Application Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (b *jsiiProxy_BaseService) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		b,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Registers the service as a target of a Classic Load Balancer (CLB).
//
// Don't call this. Call `loadBalancer.addTarget()` instead.
func (b *jsiiProxy_BaseService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		b,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// This method is called to attach this service to a Network Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (b *jsiiProxy_BaseService) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		b,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
func (b *jsiiProxy_BaseService) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	var returns ScalableTaskCount

	_jsii_.Invoke(
		b,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method is called to create a networkConfiguration.
func (b *jsiiProxy_BaseService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		b,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

// Enable CloudMap service discovery for the service.
//
// Returns: The created CloudMap service
func (b *jsiiProxy_BaseService) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		b,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BaseService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BaseService) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return a load balancing target for a specific container and port.
//
// Use this function to create a load balancer target if you want to load balance to
// another container than the first essential container or the first mapped port on
// the container.
//
// Use the return value of this function where you would normally use a load balancer
// target, instead of the `Service` object itself.
//
// TODO: EXAMPLE
//
func (b *jsiiProxy_BaseService) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		b,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// This method returns the specified CloudWatch metric name for this service.
func (b *jsiiProxy_BaseService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's CPU utilization.
func (b *jsiiProxy_BaseService) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's memory utilization.
func (b *jsiiProxy_BaseService) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
//
// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
//
// TODO: EXAMPLE
//
func (b *jsiiProxy_BaseService) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"registerLoadBalancerTargets",
		args,
	)
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BaseService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the base Ec2Service or FargateService service.
//
// TODO: EXAMPLE
//
type BaseServiceOptions struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `json:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
}

// Complete base service properties that are required to be supplied by the implementation of the BaseService class.
//
// TODO: EXAMPLE
//
type BaseServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `json:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
	// The launch type on which to run your service.
	//
	// LaunchType will be omitted if capacity provider strategies are specified on the service.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
	//
	// Valid values are: LaunchType.ECS or LaunchType.FARGATE or LaunchType.EXTERNAL
	//
	LaunchType LaunchType `json:"launchType"`
}

// Instance resource used for bin packing.
type BinPackResource string

const (
	BinPackResource_CPU BinPackResource = "CPU"
	BinPackResource_MEMORY BinPackResource = "MEMORY"
)

// Construct an Bottlerocket image from the latest AMI published in SSM.
//
// TODO: EXAMPLE
//
type BottleRocketImage interface {
	awsec2.IMachineImage
	GetImage(scope constructs.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for BottleRocketImage
type jsiiProxy_BottleRocketImage struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the BottleRocketImage class.
func NewBottleRocketImage(props *BottleRocketImageProps) BottleRocketImage {
	_init_.Initialize()

	j := jsiiProxy_BottleRocketImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the BottleRocketImage class.
func NewBottleRocketImage_Override(b BottleRocketImage, props *BottleRocketImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		b,
	)
}

// Return the correct image.
func (b *jsiiProxy_BottleRocketImage) GetImage(scope constructs.Construct) *awsec2.MachineImageConfig {
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		b,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Properties for BottleRocketImage.
//
// TODO: EXAMPLE
//
type BottleRocketImageProps struct {
	// The CPU architecture.
	Architecture awsec2.InstanceArchitecture `json:"architecture"`
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	CachedInContext *bool `json:"cachedInContext"`
	// The Amazon ECS variant to use.
	//
	// Only `aws-ecs-1` is currently available
	Variant BottlerocketEcsVariant `json:"variant"`
}

// Amazon ECS variant.
type BottlerocketEcsVariant string

const (
	BottlerocketEcsVariant_AWS_ECS_1 BottlerocketEcsVariant = "AWS_ECS_1"
)

// The built-in container instance attributes.
//
// TODO: EXAMPLE
//
type BuiltInAttributes interface {
}

// The jsii proxy struct for BuiltInAttributes
type jsiiProxy_BuiltInAttributes struct {
	_ byte // padding
}

func NewBuiltInAttributes() BuiltInAttributes {
	_init_.Initialize()

	j := jsiiProxy_BuiltInAttributes{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewBuiltInAttributes_Override(b BuiltInAttributes) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		nil, // no parameters
		b,
	)
}

func BuiltInAttributes_AMI_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"AMI_ID",
		&returns,
	)
	return returns
}

func BuiltInAttributes_AVAILABILITY_ZONE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"AVAILABILITY_ZONE",
		&returns,
	)
	return returns
}

func BuiltInAttributes_INSTANCE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"INSTANCE_ID",
		&returns,
	)
	return returns
}

func BuiltInAttributes_INSTANCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"INSTANCE_TYPE",
		&returns,
	)
	return returns
}

func BuiltInAttributes_OS_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.BuiltInAttributes",
		"OS_TYPE",
		&returns,
	)
	return returns
}

// A Linux capability.
type Capability string

const (
	Capability_ALL Capability = "ALL"
	Capability_AUDIT_CONTROL Capability = "AUDIT_CONTROL"
	Capability_AUDIT_WRITE Capability = "AUDIT_WRITE"
	Capability_BLOCK_SUSPEND Capability = "BLOCK_SUSPEND"
	Capability_CHOWN Capability = "CHOWN"
	Capability_DAC_OVERRIDE Capability = "DAC_OVERRIDE"
	Capability_DAC_READ_SEARCH Capability = "DAC_READ_SEARCH"
	Capability_FOWNER Capability = "FOWNER"
	Capability_FSETID Capability = "FSETID"
	Capability_IPC_LOCK Capability = "IPC_LOCK"
	Capability_IPC_OWNER Capability = "IPC_OWNER"
	Capability_KILL Capability = "KILL"
	Capability_LEASE Capability = "LEASE"
	Capability_LINUX_IMMUTABLE Capability = "LINUX_IMMUTABLE"
	Capability_MAC_ADMIN Capability = "MAC_ADMIN"
	Capability_MAC_OVERRIDE Capability = "MAC_OVERRIDE"
	Capability_MKNOD Capability = "MKNOD"
	Capability_NET_ADMIN Capability = "NET_ADMIN"
	Capability_NET_BIND_SERVICE Capability = "NET_BIND_SERVICE"
	Capability_NET_BROADCAST Capability = "NET_BROADCAST"
	Capability_NET_RAW Capability = "NET_RAW"
	Capability_SETFCAP Capability = "SETFCAP"
	Capability_SETGID Capability = "SETGID"
	Capability_SETPCAP Capability = "SETPCAP"
	Capability_SETUID Capability = "SETUID"
	Capability_SYS_ADMIN Capability = "SYS_ADMIN"
	Capability_SYS_BOOT Capability = "SYS_BOOT"
	Capability_SYS_CHROOT Capability = "SYS_CHROOT"
	Capability_SYS_MODULE Capability = "SYS_MODULE"
	Capability_SYS_NICE Capability = "SYS_NICE"
	Capability_SYS_PACCT Capability = "SYS_PACCT"
	Capability_SYS_PTRACE Capability = "SYS_PTRACE"
	Capability_SYS_RAWIO Capability = "SYS_RAWIO"
	Capability_SYS_RESOURCE Capability = "SYS_RESOURCE"
	Capability_SYS_TIME Capability = "SYS_TIME"
	Capability_SYS_TTY_CONFIG Capability = "SYS_TTY_CONFIG"
	Capability_SYSLOG Capability = "SYSLOG"
	Capability_WAKE_ALARM Capability = "WAKE_ALARM"
)

// A Capacity Provider strategy to use for the service.
//
// NOTE: defaultCapacityProviderStrategy on cluster not currently supported.
//
// TODO: EXAMPLE
//
type CapacityProviderStrategy struct {
	// The name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider"`
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one
	// capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default
	// value of 0 is used.
	Base *float64 `json:"base"`
	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The weight value is taken into consideration after the base value, if defined, is satisfied.
	Weight *float64 `json:"weight"`
}

// A CloudFormation `AWS::ECS::CapacityProvider`.
//
// The `AWS::ECS::CapacityProvider` resource creates an Amazon Elastic Container Service (Amazon ECS) capacity provider. Capacity providers are associated with an Amazon ECS cluster and are used in capacity provider strategies to facilitate cluster auto scaling.
//
// Only capacity providers using an Auto Scaling group can be created. Amazon ECS tasks on AWS Fargate use the `FARGATE` and `FARGATE_SPOT` capacity providers which are already created and available to all accounts in Regions supported by AWS Fargate .
//
// TODO: EXAMPLE
//
type CfnCapacityProvider interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingGroupProvider() interface{}
	SetAutoScalingGroupProvider(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCapacityProvider
type jsiiProxy_CfnCapacityProvider struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCapacityProvider) AutoScalingGroupProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingGroupProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapacityProvider) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::CapacityProvider`.
func NewCfnCapacityProvider(scope constructs.Construct, id *string, props *CfnCapacityProviderProps) CfnCapacityProvider {
	_init_.Initialize()

	j := jsiiProxy_CfnCapacityProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::CapacityProvider`.
func NewCfnCapacityProvider_Override(c CfnCapacityProvider, scope constructs.Construct, id *string, props *CfnCapacityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCapacityProvider) SetAutoScalingGroupProvider(val interface{}) {
	_jsii_.Set(
		j,
		"autoScalingGroupProvider",
		val,
	)
}

func (j *jsiiProxy_CfnCapacityProvider) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCapacityProvider_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCapacityProvider_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
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
func CfnCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapacityProvider_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnCapacityProvider",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCapacityProvider) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCapacityProvider) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCapacityProvider) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCapacityProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCapacityProvider) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCapacityProvider) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnCapacityProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCapacityProvider) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCapacityProvider) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCapacityProvider) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCapacityProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCapacityProvider) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCapacityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCapacityProvider) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AutoScalingGroupProvider` property specifies the details of the Auto Scaling group for the capacity provider.
//
// TODO: EXAMPLE
//
type CfnCapacityProvider_AutoScalingGroupProviderProperty struct {
	// The Amazon Resource Name (ARN) or short name that identifies the Auto Scaling group.
	AutoScalingGroupArn *string `json:"autoScalingGroupArn"`
	// The managed scaling settings for the Auto Scaling group capacity provider.
	ManagedScaling interface{} `json:"managedScaling"`
	// The managed termination protection setting to use for the Auto Scaling group capacity provider.
	//
	// This determines whether the Auto Scaling group has managed termination protection. The default is disabled.
	//
	// > When using managed termination protection, managed scaling must also be used otherwise managed termination protection doesn't work.
	//
	// When managed termination protection is enabled, Amazon ECS prevents the Amazon EC2 instances in an Auto Scaling group that contain tasks from being terminated during a scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must have instance protection from scale-in actions enabled as well. For more information, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *AWS Auto Scaling User Guide* .
	//
	// When managed termination protection is disabled, your Amazon EC2 instances aren't protected from termination when the Auto Scaling group scales in.
	ManagedTerminationProtection *string `json:"managedTerminationProtection"`
}

// The `ManagedScaling` property specifies the settings for the Auto Scaling group capacity provider.
//
// When managed scaling is enabled, Amazon ECS manages the scale-in and scale-out actions of the Auto Scaling group. Amazon ECS manages a target tracking scaling policy using an Amazon ECS-managed CloudWatch metric with the specified `targetCapacity` value as the target value for the metric. For more information, see [Using Managed Scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling) in the *Amazon Elastic Container Service Developer Guide* .
//
// If managed scaling is disabled, the user must manage the scaling of the Auto Scaling group.
//
// TODO: EXAMPLE
//
type CfnCapacityProvider_ManagedScalingProperty struct {
	// The period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group.
	//
	// If this parameter is omitted, the default value of `300` seconds is used.
	InstanceWarmupPeriod *float64 `json:"instanceWarmupPeriod"`
	// The maximum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `10000` is used.
	MaximumScalingStepSize *float64 `json:"maximumScalingStepSize"`
	// The minimum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `1` is used.
	MinimumScalingStepSize *float64 `json:"minimumScalingStepSize"`
	// Determines whether to enable managed scaling for the capacity provider.
	Status *string `json:"status"`
	// The target capacity value for the capacity provider.
	//
	// The specified value must be greater than `0` and less than or equal to `100` . A value of `100` results in the Amazon EC2 instances in your Auto Scaling group being completely used.
	TargetCapacity *float64 `json:"targetCapacity"`
}

// Properties for defining a `CfnCapacityProvider`.
//
// TODO: EXAMPLE
//
type CfnCapacityProviderProps struct {
	// The Auto Scaling group settings for the capacity provider.
	AutoScalingGroupProvider interface{} `json:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
	Name *string `json:"name"`
	// The metadata that you apply to the capacity provider to help you categorize and organize it.
	//
	// Each tag consists of a key and an optional value. You define both.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::ECS::Cluster`.
//
// The `AWS::ECS::Cluster` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.
//
// TODO: EXAMPLE
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CapacityProviders() *[]*string
	SetCapacityProviders(val *[]*string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterSettings() interface{}
	SetClusterSettings(val interface{})
	Configuration() interface{}
	SetConfiguration(val interface{})
	CreationStack() *[]*string
	DefaultCapacityProviderStrategy() interface{}
	SetDefaultCapacityProviderStrategy(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnCluster) CapacityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capacityProviders",
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

func (j *jsiiProxy_CfnCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
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

func (j *jsiiProxy_CfnCluster) DefaultCapacityProviderStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCapacityProviderStrategy",
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

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
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


// Create a new `AWS::ECS::Cluster`.
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetCapacityProviders(val *[]*string) {
	_jsii_.Set(
		j,
		"capacityProviders",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterSettings(val interface{}) {
	_jsii_.Set(
		j,
		"clusterSettings",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDefaultCapacityProviderStrategy(val interface{}) {
	_jsii_.Set(
		j,
		"defaultCapacityProviderStrategy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCluster",
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnCluster",
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
		"aws-cdk-lib.aws_ecs.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `CapacityProviderStrategyItem` property specifies the details of the default capacity provider strategy for the cluster.
//
// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
//
// TODO: EXAMPLE
//
type CfnCluster_CapacityProviderStrategyItemProperty struct {
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `json:"base"`
	// The short name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that's run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `json:"weight"`
}

// The execute command configuration for the cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_ClusterConfigurationProperty struct {
	// The details of the execute command configuration.
	ExecuteCommandConfiguration interface{} `json:"executeCommandConfiguration"`
}

// The settings to use when creating a cluster.
//
// This parameter is used to enable CloudWatch Container Insights for a cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_ClusterSettingsProperty struct {
	// The name of the cluster setting.
	//
	// The only supported value is `containerInsights` .
	Name *string `json:"name"`
	// The value to set for the cluster setting.
	//
	// The supported values are `enabled` and `disabled` . If `enabled` is specified, CloudWatch Container Insights will be enabled for the cluster, otherwise it will be disabled unless the `containerInsights` account setting is enabled. If a cluster value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	Value *string `json:"value"`
}

// The details of the execute command configuration.
//
// TODO: EXAMPLE
//
type CfnCluster_ExecuteCommandConfigurationProperty struct {
	// Specify an AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKeyId *string `json:"kmsKeyId"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When `logging=OVERRIDE` is specified, a `logConfiguration` must be provided.
	LogConfiguration interface{} `json:"logConfiguration"`
	// The log setting to use for redirecting logs for your execute command results. The following log settings are available.
	//
	// - `NONE` : The execute command session is not logged.
	// - `DEFAULT` : The `awslogs` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no `awslogs` log driver is configured in the task definition, the output won't be logged.
	// - `OVERRIDE` : Specify the logging details as a part of `logConfiguration` . If the `OVERRIDE` logging option is specified, the `logConfiguration` is required.
	Logging *string `json:"logging"`
}

// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
type CfnCluster_ExecuteCommandLogConfigurationProperty struct {
	// Determines whether to enable encryption on the CloudWatch logs.
	//
	// If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled interface{} `json:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// > The CloudWatch log group must already be created.
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName"`
	// The name of the S3 bucket to send logs to.
	//
	// > The S3 bucket must already be created.
	S3BucketName *string `json:"s3BucketName"`
	// Determines whether to use encryption on the S3 logs.
	//
	// If not specified, encryption is not used.
	S3EncryptionEnabled interface{} `json:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
}

// A CloudFormation `AWS::ECS::ClusterCapacityProviderAssociations`.
//
// The `AWS::ECS::ClusterCapacityProviderAssociations` resource associates one or more capacity providers and a default capacity provider strategy with a cluster.
//
// TODO: EXAMPLE
//
type CfnClusterCapacityProviderAssociations interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CapacityProviders() *[]*string
	SetCapacityProviders(val *[]*string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cluster() *string
	SetCluster(val *string)
	CreationStack() *[]*string
	DefaultCapacityProviderStrategy() interface{}
	SetDefaultCapacityProviderStrategy(val interface{})
	LogicalId() *string
	Node() constructs.Node
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

// The jsii proxy struct for CfnClusterCapacityProviderAssociations
type jsiiProxy_CfnClusterCapacityProviderAssociations struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) CapacityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capacityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) DefaultCapacityProviderStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCapacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::ClusterCapacityProviderAssociations`.
func NewCfnClusterCapacityProviderAssociations(scope constructs.Construct, id *string, props *CfnClusterCapacityProviderAssociationsProps) CfnClusterCapacityProviderAssociations {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterCapacityProviderAssociations{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::ClusterCapacityProviderAssociations`.
func NewCfnClusterCapacityProviderAssociations_Override(c CfnClusterCapacityProviderAssociations, scope constructs.Construct, id *string, props *CfnClusterCapacityProviderAssociationsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) SetCapacityProviders(val *[]*string) {
	_jsii_.Set(
		j,
		"capacityProviders",
		val,
	)
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnClusterCapacityProviderAssociations) SetDefaultCapacityProviderStrategy(val interface{}) {
	_jsii_.Set(
		j,
		"defaultCapacityProviderStrategy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClusterCapacityProviderAssociations_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClusterCapacityProviderAssociations_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
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
func CfnClusterCapacityProviderAssociations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterCapacityProviderAssociations_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnClusterCapacityProviderAssociations",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `CapacityProviderStrategy` property specifies the details of the default capacity provider strategy for the cluster.
//
// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
//
// TODO: EXAMPLE
//
type CfnClusterCapacityProviderAssociations_CapacityProviderStrategyProperty struct {
	// The short name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider"`
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `json:"base"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` will not be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that is run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `json:"weight"`
}

// Properties for defining a `CfnClusterCapacityProviderAssociations`.
//
// TODO: EXAMPLE
//
type CfnClusterCapacityProviderAssociationsProps struct {
	// The capacity providers to associate with the cluster.
	CapacityProviders *[]*string `json:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	Cluster *string `json:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy interface{} `json:"defaultCapacityProviderStrategy"`
}

// Properties for defining a `CfnCluster`.
//
// TODO: EXAMPLE
//
type CfnClusterProps struct {
	// The short name of one or more capacity providers to associate with the cluster.
	//
	// A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created and not already associated with another cluster.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	CapacityProviders *[]*string `json:"capacityProviders"`
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	ClusterName *string `json:"clusterName"`
	// The setting to use when creating a cluster.
	//
	// This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	ClusterSettings interface{} `json:"clusterSettings"`
	// The execute command configuration for the cluster.
	Configuration interface{} `json:"configuration"`
	// The default capacity provider strategy for the cluster.
	//
	// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
	DefaultCapacityProviderStrategy interface{} `json:"defaultCapacityProviderStrategy"`
	// The metadata that you apply to the cluster to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value. You define both.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::ECS::PrimaryTaskSet`.
//
// Specifies which task set in a service is the primary task set. Any parameters that are updated on the primary task set in a service will transition to the service. This is used when a service uses the `EXTERNAL` deployment controller type. For more information, see [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnPrimaryTaskSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cluster() *string
	SetCluster(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Service() *string
	SetService(val *string)
	Stack() awscdk.Stack
	TaskSetId() *string
	SetTaskSetId(val *string)
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

// The jsii proxy struct for CfnPrimaryTaskSet
type jsiiProxy_CfnPrimaryTaskSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPrimaryTaskSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) TaskSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrimaryTaskSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::PrimaryTaskSet`.
func NewCfnPrimaryTaskSet(scope constructs.Construct, id *string, props *CfnPrimaryTaskSetProps) CfnPrimaryTaskSet {
	_init_.Initialize()

	j := jsiiProxy_CfnPrimaryTaskSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::PrimaryTaskSet`.
func NewCfnPrimaryTaskSet_Override(c CfnPrimaryTaskSet, scope constructs.Construct, id *string, props *CfnPrimaryTaskSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPrimaryTaskSet) SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnPrimaryTaskSet) SetService(val *string) {
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_CfnPrimaryTaskSet) SetTaskSetId(val *string) {
	_jsii_.Set(
		j,
		"taskSetId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPrimaryTaskSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPrimaryTaskSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
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
func CfnPrimaryTaskSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrimaryTaskSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnPrimaryTaskSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPrimaryTaskSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPrimaryTaskSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPrimaryTaskSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPrimaryTaskSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPrimaryTaskSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrimaryTaskSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPrimaryTaskSet`.
//
// TODO: EXAMPLE
//
type CfnPrimaryTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set exists in.
	Cluster *string `json:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service that the task set exists in.
	Service *string `json:"service"`
	// The short name or full Amazon Resource Name (ARN) of the task set to set as the primary task set in the deployment.
	TaskSetId *string `json:"taskSetId"`
}

// A CloudFormation `AWS::ECS::Service`.
//
// The `AWS::ECS::Service` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
//
// TODO: EXAMPLE
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrName() *string
	AttrServiceArn() *string
	CapacityProviderStrategy() interface{}
	SetCapacityProviderStrategy(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cluster() *string
	SetCluster(val *string)
	CreationStack() *[]*string
	DeploymentConfiguration() interface{}
	SetDeploymentConfiguration(val interface{})
	DeploymentController() interface{}
	SetDeploymentController(val interface{})
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	HealthCheckGracePeriodSeconds() *float64
	SetHealthCheckGracePeriodSeconds(val *float64)
	LaunchType() *string
	SetLaunchType(val *string)
	LoadBalancers() interface{}
	SetLoadBalancers(val interface{})
	LogicalId() *string
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	Node() constructs.Node
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	PlacementStrategies() interface{}
	SetPlacementStrategies(val interface{})
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PropagateTags() *string
	SetPropagateTags(val *string)
	Ref() *string
	Role() *string
	SetRole(val *string)
	SchedulingStrategy() *string
	SetSchedulingStrategy(val *string)
	ServiceName() *string
	SetServiceName(val *string)
	ServiceRegistries() interface{}
	SetServiceRegistries(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TaskDefinition() *string
	SetTaskDefinition(val *string)
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

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnService) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CapacityProviderStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DeploymentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DeploymentController() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LoadBalancers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlacementConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlacementStrategies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceRegistries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::Service`.
func NewCfnService(scope constructs.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::Service`.
func NewCfnService_Override(c CfnService, scope constructs.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService) SetCapacityProviderStrategy(val interface{}) {
	_jsii_.Set(
		j,
		"capacityProviderStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetDeploymentConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetDeploymentController(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentController",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetDesiredCount(val *float64) {
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetEnableEcsManagedTags(val interface{}) {
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetEnableExecuteCommand(val interface{}) {
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckGracePeriodSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetLaunchType(val *string) {
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetLoadBalancers(val interface{}) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetPlacementConstraints(val interface{}) {
	_jsii_.Set(
		j,
		"placementConstraints",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetPlacementStrategies(val interface{}) {
	_jsii_.Set(
		j,
		"placementStrategies",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetPlatformVersion(val *string) {
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetPropagateTags(val *string) {
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetSchedulingStrategy(val *string) {
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetServiceRegistries(val interface{}) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetTaskDefinition(val *string) {
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
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
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnService) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnService) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object representing the networking details for a task or service.
//
// TODO: EXAMPLE
//
type CfnService_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	Subnets *[]*string `json:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	AssignPublicIp *string `json:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	SecurityGroups *[]*string `json:"securityGroups"`
}

// The details of a capacity provider strategy.
//
// A capacity provider strategy can be set when using the `RunTask` or `CreateService` APIs or as the default capacity provider strategy for a cluster with the `CreateCluster` API.
//
// Only capacity providers that are already associated with a cluster and have an `ACTIVE` or `UPDATING` status can be used in a capacity provider strategy. The `PutClusterCapacityProviders` API is used to associate a capacity provider with a cluster.
//
// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created. New Auto Scaling group capacity providers can be created with the `CreateCapacityProvider` API operation.
//
// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used in a capacity provider strategy.
//
// TODO: EXAMPLE
//
type CfnService_CapacityProviderStrategyItemProperty struct {
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `json:"base"`
	// The short name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that's run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `json:"weight"`
}

// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type.
//
// The `DeploymentCircuitBreaker` property determines whether a service deployment will fail if the service can't reach a steady state. If deployment circuit breaker is enabled, a service deployment will transition to a failed state and stop launching new tasks. If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
//
// TODO: EXAMPLE
//
type CfnService_DeploymentCircuitBreakerProperty struct {
	// Determines whether to enable the deployment circuit breaker logic for the service.
	Enable interface{} `json:"enable"`
	// Determines whether to enable Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	Rollback interface{} `json:"rollback"`
}

// The `DeploymentConfiguration` property specifies optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
//
// TODO: EXAMPLE
//
type CfnService_DeploymentConfigurationProperty struct {
	// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type that are not behind a Classic Load Balancer.
	//
	// The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If enabled, a service deployment will transition to a failed state and stop launching new tasks. You can also enable Amazon ECS to roll back your service to the last completed deployment after a failure. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide* .
	DeploymentCircuitBreaker interface{} `json:"deploymentCircuitBreaker"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the *maximum percent* parameter represents an upper limit on the number of tasks in a service that are allowed in the `RUNNING` or `PENDING` state during a deployment, as a percentage of the desired number of tasks (rounded down to the nearest integer), and while any container instances are in the `DRAINING` state if the service contains tasks using the EC2 launch type.
	//
	// This parameter enables you to define the deployment batch size. For example, if your service has a desired number of four tasks and a maximum percent value of 200%, the scheduler may start four new tasks before stopping the four older tasks (provided that the cluster resources required to do this are available). The default value for maximum percent is 200%.
	//
	// If a service is using the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and tasks that use the EC2 launch type, the *maximum percent* value is set to the default value and is used to define the upper limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state. If the tasks in the service use the Fargate launch type, the maximum percent value is not used, although it is returned when describing your service.
	MaximumPercent *float64 `json:"maximumPercent"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the *minimum healthy percent* represents a lower limit on the number of tasks in a service that must remain in the `RUNNING` state during a deployment, as a percentage of the desired number of tasks (rounded up to the nearest integer), and while any container instances are in the `DRAINING` state if the service contains tasks using the EC2 launch type.
	//
	// This parameter enables you to deploy without using additional cluster capacity. For example, if your service has a desired number of four tasks and a minimum healthy percent of 50%, the scheduler may stop two existing tasks to free up cluster capacity before starting two new tasks. Tasks for services that *do not* use a load balancer are considered healthy if they're in the `RUNNING` state; tasks for services that *do* use a load balancer are considered healthy if they're in the `RUNNING` state and they're reported as healthy by the load balancer. The default value for minimum healthy percent is 100%.
	//
	// If a service is using the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and tasks that use the EC2 launch type, the *minimum healthy percent* value is set to the default value and is used to define the lower limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state. If the tasks in the service use the Fargate launch type, the minimum healthy percent value is not used, although it is returned when describing your service.
	MinimumHealthyPercent *float64 `json:"minimumHealthyPercent"`
}

// The deployment controller to use for the service.
//
// For more information, see [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnService_DeploymentControllerProperty struct {
	// The deployment controller type to use. There are three deployment controller types available:.
	//
	// - **ECS** - The rolling update ( `ECS` ) deployment type involves replacing the current running version of the container with the latest version. The number of containers Amazon ECS adds or removes from the service during a rolling update is controlled by adjusting the minimum and maximum number of healthy tasks allowed during a service deployment, as specified in the [DeploymentConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeploymentConfiguration.html) .
	// - **CODE_DEPLOY** - The blue/green ( `CODE_DEPLOY` ) deployment type uses the blue/green deployment model powered by AWS CodeDeploy , which allows you to verify a new deployment of a service before sending production traffic to it.
	// - **EXTERNAL** - The external ( `EXTERNAL` ) deployment type enables you to use any third-party deployment controller for full control over the deployment process for an Amazon ECS service.
	Type *string `json:"type"`
}

// The `LoadBalancer` property specifies details on a load balancer that is used with a service.
//
// If the service is using the `CODE_DEPLOY` deployment controller, the service is required to use either an Application Load Balancer or Network Load Balancer. When you are creating an AWS CodeDeploy deployment group, you specify two target groups (referred to as a `targetGroupPair` ). Each target group binds to a separate task set in the deployment. The load balancer can also have up to two listeners, a required listener for production traffic and an optional listener that allows you to test new revisions of the service before routing production traffic to it.
//
// Services with tasks that use the `awsvpc` network mode (for example, those with the Fargate launch type) only support Application Load Balancers and Network Load Balancers. Classic Load Balancers are not supported. Also, when you create any target groups for these services, you must choose `ip` as the target type, not `instance` . Tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance.
//
// TODO: EXAMPLE
//
type CfnService_LoadBalancerProperty struct {
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	ContainerPort *float64 `json:"containerPort"`
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	ContainerName *string `json:"containerName"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you're using a Classic Load Balancer, omit the target group ARN.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering Multiple Target Groups with a Service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/Green Deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	TargetGroupArn *string `json:"targetGroupArn"`
}

// The `NetworkConfiguration` property specifies an object representing the network configuration for a task or service.
//
// TODO: EXAMPLE
//
type CfnService_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	AwsvpcConfiguration interface{} `json:"awsvpcConfiguration"`
}

// The `PlacementConstraint` property specifies an object representing a constraint on task placement in the task definition.
//
// For more information, see [Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnService_PlacementConstraintProperty struct {
	// The type of constraint.
	//
	// Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
	Type *string `json:"type"`
	// A cluster query language expression to apply to the constraint.
	//
	// The expression can have a maximum length of 2000 characters. You can't specify an expression if the constraint type is `distinctInstance` . For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	Expression *string `json:"expression"`
}

// The `PlacementStrategy` property specifies the task placement strategy for a task or service.
//
// For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnService_PlacementStrategyProperty struct {
	// The type of placement strategy.
	//
	// The `random` placement strategy randomly places tasks on available candidates. The `spread` placement strategy spreads placement across available candidates evenly based on the `field` parameter. The `binpack` strategy places tasks on available candidates that have the least available amount of the resource that's specified with the `field` parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory but still enough to run the task.
	Type *string `json:"type"`
	// The field to apply the placement strategy against.
	//
	// For the `spread` placement strategy, valid values are `instanceId` (or `host` , which has the same effect), or any platform or custom attribute that's applied to a container instance, such as `attribute:ecs.availability-zone` . For the `binpack` placement strategy, valid values are `cpu` and `memory` . For the `random` placement strategy, this field is not used.
	Field *string `json:"field"`
}

// The `ServiceRegistry` property specifies details of the service registry.
//
// For more information, see [Service Discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnService_ServiceRegistryProperty struct {
	// The container name value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition that your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition that your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerName *string `json:"containerName"`
	// The port value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerPort *float64 `json:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field might be used if both the `awsvpc` network mode and SRV records are used.
	Port *float64 `json:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map . For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) .
	RegistryArn *string `json:"registryArn"`
}

// Properties for defining a `CfnService`.
//
// TODO: EXAMPLE
//
type CfnServiceProps struct {
	// The capacity provider strategy to use for the service.
	//
	// A capacity provider strategy consists of one or more capacity providers along with the `base` and `weight` to assign to them. A capacity provider must be associated with the cluster to be used in a capacity provider strategy. The PutClusterCapacityProviders API is used to associate a capacity provider with a cluster. Only capacity providers with an `ACTIVE` or `UPDATING` status can be used.
	//
	// Review the [Capacity provider considerations](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html#capacity-providers-considerations) in the *Amazon Elastic Container Service Developer Guide.*
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or `launchType` is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created. New capacity providers can be created with the CreateCapacityProvider API operation.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	//
	// The PutClusterCapacityProviders API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
	CapacityProviderStrategy interface{} `json:"capacityProviderStrategy"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	Cluster *string `json:"cluster"`
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration interface{} `json:"deploymentConfiguration"`
	// The deployment controller to use for the service.
	//
	// If no deployment controller is specified, the default value of `ECS` is used.
	DeploymentController interface{} `json:"deploymentController"`
	// The number of instantiations of the specified task definition to place and keep running on your cluster.
	//
	// For new services, if a desired count is not specified, a default value of `1` is used. When using the `DAEMON` scheduling strategy, the desired count is not required.
	//
	// For existing services, if a desired count is not specified, it is omitted from the operation.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnableEcsManagedTags interface{} `json:"enableEcsManagedTags"`
	// Determines whether the execute command functionality is enabled for the service.
	//
	// If `true` , the execute command functionality is enabled for all containers in tasks as part of the service.
	EnableExecuteCommand interface{} `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	//
	// This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of `0` is used.
	//
	// If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
	HealthCheckGracePeriodSeconds *float64 `json:"healthCheckGracePeriodSeconds"`
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType *string `json:"launchType"`
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the `Role` property, `LoadBalancers` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide* .
	LoadBalancers interface{} `json:"loadBalancers"`
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the `awsvpc` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	NetworkConfiguration interface{} `json:"networkConfiguration"`
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	PlacementConstraints interface{} `json:"placementConstraints"`
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of five strategy rules per service. For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlacementStrategies interface{} `json:"placementStrategies"`
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion *string `json:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the tasks within the service during service creation. To add tags to a task after service creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	PropagateTags *string `json:"propagateTags"`
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the `awsvpc` network mode. If you specify the `role` parameter, you must also specify a load balancer object with the `loadBalancers` parameter.
	//
	// > If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the `awsvpc` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If your specified role has a path other than `/` , then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	Role *string `json:"role"`
	// The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// There are two service scheduler strategies available:
	//
	// - `REPLICA` -The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the `CODE_DEPLOY` or `EXTERNAL` deployment controller types.
	// - `DAEMON` -The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//
	// > Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy.
	SchedulingStrategy *string `json:"schedulingStrategy"`
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	ServiceName *string `json:"serviceName"`
	// The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	//
	// > Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	ServiceRegistries interface{} `json:"serviceRegistries"`
	// The metadata that you apply to the service to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	//
	// The `revision` is required in order for the resource to stabilize.
	//
	// A task definition must be specified if the service is using either the `ECS` or `CODE_DEPLOY` deployment controllers.
	TaskDefinition *string `json:"taskDefinition"`
}

// A CloudFormation `AWS::ECS::TaskDefinition`.
//
// The `AWS::ECS::TaskDefinition` resource describes the container and volume definitions of an Amazon Elastic Container Service (Amazon ECS) task. You can specify which Docker images to use, the required resources, and other configurations related to launching the task definition through an Amazon ECS service or task.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrTaskDefinitionArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContainerDefinitions() interface{}
	SetContainerDefinitions(val interface{})
	Cpu() *string
	SetCpu(val *string)
	CreationStack() *[]*string
	EphemeralStorage() interface{}
	SetEphemeralStorage(val interface{})
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	Family() *string
	SetFamily(val *string)
	InferenceAccelerators() interface{}
	SetInferenceAccelerators(val interface{})
	IpcMode() *string
	SetIpcMode(val *string)
	LogicalId() *string
	Memory() *string
	SetMemory(val *string)
	NetworkMode() *string
	SetNetworkMode(val *string)
	Node() constructs.Node
	PidMode() *string
	SetPidMode(val *string)
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	ProxyConfiguration() interface{}
	SetProxyConfiguration(val interface{})
	Ref() *string
	RequiresCompatibilities() *[]*string
	SetRequiresCompatibilities(val *[]*string)
	RuntimePlatform() interface{}
	SetRuntimePlatform(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	UpdatedProperites() *map[string]interface{}
	Volumes() interface{}
	SetVolumes(val interface{})
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

// The jsii proxy struct for CfnTaskDefinition
type jsiiProxy_CfnTaskDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTaskDefinition) AttrTaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTaskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ContainerDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) EphemeralStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) InferenceAccelerators() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) IpcMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) PidMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) PlacementConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ProxyConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) RequiresCompatibilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) RuntimePlatform() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Volumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinition(scope constructs.Construct, id *string, props *CfnTaskDefinitionProps) CfnTaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnTaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinition_Override(c CfnTaskDefinition, scope constructs.Construct, id *string, props *CfnTaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetContainerDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"containerDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetCpu(val *string) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetEphemeralStorage(val interface{}) {
	_jsii_.Set(
		j,
		"ephemeralStorage",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetInferenceAccelerators(val interface{}) {
	_jsii_.Set(
		j,
		"inferenceAccelerators",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetIpcMode(val *string) {
	_jsii_.Set(
		j,
		"ipcMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetMemory(val *string) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetNetworkMode(val *string) {
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetPidMode(val *string) {
	_jsii_.Set(
		j,
		"pidMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetPlacementConstraints(val interface{}) {
	_jsii_.Set(
		j,
		"placementConstraints",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetProxyConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"proxyConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetRequiresCompatibilities(val *[]*string) {
	_jsii_.Set(
		j,
		"requiresCompatibilities",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetRuntimePlatform(val interface{}) {
	_jsii_.Set(
		j,
		"runtimePlatform",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetTaskRoleArn(val *string) {
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition) SetVolumes(val interface{}) {
	_jsii_.Set(
		j,
		"volumes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTaskDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTaskDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
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
func CfnTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnTaskDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTaskDefinition) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTaskDefinition) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTaskDefinition) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTaskDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTaskDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTaskDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTaskDefinition) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTaskDefinition) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTaskDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTaskDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTaskDefinition) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The authorization configuration details for the Amazon EFS file system.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_AuthorizationConfigProperty struct {
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, the root directory value specified in the `EFSVolumeConfiguration` must either be omitted or set to `/` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be enabled in the `EFSVolumeConfiguration` . For more information, see [Working with Amazon EFS Access Points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide* .
	AccessPointId *string `json:"accessPointId"`
	// Determines whether to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the `EFSVolumeConfiguration` . If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Using Amazon EFS Access Points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide* .
	Iam *string `json:"iam"`
}

// The `ContainerDefinition` property specifies a container definition.
//
// Container definitions are used in task definitions to describe the different containers that are launched as part of a task.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_ContainerDefinitionProperty struct {
	// The command that's passed to the container.
	//
	// This parameter maps to `Cmd` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `COMMAND` parameter to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . For more information, see [https://docs.docker.com/engine/reference/builder/#cmd](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#cmd) . If there are multiple arguments, each argument is a separated string in the array.
	Command *[]*string `json:"command"`
	// The number of `cpu` units reserved for the container.
	//
	// This parameter maps to `CpuShares` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--cpu-shares` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// This field is optional for tasks using the Fargate launch type, and the only requirement is that the total amount of CPU reserved for all containers within a task be lower than the task-level `cpu` value.
	//
	// > You can determine the number of CPU units that are available per EC2 instance type by multiplying the vCPUs listed for that instance type on the [Amazon EC2 Instances](https://docs.aws.amazon.com/ec2/instance-types/) detail page by 1,024.
	//
	// Linux containers share unallocated CPU units with other containers on the container instance with the same ratio as their allocated amount. For example, if you run a single-container task on a single-core instance type with 512 CPU units specified for that container, and that's the only task running on the container instance, that container could use the full 1,024 CPU unit share at any given time. However, if you launched another copy of the same task on that container instance, each task is guaranteed a minimum of 512 CPU units when needed. Moreover, each container could float to higher CPU usage if the other container was not using it. If both tasks were 100% active all of the time, they would be limited to 512 CPU units.
	//
	// On Linux container instances, the Docker daemon on the container instance uses the CPU value to calculate the relative CPU share ratios for running containers. For more information, see [CPU share constraint](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#cpu-share-constraint) in the Docker documentation. The minimum valid CPU share value that the Linux kernel allows is 2. However, the CPU parameter isn't required, and you can use CPU values below 2 in your container definitions. For CPU values below 2 (including null), the behavior varies based on your Amazon ECS container agent version:
	//
	// - *Agent versions less than or equal to 1.1.0:* Null and zero CPU values are passed to Docker as 0, which Docker then converts to 1,024 CPU shares. CPU values of 1 are passed to Docker as 1, which the Linux kernel converts to two CPU shares.
	// - *Agent versions greater than or equal to 1.2.0:* Null, zero, and CPU values of 1 are passed to Docker as 2.
	//
	// On Windows container instances, the CPU limit is enforced as an absolute limit, or a quota. Windows containers only have access to the specified amount of CPU that's described in the task definition. A null or zero CPU value is passed to Docker as `0` , which Windows interprets as 1% of one CPU.
	Cpu *float64 `json:"cpu"`
	// The dependencies defined for container startup and shutdown.
	//
	// A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.
	//
	// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent to enable container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For tasks using the Fargate launch type, the task or service requires the following platforms:
	//
	// - Linux platform version `1.3.0` or later.
	// - Windows platform version `1.0.0` or later.
	DependsOn interface{} `json:"dependsOn"`
	// When this parameter is true, networking is disabled within the container.
	//
	// This parameter maps to `NetworkDisabled` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) .
	//
	// > This parameter is not supported for Windows containers.
	DisableNetworking interface{} `json:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	//
	// This parameter maps to `DnsSearch` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--dns-search` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	DnsSearchDomains *[]*string `json:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	//
	// This parameter maps to `Dns` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--dns` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	DnsServers *[]*string `json:"dnsServers"`
	// A key/value map of labels to add to the container.
	//
	// This parameter maps to `Labels` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--label` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	DockerLabels interface{} `json:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	//
	// This field isn't valid for containers in tasks using the Fargate launch type.
	//
	// With Windows containers, this parameter can be used to reference a credential spec file when configuring a container for Active Directory authentication. For more information, see [Using gMSAs for Windows Containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows-gmsa.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// This parameter maps to `SecurityOpt` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--security-opt` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > The Amazon ECS container agent running on a container instance must register with the `ECS_SELINUX_CAPABLE=true` or `ECS_APPARMOR_CAPABLE=true` environment variables before containers placed on that instance can use these security options. For more information, see [Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For more information about valid values, see [Docker Run Security Configuration](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// Valid values: "no-new-privileges" | "apparmor:PROFILE" | "label:value" | "credentialspec:CredentialSpecFilePath"
	DockerSecurityOptions *[]*string `json:"dockerSecurityOptions"`
	// > Early versions of the Amazon ECS container agent don't properly handle `entryPoint` parameters.
	//
	// If you have problems using `entryPoint` , update your container agent or enter your commands and arguments as `command` array items instead.
	//
	// The entry point that's passed to the container. This parameter maps to `Entrypoint` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--entrypoint` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . For more information, see [https://docs.docker.com/engine/reference/builder/#entrypoint](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#entrypoint) .
	EntryPoint *[]*string `json:"entryPoint"`
	// The environment variables to pass to a container.
	//
	// This parameter maps to `Env` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--env` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > We don't recommend that you use plaintext environment variables for sensitive information, such as credential data.
	Environment interface{} `json:"environment"`
	// A list of files containing the environment variables to pass to a container.
	//
	// This parameter maps to the `--env-file` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file contains an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .
	//
	// If there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying Environment Variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnvironmentFiles interface{} `json:"environmentFiles"`
	// If the `essential` parameter of a container is marked as `true` , and that container fails or stops for any reason, all other containers that are part of the task are stopped.
	//
	// If the `essential` parameter of a container is marked as `false` , its failure doesn't affect the rest of the containers in a task. If this parameter is omitted, a container is assumed to be essential.
	//
	// All tasks must have at least one essential container. If you have an application that's composed of multiple containers, group containers that are used for a common purpose into components, and separate the different components into multiple task definitions. For more information, see [Application Architecture](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/application_architecture.html) in the *Amazon Elastic Container Service Developer Guide* .
	Essential interface{} `json:"essential"`
	// A list of hostnames and IP address mappings to append to the `/etc/hosts` file on the container.
	//
	// This parameter maps to `ExtraHosts` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--add-host` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter isn't supported for Windows containers or tasks that use the `awsvpc` network mode.
	ExtraHosts interface{} `json:"extraHosts"`
	// The FireLens configuration for the container.
	//
	// This is used to specify and configure a log router for container logs. For more information, see [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
	FirelensConfiguration interface{} `json:"firelensConfiguration"`
	// The container health check command and associated configuration parameters for the container.
	//
	// This parameter maps to `HealthCheck` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `HEALTHCHECK` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	HealthCheck interface{} `json:"healthCheck"`
	// The hostname to use for your container.
	//
	// This parameter maps to `Hostname` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--hostname` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > The `hostname` parameter is not supported if you're using the `awsvpc` network mode.
	Hostname *string `json:"hostname"`
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon. By default, images in the Docker Hub registry are available. Other repositories are specified with either `*repository-url* / *image* : *tag*` or `*repository-url* / *image* @ *digest*` . Up to 255 letters (uppercase and lowercase), numbers, hyphens, underscores, colons, periods, forward slashes, and number signs are allowed. This parameter maps to `Image` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `IMAGE` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// - When a new task starts, the Amazon ECS container agent pulls the latest version of the specified image and tag for the container to use. However, subsequent updates to a repository image aren't propagated to already running tasks.
	// - Images in Amazon ECR repositories can be specified by either using the full `registry/repository:tag` or `registry/repository@digest` . For example, `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest` or `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE` .
	// - Images in official repositories on Docker Hub use a single name (for example, `ubuntu` or `mongo` ).
	// - Images in other repositories on Docker Hub are qualified with an organization name (for example, `amazon/amazon-ecs-agent` ).
	// - Images in other online repositories are qualified further by a domain name (for example, `quay.io/assemblyline/ubuntu` ).
	Image *string `json:"image"`
	// When this parameter is `true` , you can deploy containerized applications that require `stdin` or a `tty` to be allocated.
	//
	// This parameter maps to `OpenStdin` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--interactive` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	Interactive interface{} `json:"interactive"`
	// The `links` parameter allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the network mode of a task definition is `bridge` . The `name:internalName` construct is analogous to `name:alias` in Docker links. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. For more information about linking Docker containers, go to [Legacy container links](https://docs.aws.amazon.com/https://docs.docker.com/network/links/) in the Docker documentation. This parameter maps to `Links` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--link` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers. > Containers that are collocated on a single container instance may be able to communicate with each other without requiring links or host port mappings. Network isolation is achieved on the container instance using security groups and VPC settings.
	Links *[]*string `json:"links"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities. For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html) .
	//
	// > This parameter is not supported for Windows containers.
	LinuxParameters interface{} `json:"linuxParameters"`
	// The log configuration specification for the container.
	//
	// This parameter maps to `LogConfig` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--log-driver` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . By default, containers use the same logging driver that the Docker daemon uses. However, the container may use a different logging driver than the Docker daemon by specifying a log driver with this parameter in the container definition. To use a different logging driver for a container, the log system must be configured properly on the container instance (or on a different log server for remote logging options). For more information on the options for different supported log drivers, see [Configure logging drivers](https://docs.aws.amazon.com/https://docs.docker.com/engine/admin/logging/overview/) in the Docker documentation.
	//
	// > Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon (shown in the [LogConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_LogConfiguration.html) data type). Additional log drivers may be available in future releases of the Amazon ECS container agent.
	//
	// This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	//
	// > The Amazon ECS container agent running on a container instance must register the logging drivers available on that instance with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide* .
	LogConfiguration interface{} `json:"logConfiguration"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the memory specified here, the container is killed. The total amount of memory reserved for all containers within a task must be lower than the task `memory` value, if one is specified. This parameter maps to `Memory` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--memory` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// If using the Fargate launch type, this parameter is optional.
	//
	// If using the EC2 launch type, you must specify either a task-level memory value or a container-level memory value. If you specify both a container-level `memory` and `memoryReservation` value, `memory` must be greater than `memoryReservation` . If you specify `memoryReservation` , then that value is subtracted from the available memory resources for the container instance where the container is placed. Otherwise, the value of `memory` is used.
	//
	// The Docker 20.10.0 or later daemon reserves a minimum of 6 MiB of memory for a container, so you should not specify fewer than 6 MiB of memory for your containers.
	//
	// The Docker 19.03.13-ce or earlier daemon reserves a minimum of 4 MiB of memory for a container, so you should not specify fewer than 4 MiB of memory for your containers.
	Memory *float64 `json:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the container memory to this soft limit. However, your container can consume more memory when it needs to, up to either the hard limit specified with the `memory` parameter (if applicable), or all of the available memory on the container instance, whichever comes first. This parameter maps to `MemoryReservation` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--memory-reservation` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// If a task-level memory value is not specified, you must specify a non-zero integer for one or both of `memory` or `memoryReservation` in a container definition. If you specify both, `memory` must be greater than `memoryReservation` . If you specify `memoryReservation` , then that value is subtracted from the available memory resources for the container instance where the container is placed. Otherwise, the value of `memory` is used.
	//
	// For example, if your container normally uses 128 MiB of memory, but occasionally bursts to 256 MiB of memory for short periods of time, you can set a `memoryReservation` of 128 MiB, and a `memory` hard limit of 300 MiB. This configuration would allow the container to only reserve 128 MiB of memory from the remaining resources on the container instance, but also allow the container to consume more memory resources when needed.
	//
	// The Docker daemon reserves a minimum of 4 MiB of memory for a container. Therefore, we recommend that you specify fewer than 4 MiB of memory for your containers.
	MemoryReservation *float64 `json:"memoryReservation"`
	// The mount points for data volumes in your container.
	//
	// This parameter maps to `Volumes` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--volume` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// Windows containers can mount whole directories on the same drive as `$env:ProgramData` . Windows containers can't mount directories on a different drive, and mount point can't be across drives.
	MountPoints interface{} `json:"mountPoints"`
	// The name of a container.
	//
	// If you're linking multiple containers together in a task definition, the `name` of one container can be entered in the `links` of another container to connect the containers. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This parameter maps to `name` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--name` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	Name *string `json:"name"`
	// The list of port mappings for the container.
	//
	// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
	//
	// For task definitions that use the `awsvpc` network mode, you should only specify the `containerPort` . The `hostPort` can be left blank or it must be the same value as the `containerPort` .
	//
	// Port mappings on Windows use the `NetNAT` gateway address rather than `localhost` . There is no loopback for port mappings on Windows, so you cannot access a container's mapped port from the host itself.
	//
	// This parameter maps to `PortBindings` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--publish` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . If the network mode of a task definition is set to `none` , then you can't specify port mappings. If the network mode of a task definition is set to `host` , then host ports must either be undefined or they must match the container port in the port mapping.
	//
	// > After a task reaches the `RUNNING` status, manual and automatic host and container port assignments are visible in the *Network Bindings* section of a container description for a selected task in the Amazon ECS console. The assignments are also visible in the `networkBindings` section [DescribeTasks](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTasks.html) responses.
	PortMappings interface{} `json:"portMappings"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the `root` user).
	//
	// This parameter maps to `Privileged` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--privileged` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	Privileged interface{} `json:"privileged"`
	// When this parameter is `true` , a TTY is allocated.
	//
	// This parameter maps to `Tty` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--tty` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	PseudoTerminal interface{} `json:"pseudoTerminal"`
	// When this parameter is true, the container is given read-only access to its root file system.
	//
	// This parameter maps to `ReadonlyRootfs` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--read-only` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	ReadonlyRootFilesystem interface{} `json:"readonlyRootFilesystem"`
	// The private repository authentication credentials to use.
	RepositoryCredentials interface{} `json:"repositoryCredentials"`
	// The type and amount of a resource to assign to a container.
	//
	// The only supported resource is a GPU.
	ResourceRequirements interface{} `json:"resourceRequirements"`
	// The secrets to pass to the container.
	//
	// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	Secrets interface{} `json:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	//
	// For example, you specify two containers in a task definition with containerA having a dependency on containerB reaching a `COMPLETE` , `SUCCESS` , or `HEALTHY` status. If a `startTimeout` value is specified for containerB and it doesn't reach the desired status within that time then containerA gives up and not start. This results in the task transitioning to a `STOPPED` state.
	//
	// > When the `ECS_CONTAINER_START_TIMEOUT` container agent configuration variable is used, it's enforced independently from this start timeout value.
	//
	// For tasks using the Fargate launch type, the task or service requires the following platforms:
	//
	// - Linux platform version `1.3.0` or later.
	// - Windows platform version `1.0.0` or later.
	//
	// For tasks using the EC2 launch type, your container instances require at least version `1.26.0` of the container agent to enable a container start timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version `1.26.0-1` of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	StartTimeout *float64 `json:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	//
	// For tasks using the Fargate launch type, the task or service requires the following platforms:
	//
	// - Linux platform version `1.3.0` or later.
	// - Windows platform version `1.0.0` or later.
	//
	// The max stop timeout value is 120 seconds and if the parameter is not specified, the default value of 30 seconds is used.
	//
	// For tasks that use the EC2 launch type, if the `stopTimeout` parameter isn't specified, the value set for the Amazon ECS container agent configuration variable `ECS_CONTAINER_STOP_TIMEOUT` is used. If neither the `stopTimeout` parameter or the `ECS_CONTAINER_STOP_TIMEOUT` agent configuration variable are set, then the default values of 30 seconds for Linux containers and 30 seconds on Windows containers are used. Your container instances require at least version 1.26.0 of the container agent to enable a container stop timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	StopTimeout *float64 `json:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	//
	// This parameter maps to `Sysctls` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--sysctl` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > We don't recommended that you specify network-related `systemControls` parameters for multiple containers in a single task that also uses either the `awsvpc` or `host` network modes. For tasks that use the `awsvpc` network mode, the container that's started last determines which `systemControls` parameters take effect. For tasks that use the `host` network mode, it changes the container instance's namespaced kernel parameters as well as the containers.
	SystemControls interface{} `json:"systemControls"`
	// A list of `ulimits` to set in the container.
	//
	// This parameter maps to `Ulimits` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--ulimit` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . Valid naming values are displayed in the [Ulimit](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Ulimit.html) data type. This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	//
	// > This parameter is not supported for Windows containers.
	Ulimits interface{} `json:"ulimits"`
	// The user to use inside the container.
	//
	// This parameter maps to `User` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--user` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > When running tasks using the `host` network mode, don't run containers using the root user (UID 0). We recommend using a non-root user for better security.
	//
	// You can specify the `user` using the following formats. If specifying a UID or GID, you must specify it as a positive integer.
	//
	// - `user`
	// - `user:group`
	// - `uid`
	// - `uid:gid`
	// - `user:gid`
	// - `uid:group`
	//
	// > This parameter is not supported for Windows containers.
	User *string `json:"user"`
	// Data volumes to mount from another container.
	//
	// This parameter maps to `VolumesFrom` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--volumes-from` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	VolumesFrom interface{} `json:"volumesFrom"`
	// The working directory to run commands inside the container in.
	//
	// This parameter maps to `WorkingDir` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--workdir` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	WorkingDirectory *string `json:"workingDirectory"`
}

// The `ContainerDependency` property specifies the dependencies defined for container startup and shutdown.
//
// A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.
//
// Your Amazon ECS container instances require at least version 1.26.0 of the container agent to enable container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you are using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > For tasks using the Fargate launch type, this parameter requires that the task or service uses platform version 1.3.0 or later.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_ContainerDependencyProperty struct {
	// The dependency condition of the container. The following are the available conditions and their behavior:.
	//
	// - `START` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.
	// - `COMPLETE` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.
	// - `SUCCESS` - This condition is the same as `COMPLETE` , but it also requires that the container exits with a `zero` status. This condition can't be set on an essential container.
	// - `HEALTHY` - This condition validates that the dependent container passes its Docker health check before permitting other containers to start. This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.
	Condition *string `json:"condition"`
	// The name of a container.
	ContainerName *string `json:"containerName"`
}

// The `Device` property specifies an object representing a container instance host device.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_DeviceProperty struct {
	// The path inside the container at which to expose the host device.
	ContainerPath *string `json:"containerPath"`
	// The path for the device on the host container instance.
	HostPath *string `json:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	Permissions *[]*string `json:"permissions"`
}

// The `DockerVolumeConfiguration` property specifies a Docker volume configuration and is used when you use Docker volumes.
//
// Docker volumes are only supported when you are using the EC2 launch type. Windows containers only support the use of the `local` driver. To use bind mounts, specify a `host` instead.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_DockerVolumeConfigurationProperty struct {
	// If this value is `true` , the Docker volume is created if it doesn't already exist.
	//
	// > This field is only used if the `scope` is `shared` .
	Autoprovision interface{} `json:"autoprovision"`
	// The Docker volume driver to use.
	//
	// The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use `docker plugin ls` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. For more information, see [Docker plugin discovery](https://docs.aws.amazon.com/https://docs.docker.com/engine/extend/plugin_api/#plugin-discovery) . This parameter maps to `Driver` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxdriver` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Driver *string `json:"driver"`
	// A map of Docker driver-specific options passed through.
	//
	// This parameter maps to `DriverOpts` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxopt` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	DriverOpts interface{} `json:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	//
	// This parameter maps to `Labels` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxlabel` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Labels interface{} `json:"labels"`
	// The scope for the Docker volume that determines its lifecycle.
	//
	// Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.
	Scope *string `json:"scope"`
}

// TODO: EXAMPLE
//
type CfnTaskDefinition_EfsVolumeConfigurationProperty struct {
	// `CfnTaskDefinition.EfsVolumeConfigurationProperty.FileSystemId`.
	FileSystemId *string `json:"fileSystemId"`
	// `CfnTaskDefinition.EfsVolumeConfigurationProperty.AuthorizationConfig`.
	AuthorizationConfig interface{} `json:"authorizationConfig"`
	// `CfnTaskDefinition.EfsVolumeConfigurationProperty.RootDirectory`.
	RootDirectory *string `json:"rootDirectory"`
	// `CfnTaskDefinition.EfsVolumeConfigurationProperty.TransitEncryption`.
	TransitEncryption *string `json:"transitEncryption"`
	// `CfnTaskDefinition.EfsVolumeConfigurationProperty.TransitEncryptionPort`.
	TransitEncryptionPort *float64 `json:"transitEncryptionPort"`
}

// A list of files containing the environment variables to pass to a container.
//
// You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file should contain an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .
//
// If there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// This parameter is only supported for tasks hosted on Fargate using the following platform versions:
//
// - Linux platform version `1.4.0` or later.
// - Windows platform version `1.0.0` or later.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_EnvironmentFileProperty struct {
	// The file type to use.
	//
	// The only supported value is `s3` .
	Type *string `json:"type"`
	// The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.
	Value *string `json:"value"`
}

// The amount of ephemeral storage to allocate for the task.
//
// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate . For more information, see [Fargate task storage](https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html) in the *Amazon ECS User Guide for AWS Fargate* .
//
// > This parameter is only supported for tasks hosted on Fargate using the following platform versions:
// >
// > - Linux platform version `1.4.0` or later.
// > - Windows platform version `1.0.0` or later.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_EphemeralStorageProperty struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.
	SizeInGiB *float64 `json:"sizeInGiB"`
}

// The FireLens configuration for the container.
//
// This is used to specify and configure a log router for container logs. For more information, see [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_FirelensConfigurationProperty struct {
	// The options to use when configuring the log router.
	//
	// This field is optional and can be used to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event.
	//
	// If specified, valid option keys are:
	//
	// - `enable-ecs-log-metadata` , which can be `true` or `false`
	// - `config-file-type` , which can be `s3` or `file`
	// - `config-file-value` , which is either an S3 ARN or a file path
	Options interface{} `json:"options"`
	// The log router to use.
	//
	// The valid values are `fluentd` or `fluentbit` .
	Type *string `json:"type"`
}

// The `HealthCheck` property specifies an object representing a container health check.
//
// Health check parameters that are specified in a container definition override any Docker health checks that exist in the container image (such as those specified in a parent image or from the image's Dockerfile).
//
// The following are notes about container health check support:
//
// - Container health checks require version 1.17.0 or greater of the Amazon ECS container agent. For more information, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) .
// - Container health checks are supported for Fargate tasks if you are using platform version 1.1.0 or greater. For more information, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) .
// - Container health checks are not supported for tasks that are part of a service that is configured to use a Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_HealthCheckProperty struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with `CMD` to execute the command arguments directly, or `CMD-SHELL` to run the command with the container's default shell.
	//
	// When you use the AWS Management Console JSON panel, the AWS Command Line Interface , or the APIs, enclose the list of commands in brackets.
	//
	// `[ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ]`
	//
	// You don't need to include the brackets when you use the AWS Management Console.
	//
	// `"CMD-SHELL", "curl -f http://localhost/ || exit 1"`
	//
	// An exit code of 0 indicates success, and non-zero exit code indicates failure. For more information, see `HealthCheck` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) .
	Command *[]*string `json:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds. The default value is 30 seconds.
	Interval *float64 `json:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries. The default value is 3.
	Retries *float64 `json:"retries"`
	// The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You can specify between 0 and 300 seconds. By default, the `startPeriod` is disabled.
	//
	// > If a health check succeeds within the `startPeriod` , then the container is considered healthy and any subsequent failures count toward the maximum number of retries.
	StartPeriod *float64 `json:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds. The default value is 5.
	Timeout *float64 `json:"timeout"`
}

// The `HostEntry` property specifies a hostname and an IP address that are added to the `/etc/hosts` file of a container through the `extraHosts` parameter of its `ContainerDefinition` resource.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_HostEntryProperty struct {
	// The hostname to use in the `/etc/hosts` entry.
	Hostname *string `json:"hostname"`
	// The IP address to use in the `/etc/hosts` entry.
	IpAddress *string `json:"ipAddress"`
}

// The `HostVolumeProperties` property specifies details on a container instance bind mount host volume.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_HostVolumePropertiesProperty struct {
	// When the `host` parameter is used, specify a `sourcePath` to declare the path on the host container instance that's presented to the container.
	//
	// If this parameter is empty, then the Docker daemon has assigned a host path for you. If the `host` parameter contains a `sourcePath` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the `sourcePath` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
	//
	// If you're using the Fargate launch type, the `sourcePath` parameter is not supported.
	SourcePath *string `json:"sourcePath"`
}

// Details on an Elastic Inference accelerator.
//
// For more information, see [Working with Amazon Elastic Inference on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-eia.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_InferenceAcceleratorProperty struct {
	// The Elastic Inference accelerator device name.
	//
	// The `deviceName` must also be referenced in a container definition as a [ResourceRequirement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-resourcerequirement.html) .
	DeviceName *string `json:"deviceName"`
	// The Elastic Inference accelerator type to use.
	DeviceType *string `json:"deviceType"`
}

// The `KernelCapabilities` property specifies the Linux capabilities for the container that are added to or dropped from the default configuration that is provided by Docker.
//
// For more information on the default capabilities and the non-default available capabilities, see [Runtime privilege and Linux capabilities](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities) in the *Docker run reference* . For more detailed information on these Linux capabilities, see the [capabilities(7)](https://docs.aws.amazon.com/http://man7.org/linux/man-pages/man7/capabilities.7.html) Linux manual page.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_KernelCapabilitiesProperty struct {
	// The Linux capabilities for the container that have been added to the default configuration provided by Docker.
	//
	// This parameter maps to `CapAdd` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--cap-add` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > Tasks launched on AWS Fargate only support adding the `SYS_PTRACE` kernel capability.
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`
	Add *[]*string `json:"add"`
	// The Linux capabilities for the container that have been removed from the default configuration provided by Docker.
	//
	// This parameter maps to `CapDrop` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--cap-drop` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`
	Drop *[]*string `json:"drop"`
}

// The `KeyValuePair` property specifies a key-value pair object.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_KeyValuePairProperty struct {
	// The name of the key-value pair.
	//
	// For environment variables, this is the name of the environment variable.
	Name *string `json:"name"`
	// The value of the key-value pair.
	//
	// For environment variables, this is the value of the environment variable.
	Value *string `json:"value"`
}

// The `LinuxParameters` property specifies Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html) .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_LinuxParametersProperty struct {
	// The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.
	//
	// > For tasks that use the Fargate launch type, `capabilities` is supported for all platform versions but the `add` parameter is only supported if using platform version 1.4.0 or later.
	Capabilities interface{} `json:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to `Devices` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--device` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `devices` parameter isn't supported.
	Devices interface{} `json:"devices"`
	// Run an `init` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the `--init` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	InitProcessEnabled interface{} `json:"initProcessEnabled"`
	// The total amount of swap memory (in MiB) a container can use.
	//
	// This parameter will be translated to the `--memory-swap` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) where the value would be the sum of the container memory plus the `maxSwap` value.
	//
	// If a `maxSwap` value of `0` is specified, the container will not use swap. Accepted values are `0` or any positive integer. If the `maxSwap` parameter is omitted, the container will use the swap configuration for the container instance it is running on. A `maxSwap` value must be set for the `swappiness` parameter to be used.
	//
	// > If you're using tasks that use the Fargate launch type, the `maxSwap` parameter isn't supported.
	MaxSwap *float64 `json:"maxSwap"`
	// The value for the size (in MiB) of the `/dev/shm` volume.
	//
	// This parameter maps to the `--shm-size` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you are using tasks that use the Fargate launch type, the `sharedMemorySize` parameter is not supported.
	SharedMemorySize *float64 `json:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// A `swappiness` value of `0` will cause swapping to not happen unless absolutely necessary. A `swappiness` value of `100` will cause pages to be swapped very aggressively. Accepted values are whole numbers between `0` and `100` . If the `swappiness` parameter is not specified, a default value of `60` is used. If a value is not specified for `maxSwap` then this parameter is ignored. This parameter maps to the `--memory-swappiness` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `swappiness` parameter isn't supported.
	Swappiness *float64 `json:"swappiness"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the `--tmpfs` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `tmpfs` parameter isn't supported.
	Tmpfs interface{} `json:"tmpfs"`
}

// The `LogConfiguration` property specifies log configuration options to send to a custom log driver for the container.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_LogConfigurationProperty struct {
	// The log driver to use for the container.
	//
	// For tasks on AWS Fargate , the supported log drivers are `awslogs` , `splunk` , and `awsfirelens` .
	//
	// For tasks hosted on Amazon EC2 instances, the supported log drivers are `awslogs` , `fluentd` , `gelf` , `json-file` , `journald` , `logentries` , `syslog` , `splunk` , and `awsfirelens` .
	//
	// For more information about using the `awslogs` log driver, see [Using the awslogs log driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For more information about using the `awsfirelens` log driver, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.
	LogDriver *string `json:"logDriver"`
	// The configuration options to send to the log driver.
	//
	// This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	Options interface{} `json:"options"`
	// The secrets to pass to the log configuration.
	//
	// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	SecretOptions interface{} `json:"secretOptions"`
}

// The `MountPoint` property specifies details on a volume mount point that is used in a container definition.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_MountPointProperty struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `json:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `json:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the `name` parameter of task definition `volume` .
	SourceVolume *string `json:"sourceVolume"`
}

// The `PortMapping` property specifies a port mapping.
//
// Port mappings allow containers to access ports on the host container instance to send or receive traffic. Port mappings are specified as part of the container definition.
//
// If you are using containers in a task with the `awsvpc` or `host` network mode, exposed ports should be specified using `containerPort` . The `hostPort` can be left blank or it must be the same value as the `containerPort` .
//
// After a task reaches the `RUNNING` status, manual and automatic host and container port assignments are visible in the `networkBindings` section of [DescribeTasks](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTasks.html) API responses.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_PortMappingProperty struct {
	// The port number on the container that's bound to the user-specified or automatically assigned host port.
	//
	// If you use containers in a task with the `awsvpc` or `host` network mode, specify the exposed ports using `containerPort` .
	//
	// If you use containers in a task with the `bridge` network mode and you specify a container port and not a host port, your container automatically receives a host port in the ephemeral port range. For more information, see `hostPort` . Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	ContainerPort *float64 `json:"containerPort"`
	// The port number on the container instance to reserve for your container.
	//
	// If you are using containers in a task with the `awsvpc` or `host` network mode, the `hostPort` can either be left blank or set to the same value as the `containerPort` .
	//
	// If you are using containers in a task with the `bridge` network mode, you can specify a non-reserved host port for your container port mapping, or you can omit the `hostPort` (or set it to `0` ) while specifying a `containerPort` and your container automatically receives a port in the ephemeral port range for your container instance operating system and Docker version.
	//
	// The default ephemeral port range for Docker version 1.6.0 and later is listed on the instance under `/proc/sys/net/ipv4/ip_local_port_range` . If this kernel parameter is unavailable, the default ephemeral port range from 49153 through 65535 is used. Do not attempt to specify a host port in the ephemeral port range as these are reserved for automatic assignment. In general, ports below 32768 are outside of the ephemeral port range.
	//
	// > The default ephemeral port range from 49153 through 65535 is always used for Docker versions before 1.6.0.
	//
	// The default reserved ports are 22 for SSH, the Docker ports 2375 and 2376, and the Amazon ECS container agent ports 51678-51680. Any host port that was previously specified in a running task is also reserved while the task is running (after a task stops, the host port is released). The current reserved ports are displayed in the `remainingResources` of [DescribeContainerInstances](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeContainerInstances.html) output. A container instance can have up to 100 reserved ports at a time, including the default reserved ports. Automatically assigned ports don't count toward the 100 reserved ports limit.
	HostPort *float64 `json:"hostPort"`
	// The protocol used for the port mapping.
	//
	// Valid values are `tcp` and `udp` . The default is `tcp` .
	Protocol *string `json:"protocol"`
}

// The `ProxyConfiguration` property specifies the details for the App Mesh proxy.
//
// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// For tasks using the Fargate launch type, the task or service requires platform version 1.3.0 or later.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_ProxyConfigurationProperty struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `json:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.
	//
	// - `IgnoredUID` - (Required) The user ID (UID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredGID` is specified, this field can be empty.
	// - `IgnoredGID` - (Required) The group ID (GID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredUID` is specified, this field can be empty.
	// - `AppPorts` - (Required) The list of ports that the application uses. Network traffic to these ports is forwarded to the `ProxyIngressPort` and `ProxyEgressPort` .
	// - `ProxyIngressPort` - (Required) Specifies the port that incoming traffic to the `AppPorts` is directed to.
	// - `ProxyEgressPort` - (Required) Specifies the port that outgoing traffic from the `AppPorts` is directed to.
	// - `EgressIgnoredPorts` - (Required) The egress traffic going to the specified ports is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	// - `EgressIgnoredIPs` - (Required) The egress traffic going to the specified IP addresses is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	ProxyConfigurationProperties interface{} `json:"proxyConfigurationProperties"`
	// The proxy type.
	//
	// The only supported value is `APPMESH` .
	Type *string `json:"type"`
}

// The `RepositoryCredentials` property specifies the repository credentials for private registry authentication.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_RepositoryCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	//
	// > When you use the Amazon ECS API, AWS CLI , or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.
	CredentialsParameter *string `json:"credentialsParameter"`
}

// The `ResourceRequirement` property specifies the type and amount of a resource to assign to a container.
//
// The only supported resource is a GPU. For more information, see [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html) in the *Amazon Elastic Container Service Developer Guide*
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_ResourceRequirementProperty struct {
	// The type of resource to assign to a container.
	//
	// The supported values are `GPU` or `InferenceAccelerator` .
	Type *string `json:"type"`
	// The value for the specified resource type.
	//
	// If the `GPU` type is used, the value is the number of physical `GPUs` the Amazon ECS container agent will reserve for the container. The number of GPUs reserved for all containers in a task should not exceed the number of available GPUs on the container instance the task is launched on.
	//
	// If the `InferenceAccelerator` type is used, the `value` should match the `DeviceName` for an [InferenceAccelerator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html) specified in a task definition.
	Value *string `json:"value"`
}

// Information about the platform for the Amazon ECS service or task.
//
// For more informataion about `RuntimePlatform` , see [RuntimePlatform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_RuntimePlatformProperty struct {
	// The CPU architecture.
	//
	// You can run your Linux tasks on an ARM-based platform by setting the value to `ARM64` . This option is avaiable for tasks that run on Linuc Amazon EC2 instance or Linux containers on Fargate.
	CpuArchitecture *string `json:"cpuArchitecture"`
	// The operating system.
	OperatingSystemFamily *string `json:"operatingSystemFamily"`
}

// The `Secret` property specifies an object representing the secret to expose to your container.
//
// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_SecretProperty struct {
	// The name of the secret.
	Name *string `json:"name"`
	// The secret to expose to the container.
	//
	// The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter Store.
	//
	// > If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.
	ValueFrom *string `json:"valueFrom"`
}

// A list of namespaced kernel parameters to set in the container.
//
// This parameter maps to `Sysctls` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--sysctl` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
//
// We don't recommend that you specify network-related `systemControls` parameters for multiple containers in a single task. This task also uses either the `awsvpc` or `host` network mode. It does it for the following reasons.
//
// - For tasks that use the `awsvpc` network mode, if you set `systemControls` for any container, it applies to all containers in the task. If you set different `systemControls` for multiple containers in a single task, the container that's started last determines which `systemControls` take effect.
// - For tasks that use the `host` network mode, the `systemControls` parameter applies to the container instance's kernel parameter and that of all containers of any tasks running on that container instance.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_SystemControlProperty struct {
	// The namespaced kernel parameter to set a `value` for.
	Namespace *string `json:"namespace"`
	// The value for the namespaced kernel parameter that's specified in `namespace` .
	Value *string `json:"value"`
}

// The `TaskDefinitionPlacementConstraint` property specifies an object representing a constraint on task placement in the task definition.
//
// If you are using the Fargate launch type, task placement constraints are not supported.
//
// For more information, see [Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_TaskDefinitionPlacementConstraintProperty struct {
	// The type of constraint.
	//
	// The `MemberOf` constraint restricts selection to be from a group of valid candidates.
	Type *string `json:"type"`
	// A cluster query language expression to apply to the constraint.
	//
	// For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	Expression *string `json:"expression"`
}

// The `Tmpfs` property specifies the container path, mount options, and size of the tmpfs mount.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_TmpfsProperty struct {
	// The maximum size (in MiB) of the tmpfs volume.
	Size *float64 `json:"size"`
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `json:"containerPath"`
	// The list of tmpfs volume mount options.
	//
	// Valid values: `"defaults" | "ro" | "rw" | "suid" | "nosuid" | "dev" | "nodev" | "exec" | "noexec" | "sync" | "async" | "dirsync" | "remount" | "mand" | "nomand" | "atime" | "noatime" | "diratime" | "nodiratime" | "bind" | "rbind" | "unbindable" | "runbindable" | "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime" | "norelatime" | "strictatime" | "nostrictatime" | "mode" | "uid" | "gid" | "nr_inodes" | "nr_blocks" | "mpol"`
	MountOptions *[]*string `json:"mountOptions"`
}

// The `Ulimit` property specifies the `ulimit` settings to pass to the container.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_UlimitProperty struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `json:"hardLimit"`
	// The `type` of the `ulimit` .
	Name *string `json:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `json:"softLimit"`
}

// The `VolumeFrom` property specifies details on a data volume from another container in the same task definition.
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_VolumeFromProperty struct {
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `json:"readOnly"`
	// The name of another container within the same task definition to mount volumes from.
	SourceContainer *string `json:"sourceContainer"`
}

// The `Volume` property specifies a data volume used in a task definition.
//
// For tasks that use a Docker volume, specify a `DockerVolumeConfiguration` . For tasks that use a bind mount host volume, specify a `host` and optional `sourcePath` . For more information, see [Using Data Volumes in Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) .
//
// TODO: EXAMPLE
//
type CfnTaskDefinition_VolumeProperty struct {
	// This parameter is specified when you use Docker volumes.
	//
	// Windows containers only support the use of the `local` driver. To use bind mounts, specify the `host` parameter instead.
	//
	// > Docker volumes aren't supported by tasks run on AWS Fargate .
	DockerVolumeConfiguration interface{} `json:"dockerVolumeConfiguration"`
	// `CfnTaskDefinition.VolumeProperty.EfsVolumeConfiguration`.
	EfsVolumeConfiguration interface{} `json:"efsVolumeConfiguration"`
	// This parameter is specified when you use bind mount host volumes.
	//
	// The contents of the `host` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the `host` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//
	// Windows containers can mount whole directories on the same drive as `$env:ProgramData` . Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount `C:\my\path:C:\my\path` and `D:\:D:\` , but not `D:\my\path:C:\my\path` or `D:\:C:\my\path` .
	Host interface{} `json:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	Name *string `json:"name"`
}

// Properties for defining a `CfnTaskDefinition`.
//
// TODO: EXAMPLE
//
type CfnTaskDefinitionProps struct {
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
	ContainerDefinitions interface{} `json:"containerDefinitions"`
	// The number of `cpu` units used by the task.
	//
	// If you use the EC2 launch type, this field is optional. Any value can be used. If you use the Fargate launch type, this field is required. You must use one of the following values. The value that you choose determines your range of valid values for the `memory` parameter.
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers on Fargate.
	//
	// - 256 (.25 vCPU) - Available `memory` values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	// - 512 (.5 vCPU) - Available `memory` values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	// - 1024 (1 vCPU) - Available `memory` values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	// - 2048 (2 vCPU) - Available `memory` values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	// - 4096 (4 vCPU) - Available `memory` values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)
	Cpu *string `json:"cpu"`
	// The ephemeral storage settings to use for tasks run with the task definition.
	EphemeralStorage interface{} `json:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	Family *string `json:"family"`
	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators interface{} `json:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// The valid values are `host` , `task` , or `none` . If `host` is specified, then all containers within the tasks that specified the `host` IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same IPC resources. If `none` is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance. If no value is specified, then the IPC resource namespace sharing depends on the Docker daemon setting on the container instance. For more information, see [IPC settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#ipc-settings---ipc) in the *Docker run reference* .
	//
	// If the `host` IPC mode is used, be aware that there is a heightened risk of undesired IPC namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// If you are setting namespaced kernel parameters using `systemControls` for the containers in the task, the following will apply to your IPC resource namespace. For more information, see [System Controls](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// - For tasks that use the `host` IPC mode, IPC namespace related `systemControls` are not supported.
	// - For tasks that use the `task` IPC mode, IPC namespace related `systemControls` will apply to all containers within a task.
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	IpcMode *string `json:"ipcMode"`
	// The amount (in MiB) of memory used by the task.
	//
	// If your tasks runs on Amazon EC2 instances, you must specify either a task-level memory value or a container-level memory value. This field is optional and any value can be used. If a task-level memory value is specified, the container-level memory value is optional. For more information regarding container-level memory and memory reservation, see [ContainerDefinition](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) .
	//
	// If your tasks runs on AWS Fargate , this field is required. You must use one of the following values. The value you choose determines your range of valid values for the `cpu` parameter.
	//
	// - 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available `cpu` values: 256 (.25 vCPU)
	// - 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available `cpu` values: 512 (.5 vCPU)
	// - 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available `cpu` values: 1024 (1 vCPU)
	// - Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available `cpu` values: 2048 (2 vCPU)
	// - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available `cpu` values: 4096 (4 vCPU)
	Memory *string `json:"memory"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are `none` , `bridge` , `awsvpc` , and `host` . The default Docker network mode is `bridge` . If you are using the Fargate launch type, the `awsvpc` network mode is required. If you are using the EC2 launch type, any network mode can be used. If the network mode is set to `none` , you cannot specify port mappings in your container definitions, and the tasks containers do not have external connectivity. The `host` and `awsvpc` network modes offer the highest networking performance for containers because they use the EC2 network stack instead of the virtualized network stack provided by the `bridge` mode.
	//
	// With the `host` and `awsvpc` network modes, exposed container ports are mapped directly to the corresponding host port (for the `host` network mode) or the attached elastic network interface port (for the `awsvpc` network mode), so you cannot take advantage of dynamic host port mappings.
	//
	// If the network mode is `awsvpc` , the task is allocated an elastic network interface, and you must specify a [NetworkConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkConfiguration.html) value when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > Currently, only Amazon ECS-optimized AMIs, other Amazon Linux variants with the `ecs-init` package, or AWS Fargate infrastructure support the `awsvpc` network mode.
	//
	// If the network mode is `host` , you cannot run multiple instantiations of the same task on a single container instance when port mappings are used.
	//
	// Docker for Windows uses different network modes than Docker for Linux. When you register a task definition with Windows containers, you must not specify a network mode. If you use the console to register a task definition with Windows containers, you must choose the `<default>` network mode object.
	//
	// For more information, see [Network settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#network-settings) in the *Docker run reference* .
	NetworkMode *string `json:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . If `host` is specified, then all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same process namespace. If no value is specified, the default is a private namespace. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, be aware that there is a heightened risk of undesired process namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	PidMode *string `json:"pidMode"`
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	PlacementConstraints interface{} `json:"placementConstraints"`
	// The `ProxyConfiguration` property specifies the configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	ProxyConfiguration interface{} `json:"proxyConfiguration"`
	// The task launch types the task definition was validated against.
	//
	// To determine which task launch types the task definition is validated for, see the `TaskDefinition$compatibilities` parameter.
	RequiresCompatibilities *[]*string `json:"requiresCompatibilities"`
	// The operating system that your task definitions are running on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task in a service, this value must match the `runtimePlatform` value of the service.
	RuntimePlatform interface{} `json:"runtimePlatform"`
	// The metadata that you apply to the task definition to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value. You define both of them.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn *string `json:"taskRoleArn"`
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	Volumes interface{} `json:"volumes"`
}

// A CloudFormation `AWS::ECS::TaskSet`.
//
// Create a task set in the specified cluster and service. This is used when a service uses the `EXTERNAL` deployment controller type. For more information, see [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTaskSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Cluster() *string
	SetCluster(val *string)
	CreationStack() *[]*string
	ExternalId() *string
	SetExternalId(val *string)
	LaunchType() *string
	SetLaunchType(val *string)
	LoadBalancers() interface{}
	SetLoadBalancers(val interface{})
	LogicalId() *string
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	Node() constructs.Node
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	Ref() *string
	Scale() interface{}
	SetScale(val interface{})
	Service() *string
	SetService(val *string)
	ServiceRegistries() interface{}
	SetServiceRegistries(val interface{})
	Stack() awscdk.Stack
	TaskDefinition() *string
	SetTaskDefinition(val *string)
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

// The jsii proxy struct for CfnTaskSet
type jsiiProxy_CfnTaskSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTaskSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) LoadBalancers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Scale() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) ServiceRegistries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::TaskSet`.
func NewCfnTaskSet(scope constructs.Construct, id *string, props *CfnTaskSetProps) CfnTaskSet {
	_init_.Initialize()

	j := jsiiProxy_CfnTaskSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::TaskSet`.
func NewCfnTaskSet_Override(c CfnTaskSet, scope constructs.Construct, id *string, props *CfnTaskSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetExternalId(val *string) {
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetLaunchType(val *string) {
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetLoadBalancers(val interface{}) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetPlatformVersion(val *string) {
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetScale(val interface{}) {
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetService(val *string) {
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetServiceRegistries(val interface{}) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

func (j *jsiiProxy_CfnTaskSet) SetTaskDefinition(val *string) {
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTaskSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTaskSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
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
func CfnTaskSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnTaskSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTaskSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTaskSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTaskSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTaskSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTaskSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTaskSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTaskSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTaskSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTaskSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTaskSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTaskSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTaskSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTaskSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTaskSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The networking details for a task.
//
// TODO: EXAMPLE
//
type CfnTaskSet_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	Subnets *[]*string `json:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	AssignPublicIp *string `json:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	SecurityGroups *[]*string `json:"securityGroups"`
}

// Details on the load balancer or load balancers to use with a task set.
//
// TODO: EXAMPLE
//
type CfnTaskSet_LoadBalancerProperty struct {
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	ContainerName *string `json:"containerName"`
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	ContainerPort *float64 `json:"containerPort"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you're using a Classic Load Balancer, omit the target group ARN.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering Multiple Target Groups with a Service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/Green Deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	TargetGroupArn *string `json:"targetGroupArn"`
}

// The network configuration for a task.
//
// TODO: EXAMPLE
//
type CfnTaskSet_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	AwsVpcConfiguration interface{} `json:"awsVpcConfiguration"`
}

// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
//
// TODO: EXAMPLE
//
type CfnTaskSet_ScaleProperty struct {
	// The unit of measure for the scale value.
	Unit *string `json:"unit"`
	// The value, specified as a percent total of a service's `desiredCount` , to scale the task set.
	//
	// Accepted values are numbers between 0 and 100.
	Value *float64 `json:"value"`
}

// The details for the service registry.
//
// TODO: EXAMPLE
//
type CfnTaskSet_ServiceRegistryProperty struct {
	// The container name value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition that your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition that your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerName *string `json:"containerName"`
	// The port value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerPort *float64 `json:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field might be used if both the `awsvpc` network mode and SRV records are used.
	Port *float64 `json:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map . For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) .
	RegistryArn *string `json:"registryArn"`
}

// Properties for defining a `CfnTaskSet`.
//
// TODO: EXAMPLE
//
type CfnTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	Cluster *string `json:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	Service *string `json:"service"`
	// The task definition for the tasks in the task set to use.
	TaskDefinition *string `json:"taskDefinition"`
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map attribute set to the provided value.
	ExternalId *string `json:"externalId"`
	// The launch type that new tasks in the task set uses.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If a `launchType` is specified, the `capacityProviderStrategy` parameter must be omitted.
	LaunchType *string `json:"launchType"`
	// A load balancer object representing the load balancer to use with the task set.
	//
	// The supported load balancer types are either an Application Load Balancer or a Network Load Balancer.
	LoadBalancers interface{} `json:"loadBalancers"`
	// The network configuration for the task set.
	NetworkConfiguration interface{} `json:"networkConfiguration"`
	// The platform version that the tasks in the task set uses.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used.
	PlatformVersion *string `json:"platformVersion"`
	// A floating-point percentage of your desired number of tasks to place and keep running in the task set.
	Scale interface{} `json:"scale"`
	// The details of the service discovery registries to assign to this task set.
	//
	// For more information, see [Service Discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	ServiceRegistries interface{} `json:"serviceRegistries"`
}

// The options for creating an AWS Cloud Map namespace.
//
// TODO: EXAMPLE
//
type CloudMapNamespaceOptions struct {
	// The name of the namespace, such as example.com.
	Name *string `json:"name"`
	// The type of CloudMap Namespace to create.
	Type awsservicediscovery.NamespaceType `json:"type"`
	// The VPC to associate the namespace with.
	//
	// This property is required for private DNS namespaces.
	Vpc awsec2.IVpc `json:"vpc"`
}

// The options to enabling AWS Cloud Map for an Amazon ECS service.
//
// TODO: EXAMPLE
//
type CloudMapOptions struct {
	// The service discovery namespace for the Cloud Map service to attach to the ECS service.
	CloudMapNamespace awsservicediscovery.INamespace `json:"cloudMapNamespace"`
	// The container to point to for a SRV record.
	Container ContainerDefinition `json:"container"`
	// The port to point to for a SRV record.
	ContainerPort *float64 `json:"containerPort"`
	// The DNS record type that you want AWS Cloud Map to create.
	//
	// The supported record types are A or SRV.
	DnsRecordType awsservicediscovery.DnsRecordType `json:"dnsRecordType"`
	// The amount of time that you want DNS resolvers to cache the settings for this record.
	DnsTtl awscdk.Duration `json:"dnsTtl"`
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	//
	// NOTE: This is used for HealthCheckCustomConfig
	FailureThreshold *float64 `json:"failureThreshold"`
	// The name of the Cloud Map service to attach to the ECS service.
	Name *string `json:"name"`
}

// A regional grouping of one or more container instances on which you can run tasks and services.
//
// TODO: EXAMPLE
//
type Cluster interface {
	awscdk.Resource
	ICluster
	AutoscalingGroup() awsautoscaling.IAutoScalingGroup
	ClusterArn() *string
	ClusterName() *string
	Connections() awsec2.Connections
	DefaultCloudMapNamespace() awsservicediscovery.INamespace
	Env() *awscdk.ResourceEnvironment
	ExecuteCommandConfiguration() *ExecuteCommandConfiguration
	HasEc2Capacity() *bool
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions)
	AddCapacity(id *string, options *AddCapacityOptions) awsautoscaling.AutoScalingGroup
	AddDefaultCloudMapNamespace(options *CloudMapNamespaceOptions) awsservicediscovery.INamespace
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	EnableFargateCapacityProviders()
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) AutoscalingGroup() awsautoscaling.IAutoScalingGroup {
	var returns awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoscalingGroup",
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

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
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

func (j *jsiiProxy_Cluster) DefaultCloudMapNamespace() awsservicediscovery.INamespace {
	var returns awsservicediscovery.INamespace
	_jsii_.Get(
		j,
		"defaultCloudMapNamespace",
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

func (j *jsiiProxy_Cluster) ExecuteCommandConfiguration() *ExecuteCommandConfiguration {
	var returns *ExecuteCommandConfiguration
	_jsii_.Get(
		j,
		"executeCommandConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) HasEc2Capacity() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasEc2Capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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


// Constructs a new instance of the Cluster class.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the Cluster class.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing cluster to the stack from its attributes.
func Cluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"fromClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Cluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// This method adds an Auto Scaling Group Capacity Provider to a cluster.
func (c *jsiiProxy_Cluster) AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions) {
	_jsii_.InvokeVoid(
		c,
		"addAsgCapacityProvider",
		[]interface{}{provider, options},
	)
}

// It is highly recommended to use {@link Cluster.addAsgCapacityProvider} instead of this method.
//
// This method adds compute capacity to a cluster by creating an AutoScalingGroup with the specified options.
//
// Returns the AutoScalingGroup so you can add autoscaling settings to it.
func (c *jsiiProxy_Cluster) AddCapacity(id *string, options *AddCapacityOptions) awsautoscaling.AutoScalingGroup {
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		c,
		"addCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Add an AWS Cloud Map DNS namespace for this cluster.
//
// NOTE: HttpNamespaces are not supported, as ECS always requires a DNSConfig when registering an instance to a Cloud
// Map service.
func (c *jsiiProxy_Cluster) AddDefaultCloudMapNamespace(options *CloudMapNamespaceOptions) awsservicediscovery.INamespace {
	var returns awsservicediscovery.INamespace

	_jsii_.Invoke(
		c,
		"addDefaultCloudMapNamespace",
		[]interface{}{options},
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
func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Enable the Fargate capacity providers for this cluster.
func (c *jsiiProxy_Cluster) EnableFargateCapacityProviders() {
	_jsii_.InvokeVoid(
		c,
		"enableFargateCapacityProviders",
		nil, // no parameters
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
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

// This method returns the specifed CloudWatch metric for this cluster.
func (c *jsiiProxy_Cluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this clusters CPU reservation.
func (c *jsiiProxy_Cluster) MetricCpuReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricCpuReservation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this clusters CPU utilization.
func (c *jsiiProxy_Cluster) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this clusters memory reservation.
func (c *jsiiProxy_Cluster) MetricMemoryReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricMemoryReservation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this clusters memory utilization.
func (c *jsiiProxy_Cluster) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
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

// The properties to import from the ECS cluster.
//
// TODO: EXAMPLE
//
type ClusterAttributes struct {
	// The name of the cluster.
	ClusterName *string `json:"clusterName"`
	// The security groups associated with the container instances registered to the cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The VPC associated with the cluster.
	Vpc awsec2.IVpc `json:"vpc"`
	// Autoscaling group added to the cluster if capacity is added.
	AutoscalingGroup awsautoscaling.IAutoScalingGroup `json:"autoscalingGroup"`
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn *string `json:"clusterArn"`
	// The AWS Cloud Map namespace to associate with the cluster.
	DefaultCloudMapNamespace awsservicediscovery.INamespace `json:"defaultCloudMapNamespace"`
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `json:"executeCommandConfiguration"`
	// Specifies whether the cluster has EC2 instance capacity.
	HasEc2Capacity *bool `json:"hasEc2Capacity"`
}

// The properties used to define an ECS cluster.
//
// TODO: EXAMPLE
//
type ClusterProps struct {
	// The ec2 capacity to add to the cluster.
	Capacity *AddCapacityOptions `json:"capacity"`
	// The name for the cluster.
	ClusterName *string `json:"clusterName"`
	// If true CloudWatch Container Insights will be enabled for the cluster.
	ContainerInsights *bool `json:"containerInsights"`
	// The service discovery namespace created in this cluster.
	DefaultCloudMapNamespace *CloudMapNamespaceOptions `json:"defaultCloudMapNamespace"`
	// Whether to enable Fargate Capacity Providers.
	EnableFargateCapacityProviders *bool `json:"enableFargateCapacityProviders"`
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `json:"executeCommandConfiguration"`
	// The VPC where your ECS instances will be running or your ENIs will be deployed.
	Vpc awsec2.IVpc `json:"vpc"`
}

// The common task definition attributes used across all types of task definitions.
//
// TODO: EXAMPLE
//
type CommonTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `json:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `json:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// The common properties for all task definitions.
//
// For more information, see
// [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html).
//
// TODO: EXAMPLE
//
type CommonTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `json:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `json:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `json:"volumes"`
}

// The task launch type compatibility requirement.
//
// TODO: EXAMPLE
//
type Compatibility string

const (
	Compatibility_EC2 Compatibility = "EC2"
	Compatibility_FARGATE Compatibility = "FARGATE"
	Compatibility_EC2_AND_FARGATE Compatibility = "EC2_AND_FARGATE"
	Compatibility_EXTERNAL Compatibility = "EXTERNAL"
)

// A container definition is used in a task definition to describe the containers that are launched as part of a task.
//
// TODO: EXAMPLE
//
type ContainerDefinition interface {
	constructs.Construct
	ContainerDependencies() *[]*ContainerDependency
	ContainerName() *string
	ContainerPort() *float64
	EnvironmentFiles() *[]*EnvironmentFileConfig
	Essential() *bool
	IngressPort() *float64
	LinuxParameters() LinuxParameters
	LogDriverConfig() *LogDriverConfig
	MemoryLimitSpecified() *bool
	MountPoints() *[]*MountPoint
	Node() constructs.Node
	PortMappings() *[]*PortMapping
	ReferencesSecretJsonField() *bool
	TaskDefinition() TaskDefinition
	Ulimits() *[]*Ulimit
	VolumesFrom() *[]*VolumeFrom
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	AddEnvironment(name *string, value *string)
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	AddLink(container ContainerDefinition, alias *string)
	AddMountPoints(mountPoints ...*MountPoint)
	AddPortMappings(portMappings ...*PortMapping)
	AddScratch(scratch *ScratchSpace)
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	AddUlimits(ulimits ...*Ulimit)
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	ToString() *string
}

// The jsii proxy struct for ContainerDefinition
type jsiiProxy_ContainerDefinition struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ContainerDefinition) ContainerDependencies() *[]*ContainerDependency {
	var returns *[]*ContainerDependency
	_jsii_.Get(
		j,
		"containerDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) EnvironmentFiles() *[]*EnvironmentFileConfig {
	var returns *[]*EnvironmentFileConfig
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Essential() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) IngressPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) LogDriverConfig() *LogDriverConfig {
	var returns *LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) MemoryLimitSpecified() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"memoryLimitSpecified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) MountPoints() *[]*MountPoint {
	var returns *[]*MountPoint
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) PortMappings() *[]*PortMapping {
	var returns *[]*PortMapping
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) VolumesFrom() *[]*VolumeFrom {
	var returns *[]*VolumeFrom
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ContainerDefinition class.
func NewContainerDefinition(scope constructs.Construct, id *string, props *ContainerDefinitionProps) ContainerDefinition {
	_init_.Initialize()

	j := jsiiProxy_ContainerDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ContainerDefinition class.
func NewContainerDefinition_Override(c ContainerDefinition, scope constructs.Construct, id *string, props *ContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// This method adds one or more container dependencies to the container.
func (c *jsiiProxy_ContainerDefinition) AddContainerDependencies(containerDependencies ...*ContainerDependency) {
	args := []interface{}{}
	for _, a := range containerDependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addContainerDependencies",
		args,
	)
}

// This method adds an environment variable to the container.
func (c *jsiiProxy_ContainerDefinition) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

// This method adds one or more resources to the container.
func (c *jsiiProxy_ContainerDefinition) AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string) {
	args := []interface{}{}
	for _, a := range inferenceAcceleratorResources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addInferenceAcceleratorResource",
		args,
	)
}

// This method adds a link which allows containers to communicate with each other without the need for port mappings.
//
// This parameter is only supported if the task definition is using the bridge network mode.
// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
func (c *jsiiProxy_ContainerDefinition) AddLink(container ContainerDefinition, alias *string) {
	_jsii_.InvokeVoid(
		c,
		"addLink",
		[]interface{}{container, alias},
	)
}

// This method adds one or more mount points for data volumes to the container.
func (c *jsiiProxy_ContainerDefinition) AddMountPoints(mountPoints ...*MountPoint) {
	args := []interface{}{}
	for _, a := range mountPoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addMountPoints",
		args,
	)
}

// This method adds one or more port mappings to the container.
func (c *jsiiProxy_ContainerDefinition) AddPortMappings(portMappings ...*PortMapping) {
	args := []interface{}{}
	for _, a := range portMappings {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addPortMappings",
		args,
	)
}

// This method mounts temporary disk space to the container.
//
// This adds the correct container mountPoint and task definition volume.
func (c *jsiiProxy_ContainerDefinition) AddScratch(scratch *ScratchSpace) {
	_jsii_.InvokeVoid(
		c,
		"addScratch",
		[]interface{}{scratch},
	)
}

// This method adds the specified statement to the IAM task execution policy in the task definition.
func (c *jsiiProxy_ContainerDefinition) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		c,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

// This method adds one or more ulimits to the container.
func (c *jsiiProxy_ContainerDefinition) AddUlimits(ulimits ...*Ulimit) {
	args := []interface{}{}
	for _, a := range ulimits {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUlimits",
		args,
	)
}

// This method adds one or more volumes to the container.
func (c *jsiiProxy_ContainerDefinition) AddVolumesFrom(volumesFrom ...*VolumeFrom) {
	args := []interface{}{}
	for _, a := range volumesFrom {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addVolumesFrom",
		args,
	)
}

// Returns the host port for the requested container port if it exists.
func (c *jsiiProxy_ContainerDefinition) FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping {
	var returns *PortMapping

	_jsii_.Invoke(
		c,
		"findPortMapping",
		[]interface{}{containerPort, protocol},
		&returns,
	)

	return returns
}

// Render this container definition to a CloudFormation object.
func (c *jsiiProxy_ContainerDefinition) RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty {
	var returns *CfnTaskDefinition_ContainerDefinitionProperty

	_jsii_.Invoke(
		c,
		"renderContainerDefinition",
		[]interface{}{_taskDefinition},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type ContainerDefinitionOptions struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage
	Image ContainerImage `json:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `json:"command"`
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `json:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `json:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `json:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `json:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `json:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `json:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `json:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `json:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `json:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `json:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `json:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `json:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `json:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `json:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `json:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `json:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `json:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `json:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `json:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `json:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `json:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `json:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `json:"systemControls"`
	// The user name to use inside the container.
	User *string `json:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `json:"workingDirectory"`
}

// The properties in a container definition.
//
// TODO: EXAMPLE
//
type ContainerDefinitionProps struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage
	Image ContainerImage `json:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `json:"command"`
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `json:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `json:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `json:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `json:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `json:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `json:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `json:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `json:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `json:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `json:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `json:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `json:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `json:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `json:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `json:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `json:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `json:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `json:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `json:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `json:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `json:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `json:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `json:"systemControls"`
	// The user name to use inside the container.
	User *string `json:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `json:"workingDirectory"`
	// The name of the task definition that includes this container definition.
	//
	// [disable-awslint:ref-via-interface]
	TaskDefinition TaskDefinition `json:"taskDefinition"`
}

// The details of a dependency on another container in the task definition.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDependency.html
//
type ContainerDependency struct {
	// The container to depend on.
	Container ContainerDefinition `json:"container"`
	// The state the container needs to be in to satisfy the dependency and proceed with startup.
	//
	// Valid values are ContainerDependencyCondition.START, ContainerDependencyCondition.COMPLETE,
	// ContainerDependencyCondition.SUCCESS and ContainerDependencyCondition.HEALTHY.
	Condition ContainerDependencyCondition `json:"condition"`
}

type ContainerDependencyCondition string

const (
	ContainerDependencyCondition_START ContainerDependencyCondition = "START"
	ContainerDependencyCondition_COMPLETE ContainerDependencyCondition = "COMPLETE"
	ContainerDependencyCondition_SUCCESS ContainerDependencyCondition = "SUCCESS"
	ContainerDependencyCondition_HEALTHY ContainerDependencyCondition = "HEALTHY"
)

// Constructs for types of container images.
//
// TODO: EXAMPLE
//
type ContainerImage interface {
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for ContainerImage
type jsiiProxy_ContainerImage struct {
	_ byte // padding
}

func NewContainerImage_Override(c ContainerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		nil, // no parameters
		c,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func ContainerImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func ContainerImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func ContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func ContainerImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func ContainerImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

// Called when the image is used by a ContainerDefinition.
func (c *jsiiProxy_ContainerImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// The configuration for creating a container image.
//
// TODO: EXAMPLE
//
type ContainerImageConfig struct {
	// Specifies the name of the container image.
	ImageName *string `json:"imageName"`
	// Specifies the credentials used to access the image repository.
	RepositoryCredentials *CfnTaskDefinition_RepositoryCredentialsProperty `json:"repositoryCredentials"`
}

// The CpuArchitecture for Fargate Runtime Platform.
//
// TODO: EXAMPLE
//
type CpuArchitecture interface {
}

// The jsii proxy struct for CpuArchitecture
type jsiiProxy_CpuArchitecture struct {
	_ byte // padding
}

// Other cpu architecture.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-cpuarchitecture for all available cpu architecture.
//
func CpuArchitecture_Of(cpuArchitecture *string) CpuArchitecture {
	_init_.Initialize()

	var returns CpuArchitecture

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"of",
		[]interface{}{cpuArchitecture},
		&returns,
	)

	return returns
}

func CpuArchitecture_ARM64() CpuArchitecture {
	_init_.Initialize()
	var returns CpuArchitecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"ARM64",
		&returns,
	)
	return returns
}

func CpuArchitecture_X86_64() CpuArchitecture {
	_init_.Initialize()
	var returns CpuArchitecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"X86_64",
		&returns,
	)
	return returns
}

// The properties for enabling scaling based on CPU utilization.
//
// TODO: EXAMPLE
//
type CpuUtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// The target value for CPU utilization across all tasks in the service.
	TargetUtilizationPercent *float64 `json:"targetUtilizationPercent"`
}

// The deployment circuit breaker to use for the service.
//
// TODO: EXAMPLE
//
type DeploymentCircuitBreaker struct {
	// Whether to enable rollback on deployment failure.
	Rollback *bool `json:"rollback"`
}

// The deployment controller to use for the service.
//
// TODO: EXAMPLE
//
type DeploymentController struct {
	// The deployment controller type to use.
	Type DeploymentControllerType `json:"type"`
}

// The deployment controller type to use for the service.
//
// TODO: EXAMPLE
//
type DeploymentControllerType string

const (
	DeploymentControllerType_ECS DeploymentControllerType = "ECS"
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerType_EXTERNAL DeploymentControllerType = "EXTERNAL"
)

// A container instance host device.
//
// TODO: EXAMPLE
//
type Device struct {
	// The path for the device on the host container instance.
	HostPath *string `json:"hostPath"`
	// The path inside the container at which to expose the host device.
	ContainerPath *string `json:"containerPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for read, write, and mknod for the device.
	Permissions *[]DevicePermission `json:"permissions"`
}

// Permissions for device access.
type DevicePermission string

const (
	DevicePermission_READ DevicePermission = "READ"
	DevicePermission_WRITE DevicePermission = "WRITE"
	DevicePermission_MKNOD DevicePermission = "MKNOD"
)

// The configuration for a Docker volume.
//
// Docker volumes are only supported when you are using the EC2 launch type.
//
// TODO: EXAMPLE
//
type DockerVolumeConfiguration struct {
	// The Docker volume driver to use.
	Driver *string `json:"driver"`
	// The scope for the Docker volume that determines its lifecycle.
	Scope Scope `json:"scope"`
	// Specifies whether the Docker volume should be created if it does not already exist.
	//
	// If true is specified, the Docker volume will be created for you.
	Autoprovision *bool `json:"autoprovision"`
	// A map of Docker driver-specific options passed through.
	DriverOpts *map[string]*string `json:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	Labels *map[string]*string `json:"labels"`
}

// This creates a service using the EC2 launch type on an ECS cluster.
//
// TODO: EXAMPLE
//
type Ec2Service interface {
	BaseService
	IEc2Service
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	CloudMapService() awsservicediscovery.IService
	Cluster() ICluster
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	Node() constructs.Node
	PhysicalName() *string
	ServiceArn() *string
	ServiceName() *string
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	Stack() awscdk.Stack
	TaskDefinition() TaskDefinition
	AddPlacementConstraints(constraints ...PlacementConstraint)
	AddPlacementStrategies(strategies ...PlacementStrategy)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	ToString() *string
}

// The jsii proxy struct for Ec2Service
type jsiiProxy_Ec2Service struct {
	jsiiProxy_BaseService
	jsiiProxy_IEc2Service
}

func (j *jsiiProxy_Ec2Service) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the Ec2Service class.
func NewEc2Service(scope constructs.Construct, id *string, props *Ec2ServiceProps) Ec2Service {
	_init_.Initialize()

	j := jsiiProxy_Ec2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the Ec2Service class.
func NewEc2Service_Override(e Ec2Service, scope constructs.Construct, id *string, props *Ec2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_Ec2Service) SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_Ec2Service) SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_Ec2Service) SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_Ec2Service) SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Imports from the specified service ARN.
func Ec2Service_FromEc2ServiceArn(scope constructs.Construct, id *string, ec2ServiceArn *string) IEc2Service {
	_init_.Initialize()

	var returns IEc2Service

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"fromEc2ServiceArn",
		[]interface{}{scope, id, ec2ServiceArn},
		&returns,
	)

	return returns
}

// Imports from the specified service attrributes.
func Ec2Service_FromEc2ServiceAttributes(scope constructs.Construct, id *string, attrs *Ec2ServiceAttributes) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"fromEc2ServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Ec2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Ec2Service_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds one or more placement constraints to use for tasks in the service.
//
// For more information, see
// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
func (e *jsiiProxy_Ec2Service) AddPlacementConstraints(constraints ...PlacementConstraint) {
	args := []interface{}{}
	for _, a := range constraints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraints",
		args,
	)
}

// Adds one or more placement strategies to use for tasks in the service.
//
// For more information, see
// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
func (e *jsiiProxy_Ec2Service) AddPlacementStrategies(strategies ...PlacementStrategy) {
	args := []interface{}{}
	for _, a := range strategies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlacementStrategies",
		args,
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
func (e *jsiiProxy_Ec2Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associates this service with a CloudMap service.
func (e *jsiiProxy_Ec2Service) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		e,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

// This method is called to attach this service to an Application Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (e *jsiiProxy_Ec2Service) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Registers the service as a target of a Classic Load Balancer (CLB).
//
// Don't call this. Call `loadBalancer.addTarget()` instead.
func (e *jsiiProxy_Ec2Service) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		e,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// This method is called to attach this service to a Network Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (e *jsiiProxy_Ec2Service) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
func (e *jsiiProxy_Ec2Service) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	var returns ScalableTaskCount

	_jsii_.Invoke(
		e,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method is called to create a networkConfiguration.
func (e *jsiiProxy_Ec2Service) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

// Enable CloudMap service discovery for the service.
//
// Returns: The created CloudMap service
func (e *jsiiProxy_Ec2Service) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		e,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_Ec2Service) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_Ec2Service) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return a load balancing target for a specific container and port.
//
// Use this function to create a load balancer target if you want to load balance to
// another container than the first essential container or the first mapped port on
// the container.
//
// Use the return value of this function where you would normally use a load balancer
// target, instead of the `Service` object itself.
//
// TODO: EXAMPLE
//
func (e *jsiiProxy_Ec2Service) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		e,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// This method returns the specified CloudWatch metric name for this service.
func (e *jsiiProxy_Ec2Service) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's CPU utilization.
func (e *jsiiProxy_Ec2Service) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's memory utilization.
func (e *jsiiProxy_Ec2Service) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
//
// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
//
// TODO: EXAMPLE
//
func (e *jsiiProxy_Ec2Service) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"registerLoadBalancerTargets",
		args,
	)
}

// Returns a string representation of this construct.
func (e *jsiiProxy_Ec2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties to import from the service using the EC2 launch type.
//
// TODO: EXAMPLE
//
type Ec2ServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// The service ARN.
	ServiceArn *string `json:"serviceArn"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
}

// The properties for defining a service using the EC2 launch type.
//
// TODO: EXAMPLE
//
type Ec2ServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `json:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface]
	TaskDefinition TaskDefinition `json:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// Specifies whether the service will use the daemon scheduling strategy.
	//
	// If true, the service scheduler deploys exactly one task on each container instance in your cluster.
	//
	// When you are using this strategy, do not specify a desired number of tasks orany task placement strategies.
	Daemon *bool `json:"daemon"`
	// The placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	PlacementConstraints *[]PlacementConstraint `json:"placementConstraints"`
	// The placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	PlacementStrategies *[]PlacementStrategy `json:"placementStrategies"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The subnets to associate with the service.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// The details of a task definition run on an EC2 cluster.
//
// TODO: EXAMPLE
//
type Ec2TaskDefinition interface {
	TaskDefinition
	IEc2TaskDefinition
	Compatibility() Compatibility
	Containers() *[]ContainerDefinition
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	Env() *awscdk.ResourceEnvironment
	EphemeralStorageGiB() *float64
	ExecutionRole() awsiam.IRole
	Family() *string
	InferenceAccelerators() *[]*InferenceAccelerator
	IsEc2Compatible() *bool
	IsExternalCompatible() *bool
	IsFargateCompatible() *bool
	NetworkMode() NetworkMode
	Node() constructs.Node
	PhysicalName() *string
	ReferencesSecretJsonField() *bool
	Stack() awscdk.Stack
	TaskDefinitionArn() *string
	TaskRole() awsiam.IRole
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	AddExtension(extension ITaskDefinitionExtension)
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	AddPlacementConstraint(constraint PlacementConstraint)
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	AddVolume(volume *Volume)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	FindContainer(containerName *string) ContainerDefinition
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ObtainExecutionRole() awsiam.IRole
	ToString() *string
}

// The jsii proxy struct for Ec2TaskDefinition
type jsiiProxy_Ec2TaskDefinition struct {
	jsiiProxy_TaskDefinition
	jsiiProxy_IEc2TaskDefinition
}

func (j *jsiiProxy_Ec2TaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) Containers() *[]ContainerDefinition {
	var returns *[]ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) DefaultContainer() ContainerDefinition {
	var returns ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) InferenceAccelerators() *[]*InferenceAccelerator {
	var returns *[]*InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


// Constructs a new instance of the Ec2TaskDefinition class.
func NewEc2TaskDefinition(scope constructs.Construct, id *string, props *Ec2TaskDefinitionProps) Ec2TaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_Ec2TaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the Ec2TaskDefinition class.
func NewEc2TaskDefinition_Override(e Ec2TaskDefinition, scope constructs.Construct, id *string, props *Ec2TaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_Ec2TaskDefinition) SetDefaultContainer(val ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
func Ec2TaskDefinition_FromEc2TaskDefinitionArn(scope constructs.Construct, id *string, ec2TaskDefinitionArn *string) IEc2TaskDefinition {
	_init_.Initialize()

	var returns IEc2TaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"fromEc2TaskDefinitionArn",
		[]interface{}{scope, id, ec2TaskDefinitionArn},
		&returns,
	)

	return returns
}

// Imports an existing Ec2 task definition from its attributes.
func Ec2TaskDefinition_FromEc2TaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *Ec2TaskDefinitionAttributes) IEc2TaskDefinition {
	_init_.Initialize()

	var returns IEc2TaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"fromEc2TaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func Ec2TaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func Ec2TaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"fromTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Ec2TaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Ec2TaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2TaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a new container to the task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		e,
		"addContainer",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds the specified extension to the task definition.
//
// Extension can be used to apply a packaged modification to
// a task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		e,
		"addExtension",
		[]interface{}{extension},
	)
}

// Adds a firelens log router to the task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter {
	var returns FirelensLogRouter

	_jsii_.Invoke(
		e,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds an inference accelerator to the task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		e,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

// Adds the specified placement constraint to the task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

// Adds a policy statement to the task execution IAM role.
func (e *jsiiProxy_Ec2TaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a policy statement to the task IAM role.
func (e *jsiiProxy_Ec2TaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a volume to the task definition.
func (e *jsiiProxy_Ec2TaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
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
func (e *jsiiProxy_Ec2TaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns the container that match the provided containerName.
func (e *jsiiProxy_Ec2TaskDefinition) FindContainer(containerName *string) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		e,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_Ec2TaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_Ec2TaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Creates the task execution IAM role if it doesn't already exist.
func (e *jsiiProxy_Ec2TaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		e,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_Ec2TaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes used to import an existing EC2 task definition.
//
// TODO: EXAMPLE
//
type Ec2TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `json:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `json:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// The properties for a task definition run on an EC2 cluster.
//
// TODO: EXAMPLE
//
type Ec2TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `json:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `json:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `json:"volumes"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	InferenceAccelerators *[]*InferenceAccelerator `json:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	IpcMode IpcMode `json:"ipcMode"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are none, bridge, awsvpc, and host.
	NetworkMode NetworkMode `json:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	PidMode PidMode `json:"pidMode"`
	// An array of placement constraint objects to use for the task.
	//
	// You can
	// specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	PlacementConstraints *[]PlacementConstraint `json:"placementConstraints"`
}

// An image from an Amazon ECR repository.
//
// TODO: EXAMPLE
//
type EcrImage interface {
	ContainerImage
	ImageName() *string
	Bind(_scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for EcrImage
type jsiiProxy_EcrImage struct {
	jsiiProxy_ContainerImage
}

func (j *jsiiProxy_EcrImage) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}


// Constructs a new instance of the EcrImage class.
func NewEcrImage(repository awsecr.IRepository, tagOrDigest *string) EcrImage {
	_init_.Initialize()

	j := jsiiProxy_EcrImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcrImage class.
func NewEcrImage_Override(e EcrImage, repository awsecr.IRepository, tagOrDigest *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		e,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func EcrImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func EcrImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func EcrImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func EcrImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func EcrImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

// Called when the image is used by a ContainerDefinition.
func (e *jsiiProxy_EcrImage) Bind(_scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, containerDefinition},
		&returns,
	)

	return returns
}

// Construct a Linux or Windows machine image from the latest ECS Optimized AMI published in SSM.
//
// TODO: EXAMPLE
//
type EcsOptimizedImage interface {
	awsec2.IMachineImage
	GetImage(scope constructs.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EcsOptimizedImage
type jsiiProxy_EcsOptimizedImage struct {
	internal.Type__awsec2IMachineImage
}

// Construct an Amazon Linux AMI image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_AmazonLinux(options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"amazonLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct an Amazon Linux 2 image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_AmazonLinux2(hardwareType AmiHardwareType, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"amazonLinux2",
		[]interface{}{hardwareType, options},
		&returns,
	)

	return returns
}

// Construct a Windows image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_Windows(windowsVersion WindowsOptimizedVersion, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"windows",
		[]interface{}{windowsVersion, options},
		&returns,
	)

	return returns
}

// Return the correct image.
func (e *jsiiProxy_EcsOptimizedImage) GetImage(scope constructs.Construct) *awsec2.MachineImageConfig {
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		e,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Additional configuration properties for EcsOptimizedImage factory functions.
//
// TODO: EXAMPLE
//
type EcsOptimizedImageOptions struct {
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	CachedInContext *bool `json:"cachedInContext"`
}

// TODO: EXAMPLE
//
type EcsTarget struct {
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// Listener and properties for adding target group to the listener.
	Listener ListenerConfig `json:"listener"`
	// ID for a target group to be created.
	NewTargetGroupId *string `json:"newTargetGroupId"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `json:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	Protocol Protocol `json:"protocol"`
}

// The configuration for an Elastic FileSystem volume.
//
// TODO: EXAMPLE
//
type EfsVolumeConfiguration struct {
	// The Amazon EFS file system ID to use.
	FileSystemId *string `json:"fileSystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig *AuthorizationConfig `json:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// Specifying / will have the same effect as omitting this parameter.
	RootDirectory *string `json:"rootDirectory"`
	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be enabled if Amazon EFS IAM authorization is used.
	//
	// Valid values: ENABLED | DISABLED
	TransitEncryption *string `json:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// EFS mount helper uses.
	TransitEncryptionPort *float64 `json:"transitEncryptionPort"`
}

// Constructs for types of environment files.
//
// TODO: EXAMPLE
//
type EnvironmentFile interface {
	Bind(scope constructs.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for EnvironmentFile
type jsiiProxy_EnvironmentFile struct {
	_ byte // padding
}

func NewEnvironmentFile_Override(e EnvironmentFile) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		nil, // no parameters
		e,
	)
}

// Loads the environment file from a local disk path.
func EnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func EnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Called when the container is initialized to allow this object to bind to the stack.
func (e *jsiiProxy_EnvironmentFile) Bind(scope constructs.Construct) *EnvironmentFileConfig {
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Configuration for the environment file.
//
// TODO: EXAMPLE
//
type EnvironmentFileConfig struct {
	// The type of environment file.
	FileType EnvironmentFileType `json:"fileType"`
	// The location of the environment file in S3.
	S3Location *awss3.Location `json:"s3Location"`
}

// Type of environment file to be included in the container definition.
type EnvironmentFileType string

const (
	EnvironmentFileType_S3 EnvironmentFileType = "S3"
)

// The details of the execute command configuration.
//
// For more information, see
// [ExecuteCommandConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html
//
// TODO: EXAMPLE
//
type ExecuteCommandConfiguration struct {
	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKey awskms.IKey `json:"kmsKey"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
	LogConfiguration *ExecuteCommandLogConfiguration `json:"logConfiguration"`
	// The log settings to use for logging the execute command session.
	Logging ExecuteCommandLogging `json:"logging"`
}

// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs and/ or an Amazon S3 bucket.
// For more information, see [ExecuteCommandLogConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandlogconfiguration.html
//
// TODO: EXAMPLE
//
type ExecuteCommandLogConfiguration struct {
	// Whether or not to enable encryption on the CloudWatch logs.
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// The CloudWatch log group must already be created.
	CloudWatchLogGroup awslogs.ILogGroup `json:"cloudWatchLogGroup"`
	// The name of the S3 bucket to send logs to.
	//
	// The S3 bucket must already be created.
	S3Bucket awss3.IBucket `json:"s3Bucket"`
	// Whether or not to enable encryption on the CloudWatch logs.
	S3EncryptionEnabled *bool `json:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
}

// The log settings to use to for logging the execute command session.
//
// For more information, see
// [Logging] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logging
//
// TODO: EXAMPLE
//
type ExecuteCommandLogging string

const (
	ExecuteCommandLogging_NONE ExecuteCommandLogging = "NONE"
	ExecuteCommandLogging_DEFAULT ExecuteCommandLogging = "DEFAULT"
	ExecuteCommandLogging_OVERRIDE ExecuteCommandLogging = "OVERRIDE"
)

// This creates a service using the External launch type on an ECS cluster.
//
// TODO: EXAMPLE
//
type ExternalService interface {
	BaseService
	IExternalService
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	CloudMapService() awsservicediscovery.IService
	Cluster() ICluster
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	Node() constructs.Node
	PhysicalName() *string
	ServiceArn() *string
	ServiceName() *string
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	Stack() awscdk.Stack
	TaskDefinition() TaskDefinition
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateCloudMapService(_options *AssociateCloudMapServiceOptions)
	AttachToApplicationTargetGroup(_targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AutoScaleTaskCount(_props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	ConfigureAwsVpcNetworkingWithSecurityGroups(_vpc awsec2.IVpc, _assignPublicIp *bool, _vpcSubnets *awsec2.SubnetSelection, _securityGroups *[]awsec2.ISecurityGroup)
	EnableCloudMap(_options *CloudMapOptions) awsservicediscovery.Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LoadBalancerTarget(_options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterLoadBalancerTargets(_targets ...*EcsTarget)
	ToString() *string
}

// The jsii proxy struct for ExternalService
type jsiiProxy_ExternalService struct {
	jsiiProxy_BaseService
	jsiiProxy_IExternalService
}

func (j *jsiiProxy_ExternalService) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ExternalService class.
func NewExternalService(scope constructs.Construct, id *string, props *ExternalServiceProps) ExternalService {
	_init_.Initialize()

	j := jsiiProxy_ExternalService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ExternalService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ExternalService class.
func NewExternalService_Override(e ExternalService, scope constructs.Construct, id *string, props *ExternalServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ExternalService",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ExternalService) SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_ExternalService) SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_ExternalService) SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_ExternalService) SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Imports from the specified service ARN.
func ExternalService_FromExternalServiceArn(scope constructs.Construct, id *string, externalServiceArn *string) IExternalService {
	_init_.Initialize()

	var returns IExternalService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalService",
		"fromExternalServiceArn",
		[]interface{}{scope, id, externalServiceArn},
		&returns,
	)

	return returns
}

// Imports from the specified service attrributes.
func ExternalService_FromExternalServiceAttributes(scope constructs.Construct, id *string, attrs *ExternalServiceAttributes) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalService",
		"fromExternalServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ExternalService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ExternalService_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalService",
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
func (e *jsiiProxy_ExternalService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Overriden method to throw error as `associateCloudMapService` is not supported for external service.
func (e *jsiiProxy_ExternalService) AssociateCloudMapService(_options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		e,
		"associateCloudMapService",
		[]interface{}{_options},
	)
}

// Overriden method to throw error as `attachToApplicationTargetGroup` is not supported for external service.
func (e *jsiiProxy_ExternalService) AttachToApplicationTargetGroup(_targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToApplicationTargetGroup",
		[]interface{}{_targetGroup},
		&returns,
	)

	return returns
}

// Registers the service as a target of a Classic Load Balancer (CLB).
//
// Don't call this. Call `loadBalancer.addTarget()` instead.
func (e *jsiiProxy_ExternalService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		e,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// This method is called to attach this service to a Network Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (e *jsiiProxy_ExternalService) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Overriden method to throw error as `autoScaleTaskCount` is not supported for external service.
func (e *jsiiProxy_ExternalService) AutoScaleTaskCount(_props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	var returns ScalableTaskCount

	_jsii_.Invoke(
		e,
		"autoScaleTaskCount",
		[]interface{}{_props},
		&returns,
	)

	return returns
}

// Overriden method to throw error as `configureAwsVpcNetworkingWithSecurityGroups` is not supported for external service.
func (e *jsiiProxy_ExternalService) ConfigureAwsVpcNetworkingWithSecurityGroups(_vpc awsec2.IVpc, _assignPublicIp *bool, _vpcSubnets *awsec2.SubnetSelection, _securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{_vpc, _assignPublicIp, _vpcSubnets, _securityGroups},
	)
}

// Overriden method to throw error as `enableCloudMap` is not supported for external service.
func (e *jsiiProxy_ExternalService) EnableCloudMap(_options *CloudMapOptions) awsservicediscovery.Service {
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		e,
		"enableCloudMap",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_ExternalService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_ExternalService) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Overriden method to throw error as `loadBalancerTarget` is not supported for external service.
func (e *jsiiProxy_ExternalService) LoadBalancerTarget(_options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		e,
		"loadBalancerTarget",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

// This method returns the specified CloudWatch metric name for this service.
func (e *jsiiProxy_ExternalService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's CPU utilization.
func (e *jsiiProxy_ExternalService) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's memory utilization.
func (e *jsiiProxy_ExternalService) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Overriden method to throw error as `registerLoadBalancerTargets` is not supported for external service.
func (e *jsiiProxy_ExternalService) RegisterLoadBalancerTargets(_targets ...*EcsTarget) {
	args := []interface{}{}
	for _, a := range _targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"registerLoadBalancerTargets",
		args,
	)
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ExternalService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties to import from the service using the External launch type.
//
// TODO: EXAMPLE
//
type ExternalServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// The service ARN.
	ServiceArn *string `json:"serviceArn"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
}

// The properties for defining a service using the External launch type.
//
// TODO: EXAMPLE
//
type ExternalServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `json:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface]
	TaskDefinition TaskDefinition `json:"taskDefinition"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
}

// The details of a task definition run on an External cluster.
//
// TODO: EXAMPLE
//
type ExternalTaskDefinition interface {
	TaskDefinition
	IExternalTaskDefinition
	Compatibility() Compatibility
	Containers() *[]ContainerDefinition
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	Env() *awscdk.ResourceEnvironment
	EphemeralStorageGiB() *float64
	ExecutionRole() awsiam.IRole
	Family() *string
	InferenceAccelerators() *[]*InferenceAccelerator
	IsEc2Compatible() *bool
	IsExternalCompatible() *bool
	IsFargateCompatible() *bool
	NetworkMode() NetworkMode
	Node() constructs.Node
	PhysicalName() *string
	ReferencesSecretJsonField() *bool
	Stack() awscdk.Stack
	TaskDefinitionArn() *string
	TaskRole() awsiam.IRole
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	AddExtension(extension ITaskDefinitionExtension)
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	AddInferenceAccelerator(_inferenceAccelerator *InferenceAccelerator)
	AddPlacementConstraint(constraint PlacementConstraint)
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	AddVolume(_volume *Volume)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	FindContainer(containerName *string) ContainerDefinition
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ObtainExecutionRole() awsiam.IRole
	ToString() *string
}

// The jsii proxy struct for ExternalTaskDefinition
type jsiiProxy_ExternalTaskDefinition struct {
	jsiiProxy_TaskDefinition
	jsiiProxy_IExternalTaskDefinition
}

func (j *jsiiProxy_ExternalTaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) Containers() *[]ContainerDefinition {
	var returns *[]ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) DefaultContainer() ContainerDefinition {
	var returns ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) InferenceAccelerators() *[]*InferenceAccelerator {
	var returns *[]*InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ExternalTaskDefinition class.
func NewExternalTaskDefinition(scope constructs.Construct, id *string, props *ExternalTaskDefinitionProps) ExternalTaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_ExternalTaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ExternalTaskDefinition class.
func NewExternalTaskDefinition_Override(e ExternalTaskDefinition, scope constructs.Construct, id *string, props *ExternalTaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ExternalTaskDefinition) SetDefaultContainer(val ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
func ExternalTaskDefinition_FromEc2TaskDefinitionArn(scope constructs.Construct, id *string, externalTaskDefinitionArn *string) IExternalTaskDefinition {
	_init_.Initialize()

	var returns IExternalTaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"fromEc2TaskDefinitionArn",
		[]interface{}{scope, id, externalTaskDefinitionArn},
		&returns,
	)

	return returns
}

// Imports an existing External task definition from its attributes.
func ExternalTaskDefinition_FromExternalTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *ExternalTaskDefinitionAttributes) IExternalTaskDefinition {
	_init_.Initialize()

	var returns IExternalTaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"fromExternalTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func ExternalTaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func ExternalTaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"fromTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ExternalTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ExternalTaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalTaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a new container to the task definition.
func (e *jsiiProxy_ExternalTaskDefinition) AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		e,
		"addContainer",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds the specified extension to the task definition.
//
// Extension can be used to apply a packaged modification to
// a task definition.
func (e *jsiiProxy_ExternalTaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		e,
		"addExtension",
		[]interface{}{extension},
	)
}

// Adds a firelens log router to the task definition.
func (e *jsiiProxy_ExternalTaskDefinition) AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter {
	var returns FirelensLogRouter

	_jsii_.Invoke(
		e,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Overriden method to throw error as interface accelerators are not supported for external tasks.
func (e *jsiiProxy_ExternalTaskDefinition) AddInferenceAccelerator(_inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		e,
		"addInferenceAccelerator",
		[]interface{}{_inferenceAccelerator},
	)
}

// Adds the specified placement constraint to the task definition.
func (e *jsiiProxy_ExternalTaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

// Adds a policy statement to the task execution IAM role.
func (e *jsiiProxy_ExternalTaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a policy statement to the task IAM role.
func (e *jsiiProxy_ExternalTaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

// Overridden method to throw error, as volumes are not supported for external task definitions.
func (e *jsiiProxy_ExternalTaskDefinition) AddVolume(_volume *Volume) {
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{_volume},
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
func (e *jsiiProxy_ExternalTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns the container that match the provided containerName.
func (e *jsiiProxy_ExternalTaskDefinition) FindContainer(containerName *string) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		e,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_ExternalTaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
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
func (e *jsiiProxy_ExternalTaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Creates the task execution IAM role if it doesn't already exist.
func (e *jsiiProxy_ExternalTaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		e,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ExternalTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes used to import an existing External task definition.
//
// TODO: EXAMPLE
//
type ExternalTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `json:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `json:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// The properties for a task definition run on an External cluster.
//
// TODO: EXAMPLE
//
type ExternalTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `json:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `json:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `json:"volumes"`
}

// The platform version on which to run your service.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
//
type FargatePlatformVersion string

const (
	FargatePlatformVersion_LATEST FargatePlatformVersion = "LATEST"
	FargatePlatformVersion_VERSION1_4 FargatePlatformVersion = "VERSION1_4"
	FargatePlatformVersion_VERSION1_3 FargatePlatformVersion = "VERSION1_3"
	FargatePlatformVersion_VERSION1_2 FargatePlatformVersion = "VERSION1_2"
	FargatePlatformVersion_VERSION1_1 FargatePlatformVersion = "VERSION1_1"
	FargatePlatformVersion_VERSION1_0 FargatePlatformVersion = "VERSION1_0"
)

// This creates a service using the Fargate launch type on an ECS cluster.
//
// TODO: EXAMPLE
//
type FargateService interface {
	BaseService
	IFargateService
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	CloudMapService() awsservicediscovery.IService
	Cluster() ICluster
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	Node() constructs.Node
	PhysicalName() *string
	ServiceArn() *string
	ServiceName() *string
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	Stack() awscdk.Stack
	TaskDefinition() TaskDefinition
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	ToString() *string
}

// The jsii proxy struct for FargateService
type jsiiProxy_FargateService struct {
	jsiiProxy_BaseService
	jsiiProxy_IFargateService
}

func (j *jsiiProxy_FargateService) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateService) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the FargateService class.
func NewFargateService(scope constructs.Construct, id *string, props *FargateServiceProps) FargateService {
	_init_.Initialize()

	j := jsiiProxy_FargateService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FargateService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FargateService class.
func NewFargateService_Override(f FargateService, scope constructs.Construct, id *string, props *FargateServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FargateService",
		[]interface{}{scope, id, props},
		f,
	)
}

func (j *jsiiProxy_FargateService) SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_FargateService) SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_FargateService) SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_FargateService) SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Imports from the specified service ARN.
func FargateService_FromFargateServiceArn(scope constructs.Construct, id *string, fargateServiceArn *string) IFargateService {
	_init_.Initialize()

	var returns IFargateService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateService",
		"fromFargateServiceArn",
		[]interface{}{scope, id, fargateServiceArn},
		&returns,
	)

	return returns
}

// Imports from the specified service attrributes.
func FargateService_FromFargateServiceAttributes(scope constructs.Construct, id *string, attrs *FargateServiceAttributes) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateService",
		"fromFargateServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FargateService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FargateService_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateService",
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
func (f *jsiiProxy_FargateService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associates this service with a CloudMap service.
func (f *jsiiProxy_FargateService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		f,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

// This method is called to attach this service to an Application Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (f *jsiiProxy_FargateService) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		f,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// Registers the service as a target of a Classic Load Balancer (CLB).
//
// Don't call this. Call `loadBalancer.addTarget()` instead.
func (f *jsiiProxy_FargateService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		f,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// This method is called to attach this service to a Network Load Balancer.
//
// Don't call this function directly. Instead, call `listener.addTargets()`
// to add this service to a load balancer.
func (f *jsiiProxy_FargateService) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		f,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
func (f *jsiiProxy_FargateService) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	var returns ScalableTaskCount

	_jsii_.Invoke(
		f,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method is called to create a networkConfiguration.
func (f *jsiiProxy_FargateService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		f,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

// Enable CloudMap service discovery for the service.
//
// Returns: The created CloudMap service
func (f *jsiiProxy_FargateService) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		f,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_FargateService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (f *jsiiProxy_FargateService) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Return a load balancing target for a specific container and port.
//
// Use this function to create a load balancer target if you want to load balance to
// another container than the first essential container or the first mapped port on
// the container.
//
// Use the return value of this function where you would normally use a load balancer
// target, instead of the `Service` object itself.
//
// TODO: EXAMPLE
//
func (f *jsiiProxy_FargateService) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		f,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// This method returns the specified CloudWatch metric name for this service.
func (f *jsiiProxy_FargateService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's CPU utilization.
func (f *jsiiProxy_FargateService) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// This method returns the CloudWatch metric for this service's memory utilization.
func (f *jsiiProxy_FargateService) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
//
// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
//
// TODO: EXAMPLE
//
func (f *jsiiProxy_FargateService) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"registerLoadBalancerTargets",
		args,
	)
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FargateService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties to import from the service using the Fargate launch type.
//
// TODO: EXAMPLE
//
type FargateServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// The service ARN.
	ServiceArn *string `json:"serviceArn"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
}

// The properties for defining a service using the Fargate launch type.
//
// TODO: EXAMPLE
//
type FargateServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `json:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `json:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `json:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `json:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `json:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `json:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `json:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `json:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `json:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `json:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `json:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `json:"propagateTags"`
	// The name of the service.
	ServiceName *string `json:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface]
	TaskDefinition TaskDefinition `json:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion FargatePlatformVersion `json:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The subnets to associate with the service.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// The details of a task definition run on a Fargate cluster.
//
// TODO: EXAMPLE
//
type FargateTaskDefinition interface {
	TaskDefinition
	IFargateTaskDefinition
	Compatibility() Compatibility
	Containers() *[]ContainerDefinition
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	Env() *awscdk.ResourceEnvironment
	EphemeralStorageGiB() *float64
	ExecutionRole() awsiam.IRole
	Family() *string
	InferenceAccelerators() *[]*InferenceAccelerator
	IsEc2Compatible() *bool
	IsExternalCompatible() *bool
	IsFargateCompatible() *bool
	NetworkMode() NetworkMode
	Node() constructs.Node
	PhysicalName() *string
	ReferencesSecretJsonField() *bool
	Stack() awscdk.Stack
	TaskDefinitionArn() *string
	TaskRole() awsiam.IRole
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	AddExtension(extension ITaskDefinitionExtension)
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	AddPlacementConstraint(constraint PlacementConstraint)
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	AddVolume(volume *Volume)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	FindContainer(containerName *string) ContainerDefinition
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ObtainExecutionRole() awsiam.IRole
	ToString() *string
}

// The jsii proxy struct for FargateTaskDefinition
type jsiiProxy_FargateTaskDefinition struct {
	jsiiProxy_TaskDefinition
	jsiiProxy_IFargateTaskDefinition
}

func (j *jsiiProxy_FargateTaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) Containers() *[]ContainerDefinition {
	var returns *[]ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) DefaultContainer() ContainerDefinition {
	var returns ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) InferenceAccelerators() *[]*InferenceAccelerator {
	var returns *[]*InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateTaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


// Constructs a new instance of the FargateTaskDefinition class.
func NewFargateTaskDefinition(scope constructs.Construct, id *string, props *FargateTaskDefinitionProps) FargateTaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_FargateTaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FargateTaskDefinition class.
func NewFargateTaskDefinition_Override(f FargateTaskDefinition, scope constructs.Construct, id *string, props *FargateTaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		[]interface{}{scope, id, props},
		f,
	)
}

func (j *jsiiProxy_FargateTaskDefinition) SetDefaultContainer(val ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
func FargateTaskDefinition_FromFargateTaskDefinitionArn(scope constructs.Construct, id *string, fargateTaskDefinitionArn *string) IFargateTaskDefinition {
	_init_.Initialize()

	var returns IFargateTaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"fromFargateTaskDefinitionArn",
		[]interface{}{scope, id, fargateTaskDefinitionArn},
		&returns,
	)

	return returns
}

// Import an existing Fargate task definition from its attributes.
func FargateTaskDefinition_FromFargateTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *FargateTaskDefinitionAttributes) IFargateTaskDefinition {
	_init_.Initialize()

	var returns IFargateTaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"fromFargateTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func FargateTaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func FargateTaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"fromTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FargateTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FargateTaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateTaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a new container to the task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		f,
		"addContainer",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds the specified extension to the task definition.
//
// Extension can be used to apply a packaged modification to
// a task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		f,
		"addExtension",
		[]interface{}{extension},
	)
}

// Adds a firelens log router to the task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter {
	var returns FirelensLogRouter

	_jsii_.Invoke(
		f,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds an inference accelerator to the task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		f,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

// Adds the specified placement constraint to the task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		f,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

// Adds a policy statement to the task execution IAM role.
func (f *jsiiProxy_FargateTaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a policy statement to the task IAM role.
func (f *jsiiProxy_FargateTaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a volume to the task definition.
func (f *jsiiProxy_FargateTaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		f,
		"addVolume",
		[]interface{}{volume},
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
func (f *jsiiProxy_FargateTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns the container that match the provided containerName.
func (f *jsiiProxy_FargateTaskDefinition) FindContainer(containerName *string) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		f,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateTaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_FargateTaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (f *jsiiProxy_FargateTaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Creates the task execution IAM role if it doesn't already exist.
func (f *jsiiProxy_FargateTaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		f,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FargateTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes used to import an existing Fargate task definition.
//
// TODO: EXAMPLE
//
type FargateTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `json:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `json:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
}

// The properties for a task definition.
//
// TODO: EXAMPLE
//
type FargateTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `json:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `json:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `json:"volumes"`
	// The number of cpu units used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values,
	// which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	//
	// 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	//
	// 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	// 2048 (2 vCPU) - Available memory values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	//
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)
	Cpu *float64 `json:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The maximum supported value is 200 GiB.
	//
	// NOTE: This parameter is only supported for tasks hosted on AWS Fargate using platform version 1.4.0 or later.
	EphemeralStorageGiB *float64 `json:"ephemeralStorageGiB"`
	// The amount (in MiB) of memory used by the task.
	//
	// For tasks using the Fargate launch type,
	// this field is required and you must use one of the following values, which determines your range of valid values for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The operating system that your task definitions are running on.
	//
	// A runtimePlatform is supported only for tasks using the Fargate launch type.
	RuntimePlatform *RuntimePlatform `json:"runtimePlatform"`
}

// FireLens enables you to use task definition parameters to route logs to an AWS service   or AWS Partner Network (APN) destination for log storage and analytics.
//
// TODO: EXAMPLE
//
type FireLensLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for FireLensLogDriver
type jsiiProxy_FireLensLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the FireLensLogDriver class.
func NewFireLensLogDriver(props *FireLensLogDriverProps) FireLensLogDriver {
	_init_.Initialize()

	j := jsiiProxy_FireLensLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FireLensLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FireLensLogDriver class.
func NewFireLensLogDriver_Override(f FireLensLogDriver, props *FireLensLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FireLensLogDriver",
		[]interface{}{props},
		f,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func FireLensLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FireLensLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (f *jsiiProxy_FireLensLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the firelens log driver configuration options.
//
// TODO: EXAMPLE
//
type FireLensLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `json:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *map[string]Secret `json:"secretOptions"`
}

// Firelens Configuration https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef.
//
// TODO: EXAMPLE
//
type FirelensConfig struct {
	// The log router to use.
	Type FirelensLogRouterType `json:"type"`
	// Firelens options.
	Options *FirelensOptions `json:"options"`
}

// Firelens configuration file type, s3 or file path.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig
type FirelensConfigFileType string

const (
	FirelensConfigFileType_S3 FirelensConfigFileType = "S3"
	FirelensConfigFileType_FILE FirelensConfigFileType = "FILE"
)

// Firelens log router.
//
// TODO: EXAMPLE
//
type FirelensLogRouter interface {
	ContainerDefinition
	ContainerDependencies() *[]*ContainerDependency
	ContainerName() *string
	ContainerPort() *float64
	EnvironmentFiles() *[]*EnvironmentFileConfig
	Essential() *bool
	FirelensConfig() *FirelensConfig
	IngressPort() *float64
	LinuxParameters() LinuxParameters
	LogDriverConfig() *LogDriverConfig
	MemoryLimitSpecified() *bool
	MountPoints() *[]*MountPoint
	Node() constructs.Node
	PortMappings() *[]*PortMapping
	ReferencesSecretJsonField() *bool
	TaskDefinition() TaskDefinition
	Ulimits() *[]*Ulimit
	VolumesFrom() *[]*VolumeFrom
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	AddEnvironment(name *string, value *string)
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	AddLink(container ContainerDefinition, alias *string)
	AddMountPoints(mountPoints ...*MountPoint)
	AddPortMappings(portMappings ...*PortMapping)
	AddScratch(scratch *ScratchSpace)
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	AddUlimits(ulimits ...*Ulimit)
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	ToString() *string
}

// The jsii proxy struct for FirelensLogRouter
type jsiiProxy_FirelensLogRouter struct {
	jsiiProxy_ContainerDefinition
}

func (j *jsiiProxy_FirelensLogRouter) ContainerDependencies() *[]*ContainerDependency {
	var returns *[]*ContainerDependency
	_jsii_.Get(
		j,
		"containerDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) EnvironmentFiles() *[]*EnvironmentFileConfig {
	var returns *[]*EnvironmentFileConfig
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Essential() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) FirelensConfig() *FirelensConfig {
	var returns *FirelensConfig
	_jsii_.Get(
		j,
		"firelensConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) IngressPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) LogDriverConfig() *LogDriverConfig {
	var returns *LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) MemoryLimitSpecified() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"memoryLimitSpecified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) MountPoints() *[]*MountPoint {
	var returns *[]*MountPoint
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) PortMappings() *[]*PortMapping {
	var returns *[]*PortMapping
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) VolumesFrom() *[]*VolumeFrom {
	var returns *[]*VolumeFrom
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}


// Constructs a new instance of the FirelensLogRouter class.
func NewFirelensLogRouter(scope constructs.Construct, id *string, props *FirelensLogRouterProps) FirelensLogRouter {
	_init_.Initialize()

	j := jsiiProxy_FirelensLogRouter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FirelensLogRouter class.
func NewFirelensLogRouter_Override(f FirelensLogRouter, scope constructs.Construct, id *string, props *FirelensLogRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FirelensLogRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// This method adds one or more container dependencies to the container.
func (f *jsiiProxy_FirelensLogRouter) AddContainerDependencies(containerDependencies ...*ContainerDependency) {
	args := []interface{}{}
	for _, a := range containerDependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addContainerDependencies",
		args,
	)
}

// This method adds an environment variable to the container.
func (f *jsiiProxy_FirelensLogRouter) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		f,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

// This method adds one or more resources to the container.
func (f *jsiiProxy_FirelensLogRouter) AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string) {
	args := []interface{}{}
	for _, a := range inferenceAcceleratorResources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addInferenceAcceleratorResource",
		args,
	)
}

// This method adds a link which allows containers to communicate with each other without the need for port mappings.
//
// This parameter is only supported if the task definition is using the bridge network mode.
// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
func (f *jsiiProxy_FirelensLogRouter) AddLink(container ContainerDefinition, alias *string) {
	_jsii_.InvokeVoid(
		f,
		"addLink",
		[]interface{}{container, alias},
	)
}

// This method adds one or more mount points for data volumes to the container.
func (f *jsiiProxy_FirelensLogRouter) AddMountPoints(mountPoints ...*MountPoint) {
	args := []interface{}{}
	for _, a := range mountPoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addMountPoints",
		args,
	)
}

// This method adds one or more port mappings to the container.
func (f *jsiiProxy_FirelensLogRouter) AddPortMappings(portMappings ...*PortMapping) {
	args := []interface{}{}
	for _, a := range portMappings {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addPortMappings",
		args,
	)
}

// This method mounts temporary disk space to the container.
//
// This adds the correct container mountPoint and task definition volume.
func (f *jsiiProxy_FirelensLogRouter) AddScratch(scratch *ScratchSpace) {
	_jsii_.InvokeVoid(
		f,
		"addScratch",
		[]interface{}{scratch},
	)
}

// This method adds the specified statement to the IAM task execution policy in the task definition.
func (f *jsiiProxy_FirelensLogRouter) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

// This method adds one or more ulimits to the container.
func (f *jsiiProxy_FirelensLogRouter) AddUlimits(ulimits ...*Ulimit) {
	args := []interface{}{}
	for _, a := range ulimits {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addUlimits",
		args,
	)
}

// This method adds one or more volumes to the container.
func (f *jsiiProxy_FirelensLogRouter) AddVolumesFrom(volumesFrom ...*VolumeFrom) {
	args := []interface{}{}
	for _, a := range volumesFrom {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addVolumesFrom",
		args,
	)
}

// Returns the host port for the requested container port if it exists.
func (f *jsiiProxy_FirelensLogRouter) FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping {
	var returns *PortMapping

	_jsii_.Invoke(
		f,
		"findPortMapping",
		[]interface{}{containerPort, protocol},
		&returns,
	)

	return returns
}

// Render this container definition to a CloudFormation object.
func (f *jsiiProxy_FirelensLogRouter) RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty {
	var returns *CfnTaskDefinition_ContainerDefinitionProperty

	_jsii_.Invoke(
		f,
		"renderContainerDefinition",
		[]interface{}{_taskDefinition},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FirelensLogRouter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The options for creating a firelens log router.
//
// TODO: EXAMPLE
//
type FirelensLogRouterDefinitionOptions struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage
	Image ContainerImage `json:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `json:"command"`
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `json:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `json:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `json:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `json:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `json:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `json:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `json:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `json:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `json:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `json:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `json:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `json:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `json:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `json:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `json:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `json:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `json:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `json:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `json:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `json:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `json:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `json:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `json:"systemControls"`
	// The user name to use inside the container.
	User *string `json:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `json:"workingDirectory"`
	// Firelens configuration.
	FirelensConfig *FirelensConfig `json:"firelensConfig"`
}

// The properties in a firelens log router.
//
// TODO: EXAMPLE
//
type FirelensLogRouterProps struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage
	Image ContainerImage `json:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `json:"command"`
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `json:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `json:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `json:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `json:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `json:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `json:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `json:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `json:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `json:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `json:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `json:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `json:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `json:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `json:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `json:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `json:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `json:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `json:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `json:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `json:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `json:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `json:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `json:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `json:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `json:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `json:"systemControls"`
	// The user name to use inside the container.
	User *string `json:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `json:"workingDirectory"`
	// The name of the task definition that includes this container definition.
	//
	// [disable-awslint:ref-via-interface]
	TaskDefinition TaskDefinition `json:"taskDefinition"`
	// Firelens configuration.
	FirelensConfig *FirelensConfig `json:"firelensConfig"`
}

// Firelens log router type, fluentbit or fluentd.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html
type FirelensLogRouterType string

const (
	FirelensLogRouterType_FLUENTBIT FirelensLogRouterType = "FLUENTBIT"
	FirelensLogRouterType_FLUENTD FirelensLogRouterType = "FLUENTD"
)

// The options for firelens log router https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig.
//
// TODO: EXAMPLE
//
type FirelensOptions struct {
	// Custom configuration file, S3 ARN or a file path.
	ConfigFileValue *string `json:"configFileValue"`
	// Custom configuration file, s3 or file.
	ConfigFileType FirelensConfigFileType `json:"configFileType"`
	// By default, Amazon ECS adds additional fields in your log entries that help identify the source of the logs.
	//
	// You can disable this action by setting enable-ecs-log-metadata to false.
	EnableECSLogMetadata *bool `json:"enableECSLogMetadata"`
}

// A log driver that sends log information to journald Logs.
//
// TODO: EXAMPLE
//
type FluentdLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for FluentdLogDriver
type jsiiProxy_FluentdLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the FluentdLogDriver class.
func NewFluentdLogDriver(props *FluentdLogDriverProps) FluentdLogDriver {
	_init_.Initialize()

	j := jsiiProxy_FluentdLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FluentdLogDriver class.
func NewFluentdLogDriver_Override(f FluentdLogDriver, props *FluentdLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		[]interface{}{props},
		f,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func FluentdLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (f *jsiiProxy_FluentdLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the fluentd log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/fluentd/)
//
// TODO: EXAMPLE
//
type FluentdLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// By default, the logging driver connects to localhost:24224.
	//
	// Supply the
	// address option to connect to a different address. tcp(default) and unix
	// sockets are supported.
	Address *string `json:"address"`
	// Docker connects to Fluentd in the background.
	//
	// Messages are buffered until
	// the connection is established.
	AsyncConnect *bool `json:"asyncConnect"`
	// The amount of data to buffer before flushing to disk.
	BufferLimit *float64 `json:"bufferLimit"`
	// The maximum number of retries.
	MaxRetries *float64 `json:"maxRetries"`
	// How long to wait between retries.
	RetryWait awscdk.Duration `json:"retryWait"`
	// Generates event logs in nanosecond resolution.
	SubSecondPrecision *bool `json:"subSecondPrecision"`
}

// The type of compression the GELF driver uses to compress each log message.
type GelfCompressionType string

const (
	GelfCompressionType_GZIP GelfCompressionType = "GZIP"
	GelfCompressionType_ZLIB GelfCompressionType = "ZLIB"
	GelfCompressionType_NONE GelfCompressionType = "NONE"
)

// A log driver that sends log information to journald Logs.
//
// TODO: EXAMPLE
//
type GelfLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for GelfLogDriver
type jsiiProxy_GelfLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the GelfLogDriver class.
func NewGelfLogDriver(props *GelfLogDriverProps) GelfLogDriver {
	_init_.Initialize()

	j := jsiiProxy_GelfLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the GelfLogDriver class.
func NewGelfLogDriver_Override(g GelfLogDriver, props *GelfLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		[]interface{}{props},
		g,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func GelfLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.GelfLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (g *jsiiProxy_GelfLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the journald log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/gelf/)
//
// TODO: EXAMPLE
//
type GelfLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// The address of the GELF server.
	//
	// tcp and udp are the only supported URI
	// specifier and you must specify the port.
	Address *string `json:"address"`
	// UDP Only The level of compression when gzip or zlib is the gelf-compression-type.
	//
	// An integer in the range of -1 to 9 (BestCompression). Higher levels provide more
	// compression at lower speed. Either -1 or 0 disables compression.
	CompressionLevel *float64 `json:"compressionLevel"`
	// UDP Only The type of compression the GELF driver uses to compress each log message.
	//
	// Allowed values are gzip, zlib and none.
	CompressionType GelfCompressionType `json:"compressionType"`
	// TCP Only The maximum number of reconnection attempts when the connection drop.
	//
	// A positive integer.
	TcpMaxReconnect *float64 `json:"tcpMaxReconnect"`
	// TCP Only The number of seconds to wait between reconnection attempts.
	//
	// A positive integer.
	TcpReconnectDelay awscdk.Duration `json:"tcpReconnectDelay"`
}

// A log driver that sends logs to the specified driver.
//
// TODO: EXAMPLE
//
type GenericLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for GenericLogDriver
type jsiiProxy_GenericLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the GenericLogDriver class.
func NewGenericLogDriver(props *GenericLogDriverProps) GenericLogDriver {
	_init_.Initialize()

	j := jsiiProxy_GenericLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the GenericLogDriver class.
func NewGenericLogDriver_Override(g GenericLogDriver, props *GenericLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		[]interface{}{props},
		g,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func GenericLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (g *jsiiProxy_GenericLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// The configuration to use when creating a log driver.
//
// TODO: EXAMPLE
//
type GenericLogDriverProps struct {
	// The log driver to use for the container.
	//
	// The valid values listed for this parameter are log drivers
	// that the Amazon ECS container agent can communicate with by default.
	//
	// For tasks using the Fargate launch type, the supported log drivers are awslogs and splunk.
	// For tasks using the EC2 launch type, the supported log drivers are awslogs, syslog, gelf, fluentd, splunk, journald, and json-file.
	//
	// For more information about using the awslogs log driver, see
	// [Using the awslogs Log Driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LogDriver *string `json:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `json:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *map[string]Secret `json:"secretOptions"`
}

// The health check command and associated configuration parameters for the container.
//
// TODO: EXAMPLE
//
type HealthCheck struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with CMD to execute the command arguments directly, or
	// CMD-SHELL to run the command with the container's default shell.
	//
	// For example: [ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ]
	Command *[]*string `json:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds.
	Interval awscdk.Duration `json:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries.
	Retries *float64 `json:"retries"`
	// The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You may specify between 0 and 300 seconds.
	StartPeriod awscdk.Duration `json:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds.
	Timeout awscdk.Duration `json:"timeout"`
}

// The details on a container instance bind mount host volume.
//
// TODO: EXAMPLE
//
type Host struct {
	// Specifies the path on the host container instance that is presented to the container.
	//
	// If the sourcePath value does not exist on the host container instance, the Docker daemon creates it.
	// If the location does exist, the contents of the source path folder are exported.
	//
	// This property is not supported for tasks that use the Fargate launch type.
	SourcePath *string `json:"sourcePath"`
}

// The interface for BaseService.
type IBaseService interface {
	IService
	// The cluster that hosts the service.
	Cluster() ICluster
}

// The jsii proxy for IBaseService
type jsiiProxy_IBaseService struct {
	jsiiProxy_IService
}

func (j *jsiiProxy_IBaseService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

// A regional grouping of one or more container instances on which you can run tasks and services.
type ICluster interface {
	awscdk.IResource
	// The autoscaling group added to the cluster if capacity is associated to the cluster.
	AutoscalingGroup() awsautoscaling.IAutoScalingGroup
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn() *string
	// The name of the cluster.
	ClusterName() *string
	// Manage the allowed network connections for the cluster with Security Groups.
	Connections() awsec2.Connections
	// The AWS Cloud Map namespace to associate with the cluster.
	DefaultCloudMapNamespace() awsservicediscovery.INamespace
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration() *ExecuteCommandConfiguration
	// Specifies whether the cluster has EC2 instance capacity.
	HasEc2Capacity() *bool
	// The VPC associated with the cluster.
	Vpc() awsec2.IVpc
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICluster) AutoscalingGroup() awsautoscaling.IAutoScalingGroup {
	var returns awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoscalingGroup",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
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

func (j *jsiiProxy_ICluster) DefaultCloudMapNamespace() awsservicediscovery.INamespace {
	var returns awsservicediscovery.INamespace
	_jsii_.Get(
		j,
		"defaultCloudMapNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ExecuteCommandConfiguration() *ExecuteCommandConfiguration {
	var returns *ExecuteCommandConfiguration
	_jsii_.Get(
		j,
		"executeCommandConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) HasEc2Capacity() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasEc2Capacity",
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

// The interface for a service using the EC2 launch type on an ECS cluster.
type IEc2Service interface {
	IService
}

// The jsii proxy for IEc2Service
type jsiiProxy_IEc2Service struct {
	jsiiProxy_IService
}

// The interface of a task definition run on an EC2 cluster.
type IEc2TaskDefinition interface {
	ITaskDefinition
}

// The jsii proxy for IEc2TaskDefinition
type jsiiProxy_IEc2TaskDefinition struct {
	jsiiProxy_ITaskDefinition
}

// Interface for ECS load balancer target.
type IEcsLoadBalancerTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
}

// The jsii proxy for IEcsLoadBalancerTarget
type jsiiProxy_IEcsLoadBalancerTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

// The interface for a service using the External launch type on an ECS cluster.
type IExternalService interface {
	IService
}

// The jsii proxy for IExternalService
type jsiiProxy_IExternalService struct {
	jsiiProxy_IService
}

// The interface of a task definition run on an External cluster.
type IExternalTaskDefinition interface {
	ITaskDefinition
}

// The jsii proxy for IExternalTaskDefinition
type jsiiProxy_IExternalTaskDefinition struct {
	jsiiProxy_ITaskDefinition
}

// The interface for a service using the Fargate launch type on an ECS cluster.
type IFargateService interface {
	IService
}

// The jsii proxy for IFargateService
type jsiiProxy_IFargateService struct {
	jsiiProxy_IService
}

// The interface of a task definition run on a Fargate cluster.
type IFargateTaskDefinition interface {
	ITaskDefinition
}

// The jsii proxy for IFargateTaskDefinition
type jsiiProxy_IFargateTaskDefinition struct {
	jsiiProxy_ITaskDefinition
}

// The interface for a service.
type IService interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
}

// The jsii proxy for IService
type jsiiProxy_IService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

// The interface for all task definitions.
type ITaskDefinition interface {
	awscdk.IResource
	// What launch types this task definition should be compatible with.
	Compatibility() Compatibility
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS Anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	NetworkMode() NetworkMode
	// ARN of this task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
}

// The jsii proxy for ITaskDefinition
type jsiiProxy_ITaskDefinition struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}

// An extension for Task Definitions.
//
// Classes that want to make changes to a TaskDefinition (such as
// adding helper containers) can implement this interface, and can
// then be "added" to a TaskDefinition like so:
//
//     taskDefinition.addExtension(new MyExtension("some_parameter"));
type ITaskDefinitionExtension interface {
	// Apply the extension to the given TaskDefinition.
	Extend(taskDefinition TaskDefinition)
}

// The jsii proxy for ITaskDefinitionExtension
type jsiiProxy_ITaskDefinitionExtension struct {
	_ byte // padding
}

func (i *jsiiProxy_ITaskDefinitionExtension) Extend(taskDefinition TaskDefinition) {
	_jsii_.InvokeVoid(
		i,
		"extend",
		[]interface{}{taskDefinition},
	)
}

// Elastic Inference Accelerator.
//
// For more information, see [Elastic Inference Basics](https://docs.aws.amazon.com/elastic-inference/latest/developerguide/basics.html)
//
// TODO: EXAMPLE
//
type InferenceAccelerator struct {
	// The Elastic Inference accelerator device name.
	DeviceName *string `json:"deviceName"`
	// The Elastic Inference accelerator type to use.
	//
	// The allowed values are: eia2.medium, eia2.large and eia2.xlarge.
	DeviceType *string `json:"deviceType"`
}

// The IPC resource namespace to use for the containers in the task.
type IpcMode string

const (
	IpcMode_NONE IpcMode = "NONE"
	IpcMode_HOST IpcMode = "HOST"
	IpcMode_TASK IpcMode = "TASK"
)

// A log driver that sends log information to journald Logs.
//
// TODO: EXAMPLE
//
type JournaldLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JournaldLogDriver
type jsiiProxy_JournaldLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JournaldLogDriver class.
func NewJournaldLogDriver(props *JournaldLogDriverProps) JournaldLogDriver {
	_init_.Initialize()

	j := jsiiProxy_JournaldLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JournaldLogDriver class.
func NewJournaldLogDriver_Override(j JournaldLogDriver, props *JournaldLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func JournaldLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (j *jsiiProxy_JournaldLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the journald log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/journald/)
//
// TODO: EXAMPLE
//
type JournaldLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
}

// A log driver that sends log information to json-file Logs.
//
// TODO: EXAMPLE
//
type JsonFileLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JsonFileLogDriver
type jsiiProxy_JsonFileLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JsonFileLogDriver class.
func NewJsonFileLogDriver(props *JsonFileLogDriverProps) JsonFileLogDriver {
	_init_.Initialize()

	j := jsiiProxy_JsonFileLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JsonFileLogDriver class.
func NewJsonFileLogDriver_Override(j JsonFileLogDriver, props *JsonFileLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func JsonFileLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (j *jsiiProxy_JsonFileLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the json-file log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/json-file/)
//
// TODO: EXAMPLE
//
type JsonFileLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// Toggles compression for rotated logs.
	Compress *bool `json:"compress"`
	// The maximum number of log files that can be present.
	//
	// If rolling the logs creates
	// excess files, the oldest file is removed. Only effective when max-size is also set.
	// A positive integer.
	MaxFile *float64 `json:"maxFile"`
	// The maximum size of the log before it is rolled.
	//
	// A positive integer plus a modifier
	// representing the unit of measure (k, m, or g).
	MaxSize *string `json:"maxSize"`
}

// The launch type of an ECS service.
type LaunchType string

const (
	LaunchType_EC2 LaunchType = "EC2"
	LaunchType_FARGATE LaunchType = "FARGATE"
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

// Linux-specific options that are applied to the container.
//
// TODO: EXAMPLE
//
type LinuxParameters interface {
	constructs.Construct
	Node() constructs.Node
	AddCapabilities(cap ...Capability)
	AddDevices(device ...*Device)
	AddTmpfs(tmpfs ...*Tmpfs)
	DropCapabilities(cap ...Capability)
	RenderLinuxParameters() *CfnTaskDefinition_LinuxParametersProperty
	ToString() *string
}

// The jsii proxy struct for LinuxParameters
type jsiiProxy_LinuxParameters struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LinuxParameters) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters(scope constructs.Construct, id *string, props *LinuxParametersProps) LinuxParameters {
	_init_.Initialize()

	j := jsiiProxy_LinuxParameters{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LinuxParameters",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters_Override(l LinuxParameters, scope constructs.Construct, id *string, props *LinuxParametersProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LinuxParameters",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LinuxParameters_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LinuxParameters",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Adds one or more Linux capabilities to the Docker configuration of a container.
//
// Only works with EC2 launch type.
func (l *jsiiProxy_LinuxParameters) AddCapabilities(cap ...Capability) {
	args := []interface{}{}
	for _, a := range cap {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addCapabilities",
		args,
	)
}

// Adds one or more host devices to a container.
func (l *jsiiProxy_LinuxParameters) AddDevices(device ...*Device) {
	args := []interface{}{}
	for _, a := range device {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addDevices",
		args,
	)
}

// Specifies the container path, mount options, and size (in MiB) of the tmpfs mount for a container.
//
// Only works with EC2 launch type.
func (l *jsiiProxy_LinuxParameters) AddTmpfs(tmpfs ...*Tmpfs) {
	args := []interface{}{}
	for _, a := range tmpfs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addTmpfs",
		args,
	)
}

// Removes one or more Linux capabilities to the Docker configuration of a container.
//
// Only works with EC2 launch type.
func (l *jsiiProxy_LinuxParameters) DropCapabilities(cap ...Capability) {
	args := []interface{}{}
	for _, a := range cap {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"dropCapabilities",
		args,
	)
}

// Renders the Linux parameters to a CloudFormation object.
func (l *jsiiProxy_LinuxParameters) RenderLinuxParameters() *CfnTaskDefinition_LinuxParametersProperty {
	var returns *CfnTaskDefinition_LinuxParametersProperty

	_jsii_.Invoke(
		l,
		"renderLinuxParameters",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LinuxParameters) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for defining Linux-specific options that are applied to the container.
//
// TODO: EXAMPLE
//
type LinuxParametersProps struct {
	// Specifies whether to run an init process inside the container that forwards signals and reaps processes.
	InitProcessEnabled *bool `json:"initProcessEnabled"`
	// The value for the size (in MiB) of the /dev/shm volume.
	SharedMemorySize *float64 `json:"sharedMemorySize"`
}

// Base class for configuring listener when registering targets.
//
// TODO: EXAMPLE
//
type ListenerConfig interface {
	AddTargets(id *string, target *LoadBalancerTargetOptions, service BaseService)
}

// The jsii proxy struct for ListenerConfig
type jsiiProxy_ListenerConfig struct {
	_ byte // padding
}

func NewListenerConfig_Override(l ListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ListenerConfig",
		nil, // no parameters
		l,
	)
}

// Create a config for adding target group to ALB listener.
func ListenerConfig_ApplicationListener(listener awselasticloadbalancingv2.ApplicationListener, props *awselasticloadbalancingv2.AddApplicationTargetsProps) ListenerConfig {
	_init_.Initialize()

	var returns ListenerConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ListenerConfig",
		"applicationListener",
		[]interface{}{listener, props},
		&returns,
	)

	return returns
}

// Create a config for adding target group to NLB listener.
func ListenerConfig_NetworkListener(listener awselasticloadbalancingv2.NetworkListener, props *awselasticloadbalancingv2.AddNetworkTargetsProps) ListenerConfig {
	_init_.Initialize()

	var returns ListenerConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ListenerConfig",
		"networkListener",
		[]interface{}{listener, props},
		&returns,
	)

	return returns
}

// Create and attach a target group to listener.
func (l *jsiiProxy_ListenerConfig) AddTargets(id *string, target *LoadBalancerTargetOptions, service BaseService) {
	_jsii_.InvokeVoid(
		l,
		"addTargets",
		[]interface{}{id, target, service},
	)
}

// Properties for defining an ECS target.
//
// The port mapping for it must already have been created through addPortMapping().
//
// TODO: EXAMPLE
//
type LoadBalancerTargetOptions struct {
	// The name of the container.
	ContainerName *string `json:"containerName"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `json:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	Protocol Protocol `json:"protocol"`
}

// The base class for log drivers.
//
// TODO: EXAMPLE
//
type LogDriver interface {
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for LogDriver
type jsiiProxy_LogDriver struct {
	_ byte // padding
}

func NewLogDriver_Override(l LogDriver) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDriver",
		nil, // no parameters
		l,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func LogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (l *jsiiProxy_LogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// The configuration to use when creating a log driver.
//
// TODO: EXAMPLE
//
type LogDriverConfig struct {
	// The log driver to use for the container.
	//
	// The valid values listed for this parameter are log drivers
	// that the Amazon ECS container agent can communicate with by default.
	//
	// For tasks using the Fargate launch type, the supported log drivers are awslogs, splunk, and awsfirelens.
	// For tasks using the EC2 launch type, the supported log drivers are awslogs, fluentd, gelf, json-file, journald,
	// logentries,syslog, splunk, and awsfirelens.
	//
	// For more information about using the awslogs log driver, see
	// [Using the awslogs Log Driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html)
	// in the Amazon Elastic Container Service Developer Guide.
	//
	// For more information about using the awsfirelens log driver, see
	// [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LogDriver *string `json:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `json:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *[]*CfnTaskDefinition_SecretProperty `json:"secretOptions"`
}

// The base class for log drivers.
//
// TODO: EXAMPLE
//
type LogDrivers interface {
}

// The jsii proxy struct for LogDrivers
type jsiiProxy_LogDrivers struct {
	_ byte // padding
}

func NewLogDrivers() LogDrivers {
	_init_.Initialize()

	j := jsiiProxy_LogDrivers{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLogDrivers_Override(l LogDrivers) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		nil, // no parameters
		l,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func LogDrivers_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to firelens log router.
//
// For detail configurations, please refer to Amazon ECS FireLens Examples:
// https://github.com/aws-samples/amazon-ecs-firelens-examples
func LogDrivers_Firelens(props *FireLensLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"firelens",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to fluentd Logs.
func LogDrivers_Fluentd(props *FluentdLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"fluentd",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to gelf Logs.
func LogDrivers_Gelf(props *GelfLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"gelf",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to journald Logs.
func LogDrivers_Journald(props *JournaldLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"journald",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to json-file Logs.
func LogDrivers_JsonFile(props *JsonFileLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"jsonFile",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to splunk Logs.
func LogDrivers_Splunk(props *SplunkLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"splunk",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to syslog Logs.
func LogDrivers_Syslog(props *SyslogLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"syslog",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The machine image type.
//
// TODO: EXAMPLE
//
type MachineImageType string

const (
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

// The properties for enabling scaling based on memory utilization.
//
// TODO: EXAMPLE
//
type MemoryUtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// The target value for memory utilization across all tasks in the service.
	TargetUtilizationPercent *float64 `json:"targetUtilizationPercent"`
}

// The details of data volume mount points for a container.
//
// TODO: EXAMPLE
//
type MountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `json:"containerPath"`
	// Specifies whether to give the container read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `json:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	SourceVolume *string `json:"sourceVolume"`
}

// The networking mode to use for the containers in the task.
//
// TODO: EXAMPLE
//
type NetworkMode string

const (
	NetworkMode_NONE NetworkMode = "NONE"
	NetworkMode_BRIDGE NetworkMode = "BRIDGE"
	NetworkMode_AWS_VPC NetworkMode = "AWS_VPC"
	NetworkMode_HOST NetworkMode = "HOST"
	NetworkMode_NAT NetworkMode = "NAT"
)

// The operating system for Fargate Runtime Platform.
//
// TODO: EXAMPLE
//
type OperatingSystemFamily interface {
}

// The jsii proxy struct for OperatingSystemFamily
type jsiiProxy_OperatingSystemFamily struct {
	_ byte // padding
}

// Other operating system family.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-operatingsystemfamily for all available operating system family.
//
func OperatingSystemFamily_Of(family *string) OperatingSystemFamily {
	_init_.Initialize()

	var returns OperatingSystemFamily

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"of",
		[]interface{}{family},
		&returns,
	)

	return returns
}

func OperatingSystemFamily_LINUX() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"LINUX",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2004_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2004_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2016_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2016_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_20H2_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_20H2_CORE",
		&returns,
	)
	return returns
}

// The process namespace to use for the containers in the task.
type PidMode string

const (
	PidMode_HOST PidMode = "HOST"
	PidMode_TASK PidMode = "TASK"
)

// The placement constraints to use for tasks in the service. For more information, see [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
//
// Tasks will only be placed on instances that match these rules.
//
// TODO: EXAMPLE
//
type PlacementConstraint interface {
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

// Return the placement JSON.
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

// The placement strategies to use for tasks in the service. For more information, see [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
//
// Tasks will preferentially be placed on instances that match these rules.
//
// TODO: EXAMPLE
//
type PlacementStrategy interface {
	ToJson() *[]*CfnService_PlacementStrategyProperty
}

// The jsii proxy struct for PlacementStrategy
type jsiiProxy_PlacementStrategy struct {
	_ byte // padding
}

// Places tasks on the container instances with the least available capacity of the specified resource.
func PlacementStrategy_PackedBy(resource BinPackResource) PlacementStrategy {
	_init_.Initialize()

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

// Return the placement JSON.
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

// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
//
// TODO: EXAMPLE
//
type PortMapping struct {
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// For more information, see hostPort.
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	ContainerPort *float64 `json:"containerPort"`
	// The port number on the container instance to reserve for your container.
	//
	// If you are using containers in a task with the awsvpc or host network mode,
	// the hostPort can either be left blank or set to the same value as the containerPort.
	//
	// If you are using containers in a task with the bridge network mode,
	// you can specify a non-reserved host port for your container port mapping, or
	// you can omit the hostPort (or set it to 0) while specifying a containerPort and
	// your container automatically receives a port in the ephemeral port range for
	// your container instance operating system and Docker version.
	HostPort *float64 `json:"hostPort"`
	// The protocol used for the port mapping.
	//
	// Valid values are Protocol.TCP and Protocol.UDP.
	Protocol Protocol `json:"protocol"`
}

// Propagate tags from either service or task definition.
type PropagatedTagSource string

const (
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

// Network protocol.
//
// TODO: EXAMPLE
//
type Protocol string

const (
	Protocol_TCP Protocol = "TCP"
	Protocol_UDP Protocol = "UDP"
)

// The base class for proxy configurations.
type ProxyConfiguration interface {
	Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for ProxyConfiguration
type jsiiProxy_ProxyConfiguration struct {
	_ byte // padding
}

func NewProxyConfiguration_Override(p ProxyConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfiguration",
		nil, // no parameters
		p,
	)
}

// Called when the proxy configuration is configured on a task definition.
func (p *jsiiProxy_ProxyConfiguration) Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

// The base class for proxy configurations.
//
// TODO: EXAMPLE
//
type ProxyConfigurations interface {
}

// The jsii proxy struct for ProxyConfigurations
type jsiiProxy_ProxyConfigurations struct {
	_ byte // padding
}

func NewProxyConfigurations() ProxyConfigurations {
	_init_.Initialize()

	j := jsiiProxy_ProxyConfigurations{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewProxyConfigurations_Override(p ProxyConfigurations) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		nil, // no parameters
		p,
	)
}

// Constructs a new instance of the ProxyConfiguration class.
func ProxyConfigurations_AppMeshProxyConfiguration(props *AppMeshProxyConfigurationConfigProps) ProxyConfiguration {
	_init_.Initialize()

	var returns ProxyConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		"appMeshProxyConfiguration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// An image hosted in a public or private repository.
//
// For images hosted in Amazon ECR, see
// [EcrImage](https://docs.aws.amazon.com/AmazonECR/latest/userguide/images.html).
//
// TODO: EXAMPLE
//
type RepositoryImage interface {
	ContainerImage
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for RepositoryImage
type jsiiProxy_RepositoryImage struct {
	jsiiProxy_ContainerImage
}

// Constructs a new instance of the RepositoryImage class.
func NewRepositoryImage(imageName *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	j := jsiiProxy_RepositoryImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		[]interface{}{imageName, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the RepositoryImage class.
func NewRepositoryImage_Override(r RepositoryImage, imageName *string, props *RepositoryImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		[]interface{}{imageName, props},
		r,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func RepositoryImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func RepositoryImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func RepositoryImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func RepositoryImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func RepositoryImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

// Called when the image is used by a ContainerDefinition.
func (r *jsiiProxy_RepositoryImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// The properties for an image hosted in a public or private repository.
//
// TODO: EXAMPLE
//
type RepositoryImageProps struct {
	// The secret to expose to the container that contains the credentials for the image repository.
	//
	// The supported value is the full ARN of an AWS Secrets Manager secret.
	Credentials awssecretsmanager.ISecret `json:"credentials"`
}

// The properties for enabling scaling based on Application Load Balancer (ALB) request counts.
//
// TODO: EXAMPLE
//
type RequestCountScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// The number of ALB requests per target.
	RequestsPerTarget *float64 `json:"requestsPerTarget"`
	// The ALB target group name.
	TargetGroup awselasticloadbalancingv2.ApplicationTargetGroup `json:"targetGroup"`
}

// The interface for Runtime Platform.
//
// TODO: EXAMPLE
//
type RuntimePlatform struct {
	// The CpuArchitecture for Fargate Runtime Platform.
	CpuArchitecture CpuArchitecture `json:"cpuArchitecture"`
	// The operating system for Fargate Runtime Platform.
	OperatingSystemFamily OperatingSystemFamily `json:"operatingSystemFamily"`
}

// Environment file from S3.
//
// TODO: EXAMPLE
//
type S3EnvironmentFile interface {
	EnvironmentFile
	Bind(_scope constructs.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for S3EnvironmentFile
type jsiiProxy_S3EnvironmentFile struct {
	jsiiProxy_EnvironmentFile
}

func NewS3EnvironmentFile(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	j := jsiiProxy_S3EnvironmentFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3EnvironmentFile_Override(s S3EnvironmentFile, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the environment file from a local disk path.
func S3EnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func S3EnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.S3EnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Called when the container is initialized to allow this object to bind to the stack.
func (s *jsiiProxy_S3EnvironmentFile) Bind(_scope constructs.Construct) *EnvironmentFileConfig {
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// The scalable attribute representing task count.
//
// TODO: EXAMPLE
//
type ScalableTaskCount interface {
	awsapplicationautoscaling.BaseScalableAttribute
	Node() constructs.Node
	Props() *awsapplicationautoscaling.BaseScalableAttributeProps
	DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps)
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps)
	ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps)
	ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps)
	ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps)
	ToString() *string
}

// The jsii proxy struct for ScalableTaskCount
type jsiiProxy_ScalableTaskCount struct {
	internal.Type__awsapplicationautoscalingBaseScalableAttribute
}

func (j *jsiiProxy_ScalableTaskCount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTaskCount) Props() *awsapplicationautoscaling.BaseScalableAttributeProps {
	var returns *awsapplicationautoscaling.BaseScalableAttributeProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScalableTaskCount class.
func NewScalableTaskCount(scope constructs.Construct, id *string, props *ScalableTaskCountProps) ScalableTaskCount {
	_init_.Initialize()

	j := jsiiProxy_ScalableTaskCount{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScalableTaskCount class.
func NewScalableTaskCount_Override(s ScalableTaskCount, scope constructs.Construct, id *string, props *ScalableTaskCountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ScalableTaskCount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Scale out or in based on a metric value.
func (s *jsiiProxy_ScalableTaskCount) DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

// Scale out or in based on time.
func (s *jsiiProxy_ScalableTaskCount) DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

// Scale out or in in order to keep a metric around a target value.
func (s *jsiiProxy_ScalableTaskCount) DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

// Scales in or out to achieve a target CPU utilization.
func (s *jsiiProxy_ScalableTaskCount) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
	)
}

// Scales in or out to achieve a target memory utilization.
func (s *jsiiProxy_ScalableTaskCount) ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMemoryUtilization",
		[]interface{}{id, props},
	)
}

// Scales in or out based on a specified metric value.
func (s *jsiiProxy_ScalableTaskCount) ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
	)
}

// Scales in or out to achieve a target Application Load Balancer request count per target.
func (s *jsiiProxy_ScalableTaskCount) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnRequestCount",
		[]interface{}{id, props},
	)
}

// Scales in or out based on a specified scheduled time.
func (s *jsiiProxy_ScalableTaskCount) ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, props},
	)
}

// Scales in or out to achieve a target on a custom metric.
func (s *jsiiProxy_ScalableTaskCount) ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleToTrackCustomMetric",
		[]interface{}{id, props},
	)
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ScalableTaskCount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties of a scalable attribute representing task count.
//
// TODO: EXAMPLE
//
type ScalableTaskCountProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `json:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `json:"minCapacity"`
	// Scalable dimension of the attribute.
	Dimension *string `json:"dimension"`
	// Resource ID of the attribute.
	ResourceId *string `json:"resourceId"`
	// Role to use for scaling.
	Role awsiam.IRole `json:"role"`
	// Service namespace of the scalable attribute.
	ServiceNamespace awsapplicationautoscaling.ServiceNamespace `json:"serviceNamespace"`
}

// The scope for the Docker volume that determines its lifecycle.
//
// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
// Docker volumes that are scoped as shared persist after the task stops.
type Scope string

const (
	Scope_TASK Scope = "TASK"
	Scope_SHARED Scope = "SHARED"
)

// The temporary disk space mounted to the container.
//
// TODO: EXAMPLE
//
type ScratchSpace struct {
	// The path on the container to mount the scratch volume at.
	ContainerPath *string `json:"containerPath"`
	// The name of the scratch volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	Name *string `json:"name"`
	// Specifies whether to give the container read-only access to the scratch volume.
	//
	// If this value is true, the container has read-only access to the scratch volume.
	// If this value is false, then the container can write to the scratch volume.
	ReadOnly *bool `json:"readOnly"`
	SourcePath *string `json:"sourcePath"`
}

// A secret environment variable.
//
// TODO: EXAMPLE
//
type Secret interface {
	Arn() *string
	HasField() *bool
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	_ byte // padding
}

func (j *jsiiProxy_Secret) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) HasField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasField",
		&returns,
	)
	return returns
}


func NewSecret_Override(s Secret) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Secret",
		nil, // no parameters
		s,
	)
}

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
func Secret_FromSecretsManager(secret awssecretsmanager.ISecret, field *string) Secret {
	_init_.Initialize()

	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
		"fromSecretsManager",
		[]interface{}{secret, field},
		&returns,
	)

	return returns
}

// Creates an environment variable value from a parameter stored in AWS Systems Manager Parameter Store.
func Secret_FromSsmParameter(parameter awsssm.IParameter) Secret {
	_init_.Initialize()

	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

// Grants reading the secret to a principal.
func (s *jsiiProxy_Secret) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// A log driver that sends log information to splunk Logs.
//
// TODO: EXAMPLE
//
type SplunkLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for SplunkLogDriver
type jsiiProxy_SplunkLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the SplunkLogDriver class.
func NewSplunkLogDriver(props *SplunkLogDriverProps) SplunkLogDriver {
	_init_.Initialize()

	j := jsiiProxy_SplunkLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SplunkLogDriver class.
func NewSplunkLogDriver_Override(s SplunkLogDriver, props *SplunkLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		[]interface{}{props},
		s,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func SplunkLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.SplunkLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (s *jsiiProxy_SplunkLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the splunk log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/splunk/)
//
// TODO: EXAMPLE
//
type SplunkLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// Path to your Splunk Enterprise, self-service Splunk Cloud instance, or Splunk Cloud managed cluster (including port and scheme used by HTTP Event Collector) in one of the following formats: https://your_splunk_instance:8088 or https://input-prd-p-XXXXXXX.cloud.splunk.com:8088 or https://http-inputs-XXXXXXXX.splunkcloud.com.
	Url *string `json:"url"`
	// Name to use for validating server certificate.
	CaName *string `json:"caName"`
	// Path to root certificate.
	CaPath *string `json:"caPath"`
	// Message format.
	//
	// Can be inline, json or raw.
	Format SplunkLogFormat `json:"format"`
	// Enable/disable gzip compression to send events to Splunk Enterprise or Splunk Cloud instance.
	Gzip *bool `json:"gzip"`
	// Set compression level for gzip.
	//
	// Valid values are -1 (default), 0 (no compression),
	// 1 (best speed) ... 9 (best compression).
	GzipLevel *float64 `json:"gzipLevel"`
	// Event index.
	Index *string `json:"index"`
	// Ignore server certificate validation.
	InsecureSkipVerify *string `json:"insecureSkipVerify"`
	// Splunk HTTP Event Collector token (Secret).
	//
	// The splunk-token is added to the SecretOptions property of the Log Driver Configuration. So the secret value will not be
	// resolved or viewable as plain text.
	//
	// Please provide at least one of `token` or `secretToken`.
	SecretToken Secret `json:"secretToken"`
	// Event source.
	Source *string `json:"source"`
	// Event source type.
	SourceType *string `json:"sourceType"`
	// Verify on start, that docker can connect to Splunk server.
	VerifyConnection *bool `json:"verifyConnection"`
}

// Log Message Format.
type SplunkLogFormat string

const (
	SplunkLogFormat_INLINE SplunkLogFormat = "INLINE"
	SplunkLogFormat_JSON SplunkLogFormat = "JSON"
	SplunkLogFormat_RAW SplunkLogFormat = "RAW"
)

// A log driver that sends log information to syslog Logs.
//
// TODO: EXAMPLE
//
type SyslogLogDriver interface {
	LogDriver
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for SyslogLogDriver
type jsiiProxy_SyslogLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the SyslogLogDriver class.
func NewSyslogLogDriver(props *SyslogLogDriverProps) SyslogLogDriver {
	_init_.Initialize()

	j := jsiiProxy_SyslogLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SyslogLogDriver class.
func NewSyslogLogDriver_Override(s SyslogLogDriver, props *SyslogLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		[]interface{}{props},
		s,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func SyslogLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called when the log driver is configured on a container.
func (s *jsiiProxy_SyslogLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

// Specifies the syslog log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/syslog/)
//
// TODO: EXAMPLE
//
type SyslogLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `json:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `json:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `json:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `json:"tag"`
	// The address of an external syslog server.
	//
	// The URI specifier may be
	// [tcp|udp|tcp+tls]://host:port, unix://path, or unixgram://path.
	Address *string `json:"address"`
	// The syslog facility to use.
	//
	// Can be the number or name for any valid
	// syslog facility. See the syslog documentation:
	// https://tools.ietf.org/html/rfc5424#section-6.2.1.
	Facility *string `json:"facility"`
	// The syslog message format to use.
	//
	// If not specified the local UNIX syslog
	// format is used, without a specified hostname. Specify rfc3164 for the RFC-3164
	// compatible format, rfc5424 for RFC-5424 compatible format, or rfc5424micro
	// for RFC-5424 compatible format with microsecond timestamp resolution.
	Format *string `json:"format"`
	// The absolute path to the trust certificates signed by the CA.
	//
	// Ignored
	// if the address protocol is not tcp+tls.
	TlsCaCert *string `json:"tlsCaCert"`
	// The absolute path to the TLS certificate file.
	//
	// Ignored if the address
	// protocol is not tcp+tls.
	TlsCert *string `json:"tlsCert"`
	// The absolute path to the TLS key file.
	//
	// Ignored if the address protocol
	// is not tcp+tls.
	TlsKey *string `json:"tlsKey"`
	// If set to true, TLS verification is skipped when connecting to the syslog daemon.
	//
	// Ignored if the address protocol is not tcp+tls.
	TlsSkipVerify *bool `json:"tlsSkipVerify"`
}

// Kernel parameters to set in the container.
//
// TODO: EXAMPLE
//
type SystemControl struct {
	// The namespaced kernel parameter for which to set a value.
	Namespace *string `json:"namespace"`
	// The value for the namespaced kernel parameter specified in namespace.
	Value *string `json:"value"`
}

// A special type of {@link ContainerImage} that uses an ECR repository for the image, but a CloudFormation Parameter for the tag of the image in that repository.
//
// This allows providing this tag through the Parameter at deploy time,
// for example in a CodePipeline that pushes a new tag of the image to the repository during a build step,
// and then provides that new tag through the CloudFormation Parameter in the deploy step.
//
// TODO: EXAMPLE
//
// See: #tagParameterName
//
type TagParameterContainerImage interface {
	ContainerImage
	TagParameterName() *string
	TagParameterValue() *string
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for TagParameterContainerImage
type jsiiProxy_TagParameterContainerImage struct {
	jsiiProxy_ContainerImage
}

func (j *jsiiProxy_TagParameterContainerImage) TagParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagParameterContainerImage) TagParameterValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagParameterValue",
		&returns,
	)
	return returns
}


func NewTagParameterContainerImage(repository awsecr.IRepository) TagParameterContainerImage {
	_init_.Initialize()

	j := jsiiProxy_TagParameterContainerImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		[]interface{}{repository},
		&j,
	)

	return &j
}

func NewTagParameterContainerImage_Override(t TagParameterContainerImage, repository awsecr.IRepository) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		[]interface{}{repository},
		t,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func TagParameterContainerImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func TagParameterContainerImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func TagParameterContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func TagParameterContainerImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func TagParameterContainerImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

// Called when the image is used by a ContainerDefinition.
func (t *jsiiProxy_TagParameterContainerImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

// The base class for all task definitions.
//
// TODO: EXAMPLE
//
type TaskDefinition interface {
	awscdk.Resource
	ITaskDefinition
	Compatibility() Compatibility
	Containers() *[]ContainerDefinition
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	Env() *awscdk.ResourceEnvironment
	EphemeralStorageGiB() *float64
	ExecutionRole() awsiam.IRole
	Family() *string
	InferenceAccelerators() *[]*InferenceAccelerator
	IsEc2Compatible() *bool
	IsExternalCompatible() *bool
	IsFargateCompatible() *bool
	NetworkMode() NetworkMode
	Node() constructs.Node
	PhysicalName() *string
	ReferencesSecretJsonField() *bool
	Stack() awscdk.Stack
	TaskDefinitionArn() *string
	TaskRole() awsiam.IRole
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	AddExtension(extension ITaskDefinitionExtension)
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	AddPlacementConstraint(constraint PlacementConstraint)
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	AddVolume(volume *Volume)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	FindContainer(containerName *string) ContainerDefinition
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ObtainExecutionRole() awsiam.IRole
	ToString() *string
}

// The jsii proxy struct for TaskDefinition
type jsiiProxy_TaskDefinition struct {
	internal.Type__awscdkResource
	jsiiProxy_ITaskDefinition
}

func (j *jsiiProxy_TaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Containers() *[]ContainerDefinition {
	var returns *[]ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) DefaultContainer() ContainerDefinition {
	var returns ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) InferenceAccelerators() *[]*InferenceAccelerator {
	var returns *[]*InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


// Constructs a new instance of the TaskDefinition class.
func NewTaskDefinition(scope constructs.Construct, id *string, props *TaskDefinitionProps) TaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_TaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the TaskDefinition class.
func NewTaskDefinition_Override(t TaskDefinition, scope constructs.Construct, id *string, props *TaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TaskDefinition) SetDefaultContainer(val ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func TaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func TaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"fromTaskDefinitionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a new container to the task definition.
func (t *jsiiProxy_TaskDefinition) AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		t,
		"addContainer",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds the specified extension to the task definition.
//
// Extension can be used to apply a packaged modification to
// a task definition.
func (t *jsiiProxy_TaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		t,
		"addExtension",
		[]interface{}{extension},
	)
}

// Adds a firelens log router to the task definition.
func (t *jsiiProxy_TaskDefinition) AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter {
	var returns FirelensLogRouter

	_jsii_.Invoke(
		t,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Adds an inference accelerator to the task definition.
func (t *jsiiProxy_TaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		t,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

// Adds the specified placement constraint to the task definition.
func (t *jsiiProxy_TaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		t,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

// Adds a policy statement to the task execution IAM role.
func (t *jsiiProxy_TaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a policy statement to the task IAM role.
func (t *jsiiProxy_TaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

// Adds a volume to the task definition.
func (t *jsiiProxy_TaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		t,
		"addVolume",
		[]interface{}{volume},
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
func (t *jsiiProxy_TaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns the container that match the provided containerName.
func (t *jsiiProxy_TaskDefinition) FindContainer(containerName *string) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		t,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
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
func (t *jsiiProxy_TaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
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
func (t *jsiiProxy_TaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Creates the task execution IAM role if it doesn't already exist.
func (t *jsiiProxy_TaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		t,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A reference to an existing task definition.
//
// TODO: EXAMPLE
//
type TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `json:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `json:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// What launch types this task definition should be compatible with.
	Compatibility Compatibility `json:"compatibility"`
}

// The properties for task definitions.
//
// TODO: EXAMPLE
//
type TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `json:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `json:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `json:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `json:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `json:"volumes"`
	// The task launch type compatiblity requirement.
	Compatibility Compatibility `json:"compatibility"`
	// The number of cpu units used by the task.
	//
	// If you are using the EC2 launch type, this field is optional and any value can be used.
	// If you are using the Fargate launch type, this field is required and you must use one of the following values,
	// which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	//
	// 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	//
	// 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	// 2048 (2 vCPU) - Available memory values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	//
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)
	Cpu *string `json:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB *float64 `json:"ephemeralStorageGiB"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	InferenceAccelerators *[]*InferenceAccelerator `json:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	IpcMode IpcMode `json:"ipcMode"`
	// The amount (in MiB) of memory used by the task.
	//
	// If using the EC2 launch type, this field is optional and any value can be used.
	// If using the Fargate launch type, this field is required and you must use one of the following values,
	// which determines your range of valid values for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	MemoryMiB *string `json:"memoryMiB"`
	// The networking mode to use for the containers in the task.
	//
	// On Fargate, the only supported networking mode is AwsVpc.
	NetworkMode NetworkMode `json:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	PidMode PidMode `json:"pidMode"`
	// The placement constraints to use for tasks in the service.
	//
	// You can specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	//
	// Not supported in Fargate.
	PlacementConstraints *[]PlacementConstraint `json:"placementConstraints"`
	// The operating system that your task definitions are running on.
	//
	// A runtimePlatform is supported only for tasks using the Fargate launch type.
	RuntimePlatform *RuntimePlatform `json:"runtimePlatform"`
}

// The details of a tmpfs mount for a container.
//
// TODO: EXAMPLE
//
type Tmpfs struct {
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `json:"containerPath"`
	// The size (in MiB) of the tmpfs volume.
	Size *float64 `json:"size"`
	// The list of tmpfs volume mount options.
	//
	// For more information, see
	// [TmpfsMountOptions](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tmpfs.html).
	MountOptions *[]TmpfsMountOption `json:"mountOptions"`
}

// The supported options for a tmpfs mount for a container.
type TmpfsMountOption string

const (
	TmpfsMountOption_DEFAULTS TmpfsMountOption = "DEFAULTS"
	TmpfsMountOption_RO TmpfsMountOption = "RO"
	TmpfsMountOption_RW TmpfsMountOption = "RW"
	TmpfsMountOption_SUID TmpfsMountOption = "SUID"
	TmpfsMountOption_NOSUID TmpfsMountOption = "NOSUID"
	TmpfsMountOption_DEV TmpfsMountOption = "DEV"
	TmpfsMountOption_NODEV TmpfsMountOption = "NODEV"
	TmpfsMountOption_EXEC TmpfsMountOption = "EXEC"
	TmpfsMountOption_NOEXEC TmpfsMountOption = "NOEXEC"
	TmpfsMountOption_SYNC TmpfsMountOption = "SYNC"
	TmpfsMountOption_ASYNC TmpfsMountOption = "ASYNC"
	TmpfsMountOption_DIRSYNC TmpfsMountOption = "DIRSYNC"
	TmpfsMountOption_REMOUNT TmpfsMountOption = "REMOUNT"
	TmpfsMountOption_MAND TmpfsMountOption = "MAND"
	TmpfsMountOption_NOMAND TmpfsMountOption = "NOMAND"
	TmpfsMountOption_ATIME TmpfsMountOption = "ATIME"
	TmpfsMountOption_NOATIME TmpfsMountOption = "NOATIME"
	TmpfsMountOption_DIRATIME TmpfsMountOption = "DIRATIME"
	TmpfsMountOption_NODIRATIME TmpfsMountOption = "NODIRATIME"
	TmpfsMountOption_BIND TmpfsMountOption = "BIND"
	TmpfsMountOption_RBIND TmpfsMountOption = "RBIND"
	TmpfsMountOption_UNBINDABLE TmpfsMountOption = "UNBINDABLE"
	TmpfsMountOption_RUNBINDABLE TmpfsMountOption = "RUNBINDABLE"
	TmpfsMountOption_PRIVATE TmpfsMountOption = "PRIVATE"
	TmpfsMountOption_RPRIVATE TmpfsMountOption = "RPRIVATE"
	TmpfsMountOption_SHARED TmpfsMountOption = "SHARED"
	TmpfsMountOption_RSHARED TmpfsMountOption = "RSHARED"
	TmpfsMountOption_SLAVE TmpfsMountOption = "SLAVE"
	TmpfsMountOption_RSLAVE TmpfsMountOption = "RSLAVE"
	TmpfsMountOption_RELATIME TmpfsMountOption = "RELATIME"
	TmpfsMountOption_NORELATIME TmpfsMountOption = "NORELATIME"
	TmpfsMountOption_STRICTATIME TmpfsMountOption = "STRICTATIME"
	TmpfsMountOption_NOSTRICTATIME TmpfsMountOption = "NOSTRICTATIME"
	TmpfsMountOption_MODE TmpfsMountOption = "MODE"
	TmpfsMountOption_UID TmpfsMountOption = "UID"
	TmpfsMountOption_GID TmpfsMountOption = "GID"
	TmpfsMountOption_NR_INODES TmpfsMountOption = "NR_INODES"
	TmpfsMountOption_NR_BLOCKS TmpfsMountOption = "NR_BLOCKS"
	TmpfsMountOption_MPOL TmpfsMountOption = "MPOL"
)

// The properties for enabling target tracking scaling based on a custom CloudWatch metric.
//
// TODO: EXAMPLE
//
type TrackCustomMetricProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// The custom CloudWatch metric to track.
	//
	// The metric must represent utilization; that is, you will always get the following behavior:
	//
	// - metric > targetValue => scale out
	// - metric < targetValue => scale in
	Metric awscloudwatch.IMetric `json:"metric"`
	// The target value for the custom CloudWatch metric.
	TargetValue *float64 `json:"targetValue"`
}

// The ulimit settings to pass to the container.
//
// NOTE: Does not work for Windows containers.
//
// TODO: EXAMPLE
//
type Ulimit struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `json:"hardLimit"`
	// The type of the ulimit.
	//
	// For more information, see [UlimitName](https://docs.aws.amazon.com/cdk/api/latest/typescript/api/aws-ecs/ulimitname.html#aws_ecs_UlimitName).
	Name UlimitName `json:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `json:"softLimit"`
}

// Type of resource to set a limit on.
type UlimitName string

const (
	UlimitName_CORE UlimitName = "CORE"
	UlimitName_CPU UlimitName = "CPU"
	UlimitName_DATA UlimitName = "DATA"
	UlimitName_FSIZE UlimitName = "FSIZE"
	UlimitName_LOCKS UlimitName = "LOCKS"
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	UlimitName_NICE UlimitName = "NICE"
	UlimitName_NOFILE UlimitName = "NOFILE"
	UlimitName_NPROC UlimitName = "NPROC"
	UlimitName_RSS UlimitName = "RSS"
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	UlimitName_RTTIME UlimitName = "RTTIME"
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	UlimitName_STACK UlimitName = "STACK"
)

// A data volume used in a task definition.
//
// For tasks that use a Docker volume, specify a DockerVolumeConfiguration.
// For tasks that use a bind mount host volume, specify a host and optional sourcePath.
//
// For more information, see [Using Data Volumes in Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html).
//
// TODO: EXAMPLE
//
type Volume struct {
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.
	// This name is referenced in the sourceVolume parameter of container definition mountPoints.
	Name *string `json:"name"`
	// This property is specified when you are using Docker volumes.
	//
	// Docker volumes are only supported when you are using the EC2 launch type.
	// Windows containers only support the use of the local driver.
	// To use bind mounts, specify a host instead.
	DockerVolumeConfiguration *DockerVolumeConfiguration `json:"dockerVolumeConfiguration"`
	// This property is specified when you are using Amazon EFS.
	//
	// When specifying Amazon EFS volumes in tasks using the Fargate launch type,
	// Fargate creates a supervisor container that is responsible for managing the Amazon EFS volume.
	// The supervisor container uses a small amount of the task's memory.
	// The supervisor container is visible when querying the task metadata version 4 endpoint,
	// but is not visible in CloudWatch Container Insights.
	EfsVolumeConfiguration *EfsVolumeConfiguration `json:"efsVolumeConfiguration"`
	// This property is specified when you are using bind mount host volumes.
	//
	// Bind mount host volumes are supported when you are using either the EC2 or Fargate launch types.
	// The contents of the host parameter determine whether your bind mount host volume persists on the
	// host container instance and where it is stored. If the host parameter is empty, then the Docker
	// daemon assigns a host path for your data volume. However, the data is not guaranteed to persist
	// after the containers associated with it stop running.
	Host *Host `json:"host"`
}

// The details on a data volume from another container in the same task definition.
//
// TODO: EXAMPLE
//
type VolumeFrom struct {
	// Specifies whether the container has read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `json:"readOnly"`
	// The name of another container within the same task definition from which to mount volumes.
	SourceContainer *string `json:"sourceContainer"`
}

// ECS-optimized Windows version list.
type WindowsOptimizedVersion string

const (
	WindowsOptimizedVersion_SERVER_2019 WindowsOptimizedVersion = "SERVER_2019"
	WindowsOptimizedVersion_SERVER_2016 WindowsOptimizedVersion = "SERVER_2016"
)


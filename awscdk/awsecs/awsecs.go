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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   addAutoScalingGroupCapacityOptions := &addAutoScalingGroupCapacityOptions{
//   	canContainersAccessInstanceRole: jsii.Boolean(false),
//   	machineImageType: awscdk.Aws_ecs.machineImageType_AMAZON_LINUX_2,
//   	spotInstanceDraining: jsii.Boolean(false),
//   	topicEncryptionKey: key,
//   }
//
type AddAutoScalingGroupCapacityOptions struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `field:"optional" json:"canContainersAccessInstanceRole" yaml:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `field:"optional" json:"spotInstanceDraining" yaml:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `field:"optional" json:"topicEncryptionKey" yaml:"topicEncryptionKey"`
}

// The properties for adding instance capacity to an AutoScalingGroup.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   // Either add default capacity
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   // Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux(),
//   	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
//   	// machineImage: EcsOptimizedImage.amazonLinux2(),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   cluster.addAutoScalingGroup(autoScalingGroup)
//
type AddCapacityOptions struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `field:"optional" json:"canContainersAccessInstanceRole" yaml:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `field:"optional" json:"spotInstanceDraining" yaml:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `field:"optional" json:"topicEncryptionKey" yaml:"topicEncryptionKey"`
	// Whether the instances can initiate connections to anywhere by default.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*awsautoscaling.BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	GroupMetrics *[]awsautoscaling.GroupMetrics `field:"optional" json:"groupMetrics" yaml:"groupMetrics"`
	// Configuration for health checks.
	HealthCheck awsautoscaling.HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	IgnoreUnmodifiedSizeProperties *bool `field:"optional" json:"ignoreUnmodifiedSizeProperties" yaml:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	InstanceMonitoring awsautoscaling.Monitoring `field:"optional" json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Maximum number of instances in the fleet.
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
	MaxInstanceLifetime awscdk.Duration `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
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
	NewInstancesProtectedFromScaleIn *bool `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	Notifications *[]*awsautoscaling.NotificationConfiguration `field:"optional" json:"notifications" yaml:"notifications"`
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
	Signals awsautoscaling.Signals `field:"optional" json:"signals" yaml:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	TerminationPolicies *[]awsautoscaling.TerminationPolicy `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	UpdatePolicy awsautoscaling.UpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Where to place instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The EC2 instance type to use when launching instances into the AutoScalingGroup.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
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
	MachineImage awsec2.IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
}

// The ECS-optimized AMI variant to use.
//
// For more information, see
// [Amazon ECS-optimized AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
//
// Example:
//   var cluster cluster
//
//
//   cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(ecs.amiHardwareType_ARM),
//   })
//
type AmiHardwareType string

const (
	// Use the standard Amazon ECS-optimized AMI.
	AmiHardwareType_STANDARD AmiHardwareType = "STANDARD"
	// Use the Amazon ECS GPU-optimized AMI.
	AmiHardwareType_GPU AmiHardwareType = "GPU"
	// Use the Amazon ECS-optimized Amazon Linux 2 (arm64) AMI.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfiguration := awscdk.Aws_ecs.NewAppMeshProxyConfiguration(&appMeshProxyConfigurationConfigProps{
//   	containerName: jsii.String("containerName"),
//   	properties: &appMeshProxyConfigurationProps{
//   		appPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		proxyEgressPort: jsii.Number(123),
//   		proxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		egressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		egressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ignoredGID: jsii.Number(123),
//   		ignoredUID: jsii.Number(123),
//   	},
//   })
//
type AppMeshProxyConfiguration interface {
	ProxyConfiguration
	// Called when the proxy configuration is configured on a task definition.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationConfigProps := &appMeshProxyConfigurationConfigProps{
//   	containerName: jsii.String("containerName"),
//   	properties: &appMeshProxyConfigurationProps{
//   		appPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		proxyEgressPort: jsii.Number(123),
//   		proxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		egressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		egressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ignoredGID: jsii.Number(123),
//   		ignoredUID: jsii.Number(123),
//   	},
//   }
//
type AppMeshProxyConfigurationConfigProps struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin.
	Properties *AppMeshProxyConfigurationProps `field:"required" json:"properties" yaml:"properties"`
}

// Interface for setting the properties of proxy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationProps := &appMeshProxyConfigurationProps{
//   	appPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	proxyEgressPort: jsii.Number(123),
//   	proxyIngressPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	egressIgnoredIPs: []*string{
//   		jsii.String("egressIgnoredIPs"),
//   	},
//   	egressIgnoredPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	ignoredGID: jsii.Number(123),
//   	ignoredUID: jsii.Number(123),
//   }
//
type AppMeshProxyConfigurationProps struct {
	// The list of ports that the application uses.
	//
	// Network traffic to these ports is forwarded to the ProxyIngressPort and ProxyEgressPort.
	AppPorts *[]*float64 `field:"required" json:"appPorts" yaml:"appPorts"`
	// Specifies the port that outgoing traffic from the AppPorts is directed to.
	ProxyEgressPort *float64 `field:"required" json:"proxyEgressPort" yaml:"proxyEgressPort"`
	// Specifies the port that incoming traffic to the AppPorts is directed to.
	ProxyIngressPort *float64 `field:"required" json:"proxyIngressPort" yaml:"proxyIngressPort"`
	// The egress traffic going to these specified IP addresses is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredIPs *[]*string `field:"optional" json:"egressIgnoredIPs" yaml:"egressIgnoredIPs"`
	// The egress traffic going to these specified ports is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredPorts *[]*float64 `field:"optional" json:"egressIgnoredPorts" yaml:"egressIgnoredPorts"`
	// The group ID (GID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredUID is specified, this field can be empty.
	IgnoredGID *float64 `field:"optional" json:"ignoredGID" yaml:"ignoredGID"`
	// The user ID (UID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredGID is specified, this field can be empty.
	IgnoredUID *float64 `field:"optional" json:"ignoredUID" yaml:"ignoredUID"`
}

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
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
//   	minCapacity: jsii.Number(0),
//   	maxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   })
//   cluster.addAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: capacityProvider.capacityProviderName,
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type AsgCapacityProvider interface {
	constructs.Construct
	// Auto Scaling Group.
	AutoScalingGroup() awsautoscaling.AutoScalingGroup
	// Capacity provider name.
	CapacityProviderName() *string
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
//   	minCapacity: jsii.Number(0),
//   	maxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   })
//   cluster.addAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: capacityProvider.capacityProviderName,
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type AsgCapacityProviderProps struct {
	// Specifies whether the containers can access the container instance role.
	CanContainersAccessInstanceRole *bool `field:"optional" json:"canContainersAccessInstanceRole" yaml:"canContainersAccessInstanceRole"`
	// What type of machine image this is.
	//
	// Depending on the setting, different UserData will automatically be added
	// to the `AutoScalingGroup` to configure it properly for use with ECS.
	//
	// If you create an `AutoScalingGroup` yourself and are adding it via
	// `addAutoScalingGroup()`, you must specify this value. If you are adding an
	// `autoScalingGroup` via `addCapacity`, this value will be determined
	// from the `machineImage` you pass.
	MachineImageType MachineImageType `field:"optional" json:"machineImageType" yaml:"machineImageType"`
	// Specify whether to enable Automated Draining for Spot Instances running Amazon ECS Services.
	//
	// For more information, see [Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html).
	SpotInstanceDraining *bool `field:"optional" json:"spotInstanceDraining" yaml:"spotInstanceDraining"`
	// If {@link AddAutoScalingGroupCapacityOptions.taskDrainTime} is non-zero, then the ECS cluster creates an SNS Topic to as part of a system to drain instances of tasks when the instance is being shut down. If this property is provided, then this key will be used to encrypt the contents of that SNS Topic. See [SNS Data Encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for more information.
	TopicEncryptionKey awskms.IKey `field:"optional" json:"topicEncryptionKey" yaml:"topicEncryptionKey"`
	// The autoscaling group to add as a Capacity Provider.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// The name of the capacity provider.
	//
	// If a name is specified,
	// it cannot start with `aws`, `ecs`, or `fargate`. If no name is specified,
	// a default name in the CFNStackName-CFNResourceName-RandomString format is used.
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// Whether to enable managed scaling.
	EnableManagedScaling *bool `field:"optional" json:"enableManagedScaling" yaml:"enableManagedScaling"`
	// Whether to enable managed termination protection.
	EnableManagedTerminationProtection *bool `field:"optional" json:"enableManagedTerminationProtection" yaml:"enableManagedTerminationProtection"`
	// Maximum scaling step size.
	//
	// In most cases this should be left alone.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// Minimum scaling step size.
	//
	// In most cases this should be left alone.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Target capacity percent.
	//
	// In most cases this should be left alone.
	TargetCapacityPercent *float64 `field:"optional" json:"targetCapacityPercent" yaml:"targetCapacityPercent"`
}

// Environment file from a local directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetEnvironmentFile := awscdk.Aws_ecs.NewAssetEnvironmentFile(jsii.String("path"), &assetOptions{
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	readers: []*iGrantable{
//   		grantable,
//   	},
//   })
//
type AssetEnvironmentFile interface {
	EnvironmentFile
	// The path to the asset file or directory.
	Path() *string
	// Called when the container is initialized to allow this object to bind to the stack.
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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import ecsPatterns "github.com/aws/aws-cdk-go/awscdk"
//   import cxapi "github.com/aws/aws-cdk-go/awscdk"
//   import path "github.com/aws-samples/dummy/path"
//
//   app := awscdk.NewApp()
//
//   stack := awscdk.NewStack(app, jsii.String("aws-ecs-patterns-queue"))
//   stack.node.setContext(cxapi.eCS_REMOVE_DEFAULT_DESIRED_COUNT, jsii.Boolean(true))
//
//   vpc := ec2.NewVpc(stack, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(2),
//   })
//
//   ecsPatterns.NewQueueProcessingFargateService(stack, jsii.String("QueueProcessingService"), &queueProcessingFargateServiceProps{
//   	vpc: vpc,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.NewAssetImage(path.join(__dirname, jsii.String(".."), jsii.String("sqs-reader"))),
//   })
//
type AssetImage interface {
	ContainerImage
	// Called when the image is used by a ContainerDefinition.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//
//   assetImageProps := &assetImageProps{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	target: jsii.String("target"),
//   }
//
type AssetImageProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *awsecrassets.DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	NetworkMode awsecrassets.NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Docker target to build to.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

// The options for using a cloudmap service.
//
// Example:
//   var cloudMapService service
//   var ecsService fargateService
//
//
//   ecsService.associateCloudMapService(&associateCloudMapServiceOptions{
//   	service: cloudMapService,
//   })
//
type AssociateCloudMapServiceOptions struct {
	// The cloudmap service to register with.
	Service awsservicediscovery.IService `field:"required" json:"service" yaml:"service"`
	// The container to point to for a SRV record.
	Container ContainerDefinition `field:"optional" json:"container" yaml:"container"`
	// The port to point to for a SRV record.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
}

// The authorization configuration details for the Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfig := &authorizationConfig{
//   	accessPointId: jsii.String("accessPointId"),
//   	iam: jsii.String("iam"),
//   }
//
type AuthorizationConfig struct {
	// The access point ID to use.
	//
	// If an access point is specified, the root directory value will be
	// relative to the directory set for the access point.
	// If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the EFSVolumeConfiguration.
	//
	// Valid values: ENABLED | DISABLED.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

// A log driver that sends log information to CloudWatch Logs.
//
// Example:
//   var cluster cluster
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewAwsLogDriver(&awsLogDriverProps{
//   		streamPrefix: jsii.String("EventDemo"),
//   		mode: ecs.awsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.addTarget(targets.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	taskCount: jsii.Number(1),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerName: jsii.String("TheContainer"),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("I_WAS_TRIGGERED"),
//   					value: jsii.String("From CloudWatch Events"),
//   				},
//   			},
//   		},
//   	},
//   }))
//
type AwsLogDriver interface {
	LogDriver
	// The log group to send log streams to.
	//
	// Only available after the LogDriver has been bound to a ContainerDefinition.
	LogGroup() awslogs.ILogGroup
	SetLogGroup(val awslogs.ILogGroup)
	// Called when the log driver is configured on a container.
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
// Example:
//   var cluster cluster
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewAwsLogDriver(&awsLogDriverProps{
//   		streamPrefix: jsii.String("EventDemo"),
//   		mode: ecs.awsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.addTarget(targets.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	taskCount: jsii.Number(1),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerName: jsii.String("TheContainer"),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("I_WAS_TRIGGERED"),
//   					value: jsii.String("From CloudWatch Events"),
//   				},
//   			},
//   		},
//   	},
//   }))
//
type AwsLogDriverMode string

const (
	// (default) direct, blocking delivery from container to driver.
	AwsLogDriverMode_BLOCKING AwsLogDriverMode = "BLOCKING"
	// The non-blocking message delivery mode prevents applications from blocking due to logging back pressure.
	//
	// Applications are likely to fail in unexpected ways when STDERR or STDOUT streams block.
	AwsLogDriverMode_NON_BLOCKING AwsLogDriverMode = "NON_BLOCKING"
)

// Specifies the awslogs log driver configuration options.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
//
type AwsLogDriverProps struct {
	// Prefix for the log streams.
	//
	// The awslogs-stream-prefix option allows you to associate a log stream
	// with the specified prefix, the container name, and the ID of the Amazon
	// ECS task to which the container belongs. If you specify a prefix with
	// this option, then the log stream takes the following format:
	//
	// prefix-name/container-name/ecs-task-id.
	StreamPrefix *string `field:"required" json:"streamPrefix" yaml:"streamPrefix"`
	// This option defines a multiline start pattern in Python strftime format.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// The log group to log to.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs when the log group is automatically created by this construct.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The delivery mode of log messages from the container to awslogs.
	Mode AwsLogDriverMode `field:"optional" json:"mode" yaml:"mode"`
	// This option defines a multiline start pattern using a regular expression.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	//
	// This option is ignored if datetimeFormat is also configured.
	MultilinePattern *string `field:"optional" json:"multilinePattern" yaml:"multilinePattern"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseLogDriverProps := &baseLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   }
//
type BaseLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

// The base class for Ec2Service and FargateService services.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   service := ecs.baseService.fromServiceArnWithCluster(this, jsii.String("EcsService"), jsii.String("arn:aws:ecs:us-east-1:123456789012:service/myClusterName/myServiceName"))
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   buildOutput := codepipeline.NewArtifact()
//   // add source and build stages to the pipeline as usual...
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		codepipeline_actions.NewEcsDeployAction(&ecsDeployActionProps{
//   			actionName: jsii.String("DeployAction"),
//   			service: service,
//   			input: buildOutput,
//   		}),
//   	},
//   })
//
type BaseService interface {
	awscdk.Resource
	IBaseService
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// The details of the AWS Cloud Map service.
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
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
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	TaskDefinition() TaskDefinition
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
	// Associates this service with a CloudMap service.
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	// This method is called to attach this service to an Application Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// This method is called to create a networkConfiguration.
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	// Enable CloudMap service discovery for the service.
	//
	// Returns: The created CloudMap service.
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
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
	// Return a load balancing target for a specific container and port.
	//
	// Use this function to create a load balancer target if you want to load balance to
	// another container than the first essential container or the first mapped port on
	// the container.
	//
	// Use the return value of this function where you would normally use a load balancer
	// target, instead of the `Service` object itself.
	//
	// Example:
	//   var listener applicationListener
	//   var service baseService
	//
	//   listener.addTargets(jsii.String("ECS"), &addApplicationTargetsProps{
	//   	port: jsii.Number(80),
	//   	targets: []iApplicationLoadBalancerTarget{
	//   		service.loadBalancerTarget(&loadBalancerTargetOptions{
	//   			containerName: jsii.String("MyContainer"),
	//   			containerPort: jsii.Number(1234),
	//   		}),
	//   	},
	//   })
	//
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
	//
	// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
	//
	// Example:
	//   var listener applicationListener
	//   var service baseService
	//
	//   service.registerLoadBalancerTargets(&ecsTarget{
	//   	containerName: jsii.String("web"),
	//   	containerPort: jsii.Number(80),
	//   	newTargetGroupId: jsii.String("ECS"),
	//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
	//   		protocol: elbv2.applicationProtocol_HTTPS,
	//   	}),
	//   })
	//
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	// Returns a string representation of this construct.
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

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func BaseService_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
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

func (b *jsiiProxy_BaseService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BaseService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		b,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

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

func (b *jsiiProxy_BaseService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		b,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

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

func (b *jsiiProxy_BaseService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		b,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var namespace iNamespace
//
//   baseServiceOptions := &baseServiceOptions{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(false),
//   	},
//   	cloudMapOptions: &cloudMapOptions{
//   		cloudMapNamespace: namespace,
//   		container: containerDefinition,
//   		containerPort: jsii.Number(123),
//   		dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   		dnsTtl: cdk.duration.minutes(jsii.Number(30)),
//   		failureThreshold: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	deploymentController: &deploymentController{
//   		type: awscdk.Aws_ecs.deploymentControllerType_ECS,
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableECSManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriod: cdk.*duration.minutes(jsii.Number(30)),
//   	maxHealthyPercent: jsii.Number(123),
//   	minHealthyPercent: jsii.Number(123),
//   	propagateTags: awscdk.*Aws_ecs.propagatedTagSource_SERVICE,
//   	serviceName: jsii.String("serviceName"),
//   }
//
type BaseServiceOptions struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

// Complete base service properties that are required to be supplied by the implementation of the BaseService class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var namespace iNamespace
//
//   baseServiceProps := &baseServiceProps{
//   	cluster: cluster,
//   	launchType: awscdk.Aws_ecs.launchType_EC2,
//
//   	// the properties below are optional
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(false),
//   	},
//   	cloudMapOptions: &cloudMapOptions{
//   		cloudMapNamespace: namespace,
//   		container: containerDefinition,
//   		containerPort: jsii.Number(123),
//   		dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   		dnsTtl: cdk.duration.minutes(jsii.Number(30)),
//   		failureThreshold: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	deploymentController: &deploymentController{
//   		type: awscdk.*Aws_ecs.deploymentControllerType_ECS,
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableECSManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriod: cdk.*duration.minutes(jsii.Number(30)),
//   	maxHealthyPercent: jsii.Number(123),
//   	minHealthyPercent: jsii.Number(123),
//   	propagateTags: awscdk.*Aws_ecs.propagatedTagSource_SERVICE,
//   	serviceName: jsii.String("serviceName"),
//   }
//
type BaseServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The launch type on which to run your service.
	//
	// LaunchType will be omitted if capacity provider strategies are specified on the service.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
	//
	// Valid values are: LaunchType.ECS or LaunchType.FARGATE or LaunchType.EXTERNAL
	//
	LaunchType LaunchType `field:"required" json:"launchType" yaml:"launchType"`
}

// Instance resource used for bin packing.
type BinPackResource string

const (
	// Fill up hosts' CPU allocations first.
	BinPackResource_CPU BinPackResource = "CPU"
	// Fill up hosts' memory allocations first.
	BinPackResource_MEMORY BinPackResource = "MEMORY"
)

// Construct an Bottlerocket image from the latest AMI published in SSM.
//
// Example:
//   var cluster cluster
//
//
//   cluster.addCapacity(jsii.String("bottlerocket-asg"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c5.large")),
//   	machineImage: ecs.NewBottleRocketImage(),
//   })
//
type BottleRocketImage interface {
	awsec2.IMachineImage
	// Return the correct image.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bottleRocketImageProps := &bottleRocketImageProps{
//   	architecture: awscdk.Aws_ec2.instanceArchitecture_ARM_64,
//   	cachedInContext: jsii.Boolean(false),
//   	variant: awscdk.Aws_ecs.bottlerocketEcsVariant_AWS_ECS_1,
//   }
//
type BottleRocketImageProps struct {
	// The CPU architecture.
	Architecture awsec2.InstanceArchitecture `field:"optional" json:"architecture" yaml:"architecture"`
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
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// The Amazon ECS variant to use.
	//
	// Only `aws-ecs-1` is currently available.
	Variant BottlerocketEcsVariant `field:"optional" json:"variant" yaml:"variant"`
}

// Amazon ECS variant.
type BottlerocketEcsVariant string

const (
	// aws-ecs-1 variant.
	BottlerocketEcsVariant_AWS_ECS_1 BottlerocketEcsVariant = "AWS_ECS_1"
)

// The built-in container instance attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   builtInAttributes := awscdk.Aws_ecs.NewBuiltInAttributes()
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategy := &capacityProviderStrategy{
//   	capacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	base: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type CapacityProviderStrategy struct {
	// The name of the capacity provider.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one
	// capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default
	// value of 0 is used.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The weight value is taken into consideration after the base value, if defined, is satisfied.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// A CloudFormation `AWS::ECS::CapacityProvider`.
//
// The `AWS::ECS::CapacityProvider` resource creates an Amazon Elastic Container Service (Amazon ECS) capacity provider. Capacity providers are associated with an Amazon ECS cluster and are used in capacity provider strategies to facilitate cluster auto scaling.
//
// Only capacity providers using an Auto Scaling group can be created. Amazon ECS tasks on AWS Fargate use the `FARGATE` and `FARGATE_SPOT` capacity providers which are already created and available to all accounts in Regions supported by AWS Fargate .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProvider := awscdk.Aws_ecs.NewCfnCapacityProvider(this, jsii.String("MyCfnCapacityProvider"), &cfnCapacityProviderProps{
//   	autoScalingGroupProvider: &autoScalingGroupProviderProperty{
//   		autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   		// the properties below are optional
//   		managedScaling: &managedScalingProperty{
//   			instanceWarmupPeriod: jsii.Number(123),
//   			maximumScalingStepSize: jsii.Number(123),
//   			minimumScalingStepSize: jsii.Number(123),
//   			status: jsii.String("status"),
//   			targetCapacity: jsii.Number(123),
//   		},
//   		managedTerminationProtection: jsii.String("managedTerminationProtection"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCapacityProvider interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Auto Scaling group settings for the capacity provider.
	AutoScalingGroupProvider() interface{}
	SetAutoScalingGroupProvider(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnCapacityProvider) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCapacityProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCapacityProvider) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupProviderProperty := &autoScalingGroupProviderProperty{
//   	autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   	// the properties below are optional
//   	managedScaling: &managedScalingProperty{
//   		instanceWarmupPeriod: jsii.Number(123),
//   		maximumScalingStepSize: jsii.Number(123),
//   		minimumScalingStepSize: jsii.Number(123),
//   		status: jsii.String("status"),
//   		targetCapacity: jsii.Number(123),
//   	},
//   	managedTerminationProtection: jsii.String("managedTerminationProtection"),
//   }
//
type CfnCapacityProvider_AutoScalingGroupProviderProperty struct {
	// The Amazon Resource Name (ARN) or short name that identifies the Auto Scaling group.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The managed scaling settings for the Auto Scaling group capacity provider.
	ManagedScaling interface{} `field:"optional" json:"managedScaling" yaml:"managedScaling"`
	// The managed termination protection setting to use for the Auto Scaling group capacity provider.
	//
	// This determines whether the Auto Scaling group has managed termination protection. The default is disabled.
	//
	// > When using managed termination protection, managed scaling must also be used otherwise managed termination protection doesn't work.
	//
	// When managed termination protection is enabled, Amazon ECS prevents the Amazon EC2 instances in an Auto Scaling group that contain tasks from being terminated during a scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must have instance protection from scale-in actions enabled as well. For more information, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *AWS Auto Scaling User Guide* .
	//
	// When managed termination protection is disabled, your Amazon EC2 instances aren't protected from termination when the Auto Scaling group scales in.
	ManagedTerminationProtection *string `field:"optional" json:"managedTerminationProtection" yaml:"managedTerminationProtection"`
}

// The `ManagedScaling` property specifies the settings for the Auto Scaling group capacity provider.
//
// When managed scaling is enabled, Amazon ECS manages the scale-in and scale-out actions of the Auto Scaling group. Amazon ECS manages a target tracking scaling policy using an Amazon ECS-managed CloudWatch metric with the specified `targetCapacity` value as the target value for the metric. For more information, see [Using Managed Scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling) in the *Amazon Elastic Container Service Developer Guide* .
//
// If managed scaling is disabled, the user must manage the scaling of the Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingProperty := &managedScalingProperty{
//   	instanceWarmupPeriod: jsii.Number(123),
//   	maximumScalingStepSize: jsii.Number(123),
//   	minimumScalingStepSize: jsii.Number(123),
//   	status: jsii.String("status"),
//   	targetCapacity: jsii.Number(123),
//   }
//
type CfnCapacityProvider_ManagedScalingProperty struct {
	// The period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group.
	//
	// If this parameter is omitted, the default value of `300` seconds is used.
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// The maximum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `10000` is used.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// The minimum number of container instances that Amazon ECS scales in or scales out at one time.
	//
	// If this parameter is omitted, the default value of `1` is used.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Determines whether to use managed scaling for the capacity provider.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target capacity value for the capacity provider.
	//
	// The specified value must be greater than `0` and less than or equal to `100` . A value of `100` results in the Amazon EC2 instances in your Auto Scaling group being completely used.
	TargetCapacity *float64 `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

// Properties for defining a `CfnCapacityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProviderProps := &cfnCapacityProviderProps{
//   	autoScalingGroupProvider: &autoScalingGroupProviderProperty{
//   		autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   		// the properties below are optional
//   		managedScaling: &managedScalingProperty{
//   			instanceWarmupPeriod: jsii.Number(123),
//   			maximumScalingStepSize: jsii.Number(123),
//   			minimumScalingStepSize: jsii.Number(123),
//   			status: jsii.String("status"),
//   			targetCapacity: jsii.Number(123),
//   		},
//   		managedTerminationProtection: jsii.String("managedTerminationProtection"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCapacityProviderProps struct {
	// The Auto Scaling group settings for the capacity provider.
	AutoScalingGroupProvider interface{} `field:"required" json:"autoScalingGroupProvider" yaml:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
	Name *string `field:"optional" json:"name" yaml:"name"`
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
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ECS::Cluster`.
//
// The `AWS::ECS::Cluster` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_ecs.NewCfnCluster(this, jsii.String("MyCfnCluster"), &cfnClusterProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	clusterSettings: []interface{}{
//   		&clusterSettingsProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	configuration: &clusterConfigurationProperty{
//   		executeCommandConfiguration: &executeCommandConfigurationProperty{
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			logConfiguration: &executeCommandLogConfigurationProperty{
//   				cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				s3BucketName: jsii.String("s3BucketName"),
//   				s3EncryptionEnabled: jsii.Boolean(false),
//   				s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   			},
//   			logging: jsii.String("logging"),
//   		},
//   	},
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as `arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster` .
	AttrArn() *string
	// The short name of one or more capacity providers to associate with the cluster.
	//
	// A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created and not already associated with another cluster.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	CapacityProviders() *[]*string
	SetCapacityProviders(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	ClusterName() *string
	SetClusterName(val *string)
	// The setting to use when creating a cluster.
	//
	// This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	ClusterSettings() interface{}
	SetClusterSettings(val interface{})
	// The execute command configuration for the cluster.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default capacity provider strategy for the cluster.
	//
	// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
	DefaultCapacityProviderStrategy() interface{}
	SetDefaultCapacityProviderStrategy(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategyItemProperty := &capacityProviderStrategyItemProperty{
//   	base: jsii.Number(123),
//   	capacityProvider: jsii.String("capacityProvider"),
//   	weight: jsii.Number(123),
//   }
//
type CfnCluster_CapacityProviderStrategyItemProperty struct {
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The short name of the capacity provider.
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that's run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// The execute command configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigurationProperty := &clusterConfigurationProperty{
//   	executeCommandConfiguration: &executeCommandConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		logConfiguration: &executeCommandLogConfigurationProperty{
//   			cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3EncryptionEnabled: jsii.Boolean(false),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		logging: jsii.String("logging"),
//   	},
//   }
//
type CfnCluster_ClusterConfigurationProperty struct {
	// The details of the execute command configuration.
	ExecuteCommandConfiguration interface{} `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

// The settings to use when creating a cluster.
//
// This parameter is used to turn on CloudWatch Container Insights for a cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSettingsProperty := &clusterSettingsProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnCluster_ClusterSettingsProperty struct {
	// The name of the cluster setting.
	//
	// The only supported value is `containerInsights` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value to set for the cluster setting.
	//
	// The supported values are `enabled` and `disabled` . If `enabled` is specified, CloudWatch Container Insights will be enabled for the cluster, otherwise it will be disabled unless the `containerInsights` account setting is enabled. If a cluster value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The details of the execute command configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executeCommandConfigurationProperty := &executeCommandConfigurationProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logConfiguration: &executeCommandLogConfigurationProperty{
//   		cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   		cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3EncryptionEnabled: jsii.Boolean(false),
//   		s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	logging: jsii.String("logging"),
//   }
//
type CfnCluster_ExecuteCommandConfigurationProperty struct {
	// Specify an AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When `logging=OVERRIDE` is specified, a `logConfiguration` must be provided.
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The log setting to use for redirecting logs for your execute command results. The following log settings are available.
	//
	// - `NONE` : The execute command session is not logged.
	// - `DEFAULT` : The `awslogs` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no `awslogs` log driver is configured in the task definition, the output won't be logged.
	// - `OVERRIDE` : Specify the logging details as a part of `logConfiguration` . If the `OVERRIDE` logging option is specified, the `logConfiguration` is required.
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
}

// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executeCommandLogConfigurationProperty := &executeCommandLogConfigurationProperty{
//   	cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   	cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3EncryptionEnabled: jsii.Boolean(false),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnCluster_ExecuteCommandLogConfigurationProperty struct {
	// Determines whether to use encryption on the CloudWatch logs.
	//
	// If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled interface{} `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// > The CloudWatch log group must already be created.
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// The name of the S3 bucket to send logs to.
	//
	// > The S3 bucket must already be created.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Determines whether to use encryption on the S3 logs.
	//
	// If not specified, encryption is not used.
	S3EncryptionEnabled interface{} `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

// A CloudFormation `AWS::ECS::ClusterCapacityProviderAssociations`.
//
// The `AWS::ECS::ClusterCapacityProviderAssociations` resource associates one or more capacity providers and a default capacity provider strategy with a cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterCapacityProviderAssociations := awscdk.Aws_ecs.NewCfnClusterCapacityProviderAssociations(this, jsii.String("MyCfnClusterCapacityProviderAssociations"), &cfnClusterCapacityProviderAssociationsProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	cluster: jsii.String("cluster"),
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyProperty{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   })
//
type CfnClusterCapacityProviderAssociations interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The capacity providers to associate with the cluster.
	CapacityProviders() *[]*string
	SetCapacityProviders(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The cluster the capacity provider association is the target of.
	Cluster() *string
	SetCluster(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy() interface{}
	SetDefaultCapacityProviderStrategy(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnClusterCapacityProviderAssociations) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategyProperty := &capacityProviderStrategyProperty{
//   	capacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	base: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type CfnClusterCapacityProviderAssociations_CapacityProviderStrategyProperty struct {
	// The short name of the capacity provider.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` will not be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that is run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Properties for defining a `CfnClusterCapacityProviderAssociations`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterCapacityProviderAssociationsProps := &cfnClusterCapacityProviderAssociationsProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	cluster: jsii.String("cluster"),
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyProperty{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnClusterCapacityProviderAssociationsProps struct {
	// The capacity providers to associate with the cluster.
	CapacityProviders *[]*string `field:"required" json:"capacityProviders" yaml:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy interface{} `field:"required" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
}

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	clusterSettings: []interface{}{
//   		&clusterSettingsProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	configuration: &clusterConfigurationProperty{
//   		executeCommandConfiguration: &executeCommandConfigurationProperty{
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			logConfiguration: &executeCommandLogConfigurationProperty{
//   				cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				s3BucketName: jsii.String("s3BucketName"),
//   				s3EncryptionEnabled: jsii.Boolean(false),
//   				s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   			},
//   			logging: jsii.String("logging"),
//   		},
//   	},
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterProps struct {
	// The short name of one or more capacity providers to associate with the cluster.
	//
	// A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created and not already associated with another cluster.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The setting to use when creating a cluster.
	//
	// This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	ClusterSettings interface{} `field:"optional" json:"clusterSettings" yaml:"clusterSettings"`
	// The execute command configuration for the cluster.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The default capacity provider strategy for the cluster.
	//
	// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
	DefaultCapacityProviderStrategy interface{} `field:"optional" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
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
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ECS::PrimaryTaskSet`.
//
// Specifies which task set in a service is the primary task set. Any parameters that are updated on the primary task set in a service will transition to the service. This is used when a service uses the `EXTERNAL` deployment controller type. For more information, see [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrimaryTaskSet := awscdk.Aws_ecs.NewCfnPrimaryTaskSet(this, jsii.String("MyCfnPrimaryTaskSet"), &cfnPrimaryTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskSetId: jsii.String("taskSetId"),
//   })
//
type CfnPrimaryTaskSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set exists in.
	Cluster() *string
	SetCluster(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The short name or full Amazon Resource Name (ARN) of the service that the task set exists in.
	Service() *string
	SetService(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The short name or full Amazon Resource Name (ARN) of the task set to set as the primary task set in the deployment.
	TaskSetId() *string
	SetTaskSetId(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnPrimaryTaskSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPrimaryTaskSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnPrimaryTaskSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrimaryTaskSetProps := &cfnPrimaryTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskSetId: jsii.String("taskSetId"),
//   }
//
type CfnPrimaryTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set exists in.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service that the task set exists in.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The short name or full Amazon Resource Name (ARN) of the task set to set as the primary task set in the deployment.
	TaskSetId *string `field:"required" json:"taskSetId" yaml:"taskSetId"`
}

// A CloudFormation `AWS::ECS::Service`.
//
// The `AWS::ECS::Service` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnService := awscdk.Aws_ecs.NewCfnService(this, jsii.String("MyCfnService"), &cfnServiceProps{
//   	capacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	cluster: jsii.String("cluster"),
//   	deploymentConfiguration: &deploymentConfigurationProperty{
//   		deploymentCircuitBreaker: &deploymentCircuitBreakerProperty{
//   			enable: jsii.Boolean(false),
//   			rollback: jsii.Boolean(false),
//   		},
//   		maximumPercent: jsii.Number(123),
//   		minimumHealthyPercent: jsii.Number(123),
//   	},
//   	deploymentController: &deploymentControllerProperty{
//   		type: jsii.String("type"),
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableEcsManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriodSeconds: jsii.Number(123),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			containerName: jsii.String("containerName"),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsvpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	placementStrategies: []interface{}{
//   		&placementStrategyProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	role: jsii.String("role"),
//   	schedulingStrategy: jsii.String("schedulingStrategy"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskDefinition: jsii.String("taskDefinition"),
//   })
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the Amazon ECS service, such as `sample-webapp` .
	AttrName() *string
	// Not currently supported in AWS CloudFormation .
	AttrServiceArn() *string
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
	CapacityProviderStrategy() interface{}
	SetCapacityProviderStrategy(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	Cluster() *string
	SetCluster(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration() interface{}
	SetDeploymentConfiguration(val interface{})
	// The deployment controller to use for the service.
	//
	// If no deployment controller is specified, the default value of `ECS` is used.
	DeploymentController() interface{}
	SetDeploymentController(val interface{})
	// The number of instantiations of the specified task definition to place and keep running on your cluster.
	//
	// For new services, if a desired count is not specified, a default value of `1` is used. When using the `DAEMON` scheduling strategy, the desired count is not required.
	//
	// For existing services, if a desired count is not specified, it is omitted from the operation.
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	// Determines whether the execute command functionality is enabled for the service.
	//
	// If `true` , the execute command functionality is enabled for all containers in tasks as part of the service.
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	//
	// This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of `0` is used.
	//
	// If you do not use an Elastic Load Balancing, we recomend that you use the `startPeriod` in the task definition healtch check parameters. For more information, see [Health check](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html) .
	//
	// If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
	HealthCheckGracePeriodSeconds() *float64
	SetHealthCheckGracePeriodSeconds(val *float64)
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType() *string
	SetLaunchType(val *string)
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the `Role` property, `LoadBalancers` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide* .
	LoadBalancers() interface{}
	SetLoadBalancers(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the `awsvpc` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of five strategy rules per service. For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlacementStrategies() interface{}
	SetPlacementStrategies(val interface{})
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the tasks within the service during service creation. To add tags to a task after service creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	PropagateTags() *string
	SetPropagateTags(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the `awsvpc` network mode. If you specify the `role` parameter, you must also specify a load balancer object with the `loadBalancers` parameter.
	//
	// > If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the `awsvpc` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If your specified role has a path other than `/` , then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	Role() *string
	SetRole(val *string)
	// The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// There are two service scheduler strategies available:
	//
	// - `REPLICA` -The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the `CODE_DEPLOY` or `EXTERNAL` deployment controller types.
	// - `DAEMON` -The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//
	// > Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy.
	SchedulingStrategy() *string
	SetSchedulingStrategy(val *string)
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	ServiceName() *string
	SetServiceName(val *string)
	// The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	//
	// > Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	ServiceRegistries() interface{}
	SetServiceRegistries(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	Tags() awscdk.TagManager
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	//
	// The `revision` is required in order for the resource to stabilize.
	//
	// A task definition must be specified if the service is using either the `ECS` or `CODE_DEPLOY` deployment controllers.
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsVpcConfigurationProperty := &awsVpcConfigurationProperty{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnService_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategyItemProperty := &capacityProviderStrategyItemProperty{
//   	base: jsii.Number(123),
//   	capacityProvider: jsii.String("capacityProvider"),
//   	weight: jsii.Number(123),
//   }
//
type CfnService_CapacityProviderStrategyItemProperty struct {
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The short name of the capacity provider.
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that's run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type.
//
// The `DeploymentCircuitBreaker` property determines whether a service deployment will fail if the service can't reach a steady state. If deployment circuit breaker is enabled, a service deployment will transition to a failed state and stop launching new tasks. If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentCircuitBreakerProperty := &deploymentCircuitBreakerProperty{
//   	enable: jsii.Boolean(false),
//   	rollback: jsii.Boolean(false),
//   }
//
type CfnService_DeploymentCircuitBreakerProperty struct {
	// Determines whether to use the deployment circuit breaker logic for the service.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

// The `DeploymentConfiguration` property specifies optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationProperty := &deploymentConfigurationProperty{
//   	deploymentCircuitBreaker: &deploymentCircuitBreakerProperty{
//   		enable: jsii.Boolean(false),
//   		rollback: jsii.Boolean(false),
//   	},
//   	maximumPercent: jsii.Number(123),
//   	minimumHealthyPercent: jsii.Number(123),
//   }
//
type CfnService_DeploymentConfigurationProperty struct {
	// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type that are not behind a Classic Load Balancer.
	//
	// The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If enabled, a service deployment will transition to a failed state and stop launching new tasks. You can also enable Amazon ECS to roll back your service to the last completed deployment after a failure. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide* .
	DeploymentCircuitBreaker interface{} `field:"optional" json:"deploymentCircuitBreaker" yaml:"deploymentCircuitBreaker"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the `maximumPercent` parameter represents an upper limit on the number of your service's tasks that are allowed in the `RUNNING` or `PENDING` state during a deployment, as a percentage of the `desiredCount` (rounded down to the nearest integer).
	//
	// This parameter enables you to define the deployment batch size. For example, if your service is using the `REPLICA` service scheduler and has a `desiredCount` of four tasks and a `maximumPercent` value of 200%, the scheduler may start four new tasks before stopping the four older tasks (provided that the cluster resources required to do this are available). The default `maximumPercent` value for a service using the `REPLICA` service scheduler is 200%.
	//
	// If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and tasks that use the EC2 launch type, the *maximum percent* value is set to the default value and is used to define the upper limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state. If the tasks in the service use the Fargate launch type, the maximum percent value is not used, although it is returned when describing your service.
	MaximumPercent *float64 `field:"optional" json:"maximumPercent" yaml:"maximumPercent"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the `minimumHealthyPercent` represents a lower limit on the number of your service's tasks that must remain in the `RUNNING` state during a deployment, as a percentage of the `desiredCount` (rounded up to the nearest integer).
	//
	// This parameter enables you to deploy without using additional cluster capacity. For example, if your service has a `desiredCount` of four tasks and a `minimumHealthyPercent` of 50%, the service scheduler may stop two existing tasks to free up cluster capacity before starting two new tasks.
	//
	// For services that *do not* use a load balancer, the following should be noted:
	//
	// - A service is considered healthy if all essential containers within the tasks in the service pass their health checks.
	// - If a task has no essential containers with a health check defined, the service scheduler will wait for 40 seconds after a task reaches a `RUNNING` state before the task is counted towards the minimum healthy percent total.
	// - If a task has one or more essential containers with a health check defined, the service scheduler will wait for the task to reach a healthy status before counting it towards the minimum healthy percent total. A task is considered healthy when all essential containers within the task have passed their health checks. The amount of time the service scheduler can wait for is determined by the container health check settings.
	//
	// For services are that *do* use a load balancer, the following should be noted:
	//
	// - If a task has no essential containers with a health check defined, the service scheduler will wait for the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.
	// - If a task has an essential container with a health check defined, the service scheduler will wait for both the task to reach a healthy status and the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.
	//
	// If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and is running tasks that use the EC2 launch type, the *minimum healthy percent* value is set to the default value and is used to define the lower limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state. If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and is running tasks that use the Fargate launch type, the minimum healthy percent value is not used, although it is returned when describing your service.
	MinimumHealthyPercent *float64 `field:"optional" json:"minimumHealthyPercent" yaml:"minimumHealthyPercent"`
}

// The deployment controller to use for the service.
//
// For more information, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentControllerProperty := &deploymentControllerProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnService_DeploymentControllerProperty struct {
	// The deployment controller type to use. There are three deployment controller types available:.
	//
	// - **ECS** - The rolling update ( `ECS` ) deployment type involves replacing the current running version of the container with the latest version. The number of containers Amazon ECS adds or removes from the service during a rolling update is controlled by adjusting the minimum and maximum number of healthy tasks allowed during a service deployment, as specified in the [DeploymentConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeploymentConfiguration.html) .
	// - **CODE_DEPLOY** - The blue/green ( `CODE_DEPLOY` ) deployment type uses the blue/green deployment model powered by AWS CodeDeploy , which allows you to verify a new deployment of a service before sending production traffic to it.
	// - **EXTERNAL** - The external ( `EXTERNAL` ) deployment type enables you to use any third-party deployment controller for full control over the deployment process for an Amazon ECS service.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The `LoadBalancer` property specifies details on a load balancer that is used with a service.
//
// If the service is using the `CODE_DEPLOY` deployment controller, the service is required to use either an Application Load Balancer or Network Load Balancer. When you are creating an AWS CodeDeploy deployment group, you specify two target groups (referred to as a `targetGroupPair` ). Each target group binds to a separate task set in the deployment. The load balancer can also have up to two listeners, a required listener for production traffic and an optional listener that allows you to test new revisions of the service before routing production traffic to it.
//
// Services with tasks that use the `awsvpc` network mode (for example, those with the Fargate launch type) only support Application Load Balancers and Network Load Balancers. Classic Load Balancers are not supported. Also, when you create any target groups for these services, you must choose `ip` as the target type, not `instance` . Tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerProperty := &loadBalancerProperty{
//   	containerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	containerName: jsii.String("containerName"),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnService_LoadBalancerProperty struct {
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you're using a Classic Load Balancer, omit the target group ARN.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

// The `NetworkConfiguration` property specifies an object representing the network configuration for a task or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	awsvpcConfiguration: &awsVpcConfigurationProperty{
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//
//   		// the properties below are optional
//   		assignPublicIp: jsii.String("assignPublicIp"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   }
//
type CfnService_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	AwsvpcConfiguration interface{} `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

// The `PlacementConstraint` property specifies an object representing a constraint on task placement in the task definition.
//
// For more information, see [Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementConstraintProperty := &placementConstraintProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	expression: jsii.String("expression"),
//   }
//
type CfnService_PlacementConstraintProperty struct {
	// The type of constraint.
	//
	// Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A cluster query language expression to apply to the constraint.
	//
	// The expression can have a maximum length of 2000 characters. You can't specify an expression if the constraint type is `distinctInstance` . For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

// The `PlacementStrategy` property specifies the task placement strategy for a task or service.
//
// For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementStrategyProperty := &placementStrategyProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	field: jsii.String("field"),
//   }
//
type CfnService_PlacementStrategyProperty struct {
	// The type of placement strategy.
	//
	// The `random` placement strategy randomly places tasks on available candidates. The `spread` placement strategy spreads placement across available candidates evenly based on the `field` parameter. The `binpack` strategy places tasks on available candidates that have the least available amount of the resource that's specified with the `field` parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory but still enough to run the task.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The field to apply the placement strategy against.
	//
	// For the `spread` placement strategy, valid values are `instanceId` (or `host` , which has the same effect), or any platform or custom attribute that's applied to a container instance, such as `attribute:ecs.availability-zone` . For the `binpack` placement strategy, valid values are `cpu` and `memory` . For the `random` placement strategy, this field is not used.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

// The `ServiceRegistry` property specifies details of the service registry.
//
// For more information, see [Service Discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceRegistryProperty := &serviceRegistryProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	port: jsii.Number(123),
//   	registryArn: jsii.String("registryArn"),
//   }
//
type CfnService_ServiceRegistryProperty struct {
	// The container name value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition that your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition that your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field might be used if both the `awsvpc` network mode and SRV records are used.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map . For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) .
	RegistryArn *string `field:"optional" json:"registryArn" yaml:"registryArn"`
}

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	capacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	cluster: jsii.String("cluster"),
//   	deploymentConfiguration: &deploymentConfigurationProperty{
//   		deploymentCircuitBreaker: &deploymentCircuitBreakerProperty{
//   			enable: jsii.Boolean(false),
//   			rollback: jsii.Boolean(false),
//   		},
//   		maximumPercent: jsii.Number(123),
//   		minimumHealthyPercent: jsii.Number(123),
//   	},
//   	deploymentController: &deploymentControllerProperty{
//   		type: jsii.String("type"),
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableEcsManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriodSeconds: jsii.Number(123),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			containerName: jsii.String("containerName"),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsvpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	placementStrategies: []interface{}{
//   		&placementStrategyProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	role: jsii.String("role"),
//   	schedulingStrategy: jsii.String("schedulingStrategy"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskDefinition: jsii.String("taskDefinition"),
//   }
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
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// The deployment controller to use for the service.
	//
	// If no deployment controller is specified, the default value of `ECS` is used.
	DeploymentController interface{} `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The number of instantiations of the specified task definition to place and keep running on your cluster.
	//
	// For new services, if a desired count is not specified, a default value of `1` is used. When using the `DAEMON` scheduling strategy, the desired count is not required.
	//
	// For existing services, if a desired count is not specified, it is omitted from the operation.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Determines whether the execute command functionality is enabled for the service.
	//
	// If `true` , the execute command functionality is enabled for all containers in tasks as part of the service.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	//
	// This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of `0` is used.
	//
	// If you do not use an Elastic Load Balancing, we recomend that you use the `startPeriod` in the task definition healtch check parameters. For more information, see [Health check](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html) .
	//
	// If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
	HealthCheckGracePeriodSeconds *float64 `field:"optional" json:"healthCheckGracePeriodSeconds" yaml:"healthCheckGracePeriodSeconds"`
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the `Role` property, `LoadBalancers` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide* .
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the `awsvpc` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of five strategy rules per service. For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlacementStrategies interface{} `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the tasks within the service during service creation. To add tags to a task after service creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the `awsvpc` network mode. If you specify the `role` parameter, you must also specify a load balancer object with the `loadBalancers` parameter.
	//
	// > If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the `awsvpc` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If your specified role has a path other than `/` , then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name `bar` has a path of `/foo/` then you would specify `/foo/bar` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// There are two service scheduler strategies available:
	//
	// - `REPLICA` -The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the `CODE_DEPLOY` or `EXTERNAL` deployment controller types.
	// - `DAEMON` -The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//
	// > Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy.
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	//
	// > Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
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
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The `family` and `revision` ( `family:revision` ) or full ARN of the task definition to run in your service.
	//
	// The `revision` is required in order for the resource to stabilize.
	//
	// A task definition must be specified if the service is using either the `ECS` or `CODE_DEPLOY` deployment controllers.
	TaskDefinition *string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
}

// A CloudFormation `AWS::ECS::TaskDefinition`.
//
// The `AWS::ECS::TaskDefinition` resource describes the container and volume definitions of an Amazon Elastic Container Service (Amazon ECS) task. You can specify which Docker images to use, the required resources, and other configurations related to launching the task definition through an Amazon ECS service or task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinition := awscdk.Aws_ecs.NewCfnTaskDefinition(this, jsii.String("MyCfnTaskDefinition"), &cfnTaskDefinitionProps{
//   	containerDefinitions: []interface{}{
//   		&containerDefinitionProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			dependsOn: []interface{}{
//   				&containerDependencyProperty{
//   					condition: jsii.String("condition"),
//   					containerName: jsii.String("containerName"),
//   				},
//   			},
//   			disableNetworking: jsii.Boolean(false),
//   			dnsSearchDomains: []*string{
//   				jsii.String("dnsSearchDomains"),
//   			},
//   			dnsServers: []*string{
//   				jsii.String("dnsServers"),
//   			},
//   			dockerLabels: map[string]*string{
//   				"dockerLabelsKey": jsii.String("dockerLabels"),
//   			},
//   			dockerSecurityOptions: []*string{
//   				jsii.String("dockerSecurityOptions"),
//   			},
//   			entryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			environment: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			environmentFiles: []interface{}{
//   				&environmentFileProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			essential: jsii.Boolean(false),
//   			extraHosts: []interface{}{
//   				&hostEntryProperty{
//   					hostname: jsii.String("hostname"),
//   					ipAddress: jsii.String("ipAddress"),
//   				},
//   			},
//   			firelensConfiguration: &firelensConfigurationProperty{
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				type: jsii.String("type"),
//   			},
//   			healthCheck: &healthCheckProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				interval: jsii.Number(123),
//   				retries: jsii.Number(123),
//   				startPeriod: jsii.Number(123),
//   				timeout: jsii.Number(123),
//   			},
//   			hostname: jsii.String("hostname"),
//   			image: jsii.String("image"),
//   			interactive: jsii.Boolean(false),
//   			links: []*string{
//   				jsii.String("links"),
//   			},
//   			linuxParameters: &linuxParametersProperty{
//   				capabilities: &kernelCapabilitiesProperty{
//   					add: []*string{
//   						jsii.String("add"),
//   					},
//   					drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				devices: []interface{}{
//   					&deviceProperty{
//   						containerPath: jsii.String("containerPath"),
//   						hostPath: jsii.String("hostPath"),
//   						permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				initProcessEnabled: jsii.Boolean(false),
//   				maxSwap: jsii.Number(123),
//   				sharedMemorySize: jsii.Number(123),
//   				swappiness: jsii.Number(123),
//   				tmpfs: []interface{}{
//   					&tmpfsProperty{
//   						size: jsii.Number(123),
//
//   						// the properties below are optional
//   						containerPath: jsii.String("containerPath"),
//   						mountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   					},
//   				},
//   			},
//   			logConfiguration: &logConfigurationProperty{
//   				logDriver: jsii.String("logDriver"),
//
//   				// the properties below are optional
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				secretOptions: []interface{}{
//   					&secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			memory: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   			mountPoints: []interface{}{
//   				&mountPointProperty{
//   					containerPath: jsii.String("containerPath"),
//   					readOnly: jsii.Boolean(false),
//   					sourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			name: jsii.String("name"),
//   			portMappings: []interface{}{
//   				&portMappingProperty{
//   					containerPort: jsii.Number(123),
//   					hostPort: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   				},
//   			},
//   			privileged: jsii.Boolean(false),
//   			pseudoTerminal: jsii.Boolean(false),
//   			readonlyRootFilesystem: jsii.Boolean(false),
//   			repositoryCredentials: &repositoryCredentialsProperty{
//   				credentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			resourceRequirements: []interface{}{
//   				&resourceRequirementProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			secrets: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			startTimeout: jsii.Number(123),
//   			stopTimeout: jsii.Number(123),
//   			systemControls: []interface{}{
//   				&systemControlProperty{
//   					namespace: jsii.String("namespace"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			ulimits: []interface{}{
//   				&ulimitProperty{
//   					hardLimit: jsii.Number(123),
//   					name: jsii.String("name"),
//   					softLimit: jsii.Number(123),
//   				},
//   			},
//   			user: jsii.String("user"),
//   			volumesFrom: []interface{}{
//   				&volumeFromProperty{
//   					readOnly: jsii.Boolean(false),
//   					sourceContainer: jsii.String("sourceContainer"),
//   				},
//   			},
//   			workingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	cpu: jsii.String("cpu"),
//   	ephemeralStorage: &ephemeralStorageProperty{
//   		sizeInGiB: jsii.Number(123),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	family: jsii.String("family"),
//   	inferenceAccelerators: []interface{}{
//   		&inferenceAcceleratorProperty{
//   			deviceName: jsii.String("deviceName"),
//   			deviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	ipcMode: jsii.String("ipcMode"),
//   	memory: jsii.String("memory"),
//   	networkMode: jsii.String("networkMode"),
//   	pidMode: jsii.String("pidMode"),
//   	placementConstraints: []interface{}{
//   		&taskDefinitionPlacementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	proxyConfiguration: &proxyConfigurationProperty{
//   		containerName: jsii.String("containerName"),
//
//   		// the properties below are optional
//   		proxyConfigurationProperties: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		type: jsii.String("type"),
//   	},
//   	requiresCompatibilities: []*string{
//   		jsii.String("requiresCompatibilities"),
//   	},
//   	runtimePlatform: &runtimePlatformProperty{
//   		cpuArchitecture: jsii.String("cpuArchitecture"),
//   		operatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskRoleArn: jsii.String("taskRoleArn"),
//   	volumes: []interface{}{
//   		&volumeProperty{
//   			dockerVolumeConfiguration: &dockerVolumeConfigurationProperty{
//   				autoprovision: jsii.Boolean(false),
//   				driver: jsii.String("driver"),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				scope: jsii.String("scope"),
//   			},
//   			efsVolumeConfiguration: &eFSVolumeConfigurationProperty{
//   				filesystemId: jsii.String("filesystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfigProperty{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &hostVolumePropertiesProperty{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   })
//
type CfnTaskDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrTaskDefinitionArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
	ContainerDefinitions() interface{}
	SetContainerDefinitions(val interface{})
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
	// - 4096 (4 vCPU) - Available `memory` values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	Cpu() *string
	SetCpu(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ephemeral storage settings to use for tasks run with the task definition.
	EphemeralStorage() interface{}
	SetEphemeralStorage(val interface{})
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	Family() *string
	SetFamily(val *string)
	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators() interface{}
	SetInferenceAccelerators(val interface{})
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
	IpcMode() *string
	SetIpcMode(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
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
	// - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available `cpu` values: 4096 (4 vCPU).
	Memory() *string
	SetMemory(val *string)
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
	NetworkMode() *string
	SetNetworkMode(val *string)
	// The tree node.
	Node() constructs.Node
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . If `host` is specified, then all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same process namespace. If no value is specified, the default is a private namespace. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, be aware that there is a heightened risk of undesired process namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	PidMode() *string
	SetPidMode(val *string)
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	// The `ProxyConfiguration` property specifies the configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	ProxyConfiguration() interface{}
	SetProxyConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The task launch types the task definition was validated against.
	//
	// To determine which task launch types the task definition is validated for, see the `TaskDefinition$compatibilities` parameter.
	RequiresCompatibilities() *[]*string
	SetRequiresCompatibilities(val *[]*string)
	// The operating system that your tasks definitions run on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the `runtimePlatform` value of the service.
	RuntimePlatform() interface{}
	SetRuntimePlatform(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	Tags() awscdk.TagManager
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	Volumes() interface{}
	SetVolumes(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnTaskDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnTaskDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigProperty := &authorizationConfigProperty{
//   	accessPointId: jsii.String("accessPointId"),
//   	iam: jsii.String("iam"),
//   }
//
type CfnTaskDefinition_AuthorizationConfigProperty struct {
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, the root directory value specified in the `EFSVolumeConfiguration` must either be omitted or set to `/` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be enabled in the `EFSVolumeConfiguration` . For more information, see [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide* .
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Determines whether to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the `EFSVolumeConfiguration` . If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide* .
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

// The `ContainerDefinition` property specifies a container definition.
//
// Container definitions are used in task definitions to describe the different containers that are launched as part of a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDefinitionProperty := &containerDefinitionProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	cpu: jsii.Number(123),
//   	dependsOn: []interface{}{
//   		&containerDependencyProperty{
//   			condition: jsii.String("condition"),
//   			containerName: jsii.String("containerName"),
//   		},
//   	},
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	environmentFiles: []interface{}{
//   		&environmentFileProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: []interface{}{
//   		&hostEntryProperty{
//   			hostname: jsii.String("hostname"),
//   			ipAddress: jsii.String("ipAddress"),
//   		},
//   	},
//   	firelensConfiguration: &firelensConfigurationProperty{
//   		options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		type: jsii.String("type"),
//   	},
//   	healthCheck: &healthCheckProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		interval: jsii.Number(123),
//   		retries: jsii.Number(123),
//   		startPeriod: jsii.Number(123),
//   		timeout: jsii.Number(123),
//   	},
//   	hostname: jsii.String("hostname"),
//   	image: jsii.String("image"),
//   	interactive: jsii.Boolean(false),
//   	links: []*string{
//   		jsii.String("links"),
//   	},
//   	linuxParameters: &linuxParametersProperty{
//   		capabilities: &kernelCapabilitiesProperty{
//   			add: []*string{
//   				jsii.String("add"),
//   			},
//   			drop: []*string{
//   				jsii.String("drop"),
//   			},
//   		},
//   		devices: []interface{}{
//   			&deviceProperty{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//   				permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   			},
//   		},
//   		initProcessEnabled: jsii.Boolean(false),
//   		maxSwap: jsii.Number(123),
//   		sharedMemorySize: jsii.Number(123),
//   		swappiness: jsii.Number(123),
//   		tmpfs: []interface{}{
//   			&tmpfsProperty{
//   				size: jsii.Number(123),
//
//   				// the properties below are optional
//   				containerPath: jsii.String("containerPath"),
//   				mountOptions: []*string{
//   					jsii.String("mountOptions"),
//   				},
//   			},
//   		},
//   	},
//   	logConfiguration: &logConfigurationProperty{
//   		logDriver: jsii.String("logDriver"),
//
//   		// the properties below are optional
//   		options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		secretOptions: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	memory: jsii.Number(123),
//   	memoryReservation: jsii.Number(123),
//   	mountPoints: []interface{}{
//   		&mountPointProperty{
//   			containerPath: jsii.String("containerPath"),
//   			readOnly: jsii.Boolean(false),
//   			sourceVolume: jsii.String("sourceVolume"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	portMappings: []interface{}{
//   		&portMappingProperty{
//   			containerPort: jsii.Number(123),
//   			hostPort: jsii.Number(123),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	pseudoTerminal: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	repositoryCredentials: &repositoryCredentialsProperty{
//   		credentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   	resourceRequirements: []interface{}{
//   		&resourceRequirementProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	secrets: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   	startTimeout: jsii.Number(123),
//   	stopTimeout: jsii.Number(123),
//   	systemControls: []interface{}{
//   		&systemControlProperty{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ulimits: []interface{}{
//   		&ulimitProperty{
//   			hardLimit: jsii.Number(123),
//   			name: jsii.String("name"),
//   			softLimit: jsii.Number(123),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	volumesFrom: []interface{}{
//   		&volumeFromProperty{
//   			readOnly: jsii.Boolean(false),
//   			sourceContainer: jsii.String("sourceContainer"),
//   		},
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnTaskDefinition_ContainerDefinitionProperty struct {
	// The command that's passed to the container.
	//
	// This parameter maps to `Cmd` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `COMMAND` parameter to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . For more information, see [https://docs.docker.com/engine/reference/builder/#cmd](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#cmd) . If there are multiple arguments, each argument is a separated string in the array.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
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
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The dependencies defined for container startup and shutdown.
	//
	// A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.
	//
	// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent to turn on container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For tasks using the Fargate launch type, the task or service requires the following platforms:
	//
	// - Linux platform version `1.3.0` or later.
	// - Windows platform version `1.0.0` or later.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// When this parameter is true, networking is disabled within the container.
	//
	// This parameter maps to `NetworkDisabled` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) .
	//
	// > This parameter is not supported for Windows containers.
	DisableNetworking interface{} `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	//
	// This parameter maps to `DnsSearch` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--dns-search` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	//
	// This parameter maps to `Dns` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--dns` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	//
	// This parameter maps to `Labels` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--label` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	DockerLabels interface{} `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
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
	// Valid values: "no-new-privileges" | "apparmor:PROFILE" | "label:value" | "credentialspec:CredentialSpecFilePath".
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// > Early versions of the Amazon ECS container agent don't properly handle `entryPoint` parameters.
	//
	// If you have problems using `entryPoint` , update your container agent or enter your commands and arguments as `command` array items instead.
	//
	// The entry point that's passed to the container. This parameter maps to `Entrypoint` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--entrypoint` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . For more information, see [https://docs.docker.com/engine/reference/builder/#entrypoint](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#entrypoint) .
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to a container.
	//
	// This parameter maps to `Env` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--env` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > We don't recommend that you use plaintext environment variables for sensitive information, such as credential data.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A list of files containing the environment variables to pass to a container.
	//
	// This parameter maps to the `--env-file` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file contains an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .
	//
	// If there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying Environment Variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// If the `essential` parameter of a container is marked as `true` , and that container fails or stops for any reason, all other containers that are part of the task are stopped.
	//
	// If the `essential` parameter of a container is marked as `false` , its failure doesn't affect the rest of the containers in a task. If this parameter is omitted, a container is assumed to be essential.
	//
	// All tasks must have at least one essential container. If you have an application that's composed of multiple containers, group containers that are used for a common purpose into components, and separate the different components into multiple task definitions. For more information, see [Application Architecture](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/application_architecture.html) in the *Amazon Elastic Container Service Developer Guide* .
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the `/etc/hosts` file on the container.
	//
	// This parameter maps to `ExtraHosts` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--add-host` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter isn't supported for Windows containers or tasks that use the `awsvpc` network mode.
	ExtraHosts interface{} `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The FireLens configuration for the container.
	//
	// This is used to specify and configure a log router for container logs. For more information, see [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
	FirelensConfiguration interface{} `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// The container health check command and associated configuration parameters for the container.
	//
	// This parameter maps to `HealthCheck` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `HEALTHCHECK` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	//
	// This parameter maps to `Hostname` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--hostname` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > The `hostname` parameter is not supported if you're using the `awsvpc` network mode.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon. By default, images in the Docker Hub registry are available. Other repositories are specified with either `*repository-url* / *image* : *tag*` or `*repository-url* / *image* @ *digest*` . Up to 255 letters (uppercase and lowercase), numbers, hyphens, underscores, colons, periods, forward slashes, and number signs are allowed. This parameter maps to `Image` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `IMAGE` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// - When a new task starts, the Amazon ECS container agent pulls the latest version of the specified image and tag for the container to use. However, subsequent updates to a repository image aren't propagated to already running tasks.
	// - Images in Amazon ECR repositories can be specified by either using the full `registry/repository:tag` or `registry/repository@digest` . For example, `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest` or `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE` .
	// - Images in official repositories on Docker Hub use a single name (for example, `ubuntu` or `mongo` ).
	// - Images in other repositories on Docker Hub are qualified with an organization name (for example, `amazon/amazon-ecs-agent` ).
	// - Images in other online repositories are qualified further by a domain name (for example, `quay.io/assemblyline/ubuntu` ).
	Image *string `field:"optional" json:"image" yaml:"image"`
	// When this parameter is `true` , you can deploy containerized applications that require `stdin` or a `tty` to be allocated.
	//
	// This parameter maps to `OpenStdin` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--interactive` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// The `links` parameter allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the network mode of a task definition is `bridge` . The `name:internalName` construct is analogous to `name:alias` in Docker links. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. For more information about linking Docker containers, go to [Legacy container links](https://docs.aws.amazon.com/https://docs.docker.com/network/links/) in the Docker documentation. This parameter maps to `Links` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--link` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers. > Containers that are collocated on a single container instance may be able to communicate with each other without requiring links or host port mappings. Network isolation is achieved on the container instance using security groups and VPC settings.
	Links *[]*string `field:"optional" json:"links" yaml:"links"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities. For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html) .
	//
	// > This parameter is not supported for Windows containers.
	LinuxParameters interface{} `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	//
	// This parameter maps to `LogConfig` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--log-driver` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . By default, containers use the same logging driver that the Docker daemon uses. However, the container may use a different logging driver than the Docker daemon by specifying a log driver with this parameter in the container definition. To use a different logging driver for a container, the log system must be configured properly on the container instance (or on a different log server for remote logging options). For more information on the options for different supported log drivers, see [Configure logging drivers](https://docs.aws.amazon.com/https://docs.docker.com/engine/admin/logging/overview/) in the Docker documentation.
	//
	// > Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon (shown in the [LogConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_LogConfiguration.html) data type). Additional log drivers may be available in future releases of the Amazon ECS container agent.
	//
	// This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	//
	// > The Amazon ECS container agent running on a container instance must register the logging drivers available on that instance with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide* .
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
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
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the container memory to this soft limit. However, your container can consume more memory when it needs to, up to either the hard limit specified with the `memory` parameter (if applicable), or all of the available memory on the container instance, whichever comes first. This parameter maps to `MemoryReservation` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--memory-reservation` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// If a task-level memory value is not specified, you must specify a non-zero integer for one or both of `memory` or `memoryReservation` in a container definition. If you specify both, `memory` must be greater than `memoryReservation` . If you specify `memoryReservation` , then that value is subtracted from the available memory resources for the container instance where the container is placed. Otherwise, the value of `memory` is used.
	//
	// For example, if your container normally uses 128 MiB of memory, but occasionally bursts to 256 MiB of memory for short periods of time, you can set a `memoryReservation` of 128 MiB, and a `memory` hard limit of 300 MiB. This configuration would allow the container to only reserve 128 MiB of memory from the remaining resources on the container instance, but also allow the container to consume more memory resources when needed.
	//
	// The Docker daemon reserves a minimum of 4 MiB of memory for a container. Therefore, we recommend that you specify fewer than 4 MiB of memory for your containers.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The mount points for data volumes in your container.
	//
	// This parameter maps to `Volumes` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--volume` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// Windows containers can mount whole directories on the same drive as `$env:ProgramData` . Windows containers can't mount directories on a different drive, and mount point can't be across drives.
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// The name of a container.
	//
	// If you're linking multiple containers together in a task definition, the `name` of one container can be entered in the `links` of another container to connect the containers. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This parameter maps to `name` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--name` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	Name *string `field:"optional" json:"name" yaml:"name"`
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
	PortMappings interface{} `field:"optional" json:"portMappings" yaml:"portMappings"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the `root` user).
	//
	// This parameter maps to `Privileged` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--privileged` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is `true` , a TTY is allocated.
	//
	// This parameter maps to `Tty` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--tty` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// When this parameter is true, the container is given read-only access to its root file system.
	//
	// This parameter maps to `ReadonlyRootfs` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--read-only` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > This parameter is not supported for Windows containers.
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The private repository authentication credentials to use.
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// The type and amount of a resource to assign to a container.
	//
	// The only supported resource is a GPU.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// The secrets to pass to the container.
	//
	// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
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
	// For tasks using the EC2 launch type, your container instances require at least version `1.26.0` of the container agent to use a container start timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version `1.26.0-1` of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	//
	// For tasks using the Fargate launch type, the task or service requires the following platforms:
	//
	// - Linux platform version `1.3.0` or later.
	// - Windows platform version `1.0.0` or later.
	//
	// The max stop timeout value is 120 seconds and if the parameter is not specified, the default value of 30 seconds is used.
	//
	// For tasks that use the EC2 launch type, if the `stopTimeout` parameter isn't specified, the value set for the Amazon ECS container agent configuration variable `ECS_CONTAINER_STOP_TIMEOUT` is used. If neither the `stopTimeout` parameter or the `ECS_CONTAINER_STOP_TIMEOUT` agent configuration variable are set, then the default values of 30 seconds for Linux containers and 30 seconds on Windows containers are used. Your container instances require at least version 1.26.0 of the container agent to use a container stop timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	//
	// This parameter maps to `Sysctls` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--sysctl` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > We don't recommended that you specify network-related `systemControls` parameters for multiple containers in a single task that also uses either the `awsvpc` or `host` network modes. For tasks that use the `awsvpc` network mode, the container that's started last determines which `systemControls` parameters take effect. For tasks that use the `host` network mode, it changes the container instance's namespaced kernel parameters as well as the containers.
	SystemControls interface{} `field:"optional" json:"systemControls" yaml:"systemControls"`
	// A list of `ulimits` to set in the container.
	//
	// This parameter maps to `Ulimits` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--ulimit` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . Valid naming values are displayed in the [Ulimit](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Ulimit.html) data type. This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	//
	// > This parameter is not supported for Windows containers.
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
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
	User *string `field:"optional" json:"user" yaml:"user"`
	// Data volumes to mount from another container.
	//
	// This parameter maps to `VolumesFrom` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--volumes-from` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	VolumesFrom interface{} `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// The working directory to run commands inside the container in.
	//
	// This parameter maps to `WorkingDir` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--workdir` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// The `ContainerDependency` property specifies the dependencies defined for container startup and shutdown.
//
// A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.
//
// Your Amazon ECS container instances require at least version 1.26.0 of the container agent to enable container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you are using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > For tasks using the Fargate launch type, this parameter requires that the task or service uses platform version 1.3.0 or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDependencyProperty := &containerDependencyProperty{
//   	condition: jsii.String("condition"),
//   	containerName: jsii.String("containerName"),
//   }
//
type CfnTaskDefinition_ContainerDependencyProperty struct {
	// The dependency condition of the container. The following are the available conditions and their behavior:.
	//
	// - `START` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.
	// - `COMPLETE` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.
	// - `SUCCESS` - This condition is the same as `COMPLETE` , but it also requires that the container exits with a `zero` status. This condition can't be set on an essential container.
	// - `HEALTHY` - This condition validates that the dependent container passes its Docker health check before permitting other containers to start. This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The name of a container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

// The `Device` property specifies an object representing a container instance host device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	containerPath: jsii.String("containerPath"),
//   	hostPath: jsii.String("hostPath"),
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   }
//
type CfnTaskDefinition_DeviceProperty struct {
	// The path inside the container at which to expose the host device.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

// The `DockerVolumeConfiguration` property specifies a Docker volume configuration and is used when you use Docker volumes.
//
// Docker volumes are only supported when you are using the EC2 launch type. Windows containers only support the use of the `local` driver. To use bind mounts, specify a `host` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerVolumeConfigurationProperty := &dockerVolumeConfigurationProperty{
//   	autoprovision: jsii.Boolean(false),
//   	driver: jsii.String("driver"),
//   	driverOpts: map[string]*string{
//   		"driverOptsKey": jsii.String("driverOpts"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	scope: jsii.String("scope"),
//   }
//
type CfnTaskDefinition_DockerVolumeConfigurationProperty struct {
	// If this value is `true` , the Docker volume is created if it doesn't already exist.
	//
	// > This field is only used if the `scope` is `shared` .
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// The Docker volume driver to use.
	//
	// The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use `docker plugin ls` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. For more information, see [Docker plugin discovery](https://docs.aws.amazon.com/https://docs.docker.com/engine/extend/plugin_api/#plugin-discovery) . This parameter maps to `Driver` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxdriver` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// A map of Docker driver-specific options passed through.
	//
	// This parameter maps to `DriverOpts` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxopt` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	DriverOpts interface{} `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	//
	// This parameter maps to `Labels` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxlabel` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The scope for the Docker volume that determines its lifecycle.
	//
	// Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

// This parameter is specified when you're using an Amazon Elastic File System file system for task storage.
//
// For more information, see [Amazon EFS volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eFSVolumeConfigurationProperty := &eFSVolumeConfigurationProperty{
//   	filesystemId: jsii.String("filesystemId"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfigProperty{
//   		accessPointId: jsii.String("accessPointId"),
//   		iam: jsii.String("iam"),
//   	},
//   	rootDirectory: jsii.String("rootDirectory"),
//   	transitEncryption: jsii.String("transitEncryption"),
//   	transitEncryptionPort: jsii.Number(123),
//   }
//
type CfnTaskDefinition_EFSVolumeConfigurationProperty struct {
	// The Amazon EFS file system ID to use.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying `/` will have the same effect as omitting this parameter.
	//
	// > If an EFS access point is specified in the `authorizationConfig` , the root directory parameter must either be omitted or set to `/` which will enforce the path set on the EFS access point.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Determines whether to use encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be enabled if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide* .
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses. For more information, see [EFS mount helper](https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html) in the *Amazon Elastic File System User Guide* .
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileProperty := &environmentFileProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_EnvironmentFileProperty struct {
	// The file type to use.
	//
	// The only supported value is `s3` .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The amount of ephemeral storage to allocate for the task.
//
// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate . For more information, see [Fargate task storage](https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html) in the *Amazon ECS User Guide for AWS Fargate* .
//
// > This parameter is only supported for tasks hosted on Fargate using Linux platform version `1.4.0` or later. This parameter is not supported for Windows containers on Fargate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &ephemeralStorageProperty{
//   	sizeInGiB: jsii.Number(123),
//   }
//
type CfnTaskDefinition_EphemeralStorageProperty struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.
	SizeInGiB *float64 `field:"optional" json:"sizeInGiB" yaml:"sizeInGiB"`
}

// The FireLens configuration for the container.
//
// This is used to specify and configure a log router for container logs. For more information, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfigurationProperty := &firelensConfigurationProperty{
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	type: jsii.String("type"),
//   }
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
	// - `config-file-value` , which is either an S3 ARN or a file path.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The log router to use.
	//
	// The valid values are `fluentd` or `fluentbit` .
	Type *string `field:"optional" json:"type" yaml:"type"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &healthCheckProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	interval: jsii.Number(123),
//   	retries: jsii.Number(123),
//   	startPeriod: jsii.Number(123),
//   	timeout: jsii.Number(123),
//   }
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
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds. The default value is 30 seconds.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries. The default value is 3.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You can specify between 0 and 300 seconds. By default, the `startPeriod` is disabled.
	//
	// > If a health check succeeds within the `startPeriod` , then the container is considered healthy and any subsequent failures count toward the maximum number of retries.
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds. The default value is 5.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

// The `HostEntry` property specifies a hostname and an IP address that are added to the `/etc/hosts` file of a container through the `extraHosts` parameter of its `ContainerDefinition` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostEntryProperty := &hostEntryProperty{
//   	hostname: jsii.String("hostname"),
//   	ipAddress: jsii.String("ipAddress"),
//   }
//
type CfnTaskDefinition_HostEntryProperty struct {
	// The hostname to use in the `/etc/hosts` entry.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The IP address to use in the `/etc/hosts` entry.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

// The `HostVolumeProperties` property specifies details on a container instance bind mount host volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostVolumePropertiesProperty := &hostVolumePropertiesProperty{
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type CfnTaskDefinition_HostVolumePropertiesProperty struct {
	// When the `host` parameter is used, specify a `sourcePath` to declare the path on the host container instance that's presented to the container.
	//
	// If this parameter is empty, then the Docker daemon has assigned a host path for you. If the `host` parameter contains a `sourcePath` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the `sourcePath` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
	//
	// If you're using the Fargate launch type, the `sourcePath` parameter is not supported.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

// Details on an Elastic Inference accelerator.
//
// For more information, see [Working with Amazon Elastic Inference on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-eia.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceAcceleratorProperty := &inferenceAcceleratorProperty{
//   	deviceName: jsii.String("deviceName"),
//   	deviceType: jsii.String("deviceType"),
//   }
//
type CfnTaskDefinition_InferenceAcceleratorProperty struct {
	// The Elastic Inference accelerator device name.
	//
	// The `deviceName` must also be referenced in a container definition as a [ResourceRequirement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-resourcerequirement.html) .
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

// The `KernelCapabilities` property specifies the Linux capabilities for the container that are added to or dropped from the default configuration that is provided by Docker.
//
// For more information on the default capabilities and the non-default available capabilities, see [Runtime privilege and Linux capabilities](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities) in the *Docker run reference* . For more detailed information on these Linux capabilities, see the [capabilities(7)](https://docs.aws.amazon.com/http://man7.org/linux/man-pages/man7/capabilities.7.html) Linux manual page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelCapabilitiesProperty := &kernelCapabilitiesProperty{
//   	add: []*string{
//   		jsii.String("add"),
//   	},
//   	drop: []*string{
//   		jsii.String("drop"),
//   	},
//   }
//
type CfnTaskDefinition_KernelCapabilitiesProperty struct {
	// The Linux capabilities for the container that have been added to the default configuration provided by Docker.
	//
	// This parameter maps to `CapAdd` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--cap-add` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > Tasks launched on AWS Fargate only support adding the `SYS_PTRACE` kernel capability.
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// The Linux capabilities for the container that have been removed from the default configuration provided by Docker.
	//
	// This parameter maps to `CapDrop` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--cap-drop` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

// The `KeyValuePair` property specifies a key-value pair object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValuePairProperty := &keyValuePairProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_KeyValuePairProperty struct {
	// The name of the key-value pair.
	//
	// For environment variables, this is the name of the environment variable.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair.
	//
	// For environment variables, this is the value of the environment variable.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `LinuxParameters` property specifies Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParametersProperty := &linuxParametersProperty{
//   	capabilities: &kernelCapabilitiesProperty{
//   		add: []*string{
//   			jsii.String("add"),
//   		},
//   		drop: []*string{
//   			jsii.String("drop"),
//   		},
//   	},
//   	devices: []interface{}{
//   		&deviceProperty{
//   			containerPath: jsii.String("containerPath"),
//   			hostPath: jsii.String("hostPath"),
//   			permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   		},
//   	},
//   	initProcessEnabled: jsii.Boolean(false),
//   	maxSwap: jsii.Number(123),
//   	sharedMemorySize: jsii.Number(123),
//   	swappiness: jsii.Number(123),
//   	tmpfs: []interface{}{
//   		&tmpfsProperty{
//   			size: jsii.Number(123),
//
//   			// the properties below are optional
//   			containerPath: jsii.String("containerPath"),
//   			mountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   		},
//   	},
//   }
//
type CfnTaskDefinition_LinuxParametersProperty struct {
	// The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.
	//
	// > For tasks that use the Fargate launch type, `capabilities` is supported for all platform versions but the `add` parameter is only supported if using platform version 1.4.0 or later.
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to `Devices` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--device` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `devices` parameter isn't supported.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Run an `init` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the `--init` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory (in MiB) a container can use.
	//
	// This parameter will be translated to the `--memory-swap` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) where the value would be the sum of the container memory plus the `maxSwap` value.
	//
	// If a `maxSwap` value of `0` is specified, the container will not use swap. Accepted values are `0` or any positive integer. If the `maxSwap` parameter is omitted, the container will use the swap configuration for the container instance it is running on. A `maxSwap` value must be set for the `swappiness` parameter to be used.
	//
	// > If you're using tasks that use the Fargate launch type, the `maxSwap` parameter isn't supported.
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size (in MiB) of the `/dev/shm` volume.
	//
	// This parameter maps to the `--shm-size` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you are using tasks that use the Fargate launch type, the `sharedMemorySize` parameter is not supported.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// A `swappiness` value of `0` will cause swapping to not happen unless absolutely necessary. A `swappiness` value of `100` will cause pages to be swapped very aggressively. Accepted values are whole numbers between `0` and `100` . If the `swappiness` parameter is not specified, a default value of `60` is used. If a value is not specified for `maxSwap` then this parameter is ignored. This parameter maps to the `--memory-swappiness` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `swappiness` parameter isn't supported.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the `--tmpfs` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `tmpfs` parameter isn't supported.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

// The `LogConfiguration` property specifies log configuration options to send to a custom log driver for the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
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
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	//
	// This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	//
	// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

// The `MountPoint` property specifies details on a volume mount point that is used in a container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPointProperty := &mountPointProperty{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type CfnTaskDefinition_MountPointProperty struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the `name` parameter of task definition `volume` .
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

// The `PortMapping` property specifies a port mapping.
//
// Port mappings allow containers to access ports on the host container instance to send or receive traffic. Port mappings are specified as part of the container definition.
//
// If you are using containers in a task with the `awsvpc` or `host` network mode, exposed ports should be specified using `containerPort` . The `hostPort` can be left blank or it must be the same value as the `containerPort` .
//
// After a task reaches the `RUNNING` status, manual and automatic host and container port assignments are visible in the `networkBindings` section of [DescribeTasks](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTasks.html) API responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portMappingProperty := &portMappingProperty{
//   	containerPort: jsii.Number(123),
//   	hostPort: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnTaskDefinition_PortMappingProperty struct {
	// The port number on the container that's bound to the user-specified or automatically assigned host port.
	//
	// If you use containers in a task with the `awsvpc` or `host` network mode, specify the exposed ports using `containerPort` .
	//
	// If you use containers in a task with the `bridge` network mode and you specify a container port and not a host port, your container automatically receives a host port in the ephemeral port range. For more information, see `hostPort` . Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
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
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The protocol used for the port mapping.
	//
	// Valid values are `tcp` and `udp` . The default is `tcp` .
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

// The `ProxyConfiguration` property specifies the details for the App Mesh proxy.
//
// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// For tasks using the Fargate launch type, the task or service requires platform version 1.3.0 or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurationProperty := &proxyConfigurationProperty{
//   	containerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	proxyConfigurationProperties: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnTaskDefinition_ProxyConfigurationProperty struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.
	//
	// - `IgnoredUID` - (Required) The user ID (UID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredGID` is specified, this field can be empty.
	// - `IgnoredGID` - (Required) The group ID (GID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredUID` is specified, this field can be empty.
	// - `AppPorts` - (Required) The list of ports that the application uses. Network traffic to these ports is forwarded to the `ProxyIngressPort` and `ProxyEgressPort` .
	// - `ProxyIngressPort` - (Required) Specifies the port that incoming traffic to the `AppPorts` is directed to.
	// - `ProxyEgressPort` - (Required) Specifies the port that outgoing traffic from the `AppPorts` is directed to.
	// - `EgressIgnoredPorts` - (Required) The egress traffic going to the specified ports is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	// - `EgressIgnoredIPs` - (Required) The egress traffic going to the specified IP addresses is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	ProxyConfigurationProperties interface{} `field:"optional" json:"proxyConfigurationProperties" yaml:"proxyConfigurationProperties"`
	// The proxy type.
	//
	// The only supported value is `APPMESH` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The `RepositoryCredentials` property specifies the repository credentials for private registry authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCredentialsProperty := &repositoryCredentialsProperty{
//   	credentialsParameter: jsii.String("credentialsParameter"),
//   }
//
type CfnTaskDefinition_RepositoryCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	//
	// > When you use the Amazon ECS API, AWS CLI , or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}

// The `ResourceRequirement` property specifies the type and amount of a resource to assign to a container.
//
// The only supported resource is a GPU. For more information, see [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html) in the *Amazon Elastic Container Service Developer Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceRequirementProperty := &resourceRequirementProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_ResourceRequirementProperty struct {
	// The type of resource to assign to a container.
	//
	// The supported values are `GPU` or `InferenceAccelerator` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The value for the specified resource type.
	//
	// If the `GPU` type is used, the value is the number of physical `GPUs` the Amazon ECS container agent will reserve for the container. The number of GPUs reserved for all containers in a task should not exceed the number of available GPUs on the container instance the task is launched on.
	//
	// If the `InferenceAccelerator` type is used, the `value` should match the `DeviceName` for an [InferenceAccelerator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html) specified in a task definition.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Information about the platform for the Amazon ECS service or task.
//
// For more informataion about `RuntimePlatform` , see [RuntimePlatform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimePlatformProperty := &runtimePlatformProperty{
//   	cpuArchitecture: jsii.String("cpuArchitecture"),
//   	operatingSystemFamily: jsii.String("operatingSystemFamily"),
//   }
//
type CfnTaskDefinition_RuntimePlatformProperty struct {
	// The CPU architecture.
	//
	// You can run your Linux tasks on an ARM-based platform by setting the value to `ARM64` . This option is avaiable for tasks that run on Linuc Amazon EC2 instance or Linux containers on Fargate.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

// The `Secret` property specifies an object representing the secret to expose to your container.
//
// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretProperty := &secretProperty{
//   	name: jsii.String("name"),
//   	valueFrom: jsii.String("valueFrom"),
//   }
//
type CfnTaskDefinition_SecretProperty struct {
	// The name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The secret to expose to the container.
	//
	// The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter Store.
	//
	// For information about the require AWS Identity and Access Management permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemControlProperty := &systemControlProperty{
//   	namespace: jsii.String("namespace"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_SystemControlProperty struct {
	// The namespaced kernel parameter to set a `value` for.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The value for the namespaced kernel parameter that's specified in `namespace` .
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `TaskDefinitionPlacementConstraint` property specifies an object representing a constraint on task placement in the task definition.
//
// If you are using the Fargate launch type, task placement constraints are not supported.
//
// For more information, see [Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskDefinitionPlacementConstraintProperty := &taskDefinitionPlacementConstraintProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	expression: jsii.String("expression"),
//   }
//
type CfnTaskDefinition_TaskDefinitionPlacementConstraintProperty struct {
	// The type of constraint.
	//
	// The `MemberOf` constraint restricts selection to be from a group of valid candidates.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A cluster query language expression to apply to the constraint.
	//
	// For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

// The `Tmpfs` property specifies the container path, mount options, and size of the tmpfs mount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfsProperty := &tmpfsProperty{
//   	size: jsii.Number(123),
//
//   	// the properties below are optional
//   	containerPath: jsii.String("containerPath"),
//   	mountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   }
//
type CfnTaskDefinition_TmpfsProperty struct {
	// The maximum size (in MiB) of the tmpfs volume.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The list of tmpfs volume mount options.
	//
	// Valid values: `"defaults" | "ro" | "rw" | "suid" | "nosuid" | "dev" | "nodev" | "exec" | "noexec" | "sync" | "async" | "dirsync" | "remount" | "mand" | "nomand" | "atime" | "noatime" | "diratime" | "nodiratime" | "bind" | "rbind" | "unbindable" | "runbindable" | "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime" | "norelatime" | "strictatime" | "nostrictatime" | "mode" | "uid" | "gid" | "nr_inodes" | "nr_blocks" | "mpol"`.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

// The `Ulimit` property specifies the `ulimit` settings to pass to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimitProperty := &ulimitProperty{
//   	hardLimit: jsii.Number(123),
//   	name: jsii.String("name"),
//   	softLimit: jsii.Number(123),
//   }
//
type CfnTaskDefinition_UlimitProperty struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The `type` of the `ulimit` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

// The `VolumeFrom` property specifies details on a data volume from another container in the same task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeFromProperty := &volumeFromProperty{
//   	readOnly: jsii.Boolean(false),
//   	sourceContainer: jsii.String("sourceContainer"),
//   }
//
type CfnTaskDefinition_VolumeFromProperty struct {
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition to mount volumes from.
	SourceContainer *string `field:"optional" json:"sourceContainer" yaml:"sourceContainer"`
}

// The `Volume` property specifies a data volume used in a task definition.
//
// For tasks that use a Docker volume, specify a `DockerVolumeConfiguration` . For tasks that use a bind mount host volume, specify a `host` and optional `sourcePath` . For more information, see [Using Data Volumes in Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeProperty := &volumeProperty{
//   	dockerVolumeConfiguration: &dockerVolumeConfigurationProperty{
//   		autoprovision: jsii.Boolean(false),
//   		driver: jsii.String("driver"),
//   		driverOpts: map[string]*string{
//   			"driverOptsKey": jsii.String("driverOpts"),
//   		},
//   		labels: map[string]*string{
//   			"labelsKey": jsii.String("labels"),
//   		},
//   		scope: jsii.String("scope"),
//   	},
//   	efsVolumeConfiguration: &eFSVolumeConfigurationProperty{
//   		filesystemId: jsii.String("filesystemId"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			accessPointId: jsii.String("accessPointId"),
//   			iam: jsii.String("iam"),
//   		},
//   		rootDirectory: jsii.String("rootDirectory"),
//   		transitEncryption: jsii.String("transitEncryption"),
//   		transitEncryptionPort: jsii.Number(123),
//   	},
//   	host: &hostVolumePropertiesProperty{
//   		sourcePath: jsii.String("sourcePath"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnTaskDefinition_VolumeProperty struct {
	// This parameter is specified when you use Docker volumes.
	//
	// Windows containers only support the use of the `local` driver. To use bind mounts, specify the `host` parameter instead.
	//
	// > Docker volumes aren't supported by tasks run on AWS Fargate .
	DockerVolumeConfiguration interface{} `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// This parameter is specified when you use an Amazon Elastic File System file system for task storage.
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// This parameter is specified when you use bind mount host volumes.
	//
	// The contents of the `host` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the `host` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//
	// Windows containers can mount whole directories on the same drive as `$env:ProgramData` . Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount `C:\my\path:C:\my\path` and `D:\:D:\` , but not `D:\my\path:C:\my\path` or `D:\:C:\my\path` .
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnTaskDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinitionProps := &cfnTaskDefinitionProps{
//   	containerDefinitions: []interface{}{
//   		&containerDefinitionProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			dependsOn: []interface{}{
//   				&containerDependencyProperty{
//   					condition: jsii.String("condition"),
//   					containerName: jsii.String("containerName"),
//   				},
//   			},
//   			disableNetworking: jsii.Boolean(false),
//   			dnsSearchDomains: []*string{
//   				jsii.String("dnsSearchDomains"),
//   			},
//   			dnsServers: []*string{
//   				jsii.String("dnsServers"),
//   			},
//   			dockerLabels: map[string]*string{
//   				"dockerLabelsKey": jsii.String("dockerLabels"),
//   			},
//   			dockerSecurityOptions: []*string{
//   				jsii.String("dockerSecurityOptions"),
//   			},
//   			entryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			environment: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			environmentFiles: []interface{}{
//   				&environmentFileProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			essential: jsii.Boolean(false),
//   			extraHosts: []interface{}{
//   				&hostEntryProperty{
//   					hostname: jsii.String("hostname"),
//   					ipAddress: jsii.String("ipAddress"),
//   				},
//   			},
//   			firelensConfiguration: &firelensConfigurationProperty{
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				type: jsii.String("type"),
//   			},
//   			healthCheck: &healthCheckProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				interval: jsii.Number(123),
//   				retries: jsii.Number(123),
//   				startPeriod: jsii.Number(123),
//   				timeout: jsii.Number(123),
//   			},
//   			hostname: jsii.String("hostname"),
//   			image: jsii.String("image"),
//   			interactive: jsii.Boolean(false),
//   			links: []*string{
//   				jsii.String("links"),
//   			},
//   			linuxParameters: &linuxParametersProperty{
//   				capabilities: &kernelCapabilitiesProperty{
//   					add: []*string{
//   						jsii.String("add"),
//   					},
//   					drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				devices: []interface{}{
//   					&deviceProperty{
//   						containerPath: jsii.String("containerPath"),
//   						hostPath: jsii.String("hostPath"),
//   						permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				initProcessEnabled: jsii.Boolean(false),
//   				maxSwap: jsii.Number(123),
//   				sharedMemorySize: jsii.Number(123),
//   				swappiness: jsii.Number(123),
//   				tmpfs: []interface{}{
//   					&tmpfsProperty{
//   						size: jsii.Number(123),
//
//   						// the properties below are optional
//   						containerPath: jsii.String("containerPath"),
//   						mountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   					},
//   				},
//   			},
//   			logConfiguration: &logConfigurationProperty{
//   				logDriver: jsii.String("logDriver"),
//
//   				// the properties below are optional
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				secretOptions: []interface{}{
//   					&secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			memory: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   			mountPoints: []interface{}{
//   				&mountPointProperty{
//   					containerPath: jsii.String("containerPath"),
//   					readOnly: jsii.Boolean(false),
//   					sourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			name: jsii.String("name"),
//   			portMappings: []interface{}{
//   				&portMappingProperty{
//   					containerPort: jsii.Number(123),
//   					hostPort: jsii.Number(123),
//   					protocol: jsii.String("protocol"),
//   				},
//   			},
//   			privileged: jsii.Boolean(false),
//   			pseudoTerminal: jsii.Boolean(false),
//   			readonlyRootFilesystem: jsii.Boolean(false),
//   			repositoryCredentials: &repositoryCredentialsProperty{
//   				credentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			resourceRequirements: []interface{}{
//   				&resourceRequirementProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			secrets: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			startTimeout: jsii.Number(123),
//   			stopTimeout: jsii.Number(123),
//   			systemControls: []interface{}{
//   				&systemControlProperty{
//   					namespace: jsii.String("namespace"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			ulimits: []interface{}{
//   				&ulimitProperty{
//   					hardLimit: jsii.Number(123),
//   					name: jsii.String("name"),
//   					softLimit: jsii.Number(123),
//   				},
//   			},
//   			user: jsii.String("user"),
//   			volumesFrom: []interface{}{
//   				&volumeFromProperty{
//   					readOnly: jsii.Boolean(false),
//   					sourceContainer: jsii.String("sourceContainer"),
//   				},
//   			},
//   			workingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	cpu: jsii.String("cpu"),
//   	ephemeralStorage: &ephemeralStorageProperty{
//   		sizeInGiB: jsii.Number(123),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	family: jsii.String("family"),
//   	inferenceAccelerators: []interface{}{
//   		&inferenceAcceleratorProperty{
//   			deviceName: jsii.String("deviceName"),
//   			deviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	ipcMode: jsii.String("ipcMode"),
//   	memory: jsii.String("memory"),
//   	networkMode: jsii.String("networkMode"),
//   	pidMode: jsii.String("pidMode"),
//   	placementConstraints: []interface{}{
//   		&taskDefinitionPlacementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	proxyConfiguration: &proxyConfigurationProperty{
//   		containerName: jsii.String("containerName"),
//
//   		// the properties below are optional
//   		proxyConfigurationProperties: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		type: jsii.String("type"),
//   	},
//   	requiresCompatibilities: []*string{
//   		jsii.String("requiresCompatibilities"),
//   	},
//   	runtimePlatform: &runtimePlatformProperty{
//   		cpuArchitecture: jsii.String("cpuArchitecture"),
//   		operatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskRoleArn: jsii.String("taskRoleArn"),
//   	volumes: []interface{}{
//   		&volumeProperty{
//   			dockerVolumeConfiguration: &dockerVolumeConfigurationProperty{
//   				autoprovision: jsii.Boolean(false),
//   				driver: jsii.String("driver"),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				scope: jsii.String("scope"),
//   			},
//   			efsVolumeConfiguration: &eFSVolumeConfigurationProperty{
//   				filesystemId: jsii.String("filesystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfigProperty{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &hostVolumePropertiesProperty{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTaskDefinitionProps struct {
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
	ContainerDefinitions interface{} `field:"optional" json:"containerDefinitions" yaml:"containerDefinitions"`
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
	// - 4096 (4 vCPU) - Available `memory` values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ephemeral storage settings to use for tasks run with the task definition.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators interface{} `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
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
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
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
	// - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available `cpu` values: 4096 (4 vCPU).
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
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
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . If `host` is specified, then all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same process namespace. If no value is specified, the default is a private namespace. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, be aware that there is a heightened risk of undesired process namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The `ProxyConfiguration` property specifies the configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The task launch types the task definition was validated against.
	//
	// To determine which task launch types the task definition is validated for, see the `TaskDefinition$compatibilities` parameter.
	RequiresCompatibilities *[]*string `field:"optional" json:"requiresCompatibilities" yaml:"requiresCompatibilities"`
	// The operating system that your tasks definitions run on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the `runtimePlatform` value of the service.
	RuntimePlatform interface{} `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
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
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

// A CloudFormation `AWS::ECS::TaskSet`.
//
// Create a task set in the specified cluster and service. This is used when a service uses the `EXTERNAL` deployment controller type. For more information, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskSet := awscdk.Aws_ecs.NewCfnTaskSet(this, jsii.String("MyCfnTaskSet"), &cfnTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskDefinition: jsii.String("taskDefinition"),
//
//   	// the properties below are optional
//   	externalId: jsii.String("externalId"),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsVpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	scale: &scaleProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   })
//
type CfnTaskSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the task set.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	Cluster() *string
	SetCluster(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map attribute set to the provided value.
	ExternalId() *string
	SetExternalId(val *string)
	// The launch type that new tasks in the task set uses.
	//
	// For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If a `launchType` is specified, the `capacityProviderStrategy` parameter must be omitted.
	LaunchType() *string
	SetLaunchType(val *string)
	// A load balancer object representing the load balancer to use with the task set.
	//
	// The supported load balancer types are either an Application Load Balancer or a Network Load Balancer.
	LoadBalancers() interface{}
	SetLoadBalancers(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The network configuration for the task set.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// The platform version that the tasks in the task set uses.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used.
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A floating-point percentage of your desired number of tasks to place and keep running in the task set.
	Scale() interface{}
	SetScale(val interface{})
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	Service() *string
	SetService(val *string)
	// The details of the service discovery registries to assign to this task set.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	ServiceRegistries() interface{}
	SetServiceRegistries(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The task definition for the tasks in the task set to use.
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnTaskSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTaskSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTaskSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTaskSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTaskSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTaskSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTaskSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnTaskSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsVpcConfigurationProperty := &awsVpcConfigurationProperty{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnTaskSet_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

// Details on the load balancer or load balancers to use with a task set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerProperty := &loadBalancerProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnTaskSet_LoadBalancerProperty struct {
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you're using a Classic Load Balancer, omit the target group ARN.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

// The network configuration for a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	awsVpcConfiguration: &awsVpcConfigurationProperty{
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//
//   		// the properties below are optional
//   		assignPublicIp: jsii.String("assignPublicIp"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   }
//
type CfnTaskSet_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	AwsVpcConfiguration interface{} `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleProperty := &scaleProperty{
//   	unit: jsii.String("unit"),
//   	value: jsii.Number(123),
//   }
//
type CfnTaskSet_ScaleProperty struct {
	// The unit of measure for the scale value.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value, specified as a percent total of a service's `desiredCount` , to scale the task set.
	//
	// Accepted values are numbers between 0 and 100.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// The details for the service registry.
//
// Each service may be associated with one service registry. Multiple service registries for each service are not supported.
//
// When you add, update, or remove the service registries configuration, Amazon ECS starts a new deployment. New tasks are registered and deregistered to the updated service registry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceRegistryProperty := &serviceRegistryProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	port: jsii.Number(123),
//   	registryArn: jsii.String("registryArn"),
//   }
//
type CfnTaskSet_ServiceRegistryProperty struct {
	// The container name value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition that your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition that your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field might be used if both the `awsvpc` network mode and SRV records are used.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map . For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) .
	RegistryArn *string `field:"optional" json:"registryArn" yaml:"registryArn"`
}

// Properties for defining a `CfnTaskSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskSetProps := &cfnTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskDefinition: jsii.String("taskDefinition"),
//
//   	// the properties below are optional
//   	externalId: jsii.String("externalId"),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsVpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	scale: &scaleProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   }
//
type CfnTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The task definition for the tasks in the task set to use.
	TaskDefinition *string `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map attribute set to the provided value.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The launch type that new tasks in the task set uses.
	//
	// For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If a `launchType` is specified, the `capacityProviderStrategy` parameter must be omitted.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A load balancer object representing the load balancer to use with the task set.
	//
	// The supported load balancer types are either an Application Load Balancer or a Network Load Balancer.
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the task set.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The platform version that the tasks in the task set uses.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// A floating-point percentage of your desired number of tasks to place and keep running in the task set.
	Scale interface{} `field:"optional" json:"scale" yaml:"scale"`
	// The details of the service discovery registries to assign to this task set.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
}

// The options for creating an AWS Cloud Map namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cloudMapNamespaceOptions := &cloudMapNamespaceOptions{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	type: awscdk.Aws_servicediscovery.namespaceType_HTTP,
//   	vpc: vpc,
//   }
//
type CloudMapNamespaceOptions struct {
	// The name of the namespace, such as example.com.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of CloudMap Namespace to create.
	Type awsservicediscovery.NamespaceType `field:"optional" json:"type" yaml:"type"`
	// The VPC to associate the namespace with.
	//
	// This property is required for private DNS namespaces.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// The options to enabling AWS Cloud Map for an Amazon ECS service.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create A records - useful for AWSVPC network mode.
//   		dnsRecordType: cloudmap.dnsRecordType_A,
//   	},
//   })
//
type CloudMapOptions struct {
	// The service discovery namespace for the Cloud Map service to attach to the ECS service.
	CloudMapNamespace awsservicediscovery.INamespace `field:"optional" json:"cloudMapNamespace" yaml:"cloudMapNamespace"`
	// The container to point to for a SRV record.
	Container ContainerDefinition `field:"optional" json:"container" yaml:"container"`
	// The port to point to for a SRV record.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The DNS record type that you want AWS Cloud Map to create.
	//
	// The supported record types are A or SRV.
	DnsRecordType awsservicediscovery.DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time that you want DNS resolvers to cache the settings for this record.
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	//
	// NOTE: This is used for HealthCheckCustomConfig.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The name of the Cloud Map service to attach to the ECS service.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A regional grouping of one or more container instances on which you can run tasks and services.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
//   	minCapacity: jsii.Number(0),
//   	maxCapacity: jsii.Number(100),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   })
//   cluster.addAsgCapacityProvider(capacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: capacityProvider.capacityProviderName,
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type Cluster interface {
	awscdk.Resource
	ICluster
	// Getter for autoscaling group added to cluster.
	AutoscalingGroup() awsautoscaling.IAutoScalingGroup
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn() *string
	// The name of the cluster.
	ClusterName() *string
	// Manage the allowed network connections for the cluster with Security Groups.
	Connections() awsec2.Connections
	// Getter for namespace added to cluster.
	DefaultCloudMapNamespace() awsservicediscovery.INamespace
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Getter for execute command configuration associated with the cluster.
	ExecuteCommandConfiguration() *ExecuteCommandConfiguration
	// Whether the cluster has EC2 capacity associated with it.
	HasEc2Capacity() *bool
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC associated with the cluster.
	Vpc() awsec2.IVpc
	// This method adds an Auto Scaling Group Capacity Provider to a cluster.
	AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions)
	// It is highly recommended to use {@link Cluster.addAsgCapacityProvider} instead of this method.
	//
	// This method adds compute capacity to a cluster by creating an AutoScalingGroup with the specified options.
	//
	// Returns the AutoScalingGroup so you can add autoscaling settings to it.
	AddCapacity(id *string, options *AddCapacityOptions) awsautoscaling.AutoScalingGroup
	// Add an AWS Cloud Map DNS namespace for this cluster.
	//
	// NOTE: HttpNamespaces are not supported, as ECS always requires a DNSConfig when registering an instance to a Cloud
	// Map service.
	AddDefaultCloudMapNamespace(options *CloudMapNamespaceOptions) awsservicediscovery.INamespace
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
	// Enable the Fargate capacity providers for this cluster.
	EnableFargateCapacityProviders()
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
	// This method returns the specifed CloudWatch metric for this cluster.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters CPU reservation.
	MetricCpuReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters memory reservation.
	MetricMemoryReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
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

// Import an existing cluster to the stack from the cluster ARN.
//
// This does not provide access to the vpc, hasEc2Capacity, or connections -
// use the `fromClusterAttributes` method to access those properties.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
		&returns,
	)

	return returns
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_Cluster) AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions) {
	_jsii_.InvokeVoid(
		c,
		"addAsgCapacityProvider",
		[]interface{}{provider, options},
	)
}

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

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var bucket bucket
//   var key key
//   var logGroup logGroup
//   var namespace iNamespace
//   var securityGroup securityGroup
//   var vpc vpc
//
//   clusterAttributes := &clusterAttributes{
//   	clusterName: jsii.String("clusterName"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	autoscalingGroup: autoScalingGroup,
//   	clusterArn: jsii.String("clusterArn"),
//   	defaultCloudMapNamespace: namespace,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: key,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			cloudWatchLogGroup: logGroup,
//   			s3Bucket: bucket,
//   			s3EncryptionEnabled: jsii.Boolean(false),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		logging: awscdk.Aws_ecs.executeCommandLogging_NONE,
//   	},
//   	hasEc2Capacity: jsii.Boolean(false),
//   }
//
type ClusterAttributes struct {
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The security groups associated with the container instances registered to the cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The VPC associated with the cluster.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Autoscaling group added to the cluster if capacity is added.
	AutoscalingGroup awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoscalingGroup" yaml:"autoscalingGroup"`
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The AWS Cloud Map namespace to associate with the cluster.
	DefaultCloudMapNamespace awsservicediscovery.INamespace `field:"optional" json:"defaultCloudMapNamespace" yaml:"defaultCloudMapNamespace"`
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// Specifies whether the cluster has EC2 instance capacity.
	HasEc2Capacity *bool `field:"optional" json:"hasEc2Capacity" yaml:"hasEc2Capacity"`
}

// The properties used to define an ECS cluster.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	memoryMiB: jsii.String("512"),
//   	cpu: jsii.String("256"),
//   	compatibility: ecs.compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	assignPublicIp: jsii.Boolean(true),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("SOME_KEY"),
//   					value: sfn.jsonPath.stringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	launchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
type ClusterProps struct {
	// The ec2 capacity to add to the cluster.
	Capacity *AddCapacityOptions `field:"optional" json:"capacity" yaml:"capacity"`
	// The name for the cluster.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// If true CloudWatch Container Insights will be enabled for the cluster.
	ContainerInsights *bool `field:"optional" json:"containerInsights" yaml:"containerInsights"`
	// The service discovery namespace created in this cluster.
	DefaultCloudMapNamespace *CloudMapNamespaceOptions `field:"optional" json:"defaultCloudMapNamespace" yaml:"defaultCloudMapNamespace"`
	// Whether to enable Fargate Capacity Providers.
	EnableFargateCapacityProviders *bool `field:"optional" json:"enableFargateCapacityProviders" yaml:"enableFargateCapacityProviders"`
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// The VPC where your ECS instances will be running or your ENIs will be deployed.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// The common task definition attributes used across all types of task definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   commonTaskDefinitionAttributes := &commonTaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type CommonTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

// The common properties for all task definitions.
//
// For more information, see
// [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var proxyConfiguration proxyConfiguration
//   var role role
//
//   commonTaskDefinitionProps := &commonTaskDefinitionProps{
//   	executionRole: role,
//   	family: jsii.String("family"),
//   	proxyConfiguration: proxyConfiguration,
//   	taskRole: role,
//   	volumes: []volume{
//   		&volume{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			dockerVolumeConfiguration: &dockerVolumeConfiguration{
//   				driver: jsii.String("driver"),
//   				scope: awscdk.Aws_ecs.scope_TASK,
//
//   				// the properties below are optional
//   				autoprovision: jsii.Boolean(false),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   			},
//   			efsVolumeConfiguration: &efsVolumeConfiguration{
//   				fileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfig{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &host{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   }
//
type CommonTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

// The task launch type compatibility requirement.
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
type Compatibility string

const (
	// The task should specify the EC2 launch type.
	Compatibility_EC2 Compatibility = "EC2"
	// The task should specify the Fargate launch type.
	Compatibility_FARGATE Compatibility = "FARGATE"
	// The task can specify either the EC2 or Fargate launch types.
	Compatibility_EC2_AND_FARGATE Compatibility = "EC2_AND_FARGATE"
	// The task should specify the External launch type.
	Compatibility_EXTERNAL Compatibility = "EXTERNAL"
)

// A container definition is used in a task definition to describe the containers that are launched as part of a task.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
type ContainerDefinition interface {
	constructs.Construct
	// An array dependencies defined for container startup and shutdown.
	ContainerDependencies() *[]*ContainerDependency
	// The name of this container.
	ContainerName() *string
	// The port the container will listen on.
	ContainerPort() *float64
	// The environment files for this container.
	EnvironmentFiles() *[]*EnvironmentFileConfig
	// Specifies whether the container will be marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container
	// fails or stops for any reason, all other containers that are part of the task are
	// stopped. If the essential parameter of a container is marked as false, then its
	// failure does not affect the rest of the containers in a task.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential() *bool
	// The name of the image referenced by this container.
	ImageName() *string
	// The inbound rules associated with the security group the task or service will use.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	IngressPort() *float64
	// The Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	LinuxParameters() LinuxParameters
	// The log configuration specification for the container.
	LogDriverConfig() *LogDriverConfig
	// Whether there was at least one memory limit specified in this definition.
	MemoryLimitSpecified() *bool
	// The mount points for data volumes in your container.
	MountPoints() *[]*MountPoint
	// The tree node.
	Node() constructs.Node
	// The list of port mappings for the container.
	//
	// Port mappings allow containers to access ports
	// on the host container instance to send or receive traffic.
	PortMappings() *[]*PortMapping
	// Whether this container definition references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The name of the task definition that includes this container definition.
	TaskDefinition() TaskDefinition
	// An array of ulimits to set in the container.
	Ulimits() *[]*Ulimit
	// The data volumes to mount from another container in the same task definition.
	VolumesFrom() *[]*VolumeFrom
	// This method adds one or more container dependencies to the container.
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	// This method adds an environment variable to the container.
	AddEnvironment(name *string, value *string)
	// This method adds one or more resources to the container.
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	// This method adds a link which allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the task definition is using the bridge network mode.
	// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
	AddLink(container ContainerDefinition, alias *string)
	// This method adds one or more mount points for data volumes to the container.
	AddMountPoints(mountPoints ...*MountPoint)
	// This method adds one or more port mappings to the container.
	AddPortMappings(portMappings ...*PortMapping)
	// This method mounts temporary disk space to the container.
	//
	// This adds the correct container mountPoint and task definition volume.
	AddScratch(scratch *ScratchSpace)
	// This method adds the specified statement to the IAM task execution policy in the task definition.
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	// This method adds one or more ulimits to the container.
	AddUlimits(ulimits ...*Ulimit)
	// This method adds one or more volumes to the container.
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	// Returns the host port for the requested container port if it exists.
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	// Render this container definition to a CloudFormation object.
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	// Returns a string representation of this construct.
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

func (j *jsiiProxy_ContainerDefinition) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_ContainerDefinition) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		c,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

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

func (c *jsiiProxy_ContainerDefinition) AddLink(container ContainerDefinition, alias *string) {
	_jsii_.InvokeVoid(
		c,
		"addLink",
		[]interface{}{container, alias},
	)
}

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

func (c *jsiiProxy_ContainerDefinition) AddScratch(scratch *ScratchSpace) {
	_jsii_.InvokeVoid(
		c,
		"addScratch",
		[]interface{}{scratch},
	)
}

func (c *jsiiProxy_ContainerDefinition) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		c,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

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

// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
type ContainerDefinitionOptions struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// The user name to use inside the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// The properties in a container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//   var taskDefinition taskDefinition
//
//   containerDefinitionProps := &containerDefinitionProps{
//   	image: containerImage,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: cdk.duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: cdk.*duration.minutes(jsii.Number(30)),
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostPort: jsii.Number(123),
//   			protocol: awscdk.Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	stopTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type ContainerDefinitionProps struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// The user name to use inside the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// The name of the task definition that includes this container definition.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
}

// The details of a dependency on another container in the task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerDefinition containerDefinition
//
//   containerDependency := &containerDependency{
//   	container: containerDefinition,
//
//   	// the properties below are optional
//   	condition: awscdk.Aws_ecs.containerDependencyCondition_START,
//   }
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDependency.html
//
type ContainerDependency struct {
	// The container to depend on.
	Container ContainerDefinition `field:"required" json:"container" yaml:"container"`
	// The state the container needs to be in to satisfy the dependency and proceed with startup.
	//
	// Valid values are ContainerDependencyCondition.START, ContainerDependencyCondition.COMPLETE,
	// ContainerDependencyCondition.SUCCESS and ContainerDependencyCondition.HEALTHY.
	Condition ContainerDependencyCondition `field:"optional" json:"condition" yaml:"condition"`
}

type ContainerDependencyCondition string

const (
	// This condition emulates the behavior of links and volumes today.
	//
	// It validates that a dependent container is started before permitting other containers to start.
	ContainerDependencyCondition_START ContainerDependencyCondition = "START"
	// This condition validates that a dependent container runs to completion (exits) before permitting other containers to start.
	//
	// This can be useful for nonessential containers that run a script and then exit.
	ContainerDependencyCondition_COMPLETE ContainerDependencyCondition = "COMPLETE"
	// This condition is the same as COMPLETE, but it also requires that the container exits with a zero status.
	ContainerDependencyCondition_SUCCESS ContainerDependencyCondition = "SUCCESS"
	// This condition validates that the dependent container passes its Docker health check before permitting other containers to start.
	//
	// This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.
	ContainerDependencyCondition_HEALTHY ContainerDependencyCondition = "HEALTHY"
)

// Constructs for types of container images.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCPCluster"), &clusterProps{
//   	vpc: vpc,
//   	enableFargateCapacityProviders: jsii.Boolean(true),
//   })
//
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   })
//
//   ecs.NewFargateService(this, jsii.String("FargateService"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("FARGATE_SPOT"),
//   			weight: jsii.Number(2),
//   		},
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("FARGATE"),
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type ContainerImage interface {
	// Called when the image is used by a ContainerDefinition.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerImageConfig := &containerImageConfig{
//   	imageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	repositoryCredentials: &repositoryCredentialsProperty{
//   		credentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   }
//
type ContainerImageConfig struct {
	// Specifies the name of the container image.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Specifies the credentials used to access the image repository.
	RepositoryCredentials *CfnTaskDefinition_RepositoryCredentialsProperty `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
}

// The CpuArchitecture for Fargate Runtime Platform.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
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
// Example:
//   var target applicationTargetGroup
//   var service baseService
//
//   scaling := service.autoScaleTaskCount(&enableScalingProps{
//   	maxCapacity: jsii.Number(10),
//   })
//   scaling.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scaling.scaleOnRequestCount(jsii.String("RequestScaling"), &requestCountScalingProps{
//   	requestsPerTarget: jsii.Number(10000),
//   	targetGroup: target,
//   })
//
type CpuUtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The target value for CPU utilization across all tasks in the service.
	TargetUtilizationPercent *float64 `field:"required" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}

// The deployment circuit breaker to use for the service.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(true),
//   	},
//   })
//
type DeploymentCircuitBreaker struct {
	// Whether to enable rollback on deployment failure.
	Rollback *bool `field:"optional" json:"rollback" yaml:"rollback"`
}

// The deployment controller to use for the service.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	deploymentController: &deploymentController{
//   		type: ecs.deploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
type DeploymentController struct {
	// The deployment controller type to use.
	Type DeploymentControllerType `field:"optional" json:"type" yaml:"type"`
}

// The deployment controller type to use for the service.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	deploymentController: &deploymentController{
//   		type: ecs.deploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
type DeploymentControllerType string

const (
	// The rolling update (ECS) deployment type involves replacing the current running version of the container with the latest version.
	DeploymentControllerType_ECS DeploymentControllerType = "ECS"
	// The blue/green (CODE_DEPLOY) deployment type uses the blue/green deployment model powered by AWS CodeDeploy.
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	// The external (EXTERNAL) deployment type enables you to use any third-party deployment controller.
	DeploymentControllerType_EXTERNAL DeploymentControllerType = "EXTERNAL"
)

// A container instance host device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   device := &device{
//   	hostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	containerPath: jsii.String("containerPath"),
//   	permissions: []devicePermission{
//   		awscdk.Aws_ecs.*devicePermission_READ,
//   	},
//   }
//
type Device struct {
	// The path for the device on the host container instance.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// The path inside the container at which to expose the host device.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for read, write, and mknod for the device.
	Permissions *[]DevicePermission `field:"optional" json:"permissions" yaml:"permissions"`
}

// Permissions for device access.
type DevicePermission string

const (
	// Read.
	DevicePermission_READ DevicePermission = "READ"
	// Write.
	DevicePermission_WRITE DevicePermission = "WRITE"
	// Make a node.
	DevicePermission_MKNOD DevicePermission = "MKNOD"
)

// The configuration for a Docker volume.
//
// Docker volumes are only supported when you are using the EC2 launch type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerVolumeConfiguration := &dockerVolumeConfiguration{
//   	driver: jsii.String("driver"),
//   	scope: awscdk.Aws_ecs.scope_TASK,
//
//   	// the properties below are optional
//   	autoprovision: jsii.Boolean(false),
//   	driverOpts: map[string]*string{
//   		"driverOptsKey": jsii.String("driverOpts"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   }
//
type DockerVolumeConfiguration struct {
	// The Docker volume driver to use.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// The scope for the Docker volume that determines its lifecycle.
	Scope Scope `field:"required" json:"scope" yaml:"scope"`
	// Specifies whether the Docker volume should be created if it does not already exist.
	//
	// If true is specified, the Docker volume will be created for you.
	Autoprovision *bool `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// A map of Docker driver-specific options passed through.
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

// This creates a service using the EC2 launch type on an ECS cluster.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   })
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//   lb.addTarget(service.loadBalancerTarget(&loadBalancerTargetOptions{
//   	containerName: jsii.String("MyContainer"),
//   	containerPort: jsii.Number(80),
//   }))
//
type Ec2Service interface {
	BaseService
	IEc2Service
	// The details of the AWS Cloud Map service.
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
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
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	TaskDefinition() TaskDefinition
	// Adds one or more placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	AddPlacementConstraints(constraints ...PlacementConstraint)
	// Adds one or more placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	AddPlacementStrategies(strategies ...PlacementStrategy)
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
	// Associates this service with a CloudMap service.
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	// This method is called to attach this service to an Application Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// This method is called to create a networkConfiguration.
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	// Enable CloudMap service discovery for the service.
	//
	// Returns: The created CloudMap service.
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
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
	// Return a load balancing target for a specific container and port.
	//
	// Use this function to create a load balancer target if you want to load balance to
	// another container than the first essential container or the first mapped port on
	// the container.
	//
	// Use the return value of this function where you would normally use a load balancer
	// target, instead of the `Service` object itself.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   listener.addTargets('ECS', {
	//     port: 80,
	//     targets: [service.loadBalancerTarget({
	//       containerName: 'MyContainer',
	//       containerPort: 1234,
	//     })],
	//   });
	//
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
	//
	// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   service.registerLoadBalancerTargets(
	//     {
	//       containerName: 'web',
	//       containerPort: 80,
	//       newTargetGroupId: 'ECS',
	//       listener: ecs.ListenerConfig.applicationListener(listener, {
	//         protocol: elbv2.ApplicationProtocol.HTTPS
	//       }),
	//     },
	//   )
	//
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	// Returns a string representation of this construct.
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

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func Ec2Service_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
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

func (e *jsiiProxy_Ec2Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_Ec2Service) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		e,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

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

func (e *jsiiProxy_Ec2Service) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		e,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

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

func (e *jsiiProxy_Ec2Service) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   ec2ServiceAttributes := &ec2ServiceAttributes{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type Ec2ServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The service ARN.
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

// The properties for defining a service using the EC2 launch type.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   })
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//   lb.addTarget(service.loadBalancerTarget(&loadBalancerTargetOptions{
//   	containerName: jsii.String("MyContainer"),
//   	containerPort: jsii.Number(80),
//   }))
//
type Ec2ServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Specifies whether the service will use the daemon scheduling strategy.
	//
	// If true, the service scheduler deploys exactly one task on each container instance in your cluster.
	//
	// When you are using this strategy, do not specify a desired number of tasks orany task placement strategies.
	Daemon *bool `field:"optional" json:"daemon" yaml:"daemon"`
	// The placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	PlacementStrategies *[]PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// The details of a task definition run on an EC2 cluster.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		token: awscdk.SecretValue.secretsManager(jsii.String("my-splunk-token")),
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type Ec2TaskDefinition interface {
	TaskDefinition
	IEc2TaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() Compatibility
	// The container definitions.
	Containers() *[]ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	NetworkMode() NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	// Adds an inference accelerator to the task definition.
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *Volume)
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
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) ContainerDefinition
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
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (e *jsiiProxy_Ec2TaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		e,
		"addExtension",
		[]interface{}{extension},
	)
}

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

func (e *jsiiProxy_Ec2TaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		e,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

func (e *jsiiProxy_Ec2TaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (e *jsiiProxy_Ec2TaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (e *jsiiProxy_Ec2TaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (e *jsiiProxy_Ec2TaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_Ec2TaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   ec2TaskDefinitionAttributes := &ec2TaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type Ec2TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

// The properties for a task definition run on an EC2 cluster.
//
// Example:
//   ec2TaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"), &ec2TaskDefinitionProps{
//   	networkMode: ecs.networkMode_BRIDGE,
//   })
//
//   container := ec2TaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
//   	// Use an image from DockerHub
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   })
//
type Ec2TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	InferenceAccelerators *[]*InferenceAccelerator `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	IpcMode IpcMode `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are NONE, BRIDGE, AWS_VPC, and HOST.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	PidMode PidMode `field:"optional" json:"pidMode" yaml:"pidMode"`
	// An array of placement constraint objects to use for the task.
	//
	// You can
	// specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
}

// An image from an Amazon ECR repository.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repo := ecr.repository.fromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))
//
//   batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.NewEcrImage(repo, jsii.String("latest")),
//   	},
//   })
//
type EcrImage interface {
	ContainerImage
	// The image name. Images in Amazon ECR repositories can be specified by either using the full registry/repository:tag or registry/repository@digest.
	//
	// For example, 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest or
	// 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE.
	ImageName() *string
	// Called when the image is used by a ContainerDefinition.
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
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   // Either add default capacity
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   // Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux(),
//   	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
//   	// machineImage: EcsOptimizedImage.amazonLinux2(),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   cluster.addAutoScalingGroup(autoScalingGroup)
//
type EcsOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
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
// Example:
//   var vpc vpc
//
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux(&ecsOptimizedImageOptions{
//   		cachedInContext: jsii.Boolean(true),
//   	}),
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   })
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
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
}

// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type EcsTarget struct {
	// The name of the container.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Listener and properties for adding target group to the listener.
	Listener ListenerConfig `field:"required" json:"listener" yaml:"listener"`
	// ID for a target group to be created.
	NewTargetGroupId *string `field:"required" json:"newTargetGroupId" yaml:"newTargetGroupId"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

// The configuration for an Elastic FileSystem volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsVolumeConfiguration := &efsVolumeConfiguration{
//   	fileSystemId: jsii.String("fileSystemId"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfig{
//   		accessPointId: jsii.String("accessPointId"),
//   		iam: jsii.String("iam"),
//   	},
//   	rootDirectory: jsii.String("rootDirectory"),
//   	transitEncryption: jsii.String("transitEncryption"),
//   	transitEncryptionPort: jsii.Number(123),
//   }
//
type EfsVolumeConfiguration struct {
	// The Amazon EFS file system ID to use.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// Specifying / will have the same effect as omitting this parameter.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be enabled if Amazon EFS IAM authorization is used.
	//
	// Valid values: ENABLED | DISABLED.
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// EFS mount helper uses.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

// Constructs for types of environment files.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   	environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	environmentFiles: []environmentFile{
//   		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret.fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//
type EnvironmentFile interface {
	// Called when the container is initialized to allow this object to bind to the stack.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileConfig := &environmentFileConfig{
//   	fileType: awscdk.Aws_ecs.environmentFileType_S3,
//   	s3Location: &location{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type EnvironmentFileConfig struct {
	// The type of environment file.
	FileType EnvironmentFileType `field:"required" json:"fileType" yaml:"fileType"`
	// The location of the environment file in S3.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

// Type of environment file to be included in the container definition.
type EnvironmentFileType string

const (
	// Environment file hosted on S3, referenced by object ARN.
	EnvironmentFileType_S3 EnvironmentFileType = "S3"
)

// The details of the execute command configuration.
//
// For more information, see
// [ExecuteCommandConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
//   	encryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
//   	encryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: kmsKey,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchLogGroup: logGroup,
//   			cloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			s3Bucket: execBucket,
//   			s3EncryptionEnabled: jsii.Boolean(true),
//   			s3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		logging: ecs.executeCommandLogging_OVERRIDE,
//   	},
//   })
//
type ExecuteCommandConfiguration struct {
	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The log configuration for the results of the execute command actions.
	//
	// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
	LogConfiguration *ExecuteCommandLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The log settings to use for logging the execute command session.
	Logging ExecuteCommandLogging `field:"optional" json:"logging" yaml:"logging"`
}

// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs and/ or an Amazon S3 bucket.
// For more information, see [ExecuteCommandLogConfiguration] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandlogconfiguration.html
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
//   	encryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
//   	encryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: kmsKey,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchLogGroup: logGroup,
//   			cloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			s3Bucket: execBucket,
//   			s3EncryptionEnabled: jsii.Boolean(true),
//   			s3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		logging: ecs.executeCommandLogging_OVERRIDE,
//   	},
//   })
//
type ExecuteCommandLogConfiguration struct {
	// Whether or not to enable encryption on the CloudWatch logs.
	CloudWatchEncryptionEnabled *bool `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// The CloudWatch log group must already be created.
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// The name of the S3 bucket to send logs to.
	//
	// The S3 bucket must already be created.
	S3Bucket awss3.IBucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Whether or not to enable encryption on the CloudWatch logs.
	S3EncryptionEnabled *bool `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

// The log settings to use to for logging the execute command session.
//
// For more information, see
// [Logging] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logging
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
//   	encryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
//   	encryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: kmsKey,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchLogGroup: logGroup,
//   			cloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			s3Bucket: execBucket,
//   			s3EncryptionEnabled: jsii.Boolean(true),
//   			s3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		logging: ecs.executeCommandLogging_OVERRIDE,
//   	},
//   })
//
type ExecuteCommandLogging string

const (
	// The execute command session is not logged.
	ExecuteCommandLogging_NONE ExecuteCommandLogging = "NONE"
	// The awslogs configuration in the task definition is used.
	//
	// If no logging parameter is specified, it defaults to this value. If no awslogs log driver is configured in the task definition, the output won't be logged.
	ExecuteCommandLogging_DEFAULT ExecuteCommandLogging = "DEFAULT"
	// Specify the logging details as a part of logConfiguration.
	ExecuteCommandLogging_OVERRIDE ExecuteCommandLogging = "OVERRIDE"
)

// This creates a service using the External launch type on an ECS cluster.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//
//   service := ecs.NewExternalService(this, jsii.String("Service"), &externalServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	desiredCount: jsii.Number(5),
//   })
//
type ExternalService interface {
	BaseService
	IExternalService
	// The details of the AWS Cloud Map service.
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
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
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	TaskDefinition() TaskDefinition
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
	// Overriden method to throw error as `associateCloudMapService` is not supported for external service.
	AssociateCloudMapService(_options *AssociateCloudMapServiceOptions)
	// Overriden method to throw error as `attachToApplicationTargetGroup` is not supported for external service.
	AttachToApplicationTargetGroup(_targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Overriden method to throw error as `autoScaleTaskCount` is not supported for external service.
	AutoScaleTaskCount(_props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// Overriden method to throw error as `configureAwsVpcNetworkingWithSecurityGroups` is not supported for external service.
	ConfigureAwsVpcNetworkingWithSecurityGroups(_vpc awsec2.IVpc, _assignPublicIp *bool, _vpcSubnets *awsec2.SubnetSelection, _securityGroups *[]awsec2.ISecurityGroup)
	// Overriden method to throw error as `enableCloudMap` is not supported for external service.
	EnableCloudMap(_options *CloudMapOptions) awsservicediscovery.Service
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
	// Overriden method to throw error as `loadBalancerTarget` is not supported for external service.
	LoadBalancerTarget(_options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Overriden method to throw error as `registerLoadBalancerTargets` is not supported for external service.
	RegisterLoadBalancerTargets(_targets ...*EcsTarget)
	// Returns a string representation of this construct.
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

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func ExternalService_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ExternalService",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
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

func (e *jsiiProxy_ExternalService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_ExternalService) AssociateCloudMapService(_options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		e,
		"associateCloudMapService",
		[]interface{}{_options},
	)
}

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

func (e *jsiiProxy_ExternalService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		e,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

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

func (e *jsiiProxy_ExternalService) ConfigureAwsVpcNetworkingWithSecurityGroups(_vpc awsec2.IVpc, _assignPublicIp *bool, _vpcSubnets *awsec2.SubnetSelection, _securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{_vpc, _assignPublicIp, _vpcSubnets, _securityGroups},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   externalServiceAttributes := &externalServiceAttributes{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type ExternalServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The service ARN.
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

// The properties for defining a service using the External launch type.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//
//   service := ecs.NewExternalService(this, jsii.String("Service"), &externalServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	desiredCount: jsii.Number(5),
//   })
//
type ExternalServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

// The details of a task definition run on an External cluster.
//
// Example:
//   externalTaskDefinition := ecs.NewExternalTaskDefinition(this, jsii.String("TaskDef"))
//
//   container := externalTaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
//   	// Use an image from DockerHub
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   })
//
type ExternalTaskDefinition interface {
	TaskDefinition
	IExternalTaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() Compatibility
	// The container definitions.
	Containers() *[]ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	NetworkMode() NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	// Overriden method to throw error as interface accelerators are not supported for external tasks.
	AddInferenceAccelerator(_inferenceAccelerator *InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *Volume)
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
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) ContainerDefinition
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
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (e *jsiiProxy_ExternalTaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		e,
		"addExtension",
		[]interface{}{extension},
	)
}

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

func (e *jsiiProxy_ExternalTaskDefinition) AddInferenceAccelerator(_inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		e,
		"addInferenceAccelerator",
		[]interface{}{_inferenceAccelerator},
	)
}

func (e *jsiiProxy_ExternalTaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (e *jsiiProxy_ExternalTaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (e *jsiiProxy_ExternalTaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (e *jsiiProxy_ExternalTaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_ExternalTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   externalTaskDefinitionAttributes := &externalTaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type ExternalTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

// The properties for a task definition run on an External cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var proxyConfiguration proxyConfiguration
//   var role role
//
//   externalTaskDefinitionProps := &externalTaskDefinitionProps{
//   	executionRole: role,
//   	family: jsii.String("family"),
//   	proxyConfiguration: proxyConfiguration,
//   	taskRole: role,
//   	volumes: []volume{
//   		&volume{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			dockerVolumeConfiguration: &dockerVolumeConfiguration{
//   				driver: jsii.String("driver"),
//   				scope: awscdk.Aws_ecs.scope_TASK,
//
//   				// the properties below are optional
//   				autoprovision: jsii.Boolean(false),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   			},
//   			efsVolumeConfiguration: &efsVolumeConfiguration{
//   				fileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfig{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &host{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   }
//
type ExternalTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

// The platform version on which to run your service.
//
// Example:
//   var cluster cluster
//
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
//   	cluster: cluster,
//   	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		memoryLimitMiB: jsii.Number(512),
//   	},
//   	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
//   	platformVersion: ecs.fargatePlatformVersion_LATEST,
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
//
type FargatePlatformVersion string

const (
	// The latest, recommended platform version.
	FargatePlatformVersion_LATEST FargatePlatformVersion = "LATEST"
	// Version 1.4.0.
	//
	// Supports EFS endpoints, CAP_SYS_PTRACE Linux capability,
	// network performance metrics in CloudWatch Container Insights,
	// consolidated 20 GB ephemeral volume.
	FargatePlatformVersion_VERSION1_4 FargatePlatformVersion = "VERSION1_4"
	// Version 1.3.0.
	//
	// Supports secrets, task recycling.
	FargatePlatformVersion_VERSION1_3 FargatePlatformVersion = "VERSION1_3"
	// Version 1.2.0.
	//
	// Supports private registries.
	FargatePlatformVersion_VERSION1_2 FargatePlatformVersion = "VERSION1_2"
	// Version 1.1.0.
	//
	// Supports task metadata, health checks, service discovery.
	FargatePlatformVersion_VERSION1_1 FargatePlatformVersion = "VERSION1_1"
	// Initial release.
	//
	// Based on Amazon Linux 2017.09.
	FargatePlatformVersion_VERSION1_0 FargatePlatformVersion = "VERSION1_0"
)

// This creates a service using the Fargate launch type on an ECS cluster.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type FargateService interface {
	BaseService
	IFargateService
	// The details of the AWS Cloud Map service.
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
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
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	TaskDefinition() TaskDefinition
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
	// Associates this service with a CloudMap service.
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	// This method is called to attach this service to an Application Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// This method is called to create a networkConfiguration.
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	// Enable CloudMap service discovery for the service.
	//
	// Returns: The created CloudMap service.
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
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
	// Return a load balancing target for a specific container and port.
	//
	// Use this function to create a load balancer target if you want to load balance to
	// another container than the first essential container or the first mapped port on
	// the container.
	//
	// Use the return value of this function where you would normally use a load balancer
	// target, instead of the `Service` object itself.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   listener.addTargets('ECS', {
	//     port: 80,
	//     targets: [service.loadBalancerTarget({
	//       containerName: 'MyContainer',
	//       containerPort: 1234,
	//     })],
	//   });
	//
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
	//
	// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
	//
	// Example:
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   service.registerLoadBalancerTargets(
	//     {
	//       containerName: 'web',
	//       containerPort: 80,
	//       newTargetGroupId: 'ECS',
	//       listener: ecs.ListenerConfig.applicationListener(listener, {
	//         protocol: elbv2.ApplicationProtocol.HTTPS
	//       }),
	//     },
	//   )
	//
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	// Returns a string representation of this construct.
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

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func FargateService_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FargateService",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
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

func (f *jsiiProxy_FargateService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FargateService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	_jsii_.InvokeVoid(
		f,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

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

func (f *jsiiProxy_FargateService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	_jsii_.InvokeVoid(
		f,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

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

func (f *jsiiProxy_FargateService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		f,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   fargateServiceAttributes := &fargateServiceAttributes{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type FargateServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The service ARN.
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

// The properties for defining a service using the Fargate launch type.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type FargateServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The task definition to use for tasks in the service.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// The details of a task definition run on a Fargate cluster.
//
// Example:
//   fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	memoryLimitMiB: jsii.Number(512),
//   	cpu: jsii.Number(256),
//   })
//   container := fargateTaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
//   	// Use an image from DockerHub
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   })
//
type FargateTaskDefinition interface {
	TaskDefinition
	IFargateTaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() Compatibility
	// The container definitions.
	Containers() *[]ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The Docker networking mode to use for the containers in the task.
	//
	// Fargate tasks require the awsvpc network mode.
	NetworkMode() NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	// Adds an inference accelerator to the task definition.
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *Volume)
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
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) ContainerDefinition
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
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (f *jsiiProxy_FargateTaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		f,
		"addExtension",
		[]interface{}{extension},
	)
}

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

func (f *jsiiProxy_FargateTaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		f,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

func (f *jsiiProxy_FargateTaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		f,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (f *jsiiProxy_FargateTaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (f *jsiiProxy_FargateTaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (f *jsiiProxy_FargateTaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		f,
		"addVolume",
		[]interface{}{volume},
	)
}

func (f *jsiiProxy_FargateTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   fargateTaskDefinitionAttributes := &fargateTaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type FargateTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

// The properties for a task definition.
//
// Example:
//   fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	memoryLimitMiB: jsii.Number(512),
//   	cpu: jsii.Number(256),
//   })
//   container := fargateTaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
//   	// Use an image from DockerHub
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   })
//
type FargateTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
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
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// The maximum supported value is 200 GiB.
	//
	// NOTE: This parameter is only supported for tasks hosted on AWS Fargate using platform version 1.4.0 or later.
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
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
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU).
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The operating system that your task definitions are running on.
	//
	// A runtimePlatform is supported only for tasks using the Fargate launch type.
	RuntimePlatform *RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
}

// FireLens enables you to use task definition parameters to route logs to an AWS service   or AWS Partner Network (APN) destination for log storage and analytics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   fireLensLogDriver := awscdk.Aws_ecs.NewFireLensLogDriver(&fireLensLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: map[string]*secret{
//   		"secretOptionsKey": secret,
//   	},
//   	tag: jsii.String("tag"),
//   })
//
type FireLensLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.firelens(&fireLensLogDriverProps{
//   		options: map[string]*string{
//   			"Name": jsii.String("firehose"),
//   			"region": jsii.String("us-west-2"),
//   			"delivery_stream": jsii.String("my-stream"),
//   		},
//   	}),
//   })
//
type FireLensLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *map[string]Secret `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

// Firelens Configuration https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfig := &firelensConfig{
//   	type: awscdk.Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   	// the properties below are optional
//   	options: &firelensOptions{
//   		configFileValue: jsii.String("configFileValue"),
//
//   		// the properties below are optional
//   		configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   		enableECSLogMetadata: jsii.Boolean(false),
//   	},
//   }
//
type FirelensConfig struct {
	// The log router to use.
	Type FirelensLogRouterType `field:"required" json:"type" yaml:"type"`
	// Firelens options.
	Options *FirelensOptions `field:"optional" json:"options" yaml:"options"`
}

// Firelens configuration file type, s3 or file path.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig
type FirelensConfigFileType string

const (
	// s3.
	FirelensConfigFileType_S3 FirelensConfigFileType = "S3"
	// fluentd.
	FirelensConfigFileType_FILE FirelensConfigFileType = "FILE"
)

// Firelens log router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//   var taskDefinition taskDefinition
//
//   firelensLogRouter := awscdk.Aws_ecs.NewFirelensLogRouter(this, jsii.String("MyFirelensLogRouter"), &firelensLogRouterProps{
//   	firelensConfig: &firelensConfig{
//   		type: awscdk.*Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		options: &firelensOptions{
//   			configFileValue: jsii.String("configFileValue"),
//
//   			// the properties below are optional
//   			configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   			enableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	image: containerImage,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: cdk.duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: cdk.*duration.minutes(jsii.Number(30)),
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostPort: jsii.Number(123),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	stopTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
type FirelensLogRouter interface {
	ContainerDefinition
	// An array dependencies defined for container startup and shutdown.
	ContainerDependencies() *[]*ContainerDependency
	// The name of this container.
	ContainerName() *string
	// The port the container will listen on.
	ContainerPort() *float64
	// The environment files for this container.
	EnvironmentFiles() *[]*EnvironmentFileConfig
	// Specifies whether the container will be marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container
	// fails or stops for any reason, all other containers that are part of the task are
	// stopped. If the essential parameter of a container is marked as false, then its
	// failure does not affect the rest of the containers in a task.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential() *bool
	// Firelens configuration.
	FirelensConfig() *FirelensConfig
	// The name of the image referenced by this container.
	ImageName() *string
	// The inbound rules associated with the security group the task or service will use.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	IngressPort() *float64
	// The Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	LinuxParameters() LinuxParameters
	// The log configuration specification for the container.
	LogDriverConfig() *LogDriverConfig
	// Whether there was at least one memory limit specified in this definition.
	MemoryLimitSpecified() *bool
	// The mount points for data volumes in your container.
	MountPoints() *[]*MountPoint
	// The tree node.
	Node() constructs.Node
	// The list of port mappings for the container.
	//
	// Port mappings allow containers to access ports
	// on the host container instance to send or receive traffic.
	PortMappings() *[]*PortMapping
	// Whether this container definition references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The name of the task definition that includes this container definition.
	TaskDefinition() TaskDefinition
	// An array of ulimits to set in the container.
	Ulimits() *[]*Ulimit
	// The data volumes to mount from another container in the same task definition.
	VolumesFrom() *[]*VolumeFrom
	// This method adds one or more container dependencies to the container.
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	// This method adds an environment variable to the container.
	AddEnvironment(name *string, value *string)
	// This method adds one or more resources to the container.
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	// This method adds a link which allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the task definition is using the bridge network mode.
	// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
	AddLink(container ContainerDefinition, alias *string)
	// This method adds one or more mount points for data volumes to the container.
	AddMountPoints(mountPoints ...*MountPoint)
	// This method adds one or more port mappings to the container.
	AddPortMappings(portMappings ...*PortMapping)
	// This method mounts temporary disk space to the container.
	//
	// This adds the correct container mountPoint and task definition volume.
	AddScratch(scratch *ScratchSpace)
	// This method adds the specified statement to the IAM task execution policy in the task definition.
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	// This method adds one or more ulimits to the container.
	AddUlimits(ulimits ...*Ulimit)
	// This method adds one or more volumes to the container.
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	// Returns the host port for the requested container port if it exists.
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	// Render this container definition to a CloudFormation object.
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	// Returns a string representation of this construct.
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

func (j *jsiiProxy_FirelensLogRouter) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
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
// Deprecated: use `x instanceof Construct` instead.
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

func (f *jsiiProxy_FirelensLogRouter) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		f,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

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

func (f *jsiiProxy_FirelensLogRouter) AddLink(container ContainerDefinition, alias *string) {
	_jsii_.InvokeVoid(
		f,
		"addLink",
		[]interface{}{container, alias},
	)
}

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

func (f *jsiiProxy_FirelensLogRouter) AddScratch(scratch *ScratchSpace) {
	_jsii_.InvokeVoid(
		f,
		"addScratch",
		[]interface{}{scratch},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//
//   firelensLogRouterDefinitionOptions := &firelensLogRouterDefinitionOptions{
//   	firelensConfig: &firelensConfig{
//   		type: awscdk.Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		options: &firelensOptions{
//   			configFileValue: jsii.String("configFileValue"),
//
//   			// the properties below are optional
//   			configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   			enableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	image: containerImage,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: cdk.duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: cdk.*duration.minutes(jsii.Number(30)),
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostPort: jsii.Number(123),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	stopTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type FirelensLogRouterDefinitionOptions struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// The user name to use inside the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// Firelens configuration.
	FirelensConfig *FirelensConfig `field:"required" json:"firelensConfig" yaml:"firelensConfig"`
}

// The properties in a firelens log router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//   var taskDefinition taskDefinition
//
//   firelensLogRouterProps := &firelensLogRouterProps{
//   	firelensConfig: &firelensConfig{
//   		type: awscdk.Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		options: &firelensOptions{
//   			configFileValue: jsii.String("configFileValue"),
//
//   			// the properties below are optional
//   			configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   			enableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	image: containerImage,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: cdk.duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: cdk.*duration.minutes(jsii.Number(30)),
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostPort: jsii.Number(123),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	stopTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type FirelensLogRouterProps struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// The user name to use inside the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// The name of the task definition that includes this container definition.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Firelens configuration.
	FirelensConfig *FirelensConfig `field:"required" json:"firelensConfig" yaml:"firelensConfig"`
}

// Firelens log router type, fluentbit or fluentd.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html
type FirelensLogRouterType string

const (
	// fluentbit.
	FirelensLogRouterType_FLUENTBIT FirelensLogRouterType = "FLUENTBIT"
	// fluentd.
	FirelensLogRouterType_FLUENTD FirelensLogRouterType = "FLUENTD"
)

// The options for firelens log router https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensOptions := &firelensOptions{
//   	configFileValue: jsii.String("configFileValue"),
//
//   	// the properties below are optional
//   	configFileType: awscdk.Aws_ecs.firelensConfigFileType_S3,
//   	enableECSLogMetadata: jsii.Boolean(false),
//   }
//
type FirelensOptions struct {
	// Custom configuration file, S3 ARN or a file path.
	ConfigFileValue *string `field:"required" json:"configFileValue" yaml:"configFileValue"`
	// Custom configuration file, s3 or file.
	ConfigFileType FirelensConfigFileType `field:"optional" json:"configFileType" yaml:"configFileType"`
	// By default, Amazon ECS adds additional fields in your log entries that help identify the source of the logs.
	//
	// You can disable this action by setting enable-ecs-log-metadata to false.
	EnableECSLogMetadata *bool `field:"optional" json:"enableECSLogMetadata" yaml:"enableECSLogMetadata"`
}

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fluentdLogDriver := awscdk.Aws_ecs.NewFluentdLogDriver(&fluentdLogDriverProps{
//   	address: jsii.String("address"),
//   	asyncConnect: jsii.Boolean(false),
//   	bufferLimit: jsii.Number(123),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxRetries: jsii.Number(123),
//   	retryWait: cdk.duration.minutes(jsii.Number(30)),
//   	subSecondPrecision: jsii.Boolean(false),
//   	tag: jsii.String("tag"),
//   })
//
type FluentdLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fluentdLogDriverProps := &fluentdLogDriverProps{
//   	address: jsii.String("address"),
//   	asyncConnect: jsii.Boolean(false),
//   	bufferLimit: jsii.Number(123),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxRetries: jsii.Number(123),
//   	retryWait: cdk.duration.minutes(jsii.Number(30)),
//   	subSecondPrecision: jsii.Boolean(false),
//   	tag: jsii.String("tag"),
//   }
//
type FluentdLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// By default, the logging driver connects to localhost:24224.
	//
	// Supply the
	// address option to connect to a different address. tcp(default) and unix
	// sockets are supported.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docker connects to Fluentd in the background.
	//
	// Messages are buffered until
	// the connection is established.
	AsyncConnect *bool `field:"optional" json:"asyncConnect" yaml:"asyncConnect"`
	// The amount of data to buffer before flushing to disk.
	BufferLimit *float64 `field:"optional" json:"bufferLimit" yaml:"bufferLimit"`
	// The maximum number of retries.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// How long to wait between retries.
	RetryWait awscdk.Duration `field:"optional" json:"retryWait" yaml:"retryWait"`
	// Generates event logs in nanosecond resolution.
	SubSecondPrecision *bool `field:"optional" json:"subSecondPrecision" yaml:"subSecondPrecision"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gelfLogDriver := awscdk.Aws_ecs.NewGelfLogDriver(&gelfLogDriverProps{
//   	address: jsii.String("address"),
//
//   	// the properties below are optional
//   	compressionLevel: jsii.Number(123),
//   	compressionType: awscdk.*Aws_ecs.gelfCompressionType_GZIP,
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tcpMaxReconnect: jsii.Number(123),
//   	tcpReconnectDelay: cdk.duration.minutes(jsii.Number(30)),
//   })
//
type GelfLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.gelf(&gelfLogDriverProps{
//   		address: jsii.String("my-gelf-address"),
//   	}),
//   })
//
type GelfLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The address of the GELF server.
	//
	// tcp and udp are the only supported URI
	// specifier and you must specify the port.
	Address *string `field:"required" json:"address" yaml:"address"`
	// UDP Only The level of compression when gzip or zlib is the gelf-compression-type.
	//
	// An integer in the range of -1 to 9 (BestCompression). Higher levels provide more
	// compression at lower speed. Either -1 or 0 disables compression.
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// UDP Only The type of compression the GELF driver uses to compress each log message.
	//
	// Allowed values are gzip, zlib and none.
	CompressionType GelfCompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// TCP Only The maximum number of reconnection attempts when the connection drop.
	//
	// A positive integer.
	TcpMaxReconnect *float64 `field:"optional" json:"tcpMaxReconnect" yaml:"tcpMaxReconnect"`
	// TCP Only The number of seconds to wait between reconnection attempts.
	//
	// A positive integer.
	TcpReconnectDelay awscdk.Duration `field:"optional" json:"tcpReconnectDelay" yaml:"tcpReconnectDelay"`
}

// A log driver that sends logs to the specified driver.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewGenericLogDriver(&genericLogDriverProps{
//   		logDriver: jsii.String("fluentd"),
//   		options: map[string]*string{
//   			"tag": jsii.String("example-tag"),
//   		},
//   	}),
//   })
//
type GenericLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewGenericLogDriver(&genericLogDriverProps{
//   		logDriver: jsii.String("fluentd"),
//   		options: map[string]*string{
//   			"tag": jsii.String("example-tag"),
//   		},
//   	}),
//   })
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
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *map[string]Secret `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

// The health check command and associated configuration parameters for the container.
//
// Example:
//   var vpc vpc
//   var securityGroup securityGroup
//
//   queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
//   	vpc: vpc,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.containerImage.fromRegistry(jsii.String("test")),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("CMD-SHELL"),
//   			jsii.String("curl -f http://localhost/ || exit 1"),
//   		},
//   		// the properties below are optional
//   		interval: awscdk.Duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: awscdk.Duration.minutes(jsii.Number(30)),
//   		timeout: awscdk.Duration.minutes(jsii.Number(30)),
//   	},
//   })
//
type HealthCheck struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with CMD to execute the command arguments directly, or
	// CMD-SHELL to run the command with the container's default shell.
	//
	// For example: [ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ].
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You may specify between 0 and 300 seconds.
	StartPeriod awscdk.Duration `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// The details on a container instance bind mount host volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   host := &host{
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type Host struct {
	// Specifies the path on the host container instance that is presented to the container.
	//
	// If the sourcePath value does not exist on the host container instance, the Docker daemon creates it.
	// If the location does exist, the contents of the source path folder are exported.
	//
	// This property is not supported for tasks that use the Fargate launch type.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceAccelerator := &inferenceAccelerator{
//   	deviceName: jsii.String("deviceName"),
//   	deviceType: jsii.String("deviceType"),
//   }
//
type InferenceAccelerator struct {
	// The Elastic Inference accelerator device name.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	//
	// The allowed values are: eia2.medium, eia2.large and eia2.xlarge.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

// The IPC resource namespace to use for the containers in the task.
type IpcMode string

const (
	// If none is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance.
	IpcMode_NONE IpcMode = "NONE"
	// If host is specified, then all containers within the tasks that specified the host IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance.
	IpcMode_HOST IpcMode = "HOST"
	// If task is specified, all containers within the specified task share the same IPC resources.
	IpcMode_TASK IpcMode = "TASK"
)

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   journaldLogDriver := awscdk.Aws_ecs.NewJournaldLogDriver(&journaldLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   })
//
type JournaldLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   journaldLogDriverProps := &journaldLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   }
//
type JournaldLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

// A log driver that sends log information to json-file Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFileLogDriver := awscdk.Aws_ecs.NewJsonFileLogDriver(&jsonFileLogDriverProps{
//   	compress: jsii.Boolean(false),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxFile: jsii.Number(123),
//   	maxSize: jsii.String("maxSize"),
//   	tag: jsii.String("tag"),
//   })
//
type JsonFileLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFileLogDriverProps := &jsonFileLogDriverProps{
//   	compress: jsii.Boolean(false),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxFile: jsii.Number(123),
//   	maxSize: jsii.String("maxSize"),
//   	tag: jsii.String("tag"),
//   }
//
type JsonFileLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Toggles compression for rotated logs.
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The maximum number of log files that can be present.
	//
	// If rolling the logs creates
	// excess files, the oldest file is removed. Only effective when max-size is also set.
	// A positive integer.
	MaxFile *float64 `field:"optional" json:"maxFile" yaml:"maxFile"`
	// The maximum size of the log before it is rolled.
	//
	// A positive integer plus a modifier
	// representing the unit of measure (k, m, or g).
	MaxSize *string `field:"optional" json:"maxSize" yaml:"maxSize"`
}

// The launch type of an ECS service.
type LaunchType string

const (
	// The service will be launched using the EC2 launch type.
	LaunchType_EC2 LaunchType = "EC2"
	// The service will be launched using the FARGATE launch type.
	LaunchType_FARGATE LaunchType = "FARGATE"
	// The service will be launched using the EXTERNAL launch type.
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

// Linux-specific options that are applied to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParameters := awscdk.Aws_ecs.NewLinuxParameters(this, jsii.String("MyLinuxParameters"), &linuxParametersProps{
//   	initProcessEnabled: jsii.Boolean(false),
//   	sharedMemorySize: jsii.Number(123),
//   })
//
type LinuxParameters interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Adds one or more Linux capabilities to the Docker configuration of a container.
	//
	// Only works with EC2 launch type.
	AddCapabilities(cap ...Capability)
	// Adds one or more host devices to a container.
	AddDevices(device ...*Device)
	// Specifies the container path, mount options, and size (in MiB) of the tmpfs mount for a container.
	//
	// Only works with EC2 launch type.
	AddTmpfs(tmpfs ...*Tmpfs)
	// Removes one or more Linux capabilities to the Docker configuration of a container.
	//
	// Only works with EC2 launch type.
	DropCapabilities(cap ...Capability)
	// Renders the Linux parameters to a CloudFormation object.
	RenderLinuxParameters() *CfnTaskDefinition_LinuxParametersProperty
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParametersProps := &linuxParametersProps{
//   	initProcessEnabled: jsii.Boolean(false),
//   	sharedMemorySize: jsii.Number(123),
//   }
//
type LinuxParametersProps struct {
	// Specifies whether to run an init process inside the container that forwards signals and reaps processes.
	InitProcessEnabled *bool `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The value for the size (in MiB) of the /dev/shm volume.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
}

// Base class for configuring listener when registering targets.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type ListenerConfig interface {
	// Create and attach a target group to listener.
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
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   })
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//   lb.addTarget(service.loadBalancerTarget(&loadBalancerTargetOptions{
//   	containerName: jsii.String("MyContainer"),
//   	containerPort: jsii.Number(80),
//   }))
//
type LoadBalancerTargetOptions struct {
	// The name of the container.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

// The base class for log drivers.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		token: awscdk.SecretValue.secretsManager(jsii.String("my-splunk-token")),
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type LogDriver interface {
	// Called when the log driver is configured on a container.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDriverConfig := &logDriverConfig{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: []secretProperty{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
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
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *[]*CfnTaskDefinition_SecretProperty `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

// The base class for log drivers.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("EventDemo"),
//   	}),
//   })
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
// Example:
//   var cluster cluster
//
//
//   cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	machineImageType: ecs.machineImageType_BOTTLEROCKET,
//   })
//
type MachineImageType string

const (
	// Amazon ECS-optimized Amazon Linux 2 AMI.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

// The properties for enabling scaling based on memory utilization.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
type MemoryUtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The target value for memory utilization across all tasks in the service.
	TargetUtilizationPercent *float64 `field:"required" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}

// The details of data volume mount points for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPoint := &mountPoint{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type MountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Specifies whether to give the container read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	SourceVolume *string `field:"required" json:"sourceVolume" yaml:"sourceVolume"`
}

// The networking mode to use for the containers in the task.
//
// Example:
//   ec2TaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"), &ec2TaskDefinitionProps{
//   	networkMode: ecs.networkMode_BRIDGE,
//   })
//
//   container := ec2TaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
//   	// Use an image from DockerHub
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   })
//
type NetworkMode string

const (
	// The task's containers do not have external connectivity and port mappings can't be specified in the container definition.
	NetworkMode_NONE NetworkMode = "NONE"
	// The task utilizes Docker's built-in virtual network which runs inside each container instance.
	NetworkMode_BRIDGE NetworkMode = "BRIDGE"
	// The task is allocated an elastic network interface.
	NetworkMode_AWS_VPC NetworkMode = "AWS_VPC"
	// The task bypasses Docker's built-in virtual network and maps container ports directly to the EC2 instance's network interface directly.
	//
	// In this mode, you can't run multiple instantiations of the same task on a
	// single container instance when port mappings are used.
	NetworkMode_HOST NetworkMode = "HOST"
	// The task utilizes NAT network mode required by Windows containers.
	//
	// This is the only supported network mode for Windows containers. For more information, see
	// [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#network_mode).
	NetworkMode_NAT NetworkMode = "NAT"
)

// The operating system for Fargate Runtime Platform.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
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
	// If host is specified, then all containers within the tasks that specified the host PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance.
	PidMode_HOST PidMode = "HOST"
	// If task is specified, all containers within the specified task share the same process namespace.
	PidMode_TASK PidMode = "TASK"
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

// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
//
// Example:
//   var container containerDefinition
//
//
//   container.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(3000),
//   })
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
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
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
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The protocol used for the port mapping.
	//
	// Valid values are Protocol.TCP and Protocol.UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

// Propagate tags from either service or task definition.
type PropagatedTagSource string

const (
	// Propagate tags from service.
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	// Propagate tags from task definition.
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	// Do not propagate.
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

// Network protocol.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
type Protocol string

const (
	// TCP.
	Protocol_TCP Protocol = "TCP"
	// UDP.
	Protocol_UDP Protocol = "UDP"
)

// The base class for proxy configurations.
type ProxyConfiguration interface {
	// Called when the proxy configuration is configured on a task definition.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurations := awscdk.Aws_ecs.NewProxyConfigurations()
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
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import batch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
//   	"computeEnvironments": []map[string]interface{}{
//   		map[string]interface{}{
//   			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			"order": jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
//   })
//
//   rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
//   	deadLetterQueue: queue,
//   	event: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	retryAttempts: jsii.Number(2),
//   	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
//   }))
//
type RepositoryImage interface {
	ContainerImage
	// Called when the image is used by a ContainerDefinition.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   repositoryImageProps := &repositoryImageProps{
//   	credentials: secret,
//   }
//
type RepositoryImageProps struct {
	// The secret to expose to the container that contains the credentials for the image repository.
	//
	// The supported value is the full ARN of an AWS Secrets Manager secret.
	Credentials awssecretsmanager.ISecret `field:"optional" json:"credentials" yaml:"credentials"`
}

// The properties for enabling scaling based on Application Load Balancer (ALB) request counts.
//
// Example:
//   var target applicationTargetGroup
//   var service baseService
//
//   scaling := service.autoScaleTaskCount(&enableScalingProps{
//   	maxCapacity: jsii.Number(10),
//   })
//   scaling.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scaling.scaleOnRequestCount(jsii.String("RequestScaling"), &requestCountScalingProps{
//   	requestsPerTarget: jsii.Number(10000),
//   	targetGroup: target,
//   })
//
type RequestCountScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The number of ALB requests per target.
	RequestsPerTarget *float64 `field:"required" json:"requestsPerTarget" yaml:"requestsPerTarget"`
	// The ALB target group name.
	TargetGroup awselasticloadbalancingv2.ApplicationTargetGroup `field:"required" json:"targetGroup" yaml:"targetGroup"`
}

// The interface for Runtime Platform.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
//
type RuntimePlatform struct {
	// The CpuArchitecture for Fargate Runtime Platform.
	CpuArchitecture CpuArchitecture `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system for Fargate Runtime Platform.
	OperatingSystemFamily OperatingSystemFamily `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

// Environment file from S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3EnvironmentFile := awscdk.Aws_ecs.NewS3EnvironmentFile(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
type S3EnvironmentFile interface {
	EnvironmentFile
	// Called when the container is initialized to allow this object to bind to the stack.
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
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
type ScalableTaskCount interface {
	awsapplicationautoscaling.BaseScalableAttribute
	// The tree node.
	Node() constructs.Node
	Props() *awsapplicationautoscaling.BaseScalableAttributeProps
	// Scale out or in based on a metric value.
	DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scale out or in based on time.
	DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps)
	// Scales in or out to achieve a target CPU utilization.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps)
	// Scales in or out to achieve a target memory utilization.
	ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps)
	// Scales in or out based on a specified metric value.
	ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scales in or out to achieve a target Application Load Balancer request count per target.
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps)
	// Scales in or out based on a specified scheduled time.
	ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scales in or out to achieve a target on a custom metric.
	ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps)
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMemoryUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnRequestCount",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleToTrackCustomMetric",
		[]interface{}{id, props},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   scalableTaskCountProps := &scalableTaskCountProps{
//   	dimension: jsii.String("dimension"),
//   	maxCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	role: role,
//   	serviceNamespace: awscdk.Aws_applicationautoscaling.serviceNamespace_ECS,
//
//   	// the properties below are optional
//   	minCapacity: jsii.Number(123),
//   }
//
type ScalableTaskCountProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Scalable dimension of the attribute.
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// Resource ID of the attribute.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to use for scaling.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Service namespace of the scalable attribute.
	ServiceNamespace awsapplicationautoscaling.ServiceNamespace `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

// The scope for the Docker volume that determines its lifecycle.
//
// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
// Docker volumes that are scoped as shared persist after the task stops.
type Scope string

const (
	// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
	Scope_TASK Scope = "TASK"
	// Docker volumes that are scoped as shared persist after the task stops.
	Scope_SHARED Scope = "SHARED"
)

// The temporary disk space mounted to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scratchSpace := &scratchSpace{
//   	containerPath: jsii.String("containerPath"),
//   	name: jsii.String("name"),
//   	readOnly: jsii.Boolean(false),
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type ScratchSpace struct {
	// The path on the container to mount the scratch volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The name of the scratch volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether to give the container read-only access to the scratch volume.
	//
	// If this value is true, the container has read-only access to the scratch volume.
	// If this value is false, then the container can write to the scratch volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
}

// A secret environment variable.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   	environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	environmentFiles: []environmentFile{
//   		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret.fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//
type Secret interface {
	// The ARN of the secret.
	Arn() *string
	// Whether this secret uses a specific JSON field.
	HasField() *bool
	// Grants reading the secret to a principal.
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

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
func Secret_FromSecretsManagerVersion(secret awssecretsmanager.ISecret, versionInfo *SecretVersionInfo, field *string) Secret {
	_init_.Initialize()

	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
		"fromSecretsManagerVersion",
		[]interface{}{secret, versionInfo, field},
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

// Specify the secret's version id or version stage.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   	environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	environmentFiles: []environmentFile{
//   		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret.fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//
type SecretVersionInfo struct {
	// version id of the secret.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// version stage of the secret.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}

// A log driver that sends log information to splunk Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   splunkLogDriver := awscdk.Aws_ecs.NewSplunkLogDriver(&splunkLogDriverProps{
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	caName: jsii.String("caName"),
//   	caPath: jsii.String("caPath"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	format: awscdk.*Aws_ecs.splunkLogFormat_INLINE,
//   	gzip: jsii.Boolean(false),
//   	gzipLevel: jsii.Number(123),
//   	index: jsii.String("index"),
//   	insecureSkipVerify: jsii.String("insecureSkipVerify"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	secretToken: secret,
//   	source: jsii.String("source"),
//   	sourceType: jsii.String("sourceType"),
//   	tag: jsii.String("tag"),
//   	verifyConnection: jsii.Boolean(false),
//   })
//
type SplunkLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		token: awscdk.SecretValue.secretsManager(jsii.String("my-splunk-token")),
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type SplunkLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Path to your Splunk Enterprise, self-service Splunk Cloud instance, or Splunk Cloud managed cluster (including port and scheme used by HTTP Event Collector) in one of the following formats: https://your_splunk_instance:8088 or https://input-prd-p-XXXXXXX.cloud.splunk.com:8088 or https://http-inputs-XXXXXXXX.splunkcloud.com.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Name to use for validating server certificate.
	CaName *string `field:"optional" json:"caName" yaml:"caName"`
	// Path to root certificate.
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// Message format.
	//
	// Can be inline, json or raw.
	Format SplunkLogFormat `field:"optional" json:"format" yaml:"format"`
	// Enable/disable gzip compression to send events to Splunk Enterprise or Splunk Cloud instance.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// Set compression level for gzip.
	//
	// Valid values are -1 (default), 0 (no compression),
	// 1 (best speed) ... 9 (best compression).
	GzipLevel *float64 `field:"optional" json:"gzipLevel" yaml:"gzipLevel"`
	// Event index.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Ignore server certificate validation.
	InsecureSkipVerify *string `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// Splunk HTTP Event Collector token (Secret).
	//
	// The splunk-token is added to the SecretOptions property of the Log Driver Configuration. So the secret value will not be
	// resolved or viewable as plain text.
	//
	// Please provide at least one of `token` or `secretToken`.
	SecretToken Secret `field:"optional" json:"secretToken" yaml:"secretToken"`
	// Event source.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Event source type.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Verify on start, that docker can connect to Splunk server.
	VerifyConnection *bool `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syslogLogDriver := awscdk.Aws_ecs.NewSyslogLogDriver(&syslogLogDriverProps{
//   	address: jsii.String("address"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	facility: jsii.String("facility"),
//   	format: jsii.String("format"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tlsCaCert: jsii.String("tlsCaCert"),
//   	tlsCert: jsii.String("tlsCert"),
//   	tlsKey: jsii.String("tlsKey"),
//   	tlsSkipVerify: jsii.Boolean(false),
//   })
//
type SyslogLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syslogLogDriverProps := &syslogLogDriverProps{
//   	address: jsii.String("address"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	facility: jsii.String("facility"),
//   	format: jsii.String("format"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tlsCaCert: jsii.String("tlsCaCert"),
//   	tlsCert: jsii.String("tlsCert"),
//   	tlsKey: jsii.String("tlsKey"),
//   	tlsSkipVerify: jsii.Boolean(false),
//   }
//
type SyslogLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The address of an external syslog server.
	//
	// The URI specifier may be
	// [tcp|udp|tcp+tls]://host:port, unix://path, or unixgram://path.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The syslog facility to use.
	//
	// Can be the number or name for any valid
	// syslog facility. See the syslog documentation:
	// https://tools.ietf.org/html/rfc5424#section-6.2.1.
	Facility *string `field:"optional" json:"facility" yaml:"facility"`
	// The syslog message format to use.
	//
	// If not specified the local UNIX syslog
	// format is used, without a specified hostname. Specify rfc3164 for the RFC-3164
	// compatible format, rfc5424 for RFC-5424 compatible format, or rfc5424micro
	// for RFC-5424 compatible format with microsecond timestamp resolution.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The absolute path to the trust certificates signed by the CA.
	//
	// Ignored
	// if the address protocol is not tcp+tls.
	TlsCaCert *string `field:"optional" json:"tlsCaCert" yaml:"tlsCaCert"`
	// The absolute path to the TLS certificate file.
	//
	// Ignored if the address
	// protocol is not tcp+tls.
	TlsCert *string `field:"optional" json:"tlsCert" yaml:"tlsCert"`
	// The absolute path to the TLS key file.
	//
	// Ignored if the address protocol
	// is not tcp+tls.
	TlsKey *string `field:"optional" json:"tlsKey" yaml:"tlsKey"`
	// If set to true, TLS verification is skipped when connecting to the syslog daemon.
	//
	// Ignored if the address protocol is not tcp+tls.
	TlsSkipVerify *bool `field:"optional" json:"tlsSkipVerify" yaml:"tlsSkipVerify"`
}

// Kernel parameters to set in the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemControl := &systemControl{
//   	namespace: jsii.String("namespace"),
//   	value: jsii.String("value"),
//   }
//
type SystemControl struct {
	// The namespaced kernel parameter for which to set a value.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The value for the namespaced kernel parameter specified in namespace.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// A special type of {@link ContainerImage} that uses an ECR repository for the image, but a CloudFormation Parameter for the tag of the image in that repository.
//
// This allows providing this tag through the Parameter at deploy time,
// for example in a CodePipeline that pushes a new tag of the image to the repository during a build step,
// and then provides that new tag through the CloudFormation Parameter in the deploy step.
//
// Example:
//   /**
//    * These are the construction properties for {@link EcsAppStack}.
//    * They extend the standard Stack properties,
//    * but also require providing the ContainerImage that the service will use.
//    * That Image will be provided from the Stack containing the CodePipeline.
//    */
//   type ecsAppStackProps struct {
//   	stackProps
//   	image containerImage
//   }
//
//   /**
//    * This is the Stack containing a simple ECS Service that uses the provided ContainerImage.
//    */
//   type EcsAppStack struct {
//   	stack
//   }
//
//   func NewEcsAppStack(scope construct, id *string, props ecsAppStackProps) *EcsAppStack {
//   	this := &EcsAppStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDefinition"), &taskDefinitionProps{
//   		compatibility: ecs.compatibility_FARGATE,
//   		cpu: jsii.String("1024"),
//   		memoryMiB: jsii.String("2048"),
//   	})
//   	taskDefinition.addContainer(jsii.String("AppContainer"), &containerDefinitionOptions{
//   		image: props.image,
//   	})
//   	ecs.NewFargateService(this, jsii.String("EcsService"), &fargateServiceProps{
//   		taskDefinition: taskDefinition,
//   		cluster: ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   			vpc: ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   				maxAzs: jsii.Number(1),
//   			}),
//   		}),
//   	})
//   	return this
//   }
//
//   /**
//    * This is the Stack containing the CodePipeline definition that deploys an ECS Service.
//    */
//   type PipelineStack struct {
//   	stack
//   	tagParameterContainerImage tagParameterContainerImage
//   }tagParameterContainerImage tagParameterContainerImage
//
//   func NewPipelineStack(scope construct, id *string, props stackProps) *PipelineStack {
//   	this := &PipelineStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	/* ********** ECS part **************** */
//
//   	// this is the ECR repository where the built Docker image will be pushed
//   	appEcrRepo := ecr.NewRepository(this, jsii.String("EcsDeployRepository"))
//   	// the build that creates the Docker image, and pushes it to the ECR repo
//   	appCodeDockerBuild := codebuild.NewPipelineProject(this, jsii.String("AppCodeDockerImageBuildAndPushProject"), &pipelineProjectProps{
//   		environment: &buildEnvironment{
//   			// we need to run Docker
//   			privileged: jsii.Boolean(true),
//   		},
//   		buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   			"phases": map[string]map[string][]*string{
//   				"build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("$(aws ecr get-login --region $AWS_DEFAULT_REGION --no-include-email)"),
//   						jsii.String("docker build -t $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION ."),
//   					},
//   				},
//   				"post_build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("docker push $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION"),
//   						jsii.String("export imageTag=$CODEBUILD_RESOLVED_SOURCE_VERSION"),
//   					},
//   				},
//   			},
//   			"env": map[string][]*string{
//   				// save the imageTag environment variable as a CodePipeline Variable
//   				"exported-variables": []*string{
//   					jsii.String("imageTag"),
//   				},
//   			},
//   		}),
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"REPOSITORY_URI": &buildEnvironmentVariable{
//   				"value": appEcrRepo.repositoryUri,
//   			},
//   		},
//   	})
//   	// needed for `docker push`
//   	appEcrRepo.grantPullPush(appCodeDockerBuild)
//   	// create the ContainerImage used for the ECS application Stack
//   	this.tagParameterContainerImage = ecs.NewTagParameterContainerImage(appEcrRepo)
//
//   	cdkCodeBuild := codebuild.NewPipelineProject(this, jsii.String("CdkCodeBuildProject"), &pipelineProjectProps{
//   		buildSpec: codebuild.*buildSpec.fromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   			"phases": map[string]map[string][]*string{
//   				"install": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("npm install"),
//   					},
//   				},
//   				"build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("npx cdk synth --verbose"),
//   					},
//   				},
//   			},
//   			"artifacts": map[string]*string{
//   				// store the entire Cloud Assembly as the output artifact
//   				"base-directory": jsii.String("cdk.out"),
//   				"files": jsii.String("**/*"),
//   			},
//   		}),
//   	})
//
//   	/* ********** Pipeline part **************** */
//
//   	appCodeSourceOutput := codepipeline.NewArtifact()
//   	cdkCodeSourceOutput := codepipeline.NewArtifact()
//   	cdkCodeBuildOutput := codepipeline.NewArtifact()
//   	appCodeBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   		actionName: jsii.String("AppCodeDockerImageBuildAndPush"),
//   		project: appCodeDockerBuild,
//   		input: appCodeSourceOutput,
//   	})
//   	codepipeline.NewPipeline(this, jsii.String("CodePipelineDeployingEcsApplication"), &pipelineProps{
//   		artifactBucket: s3.NewBucket(this, jsii.String("ArtifactBucket"), &bucketProps{
//   			removalPolicy: cdk.removalPolicy_DESTROY,
//   		}),
//   		stages: []stageProps{
//   			&stageProps{
//   				stageName: jsii.String("Source"),
//   				actions: []iAction{
//   					// this is the Action that takes the source of your application code
//   					codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   						actionName: jsii.String("AppCodeSource"),
//   						repository: codecommit.NewRepository(this, jsii.String("AppCodeSourceRepository"), &repositoryProps{
//   							repositoryName: jsii.String("AppCodeSourceRepository"),
//   						}),
//   						output: appCodeSourceOutput,
//   					}),
//   					// this is the Action that takes the source of your CDK code
//   					// (which would probably include this Pipeline code as well)
//   					codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   						actionName: jsii.String("CdkCodeSource"),
//   						repository: codecommit.NewRepository(this, jsii.String("CdkCodeSourceRepository"), &repositoryProps{
//   							repositoryName: jsii.String("CdkCodeSourceRepository"),
//   						}),
//   						output: cdkCodeSourceOutput,
//   					}),
//   				},
//   			},
//   			&stageProps{
//   				stageName: jsii.String("Build"),
//   				actions: []*iAction{
//   					appCodeBuildAction,
//   					codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   						actionName: jsii.String("CdkCodeBuildAndSynth"),
//   						project: cdkCodeBuild,
//   						input: cdkCodeSourceOutput,
//   						outputs: []artifact{
//   							cdkCodeBuildOutput,
//   						},
//   					}),
//   				},
//   			},
//   			&stageProps{
//   				stageName: jsii.String("Deploy"),
//   				actions: []*iAction{
//   					codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   						actionName: jsii.String("CFN_Deploy"),
//   						stackName: jsii.String("SampleEcsStackDeployedFromCodePipeline"),
//   						// this name has to be the same name as used below in the CDK code for the application Stack
//   						templatePath: cdkCodeBuildOutput.atPath(jsii.String("EcsStackDeployedInPipeline.template.json")),
//   						adminPermissions: jsii.Boolean(true),
//   						parameterOverrides: map[string]interface{}{
//   							// read the tag pushed to the ECR repository from the CodePipeline Variable saved by the application build step,
//   							// and pass it as the CloudFormation Parameter for the tag
//   							this.tagParameterContainerImage.tagParameterName: appCodeBuildAction.variable(jsii.String("imageTag")),
//   						},
//   					}),
//   				},
//   			},
//   		},
//   	})
//   	return this
//   }
//
//   app := cdk.NewApp()
//
//   // the CodePipeline Stack needs to be created first
//   pipelineStack := NewPipelineStack(app, jsii.String("aws-cdk-pipeline-ecs-separate-sources"))
//   // we supply the image to the ECS application Stack from the CodePipeline Stack
//   // we supply the image to the ECS application Stack from the CodePipeline Stack
//   NewEcsAppStack(app, jsii.String("EcsStackDeployedInPipeline"), &ecsAppStackProps{
//   	image: pipelineStack.tagParameterContainerImage,
//   })
//
// See: #tagParameterName.
//
type TagParameterContainerImage interface {
	ContainerImage
	// Returns the name of the CloudFormation Parameter that represents the tag of the image in the ECR repository.
	TagParameterName() *string
	// Returns the value of the CloudFormation Parameter that represents the tag of the image in the ECR repository.
	TagParameterValue() *string
	// Called when the image is used by a ContainerDefinition.
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
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type TaskDefinition interface {
	awscdk.Resource
	ITaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() Compatibility
	// The container definitions.
	Containers() *[]ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	NetworkMode() NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	// Adds an inference accelerator to the task definition.
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *Volume)
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
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) ContainerDefinition
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
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (t *jsiiProxy_TaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		t,
		"addExtension",
		[]interface{}{extension},
	)
}

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

func (t *jsiiProxy_TaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		t,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

func (t *jsiiProxy_TaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		t,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (t *jsiiProxy_TaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (t *jsiiProxy_TaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (t *jsiiProxy_TaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		t,
		"addVolume",
		[]interface{}{volume},
	)
}

func (t *jsiiProxy_TaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   taskDefinitionAttributes := &taskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	compatibility: awscdk.Aws_ecs.compatibility_EC2,
//   	networkMode: awscdk.*Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// What launch types this task definition should be compatible with.
	Compatibility Compatibility `field:"optional" json:"compatibility" yaml:"compatibility"`
}

// The properties for task definitions.
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
type TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The task launch type compatiblity requirement.
	Compatibility Compatibility `field:"required" json:"compatibility" yaml:"compatibility"`
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
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	InferenceAccelerators *[]*InferenceAccelerator `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	IpcMode IpcMode `field:"optional" json:"ipcMode" yaml:"ipcMode"`
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
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU).
	MemoryMiB *string `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The networking mode to use for the containers in the task.
	//
	// On Fargate, the only supported networking mode is AwsVpc.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	PidMode PidMode `field:"optional" json:"pidMode" yaml:"pidMode"`
	// The placement constraints to use for tasks in the service.
	//
	// You can specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	//
	// Not supported in Fargate.
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The operating system that your task definitions are running on.
	//
	// A runtimePlatform is supported only for tasks using the Fargate launch type.
	RuntimePlatform *RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
}

// The details of a tmpfs mount for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfs := &tmpfs{
//   	containerPath: jsii.String("containerPath"),
//   	size: jsii.Number(123),
//
//   	// the properties below are optional
//   	mountOptions: []tmpfsMountOption{
//   		awscdk.Aws_ecs.*tmpfsMountOption_DEFAULTS,
//   	},
//   }
//
type Tmpfs struct {
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The size (in MiB) of the tmpfs volume.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The list of tmpfs volume mount options.
	//
	// For more information, see
	// [TmpfsMountOptions](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tmpfs.html).
	MountOptions *[]TmpfsMountOption `field:"optional" json:"mountOptions" yaml:"mountOptions"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric metric
//
//   trackCustomMetricProps := &trackCustomMetricProps{
//   	metric: metric,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	scaleInCooldown: cdk.duration.minutes(jsii.Number(30)),
//   	scaleOutCooldown: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type TrackCustomMetricProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The custom CloudWatch metric to track.
	//
	// The metric must represent utilization; that is, you will always get the following behavior:
	//
	// - metric > targetValue => scale out
	// - metric < targetValue => scale in.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The target value for the custom CloudWatch metric.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

// The ulimit settings to pass to the container.
//
// NOTE: Does not work for Windows containers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimit := &ulimit{
//   	hardLimit: jsii.Number(123),
//   	name: awscdk.Aws_ecs.ulimitName_CORE,
//   	softLimit: jsii.Number(123),
//   }
//
type Ulimit struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The type of the ulimit.
	//
	// For more information, see [UlimitName](https://docs.aws.amazon.com/cdk/api/latest/typescript/api/aws-ecs/ulimitname.html#aws_ecs_UlimitName).
	Name UlimitName `field:"required" json:"name" yaml:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
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
// Example:
//   fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	memoryLimitMiB: jsii.Number(512),
//   	cpu: jsii.Number(256),
//   })
//   volume := map[string]interface{}{
//   	// Use an Elastic FileSystem
//   	"name": jsii.String("mydatavolume"),
//   	"efsVolumeConfiguration": map[string]*string{
//   		"fileSystemId": jsii.String("EFS"),
//   	},
//   }
//
//   container := fargateTaskDefinition.addVolume(volume)
//
type Volume struct {
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.
	// This name is referenced in the sourceVolume parameter of container definition mountPoints.
	Name *string `field:"required" json:"name" yaml:"name"`
	// This property is specified when you are using Docker volumes.
	//
	// Docker volumes are only supported when you are using the EC2 launch type.
	// Windows containers only support the use of the local driver.
	// To use bind mounts, specify a host instead.
	DockerVolumeConfiguration *DockerVolumeConfiguration `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// This property is specified when you are using Amazon EFS.
	//
	// When specifying Amazon EFS volumes in tasks using the Fargate launch type,
	// Fargate creates a supervisor container that is responsible for managing the Amazon EFS volume.
	// The supervisor container uses a small amount of the task's memory.
	// The supervisor container is visible when querying the task metadata version 4 endpoint,
	// but is not visible in CloudWatch Container Insights.
	EfsVolumeConfiguration *EfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// This property is specified when you are using bind mount host volumes.
	//
	// Bind mount host volumes are supported when you are using either the EC2 or Fargate launch types.
	// The contents of the host parameter determine whether your bind mount host volume persists on the
	// host container instance and where it is stored. If the host parameter is empty, then the Docker
	// daemon assigns a host path for your data volume. However, the data is not guaranteed to persist
	// after the containers associated with it stop running.
	Host *Host `field:"optional" json:"host" yaml:"host"`
}

// The details on a data volume from another container in the same task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeFrom := &volumeFrom{
//   	readOnly: jsii.Boolean(false),
//   	sourceContainer: jsii.String("sourceContainer"),
//   }
//
type VolumeFrom struct {
	// Specifies whether the container has read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition from which to mount volumes.
	SourceContainer *string `field:"required" json:"sourceContainer" yaml:"sourceContainer"`
}

// ECS-optimized Windows version list.
type WindowsOptimizedVersion string

const (
	WindowsOptimizedVersion_SERVER_2019 WindowsOptimizedVersion = "SERVER_2019"
	WindowsOptimizedVersion_SERVER_2016 WindowsOptimizedVersion = "SERVER_2016"
)


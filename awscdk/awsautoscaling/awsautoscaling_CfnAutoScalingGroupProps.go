package awsautoscaling


// Properties for defining a `CfnAutoScalingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingGroupProps := &cfnAutoScalingGroupProps{
//   	maxSize: jsii.String("maxSize"),
//   	minSize: jsii.String("minSize"),
//
//   	// the properties below are optional
//   	autoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	capacityRebalance: jsii.Boolean(false),
//   	context: jsii.String("context"),
//   	cooldown: jsii.String("cooldown"),
//   	defaultInstanceWarmup: jsii.Number(123),
//   	desiredCapacity: jsii.String("desiredCapacity"),
//   	desiredCapacityType: jsii.String("desiredCapacityType"),
//   	healthCheckGracePeriod: jsii.Number(123),
//   	healthCheckType: jsii.String("healthCheckType"),
//   	instanceId: jsii.String("instanceId"),
//   	launchConfigurationName: jsii.String("launchConfigurationName"),
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	lifecycleHookSpecificationList: []interface{}{
//   		&lifecycleHookSpecificationProperty{
//   			lifecycleHookName: jsii.String("lifecycleHookName"),
//   			lifecycleTransition: jsii.String("lifecycleTransition"),
//
//   			// the properties below are optional
//   			defaultResult: jsii.String("defaultResult"),
//   			heartbeatTimeout: jsii.Number(123),
//   			notificationMetadata: jsii.String("notificationMetadata"),
//   			notificationTargetArn: jsii.String("notificationTargetArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	loadBalancerNames: []*string{
//   		jsii.String("loadBalancerNames"),
//   	},
//   	maxInstanceLifetime: jsii.Number(123),
//   	metricsCollection: []interface{}{
//   		&metricsCollectionProperty{
//   			granularity: jsii.String("granularity"),
//
//   			// the properties below are optional
//   			metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   		},
//   	},
//   	mixedInstancesPolicy: &mixedInstancesPolicyProperty{
//   		launchTemplate: &launchTemplateProperty{
//   			launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   				version: jsii.String("version"),
//
//   				// the properties below are optional
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   			},
//
//   			// the properties below are optional
//   			overrides: []interface{}{
//   				&launchTemplateOverridesProperty{
//   					instanceRequirements: &instanceRequirementsProperty{
//   						acceleratorCount: &acceleratorCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						acceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						allowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   						bareMetal: jsii.String("bareMetal"),
//   						baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						burstablePerformance: jsii.String("burstablePerformance"),
//   						cpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						excludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						instanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						localStorage: jsii.String("localStorage"),
//   						localStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						memoryMiB: &memoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						requireHibernateSupport: jsii.Boolean(false),
//   						spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						vCpuCount: &vCpuCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   						version: jsii.String("version"),
//
//   						// the properties below are optional
//   						launchTemplateId: jsii.String("launchTemplateId"),
//   						launchTemplateName: jsii.String("launchTemplateName"),
//   					},
//   					weightedCapacity: jsii.String("weightedCapacity"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		instancesDistribution: &instancesDistributionProperty{
//   			onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   			onDemandBaseCapacity: jsii.Number(123),
//   			onDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   			spotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   			spotInstancePools: jsii.Number(123),
//   			spotMaxPrice: jsii.String("spotMaxPrice"),
//   		},
//   	},
//   	newInstancesProtectedFromScaleIn: jsii.Boolean(false),
//   	notificationConfigurations: []interface{}{
//   		&notificationConfigurationProperty{
//   			topicArn: jsii.String("topicArn"),
//
//   			// the properties below are optional
//   			notificationTypes: []*string{
//   				jsii.String("notificationTypes"),
//   			},
//   		},
//   	},
//   	placementGroup: jsii.String("placementGroup"),
//   	serviceLinkedRoleArn: jsii.String("serviceLinkedRoleArn"),
//   	tags: []tagPropertyProperty{
//   		&tagPropertyProperty{
//   			key: jsii.String("key"),
//   			propagateAtLaunch: jsii.Boolean(false),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   	terminationPolicies: []*string{
//   		jsii.String("terminationPolicies"),
//   	},
//   	vpcZoneIdentifier: []*string{
//   		jsii.String("vpcZoneIdentifier"),
//   	},
//   }
//
type CfnAutoScalingGroupProps struct {
	// The maximum size of the group.
	//
	// > With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to go above `MaxSize` to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will never go above `MaxSize` by more than your largest instance weight (weights that define how many units each instance contributes to the desired capacity of the group).
	MaxSize *string `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum size of the group.
	MinSize *string `field:"required" json:"minSize" yaml:"minSize"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// A list of Availability Zones where instances in the Auto Scaling group can be created.
	//
	// Used for launching into EC2-Classic or the default VPC subnet in each Availability Zone when not using the `VPCZoneIdentifier` property, or for attaching a network interface when an existing network interface ID is specified in a launch template.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Indicates whether Capacity Rebalancing is enabled.
	//
	// Otherwise, Capacity Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of interruption. After launching a new instance, it then terminates an old instance. For more information, see [Use Capacity Rebalancing to handle Amazon EC2 Spot Interruptions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the in the *Amazon EC2 Auto Scaling User Guide* .
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
	// Reserved.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// *Only needed if you use simple scaling policies.*.
	//
	// The amount of time, in seconds, between one scaling activity ending and another one starting due to simple scaling policies. For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: `300` seconds.
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Not currently supported by CloudFormation.
	DefaultInstanceWarmup *float64 `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain.
	//
	// It can scale beyond this capacity if you configure automatic scaling.
	//
	// The number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity when creating the stack, the default is the minimum size of the group.
	//
	// CloudFormation marks the Auto Scaling group as successful (by setting its status to CREATE_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set in the launch template or launch configuration that you specified, then desired capacity is not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance capacity and your maximum price.
	DesiredCapacity *string `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The unit of measurement for the value specified for desired capacity.
	//
	// Amazon EC2 Auto Scaling supports `DesiredCapacityType` for attribute-based instance type selection only. For more information, see [Creating an Auto Scaling group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// By default, Amazon EC2 Auto Scaling specifies `units` , which translates into number of instances.
	//
	// Valid values: `units` | `vcpu` | `memory-mib`.
	DesiredCapacityType *string `field:"optional" json:"desiredCapacityType" yaml:"desiredCapacityType"`
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed Elastic Load Balancing or custom health check.
	//
	// This is useful if your instances do not immediately pass these health checks after they enter the `InService` state. For more information, see [Health check grace period](https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: `0` seconds.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The service to use for the health checks.
	//
	// The valid values are `EC2` (default) and `ELB` . If you configure an Auto Scaling group to use load balancer (ELB) health checks, it considers the instance unhealthy if it fails either the EC2 status checks or the load balancer health checks. For more information, see [Health checks for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the *Amazon EC2 Auto Scaling User Guide* .
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// The ID of the instance used to base the launch configuration on.
	//
	// For more information, see [Create an Auto Scaling group using an EC2 instance](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify `LaunchTemplate` , `MixedInstancesPolicy` , or `LaunchConfigurationName` , don't specify `InstanceId` .
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The name of the launch configuration to use to launch instances.
	//
	// Required only if you don't specify `LaunchTemplate` , `MixedInstancesPolicy` , or `InstanceId` .
	LaunchConfigurationName *string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// Information used to specify the launch template and version to use to launch instances.
	//
	// You can alternatively associate a launch template to the Auto Scaling group by specifying a `MixedInstancesPolicy` . For more information about creating launch templates, see [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you omit this property, you must specify `MixedInstancesPolicy` , `LaunchConfigurationName` , or `InstanceId` .
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// One or more lifecycle hooks to add to the Auto Scaling group before instances are launched.
	LifecycleHookSpecificationList interface{} `field:"optional" json:"lifecycleHookSpecificationList" yaml:"lifecycleHookSpecificationList"`
	// A list of Classic Load Balancers associated with this Auto Scaling group.
	//
	// For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers, specify the `TargetGroupARNs` property instead.
	LoadBalancerNames *[]*string `field:"optional" json:"loadBalancerNames" yaml:"loadBalancerNames"`
	// The maximum amount of time, in seconds, that an instance can be in service.
	//
	// The default is null. If specified, the value must be either 0 or a number equal to or greater than 86,400 seconds (1 day). For more information, see [Replacing Auto Scaling instances based on maximum instance lifetime](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html) in the *Amazon EC2 Auto Scaling User Guide* .
	MaxInstanceLifetime *float64 `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Enables the monitoring of group metrics of an Auto Scaling group.
	//
	// By default, these metrics are disabled.
	MetricsCollection interface{} `field:"optional" json:"metricsCollection" yaml:"metricsCollection"`
	// An embedded object that specifies a mixed instances policy.
	//
	// The policy includes properties that not only define the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but also the properties that specify the instance configuration information—the launch template and instance types. The policy can also include a weight for each instance type and different launch templates for individual instance types.
	//
	// For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
	MixedInstancesPolicy interface{} `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// For more information about preventing instances from terminating on scale in, see [Using instance scale-in protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html) in the *Amazon EC2 Auto Scaling User Guide* .
	NewInstancesProtectedFromScaleIn interface{} `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Configures an Auto Scaling group to send notifications when specified events take place.
	NotificationConfigurations interface{} `field:"optional" json:"notificationConfigurations" yaml:"notificationConfigurations"`
	// The name of the placement group into which to launch your instances.
	//
	// For more information, see [Placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// > A *cluster* placement group is a logical grouping of instances within a single Availability Zone. You cannot specify multiple Availability Zones and a cluster placement group.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS service on your behalf.
	//
	// By default, Amazon EC2 Auto Scaling uses a service-linked role named `AWSServiceRoleForAutoScaling` , which it creates if it does not exist. For more information, see [Service-linked roles](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	ServiceLinkedRoleArn *string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// One or more tags.
	//
	// You can tag your Auto Scaling group and propagate the tags to the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS volumes, specify the tags in a launch template but use caution. If the launch template specifies an instance tag with a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling overrides the value of that instance tag with the value specified by the Auto Scaling group. For more information, see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide* .
	Tags *[]*CfnAutoScalingGroup_TagPropertyProperty `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Names (ARN) of the target groups to associate with the Auto Scaling group.
	//
	// Instances are registered as targets in a target group, and traffic is routed to the target group. For more information, see [Elastic Load Balancing and Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html) in the *Amazon EC2 Auto Scaling User Guide* .
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// A policy or a list of policies that are used to select the instance to terminate.
	//
	// These policies are executed in the order that you list them. For more information, see [Work with Amazon EC2 Auto Scaling termination policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid values: `Default` | `AllocationStrategy` | `ClosestToNextInstanceHour` | `NewestInstance` | `OldestInstance` | `OldestLaunchConfiguration` | `OldestLaunchTemplate` | `arn:aws:lambda:region:account-id:function:my-function:my-alias`.
	TerminationPolicies *[]*string `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.
	//
	// If you specify `VPCZoneIdentifier` with `AvailabilityZones` , the subnets that you specify for this property must reside in those Availability Zones.
	//
	// If this resource specifies public subnets and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html) .
	//
	// Conditional: If your account supports EC2-Classic and VPC, this property is required to launch instances into a VPC.
	//
	// > When you update `VPCZoneIdentifier` , this retains the same Auto Scaling group and replaces old instances with new ones, according to the specified subnets. You can optionally specify how CloudFormation handles these updates by using an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) .
	VpcZoneIdentifier *[]*string `field:"optional" json:"vpcZoneIdentifier" yaml:"vpcZoneIdentifier"`
}


package awsautoscaling


// Properties for defining a `CfnAutoScalingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingGroupProps := &CfnAutoScalingGroupProps{
//   	MaxSize: jsii.String("maxSize"),
//   	MinSize: jsii.String("minSize"),
//
//   	// the properties below are optional
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	CapacityRebalance: jsii.Boolean(false),
//   	Context: jsii.String("context"),
//   	Cooldown: jsii.String("cooldown"),
//   	DefaultInstanceWarmup: jsii.Number(123),
//   	DesiredCapacity: jsii.String("desiredCapacity"),
//   	DesiredCapacityType: jsii.String("desiredCapacityType"),
//   	HealthCheckGracePeriod: jsii.Number(123),
//   	HealthCheckType: jsii.String("healthCheckType"),
//   	InstanceId: jsii.String("instanceId"),
//   	InstanceMaintenancePolicy: &InstanceMaintenancePolicyProperty{
//   		MaxHealthyPercentage: jsii.Number(123),
//   		MinHealthyPercentage: jsii.Number(123),
//   	},
//   	LaunchConfigurationName: jsii.String("launchConfigurationName"),
//   	LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	LifecycleHookSpecificationList: []interface{}{
//   		&LifecycleHookSpecificationProperty{
//   			LifecycleHookName: jsii.String("lifecycleHookName"),
//   			LifecycleTransition: jsii.String("lifecycleTransition"),
//
//   			// the properties below are optional
//   			DefaultResult: jsii.String("defaultResult"),
//   			HeartbeatTimeout: jsii.Number(123),
//   			NotificationMetadata: jsii.String("notificationMetadata"),
//   			NotificationTargetArn: jsii.String("notificationTargetArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	LoadBalancerNames: []*string{
//   		jsii.String("loadBalancerNames"),
//   	},
//   	MaxInstanceLifetime: jsii.Number(123),
//   	MetricsCollection: []interface{}{
//   		&MetricsCollectionProperty{
//   			Granularity: jsii.String("granularity"),
//
//   			// the properties below are optional
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   		},
//   	},
//   	MixedInstancesPolicy: &MixedInstancesPolicyProperty{
//   		LaunchTemplate: &LaunchTemplateProperty{
//   			LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   				Version: jsii.String("version"),
//
//   				// the properties below are optional
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   			},
//
//   			// the properties below are optional
//   			Overrides: []interface{}{
//   				&LaunchTemplateOverridesProperty{
//   					InstanceRequirements: &InstanceRequirementsProperty{
//   						MemoryMiB: &MemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						VCpuCount: &VCpuCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						AcceleratorCount: &AcceleratorCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						AcceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						AllowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   						BareMetal: jsii.String("bareMetal"),
//   						BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						BurstablePerformance: jsii.String("burstablePerformance"),
//   						CpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						ExcludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						InstanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						LocalStorage: jsii.String("localStorage"),
//   						LocalStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   						MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						RequireHibernateSupport: jsii.Boolean(false),
//   						SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   						Version: jsii.String("version"),
//
//   						// the properties below are optional
//   						LaunchTemplateId: jsii.String("launchTemplateId"),
//   						LaunchTemplateName: jsii.String("launchTemplateName"),
//   					},
//   					WeightedCapacity: jsii.String("weightedCapacity"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		InstancesDistribution: &InstancesDistributionProperty{
//   			OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   			OnDemandBaseCapacity: jsii.Number(123),
//   			OnDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   			SpotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   			SpotInstancePools: jsii.Number(123),
//   			SpotMaxPrice: jsii.String("spotMaxPrice"),
//   		},
//   	},
//   	NewInstancesProtectedFromScaleIn: jsii.Boolean(false),
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		TopicArn: jsii.String("topicArn"),
//
//   		// the properties below are optional
//   		NotificationTypes: []*string{
//   			jsii.String("notificationTypes"),
//   		},
//   	},
//   	NotificationConfigurations: []interface{}{
//   		&NotificationConfigurationProperty{
//   			TopicArn: jsii.String("topicArn"),
//
//   			// the properties below are optional
//   			NotificationTypes: []*string{
//   				jsii.String("notificationTypes"),
//   			},
//   		},
//   	},
//   	PlacementGroup: jsii.String("placementGroup"),
//   	ServiceLinkedRoleArn: jsii.String("serviceLinkedRoleArn"),
//   	Tags: []tagPropertyProperty{
//   		&tagPropertyProperty{
//   			Key: jsii.String("key"),
//   			PropagateAtLaunch: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   	TerminationPolicies: []*string{
//   		jsii.String("terminationPolicies"),
//   	},
//   	VpcZoneIdentifier: []*string{
//   		jsii.String("vpcZoneIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html
//
type CfnAutoScalingGroupProps struct {
	// The maximum size of the group.
	//
	// > With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to go above `MaxSize` to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will never go above `MaxSize` by more than your largest instance weight (weights that define how many units each instance contributes to the desired capacity of the group).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-maxsize
	//
	MaxSize *string `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum size of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-minsize
	//
	MinSize *string `field:"required" json:"minSize" yaml:"minSize"`
	// The name of the Auto Scaling group. This name must be unique per Region per account.
	//
	// The name can contain any ASCII character 33 to 126 including most punctuation characters, digits, and upper and lowercased letters.
	//
	// > You cannot use a colon (:) in the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-autoscalinggroupname
	//
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// A list of Availability Zones where instances in the Auto Scaling group can be created.
	//
	// Used for launching into the default VPC subnet in each Availability Zone when not using the `VPCZoneIdentifier` property, or for attaching a network interface when an existing network interface ID is specified in a launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Indicates whether Capacity Rebalancing is enabled.
	//
	// Otherwise, Capacity Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of interruption. After launching a new instance, it then terminates an old instance. For more information, see [Use Capacity Rebalancing to handle Amazon EC2 Spot Interruptions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-capacityrebalance
	//
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
	// Reserved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-context
	//
	Context *string `field:"optional" json:"context" yaml:"context"`
	// *Only needed if you use simple scaling policies.*.
	//
	// The amount of time, in seconds, between one scaling activity ending and another one starting due to simple scaling policies. For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: `300` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-cooldown
	//
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the `InService` state.
	//
	// During an instance refresh, Amazon EC2 Auto Scaling waits for the warm-up period after it replaces an instance before it moves on to replacing the next instance. Amazon EC2 Auto Scaling also waits for the warm-up period before aggregating the metrics for new instances with existing instances in the Amazon CloudWatch metrics that are used for scaling, resulting in more reliable usage data. For more information, see [Set the default instance warmup for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// > To manage various warm-up settings at the group level, we recommend that you set the default instance warmup, *even if it is set to 0 seconds* . To remove a value that you previously set, include the property but specify `-1` for the value. However, we strongly recommend keeping the default instance warmup enabled by specifying a value of `0` or other nominal value.
	//
	// Default: None.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-defaultinstancewarmup
	//
	DefaultInstanceWarmup *float64 `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain.
	//
	// It can scale beyond this capacity if you configure automatic scaling.
	//
	// The number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity when creating the stack, the default is the minimum size of the group.
	//
	// CloudFormation marks the Auto Scaling group as successful (by setting its status to CREATE_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set in the launch template or launch configuration that you specified, then desired capacity is not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance capacity and your maximum price.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-desiredcapacity
	//
	DesiredCapacity *string `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The unit of measurement for the value specified for desired capacity.
	//
	// Amazon EC2 Auto Scaling supports `DesiredCapacityType` for attribute-based instance type selection only. For more information, see [Creating an Auto Scaling group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// By default, Amazon EC2 Auto Scaling specifies `units` , which translates into number of instances.
	//
	// Valid values: `units` | `vcpu` | `memory-mib`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-desiredcapacitytype
	//
	DesiredCapacityType *string `field:"optional" json:"desiredCapacityType" yaml:"desiredCapacityType"`
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.
	//
	// This is useful if your instances do not immediately pass their health checks after they enter the `InService` state. For more information, see [Set the health check grace period for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: `0` seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-healthcheckgraceperiod
	//
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// A comma-separated value string of one or more health check types.
	//
	// The valid values are `EC2` , `ELB` , and `VPC_LATTICE` . `EC2` is the default health check and cannot be disabled. For more information, see [Health checks for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Only specify `EC2` if you must clear a value that was previously set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-healthchecktype
	//
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// The ID of the instance used to base the launch configuration on.
	//
	// For more information, see [Create an Auto Scaling group using an EC2 instance](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify `LaunchTemplate` , `MixedInstancesPolicy` , or `LaunchConfigurationName` , don't specify `InstanceId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// An instance maintenance policy.
	//
	// For more information, see [Set instance maintenance policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-instancemaintenancepolicy
	//
	InstanceMaintenancePolicy interface{} `field:"optional" json:"instanceMaintenancePolicy" yaml:"instanceMaintenancePolicy"`
	// The name of the launch configuration to use to launch instances.
	//
	// Required only if you don't specify `LaunchTemplate` , `MixedInstancesPolicy` , or `InstanceId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-launchconfigurationname
	//
	LaunchConfigurationName *string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// Information used to specify the launch template and version to use to launch instances.
	//
	// You can alternatively associate a launch template to the Auto Scaling group by specifying a `MixedInstancesPolicy` . For more information about creating launch templates, see [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you omit this property, you must specify `MixedInstancesPolicy` , `LaunchConfigurationName` , or `InstanceId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-launchtemplate
	//
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// One or more lifecycle hooks to add to the Auto Scaling group before instances are launched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-lifecyclehookspecificationlist
	//
	LifecycleHookSpecificationList interface{} `field:"optional" json:"lifecycleHookSpecificationList" yaml:"lifecycleHookSpecificationList"`
	// A list of Classic Load Balancers associated with this Auto Scaling group.
	//
	// For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers, specify the `TargetGroupARNs` property instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-loadbalancernames
	//
	LoadBalancerNames *[]*string `field:"optional" json:"loadBalancerNames" yaml:"loadBalancerNames"`
	// The maximum amount of time, in seconds, that an instance can be in service.
	//
	// The default is null. If specified, the value must be either 0 or a number equal to or greater than 86,400 seconds (1 day). For more information, see [Replacing Auto Scaling instances based on maximum instance lifetime](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-maxinstancelifetime
	//
	MaxInstanceLifetime *float64 `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Enables the monitoring of group metrics of an Auto Scaling group.
	//
	// By default, these metrics are disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-metricscollection
	//
	MetricsCollection interface{} `field:"optional" json:"metricsCollection" yaml:"metricsCollection"`
	// An embedded object that specifies a mixed instances policy.
	//
	// The policy includes properties that not only define the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but also the properties that specify the instance configuration information—the launch template and instance types. The policy can also include a weight for each instance type and different launch templates for individual instance types.
	//
	// For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-mixedinstancespolicy
	//
	MixedInstancesPolicy interface{} `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// For more information about preventing instances from terminating on scale in, see [Using instance scale-in protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-newinstancesprotectedfromscalein
	//
	NewInstancesProtectedFromScaleIn interface{} `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// A structure that specifies an Amazon SNS notification configuration for the ``NotificationConfigurations`` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.  For an example template snippet, see [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html).  For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-notificationconfiguration
	//
	// Deprecated: this property has been deprecated.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Configures an Auto Scaling group to send notifications when specified events take place.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-notificationconfigurations
	//
	NotificationConfigurations interface{} `field:"optional" json:"notificationConfigurations" yaml:"notificationConfigurations"`
	// The name of the placement group into which to launch your instances.
	//
	// For more information, see [Placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// > A *cluster* placement group is a logical grouping of instances within a single Availability Zone. You cannot specify multiple Availability Zones and a cluster placement group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-placementgroup
	//
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS service on your behalf.
	//
	// By default, Amazon EC2 Auto Scaling uses a service-linked role named `AWSServiceRoleForAutoScaling` , which it creates if it does not exist. For more information, see [Service-linked roles](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-servicelinkedrolearn
	//
	ServiceLinkedRoleArn *string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// One or more tags.
	//
	// You can tag your Auto Scaling group and propagate the tags to the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS volumes, specify the tags in a launch template but use caution. If the launch template specifies an instance tag with a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling overrides the value of that instance tag with the value specified by the Auto Scaling group. For more information, see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-tags
	//
	Tags *[]*CfnAutoScalingGroup_TagPropertyProperty `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling group.
	//
	// Instances are registered as targets with the target groups. The target groups receive incoming traffic and route requests to one or more registered targets. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-targetgrouparns
	//
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// A policy or a list of policies that are used to select the instance to terminate.
	//
	// These policies are executed in the order that you list them. For more information, see [Work with Amazon EC2 Auto Scaling termination policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid values: `Default` | `AllocationStrategy` | `ClosestToNextInstanceHour` | `NewestInstance` | `OldestInstance` | `OldestLaunchConfiguration` | `OldestLaunchTemplate` | `arn:aws:lambda:region:account-id:function:my-function:my-alias`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-terminationpolicies
	//
	TerminationPolicies *[]*string `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.
	//
	// If this resource specifies public subnets and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html) .
	//
	// > When you update `VPCZoneIdentifier` , this retains the same Auto Scaling group and replaces old instances with new ones, according to the specified subnets. You can optionally specify how CloudFormation handles these updates by using an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) .
	//
	// Required to launch instances into a nondefault VPC. If you specify `VPCZoneIdentifier` with `AvailabilityZones` , the subnets that you specify for this property must reside in those Availability Zones.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#cfn-autoscaling-autoscalinggroup-vpczoneidentifier
	//
	VpcZoneIdentifier *[]*string `field:"optional" json:"vpcZoneIdentifier" yaml:"vpcZoneIdentifier"`
}


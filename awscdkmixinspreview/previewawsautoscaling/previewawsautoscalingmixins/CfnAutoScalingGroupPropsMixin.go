package previewawsautoscalingmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsautoscaling/previewawsautoscalingmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScaling::AutoScalingGroup` resource defines an Amazon EC2 Auto Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical grouping for the purposes of automatic scaling and management.
//
// For more information about Amazon EC2 Auto Scaling, see the [Amazon EC2 Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html) .
//
// > Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use launch configurations. For more information, see [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html) in the *Amazon EC2 Auto Scaling User Guide* .
// >
// > For help migrating from launch configurations to launch templates, see [Migrate AWS CloudFormation stacks from launch configurations to launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutoScalingGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnAutoScalingGroupPropsMixin(&CfnAutoScalingGroupMixinProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	AvailabilityZoneDistribution: &AvailabilityZoneDistributionProperty{
//   		CapacityDistributionStrategy: jsii.String("capacityDistributionStrategy"),
//   	},
//   	AvailabilityZoneImpairmentPolicy: &AvailabilityZoneImpairmentPolicyProperty{
//   		ImpairedZoneHealthCheckBehavior: jsii.String("impairedZoneHealthCheckBehavior"),
//   		ZonalShiftEnabled: jsii.Boolean(false),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	CapacityRebalance: jsii.Boolean(false),
//   	CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   		CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		CapacityReservationTarget: &CapacityReservationTargetProperty{
//   			CapacityReservationIds: []*string{
//   				jsii.String("capacityReservationIds"),
//   			},
//   			CapacityReservationResourceGroupArns: []*string{
//   				jsii.String("capacityReservationResourceGroupArns"),
//   			},
//   		},
//   	},
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
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   		Version: jsii.String("version"),
//   	},
//   	LifecycleHookSpecificationList: []interface{}{
//   		&LifecycleHookSpecificationProperty{
//   			DefaultResult: jsii.String("defaultResult"),
//   			HeartbeatTimeout: jsii.Number(123),
//   			LifecycleHookName: jsii.String("lifecycleHookName"),
//   			LifecycleTransition: jsii.String("lifecycleTransition"),
//   			NotificationMetadata: jsii.String("notificationMetadata"),
//   			NotificationTargetArn: jsii.String("notificationTargetArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	LoadBalancerNames: []*string{
//   		jsii.String("loadBalancerNames"),
//   	},
//   	MaxInstanceLifetime: jsii.Number(123),
//   	MaxSize: jsii.String("maxSize"),
//   	MetricsCollection: []interface{}{
//   		&MetricsCollectionProperty{
//   			Granularity: jsii.String("granularity"),
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   		},
//   	},
//   	MinSize: jsii.String("minSize"),
//   	MixedInstancesPolicy: &MixedInstancesPolicyProperty{
//   		InstancesDistribution: &InstancesDistributionProperty{
//   			OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   			OnDemandBaseCapacity: jsii.Number(123),
//   			OnDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   			SpotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   			SpotInstancePools: jsii.Number(123),
//   			SpotMaxPrice: jsii.String("spotMaxPrice"),
//   		},
//   		LaunchTemplate: &LaunchTemplateProperty{
//   			LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   				Version: jsii.String("version"),
//   			},
//   			Overrides: []interface{}{
//   				&LaunchTemplateOverridesProperty{
//   					InstanceRequirements: &InstanceRequirementsProperty{
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
//   						BaselinePerformanceFactors: &BaselinePerformanceFactorsRequestProperty{
//   							Cpu: &CpuPerformanceFactorRequestProperty{
//   								References: []interface{}{
//   									&PerformanceFactorReferenceRequestProperty{
//   										InstanceFamily: jsii.String("instanceFamily"),
//   									},
//   								},
//   							},
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
//   						MemoryMiB: &MemoryMiBRequestProperty{
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
//   						VCpuCount: &VCpuCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   						LaunchTemplateId: jsii.String("launchTemplateId"),
//   						LaunchTemplateName: jsii.String("launchTemplateName"),
//   						Version: jsii.String("version"),
//   					},
//   					WeightedCapacity: jsii.String("weightedCapacity"),
//   				},
//   			},
//   		},
//   	},
//   	NewInstancesProtectedFromScaleIn: jsii.Boolean(false),
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		NotificationTypes: []*string{
//   			jsii.String("notificationTypes"),
//   		},
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	NotificationConfigurations: []interface{}{
//   		&NotificationConfigurationProperty{
//   			NotificationTypes: []*string{
//   				jsii.String("notificationTypes"),
//   			},
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   	PlacementGroup: jsii.String("placementGroup"),
//   	ServiceLinkedRoleArn: jsii.String("serviceLinkedRoleArn"),
//   	SkipZonalShiftValidation: jsii.Boolean(false),
//   	Tags: []TagPropertyProperty{
//   		&TagPropertyProperty{
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
//   	TrafficSources: []interface{}{
//   		&TrafficSourceIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	VpcZoneIdentifier: []*string{
//   		jsii.String("vpcZoneIdentifier"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html
//
type CfnAutoScalingGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAutoScalingGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutoScalingGroupPropsMixin
type jsiiProxy_CfnAutoScalingGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAutoScalingGroupPropsMixin) Props() *CfnAutoScalingGroupMixinProps {
	var returns *CfnAutoScalingGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AutoScaling::AutoScalingGroup`.
func NewCfnAutoScalingGroupPropsMixin(props *CfnAutoScalingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutoScalingGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutoScalingGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutoScalingGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnAutoScalingGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AutoScaling::AutoScalingGroup`.
func NewCfnAutoScalingGroupPropsMixin_Override(c CfnAutoScalingGroupPropsMixin, props *CfnAutoScalingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnAutoScalingGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAutoScalingGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoScalingGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnAutoScalingGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutoScalingGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnAutoScalingGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


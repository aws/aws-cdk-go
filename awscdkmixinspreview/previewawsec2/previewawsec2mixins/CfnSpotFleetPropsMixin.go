package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Spot Fleet request.
//
// The Spot Fleet request specifies the total target capacity and the On-Demand target capacity. Amazon EC2 calculates the difference between the total capacity and On-Demand capacity, and launches the difference as Spot capacity.
//
// You can submit a single request that includes multiple launch specifications that vary by instance type, AMI, Availability Zone, or subnet.
//
// By default, the Spot Fleet requests Spot Instances in the Spot Instance pool where the price per unit is the lowest. Each launch specification can include its own instance weighting that reflects the value of the instance type to your application workload.
//
// Alternatively, you can specify that the Spot Fleet distribute the target capacity across the Spot pools included in its launch specifications. By ensuring that the Spot Instances in your Spot Fleet are in different Spot pools, you can improve the availability of your fleet.
//
// You can specify tags for the Spot Fleet request and instances launched by the fleet. You cannot tag other resource types in a Spot Fleet request because only the `spot-fleet-request` and `instance` resource types are supported.
//
// For more information, see [Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSpotFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnSpotFleetPropsMixin(&CfnSpotFleetMixinProps{
//   	SpotFleetRequestConfigData: &SpotFleetRequestConfigDataProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		Context: jsii.String("context"),
//   		ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   		IamFleetRole: jsii.String("iamFleetRole"),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		InstancePoolsToUseCount: jsii.Number(123),
//   		LaunchSpecifications: []interface{}{
//   			&SpotFleetLaunchSpecificationProperty{
//   				BlockDeviceMappings: []interface{}{
//   					&BlockDeviceMappingProperty{
//   						DeviceName: jsii.String("deviceName"),
//   						Ebs: &EbsBlockDeviceProperty{
//   							DeleteOnTermination: jsii.Boolean(false),
//   							Encrypted: jsii.Boolean(false),
//   							Iops: jsii.Number(123),
//   							SnapshotId: jsii.String("snapshotId"),
//   							VolumeSize: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						NoDevice: jsii.String("noDevice"),
//   						VirtualName: jsii.String("virtualName"),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   				IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   				ImageId: jsii.String("imageId"),
//   				InstanceRequirements: &InstanceRequirementsRequestProperty{
//   					AcceleratorCount: &AcceleratorCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					AcceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					AllowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					BareMetal: jsii.String("bareMetal"),
//   					BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					BaselinePerformanceFactors: &BaselinePerformanceFactorsRequestProperty{
//   						Cpu: &CpuPerformanceFactorRequestProperty{
//   							References: []interface{}{
//   								&PerformanceFactorReferenceRequestProperty{
//   									InstanceFamily: jsii.String("instanceFamily"),
//   								},
//   							},
//   						},
//   					},
//   					BurstablePerformance: jsii.String("burstablePerformance"),
//   					CpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					ExcludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					InstanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					LocalStorage: jsii.String("localStorage"),
//   					LocalStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   					MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					MemoryMiB: &MemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					RequireHibernateSupport: jsii.Boolean(false),
//   					SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					VCpuCount: &VCpuCountRangeRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				KernelId: jsii.String("kernelId"),
//   				KeyName: jsii.String("keyName"),
//   				Monitoring: &SpotFleetMonitoringProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   				NetworkInterfaces: []interface{}{
//   					&InstanceNetworkInterfaceSpecificationProperty{
//   						AssociatePublicIpAddress: jsii.Boolean(false),
//   						DeleteOnTermination: jsii.Boolean(false),
//   						Description: jsii.String("description"),
//   						DeviceIndex: jsii.Number(123),
//   						Groups: []*string{
//   							jsii.String("groups"),
//   						},
//   						Ipv6AddressCount: jsii.Number(123),
//   						Ipv6Addresses: []interface{}{
//   							&InstanceIpv6AddressProperty{
//   								Ipv6Address: jsii.String("ipv6Address"),
//   							},
//   						},
//   						NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   						PrivateIpAddresses: []interface{}{
//   							&PrivateIpAddressSpecificationProperty{
//   								Primary: jsii.Boolean(false),
//   								PrivateIpAddress: jsii.String("privateIpAddress"),
//   							},
//   						},
//   						SecondaryPrivateIpAddressCount: jsii.Number(123),
//   						SubnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				Placement: &SpotPlacementProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					GroupName: jsii.String("groupName"),
//   					Tenancy: jsii.String("tenancy"),
//   				},
//   				RamdiskId: jsii.String("ramdiskId"),
//   				SecurityGroups: []interface{}{
//   					&GroupIdentifierProperty{
//   						GroupId: jsii.String("groupId"),
//   					},
//   				},
//   				SpotPrice: jsii.String("spotPrice"),
//   				SubnetId: jsii.String("subnetId"),
//   				TagSpecifications: []interface{}{
//   					&SpotFleetTagSpecificationProperty{
//   						ResourceType: jsii.String("resourceType"),
//   						Tags: []CfnTag{
//   							&CfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				UserData: jsii.String("userData"),
//   				WeightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		LaunchTemplateConfigs: []interface{}{
//   			&LaunchTemplateConfigProperty{
//   				LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationProperty{
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					LaunchTemplateName: jsii.String("launchTemplateName"),
//   					Version: jsii.String("version"),
//   				},
//   				Overrides: []interface{}{
//   					&LaunchTemplateOverridesProperty{
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						InstanceRequirements: &InstanceRequirementsRequestProperty{
//   							AcceleratorCount: &AcceleratorCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorManufacturers: []*string{
//   								jsii.String("acceleratorManufacturers"),
//   							},
//   							AcceleratorNames: []*string{
//   								jsii.String("acceleratorNames"),
//   							},
//   							AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorTypes: []*string{
//   								jsii.String("acceleratorTypes"),
//   							},
//   							AllowedInstanceTypes: []*string{
//   								jsii.String("allowedInstanceTypes"),
//   							},
//   							BareMetal: jsii.String("bareMetal"),
//   							BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							BaselinePerformanceFactors: &BaselinePerformanceFactorsRequestProperty{
//   								Cpu: &CpuPerformanceFactorRequestProperty{
//   									References: []interface{}{
//   										&PerformanceFactorReferenceRequestProperty{
//   											InstanceFamily: jsii.String("instanceFamily"),
//   										},
//   									},
//   								},
//   							},
//   							BurstablePerformance: jsii.String("burstablePerformance"),
//   							CpuManufacturers: []*string{
//   								jsii.String("cpuManufacturers"),
//   							},
//   							ExcludedInstanceTypes: []*string{
//   								jsii.String("excludedInstanceTypes"),
//   							},
//   							InstanceGenerations: []*string{
//   								jsii.String("instanceGenerations"),
//   							},
//   							LocalStorage: jsii.String("localStorage"),
//   							LocalStorageTypes: []*string{
//   								jsii.String("localStorageTypes"),
//   							},
//   							MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   							MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							MemoryMiB: &MemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							RequireHibernateSupport: jsii.Boolean(false),
//   							SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							VCpuCount: &VCpuCountRangeRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   						},
//   						InstanceType: jsii.String("instanceType"),
//   						Priority: jsii.Number(123),
//   						SpotPrice: jsii.String("spotPrice"),
//   						SubnetId: jsii.String("subnetId"),
//   						WeightedCapacity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		LoadBalancersConfig: &LoadBalancersConfigProperty{
//   			ClassicLoadBalancersConfig: &ClassicLoadBalancersConfigProperty{
//   				ClassicLoadBalancers: []interface{}{
//   					&ClassicLoadBalancerProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			TargetGroupsConfig: &TargetGroupsConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupProperty{
//   						Arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   		OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		OnDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   		OnDemandTargetCapacity: jsii.Number(123),
//   		ReplaceUnhealthyInstances: jsii.Boolean(false),
//   		SpotMaintenanceStrategies: &SpotMaintenanceStrategiesProperty{
//   			CapacityRebalance: &SpotCapacityRebalanceProperty{
//   				ReplacementStrategy: jsii.String("replacementStrategy"),
//   				TerminationDelay: jsii.Number(123),
//   			},
//   		},
//   		SpotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   		SpotPrice: jsii.String("spotPrice"),
//   		TagSpecifications: []interface{}{
//   			&SpotFleetTagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		TargetCapacity: jsii.Number(123),
//   		TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		TerminateInstancesWithExpiration: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   		ValidFrom: jsii.String("validFrom"),
//   		ValidUntil: jsii.String("validUntil"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
//
type CfnSpotFleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSpotFleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSpotFleetPropsMixin
type jsiiProxy_CfnSpotFleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSpotFleetPropsMixin) Props() *CfnSpotFleetMixinProps {
	var returns *CfnSpotFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::SpotFleet`.
func NewCfnSpotFleetPropsMixin(props *CfnSpotFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSpotFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSpotFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSpotFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSpotFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::SpotFleet`.
func NewCfnSpotFleetPropsMixin_Override(c CfnSpotFleetPropsMixin, props *CfnSpotFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSpotFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSpotFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSpotFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSpotFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSpotFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSpotFleetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSpotFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


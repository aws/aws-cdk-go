package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the configuration information to launch a fleet--or group--of instances.
//
// An EC2 Fleet can launch multiple instance types across multiple Availability Zones, using the On-Demand Instance, Reserved Instance, and Spot Instance purchasing models together. Using EC2 Fleet, you can define separate On-Demand and Spot capacity targets, specify the instance types that work best for your applications, and specify how Amazon EC2 should distribute your fleet capacity within each purchasing model. For more information, see [Launching an EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEC2FleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnEC2FleetPropsMixin(&CfnEC2FleetMixinProps{
//   	Context: jsii.String("context"),
//   	ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	LaunchTemplateConfigs: []interface{}{
//   		&FleetLaunchTemplateConfigRequestProperty{
//   			LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationRequestProperty{
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   				Version: jsii.String("version"),
//   			},
//   			Overrides: []interface{}{
//   				&FleetLaunchTemplateOverridesRequestProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   					BlockDeviceMappings: []interface{}{
//   						&BlockDeviceMappingProperty{
//   							DeviceName: jsii.String("deviceName"),
//   							Ebs: &EbsBlockDeviceProperty{
//   								DeleteOnTermination: jsii.Boolean(false),
//   								Encrypted: jsii.Boolean(false),
//   								Iops: jsii.Number(123),
//   								KmsKeyId: jsii.String("kmsKeyId"),
//   								SnapshotId: jsii.String("snapshotId"),
//   								VolumeSize: jsii.Number(123),
//   								VolumeType: jsii.String("volumeType"),
//   							},
//   							NoDevice: jsii.String("noDevice"),
//   							VirtualName: jsii.String("virtualName"),
//   						},
//   					},
//   					InstanceRequirements: &InstanceRequirementsRequestProperty{
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
//   						RequireEncryptionInTransit: jsii.Boolean(false),
//   						RequireHibernateSupport: jsii.Boolean(false),
//   						SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						VCpuCount: &VCpuCountRangeRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					MaxPrice: jsii.String("maxPrice"),
//   					Placement: &PlacementProperty{
//   						Affinity: jsii.String("affinity"),
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						GroupName: jsii.String("groupName"),
//   						HostId: jsii.String("hostId"),
//   						HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   						PartitionNumber: jsii.Number(123),
//   						SpreadDomain: jsii.String("spreadDomain"),
//   						Tenancy: jsii.String("tenancy"),
//   					},
//   					Priority: jsii.Number(123),
//   					SubnetId: jsii.String("subnetId"),
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	OnDemandOptions: &OnDemandOptionsRequestProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		CapacityReservationOptions: &CapacityReservationOptionsRequestProperty{
//   			UsageStrategy: jsii.String("usageStrategy"),
//   		},
//   		MaxTotalPrice: jsii.String("maxTotalPrice"),
//   		MinTargetCapacity: jsii.Number(123),
//   		SingleAvailabilityZone: jsii.Boolean(false),
//   		SingleInstanceType: jsii.Boolean(false),
//   	},
//   	ReplaceUnhealthyInstances: jsii.Boolean(false),
//   	SpotOptions: &SpotOptionsRequestProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		InstancePoolsToUseCount: jsii.Number(123),
//   		MaintenanceStrategies: &MaintenanceStrategiesProperty{
//   			CapacityRebalance: &CapacityRebalanceProperty{
//   				ReplacementStrategy: jsii.String("replacementStrategy"),
//   				TerminationDelay: jsii.Number(123),
//   			},
//   		},
//   		MaxTotalPrice: jsii.String("maxTotalPrice"),
//   		MinTargetCapacity: jsii.Number(123),
//   		SingleAvailabilityZone: jsii.Boolean(false),
//   		SingleInstanceType: jsii.Boolean(false),
//   	},
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TargetCapacitySpecification: &TargetCapacitySpecificationRequestProperty{
//   		DefaultTargetCapacityType: jsii.String("defaultTargetCapacityType"),
//   		OnDemandTargetCapacity: jsii.Number(123),
//   		SpotTargetCapacity: jsii.Number(123),
//   		TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		TotalTargetCapacity: jsii.Number(123),
//   	},
//   	TerminateInstancesWithExpiration: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	ValidFrom: jsii.String("validFrom"),
//   	ValidUntil: jsii.String("validUntil"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html
//
type CfnEC2FleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEC2FleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEC2FleetPropsMixin
type jsiiProxy_CfnEC2FleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEC2FleetPropsMixin) Props() *CfnEC2FleetMixinProps {
	var returns *CfnEC2FleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2FleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::EC2Fleet`.
func NewCfnEC2FleetPropsMixin(props *CfnEC2FleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEC2FleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEC2FleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEC2FleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEC2FleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::EC2Fleet`.
func NewCfnEC2FleetPropsMixin_Override(c CfnEC2FleetPropsMixin, props *CfnEC2FleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEC2FleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEC2FleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEC2FleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEC2FleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEC2FleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnEC2FleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEC2FleetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEC2FleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


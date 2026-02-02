package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the properties for creating a launch template.
//
// The minimum required properties for specifying a launch template are as follows:
//
// - You must specify at least one property for the launch template data.
// - You can optionally specify a name for the launch template. If you do not specify a name, CloudFormation creates a name for you.
//
// A launch template can contain some or all of the configuration information to launch an instance. When you launch an instance using a launch template, instance properties that are not specified in the launch template use default values, except the `ImageId` property, which has no default value. If you do not specify an AMI ID for the launch template `ImageId` property, you must specify an AMI ID for the instance `ImageId` property.
//
// For more information, see [Launch an instance from a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnLaunchTemplatePropsMixin(&CfnLaunchTemplateMixinProps{
//   	LaunchTemplateData: &LaunchTemplateDataProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&BlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsProperty{
//   					DeleteOnTermination: jsii.Boolean(false),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeInitializationRate: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   			CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   			CapacityReservationTarget: &CapacityReservationTargetProperty{
//   				CapacityReservationId: jsii.String("capacityReservationId"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   			},
//   		},
//   		CpuOptions: &CpuOptionsProperty{
//   			AmdSevSnp: jsii.String("amdSevSnp"),
//   			CoreCount: jsii.Number(123),
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   		CreditSpecification: &CreditSpecificationProperty{
//   			CpuCredits: jsii.String("cpuCredits"),
//   		},
//   		DisableApiStop: jsii.Boolean(false),
//   		DisableApiTermination: jsii.Boolean(false),
//   		EbsOptimized: jsii.Boolean(false),
//   		ElasticGpuSpecifications: []interface{}{
//   			&ElasticGpuSpecificationProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		ElasticInferenceAccelerators: []interface{}{
//   			&LaunchTemplateElasticInferenceAcceleratorProperty{
//   				Count: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		EnclaveOptions: &EnclaveOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		HibernationOptions: &HibernationOptionsProperty{
//   			Configured: jsii.Boolean(false),
//   		},
//   		IamInstanceProfile: &IamInstanceProfileProperty{
//   			Arn: jsii.String("arn"),
//   			Name: jsii.String("name"),
//   		},
//   		ImageId: jsii.String("imageId"),
//   		InstanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   		InstanceMarketOptions: &InstanceMarketOptionsProperty{
//   			MarketType: jsii.String("marketType"),
//   			SpotOptions: &SpotOptionsProperty{
//   				BlockDurationMinutes: jsii.Number(123),
//   				InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   				MaxPrice: jsii.String("maxPrice"),
//   				SpotInstanceType: jsii.String("spotInstanceType"),
//   				ValidUntil: jsii.String("validUntil"),
//   			},
//   		},
//   		InstanceRequirements: &InstanceRequirementsProperty{
//   			AcceleratorCount: &AcceleratorCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorManufacturers: []*string{
//   				jsii.String("acceleratorManufacturers"),
//   			},
//   			AcceleratorNames: []*string{
//   				jsii.String("acceleratorNames"),
//   			},
//   			AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorTypes: []*string{
//   				jsii.String("acceleratorTypes"),
//   			},
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   			BareMetal: jsii.String("bareMetal"),
//   			BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			BaselinePerformanceFactors: &BaselinePerformanceFactorsProperty{
//   				Cpu: &CpuProperty{
//   					References: []interface{}{
//   						&ReferenceProperty{
//   							InstanceFamily: jsii.String("instanceFamily"),
//   						},
//   					},
//   				},
//   			},
//   			BurstablePerformance: jsii.String("burstablePerformance"),
//   			CpuManufacturers: []*string{
//   				jsii.String("cpuManufacturers"),
//   			},
//   			ExcludedInstanceTypes: []*string{
//   				jsii.String("excludedInstanceTypes"),
//   			},
//   			InstanceGenerations: []*string{
//   				jsii.String("instanceGenerations"),
//   			},
//   			LocalStorage: jsii.String("localStorage"),
//   			LocalStorageTypes: []*string{
//   				jsii.String("localStorageTypes"),
//   			},
//   			MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   			MemoryGiBPerVCpu: &MemoryGiBPerVCpuProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			MemoryMiB: &MemoryMiBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkBandwidthGbps: &NetworkBandwidthGbpsProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkInterfaceCount: &NetworkInterfaceCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			RequireHibernateSupport: jsii.Boolean(false),
//   			SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			TotalLocalStorageGb: &TotalLocalStorageGBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			VCpuCount: &VCpuCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		KernelId: jsii.String("kernelId"),
//   		KeyName: jsii.String("keyName"),
//   		LicenseSpecifications: []interface{}{
//   			&LicenseSpecificationProperty{
//   				LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   			},
//   		},
//   		MaintenanceOptions: &MaintenanceOptionsProperty{
//   			AutoRecovery: jsii.String("autoRecovery"),
//   		},
//   		MetadataOptions: &MetadataOptionsProperty{
//   			HttpEndpoint: jsii.String("httpEndpoint"),
//   			HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   			HttpPutResponseHopLimit: jsii.Number(123),
//   			HttpTokens: jsii.String("httpTokens"),
//   			InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   		},
//   		Monitoring: &MonitoringProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			&NetworkInterfaceProperty{
//   				AssociateCarrierIpAddress: jsii.Boolean(false),
//   				AssociatePublicIpAddress: jsii.Boolean(false),
//   				ConnectionTrackingSpecification: &ConnectionTrackingSpecificationProperty{
//   					TcpEstablishedTimeout: jsii.Number(123),
//   					UdpStreamTimeout: jsii.Number(123),
//   					UdpTimeout: jsii.Number(123),
//   				},
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Description: jsii.String("description"),
//   				DeviceIndex: jsii.Number(123),
//   				EnaQueueCount: jsii.Number(123),
//   				EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   					EnaSrdEnabled: jsii.Boolean(false),
//   					EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   						EnaSrdUdpEnabled: jsii.Boolean(false),
//   					},
//   				},
//   				Groups: []*string{
//   					jsii.String("groups"),
//   				},
//   				InterfaceType: jsii.String("interfaceType"),
//   				Ipv4PrefixCount: jsii.Number(123),
//   				Ipv4Prefixes: []interface{}{
//   					&Ipv4PrefixSpecificationProperty{
//   						Ipv4Prefix: jsii.String("ipv4Prefix"),
//   					},
//   				},
//   				Ipv6AddressCount: jsii.Number(123),
//   				Ipv6Addresses: []interface{}{
//   					&Ipv6AddProperty{
//   						Ipv6Address: jsii.String("ipv6Address"),
//   					},
//   				},
//   				Ipv6PrefixCount: jsii.Number(123),
//   				Ipv6Prefixes: []interface{}{
//   					&Ipv6PrefixSpecificationProperty{
//   						Ipv6Prefix: jsii.String("ipv6Prefix"),
//   					},
//   				},
//   				NetworkCardIndex: jsii.Number(123),
//   				NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   				PrimaryIpv6: jsii.Boolean(false),
//   				PrivateIpAddress: jsii.String("privateIpAddress"),
//   				PrivateIpAddresses: []interface{}{
//   					&PrivateIpAddProperty{
//   						Primary: jsii.Boolean(false),
//   						PrivateIpAddress: jsii.String("privateIpAddress"),
//   					},
//   				},
//   				SecondaryPrivateIpAddressCount: jsii.Number(123),
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		NetworkPerformanceOptions: &NetworkPerformanceOptionsProperty{
//   			BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   		},
//   		Placement: &PlacementProperty{
//   			Affinity: jsii.String("affinity"),
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			GroupId: jsii.String("groupId"),
//   			GroupName: jsii.String("groupName"),
//   			HostId: jsii.String("hostId"),
//   			HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   			PartitionNumber: jsii.Number(123),
//   			SpreadDomain: jsii.String("spreadDomain"),
//   			Tenancy: jsii.String("tenancy"),
//   		},
//   		PrivateDnsNameOptions: &PrivateDnsNameOptionsProperty{
//   			EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   			EnableResourceNameDnsARecord: jsii.Boolean(false),
//   			HostnameType: jsii.String("hostnameType"),
//   		},
//   		RamDiskId: jsii.String("ramDiskId"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		TagSpecifications: []interface{}{
//   			&TagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UserData: jsii.String("userData"),
//   	},
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	TagSpecifications: []interface{}{
//   		&LaunchTemplateTagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	VersionDescription: jsii.String("versionDescription"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html
//
type CfnLaunchTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLaunchTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLaunchTemplatePropsMixin
type jsiiProxy_CfnLaunchTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLaunchTemplatePropsMixin) Props() *CfnLaunchTemplateMixinProps {
	var returns *CfnLaunchTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::LaunchTemplate`.
func NewCfnLaunchTemplatePropsMixin(props *CfnLaunchTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLaunchTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLaunchTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnLaunchTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::LaunchTemplate`.
func NewCfnLaunchTemplatePropsMixin_Override(c CfnLaunchTemplatePropsMixin, props *CfnLaunchTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnLaunchTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLaunchTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnLaunchTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnLaunchTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLaunchTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawsworkspacesinstancesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsworkspacesinstances/previewawsworkspacesinstancesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WorkspacesInstances::WorkspaceInstance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspaceInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkspaceInstancePropsMixin(&CfnWorkspaceInstanceMixinProps{
//   	ManagedInstance: &ManagedInstanceProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&BlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsBlockDeviceProperty{
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					Throughput: jsii.Number(123),
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
//   		CpuOptions: &CpuOptionsRequestProperty{
//   			CoreCount: jsii.Number(123),
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   		CreditSpecification: &CreditSpecificationRequestProperty{
//   			CpuCredits: jsii.String("cpuCredits"),
//   		},
//   		DisableApiStop: jsii.Boolean(false),
//   		EbsOptimized: jsii.Boolean(false),
//   		EnablePrimaryIpv6: jsii.Boolean(false),
//   		EnclaveOptions: &EnclaveOptionsRequestProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		HibernationOptions: &HibernationOptionsRequestProperty{
//   			Configured: jsii.Boolean(false),
//   		},
//   		IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   			Arn: jsii.String("arn"),
//   			Name: jsii.String("name"),
//   		},
//   		ImageId: jsii.String("imageId"),
//   		InstanceMarketOptions: &InstanceMarketOptionsRequestProperty{
//   			MarketType: jsii.String("marketType"),
//   			SpotOptions: &SpotMarketOptionsProperty{
//   				InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   				MaxPrice: jsii.String("maxPrice"),
//   				SpotInstanceType: jsii.String("spotInstanceType"),
//   				ValidUntilUtc: jsii.String("validUntilUtc"),
//   			},
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		Ipv6AddressCount: jsii.Number(123),
//   		KeyName: jsii.String("keyName"),
//   		LicenseSpecifications: []interface{}{
//   			&LicenseConfigurationRequestProperty{
//   				LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   			},
//   		},
//   		MaintenanceOptions: &InstanceMaintenanceOptionsRequestProperty{
//   			AutoRecovery: jsii.String("autoRecovery"),
//   		},
//   		MetadataOptions: &InstanceMetadataOptionsRequestProperty{
//   			HttpEndpoint: jsii.String("httpEndpoint"),
//   			HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   			HttpPutResponseHopLimit: jsii.Number(123),
//   			HttpTokens: jsii.String("httpTokens"),
//   			InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   		},
//   		Monitoring: &RunInstancesMonitoringEnabledProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			&InstanceNetworkInterfaceSpecificationProperty{
//   				Description: jsii.String("description"),
//   				DeviceIndex: jsii.Number(123),
//   				Groups: []*string{
//   					jsii.String("groups"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		NetworkPerformanceOptions: &InstanceNetworkPerformanceOptionsRequestProperty{
//   			BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   		},
//   		Placement: &PlacementProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			GroupId: jsii.String("groupId"),
//   			GroupName: jsii.String("groupName"),
//   			PartitionNumber: jsii.Number(123),
//   			Tenancy: jsii.String("tenancy"),
//   		},
//   		PrivateDnsNameOptions: &PrivateDnsNameOptionsRequestProperty{
//   			EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   			EnableResourceNameDnsARecord: jsii.Boolean(false),
//   			HostnameType: jsii.String("hostnameType"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-workspaceinstance.html
//
type CfnWorkspaceInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkspaceInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspaceInstancePropsMixin
type jsiiProxy_CfnWorkspaceInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkspaceInstancePropsMixin) Props() *CfnWorkspaceInstanceMixinProps {
	var returns *CfnWorkspaceInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkspacesInstances::WorkspaceInstance`.
func NewCfnWorkspaceInstancePropsMixin(props *CfnWorkspaceInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkspaceInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspaceInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspaceInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesinstances.mixins.CfnWorkspaceInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkspacesInstances::WorkspaceInstance`.
func NewCfnWorkspaceInstancePropsMixin_Override(c CfnWorkspaceInstancePropsMixin, props *CfnWorkspaceInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesinstances.mixins.CfnWorkspaceInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkspaceInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspacesinstances.mixins.CfnWorkspaceInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspaceInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspacesinstances.mixins.CfnWorkspaceInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkspaceInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package mixinsawsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsfsx/mixinsawsfsx/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an FSx for ONTAP or Amazon FSx for OpenZFS storage volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVolumePropsMixin := awscdkmixinspreview.Mixins.NewCfnVolumePropsMixin(&CfnVolumeMixinProps{
//   	BackupId: jsii.String("backupId"),
//   	Name: jsii.String("name"),
//   	OntapConfiguration: &OntapConfigurationProperty{
//   		AggregateConfiguration: &AggregateConfigurationProperty{
//   			Aggregates: []*string{
//   				jsii.String("aggregates"),
//   			},
//   			ConstituentsPerAggregate: jsii.Number(123),
//   		},
//   		CopyTagsToBackups: jsii.String("copyTagsToBackups"),
//   		JunctionPath: jsii.String("junctionPath"),
//   		OntapVolumeType: jsii.String("ontapVolumeType"),
//   		SecurityStyle: jsii.String("securityStyle"),
//   		SizeInBytes: jsii.String("sizeInBytes"),
//   		SizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		SnaplockConfiguration: &SnaplockConfigurationProperty{
//   			AuditLogVolume: jsii.String("auditLogVolume"),
//   			AutocommitPeriod: &AutocommitPeriodProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			PrivilegedDelete: jsii.String("privilegedDelete"),
//   			RetentionPeriod: &SnaplockRetentionPeriodProperty{
//   				DefaultRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   				MaximumRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   				MinimumRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			SnaplockType: jsii.String("snaplockType"),
//   			VolumeAppendModeEnabled: jsii.String("volumeAppendModeEnabled"),
//   		},
//   		SnapshotPolicy: jsii.String("snapshotPolicy"),
//   		StorageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		StorageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//   		TieringPolicy: &TieringPolicyProperty{
//   			CoolingPeriod: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   		VolumeStyle: jsii.String("volumeStyle"),
//   	},
//   	OpenZfsConfiguration: &OpenZFSConfigurationProperty{
//   		CopyTagsToSnapshots: jsii.Boolean(false),
//   		DataCompressionType: jsii.String("dataCompressionType"),
//   		NfsExports: []interface{}{
//   			&NfsExportsProperty{
//   				ClientConfigurations: []interface{}{
//   					&ClientConfigurationsProperty{
//   						Clients: jsii.String("clients"),
//   						Options: []*string{
//   							jsii.String("options"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Options: []*string{
//   			jsii.String("options"),
//   		},
//   		OriginSnapshot: &OriginSnapshotProperty{
//   			CopyStrategy: jsii.String("copyStrategy"),
//   			SnapshotArn: jsii.String("snapshotArn"),
//   		},
//   		ParentVolumeId: jsii.String("parentVolumeId"),
//   		ReadOnly: jsii.Boolean(false),
//   		RecordSizeKiB: jsii.Number(123),
//   		StorageCapacityQuotaGiB: jsii.Number(123),
//   		StorageCapacityReservationGiB: jsii.Number(123),
//   		UserAndGroupQuotas: []interface{}{
//   			&UserAndGroupQuotasProperty{
//   				Id: jsii.Number(123),
//   				StorageCapacityQuotaGiB: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeType: jsii.String("volumeType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html
//
type CfnVolumePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVolumeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVolumePropsMixin
type jsiiProxy_CfnVolumePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVolumePropsMixin) Props() *CfnVolumeMixinProps {
	var returns *CfnVolumeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVolumePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FSx::Volume`.
func NewCfnVolumePropsMixin(props *CfnVolumeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVolumePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVolumePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVolumePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnVolumePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FSx::Volume`.
func NewCfnVolumePropsMixin_Override(c CfnVolumePropsMixin, props *CfnVolumeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnVolumePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVolumePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVolumePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnVolumePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVolumePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnVolumePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVolumePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVolumePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


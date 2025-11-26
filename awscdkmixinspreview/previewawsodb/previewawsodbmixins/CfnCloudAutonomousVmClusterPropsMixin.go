package previewawsodbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsodb/previewawsodbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::CloudAutonomousVmCluster` resource creates an Autonomous VM cluster.
//
// An Autonomous VM cluster provides the infrastructure for running Autonomous Databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudAutonomousVmClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudAutonomousVmClusterPropsMixin(&CfnCloudAutonomousVmClusterMixinProps{
//   	AutonomousDataStorageSizeInTBs: jsii.Number(123),
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	CpuCoreCountPerNode: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IsMtlsEnabledVmCluster: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		DaysOfWeek: []*string{
//   			jsii.String("daysOfWeek"),
//   		},
//   		HoursOfDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		LeadTimeInWeeks: jsii.Number(123),
//   		Months: []*string{
//   			jsii.String("months"),
//   		},
//   		Preference: jsii.String("preference"),
//   		WeeksOfMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	MemoryPerOracleComputeUnitInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortNonTls: jsii.Number(123),
//   	ScanListenerPortTls: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   	TotalContainerDatabases: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html
//
type CfnCloudAutonomousVmClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudAutonomousVmClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudAutonomousVmClusterPropsMixin
type jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin) Props() *CfnCloudAutonomousVmClusterMixinProps {
	var returns *CfnCloudAutonomousVmClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ODB::CloudAutonomousVmCluster`.
func NewCfnCloudAutonomousVmClusterPropsMixin(props *CfnCloudAutonomousVmClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudAutonomousVmClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudAutonomousVmClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ODB::CloudAutonomousVmCluster`.
func NewCfnCloudAutonomousVmClusterPropsMixin_Override(c CfnCloudAutonomousVmClusterPropsMixin, props *CfnCloudAutonomousVmClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudAutonomousVmClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudAutonomousVmClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudAutonomousVmClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


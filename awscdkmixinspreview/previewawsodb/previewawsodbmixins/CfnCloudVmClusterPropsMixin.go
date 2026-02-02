package previewawsodbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsodb/previewawsodbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::CloudVmCluster` resource creates a VM cluster on the specified Exadata infrastructure in the Oracle Database.
//
// A VM cluster provides the compute resources for Oracle Database workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudVmClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudVmClusterPropsMixin(&CfnCloudVmClusterMixinProps{
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	ClusterName: jsii.String("clusterName"),
//   	CpuCoreCount: jsii.Number(123),
//   	DataCollectionOptions: &DataCollectionOptionsProperty{
//   		IsDiagnosticsEventsEnabled: jsii.Boolean(false),
//   		IsHealthMonitoringEnabled: jsii.Boolean(false),
//   		IsIncidentLogsEnabled: jsii.Boolean(false),
//   	},
//   	DataStorageSizeInTBs: jsii.Number(123),
//   	DbNodes: []interface{}{
//   		&DbNodeProperty{
//   			BackupIpId: jsii.String("backupIpId"),
//   			BackupVnic2Id: jsii.String("backupVnic2Id"),
//   			CpuCoreCount: jsii.Number(123),
//   			DbNodeArn: jsii.String("dbNodeArn"),
//   			DbNodeId: jsii.String("dbNodeId"),
//   			DbNodeStorageSizeInGBs: jsii.Number(123),
//   			DbServerId: jsii.String("dbServerId"),
//   			DbSystemId: jsii.String("dbSystemId"),
//   			HostIpId: jsii.String("hostIpId"),
//   			Hostname: jsii.String("hostname"),
//   			MemorySizeInGBs: jsii.Number(123),
//   			Ocid: jsii.String("ocid"),
//   			Status: jsii.String("status"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Vnic2Id: jsii.String("vnic2Id"),
//   			VnicId: jsii.String("vnicId"),
//   		},
//   	},
//   	DbNodeStorageSizeInGBs: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	GiVersion: jsii.String("giVersion"),
//   	Hostname: jsii.String("hostname"),
//   	IsLocalBackupEnabled: jsii.Boolean(false),
//   	IsSparseDiskgroupEnabled: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MemorySizeInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortTcp: jsii.Number(123),
//   	SshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	SystemVersion: jsii.String("systemVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html
//
type CfnCloudVmClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudVmClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudVmClusterPropsMixin
type jsiiProxy_CfnCloudVmClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudVmClusterPropsMixin) Props() *CfnCloudVmClusterMixinProps {
	var returns *CfnCloudVmClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ODB::CloudVmCluster`.
func NewCfnCloudVmClusterPropsMixin(props *CfnCloudVmClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudVmClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudVmClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudVmClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ODB::CloudVmCluster`.
func NewCfnCloudVmClusterPropsMixin_Override(c CfnCloudVmClusterPropsMixin, props *CfnCloudVmClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudVmClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudVmClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudVmClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudVmClusterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCloudVmClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


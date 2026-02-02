package previewawsfsxmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfsx/previewawsfsxmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::FSx::FileSystem` resource is an Amazon FSx resource type that specifies an Amazon FSx file system.
//
// You can create any of the following supported file system types:
//
// - Amazon FSx for Lustre
// - Amazon FSx for NetApp ONTAP
// - FSx for OpenZFS
// - Amazon FSx for Windows File Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFileSystemPropsMixin := awscdkmixinspreview.Mixins.NewCfnFileSystemPropsMixin(&CfnFileSystemMixinProps{
//   	BackupId: jsii.String("backupId"),
//   	FileSystemType: jsii.String("fileSystemType"),
//   	FileSystemTypeVersion: jsii.String("fileSystemTypeVersion"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LustreConfiguration: &LustreConfigurationProperty{
//   		AutoImportPolicy: jsii.String("autoImportPolicy"),
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DataCompressionType: jsii.String("dataCompressionType"),
//   		DataReadCacheConfiguration: &DataReadCacheConfigurationProperty{
//   			SizeGiB: jsii.Number(123),
//   			SizingMode: jsii.String("sizingMode"),
//   		},
//   		DeploymentType: jsii.String("deploymentType"),
//   		DriveCacheType: jsii.String("driveCacheType"),
//   		EfaEnabled: jsii.Boolean(false),
//   		ExportPath: jsii.String("exportPath"),
//   		ImportedFileChunkSize: jsii.Number(123),
//   		ImportPath: jsii.String("importPath"),
//   		MetadataConfiguration: &MetadataConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	NetworkType: jsii.String("networkType"),
//   	OntapConfiguration: &OntapConfigurationProperty{
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DeploymentType: jsii.String("deploymentType"),
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		EndpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   		EndpointIpv6AddressRange: jsii.String("endpointIpv6AddressRange"),
//   		FsxAdminPassword: jsii.String("fsxAdminPassword"),
//   		HaPairs: jsii.Number(123),
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		RouteTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		ThroughputCapacityPerHaPair: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	OpenZfsConfiguration: &OpenZFSConfigurationProperty{
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		CopyTagsToVolumes: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DeploymentType: jsii.String("deploymentType"),
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		EndpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   		EndpointIpv6AddressRange: jsii.String("endpointIpv6AddressRange"),
//   		Options: []*string{
//   			jsii.String("options"),
//   		},
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		ReadCacheConfiguration: &ReadCacheConfigurationProperty{
//   			SizeGiB: jsii.Number(123),
//   			SizingMode: jsii.String("sizingMode"),
//   		},
//   		RootVolumeConfiguration: &RootVolumeConfigurationProperty{
//   			CopyTagsToSnapshots: jsii.Boolean(false),
//   			DataCompressionType: jsii.String("dataCompressionType"),
//   			NfsExports: []interface{}{
//   				&NfsExportsProperty{
//   					ClientConfigurations: []interface{}{
//   						&ClientConfigurationsProperty{
//   							Clients: jsii.String("clients"),
//   							Options: []*string{
//   								jsii.String("options"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			ReadOnly: jsii.Boolean(false),
//   			RecordSizeKiB: jsii.Number(123),
//   			UserAndGroupQuotas: []interface{}{
//   				&UserAndGroupQuotasProperty{
//   					Id: jsii.Number(123),
//   					StorageCapacityQuotaGiB: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		RouteTableIds: []*string{
//   			jsii.String("routeTableIds"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WindowsConfiguration: &WindowsConfigurationProperty{
//   		ActiveDirectoryId: jsii.String("activeDirectoryId"),
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		AuditLogConfiguration: &AuditLogConfigurationProperty{
//   			AuditLogDestination: jsii.String("auditLogDestination"),
//   			FileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   			FileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//   		},
//   		AutomaticBackupRetentionDays: jsii.Number(123),
//   		CopyTagsToBackups: jsii.Boolean(false),
//   		DailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   		DeploymentType: jsii.String("deploymentType"),
//   		DiskIopsConfiguration: &DiskIopsConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		PreferredSubnetId: jsii.String("preferredSubnetId"),
//   		SelfManagedActiveDirectoryConfiguration: &SelfManagedActiveDirectoryConfigurationProperty{
//   			DnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			DomainJoinServiceAccountSecret: jsii.String("domainJoinServiceAccountSecret"),
//   			DomainName: jsii.String("domainName"),
//   			FileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   			Password: jsii.String("password"),
//   			UserName: jsii.String("userName"),
//   		},
//   		ThroughputCapacity: jsii.Number(123),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
type CfnFileSystemPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFileSystemMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFileSystemPropsMixin
type jsiiProxy_CfnFileSystemPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFileSystemPropsMixin) Props() *CfnFileSystemMixinProps {
	var returns *CfnFileSystemMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileSystemPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FSx::FileSystem`.
func NewCfnFileSystemPropsMixin(props *CfnFileSystemMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFileSystemPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFileSystemPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFileSystemPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnFileSystemPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FSx::FileSystem`.
func NewCfnFileSystemPropsMixin_Override(c CfnFileSystemPropsMixin, props *CfnFileSystemMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnFileSystemPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFileSystemPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileSystemPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnFileSystemPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFileSystemPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fsx.mixins.CfnFileSystemPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFileSystemPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFileSystemPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


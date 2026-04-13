package awsfsx

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnDataRepositoryAssociationMixinProps",
		reflect.TypeOf((*CfnDataRepositoryAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnDataRepositoryAssociationPropsMixin",
		reflect.TypeOf((*CfnDataRepositoryAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataRepositoryAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnDataRepositoryAssociationPropsMixin.AutoExportPolicyProperty",
		reflect.TypeOf((*CfnDataRepositoryAssociationPropsMixin_AutoExportPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnDataRepositoryAssociationPropsMixin.AutoImportPolicyProperty",
		reflect.TypeOf((*CfnDataRepositoryAssociationPropsMixin_AutoImportPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnDataRepositoryAssociationPropsMixin.S3Property",
		reflect.TypeOf((*CfnDataRepositoryAssociationPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemMixinProps",
		reflect.TypeOf((*CfnFileSystemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin",
		reflect.TypeOf((*CfnFileSystemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFileSystemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.AuditLogConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_AuditLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.ClientConfigurationsProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ClientConfigurationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.DataReadCacheConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_DataReadCacheConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.DiskIopsConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_DiskIopsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.FsrmConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_FsrmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.LustreConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_LustreConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.NfsExportsProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_NfsExportsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.OntapConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_OntapConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.OpenZFSConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_OpenZFSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.ReadCacheConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ReadCacheConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.RootVolumeConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_RootVolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.SelfManagedActiveDirectoryConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_SelfManagedActiveDirectoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.UserAndGroupQuotasProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_UserAndGroupQuotasProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnFileSystemPropsMixin.WindowsConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_WindowsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentMixinProps",
		reflect.TypeOf((*CfnS3AccessPointAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnS3AccessPointAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.FileSystemGIDProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_FileSystemGIDProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.OntapFileSystemIdentityProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_OntapFileSystemIdentityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.OntapUnixFileSystemUserProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_OntapUnixFileSystemUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.OntapWindowsFileSystemUserProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_OntapWindowsFileSystemUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.OpenZFSFileSystemIdentityProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_OpenZFSFileSystemIdentityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.OpenZFSPosixFileSystemUserProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_OpenZFSPosixFileSystemUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.S3AccessPointOntapConfigurationProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_S3AccessPointOntapConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.S3AccessPointOpenZFSConfigurationProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_S3AccessPointOpenZFSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.S3AccessPointProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_S3AccessPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnS3AccessPointAttachmentPropsMixin.S3AccessPointVpcConfigurationProperty",
		reflect.TypeOf((*CfnS3AccessPointAttachmentPropsMixin_S3AccessPointVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnSnapshotMixinProps",
		reflect.TypeOf((*CfnSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnSnapshotPropsMixin",
		reflect.TypeOf((*CfnSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnStorageVirtualMachineMixinProps",
		reflect.TypeOf((*CfnStorageVirtualMachineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnStorageVirtualMachinePropsMixin",
		reflect.TypeOf((*CfnStorageVirtualMachinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStorageVirtualMachinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnStorageVirtualMachinePropsMixin.ActiveDirectoryConfigurationProperty",
		reflect.TypeOf((*CfnStorageVirtualMachinePropsMixin_ActiveDirectoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnStorageVirtualMachinePropsMixin.SelfManagedActiveDirectoryConfigurationProperty",
		reflect.TypeOf((*CfnStorageVirtualMachinePropsMixin_SelfManagedActiveDirectoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumeMixinProps",
		reflect.TypeOf((*CfnVolumeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin",
		reflect.TypeOf((*CfnVolumePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVolumePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.AggregateConfigurationProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_AggregateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.AutocommitPeriodProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_AutocommitPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.ClientConfigurationsProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_ClientConfigurationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.NfsExportsProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_NfsExportsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.OntapConfigurationProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_OntapConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.OpenZFSConfigurationProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_OpenZFSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.OriginSnapshotProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_OriginSnapshotProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.RetentionPeriodProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_RetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.SnaplockConfigurationProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_SnaplockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.SnaplockRetentionPeriodProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_SnaplockRetentionPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.TieringPolicyProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_TieringPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fsx.CfnVolumePropsMixin.UserAndGroupQuotasProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_UserAndGroupQuotasProperty)(nil)).Elem(),
	)
}

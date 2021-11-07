package awslightsail

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_lightsail.CfnDatabase",
		reflect.TypeOf((*CfnDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDatabaseArn", GoGetter: "AttrDatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetention", GoGetter: "BackupRetention"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificateIdentifier", GoGetter: "CaCertificateIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "masterDatabaseName", GoGetter: "MasterDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "masterUsername", GoGetter: "MasterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPassword", GoGetter: "MasterUserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preferredBackupWindow", GoGetter: "PreferredBackupWindow"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMaintenanceWindow", GoGetter: "PreferredMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseBlueprintId", GoGetter: "RelationalDatabaseBlueprintId"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseBundleId", GoGetter: "RelationalDatabaseBundleId"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseName", GoGetter: "RelationalDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "relationalDatabaseParameters", GoGetter: "RelationalDatabaseParameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rotateMasterUserPassword", GoGetter: "RotateMasterUserPassword"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnDatabase.RelationalDatabaseParameterProperty",
		reflect.TypeOf((*CfnDatabase_RelationalDatabaseParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnDatabaseProps",
		reflect.TypeOf((*CfnDatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_lightsail.CfnDisk",
		reflect.TypeOf((*CfnDisk)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addOns", GoGetter: "AddOns"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAttachedTo", GoGetter: "AttrAttachedTo"},
			_jsii_.MemberProperty{JsiiProperty: "attrAttachmentState", GoGetter: "AttrAttachmentState"},
			_jsii_.MemberProperty{JsiiProperty: "attrDiskArn", GoGetter: "AttrDiskArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIops", GoGetter: "AttrIops"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsAttached", GoGetter: "AttrIsAttached"},
			_jsii_.MemberProperty{JsiiProperty: "attrPath", GoGetter: "AttrPath"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceType", GoGetter: "AttrResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrSupportCode", GoGetter: "AttrSupportCode"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskName", GoGetter: "DiskName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInGb", GoGetter: "SizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDisk{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnDisk.AddOnProperty",
		reflect.TypeOf((*CfnDisk_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnDisk.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnDisk_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnDiskProps",
		reflect.TypeOf((*CfnDiskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_lightsail.CfnInstance",
		reflect.TypeOf((*CfnInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addOns", GoGetter: "AddOns"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrHardwareCpuCount", GoGetter: "AttrHardwareCpuCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrHardwareRamSizeInGb", GoGetter: "AttrHardwareRamSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "attrInstanceArn", GoGetter: "AttrInstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsStaticIp", GoGetter: "AttrIsStaticIp"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocationAvailabilityZone", GoGetter: "AttrLocationAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocationRegionName", GoGetter: "AttrLocationRegionName"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkingMonthlyTransferGbPerMonthAllocated", GoGetter: "AttrNetworkingMonthlyTransferGbPerMonthAllocated"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrivateIpAddress", GoGetter: "AttrPrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicIpAddress", GoGetter: "AttrPublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceType", GoGetter: "AttrResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrSshKeyName", GoGetter: "AttrSshKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateCode", GoGetter: "AttrStateCode"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateName", GoGetter: "AttrStateName"},
			_jsii_.MemberProperty{JsiiProperty: "attrSupportCode", GoGetter: "AttrSupportCode"},
			_jsii_.MemberProperty{JsiiProperty: "attrUserName", GoGetter: "AttrUserName"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "blueprintId", GoGetter: "BlueprintId"},
			_jsii_.MemberProperty{JsiiProperty: "bundleId", GoGetter: "BundleId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hardware", GoGetter: "Hardware"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceName", GoGetter: "InstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairName", GoGetter: "KeyPairName"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networking", GoGetter: "Networking"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.AddOnProperty",
		reflect.TypeOf((*CfnInstance_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnInstance_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.DiskProperty",
		reflect.TypeOf((*CfnInstance_DiskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.HardwareProperty",
		reflect.TypeOf((*CfnInstance_HardwareProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.LocationProperty",
		reflect.TypeOf((*CfnInstance_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.MonthlyTransferProperty",
		reflect.TypeOf((*CfnInstance_MonthlyTransferProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.NetworkingProperty",
		reflect.TypeOf((*CfnInstance_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.PortProperty",
		reflect.TypeOf((*CfnInstance_PortProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstance.StateProperty",
		reflect.TypeOf((*CfnInstance_StateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnInstanceProps",
		reflect.TypeOf((*CfnInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_lightsail.CfnStaticIp",
		reflect.TypeOf((*CfnStaticIp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attachedTo", GoGetter: "AttachedTo"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpAddress", GoGetter: "AttrIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsAttached", GoGetter: "AttrIsAttached"},
			_jsii_.MemberProperty{JsiiProperty: "attrStaticIpArn", GoGetter: "AttrStaticIpArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "staticIpName", GoGetter: "StaticIpName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStaticIp{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_lightsail.CfnStaticIpProps",
		reflect.TypeOf((*CfnStaticIpProps)(nil)).Elem(),
	)
}

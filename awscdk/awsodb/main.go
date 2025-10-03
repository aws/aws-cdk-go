package awsodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		reflect.TypeOf((*CfnCloudAutonomousVmCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAutonomousDataStoragePercentage", GoGetter: "AttrAutonomousDataStoragePercentage"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailableAutonomousDataStorageSizeInTBs", GoGetter: "AttrAvailableAutonomousDataStorageSizeInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailableContainerDatabases", GoGetter: "AttrAvailableContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailableCpus", GoGetter: "AttrAvailableCpus"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudAutonomousVmClusterArn", GoGetter: "AttrCloudAutonomousVmClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudAutonomousVmClusterId", GoGetter: "AttrCloudAutonomousVmClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "attrComputeModel", GoGetter: "AttrComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "attrCpuCoreCount", GoGetter: "AttrCpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrCpuPercentage", GoGetter: "AttrCpuPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataStorageSizeInGBs", GoGetter: "AttrDataStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataStorageSizeInTBs", GoGetter: "AttrDataStorageSizeInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbNodeStorageSizeInGBs", GoGetter: "AttrDbNodeStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomain", GoGetter: "AttrDomain"},
			_jsii_.MemberProperty{JsiiProperty: "attrExadataStorageInTBsLowestScaledValue", GoGetter: "AttrExadataStorageInTBsLowestScaledValue"},
			_jsii_.MemberProperty{JsiiProperty: "attrHostname", GoGetter: "AttrHostname"},
			_jsii_.MemberProperty{JsiiProperty: "attrMaxAcdsLowestScaledValue", GoGetter: "AttrMaxAcdsLowestScaledValue"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemorySizeInGBs", GoGetter: "AttrMemorySizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrNodeCount", GoGetter: "AttrNodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrNonProvisionableAutonomousContainerDatabases", GoGetter: "AttrNonProvisionableAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "attrOcid", GoGetter: "AttrOcid"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciResourceAnchorName", GoGetter: "AttrOciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciUrl", GoGetter: "AttrOciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrProvisionableAutonomousContainerDatabases", GoGetter: "AttrProvisionableAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "attrProvisionedAutonomousContainerDatabases", GoGetter: "AttrProvisionedAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "attrProvisionedCpus", GoGetter: "AttrProvisionedCpus"},
			_jsii_.MemberProperty{JsiiProperty: "attrReclaimableCpus", GoGetter: "AttrReclaimableCpus"},
			_jsii_.MemberProperty{JsiiProperty: "attrReservedCpus", GoGetter: "AttrReservedCpus"},
			_jsii_.MemberProperty{JsiiProperty: "attrShape", GoGetter: "AttrShape"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousDataStorageSizeInTBs", GoGetter: "AutonomousDataStorageSizeInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAutonomousVmClusterRef", GoGetter: "CloudAutonomousVmClusterRef"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureId", GoGetter: "CloudExadataInfrastructureId"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountPerNode", GoGetter: "CpuCoreCountPerNode"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbServers", GoGetter: "DbServers"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "isMtlsEnabledVmCluster", GoGetter: "IsMtlsEnabledVmCluster"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindow", GoGetter: "MaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "memoryPerOracleComputeUnitInGBs", GoGetter: "MemoryPerOracleComputeUnitInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkId", GoGetter: "OdbNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortNonTls", GoGetter: "ScanListenerPortNonTls"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTls", GoGetter: "ScanListenerPortTls"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalContainerDatabases", GoGetter: "TotalContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudAutonomousVmCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICloudAutonomousVmClusterRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster.MaintenanceWindowProperty",
		reflect.TypeOf((*CfnCloudAutonomousVmCluster_MaintenanceWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmClusterProps",
		reflect.TypeOf((*CfnCloudAutonomousVmClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		reflect.TypeOf((*CfnCloudExadataInfrastructure)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrActivatedStorageCount", GoGetter: "AttrActivatedStorageCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrAdditionalStorageCount", GoGetter: "AttrAdditionalStorageCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailableStorageSizeInGBs", GoGetter: "AttrAvailableStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudExadataInfrastructureArn", GoGetter: "AttrCloudExadataInfrastructureArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudExadataInfrastructureId", GoGetter: "AttrCloudExadataInfrastructureId"},
			_jsii_.MemberProperty{JsiiProperty: "attrComputeModel", GoGetter: "AttrComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "attrCpuCount", GoGetter: "AttrCpuCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataStorageSizeInTBs", GoGetter: "AttrDataStorageSizeInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbNodeStorageSizeInGBs", GoGetter: "AttrDbNodeStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbServerIds", GoGetter: "AttrDbServerIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbServerVersion", GoGetter: "AttrDbServerVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrMaxCpuCount", GoGetter: "AttrMaxCpuCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrMaxDataStorageInTBs", GoGetter: "AttrMaxDataStorageInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrMaxDbNodeStorageSizeInGBs", GoGetter: "AttrMaxDbNodeStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrMaxMemoryInGBs", GoGetter: "AttrMaxMemoryInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemorySizeInGBs", GoGetter: "AttrMemorySizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrOcid", GoGetter: "AttrOcid"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciResourceAnchorName", GoGetter: "AttrOciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciUrl", GoGetter: "AttrOciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrStorageServerVersion", GoGetter: "AttrStorageServerVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalStorageSizeInGBs", GoGetter: "AttrTotalStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneId", GoGetter: "AvailabilityZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureRef", GoGetter: "CloudExadataInfrastructureRef"},
			_jsii_.MemberProperty{JsiiProperty: "computeCount", GoGetter: "ComputeCount"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerContactsToSendToOci", GoGetter: "CustomerContactsToSendToOci"},
			_jsii_.MemberProperty{JsiiProperty: "databaseServerType", GoGetter: "DatabaseServerType"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "shape", GoGetter: "Shape"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageCount", GoGetter: "StorageCount"},
			_jsii_.MemberProperty{JsiiProperty: "storageServerType", GoGetter: "StorageServerType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudExadataInfrastructure{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICloudExadataInfrastructureRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure.CustomerContactProperty",
		reflect.TypeOf((*CfnCloudExadataInfrastructure_CustomerContactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructureProps",
		reflect.TypeOf((*CfnCloudExadataInfrastructureProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		reflect.TypeOf((*CfnCloudVmCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudVmClusterArn", GoGetter: "AttrCloudVmClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudVmClusterId", GoGetter: "AttrCloudVmClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "attrComputeModel", GoGetter: "AttrComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "attrDiskRedundancy", GoGetter: "AttrDiskRedundancy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomain", GoGetter: "AttrDomain"},
			_jsii_.MemberProperty{JsiiProperty: "attrListenerPort", GoGetter: "AttrListenerPort"},
			_jsii_.MemberProperty{JsiiProperty: "attrNodeCount", GoGetter: "AttrNodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrOcid", GoGetter: "AttrOcid"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciResourceAnchorName", GoGetter: "AttrOciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciUrl", GoGetter: "AttrOciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrScanDnsName", GoGetter: "AttrScanDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "attrScanIpIds", GoGetter: "AttrScanIpIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrShape", GoGetter: "AttrShape"},
			_jsii_.MemberProperty{JsiiProperty: "attrStorageSizeInGBs", GoGetter: "AttrStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "attrVipIds", GoGetter: "AttrVipIds"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureId", GoGetter: "CloudExadataInfrastructureId"},
			_jsii_.MemberProperty{JsiiProperty: "cloudVmClusterRef", GoGetter: "CloudVmClusterRef"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataCollectionOptions", GoGetter: "DataCollectionOptions"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTBs", GoGetter: "DataStorageSizeInTBs"},
			_jsii_.MemberProperty{JsiiProperty: "dbNodeStorageSizeInGBs", GoGetter: "DbNodeStorageSizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "dbServers", GoGetter: "DbServers"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "giVersion", GoGetter: "GiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "isLocalBackupEnabled", GoGetter: "IsLocalBackupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isSparseDiskgroupEnabled", GoGetter: "IsSparseDiskgroupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "memorySizeInGBs", GoGetter: "MemorySizeInGBs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkId", GoGetter: "OdbNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTcp", GoGetter: "ScanListenerPortTcp"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sshPublicKeys", GoGetter: "SshPublicKeys"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "systemVersion", GoGetter: "SystemVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudVmCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICloudVmClusterRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster.DataCollectionOptionsProperty",
		reflect.TypeOf((*CfnCloudVmCluster_DataCollectionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnCloudVmClusterProps",
		reflect.TypeOf((*CfnCloudVmClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		reflect.TypeOf((*CfnOdbNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciNetworkAnchorId", GoGetter: "AttrOciNetworkAnchorId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciResourceAnchorName", GoGetter: "AttrOciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "attrOciVcnUrl", GoGetter: "AttrOciVcnUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrOdbNetworkArn", GoGetter: "AttrOdbNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrOdbNetworkId", GoGetter: "AttrOdbNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneId", GoGetter: "AvailabilityZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "backupSubnetCidr", GoGetter: "BackupSubnetCidr"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientSubnetCidr", GoGetter: "ClientSubnetCidr"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDnsPrefix", GoGetter: "DefaultDnsPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAssociatedResources", GoGetter: "DeleteAssociatedResources"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkRef", GoGetter: "OdbNetworkRef"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOdbNetwork{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOdbNetworkRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CfnOdbNetworkProps",
		reflect.TypeOf((*CfnOdbNetworkProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CloudAutonomousVmClusterReference",
		reflect.TypeOf((*CloudAutonomousVmClusterReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CloudExadataInfrastructureReference",
		reflect.TypeOf((*CloudExadataInfrastructureReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.CloudVmClusterReference",
		reflect.TypeOf((*CloudVmClusterReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_odb.ICloudAutonomousVmClusterRef",
		reflect.TypeOf((*ICloudAutonomousVmClusterRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudAutonomousVmClusterRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_odb.ICloudExadataInfrastructureRef",
		reflect.TypeOf((*ICloudExadataInfrastructureRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudExadataInfrastructureRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_odb.ICloudVmClusterRef",
		reflect.TypeOf((*ICloudVmClusterRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudVmClusterRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_odb.IOdbNetworkRef",
		reflect.TypeOf((*IOdbNetworkRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IOdbNetworkRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_odb.OdbNetworkReference",
		reflect.TypeOf((*OdbNetworkReference)(nil)).Elem(),
	)
}

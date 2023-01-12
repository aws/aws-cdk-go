package awsrds

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.AuroraCapacityUnit",
		reflect.TypeOf((*AuroraCapacityUnit)(nil)).Elem(),
		map[string]interface{}{
			"ACU_1": AuroraCapacityUnit_ACU_1,
			"ACU_2": AuroraCapacityUnit_ACU_2,
			"ACU_4": AuroraCapacityUnit_ACU_4,
			"ACU_8": AuroraCapacityUnit_ACU_8,
			"ACU_16": AuroraCapacityUnit_ACU_16,
			"ACU_32": AuroraCapacityUnit_ACU_32,
			"ACU_64": AuroraCapacityUnit_ACU_64,
			"ACU_128": AuroraCapacityUnit_ACU_128,
			"ACU_192": AuroraCapacityUnit_ACU_192,
			"ACU_256": AuroraCapacityUnit_ACU_256,
			"ACU_384": AuroraCapacityUnit_ACU_384,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.AuroraClusterEngineProps",
		reflect.TypeOf((*AuroraClusterEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.AuroraEngineVersion",
		reflect.TypeOf((*AuroraEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auroraFullVersion", GoGetter: "AuroraFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "auroraMajorVersion", GoGetter: "AuroraMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_AuroraEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.AuroraMysqlClusterEngineProps",
		reflect.TypeOf((*AuroraMysqlClusterEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.AuroraMysqlEngineVersion",
		reflect.TypeOf((*AuroraMysqlEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auroraMysqlFullVersion", GoGetter: "AuroraMysqlFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "auroraMysqlMajorVersion", GoGetter: "AuroraMysqlMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_AuroraMysqlEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.AuroraPostgresClusterEngineProps",
		reflect.TypeOf((*AuroraPostgresClusterEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineFeatures",
		reflect.TypeOf((*AuroraPostgresEngineFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.AuroraPostgresEngineVersion",
		reflect.TypeOf((*AuroraPostgresEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auroraPostgresFullVersion", GoGetter: "AuroraPostgresFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "auroraPostgresMajorVersion", GoGetter: "AuroraPostgresMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_AuroraPostgresEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.BackupProps",
		reflect.TypeOf((*BackupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBCluster",
		reflect.TypeOf((*CfnDBCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocatedStorage", GoGetter: "AllocatedStorage"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associatedRoles", GoGetter: "AssociatedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbClusterArn", GoGetter: "AttrDbClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbClusterResourceId", GoGetter: "AttrDbClusterResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointAddress", GoGetter: "AttrEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointPort", GoGetter: "AttrEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "attrReadEndpointAddress", GoGetter: "AttrReadEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgrade", GoGetter: "AutoMinorVersionUpgrade"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "backtrackWindow", GoGetter: "BacktrackWindow"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetentionPeriod", GoGetter: "BackupRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToSnapshot", GoGetter: "CopyTagsToSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterIdentifier", GoGetter: "DbClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterInstanceClass", GoGetter: "DbClusterInstanceClass"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterParameterGroupName", GoGetter: "DbClusterParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceParameterGroupName", GoGetter: "DbInstanceParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupName", GoGetter: "DbSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbSystemId", GoGetter: "DbSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtection", GoGetter: "DeletionProtection"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainIamRoleName", GoGetter: "DomainIamRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "enableCloudwatchLogsExports", GoGetter: "EnableCloudwatchLogsExports"},
			_jsii_.MemberProperty{JsiiProperty: "enableHttpEndpoint", GoGetter: "EnableHttpEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamDatabaseAuthentication", GoGetter: "EnableIamDatabaseAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "engineMode", GoGetter: "EngineMode"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "globalClusterIdentifier", GoGetter: "GlobalClusterIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "masterUsername", GoGetter: "MasterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPassword", GoGetter: "MasterUserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringInterval", GoGetter: "MonitoringInterval"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringRoleArn", GoGetter: "MonitoringRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "networkType", GoGetter: "NetworkType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsEnabled", GoGetter: "PerformanceInsightsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsKmsKeyId", GoGetter: "PerformanceInsightsKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsRetentionPeriod", GoGetter: "PerformanceInsightsRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "preferredBackupWindow", GoGetter: "PreferredBackupWindow"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMaintenanceWindow", GoGetter: "PreferredMaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSourceIdentifier", GoGetter: "ReplicationSourceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "restoreType", GoGetter: "RestoreType"},
			_jsii_.MemberProperty{JsiiProperty: "scalingConfiguration", GoGetter: "ScalingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessV2ScalingConfiguration", GoGetter: "ServerlessV2ScalingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotIdentifier", GoGetter: "SnapshotIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDbClusterIdentifier", GoGetter: "SourceDbClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRegion", GoGetter: "SourceRegion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageEncrypted", GoGetter: "StorageEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "useLatestRestorableTime", GoGetter: "UseLatestRestorableTime"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroupIds", GoGetter: "VpcSecurityGroupIds"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBCluster.DBClusterRoleProperty",
		reflect.TypeOf((*CfnDBCluster_DBClusterRoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBCluster.EndpointProperty",
		reflect.TypeOf((*CfnDBCluster_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBCluster.ReadEndpointProperty",
		reflect.TypeOf((*CfnDBCluster_ReadEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBCluster.ScalingConfigurationProperty",
		reflect.TypeOf((*CfnDBCluster_ScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBCluster.ServerlessV2ScalingConfigurationProperty",
		reflect.TypeOf((*CfnDBCluster_ServerlessV2ScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroup",
		reflect.TypeOf((*CfnDBClusterParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterParameterGroupName", GoGetter: "DbClusterParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
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
			j := jsiiProxy_CfnDBClusterParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBClusterParameterGroupProps",
		reflect.TypeOf((*CfnDBClusterParameterGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBClusterProps",
		reflect.TypeOf((*CfnDBClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		reflect.TypeOf((*CfnDBInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocatedStorage", GoGetter: "AllocatedStorage"},
			_jsii_.MemberProperty{JsiiProperty: "allowMajorVersionUpgrade", GoGetter: "AllowMajorVersionUpgrade"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associatedRoles", GoGetter: "AssociatedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbInstanceArn", GoGetter: "AttrDbInstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbiResourceId", GoGetter: "AttrDbiResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbSystemId", GoGetter: "AttrDbSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointAddress", GoGetter: "AttrEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointHostedZoneId", GoGetter: "AttrEndpointHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointPort", GoGetter: "AttrEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgrade", GoGetter: "AutoMinorVersionUpgrade"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetentionPeriod", GoGetter: "BackupRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificateIdentifier", GoGetter: "CaCertificateIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "characterSetName", GoGetter: "CharacterSetName"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToSnapshot", GoGetter: "CopyTagsToSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customIamInstanceProfile", GoGetter: "CustomIamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterIdentifier", GoGetter: "DbClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterSnapshotIdentifier", GoGetter: "DbClusterSnapshotIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceClass", GoGetter: "DbInstanceClass"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceIdentifier", GoGetter: "DbInstanceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dbName", GoGetter: "DbName"},
			_jsii_.MemberProperty{JsiiProperty: "dbParameterGroupName", GoGetter: "DbParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbSecurityGroups", GoGetter: "DbSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "dbSnapshotIdentifier", GoGetter: "DbSnapshotIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupName", GoGetter: "DbSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAutomatedBackups", GoGetter: "DeleteAutomatedBackups"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtection", GoGetter: "DeletionProtection"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainIamRoleName", GoGetter: "DomainIamRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "enableCloudwatchLogsExports", GoGetter: "EnableCloudwatchLogsExports"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamDatabaseAuthentication", GoGetter: "EnableIamDatabaseAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "enablePerformanceInsights", GoGetter: "EnablePerformanceInsights"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "masterUsername", GoGetter: "MasterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPassword", GoGetter: "MasterUserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "maxAllocatedStorage", GoGetter: "MaxAllocatedStorage"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringInterval", GoGetter: "MonitoringInterval"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringRoleArn", GoGetter: "MonitoringRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "multiAz", GoGetter: "MultiAz"},
			_jsii_.MemberProperty{JsiiProperty: "ncharCharacterSetName", GoGetter: "NcharCharacterSetName"},
			_jsii_.MemberProperty{JsiiProperty: "networkType", GoGetter: "NetworkType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "optionGroupName", GoGetter: "OptionGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsKmsKeyId", GoGetter: "PerformanceInsightsKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsRetentionPeriod", GoGetter: "PerformanceInsightsRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "preferredBackupWindow", GoGetter: "PreferredBackupWindow"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMaintenanceWindow", GoGetter: "PreferredMaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "processorFeatures", GoGetter: "ProcessorFeatures"},
			_jsii_.MemberProperty{JsiiProperty: "promotionTier", GoGetter: "PromotionTier"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "replicaMode", GoGetter: "ReplicaMode"},
			_jsii_.MemberProperty{JsiiProperty: "restoreTime", GoGetter: "RestoreTime"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDbInstanceAutomatedBackupsArn", GoGetter: "SourceDbInstanceAutomatedBackupsArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDbInstanceIdentifier", GoGetter: "SourceDbInstanceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDbiResourceId", GoGetter: "SourceDbiResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRegion", GoGetter: "SourceRegion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageEncrypted", GoGetter: "StorageEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "storageThroughput", GoGetter: "StorageThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timezone", GoGetter: "Timezone"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultProcessorFeatures", GoGetter: "UseDefaultProcessorFeatures"},
			_jsii_.MemberProperty{JsiiProperty: "useLatestRestorableTime", GoGetter: "UseLatestRestorableTime"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroups", GoGetter: "VpcSecurityGroups"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBInstance.DBInstanceRoleProperty",
		reflect.TypeOf((*CfnDBInstance_DBInstanceRoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBInstance.EndpointProperty",
		reflect.TypeOf((*CfnDBInstance_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBInstance.ProcessorFeatureProperty",
		reflect.TypeOf((*CfnDBInstance_ProcessorFeatureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBInstanceProps",
		reflect.TypeOf((*CfnDBInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroup",
		reflect.TypeOf((*CfnDBParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbParameterGroupName", GoGetter: "AttrDbParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbParameterGroupName", GoGetter: "DbParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
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
			j := jsiiProxy_CfnDBParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBParameterGroupProps",
		reflect.TypeOf((*CfnDBParameterGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBProxy",
		reflect.TypeOf((*CfnDBProxy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbProxyArn", GoGetter: "AttrDbProxyArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpoint", GoGetter: "AttrEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcId", GoGetter: "AttrVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "auth", GoGetter: "Auth"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
			_jsii_.MemberProperty{JsiiProperty: "debugLogging", GoGetter: "DebugLogging"},
			_jsii_.MemberProperty{JsiiProperty: "engineFamily", GoGetter: "EngineFamily"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "idleClientTimeout", GoGetter: "IdleClientTimeout"},
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
			_jsii_.MemberProperty{JsiiProperty: "requireTls", GoGetter: "RequireTls"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroupIds", GoGetter: "VpcSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnetIds", GoGetter: "VpcSubnetIds"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBProxy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxy.AuthFormatProperty",
		reflect.TypeOf((*CfnDBProxy_AuthFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxy.TagFormatProperty",
		reflect.TypeOf((*CfnDBProxy_TagFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint",
		reflect.TypeOf((*CfnDBProxyEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDbProxyEndpointArn", GoGetter: "AttrDbProxyEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpoint", GoGetter: "AttrEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcId", GoGetter: "AttrVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyEndpointName", GoGetter: "DbProxyEndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "targetRole", GoGetter: "TargetRole"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecurityGroupIds", GoGetter: "VpcSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnetIds", GoGetter: "VpcSubnetIds"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBProxyEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpoint.TagFormatProperty",
		reflect.TypeOf((*CfnDBProxyEndpoint_TagFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxyEndpointProps",
		reflect.TypeOf((*CfnDBProxyEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxyProps",
		reflect.TypeOf((*CfnDBProxyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup",
		reflect.TypeOf((*CfnDBProxyTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetGroupArn", GoGetter: "AttrTargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionPoolConfigurationInfo", GoGetter: "ConnectionPoolConfigurationInfo"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbClusterIdentifiers", GoGetter: "DbClusterIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceIdentifiers", GoGetter: "DbInstanceIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBProxyTargetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroup.ConnectionPoolConfigurationInfoFormatProperty",
		reflect.TypeOf((*CfnDBProxyTargetGroup_ConnectionPoolConfigurationInfoFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBProxyTargetGroupProps",
		reflect.TypeOf((*CfnDBProxyTargetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup",
		reflect.TypeOf((*CfnDBSecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbSecurityGroupIngress", GoGetter: "DbSecurityGroupIngress"},
			_jsii_.MemberProperty{JsiiProperty: "ec2VpcId", GoGetter: "Ec2VpcId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupDescription", GoGetter: "GroupDescription"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBSecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroup.IngressProperty",
		reflect.TypeOf((*CfnDBSecurityGroup_IngressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngress",
		reflect.TypeOf((*CfnDBSecurityGroupIngress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrip", GoGetter: "Cidrip"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbSecurityGroupName", GoGetter: "DbSecurityGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "ec2SecurityGroupId", GoGetter: "Ec2SecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "ec2SecurityGroupName", GoGetter: "Ec2SecurityGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "ec2SecurityGroupOwnerId", GoGetter: "Ec2SecurityGroupOwnerId"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBSecurityGroupIngress{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupIngressProps",
		reflect.TypeOf((*CfnDBSecurityGroupIngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBSecurityGroupProps",
		reflect.TypeOf((*CfnDBSecurityGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroup",
		reflect.TypeOf((*CfnDBSubnetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupDescription", GoGetter: "DbSubnetGroupDescription"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupName", GoGetter: "DbSubnetGroupName"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDBSubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnDBSubnetGroupProps",
		reflect.TypeOf((*CfnDBSubnetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnEventSubscription",
		reflect.TypeOf((*CfnEventSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "eventCategories", GoGetter: "EventCategories"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicArn", GoGetter: "SnsTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIds", GoGetter: "SourceIds"},
			_jsii_.MemberProperty{JsiiProperty: "sourceType", GoGetter: "SourceType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionName", GoGetter: "SubscriptionName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnEventSubscriptionProps",
		reflect.TypeOf((*CfnEventSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnGlobalCluster",
		reflect.TypeOf((*CfnGlobalCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtection", GoGetter: "DeletionProtection"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "globalClusterIdentifier", GoGetter: "GlobalClusterIdentifier"},
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
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDbClusterIdentifier", GoGetter: "SourceDbClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageEncrypted", GoGetter: "StorageEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnGlobalClusterProps",
		reflect.TypeOf((*CfnGlobalClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.CfnOptionGroup",
		reflect.TypeOf((*CfnOptionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "engineName", GoGetter: "EngineName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "majorEngineVersion", GoGetter: "MajorEngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "optionConfigurations", GoGetter: "OptionConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "optionGroupDescription", GoGetter: "OptionGroupDescription"},
			_jsii_.MemberProperty{JsiiProperty: "optionGroupName", GoGetter: "OptionGroupName"},
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
			j := jsiiProxy_CfnOptionGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnOptionGroup.OptionConfigurationProperty",
		reflect.TypeOf((*CfnOptionGroup_OptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnOptionGroup.OptionSettingProperty",
		reflect.TypeOf((*CfnOptionGroup_OptionSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CfnOptionGroupProps",
		reflect.TypeOf((*CfnOptionGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ClusterEngineBindOptions",
		reflect.TypeOf((*ClusterEngineBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ClusterEngineConfig",
		reflect.TypeOf((*ClusterEngineConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ClusterEngineFeatures",
		reflect.TypeOf((*ClusterEngineFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CommonRotationUserOptions",
		reflect.TypeOf((*CommonRotationUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.Credentials",
		reflect.TypeOf((*Credentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "replicaRegions", GoGetter: "ReplicaRegions"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameAsString", GoGetter: "UsernameAsString"},
		},
		func() interface{} {
			return &jsiiProxy_Credentials{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CredentialsBaseOptions",
		reflect.TypeOf((*CredentialsBaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.CredentialsFromUsernameOptions",
		reflect.TypeOf((*CredentialsFromUsernameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseCluster",
		reflect.TypeOf((*DatabaseCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoints", GoGetter: "InstanceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifiers", GoGetter: "InstanceIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeadlocks", GoMethod: "MetricDeadlocks"},
			_jsii_.MemberMethod{JsiiMethod: "metricEngineUptime", GoMethod: "MetricEngineUptime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeLocalStorage", GoMethod: "MetricFreeLocalStorage"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkReceiveThroughput", GoMethod: "MetricNetworkReceiveThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkThroughput", GoMethod: "MetricNetworkThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkTransmitThroughput", GoMethod: "MetricNetworkTransmitThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricSnapshotStorageUsed", GoMethod: "MetricSnapshotStorageUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalBackupStorageBilled", GoMethod: "MetricTotalBackupStorageBilled"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeBytesUsed", GoMethod: "MetricVolumeBytesUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadIOPs", GoMethod: "MetricVolumeReadIOPs"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteIOPs", GoMethod: "MetricVolumeWriteIOPs"},
			_jsii_.MemberProperty{JsiiProperty: "multiUserRotationApplication", GoGetter: "MultiUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "singleUserRotationApplication", GoGetter: "SingleUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroup", GoGetter: "SubnetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseCluster{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DatabaseClusterBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseClusterAttributes",
		reflect.TypeOf((*DatabaseClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseClusterBase",
		reflect.TypeOf((*DatabaseClusterBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoints", GoGetter: "InstanceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifiers", GoGetter: "InstanceIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeadlocks", GoMethod: "MetricDeadlocks"},
			_jsii_.MemberMethod{JsiiMethod: "metricEngineUptime", GoMethod: "MetricEngineUptime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeLocalStorage", GoMethod: "MetricFreeLocalStorage"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkReceiveThroughput", GoMethod: "MetricNetworkReceiveThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkThroughput", GoMethod: "MetricNetworkThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkTransmitThroughput", GoMethod: "MetricNetworkTransmitThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricSnapshotStorageUsed", GoMethod: "MetricSnapshotStorageUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalBackupStorageBilled", GoMethod: "MetricTotalBackupStorageBilled"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeBytesUsed", GoMethod: "MetricVolumeBytesUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadIOPs", GoMethod: "MetricVolumeReadIOPs"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteIOPs", GoMethod: "MetricVolumeWriteIOPs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseClusterBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseCluster)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		reflect.TypeOf((*DatabaseClusterEngine)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatabaseClusterEngine{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshot",
		reflect.TypeOf((*DatabaseClusterFromSnapshot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoints", GoGetter: "InstanceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifiers", GoGetter: "InstanceIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeadlocks", GoMethod: "MetricDeadlocks"},
			_jsii_.MemberMethod{JsiiMethod: "metricEngineUptime", GoMethod: "MetricEngineUptime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeLocalStorage", GoMethod: "MetricFreeLocalStorage"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkReceiveThroughput", GoMethod: "MetricNetworkReceiveThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkThroughput", GoMethod: "MetricNetworkThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkTransmitThroughput", GoMethod: "MetricNetworkTransmitThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricSnapshotStorageUsed", GoMethod: "MetricSnapshotStorageUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalBackupStorageBilled", GoMethod: "MetricTotalBackupStorageBilled"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeBytesUsed", GoMethod: "MetricVolumeBytesUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadIOPs", GoMethod: "MetricVolumeReadIOPs"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteIOPs", GoMethod: "MetricVolumeWriteIOPs"},
			_jsii_.MemberProperty{JsiiProperty: "multiUserRotationApplication", GoGetter: "MultiUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "singleUserRotationApplication", GoGetter: "SingleUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroup", GoGetter: "SubnetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseClusterFromSnapshot{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DatabaseClusterBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseClusterFromSnapshotProps",
		reflect.TypeOf((*DatabaseClusterFromSnapshotProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseClusterProps",
		reflect.TypeOf((*DatabaseClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		reflect.TypeOf((*DatabaseInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointAddress", GoGetter: "DbInstanceEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointPort", GoGetter: "DbInstanceEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamAuthentication", GoGetter: "EnableIamAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoint", GoGetter: "InstanceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifier", GoGetter: "InstanceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadIOPS", GoMethod: "MetricReadIOPS"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteIOPS", GoMethod: "MetricWriteIOPS"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberMethod{JsiiMethod: "setLogRetention", GoMethod: "SetLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCfnProps", GoGetter: "SourceCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPlacement", GoGetter: "VpcPlacement"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseInstance{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DatabaseInstanceBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseInstance)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceAttributes",
		reflect.TypeOf((*DatabaseInstanceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseInstanceBase",
		reflect.TypeOf((*DatabaseInstanceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointAddress", GoGetter: "DbInstanceEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointPort", GoGetter: "DbInstanceEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamAuthentication", GoGetter: "EnableIamAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoint", GoGetter: "InstanceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifier", GoGetter: "InstanceIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadIOPS", GoMethod: "MetricReadIOPS"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteIOPS", GoMethod: "MetricWriteIOPS"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseInstanceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseInstance)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseInstanceEngine",
		reflect.TypeOf((*DatabaseInstanceEngine)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatabaseInstanceEngine{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshot",
		reflect.TypeOf((*DatabaseInstanceFromSnapshot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointAddress", GoGetter: "DbInstanceEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointPort", GoGetter: "DbInstanceEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamAuthentication", GoGetter: "EnableIamAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoint", GoGetter: "InstanceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifier", GoGetter: "InstanceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadIOPS", GoMethod: "MetricReadIOPS"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteIOPS", GoMethod: "MetricWriteIOPS"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberMethod{JsiiMethod: "setLogRetention", GoMethod: "SetLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCfnProps", GoGetter: "SourceCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPlacement", GoGetter: "VpcPlacement"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseInstanceFromSnapshot{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DatabaseInstanceBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseInstance)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceFromSnapshotProps",
		reflect.TypeOf((*DatabaseInstanceFromSnapshotProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceNewProps",
		reflect.TypeOf((*DatabaseInstanceNewProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceProps",
		reflect.TypeOf((*DatabaseInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplica",
		reflect.TypeOf((*DatabaseInstanceReadReplica)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointAddress", GoGetter: "DbInstanceEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointPort", GoGetter: "DbInstanceEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "enableIamAuthentication", GoGetter: "EnableIamAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoint", GoGetter: "InstanceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifier", GoGetter: "InstanceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadIOPS", GoMethod: "MetricReadIOPS"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteIOPS", GoMethod: "MetricWriteIOPS"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "setLogRetention", GoMethod: "SetLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPlacement", GoGetter: "VpcPlacement"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseInstanceReadReplica{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DatabaseInstanceBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseInstance)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceReadReplicaProps",
		reflect.TypeOf((*DatabaseInstanceReadReplicaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseInstanceSourceProps",
		reflect.TypeOf((*DatabaseInstanceSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseProxy",
		reflect.TypeOf((*DatabaseProxy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyArn", GoGetter: "DbProxyArn"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseProxy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabaseProxy)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecretAttachmentTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseProxyAttributes",
		reflect.TypeOf((*DatabaseProxyAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseProxyOptions",
		reflect.TypeOf((*DatabaseProxyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseProxyProps",
		reflect.TypeOf((*DatabaseProxyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.DatabaseSecret",
		reflect.TypeOf((*DatabaseSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addReplicaRegion", GoMethod: "AddReplicaRegion"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSchedule", GoMethod: "AddRotationSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arnForPolicies", GoGetter: "ArnForPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "attach", GoMethod: "Attach"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "denyAccountRootDelete", GoMethod: "DenyAccountRootDelete"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretFullArn", GoGetter: "SecretFullArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretValue", GoGetter: "SecretValue"},
			_jsii_.MemberMethod{JsiiMethod: "secretValueFromJson", GoMethod: "SecretValueFromJson"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseSecret{}
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerSecret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.DatabaseSecretProps",
		reflect.TypeOf((*DatabaseSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "socketAddress", GoGetter: "SocketAddress"},
		},
		func() interface{} {
			return &jsiiProxy_Endpoint{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.EngineVersion",
		reflect.TypeOf((*EngineVersion)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IClusterEngine",
		reflect.TypeOf((*IClusterEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToCluster", GoMethod: "BindToCluster"},
			_jsii_.MemberProperty{JsiiProperty: "combineImportAndExportRoles", GoGetter: "CombineImportAndExportRoles"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUsername", GoGetter: "DefaultUsername"},
			_jsii_.MemberProperty{JsiiProperty: "engineFamily", GoGetter: "EngineFamily"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "multiUserRotationApplication", GoGetter: "MultiUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroupFamily", GoGetter: "ParameterGroupFamily"},
			_jsii_.MemberProperty{JsiiProperty: "singleUserRotationApplication", GoGetter: "SingleUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "supportedLogTypes", GoGetter: "SupportedLogTypes"},
		},
		func() interface{} {
			j := jsiiProxy_IClusterEngine{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEngine)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IDatabaseCluster",
		reflect.TypeOf((*IDatabaseCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoints", GoGetter: "InstanceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifiers", GoGetter: "InstanceIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricDeadlocks", GoMethod: "MetricDeadlocks"},
			_jsii_.MemberMethod{JsiiMethod: "metricEngineUptime", GoMethod: "MetricEngineUptime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeLocalStorage", GoMethod: "MetricFreeLocalStorage"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkReceiveThroughput", GoMethod: "MetricNetworkReceiveThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkThroughput", GoMethod: "MetricNetworkThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkTransmitThroughput", GoMethod: "MetricNetworkTransmitThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "metricSnapshotStorageUsed", GoMethod: "MetricSnapshotStorageUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalBackupStorageBilled", GoMethod: "MetricTotalBackupStorageBilled"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeBytesUsed", GoMethod: "MetricVolumeBytesUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadIOPs", GoMethod: "MetricVolumeReadIOPs"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteIOPs", GoMethod: "MetricVolumeWriteIOPs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDatabaseCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecretAttachmentTarget)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IDatabaseInstance",
		reflect.TypeOf((*IDatabaseInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProxy", GoMethod: "AddProxy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointAddress", GoGetter: "DbInstanceEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dbInstanceEndpointPort", GoGetter: "DbInstanceEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceArn", GoGetter: "InstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceEndpoint", GoGetter: "InstanceEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdentifier", GoGetter: "InstanceIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDatabaseConnections", GoMethod: "MetricDatabaseConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeableMemory", GoMethod: "MetricFreeableMemory"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadIOPS", GoMethod: "MetricReadIOPS"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteIOPS", GoMethod: "MetricWriteIOPS"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDatabaseInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecretAttachmentTarget)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IDatabaseProxy",
		reflect.TypeOf((*IDatabaseProxy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyArn", GoGetter: "DbProxyArn"},
			_jsii_.MemberProperty{JsiiProperty: "dbProxyName", GoGetter: "DbProxyName"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDatabaseProxy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IEngine",
		reflect.TypeOf((*IEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultUsername", GoGetter: "DefaultUsername"},
			_jsii_.MemberProperty{JsiiProperty: "engineFamily", GoGetter: "EngineFamily"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroupFamily", GoGetter: "ParameterGroupFamily"},
		},
		func() interface{} {
			return &jsiiProxy_IEngine{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IInstanceEngine",
		reflect.TypeOf((*IInstanceEngine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToInstance", GoMethod: "BindToInstance"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUsername", GoGetter: "DefaultUsername"},
			_jsii_.MemberProperty{JsiiProperty: "engineFamily", GoGetter: "EngineFamily"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "multiUserRotationApplication", GoGetter: "MultiUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroupFamily", GoGetter: "ParameterGroupFamily"},
			_jsii_.MemberProperty{JsiiProperty: "singleUserRotationApplication", GoGetter: "SingleUserRotationApplication"},
			_jsii_.MemberProperty{JsiiProperty: "supportsReadReplicaBackups", GoGetter: "SupportsReadReplicaBackups"},
		},
		func() interface{} {
			j := jsiiProxy_IInstanceEngine{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEngine)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IOptionGroup",
		reflect.TypeOf((*IOptionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfiguration", GoMethod: "AddConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optionGroupName", GoGetter: "OptionGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IOptionGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IParameterGroup",
		reflect.TypeOf((*IParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindToCluster", GoMethod: "BindToCluster"},
			_jsii_.MemberMethod{JsiiMethod: "bindToInstance", GoMethod: "BindToInstance"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.IServerlessCluster",
		reflect.TypeOf((*IServerlessCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantDataApiAccess", GoMethod: "GrantDataApiAccess"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IServerlessCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecretAttachmentTarget)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_rds.ISubnetGroup",
		reflect.TypeOf((*ISubnetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroupName", GoGetter: "SubnetGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_ISubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.InstanceEngineBindOptions",
		reflect.TypeOf((*InstanceEngineBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.InstanceEngineConfig",
		reflect.TypeOf((*InstanceEngineConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.InstanceEngineFeatures",
		reflect.TypeOf((*InstanceEngineFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.InstanceProps",
		reflect.TypeOf((*InstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.InstanceUpdateBehaviour",
		reflect.TypeOf((*InstanceUpdateBehaviour)(nil)).Elem(),
		map[string]interface{}{
			"BULK": InstanceUpdateBehaviour_BULK,
			"ROLLING": InstanceUpdateBehaviour_ROLLING,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.LicenseModel",
		reflect.TypeOf((*LicenseModel)(nil)).Elem(),
		map[string]interface{}{
			"LICENSE_INCLUDED": LicenseModel_LICENSE_INCLUDED,
			"BRING_YOUR_OWN_LICENSE": LicenseModel_BRING_YOUR_OWN_LICENSE,
			"GENERAL_PUBLIC_LICENSE": LicenseModel_GENERAL_PUBLIC_LICENSE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.MariaDbEngineVersion",
		reflect.TypeOf((*MariaDbEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "mariaDbFullVersion", GoGetter: "MariaDbFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "mariaDbMajorVersion", GoGetter: "MariaDbMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_MariaDbEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.MariaDbInstanceEngineProps",
		reflect.TypeOf((*MariaDbInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.MySqlInstanceEngineProps",
		reflect.TypeOf((*MySqlInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.MysqlEngineVersion",
		reflect.TypeOf((*MysqlEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "mysqlFullVersion", GoGetter: "MysqlFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlMajorVersion", GoGetter: "MysqlMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_MysqlEngineVersion{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.NetworkType",
		reflect.TypeOf((*NetworkType)(nil)).Elem(),
		map[string]interface{}{
			"IPV4": NetworkType_IPV4,
			"DUAL": NetworkType_DUAL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OptionConfiguration",
		reflect.TypeOf((*OptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.OptionGroup",
		reflect.TypeOf((*OptionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfiguration", GoMethod: "AddConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optionConnections", GoGetter: "OptionConnections"},
			_jsii_.MemberProperty{JsiiProperty: "optionGroupName", GoGetter: "OptionGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OptionGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOptionGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OptionGroupProps",
		reflect.TypeOf((*OptionGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OracleEeCdbInstanceEngineProps",
		reflect.TypeOf((*OracleEeCdbInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OracleEeInstanceEngineProps",
		reflect.TypeOf((*OracleEeInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.OracleEngineVersion",
		reflect.TypeOf((*OracleEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "oracleFullVersion", GoGetter: "OracleFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "oracleMajorVersion", GoGetter: "OracleMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_OracleEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OracleSe2CdbInstanceEngineProps",
		reflect.TypeOf((*OracleSe2CdbInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.OracleSe2InstanceEngineProps",
		reflect.TypeOf((*OracleSe2InstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		reflect.TypeOf((*ParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindToCluster", GoMethod: "BindToCluster"},
			_jsii_.MemberMethod{JsiiMethod: "bindToInstance", GoMethod: "BindToInstance"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IParameterGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ParameterGroupClusterBindOptions",
		reflect.TypeOf((*ParameterGroupClusterBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ParameterGroupClusterConfig",
		reflect.TypeOf((*ParameterGroupClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ParameterGroupInstanceBindOptions",
		reflect.TypeOf((*ParameterGroupInstanceBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ParameterGroupInstanceConfig",
		reflect.TypeOf((*ParameterGroupInstanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ParameterGroupProps",
		reflect.TypeOf((*ParameterGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.PerformanceInsightRetention",
		reflect.TypeOf((*PerformanceInsightRetention)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": PerformanceInsightRetention_DEFAULT,
			"LONG_TERM": PerformanceInsightRetention_LONG_TERM,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.PostgresEngineFeatures",
		reflect.TypeOf((*PostgresEngineFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.PostgresEngineVersion",
		reflect.TypeOf((*PostgresEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "postgresFullVersion", GoGetter: "PostgresFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "postgresMajorVersion", GoGetter: "PostgresMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_PostgresEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.PostgresInstanceEngineProps",
		reflect.TypeOf((*PostgresInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ProcessorFeatures",
		reflect.TypeOf((*ProcessorFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.ProxyTarget",
		reflect.TypeOf((*ProxyTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ProxyTarget{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ProxyTargetConfig",
		reflect.TypeOf((*ProxyTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.RotationMultiUserOptions",
		reflect.TypeOf((*RotationMultiUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.RotationSingleUserOptions",
		reflect.TypeOf((*RotationSingleUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.ServerlessCluster",
		reflect.TypeOf((*ServerlessCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enableDataApi", GoGetter: "EnableDataApi"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDataApiAccess", GoMethod: "GrantDataApiAccess"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServerlessCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerlessCluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ServerlessClusterAttributes",
		reflect.TypeOf((*ServerlessClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshot",
		reflect.TypeOf((*ServerlessClusterFromSnapshot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterReadEndpoint", GoGetter: "ClusterReadEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enableDataApi", GoGetter: "EnableDataApi"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDataApiAccess", GoMethod: "GrantDataApiAccess"},
			_jsii_.MemberProperty{JsiiProperty: "newCfnProps", GoGetter: "NewCfnProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServerlessClusterFromSnapshot{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerlessCluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ServerlessClusterFromSnapshotProps",
		reflect.TypeOf((*ServerlessClusterFromSnapshotProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ServerlessClusterProps",
		reflect.TypeOf((*ServerlessClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.ServerlessScalingOptions",
		reflect.TypeOf((*ServerlessScalingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.SessionPinningFilter",
		reflect.TypeOf((*SessionPinningFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filterName", GoGetter: "FilterName"},
		},
		func() interface{} {
			return &jsiiProxy_SessionPinningFilter{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		reflect.TypeOf((*SnapshotCredentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "generatePassword", GoGetter: "GeneratePassword"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "replaceOnPasswordCriteriaChanges", GoGetter: "ReplaceOnPasswordCriteriaChanges"},
			_jsii_.MemberProperty{JsiiProperty: "replicaRegions", GoGetter: "ReplicaRegions"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			return &jsiiProxy_SnapshotCredentials{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SnapshotCredentialsFromGeneratedPasswordOptions",
		reflect.TypeOf((*SnapshotCredentialsFromGeneratedPasswordOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SqlServerEeInstanceEngineProps",
		reflect.TypeOf((*SqlServerEeInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.SqlServerEngineVersion",
		reflect.TypeOf((*SqlServerEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "sqlServerFullVersion", GoGetter: "SqlServerFullVersion"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerMajorVersion", GoGetter: "SqlServerMajorVersion"},
		},
		func() interface{} {
			return &jsiiProxy_SqlServerEngineVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SqlServerExInstanceEngineProps",
		reflect.TypeOf((*SqlServerExInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SqlServerSeInstanceEngineProps",
		reflect.TypeOf((*SqlServerSeInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SqlServerWebInstanceEngineProps",
		reflect.TypeOf((*SqlServerWebInstanceEngineProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_rds.StorageType",
		reflect.TypeOf((*StorageType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": StorageType_STANDARD,
			"GP2": StorageType_GP2,
			"GP3": StorageType_GP3,
			"IO1": StorageType_IO1,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_rds.SubnetGroup",
		reflect.TypeOf((*SubnetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroupName", GoGetter: "SubnetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnetGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_rds.SubnetGroupProps",
		reflect.TypeOf((*SubnetGroupProps)(nil)).Elem(),
	)
}

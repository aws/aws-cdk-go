package awsquicksight

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		reflect.TypeOf((*CfnAnalysis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "analysisId", GoGetter: "AnalysisId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSetArns", GoGetter: "AttrDataSetArns"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrSheets", GoGetter: "AttrSheets"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errors", GoGetter: "Errors"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceEntity", GoGetter: "SourceEntity"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "themeArn", GoGetter: "ThemeArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnalysis{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.AnalysisErrorProperty",
		reflect.TypeOf((*CfnAnalysis_AnalysisErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.AnalysisSourceEntityProperty",
		reflect.TypeOf((*CfnAnalysis_AnalysisSourceEntityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.AnalysisSourceTemplateProperty",
		reflect.TypeOf((*CfnAnalysis_AnalysisSourceTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.DataSetReferenceProperty",
		reflect.TypeOf((*CfnAnalysis_DataSetReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.DateTimeParameterProperty",
		reflect.TypeOf((*CfnAnalysis_DateTimeParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.DecimalParameterProperty",
		reflect.TypeOf((*CfnAnalysis_DecimalParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.IntegerParameterProperty",
		reflect.TypeOf((*CfnAnalysis_IntegerParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.ParametersProperty",
		reflect.TypeOf((*CfnAnalysis_ParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.ResourcePermissionProperty",
		reflect.TypeOf((*CfnAnalysis_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.SheetProperty",
		reflect.TypeOf((*CfnAnalysis_SheetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis.StringParameterProperty",
		reflect.TypeOf((*CfnAnalysis_StringParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnAnalysisProps",
		reflect.TypeOf((*CfnAnalysisProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		reflect.TypeOf((*CfnDashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastPublishedTime", GoGetter: "AttrLastPublishedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionArn", GoGetter: "AttrVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionCreatedTime", GoGetter: "AttrVersionCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionDataSetArns", GoGetter: "AttrVersionDataSetArns"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionDescription", GoGetter: "AttrVersionDescription"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionErrors", GoGetter: "AttrVersionErrors"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionSheets", GoGetter: "AttrVersionSheets"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionSourceEntityArn", GoGetter: "AttrVersionSourceEntityArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionStatus", GoGetter: "AttrVersionStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionThemeArn", GoGetter: "AttrVersionThemeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionVersionNumber", GoGetter: "AttrVersionVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardId", GoGetter: "DashboardId"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardPublishOptions", GoGetter: "DashboardPublishOptions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceEntity", GoGetter: "SourceEntity"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "themeArn", GoGetter: "ThemeArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.AdHocFilteringOptionProperty",
		reflect.TypeOf((*CfnDashboard_AdHocFilteringOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DashboardErrorProperty",
		reflect.TypeOf((*CfnDashboard_DashboardErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DashboardPublishOptionsProperty",
		reflect.TypeOf((*CfnDashboard_DashboardPublishOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DashboardSourceEntityProperty",
		reflect.TypeOf((*CfnDashboard_DashboardSourceEntityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DashboardSourceTemplateProperty",
		reflect.TypeOf((*CfnDashboard_DashboardSourceTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DashboardVersionProperty",
		reflect.TypeOf((*CfnDashboard_DashboardVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DataSetReferenceProperty",
		reflect.TypeOf((*CfnDashboard_DataSetReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DateTimeParameterProperty",
		reflect.TypeOf((*CfnDashboard_DateTimeParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.DecimalParameterProperty",
		reflect.TypeOf((*CfnDashboard_DecimalParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.ExportToCSVOptionProperty",
		reflect.TypeOf((*CfnDashboard_ExportToCSVOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.IntegerParameterProperty",
		reflect.TypeOf((*CfnDashboard_IntegerParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.ParametersProperty",
		reflect.TypeOf((*CfnDashboard_ParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.ResourcePermissionProperty",
		reflect.TypeOf((*CfnDashboard_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.SheetControlsOptionProperty",
		reflect.TypeOf((*CfnDashboard_SheetControlsOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.SheetProperty",
		reflect.TypeOf((*CfnDashboard_SheetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboard.StringParameterProperty",
		reflect.TypeOf((*CfnDashboard_StringParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDashboardProps",
		reflect.TypeOf((*CfnDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		reflect.TypeOf((*CfnDataSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrConsumedSpiceCapacityInBytes", GoGetter: "AttrConsumedSpiceCapacityInBytes"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrOutputColumns", GoGetter: "AttrOutputColumns"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "columnGroups", GoGetter: "ColumnGroups"},
			_jsii_.MemberProperty{JsiiProperty: "columnLevelPermissionRules", GoGetter: "ColumnLevelPermissionRules"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataSetId", GoGetter: "DataSetId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSetUsageConfiguration", GoGetter: "DataSetUsageConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "fieldFolders", GoGetter: "FieldFolders"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "importMode", GoGetter: "ImportMode"},
			_jsii_.MemberProperty{JsiiProperty: "ingestionWaitPolicy", GoGetter: "IngestionWaitPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalTableMap", GoGetter: "LogicalTableMap"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "physicalTableMap", GoGetter: "PhysicalTableMap"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "rowLevelPermissionDataSet", GoGetter: "RowLevelPermissionDataSet"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.CalculatedColumnProperty",
		reflect.TypeOf((*CfnDataSet_CalculatedColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.CastColumnTypeOperationProperty",
		reflect.TypeOf((*CfnDataSet_CastColumnTypeOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ColumnDescriptionProperty",
		reflect.TypeOf((*CfnDataSet_ColumnDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ColumnGroupProperty",
		reflect.TypeOf((*CfnDataSet_ColumnGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ColumnLevelPermissionRuleProperty",
		reflect.TypeOf((*CfnDataSet_ColumnLevelPermissionRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ColumnTagProperty",
		reflect.TypeOf((*CfnDataSet_ColumnTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.CreateColumnsOperationProperty",
		reflect.TypeOf((*CfnDataSet_CreateColumnsOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.CustomSqlProperty",
		reflect.TypeOf((*CfnDataSet_CustomSqlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.DataSetUsageConfigurationProperty",
		reflect.TypeOf((*CfnDataSet_DataSetUsageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.FieldFolderProperty",
		reflect.TypeOf((*CfnDataSet_FieldFolderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.FilterOperationProperty",
		reflect.TypeOf((*CfnDataSet_FilterOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.GeoSpatialColumnGroupProperty",
		reflect.TypeOf((*CfnDataSet_GeoSpatialColumnGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.IngestionWaitPolicyProperty",
		reflect.TypeOf((*CfnDataSet_IngestionWaitPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.InputColumnProperty",
		reflect.TypeOf((*CfnDataSet_InputColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.JoinInstructionProperty",
		reflect.TypeOf((*CfnDataSet_JoinInstructionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.JoinKeyPropertiesProperty",
		reflect.TypeOf((*CfnDataSet_JoinKeyPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.LogicalTableProperty",
		reflect.TypeOf((*CfnDataSet_LogicalTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.LogicalTableSourceProperty",
		reflect.TypeOf((*CfnDataSet_LogicalTableSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.OutputColumnProperty",
		reflect.TypeOf((*CfnDataSet_OutputColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.PhysicalTableProperty",
		reflect.TypeOf((*CfnDataSet_PhysicalTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ProjectOperationProperty",
		reflect.TypeOf((*CfnDataSet_ProjectOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.RelationalTableProperty",
		reflect.TypeOf((*CfnDataSet_RelationalTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.RenameColumnOperationProperty",
		reflect.TypeOf((*CfnDataSet_RenameColumnOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.ResourcePermissionProperty",
		reflect.TypeOf((*CfnDataSet_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.RowLevelPermissionDataSetProperty",
		reflect.TypeOf((*CfnDataSet_RowLevelPermissionDataSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.S3SourceProperty",
		reflect.TypeOf((*CfnDataSet_S3SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.TagColumnOperationProperty",
		reflect.TypeOf((*CfnDataSet_TagColumnOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.TransformOperationProperty",
		reflect.TypeOf((*CfnDataSet_TransformOperationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSet.UploadSettingsProperty",
		reflect.TypeOf((*CfnDataSet_UploadSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSetProps",
		reflect.TypeOf((*CfnDataSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		reflect.TypeOf((*CfnDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alternateDataSourceParameters", GoGetter: "AlternateDataSourceParameters"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceParameters", GoGetter: "DataSourceParameters"},
			_jsii_.MemberProperty{JsiiProperty: "errorInfo", GoGetter: "ErrorInfo"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sslProperties", GoGetter: "SslProperties"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectionProperties", GoGetter: "VpcConnectionProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.AmazonElasticsearchParametersProperty",
		reflect.TypeOf((*CfnDataSource_AmazonElasticsearchParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.AmazonOpenSearchParametersProperty",
		reflect.TypeOf((*CfnDataSource_AmazonOpenSearchParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.AthenaParametersProperty",
		reflect.TypeOf((*CfnDataSource_AthenaParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.AuroraParametersProperty",
		reflect.TypeOf((*CfnDataSource_AuroraParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.AuroraPostgreSqlParametersProperty",
		reflect.TypeOf((*CfnDataSource_AuroraPostgreSqlParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.CredentialPairProperty",
		reflect.TypeOf((*CfnDataSource_CredentialPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.DataSourceCredentialsProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.DataSourceErrorInfoProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.DataSourceParametersProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.DatabricksParametersProperty",
		reflect.TypeOf((*CfnDataSource_DatabricksParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.ManifestFileLocationProperty",
		reflect.TypeOf((*CfnDataSource_ManifestFileLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.MariaDbParametersProperty",
		reflect.TypeOf((*CfnDataSource_MariaDbParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.MySqlParametersProperty",
		reflect.TypeOf((*CfnDataSource_MySqlParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.OracleParametersProperty",
		reflect.TypeOf((*CfnDataSource_OracleParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.PostgreSqlParametersProperty",
		reflect.TypeOf((*CfnDataSource_PostgreSqlParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.PrestoParametersProperty",
		reflect.TypeOf((*CfnDataSource_PrestoParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.RdsParametersProperty",
		reflect.TypeOf((*CfnDataSource_RdsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.RedshiftParametersProperty",
		reflect.TypeOf((*CfnDataSource_RedshiftParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.ResourcePermissionProperty",
		reflect.TypeOf((*CfnDataSource_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.S3ParametersProperty",
		reflect.TypeOf((*CfnDataSource_S3ParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.SnowflakeParametersProperty",
		reflect.TypeOf((*CfnDataSource_SnowflakeParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.SparkParametersProperty",
		reflect.TypeOf((*CfnDataSource_SparkParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.SqlServerParametersProperty",
		reflect.TypeOf((*CfnDataSource_SqlServerParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.SslPropertiesProperty",
		reflect.TypeOf((*CfnDataSource_SslPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.TeradataParametersProperty",
		reflect.TypeOf((*CfnDataSource_TeradataParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSource.VpcConnectionPropertiesProperty",
		reflect.TypeOf((*CfnDataSource_VpcConnectionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnDataSourceProps",
		reflect.TypeOf((*CfnDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		reflect.TypeOf((*CfnTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionCreatedTime", GoGetter: "AttrVersionCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionDataSetConfigurations", GoGetter: "AttrVersionDataSetConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionDescription", GoGetter: "AttrVersionDescription"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionErrors", GoGetter: "AttrVersionErrors"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionSheets", GoGetter: "AttrVersionSheets"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionSourceEntityArn", GoGetter: "AttrVersionSourceEntityArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionStatus", GoGetter: "AttrVersionStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionThemeArn", GoGetter: "AttrVersionThemeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionVersionNumber", GoGetter: "AttrVersionVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceEntity", GoGetter: "SourceEntity"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateId", GoGetter: "TemplateId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.ColumnGroupColumnSchemaProperty",
		reflect.TypeOf((*CfnTemplate_ColumnGroupColumnSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.ColumnGroupSchemaProperty",
		reflect.TypeOf((*CfnTemplate_ColumnGroupSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.ColumnSchemaProperty",
		reflect.TypeOf((*CfnTemplate_ColumnSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.DataSetConfigurationProperty",
		reflect.TypeOf((*CfnTemplate_DataSetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.DataSetReferenceProperty",
		reflect.TypeOf((*CfnTemplate_DataSetReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.DataSetSchemaProperty",
		reflect.TypeOf((*CfnTemplate_DataSetSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.ResourcePermissionProperty",
		reflect.TypeOf((*CfnTemplate_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.SheetProperty",
		reflect.TypeOf((*CfnTemplate_SheetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.TemplateErrorProperty",
		reflect.TypeOf((*CfnTemplate_TemplateErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.TemplateSourceAnalysisProperty",
		reflect.TypeOf((*CfnTemplate_TemplateSourceAnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.TemplateSourceEntityProperty",
		reflect.TypeOf((*CfnTemplate_TemplateSourceEntityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.TemplateSourceTemplateProperty",
		reflect.TypeOf((*CfnTemplate_TemplateSourceTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplate.TemplateVersionProperty",
		reflect.TypeOf((*CfnTemplate_TemplateVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTemplateProps",
		reflect.TypeOf((*CfnTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		reflect.TypeOf((*CfnTheme)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedTime", GoGetter: "AttrCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionArn", GoGetter: "AttrVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionBaseThemeId", GoGetter: "AttrVersionBaseThemeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionCreatedTime", GoGetter: "AttrVersionCreatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionDescription", GoGetter: "AttrVersionDescription"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionErrors", GoGetter: "AttrVersionErrors"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionStatus", GoGetter: "AttrVersionStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionVersionNumber", GoGetter: "AttrVersionVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "baseThemeId", GoGetter: "BaseThemeId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "themeId", GoGetter: "ThemeId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTheme{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.BorderStyleProperty",
		reflect.TypeOf((*CfnTheme_BorderStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.DataColorPaletteProperty",
		reflect.TypeOf((*CfnTheme_DataColorPaletteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.FontProperty",
		reflect.TypeOf((*CfnTheme_FontProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.GutterStyleProperty",
		reflect.TypeOf((*CfnTheme_GutterStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.MarginStyleProperty",
		reflect.TypeOf((*CfnTheme_MarginStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.ResourcePermissionProperty",
		reflect.TypeOf((*CfnTheme_ResourcePermissionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.SheetStyleProperty",
		reflect.TypeOf((*CfnTheme_SheetStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.ThemeConfigurationProperty",
		reflect.TypeOf((*CfnTheme_ThemeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.ThemeErrorProperty",
		reflect.TypeOf((*CfnTheme_ThemeErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.ThemeVersionProperty",
		reflect.TypeOf((*CfnTheme_ThemeVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.TileLayoutStyleProperty",
		reflect.TypeOf((*CfnTheme_TileLayoutStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.TileStyleProperty",
		reflect.TypeOf((*CfnTheme_TileStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.TypographyProperty",
		reflect.TypeOf((*CfnTheme_TypographyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnTheme.UIColorPaletteProperty",
		reflect.TypeOf((*CfnTheme_UIColorPaletteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_quicksight.CfnThemeProps",
		reflect.TypeOf((*CfnThemeProps)(nil)).Elem(),
	)
}

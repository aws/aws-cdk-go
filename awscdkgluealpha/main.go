// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		reflect.TypeOf((*AssetCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AssetCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		reflect.TypeOf((*ClassificationString)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ClassificationString{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.CloudWatchEncryption",
		reflect.TypeOf((*CloudWatchEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.CloudWatchEncryptionMode",
		reflect.TypeOf((*CloudWatchEncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"KMS": CloudWatchEncryptionMode_KMS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.Column",
		reflect.TypeOf((*Column)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.ColumnCountMismatchHandlingAction",
		reflect.TypeOf((*ColumnCountMismatchHandlingAction)(nil)).Elem(),
		map[string]interface{}{
			"DISABLED": ColumnCountMismatchHandlingAction_DISABLED,
			"FAIL": ColumnCountMismatchHandlingAction_FAIL,
			"SET_TO_NULL": ColumnCountMismatchHandlingAction_SET_TO_NULL,
			"DROP_ROW": ColumnCountMismatchHandlingAction_DROP_ROW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.CompressionType",
		reflect.TypeOf((*CompressionType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompressionType_NONE,
			"BZIP2": CompressionType_BZIP2,
			"GZIP": CompressionType_GZIP,
			"SNAPPY": CompressionType_SNAPPY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Connection",
		reflect.TypeOf((*Connection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProperty", GoMethod: "AddProperty"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionName", GoGetter: "ConnectionName"},
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
			j := jsiiProxy_Connection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnection)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.ConnectionOptions",
		reflect.TypeOf((*ConnectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.ConnectionProps",
		reflect.TypeOf((*ConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		reflect.TypeOf((*ConnectionType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ConnectionType{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.ContinuousLoggingProps",
		reflect.TypeOf((*ContinuousLoggingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.DataFormat",
		reflect.TypeOf((*DataFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "classificationString", GoGetter: "ClassificationString"},
			_jsii_.MemberProperty{JsiiProperty: "inputFormat", GoGetter: "InputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "serializationLibrary", GoGetter: "SerializationLibrary"},
		},
		func() interface{} {
			return &jsiiProxy_DataFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.DataFormatProps",
		reflect.TypeOf((*DataFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.DataQualityRuleset",
		reflect.TypeOf((*DataQualityRuleset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "rulesetArn", GoGetter: "RulesetArn"},
			_jsii_.MemberProperty{JsiiProperty: "rulesetName", GoGetter: "RulesetName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataQualityRuleset{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataQualityRuleset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.DataQualityRulesetProps",
		reflect.TypeOf((*DataQualityRulesetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.DataQualityTargetTable",
		reflect.TypeOf((*DataQualityTargetTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
		},
		func() interface{} {
			return &jsiiProxy_DataQualityTargetTable{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Database",
		reflect.TypeOf((*Database)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "catalogArn", GoGetter: "CatalogArn"},
			_jsii_.MemberProperty{JsiiProperty: "catalogId", GoGetter: "CatalogId"},
			_jsii_.MemberProperty{JsiiProperty: "databaseArn", GoGetter: "DatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "locationUri", GoGetter: "LocationUri"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Database{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDatabase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.DatabaseProps",
		reflect.TypeOf((*DatabaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.ExecutionClass",
		reflect.TypeOf((*ExecutionClass)(nil)).Elem(),
		map[string]interface{}{
			"FLEX": ExecutionClass_FLEX,
			"STANDARD": ExecutionClass_STANDARD,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.ExternalTable",
		reflect.TypeOf((*ExternalTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPartitionIndex", GoMethod: "AddPartitionIndex"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantToUnderlyingResources", GoMethod: "GrantToUnderlyingResources"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitionIndexes", GoGetter: "PartitionIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "partitionKeys", GoGetter: "PartitionKeys"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageParameters", GoGetter: "StorageParameters"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableResource", GoGetter: "TableResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalTable{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TableBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.ExternalTableProps",
		reflect.TypeOf((*ExternalTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		reflect.TypeOf((*GlueVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_GlueVersion{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.IConnection",
		reflect.TypeOf((*IConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionName", GoGetter: "ConnectionName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.IDataQualityRuleset",
		reflect.TypeOf((*IDataQualityRuleset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "rulesetArn", GoGetter: "RulesetArn"},
			_jsii_.MemberProperty{JsiiProperty: "rulesetName", GoGetter: "RulesetName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDataQualityRuleset{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.IDatabase",
		reflect.TypeOf((*IDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "catalogArn", GoGetter: "CatalogArn"},
			_jsii_.MemberProperty{JsiiProperty: "catalogId", GoGetter: "CatalogId"},
			_jsii_.MemberProperty{JsiiProperty: "databaseArn", GoGetter: "DatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.IJob",
		reflect.TypeOf((*IJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "jobArn", GoGetter: "JobArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobName", GoGetter: "JobName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailure", GoMethod: "MetricFailure"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccess", GoMethod: "MetricSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimeout", GoMethod: "MetricTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onFailure", GoMethod: "OnFailure"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onSuccess", GoMethod: "OnSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "onTimeout", GoMethod: "OnTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IJob{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.ISecurityConfiguration",
		reflect.TypeOf((*ISecurityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigurationName", GoGetter: "SecurityConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ISecurityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-glue-alpha.ITable",
		reflect.TypeOf((*ITable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
		},
		func() interface{} {
			j := jsiiProxy_ITable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.InputFormat",
		reflect.TypeOf((*InputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_InputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.InvalidCharHandlingAction",
		reflect.TypeOf((*InvalidCharHandlingAction)(nil)).Elem(),
		map[string]interface{}{
			"DISABLED": InvalidCharHandlingAction_DISABLED,
			"FAIL": InvalidCharHandlingAction_FAIL,
			"SET_TO_NULL": InvalidCharHandlingAction_SET_TO_NULL,
			"DROP_ROW": InvalidCharHandlingAction_DROP_ROW,
			"REPLACE": InvalidCharHandlingAction_REPLACE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "jobArn", GoGetter: "JobArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobName", GoGetter: "JobName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailure", GoMethod: "MetricFailure"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccess", GoMethod: "MetricSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimeout", GoMethod: "MetricTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onFailure", GoMethod: "OnFailure"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onSuccess", GoMethod: "OnSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "onTimeout", GoMethod: "OnTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "sparkUILoggingLocation", GoGetter: "SparkUILoggingLocation"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Job{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJob)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.JobAttributes",
		reflect.TypeOf((*JobAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.JobBookmarksEncryption",
		reflect.TypeOf((*JobBookmarksEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.JobBookmarksEncryptionMode",
		reflect.TypeOf((*JobBookmarksEncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"CLIENT_SIDE_KMS": JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		reflect.TypeOf((*JobExecutable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_JobExecutable{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.JobExecutableConfig",
		reflect.TypeOf((*JobExecutableConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.JobLanguage",
		reflect.TypeOf((*JobLanguage)(nil)).Elem(),
		map[string]interface{}{
			"SCALA": JobLanguage_SCALA,
			"PYTHON": JobLanguage_PYTHON,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.JobState",
		reflect.TypeOf((*JobState)(nil)).Elem(),
		map[string]interface{}{
			"SUCCEEDED": JobState_SUCCEEDED,
			"FAILED": JobState_FAILED,
			"TIMEOUT": JobState_TIMEOUT,
			"STARTING": JobState_STARTING,
			"RUNNING": JobState_RUNNING,
			"STOPPING": JobState_STOPPING,
			"STOPPED": JobState_STOPPED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.JobType",
		reflect.TypeOf((*JobType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_JobType{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.MetricType",
		reflect.TypeOf((*MetricType)(nil)).Elem(),
		map[string]interface{}{
			"GAUGE": MetricType_GAUGE,
			"COUNT": MetricType_COUNT,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.NumericOverflowHandlingAction",
		reflect.TypeOf((*NumericOverflowHandlingAction)(nil)).Elem(),
		map[string]interface{}{
			"DISABLED": NumericOverflowHandlingAction_DISABLED,
			"FAIL": NumericOverflowHandlingAction_FAIL,
			"SET_TO_NULL": NumericOverflowHandlingAction_SET_TO_NULL,
			"DROP_ROW": NumericOverflowHandlingAction_DROP_ROW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.OrcColumnMappingType",
		reflect.TypeOf((*OrcColumnMappingType)(nil)).Elem(),
		map[string]interface{}{
			"NAME": OrcColumnMappingType_NAME,
			"POSITION": OrcColumnMappingType_POSITION,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.OutputFormat",
		reflect.TypeOf((*OutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_OutputFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.PartitionIndex",
		reflect.TypeOf((*PartitionIndex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.PythonRayExecutableProps",
		reflect.TypeOf((*PythonRayExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.PythonShellExecutableProps",
		reflect.TypeOf((*PythonShellExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.PythonSparkJobExecutableProps",
		reflect.TypeOf((*PythonSparkJobExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.PythonVersion",
		reflect.TypeOf((*PythonVersion)(nil)).Elem(),
		map[string]interface{}{
			"TWO": PythonVersion_TWO,
			"THREE": PythonVersion_THREE,
			"THREE_NINE": PythonVersion_THREE_NINE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.S3Code",
		reflect.TypeOf((*S3Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Code{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.S3Encryption",
		reflect.TypeOf((*S3Encryption)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.S3EncryptionMode",
		reflect.TypeOf((*S3EncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"S3_MANAGED": S3EncryptionMode_S3_MANAGED,
			"KMS": S3EncryptionMode_KMS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.S3Table",
		reflect.TypeOf((*S3Table)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPartitionIndex", GoMethod: "AddPartitionIndex"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "generateS3PrefixForGrant", GoMethod: "GenerateS3PrefixForGrant"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantToUnderlyingResources", GoMethod: "GrantToUnderlyingResources"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitionIndexes", GoGetter: "PartitionIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "partitionKeys", GoGetter: "PartitionKeys"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "s3Prefix", GoGetter: "S3Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageParameters", GoGetter: "StorageParameters"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableResource", GoGetter: "TableResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3Table{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TableBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.S3TableProps",
		reflect.TypeOf((*S3TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.ScalaJobExecutableProps",
		reflect.TypeOf((*ScalaJobExecutableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Schema",
		reflect.TypeOf((*Schema)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Schema{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.SecurityConfiguration",
		reflect.TypeOf((*SecurityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchEncryptionKey", GoGetter: "CloudWatchEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobBookmarksEncryptionKey", GoGetter: "JobBookmarksEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "s3EncryptionKey", GoGetter: "S3EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigurationName", GoGetter: "SecurityConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecurityConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.SecurityConfigurationProps",
		reflect.TypeOf((*SecurityConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.SerializationLibrary",
		reflect.TypeOf((*SerializationLibrary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
		},
		func() interface{} {
			return &jsiiProxy_SerializationLibrary{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.SparkUILoggingLocation",
		reflect.TypeOf((*SparkUILoggingLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.SparkUIProps",
		reflect.TypeOf((*SparkUIProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		reflect.TypeOf((*StorageParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_StorageParameter{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.StorageParameters",
		reflect.TypeOf((*StorageParameters)(nil)).Elem(),
		map[string]interface{}{
			"SKIP_HEADER_LINE_COUNT": StorageParameters_SKIP_HEADER_LINE_COUNT,
			"DATA_CLEANSING_ENABLED": StorageParameters_DATA_CLEANSING_ENABLED,
			"COMPRESSION_TYPE": StorageParameters_COMPRESSION_TYPE,
			"INVALID_CHAR_HANDLING": StorageParameters_INVALID_CHAR_HANDLING,
			"REPLACEMENT_CHAR": StorageParameters_REPLACEMENT_CHAR,
			"NUMERIC_OVERFLOW_HANDLING": StorageParameters_NUMERIC_OVERFLOW_HANDLING,
			"SURPLUS_BYTES_HANDLING": StorageParameters_SURPLUS_BYTES_HANDLING,
			"SURPLUS_CHAR_HANDLING": StorageParameters_SURPLUS_CHAR_HANDLING,
			"COLUMN_COUNT_MISMATCH_HANDLING": StorageParameters_COLUMN_COUNT_MISMATCH_HANDLING,
			"NUM_ROWS": StorageParameters_NUM_ROWS,
			"SERIALIZATION_NULL_FORMAT": StorageParameters_SERIALIZATION_NULL_FORMAT,
			"ORC_SCHEMA_RESOLUTION": StorageParameters_ORC_SCHEMA_RESOLUTION,
			"WRITE_PARALLEL": StorageParameters_WRITE_PARALLEL,
			"WRITE_MAX_FILESIZE_MB": StorageParameters_WRITE_MAX_FILESIZE_MB,
			"WRITE_KMS_KEY_ID": StorageParameters_WRITE_KMS_KEY_ID,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.SurplusBytesHandlingAction",
		reflect.TypeOf((*SurplusBytesHandlingAction)(nil)).Elem(),
		map[string]interface{}{
			"SET_TO_NULL": SurplusBytesHandlingAction_SET_TO_NULL,
			"DISABLED": SurplusBytesHandlingAction_DISABLED,
			"FAIL": SurplusBytesHandlingAction_FAIL,
			"DROP_ROW": SurplusBytesHandlingAction_DROP_ROW,
			"TRUNCATE": SurplusBytesHandlingAction_TRUNCATE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.SurplusCharHandlingAction",
		reflect.TypeOf((*SurplusCharHandlingAction)(nil)).Elem(),
		map[string]interface{}{
			"SET_TO_NULL": SurplusCharHandlingAction_SET_TO_NULL,
			"DISABLED": SurplusCharHandlingAction_DISABLED,
			"FAIL": SurplusCharHandlingAction_FAIL,
			"DROP_ROW": SurplusCharHandlingAction_DROP_ROW,
			"TRUNCATE": SurplusCharHandlingAction_TRUNCATE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.Table",
		reflect.TypeOf((*Table)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPartitionIndex", GoMethod: "AddPartitionIndex"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "generateS3PrefixForGrant", GoMethod: "GenerateS3PrefixForGrant"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantToUnderlyingResources", GoMethod: "GrantToUnderlyingResources"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitionIndexes", GoGetter: "PartitionIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "partitionKeys", GoGetter: "PartitionKeys"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "s3Prefix", GoGetter: "S3Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageParameters", GoGetter: "StorageParameters"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableResource", GoGetter: "TableResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Table{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_S3Table)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.TableAttributes",
		reflect.TypeOf((*TableAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.TableBase",
		reflect.TypeOf((*TableBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPartitionIndex", GoMethod: "AddPartitionIndex"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "compressed", GoGetter: "Compressed"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormat", GoGetter: "DataFormat"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantToUnderlyingResources", GoMethod: "GrantToUnderlyingResources"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitionIndexes", GoGetter: "PartitionIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "partitionKeys", GoGetter: "PartitionKeys"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageParameters", GoGetter: "StorageParameters"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableResource", GoGetter: "TableResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TableBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.TableBaseProps",
		reflect.TypeOf((*TableBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.TableEncryption",
		reflect.TypeOf((*TableEncryption)(nil)).Elem(),
		map[string]interface{}{
			"S3_MANAGED": TableEncryption_S3_MANAGED,
			"KMS": TableEncryption_KMS,
			"KMS_MANAGED": TableEncryption_KMS_MANAGED,
			"CLIENT_SIDE_KMS": TableEncryption_CLIENT_SIDE_KMS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.TableProps",
		reflect.TypeOf((*TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-glue-alpha.Type",
		reflect.TypeOf((*Type)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		reflect.TypeOf((*WorkerType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_WorkerType{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-glue-alpha.WriteParallel",
		reflect.TypeOf((*WriteParallel)(nil)).Elem(),
		map[string]interface{}{
			"ON": WriteParallel_ON,
			"OFF": WriteParallel_OFF,
		},
	)
}

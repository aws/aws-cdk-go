package awsdynamodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.Attribute",
		reflect.TypeOf((*Attribute)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.AttributeType",
		reflect.TypeOf((*AttributeType)(nil)).Elem(),
		map[string]interface{}{
			"BINARY": AttributeType_BINARY,
			"NUMBER": AttributeType_NUMBER,
			"STRING": AttributeType_STRING,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.BillingMode",
		reflect.TypeOf((*BillingMode)(nil)).Elem(),
		map[string]interface{}{
			"PAY_PER_REQUEST": BillingMode_PAY_PER_REQUEST,
			"PROVISIONED": BillingMode_PROVISIONED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		reflect.TypeOf((*CfnGlobalTable)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attributeDefinitions", GoGetter: "AttributeDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "attrStreamArn", GoGetter: "AttrStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTableId", GoGetter: "AttrTableId"},
			_jsii_.MemberProperty{JsiiProperty: "billingMode", GoGetter: "BillingMode"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "globalSecondaryIndexes", GoGetter: "GlobalSecondaryIndexes"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keySchema", GoGetter: "KeySchema"},
			_jsii_.MemberProperty{JsiiProperty: "localSecondaryIndexes", GoGetter: "LocalSecondaryIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "replicas", GoGetter: "Replicas"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamSpecification", GoGetter: "StreamSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "timeToLiveSpecification", GoGetter: "TimeToLiveSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "writeProvisionedThroughputSettings", GoGetter: "WriteProvisionedThroughputSettings"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.AttributeDefinitionProperty",
		reflect.TypeOf((*CfnGlobalTable_AttributeDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.CapacityAutoScalingSettingsProperty",
		reflect.TypeOf((*CfnGlobalTable_CapacityAutoScalingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ContributorInsightsSpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_ContributorInsightsSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.GlobalSecondaryIndexProperty",
		reflect.TypeOf((*CfnGlobalTable_GlobalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.KeySchemaProperty",
		reflect.TypeOf((*CfnGlobalTable_KeySchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.LocalSecondaryIndexProperty",
		reflect.TypeOf((*CfnGlobalTable_LocalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.PointInTimeRecoverySpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_PointInTimeRecoverySpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ProjectionProperty",
		reflect.TypeOf((*CfnGlobalTable_ProjectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ReadProvisionedThroughputSettingsProperty",
		reflect.TypeOf((*CfnGlobalTable_ReadProvisionedThroughputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ReplicaGlobalSecondaryIndexSpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_ReplicaGlobalSecondaryIndexSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ReplicaSSESpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_ReplicaSSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.ReplicaSpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_ReplicaSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.SSESpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.StreamSpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_StreamSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.TargetTrackingScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnGlobalTable_TargetTrackingScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.TimeToLiveSpecificationProperty",
		reflect.TypeOf((*CfnGlobalTable_TimeToLiveSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable.WriteProvisionedThroughputSettingsProperty",
		reflect.TypeOf((*CfnGlobalTable_WriteProvisionedThroughputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTableProps",
		reflect.TypeOf((*CfnGlobalTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_dynamodb.CfnTable",
		reflect.TypeOf((*CfnTable)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attributeDefinitions", GoGetter: "AttributeDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "attrStreamArn", GoGetter: "AttrStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "billingMode", GoGetter: "BillingMode"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contributorInsightsSpecification", GoGetter: "ContributorInsightsSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "globalSecondaryIndexes", GoGetter: "GlobalSecondaryIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "importSourceSpecification", GoGetter: "ImportSourceSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keySchema", GoGetter: "KeySchema"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamSpecification", GoGetter: "KinesisStreamSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "localSecondaryIndexes", GoGetter: "LocalSecondaryIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeRecoverySpecification", GoGetter: "PointInTimeRecoverySpecification"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedThroughput", GoGetter: "ProvisionedThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamSpecification", GoGetter: "StreamSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "tableClass", GoGetter: "TableClass"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeToLiveSpecification", GoGetter: "TimeToLiveSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.AttributeDefinitionProperty",
		reflect.TypeOf((*CfnTable_AttributeDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.ContributorInsightsSpecificationProperty",
		reflect.TypeOf((*CfnTable_ContributorInsightsSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.CsvProperty",
		reflect.TypeOf((*CfnTable_CsvProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.GlobalSecondaryIndexProperty",
		reflect.TypeOf((*CfnTable_GlobalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.ImportSourceSpecificationProperty",
		reflect.TypeOf((*CfnTable_ImportSourceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.InputFormatOptionsProperty",
		reflect.TypeOf((*CfnTable_InputFormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.KeySchemaProperty",
		reflect.TypeOf((*CfnTable_KeySchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.KinesisStreamSpecificationProperty",
		reflect.TypeOf((*CfnTable_KinesisStreamSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.LocalSecondaryIndexProperty",
		reflect.TypeOf((*CfnTable_LocalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.PointInTimeRecoverySpecificationProperty",
		reflect.TypeOf((*CfnTable_PointInTimeRecoverySpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.ProjectionProperty",
		reflect.TypeOf((*CfnTable_ProjectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnTable_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.S3BucketSourceProperty",
		reflect.TypeOf((*CfnTable_S3BucketSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.SSESpecificationProperty",
		reflect.TypeOf((*CfnTable_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.StreamSpecificationProperty",
		reflect.TypeOf((*CfnTable_StreamSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTable.TimeToLiveSpecificationProperty",
		reflect.TypeOf((*CfnTable_TimeToLiveSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.CfnTableProps",
		reflect.TypeOf((*CfnTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.EnableScalingProps",
		reflect.TypeOf((*EnableScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.GlobalSecondaryIndexProps",
		reflect.TypeOf((*GlobalSecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_dynamodb.IScalableTableAttribute",
		reflect.TypeOf((*IScalableTableAttribute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "scaleOnSchedule", GoMethod: "ScaleOnSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnUtilization", GoMethod: "ScaleOnUtilization"},
		},
		func() interface{} {
			return &jsiiProxy_IScalableTableAttribute{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_dynamodb.ITable",
		reflect.TypeOf((*ITable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadData", GoMethod: "GrantReadData"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWriteData", GoMethod: "GrantReadWriteData"},
			_jsii_.MemberMethod{JsiiMethod: "grantStream", GoMethod: "GrantStream"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamRead", GoMethod: "GrantStreamRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantTableListStreams", GoMethod: "GrantTableListStreams"},
			_jsii_.MemberMethod{JsiiMethod: "grantWriteData", GoMethod: "GrantWriteData"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricConditionalCheckFailedRequests", GoMethod: "MetricConditionalCheckFailedRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedReadCapacityUnits", GoMethod: "MetricConsumedReadCapacityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedWriteCapacityUnits", GoMethod: "MetricConsumedWriteCapacityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccessfulRequestLatency", GoMethod: "MetricSuccessfulRequestLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForOperations", GoMethod: "MetricSystemErrorsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequests", GoMethod: "MetricThrottledRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequestsForOperations", GoMethod: "MetricThrottledRequestsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableStreamArn", GoGetter: "TableStreamArn"},
		},
		func() interface{} {
			j := jsiiProxy_ITable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.LocalSecondaryIndexProps",
		reflect.TypeOf((*LocalSecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.Operation",
		reflect.TypeOf((*Operation)(nil)).Elem(),
		map[string]interface{}{
			"GET_ITEM": Operation_GET_ITEM,
			"BATCH_GET_ITEM": Operation_BATCH_GET_ITEM,
			"SCAN": Operation_SCAN,
			"QUERY": Operation_QUERY,
			"GET_RECORDS": Operation_GET_RECORDS,
			"PUT_ITEM": Operation_PUT_ITEM,
			"DELETE_ITEM": Operation_DELETE_ITEM,
			"UPDATE_ITEM": Operation_UPDATE_ITEM,
			"BATCH_WRITE_ITEM": Operation_BATCH_WRITE_ITEM,
			"TRANSACT_WRITE_ITEMS": Operation_TRANSACT_WRITE_ITEMS,
			"TRANSACT_GET_ITEMS": Operation_TRANSACT_GET_ITEMS,
			"EXECUTE_TRANSACTION": Operation_EXECUTE_TRANSACTION,
			"BATCH_EXECUTE_STATEMENT": Operation_BATCH_EXECUTE_STATEMENT,
			"EXECUTE_STATEMENT": Operation_EXECUTE_STATEMENT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.OperationsMetricOptions",
		reflect.TypeOf((*OperationsMetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.ProjectionType",
		reflect.TypeOf((*ProjectionType)(nil)).Elem(),
		map[string]interface{}{
			"KEYS_ONLY": ProjectionType_KEYS_ONLY,
			"INCLUDE": ProjectionType_INCLUDE,
			"ALL": ProjectionType_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.SchemaOptions",
		reflect.TypeOf((*SchemaOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.SecondaryIndexProps",
		reflect.TypeOf((*SecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.StreamViewType",
		reflect.TypeOf((*StreamViewType)(nil)).Elem(),
		map[string]interface{}{
			"NEW_IMAGE": StreamViewType_NEW_IMAGE,
			"OLD_IMAGE": StreamViewType_OLD_IMAGE,
			"NEW_AND_OLD_IMAGES": StreamViewType_NEW_AND_OLD_IMAGES,
			"KEYS_ONLY": StreamViewType_KEYS_ONLY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.SystemErrorsForOperationsMetricOptions",
		reflect.TypeOf((*SystemErrorsForOperationsMetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_dynamodb.Table",
		reflect.TypeOf((*Table)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGlobalSecondaryIndex", GoMethod: "AddGlobalSecondaryIndex"},
			_jsii_.MemberMethod{JsiiMethod: "addLocalSecondaryIndex", GoMethod: "AddLocalSecondaryIndex"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "autoScaleGlobalSecondaryIndexReadCapacity", GoMethod: "AutoScaleGlobalSecondaryIndexReadCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "autoScaleGlobalSecondaryIndexWriteCapacity", GoMethod: "AutoScaleGlobalSecondaryIndexWriteCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "autoScaleReadCapacity", GoMethod: "AutoScaleReadCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "autoScaleWriteCapacity", GoMethod: "AutoScaleWriteCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadData", GoMethod: "GrantReadData"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWriteData", GoMethod: "GrantReadWriteData"},
			_jsii_.MemberMethod{JsiiMethod: "grantStream", GoMethod: "GrantStream"},
			_jsii_.MemberMethod{JsiiMethod: "grantStreamRead", GoMethod: "GrantStreamRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantTableListStreams", GoMethod: "GrantTableListStreams"},
			_jsii_.MemberMethod{JsiiMethod: "grantWriteData", GoMethod: "GrantWriteData"},
			_jsii_.MemberProperty{JsiiProperty: "hasIndex", GoGetter: "HasIndex"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricConditionalCheckFailedRequests", GoMethod: "MetricConditionalCheckFailedRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedReadCapacityUnits", GoMethod: "MetricConsumedReadCapacityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedWriteCapacityUnits", GoMethod: "MetricConsumedWriteCapacityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccessfulRequestLatency", GoMethod: "MetricSuccessfulRequestLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForOperations", GoMethod: "MetricSystemErrorsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequests", GoMethod: "MetricThrottledRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequestsForOperation", GoMethod: "MetricThrottledRequestsForOperation"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequestsForOperations", GoMethod: "MetricThrottledRequestsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "regionalArns", GoGetter: "RegionalArns"},
			_jsii_.MemberMethod{JsiiMethod: "schema", GoMethod: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableStreamArn", GoGetter: "TableStreamArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Table{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.TableAttributes",
		reflect.TypeOf((*TableAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.TableClass",
		reflect.TypeOf((*TableClass)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": TableClass_STANDARD,
			"STANDARD_INFREQUENT_ACCESS": TableClass_STANDARD_INFREQUENT_ACCESS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_dynamodb.TableEncryption",
		reflect.TypeOf((*TableEncryption)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": TableEncryption_DEFAULT,
			"CUSTOMER_MANAGED": TableEncryption_CUSTOMER_MANAGED,
			"AWS_MANAGED": TableEncryption_AWS_MANAGED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.TableOptions",
		reflect.TypeOf((*TableOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.TableProps",
		reflect.TypeOf((*TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_dynamodb.UtilizationScalingProps",
		reflect.TypeOf((*UtilizationScalingProps)(nil)).Elem(),
	)
}

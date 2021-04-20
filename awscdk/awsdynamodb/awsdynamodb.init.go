package awsdynamodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.Attribute",
		reflect.TypeOf((*Attribute)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.AttributeType",
		reflect.TypeOf((*AttributeType)(nil)).Elem(),
		map[string]interface{}{
			"BINARY": AttributeType_BINARY,
			"NUMBER": AttributeType_NUMBER,
			"STRING": AttributeType_STRING,
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.BillingMode",
		reflect.TypeOf((*BillingMode)(nil)).Elem(),
		map[string]interface{}{
			"PAY_PER_REQUEST": BillingMode_PAY_PER_REQUEST,
			"PROVISIONED": BillingMode_PROVISIONED,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_dynamodb.CfnTable",
		reflect.TypeOf((*CfnTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
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
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keySchema", GoGetter: "KeySchema"},
			_jsii_.MemberProperty{JsiiProperty: "localSecondaryIndexes", GoGetter: "LocalSecondaryIndexes"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeRecoverySpecification", GoGetter: "PointInTimeRecoverySpecification"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedThroughput", GoGetter: "ProvisionedThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "streamSpecification", GoGetter: "StreamSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeToLiveSpecification", GoGetter: "TimeToLiveSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
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
		"monocdk.aws_dynamodb.CfnTable.AttributeDefinitionProperty",
		reflect.TypeOf((*CfnTable_AttributeDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.ContributorInsightsSpecificationProperty",
		reflect.TypeOf((*CfnTable_ContributorInsightsSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.GlobalSecondaryIndexProperty",
		reflect.TypeOf((*CfnTable_GlobalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.KeySchemaProperty",
		reflect.TypeOf((*CfnTable_KeySchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.LocalSecondaryIndexProperty",
		reflect.TypeOf((*CfnTable_LocalSecondaryIndexProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.PointInTimeRecoverySpecificationProperty",
		reflect.TypeOf((*CfnTable_PointInTimeRecoverySpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.ProjectionProperty",
		reflect.TypeOf((*CfnTable_ProjectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnTable_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.SSESpecificationProperty",
		reflect.TypeOf((*CfnTable_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.StreamSpecificationProperty",
		reflect.TypeOf((*CfnTable_StreamSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTable.TimeToLiveSpecificationProperty",
		reflect.TypeOf((*CfnTable_TimeToLiveSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.CfnTableProps",
		reflect.TypeOf((*CfnTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.EnableScalingProps",
		reflect.TypeOf((*EnableScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.GlobalSecondaryIndexProps",
		reflect.TypeOf((*GlobalSecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_dynamodb.IScalableTableAttribute",
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
		"monocdk.aws_dynamodb.ITable",
		reflect.TypeOf((*ITable)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForOperations", GoMethod: "MetricSystemErrorsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequests", GoMethod: "MetricThrottledRequests"},
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
		"monocdk.aws_dynamodb.LocalSecondaryIndexProps",
		reflect.TypeOf((*LocalSecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.Operation",
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
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.ProjectionType",
		reflect.TypeOf((*ProjectionType)(nil)).Elem(),
		map[string]interface{}{
			"KEYS_ONLY": ProjectionType_KEYS_ONLY,
			"INCLUDE": ProjectionType_INCLUDE,
			"ALL": ProjectionType_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.SecondaryIndexProps",
		reflect.TypeOf((*SecondaryIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.StreamViewType",
		reflect.TypeOf((*StreamViewType)(nil)).Elem(),
		map[string]interface{}{
			"NEW_IMAGE": StreamViewType_NEW_IMAGE,
			"OLD_IMAGE": StreamViewType_OLD_IMAGE,
			"NEW_AND_OLD_IMAGES": StreamViewType_NEW_AND_OLD_IMAGES,
			"KEYS_ONLY": StreamViewType_KEYS_ONLY,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.SystemErrorsForOperationsMetricOptions",
		reflect.TypeOf((*SystemErrorsForOperationsMetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_dynamodb.Table",
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
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrors", GoMethod: "MetricSystemErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemErrorsForOperations", GoMethod: "MetricSystemErrorsForOperations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottledRequests", GoMethod: "MetricThrottledRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricUserErrors", GoMethod: "MetricUserErrors"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "regionalArns", GoGetter: "RegionalArns"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tableArn", GoGetter: "TableArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableStreamArn", GoGetter: "TableStreamArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Table{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.TableAttributes",
		reflect.TypeOf((*TableAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_dynamodb.TableEncryption",
		reflect.TypeOf((*TableEncryption)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": TableEncryption_DEFAULT,
			"CUSTOMER_MANAGED": TableEncryption_CUSTOMER_MANAGED,
			"AWS_MANAGED": TableEncryption_AWS_MANAGED,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.TableOptions",
		reflect.TypeOf((*TableOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.TableProps",
		reflect.TypeOf((*TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_dynamodb.UtilizationScalingProps",
		reflect.TypeOf((*UtilizationScalingProps)(nil)).Elem(),
	)
}

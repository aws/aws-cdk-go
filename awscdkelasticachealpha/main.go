// The CDK Construct Library for AWS::ElastiCache
package awscdkelasticachealpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.AccessControl",
		reflect.TypeOf((*AccessControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessString", GoGetter: "AccessString"},
		},
		func() interface{} {
			return &jsiiProxy_AccessControl{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.BackupSettings",
		reflect.TypeOf((*BackupSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		reflect.TypeOf((*CacheEngine)(nil)).Elem(),
		map[string]interface{}{
			"VALKEY_LATEST": CacheEngine_VALKEY_LATEST,
			"VALKEY_7": CacheEngine_VALKEY_7,
			"VALKEY_8": CacheEngine_VALKEY_8,
			"REDIS_LATEST": CacheEngine_REDIS_LATEST,
			"REDIS_7": CacheEngine_REDIS_7,
			"MEMCACHED_LATEST": CacheEngine_MEMCACHED_LATEST,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.CacheUsageLimitsProperty",
		reflect.TypeOf((*CacheUsageLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-elasticache-alpha.IServerlessCache",
		reflect.TypeOf((*IServerlessCache)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupArnsToRestore", GoGetter: "BackupArnsToRestore"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveConnections", GoMethod: "MetricActiveConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitCount", GoMethod: "MetricCacheHitCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitRate", GoMethod: "MetricCacheHitRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheMissCount", GoMethod: "MetricCacheMissCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataStored", GoMethod: "MetricDataStored"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesIn", GoMethod: "MetricNetworkBytesIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesOut", GoMethod: "MetricNetworkBytesOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricProcessingUnitsConsumed", GoMethod: "MetricProcessingUnitsConsumed"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadRequestLatency", GoMethod: "MetricReadRequestLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteRequestLatency", GoMethod: "MetricWriteRequestLatency"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheArn", GoGetter: "ServerlessCacheArn"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheName", GoGetter: "ServerlessCacheName"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheRef", GoGetter: "ServerlessCacheRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberProperty{JsiiProperty: "userGroup", GoGetter: "UserGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_IServerlessCache{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawselasticacheIServerlessCacheRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-elasticache-alpha.IUser",
		reflect.TypeOf((*IUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "userArn", GoGetter: "UserArn"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
		},
		func() interface{} {
			j := jsiiProxy_IUser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-elasticache-alpha.IUserGroup",
		reflect.TypeOf((*IUserGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUser", GoMethod: "AddUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupArn", GoGetter: "UserGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupName", GoGetter: "UserGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
		},
		func() interface{} {
			j := jsiiProxy_IUserGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.IamUser",
		reflect.TypeOf((*IamUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessString", GoGetter: "AccessString"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userArn", GoGetter: "UserArn"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userStatus", GoGetter: "UserStatus"},
		},
		func() interface{} {
			j := jsiiProxy_IamUser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.IamUserProps",
		reflect.TypeOf((*IamUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.NoPasswordUser",
		reflect.TypeOf((*NoPasswordUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessString", GoGetter: "AccessString"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userArn", GoGetter: "UserArn"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userStatus", GoGetter: "UserStatus"},
		},
		func() interface{} {
			j := jsiiProxy_NoPasswordUser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.NoPasswordUserProps",
		reflect.TypeOf((*NoPasswordUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.PasswordUser",
		reflect.TypeOf((*PasswordUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessString", GoGetter: "AccessString"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userArn", GoGetter: "UserArn"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userStatus", GoGetter: "UserStatus"},
		},
		func() interface{} {
			j := jsiiProxy_PasswordUser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.PasswordUserProps",
		reflect.TypeOf((*PasswordUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCache",
		reflect.TypeOf((*ServerlessCache)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupArnsToRestore", GoGetter: "BackupArnsToRestore"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveConnections", GoMethod: "MetricActiveConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitCount", GoMethod: "MetricCacheHitCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitRate", GoMethod: "MetricCacheHitRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheMissCount", GoMethod: "MetricCacheMissCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataStored", GoMethod: "MetricDataStored"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesIn", GoMethod: "MetricNetworkBytesIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesOut", GoMethod: "MetricNetworkBytesOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricProcessingUnitsConsumed", GoMethod: "MetricProcessingUnitsConsumed"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadRequestLatency", GoMethod: "MetricReadRequestLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteRequestLatency", GoMethod: "MetricWriteRequestLatency"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheArn", GoGetter: "ServerlessCacheArn"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheEndpointAddress", GoGetter: "ServerlessCacheEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheEndpointPort", GoGetter: "ServerlessCacheEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheName", GoGetter: "ServerlessCacheName"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheReaderEndpointAddress", GoGetter: "ServerlessCacheReaderEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheReaderEndpointPort", GoGetter: "ServerlessCacheReaderEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheRef", GoGetter: "ServerlessCacheRef"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheStatus", GoGetter: "ServerlessCacheStatus"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroup", GoGetter: "UserGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ServerlessCache{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServerlessCacheBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCacheAttributes",
		reflect.TypeOf((*ServerlessCacheAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCacheBase",
		reflect.TypeOf((*ServerlessCacheBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupArnsToRestore", GoGetter: "BackupArnsToRestore"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveConnections", GoMethod: "MetricActiveConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitCount", GoMethod: "MetricCacheHitCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheHitRate", GoMethod: "MetricCacheHitRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricCacheMissCount", GoMethod: "MetricCacheMissCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDataStored", GoMethod: "MetricDataStored"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesIn", GoMethod: "MetricNetworkBytesIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkBytesOut", GoMethod: "MetricNetworkBytesOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricProcessingUnitsConsumed", GoMethod: "MetricProcessingUnitsConsumed"},
			_jsii_.MemberMethod{JsiiMethod: "metricReadRequestLatency", GoMethod: "MetricReadRequestLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricWriteRequestLatency", GoMethod: "MetricWriteRequestLatency"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheArn", GoGetter: "ServerlessCacheArn"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheName", GoGetter: "ServerlessCacheName"},
			_jsii_.MemberProperty{JsiiProperty: "serverlessCacheRef", GoGetter: "ServerlessCacheRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroup", GoGetter: "UserGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ServerlessCacheBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerlessCache)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCacheGrants",
		reflect.TypeOf((*ServerlessCacheGrants)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "connect", GoMethod: "Connect"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
		},
		func() interface{} {
			return &jsiiProxy_ServerlessCacheGrants{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCacheProps",
		reflect.TypeOf((*ServerlessCacheProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.UserBase",
		reflect.TypeOf((*UserBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userArn", GoGetter: "UserArn"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
		},
		func() interface{} {
			j := jsiiProxy_UserBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.UserBaseAttributes",
		reflect.TypeOf((*UserBaseAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.UserBaseProps",
		reflect.TypeOf((*UserBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-elasticache-alpha.UserEngine",
		reflect.TypeOf((*UserEngine)(nil)).Elem(),
		map[string]interface{}{
			"VALKEY": UserEngine_VALKEY,
			"REDIS": UserEngine_REDIS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.UserGroup",
		reflect.TypeOf((*UserGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUser", GoMethod: "AddUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupArn", GoGetter: "UserGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupName", GoGetter: "UserGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupStatus", GoGetter: "UserGroupStatus"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
		},
		func() interface{} {
			j := jsiiProxy_UserGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserGroupBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.UserGroupAttributes",
		reflect.TypeOf((*UserGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-elasticache-alpha.UserGroupBase",
		reflect.TypeOf((*UserGroupBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUser", GoMethod: "AddUser"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupArn", GoGetter: "UserGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupName", GoGetter: "UserGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
		},
		func() interface{} {
			j := jsiiProxy_UserGroupBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUserGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-elasticache-alpha.UserGroupProps",
		reflect.TypeOf((*UserGroupProps)(nil)).Elem(),
	)
}

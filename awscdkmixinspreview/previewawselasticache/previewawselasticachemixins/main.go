package previewawselasticachemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterElasticacheLogs",
		reflect.TypeOf((*CfnCacheClusterElasticacheLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCacheClusterElasticacheLogs{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterLogsMixin",
		reflect.TypeOf((*CfnCacheClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCacheClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterMixinProps",
		reflect.TypeOf((*CfnCacheClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin",
		reflect.TypeOf((*CfnCacheClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCacheClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin.CloudWatchLogsDestinationDetailsProperty",
		reflect.TypeOf((*CfnCacheClusterPropsMixin_CloudWatchLogsDestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin.DestinationDetailsProperty",
		reflect.TypeOf((*CfnCacheClusterPropsMixin_DestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin.KinesisFirehoseDestinationDetailsProperty",
		reflect.TypeOf((*CfnCacheClusterPropsMixin_KinesisFirehoseDestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin.LogDeliveryConfigurationRequestProperty",
		reflect.TypeOf((*CfnCacheClusterPropsMixin_LogDeliveryConfigurationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupMixinProps",
		reflect.TypeOf((*CfnGlobalReplicationGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin",
		reflect.TypeOf((*CfnGlobalReplicationGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalReplicationGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin.GlobalReplicationGroupMemberProperty",
		reflect.TypeOf((*CfnGlobalReplicationGroupPropsMixin_GlobalReplicationGroupMemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin.RegionalConfigurationProperty",
		reflect.TypeOf((*CfnGlobalReplicationGroupPropsMixin_RegionalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin.ReshardingConfigurationProperty",
		reflect.TypeOf((*CfnGlobalReplicationGroupPropsMixin_ReshardingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnParameterGroupMixinProps",
		reflect.TypeOf((*CfnParameterGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnParameterGroupPropsMixin",
		reflect.TypeOf((*CfnParameterGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnParameterGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupMixinProps",
		reflect.TypeOf((*CfnReplicationGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin.CloudWatchLogsDestinationDetailsProperty",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin_CloudWatchLogsDestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin.DestinationDetailsProperty",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin_DestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin.KinesisFirehoseDestinationDetailsProperty",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin_KinesisFirehoseDestinationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin.LogDeliveryConfigurationRequestProperty",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin_LogDeliveryConfigurationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin.NodeGroupConfigurationProperty",
		reflect.TypeOf((*CfnReplicationGroupPropsMixin_NodeGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSecurityGroupIngressMixinProps",
		reflect.TypeOf((*CfnSecurityGroupIngressMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSecurityGroupIngressPropsMixin",
		reflect.TypeOf((*CfnSecurityGroupIngressPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityGroupIngressPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSecurityGroupMixinProps",
		reflect.TypeOf((*CfnSecurityGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSecurityGroupPropsMixin",
		reflect.TypeOf((*CfnSecurityGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCacheMixinProps",
		reflect.TypeOf((*CfnServerlessCacheMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCachePropsMixin",
		reflect.TypeOf((*CfnServerlessCachePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessCachePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCachePropsMixin.CacheUsageLimitsProperty",
		reflect.TypeOf((*CfnServerlessCachePropsMixin_CacheUsageLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCachePropsMixin.DataStorageProperty",
		reflect.TypeOf((*CfnServerlessCachePropsMixin_DataStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCachePropsMixin.ECPUPerSecondProperty",
		reflect.TypeOf((*CfnServerlessCachePropsMixin_ECPUPerSecondProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnServerlessCachePropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnServerlessCachePropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSubnetGroupMixinProps",
		reflect.TypeOf((*CfnSubnetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnSubnetGroupPropsMixin",
		reflect.TypeOf((*CfnSubnetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubnetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnUserGroupMixinProps",
		reflect.TypeOf((*CfnUserGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnUserGroupPropsMixin",
		reflect.TypeOf((*CfnUserGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnUserMixinProps",
		reflect.TypeOf((*CfnUserMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnUserPropsMixin",
		reflect.TypeOf((*CfnUserPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnUserPropsMixin.AuthenticationModeProperty",
		reflect.TypeOf((*CfnUserPropsMixin_AuthenticationModeProperty)(nil)).Elem(),
	)
}

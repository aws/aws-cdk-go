package previewawsroute53profilesmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileAssociationMixinProps",
		reflect.TypeOf((*CfnProfileAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileAssociationPropsMixin",
		reflect.TypeOf((*CfnProfileAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfileAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileLogsMixin",
		reflect.TypeOf((*CfnProfileLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfileLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileMixinProps",
		reflect.TypeOf((*CfnProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfilePropsMixin",
		reflect.TypeOf((*CfnProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileResourceAssociationMixinProps",
		reflect.TypeOf((*CfnProfileResourceAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileResourceAssociationPropsMixin",
		reflect.TypeOf((*CfnProfileResourceAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfileResourceAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogs",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.S3",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3_JSON,
			"PLAIN": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3_PLAIN,
			"W3C": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3_W3C,
			"PARQUET": CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsS3Props",
		reflect.TypeOf((*CfnProfileRoute53ProfilesResolverQueryLogsS3Props)(nil)).Elem(),
	)
}

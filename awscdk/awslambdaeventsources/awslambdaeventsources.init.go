package awslambdaeventsources

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		reflect.TypeOf((*ApiEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ApiEventSource{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_lambda_event_sources.AuthenticationMethod",
		reflect.TypeOf((*AuthenticationMethod)(nil)).Elem(),
		map[string]interface{}{
			"SASL_SCRAM_512_AUTH": AuthenticationMethod_SASL_SCRAM_512_AUTH,
			"SASL_SCRAM_256_AUTH": AuthenticationMethod_SASL_SCRAM_256_AUTH,
			"BASIC_AUTH": AuthenticationMethod_BASIC_AUTH,
			"CLIENT_CERTIFICATE_TLS_AUTH": AuthenticationMethod_CLIENT_CERTIFICATE_TLS_AUTH,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.BaseStreamEventSourceProps",
		reflect.TypeOf((*BaseStreamEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		reflect.TypeOf((*DynamoEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "enrichMappingOptions", GoMethod: "EnrichMappingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "eventSourceMappingId", GoGetter: "EventSourceMappingId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSourceProps",
		reflect.TypeOf((*DynamoEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.KafkaEventSourceProps",
		reflect.TypeOf((*KafkaEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisEventSource",
		reflect.TypeOf((*KinesisEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "enrichMappingOptions", GoMethod: "EnrichMappingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "eventSourceMappingId", GoGetter: "EventSourceMappingId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "stream", GoGetter: "Stream"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisEventSourceProps",
		reflect.TypeOf((*KinesisEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		reflect.TypeOf((*ManagedKafkaEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "enrichMappingOptions", GoMethod: "EnrichMappingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "eventSourceMappingId", GoGetter: "EventSourceMappingId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKafkaEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSourceProps",
		reflect.TypeOf((*ManagedKafkaEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		reflect.TypeOf((*S3EventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
		},
		func() interface{} {
			j := jsiiProxy_S3EventSource{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSourceProps",
		reflect.TypeOf((*S3EventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		reflect.TypeOf((*SelfManagedKafkaEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "enrichMappingOptions", GoMethod: "EnrichMappingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_SelfManagedKafkaEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSourceProps",
		reflect.TypeOf((*SelfManagedKafkaEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		reflect.TypeOf((*SnsDlq)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SnsDlq{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSourceDlq)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		reflect.TypeOf((*SnsEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
		},
		func() interface{} {
			j := jsiiProxy_SnsEventSource{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSourceProps",
		reflect.TypeOf((*SnsEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		reflect.TypeOf((*SqsDlq)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsDlq{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSourceDlq)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		reflect.TypeOf((*SqsEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "eventSourceMappingId", GoGetter: "EventSourceMappingId"},
			_jsii_.MemberProperty{JsiiProperty: "queue", GoGetter: "Queue"},
		},
		func() interface{} {
			j := jsiiProxy_SqsEventSource{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSourceProps",
		reflect.TypeOf((*SqsEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_event_sources.StreamEventSource",
		reflect.TypeOf((*StreamEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "enrichMappingOptions", GoMethod: "EnrichMappingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_StreamEventSource{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_event_sources.StreamEventSourceProps",
		reflect.TypeOf((*StreamEventSourceProps)(nil)).Elem(),
	)
}

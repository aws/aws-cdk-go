package awscdkiotalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.ActionConfig",
		reflect.TypeOf((*ActionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.IAction",
		reflect.TypeOf((*IAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAction{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.ITopicRule",
		reflect.TypeOf((*ITopicRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleArn", GoGetter: "TopicRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleName", GoGetter: "TopicRuleName"},
		},
		func() interface{} {
			j := jsiiProxy_ITopicRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.IotSql",
		reflect.TypeOf((*IotSql)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IotSql{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.IotSqlConfig",
		reflect.TypeOf((*IotSqlConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		reflect.TypeOf((*TopicRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleArn", GoGetter: "TopicRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleName", GoGetter: "TopicRuleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TopicRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITopicRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.TopicRuleProps",
		reflect.TypeOf((*TopicRuleProps)(nil)).Elem(),
	)
}

package previewawsconnectevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents",
		reflect.TypeOf((*InstanceEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "codeConnectContactPattern", GoMethod: "CodeConnectContactPattern"},
			_jsii_.MemberMethod{JsiiMethod: "contactLensPostCallRulesMatchedPattern", GoMethod: "ContactLensPostCallRulesMatchedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "contactLensRealtimeRulesMatchedPattern", GoMethod: "ContactLensRealtimeRulesMatchedPattern"},
		},
		func() interface{} {
			return &jsiiProxy_InstanceEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact",
		reflect.TypeOf((*InstanceEvents_CodeConnectContact)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_CodeConnectContact{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact.AgentInfo",
		reflect.TypeOf((*InstanceEvents_CodeConnectContact_AgentInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact.CodeConnectContactProps",
		reflect.TypeOf((*InstanceEvents_CodeConnectContact_CodeConnectContactProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact.QueueInfo",
		reflect.TypeOf((*InstanceEvents_CodeConnectContact_QueueInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensPostCallRulesMatched",
		reflect.TypeOf((*InstanceEvents_ContactLensPostCallRulesMatched)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_ContactLensPostCallRulesMatched{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensPostCallRulesMatched.ContactLensPostCallRulesMatchedProps",
		reflect.TypeOf((*InstanceEvents_ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensRealtimeRulesMatched",
		reflect.TypeOf((*InstanceEvents_ContactLensRealtimeRulesMatched)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_ContactLensRealtimeRulesMatched{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensRealtimeRulesMatched.ContactLensRealtimeRulesMatchedProps",
		reflect.TypeOf((*InstanceEvents_ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps)(nil)).Elem(),
	)
}

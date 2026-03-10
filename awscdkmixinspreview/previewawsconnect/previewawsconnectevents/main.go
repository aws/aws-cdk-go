package previewawsconnectevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact",
		reflect.TypeOf((*CodeConnectContact)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodeConnectContact{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact.AgentInfo",
		reflect.TypeOf((*CodeConnectContact_AgentInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact.CodeConnectContactProps",
		reflect.TypeOf((*CodeConnectContact_CodeConnectContactProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact.QueueInfo",
		reflect.TypeOf((*CodeConnectContact_QueueInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensPostCallRulesMatched",
		reflect.TypeOf((*ContactLensPostCallRulesMatched)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ContactLensPostCallRulesMatched{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensPostCallRulesMatched.ContactLensPostCallRulesMatchedProps",
		reflect.TypeOf((*ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensRealtimeRulesMatched",
		reflect.TypeOf((*ContactLensRealtimeRulesMatched)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ContactLensRealtimeRulesMatched{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensRealtimeRulesMatched.ContactLensRealtimeRulesMatchedProps",
		reflect.TypeOf((*ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps)(nil)).Elem(),
	)
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
}

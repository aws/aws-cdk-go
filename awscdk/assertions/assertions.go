package assertions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.Annotations",
		reflect.TypeOf((*Annotations)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "findError", GoMethod: "FindError"},
			_jsii_.MemberMethod{JsiiMethod: "findInfo", GoMethod: "FindInfo"},
			_jsii_.MemberMethod{JsiiMethod: "findWarning", GoMethod: "FindWarning"},
			_jsii_.MemberMethod{JsiiMethod: "hasError", GoMethod: "HasError"},
			_jsii_.MemberMethod{JsiiMethod: "hasInfo", GoMethod: "HasInfo"},
			_jsii_.MemberMethod{JsiiMethod: "hasNoError", GoMethod: "HasNoError"},
			_jsii_.MemberMethod{JsiiMethod: "hasNoInfo", GoMethod: "HasNoInfo"},
			_jsii_.MemberMethod{JsiiMethod: "hasNoWarning", GoMethod: "HasNoWarning"},
			_jsii_.MemberMethod{JsiiMethod: "hasWarning", GoMethod: "HasWarning"},
		},
		func() interface{} {
			return &jsiiProxy_Annotations{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.Capture",
		reflect.TypeOf((*Capture)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asArray", GoMethod: "AsArray"},
			_jsii_.MemberMethod{JsiiMethod: "asBoolean", GoMethod: "AsBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "asNumber", GoMethod: "AsNumber"},
			_jsii_.MemberMethod{JsiiMethod: "asObject", GoMethod: "AsObject"},
			_jsii_.MemberMethod{JsiiMethod: "asString", GoMethod: "AsString"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberMethod{JsiiMethod: "test", GoMethod: "Test"},
		},
		func() interface{} {
			j := jsiiProxy_Capture{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Matcher)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.Match",
		reflect.TypeOf((*Match)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Match{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assertions.MatchCapture",
		reflect.TypeOf((*MatchCapture)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assertions.MatchFailure",
		reflect.TypeOf((*MatchFailure)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.MatchResult",
		reflect.TypeOf((*MatchResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "compose", GoMethod: "Compose"},
			_jsii_.MemberProperty{JsiiProperty: "failCount", GoGetter: "FailCount"},
			_jsii_.MemberMethod{JsiiMethod: "finished", GoMethod: "Finished"},
			_jsii_.MemberMethod{JsiiMethod: "hasFailed", GoMethod: "HasFailed"},
			_jsii_.MemberMethod{JsiiMethod: "push", GoMethod: "Push"},
			_jsii_.MemberMethod{JsiiMethod: "recordCapture", GoMethod: "RecordCapture"},
			_jsii_.MemberMethod{JsiiMethod: "recordFailure", GoMethod: "RecordFailure"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toHumanStrings", GoMethod: "ToHumanStrings"},
		},
		func() interface{} {
			return &jsiiProxy_MatchResult{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.Matcher",
		reflect.TypeOf((*Matcher)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "test", GoMethod: "Test"},
		},
		func() interface{} {
			return &jsiiProxy_Matcher{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assertions.Template",
		reflect.TypeOf((*Template)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allResources", GoMethod: "AllResources"},
			_jsii_.MemberMethod{JsiiMethod: "allResourcesProperties", GoMethod: "AllResourcesProperties"},
			_jsii_.MemberMethod{JsiiMethod: "findConditions", GoMethod: "FindConditions"},
			_jsii_.MemberMethod{JsiiMethod: "findMappings", GoMethod: "FindMappings"},
			_jsii_.MemberMethod{JsiiMethod: "findOutputs", GoMethod: "FindOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "findParameters", GoMethod: "FindParameters"},
			_jsii_.MemberMethod{JsiiMethod: "findResources", GoMethod: "FindResources"},
			_jsii_.MemberMethod{JsiiMethod: "hasCondition", GoMethod: "HasCondition"},
			_jsii_.MemberMethod{JsiiMethod: "hasMapping", GoMethod: "HasMapping"},
			_jsii_.MemberMethod{JsiiMethod: "hasOutput", GoMethod: "HasOutput"},
			_jsii_.MemberMethod{JsiiMethod: "hasParameter", GoMethod: "HasParameter"},
			_jsii_.MemberMethod{JsiiMethod: "hasResource", GoMethod: "HasResource"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceProperties", GoMethod: "HasResourceProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resourceCountIs", GoMethod: "ResourceCountIs"},
			_jsii_.MemberMethod{JsiiMethod: "resourcePropertiesCountIs", GoMethod: "ResourcePropertiesCountIs"},
			_jsii_.MemberMethod{JsiiMethod: "templateMatches", GoMethod: "TemplateMatches"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
		},
		func() interface{} {
			return &jsiiProxy_Template{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assertions.TemplateParsingOptions",
		reflect.TypeOf((*TemplateParsingOptions)(nil)).Elem(),
	)
}

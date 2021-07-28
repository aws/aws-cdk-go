package assertions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.assertions.Match",
		reflect.TypeOf((*Match)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Match{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.assertions.MatchResult",
		reflect.TypeOf((*MatchResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "compose", GoMethod: "Compose"},
			_jsii_.MemberProperty{JsiiProperty: "failCount", GoGetter: "FailCount"},
			_jsii_.MemberMethod{JsiiMethod: "hasFailed", GoMethod: "HasFailed"},
			_jsii_.MemberMethod{JsiiMethod: "push", GoMethod: "Push"},
			_jsii_.MemberMethod{JsiiMethod: "toHumanStrings", GoMethod: "ToHumanStrings"},
		},
		func() interface{} {
			return &jsiiProxy_MatchResult{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.assertions.Matcher",
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
		"monocdk.assertions.TemplateAssertions",
		reflect.TypeOf((*TemplateAssertions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "findResources", GoMethod: "FindResources"},
			_jsii_.MemberMethod{JsiiMethod: "hasResource", GoMethod: "HasResource"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceProperties", GoMethod: "HasResourceProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resourceCountIs", GoMethod: "ResourceCountIs"},
			_jsii_.MemberMethod{JsiiMethod: "templateMatches", GoMethod: "TemplateMatches"},
		},
		func() interface{} {
			return &jsiiProxy_TemplateAssertions{}
		},
	)
}

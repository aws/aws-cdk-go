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
		"monocdk.assertions.TemplateAssertions",
		reflect.TypeOf((*TemplateAssertions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "hasResourceDefinition", GoMethod: "HasResourceDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceProperties", GoMethod: "HasResourceProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resourceCountIs", GoMethod: "ResourceCountIs"},
			_jsii_.MemberMethod{JsiiMethod: "templateMatches", GoMethod: "TemplateMatches"},
		},
		func() interface{} {
			return &jsiiProxy_TemplateAssertions{}
		},
	)
}

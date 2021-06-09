package cloudformationinclude

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.cloudformation_include.CfnInclude",
		reflect.TypeOf((*CfnInclude)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getCondition", GoMethod: "GetCondition"},
			_jsii_.MemberMethod{JsiiMethod: "getHook", GoMethod: "GetHook"},
			_jsii_.MemberMethod{JsiiMethod: "getMapping", GoMethod: "GetMapping"},
			_jsii_.MemberMethod{JsiiMethod: "getNestedStack", GoMethod: "GetNestedStack"},
			_jsii_.MemberMethod{JsiiMethod: "getOutput", GoMethod: "GetOutput"},
			_jsii_.MemberMethod{JsiiMethod: "getParameter", GoMethod: "GetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "getResource", GoMethod: "GetResource"},
			_jsii_.MemberMethod{JsiiMethod: "getRule", GoMethod: "GetRule"},
			_jsii_.MemberMethod{JsiiMethod: "loadNestedStack", GoMethod: "LoadNestedStack"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInclude{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.cloudformation_include.CfnIncludeProps",
		reflect.TypeOf((*CfnIncludeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.cloudformation_include.IncludedNestedStack",
		reflect.TypeOf((*IncludedNestedStack)(nil)).Elem(),
	)
}

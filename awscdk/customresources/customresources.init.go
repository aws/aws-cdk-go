package customresources

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.custom_resources.AwsCustomResource",
		reflect.TypeOf((*AwsCustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getResponseField", GoMethod: "GetResponseField"},
			_jsii_.MemberMethod{JsiiMethod: "getResponseFieldReference", GoMethod: "GetResponseFieldReference"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCustomResource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.custom_resources.AwsCustomResourcePolicy",
		reflect.TypeOf((*AwsCustomResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "resources", GoGetter: "Resources"},
			_jsii_.MemberProperty{JsiiProperty: "statements", GoGetter: "Statements"},
		},
		func() interface{} {
			return &jsiiProxy_AwsCustomResourcePolicy{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.custom_resources.AwsCustomResourceProps",
		reflect.TypeOf((*AwsCustomResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.custom_resources.AwsSdkCall",
		reflect.TypeOf((*AwsSdkCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.custom_resources.PhysicalResourceId",
		reflect.TypeOf((*PhysicalResourceId)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "responsePath", GoGetter: "ResponsePath"},
		},
		func() interface{} {
			return &jsiiProxy_PhysicalResourceId{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.custom_resources.PhysicalResourceIdReference",
		reflect.TypeOf((*PhysicalResourceIdReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PhysicalResourceIdReference{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResolvable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.custom_resources.Provider",
		reflect.TypeOf((*Provider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "isCompleteHandler", GoGetter: "IsCompleteHandler"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onEventHandler", GoGetter: "OnEventHandler"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Provider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscloudformationICustomResourceProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.custom_resources.ProviderProps",
		reflect.TypeOf((*ProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.custom_resources.SdkCallsPolicyOptions",
		reflect.TypeOf((*SdkCallsPolicyOptions)(nil)).Elem(),
	)
}

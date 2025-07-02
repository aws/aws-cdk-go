package customresources

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		reflect.TypeOf((*AwsCustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getResponseField", GoMethod: "GetResponseField"},
			_jsii_.MemberMethod{JsiiMethod: "getResponseFieldReference", GoMethod: "GetResponseFieldReference"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCustomResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.AwsCustomResourcePolicy",
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
		"aws-cdk-lib.custom_resources.AwsCustomResourceProps",
		reflect.TypeOf((*AwsCustomResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.AwsSdkCall",
		reflect.TypeOf((*AwsSdkCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.CustomResourceConfig",
		reflect.TypeOf((*CustomResourceConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLambdaRuntime", GoMethod: "AddLambdaRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "addLogRetentionLifetime", GoMethod: "AddLogRetentionLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "addRemovalPolicy", GoMethod: "AddRemovalPolicy"},
		},
		func() interface{} {
			return &jsiiProxy_CustomResourceConfig{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.CustomResourceLambdaRuntime",
		reflect.TypeOf((*CustomResourceLambdaRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceLambdaRuntime{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.CustomResourceLogRetention",
		reflect.TypeOf((*CustomResourceLogRetention)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceLogRetention{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.CustomResourceRemovalPolicy",
		reflect.TypeOf((*CustomResourceRemovalPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceRemovalPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.LogOptions",
		reflect.TypeOf((*LogOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.Logging",
		reflect.TypeOf((*Logging)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Logging{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.LoggingProps",
		reflect.TypeOf((*LoggingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
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
		"aws-cdk-lib.custom_resources.PhysicalResourceIdReference",
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
		"aws-cdk-lib.custom_resources.Provider",
		reflect.TypeOf((*Provider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isCompleteHandler", GoGetter: "IsCompleteHandler"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onEventHandler", GoGetter: "OnEventHandler"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Provider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.ProviderProps",
		reflect.TypeOf((*ProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.SdkCallsPolicyOptions",
		reflect.TypeOf((*SdkCallsPolicyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.custom_resources.WaiterStateMachine",
		reflect.TypeOf((*WaiterStateMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantStartExecution", GoMethod: "GrantStartExecution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachineArn", GoGetter: "StateMachineArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WaiterStateMachine{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.custom_resources.WaiterStateMachineProps",
		reflect.TypeOf((*WaiterStateMachineProps)(nil)).Elem(),
	)
}

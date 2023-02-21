package awsglobalacceleratorendpoints

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpoint",
		reflect.TypeOf((*ApplicationLoadBalancerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderEndpointConfiguration", GoMethod: "RenderEndpointConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationLoadBalancerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsglobalacceleratorIEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.ApplicationLoadBalancerEndpointOptions",
		reflect.TypeOf((*ApplicationLoadBalancerEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.CfnEipEndpoint",
		reflect.TypeOf((*CfnEipEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderEndpointConfiguration", GoMethod: "RenderEndpointConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEipEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsglobalacceleratorIEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.CfnEipEndpointProps",
		reflect.TypeOf((*CfnEipEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.InstanceEndpoint",
		reflect.TypeOf((*InstanceEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderEndpointConfiguration", GoMethod: "RenderEndpointConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsglobalacceleratorIEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.InstanceEndpointProps",
		reflect.TypeOf((*InstanceEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpoint",
		reflect.TypeOf((*NetworkLoadBalancerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "renderEndpointConfiguration", GoMethod: "RenderEndpointConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkLoadBalancerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsglobalacceleratorIEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_globalaccelerator_endpoints.NetworkLoadBalancerEndpointProps",
		reflect.TypeOf((*NetworkLoadBalancerEndpointProps)(nil)).Elem(),
	)
}

package awselasticloadbalancingv2targets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbArnTarget",
		reflect.TypeOf((*AlbArnTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_AlbArnTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.AlbTarget",
		reflect.TypeOf((*AlbTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_AlbTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AlbArnTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceIdTarget",
		reflect.TypeOf((*InstanceIdTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceIdTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget)
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.InstanceTarget",
		reflect.TypeOf((*InstanceTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InstanceIdTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.IpTarget",
		reflect.TypeOf((*IpTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_IpTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget)
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.LambdaTarget",
		reflect.TypeOf((*LambdaTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget)
			return &j
		},
	)
}

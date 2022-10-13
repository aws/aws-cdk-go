package awsroute53targets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.ApiGateway",
		reflect.TypeOf((*ApiGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGateway{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiGatewayDomain)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayDomain",
		reflect.TypeOf((*ApiGatewayDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayDomain{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		reflect.TypeOf((*ApiGatewayv2DomainProperties)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayv2DomainProperties{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.BucketWebsiteTarget",
		reflect.TypeOf((*BucketWebsiteTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_BucketWebsiteTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.ClassicLoadBalancerTarget",
		reflect.TypeOf((*ClassicLoadBalancerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ClassicLoadBalancerTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.CloudFrontTarget",
		reflect.TypeOf((*CloudFrontTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFrontTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		reflect.TypeOf((*ElasticBeanstalkEnvironmentEndpointTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		reflect.TypeOf((*GlobalAcceleratorDomainTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalAcceleratorDomainTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		reflect.TypeOf((*GlobalAcceleratorTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalAcceleratorTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GlobalAcceleratorDomainTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		reflect.TypeOf((*InterfaceVpcEndpointTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InterfaceVpcEndpointTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.LoadBalancerTarget",
		reflect.TypeOf((*LoadBalancerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LoadBalancerTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.Route53RecordTarget",
		reflect.TypeOf((*Route53RecordTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Route53RecordTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		reflect.TypeOf((*UserPoolDomainTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_UserPoolDomainTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awsroute53IAliasRecordTarget)
			return &j
		},
	)
}

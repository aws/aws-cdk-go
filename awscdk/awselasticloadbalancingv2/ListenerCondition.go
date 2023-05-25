package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// ListenerCondition providers definition.
//
// Example:
//   var listener applicationListener
//   var asg autoScalingGroup
//
//
//   listener.AddTargets(jsii.String("Example.Com Fleet"), &AddApplicationTargetsProps{
//   	Priority: jsii.Number(10),
//   	Conditions: []listenerCondition{
//   		elbv2.*listenerCondition_HostHeaders([]*string{
//   			jsii.String("example.com"),
//   		}),
//   		elbv2.*listenerCondition_PathPatterns([]*string{
//   			jsii.String("/ok"),
//   			jsii.String("/path"),
//   		}),
//   	},
//   	Port: jsii.Number(8080),
//   	Targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type ListenerCondition interface {
	// Render the raw Cfn listener rule condition object.
	RenderRawCondition() interface{}
}

// The jsii proxy struct for ListenerCondition
type jsiiProxy_ListenerCondition struct {
	_ byte // padding
}

func NewListenerCondition_Override(l ListenerCondition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		nil, // no parameters
		l,
	)
}

// Create a host-header listener rule condition.
func ListenerCondition_HostHeaders(values *[]*string) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_HostHeadersParameters(values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"hostHeaders",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a http-header listener rule condition.
func ListenerCondition_HttpHeader(name *string, values *[]*string) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_HttpHeaderParameters(name, values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"httpHeader",
		[]interface{}{name, values},
		&returns,
	)

	return returns
}

// Create a http-request-method listener rule condition.
func ListenerCondition_HttpRequestMethods(values *[]*string) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_HttpRequestMethodsParameters(values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"httpRequestMethods",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a path-pattern listener rule condition.
func ListenerCondition_PathPatterns(values *[]*string) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_PathPatternsParameters(values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"pathPatterns",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a query-string listener rule condition.
func ListenerCondition_QueryStrings(values *[]*QueryStringCondition) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_QueryStringsParameters(values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"queryStrings",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a source-ip listener rule condition.
func ListenerCondition_SourceIps(values *[]*string) ListenerCondition {
	_init_.Initialize()

	if err := validateListenerCondition_SourceIpsParameters(values); err != nil {
		panic(err)
	}
	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"sourceIps",
		[]interface{}{values},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListenerCondition) RenderRawCondition() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderRawCondition",
		nil, // no parameters
		&returns,
	)

	return returns
}


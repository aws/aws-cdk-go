package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// ListenerCondition providers definition.
//
// Example:
//   var listener applicationListener
//   var asg autoScalingGroup
//
//
//   listener.addTargets(jsii.String("Example.Com Fleet"), &addApplicationTargetsProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.hostHeaders([]*string{
//   			jsii.String("example.com"),
//   		}),
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   			jsii.String("/path"),
//   		}),
//   	},
//   	port: jsii.Number(8080),
//   	targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
// Experimental.
type ListenerCondition interface {
	// Render the raw Cfn listener rule condition object.
	// Experimental.
	RenderRawCondition() interface{}
}

// The jsii proxy struct for ListenerCondition
type jsiiProxy_ListenerCondition struct {
	_ byte // padding
}

// Experimental.
func NewListenerCondition_Override(l ListenerCondition) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		nil, // no parameters
		l,
	)
}

// Create a host-header listener rule condition.
// Experimental.
func ListenerCondition_HostHeaders(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		"hostHeaders",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a http-header listener rule condition.
// Experimental.
func ListenerCondition_HttpHeader(name *string, values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		"httpHeader",
		[]interface{}{name, values},
		&returns,
	)

	return returns
}

// Create a http-request-method listener rule condition.
// Experimental.
func ListenerCondition_HttpRequestMethods(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		"httpRequestMethods",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a path-pattern listener rule condition.
// Experimental.
func ListenerCondition_PathPatterns(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		"pathPatterns",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a query-string listener rule condition.
// Experimental.
func ListenerCondition_QueryStrings(values *[]*QueryStringCondition) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
		"queryStrings",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a source-ip listener rule condition.
// Experimental.
func ListenerCondition_SourceIps(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ListenerCondition",
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


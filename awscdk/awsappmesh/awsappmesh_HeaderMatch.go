package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Used to generate header matching methods.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &httpRouteMatch{
//   			path: appmesh.httpRoutePathMatch.exactly(jsii.String("/exact")),
//   			method: appmesh.httpRouteMethod_POST,
//   			protocol: appmesh.httpRouteProtocol_HTTPS,
//   			headers: []headerMatch{
//   				appmesh.*headerMatch.valueIs(jsii.String("Content-Type"), jsii.String("application/json")),
//   				appmesh.*headerMatch.valueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
//   			},
//   			queryParameters: []queryParameterMatch{
//   				appmesh.*queryParameterMatch.valueIs(jsii.String("query-field"), jsii.String("value")),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type HeaderMatch interface {
	// Returns the header match configuration.
	// Experimental.
	Bind(scope awscdk.Construct) *HeaderMatchConfig
}

// The jsii proxy struct for HeaderMatch
type jsiiProxy_HeaderMatch struct {
	_ byte // padding
}

// Experimental.
func NewHeaderMatch_Override(h HeaderMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HeaderMatch",
		nil, // no parameters
		h,
	)
}

// The value of the header with the given name in the request must not end with the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotEndWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotEndWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not include the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotMatchRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotMatchRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not start with the specified characters.
// Experimental.
func HeaderMatch_ValueDoesNotStartWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueDoesNotStartWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must end with the specified characters.
// Experimental.
func HeaderMatch_ValueEndsWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueEndsWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must match the specified value exactly.
// Experimental.
func HeaderMatch_ValueIs(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueIs",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not match the specified value exactly.
// Experimental.
func HeaderMatch_ValueIsNot(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueIsNot",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must include the specified characters.
// Experimental.
func HeaderMatch_ValueMatchesRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueMatchesRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must be in a range of values.
// Experimental.
func HeaderMatch_ValuesIsInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valuesIsInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not be in a range of values.
// Experimental.
func HeaderMatch_ValuesIsNotInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valuesIsNotInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must start with the specified characters.
// Experimental.
func HeaderMatch_ValueStartsWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HeaderMatch",
		"valueStartsWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HeaderMatch) Bind(scope awscdk.Construct) *HeaderMatchConfig {
	var returns *HeaderMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate header matching methods.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type HeaderMatch interface {
	// Returns the header match configuration.
	Bind(scope constructs.Construct) *HeaderMatchConfig
}

// The jsii proxy struct for HeaderMatch
type jsiiProxy_HeaderMatch struct {
	_ byte // padding
}

func NewHeaderMatch_Override(h HeaderMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		nil, // no parameters
		h,
	)
}

// The value of the header with the given name in the request must not end with the specified characters.
func HeaderMatch_ValueDoesNotEndWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueDoesNotEndWithParameters(headerName, suffix); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotEndWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not include the specified characters.
func HeaderMatch_ValueDoesNotMatchRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueDoesNotMatchRegexParameters(headerName, regex); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotMatchRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not start with the specified characters.
func HeaderMatch_ValueDoesNotStartWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueDoesNotStartWithParameters(headerName, prefix); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueDoesNotStartWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must end with the specified characters.
func HeaderMatch_ValueEndsWith(headerName *string, suffix *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueEndsWithParameters(headerName, suffix); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueEndsWith",
		[]interface{}{headerName, suffix},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must match the specified value exactly.
func HeaderMatch_ValueIs(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueIsParameters(headerName, headerValue); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueIs",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not match the specified value exactly.
func HeaderMatch_ValueIsNot(headerName *string, headerValue *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueIsNotParameters(headerName, headerValue); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueIsNot",
		[]interface{}{headerName, headerValue},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must include the specified characters.
func HeaderMatch_ValueMatchesRegex(headerName *string, regex *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueMatchesRegexParameters(headerName, regex); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueMatchesRegex",
		[]interface{}{headerName, regex},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must be in a range of values.
func HeaderMatch_ValuesIsInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValuesIsInRangeParameters(headerName, start, end); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valuesIsInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must not be in a range of values.
func HeaderMatch_ValuesIsNotInRange(headerName *string, start *float64, end *float64) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValuesIsNotInRangeParameters(headerName, start, end); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valuesIsNotInRange",
		[]interface{}{headerName, start, end},
		&returns,
	)

	return returns
}

// The value of the header with the given name in the request must start with the specified characters.
func HeaderMatch_ValueStartsWith(headerName *string, prefix *string) HeaderMatch {
	_init_.Initialize()

	if err := validateHeaderMatch_ValueStartsWithParameters(headerName, prefix); err != nil {
		panic(err)
	}
	var returns HeaderMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HeaderMatch",
		"valueStartsWith",
		[]interface{}{headerName, prefix},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HeaderMatch) Bind(scope constructs.Construct) *HeaderMatchConfig {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *HeaderMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


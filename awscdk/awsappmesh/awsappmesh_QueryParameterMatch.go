package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate query parameter matching methods.
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
type QueryParameterMatch interface {
	// Returns the query parameter match configuration.
	Bind(scope constructs.Construct) *QueryParameterMatchConfig
}

// The jsii proxy struct for QueryParameterMatch
type jsiiProxy_QueryParameterMatch struct {
	_ byte // padding
}

func NewQueryParameterMatch_Override(q QueryParameterMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatch",
		nil, // no parameters
		q,
	)
}

// The value of the query parameter with the given name in the request must match the specified value exactly.
func QueryParameterMatch_ValueIs(queryParameterName *string, queryParameterValue *string) QueryParameterMatch {
	_init_.Initialize()

	var returns QueryParameterMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.QueryParameterMatch",
		"valueIs",
		[]interface{}{queryParameterName, queryParameterValue},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryParameterMatch) Bind(scope constructs.Construct) *QueryParameterMatchConfig {
	var returns *QueryParameterMatchConfig

	_jsii_.Invoke(
		q,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


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
//   router.addRoute(jsii.String("route-http2"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_Exactly(jsii.String("/exact")),
//   			Method: appmesh.HttpRouteMethod_POST,
//   			Protocol: appmesh.HttpRouteProtocol_HTTPS,
//   			Headers: []headerMatch{
//   				appmesh.*headerMatch_ValueIs(jsii.String("Content-Type"), jsii.String("application/json")),
//   				appmesh.*headerMatch_ValueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
//   			},
//   			QueryParameters: []queryParameterMatch{
//   				appmesh.*queryParameterMatch_ValueIs(jsii.String("query-field"), jsii.String("value")),
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

	if err := validateQueryParameterMatch_ValueIsParameters(queryParameterName, queryParameterValue); err != nil {
		panic(err)
	}
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
	if err := q.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *QueryParameterMatchConfig

	_jsii_.Invoke(
		q,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


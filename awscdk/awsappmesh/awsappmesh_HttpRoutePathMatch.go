package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Defines HTTP route matching based on the URL path of the request.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   				weight: jsii.Number(50),
//   			},
//   			&weightedTarget{
//   				virtualNode: node,
//   				weight: jsii.Number(50),
//   			},
//   		},
//   		match: &httpRouteMatch{
//   			path: appmesh.httpRoutePathMatch.startsWith(jsii.String("/path-to-app")),
//   		},
//   	}),
//   })
//
// Experimental.
type HttpRoutePathMatch interface {
	// Returns the route path match configuration.
	// Experimental.
	Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig
}

// The jsii proxy struct for HttpRoutePathMatch
type jsiiProxy_HttpRoutePathMatch struct {
	_ byte // padding
}

// Experimental.
func NewHttpRoutePathMatch_Override(h HttpRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
// Experimental.
func HttpRoutePathMatch_Exactly(path *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"exactly",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
// Experimental.
func HttpRoutePathMatch_Regex(regex *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"regex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
// Experimental.
func HttpRoutePathMatch_StartsWith(prefix *string) HttpRoutePathMatch {
	_init_.Initialize()

	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"startsWith",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpRoutePathMatch) Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig {
	var returns *HttpRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


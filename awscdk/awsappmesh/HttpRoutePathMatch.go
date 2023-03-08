package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines HTTP route matching based on the URL path of the request.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   			&weightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_StartsWith(jsii.String("/path-to-app")),
//   		},
//   	}),
//   })
//
type HttpRoutePathMatch interface {
	// Returns the route path match configuration.
	Bind(scope constructs.Construct) *HttpRoutePathMatchConfig
}

// The jsii proxy struct for HttpRoutePathMatch
type jsiiProxy_HttpRoutePathMatch struct {
	_ byte // padding
}

func NewHttpRoutePathMatch_Override(h HttpRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
func HttpRoutePathMatch_Exactly(path *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_ExactlyParameters(path); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"exactly",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
func HttpRoutePathMatch_Regex(regex *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_RegexParameters(regex); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"regex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
func HttpRoutePathMatch_StartsWith(prefix *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_StartsWithParameters(prefix); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpRoutePathMatch",
		"startsWith",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpRoutePathMatch) Bind(scope constructs.Construct) *HttpRoutePathMatchConfig {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *HttpRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


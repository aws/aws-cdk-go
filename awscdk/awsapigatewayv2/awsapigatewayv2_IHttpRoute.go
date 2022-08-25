package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Represents a Route for an HTTP API.
// Experimental.
type IHttpRoute interface {
	IRoute
	// Grant access to invoke the route.
	//
	// This method requires that the authorizer of the route is undefined or is
	// an `HttpIamAuthorizer`.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable, options *GrantInvokeOptions) awsiam.Grant
	// The HTTP API associated with this route.
	// Experimental.
	HttpApi() IHttpApi
	// Returns the path component of this HTTP route, `undefined` if the path is the catch-all route.
	// Experimental.
	Path() *string
	// Returns the arn of the route.
	// Experimental.
	RouteArn() *string
}

// The jsii proxy for IHttpRoute
type jsiiProxy_IHttpRoute struct {
	jsiiProxy_IRoute
}

func (i *jsiiProxy_IHttpRoute) GrantInvoke(grantee awsiam.IGrantable, options *GrantInvokeOptions) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IHttpRoute) HttpApi() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"httpApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpRoute) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpRoute) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}


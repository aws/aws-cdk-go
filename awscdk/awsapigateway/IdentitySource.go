package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an identity source.
//
// The source can be specified either as a literal value (e.g: `Auth`) which
// cannot be blank, or as an unresolved string token.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authFn,
//   	IdentitySources: []*string{
//   		apigateway.IdentitySource_Header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   })
//
type IdentitySource interface {
}

// The jsii proxy struct for IdentitySource
type jsiiProxy_IdentitySource struct {
	_ byte // padding
}

func NewIdentitySource() IdentitySource {
	_init_.Initialize()

	j := jsiiProxy_IdentitySource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIdentitySource_Override(i IdentitySource) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		nil, // no parameters
		i,
	)
}

// Provides a properly formatted request context identity source.
//
// Returns: a request context identity source.
func IdentitySource_Context(context *string) *string {
	_init_.Initialize()

	if err := validateIdentitySource_ContextParameters(context); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"context",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Provides a properly formatted header identity source.
//
// Returns: a header identity source.
func IdentitySource_Header(headerName *string) *string {
	_init_.Initialize()

	if err := validateIdentitySource_HeaderParameters(headerName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"header",
		[]interface{}{headerName},
		&returns,
	)

	return returns
}

// Provides a properly formatted query string identity source.
//
// Returns: a query string identity source.
func IdentitySource_QueryString(queryString *string) *string {
	_init_.Initialize()

	if err := validateIdentitySource_QueryStringParameters(queryString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"queryString",
		[]interface{}{queryString},
		&returns,
	)

	return returns
}

// Provides a properly formatted API Gateway stage variable identity source.
//
// Returns: an API Gateway stage variable identity source.
func IdentitySource_StageVariable(stageVariable *string) *string {
	_init_.Initialize()

	if err := validateIdentitySource_StageVariableParameters(stageVariable); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.IdentitySource",
		"stageVariable",
		[]interface{}{stageVariable},
		&returns,
	)

	return returns
}


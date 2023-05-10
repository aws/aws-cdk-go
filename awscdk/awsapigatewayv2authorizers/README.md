# AWS APIGatewayv2 Authorizers

## Table of Contents

* [Introduction](#introduction)
* [HTTP APIs](#http-apis)

  * [Default Authorization](#default-authorization)
  * [Route Authorization](#route-authorization)
  * [JWT Authorizers](#jwt-authorizers)

    * [User Pool Authorizer](#user-pool-authorizer)
  * [Lambda Authorizers](#lambda-authorizers)
  * [IAM Authorizers](#iam-authorizers)
* [WebSocket APIs](#websocket-apis)

  * [Lambda Authorizer](#lambda-authorizer)

## Introduction

API Gateway supports multiple mechanisms for controlling and managing access to your HTTP API. They are mainly
classified into Lambda Authorizers, JWT authorizers and standard AWS IAM roles and policies. More information is
available at [Controlling and managing access to an HTTP
API](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html).

## HTTP APIs

Access control for Http Apis is managed by restricting which routes can be invoked via.

Authorizers and scopes can either be applied to the api, or specifically for each route.

### Default Authorization

When using default authorization, all routes of the api will inherit the configuration.

In the example below, all routes will require the `manage:books` scope present in order to invoke the integration.

```go
import "github.com/aws/aws-cdk-go/awscdk"


issuer := "https://test.us.auth0.com"
authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("DefaultAuthorizer"), issuer, &HttpJwtAuthorizerProps{
	JwtAudience: []*string{
		jsii.String("3131231"),
	},
})

api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
	DefaultAuthorizer: authorizer,
	DefaultAuthorizationScopes: []*string{
		jsii.String("manage:books"),
	},
})
```

### Route Authorization

Authorization can also configured for each Route. When a route authorization is configured, it takes precedence over default authorization.

The example below showcases default authorization, along with route authorization. It also shows how to remove authorization entirely for a route.

* `GET /books` and `GET /books/{id}` use the default authorizer settings on the api
* `POST /books` will require the [write:books] scope
* `POST /login` removes the default authorizer (unauthenticated route)

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


issuer := "https://test.us.auth0.com"
authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("DefaultAuthorizer"), issuer, &HttpJwtAuthorizerProps{
	JwtAudience: []*string{
		jsii.String("3131231"),
	},
})

api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
	DefaultAuthorizer: authorizer,
	DefaultAuthorizationScopes: []*string{
		jsii.String("read:books"),
	},
})

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books"),
	Methods: []httpMethod{
		apigwv2.*httpMethod_GET,
	},
})

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIdIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books/{id}"),
	Methods: []*httpMethod{
		apigwv2.*httpMethod_GET,
	},
})

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books"),
	Methods: []*httpMethod{
		apigwv2.*httpMethod_POST,
	},
	AuthorizationScopes: []*string{
		jsii.String("write:books"),
	},
})

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("LoginIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/login"),
	Methods: []*httpMethod{
		apigwv2.*httpMethod_POST,
	},
	Authorizer: apigwv2.NewHttpNoneAuthorizer(),
})
```

### JWT Authorizers

JWT authorizers allow the use of JSON Web Tokens (JWTs) as part of [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html) and [OAuth 2.0](https://oauth.net/2/) frameworks to allow and restrict clients from accessing HTTP APIs.

When configured, API Gateway validates the JWT submitted by the client, and allows or denies access based on its content.

The location of the token is defined by the `identitySource` which defaults to the http `Authorization` header. However it also
[supports a number of other options](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html#http-api-lambda-authorizer.identity-sources).
It then decodes the JWT and validates the signature and claims, against the options defined in the authorizer and route (scopes).
For more information check the [JWT Authorizer documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-jwt-authorizer.html).

Clients that fail authorization are presented with either 2 responses:

* `401 - Unauthorized` - When the JWT validation fails
* `403 - Forbidden` - When the JWT validation is successful but the required scopes are not met

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


issuer := "https://test.us.auth0.com"
authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &HttpJwtAuthorizerProps{
	JwtAudience: []*string{
		jsii.String("3131231"),
	},
})

api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books"),
	Authorizer: Authorizer,
})
```

#### User Pool Authorizer

User Pool Authorizer is a type of JWT Authorizer that uses a Cognito user pool and app client to control who can access your Api. After a successful authorization from the app client, the generated access token will be used as the JWT.

Clients accessing an API that uses a user pool authorizer must first sign in to a user pool and obtain an identity or access token.
They must then use this token in the specified `identitySource` for the API call. More information is available at [using Amazon Cognito user
pools as authorizer](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html).

```go
import cognito "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


userPool := cognito.NewUserPool(this, jsii.String("UserPool"))

authorizer := awscdk.NewHttpUserPoolAuthorizer(jsii.String("BooksAuthorizer"), userPool)

api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books"),
	Authorizer: Authorizer,
})
```

### Lambda Authorizers

Lambda authorizers use a Lambda function to control access to your HTTP API. When a client calls your API, API Gateway invokes your Lambda function and uses the response to determine whether the client can access your API.

Lambda authorizers depending on their response, fall into either two types - Simple or IAM. You can learn about differences [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html#http-api-lambda-authorizer.payload-format-response).

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// This function handles your auth logic
var authHandler function


authorizer := awscdk.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &HttpLambdaAuthorizerProps{
	ResponseTypes: []httpLambdaResponseType{
		awscdk.HttpLambdaResponseType_SIMPLE,
	},
})

api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

api.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books"),
	Authorizer: Authorizer,
})
```

### IAM Authorizers

API Gateway supports IAM via the included `HttpIamAuthorizer` and grant syntax:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var principal anyPrincipal


authorizer := awscdk.NewHttpIamAuthorizer()

httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
	DefaultAuthorizer: authorizer,
})

routes := httpApi.AddRoutes(&AddRoutesOptions{
	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
	Path: jsii.String("/books/{book}"),
})

routes[0].GrantInvoke(principal)
```

## WebSocket APIs

You can set an authorizer to your WebSocket API's `$connect` route to control access to your API.

### Lambda Authorizer

Lambda authorizers use a Lambda function to control access to your WebSocket API. When a client connects to your API, API Gateway invokes your Lambda function and uses the response to determine whether the client can access your API.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// This function handles your auth logic
var authHandler function

// This function handles your WebSocket requests
var handler function


authorizer := awscdk.NewWebSocketLambdaAuthorizer(jsii.String("Authorizer"), authHandler)

integration := awscdk.NewWebSocketLambdaIntegration(jsii.String("Integration"), handler)

apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &WebSocketApiProps{
	ConnectRouteOptions: &WebSocketRouteOptions{
		Integration: *Integration,
		Authorizer: *Authorizer,
	},
})
```

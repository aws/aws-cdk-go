# AWS::APIGatewayv2 Construct Library

<!--BEGIN STABILITY BANNER-->---


Features                                   | Stability
-------------------------------------------|--------------------------------------------------------
Higher level constructs for HTTP APIs      | ![Experimental](https://img.shields.io/badge/experimental-important.svg?style=for-the-badge)
Higher level constructs for Websocket APIs | ![Experimental](https://img.shields.io/badge/experimental-important.svg?style=for-the-badge)

> **Experimental:** Higher level constructs in this module that are marked as experimental are
> under active development. They are subject to non-backward compatible changes or removal in any
> future version. These are not subject to the [Semantic Versioning](https://semver.org/) model and
> breaking changes will be announced in the release notes. This means that while you may use them,
> you may need to update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## Table of Contents

* [Introduction](#introduction)
* [HTTP API](#http-api)

  * [Defining HTTP APIs](#defining-http-apis)
  * [Cross Origin Resource Sharing (CORS)](#cross-origin-resource-sharing-cors)
  * [Publishing HTTP APIs](#publishing-http-apis)
  * [Custom Domain](#custom-domain)
  * [Mutual TLS](#mutual-tls-mtls)
  * [Managing access to HTTP APIs](#managing-access-to-http-apis)
  * [Metrics](#metrics)
  * [VPC Link](#vpc-link)
  * [Private Integration](#private-integration)
* [WebSocket API](#websocket-api)

  * [Manage Connections Permission](#manage-connections-permission)
  * [Managing access to WebSocket APIs](#managing-access-to-websocket-apis)

## Introduction

Amazon API Gateway is an AWS service for creating, publishing, maintaining, monitoring, and securing REST, HTTP, and WebSocket
APIs at any scale. API developers can create APIs that access AWS or other web services, as well as data stored in the AWS Cloud.
As an API Gateway API developer, you can create APIs for use in your own client applications. Read the
[Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html).

This module supports features under [API Gateway v2](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_ApiGatewayV2.html)
that lets users set up Websocket and HTTP APIs.
REST APIs can be created using the `@aws-cdk/aws-apigateway` module.

## HTTP API

HTTP APIs enable creation of RESTful APIs that integrate with AWS Lambda functions, known as Lambda proxy integration,
or to any routable HTTP endpoint, known as HTTP proxy integration.

### Defining HTTP APIs

HTTP APIs have two fundamental concepts - Routes and Integrations.

Routes direct incoming API requests to backend resources. Routes consist of two parts: an HTTP method and a resource
path, such as, `GET /books`. Learn more at [Working with
routes](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Use the `ANY` method
to match any methods for a route that are not explicitly defined.

Integrations define how the HTTP API responds when a client reaches a specific Route. HTTP APIs support Lambda proxy
integration, HTTP proxy integration and, AWS service integrations, also known as private integrations. Learn more at
[Configuring integrations](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations.html).

Integrations are available at the `aws-apigatewayv2-integrations` module and more information is available in that module.
As an early example, the following code snippet configures a route `GET /books` with an HTTP proxy integration all
configures all other HTTP method calls to `/books` to a lambda proxy.

```go
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"

var booksDefaultFn function


getBooksIntegration := awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal"))
booksDefaultIntegration := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)

httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

httpApi.addRoutes(&addRoutesOptions{
	path: jsii.String("/books"),
	methods: []httpMethod{
		apigwv2.*httpMethod_GET,
	},
	integration: getBooksIntegration,
})
httpApi.addRoutes(&addRoutesOptions{
	path: jsii.String("/books"),
	methods: []*httpMethod{
		apigwv2.*httpMethod_ANY,
	},
	integration: booksDefaultIntegration,
})
```

The URL to the endpoint can be retrieved via the `apiEndpoint` attribute. By default this URL is enabled for clients. Use `disableExecuteApiEndpoint` to disable it.

```go
httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &httpApiProps{
	disableExecuteApiEndpoint: jsii.Boolean(true),
})
```

The `defaultIntegration` option while defining HTTP APIs lets you create a default catch-all integration that is
matched when a client reaches a route that is not explicitly defined.

```go
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"


apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &httpApiProps{
	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("DefaultIntegration"), jsii.String("https://example.com")),
})
```

### Cross Origin Resource Sharing (CORS)

[Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security
feature that restricts HTTP requests that are initiated from scripts running in the browser. Enabling CORS will allow
requests to your API from a web application hosted in a domain different from your API domain.

When configured CORS for an HTTP API, API Gateway automatically sends a response to preflight `OPTIONS` requests, even
if there isn't an `OPTIONS` route configured. Note that, when this option is used, API Gateway will ignore CORS headers
returned from your backend integration. Learn more about [Configuring CORS for an HTTP
API](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html).

The `corsPreflight` option lets you specify a CORS configuration for an API.

```go
apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &httpApiProps{
	corsPreflight: &corsPreflightOptions{
		allowHeaders: []*string{
			jsii.String("Authorization"),
		},
		allowMethods: []corsHttpMethod{
			apigwv2.*corsHttpMethod_GET,
			apigwv2.*corsHttpMethod_HEAD,
			apigwv2.*corsHttpMethod_OPTIONS,
			apigwv2.*corsHttpMethod_POST,
		},
		allowOrigins: []*string{
			jsii.String("*"),
		},
		maxAge: awscdk.Duration.days(jsii.Number(10)),
	},
})
```

### Publishing HTTP APIs

A Stage is a logical reference to a lifecycle state of your API (for example, `dev`, `prod`, `beta`, or `v2`). API
stages are identified by their stage name. Each stage is a named reference to a deployment of the API made available for
client applications to call.

Use `HttpStage` to create a Stage resource for HTTP APIs. The following code sets up a Stage, whose URL is available at
`https://{api_id}.execute-api.{region}.amazonaws.com/beta`.

```go
var api httpApi


apigwv2.NewHttpStage(this, jsii.String("Stage"), &httpStageProps{
	httpApi: api,
	stageName: jsii.String("beta"),
})
```

If you omit the `stageName` will create a `$default` stage. A `$default` stage is one that is served from the base of
the API's URL - `https://{api_id}.execute-api.{region}.amazonaws.com/`.

Note that, `HttpApi` will always creates a `$default` stage, unless the `createDefaultStage` property is unset.

### Custom Domain

Custom domain names are simpler and more intuitive URLs that you can provide to your API users. Custom domain name are associated to API stages.

The code snippet below creates a custom domain and configures a default domain mapping for your API that maps the
custom domain to the `$default` stage of the API.

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"

var handler function


certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
domainName := "example.com"

dn := apigwv2.NewDomainName(this, jsii.String("DN"), &domainNameProps{
	domainName: jsii.String(domainName),
	certificate: acm.certificate.fromCertificateArn(this, jsii.String("cert"), certArn),
})
api := apigwv2.NewHttpApi(this, jsii.String("HttpProxyProdApi"), &httpApiProps{
	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
	// https://${dn.domainName}/foo goes to prodApi $default stage
	defaultDomainMapping: &domainMappingOptions{
		domainName: dn,
		mappingKey: jsii.String("foo"),
	},
})
```

To migrate a domain endpoint from one type to another, you can add a new endpoint configuration via `addEndpoint()`
and then configure DNS records to route traffic to the new endpoint. After that, you can remove the previous endpoint configuration.
Learn more at [Migrating a custom domain name](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-migrate.html)

To associate a specific `Stage` to a custom domain mapping -

```go
var api httpApi
var dn domainName


api.addStage(jsii.String("beta"), &httpStageOptions{
	stageName: jsii.String("beta"),
	autoDeploy: jsii.Boolean(true),
	// https://${dn.domainName}/bar goes to the beta stage
	domainMapping: &domainMappingOptions{
		domainName: dn,
		mappingKey: jsii.String("bar"),
	},
})
```

The same domain name can be associated with stages across different `HttpApi` as so -

```go
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"

var handler function
var dn domainName


apiDemo := apigwv2.NewHttpApi(this, jsii.String("DemoApi"), &httpApiProps{
	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
	// https://${dn.domainName}/demo goes to apiDemo $default stage
	defaultDomainMapping: &domainMappingOptions{
		domainName: dn,
		mappingKey: jsii.String("demo"),
	},
})
```

The `mappingKey` determines the base path of the URL with the custom domain. Each custom domain is only allowed
to have one API mapping with undefined `mappingKey`. If more than one API mappings are specified, `mappingKey` will be required for all of them. In the sample above, the custom domain is associated
with 3 API mapping resources across different APIs and Stages.

|        API     |     Stage   |   URL  |
| :------------: | :---------: | :----: |
| api | $default  |   `https://${domainName}/foo`  |
| api | beta  |   `https://${domainName}/bar`  |
| apiDemo | $default  |   `https://${domainName}/demo`  |

You can retrieve the full domain URL with mapping key using the `domainUrl` property as so -

```go
var apiDemo httpApi

demoDomainUrl := apiDemo.defaultStage.domainUrl
```

### Mutual TLS (mTLS)

Mutual TLS can be configured to limit access to your API based by using client certificates instead of (or as an extension of) using authorization headers.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
import acm "github.com/aws/aws-cdk-go/awscdk"
var bucket bucket


certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
domainName := "example.com"

apigwv2.NewDomainName(this, jsii.String("DomainName"), &domainNameProps{
	domainName: jsii.String(domainName),
	certificate: acm.certificate.fromCertificateArn(this, jsii.String("cert"), certArn),
	mtls: &mTLSConfig{
		bucket: bucket,
		key: jsii.String("someca.pem"),
		version: jsii.String("version"),
	},
})
```

Instructions for configuring your trust store can be found [here](https://aws.amazon.com/blogs/compute/introducing-mutual-tls-authentication-for-amazon-api-gateway/)

### Managing access to HTTP APIs

API Gateway supports multiple mechanisms for [controlling and managing access to your HTTP
API](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) through authorizers.

These authorizers can be found in the [APIGatewayV2-Authorizers](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-apigatewayv2-authorizers-readme.html) constructs library.

### Metrics

The API Gateway v2 service sends metrics around the performance of HTTP APIs to Amazon CloudWatch.
These metrics can be referred to using the metric APIs available on the `HttpApi` construct.
The APIs with the `metric` prefix can be used to get reference to specific metrics for this API. For example,
the method below refers to the client side errors metric for this API.

```go
api := apigwv2.NewHttpApi(this, jsii.String("my-api"))
clientErrorMetric := api.metricClientError()
```

Please note that this will return a metric for all the stages defined in the api. It is also possible to refer to metrics for a specific Stage using
the `metric` methods from the `Stage` construct.

```go
api := apigwv2.NewHttpApi(this, jsii.String("my-api"))
stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &httpStageProps{
	httpApi: api,
})
clientErrorMetric := stage.metricClientError()
```

### VPC Link

Private integrations let HTTP APIs connect with AWS resources that are placed behind a VPC. These are usually Application
Load Balancers, Network Load Balancers or a Cloud Map service. The `VpcLink` construct enables this integration.
The following code creates a `VpcLink` to a private VPC.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &vpcLinkProps{
	vpc: vpc,
})
```

Any existing `VpcLink` resource can be imported into the CDK app via the `VpcLink.fromVpcLinkAttributes()`.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

awesomeLink := apigwv2.vpcLink.fromVpcLinkAttributes(this, jsii.String("awesome-vpc-link"), &vpcLinkAttributes{
	vpcLinkId: jsii.String("us-east-1_oiuR12Abd"),
	vpc: vpc,
})
```

### Private Integration

Private integrations enable integrating an HTTP API route with private resources in a VPC, such as Application Load Balancers or
Amazon ECS container-based applications.  Using private integrations, resources in a VPC can be exposed for access by
clients outside of the VPC.

These integrations can be found in the [aws-apigatewayv2-integrations](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-apigatewayv2-integrations-readme.html) constructs library.

## WebSocket API

A WebSocket API in API Gateway is a collection of WebSocket routes that are integrated with backend HTTP endpoints,
Lambda functions, or other AWS services. You can use API Gateway features to help you with all aspects of the API
lifecycle, from creation through monitoring your production APIs. [Read more](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html)

WebSocket APIs have two fundamental concepts - Routes and Integrations.

WebSocket APIs direct JSON messages to backend integrations based on configured routes. (Non-JSON messages are directed
to the configured `$default` route.)

Integrations define how the WebSocket API behaves when a client reaches a specific Route. Learn more at
[Configuring integrations](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-requests.html).

Integrations are available in the `aws-apigatewayv2-integrations` module and more information is available in that module.

To add the default WebSocket routes supported by API Gateway (`$connect`, `$disconnect` and `$default`), configure them as part of api props:

```go
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"

var connectHandler function
var disconnectHandler function
var defaultHandler function


webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &webSocketApiProps{
	connectRouteOptions: &webSocketRouteOptions{
		integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("ConnectIntegration"), connectHandler),
	},
	disconnectRouteOptions: &webSocketRouteOptions{
		integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("DisconnectIntegration"), disconnectHandler),
	},
	defaultRouteOptions: &webSocketRouteOptions{
		integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("DefaultIntegration"), defaultHandler),
	},
})

apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &webSocketStageProps{
	webSocketApi: webSocketApi,
	stageName: jsii.String("dev"),
	autoDeploy: jsii.Boolean(true),
})
```

To retrieve a websocket URL and a callback URL:

```go
var webSocketStage webSocketStage


webSocketURL := webSocketStage.url
// wss://${this.api.apiId}.execute-api.${s.region}.${s.urlSuffix}/${urlPath}
callbackURL := webSocketStage.callbackUrl
```

To add any other route:

```go
import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"

var messageHandler function

webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
webSocketApi.addRoute(jsii.String("sendmessage"), &webSocketRouteOptions{
	integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
})
```

To import an existing WebSocketApi:

```go
webSocketApi := apigwv2.webSocketApi.fromWebSocketApiAttributes(this, jsii.String("mywsapi"), &webSocketApiAttributes{
	webSocketId: jsii.String("api-1234"),
})
```

### Manage Connections Permission

Grant permission to use API Gateway Management API of a WebSocket API by calling the `grantManageConnections` API.
You can use Management API to send a callback message to a connected client, get connection information, or disconnect the client. Learn more at [Use @connections commands in your backend service](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).

```go
var fn function


webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
stage := apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &webSocketStageProps{
	webSocketApi: webSocketApi,
	stageName: jsii.String("dev"),
})
// per stage permission
stage.grantManagementApiAccess(fn)
// for all the stages permission
webSocketApi.grantManageConnections(fn)
```

### Managing access to WebSocket APIs

API Gateway supports multiple mechanisms for [controlling and managing access to a WebSocket API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) through authorizers.

These authorizers can be found in the [APIGatewayV2-Authorizers](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-apigatewayv2-authorizers-readme.html) constructs library.

### API Keys

Websocket APIs also support usage of API Keys. An API Key is a key that is used to grant access to an API. These are useful for controlling and tracking access to an API, when used together with [usage plans](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html). These together allow you to configure controls around API access such as quotas and throttling, along with per-API Key metrics on usage.

To require an API Key when accessing the Websocket API:

```go
webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &webSocketApiProps{
	apiKeySelectionExpression: apigwv2.webSocketApiKeySelectionExpression_HEADER_X_API_KEY(),
})
```

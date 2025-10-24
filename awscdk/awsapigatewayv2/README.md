# AWS APIGatewayv2 Construct Library

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
  * [Generating ARN for Execute API](#generating-arn-for-execute-api)
  * [Access Logging](#access-logging)
* [WebSocket API](#websocket-api)

  * [Manage Connections Permission](#manage-connections-permission)
  * [Managing access to WebSocket APIs](#managing-access-to-websocket-apis)
  * [Usage Plan and API Keys](#usage-plan-and-api-keys)

## Introduction

Amazon API Gateway is an AWS service for creating, publishing, maintaining, monitoring, and securing REST, HTTP, and WebSocket
APIs at any scale. API developers can create APIs that access AWS or other web services, as well as data stored in the AWS Cloud.
As an API Gateway API developer, you can create APIs for use in your own client applications. Read the
[Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html).

This module supports features under [API Gateway v2](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_ApiGatewayV2.html)
that lets users set up Websocket and HTTP APIs.
REST APIs can be created using the `aws-cdk-lib/aws-apigateway` module.

HTTP and Websocket APIs use the same CloudFormation resources under the hood. However, this module separates them into two separate constructs for a more efficient abstraction since there are a number of CloudFormation properties that specifically apply only to each type of API.

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
As an early example, we have a website for a bookstore where the following code snippet configures a route `GET /books` with an HTTP proxy integration. All other HTTP method calls to `/books` route to a default lambda proxy for the bookstore.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var bookStoreDefaultFn Function


getBooksIntegration := awscdk.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
bookStoreDefaultIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), bookStoreDefaultFn)

httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))

httpApi.AddRoutes(&AddRoutesOptions{
	Path: jsii.String("/books"),
	Methods: []HttpMethod{
		apigwv2.HttpMethod_GET,
	},
	Integration: getBooksIntegration,
})
httpApi.AddRoutes(&AddRoutesOptions{
	Path: jsii.String("/books"),
	Methods: []HttpMethod{
		apigwv2.HttpMethod_ANY,
	},
	Integration: bookStoreDefaultIntegration,
})
```

The URL to the endpoint can be retrieved via the `apiEndpoint` attribute. By default this URL is enabled for clients. Use `disableExecuteApiEndpoint` to disable it.

```go
httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
	DisableExecuteApiEndpoint: jsii.Boolean(true),
})
```

The `defaultIntegration` option while defining HTTP APIs lets you create a default catch-all integration that is
matched when a client reaches a route that is not explicitly defined.

```go
import "github.com/aws/aws-cdk-go/awscdk"


apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpUrlIntegration(jsii.String("DefaultIntegration"), jsii.String("https://example.com")),
})
```

The `routeSelectionExpression` option allows configuring the HTTP API to accept only `${request.method} ${request.path}`. Setting it to `true` automatically applies this value.

```go
apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &HttpApiProps{
	RouteSelectionExpression: jsii.Boolean(true),
})
```

You can configure IP address type for the API endpoint using `ipAddressType` property.
Valid values are `IPV4` (default) and `DUAL_STACK`.

```go
apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
	IpAddressType: apigwv2.IpAddressType_DUAL_STACK,
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
apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &HttpApiProps{
	CorsPreflight: &CorsPreflightOptions{
		AllowHeaders: []*string{
			jsii.String("Authorization"),
		},
		AllowMethods: []CorsHttpMethod{
			apigwv2.CorsHttpMethod_GET,
			apigwv2.CorsHttpMethod_HEAD,
			apigwv2.CorsHttpMethod_OPTIONS,
			apigwv2.CorsHttpMethod_POST,
		},
		AllowOrigins: []*string{
			jsii.String("*"),
		},
		MaxAge: awscdk.Duration_Days(jsii.Number(10)),
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
var api HttpApi


apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
	StageName: jsii.String("beta"),
	Description: jsii.String("My Stage"),
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
import "github.com/aws/aws-cdk-go/awscdk"

var handler Function


certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
domainName := "example.com"

dn := apigwv2.NewDomainName(this, jsii.String("DN"), &DomainNameProps{
	DomainName: domainName,
	Certificate: acm.Certificate_FromCertificateArn(this, jsii.String("cert"), certArn),
})
api := apigwv2.NewHttpApi(this, jsii.String("HttpProxyProdApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
	// https://${dn.domainName}/foo goes to prodApi $default stage
	DefaultDomainMapping: &DomainMappingOptions{
		DomainName: dn,
		MappingKey: jsii.String("foo"),
	},
})
```

The IP address type for the domain name can be configured by using the `ipAddressType`
property. Valid values are `IPV4` (default) and `DUAL_STACK`.

```go
import acm "github.com/aws/aws-cdk-go/awscdk"

var certificate ICertificate
var domainName string


dn := apigwv2.NewDomainName(this, jsii.String("DN"), &DomainNameProps{
	DomainName: domainName,
	Certificate: certificate,
	IpAddressType: apigwv2.IpAddressType_DUAL_STACK,
})
```

To migrate a domain endpoint from one type to another, you can add a new endpoint configuration via `addEndpoint()`
and then configure DNS records to route traffic to the new endpoint. After that, you can remove the previous endpoint configuration.
Learn more at [Migrating a custom domain name](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-migrate.html)

To associate a specific `Stage` to a custom domain mapping -

```go
var api HttpApi
var dn DomainName


api.AddStage(jsii.String("beta"), &HttpStageOptions{
	StageName: jsii.String("beta"),
	AutoDeploy: jsii.Boolean(true),
	// https://${dn.domainName}/bar goes to the beta stage
	DomainMapping: &DomainMappingOptions{
		DomainName: dn,
		MappingKey: jsii.String("bar"),
	},
})
```

The same domain name can be associated with stages across different `HttpApi` as so -

```go
import "github.com/aws/aws-cdk-go/awscdk"

var handler Function
var dn DomainName


apiDemo := apigwv2.NewHttpApi(this, jsii.String("DemoApi"), &HttpApiProps{
	DefaultIntegration: awscdk.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
	// https://${dn.domainName}/demo goes to apiDemo $default stage
	DefaultDomainMapping: &DomainMappingOptions{
		DomainName: dn,
		MappingKey: jsii.String("demo"),
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
var apiDemo HttpApi

demoDomainUrl := apiDemo.DefaultStage.DomainUrl
```

### Mutual TLS (mTLS)

Mutual TLS can be configured to limit access to your API based by using client certificates instead of (or as an extension of) using authorization headers.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
import acm "github.com/aws/aws-cdk-go/awscdk"
var bucket Bucket


certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
domainName := "example.com"

apigwv2.NewDomainName(this, jsii.String("DomainName"), &DomainNameProps{
	DomainName: jsii.String(DomainName),
	Certificate: acm.Certificate_FromCertificateArn(this, jsii.String("cert"), certArn),
	Mtls: &MTLSConfig{
		Bucket: *Bucket,
		Key: jsii.String("someca.pem"),
		Version: jsii.String("version"),
	},
})
```

Instructions for configuring your trust store can be found [here](https://aws.amazon.com/blogs/compute/introducing-mutual-tls-authentication-for-amazon-api-gateway/)

### Managing access to HTTP APIs

API Gateway supports multiple mechanisms for [controlling and managing access to your HTTP
API](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) through authorizers.

These authorizers can be found in the [APIGatewayV2-Authorizers](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_apigatewayv2_authorizers-readme.html) constructs library.

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
stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
})
clientErrorMetric := stage.metricClientError()
```

### VPC Link

Private integrations let HTTP APIs connect with AWS resources that are placed behind a VPC. These are usually Application
Load Balancers, Network Load Balancers or a Cloud Map service. The `VpcLink` construct enables this integration.
The following code creates a `VpcLink` to a private VPC.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import elb "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
alb := elb.NewApplicationLoadBalancer(this, jsii.String("AppLoadBalancer"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
})

vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
	Vpc: Vpc,
})

// Creating an HTTP ALB Integration:
albIntegration := awscdk.NewHttpAlbIntegration(jsii.String("ALBIntegration"), alb.Listeners[jsii.Number(0)], &HttpAlbIntegrationProps{
})
```

Any existing `VpcLink` resource can be imported into the CDK app via the `VpcLink.fromVpcLinkAttributes()`.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var vpc Vpc

awesomeLink := apigwv2.VpcLink_FromVpcLinkAttributes(this, jsii.String("awesome-vpc-link"), &VpcLinkAttributes{
	VpcLinkId: jsii.String("us-east-1_oiuR12Abd"),
	Vpc: Vpc,
})
```

### Private Integration

Private integrations enable integrating an HTTP API route with private resources in a VPC, such as Application Load Balancers or
Amazon ECS container-based applications.  Using private integrations, resources in a VPC can be exposed for access by
clients outside of the VPC.

These integrations can be found in the [aws-apigatewayv2-integrations](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_apigatewayv2_integrations-readme.html) constructs library.

### Generating ARN for Execute API

The arnForExecuteApi function in AWS CDK is designed to generate Amazon Resource Names (ARNs) for Execute API operations. This is particularly useful when you need to create ARNs dynamically based on different parameters like HTTP method, API path, and stage.

```go
api := apigwv2.NewHttpApi(this, jsii.String("my-api"))
arn := api.arnForExecuteApi(jsii.String("GET"), jsii.String("/myApiPath"), jsii.String("dev"))
```

* Ensure that the path parameter, if provided, starts with '/'.
* The 'ANY' method can be used for matching any HTTP methods not explicitly defined.
* The function gracefully handles undefined parameters by using wildcards, making it flexible for various API configurations.

## Access Logging

You can turn on logging to write logs to CloudWatch Logs.
Read more at [Configure logging for HTTP APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-logging.html)

```go
import logs "github.com/aws/aws-cdk-go/awscdk"

var api HttpApi
var logGroup LogGroup


stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
	AccessLogSettings: map[string]IAccessLogDestination{
		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
	},
})
```

The following code will generate the access log in the [CLF format](https://en.wikipedia.org/wiki/Common_Log_Format).

```go
import apigw "github.com/aws/aws-cdk-go/awscdk"
import logs "github.com/aws/aws-cdk-go/awscdk"

var api HttpApi
var logGroup LogGroup


stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
	AccessLogSettings: map[string]interface{}{
		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
		"format": apigw.AccessLogFormat_clf(),
	},
})
```

You can also configure your own access log format by using the `AccessLogFormat.custom()` API.
`AccessLogField` provides commonly used fields. The following code configures access log to contain.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import logs "github.com/aws/aws-cdk-go/awscdk"

var api HttpApi
var logGroup LogGroup


stage := apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
	AccessLogSettings: map[string]interface{}{
		"destination": apigwv2.NewLogGroupLogDestination(logGroup),
		"format": apigw.AccessLogFormat_custom(
		fmt.Sprintf("%v %v %v\n      %v %v", apigw.AccessLogField_contextRequestId(), apigw.AccessLogField_contextErrorMessage(), apigw.AccessLogField_contextErrorMessageString(), apigw.AccessLogField_contextAuthorizerError(), apigw.AccessLogField_contextAuthorizerIntegrationStatus())),
	},
})
```

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
import "github.com/aws/aws-cdk-go/awscdk"

var connectHandler Function
var disconnectHandler Function
var defaultHandler Function


webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
	ConnectRouteOptions: &WebSocketRouteOptions{
		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("ConnectIntegration"), connectHandler),
	},
	DisconnectRouteOptions: &WebSocketRouteOptions{
		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("DisconnectIntegration"), disconnectHandler),
	},
	DefaultRouteOptions: &WebSocketRouteOptions{
		Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("DefaultIntegration"), defaultHandler),
	},
})

apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
	WebSocketApi: WebSocketApi,
	StageName: jsii.String("dev"),
	Description: jsii.String("My Stage"),
	AutoDeploy: jsii.Boolean(true),
})
```

To retrieve a websocket URL and a callback URL:

```go
var webSocketStage WebSocketStage


webSocketURL := webSocketStage.url
// wss://${this.api.apiId}.execute-api.${s.region}.${s.urlSuffix}/${urlPath}
callbackURL := webSocketStage.callbackUrl
```

To add any other route:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var messageHandler Function

webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
})
```

To add a route that can return a result:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var messageHandler Function

webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
	ReturnResponse: jsii.Boolean(true),
})
```

To import an existing WebSocketApi:

```go
webSocketApi := apigwv2.WebSocketApi_FromWebSocketApiAttributes(this, jsii.String("mywsapi"), &WebSocketApiAttributes{
	WebSocketId: jsii.String("api-1234"),
})
```

To generate an ARN for Execute API:

```go
api := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
arn := api.ArnForExecuteApiV2(jsii.String("$connect"), jsii.String("dev"))
```

For a detailed explanation of this function, including usage and examples, please refer to the [Generating ARN for Execute API](#generating-arn-for-execute-api) section under HTTP API.

To disable schema validation, set `disableSchemaValidation` to true.

```go
apigwv2.NewWebSocketApi(this, jsii.String("api"), &WebSocketApiProps{
	DisableSchemaValidation: jsii.Boolean(true),
})
```

You can configure IP address type for the API endpoint using `ipAddressType` property.
Valid values are `IPV4` (default) and `DUAL_STACK`.

```go
apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &WebSocketApiProps{
	IpAddressType: apigwv2.IpAddressType_DUAL_STACK,
})
```

### Manage Connections Permission

Grant permission to use API Gateway Management API of a WebSocket API by calling the `grantManageConnections` API.
You can use Management API to send a callback message to a connected client, get connection information, or disconnect the client. Learn more at [Use @connections commands in your backend service](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).

```go
var fn Function


webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
stage := apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
	WebSocketApi: WebSocketApi,
	StageName: jsii.String("dev"),
})
// per stage permission
stage.GrantManagementApiAccess(fn)
// for all the stages permission
webSocketApi.GrantManageConnections(fn)
```

### Managing access to WebSocket APIs

API Gateway supports multiple mechanisms for [controlling and managing access to a WebSocket API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) through authorizers.

These authorizers can be found in the [APIGatewayV2-Authorizers](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_apigatewayv2_authorizers-readme.html) constructs library.

### API Keys

Websocket APIs also support usage of API Keys. An API Key is a key that is used to grant access to an API. These are useful for controlling and tracking access to an API, when used together with [usage plans](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html). These together allow you to configure controls around API access such as quotas and throttling, along with per-API Key metrics on usage.

To require an API Key when accessing the Websocket API:

```go
webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
	ApiKeySelectionExpression: apigwv2.WebSocketApiKeySelectionExpression_HEADER_X_API_KEY(),
})
```

## Common Config

Common config for both HTTP API and WebSocket API

### Route Settings

Represents a collection of route settings.

```go
var api HttpApi


apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
	HttpApi: api,
	Throttle: &ThrottleSettings{
		RateLimit: jsii.Number(1000),
		BurstLimit: jsii.Number(1000),
	},
	DetailedMetricsEnabled: jsii.Boolean(true),
})
```

## Usage Plan and API Keys

A usage plan specifies who can access one or more deployed WebSocket API stages, and the rate at which they can be accessed. The plan uses API keys to
identify API clients and meters access to the associated API stages for each key. Usage plans also allow configuring throttling limits and quota limits that are
enforced on individual client API keys.

The following example shows how to create and associate a usage plan and an API key for WebSocket APIs:

```go
apiKey := apigwv2.NewApiKey(this, jsii.String("ApiKey"))

usagePlan := apigwv2.NewUsagePlan(this, jsii.String("UsagePlan"), &UsagePlanProps{
	UsagePlanName: jsii.String("WebSocketUsagePlan"),
	Throttle: &ThrottleSettings{
		RateLimit: jsii.Number(10),
		BurstLimit: jsii.Number(2),
	},
})

usagePlan.addApiKey(apiKey)
```

To associate a plan to a given WebSocketAPI stage:

```go
api := apigwv2.NewWebSocketApi(this, jsii.String("my-api"))
stage := apigwv2.NewWebSocketStage(this, jsii.String("my-stage"), &WebSocketStageProps{
	WebSocketApi: api,
	StageName: jsii.String("dev"),
})

usagePlan := apigwv2.NewUsagePlan(this, jsii.String("my-usage-plan"), &UsagePlanProps{
	UsagePlanName: jsii.String("Basic"),
})

usagePlan.AddApiStage(&UsagePlanPerApiStage{
	Api: api,
	Stage: stage,
})
```

Existing usage plans can be imported into a CDK app using its id.

```go
usagePlan := apigwv2.UsagePlan_FromUsagePlanId(this, jsii.String("imported-usage-plan"), jsii.String("<usage-plan-id>"))
```

The name and value of the API Key can be specified at creation; if not provided, a name and a value will be automatically generated by API Gateway.

```go
// Auto-generated name and value
autoKey := apigwv2.NewApiKey(this, jsii.String("AutoKey"))

// Explicit name and value
explicitKey := apigwv2.NewApiKey(this, jsii.String("ExplicitKey"), &ApiKeyProps{
	ApiKeyName: jsii.String("MyWebSocketApiKey"),
	Value: jsii.String("MyApiKeyThatIsAtLeast20Characters"),
})
```

Existing API keys can also be imported into a CDK app using its id.

```go
importedKey := apigwv2.ApiKey_FromApiKeyId(this, jsii.String("imported-key"), jsii.String("<api-key-id>"))
```

The "grant" methods can be used to give prepackaged sets of permissions to other resources. The
following code provides read permission to an API key.

```go
import iam "github.com/aws/aws-cdk-go/awscdk"


user := iam.NewUser(this, jsii.String("User"))
apiKey := apigwv2.NewApiKey(this, jsii.String("ApiKey"), &ApiKeyProps{
	CustomerId: jsii.String("test-customer"),
})
apiKey.grantRead(user)
```

### Adding an API Key to an imported WebSocketApi

API Keys for WebSocket APIs are associated through Usage Plans, not directly to stages. When you import a WebSocketApi, you need to create a Usage Plan that references the
imported stage and then associate the API key with that Usage Plan.

```go
var webSocketApi IWebSocketApi


importedStage := apigwv2.WebSocketStage_FromWebSocketStageAttributes(this, jsii.String("imported-stage"), &WebSocketStageAttributes{
	StageName: jsii.String("myStage"),
	Api: webSocketApi,
})

apiKey := apigwv2.NewApiKey(this, jsii.String("MyApiKey"))

usagePlan := apigwv2.NewUsagePlan(this, jsii.String("MyUsagePlan"), &UsagePlanProps{
	ApiStages: []UsagePlanPerApiStage{
		&UsagePlanPerApiStage{
			Api: webSocketApi,
			Stage: importedStage,
		},
	},
})

usagePlan.addApiKey(apiKey)
```

### Multiple API Keys

It is possible to specify multiple API keys for a given Usage Plan, by calling `usagePlan.addApiKey()`.

When using multiple API keys, you may need to ensure that the CloudFormation logical ids of the API keys remain consistent across deployments. You can set the logical id as part of the `addApiKey()` method

```go
var usagePlan UsagePlan
var apiKey ApiKey


usagePlan.addApiKey(apiKey, &AddApiKeyOptions{
	OverrideLogicalId: jsii.String("MyCustomLogicalId"),
})
```

### Rate Limited API Key

In scenarios where you need to create a single api key and configure rate limiting for it, you can use `RateLimitedApiKey`.
This construct lets you specify rate limiting properties which should be applied only to the api key being created.
The API key created has the specified rate limits, such as quota and throttles, applied.

The following example shows how to use a rate limited api key :

```go
var api WebSocketApi
var stage WebSocketStage


key := apigwv2.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
	CustomerId: jsii.String("test-customer"),
	ApiStages: []UsagePlanPerApiStage{
		&UsagePlanPerApiStage{
			Api: api,
			Stage: stage,
		},
	},
	Quota: &QuotaSettings{
		Limit: jsii.Number(10000),
		Period: apigwv2.Period_MONTH,
	},
	Throttle: &ThrottleSettings{
		RateLimit: jsii.Number(100),
		BurstLimit: jsii.Number(200),
	},
})
```

# Amazon API Gateway Construct Library

Amazon API Gateway is a fully managed service that makes it easy for developers
to publish, maintain, monitor, and secure APIs at any scale. Create an API to
access data, business logic, or functionality from your back-end services, such
as applications running on Amazon Elastic Compute Cloud (Amazon EC2), code
running on AWS Lambda, or any web application.

## Table of Contents

* [Amazon API Gateway Construct Library](#amazon-api-gateway-construct-library)

  * [Table of Contents](#table-of-contents)
  * [Defining APIs](#defining-apis)
  * [AWS Lambda-backed APIs](#aws-lambda-backed-apis)
  * [AWS StepFunctions backed APIs](#aws-stepfunctions-backed-apis)

    * [Breaking up Methods and Resources across Stacks](#breaking-up-methods-and-resources-across-stacks)
  * [Integration Targets](#integration-targets)
  * [Usage Plan & API Keys](#usage-plan--api-keys)

    * [Adding an API Key to an imported RestApi](#adding-an-api-key-to-an-imported-restapi)
    * [⚠️ Multiple API Keys](#%EF%B8%8F-multiple-api-keys)
    * [Rate Limited API Key](#rate-limited-api-key)
  * [Working with models](#working-with-models)
  * [Default Integration and Method Options](#default-integration-and-method-options)
  * [Proxy Routes](#proxy-routes)
  * [Authorizers](#authorizers)

    * [IAM-based authorizer](#iam-based-authorizer)
    * [Lambda-based token authorizer](#lambda-based-token-authorizer)
    * [Lambda-based request authorizer](#lambda-based-request-authorizer)
    * [Cognito User Pools authorizer](#cognito-user-pools-authorizer)
  * [Mutual TLS (mTLS)](#mutual-tls-mtls)
  * [Deployments](#deployments)

    * [Deep dive: Invalidation of deployments](#deep-dive-invalidation-of-deployments)
  * [Custom Domains](#custom-domains)

    * [Custom Domains with multi-level api mapping](#custom-domains-with-multi-level-api-mapping)
  * [Access Logging](#access-logging)
  * [Cross Origin Resource Sharing (CORS)](#cross-origin-resource-sharing-cors)
  * [Endpoint Configuration](#endpoint-configuration)
  * [Private Integrations](#private-integrations)
  * [Gateway response](#gateway-response)
  * [OpenAPI Definition](#openapi-definition)

    * [Endpoint configuration](#endpoint-configuration-1)
  * [Metrics](#metrics)
  * [APIGateway v2](#apigateway-v2)

## Defining APIs

APIs are defined as a hierarchy of resources and methods. `addResource` and
`addMethod` can be used to build this hierarchy. The root resource is
`api.root`.

For example, the following code defines an API that includes the following HTTP
endpoints: `ANY /`, `GET /books`, `POST /books`, `GET /books/{book_id}`, `DELETE /books/{book_id}`.

```go
api := apigateway.NewRestApi(this, jsii.String("books-api"))

api.Root.AddMethod(jsii.String("ANY"))

books := api.Root.AddResource(jsii.String("books"))
books.AddMethod(jsii.String("GET"))
books.AddMethod(jsii.String("POST"))

book := books.AddResource(jsii.String("{book_id}"))
book.AddMethod(jsii.String("GET"))
book.AddMethod(jsii.String("DELETE"))
```

To give an IAM User or Role permission to invoke a method, use `grantExecute`:

```go
var api restApi
var user user


method := api.Root.AddResource(jsii.String("books")).AddMethod(jsii.String("GET"))
method.GrantExecute(user)
```

## AWS Lambda-backed APIs

A very common practice is to use Amazon API Gateway with AWS Lambda as the
backend integration. The `LambdaRestApi` construct makes it easy:

The following code defines a REST API that routes all requests to the
specified AWS Lambda function:

```go
var backend function

apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &LambdaRestApiProps{
	Handler: backend,
})
```

You can also supply `proxy: false`, in which case you will have to explicitly
define the API model:

```go
var backend function

api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &LambdaRestApiProps{
	Handler: backend,
	Proxy: jsii.Boolean(false),
})

items := api.Root.AddResource(jsii.String("items"))
items.AddMethod(jsii.String("GET")) // GET /items
items.AddMethod(jsii.String("POST")) // POST /items

item := items.AddResource(jsii.String("{item}"))
item.AddMethod(jsii.String("GET")) // GET /items/{item}

// the default integration for methods is "handler", but one can
// customize this behavior per method or even a sub path.
item.AddMethod(jsii.String("DELETE"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")))
```

Additionally, `integrationOptions` can be supplied to explicitly define
options of the Lambda integration:

```go
var backend function


api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &LambdaRestApiProps{
	Handler: backend,
	IntegrationOptions: &LambdaIntegrationOptions{
		AllowTestInvoke: jsii.Boolean(false),
		Timeout: awscdk.Duration_Seconds(jsii.Number(1)),
	},
})
```

## AWS StepFunctions backed APIs

You can use Amazon API Gateway with AWS Step Functions as the backend integration, specifically Synchronous Express Workflows.

The `StepFunctionsRestApi` only supports integration with Synchronous Express state machine. The `StepFunctionsRestApi` construct makes this easy by setting up input, output and error mapping.

The construct sets up an API endpoint and maps the `ANY` HTTP method and any calls to the API endpoint starts an express workflow execution for the underlying state machine.

Invoking the endpoint with any HTTP method (`GET`, `POST`, `PUT`, `DELETE`, ...) in the example below will send the request to the state machine as a new execution. On success, an HTTP code `200` is returned with the execution output as the Response Body.

If the execution fails, an HTTP `500` response is returned with the `error` and `cause` from the execution output as the Response Body. If the request is invalid (ex. bad execution input) HTTP code `400` is returned.

To disable default response models generation use the `useDefaultMethodResponses` property:

```go
var machine iStateMachine


apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &StepFunctionsRestApiProps{
	StateMachine: machine,
	UseDefaultMethodResponses: jsii.Boolean(false),
})
```

The response from the invocation contains only the `output` field from the
[StartSyncExecution](https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartSyncExecution.html#API_StartSyncExecution_ResponseSyntax) API.
In case of failures, the fields `error` and `cause` are returned as part of the response.
Other metadata such as billing details, AWS account ID and resource ARNs are not returned in the API response.

By default, a `prod` stage is provisioned.

In order to reduce the payload size sent to AWS Step Functions, `headers` are not forwarded to the Step Functions execution input. It is possible to choose whether `headers`,  `requestContext`, `path`, `querystring`, and `authorizer` are included or not. By default, `headers` are excluded in all requests.

More details about AWS Step Functions payload limit can be found at https://docs.aws.amazon.com/step-functions/latest/dg/limits-overview.html#service-limits-task-executions.

The following code defines a REST API that routes all requests to the specified AWS StepFunctions state machine:

```go
stateMachineDefinition := stepfunctions.NewPass(this, jsii.String("PassState"))

stateMachine := stepfunctions.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	Definition: stateMachineDefinition,
	StateMachineType: stepfunctions.StateMachineType_EXPRESS,
})

apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &StepFunctionsRestApiProps{
	Deploy: jsii.Boolean(true),
	StateMachine: stateMachine,
})
```

When the REST API endpoint configuration above is invoked using POST, as follows -

```bash
curl -X POST -d '{ "customerId": 1 }' https://example.com/
```

AWS Step Functions will receive the request body in its input as follows:

```json
{
  "body": {
    "customerId": 1
  },
  "path": "/",
  "querystring": {}
}
```

When the endpoint is invoked at path '/users/5' using the HTTP GET method as below:

```bash
curl -X GET https://example.com/users/5?foo=bar
```

AWS Step Functions will receive the following execution input:

```json
{
  "body": {},
  "path": {
     "users": "5"
  },
  "querystring": {
    "foo": "bar"
  }
}
```

Additional information around the request such as the request context, authorizer context, and headers can be included as part of the input
forwarded to the state machine. The following example enables headers to be included in the input but not query string.

```go
apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &StepFunctionsRestApiProps{
	StateMachine: machine,
	Headers: jsii.Boolean(true),
	Path: jsii.Boolean(false),
	Querystring: jsii.Boolean(false),
	Authorizer: jsii.Boolean(false),
	RequestContext: &RequestContext{
		Caller: jsii.Boolean(true),
		User: jsii.Boolean(true),
	},
})
```

In such a case, when the endpoint is invoked as below:

```bash
curl -X GET https://example.com/
```

AWS Step Functions will receive the following execution input:

```json
{
  "headers": {
    "Accept": "...",
    "CloudFront-Forwarded-Proto": "...",
  },
  "requestContext": {
     "accountId": "...",
     "apiKey": "...",
  },
  "body": {}
}
```

### Breaking up Methods and Resources across Stacks

It is fairly common for REST APIs with a large number of Resources and Methods to hit the [CloudFormation
limit](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html) of 500 resources per
stack.

To help with this, Resources and Methods for the same REST API can be re-organized across multiple stacks. A common
way to do this is to have a stack per Resource or groups of Resources, but this is not the only possible way.
The following example uses sets up two Resources '/pets' and '/books' in separate stacks using nested stacks:

```go
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

/**
 * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
 *
 * The root stack 'RootStack' first defines a RestApi.
 * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
 * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
 *
 * To verify this worked, go to the APIGateway
 */

type rootStack struct {
	stack
}

func newRootStack(scope construct) *rootStack {
	this := &rootStack{}
	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))

	restApi := awscdk.NewRestApi(this, jsii.String("RestApi"), &RestApiProps{
		CloudWatchRole: jsii.Boolean(true),
		Deploy: jsii.Boolean(false),
	})
	restApi.Root.AddMethod(jsii.String("ANY"))

	petsStack := NewPetsStack(this, &resourceNestedStackProps{
		restApiId: restApi.RestApiId,
		rootResourceId: restApi.RestApiRootResourceId,
	})
	booksStack := NewBooksStack(this, &resourceNestedStackProps{
		restApiId: restApi.*RestApiId,
		rootResourceId: restApi.*RestApiRootResourceId,
	})
	NewDeployStack(this, &deployStackProps{
		restApiId: restApi.*RestApiId,
		methods: petsStack.methods.concat(booksStack.methods),
	})

	awscdk.NewCfnOutput(this, jsii.String("PetsURL"), &CfnOutputProps{
		Value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.*RestApiId, this.Region),
	})

	awscdk.NewCfnOutput(this, jsii.String("BooksURL"), &CfnOutputProps{
		Value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.*RestApiId, this.*Region),
	})
	return this
}

type resourceNestedStackProps struct {
	nestedStackProps
	restApiId *string
	rootResourceId *string
}

type petsStack struct {
	nestedStack
	methods []method
}

func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
	this := &petsStack{}
	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)

	api := awscdk.RestApi_FromRestApiAttributes(this, jsii.String("RestApi"), &RestApiAttributes{
		RestApiId: props.restApiId,
		RootResourceId: props.rootResourceId,
	})

	method := api.Root.AddResource(jsii.String("pets")).AddMethod(jsii.String("GET"), awscdk.NewMockIntegration(&IntegrationOptions{
		IntegrationResponses: []integrationResponse{
			&integrationResponse{
				StatusCode: jsii.String("200"),
			},
		},
		PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
		RequestTemplates: map[string]*string{
			"application/json": jsii.String("{ \"statusCode\": 200 }"),
		},
	}), &MethodOptions{
		MethodResponses: []methodResponse{
			&methodResponse{
				StatusCode: jsii.String("200"),
			},
		},
	})

	this.methods.push(method)
	return this
}

type booksStack struct {
	nestedStack
	methods []*method
}

func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
	this := &booksStack{}
	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)

	api := awscdk.RestApi_FromRestApiAttributes(this, jsii.String("RestApi"), &RestApiAttributes{
		RestApiId: props.restApiId,
		RootResourceId: props.rootResourceId,
	})

	method := api.Root.AddResource(jsii.String("books")).AddMethod(jsii.String("GET"), awscdk.NewMockIntegration(&IntegrationOptions{
		IntegrationResponses: []*integrationResponse{
			&integrationResponse{
				StatusCode: jsii.String("200"),
			},
		},
		PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
		RequestTemplates: map[string]*string{
			"application/json": jsii.String("{ \"statusCode\": 200 }"),
		},
	}), &MethodOptions{
		MethodResponses: []*methodResponse{
			&methodResponse{
				StatusCode: jsii.String("200"),
			},
		},
	})

	this.methods.push(method)
	return this
}

type deployStackProps struct {
	nestedStackProps
	restApiId *string
	methods []*method
}

type deployStack struct {
	nestedStack
}

func newDeployStack(scope construct, props deployStackProps) *deployStack {
	this := &deployStack{}
	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)

	deployment := awscdk.NewDeployment(this, jsii.String("Deployment"), &DeploymentProps{
		Api: awscdk.RestApi_FromRestApiId(this, jsii.String("RestApi"), props.restApiId),
	})
	if *props.methods {
		for _, method := range *props.methods {
			deployment.Node.AddDependency(method)
		}
	}
	awscdk.NewStage(this, jsii.String("Stage"), &StageProps{
		Deployment: Deployment,
	})
	return this
}

NewRootStack(awscdk.NewApp())
```

## Integration Targets

Methods are associated with backend integrations, which are invoked when this
method is called. API Gateway supports the following integrations:

* `MockIntegration` - can be used to test APIs. This is the default
  integration if one is not specified.
* `AwsIntegration` - can be used to invoke arbitrary AWS service APIs.
* `HttpIntegration` - can be used to invoke HTTP endpoints.
* `LambdaIntegration` - can be used to invoke an AWS Lambda function.
* `SagemakerIntegration` - can be used to invoke Sagemaker Endpoints.

The following example shows how to integrate the `GET /book/{book_id}` method to
an AWS Lambda function:

```go
var getBookHandler function
var book resource


getBookIntegration := apigateway.NewLambdaIntegration(getBookHandler)
book.AddMethod(jsii.String("GET"), getBookIntegration)
```

Integration options can be optionally be specified:

```go
var getBookHandler function
var getBookIntegration lambdaIntegration


getBookIntegration := apigateway.NewLambdaIntegration(getBookHandler, &LambdaIntegrationOptions{
	ContentHandling: apigateway.ContentHandling_CONVERT_TO_TEXT,
	 // convert to base64
	CredentialsPassthrough: jsii.Boolean(true),
})
```

Method options can optionally be specified when adding methods:

```go
var book resource
var getBookIntegration lambdaIntegration


book.AddMethod(jsii.String("GET"), getBookIntegration, &MethodOptions{
	AuthorizationType: apigateway.AuthorizationType_IAM,
	ApiKeyRequired: jsii.Boolean(true),
})
```

It is possible to also integrate with AWS services in a different region. The following code integrates with Amazon SQS in the
`eu-west-1` region.

```go
getMessageIntegration := apigateway.NewAwsIntegration(&AwsIntegrationProps{
	Service: jsii.String("sqs"),
	Path: jsii.String("queueName"),
	Region: jsii.String("eu-west-1"),
})
```

## Usage Plan & API Keys

A usage plan specifies who can access one or more deployed API stages and methods, and the rate at which they can be
accessed. The plan uses API keys to identify API clients and meters access to the associated API stages for each key.
Usage plans also allow configuring throttling limits and quota limits that are enforced on individual client API keys.

The following example shows how to create and associate a usage plan and an API key:

```go
var integration lambdaIntegration


api := apigateway.NewRestApi(this, jsii.String("hello-api"))

v1 := api.Root.AddResource(jsii.String("v1"))
echo := v1.AddResource(jsii.String("echo"))
echoMethod := echo.AddMethod(jsii.String("GET"), integration, &MethodOptions{
	ApiKeyRequired: jsii.Boolean(true),
})

plan := api.AddUsagePlan(jsii.String("UsagePlan"), &UsagePlanProps{
	Name: jsii.String("Easy"),
	Throttle: &ThrottleSettings{
		RateLimit: jsii.Number(10),
		BurstLimit: jsii.Number(2),
	},
})

key := api.AddApiKey(jsii.String("ApiKey"))
plan.addApiKey(key)
```

To associate a plan to a given RestAPI stage:

```go
var plan usagePlan
var api restApi
var echoMethod method


plan.AddApiStage(&UsagePlanPerApiStage{
	Stage: api.DeploymentStage,
	Throttle: []throttlingPerMethod{
		&throttlingPerMethod{
			Method: echoMethod,
			Throttle: &ThrottleSettings{
				RateLimit: jsii.Number(10),
				BurstLimit: jsii.Number(2),
			},
		},
	},
})
```

Existing usage plans can be imported into a CDK app using its id.

```go
importedUsagePlan := apigateway.UsagePlan_FromUsagePlanId(this, jsii.String("imported-usage-plan"), jsii.String("<usage-plan-key-id>"))
```

The name and value of the API Key can be specified at creation; if not
provided, a name and value will be automatically generated by API Gateway.

```go
var api restApi

key := api.AddApiKey(jsii.String("ApiKey"), &ApiKeyOptions{
	ApiKeyName: jsii.String("myApiKey1"),
	Value: jsii.String("MyApiKeyThatIsAtLeast20Characters"),
})
```

Existing API keys can also be imported into a CDK app using its id.

```go
importedKey := apigateway.ApiKey_FromApiKeyId(this, jsii.String("imported-key"), jsii.String("<api-key-id>"))
```

The "grant" methods can be used to give prepackaged sets of permissions to other resources. The
following code provides read permission to an API key.

```go
var importedKey apiKey
var lambdaFn function

importedKey.grantRead(lambdaFn)
```

### Adding an API Key to an imported RestApi

API Keys are added to ApiGateway Stages, not to the API itself. When you import a RestApi
it does not have any information on the Stages that may be associated with it. Since adding an API
Key requires a stage, you should instead add the Api Key to the imported Stage.

```go
var restApi iRestApi

importedStage := apigateway.Stage_FromStageAttributes(this, jsii.String("imported-stage"), &StageAttributes{
	StageName: jsii.String("myStageName"),
	RestApi: RestApi,
})

importedStage.AddApiKey(jsii.String("MyApiKey"))
```

### ⚠️ Multiple API Keys

It is possible to specify multiple API keys for a given Usage Plan, by calling `usagePlan.addApiKey()`.

When using multiple API keys, a past bug of the CDK prevents API key associations to a Usage Plan to be deleted.
If the CDK app had the [feature flag](https://docs.aws.amazon.com/cdk/latest/guide/featureflags.html) - `@aws-cdk/aws-apigateway:usagePlanKeyOrderInsensitiveId` - enabled when the API
keys were created, then the app will not be affected by this bug.

If this is not the case, you will need to ensure that the CloudFormation [logical ids](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html) of the API keys that are not
being deleted remain unchanged.
Make note of the logical ids of these API keys before removing any, and set it as part of the `addApiKey()` method:

```go
var usageplan usagePlan
var apiKey apiKey


usageplan.addApiKey(apiKey, &AddApiKeyOptions{
	OverrideLogicalId: jsii.String("..."),
})
```

### Rate Limited API Key

In scenarios where you need to create a single api key and configure rate limiting for it, you can use `RateLimitedApiKey`.
This construct lets you specify rate limiting properties which should be applied only to the api key being created.
The API key created has the specified rate limits, such as quota and throttles, applied.

The following example shows how to use a rate limited api key :

```go
var api restApi


key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
	CustomerId: jsii.String("hello-customer"),
	Stages: []iStage{
		api.DeploymentStage,
	},
	Quota: &QuotaSettings{
		Limit: jsii.Number(10000),
		Period: apigateway.Period_MONTH,
	},
})
```

## Working with models

When you work with Lambda integrations that are not Proxy integrations, you
have to define your models and mappings for the request, response, and integration.

```go
hello := lambda.NewFunction(this, jsii.String("hello"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("hello.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda")),
})

api := apigateway.NewRestApi(this, jsii.String("hello-api"), &RestApiProps{
})
resource := api.Root.AddResource(jsii.String("v1"))
```

You can define more parameters on the integration to tune the behavior of API Gateway

```go
var hello function


integration := apigateway.NewLambdaIntegration(hello, &LambdaIntegrationOptions{
	Proxy: jsii.Boolean(false),
	RequestParameters: map[string]*string{
		// You can define mapping parameters from your method to your integration
		// - Destination parameters (the key) are the integration parameters (used in mappings)
		// - Source parameters (the value) are the source request parameters or expressions
		// @see: https://docs.aws.amazon.com/apigateway/latest/developerguide/request-response-data-mappings.html
		"integration.request.querystring.who": jsii.String("method.request.querystring.who"),
	},
	AllowTestInvoke: jsii.Boolean(true),
	RequestTemplates: map[string]*string{
		// You can define a mapping that will build a payload for your integration, based
		//  on the integration parameters that you have specified
		// Check: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
		"application/json": JSON.stringify(map[string]*string{
			"action": jsii.String("sayHello"),
			"pollId": jsii.String("$util.escapeJavaScript($input.params('who'))"),
		}),
	},
	// This parameter defines the behavior of the engine is no suitable response template is found
	PassthroughBehavior: apigateway.PassthroughBehavior_NEVER,
	IntegrationResponses: []integrationResponse{
		&integrationResponse{
			// Successful response from the Lambda function, no filter defined
			//  - the selectionPattern filter only tests the error message
			// We will set the response status code to 200
			StatusCode: jsii.String("200"),
			ResponseTemplates: map[string]*string{
				// This template takes the "message" result from the Lambda function, and embeds it in a JSON response
				// Check https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
				"application/json": JSON.stringify(map[string]*string{
					"state": jsii.String("ok"),
					"greeting": jsii.String("$util.escapeJavaScript($input.body)"),
				}),
			},
			ResponseParameters: map[string]*string{
				// We can map response parameters
				// - Destination parameters (the key) are the response parameters (used in mappings)
				// - Source parameters (the value) are the integration response parameters or expressions
				"method.response.header.Content-Type": jsii.String("'application/json'"),
				"method.response.header.Access-Control-Allow-Origin": jsii.String("'*'"),
				"method.response.header.Access-Control-Allow-Credentials": jsii.String("'true'"),
			},
		},
		&integrationResponse{
			// For errors, we check if the error message is not empty, get the error data
			SelectionPattern: jsii.String("(\n|.)+"),
			// We will set the response status code to 200
			StatusCode: jsii.String("400"),
			ResponseTemplates: map[string]*string{
				"application/json": JSON.stringify(map[string]*string{
					"state": jsii.String("error"),
					"message": jsii.String("$util.escapeJavaScript($input.path('$.errorMessage'))"),
				}),
			},
			ResponseParameters: map[string]*string{
				"method.response.header.Content-Type": jsii.String("'application/json'"),
				"method.response.header.Access-Control-Allow-Origin": jsii.String("'*'"),
				"method.response.header.Access-Control-Allow-Credentials": jsii.String("'true'"),
			},
		},
	},
})
```

You can define models for your responses (and requests)

```go
var api restApi


// We define the JSON Schema for the transformed valid response
responseModel := api.AddModel(jsii.String("ResponseModel"), &ModelOptions{
	ContentType: jsii.String("application/json"),
	ModelName: jsii.String("ResponseModel"),
	Schema: &JsonSchema{
		Schema: apigateway.JsonSchemaVersion_DRAFT4,
		Title: jsii.String("pollResponse"),
		Type: apigateway.JsonSchemaType_OBJECT,
		Properties: map[string]jsonSchema{
			"state": &jsonSchema{
				"type": apigateway.JsonSchemaType_STRING,
			},
			"greeting": &jsonSchema{
				"type": apigateway.JsonSchemaType_STRING,
			},
		},
	},
})

// We define the JSON Schema for the transformed error response
errorResponseModel := api.AddModel(jsii.String("ErrorResponseModel"), &ModelOptions{
	ContentType: jsii.String("application/json"),
	ModelName: jsii.String("ErrorResponseModel"),
	Schema: &jsonSchema{
		Schema: apigateway.JsonSchemaVersion_DRAFT4,
		Title: jsii.String("errorResponse"),
		Type: apigateway.JsonSchemaType_OBJECT,
		Properties: map[string]*jsonSchema{
			"state": &jsonSchema{
				"type": apigateway.JsonSchemaType_STRING,
			},
			"message": &jsonSchema{
				"type": apigateway.JsonSchemaType_STRING,
			},
		},
	},
})
```

And reference all on your method definition.

```go
var integration lambdaIntegration
var resource resource
var responseModel model
var errorResponseModel model


resource.AddMethod(jsii.String("GET"), integration, &MethodOptions{
	// We can mark the parameters as required
	RequestParameters: map[string]*bool{
		"method.request.querystring.who": jsii.Boolean(true),
	},
	// we can set request validator options like below
	RequestValidatorOptions: &RequestValidatorOptions{
		RequestValidatorName: jsii.String("test-validator"),
		ValidateRequestBody: jsii.Boolean(true),
		ValidateRequestParameters: jsii.Boolean(false),
	},
	MethodResponses: []methodResponse{
		&methodResponse{
			// Successful response from the integration
			StatusCode: jsii.String("200"),
			// Define what parameters are allowed or not
			ResponseParameters: map[string]*bool{
				"method.response.header.Content-Type": jsii.Boolean(true),
				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
			},
			// Validate the schema on the response
			ResponseModels: map[string]iModel{
				"application/json": responseModel,
			},
		},
		&methodResponse{
			// Same thing for the error responses
			StatusCode: jsii.String("400"),
			ResponseParameters: map[string]*bool{
				"method.response.header.Content-Type": jsii.Boolean(true),
				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
			},
			ResponseModels: map[string]*iModel{
				"application/json": errorResponseModel,
			},
		},
	},
})
```

Specifying `requestValidatorOptions` automatically creates the RequestValidator construct with the given options.
However, if you have your RequestValidator already initialized or imported, use the `requestValidator` option instead.

If you want to use `requestValidatorOptions` in multiple `addMethod()` calls
then you need to set the `@aws-cdk/aws-apigateway:requestValidatorUniqueId`
feature flag. When this feature flag is set, each `RequestValidator` will have a unique generated id.

> **Note** if you enable this feature flag when you have already used
> `addMethod()` with `requestValidatorOptions` the Logical Id of the resource
> will change causing the resource to be replaced.

```go
var integration lambdaIntegration
var resource resource
var responseModel model
var errorResponseModel model


resource.Node.SetContext(jsii.String("@aws-cdk/aws-apigateway:requestValidatorUniqueId"), jsii.Boolean(true))

resource.AddMethod(jsii.String("GET"), integration, &MethodOptions{
	// we can set request validator options like below
	RequestValidatorOptions: &RequestValidatorOptions{
		RequestValidatorName: jsii.String("test-validator"),
		ValidateRequestBody: jsii.Boolean(true),
		ValidateRequestParameters: jsii.Boolean(false),
	},
})

resource.AddMethod(jsii.String("POST"), integration, &MethodOptions{
	// we can set request validator options like below
	RequestValidatorOptions: &RequestValidatorOptions{
		RequestValidatorName: jsii.String("test-validator2"),
		ValidateRequestBody: jsii.Boolean(true),
		ValidateRequestParameters: jsii.Boolean(false),
	},
})
```

## Default Integration and Method Options

The `defaultIntegration` and `defaultMethodOptions` properties can be used to
configure a default integration at any resource level. These options will be
used when defining method under this resource (recursively) with undefined
integration or options.

> If not defined, the default integration is `MockIntegration`. See reference
> documentation for default method options.

The following example defines the `booksBackend` integration as a default
integration. This means that all API methods that do not explicitly define an
integration will be routed to this AWS Lambda function.

```go
var booksBackend lambdaIntegration

api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DefaultIntegration: booksBackend,
})

books := api.Root.AddResource(jsii.String("books"))
books.AddMethod(jsii.String("GET")) // integrated with `booksBackend`
books.AddMethod(jsii.String("POST")) // integrated with `booksBackend`

book := books.AddResource(jsii.String("{book_id}"))
book.AddMethod(jsii.String("GET"))
```

A Method can be configured with authorization scopes. Authorization scopes are
used in conjunction with an [authorizer that uses Amazon Cognito user
pools](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html#apigateway-enable-cognito-user-pool).
Read more about authorization scopes
[here](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes).

Authorization scopes for a Method can be configured using the `authorizationScopes` property as shown below -

```go
var books resource


books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
	AuthorizationType: apigateway.AuthorizationType_COGNITO,
	AuthorizationScopes: []*string{
		jsii.String("Scope1"),
		jsii.String("Scope2"),
	},
})
```

## Proxy Routes

The `addProxy` method can be used to install a greedy `{proxy+}` resource
on a path. By default, this also installs an `"ANY"` method:

```go
var resource resource
var handler function

proxy := resource.AddProxy(&ProxyResourceOptions{
	DefaultIntegration: apigateway.NewLambdaIntegration(handler),

	// "false" will require explicitly adding methods on the `proxy` resource
	AnyMethod: jsii.Boolean(true),
})
```

## Authorizers

API Gateway [supports several different authorization types](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-control-access-to-api.html)
that can be used for controlling access to your REST APIs.

### IAM-based authorizer

The following CDK code provides 'execute-api' permission to an IAM user, via IAM policies, for the 'GET' method on the `books` resource:

```go
var books resource
var iamUser user


getBooks := books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
	AuthorizationType: apigateway.AuthorizationType_IAM,
})

iamUser.AttachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowBooks"), &PolicyProps{
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("execute-api:Invoke"),
			},
			Effect: iam.Effect_ALLOW,
			Resources: []*string{
				getBooks.methodArn,
			},
		}),
	},
}))
```

### Lambda-based token authorizer

API Gateway also allows [lambda functions to be used as authorizers](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html).

This module provides support for token-based Lambda authorizers. When a client makes a request to an API's methods configured with such
an authorizer, API Gateway calls the Lambda authorizer, which takes the caller's identity as input and returns an IAM policy as output.
A token-based Lambda authorizer (also called a token authorizer) receives the caller's identity in a bearer token, such as
a JSON Web Token (JWT) or an OAuth token.

API Gateway interacts with the authorizer Lambda function handler by passing input and expecting the output in a specific format.
The event object that the handler is called with contains the `authorizationToken` and the `methodArn` from the request to the
API Gateway endpoint. The handler is expected to return the `principalId` (i.e. the client identifier) and a `policyDocument` stating
what the client is authorizer to perform.
See [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html) for a detailed specification on
inputs and outputs of the Lambda handler.

The following code attaches a token-based Lambda authorizer to the 'GET' Method of the Book resource:

```go
var authFn function
var books resource


auth := apigateway.NewTokenAuthorizer(this, jsii.String("booksAuthorizer"), &TokenAuthorizerProps{
	Handler: authFn,
})

books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
	Authorizer: auth,
})
```

A full working example is shown below.

```go
type myStack struct {
	stack
}

func newMyStack(scope construct, id *string) *myStack {
	this := &myStack{}
	newStack_Override(this, scope, id)

	authorizerFn := lambda.NewFunction(this, jsii.String("MyAuthorizerFunction"), &FunctionProps{
		Runtime: lambda.Runtime_NODEJS_LATEST(),
		Handler: jsii.String("index.handler"),
		Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("integ.token-authorizer.handler"))),
	})

	authorizer := awscdk.NewTokenAuthorizer(this, jsii.String("MyAuthorizer"), &TokenAuthorizerProps{
		Handler: authorizerFn,
	})

	restapi := awscdk.NewRestApi(this, jsii.String("MyRestApi"), &RestApiProps{
		CloudWatchRole: jsii.Boolean(true),
		DefaultMethodOptions: &MethodOptions{
			Authorizer: *Authorizer,
		},
		DefaultCorsPreflightOptions: &CorsOptions{
			AllowOrigins: awscdk.Cors_ALL_ORIGINS(),
		},
	})

	restapi.Root.AddMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&IntegrationOptions{
		IntegrationResponses: []integrationResponse{
			&integrationResponse{
				StatusCode: jsii.String("200"),
			},
		},
		PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
		RequestTemplates: map[string]*string{
			"application/json": jsii.String("{ \"statusCode\": 200 }"),
		},
	}), &MethodOptions{
		MethodResponses: []methodResponse{
			&methodResponse{
				StatusCode: jsii.String("200"),
			},
		},
	})
	return this
}
```

By default, the `TokenAuthorizer` looks for the authorization token in the request header with the key 'Authorization'. This can,
however, be modified by changing the `identitySource` property.

Authorizers can also be passed via the `defaultMethodOptions` property within the `RestApi` construct or the `Method` construct. Unless
explicitly overridden, the specified defaults will be applied across all `Method`s across the `RestApi` or across all `Resource`s,
depending on where the defaults were specified.

### Lambda-based request authorizer

This module provides support for request-based Lambda authorizers. When a client makes a request to an API's methods configured with such
an authorizer, API Gateway calls the Lambda authorizer, which takes specified parts of the request, known as identity sources,
as input and returns an IAM policy as output. A request-based Lambda authorizer (also called a request authorizer) receives
the identity sources in a series of values pulled from the request, from the headers, stage variables, query strings, and the context.

API Gateway interacts with the authorizer Lambda function handler by passing input and expecting the output in a specific format.
The event object that the handler is called with contains the body of the request and the `methodArn` from the request to the
API Gateway endpoint. The handler is expected to return the `principalId` (i.e. the client identifier) and a `policyDocument` stating
what the client is authorizer to perform.
See [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html) for a detailed specification on
inputs and outputs of the Lambda handler.

The following code attaches a request-based Lambda authorizer to the 'GET' Method of the Book resource:

```go
var authFn function
var books resource


auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &RequestAuthorizerProps{
	Handler: authFn,
	IdentitySources: []*string{
		apigateway.IdentitySource_Header(jsii.String("Authorization")),
	},
})

books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
	Authorizer: auth,
})
```

A full working example is shown below.

```go
import path "github.com/aws-samples/dummy/path"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

// Against the RestApi endpoint from the stack output, run
// `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
// `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>?allow=yes` should return 403
// `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>?allow=yes` should return 200

app := awscdk.NewApp()
stack := awscdk.NewStack(app, jsii.String("RequestAuthorizerInteg"))

authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
})

restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"), &RestApiProps{
	CloudWatchRole: jsii.Boolean(true),
})

authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &RequestAuthorizerProps{
	Handler: authorizerFn,
	IdentitySources: []*string{
		awscdk.IdentitySource_Header(jsii.String("Authorization")),
		awscdk.IdentitySource_QueryString(jsii.String("allow")),
	},
})

secondAuthorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MySecondAuthorizer"), &RequestAuthorizerProps{
	Handler: authorizerFn,
	IdentitySources: []*string{
		awscdk.IdentitySource_*Header(jsii.String("Authorization")),
		awscdk.IdentitySource_*QueryString(jsii.String("allow")),
	},
})

restapi.Root.AddMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&IntegrationOptions{
	IntegrationResponses: []integrationResponse{
		&integrationResponse{
			StatusCode: jsii.String("200"),
		},
	},
	PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
	RequestTemplates: map[string]*string{
		"application/json": jsii.String("{ \"statusCode\": 200 }"),
	},
}), &MethodOptions{
	MethodResponses: []methodResponse{
		&methodResponse{
			StatusCode: jsii.String("200"),
		},
	},
	Authorizer: Authorizer,
})

restapi.Root.ResourceForPath(jsii.String("auth")).AddMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&IntegrationOptions{
	IntegrationResponses: []*integrationResponse{
		&integrationResponse{
			StatusCode: jsii.String("200"),
		},
	},
	PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
	RequestTemplates: map[string]*string{
		"application/json": jsii.String("{ \"statusCode\": 200 }"),
	},
}), &MethodOptions{
	MethodResponses: []*methodResponse{
		&methodResponse{
			StatusCode: jsii.String("200"),
		},
	},
	Authorizer: secondAuthorizer,
})
```

By default, the `RequestAuthorizer` does not pass any kind of information from the request. This can,
however, be modified by changing the `identitySource` property, and is required when specifying a value for caching.

Authorizers can also be passed via the `defaultMethodOptions` property within the `RestApi` construct or the `Method` construct. Unless
explicitly overridden, the specified defaults will be applied across all `Method`s across the `RestApi` or across all `Resource`s,
depending on where the defaults were specified.

### Cognito User Pools authorizer

API Gateway also allows [Amazon Cognito user pools as authorizer](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)

The following snippet configures a Cognito user pool as an authorizer:

```go
var books resource
userPool := cognito.NewUserPool(this, jsii.String("UserPool"))

auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &CognitoUserPoolsAuthorizerProps{
	CognitoUserPools: []iUserPool{
		userPool,
	},
})
books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
	Authorizer: auth,
	AuthorizationType: apigateway.AuthorizationType_COGNITO,
})
```

## Mutual TLS (mTLS)

Mutual TLS can be configured to limit access to your API based by using client certificates instead of (or as an extension of) using authorization headers.

```go
var acm interface{}


apigateway.NewDomainName(this, jsii.String("domain-name"), &DomainNameProps{
	DomainName: jsii.String("example.com"),
	Certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
	Mtls: &MTLSConfig{
		Bucket: s3.NewBucket(this, jsii.String("bucket")),
		Key: jsii.String("truststore.pem"),
		Version: jsii.String("version"),
	},
})
```

Instructions for configuring your trust store can be found [here](https://aws.amazon.com/blogs/compute/introducing-mutual-tls-authentication-for-amazon-api-gateway/).

## Deployments

By default, the `RestApi` construct will automatically create an API Gateway
[Deployment](https://docs.aws.amazon.com/apigateway/api-reference/resource/deployment/) and a "prod" [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/) which represent the API configuration you
defined in your CDK app. This means that when you deploy your app, your API will
be have open access from the internet via the stage URL.

The URL of your API can be obtained from the attribute `restApi.url`, and is
also exported as an `Output` from your stack, so it's printed when you `cdk deploy` your app:

```console
$ cdk deploy
...
books.booksapiEndpointE230E8D5 = https://6lyktd4lpk.execute-api.us-east-1.amazonaws.com/prod/
```

To disable this behavior, you can set `{ deploy: false }` when creating your
API. This means that the API will not be deployed and a stage will not be
created for it. You will need to manually define a `apigateway.Deployment` and
`apigateway.Stage` resources.

Use the `deployOptions` property to customize the deployment options of your
API.

The following example will configure API Gateway to emit logs and data traces to
AWS CloudWatch for all API calls:

> Note: whether or not this is enabled or disabled by default is controlled by the
> `@aws-cdk/aws-apigateway:disableCloudWatchRole` feature flag. When this feature flag
> is set to `false` the default behavior will set `cloudWatchRole=true`

This is controlled via the `@aws-cdk/aws-apigateway:disableCloudWatchRole` feature flag and
is disabled by default. When enabled (or `@aws-cdk/aws-apigateway:disableCloudWatchRole=false`),
an IAM role will be created and associated with API Gateway to allow it to write logs and metrics to AWS CloudWatch.

```go
api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	CloudWatchRole: jsii.Boolean(true),
	DeployOptions: &StageOptions{
		LoggingLevel: apigateway.MethodLoggingLevel_INFO,
		DataTraceEnabled: jsii.Boolean(true),
	},
})
```

> Note: there can only be a single apigateway.CfnAccount per AWS environment
> so if you create multiple `RestApi`s with `cloudWatchRole=true` each new `RestApi`
> will overwrite the `CfnAccount`. It is recommended to set `cloudWatchRole=false`
> (the default behavior if `@aws-cdk/aws-apigateway:disableCloudWatchRole` is enabled)
> and only create a single CloudWatch role and account per environment.

You can specify the CloudWatch Role and Account sub-resources removal policy with the
`cloudWatchRoleRemovalPolicy` property, which defaults to `RemovalPolicy.RETAIN`.
This option requires `cloudWatchRole` to be enabled.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"


api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	CloudWatchRole: jsii.Boolean(true),
	CloudWatchRoleRemovalPolicy: cdk.RemovalPolicy_DESTROY,
})
```

### Deep dive: Invalidation of deployments

API Gateway deployments are an immutable snapshot of the API. This means that we
want to automatically create a new deployment resource every time the API model
defined in our CDK app changes.

In order to achieve that, the AWS CloudFormation logical ID of the
`AWS::ApiGateway::Deployment` resource is dynamically calculated by hashing the
API configuration (resources, methods). This means that when the configuration
changes (i.e. a resource or method are added, configuration is changed), a new
logical ID will be assigned to the deployment resource. This will cause
CloudFormation to create a new deployment resource.

By default, old deployments are *deleted*. You can set `retainDeployments: true`
to allow users revert the stage to an old deployment manually.

In order to also create a new deployment when changes are made to any authorizer attached to the API,
the `@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId` [feature flag](https://docs.aws.amazon.com/cdk/v2/guide/featureflags.html) can be enabled. This can be set
in the `cdk.json` file.

```json
{
  "context": {
    "@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId": true
  }
}
```

## Custom Domains

To associate an API with a custom domain, use the `domainName` configuration when
you define your API:

```go
var acmCertificateForExampleCom interface{}


api := apigateway.NewRestApi(this, jsii.String("MyDomain"), &RestApiProps{
	DomainName: &DomainNameOptions{
		DomainName: jsii.String("example.com"),
		Certificate: acmCertificateForExampleCom,
	},
})
```

This will define a `DomainName` resource for you, along with a `BasePathMapping`
from the root of the domain to the deployment stage of the API. This is a common
set up.

To route domain traffic to an API Gateway API, use Amazon Route 53 to create an
alias record. An alias record is a Route 53 extension to DNS. It's similar to a
CNAME record, but you can create an alias record both for the root domain, such
as `example.com`, and for subdomains, such as `www.example.com`. (You can create
CNAME records only for subdomains.)

```go
import "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var api restApi
var hostedZoneForExampleCom interface{}


route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &ARecordProps{
	Zone: hostedZoneForExampleCom,
	Target: route53.RecordTarget_FromAlias(targets.NewApiGateway(api)),
})
```

You can also define a `DomainName` resource directly in order to customize the default behavior:

```go
var acmCertificateForExampleCom interface{}


apigateway.NewDomainName(this, jsii.String("custom-domain"), &DomainNameProps{
	DomainName: jsii.String("example.com"),
	Certificate: acmCertificateForExampleCom,
	EndpointType: apigateway.EndpointType_EDGE,
	 // default is REGIONAL
	SecurityPolicy: apigateway.SecurityPolicy_TLS_1_2,
})
```

Once you have a domain, you can map base paths of the domain to APIs.
The following example will map the URL [https://example.com/go-to-api1](https://example.com/go-to-api1)
to the `api1` API and [https://example.com/boom](https://example.com/boom) to the `api2` API.

```go
var domain domainName
var api1 restApi
var api2 restApi


domain.AddBasePathMapping(api1, &BasePathMappingOptions{
	BasePath: jsii.String("go-to-api1"),
})
domain.AddBasePathMapping(api2, &BasePathMappingOptions{
	BasePath: jsii.String("boom"),
})
```

By default, the base path URL will map to the `deploymentStage` of the `RestApi`.
You can specify a different API `Stage` to which the base path URL will map to.

```go
var domain domainName
var restapi restApi


betaDeploy := apigateway.NewDeployment(this, jsii.String("beta-deployment"), &DeploymentProps{
	Api: restapi,
})
betaStage := apigateway.NewStage(this, jsii.String("beta-stage"), &StageProps{
	Deployment: betaDeploy,
})
domain.AddBasePathMapping(restapi, &BasePathMappingOptions{
	BasePath: jsii.String("api/beta"),
	Stage: betaStage,
})
```

It is possible to create a base path mapping without associating it with a
stage by using the `attachToStage` property. When set to `false`, the stage must be
included in the URL when invoking the API. For example,
[https://example.com/myapi/prod](https://example.com/myapi/prod) will invoke the stage named `prod` from the
`myapi` base path mapping.

```go
var domain domainName
var api restApi


domain.AddBasePathMapping(api, &BasePathMappingOptions{
	BasePath: jsii.String("myapi"),
	AttachToStage: jsii.Boolean(false),
})
```

If you don't specify `basePath`, all URLs under this domain will be mapped
to the API, and you won't be able to map another API to the same domain:

```go
var domain domainName
var api restApi

domain.AddBasePathMapping(api)
```

This can also be achieved through the `mapping` configuration when defining the
domain as demonstrated above.

Base path mappings can also be created with the `BasePathMapping` resource.

```go
var api restApi


domainName := apigateway.DomainName_FromDomainNameAttributes(this, jsii.String("DomainName"), &DomainNameAttributes{
	DomainName: jsii.String("domainName"),
	DomainNameAliasHostedZoneId: jsii.String("domainNameAliasHostedZoneId"),
	DomainNameAliasTarget: jsii.String("domainNameAliasTarget"),
})

apigateway.NewBasePathMapping(this, jsii.String("BasePathMapping"), &BasePathMappingProps{
	DomainName: domainName,
	RestApi: api,
})
```

If you wish to setup this domain with an Amazon Route53 alias, use the `targets.ApiGatewayDomain`:

```go
var hostedZoneForExampleCom interface{}
var domainName domainName

import "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"


route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &ARecordProps{
	Zone: hostedZoneForExampleCom,
	Target: route53.RecordTarget_FromAlias(targets.NewApiGatewayDomain(domainName)),
})
```

### Custom Domains with multi-level api mapping

Additional requirements for creating multi-level path mappings for RestApis:

(both are defaults)

* Must use `SecurityPolicy.TLS_1_2`
* DomainNames must be `EndpointType.REGIONAL`

```go
var acmCertificateForExampleCom interface{}
var restApi restApi


apigateway.NewDomainName(this, jsii.String("custom-domain"), &DomainNameProps{
	DomainName: jsii.String("example.com"),
	Certificate: acmCertificateForExampleCom,
	Mapping: restApi,
	BasePath: jsii.String("orders/v1/api"),
})
```

To then add additional mappings to a domain you can use the `addApiMapping` method.

```go
var acmCertificateForExampleCom interface{}
var restApi restApi
var secondRestApi restApi


domain := apigateway.NewDomainName(this, jsii.String("custom-domain"), &DomainNameProps{
	DomainName: jsii.String("example.com"),
	Certificate: acmCertificateForExampleCom,
	Mapping: restApi,
})

domain.AddApiMapping(secondRestApi.DeploymentStage, &ApiMappingOptions{
	BasePath: jsii.String("orders/v2/api"),
})
```

## Access Logging

Access logging creates logs every time an API method is accessed. Access logs can have information on
who has accessed the API, how the caller accessed the API and what responses were generated.
Access logs are configured on a Stage of the RestApi.
Access logs can be expressed in a format of your choosing, and can contain any access details, with a
minimum that it must include either 'requestId' or 'extendedRequestId'. The list of  variables that
can be expressed in the access log can be found
[here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference).
Read more at [Setting Up CloudWatch API Logging in API
Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)

```go
// production stage
prodLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DeployOptions: &StageOptions{
		AccessLogDestination: apigateway.NewLogGroupLogDestination(prodLogGroup),
		AccessLogFormat: apigateway.AccessLogFormat_JsonWithStandardFields(),
	},
})
deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &DeploymentProps{
	Api: Api,
})

// development stage
devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
apigateway.NewStage(this, jsii.String("dev"), &StageProps{
	Deployment: Deployment,
	AccessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
	AccessLogFormat: apigateway.AccessLogFormat_*JsonWithStandardFields(&JsonWithStandardFieldProps{
		Caller: jsii.Boolean(false),
		HttpMethod: jsii.Boolean(true),
		Ip: jsii.Boolean(true),
		Protocol: jsii.Boolean(true),
		RequestTime: jsii.Boolean(true),
		ResourcePath: jsii.Boolean(true),
		ResponseLength: jsii.Boolean(true),
		Status: jsii.Boolean(true),
		User: jsii.Boolean(true),
	}),
})
```

The following code will generate the access log in the [CLF format](https://en.wikipedia.org/wiki/Common_Log_Format).

```go
logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DeployOptions: &StageOptions{
		AccessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
		AccessLogFormat: apigateway.AccessLogFormat_Clf(),
	},
})
```

You can also configure your own access log format by using the `AccessLogFormat.custom()` API.
`AccessLogField` provides commonly used fields. The following code configures access log to contain.

```go
logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DeployOptions: &StageOptions{
		AccessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
		AccessLogFormat: apigateway.AccessLogFormat_Custom(
		fmt.Sprintf("%v %v %v\n      %v %v", apigateway.AccessLogField_ContextRequestId(), apigateway.AccessLogField_ContextErrorMessage(), apigateway.AccessLogField_ContextErrorMessageString(), apigateway.AccessLogField_ContextAuthorizerError(), apigateway.AccessLogField_ContextAuthorizerIntegrationStatus())),
	},
})
```

You can use the `methodOptions` property to configure
[default method throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html#apigateway-api-level-throttling-in-usage-plan)
for a stage. The following snippet configures the a stage that accepts
100 requests per minute, allowing burst up to 200 requests per minute.

```go
api := apigateway.NewRestApi(this, jsii.String("books"))
deployment := apigateway.NewDeployment(this, jsii.String("my-deployment"), &DeploymentProps{
	Api: Api,
})
stage := apigateway.NewStage(this, jsii.String("my-stage"), &StageProps{
	Deployment: Deployment,
	MethodOptions: map[string]methodDeploymentOptions{
		"/*/*": &methodDeploymentOptions{
			 // This special path applies to all resource paths and all HTTP methods
			"throttlingRateLimit": jsii.Number(100),
			"throttlingBurstLimit": jsii.Number(200),
		},
	},
})
```

Configuring `methodOptions` on the `deployOptions` of `RestApi` will set the
throttling behaviors on the default stage that is automatically created.

```go
api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DeployOptions: &StageOptions{
		MethodOptions: map[string]methodDeploymentOptions{
			"/*/*": &methodDeploymentOptions{
				 // This special path applies to all resource paths and all HTTP methods
				"throttlingRateLimit": jsii.Number(100),
				"throttlingBurstLimit": jsii.Number(1000),
			},
		},
	},
})
```

To write access log files to a Firehose delivery stream destination use the `FirehoseLogDestination` class:

```go
destinationBucket := s3.NewBucket(this, jsii.String("Bucket"))
deliveryStreamRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
})

stream := firehose.NewCfnDeliveryStream(this, jsii.String("MyStream"), &CfnDeliveryStreamProps{
	DeliveryStreamName: jsii.String("amazon-apigateway-delivery-stream"),
	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
		BucketArn: destinationBucket.BucketArn,
		RoleArn: deliveryStreamRole.RoleArn,
	},
})

api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
	DeployOptions: &StageOptions{
		AccessLogDestination: apigateway.NewFirehoseLogDestination(stream),
		AccessLogFormat: apigateway.AccessLogFormat_JsonWithStandardFields(),
	},
})
```

**Note:** The delivery stream name must start with `amazon-apigateway-`.

> Visit [Logging API calls to Kinesis Data Firehose](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-logging-to-kinesis.html) for more details.

## Cross Origin Resource Sharing (CORS)

[Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a mechanism
that uses additional HTTP headers to tell browsers to give a web application
running at one origin, access to selected resources from a different origin. A
web application executes a cross-origin HTTP request when it requests a resource
that has a different origin (domain, protocol, or port) from its own.

You can add the CORS [preflight](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#Preflighted_requests) OPTIONS
HTTP method to any API resource via the `defaultCorsPreflightOptions` option or by calling the `addCorsPreflight` on a specific resource.

The following example will enable CORS for all methods and all origins on all resources of the API:

```go
apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
	DefaultCorsPreflightOptions: &CorsOptions{
		AllowOrigins: apigateway.Cors_ALL_ORIGINS(),
		AllowMethods: apigateway.Cors_ALL_METHODS(),
	},
})
```

The following example will add an OPTIONS method to the `myResource` API resource, which
only allows GET and PUT HTTP requests from the origin [https://amazon.com.](https://amazon.com.)

```go
var myResource resource


myResource.AddCorsPreflight(&CorsOptions{
	AllowOrigins: []*string{
		jsii.String("https://amazon.com"),
	},
	AllowMethods: []*string{
		jsii.String("GET"),
		jsii.String("PUT"),
	},
})
```

See the
[`CorsOptions`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-apigateway.CorsOptions.html)
API reference for a detailed list of supported configuration options.

You can specify defaults this at the resource level, in which case they will be applied to the entire resource sub-tree:

```go
var resource resource


subtree := resource.AddResource(jsii.String("subtree"), &ResourceOptions{
	DefaultCorsPreflightOptions: &CorsOptions{
		AllowOrigins: []*string{
			jsii.String("https://amazon.com"),
		},
	},
})
```

This means that all resources under `subtree` (inclusive) will have a preflight
OPTIONS added to them.

See [#906](https://github.com/aws/aws-cdk/issues/906) for a list of CORS
features which are not yet supported.

## Endpoint Configuration

API gateway allows you to specify an
[API Endpoint Type](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html).
To define an endpoint type for the API gateway, use `endpointConfiguration` property:

```go
api := apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
	EndpointConfiguration: &EndpointConfiguration{
		Types: []endpointType{
			apigateway.*endpointType_EDGE,
		},
	},
})
```

You can also create an association between your Rest API and a VPC endpoint. By doing so,
API Gateway will generate a new
Route53 Alias DNS record which you can use to invoke your private APIs. More info can be found
[here](https://docs.aws.amazon.com/apigateway/latest/developerguide/associate-private-api-with-vpc-endpoint.html).

Here is an example:

```go
var someEndpoint iVpcEndpoint


api := apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
	EndpointConfiguration: &EndpointConfiguration{
		Types: []endpointType{
			apigateway.*endpointType_PRIVATE,
		},
		VpcEndpoints: []*iVpcEndpoint{
			someEndpoint,
		},
	},
})
```

By performing this association, we can invoke the API gateway using the following format:

```plaintext
https://{rest-api-id}-{vpce-id}.execute-api.{region}.amazonaws.com/{stage}
```

## Private Integrations

A private integration makes it simple to expose HTTP/HTTPS resources behind an
Amazon VPC for access by clients outside of the VPC. The private integration uses
an API Gateway resource of `VpcLink` to encapsulate connections between API
Gateway and targeted VPC resources.
The `VpcLink` is then attached to the `Integration` of a specific API Gateway
Method. The following code sets up a private integration with a network load
balancer -

```go
import elbv2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("NLB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
})
link := apigateway.NewVpcLink(this, jsii.String("link"), &VpcLinkProps{
	Targets: []iNetworkLoadBalancer{
		nlb,
	},
})

integration := apigateway.NewIntegration(&IntegrationProps{
	Type: apigateway.IntegrationType_HTTP_PROXY,
	IntegrationHttpMethod: jsii.String("ANY"),
	Options: &IntegrationOptions{
		ConnectionType: apigateway.ConnectionType_VPC_LINK,
		VpcLink: link,
	},
})
```

The uri for the private integration, in the case of a VpcLink, will be set to the DNS name of
the VPC Link's NLB. If the VPC Link has multiple NLBs or the VPC Link is imported or the DNS
name cannot be determined for any other reason, the user is expected to specify the `uri`
property.

Any existing `VpcLink` resource can be imported into the CDK app via the `VpcLink.fromVpcLinkId()`.

```go
awesomeLink := apigateway.VpcLink_FromVpcLinkId(this, jsii.String("awesome-vpc-link"), jsii.String("us-east-1_oiuR12Abd"))
```

## Gateway response

If the Rest API fails to process an incoming request, it returns to the client an error response without forwarding the
request to the integration backend. API Gateway has a set of standard response messages that are sent to the client for
each type of error. These error responses can be configured on the Rest API. The list of Gateway responses that can be
configured can be found [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html).
Learn more about [Gateway
Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-gatewayResponse-definition.html).

The following code configures a Gateway Response when the response is 'access denied':

```go
api := apigateway.NewRestApi(this, jsii.String("books-api"))
api.AddGatewayResponse(jsii.String("test-response"), &GatewayResponseOptions{
	Type: apigateway.ResponseType_ACCESS_DENIED(),
	StatusCode: jsii.String("500"),
	ResponseHeaders: map[string]*string{
		// Note that values must be enclosed within a pair of single quotes
		"Access-Control-Allow-Origin": jsii.String("'test.com'"),
		"test-key": jsii.String("'test-value'"),
	},
	Templates: map[string]*string{
		"application/json": jsii.String("{ \"message\": $context.error.messageString, \"statusCode\": \"488\", \"type\": \"$context.error.responseType\" }"),
	},
})
```

## OpenAPI Definition

CDK supports creating a REST API by importing an OpenAPI definition file. It currently supports OpenAPI v2.0 and OpenAPI
v3.0 definition files. Read more about [Configuring a REST API using
OpenAPI](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html).

The following code creates a REST API using an external OpenAPI definition JSON file -

```go
var integration integration


api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &SpecRestApiProps{
	ApiDefinition: apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json")),
})

booksResource := api.Root.AddResource(jsii.String("books"))
booksResource.AddMethod(jsii.String("GET"), integration)
```

It is possible to use the `addResource()` API to define additional API Gateway Resources.

**Note:** Deployment will fail if a Resource of the same name is already defined in the Open API specification.

**Note:** Any default properties configured, such as `defaultIntegration`, `defaultMethodOptions`, etc. will only be
applied to Resources and Methods defined in the CDK, and not the ones defined in the spec. Use the [API Gateway
extensions to OpenAPI](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html)
to configure these.

There are a number of limitations in using OpenAPI definitions in API Gateway. Read the [Amazon API Gateway important
notes for REST APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-known-issues.html#api-gateway-known-issues-rest-apis)
for more details.

**Note:** When starting off with an OpenAPI definition using `SpecRestApi`, it is not possible to configure some
properties that can be configured directly in the OpenAPI specification file. This is to prevent people duplication
of these properties and potential confusion.

### Endpoint configuration

By default, `SpecRestApi` will create an edge optimized endpoint.

This can be modified as shown below:

```go
var apiDefinition apiDefinition


api := apigateway.NewSpecRestApi(this, jsii.String("ExampleRestApi"), &SpecRestApiProps{
	ApiDefinition: ApiDefinition,
	EndpointTypes: []endpointType{
		apigateway.*endpointType_PRIVATE,
	},
})
```

**Note:** For private endpoints you will still need to provide the
[`x-amazon-apigateway-policy`](https://docs.aws.amazon.com/apigateway/latest/developerguide/openapi-extensions-policy.html) and
[`x-amazon-apigateway-endpoint-configuration`](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-endpoint-configuration.html)
in your openApi file.

## Metrics

The API Gateway service sends metrics around the performance of Rest APIs to Amazon CloudWatch.
These metrics can be referred to using the metric APIs available on the `RestApi`, `Stage` and `Method` constructs.
Note that detailed metrics must be enabled for a stage to use the `Method` metrics.
Read more about [API Gateway metrics](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-metrics-and-dimensions.html), including enabling detailed metrics.
The APIs with the `metric` prefix can be used to get reference to specific metrics for this API. For example:

```go
api := apigateway.NewRestApi(this, jsii.String("my-api"))
stage := api.DeploymentStage
method := api.Root.AddMethod(jsii.String("GET"))

clientErrorApiMetric := api.MetricClientError()
serverErrorStageMetric := stage.MetricServerError()
latencyMethodMetric := method.MetricLatency(stage)
```

## APIGateway v2

APIGateway v2 APIs are now moved to its own package named `aws-apigatewayv2`. For backwards compatibility, existing
APIGateway v2 "CFN resources" (such as `CfnApi`) that were previously exported as part of this package, are still
exported from here and have been marked deprecated. However, updates to these CloudFormation resources, such as new
properties and new resource types will not be available.

Move to using `aws-apigatewayv2` to get the latest APIs and updates.

---


This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

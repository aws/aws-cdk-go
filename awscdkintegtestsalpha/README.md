# integ-tests

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## Overview

This library is meant to be used in combination with the [integ-runner](https://github.com/aws/aws-cdk/tree/main/packages/%40aws-cdk/integ-runner) CLI
to enable users to write and execute integration tests for AWS CDK Constructs.

An integration test should be defined as a CDK application, and
there should be a 1:1 relationship between an integration test and a CDK application.

So for example, in order to create an integration test called `my-function`
we would need to create a file to contain our integration test application.

*test/integ.my-function.ts*

```go
app := awscdk.NewApp()
stack := awscdk.NewStack()
lambda.NewFunction(stack, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

This is a self contained CDK application which we could deploy by running

```bash
cdk deploy --app 'node test/integ.my-function.js'
```

In order to turn this into an integration test, all that is needed is to
use the `IntegTest` construct.

```go
var app app
var stack stack

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})
```

You will notice that the `stack` is registered to the `IntegTest` as a test case.
Each integration test can contain multiple test cases, which are just instances
of a stack. See the [Usage](#usage) section for more details.

## Usage

### IntegTest

Suppose you have a simple stack, that only encapsulates a Lambda function with a
certain handler:

```go
type stackUnderTestProps struct {
	stackProps
	architecture architecture
}

type stackUnderTest struct {
	stack
}

func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
	this := &stackUnderTest{}
	newStack_Override(this, scope, id, props)

	lambda.NewFunction(this, jsii.String("Handler"), &FunctionProps{
		Runtime: lambda.Runtime_NODEJS_LATEST(),
		Handler: jsii.String("index.handler"),
		Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		Architecture: props.architecture,
	})
	return this
}
```

You may want to test this stack under different conditions. For example, we want
this stack to be deployed correctly, regardless of the architecture we choose
for the Lambda function. In particular, it should work for both `ARM_64` and
`X86_64`. So you can create an `IntegTestCase` that exercises both scenarios:

```go
type stackUnderTestProps struct {
	stackProps
	architecture architecture
}

type stackUnderTest struct {
	stack
}

func newStackUnderTest(scope construct, id *string, props stackUnderTestProps) *stackUnderTest {
	this := &stackUnderTest{}
	newStack_Override(this, scope, id, props)

	lambda.NewFunction(this, jsii.String("Handler"), &FunctionProps{
		Runtime: lambda.Runtime_NODEJS_LATEST(),
		Handler: jsii.String("index.handler"),
		Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		Architecture: props.architecture,
	})
	return this
}

// Beginning of the test suite
app := awscdk.NewApp()

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("DifferentArchitectures"), &IntegTestProps{
	TestCases: []*stack{
		NewStackUnderTest(app, jsii.String("Stack1"), &stackUnderTestProps{
			architecture: lambda.*architecture_ARM_64(),
		}),
		NewStackUnderTest(app, jsii.String("Stack2"), &stackUnderTestProps{
			architecture: lambda.*architecture_X86_64(),
		}),
	},
})
```

This is all the instruction you need for the integration test runner to know
which stacks to synthesize, deploy and destroy. But you may also need to
customize the behavior of the runner by changing its parameters. For example:

```go
app := awscdk.NewApp()

stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))

stack := awscdk.NewStack(app, jsii.String("stack"))

testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &IntegTestProps{
	TestCases: []stack{
		stackUnderTest,
	},
	DiffAssets: jsii.Boolean(true),
	StackUpdateWorkflow: jsii.Boolean(true),
	CdkCommandOptions: &CdkCommands{
		Deploy: &DeployCommand{
			Args: &DeployOptions{
				RequireApproval: awscdklibcloudassemblyschema.RequireApproval_NEVER,
				Json: jsii.Boolean(true),
			},
		},
		Destroy: &DestroyCommand{
			Args: &DestroyOptions{
				Force: jsii.Boolean(true),
			},
		},
	},
})
```

### IntegTestCaseStack

In the majority of cases an integration test will contain a single `IntegTestCase`.
By default when you create an `IntegTest` an `IntegTestCase` is created for you
and all of your test cases are registered to this `IntegTestCase`. The `IntegTestCase`
and `IntegTestCaseStack` constructs are only needed when it is necessary to
defined different options for individual test cases.

For example, you might want to have one test case where `diffAssets` is enabled.

```go
var app app
var stackUnderTest stack

testCaseWithAssets := awscdkintegtestsalpha.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &IntegTestCaseStackProps{
	DiffAssets: jsii.Boolean(true),
})

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stackUnderTest,
		testCaseWithAssets,
	},
})
```

## Assertions

This library also provides a utility to make assertions against the infrastructure that the integration test deploys.

There are two main scenarios in which assertions are created.

* Part of an integration test using `integ-runner`

In this case you would create an integration test using the `IntegTest` construct and then make assertions using the `assert` property.
You should **not** utilize the assertion constructs directly, but should instead use the `methods` on `IntegTest.assertions`.

```go
var app app
var stack stack


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})
integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("getObject"))
```

By default an assertions stack is automatically generated for you. You may however provide your own stack to use.

```go
var app app
var stack stack
var assertionStack stack


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
	AssertionStack: assertionStack,
})
integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("getObject"))
```

* Part of a  normal CDK deployment

In this case you may be using assertions as part of a normal CDK deployment in order to make an assertion on the infrastructure
before the deployment is considered successful. In this case you can utilize the assertions constructs directly.

```go
var myAppStack stack


awscdkintegtestsalpha.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &AwsApiCallProps{
	Service: jsii.String("S3"),
	Api: jsii.String("getObject"),
})
```

### DeployAssert

Assertions are created by using the `DeployAssert` construct. This construct creates it's own `Stack` separate from
any stacks that you create as part of your integration tests. This `Stack` is treated differently from other stacks
by the `integ-runner` tool. For example, this stack will not be diffed by the `integ-runner`.

`DeployAssert` also provides utilities to register your own assertions.

```go
var myCustomResource customResource
var stack stack
var app app


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})
integ.Assertions.Expect(jsii.String("CustomAssertion"), awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"foo": jsii.String("bar"),
}), awscdkintegtestsalpha.ActualResult_FromCustomResource(myCustomResource, jsii.String("data")))
```

In the above example an assertion is created that will trigger a user defined `CustomResource`
and assert that the `data` attribute is equal to `{ foo: 'bar' }`.

### API Calls

A common method to retrieve the "actual" results to compare with what is expected is to make an
API call to receive some data. This library does this by utilizing CloudFormation custom resources
which means that CloudFormation will call out to a Lambda Function which will
make the API call.

#### HttpApiCall

Using the `HttpApiCall` will use the
[node-fetch](https://github.com/node-fetch/node-fetch) JavaScript library to
make the HTTP call.

This can be done by using the class directory (in the case of a normal deployment):

```go
var stack stack


awscdkintegtestsalpha.NewHttpApiCall(stack, jsii.String("MyAsssertion"), &HttpCallProps{
	Url: jsii.String("https://example-api.com/abc"),
})
```

Or by using the `httpApiCall` method on `DeployAssert` (when writing integration tests):

```go
var app app
var stack stack

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})
integ.Assertions.HttpApiCall(jsii.String("https://example-api.com/abc"))
```

#### AwsApiCall

Using the `AwsApiCall` construct will use the AWS JavaScript SDK to make the API call.

This can be done by using the class directory (in the case of a normal deployment):

```go
var stack stack


awscdkintegtestsalpha.NewAwsApiCall(stack, jsii.String("MyAssertion"), &AwsApiCallProps{
	Service: jsii.String("SQS"),
	Api: jsii.String("receiveMessage"),
	Parameters: map[string]*string{
		"QueueUrl": jsii.String("url"),
	},
})
```

Or by using the `awsApiCall` method on `DeployAssert` (when writing integration tests):

```go
var app app
var stack stack

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})
integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
	"QueueUrl": jsii.String("url"),
})
```

You must specify the `service` and the `api` when using The `AwsApiCall` construct.
The `service` is the name of an AWS service, in one of the following forms:

* An AWS SDK for JavaScript v3 package name (`@aws-sdk/client-api-gateway`)
* An AWS SDK for JavaScript v3 client name (`api-gateway`)
* An AWS SDK for JavaScript v2 constructor name (`APIGateway`)
* A lowercase AWS SDK for JavaScript v2 constructor name (`apigateway`)

The `api` is the name of an AWS API call, in one of the following forms:

* An API call name as found in the API Reference documentation (`GetObject`)
* The API call name starting with a lowercase letter (`getObject`)
* The AWS SDK for JavaScript v3 command class name (`GetObjectCommand`)

By default, the `AwsApiCall` construct will automatically add the correct IAM policies
to allow the Lambda function to make the API call. It does this based on the `service`
and `api` that is provided. In the above example the service is `SQS` and the api is
`receiveMessage` so it will create a policy with `Action: 'sqs:ReceiveMessage`.

There are some cases where the permissions do not exactly match the service/api call, for
example the S3 `listObjectsV2` api. In these cases it is possible to add the correct policy
by accessing the `provider` object.

```go
var app app
var stack stack
var integ integTest


apiCall := integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("listObjectsV2"), map[string]*string{
	"Bucket": jsii.String("mybucket"),
})

apiCall.Provider.AddToRolePolicy(map[string]interface{}{
	"Effect": jsii.String("Allow"),
	"Action": []*string{
		jsii.String("s3:GetObject"),
		jsii.String("s3:ListBucket"),
	},
	"Resource": []*string{
		jsii.String("*"),
	},
})
```

When executing `waitForAssertion()`, it is necessary to add an IAM policy using `waiterProvider.addToRolePolicy()`.
Because `IApiCall` does not have a `waiterProvider` property, you need to cast it to `AwsApiCall`.

```go
var integ integTest


apiCall := integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("listObjectsV2"), map[string]*string{
	"Bucket": jsii.String("mybucket"),
}).WaitForAssertions().(awsApiCall)

apiCall.WaiterProvider.AddToRolePolicy(map[string]interface{}{
	"Effect": jsii.String("Allow"),
	"Action": []*string{
		jsii.String("s3:GetObject"),
		jsii.String("s3:ListBucket"),
	},
	"Resource": []*string{
		jsii.String("*"),
	},
})
```

Note that addToRolePolicy() uses direct IAM JSON policy blobs, not a iam.PolicyStatement
object like you will see in the rest of the CDK.

### EqualsAssertion

This library currently provides the ability to assert that two values are equal
to one another by utilizing the `EqualsAssertion` class. This utilizes a Lambda
backed `CustomResource` which in tern uses the [Match](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions.Match.html) utility from the
[@aws-cdk/assertions](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.assertions-readme.html) library.

```go
var app app
var stack stack
var queue queue
var fn iFunction


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})

integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
	FunctionName: fn.FunctionName,
	InvocationType: awscdkintegtestsalpha.InvocationType_EVENT,
	Payload: jSON.stringify(map[string]*string{
		"status": jsii.String("OK"),
	}),
})

message := integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]interface{}{
	"QueueUrl": queue.queueUrl,
	"WaitTimeSeconds": jsii.Number(20),
})

message.AssertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"requestContext": map[string]*string{
		"condition": jsii.String("Success"),
	},
	"requestPayload": map[string]*string{
		"status": jsii.String("OK"),
	},
	"responseContext": map[string]*f64{
		"statusCode": jsii.Number(200),
	},
	"responsePayload": jsii.String("success"),
}))
```

#### Match

`integ-tests` also provides a `Match` utility similar to the `@aws-cdk/assertions` module. `Match`
can be used to construct the `ExpectedResult`. While the utility is similar, only a subset of methods are currently available on the `Match` utility of this module: `arrayWith`, `objectLike`, `stringLikeRegexp` and `serializedJson`.

```go
var message awsApiCall


message.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"Messages": awscdkintegtestsalpha.Match_arrayWith([]interface{}{
		map[string]map[string]map[string]interface{}{
			"Payload": awscdkintegtestsalpha.Match_serializedJson(map[string]interface{}{
				"key": jsii.String("value"),
			}),
		},
		map[string]map[string]map[string][]interface{}{
			"Body": map[string]map[string][]interface{}{
				"Values": awscdkintegtestsalpha.Match_arrayWith([]interface{}{
					map[string]*f64{
						"Asdf": jsii.Number(3),
					},
				}),
				"Message": awscdkintegtestsalpha.Match_stringLikeRegexp(jsii.String("message")),
			},
		},
	}),
}))
```

### Examples

#### Invoke a Lambda Function

In this example there is a Lambda Function that is invoked and
we assert that the payload that is returned is equal to '200'.

```go
var lambdaFunction iFunction
var app app


stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
	TestCases: []stack{
		stack,
	},
})

invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
	FunctionName: lambdaFunction.FunctionName,
})
invoke.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"Payload": jsii.String("200"),
}))
```

The above example will by default create a CloudWatch log group that's never
expired. If you want to configure it with custom log retention days, you need
to specify the `logRetention` property.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"

var lambdaFunction iFunction
var app app


stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
	TestCases: []stack{
		stack,
	},
})

invoke := integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
	FunctionName: lambdaFunction.FunctionName,
	LogRetention: logs.RetentionDays_ONE_WEEK,
})
```

#### Make an AWS API Call

In this example there is a StepFunctions state machine that is executed
and then we assert that the result of the execution is successful.

```go
var app app
var stack stack
var sm iStateMachine


testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})

// Start an execution
start := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("startExecution"), map[string]*string{
	"stateMachineArn": sm.stateMachineArn,
})

// describe the results of the execution
describe := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
})

// assert the results
describe.Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
}))
```

#### Chain ApiCalls

Sometimes it may be necessary to chain API Calls. Since each API call is its own resource, all you
need to do is add a dependency between the calls. There is an helper method `next` that can be used.

```go
var integ integTest


integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("putObject"), map[string]*string{
	"Bucket": jsii.String("amzn-s3-demo-bucket"),
	"Key": jsii.String("my-key"),
	"Body": jsii.String("helloWorld"),
}).Next(integ.Assertions.AwsApiCall(jsii.String("S3"), jsii.String("getObject"), map[string]*string{
	"Bucket": jsii.String("amzn-s3-demo-bucket"),
	"Key": jsii.String("my-key"),
}))
```

#### Wait for results

A common use case when performing assertions is to wait for a condition to pass. Sometimes the thing
that you are asserting against is not done provisioning by the time the assertion runs. In these
cases it is possible to run the assertion asynchronously by calling the `waitForAssertions()` method.

Taking the example above of executing a StepFunctions state machine, depending on the complexity of
the state machine, it might take a while for it to complete.

```go
var app app
var stack stack
var sm iStateMachine


testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &IntegTestProps{
	TestCases: []*stack{
		stack,
	},
})

// Start an execution
start := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("startExecution"), map[string]*string{
	"stateMachineArn": sm.stateMachineArn,
})

// describe the results of the execution
describe := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
}).Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
})).WaitForAssertions()
```

When you call `waitForAssertions()` the assertion provider will continuously make the `awsApiCall` until the
`ExpectedResult` is met. You can also control the parameters for waiting, for example:

```go
var testCase integTest
var start iApiCall


describe := testCase.Assertions.AwsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
}).Expect(awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
})).WaitForAssertions(&WaiterStateMachineOptions{
	TotalTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
	Interval: awscdk.Duration_Seconds(jsii.Number(15)),
	BackoffRate: jsii.Number(3),
})
```

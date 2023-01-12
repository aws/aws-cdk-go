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
lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
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

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
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

	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
		runtime: lambda.runtime_NODEJS_14_X(),
		handler: jsii.String("index.handler"),
		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		architecture: props.architecture,
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

	lambda.NewFunction(this, jsii.String("Handler"), &functionProps{
		runtime: lambda.runtime_NODEJS_14_X(),
		handler: jsii.String("index.handler"),
		code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
		architecture: props.architecture,
	})
	return this
}

// Beginning of the test suite
app := awscdk.NewApp()

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("DifferentArchitectures"), &integTestProps{
	testCases: []*stack{
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

testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &integTestProps{
	testCases: []stack{
		stackUnderTest,
	},
	diffAssets: jsii.Boolean(true),
	stackUpdateWorkflow: jsii.Boolean(true),
	cdkCommandOptions: &cdkCommands{
		deploy: &deployCommand{
			args: &deployOptions{
				requireApproval: awscdk.RequireApproval_NEVER,
				json: jsii.Boolean(true),
			},
		},
		destroy: &destroyCommand{
			args: &destroyOptions{
				force: jsii.Boolean(true),
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

testCaseWithAssets := awscdkintegtestsalpha.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &integTestCaseStackProps{
	diffAssets: jsii.Boolean(true),
})

awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
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


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})
integ.assertions.awsApiCall(jsii.String("S3"), jsii.String("getObject"))
```

By default an assertions stack is automatically generated for you. You may however provide your own stack to use.

```go
var app app
var stack stack
var assertionStack stack


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
		stack,
	},
	assertionStack: assertionStack,
})
integ.assertions.awsApiCall(jsii.String("S3"), jsii.String("getObject"))
```

* Part of a  normal CDK deployment

In this case you may be using assertions as part of a normal CDK deployment in order to make an assertion on the infrastructure
before the deployment is considered successful. In this case you can utilize the assertions constructs directly.

```go
var myAppStack stack


awscdkintegtestsalpha.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &awsApiCallProps{
	service: jsii.String("S3"),
	api: jsii.String("getObject"),
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


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})
integ.assertions.expect(jsii.String("CustomAssertion"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"foo": jsii.String("bar"),
}), awscdkintegtestsalpha.ActualResult.fromCustomResource(myCustomResource, jsii.String("data")))
```

In the above example an assertion is created that will trigger a user defined `CustomResource`
and assert that the `data` attribute is equal to `{ foo: 'bar' }`.

### AwsApiCall

A common method to retrieve the "actual" results to compare with what is expected is to make an
AWS API call to receive some data. This library does this by utilizing CloudFormation custom resources
which means that CloudFormation will call out to a Lambda Function which will
use the AWS JavaScript SDK to make the API call.

This can be done by using the class directory (in the case of a normal deployment):

```go
var stack stack


awscdkintegtestsalpha.NewAwsApiCall(stack, jsii.String("MyAssertion"), &awsApiCallProps{
	service: jsii.String("SQS"),
	api: jsii.String("receiveMessage"),
	parameters: map[string]*string{
		"QueueUrl": jsii.String("url"),
	},
})
```

Or by using the `awsApiCall` method on `DeployAssert` (when writing integration tests):

```go
var app app
var stack stack

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})
integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
	"QueueUrl": jsii.String("url"),
})
```

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


apiCall := integ.assertions.awsApiCall(jsii.String("S3"), jsii.String("listObjectsV2"), map[string]*string{
	"Bucket": jsii.String("mybucket"),
})

apiCall.provider.addToRolePolicy(map[string]interface{}{
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


integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})

integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	functionName: fn.functionName,
	invocationType: awscdkintegtestsalpha.InvocationType_EVENT,
	payload: jSON.stringify(map[string]*string{
		"status": jsii.String("OK"),
	}),
})

message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]interface{}{
	"QueueUrl": queue.queueUrl,
	"WaitTimeSeconds": jsii.Number(20),
})

message.assertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
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


message.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"Messages": awscdkintegtestsalpha.Match.arrayWith([]interface{}{
		map[string]map[string]map[string]interface{}{
			"Payload": awscdkintegtestsalpha.Match.serializedJson(map[string]interface{}{
				"key": jsii.String("value"),
			}),
		},
		map[string]map[string]map[string][]interface{}{
			"Body": map[string]map[string][]interface{}{
				"Values": awscdkintegtestsalpha.Match.arrayWith([]interface{}{
					map[string]*f64{
						"Asdf": jsii.Number(3),
					},
				}),
				"Message": awscdkintegtestsalpha.Match.stringLikeRegexp(jsii.String("message")),
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

integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &integTestProps{
	testCases: []stack{
		stack,
	},
})

invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	functionName: lambdaFunction.functionName,
})
invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"Payload": jsii.String("200"),
}))
```

#### Make an AWS API Call

In this example there is a StepFunctions state machine that is executed
and then we assert that the result of the execution is successful.

```go
var app app
var stack stack
var sm iStateMachine


testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})

// Start an execution
start := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("startExecution"), map[string]*string{
	"stateMachineArn": sm.stateMachineArn,
})

// describe the results of the execution
describe := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
})

// assert the results
describe.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
}))
```

#### Chain ApiCalls

Sometimes it may be necessary to chain API Calls. Since each API call is its own resource, all you
need to do is add a dependency between the calls. There is an helper method `next` that can be used.

```go
var integ integTest


integ.assertions.awsApiCall(jsii.String("S3"), jsii.String("putObject"), map[string]*string{
	"Bucket": jsii.String("my-bucket"),
	"Key": jsii.String("my-key"),
	"Body": jsii.String("helloWorld"),
}).next(integ.assertions.awsApiCall(jsii.String("S3"), jsii.String("getObject"), map[string]*string{
	"Bucket": jsii.String("my-bucket"),
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


testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &integTestProps{
	testCases: []*stack{
		stack,
	},
})

// Start an execution
start := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("startExecution"), map[string]*string{
	"stateMachineArn": sm.stateMachineArn,
})

// describe the results of the execution
describe := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
}).expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
})).waitForAssertions()
```

When you call `waitForAssertions()` the assertion provider will continuously make the `awsApiCall` until the
`ExpectedResult` is met. You can also control the parameters for waiting, for example:

```go
var testCase integTest
var start iApiCall


describe := testCase.assertions.awsApiCall(jsii.String("StepFunctions"), jsii.String("describeExecution"), map[string]*string{
	"executionArn": start.getAttString(jsii.String("executionArn")),
}).expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	"status": jsii.String("SUCCEEDED"),
})).waitForAssertions(&waiterStateMachineOptions{
	totalTimeout: awscdk.Duration.minutes(jsii.Number(5)),
	interval: awscdk.Duration.seconds(jsii.Number(15)),
	backoffRate: jsii.Number(3),
})
```

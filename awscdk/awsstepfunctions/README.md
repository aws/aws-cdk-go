# AWS Step Functions Construct Library

The `aws-cdk-lib/aws-stepfunctions` package contains constructs for building
serverless workflows using objects. Use this in conjunction with the
`aws-cdk-lib/aws-stepfunctions-tasks` package, which contains classes used
to call other AWS services.

Defining a workflow looks like this (for the [Step Functions Job Poller
example](https://docs.aws.amazon.com/step-functions/latest/dg/job-status-poller-sample.html)):

## Example

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var submitLambda function
var getStatusLambda function


submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &LambdaInvokeProps{
	LambdaFunction: submitLambda,
	// Lambda's result is in the attribute `guid`
	OutputPath: jsii.String("$.guid"),
})

waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &WaitProps{
	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
})

getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &LambdaInvokeProps{
	LambdaFunction: getStatusLambda,
	// Pass just the field named "guid" into the Lambda, put the
	// Lambda's result in a field called "status" in the response
	InputPath: jsii.String("$.guid"),
	OutputPath: jsii.String("$.status"),
})

jobFailed := sfn.NewFail(this, jsii.String("Job Failed"), &FailProps{
	Cause: jsii.String("AWS Batch Job Failed"),
	Error: jsii.String("DescribeJob returned FAILED"),
})

finalStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Final Job Status"), &LambdaInvokeProps{
	LambdaFunction: getStatusLambda,
	// Use "guid" field as input
	InputPath: jsii.String("$.guid"),
	OutputPath: jsii.String("$.Payload"),
})

definition := submitJob.Next(waitX).Next(getStatus).Next(sfn.NewChoice(this, jsii.String("Job Complete?")).When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("FAILED")), jobFailed).When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("SUCCEEDED")), finalStatus).Otherwise(waitX))

sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
	Comment: jsii.String("a super cool state machine"),
})
```

You can find more sample snippets and learn more about the service integrations
in the `aws-cdk-lib/aws-stepfunctions-tasks` package.

## State Machine

A `stepfunctions.StateMachine` is a resource that takes a state machine
definition. The definition is specified by its start state, and encompasses
all states reachable from the start state:

```go
startState := sfn.NewPass(this, jsii.String("StartState"))

sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(startState),
})
```

State machines are made up of a sequence of **Steps**, which represent different actions
taken in sequence. Some of these steps represent *control flow* (like `Choice`, `Map` and `Wait`)
while others represent calls made against other AWS services (like `LambdaInvoke`).
The second category are called `Task`s and they can all be found in the module [`aws-stepfunctions-tasks`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_stepfunctions_tasks-readme.html).

State machines execute using an IAM Role, which will automatically have all
permissions added that are required to make all state machine tasks execute
properly (for example, permissions to invoke any Lambda functions you add to
your workflow). A role will be created by default, but you can supply an
existing one as well.

Set the `removalPolicy` prop to `RemovalPolicy.RETAIN` if you want to retain the execution
history when CloudFormation deletes your state machine.

Alternatively you can specify an existing step functions definition by providing a string or a file that contains the ASL JSON.

```go
sfn.NewStateMachine(this, jsii.String("StateMachineFromString"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromString(jsii.String("{\"StartAt\":\"Pass\",\"States\":{\"Pass\":{\"Type\":\"Pass\",\"End\":true}}}")),
})

sfn.NewStateMachine(this, jsii.String("StateMachineFromFile"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromFile(jsii.String("./asl.json")),
})
```

## State Machine Data

An Execution represents each time the State Machine is run. Every Execution has [State Machine
Data](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-data.html):
a JSON document containing keys and values that is fed into the state machine,
gets modified by individual steps as the state machine progresses, and finally
is produced as output.

By default, the entire Data object is passed into every state, and the return data of the step
becomes new the new Data object. This behavior can be modified by supplying values for `inputPath`,
`resultSelector`, `resultPath` and `outputPath`.

### Manipulating state machine data using inputPath, resultSelector, resultPath and outputPath

These properties impact how each individual step interacts with the state machine data:

* `stateName`: the name of the state in the state machine definition. If not supplied, defaults to the construct id.
* `inputPath`: the part of the data object that gets passed to the step (`itemsPath` for `Map` states)
* `resultSelector`: the part of the step result that should be added to the state machine data
* `resultPath`: where in the state machine data the step result should be inserted
* `outputPath`: what part of the state machine data should be retained
* `errorPath`: the part of the data object that gets returned as the step error
* `causePath`: the part of the data object that gets returned as the step cause

Their values should be a string indicating a [JSON path](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-paths.html) into the State Machine Data object (like `"$.MyKey"`). If absent, the values are treated as if they were `"$"`, which means the entire object.

The following pseudocode shows how AWS Step Functions uses these parameters when executing a step:

```js
// Schematically show how Step Functions evaluates functions.
// [] represents indexing into an object by a using JSON path.

input = state[inputPath]

result = invoke_step(select_parameters(input))

state[resultPath] = result[resultSelector]

state = state[outputPath]
```

Instead of a JSON path string, each of these paths can also have the special value `JsonPath.DISCARD`, which causes the corresponding indexing expression to return an empty object (`{}`). Effectively, that means there will be an empty input object, an empty result object, no effect on the state, or an empty state, respectively.

Some steps (mostly Tasks) have *Parameters*, which are selected differently. See the next section.

See the official documentation on [input and output processing in Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html).

### Passing Parameters to Tasks

Tasks take parameters, whose values can be taken from the State Machine Data object. For example, your
workflow may want to start a CodeBuild with an environment variable that is taken from the State Machine data, or pass part of the State Machine Data into an AWS Lambda Function.

In the original JSON-based states language used by AWS Step Functions, you would
add `.$` to the end of a key to indicate that a value needs to be interpreted as
a JSON path. In the CDK API you do not change the names of any keys. Instead, you
pass special values. There are 3 types of task inputs to consider:

* Tasks that accept a "payload" type of input (like AWS Lambda invocations, or posting messages to SNS topics or SQS queues), will take an object of type `TaskInput`, like `TaskInput.fromObject()` or `TaskInput.fromJsonPathAt()`.
* When tasks expect individual string or number values to customize their behavior, you can also pass a value constructed by `JsonPath.stringAt()` or `JsonPath.numberAt()`.
* When tasks expect strongly-typed resources and you want to vary the resource that is referenced based on a name from the State Machine Data, reference the resource as if it was external (using `JsonPath.stringAt()`). For example, for a Lambda function: `Function.fromFunctionName(this, 'ReferencedFunction', JsonPath.stringAt('$.MyFunctionName'))`.

For example, to pass the value that's in the data key of `OrderId` to a Lambda
function as you invoke it, use `JsonPath.stringAt('$.OrderId')`, like so:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var orderFn function


submitJob := tasks.NewLambdaInvoke(this, jsii.String("InvokeOrderProcessor"), &LambdaInvokeProps{
	LambdaFunction: orderFn,
	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
		"OrderId": sfn.JsonPath_stringAt(jsii.String("$.OrderId")),
	}),
})
```

The following methods are available:

| Method | Purpose |
|--------|---------|
| `JsonPath.stringAt('$.Field')` | reference a field, return the type as a `string`. |
| `JsonPath.listAt('$.Field')` | reference a field, return the type as a list of strings. |
| `JsonPath.numberAt('$.Field')` | reference a field, return the type as a number. Use this for functions that expect a number argument. |
| `JsonPath.objectAt('$.Field')` | reference a field, return the type as an `IResolvable`. Use this for functions that expect an object argument. |
| `JsonPath.entirePayload` | reference the entire data object (equivalent to a path of `$`). |
| `JsonPath.taskToken` | reference the [Task Token](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token), used for integration patterns that need to run for a long time. |
| `JsonPath.executionId` | reference the Execution Id field of the context object. |
| `JsonPath.executionInput` | reference the Execution Input object of the context object. |
| `JsonPath.executionName` | reference the Execution Name field of the context object. |
| `JsonPath.executionRoleArn` | reference the Execution RoleArn field of the context object. |
| `JsonPath.executionStartTime` | reference the Execution StartTime field of the context object. |
| `JsonPath.stateEnteredTime` | reference the State EnteredTime field of the context object. |
| `JsonPath.stateName` | reference the State Name field of the context object. |
| `JsonPath.stateRetryCount` | reference the State RetryCount field of the context object. |
| `JsonPath.stateMachineId` | reference the StateMachine Id field of the context object. |
| `JsonPath.stateMachineName` | reference the StateMachine Name field of the context object. |

You can also call [intrinsic functions](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html) using the methods on `JsonPath`:

| Method | Purpose |
|--------|---------|
| `JsonPath.array(JsonPath.stringAt('$.Field'), ...)` | make an array from other elements. |
| `JsonPath.arrayPartition(JsonPath.listAt('$.inputArray'), 4)` | partition an array. |
| `JsonPath.arrayContains(JsonPath.listAt('$.inputArray'), 5)` | determine if a specific value is present in an array. |
| `JsonPath.arrayRange(1, 9, 2)` | create a new array containing a specific range of elements. |
| `JsonPath.arrayGetItem(JsonPath.listAt('$.inputArray'), 5)` | get a specified index's value in an array. |
| `JsonPath.arrayLength(JsonPath.listAt('$.inputArray'))` | get the length of an array. |
| `JsonPath.arrayUnique(JsonPath.listAt('$.inputArray'))` | remove duplicate values from an array. |
| `JsonPath.base64Encode(JsonPath.stringAt('$.input'))` | encode data based on MIME Base64 encoding scheme. |
| `JsonPath.base64Decode(JsonPath.stringAt('$.base64'))` | decode data based on MIME Base64 decoding scheme. |
| `JsonPath.hash(JsonPath.objectAt('$.Data'), JsonPath.stringAt('$.Algorithm'))` | calculate the hash value of a given input. |
| `JsonPath.jsonMerge(JsonPath.objectAt('$.Obj1'), JsonPath.objectAt('$.Obj2'))` | merge two JSON objects into a single object. |
| `JsonPath.stringToJson(JsonPath.stringAt('$.ObjStr'))` | parse a JSON string to an object |
| `JsonPath.jsonToString(JsonPath.objectAt('$.Obj'))` | stringify an object to a JSON string |
| `JsonPath.mathRandom(1, 999)` | return a random number. |
| `JsonPath.mathAdd(JsonPath.numberAt('$.value1'), JsonPath.numberAt('$.step'))` | return the sum of two numbers. |
| `JsonPath.stringSplit(JsonPath.stringAt('$.inputString'), JsonPath.stringAt('$.splitter'))` | split a string into an array of values. |
| `JsonPath.uuid()` | return a version 4 universally unique identifier (v4 UUID). |
| `JsonPath.format('The value is {}.', JsonPath.stringAt('$.Value'))` | insert elements into a format string. |

## Amazon States Language

This library comes with a set of classes that model the [Amazon States
Language](https://states-language.net/spec.html). The following State classes
are supported:

* [`Task`](#task)
* [`Pass`](#pass)
* [`Wait`](#wait)
* [`Choice`](#choice)
* [`Parallel`](#parallel)
* [`Succeed`](#succeed)
* [`Fail`](#fail)
* [`Map`](#map)
* [`Distributed Map`](#distributed-map)
* [`Custom State`](#custom-state)

An arbitrary JSON object (specified at execution start) is passed from state to
state and transformed during the execution of the workflow. For more
information, see the States Language spec.

### Task

A `Task` represents some work that needs to be done. Do not use the `Task` class directly.

Instead, use one of the classes in the `aws-cdk-lib/aws-stepfunctions-tasks` module,
which provide a much more ergonomic way to integrate with various AWS services.

### Pass

A `Pass` state passes its input to its output, without performing work.
Pass states are useful when constructing and debugging state machines.

The following example injects some fixed data into the state machine through
the `result` field. The `result` field will be added to the input and the result
will be passed as the state's output.

```go
// Makes the current JSON state { ..., "subObject": { "hello": "world" } }
pass := sfn.NewPass(this, jsii.String("Add Hello World"), &PassProps{
	Result: sfn.Result_FromObject(map[string]interface{}{
		"hello": jsii.String("world"),
	}),
	ResultPath: jsii.String("$.subObject"),
})

// Set the next state
nextState := sfn.NewPass(this, jsii.String("NextState"))
pass.Next(nextState)
```

The `Pass` state also supports passing key-value pairs as input. Values can
be static, or selected from the input with a path.

The following example filters the `greeting` field from the state input
and also injects a field called `otherData`.

```go
pass := sfn.NewPass(this, jsii.String("Filter input and inject data"), &PassProps{
	StateName: jsii.String("my-pass-state"),
	 // the custom state name for the Pass state, defaults to 'Filter input and inject data' as the state name
	Parameters: map[string]interface{}{
		 // input to the pass state
		"input": sfn.JsonPath_stringAt(jsii.String("$.input.greeting")),
		"otherData": jsii.String("some-extra-stuff"),
	},
})
```

The object specified in `parameters` will be the input of the `Pass` state.
Since neither `Result` nor `ResultPath` are supplied, the `Pass` state copies
its input through to its output.

Learn more about the [Pass state](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-pass-state.html)

### Wait

A `Wait` state waits for a given number of seconds, or until the current time
hits a particular time. The time to wait may be taken from the execution's JSON
state.

```go
// Wait until it's the time mentioned in the the state object's "triggerTime"
// field.
wait := sfn.NewWait(this, jsii.String("Wait For Trigger Time"), &WaitProps{
	Time: sfn.WaitTime_TimestampPath(jsii.String("$.triggerTime")),
})

// Set the next state
startTheWork := sfn.NewPass(this, jsii.String("StartTheWork"))
wait.Next(startTheWork)
```

### Choice

A `Choice` state can take a different path through the workflow based on the
values in the execution's JSON state:

```go
choice := sfn.NewChoice(this, jsii.String("Did it work?"))

// Add conditions with .when()
successState := sfn.NewPass(this, jsii.String("SuccessState"))
failureState := sfn.NewPass(this, jsii.String("FailureState"))
choice.When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("SUCCESS")), successState)
choice.When(sfn.Condition_NumberGreaterThan(jsii.String("$.attempts"), jsii.Number(5)), failureState)

// Use .otherwise() to indicate what should be done if none of the conditions match
tryAgainState := sfn.NewPass(this, jsii.String("TryAgainState"))
choice.Otherwise(tryAgainState)
```

If you want to temporarily branch your workflow based on a condition, but have
all branches come together and continuing as one (similar to how an `if ... then ... else` works in a programming language), use the `.afterwards()` method:

```go
choice := sfn.NewChoice(this, jsii.String("What color is it?"))
handleBlueItem := sfn.NewPass(this, jsii.String("HandleBlueItem"))
handleRedItem := sfn.NewPass(this, jsii.String("HandleRedItem"))
handleOtherItemColor := sfn.NewPass(this, jsii.String("HanldeOtherItemColor"))
choice.When(sfn.Condition_StringEquals(jsii.String("$.color"), jsii.String("BLUE")), handleBlueItem)
choice.When(sfn.Condition_StringEquals(jsii.String("$.color"), jsii.String("RED")), handleRedItem)
choice.Otherwise(handleOtherItemColor)

// Use .afterwards() to join all possible paths back together and continue
shipTheItem := sfn.NewPass(this, jsii.String("ShipTheItem"))
choice.Afterwards().Next(shipTheItem)
```

You can add comments to `Choice` states as well as conditions that use `choice.when`.

```go
choice := sfn.NewChoice(this, jsii.String("What color is it?"), &ChoiceProps{
	Comment: jsii.String("color comment"),
})
handleBlueItem := sfn.NewPass(this, jsii.String("HandleBlueItem"))
handleOtherItemColor := sfn.NewPass(this, jsii.String("HanldeOtherItemColor"))
choice.When(sfn.Condition_StringEquals(jsii.String("$.color"), jsii.String("BLUE")), handleBlueItem, &ChoiceTransitionOptions{
	Comment: jsii.String("blue item comment"),
})
choice.Otherwise(handleOtherItemColor)
```

If your `Choice` doesn't have an `otherwise()` and none of the conditions match
the JSON state, a `NoChoiceMatched` error will be thrown. Wrap the state machine
in a `Parallel` state if you want to catch and recover from this.

#### Available Conditions

see [step function comparison operators](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-choice-state.html#amazon-states-language-choice-state-rules)

* `Condition.isPresent` - matches if a json path is present
* `Condition.isNotPresent` - matches if a json path is not present
* `Condition.isString` - matches if a json path contains a string
* `Condition.isNotString` - matches if a json path is not a string
* `Condition.isNumeric` - matches if a json path is numeric
* `Condition.isNotNumeric` - matches if a json path is not numeric
* `Condition.isBoolean` - matches if a json path is boolean
* `Condition.isNotBoolean` - matches if a json path is not boolean
* `Condition.isTimestamp` - matches if a json path is a timestamp
* `Condition.isNotTimestamp` - matches if a json path is not a timestamp
* `Condition.isNotNull` - matches if a json path is not null
* `Condition.isNull` - matches if a json path is null
* `Condition.booleanEquals` - matches if a boolean field has a given value
* `Condition.booleanEqualsJsonPath` - matches if a boolean field equals a value in a given mapping path
* `Condition.stringEqualsJsonPath` - matches if a string field equals a given mapping path
* `Condition.stringEquals` - matches if a field equals a string value
* `Condition.stringLessThan` - matches if a string field sorts before a given value
* `Condition.stringLessThanJsonPath` - matches if a string field sorts before a value at given mapping path
* `Condition.stringLessThanEquals` - matches if a string field sorts equal to or before a given value
* `Condition.stringLessThanEqualsJsonPath` - matches if a string field sorts equal to or before a given mapping
* `Condition.stringGreaterThan` - matches if a string field sorts after a given value
* `Condition.stringGreaterThanJsonPath` - matches if a string field sorts after a value at a given mapping path
* `Condition.stringGreaterThanEqualsJsonPath` - matches if a string field sorts after or equal to value at a given mapping path
* `Condition.stringGreaterThanEquals` - matches if a string field sorts after or equal to a given value
* `Condition.numberEquals` - matches if a numeric field has the given value
* `Condition.numberEqualsJsonPath` - matches if a numeric field has the value in a given mapping path
* `Condition.numberLessThan` - matches if a numeric field is less than the given value
* `Condition.numberLessThanJsonPath` - matches if a numeric field is less than the value at the given mapping path
* `Condition.numberLessThanEquals` - matches if a numeric field is less than or equal to the given value
* `Condition.numberLessThanEqualsJsonPath` - matches if a numeric field is less than or equal to the numeric value at given mapping path
* `Condition.numberGreaterThan` - matches if a numeric field is greater than the given value
* `Condition.numberGreaterThanJsonPath` - matches if a numeric field is greater than the value at a given mapping path
* `Condition.numberGreaterThanEquals` - matches if a numeric field is greater than or equal to the given value
* `Condition.numberGreaterThanEqualsJsonPath` - matches if a numeric field is greater than or equal to the value at a given mapping path
* `Condition.timestampEquals` - matches if a timestamp field is the same time as the given timestamp
* `Condition.timestampEqualsJsonPath` - matches if a timestamp field is the same time as the timestamp at a given mapping path
* `Condition.timestampLessThan` - matches if a timestamp field is before the given timestamp
* `Condition.timestampLessThanJsonPath` - matches if a timestamp field is before the timestamp at a given mapping path
* `Condition.timestampLessThanEquals` - matches if a timestamp field is before or equal to the given timestamp
* `Condition.timestampLessThanEqualsJsonPath` - matches if a timestamp field is before or equal to the timestamp at a given mapping path
* `Condition.timestampGreaterThan` - matches if a timestamp field is after the timestamp at a given mapping path
* `Condition.timestampGreaterThanJsonPath` - matches if a timestamp field is after the timestamp at a given mapping path
* `Condition.timestampGreaterThanEquals` - matches if a timestamp field is after or equal to the given timestamp
* `Condition.timestampGreaterThanEqualsJsonPath` - matches if a timestamp field is after or equal to the timestamp at a given mapping path
* `Condition.stringMatches` - matches if a field matches a string pattern that can contain a wild card (*) e.g: log-*.txt or *LATEST*. No other characters other than "*" have any special meaning - * can be escaped: \\*

### Parallel

A `Parallel` state executes one or more subworkflows in parallel. It can also
be used to catch and recover from errors in subworkflows.

```go
parallel := sfn.NewParallel(this, jsii.String("Do the work in parallel"))

// Add branches to be executed in parallel
shipItem := sfn.NewPass(this, jsii.String("ShipItem"))
sendInvoice := sfn.NewPass(this, jsii.String("SendInvoice"))
restock := sfn.NewPass(this, jsii.String("Restock"))
parallel.Branch(shipItem)
parallel.Branch(sendInvoice)
parallel.Branch(restock)

// Retry the whole workflow if something goes wrong with exponential backoff
parallel.AddRetry(&RetryProps{
	MaxAttempts: jsii.Number(1),
	MaxDelay: awscdk.Duration_Seconds(jsii.Number(5)),
	JitterStrategy: sfn.JitterType_FULL,
})

// How to recover from errors
sendFailureNotification := sfn.NewPass(this, jsii.String("SendFailureNotification"))
parallel.AddCatch(sendFailureNotification)

// What to do in case everything succeeded
closeOrder := sfn.NewPass(this, jsii.String("CloseOrder"))
parallel.Next(closeOrder)
```

### Succeed

Reaching a `Succeed` state terminates the state machine execution with a
successful status.

```go
success := sfn.NewSucceed(this, jsii.String("We did it!"))
```

### Fail

Reaching a `Fail` state terminates the state machine execution with a
failure status. The fail state should report the reason for the failure.
Failures can be caught by encompassing `Parallel` states.

```go
fail := sfn.NewFail(this, jsii.String("Fail"), &FailProps{
	Error: jsii.String("WorkflowFailure"),
	Cause: jsii.String("Something went wrong"),
})
```

The `Fail` state also supports returning dynamic values as the error and cause that are selected from the input with a path.

```go
fail := sfn.NewFail(this, jsii.String("Fail"), &FailProps{
	ErrorPath: sfn.JsonPath_StringAt(jsii.String("$.someError")),
	CausePath: sfn.JsonPath_*StringAt(jsii.String("$.someCause")),
})
```

### Map

A `Map` state can be used to run a set of steps for each element of an input array.
A `Map` state will execute the same steps for multiple entries of an array in the state input.

While the `Parallel` state executes multiple branches of steps using the same input, a `Map` state will
execute the same steps for multiple entries of an array in the state input.

```go
map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
	MaxConcurrency: jsii.Number(1),
	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
	ItemSelector: map[string]interface{}{
		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
	},
	ResultPath: jsii.String("$.mapOutput"),
})

// The Map iterator can contain a IChainable, which can be an individual or multiple steps chained together.
// Below example is with a Choice and Pass step
choice := sfn.NewChoice(this, jsii.String("Choice"))
condition1 := sfn.Condition_StringEquals(jsii.String("$.item.status"), jsii.String("SUCCESS"))
step1 := sfn.NewPass(this, jsii.String("Step1"))
step2 := sfn.NewPass(this, jsii.String("Step2"))
finish := sfn.NewPass(this, jsii.String("Finish"))

definition := choice.When(condition1, step1).Otherwise(step2).Afterwards().Next(finish)

map.ItemProcessor(definition)
```

To define a distributed `Map` state set `itemProcessors` mode to `ProcessorMode.DISTRIBUTED`.
An `executionType` must be specified for the distributed `Map` workflow.

```go
map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
	MaxConcurrency: jsii.Number(1),
	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
	ItemSelector: map[string]interface{}{
		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
	},
	ResultPath: jsii.String("$.mapOutput"),
})

map.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")), &ProcessorConfig{
	Mode: sfn.ProcessorMode_DISTRIBUTED,
	ExecutionType: sfn.ProcessorType_STANDARD,
})
```

> Visit [Using Map state in Distributed mode to orchestrate large-scale parallel workloads](https://docs.aws.amazon.com/step-functions/latest/dg/use-dist-map-orchestrate-large-scale-parallel-workloads.html) for more details.

### Distributed Map

Step Functions provides a high-concurrency mode for the Map state known as Distributed mode. In this mode, the Map state can accept input from large-scale Amazon S3 data sources. For example, your input can be a JSON or CSV file stored in an Amazon S3 bucket, or a JSON array passed from a previous step in the workflow. A Map state set to Distributed is known as a Distributed Map state. In this mode, the Map state runs each iteration as a child workflow execution, which enables high concurrency of up to 10,000 parallel child workflow executions. Each child workflow execution has its own, separate execution history from that of the parent workflow.

Use the Map state in Distributed mode when you need to orchestrate large-scale parallel workloads that meet any combination of the following conditions:

* The size of your dataset exceeds 256 KB.
* The workflow's execution event history exceeds 25,000 entries.
* You need a concurrency of more than 40 parallel iterations.

A `DistributedMap` state can be used to run a set of steps for each element of an input array with high concurrency.
A `DistributedMap` state will execute the same steps for multiple entries of an array in the state input or from S3 objects.

```go
distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
	MaxConcurrency: jsii.Number(1),
	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
})
distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
```

Map states in Distributed mode support multiple sources for an array to iterate:

* JSON array from the state input payload
* objects in an S3 bucket and optional prefix
* JSON array in a JSON file stored in S3
* CSV file stored in S3
* S3 inventory manifest stored in S3

There are multiple classes that implement `IItemReader` that can be used to configure the iterator source.  These can be provided via the optional `itemReader` property.  The default behavior if `itemReader` is omitted is to use the input payload.

Map states in Distributed mode also support writing results of the iterator to an S3 bucket and optional prefix.  Use a `ResultWriter` object provided via the optional `resultWriter` property to configure which S3 location iterator results will be written. The default behavior id `resultWriter` is omitted is to use the state output payload. However, if the iterator results are larger than the 256 kb limit for Step Functions payloads then the State Machine will fail.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"


// create a bucket
bucket := s3.NewBucket(this, jsii.String("Bucket"))

distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
	ItemReader: sfn.NewS3JsonItemReader(&S3FileItemReaderProps{
		Bucket: bucket,
		Key: jsii.String("my-key.json"),
	}),
	ResultWriter: sfn.NewResultWriter(&ResultWriterProps{
		Bucket: bucket,
		Prefix: jsii.String("my-prefix"),
	}),
})
distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
```

### Custom State

It's possible that the high-level constructs for the states or `stepfunctions-tasks` do not have
the states or service integrations you are looking for. The primary reasons for this lack of
functionality are:

* A [service integration](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-service-integrations.html) is available through Amazon States Language, but not available as construct
  classes in the CDK.
* The state or state properties are available through Step Functions, but are not configurable
  through constructs

If a feature is not available, a `CustomState` can be used to supply any Amazon States Language
JSON-based object as the state definition.

[Code Snippets](https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1) are available and can be plugged in as the state definition.

Custom states can be chained together with any of the other states to create your state machine
definition. You will also need to provide any permissions that are required to the `role` that
the State Machine uses.

The Retry and Catch fields are available for error handling.
You can configure the Retry field by defining it in the JSON object or by adding it using the `addRetry` method.
However, the Catch field cannot be configured by defining it in the JSON object, so it must be added using the `addCatch` method.

The following example uses the `DynamoDB` service integration to insert data into a DynamoDB table.

```go
import "github.com/aws/aws-cdk-go/awscdk"


// create a table
table := dynamodb.NewTable(this, jsii.String("montable"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
})

finalStatus := sfn.NewPass(this, jsii.String("final step"))

// States language JSON to put an item into DynamoDB
// snippet generated from https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1
stateJson := map[string]interface{}{
	"Type": jsii.String("Task"),
	"Resource": jsii.String("arn:aws:states:::dynamodb:putItem"),
	"Parameters": map[string]interface{}{
		"TableName": table.tableName,
		"Item": map[string]map[string]*string{
			"id": map[string]*string{
				"S": jsii.String("MyEntry"),
			},
		},
	},
	"ResultPath": nil,
}

// custom state which represents a task to insert data into DynamoDB
custom := sfn.NewCustomState(this, jsii.String("my custom task"), &CustomStateProps{
	StateJson: StateJson,
})

// catch errors with addCatch
errorHandler := sfn.NewPass(this, jsii.String("handle failure"))
custom.AddCatch(errorHandler)

// retry the task if something goes wrong
custom.AddRetry(&RetryProps{
	Errors: []*string{
		sfn.Errors_ALL(),
	},
	Interval: awscdk.Duration_Seconds(jsii.Number(10)),
	MaxAttempts: jsii.Number(5),
})

chain := sfn.Chain_Start(custom).Next(finalStatus)

sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(chain),
	Timeout: awscdk.Duration_*Seconds(jsii.Number(30)),
	Comment: jsii.String("a super cool state machine"),
})

// don't forget permissions. You need to assign them
table.GrantWriteData(sm)
```

## Task Chaining

To make defining work flows as convenient (and readable in a top-to-bottom way)
as writing regular programs, it is possible to chain most methods invocations.
In particular, the `.next()` method can be repeated. The result of a series of
`.next()` calls is called a **Chain**, and can be used when defining the jump
targets of `Choice.on` or `Parallel.branch`:

```go
step1 := sfn.NewPass(this, jsii.String("Step1"))
step2 := sfn.NewPass(this, jsii.String("Step2"))
step3 := sfn.NewPass(this, jsii.String("Step3"))
step4 := sfn.NewPass(this, jsii.String("Step4"))
step5 := sfn.NewPass(this, jsii.String("Step5"))
step6 := sfn.NewPass(this, jsii.String("Step6"))
step7 := sfn.NewPass(this, jsii.String("Step7"))
step8 := sfn.NewPass(this, jsii.String("Step8"))
step9 := sfn.NewPass(this, jsii.String("Step9"))
step10 := sfn.NewPass(this, jsii.String("Step10"))
choice := sfn.NewChoice(this, jsii.String("Choice"))
condition1 := sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("SUCCESS"))
parallel := sfn.NewParallel(this, jsii.String("Parallel"))
finish := sfn.NewPass(this, jsii.String("Finish"))

definition := step1.Next(step2).Next(choice.When(condition1, step3.Next(step4).Next(step5)).Otherwise(step6).Afterwards()).Next(parallel.Branch(step7.Next(step8)).Branch(step9.Next(step10))).Next(finish)

sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})
```

If you don't like the visual look of starting a chain directly off the first
step, you can use `Chain.start`:

```go
step1 := sfn.NewPass(this, jsii.String("Step1"))
step2 := sfn.NewPass(this, jsii.String("Step2"))
step3 := sfn.NewPass(this, jsii.String("Step3"))

definition := sfn.Chain_Start(step1).Next(step2).Next(step3)
```

## Task Credentials

Tasks are executed using the State Machine's execution role. In some cases, e.g. cross-account access, an IAM role can be assumed by the State Machine's execution role to provide access to the resource.
This can be achieved by providing the optional `credentials` property which allows using a fixed role or a json expression to resolve the role at runtime from the task's inputs.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var submitLambda function
var iamRole role


// use a fixed role for all task invocations
role := sfn.TaskRole_FromRole(iamRole)
// or use a json expression to resolve the role at runtime based on task inputs
//const role = sfn.TaskRole.fromRoleArnJsonPath('$.RoleArn');

submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &LambdaInvokeProps{
	LambdaFunction: submitLambda,
	OutputPath: jsii.String("$.Payload"),
	// use credentials
	Credentials: &Credentials{
		Role: *Role,
	},
})
```

See [the AWS documentation](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html)
to learn more about AWS Step Functions support for accessing resources in other AWS accounts.

## Service Integration Patterns

AWS Step functions integrate directly with other services, either through an optimised integration pattern, or through the AWS SDK.
Therefore, it is possible to change the `integrationPattern` of services, to enable additional functionality of the said AWS Service:

```go
import glue "github.com/aws/aws-cdk-go/awscdkgluealpha"

var submitGlue job


submitJob := tasks.NewGlueStartJobRun(this, jsii.String("Submit Job"), &GlueStartJobRunProps{
	GlueJobName: submitGlue.JobName,
	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
})
```

## State Machine Fragments

It is possible to define reusable (or abstracted) mini-state machines by
defining a construct that implements `IChainable`, which requires you to define
two fields:

* `startState: State`, representing the entry point into this state machine.
* `endStates: INextable[]`, representing the (one or more) states that outgoing
  transitions will be added to if you chain onto the fragment.

Since states will be named after their construct IDs, you may need to prefix the
IDs of states if you plan to instantiate the same state machine fragment
multiples times (otherwise all states in every instantiation would have the same
name).

The class `StateMachineFragment` contains some helper functions (like
`prefixStates()`) to make it easier for you to do this. If you define your state
machine as a subclass of this, it will be convenient to use:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdk"

type myJobProps struct {
	jobFlavor *string
}

type myJob struct {
	stateMachineFragment
	startState state
	endStates []iNextable
}

func newMyJob(parent construct, id *string, props myJobProps) *myJob {
	this := &myJob{}
	sfn.NewStateMachineFragment_Override(this, parent, id)

	choice := sfn.NewChoice(this, jsii.String("Choice")).When(sfn.Condition_StringEquals(jsii.String("$.branch"), jsii.String("left")), sfn.NewPass(this, jsii.String("Left Branch"))).When(sfn.Condition_StringEquals(jsii.String("$.branch"), jsii.String("right")), sfn.NewPass(this, jsii.String("Right Branch")))

	// ...

	this.startState = choice
	this.endStates = choice.Afterwards().EndStates
	return this
}

type myStack struct {
	stack
}

func newMyStack(scope construct, id *string) *myStack {
	this := &myStack{}
	newStack_Override(this, scope, id)
	// Do 3 different variants of MyJob in parallel
	parallel := sfn.NewParallel(this, jsii.String("All jobs")).Branch(NewMyJob(this, jsii.String("Quick"), &myJobProps{
		jobFlavor: jsii.String("quick"),
	}).PrefixStates()).Branch(NewMyJob(this, jsii.String("Medium"), &myJobProps{
		jobFlavor: jsii.String("medium"),
	}).PrefixStates()).Branch(NewMyJob(this, jsii.String("Slow"), &myJobProps{
		jobFlavor: jsii.String("slow"),
	}).PrefixStates())

	sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
		DefinitionBody: sfn.DefinitionBody_FromChainable(parallel),
	})
	return this
}
```

A few utility functions are available to parse state machine fragments.

* `State.findReachableStates`: Retrieve the list of states reachable from a given state.
* `State.findReachableEndStates`: Retrieve the list of end or terminal states reachable from a given state.

## Activity

**Activities** represent work that is done on some non-Lambda worker pool. The
Step Functions workflow will submit work to this Activity, and a worker pool
that you run yourself, probably on EC2, will pull jobs from the Activity and
submit the results of individual jobs back.

You need the ARN to do so, so if you use Activities be sure to pass the Activity
ARN into your worker pool:

```go
activity := sfn.NewActivity(this, jsii.String("Activity"))

// Read this CloudFormation Output from your application and use it to poll for work on
// the activity.
// Read this CloudFormation Output from your application and use it to poll for work on
// the activity.
awscdk.NewCfnOutput(this, jsii.String("ActivityArn"), &CfnOutputProps{
	Value: activity.ActivityArn,
})
```

### Activity-Level Permissions

Granting IAM permissions to an activity can be achieved by calling the `grant(principal, actions)` API:

```go
activity := sfn.NewActivity(this, jsii.String("Activity"))

role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

activity.Grant(role, jsii.String("states:SendTaskSuccess"))
```

This will grant the IAM principal the specified actions onto the activity.

## Metrics

`Task` object expose various metrics on the execution of that particular task. For example,
to create an alarm on a particular task failing:

```go
var task task

cloudwatch.NewAlarm(this, jsii.String("TaskAlarm"), &AlarmProps{
	Metric: task.metricFailed(),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
})
```

There are also metrics on the complete state machine:

```go
var stateMachine stateMachine

cloudwatch.NewAlarm(this, jsii.String("StateMachineAlarm"), &AlarmProps{
	Metric: stateMachine.metricFailed(),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
})
```

And there are metrics on the capacity of all state machines in your account:

```go
cloudwatch.NewAlarm(this, jsii.String("ThrottledAlarm"), &AlarmProps{
	Metric: sfn.StateTransitionMetric_MetricThrottledEvents(),
	Threshold: jsii.Number(10),
	EvaluationPeriods: jsii.Number(2),
})
```

## Error names

Step Functions identifies errors in the Amazon States Language using case-sensitive strings, known as error names.
The Amazon States Language defines a set of built-in strings that name well-known errors, all beginning with the `States.` prefix.

* `States.ALL` - A wildcard that matches any known error name.
* `States.Runtime` - An execution failed due to some exception that could not be processed. Often these are caused by errors at runtime, such as attempting to apply InputPath or OutputPath on a null JSON payload. A `States.Runtime` error is not retriable, and will always cause the execution to fail. A retry or catch on `States.ALL` will NOT catch States.Runtime errors.
* `States.DataLimitExceeded` - A States.DataLimitExceeded exception will be thrown for the following:

  * When the output of a connector is larger than payload size quota.
  * When the output of a state is larger than payload size quota.
  * When, after Parameters processing, the input of a state is larger than the payload size quota.
  * See [the AWS documentation](https://docs.aws.amazon.com/step-functions/latest/dg/limits-overview.html) to learn more about AWS Step Functions Quotas.
* `States.HeartbeatTimeout` - A Task state failed to send a heartbeat for a period longer than the HeartbeatSeconds value.
* `States.Timeout` - A Task state either ran longer than the TimeoutSeconds value, or failed to send a heartbeat for a period longer than the HeartbeatSeconds value.
* `States.TaskFailed`- A Task state failed during the execution. When used in a retry or catch, `States.TaskFailed` acts as a wildcard that matches any known error name except for `States.Timeout`.

## Logging

Enable logging to CloudWatch by passing a logging configuration with a
destination LogGroup:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

definition := sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))

sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
	Logs: &LogOptions{
		Destination: logGroup,
		Level: sfn.LogLevel_ALL,
	},
})
```

## X-Ray tracing

Enable X-Ray tracing for StateMachine:

```go
definition := sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))

sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
	TracingEnabled: jsii.Boolean(true),
})
```

See [the AWS documentation](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-xray-tracing.html)
to learn more about AWS Step Functions's X-Ray support.

## State Machine Permission Grants

IAM roles, users, or groups which need to be able to work with a State Machine should be granted IAM permissions.

Any object that implements the `IGrantable` interface (has an associated principal) can be granted permissions by calling:

* `stateMachine.grantStartExecution(principal)` - grants the principal the ability to execute the state machine
* `stateMachine.grantRead(principal)` - grants the principal read access
* `stateMachine.grantTaskResponse(principal)` - grants the principal the ability to send task tokens to the state machine
* `stateMachine.grantExecution(principal, actions)` - grants the principal execution-level permissions for the IAM actions specified
* `stateMachine.grant(principal, actions)` - grants the principal state-machine-level permissions for the IAM actions specified

### Start Execution Permission

Grant permission to start an execution of a state machine by calling the `grantStartExecution()` API.

```go
var definition iChainable
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})

// Give role permission to start execution of state machine
stateMachine.grantStartExecution(role)
```

The following permission is provided to a service principal by the `grantStartExecution()` API:

* `states:StartExecution` - to state machine

### Read Permissions

Grant `read` access to a state machine by calling the `grantRead()` API.

```go
var definition iChainable
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})

// Give role read access to state machine
stateMachine.grantRead(role)
```

The following read permissions are provided to a service principal by the `grantRead()` API:

* `states:ListExecutions` - to state machine
* `states:ListStateMachines` - to state machine
* `states:DescribeExecution` - to executions
* `states:DescribeStateMachineForExecution` - to executions
* `states:GetExecutionHistory` - to executions
* `states:ListActivities` - to `*`
* `states:DescribeStateMachine` - to `*`
* `states:DescribeActivity` - to `*`

### Task Response Permissions

Grant permission to allow task responses to a state machine by calling the `grantTaskResponse()` API:

```go
var definition iChainable
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})

// Give role task response permissions to the state machine
stateMachine.grantTaskResponse(role)
```

The following read permissions are provided to a service principal by the `grantRead()` API:

* `states:SendTaskSuccess` - to state machine
* `states:SendTaskFailure` - to state machine
* `states:SendTaskHeartbeat` - to state machine

### Execution-level Permissions

Grant execution-level permissions to a state machine by calling the `grantExecution()` API:

```go
var definition iChainable
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})

// Give role permission to get execution history of ALL executions for the state machine
stateMachine.grantExecution(role, jsii.String("states:GetExecutionHistory"))
```

### Custom Permissions

You can add any set of permissions to a state machine by calling the `grant()` API.

```go
var definition iChainable
user := iam.NewUser(this, jsii.String("MyUser"))
stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
})

//give user permission to send task success to the state machine
stateMachine.grant(user, jsii.String("states:SendTaskSuccess"))
```

## Import

Any Step Functions state machine that has been created outside the stack can be imported
into your CDK stack.

State machines can be imported by their ARN via the `StateMachine.fromStateMachineArn()` API.
In addition, the StateMachine can be imported via the `StateMachine.fromStateMachineName()` method, as long as they are in the same account/region as the current construct.

```go
app := awscdk.NewApp()
stack := awscdk.Newstack(app, jsii.String("MyStack"))
sfn.StateMachine_FromStateMachineArn(this, jsii.String("ViaArnImportedStateMachine"), jsii.String("arn:aws:states:us-east-1:123456789012:stateMachine:StateMachine2E01A3A5-N5TJppzoevKQ"))

sfn.StateMachine_FromStateMachineName(this, jsii.String("ViaResourceNameImportedStateMachine"), jsii.String("StateMachine2E01A3A5-N5TJppzoevKQ"))
```

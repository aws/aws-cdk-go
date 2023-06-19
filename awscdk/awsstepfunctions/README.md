# AWS Step Functions Construct Library

The `@aws-cdk/aws-stepfunctions` package contains constructs for building
serverless workflows using objects. Use this in conjunction with the
`@aws-cdk/aws-stepfunctions-tasks` package, which contains classes used
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
	// Lambda's result is in the attribute `Payload`
	OutputPath: jsii.String("$.Payload"),
})

waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &WaitProps{
	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
})

getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &LambdaInvokeProps{
	LambdaFunction: getStatusLambda,
	// Pass just the field named "guid" into the Lambda, put the
	// Lambda's result in a field called "status" in the response
	InputPath: jsii.String("$.guid"),
	OutputPath: jsii.String("$.Payload"),
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
	Definition: Definition,
	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
})
```

You can find more sample snippets and learn more about the service integrations
in the `@aws-cdk/aws-stepfunctions-tasks` package.

## State Machine

A `stepfunctions.StateMachine` is a resource that takes a state machine
definition. The definition is specified by its start state, and encompasses
all states reachable from the start state:

```go
startState := sfn.NewPass(this, jsii.String("StartState"))

sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	Definition: startState,
})
```

State machines execute using an IAM Role, which will automatically have all
permissions added that are required to make all state machine tasks execute
properly (for example, permissions to invoke any Lambda functions you add to
your workflow). A role will be created by default, but you can supply an
existing one as well.

## Accessing State (the JsonPath class)

Every State Machine execution has [State Machine
Data](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-data.html):
a JSON document containing keys and values that is fed into the state machine,
gets modified as the state machine progresses, and finally is produced as output.

You can pass fragments of this State Machine Data into Tasks of the state machine.
To do so, use the static methods on the `JsonPath` class. For example, to pass
the value that's in the data key of `OrderId` to a Lambda function as you invoke
it, use `JsonPath.stringAt('$.OrderId')`, like so:

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

You can also call [intrinsic functions](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html) using the methods on `JsonPath`:

| Method | Purpose |
|--------|---------|
| `JsonPath.array(JsonPath.stringAt('$.Field'), ...)` | make an array from other elements. |
| `JsonPath.format('The value is {}.', JsonPath.stringAt('$.Value'))` | insert elements into a format string. |
| `JsonPath.stringToJson(JsonPath.stringAt('$.ObjStr'))` | parse a JSON string to an object |
| `JsonPath.jsonToString(JsonPath.objectAt('$.Obj'))` | stringify an object to a JSON string |

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
* [`Custom State`](#custom-state)

An arbitrary JSON object (specified at execution start) is passed from state to
state and transformed during the execution of the workflow. For more
information, see the States Language spec.

### Task

A `Task` represents some work that needs to be done. The exact work to be
done is determine by a class that implements `IStepFunctionsTask`, a collection
of which can be found in the `@aws-cdk/aws-stepfunctions-tasks` module.

The tasks in the `@aws-cdk/aws-stepfunctions-tasks` module support the
[service integration pattern](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html) that integrates Step Functions with services
directly in the Amazon States language.

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

// Retry the whole workflow if something goes wrong
parallel.AddRetry(&RetryProps{
	MaxAttempts: jsii.Number(1),
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
success := sfn.NewFail(this, jsii.String("Fail"), &FailProps{
	Error: jsii.String("WorkflowFailure"),
	Cause: jsii.String("Something went wrong"),
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
})
map.Iterator(sfn.NewPass(this, jsii.String("Pass State")))
```

### Custom State

It's possible that the high-level constructs for the states or `stepfunctions-tasks` do not have
the states or service integrations you are looking for. The primary reasons for this lack of
functionality are:

* A [service integration](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-service-integrations.html) is available through Amazon States Langauge, but not available as construct
  classes in the CDK.
* The state or state properties are available through Step Functions, but are not configurable
  through constructs

If a feature is not available, a `CustomState` can be used to supply any Amazon States Language
JSON-based object as the state definition.

[Code Snippets](https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1) are available and can be plugged in as the state definition.

Custom states can be chained together with any of the other states to create your state machine
definition. You will also need to provide any permissions that are required to the `role` that
the State Machine uses.

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

chain := sfn.Chain_Start(custom).Next(finalStatus)

sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	Definition: chain,
	Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
})

// don't forget permissions. You need to assign them
table.grantWriteData(sm)
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
	Definition: Definition,
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
		Definition: parallel,
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
	Metric: task.MetricFailed(),
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

sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass"))),
	Logs: &LogOptions{
		Destination: logGroup,
		Level: sfn.LogLevel_ALL,
	},
})
```

## X-Ray tracing

Enable X-Ray tracing for StateMachine:

```go
sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass"))),
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
	Definition: Definition,
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
	Definition: Definition,
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
	Definition: Definition,
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
	Definition: Definition,
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
	Definition: Definition,
})

//give user permission to send task success to the state machine
stateMachine.grant(user, jsii.String("states:SendTaskSuccess"))
```

## Import

Any Step Functions state machine that has been created outside the stack can be imported
into your CDK stack.

State machines can be imported by their ARN via the `StateMachine.fromStateMachineArn()` API

```go
app := awscdk.NewApp()
stack := awscdk.Newstack(app, jsii.String("MyStack"))
sfn.StateMachine_FromStateMachineArn(stack, jsii.String("ImportedStateMachine"), jsii.String("arn:aws:states:us-east-1:123456789012:stateMachine:StateMachine2E01A3A5-N5TJppzoevKQ"))
```

# Triggers

Triggers allows you to execute code during deployments. This can be used for a
variety of use cases such as:

* Self tests: validate something after a resource/construct been provisioned
* Data priming: add initial data to resources after they are created
* Preconditions: check things such as account limits or external dependencies
  before deployment.

## Usage

The `TriggerFunction` construct will define an AWS Lambda function which is
triggered *during* deployment:

```go
import triggers "github.com/aws/aws-cdk-go/awscdk"


triggers.NewTriggerFunction(this, jsii.String("MyTrigger"), &TriggerFunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String(__dirname + "/my-trigger")),
})
```

In the above example, the AWS Lambda function defined in `MyTrigger` will
be invoked when the stack is deployed.

It is also possible to trigger a predefined Lambda function by using the `Trigger` construct:

```go
import "github.com/aws/aws-cdk-go/awscdk"


func := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Code: lambda.Code_FromInline(jsii.String("foo")),
})

triggers.NewTrigger(this, jsii.String("MyTrigger"), &TriggerProps{
	Handler: func,
	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
	InvocationType: triggers.InvocationType_EVENT,
})
```

Addition properties can be used to fine-tune the behaviour of the trigger.
The `timeout` property can be used to determine how long the invocation of the function should take.
The `invocationType` property can be used to change the invocation type of the function.
This might be useful in scenarios where a fire-and-forget strategy for invoking the function is sufficient.

## Trigger Failures

If the trigger handler fails (e.g. an exception is raised), the CloudFormation
deployment will fail, as if a resource failed to provision. This makes it easy
to implement "self tests" via triggers by simply making a set of assertions on
some provisioned infrastructure.

Note that this behavior is only applied when invocationType is `REQUEST_RESPONSE`. When invocationType is `EVENT`, Lambda function is invoked asynchronously.
In that case, if Lambda function is invoked successfully, the trigger will success regardless of the result for the function execution.

## Order of Execution

By default, a trigger will be executed by CloudFormation after the associated
handler is provisioned. This means that if the handler takes an implicit
dependency on other resources (e.g. via environment variables), those resources
will be provisioned *before* the trigger is executed.

In most cases, implicit ordering should be sufficient, but you can also use
`executeAfter` and `executeBefore` to control the order of execution.

The following example defines the following order: `(hello, world) => myTrigger => goodbye`.
The resources under `hello` and `world` will be provisioned in
parallel, and then the trigger `myTrigger` will be executed. Only then the
resources under `goodbye` will be provisioned:

```go
import triggers "github.com/aws/aws-cdk-go/awscdk"

var myTrigger trigger
var hello construct
var world construct
var goodbye construct


myTrigger.ExecuteAfter(hello, world)
myTrigger.ExecuteBefore(goodbye)
```

Note that `hello` and `world` are construct *scopes*. This means that they can
be specific resources (such as an `s3.Bucket` object) or groups of resources
composed together into constructs.

## Re-execution of Triggers

By default, `executeOnHandlerChange` is enabled. This implies that the trigger
is re-executed every time the handler function code or configuration changes. If
this option is disabled, the trigger will be executed only once upon first
deployment.

In the future we will consider adding support for additional re-execution modes:

* `executeOnEveryDeployment: boolean` - re-executes every time the stack is
  deployed (add random "salt" during synthesis).
* `executeOnResourceChange: Construct[]` - re-executes when one of the resources
  under the specified scopes has changed (add the hash the CloudFormation
  resource specs).

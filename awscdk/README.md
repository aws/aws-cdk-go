# AWS Cloud Development Kit Library

The AWS CDK construct library provides APIs to define your CDK application and add
CDK constructs to the application.

## Usage

### Upgrade from CDK 1.x

When upgrading from CDK 1.x, remove all dependencies to individual CDK packages
from your dependencies file and follow the rest of the sections.

### Installation

To use this package, you need to declare this package and the `constructs` package as
dependencies.

According to the kind of project you are developing:

* For projects that are CDK libraries, declare them both under the `devDependencies`
  **and** `peerDependencies` sections.
* For CDK apps, declare them under the `dependencies` section only.

### Use in your code

#### Classic import

You can use a classic import to get access to each service namespaces:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp()
stack := awscdk.Newstack(app, jsii.String("TestStack"))

awscdk.Aws_s3.NewBucket(stack, jsii.String("TestBucket"))
```

#### Barrel import

Alternatively, you can use "barrel" imports:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp()
stack := awscdk.Newstack(app, jsii.String("TestStack"))

awscdk.NewBucket(stack, jsii.String("TestBucket"))
```

<!--BEGIN CORE DOCUMENTATION-->

## Stacks and Stages

A `Stack` is the smallest physical unit of deployment, and maps directly onto
a CloudFormation Stack. You define a Stack by defining a subclass of `Stack`
-- let's call it `MyStack` -- and instantiating the constructs that make up
your application in `MyStack`'s constructor. You then instantiate this stack
one or more times to define different instances of your application. For example,
you can instantiate it once using few and cheap EC2 instances for testing,
and once again using more and bigger EC2 instances for production.

When your application grows, you may decide that it makes more sense to split it
out across multiple `Stack` classes. This can happen for a number of reasons:

* You could be starting to reach the maximum number of resources allowed in a single
  stack (this is currently 500).
* You could decide you want to separate out stateful resources and stateless resources
  into separate stacks, so that it becomes easy to tear down and recreate the stacks
  that don't have stateful resources.
* There could be a single stack with resources (like a VPC) that are shared
  between multiple instances of other stacks containing your applications.

As soon as your conceptual application starts to encompass multiple stacks,
it is convenient to wrap them in another construct that represents your
logical application. You can then treat that new unit the same way you used
to be able to treat a single stack: by instantiating it multiple times
for different instances of your application.

You can define a custom subclass of `Stage`, holding one or more
`Stack`s, to represent a single logical instance of your application.

As a final note: `Stack`s are not a unit of reuse. They describe physical
deployment layouts, and as such are best left to application builders to
organize their deployments with. If you want to vend a reusable construct,
define it as a subclasses of `Construct`: the consumers of your construct
will decide where to place it in their own stacks.

## Stack Synthesizers

Each Stack has a *synthesizer*, an object that determines how and where
the Stack should be synthesized and deployed. The synthesizer controls
aspects like:

* How does the stack reference assets? (Either through CloudFormation
  parameters the CLI supplies, or because the Stack knows a predefined
  location where assets will be uploaded).
* What roles are used to deploy the stack? These can be bootstrapped
  roles, roles created in some other way, or just the CLI's current
  credentials.

The following synthesizers are available:

* `DefaultStackSynthesizer`: recommended. Uses predefined asset locations and
  roles created by the modern bootstrap template. Access control is done by
  controlling who can assume the deploy role. This is the default stack
  synthesizer in CDKv2.
* `LegacyStackSynthesizer`: Uses CloudFormation parameters to communicate
  asset locations, and the CLI's current permissions to deploy stacks. The
  is the default stack synthesizer in CDKv1.
* `CliCredentialsStackSynthesizer`: Uses predefined asset locations, and the
  CLI's current permissions.

Each of these synthesizers takes configuration arguments. To configure
a stack with a synthesizer, pass it as one of its properties:

```go
// Example automatically generated from non-compiling source. May contain errors.
NewMyStack(app, jsii.String("MyStack"), &stackProps{
	synthesizer: awscdk.NewDefaultStackSynthesizer(&defaultStackSynthesizerProps{
		fileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
	}),
})
```

For more information on bootstrapping accounts and customizing synthesis,
see [Bootstrapping in the CDK Developer Guide](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html).

## Nested Stacks

[Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) are stacks created as part of other stacks. You create a nested stack within another stack by using the `NestedStack` construct.

As your infrastructure grows, common patterns can emerge in which you declare the same components in multiple templates. You can separate out these common components and create dedicated templates for them. Then use the resource in your template to reference other templates, creating nested stacks.

For example, assume that you have a load balancer configuration that you use for most of your stacks. Instead of copying and pasting the same configurations into your templates, you can create a dedicated template for the load balancer. Then, you just use the resource to reference that template from within other templates.

The following example will define a single top-level stack that contains two nested stacks: each one with a single Amazon S3 bucket:

```go
// Example automatically generated from non-compiling source. May contain errors.
type myNestedStack struct {
	nestedStack
}

func newMyNestedStack(scope construct, id *string, props nestedStackProps) *myNestedStack {
	this := &myNestedStack{}
	cfn.NewNestedStack_Override(this, scope, id, props)

	s3.NewBucket(this, jsii.String("NestedBucket"))
	return this
}

type myParentStack struct {
	stack
}

func newMyParentStack(scope construct, id *string, props stackProps) *myParentStack {
	this := &myParentStack{}
	newStack_Override(this, scope, id, props)

	NewMyNestedStack(this, jsii.String("Nested1"))
	NewMyNestedStack(this, jsii.String("Nested2"))
	return this
}
```

Resources references across nested/parent boundaries (even with multiple levels of nesting) will be wired by the AWS CDK
through CloudFormation parameters and outputs. When a resource from a parent stack is referenced by a nested stack,
a CloudFormation parameter will automatically be added to the nested stack and assigned from the parent; when a resource
from a nested stack is referenced by a parent stack, a CloudFormation output will be automatically be added to the
nested stack and referenced using `Fn::GetAtt "Outputs.Xxx"` from the parent.

Nested stacks also support the use of Docker image and file assets.

## Accessing resources in a different stack

You can access resources in a different stack, as long as they are in the
same account and AWS Region (see [next section](#accessing-resources-in-a-different-stack-and-region) for an exception).
The following example defines the stack `stack1`,
which defines an Amazon S3 bucket. Then it defines a second stack, `stack2`,
which takes the bucket from stack1 as a constructor property.

```go
// Example automatically generated from non-compiling source. May contain errors.
prod := map[string]*string{
	"account": jsii.String("123456789012"),
	"region": jsii.String("us-east-1"),
}

stack1 := NewStackThatProvidesABucket(app, jsii.String("Stack1"), &stackProps{
	env: prod,
})

// stack2 will take a property { bucket: IBucket }
stack2 := NewStackThatExpectsABucket(app, jsii.String("Stack2"), &stackThatExpectsABucketProps{
	bucket: stack1.bucket,
	env: prod,
})
```

If the AWS CDK determines that the resource is in the same account and
Region, but in a different stack, it automatically synthesizes AWS
CloudFormation
[Exports](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html)
in the producing stack and an
[Fn::ImportValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html)
in the consuming stack to transfer that information from one stack to the
other.

## Accessing resources in a different stack and region

> **This feature is currently experimental**

You can enable the Stack property `crossRegionReferences`
in order to access resources in a different stack *and* region. With this feature flag
enabled it is possible to do something like creating a CloudFront distribution in `us-east-2` and
an ACM certificate in `us-east-1`.

```go
// Example automatically generated from non-compiling source. May contain errors.
stack1 := awscdk.Newstack(app, jsii.String("Stack1"), &stackProps{
	env: &environment{
		region: jsii.String("us-east-1"),
	},
	crossRegionReferences: jsii.Boolean(true),
})
cert := acm.NewCertificate(stack1, jsii.String("Cert"), &certificateProps{
	domainName: jsii.String("*.example.com"),
	validation: acm.certificateValidation.fromDns(route53.publicHostedZone.fromHostedZoneId(stack1, jsii.String("Zone"), jsii.String("Z0329774B51CGXTDQV3X"))),
})

stack2 := awscdk.Newstack(app, jsii.String("Stack2"), &stackProps{
	env: &environment{
		region: jsii.String("us-east-2"),
	},
	crossRegionReferences: jsii.Boolean(true),
})
cloudfront.NewDistribution(stack2, jsii.String("Distribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("example.com")),
	},
	domainNames: []*string{
		jsii.String("dev.example.com"),
	},
	certificate: cert,
})
```

When the AWS CDK determines that the resource is in a different stack *and* is in a different
region, it will "export" the value by creating a custom resource in the producing stack which
creates SSM Parameters in the consuming region for each exported value. The parameters will be
created with the name '/cdk/exports/${consumingStackName}/${export-name}'.
In order to "import" the exports into the consuming stack a [SSM Dynamic reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-ssm)
is used to reference the SSM parameter which was created.

In order to mimic strong references, a Custom Resource is also created in the consuming
stack which marks the SSM parameters as being "imported". When a parameter has been successfully
imported, the producing stack cannot update the value.

See the [adr](https://github.com/aws/aws-cdk/blob/main/packages/@aws-cdk/core/adr/cross-region-stack-references)
for more details on this feature.

### Removing automatic cross-stack references

The automatic references created by CDK when you use resources across stacks
are convenient, but may block your deployments if you want to remove the
resources that are referenced in this way. You will see an error like:

```text
Export Stack1:ExportsOutputFnGetAtt-****** cannot be deleted as it is in use by Stack1
```

Let's say there is a Bucket in the `stack1`, and the `stack2` references its
`bucket.bucketName`. You now want to remove the bucket and run into the error above.

It's not safe to remove `stack1.bucket` while `stack2` is still using it, so
unblocking yourself from this is a two-step process. This is how it works:

DEPLOYMENT 1: break the relationship

* Make sure `stack2` no longer references `bucket.bucketName` (maybe the consumer
  stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
  remove the Lambda Function altogether).
* In the `stack1` class, call `this.exportValue(this.bucket.bucketName)`. This
  will make sure the CloudFormation Export continues to exist while the relationship
  between the two stacks is being broken.
* Deploy (this will effectively only change the `stack2`, but it's safe to deploy both).

DEPLOYMENT 2: remove the resource

* You are now free to remove the `bucket` resource from `stack1`.
* Don't forget to remove the `exportValue()` call as well.
* Deploy again (this time only the `stack1` will be changed -- the bucket will be deleted).

## Durations

To make specifications of time intervals unambiguous, a single class called
`Duration` is used throughout the AWS Construct Library by all constructs
that that take a time interval as a parameter (be it for a timeout, a
rate, or something else).

An instance of Duration is constructed by using one of the static factory
methods on it:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.Duration.seconds(jsii.Number(300)) // 5 minutes
awscdk.Duration.minutes(jsii.Number(5)) // 5 minutes
awscdk.Duration.hours(jsii.Number(1)) // 1 hour
awscdk.Duration.days(jsii.Number(7)) // 7 days
awscdk.Duration.parse(jsii.String("PT5M"))
```

Durations can be added or subtracted together:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.Duration.minutes(jsii.Number(1)).plus(awscdk.Duration.seconds(jsii.Number(60))) // 2 minutes
awscdk.Duration.minutes(jsii.Number(5)).minus(awscdk.Duration.seconds(jsii.Number(10)))
```

## Size (Digital Information Quantity)

To make specification of digital storage quantities unambiguous, a class called
`Size` is available.

An instance of `Size` is initialized through one of its static factory methods:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.Size.kibibytes(jsii.Number(200)) // 200 KiB
awscdk.Size.mebibytes(jsii.Number(5)) // 5 MiB
awscdk.Size.gibibytes(jsii.Number(40)) // 40 GiB
awscdk.Size.tebibytes(jsii.Number(200)) // 200 TiB
awscdk.Size.pebibytes(jsii.Number(3))
```

Instances of `Size` created with one of the units can be converted into others.
By default, conversion to a higher unit will fail if the conversion does not produce
a whole number. This can be overridden by unsetting `integral` property.

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.Size.mebibytes(jsii.Number(2)).toKibibytes() // yields 2048
awscdk.Size.kibibytes(jsii.Number(2050)).toMebibytes(&sizeConversionOptions{
	rounding: awscdk.SizeRoundingBehavior_FLOOR,
})
```

## Secrets

To help avoid accidental storage of secrets as plain text, we use the `SecretValue` type to
represent secrets. Any construct that takes a value that should be a secret (such as
a password or an access key) will take a parameter of type `SecretValue`.

The best practice is to store secrets in AWS Secrets Manager and reference them using `SecretValue.secretsManager`:

```go
// Example automatically generated from non-compiling source. May contain errors.
secret := awscdk.SecretValue.secretsManager(jsii.String("secretId"), &secretsManagerSecretOptions{
	jsonField: jsii.String("password"),
	 // optional: key of a JSON field to retrieve (defaults to all content),
	versionId: jsii.String("id"),
	 // optional: id of the version (default AWSCURRENT)
	versionStage: jsii.String("stage"),
})
```

Using AWS Secrets Manager is the recommended way to reference secrets in a CDK app.
`SecretValue` also supports the following secret sources:

* `SecretValue.unsafePlainText(secret)`: stores the secret as plain text in your app and the resulting template (not recommended).
* `SecretValue.secretsManager(secret)`: refers to a secret stored in Secrets Manager
* `SecretValue.ssmSecure(param, version)`: refers to a secret stored as a SecureString in the SSM
  Parameter Store. If you don't specify the exact version, AWS CloudFormation uses the latest
  version of the parameter.
* `SecretValue.cfnParameter(param)`: refers to a secret passed through a CloudFormation parameter (must have `NoEcho: true`).
* `SecretValue.cfnDynamicReference(dynref)`: refers to a secret described by a CloudFormation dynamic reference (used by `ssmSecure` and `secretsManager`).
* `SecretValue.resourceAttribute(attr)`: refers to a secret returned from a CloudFormation resource creation.

`SecretValue`s should only be passed to constructs that accept properties of type
`SecretValue`. These constructs are written to ensure your secrets will not be
exposed where they shouldn't be. If you try to use a `SecretValue` in a
different location, an error about unsafe secret usage will be thrown at
synthesis time.

If you rotate the secret's value in Secrets Manager, you must also change at
least one property on the resource where you are using the secret, to force
CloudFormation to re-read the secret.

`SecretValue.ssmSecure()` is only supported for a limited set of resources.
[Click here for a list of supported resources and properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#template-parameters-dynamic-patterns-resources).

## ARN manipulation

Sometimes you will need to put together or pick apart Amazon Resource Names
(ARNs). The functions `stack.formatArn()` and `stack.parseArn()` exist for
this purpose.

`formatArn()` can be used to build an ARN from components. It will automatically
use the region and account of the stack you're calling it on:

```go
// Example automatically generated from non-compiling source. May contain errors.
var stack stack


// Builds "arn:<PARTITION>:lambda:<REGION>:<ACCOUNT>:function:MyFunction"
stack.formatArn(&arnComponents{
	service: jsii.String("lambda"),
	resource: jsii.String("function"),
	sep: jsii.String(":"),
	resourceName: jsii.String("MyFunction"),
})
```

`parseArn()` can be used to get a single component from an ARN. `parseArn()`
will correctly deal with both literal ARNs and deploy-time values (tokens),
but in case of a deploy-time value be aware that the result will be another
deploy-time value which cannot be inspected in the CDK application.

```go
// Example automatically generated from non-compiling source. May contain errors.
var stack stack


// Extracts the function name out of an AWS Lambda Function ARN
arnComponents := stack.parseArn(arn, jsii.String(":"))
functionName := arnComponents.resourceName
```

Note that depending on the service, the resource separator can be either
`:` or `/`, and the resource name can be either the 6th or 7th
component in the ARN. When using these functions, you will need to know
the format of the ARN you are dealing with.

For an exhaustive list of ARN formats used in AWS, see [AWS ARNs and
Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
in the AWS General Reference.

## Dependencies

### Construct Dependencies

Sometimes AWS resources depend on other resources, and the creation of one
resource must be completed before the next one can be started.

In general, CloudFormation will correctly infer the dependency relationship
between resources based on the property values that are used. In the cases where
it doesn't, the AWS Construct Library will add the dependency relationship for
you.

If you need to add an ordering dependency that is not automatically inferred,
you do so by adding a dependency relationship using
`constructA.node.addDependency(constructB)`. This will add a dependency
relationship between all resources in the scope of `constructA` and all
resources in the scope of `constructB`.

If you want a single object to represent a set of constructs that are not
necessarily in the same scope, you can use a `DependencyGroup`. The
following creates a single object that represents a dependency on two
constructs, `constructB` and `constructC`:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Declare the dependable object
bAndC := constructs.NewDependencyGroup()
bAndC.add(constructB)
bAndC.add(constructC)

// Take the dependency
constructA.node.addDependency(bAndC)
```

### Stack Dependencies

Two different stack instances can have a dependency on one another. This
happens when an resource from one stack is referenced in another stack. In
that case, CDK records the cross-stack referencing of resources,
automatically produces the right CloudFormation primitives, and adds a
dependency between the two stacks. You can also manually add a dependency
between two stacks by using the `stackA.addDependency(stackB)` method.

A stack dependency has the following implications:

* Cyclic dependencies are not allowed, so if `stackA` is using resources from
  `stackB`, the reverse is not possible anymore.
* Stacks with dependencies between them are treated specially by the CDK
  toolkit:

  * If `stackA` depends on `stackB`, running `cdk deploy stackA` will also
    automatically deploy `stackB`.
  * `stackB`'s deployment will be performed *before* `stackA`'s deployment.

### CfnResource Dependencies

To make declaring dependencies between `CfnResource` objects easier, you can declare dependencies from one `CfnResource` object on another by using the `cfnResource1.addDependency(cfnResource2)` method. This method will work for resources both within the same stack and across stacks as it detects the relative location of the two resources and adds the dependency either to the resource or between the relevant stacks, as appropriate. If more complex logic is in needed, you can similarly remove, replace, or view dependencies between `CfnResource` objects with the `CfnResource` `removeDependency`, `replaceDependency`, and `obtainDependencies` methods, respectively.

## Custom Resources

Custom Resources are CloudFormation resources that are implemented by arbitrary
user code. They can do arbitrary lookups or modifications during a
CloudFormation deployment.

Custom resources are backed by *custom resource providers*. Commonly, these are
Lambda Functions that are deployed in the same deployment as the one that
defines the custom resource itself, but they can also be backed by Lambda
Functions deployed previously, or code responding to SNS Topic events running on
EC2 instances in a completely different account. For more information on custom
resource providers, see the next section.

Once you have a provider, each definition of a `CustomResource` construct
represents one invocation. A single provider can be used for the implementation
of arbitrarily many custom resource definitions. A single definition looks like
this:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCustomResource(this, jsii.String("MyMagicalResource"), &customResourceProps{
	resourceType: jsii.String("Custom::MyCustomResource"),
	 // must start with 'Custom::'

	// the resource properties
	properties: map[string]interface{}{
		"Property1": jsii.String("foo"),
		"Property2": jsii.String("bar"),
	},

	// the ARN of the provider (SNS/Lambda) which handles
	// CREATE, UPDATE or DELETE events for this resource type
	// see next section for details
	serviceToken: jsii.String("ARN"),
})
```

### Custom Resource Providers

Custom resources are backed by a **custom resource provider** which can be
implemented in one of the following ways. The following table compares the
various provider types (ordered from low-level to high-level):

| Provider                                                             | Compute Type | Error Handling | Submit to CloudFormation | Max Timeout     | Language | Footprint |
|----------------------------------------------------------------------|:------------:|:--------------:|:------------------------:|:---------------:|:--------:|:---------:|
| [sns.Topic](#amazon-sns-topic)                                       | Self-managed | Manual         | Manual                   | Unlimited       | Any      | Depends   |
| [lambda.Function](#aws-lambda-function)                              | AWS Lambda   | Manual         | Manual                   | 15min           | Any      | Small     |
| [core.CustomResourceProvider](#the-corecustomresourceprovider-class) | AWS Lambda   | Auto           | Auto                     | 15min           | Node.js  | Small     |
| [custom-resources.Provider](#the-custom-resource-provider-framework) | AWS Lambda   | Auto           | Auto                     | Unlimited Async | Any      | Large     |

Legend:

* **Compute type**: which type of compute can is used to execute the handler.
* **Error Handling**: whether errors thrown by handler code are automatically
  trapped and a FAILED response is submitted to CloudFormation. If this is
  "Manual", developers must take care of trapping errors. Otherwise, events
  could cause stacks to hang.
* **Submit to CloudFormation**: whether the framework takes care of submitting
  SUCCESS/FAILED responses to CloudFormation through the event's response URL.
* **Max Timeout**: maximum allows/possible timeout.
* **Language**: which programming languages can be used to implement handlers.
* **Footprint**: how many resources are used by the provider framework itself.

**A NOTE ABOUT SINGLETONS**

When defining resources for a custom resource provider, you will likely want to
define them as a *stack singleton* so that only a single instance of the
provider is created in your stack and which is used by all custom resources of
that type.

Here is a basic pattern for defining stack singletons in the CDK. The following
examples ensures that only a single SNS topic is defined:

```go
// Example automatically generated from non-compiling source. May contain errors.
func getOrCreate(scope *construct) topic {
	stack := awscdk.stack.of(*scope)
	uniqueid := "GloballyUniqueIdForSingleton" // For example, a UUID from `uuidgen`
	existing := stack.node.tryFindChild(uniqueid)
	if existing {
		return existing.(topic)
	}
	return sns.NewTopic(stack, uniqueid)
}
```

#### Amazon SNS Topic

Every time a resource event occurs (CREATE/UPDATE/DELETE), an SNS notification
is sent to the SNS topic. Users must process these notifications (e.g. through a
fleet of worker hosts) and submit success/failure responses to the
CloudFormation service.

> You only need to use this type of provider if your custom resource cannot run on AWS Lambda, for reasons other than the 15
> minute timeout. If you are considering using this type of provider because you want to write a custom resource provider that may need
> to wait for more than 15 minutes for the API calls to stabilize, have a look at the [`custom-resources`](#the-custom-resource-provider-framework) module first.
>
> Refer to the [CloudFormation Custom Resource documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) for information on the contract your custom resource needs to adhere to.

Set `serviceToken` to `topic.topicArn`  in order to use this provider:

```go
// Example automatically generated from non-compiling source. May contain errors.
topic := sns.NewTopic(this, jsii.String("MyProvider"))

awscdk.NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
	serviceToken: topic.topicArn,
})
```

#### AWS Lambda Function

An AWS lambda function is called *directly* by CloudFormation for all resource
events. The handler must take care of explicitly submitting a success/failure
response to the CloudFormation service and handle various error cases.

> **We do not recommend you use this provider type.** The CDK has wrappers around Lambda Functions that make them easier to work with.
>
> If you do want to use this provider, refer to the [CloudFormation Custom Resource documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) for information on the contract your custom resource needs to adhere to.

Set `serviceToken` to `lambda.functionArn` to use this provider:

```go
// Example automatically generated from non-compiling source. May contain errors.
fn := lambda.NewFunction(this, jsii.String("MyProvider"), functionProps)

awscdk.NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
	serviceToken: fn.functionArn,
})
```

#### The `core.CustomResourceProvider` class

The class [`@aws-cdk/core.CustomResourceProvider`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_core.CustomResourceProvider.html) offers a basic low-level
framework designed to implement simple and slim custom resource providers. It
currently only supports Node.js-based user handlers, represents permissions as raw
JSON blobs instead of `iam.PolicyStatement` objects, and it does not have
support for asynchronous waiting (handler cannot exceed the 15min lambda
timeout).

> **As an application builder, we do not recommend you use this provider type.** This provider exists purely for custom resources that are part of the AWS Construct Library.
>
> The [`custom-resources`](#the-custom-resource-provider-framework) provider is more convenient to work with and more fully-featured.

The provider has a built-in singleton method which uses the resource type as a
stack-unique identifier and returns the service token:

```go
// Example automatically generated from non-compiling source. May contain errors.
serviceToken := awscdk.CustomResourceProvider.getOrCreate(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
	description: jsii.String("Lambda function created by the custom resource provider"),
})

awscdk.NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
	resourceType: jsii.String("Custom::MyCustomResourceType"),
	serviceToken: serviceToken,
})
```

The directory (`my-handler` in the above example) must include an `index.js` file. It cannot import
external dependencies or files outside this directory. It must export an async
function named `handler`. This function accepts the CloudFormation resource
event object and returns an object with the following structure:

```js
exports.handler = async function(event) {
  const id = event.PhysicalResourceId; // only for "Update" and "Delete"
  const props = event.ResourceProperties;
  const oldProps = event.OldResourceProperties; // only for "Update"s

  switch (event.RequestType) {
    case "Create":
      // ...

    case "Update":
      // ...

      // if an error is thrown, a FAILED response will be submitted to CFN
      throw new Error('Failed!');

    case "Delete":
      // ...
  }

  return {
    // (optional) the value resolved from `resource.ref`
    // defaults to "event.PhysicalResourceId" or "event.RequestId"
    PhysicalResourceId: "REF",

    // (optional) calling `resource.getAtt("Att1")` on the custom resource in the CDK app
    // will return the value "BAR".
    Data: {
      Att1: "BAR",
      Att2: "BAZ"
    },

    // (optional) user-visible message
    Reason: "User-visible message",

    // (optional) hides values from the console
    NoEcho: true
  };
}
```

Here is an complete example of a custom resource that summarizes two numbers:

`sum-handler/index.js`:

```js
exports.handler = async (e) => {
  return {
    Data: {
      Result: e.ResourceProperties.lhs + e.ResourceProperties.rhs,
    },
  };
};
```

`sum.ts`:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws-samples/dummy/awscdkcore"

type sumProps struct {
	lhs *f64
	rhs *f64
}

type Sum struct {
	Construct
	result *f64
}result *f64

func NewSum(scope Construct, id *string, props sumProps) *Sum {
	this := &Sum{}
	newConstruct_Override(this, scope, id)

	resourceType := "Custom::Sum"
	serviceToken := awscdkcore.CustomResourceProvider.getOrCreate(this, resourceType, &customResourceProviderProps{
		codeDirectory: fmt.Sprintf("%v/sum-handler", __dirname),
		runtime: *awscdkcore.CustomResourceProviderRuntime_NODEJS_14_X,
	})

	resource := awscdkcore.NewCustomResource(this, jsii.String("Resource"), &customResourceProps{
		resourceType: resourceType,
		serviceToken: serviceToken,
		properties: map[string]interface{}{
			"lhs": props.lhs,
			"rhs": props.rhs,
		},
	})

	this.result = awscdkcore.Token.asNumber(resource.getAtt(jsii.String("Result")))
	return this
}
```

Usage will look like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
sum := NewSum(this, jsii.String("MySum"), &sumProps{
	lhs: jsii.Number(40),
	rhs: jsii.Number(2),
})
awscdk.NewCfnOutput(this, jsii.String("Result"), &cfnOutputProps{
	value: awscdk.Token.asString(sum.result),
})
```

To access the ARN of the provider's AWS Lambda function role, use the `getOrCreateProvider()`
built-in singleton method:

```go
// Example automatically generated from non-compiling source. May contain errors.
provider := awscdk.CustomResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
})

roleArn := provider.roleArn
```

This role ARN can then be used in resource-based IAM policies.

To add IAM policy statements to this role, use `addToRolePolicy()`:

```go
// Example automatically generated from non-compiling source. May contain errors.
provider := awscdk.CustomResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
})
provider.addToRolePolicy(map[string]*string{
	"Effect": jsii.String("Allow"),
	"Action": jsii.String("s3:GetObject"),
	"Resource": jsii.String("*"),
})
```

Note that `addToRolePolicy()` uses direct IAM JSON policy blobs, *not* a
`iam.PolicyStatement` object like you will see in the rest of the CDK.

#### The Custom Resource Provider Framework

The [`@aws-cdk/custom-resources`](https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html) module includes an advanced framework for
implementing custom resource providers.

Handlers are implemented as AWS Lambda functions, which means that they can be
implemented in any Lambda-supported runtime. Furthermore, this provider has an
asynchronous mode, which means that users can provide an `isComplete` lambda
function which is called periodically until the operation is complete. This
allows implementing providers that can take up to two hours to stabilize.

Set `serviceToken` to `provider.serviceToken` to use this type of provider:

```go
// Example automatically generated from non-compiling source. May contain errors.
provider := customresources.NewProvider(this, jsii.String("MyProvider"), &providerProps{
	onEventHandler: onEventHandler,
	isCompleteHandler: isCompleteHandler,
})

awscdk.NewCustomResource(this, jsii.String("MyResource"), &customResourceProps{
	serviceToken: provider.serviceToken,
})
```

See the [documentation](https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html) for more details.

## AWS CloudFormation features

A CDK stack synthesizes to an AWS CloudFormation Template. This section
explains how this module allows users to access low-level CloudFormation
features when needed.

### Stack Outputs

CloudFormation [stack outputs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) and exports are created using
the `CfnOutput` class:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCfnOutput(this, jsii.String("OutputName"), &cfnOutputProps{
	value: myBucket.bucketName,
	description: jsii.String("The name of an S3 bucket"),
	 // Optional
	exportName: jsii.String("TheAwesomeBucket"),
})
```

### Parameters

CloudFormation templates support the use of [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) to
customize a template. They enable CloudFormation users to input custom values to
a template each time a stack is created or updated. While the CDK design
philosophy favors using build-time parameterization, users may need to use
CloudFormation in a number of cases (for example, when migrating an existing
stack to the AWS CDK).

Template parameters can be added to a stack by using the `CfnParameter` class:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCfnParameter(this, jsii.String("MyParameter"), &cfnParameterProps{
	type: jsii.String("Number"),
	default: jsii.Number(1337),
})
```

The value of parameters can then be obtained using one of the `value` methods.
As parameters are only resolved at deployment time, the values obtained are
placeholder tokens for the real value (`Token.isUnresolved()` would return `true`
for those):

```go
// Example automatically generated from non-compiling source. May contain errors.
param := awscdk.NewCfnParameter(this, jsii.String("ParameterName"), &cfnParameterProps{
})

// If the parameter is a String
param.valueAsString

// If the parameter is a Number
param.valueAsNumber

// If the parameter is a List
param.valueAsList
```

### Pseudo Parameters

CloudFormation supports a number of [pseudo parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html),
which resolve to useful values at deployment time. CloudFormation pseudo
parameters can be obtained from static members of the `Aws` class.

It is generally recommended to access pseudo parameters from the scope's `stack`
instead, which guarantees the values produced are qualifying the designated
stack, which is essential in cases where resources are shared cross-stack:

```go
// Example automatically generated from non-compiling source. May contain errors.
// "this" is the current construct
stack := awscdk.stack.of(this)

stack.account // Returns the AWS::AccountId for this stack (or the literal value if known)
stack.region // Returns the AWS::Region for this stack (or the literal value if known)
stack.partition
```

### Resource Options

CloudFormation resources can also specify [resource
attributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-product-attribute-reference.html). The `CfnResource` class allows
accessing those through the `cfnOptions` property:

```go
// Example automatically generated from non-compiling source. May contain errors.
rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &cfnBucketProps{
})
// -or-
rawBucketAlt := myBucket.node.defaultChild.(cfnBucket)

// then
rawBucket.cfnOptions.condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &cfnConditionProps{
})
rawBucket.cfnOptions.metadata = map[string]interface{}{
	"metadataKey": jsii.String("MetadataValue"),
}
```

Resource dependencies (the `DependsOn` attribute) is modified using the
`cfnResource.addDependency` method:

```go
// Example automatically generated from non-compiling source. May contain errors.
resourceA := awscdk.NewCfnResource(this, jsii.String("ResourceA"), resourceProps)
resourceB := awscdk.NewCfnResource(this, jsii.String("ResourceB"), resourceProps)

resourceB.addDependency(resourceA)
```

### Intrinsic Functions and Condition Expressions

CloudFormation supports [intrinsic functions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html). These functions
can be accessed from the `Fn` class, which provides type-safe methods for each
intrinsic function as well as condition expressions:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myObjectOrArray interface{}
var myArray interface{}


// To use Fn::Base64
awscdk.Fn.base64(jsii.String("SGVsbG8gQ0RLIQo="))

// To compose condition expressions:
environmentParameter := awscdk.NewCfnParameter(this, jsii.String("Environment"))
awscdk.Fn.conditionAnd(awscdk.Fn.conditionEquals(jsii.String("Production"), environmentParameter), awscdk.Fn.conditionNot(awscdk.Fn.conditionEquals(jsii.String("us-east-1"), awscdk.Aws_REGION())))

// To use Fn::ToJsonString
awscdk.Fn.toJsonString(myObjectOrArray)

// To use Fn::Length
awscdk.Fn.len(awscdk.Fn.split(jsii.String(","), myArray))
```

When working with deploy-time values (those for which `Token.isUnresolved`
returns `true`), idiomatic conditionals from the programming language cannot be
used (the value will not be known until deployment time). When conditional logic
needs to be expressed with un-resolved values, it is necessary to use
CloudFormation conditions by means of the `CfnCondition` class:

```go
// Example automatically generated from non-compiling source. May contain errors.
environmentParameter := awscdk.NewCfnParameter(this, jsii.String("Environment"))
isProd := awscdk.NewCfnCondition(this, jsii.String("IsProduction"), &cfnConditionProps{
	expression: awscdk.Fn.conditionEquals(jsii.String("Production"), environmentParameter),
})

// Configuration value that is a different string based on IsProduction
stage := awscdk.Fn.conditionIf(isProd.logicalId, jsii.String("Beta"), jsii.String("Prod")).toString()

// Make Bucket creation condition to IsProduction by accessing
// and overriding the CloudFormation resource
bucket := s3.NewBucket(this, jsii.String("Bucket"))
cfnBucket := myBucket.node.defaultChild.(cfnBucket)
cfnBucket.cfnOptions.condition = isProd
```

### Mappings

CloudFormation [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html) are created and queried using the
`CfnMappings` class:

```go
// Example automatically generated from non-compiling source. May contain errors.
regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &cfnMappingProps{
	mapping: map[string]map[string]interface{}{
		"us-east-1": map[string]interface{}{
			"regionName": jsii.String("US East (N. Virginia)"),
		},
		"us-east-2": map[string]interface{}{
			"regionName": jsii.String("US East (Ohio)"),
		},
	},
})

regionTable.findInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
```

This will yield the following template:

```yaml
Mappings:
  RegionTable:
    us-east-1:
      regionName: US East (N. Virginia)
    us-east-2:
      regionName: US East (Ohio)
```

Mappings can also be synthesized "lazily"; lazy mappings will only render a "Mappings"
section in the synthesized CloudFormation template if some `findInMap` call is unable to
immediately return a concrete value due to one or both of the keys being unresolved tokens
(some value only available at deploy-time).

For example, the following code will not produce anything in the "Mappings" section. The
call to `findInMap` will be able to resolve the value during synthesis and simply return
`'US East (Ohio)'`.

```go
// Example automatically generated from non-compiling source. May contain errors.
regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &cfnMappingProps{
	mapping: map[string]map[string]interface{}{
		"us-east-1": map[string]interface{}{
			"regionName": jsii.String("US East (N. Virginia)"),
		},
		"us-east-2": map[string]interface{}{
			"regionName": jsii.String("US East (Ohio)"),
		},
	},
	lazy: jsii.Boolean(true),
})

regionTable.findInMap(jsii.String("us-east-2"), jsii.String("regionName"))
```

On the other hand, the following code will produce the "Mappings" section shown above,
since the top-level key is an unresolved token. The call to `findInMap` will return a token that resolves to
`{ "Fn::FindInMap": [ "RegionTable", { "Ref": "AWS::Region" }, "regionName" ] }`.

```go
// Example automatically generated from non-compiling source. May contain errors.
var regionTable cfnMapping


regionTable.findInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
```

### Dynamic References

CloudFormation supports [dynamically resolving](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html) values
for SSM parameters (including secure strings) and Secrets Manager. Encoding such
references is done using the `CfnDynamicReference` class:

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCfnDynamicReference(awscdk.CfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
```

### Template Options & Transform

CloudFormation templates support a number of options, including which Macros or
[Transforms](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html) to use when deploying the stack. Those can be
configured using the `stack.templateOptions` property:

```go
// Example automatically generated from non-compiling source. May contain errors.
stack := awscdk.Newstack(app, jsii.String("StackName"))

stack.templateOptions.description = "This will appear in the AWS console"
stack.templateOptions.transforms = []*string{
	"AWS::Serverless-2016-10-31",
}
stack.templateOptions.metadata = map[string]interface{}{
	"metadataKey": jsii.String("MetadataValue"),
}
```

### Emitting Raw Resources

The `CfnResource` class allows emitting arbitrary entries in the
[Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html) section of the CloudFormation template.

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCfnResource(this, jsii.String("ResourceId"), &cfnResourceProps{
	type: jsii.String("AWS::S3::Bucket"),
	properties: map[string]interface{}{
		"BucketName": jsii.String("bucket-name"),
	},
})
```

As for any other resource, the logical ID in the CloudFormation template will be
generated by the AWS CDK, but the type and properties will be copied verbatim in
the synthesized template.

### Including raw CloudFormation template fragments

When migrating a CloudFormation stack to the AWS CDK, it can be useful to
include fragments of an existing template verbatim in the synthesized template.
This can be achieved using the `CfnInclude` class.

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewCfnInclude(this, jsii.String("ID"), &cfnIncludeProps{
	template: map[string]map[string]map[string]interface{}{
		"Resources": map[string]map[string]interface{}{
			"Bucket": map[string]interface{}{
				"Type": jsii.String("AWS::S3::Bucket"),
				"Properties": map[string]*string{
					"BucketName": jsii.String("my-shiny-bucket"),
				},
			},
		},
	},
})
```

### Termination Protection

You can prevent a stack from being accidentally deleted by enabling termination
protection on the stack. If a user attempts to delete a stack with termination
protection enabled, the deletion fails and the stack--including its status--remains
unchanged. Enabling or disabling termination protection on a stack sets it for any
nested stacks belonging to that stack as well. You can enable termination protection
on a stack by setting the `terminationProtection` prop to `true`.

```go
// Example automatically generated from non-compiling source. May contain errors.
stack := awscdk.Newstack(app, jsii.String("StackName"), &stackProps{
	terminationProtection: jsii.Boolean(true),
})
```

By default, termination protection is disabled.

### Description

You can add a description of the stack in the same way as `StackProps`.

```go
// Example automatically generated from non-compiling source. May contain errors.
stack := awscdk.Newstack(app, jsii.String("StackName"), &stackProps{
	description: jsii.String("This is a description."),
})
```

### CfnJson

`CfnJson` allows you to postpone the resolution of a JSON blob from
deployment-time. This is useful in cases where the CloudFormation JSON template
cannot express a certain value.

A common example is to use `CfnJson` in order to render a JSON map which needs
to use intrinsic functions in keys. Since JSON map keys must be strings, it is
impossible to use intrinsics in keys and `CfnJson` can help.

The following example defines an IAM role which can only be assumed by
principals that are tagged with a specific tag.

```go
// Example automatically generated from non-compiling source. May contain errors.
tagParam := awscdk.NewCfnParameter(this, jsii.String("TagName"))

stringEquals := awscdk.NewCfnJson(this, jsii.String("ConditionJson"), &cfnJsonProps{
	value: map[string]*bool{
		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
	},
})

principal := iam.NewAccountRootPrincipal().withConditions(map[string]interface{}{
	"StringEquals": stringEquals,
})

iam.NewRole(this, jsii.String("MyRole"), &roleProps{
	assumedBy: principal,
})
```

**Explanation**: since in this example we pass the tag name through a parameter, it
can only be resolved during deployment. The resolved value can be represented in
the template through a `{ "Ref": "TagName" }`. However, since we want to use
this value inside a [`aws:PrincipalTag/TAG-NAME`](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-principaltag)
IAM operator, we need it in the *key* of a `StringEquals` condition. JSON keys
*must be* strings, so to circumvent this limitation, we use `CfnJson`
to "delay" the rendition of this template section to deploy-time. This means
that the value of `StringEquals` in the template will be `{ "Fn::GetAtt": [ "ConditionJson", "Value" ] }`, and will only "expand" to the operator we synthesized during deployment.

### Stack Resource Limit

When deploying to AWS CloudFormation, it needs to keep in check the amount of resources being added inside a Stack. Currently it's possible to check the limits in the [AWS CloudFormation quotas](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html) page.

It's possible to synthesize the project with more Resources than the allowed (or even reduce the number of Resources).

Set the context key `@aws-cdk/core:stackResourceLimit` with the proper value, being 0 for disable the limit of resources.

## App Context

[Context values](https://docs.aws.amazon.com/cdk/v2/guide/context.html) are key-value pairs that can be associated with an app, stack, or construct.
One common use case for context is to use it for enabling/disabling [feature flags](https://docs.aws.amazon.com/cdk/v2/guide/featureflags.html). There are several places
where context can be specified. They are listed below in the order they are evaluated (items at the
top take precedence over those below).

* The `node.setContext()` method
* The `postCliContext` prop when you create an `App`
* The CLI via the `--context` CLI argument
* The `cdk.json` file via the `context` key:
* The `cdk.context.json` file:
* The `~/.cdk.json` file via the `context` key:
* The `context` prop when you create an `App`

### Examples of setting context

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewApp(&appProps{
	context: map[string]interface{}{
		"@aws-cdk/core:newStyleStackSynthesis": jsii.Boolean(true),
	},
})
```

```go
// Example automatically generated from non-compiling source. May contain errors.
app := awscdk.NewApp()
app.node.setContext(jsii.String("@aws-cdk/core:newStyleStackSynthesis"), jsii.Boolean(true))
```

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdk.NewApp(&appProps{
	postCliContext: map[string]interface{}{
		"@aws-cdk/core:newStyleStackSynthesis": jsii.Boolean(true),
	},
})
```

```console
cdk synth --context @aws-cdk/core:newStyleStackSynthesis=true
```

*cdk.json*

```json
{
  "context": {
    "@aws-cdk/core:newStyleStackSynthesis": true
  }
}
```

*cdk.context.json*

```json
{
  "@aws-cdk/core:newStyleStackSynthesis": true
}
```

*~/.cdk.json*

```json
{
  "context": {
    "@aws-cdk/core:newStyleStackSynthesis": true
  }
}
```

## IAM Permissions Boundary

It is possible to apply an [IAM permissions boundary](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
to all roles within a specific construct scope. The most common use case would
be to apply a permissions boundary at the `Stage` level.

```go
// Example automatically generated from non-compiling source. May contain errors.
var app app


prodStage := awscdk.NewStage(app, jsii.String("ProdStage"), &stageProps{
	permissionsBoundary: permissionsBoundary_FromName(jsii.String("cdk-${Qualifier}-PermissionsBoundary")),
})
```

Any IAM Roles or Users created within this Stage will have the default
permissions boundary attached.

For more details see the [Permissions Boundary](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_iam-readme.html#permissions-boundaries) section in the IAM guide.

<!--END CORE DOCUMENTATION-->

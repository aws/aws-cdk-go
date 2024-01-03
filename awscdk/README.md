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

For projects that are CDK libraries in NPM, declare them both under the `devDependencies` **and** `peerDependencies` sections.
To make sure your library is compatible with the widest range of CDK versions: pick the minimum `aws-cdk-lib` version
that your library requires; declare a range dependency with a caret on that version in peerDependencies, and declare a
point version dependency on that version in devDependencies.

For example, let's say the minimum version your library needs is `2.38.0`. Your `package.json` should look like this:

```javascript
{
  "peerDependencies": {
    "aws-cdk-lib": "^2.38.0",
    "constructs": "^10.0.0"
  },
  "devDependencies": {
    /* Install the oldest version for testing so we don't accidentally use features from a newer version than we declare */
    "aws-cdk-lib": "2.38.0"
  }
}
```

For CDK apps, declare them under the `dependencies` section. Use a caret so you always get the latest version:

```json
{
  "dependencies": {
    "aws-cdk-lib": "^2.38.0",
    "constructs": "^10.0.0"
  }
}
```

### Use in your code

#### Classic import

You can use a classic import to get access to each service namespaces:

```go
import "github.com/aws/aws-cdk-go/awscdk"

app := awscdk.NewApp()
stack := awscdk.NewStack(app, jsii.String("TestStack"))

awscdk.Aws_s3.NewBucket(stack, jsii.String("TestBucket"))
```

#### Barrel import

Alternatively, you can use "barrel" imports:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := awscdk.NewApp()
stack := awscdk.NewStack(app, jsii.String("TestStack"))

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
  asset locations, and the CLI's current permissions to deploy stacks. This
  is the default stack synthesizer in CDKv1.
* `CliCredentialsStackSynthesizer`: Uses predefined asset locations, and the
  CLI's current permissions.

Each of these synthesizers takes configuration arguments. To configure
a stack with a synthesizer, pass it as one of its properties:

```go
NewMyStack(app, jsii.String("MyStack"), &stackProps{
	Synthesizer: awscdk.NewDefaultStackSynthesizer(&DefaultStackSynthesizerProps{
		FileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
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
prod := map[string]*string{
	"account": jsii.String("123456789012"),
	"region": jsii.String("us-east-1"),
}

stack1 := NewStackThatProvidesABucket(app, jsii.String("Stack1"), &stackProps{
	Env: prod,
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
stack1 := awscdk.Newstack(app, jsii.String("Stack1"), &stackProps{
	Env: &Environment{
		Region: jsii.String("us-east-1"),
	},
	CrossRegionReferences: jsii.Boolean(true),
})
cert := acm.NewCertificate(stack1, jsii.String("Cert"), &CertificateProps{
	DomainName: jsii.String("*.example.com"),
	Validation: acm.CertificateValidation_FromDns(route53.PublicHostedZone_FromHostedZoneId(stack1, jsii.String("Zone"), jsii.String("Z0329774B51CGXTDQV3X"))),
})

stack2 := awscdk.Newstack(app, jsii.String("Stack2"), &stackProps{
	Env: &Environment{
		Region: jsii.String("us-east-2"),
	},
	CrossRegionReferences: jsii.Boolean(true),
})
cloudfront.NewDistribution(stack2, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("example.com")),
	},
	DomainNames: []*string{
		jsii.String("dev.example.com"),
	},
	Certificate: cert,
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
awscdk.Duration_Seconds(jsii.Number(300)) // 5 minutes
awscdk.Duration_Minutes(jsii.Number(5)) // 5 minutes
awscdk.Duration_Hours(jsii.Number(1)) // 1 hour
awscdk.Duration_Days(jsii.Number(7)) // 7 days
awscdk.Duration_Parse(jsii.String("PT5M"))
```

Durations can be added or subtracted together:

```go
awscdk.Duration_Minutes(jsii.Number(1)).Plus(awscdk.Duration_Seconds(jsii.Number(60))) // 2 minutes
awscdk.Duration_Minutes(jsii.Number(5)).Minus(awscdk.Duration_Seconds(jsii.Number(10)))
```

## Size (Digital Information Quantity)

To make specification of digital storage quantities unambiguous, a class called
`Size` is available.

An instance of `Size` is initialized through one of its static factory methods:

```go
awscdk.Size_Kibibytes(jsii.Number(200)) // 200 KiB
awscdk.Size_Mebibytes(jsii.Number(5)) // 5 MiB
awscdk.Size_Gibibytes(jsii.Number(40)) // 40 GiB
awscdk.Size_Tebibytes(jsii.Number(200)) // 200 TiB
awscdk.Size_Pebibytes(jsii.Number(3))
```

Instances of `Size` created with one of the units can be converted into others.
By default, conversion to a higher unit will fail if the conversion does not produce
a whole number. This can be overridden by unsetting `integral` property.

```go
awscdk.Size_Mebibytes(jsii.Number(2)).ToKibibytes() // yields 2048
awscdk.Size_Kibibytes(jsii.Number(2050)).ToMebibytes(&SizeConversionOptions{
	Rounding: awscdk.SizeRoundingBehavior_FLOOR,
})
```

## Secrets

To help avoid accidental storage of secrets as plain text, we use the `SecretValue` type to
represent secrets. Any construct that takes a value that should be a secret (such as
a password or an access key) will take a parameter of type `SecretValue`.

The best practice is to store secrets in AWS Secrets Manager and reference them using `SecretValue.secretsManager`:

```go
secret := awscdk.SecretValue_SecretsManager(jsii.String("secretId"), &SecretsManagerSecretOptions{
	JsonField: jsii.String("password"),
	 // optional: key of a JSON field to retrieve (defaults to all content),
	VersionId: jsii.String("id"),
	 // optional: id of the version (default AWSCURRENT)
	VersionStage: jsii.String("stage"),
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
(ARNs). The functions `stack.formatArn()` and `stack.splitArn()` exist for
this purpose.

`formatArn()` can be used to build an ARN from components. It will automatically
use the region and account of the stack you're calling it on:

```go
var stack stack


// Builds "arn:<PARTITION>:lambda:<REGION>:<ACCOUNT>:function:MyFunction"
stack.FormatArn(&ArnComponents{
	Service: jsii.String("lambda"),
	Resource: jsii.String("function"),
	ArnFormat: awscdk.ArnFormat_COLON_RESOURCE_NAME,
	ResourceName: jsii.String("MyFunction"),
})
```

`splitArn()` can be used to get a single component from an ARN. `splitArn()`
will correctly deal with both literal ARNs and deploy-time values (tokens),
but in case of a deploy-time value be aware that the result will be another
deploy-time value which cannot be inspected in the CDK application.

```go
var stack stack


// Extracts the function name out of an AWS Lambda Function ARN
arnComponents := stack.SplitArn(arn, awscdk.ArnFormat_COLON_RESOURCE_NAME)
functionName := arnComponents.ResourceName
```

Note that the format of the resource separator depends on the service and
may be any of the values supported by `ArnFormat`. When dealing with these
functions, it is important to know the format of the ARN you are dealing with.

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
// Declare the dependable object
bAndC := constructs.NewDependencyGroup()
bAndC.Add(constructB)
bAndC.Add(constructC)

// Take the dependency
constructA.Node.AddDependency(bAndC)
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
awscdk.NewCustomResource(this, jsii.String("MyMagicalResource"), &CustomResourceProps{
	ResourceType: jsii.String("Custom::MyCustomResource"),
	 // must start with 'Custom::'

	// the resource properties
	Properties: map[string]interface{}{
		"Property1": jsii.String("foo"),
		"Property2": jsii.String("bar"),
	},

	// the ARN of the provider (SNS/Lambda) which handles
	// CREATE, UPDATE or DELETE events for this resource type
	// see next section for details
	ServiceToken: jsii.String("ARN"),
})
```

### Custom Resource Providers

Custom resources are backed by a **custom resource provider** which can be
implemented in one of the following ways. The following table compares the
various provider types (ordered from low-level to high-level):

| Provider                                                             | Compute Type | Error Handling | Submit to CloudFormation |   Max Timeout   | Language | Footprint |
| -------------------------------------------------------------------- | :----------: | :------------: | :----------------------: | :-------------: | :------: | :-------: |
| [sns.Topic](#amazon-sns-topic)                                       | Self-managed |     Manual     |          Manual          |    Unlimited    |   Any    |  Depends  |
| [lambda.Function](#aws-lambda-function)                              |  AWS Lambda  |     Manual     |          Manual          |      15min      |   Any    |   Small   |
| [core.CustomResourceProvider](#the-corecustomresourceprovider-class) |  AWS Lambda  |      Auto      |           Auto           |      15min      | Node.js  |   Small   |
| [custom-resources.Provider](#the-custom-resource-provider-framework) |  AWS Lambda  |      Auto      |           Auto           | Unlimited Async |   Any    |   Large   |

Legend:

* **Compute type**: which type of compute can be used to execute the handler.
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
func getOrCreate(scope *construct) topic {
	stack := awscdk.stack_Of(*scope)
	uniqueid := "GloballyUniqueIdForSingleton" // For example, a UUID from `uuidgen`
	existing := stack.Node.TryFindChild(uniqueid)
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
topic := sns.NewTopic(this, jsii.String("MyProvider"))

awscdk.NewCustomResource(this, jsii.String("MyResource"), &CustomResourceProps{
	ServiceToken: topic.TopicArn,
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
fn := lambda.NewSingletonFunction(this, jsii.String("MyProvider"), functionProps)

awscdk.NewCustomResource(this, jsii.String("MyResource"), &CustomResourceProps{
	ServiceToken: fn.FunctionArn,
})
```

#### The `core.CustomResourceProvider` class

The class [`@aws-cdk/core.CustomResourceProvider`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_core.CustomResourceProvider.html) offers a basic low-level
framework designed to implement simple and slim custom resource providers. It
currently only supports Node.js-based user handlers, represents permissions as raw
JSON blobs instead of `iam.PolicyStatement` objects, and it does not have
support for asynchronous waiting (handler cannot exceed the 15min lambda
timeout). The `CustomResourceProviderRuntime` supports runtime `nodejs12.x`,
`nodejs14.x`, `nodejs16.x`, `nodejs18.x`.

> **As an application builder, we do not recommend you use this provider type.** This provider exists purely for custom resources that are part of the AWS Construct Library.
>
> The [`custom-resources`](#the-custom-resource-provider-framework) provider is more convenient to work with and more fully-featured.

The provider has a built-in singleton method which uses the resource type as a
stack-unique identifier and returns the service token:

```go
serviceToken := awscdk.CustomResourceProvider_GetOrCreate(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
	Description: jsii.String("Lambda function created by the custom resource provider"),
})

awscdk.NewCustomResource(this, jsii.String("MyResource"), &CustomResourceProps{
	ResourceType: jsii.String("Custom::MyCustomResourceType"),
	ServiceToken: serviceToken,
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
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdk"

type sumProps struct {
	lhs *f64
	rhs *f64
}

type Sum struct {
	construct
	result *f64
}result *f64

func NewSum(scope construct, id *string, props sumProps) *Sum {
	this := &Sum{}
	newConstruct_Override(this, scope, id)

	resourceType := "Custom::Sum"
	serviceToken := awscdk.CustomResourceProvider_GetOrCreate(this, resourceType, &CustomResourceProviderProps{
		CodeDirectory: fmt.Sprintf("%v/sum-handler", __dirname),
		Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
	})

	resource := awscdk.NewCustomResource(this, jsii.String("Resource"), &CustomResourceProps{
		ResourceType: resourceType,
		ServiceToken: serviceToken,
		Properties: map[string]interface{}{
			"lhs": props.lhs,
			"rhs": props.rhs,
		},
	})

	this.result = awscdk.Token_AsNumber(resource.GetAtt(jsii.String("Result")))
	return this
}
```

Usage will look like this:

```go
sum := NewSum(this, jsii.String("MySum"), &sumProps{
	lhs: jsii.Number(40),
	rhs: jsii.Number(2),
})
awscdk.NewCfnOutput(this, jsii.String("Result"), &CfnOutputProps{
	Value: awscdk.Token_AsString(sum.result),
})
```

To access the ARN of the provider's AWS Lambda function role, use the `getOrCreateProvider()`
built-in singleton method:

```go
provider := awscdk.CustomResourceProvider_GetOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
})

roleArn := provider.RoleArn
```

This role ARN can then be used in resource-based IAM policies.

To add IAM policy statements to this role, use `addToRolePolicy()`:

```go
provider := awscdk.CustomResourceProvider_GetOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
})
provider.AddToRolePolicy(map[string]*string{
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
provider := customresources.NewProvider(this, jsii.String("MyProvider"), &ProviderProps{
	OnEventHandler: OnEventHandler,
	IsCompleteHandler: IsCompleteHandler,
})

awscdk.NewCustomResource(this, jsii.String("MyResource"), &CustomResourceProps{
	ServiceToken: provider.ServiceToken,
})
```

See the [documentation](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-cdk-lib.custom_resources-readme.html) for more details.

## AWS CloudFormation features

A CDK stack synthesizes to an AWS CloudFormation Template. This section
explains how this module allows users to access low-level CloudFormation
features when needed.

### Stack Outputs

CloudFormation [stack outputs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) and exports are created using
the `CfnOutput` class:

```go
awscdk.NewCfnOutput(this, jsii.String("OutputName"), &CfnOutputProps{
	Value: myBucket.BucketName,
	Description: jsii.String("The name of an S3 bucket"),
	 // Optional
	ExportName: jsii.String("TheAwesomeBucket"),
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
awscdk.NewCfnParameter(this, jsii.String("MyParameter"), &CfnParameterProps{
	Type: jsii.String("Number"),
	Default: jsii.Number(1337),
})
```

The value of parameters can then be obtained using one of the `value` methods.
As parameters are only resolved at deployment time, the values obtained are
placeholder tokens for the real value (`Token.isUnresolved()` would return `true`
for those):

```go
param := awscdk.NewCfnParameter(this, jsii.String("ParameterName"), &CfnParameterProps{
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
// "this" is the current construct
stack := awscdk.stack_Of(this)

stack.Account // Returns the AWS::AccountId for this stack (or the literal value if known)
stack.Region // Returns the AWS::Region for this stack (or the literal value if known)
stack.partition
```

### Resource Options

CloudFormation resources can also specify [resource
attributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-product-attribute-reference.html). The `CfnResource` class allows
accessing those through the `cfnOptions` property:

```go
rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &CfnBucketProps{
})
// -or-
rawBucketAlt := myBucket.Node.defaultChild.(cfnBucket)

// then
rawBucket.CfnOptions.Condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &CfnConditionProps{
})
rawBucket.CfnOptions.Metadata = map[string]interface{}{
	"metadataKey": jsii.String("MetadataValue"),
}
```

Resource dependencies (the `DependsOn` attribute) is modified using the
`cfnResource.addDependency` method:

```go
resourceA := awscdk.NewCfnResource(this, jsii.String("ResourceA"), resourceProps)
resourceB := awscdk.NewCfnResource(this, jsii.String("ResourceB"), resourceProps)

resourceB.AddDependency(resourceA)
```

#### CreationPolicy

Some resources support a [CreationPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-creationpolicy.html) to be specified as a CfnOption.

The creation policy is invoked only when AWS CloudFormation creates the associated resource. Currently, the only AWS CloudFormation resources that support creation policies are `CfnAutoScalingGroup`, `CfnInstance`, `CfnWaitCondition` and `CfnFleet`.

The `CfnFleet` resource from the `aws-appstream` module supports specifying `startFleet` as
a property of the creationPolicy on the resource options. Setting it to true will make AWS CloudFormation wait until the fleet is started before continuing with the creation of
resources that depend on the fleet resource.

```go
fleet := appstream.NewCfnFleet(this, jsii.String("Fleet"), &CfnFleetProps{
	InstanceType: jsii.String("stream.standard.small"),
	Name: jsii.String("Fleet"),
	ComputeCapacity: &ComputeCapacityProperty{
		DesiredInstances: jsii.Number(1),
	},
	ImageName: jsii.String("AppStream-AmazonLinux2-09-21-2022"),
})
fleet.CfnOptions.CreationPolicy = &CfnCreationPolicy{
	StartFleet: jsii.Boolean(true),
}
```

The properties passed to the level 2 constructs `AutoScalingGroup` and `Instance` from the
`aws-ec2` module abstract what is passed into the `CfnOption` properties `resourceSignal` and
`autoScalingCreationPolicy`, but when using level 1 constructs you can specify these yourself.

The CfnWaitCondition resource from the `aws-cloudformation` module suppports the `resourceSignal`.
The format of the timeout is `PT#H#M#S`. In the example below AWS Cloudformation will wait for
3 success signals to occur within 15 minutes before the status of the resource will be set to
`CREATE_COMPLETE`.

```go
var resource cfnResource


resource.CfnOptions.CreationPolicy = &CfnCreationPolicy{
	ResourceSignal: &CfnResourceSignal{
		Count: jsii.Number(3),
		Timeout: jsii.String("PR15M"),
	},
}
```

### Intrinsic Functions and Condition Expressions

CloudFormation supports [intrinsic functions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html). These functions
can be accessed from the `Fn` class, which provides type-safe methods for each
intrinsic function as well as condition expressions:

```go
var myObjectOrArray interface{}
var myArray interface{}


// To use Fn::Base64
awscdk.Fn_Base64(jsii.String("SGVsbG8gQ0RLIQo="))

// To compose condition expressions:
environmentParameter := awscdk.NewCfnParameter(this, jsii.String("Environment"))
awscdk.Fn_ConditionAnd(awscdk.Fn_ConditionEquals(jsii.String("Production"), environmentParameter), awscdk.Fn_ConditionNot(awscdk.Fn_ConditionEquals(jsii.String("us-east-1"), awscdk.Aws_REGION())))

// To use Fn::ToJsonString
awscdk.Fn_ToJsonString(myObjectOrArray)

// To use Fn::Length
awscdk.Fn_Len(awscdk.Fn_Split(jsii.String(","), myArray))
```

When working with deploy-time values (those for which `Token.isUnresolved`
returns `true`), idiomatic conditionals from the programming language cannot be
used (the value will not be known until deployment time). When conditional logic
needs to be expressed with un-resolved values, it is necessary to use
CloudFormation conditions by means of the `CfnCondition` class:

```go
environmentParameter := awscdk.NewCfnParameter(this, jsii.String("Environment"))
isProd := awscdk.NewCfnCondition(this, jsii.String("IsProduction"), &CfnConditionProps{
	Expression: awscdk.Fn_ConditionEquals(jsii.String("Production"), environmentParameter),
})

// Configuration value that is a different string based on IsProduction
stage := awscdk.Fn_ConditionIf(isProd.LogicalId, jsii.String("Beta"), jsii.String("Prod")).ToString()

// Make Bucket creation condition to IsProduction by accessing
// and overriding the CloudFormation resource
bucket := s3.NewBucket(this, jsii.String("Bucket"))
cfnBucket := myBucket.Node.defaultChild.(cfnBucket)
cfnBucket.CfnOptions.Condition = isProd
```

### Mappings

CloudFormation [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html) are created and queried using the
`CfnMappings` class:

```go
regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &CfnMappingProps{
	Mapping: map[string]map[string]interface{}{
		"us-east-1": map[string]interface{}{
			"regionName": jsii.String("US East (N. Virginia)"),
		},
		"us-east-2": map[string]interface{}{
			"regionName": jsii.String("US East (Ohio)"),
		},
	},
})

regionTable.FindInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
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
regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &CfnMappingProps{
	Mapping: map[string]map[string]interface{}{
		"us-east-1": map[string]interface{}{
			"regionName": jsii.String("US East (N. Virginia)"),
		},
		"us-east-2": map[string]interface{}{
			"regionName": jsii.String("US East (Ohio)"),
		},
	},
	Lazy: jsii.Boolean(true),
})

regionTable.FindInMap(jsii.String("us-east-2"), jsii.String("regionName"))
```

On the other hand, the following code will produce the "Mappings" section shown above,
since the top-level key is an unresolved token. The call to `findInMap` will return a token that resolves to
`{ "Fn::FindInMap": [ "RegionTable", { "Ref": "AWS::Region" }, "regionName" ] }`.

```go
var regionTable cfnMapping


regionTable.FindInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
```

An optional default value can also be passed to `findInMap`. If either key is not found in the map and the mapping is lazy, `findInMap` will return the default value and not render the mapping.
If the mapping is not lazy or either key is an unresolved token, the call to `findInMap` will return a token that resolves to
`{ "Fn::FindInMap": [ "MapName", "TopLevelKey", "SecondLevelKey", { "DefaultValue": "DefaultValue" } ] }`, and the mapping will be rendered.
Note that the `AWS::LanguageExtentions` transform is added to enable the default value functionality.

For example, the following code will again not produce anything in the "Mappings" section. The
call to `findInMap` will be able to resolve the value during synthesis and simply return
`'Region not found'`.

```go
regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &CfnMappingProps{
	Mapping: map[string]map[string]interface{}{
		"us-east-1": map[string]interface{}{
			"regionName": jsii.String("US East (N. Virginia)"),
		},
		"us-east-2": map[string]interface{}{
			"regionName": jsii.String("US East (Ohio)"),
		},
	},
	Lazy: jsii.Boolean(true),
})

regionTable.FindInMap(jsii.String("us-west-1"), jsii.String("regionName"), jsii.String("Region not found"))
```

### Dynamic References

CloudFormation supports [dynamically resolving](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html) values
for SSM parameters (including secure strings) and Secrets Manager. Encoding such
references is done using the `CfnDynamicReference` class:

```go
awscdk.NewCfnDynamicReference(awscdk.CfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
```

### Template Options & Transform

CloudFormation templates support a number of options, including which Macros or
[Transforms](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html) to use when deploying the stack. Those can be
configured using the `stack.templateOptions` property:

```go
stack := awscdk.Newstack(app, jsii.String("StackName"))

stack.TemplateOptions.Description = "This will appear in the AWS console"
stack.TemplateOptions.Transforms = []*string{
	"AWS::Serverless-2016-10-31",
}
stack.TemplateOptions.Metadata = map[string]interface{}{
	"metadataKey": jsii.String("MetadataValue"),
}
```

### Emitting Raw Resources

The `CfnResource` class allows emitting arbitrary entries in the
[Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html) section of the CloudFormation template.

```go
awscdk.NewCfnResource(this, jsii.String("ResourceId"), &cfnResourceProps{
	Type: jsii.String("AWS::S3::Bucket"),
	Properties: map[string]interface{}{
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
stack := awscdk.Newstack(app, jsii.String("StackName"), &stackProps{
	TerminationProtection: jsii.Boolean(true),
})
```

You can also set termination protection with the setter after you've instantiated the stack.

```go
stack := awscdk.Newstack(app, jsii.String("StackName"), &stackProps{
})
stack.terminationProtection = true
```

By default, termination protection is disabled.

### Description

You can add a description of the stack in the same way as `StackProps`.

```go
stack := awscdk.Newstack(app, jsii.String("StackName"), &stackProps{
	Description: jsii.String("This is a description."),
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
tagParam := awscdk.NewCfnParameter(this, jsii.String("TagName"))

stringEquals := awscdk.NewCfnJson(this, jsii.String("ConditionJson"), &CfnJsonProps{
	Value: map[string]*bool{
		fmt.Sprintf("aws:PrincipalTag/%v", tagParam.valueAsString): jsii.Boolean(true),
	},
})

principal := iam.NewAccountRootPrincipal().WithConditions(map[string]interface{}{
	"StringEquals": stringEquals,
})

iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	AssumedBy: principal,
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

### Template Indentation

The AWS CloudFormation templates generated by CDK include indentation by default.
Indentation makes the templates more readable, but also increases their size,
and CloudFormation templates cannot exceed 1MB.

It's possible to reduce the size of your templates by suppressing indentation.

To do this for all templates, set the context key `@aws-cdk/core:suppressTemplateIndentation` to `true`.

To do this for a specific stack, add a `suppressTemplateIndentation: true` property to the
stack's `StackProps` parameter. You can also set this property to `false` to override
the context key setting.

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
awscdk.NewApp(&AppProps{
	Context: map[string]interface{}{
		"@aws-cdk/core:newStyleStackSynthesis": jsii.Boolean(true),
	},
})
```

```go
app := awscdk.NewApp()
app.Node.SetContext(jsii.String("@aws-cdk/core:newStyleStackSynthesis"), jsii.Boolean(true))
```

```go
awscdk.NewApp(&AppProps{
	PostCliContext: map[string]interface{}{
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
prodStage := awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("cdk-${Qualifier}-PermissionsBoundary")),
})
```

Any IAM Roles or Users created within this Stage will have the default
permissions boundary attached.

For more details see the [Permissions Boundary](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_iam-readme.html#permissions-boundaries) section in the IAM guide.

## Policy Validation

If you or your organization use (or would like to use) any policy validation tool, such as
[CloudFormation
Guard](https://docs.aws.amazon.com/cfn-guard/latest/ug/what-is-guard.html) or
[OPA](https://www.openpolicyagent.org/), to define constraints on your
CloudFormation template, you can incorporate them into the CDK application.
By using the appropriate plugin, you can make the CDK application check the
generated CloudFormation templates against your policies immediately after
synthesis. If there are any violations, the synthesis will fail and a report
will be printed to the console or to a file (see below).

> **Note**
> This feature is considered experimental, and both the plugin API and the
> format of the validation report are subject to change in the future.

### For application developers

To use one or more validation plugins in your application, use the
`policyValidationBeta1` property of `Stage`:

```go
// globally for the entire app (an app is a stage)
app := awscdk.NewApp(&AppProps{
	PolicyValidationBeta1: []iPolicyValidationPluginBeta1{
		// These hypothetical classes implement IPolicyValidationPluginBeta1:
		NewThirdPartyPluginX(),
		NewThirdPartyPluginY(),
	},
})

// only apply to a particular stage
prodStage := awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
	PolicyValidationBeta1: []*iPolicyValidationPluginBeta1{
		NewThirdPartyPluginX(),
	},
})
```

Immediately after synthesis, all plugins registered this way will be invoked to
validate all the templates generated in the scope you defined. In particular, if
you register the templates in the `App` object, all templates will be subject to
validation.

> **Warning**
> Other than modifying the cloud assembly, plugins can do anything that your CDK
> application can. They can read data from the filesystem, access the network
> etc. It's your responsibility as the consumer of a plugin to verify that it is
> secure to use.

By default, the report will be printed in a human readable format. If you want a
report in JSON format, enable it using the `@aws-cdk/core:validationReportJson`
context passing it directly to the application:

```go
app := awscdk.NewApp(&AppProps{
	Context: map[string]interface{}{
		"@aws-cdk/core:validationReportJson": jsii.Boolean(true),
	},
})
```

Alternatively, you can set this context key-value pair using the `cdk.json` or
`cdk.context.json` files in your project directory (see
[Runtime context](https://docs.aws.amazon.com/cdk/v2/guide/context.html)).

If you choose the JSON format, the CDK will print the policy validation report
to a file called `policy-validation-report.json` in the cloud assembly
directory. For the default, human-readable format, the report will be printed to
the standard output.

### For plugin authors

The communication protocol between the CDK core module and your policy tool is
defined by the `IPolicyValidationPluginBeta1` interface. To create a new plugin you must
write a class that implements this interface. There are two things you need to
implement: the plugin name (by overriding the `name` property), and the
`validate()` method.

The framework will call `validate()`, passing an `IPolicyValidationContextBeta1` object.
The location of the templates to be validated is given by `templatePaths`. The
plugin should return an instance of `PolicyValidationPluginReportBeta1`. This object
represents the report that the user wil receive at the end of the synthesis.

```go
type myPlugin struct {
	name
}

func (this *myPlugin) validate(context iPolicyValidationContextBeta1) policyValidationPluginReportBeta1 {
	// First read the templates using context.templatePaths...

	// ...then perform the validation, and then compose and return the report.
	// Using hard-coded values here for better clarity:
	return &policyValidationPluginReportBeta1{
		Success: jsii.Boolean(false),
		Violations: []policyViolationBeta1{
			&policyViolationBeta1{
				RuleName: jsii.String("CKV_AWS_117"),
				Description: jsii.String("Ensure that AWS Lambda function is configured inside a VPC"),
				Fix: jsii.String("https://docs.bridgecrew.io/docs/ensure-that-aws-lambda-function-is-configured-inside-a-vpc-1"),
				ViolatingResources: []policyViolatingResourceBeta1{
					&policyViolatingResourceBeta1{
						ResourceLogicalId: jsii.String("MyFunction3BAA72D1"),
						TemplatePath: jsii.String("/home/johndoe/myapp/cdk.out/MyService.template.json"),
						Locations: []*string{
							jsii.String("Properties/VpcConfig"),
						},
					},
				},
			},
		},
	}
}
```

In addition to the name, plugins may optionally report their version (`version`
property ) and a list of IDs of the rules they are going to evaluate (`ruleIds`
property).

Note that plugins are not allowed to modify anything in the cloud assembly. Any
attempt to do so will result in synthesis failure.

If your plugin depends on an external tool, keep in mind that some developers may
not have that tool installed in their workstations yet. To minimize friction, we
highly recommend that you provide some installation script along with your
plugin package, to automate the whole process. Better yet, run that script as
part of the installation of your package. With `npm`, for example, you can run
add it to the `postinstall`
[script](https://docs.npmjs.com/cli/v9/using-npm/scripts) in the `package.json`
file.

## Annotations

Construct authors can add annotations to constructs to report at three different
levels: `ERROR`, `WARN`, `INFO`.

Typically warnings are added for things that are important for the user to be
aware of, but will not cause deployment errors in all cases. Some common
scenarios are (non-exhaustive list):

* Warn when the user needs to take a manual action, e.g. IAM policy should be
  added to an referenced resource.
* Warn if the user configuration might not follow best practices (but is still
  valid)
* Warn if the user is using a deprecated API

### Acknowledging Warnings

If you would like to run with `--strict` mode enabled (warnings will throw
errors) it is possible to `acknowledge` warnings to make the warning go away.

For example, if > 10 IAM managed policies are added to an IAM Group, a warning
will be created:

```text
IAM:Group:MaxPoliciesExceeded: You added 11 to IAM Group my-group. The maximum number of managed policies attached to an IAM group is 10.
```

If you have requested a [quota increase](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entities)
you may have the ability to add > 10 managed policies which means that this
warning does not apply to you. You can acknowledge this by `acknowledging` the
warning by the `id`.

```go
awscdk.Annotations_Of(this).AcknowledgeWarning(jsii.String("IAM:Group:MaxPoliciesExceeded"), jsii.String("Account has quota increased to 20"))
```

<!--END CORE DOCUMENTATION-->

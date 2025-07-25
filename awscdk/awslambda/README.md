# AWS Lambda Construct Library

This construct library allows you to define AWS Lambda Functions.

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

When deployed, this construct creates or updates an existing
`AWS::Lambda::Function` resource. When updating, AWS CloudFormation calls the
[UpdateFunctionConfiguration](https://docs.aws.amazon.com/lambda/latest/api/API_UpdateFunctionConfiguration.html)
and [UpdateFunctionCode](https://docs.aws.amazon.com/lambda/latest/api/API_UpdateFunctionCode.html)
Lambda APIs under the hood. Because these calls happen sequentially, and
invocations can happen between these calls, your function may encounter errors
in the time between the calls. For example, if you update an existing Lambda
function by removing an environment variable and the code that references that
environment variable in the same CDK deployment, you may see invocation errors
related to a missing environment variable. To work around this, you can invoke
your function against a version or alias by default, rather than the `$LATEST`
version.

To further mitigate these issues, you can ensure consistency between your function code and infrastructure configuration by defining environment variables as a single source of truth in your CDK stack. You can define them in a separate `env.ts` file and reference them in both your handler and CDK configuration. This approach allows you to catch errors at compile time, benefit from improved IDE support, minimize the risk of mismatched configurations, and enhance maintainability.

## Handler Code

The `lambda.Code` class includes static convenience methods for various types of
runtime code.

* `lambda.Code.fromBucket(bucket, key, objectVersion)` - specify an S3 object
  that contains the archive of your runtime code.
* `lambda.Code.fromBucketV2(bucket, key, {objectVersion: version, sourceKMSKey: key})` - specify an S3 object
  that contains the archive of your runtime code.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"
var key key


bucket := s3.NewBucket(this, jsii.String("Bucket"))

options := map[string]key{
	"sourceKMSKey": key,
}
fnBucket := lambda.NewFunction(this, jsii.String("myFunction2"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromBucketV2(bucket, jsii.String("python-lambda-handler.zip"), options),
})
```

* `lambda.Code.fromInline(code)` - inline the handle code as a string. This is
  limited to supported runtimes.
* `lambda.Code.fromAsset(path)` - specify a directory or a .zip file in the local
  filesystem which will be zipped and uploaded to S3 before deployment. See also
  [bundling asset code](#bundling-asset-code).
* `lambda.Code.fromDockerBuild(path, options)` - use the result of a Docker
  build as code. The runtime code is expected to be located at `/asset` in the
  image and will be zipped and uploaded to S3 as an asset.
* `lambda.Code.fromCustomCommand(output, command, customCommandOptions)` -
  supply a command that is invoked during cdk synth. That command is meant to direct
  the generated code to output (a zip file or a directory), which is then used as the
  code for the created AWS Lambda.

The following example shows how to define a Python function and deploy the code
from the local directory `my-lambda-handler` to it:

```go
lambda.NewFunction(this, jsii.String("MyLambda"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-lambda-handler"))),
	Handler: jsii.String("index.main"),
	Runtime: lambda.Runtime_PYTHON_3_9(),
})
```

When deploying a stack that contains this code, the directory will be zip
archived and then uploaded to an S3 bucket, then the exact location of the S3
objects will be passed when the stack is deployed.

During synthesis, the CDK expects to find a directory on disk at the asset
directory specified. Note that we are referencing the asset directory relatively
to our CDK project directory. This is especially important when we want to share
this construct through a library. Different programming languages will have
different techniques for bundling resources into libraries.

## Docker Images

Lambda functions allow specifying their handlers within docker images. The docker
image can be an image from ECR or a local asset that the CDK will package and load
into ECR.

The following `DockerImageFunction` construct uses a local folder with a
Dockerfile as the asset that will be used as the function handler.

```go
lambda.NewDockerImageFunction(this, jsii.String("AssetFunction"), &DockerImageFunctionProps{
	Code: lambda.DockerImageCode_FromImageAsset(path.join(__dirname, jsii.String("docker-handler"))),
})
```

You can also specify an image that already exists in ECR as the function handler.

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"

repo := ecr.NewRepository(this, jsii.String("Repository"))

lambda.NewDockerImageFunction(this, jsii.String("ECRFunction"), &DockerImageFunctionProps{
	Code: lambda.DockerImageCode_FromEcr(repo),
})
```

The props for these docker image resources allow overriding the image's `CMD`, `ENTRYPOINT`, and `WORKDIR`
configurations as well as choosing a specific tag or digest. See their docs for more information.

To deploy a `DockerImageFunction` on Lambda `arm64` architecture, specify `Architecture.ARM_64` in `architecture`.
This will bundle docker image assets for `arm64` architecture with `--platform linux/arm64` even if build within an `x86_64` host.

With that being said, if you are bundling `DockerImageFunction` for Lambda `amd64` architecture from a `arm64` machine like a Macbook with `arm64` CPU, you would
need to specify `architecture: lambda.Architecture.X86_64` as well. This ensures the `--platform` argument is passed to the image assets
bundling process so you can bundle up `X86_64` images from the `arm64` machine.

```go
lambda.NewDockerImageFunction(this, jsii.String("AssetFunction"), &DockerImageFunctionProps{
	Code: lambda.DockerImageCode_FromImageAsset(path.join(__dirname, jsii.String("docker-arm64-handler"))),
	Architecture: lambda.Architecture_ARM_64(),
})
```

## Execution Role

Lambda functions assume an IAM role during execution. In CDK by default, Lambda
functions will use an autogenerated Role if one is not provided.

The autogenerated Role is automatically given permissions to execute the Lambda
function. To reference the autogenerated Role:

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})

role := fn.Role
```

You can also provide your own IAM role. Provided IAM roles will not automatically
be given permissions to execute the Lambda function. To provide a role and grant
it appropriate permissions:

```go
myRole := iam.NewRole(this, jsii.String("My Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	Role: myRole,
})

myRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))
myRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaVPCAccessExecutionRole")))
```

## Function Timeout

AWS Lambda functions have a default timeout of 3 seconds, but this can be increased
up to 15 minutes. The timeout is available as a property of `Function` so that
you can reference it elsewhere in your stack. For instance, you could use it to create
a CloudWatch alarm to report when your function timed out:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
})

if fn.Timeout {
	cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &AlarmProps{
		Metric: fn.metricDuration().With(&MetricOptions{
			Statistic: jsii.String("Maximum"),
		}),
		EvaluationPeriods: jsii.Number(1),
		DatapointsToAlarm: jsii.Number(1),
		Threshold: fn.*Timeout.ToMilliseconds(),
		TreatMissingData: cloudwatch.TreatMissingData_IGNORE,
		AlarmName: jsii.String("My Lambda Timeout"),
	})
}
```

## Advanced Logging

You can have more control over your function logs, by specifying the log format
(Json or plain text), the system log level, the application log level, as well
as choosing the log group:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var logGroup iLogGroup


lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("foo")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	LoggingFormat: lambda.LoggingFormat_JSON,
	SystemLogLevelV2: lambda.SystemLogLevel_INFO,
	ApplicationLogLevelV2: lambda.ApplicationLogLevel_INFO,
	LogGroup: logGroup,
})
```

To use `applicationLogLevelV2` and/or `systemLogLevelV2` you must set `loggingFormat` to `LoggingFormat.JSON`.

## Customizing Log Group Creation

### Log Group Creation Methods

| Method | Description | Tag Propagation | Prop Injection | Aspects | Notes |
|--------|-------------|-----------------|----------------|---------|-------|
| **logRetention prop** | Legacy approach using Custom Resource | False | False | False | Does not support TPA |
| **logGroup prop** | Explicitly supplied by user in CDK app | True | True | True | Full support for TPA |
| **Lazy creation** | Lambda service creates logGroup on first invocation | False | False | False | Occurs when both logRetention and logGroup are undefined and USE_CDK_MANAGED_LAMBDA_LOGGROUP is false |
| **USE_CDK_MANAGED_LAMBDA_LOGGROUP** | CDK Lambda function construct creates log group with default props | True | True | True | Feature flag must be enabled |

*TPA: Tag propagation, Prop Injection, Aspects*

#### Order of precedence

```text
                       Highest Precedence
                             |
             +---------------+---------------+
             |                               |
  +-------------------------+      +------------------------+
  |   logRetention is set   |      |     logGroup is set    |
  +-----------+-------------+      +----------+-------------+
              |                               |
              v                               v
      Create LogGroup via            Use the provided LogGroup
  Custom Resource (retention       instance (CDK-managed,
  managed, logGroup disallowed)    logRetention disallowed)
              |                               |
              +---------------+---------------+
                              |
                              v
          +-----------------------------------------------+
          |         Feature flag enabled:                 |
          | aws-cdk:aws-lambda:useCdkManagedLogGroup      |
          +------------------ +---------------------------+
                              |
                              v
              Create LogGroup at synth time
          (CDK-managed, default settings for logGroup)
                              |
                              v
                  +---------------------------+
                  | Default (no config set)   |
                  +------------+--------------+
                              |
                              v
     Lambda service creates log group on first invocation runtime
            (CDK does not manage the log group resource)
```

### Tag Propagation

Refer section `Log Group Creation Methods` to find out which modes support tag propagation.
As an example, adding the following line in your cdk app will also propagate to the logGroup.

```
const fn = new lambda.Function(this, 'MyFunctionWithFFTrue', {
  runtime: lambda.Runtime.NODEJS_20_X,
  handler: 'handler.main',
  code: lambda.Code.fromAsset('lambda'),
});
cdk.Tags.of(fn).add('env', 'dev'); // the tag is also added to the log group
```

### Log removal policy

When using the deprecated `logRetention` property for creating a LogGroup, you can configure log removal policy:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunctionWithFFTrue"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("handler.main"),
	Code: lambda.Code_FromAsset(jsii.String("lambda")),
	LogRetention: logs.RetentionDays_INFINITE,
	LogRemovalPolicy: awscdk.RemovalPolicy_RETAIN,
})
```

## Resource-based Policies

AWS Lambda supports resource-based policies for controlling access to Lambda
functions and layers on a per-resource basis. In particular, this allows you to
give permission to AWS services, AWS Organizations, or other AWS accounts to
modify and invoke your functions.

### Grant function access to AWS services

```go
// Grant permissions to a service
var fn function

principal := iam.NewServicePrincipal(jsii.String("my-service"))

fn.GrantInvoke(principal)

// Equivalent to:
fn.AddPermission(jsii.String("my-service Invocation"), &Permission{
	Principal: principal,
})
```

You can also restrict permissions given to AWS services by providing
a source account or ARN (representing the account and identifier of the resource
that accesses the function or layer).

**Important**:

> By default `fn.grantInvoke()` grants permission to the principal to invoke any version of the function, including all past ones. If you only want the principal to be granted permission to invoke the latest version or the unqualified Lambda ARN, use `grantInvokeLatestVersion(grantee)`.

```go
var fn function

principal := iam.NewServicePrincipal(jsii.String("my-service"))
// Grant invoke only to latest version and unqualified lambda arn
fn.GrantInvokeLatestVersion(principal)
```

If you want to grant access for invoking a specific version of Lambda function, you can use `fn.grantInvokeVersion(grantee, version)`

```go
var fn function
var version iVersion

principal := iam.NewServicePrincipal(jsii.String("my-service"))
// Grant invoke only to the specific version
fn.GrantInvokeVersion(principal, version)
```

For more information, see
[Granting function access to AWS services](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-serviceinvoke)
in the AWS Lambda Developer Guide.

### Grant function access to an AWS Organization

```go
// Grant permissions to an entire AWS organization
var fn function

org := iam.NewOrganizationPrincipal(jsii.String("o-xxxxxxxxxx"))

fn.GrantInvoke(org)
```

In the above example, the `principal` will be `*` and all users in the
organization `o-xxxxxxxxxx` will get function invocation permissions.

You can restrict permissions given to the organization by specifying an
AWS account or role as the `principal`:

```go
// Grant permission to an account ONLY IF they are part of the organization
var fn function

account := iam.NewAccountPrincipal(jsii.String("123456789012"))

fn.GrantInvoke(account.InOrganization(jsii.String("o-xxxxxxxxxx")))
```

For more information, see
[Granting function access to an organization](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xorginvoke)
in the AWS Lambda Developer Guide.

### Grant function access to other AWS accounts

```go
// Grant permission to other AWS account
var fn function

account := iam.NewAccountPrincipal(jsii.String("123456789012"))

fn.GrantInvoke(account)
```

For more information, see
[Granting function access to other accounts](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xaccountinvoke)
in the AWS Lambda Developer Guide.

### Grant function access to unowned principals

Providing an unowned principal (such as account principals, generic ARN
principals, service principals, and principals in other accounts) to a call to
`fn.grantInvoke` will result in a resource-based policy being created. If the
principal in question has conditions limiting the source account or ARN of the
operation (see above), these conditions will be automatically added to the
resource policy.

```go
var fn function

servicePrincipal := iam.NewServicePrincipal(jsii.String("my-service"))
sourceArn := "arn:aws:s3:::amzn-s3-demo-bucket"
sourceAccount := "111122223333"
servicePrincipalWithConditions := servicePrincipal.WithConditions(map[string]interface{}{
	"ArnLike": map[string]*string{
		"aws:SourceArn": sourceArn,
	},
	"StringEquals": map[string]*string{
		"aws:SourceAccount": sourceAccount,
	},
})

fn.GrantInvoke(servicePrincipalWithConditions)
```

### Grant function access to a CompositePrincipal

To grant invoke permissions to a `CompositePrincipal` use the `grantInvokeCompositePrincipal` method:

```go
var fn function

compositePrincipal := iam.NewCompositePrincipal(
iam.NewOrganizationPrincipal(jsii.String("o-zzzzzzzzzz")),
iam.NewServicePrincipal(jsii.String("apigateway.amazonaws.com")))

fn.GrantInvokeCompositePrincipal(compositePrincipal)
```

## Versions

You can use
[versions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html)
to manage the deployment of your AWS Lambda functions. For example, you can
publish a new version of a function for beta testing without affecting users of
the stable production version.

The function version includes the following information:

* The function code and all associated dependencies.
* The Lambda runtime that executes the function.
* All of the function settings, including the environment variables.
* A unique Amazon Resource Name (ARN) to identify this version of the function.

You could create a version to your lambda function using the `Version` construct.

```go
var fn function

version := lambda.NewVersion(this, jsii.String("MyVersion"), &VersionProps{
	Lambda: fn,
})
```

The major caveat to know here is that a function version must always point to a
specific 'version' of the function. When the function is modified, the version
will continue to point to the 'then version' of the function.

One way to ensure that the `lambda.Version` always points to the latest version
of your `lambda.Function` is to set an environment variable which changes at
least as often as your code does. This makes sure the function always has the
latest code. For instance -

```go
codeVersion := "stringOrMethodToGetCodeVersion"
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	Environment: map[string]*string{
		"CodeVersionString": codeVersion,
	},
})
```

The `fn.latestVersion` property returns a `lambda.IVersion` which represents
the `$LATEST` pseudo-version.

However, most AWS services require a specific AWS Lambda version,
and won't allow you to use `$LATEST`. Therefore, you would normally want
to use `lambda.currentVersion`.

The `fn.currentVersion` property can be used to obtain a `lambda.Version`
resource that represents the AWS Lambda function defined in your application.
Any change to your function's code or configuration will result in the creation
of a new version resource. You can specify options for this version through the
`currentVersionOptions` property.

NOTE: The `currentVersion` property is only supported when your AWS Lambda function
uses either `lambda.Code.fromAsset` or `lambda.Code.fromInline`. Other types
of code providers (such as `lambda.Code.fromBucket`) require that you define a
`lambda.Version` resource directly since the CDK is unable to determine if
their contents had changed.

### `currentVersion`: Updated hashing logic

To produce a new lambda version each time the lambda function is modified, the
`currentVersion` property under the hood, computes a new logical id based on the
properties of the function. This informs CloudFormation that a new
`AWS::Lambda::Version` resource should be created pointing to the updated Lambda
function.

However, a bug was introduced in this calculation that caused the logical id to
change when it was not required (ex: when the Function's `Tags` property, or
when the `DependsOn` clause was modified). This caused the deployment to fail
since the Lambda service does not allow creating duplicate versions.

This has been fixed in the AWS CDK but *existing* users need to opt-in via a
[feature flag](https://docs.aws.amazon.com/cdk/latest/guide/featureflags.html). Users who have run `cdk init` since this fix will be opted in,
by default.

Otherwise, you will need to enable the [feature flag](https://docs.aws.amazon.com/cdk/latest/guide/featureflags.html)
`@aws-cdk/aws-lambda:recognizeVersionProps`. Since CloudFormation does not
allow duplicate versions, you will also need to make some modification to
your function so that a new version can be created. To efficiently and trivially
modify all your lambda functions at once, you can attach the
`FunctionVersionUpgrade` aspect to the stack, which slightly alters the
function description. This aspect is intended for one-time use to upgrade the
version of all your functions at the same time, and can safely be removed after
deploying once.

```go
stack := awscdk.Newstack()
awscdk.Aspects_Of(stack).Add(lambda.NewFunctionVersionUpgrade(awscdklibcxapi.LAMBDA_RECOGNIZE_VERSION_PROPS))
```

When the new logic is in effect, you may rarely come across the following error:
`The following properties are not recognized as version properties`. This will
occur, typically when [property overrides](https://docs.aws.amazon.com/cdk/latest/guide/cfn_layer.html#cfn_layer_raw) are used, when a new property
introduced in `AWS::Lambda::Function` is used that CDK is still unaware of.

To overcome this error, use the API `Function.classifyVersionProperty()` to
record whether a new version should be generated when this property is changed.
This can be typically determined by checking whether the property can be
modified using the *[UpdateFunctionConfiguration](https://docs.aws.amazon.com/lambda/latest/dg/API_UpdateFunctionConfiguration.html)* API or not.

### `currentVersion`: Updated hashing logic for layer versions

An additional update to the hashing logic fixes two issues surrounding layers.
Prior to this change, updating the lambda layer version would have no effect on
the function version. Also, the order of lambda layers provided to the function
was unnecessarily baked into the hash.

This has been fixed in the AWS CDK starting with version 2.27. If you ran
`cdk init` with an earlier version, you will need to opt-in via a [feature flag](https://docs.aws.amazon.com/cdk/latest/guide/featureflags.html).
If you run `cdk init` with v2.27 or later, this fix will be opted in, by default.

Existing users will need to enable the [feature flag](https://docs.aws.amazon.com/cdk/latest/guide/featureflags.html)
`@aws-cdk/aws-lambda:recognizeLayerVersion`. Since CloudFormation does not
allow duplicate versions, they will also need to make some modification to
their function so that a new version can be created. To efficiently and trivially
modify all your lambda functions at once, users can attach the
`FunctionVersionUpgrade` aspect to the stack, which slightly alters the
function description. This aspect is intended for one-time use to upgrade the
version of all your functions at the same time, and can safely be removed after
deploying once.

```go
stack := awscdk.Newstack()
awscdk.Aspects_Of(stack).Add(lambda.NewFunctionVersionUpgrade(awscdklibcxapi.LAMBDA_RECOGNIZE_LAYER_VERSION))
```

## Aliases

You can define one or more
[aliases](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html)
for your AWS Lambda function. A Lambda alias is like a pointer to a specific
Lambda function version. Users can access the function version using the alias
ARN.

The `version.addAlias()` method can be used to define an AWS Lambda alias that
points to a specific version.

The following example defines an alias named `live` which will always point to a
version that represents the function as defined in your CDK app. When you change
your lambda code or configuration, a new resource will be created. You can
specify options for the current version through the `currentVersionOptions`
property.

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	CurrentVersionOptions: &VersionOptions{
		RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
		 // retain old versions
		RetryAttempts: jsii.Number(1),
	},
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})

fn.AddAlias(jsii.String("live"))
```

## Function URL

A function URL is a dedicated HTTP(S) endpoint for your Lambda function. When you create a function URL, Lambda automatically generates a unique URL endpoint for you. Function URLs can be created for the latest version Lambda Functions, or Function Aliases (but not for Versions).

Function URLs are dual stack-enabled, supporting IPv4 and IPv6, and cross-origin resource sharing (CORS) configuration. After you configure a function URL for your function, you can invoke your function through its HTTP(S) endpoint via a web browser, curl, Postman, or any HTTP client. To invoke a function using IAM authentication your HTTP client must support SigV4 signing.

See the [Invoking Function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-invocation.html) section of the AWS Lambda Developer Guide
for more information on the input and output payloads of Functions invoked in this way.

### IAM-authenticated Function URLs

To create a Function URL which can be called by an IAM identity, call `addFunctionUrl()`, followed by `grantInvokeFunctionUrl()`:

```go
// Can be a Function or an Alias
var fn function
var myRole role


fnUrl := fn.AddFunctionUrl()
fnUrl.GrantInvokeUrl(myRole)

awscdk.NewCfnOutput(this, jsii.String("TheUrl"), &CfnOutputProps{
	// The .url attributes will return the unique Function URL
	Value: fnUrl.Url,
})
```

Calls to this URL need to be signed with SigV4.

### Anonymous Function URLs

To create a Function URL which can be called anonymously, pass `authType: FunctionUrlAuthType.NONE` to `addFunctionUrl()`:

```go
// Can be a Function or an Alias
var fn function


fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_NONE,
})

awscdk.NewCfnOutput(this, jsii.String("TheUrl"), &CfnOutputProps{
	Value: fnUrl.Url,
})
```

### CORS configuration for Function URLs

If you want your Function URLs to be invokable from a web page in browser, you
will need to configure cross-origin resource sharing to allow the call (if you do
not do this, your browser will refuse to make the call):

```go
var fn function


fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_NONE,
	Cors: &FunctionUrlCorsOptions{
		// Allow this to be called from websites on https://example.com.
		// Can also be ['*'] to allow all domain.
		AllowedOrigins: []*string{
			jsii.String("https://example.com"),
		},
	},
})
```

### Invoke Mode for Function URLs

Invoke mode determines how AWS Lambda invokes your function. You can configure the invoke mode when creating a Function URL using the invokeMode property

```go
var fn function


fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_NONE,
	InvokeMode: lambda.InvokeMode_RESPONSE_STREAM,
})
```

If the invokeMode property is not specified, the default BUFFERED mode will be used.

## Layers

The `lambda.LayerVersion` class can be used to define Lambda layers and manage
granting permissions to other AWS accounts or organizations.

```go
layer := lambda.NewLayerVersion(stack, jsii.String("MyLayer"), &LayerVersionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("layer-code"))),
	CompatibleRuntimes: []runtime{
		lambda.*runtime_NODEJS_LATEST(),
	},
	License: jsii.String("Apache-2.0"),
	Description: jsii.String("A layer to test the L2 construct"),
})

// To grant usage by other AWS accounts
layer.addPermission(jsii.String("remote-account-grant"), &LayerVersionPermission{
	AccountId: awsAccountId,
})

// To grant usage to all accounts in some AWS Ogranization
// layer.grantUsage({ accountId: '*', organizationId });

// To grant usage to all accounts in some AWS Ogranization
// layer.grantUsage({ accountId: '*', organizationId });
lambda.NewFunction(stack, jsii.String("MyLayeredLambda"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("foo")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.*runtime_NODEJS_LATEST(),
	Layers: []iLayerVersion{
		layer,
	},
})
```

By default, updating a layer creates a new layer version, and CloudFormation will delete the old version as part of the stack update.

Alternatively, a removal policy can be used to retain the old version:

```go
lambda.NewLayerVersion(this, jsii.String("MyLayer"), &LayerVersionProps{
	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

## Architecture

Lambda functions, by default, run on compute systems that have the 64 bit x86 architecture.

The AWS Lambda service also runs compute on the ARM architecture, which can reduce cost
for some workloads.

A lambda function can be configured to be run on one of these platforms:

```go
lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	Architecture: lambda.Architecture_ARM_64(),
})
```

Similarly, lambda layer versions can also be tagged with architectures it is compatible with.

```go
lambda.NewLayerVersion(this, jsii.String("MyLayer"), &LayerVersionProps{
	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	CompatibleArchitectures: []architecture{
		lambda.*architecture_X86_64(),
		lambda.*architecture_ARM_64(),
	},
})
```

## Lambda Insights

Lambda functions can be configured to use CloudWatch [Lambda Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights.html)
which provides low-level runtime metrics for a Lambda functions.

```go
lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	InsightsVersion: lambda.LambdaInsightsVersion_VERSION_1_0_98_0(),
})
```

If the version of insights is not yet available in the CDK, you can also provide the ARN directly as so -

```go
layerArn := "arn:aws:lambda:us-east-1:580247275435:layer:LambdaInsightsExtension:14"
lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	InsightsVersion: lambda.LambdaInsightsVersion_FromInsightVersionArn(layerArn),
})
```

If you are deploying an ARM_64 Lambda Function, you must specify a
Lambda Insights Version >= `1_0_119_0`.

```go
lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Architecture: lambda.Architecture_ARM_64(),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	InsightsVersion: lambda.LambdaInsightsVersion_VERSION_1_0_119_0(),
})
```

### Parameters and Secrets Extension

Lambda functions can be configured to use the Parameters and Secrets Extension. The Parameters and Secrets Extension can be used to retrieve and cache [secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets_lambda.html) from Secrets Manager or [parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/ps-integration-lambda-extensions.html) from Parameter Store in Lambda functions without using an SDK.

```go
import sm "github.com/aws/aws-cdk-go/awscdk"
import ssm "github.com/aws/aws-cdk-go/awscdk"


secret := sm.NewSecret(this, jsii.String("Secret"))
parameter := ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
	ParameterName: jsii.String("mySsmParameterName"),
	StringValue: jsii.String("mySsmParameterValue"),
})

paramsAndSecrets := lambda.ParamsAndSecretsLayerVersion_FromVersion(lambda.ParamsAndSecretsVersions_V1_0_103, &ParamsAndSecretsOptions{
	CacheSize: jsii.Number(500),
	LogLevel: lambda.ParamsAndSecretsLogLevel_DEBUG,
})

lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Architecture: lambda.Architecture_ARM_64(),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	ParamsAndSecrets: ParamsAndSecrets,
})

secret.grantRead(lambdaFunction)
parameter.grantRead(lambdaFunction)
```

If the version of Parameters and Secrets Extension is not yet available in the CDK, you can also provide the ARN directly as so:

```go
import sm "github.com/aws/aws-cdk-go/awscdk"
import ssm "github.com/aws/aws-cdk-go/awscdk"


secret := sm.NewSecret(this, jsii.String("Secret"))
parameter := ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
	ParameterName: jsii.String("mySsmParameterName"),
	StringValue: jsii.String("mySsmParameterValue"),
})

layerArn := "arn:aws:lambda:us-east-1:177933569100:layer:AWS-Parameters-and-Secrets-Lambda-Extension:4"
paramsAndSecrets := lambda.ParamsAndSecretsLayerVersion_FromVersionArn(layerArn, &ParamsAndSecretsOptions{
	CacheSize: jsii.Number(500),
})

lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Architecture: lambda.Architecture_ARM_64(),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	ParamsAndSecrets: ParamsAndSecrets,
})

secret.grantRead(lambdaFunction)
parameter.grantRead(lambdaFunction)
```

## Event Rule Target

You can use an AWS Lambda function as a target for an Amazon CloudWatch event
rule:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var fn function

rule := events.NewRule(this, jsii.String("Schedule Rule"), &RuleProps{
	Schedule: events.Schedule_Cron(&CronOptions{
		Minute: jsii.String("0"),
		Hour: jsii.String("4"),
	}),
})
rule.AddTarget(targets.NewLambdaFunction(fn))
```

## Event Sources

AWS Lambda supports a [variety of event sources](https://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-function.html).

In most cases, it is possible to trigger a function as a result of an event by
using one of the `add<Event>Notification` methods on the source construct. For
example, the `s3.Bucket` construct has an `onEvent` method which can be used to
trigger a Lambda when an event, such as PutObject occurs on an S3 bucket.

An alternative way to add event sources to a function is to use `function.addEventSource(source)`.
This method accepts an `IEventSource` object. The module **@aws-cdk/aws-lambda-event-sources**
includes classes for the various event sources supported by AWS Lambda.

For example, the following code adds an SQS queue as an event source for a function:

```go
import eventsources "github.com/aws/aws-cdk-go/awscdk"
import sqs "github.com/aws/aws-cdk-go/awscdk"

var fn function

queue := sqs.NewQueue(this, jsii.String("Queue"))
fn.AddEventSource(eventsources.NewSqsEventSource(queue))
```

The following code adds an S3 bucket notification as an event source:

```go
import eventsources "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

bucket := s3.NewBucket(this, jsii.String("Bucket"))
fn.AddEventSource(eventsources.NewS3EventSource(bucket, &S3EventSourceProps{
	Events: []eventType{
		s3.*eventType_OBJECT_CREATED,
		s3.*eventType_OBJECT_REMOVED,
	},
	Filters: []notificationKeyFilter{
		&notificationKeyFilter{
			Prefix: jsii.String("subdir/"),
		},
	},
}))
```

The following code adds an DynamoDB notification as an event source filtering insert events:

```go
import eventsources "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Stream: dynamodb.StreamViewType_NEW_IMAGE,
})
fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
	StartingPosition: lambda.StartingPosition_LATEST,
	Filters: []map[string]interface{}{
		lambda.FilterCriteria_Filter(map[string]interface{}{
			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
		}),
	},
}))
```

By default, Lambda will encrypt Filter Criteria using AWS managed keys. But if you want to use a self managed KMS key to encrypt the filters, You can specify the self managed key using the `filterEncryption` property.

```go
import eventsources "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Stream: dynamodb.StreamViewType_NEW_IMAGE,
})
// Your self managed KMS key
myKey := awscdk.Key_FromKeyArn(this, jsii.String("SourceBucketEncryptionKey"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/<key-id>"))

fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
	StartingPosition: lambda.StartingPosition_LATEST,
	Filters: []map[string]interface{}{
		lambda.FilterCriteria_Filter(map[string]interface{}{
			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
		}),
	},
	FilterEncryption: myKey,
}))
```

> Lambda requires allow `kms:Decrypt` on Lambda principal `lambda.amazonaws.com` to use the key for Filter Criteria Encryption. If you create the KMS key in the stack, CDK will automatically add this permission to the Key when you creates eventSourceMapping. However, if you import the key using function like `Key.fromKeyArn` then you need to add the following permission to the KMS key before using it to encrypt Filter Criteria

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "kms:Decrypt",
            "Resource": "*"
        }
    ]
}
```

### Observability

Customers can now opt-in to get enhanced metrics for their event source mapping that capture each stage of processing using the `MetricsConfig` property.

The following code shows how to opt in for the enhanced metrics.

```go
import eventsources "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Stream: dynamodb.StreamViewType_NEW_IMAGE,
})

fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
	StartingPosition: lambda.StartingPosition_LATEST,
	MetricsConfig: &MetricsConfig{
		Metrics: []eVENT_COUNT{
			lambda.MetricType_*eVENT_COUNT,
		},
	},
}))
```

See the documentation for the **@aws-cdk/aws-lambda-event-sources** module for more details.

## Imported Lambdas

When referencing an imported lambda in the CDK, use `fromFunctionArn()` for most use cases:

```go
fn := lambda.Function_FromFunctionArn(this, jsii.String("Function"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"))
```

The `fromFunctionAttributes()` API is available for more specific use cases:

```go
fn := lambda.Function_FromFunctionAttributes(this, jsii.String("Function"), &FunctionAttributes{
	FunctionArn: jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"),
	// The following are optional properties for specific use cases and should be used with caution:

	// Use Case: imported function is in the same account as the stack. This tells the CDK that it
	// can modify the function's permissions.
	SameEnvironment: jsii.Boolean(true),

	// Use Case: imported function is in a different account and user commits to ensuring that the
	// imported function has the correct permissions outside the CDK.
	SkipPermissions: jsii.Boolean(true),
})
```

`Function.fromFunctionArn()` and `Function.fromFunctionAttributes()` will attempt to parse the Function's Region and Account ID from the ARN. `addPermissions` will only work on the `Function` object if the Region and Account ID are deterministically the same as the scope of the Stack the referenced `Function` object is created in.
If the containing Stack is environment-agnostic or the Function ARN is a Token, this comparison will fail, and calls to `Function.addPermission` will do nothing.
If you know Function permissions can safely be added, you can use `Function.fromFunctionName()` instead, or pass `sameEnvironment: true` to `Function.fromFunctionAttributes()`.

```go
fn := lambda.Function_FromFunctionName(this, jsii.String("Function"), jsii.String("MyFn"))
```

## Lambda with DLQ

A dead-letter queue can be automatically created for a Lambda function by
setting the `deadLetterQueueEnabled: true` configuration. In such case CDK creates
a `sqs.Queue` as `deadLetterQueue`.

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
	DeadLetterQueueEnabled: jsii.Boolean(true),
})
```

It is also possible to provide a dead-letter queue instead of getting a new queue created:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"


dlq := sqs.NewQueue(this, jsii.String("DLQ"))
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
	DeadLetterQueue: dlq,
})
```

You can also use a `sns.Topic` instead of an `sqs.Queue` as dead-letter queue:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


dlt := sns.NewTopic(this, jsii.String("DLQ"))
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("// your code here")),
	DeadLetterTopic: dlt,
})
```

See [the AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/dlq.html)
to learn more about AWS Lambdas and DLQs.

## Lambda with X-Ray Tracing

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
	Tracing: lambda.Tracing_ACTIVE,
})
```

See [the AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html)
to learn more about AWS Lambda's X-Ray support.

## Lambda with AWS Distro for OpenTelemetry layer

To have automatic integration with XRay without having to add dependencies or change your code, you can use the
[AWS Distro for OpenTelemetry Lambda (ADOT) layer](https://aws-otel.github.io/docs/getting-started/lambda).
Consuming the latest ADOT layer can be done with the following snippet:

```go
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
	AdotInstrumentation: &AdotInstrumentationConfig{
		LayerVersion: awscdk.AdotLayerVersion_FromJavaScriptSdkLayerVersion(awscdk.AdotLambdaLayerJavaScriptSdkVersion_LATEST()),
		ExecWrapper: awscdk.AdotLambdaExecWrapper_REGULAR_HANDLER,
	},
})
```

To use a different layer version, use one of the following helper functions for the `layerVersion` prop:

* `AdotLayerVersion.fromJavaScriptSdkLayerVersion`
* `AdotLayerVersion.fromPythonSdkLayerVersion`
* `AdotLayerVersion.fromJavaSdkLayerVersion`
* `AdotLayerVersion.fromJavaAutoInstrumentationSdkLayerVersion`
* `AdotLayerVersion.fromGenericSdkLayerVersion`

Each helper function expects a version value from a corresponding enum-like class as below:

* `AdotLambdaLayerJavaScriptSdkVersion`
* `AdotLambdaLayerPythonSdkVersion`
* `AdotLambdaLayerJavaSdkVersion`
* `AdotLambdaLayerJavaAutoInstrumentationSdkVersion`
* `AdotLambdaLayerGenericSdkVersion`

For more examples, see our [the integration test](test/integ.lambda-adot.ts).

If you want to retrieve the ARN of the ADOT Lambda layer without enabling ADOT in a Lambda function:

```go
var fn function

layerArn := lambda.AdotLambdaLayerJavaSdkVersion_V1_19_0().layerArn(fn.Stack, fn.Architecture)
```

When using the `AdotLambdaLayerPythonSdkVersion` the `AdotLambdaExecWrapper` needs to be `AdotLambdaExecWrapper.INSTRUMENT_HANDLER` as per [AWS Distro for OpenTelemetry Lambda Support For Python](https://aws-otel.github.io/docs/getting-started/lambda/lambda-python)

## Lambda with Profiling

The following code configures the lambda function with CodeGuru profiling. By default, this creates a new CodeGuru
profiling group -

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda-handler")),
	Profiling: jsii.Boolean(true),
})
```

The `profilingGroup` property can be used to configure an existing CodeGuru profiler group.

CodeGuru profiling is supported for all Java runtimes and Python3.6+ runtimes.

See [the AWS documentation](https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html)
to learn more about AWS Lambda's Profiling support.

## Lambda with Reserved Concurrent Executions

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
	ReservedConcurrentExecutions: jsii.Number(100),
})
```

https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html

## Lambda with SnapStart

SnapStart is currently supported on Python 3.12, Python 3.13, .NET 8, and Java 11 and later [Java managed runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). SnapStart does not support provisioned concurrency, Amazon Elastic File System (Amazon EFS), or ephemeral storage greater than 512 MB. After you enable Lambda SnapStart for a particular Lambda function, publishing a new version of the function will trigger an optimization process.

See [the AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) to learn more about AWS Lambda SnapStart

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("handler.zip"))),
	Runtime: lambda.Runtime_JAVA_11(),
	Handler: jsii.String("example.Handler::handleRequest"),
	SnapStart: lambda.SnapStartConf_ON_PUBLISHED_VERSIONS(),
})

version := fn.currentVersion
```

## AutoScaling

You can use Application AutoScaling to automatically configure the provisioned concurrency for your functions. AutoScaling can be set to track utilization or be based on a schedule. To configure AutoScaling on a function alias:

```go
import autoscaling "github.com/aws/aws-cdk-go/awscdk"

var fn function

alias := fn.AddAlias(jsii.String("prod"))

// Create AutoScaling target
as := alias.AddAutoScaling(&AutoScalingOptions{
	MaxCapacity: jsii.Number(50),
})

// Configure Target Tracking
as.ScaleOnUtilization(&UtilizationScalingOptions{
	UtilizationTarget: jsii.Number(0.5),
})

// Configure Scheduled Scaling
as.ScaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &ScalingSchedule{
	Schedule: autoscaling.Schedule_Cron(&CronOptions{
		Hour: jsii.String("8"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(20),
})
```

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/cxapi"
import "github.com/aws/aws-cdk-go/awscdk"

/**
 * Stack verification steps:
 * aws application-autoscaling describe-scalable-targets --service-namespace lambda --resource-ids function:<function name>:prod
 * has a minCapacity of 3 and maxCapacity of 50
 */
type testStack struct {
	stack
}

func newTestStack(scope app, id *string) *testStack {
	this := &testStack{}
	cdk.NewStack_Override(this, scope, id)

	fn := lambda.NewFunction(this, jsii.String("MyLambda"), &FunctionProps{
		Code: lambda.NewInlineCode(jsii.String("exports.handler = async () => { console.log('hello world'); };")),
		Handler: jsii.String("index.handler"),
		Runtime: lambda.Runtime_NODEJS_LATEST(),
	})

	version := fn.currentVersion

	alias := lambda.NewAlias(this, jsii.String("Alias"), &AliasProps{
		AliasName: jsii.String("prod"),
		Version: Version,
	})

	scalingTarget := alias.AddAutoScaling(&AutoScalingOptions{
		MinCapacity: jsii.Number(3),
		MaxCapacity: jsii.Number(50),
	})

	scalingTarget.ScaleOnUtilization(&UtilizationScalingOptions{
		UtilizationTarget: jsii.Number(0.5),
	})

	scalingTarget.ScaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &ScalingSchedule{
		Schedule: appscaling.Schedule_Cron(&CronOptions{
			Hour: jsii.String("8"),
			Minute: jsii.String("0"),
		}),
		MinCapacity: jsii.Number(20),
	})

	scalingTarget.ScaleOnSchedule(jsii.String("ScaleDownAtNight"), &ScalingSchedule{
		Schedule: appscaling.Schedule_*Cron(&CronOptions{
			Hour: jsii.String("20"),
			Minute: jsii.String("0"),
		}),
		MaxCapacity: jsii.Number(20),
	})

	cdk.NewCfnOutput(this, jsii.String("FunctionName"), &CfnOutputProps{
		Value: fn.FunctionName,
	})
	return this
}

app := cdk.NewApp()

stack := NewTestStack(app, jsii.String("aws-lambda-autoscaling"))

// Changes the function description when the feature flag is present
// to validate the changed function hash.
cdk.Aspects_Of(stack).Add(lambda.NewFunctionVersionUpgrade(cxapi.LAMBDA_RECOGNIZE_LAYER_VERSION))

app.Synth()
```

See [the AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/invocation-scaling.html) on autoscaling lambda functions.

## Log Group

By default, Lambda functions automatically create a log group with the name `/aws/lambda/<function-name>` upon first execution with
log data set to never expire.
This is convenient, but prevents you from changing any of the properties of this auto-created log group using the AWS CDK.
For example you cannot set log retention or assign a data protection policy.

To fully customize the logging behavior of your Lambda function, create a `logs.LogGroup` ahead of time and use the `logGroup` property to instruct the Lambda function to send logs to it.
This way you can use the full features set supported by Amazon CloudWatch Logs.

```go
import "github.com/aws/aws-cdk-go/awscdk"


myLogGroup := awscdk.NewLogGroup(this, jsii.String("MyLogGroupWithLogGroupName"), &LogGroupProps{
	LogGroupName: jsii.String("customLogGroup"),
})

lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("foo")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	LogGroup: myLogGroup,
})
```

Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
If you are deploying to another type of region, please check regional availability first.

## Lambda with Recursive Loop protection

Recursive loop protection is to stop unintended loops. The customers are opted in by default for Lambda to detect and terminate unintended loops between Lambda and other AWS Services.
The property can be assigned two values here, "Allow" and "Terminate".

The default value is set to "Terminate", which lets the Lambda to detect and terminate the recursive loops.

When the value is set to "Allow", the customers opt out of recursive loop detection and Lambda does not terminate recursive loops if any.

See [the AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html) to learn more about AWS Lambda Recusrive Loop Detection

```go
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("handler.zip"))),
	Runtime: lambda.Runtime_JAVA_11(),
	Handler: jsii.String("example.Handler::handleRequest"),
	RecursiveLoop: lambda.RecursiveLoop_TERMINATE,
})
```

### Legacy Log Retention

As an alternative to providing a custom, user controlled log group, the legacy `logRetention` property can be used to set a different expiration period.
This feature uses a Custom Resource to change the log retention of the automatically created log group.

By default, CDK uses the AWS SDK retry options when creating a log group. The `logRetentionRetryOptions` property
allows you to customize the maximum number of retries and base backoff duration.

*Note* that a [CloudFormation custom
resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html) is added
to the stack that pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the
correct log retention period (never expire, by default). This Custom Resource will also create a log group to log events of the custom resource. The log retention period for this addtional log group is hard-coded to 1 day.

*Further note* that, if the log group already exists and the `logRetention` is not set, the custom resource will reset
the log retention to never expire even if it was configured with a different value.

## FileSystem Access

You can configure a function to mount an Amazon Elastic File System (Amazon EFS) to a
directory in your runtime environment with the `filesystem` property. To access Amazon EFS
from lambda function, the Amazon EFS access point will be required.

The following sample allows the lambda function to mount the Amazon EFS access point to `/mnt/msg` in the runtime environment and access the filesystem with the POSIX identity defined in `posixUser`.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import efs "github.com/aws/aws-cdk-go/awscdk"


// create a new VPC
vpc := ec2.NewVpc(this, jsii.String("VPC"))

// create a new Amazon EFS filesystem
fileSystem := efs.NewFileSystem(this, jsii.String("Efs"), &FileSystemProps{
	Vpc: Vpc,
})

// create a new access point from the filesystem
accessPoint := fileSystem.AddAccessPoint(jsii.String("AccessPoint"), &AccessPointOptions{
	// set /export/lambda as the root of the access point
	Path: jsii.String("/export/lambda"),
	// as /export/lambda does not exist in a new efs filesystem, the efs will create the directory with the following createAcl
	CreateAcl: &Acl{
		OwnerUid: jsii.String("1001"),
		OwnerGid: jsii.String("1001"),
		Permissions: jsii.String("750"),
	},
	// enforce the POSIX identity so lambda function will access with this identity
	PosixUser: &PosixUser{
		Uid: jsii.String("1001"),
		Gid: jsii.String("1001"),
	},
})

fn := lambda.NewFunction(this, jsii.String("MyLambda"), &FunctionProps{
	// mount the access point to /mnt/msg in the lambda runtime environment
	Filesystem: lambda.FileSystem_FromEfsAccessPoint(accessPoint, jsii.String("/mnt/msg")),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	Vpc: Vpc,
})
```

## IPv6 support

You can configure IPv6 connectivity for lambda function by setting `Ipv6AllowedForDualStack` to true.
It allows Lambda functions to specify whether the IPv6 traffic should be allowed when using dual-stack VPCs.
To access IPv6 network using Lambda, Dual-stack VPC is required. Using dual-stack VPC a function communicates with subnet over either of IPv4 or IPv6.

```go
import "github.com/aws/aws-cdk-go/awscdk"


natProvider := ec2.NatProvider_Gateway()

// create dual-stack VPC
vpc := ec2.NewVpc(this, jsii.String("DualStackVpc"), &VpcProps{
	IpProtocol: ec2.IpProtocol_DUAL_STACK,
	SubnetConfiguration: []subnetConfiguration{
		&subnetConfiguration{
			Name: jsii.String("Ipv6Public1"),
			SubnetType: ec2.SubnetType_PUBLIC,
		},
		&subnetConfiguration{
			Name: jsii.String("Ipv6Public2"),
			SubnetType: ec2.SubnetType_PUBLIC,
		},
		&subnetConfiguration{
			Name: jsii.String("Ipv6Private1"),
			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
	},
	NatGatewayProvider: natProvider,
})

natGatewayId := natProvider.ConfiguredGateways[0].GatewayId
(vpc.PrivateSubnets[0].(privateSubnet)).AddIpv6Nat64Route(natGatewayId)

fn := lambda.NewFunction(this, jsii.String("Lambda_with_IPv6_VPC"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("def main(event, context): pass")),
	Handler: jsii.String("index.main"),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Vpc: Vpc,
	Ipv6AllowedForDualStack: jsii.Boolean(true),
})
```

## Outbound traffic

By default, when creating a Lambda function, it would add a security group outbound rule to allow sending all network traffic (except IPv6). This is controlled by `allowAllOutbound` in function properties, which has a default value of `true`.

To allow outbound IPv6 traffic by default, explicitly set `allowAllIpv6Outbound` to `true` in function properties as shown below (the default value for `allowAllIpv6Outbound` is `false`):

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"))

fn := lambda.NewFunction(this, jsii.String("LambdaWithIpv6Outbound"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("def main(event, context): pass")),
	Handler: jsii.String("index.main"),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Vpc: vpc,
	AllowAllIpv6Outbound: jsii.Boolean(true),
})
```

Do not specify `allowAllOutbound` or `allowAllIpv6Outbound` property if the `securityGroups` or `securityGroup` property is set. Instead, configure these properties directly on the security group.

## Ephemeral Storage

You can configure ephemeral storage on a function to control the amount of storage it gets for reading
or writing data, allowing you to use AWS Lambda for ETL jobs, ML inference, or other data-intensive workloads.
The ephemeral storage will be accessible in the functions' `/tmp` directory.

```go
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
	EphemeralStorageSize: awscdk.Size_Mebibytes(jsii.Number(1024)),
})
```

Read more about using this feature in [this AWS blog post](https://aws.amazon.com/blogs/aws/aws-lambda-now-supports-up-to-10-gb-ephemeral-storage/).

## Singleton Function

The `SingletonFunction` construct is a way to guarantee that a lambda function will be guaranteed to be part of the stack,
once and only once, irrespective of how many times the construct is declared to be part of the stack. This is guaranteed
as long as the `uuid` property and the optional `lambdaPurpose` property stay the same whenever they're declared into the
stack.

A typical use case of this function is when a higher level construct needs to declare a Lambda function as part of it but
needs to guarantee that the function is declared once. However, a user of this higher level construct can declare it any
number of times and with different properties. Using `SingletonFunction` here with a fixed `uuid` will guarantee this.

For example, the `AwsCustomResource` construct requires only one single lambda function for all api calls that are made.

## Bundling Asset Code

When using `lambda.Code.fromAsset(path)` it is possible to bundle the code by running a
command in a Docker container. The asset path will be mounted at `/asset-input`. The
Docker container is responsible for putting content at `/asset-output`. The content at
`/asset-output` will be zipped and used as Lambda code.

Example with Python:

```go
lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-python-handler")), &AssetOptions{
		Bundling: &BundlingOptions{
			Image: lambda.Runtime_PYTHON_3_9().BundlingImage,
			Command: []*string{
				jsii.String("bash"),
				jsii.String("-c"),
				jsii.String("pip install -r requirements.txt -t /asset-output && cp -au . /asset-output"),
			},
		},
	}),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Handler: jsii.String("index.handler"),
})
```

Runtimes expose a `bundlingImage` property that points to the [AWS SAM](https://github.com/awslabs/aws-sam-cli) build image.

Use `cdk.DockerImage.fromRegistry(image)` to use an existing image or
`cdk.DockerImage.fromBuild(path)` to build a specific image:

```go
lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromAsset(jsii.String("/path/to/handler"), &AssetOptions{
		Bundling: &BundlingOptions{
			Image: awscdk.DockerImage_FromBuild(jsii.String("/path/to/dir/with/DockerFile"), &DockerBuildOptions{
				BuildArgs: map[string]*string{
					"ARG1": jsii.String("value1"),
				},
			}),
			Command: []*string{
				jsii.String("my"),
				jsii.String("cool"),
				jsii.String("command"),
			},
		},
	}),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Handler: jsii.String("index.handler"),
})
```

## Language-specific APIs

Language-specific higher level constructs are provided in separate modules:

* `aws-cdk-lib/aws-lambda-nodejs`: [Github](https://github.com/aws/aws-cdk/tree/main/packages/aws-cdk-lib/aws-lambda-nodejs) & [CDK Docs](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_lambda_nodejs-readme.html)
* `@aws-cdk/aws-lambda-python-alpha`: [Github](https://github.com/aws/aws-cdk/tree/main/packages/%40aws-cdk/aws-lambda-python-alpha) & [CDK Docs](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-python-alpha-readme.html)

## Code Signing

Code signing for AWS Lambda helps to ensure that only trusted code runs in your Lambda functions.
When enabled, AWS Lambda checks every code deployment and verifies that the code package is signed by a trusted source.
For more information, see [Configuring code signing for AWS Lambda](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html).
The following code configures a function with code signing.

Please note the code will not be automatically signed before deployment. To ensure your code is properly signed, you'll need to conduct the code signing process either through the AWS CLI (Command Line Interface) [start-signing-job](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/signer/start-signing-job.html) or by accessing the AWS Signer console.

```go
import "github.com/aws/aws-cdk-go/awscdk"


signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &SigningProfileProps{
	Platform: signer.Platform_AWS_LAMBDA_SHA384_ECDSA(),
})

codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &CodeSigningConfigProps{
	SigningProfiles: []iSigningProfile{
		signingProfile,
	},
})

lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	CodeSigningConfig: CodeSigningConfig,
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

## Runtime updates

Lambda runtime management controls help reduce the risk of impact to your workloads in the rare event of a runtime version incompatibility.
For more information, see [Runtime management controls](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-controls)

```go
lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
	RuntimeManagementMode: lambda.RuntimeManagementMode_AUTO(),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

If you want to set the "Manual" setting, using the ARN of the runtime version as the argument.

```go
lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
	RuntimeManagementMode: lambda.RuntimeManagementMode_Manual(jsii.String("runtimeVersion-arn")),
	Runtime: lambda.Runtime_NODEJS_18_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

## Exclude Patterns for Assets

When using `lambda.Code.fromAsset(path)` an `exclude` property allows you to ignore particular files for assets by providing patterns for file paths to exclude. Note that this has no effect on `Assets` bundled using the `bundling` property.

The `ignoreMode` property can be used with the `exclude` property to specify the file paths to ignore based on the [.gitignore specification](https://git-scm.com/docs/gitignore) or the [.dockerignore specification](https://docs.docker.com/engine/reference/builder/#dockerignore-file). The default behavior is to ignore file paths based on simple glob patterns.

```go
lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-python-handler")), &AssetOptions{
		Exclude: []*string{
			jsii.String("*.ignore"),
		},
		IgnoreMode: awscdk.IgnoreMode_DOCKER,
	}),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Handler: jsii.String("index.handler"),
})
```

You can also write to include only certain files by using a negation.

```go
lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-python-handler")), &AssetOptions{
		Exclude: []*string{
			jsii.String("*"),
			jsii.String("!index.py"),
		},
	}),
	Runtime: lambda.Runtime_PYTHON_3_9(),
	Handler: jsii.String("index.handler"),
})
```

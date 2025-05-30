# AWS CDK CLI Library (deprecated)

<!--BEGIN STABILITY BANNER-->---


![@aws-cdk/cli-lib-lpha: Deprecated](https://img.shields.io/badge/@aws--cdk/cli--lib--alpha-deprectated-red.svg?style=for-the-badge)

> This package has been deprecated in favor of [@aws-cdk/toolkit-lib](https://github.com/aws/aws-cdk-cli/issues/155),
> a newer approach providing similar functionality to what this package offered.
> Please migrate as soon as possible.
> For any migration problems, please open [an issue](https://github.com/aws/aws-cdk-cli/issues/new/choose).
> We are committed to supporting the same feature set that this package offered.

---
<!--END STABILITY BANNER-->

## ⚠️ Deprecated module

This package is has been deprecated.
Already published versions can be used, but no support is provided whatsoever and we will soon stop publishing new versions.

Instead, please use [@aws-cdk/toolkit-lib](https://github.com/aws/aws-cdk-cli/issues/155).

## Overview

Provides a library to interact with the AWS CDK CLI programmatically from jsii supported languages.
Currently the package includes implementations for:

* `cdk deploy`
* `cdk synth`
* `cdk bootstrap`
* `cdk destroy`
* `cdk list`

## Known issues

* **JavaScript/TypeScript only**\
  The jsii packages are currently not in a working state.
* **No useful return values**\
  All output is currently printed to stdout/stderr
* **Missing or Broken options**\
  Some CLI options might not be available in this package or broken

Due to the deprecation of the package, this issues will not be resolved.

## Setup

### AWS CDK app directory

Obtain an `AwsCdkCli` class from an AWS CDK app directory (containing a `cdk.json` file):

```go
cli := awscdkclilibalpha.AwsCdkCli_FromCdkAppDirectory(jsii.String("/path/to/cdk/app"))
```

### Cloud Assembly Directory Producer

You can also create `AwsCdkCli` from a class implementing `ICloudAssemblyDirectoryProducer`.
AWS CDK apps might need to be synthesized multiple times with additional context values before they are ready.

The `produce()` method of the `ICloudAssemblyDirectoryProducer` interface provides this multi-pass ability.
It is invoked with the context values of the current iteration and should use these values to synthesize a Cloud Assembly.
The return value is the path to the assembly directory.

A basic implementation would look like this:

```go
type myProducer struct {
}

func (this *myProducer) produce(context map[string]interface{}) promise {
	app := cdk.NewApp(&AppProps{
		Context: Context,
	})
	stack := cdk.NewStack(app)
	return app.Synth().Directory
}
```

For all features (e.g. lookups) to work correctly, `cdk.App()` must be instantiated with the received `context` values.
Since it is not possible to update the context of an app, it must be created as part of the `produce()` method.

The producer can than be used like this:

```go
cli := awscdkclilibalpha.AwsCdkCli_FromCloudAssemblyDirectoryProducer(NewMyProducer())
```

## Commands

### list

```go
// await this asynchronous method call using a language feature
cli.List()
```

### synth

```go
// await this asynchronous method call using a language feature
cli.Synth(&SynthOptions{
	Stacks: []*string{
		jsii.String("MyTestStack"),
	},
})
```

### bootstrap

```go
// await this asynchronous method call using a language feature
cli.Bootstrap()
```

### deploy

```go
// await this asynchronous method call using a language feature
cli.Deploy(&DeployOptions{
	Stacks: []*string{
		jsii.String("MyTestStack"),
	},
})
```

### destroy

```go
// await this asynchronous method call using a language feature
cli.Destroy(&DestroyOptions{
	Stacks: []*string{
		jsii.String("MyTestStack"),
	},
})
```

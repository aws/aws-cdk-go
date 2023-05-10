# monocdk Experiment

[![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)

An **experiment** to bundle all of the CDK into a single module.

> :warning: Please don't use this module unless you are interested in providing
> feedback about this experience.

## Usage

### Installation

To try out `monocdk` replace all references to CDK Construct
Libraries (most `@aws-cdk/*` packages) in your `package.json` file with a single
entrey referring to `monocdk`.

You also need to add a reference to the `constructs` library, according to the
kind of project you are developing:

* For libraries, model the dependency under `devDependencies` **and** `peerDependencies`
* For apps, model the dependency under `dependencies` only

### Use in your code

#### Classic import

You can use a classic import to get access to each service namespaces:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws-samples/dummy/monocdk"


app := monocdk.Core.NewApp()
stack := monocdk.Core.NewStack(app, jsii.String("MonoCDK-Stack"))

monocdk.Aws_s3.NewBucket(stack, jsii.String("TestBucket"))
```

#### Barrel import

Alternatively, you can use "barrel" imports:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp()
stack := awscdk.Newstack(app, jsii.String("MonoCDK-Stack"))

awscdk.NewBucket(stack, jsii.String("TestBucket"))
```

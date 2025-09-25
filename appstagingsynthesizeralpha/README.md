# App Staging Synthesizer

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This library includes constructs aimed at replacing the current model of bootstrapping and providing
greater control of the bootstrap experience to the CDK user. The important constructs in this library
are as follows:

* the `IStagingResources` interface: a framework for an app-level bootstrap stack that handles
  file assets and docker assets.
* the `DefaultStagingStack`, which is a works-out-of-the-box implementation of the `IStagingResources`
  interface.
* the `AppStagingSynthesizer`, a new CDK synthesizer that will synthesize CDK applications with
  the staging resources provided.

> As this library is `experimental`, there are features that are not yet implemented. Please look
> at the list of [Known Limitations](#known-limitations) before getting started.

To get started, update your CDK App with a new `defaultStackSynthesizer`:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		 // put a unique id here
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
	}),
})
```

This will introduce a `DefaultStagingStack` in your CDK App and staging assets of your App
will live in the resources from that stack rather than the CDK Bootstrap stack.

If you are migrating from a different version of synthesis your updated CDK App will target
the resources in the `DefaultStagingStack` and no longer be tied to the bootstrapped resources
in your account.

## Bootstrap Model

In our default bootstrapping process, when you run `cdk bootstrap aws://<account>/<region>`, the following
resources are created:

* It creates Roles to assume for cross-account deployments and for Pipeline deployments;
* It creates staging resources: a global S3 bucket and global ECR repository to hold CDK assets;
* It creates Roles to write to the S3 bucket and ECR repository;

Because the bootstrapping resources include regional resources, you need to bootstrap
every region you plan to deploy to individually. All assets of all CDK apps deploying
to that account and region will be written to the single S3 Bucket and ECR repository.

By using the synthesizer in this library, instead of the
`DefaultStackSynthesizer`, a different set of staging resources will be created
for every CDK application, and they will be created automatically as part of a
regular deployment, in a separate Stack that is deployed before your application
Stacks. The staging resources will be one S3 bucket, and *one ECR repository per
image*, and Roles necessary to access those buckets and ECR repositories. The
Roles from the default bootstrap stack are still used (though their use can be
turned off).

This has the following advantages:

* Because staging resources are now application-specific, they can be fully cleaned up when you clean up
  the application.
* Because there is now one ECR repository per image instead of one ECR repository for all images, it is
  possible to effectively use ECR life cycle rules (for example, retain only the most recent 5 images)
  to cut down on storage costs.
* Resources between separate CDK Apps are separated so they can be cleaned up and lifecycle
  controlled individually.
* Because the only shared bootstrapping resources required are Roles, which are global resources,
  you now only need to bootstrap every account in one Region (instead of every Region). This makes it
  easier to do with CloudFormation StackSets.

For the deployment roles, this synthesizer still uses the Roles from the default
bootstrap stack, and nothing else. The staging resources from that bootstrap
stack will be unused. You can customize the template to remove those resources
if you prefer.  In the future, we will provide a bootstrap stack template with
only those Roles, specifically for use with this synthesizer.

## Using the Default Staging Stack per Environment

The most common use case will be to use the built-in default resources. In this scenario, the
synthesizer will create a new Staging Stack in each environment the CDK App is deployed to store
its staging resources. To use this kind of synthesizer, use `AppStagingSynthesizer.defaultResources()`.

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,

		// The following line is optional. By default it is assumed you have bootstrapped in the same
		// region(s) as the stack(s) you are deploying.
		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_DefaultBootstrapRoles(&DefaultBootstrapRolesOptions{
			BootstrapRegion: jsii.String("us-east-1"),
		}),
	}),
})
```

Every CDK App that uses the `DefaultStagingStack` must include an `appId`. This should
be an identifier unique to the app and is used to differentiate staging resources associated
with the app.

### Default Staging Stack

The Default Staging Stack includes all the staging resources necessary for CDK Assets. The below example
is of a CDK App using the `AppStagingSynthesizer` and creating a file asset for the Lambda Function
source code. As part of the `DefaultStagingStack`, an S3 bucket and IAM role will be created that will be
used to upload the asset to S3.

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
	}),
})

stack := awscdk.Newstack(app, jsii.String("my-stack"))

lambda.NewFunction(stack, jsii.String("lambda"), &FunctionProps{
	Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("assets"))),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_PYTHON_3_9(),
})

app.Synth()
```

### Custom Roles

You can customize some or all of the roles you'd like to use in the synthesizer as well,
if all you need is to supply custom roles (and not change anything else in the `DefaultStagingStack`):

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_SpecifyRoles(&BootstrapRoles{
			CloudFormationExecutionRole: appstagingsynthesizeralpha.BootstrapRole_FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Execute")),
			DeploymentRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Deploy")),
			LookupRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Lookup")),
		}),
	}),
})
```

Or, you can ask to use the CLI credentials that exist at deploy-time.
These credentials must have the ability to perform CloudFormation calls,
lookup resources in your account, and perform CloudFormation deployment.
For a full list of what is necessary, see `LookupRole`, `DeploymentActionRole`,
and `CloudFormationExecutionRole` in the
[bootstrap template](https://github.com/aws/aws-cdk-cli/blob/main/packages/aws-cdk/lib/api/bootstrap/bootstrap-template.yaml).

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_CliCredentials(),
	}),
})
```

The default staging stack will create roles to publish to the S3 bucket and ECR repositories,
assumable by the deployment role. You can also specify an existing IAM role for the
`fileAssetPublishingRole` or `imageAssetPublishingRole`:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		FileAssetPublishingRole: appstagingsynthesizeralpha.BootstrapRole_FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/S3Access")),
		ImageAssetPublishingRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/ECRAccess")),
	}),
})
```

### Deploy Time S3 Assets

There are two types of assets:

* Assets used only during deployment. These are used to hand off a large piece of data to another
  service, that will make a private copy of that data. After deployment, the asset is only necessary for
  a potential future rollback.
* Assets accessed throughout the running life time of the application.

Examples of assets that are only used at deploy time are CloudFormation Templates and Lambda Code
bundles. Examples of assets accessed throughout the life time of the application are script files
downloaded to run in a CodeBuild Project, or on EC2 instance startup. ECR images are always application
life-time assets. S3 deploy time assets are stored with a `deploy-time/` prefix, and a lifecycle rule will collect them after a configurable number of days.

Lambda assets are by default marked as deploy time assets:

```go
var stack stack

lambda.NewFunction(stack, jsii.String("lambda"), &FunctionProps{
	Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("assets"))),
	 // lambda marks deployTime = true
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_PYTHON_3_9(),
})
```

Or, if you want to create your own deploy time asset:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var stack stack

asset := awscdk.NewAsset(stack, jsii.String("deploy-time-asset"), &AssetProps{
	DeployTime: jsii.Boolean(true),
	Path: path.join(__dirname, jsii.String("deploy-time-asset")),
})
```

By default, we store deploy time assets for 30 days, but you can change this number by specifying
`deployTimeFileAssetLifetime`. The number you specify here is how long you will be able to roll back
to a previous version of an application just by doing a CloudFormation deployment with the old
template, without rebuilding and republishing assets.

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		DeployTimeFileAssetLifetime: awscdk.Duration_Days(jsii.Number(100)),
	}),
})
```

### Lifecycle Rules on ECR Repositories

By default, we store a maximum of 3 revisions of a particular docker image asset. This allows
for smooth faciliation of rollback scenarios where we may reference previous versions of an
image. When more than 3 revisions of an asset exist in the ECR repository, the oldest one is
purged.

To change the number of revisions stored, use `imageAssetVersionCount`:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		ImageAssetVersionCount: jsii.Number(10),
	}),
})
```

### Auto Delete Staging Assets on Deletion

By default, the staging resources will be cleaned up on stack deletion. That means that the
S3 Bucket and ECR Repositories are set to `RemovalPolicy.DESTROY` and have `autoDeleteObjects`
or `emptyOnDelete` turned on. This creates custom resources under the hood to facilitate
cleanup. To turn this off, specify `autoDeleteStagingAssets: false`.

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
		AppId: jsii.String("my-app-id"),
		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
		AutoDeleteStagingAssets: jsii.Boolean(false),
	}),
})
```

### Staging Bucket Encryption

You must explicitly specify the encryption type for the staging bucket via the `stagingBucketEncryption` property. In
future versions of this package, the default will be `BucketEncryption.S3_MANAGED`.

In previous versions of this package, the default was to use KMS encryption for the staging bucket. KMS keys cost
$1/month, which could result in unexpected costs for users who are not aware of this. As we stabilize this module
we intend to make the default S3-managed encryption, which is free. However, the migration path from KMS to S3
managed encryption for existing buckets is not straightforward. Therefore, for now, this property is required.

If you have an existing staging bucket encrypted with a KMS key, you will likely want to set this property to
`BucketEncryption.KMS`. If you are creating a new staging bucket, you can set this property to
`BucketEncryption.S3_MANAGED` to avoid the cost of a KMS key.

You can learn more about choosing a bucket encryption type in the
[S3 documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/serv-side-encryption.html).

## Using a Custom Staging Stack per Environment

If you want to customize some behavior that is not configurable via properties,
you can implement your own class that implements `IStagingResources`. To get a head start,
you can subclass `DefaultStagingStack`.

```go
type customStagingStackOptions struct {
	defaultStagingStackOptions
}

type customStagingStack struct {
	defaultStagingStack
}
```

Or you can roll your own staging resources from scratch, as long as it implements `IStagingResources`.

```go
type customStagingStackProps struct {
	stackProps
}

type customStagingStack struct {
	stack
}

func newCustomStagingStack(scope construct, id *string, props customStagingStackProps) *customStagingStack {
	this := &customStagingStack{}
	newStack_Override(this, scope, id, props)
	return this
}

func (this *customStagingStack) addFile(asset fileAssetSource) fileStagingLocation {
	return &fileStagingLocation{
		BucketName: jsii.String("amzn-s3-demo-bucket"),
		AssumeRoleArn: jsii.String("myArn"),
		DependencyStack: this,
	}
}

func (this *customStagingStack) addDockerImage(asset dockerImageAssetSource) imageStagingLocation {
	return &imageStagingLocation{
		RepoName: jsii.String("myRepo"),
		AssumeRoleArn: jsii.String("myArn"),
		DependencyStack: this,
	}
}
```

Using your custom staging resources means implementing a `CustomFactory` class and calling the
`AppStagingSynthesizer.customFactory()` static method. This has the benefit of providing a
custom Staging Stack that can be created in every environment the CDK App is deployed to.

```go
type customFactory struct {
}

func (this *customFactory) obtainStagingResources(stack stack, context obtainStagingResourcesContext) customStagingStack {
	myApp := awscdk.App_Of(*stack)

	return NewCustomStagingStack(myApp, fmt.Sprintf("CustomStagingStack-%v", *context.EnvironmentString), &customStagingStackProps{
	})
}

app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_CustomFactory(&CustomFactoryOptions{
		Factory: NewCustomFactory(),
		OncePerEnv: jsii.Boolean(true),
	}),
})
```

## Using an Existing Staging Stack

Use `AppStagingSynthesizer.customResources()` to supply an existing stack as the Staging Stack.
Make sure that the custom stack you provide implements `IStagingResources`.

```go
resourceApp := awscdk.NewApp()
resources := NewCustomStagingStack(resourceApp, jsii.String("CustomStagingStack"), &customStagingStackProps{
})

app := awscdk.NewApp(&AppProps{
	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_CustomResources(&CustomResourcesOptions{
		Resources: *Resources,
	}),
})
```

## Known Limitations

Since this module is experimental, there are some known limitations:

* Currently this module does not support CDK Pipelines. You must deploy CDK Apps using this
  synthesizer via `cdk deploy`. Please upvote [this issue](https://github.com/aws/aws-cdk/issues/26118)
  to indicate you want this.
* This synthesizer only needs a bootstrap stack with Roles, without staging resources. We
  haven't written such a bootstrap stack yet; at the moment you can use the existing modern
  bootstrap stack, the staging resources in them will just go unused. You can customize the
  template to remove them if desired.
* Due to limitations on the CloudFormation template size, CDK Applications can have
  at most 20 independent ECR images. Please upvote [this issue](https://github.com/aws/aws-cdk/issues/26119)
  if you need more than this.

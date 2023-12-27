# AWS CodeBuild Construct Library

AWS CodeBuild is a fully managed continuous integration service that compiles
source code, runs tests, and produces software packages that are ready to
deploy. With CodeBuild, you don’t need to provision, manage, and scale your own
build servers. CodeBuild scales continuously and processes multiple builds
concurrently, so your builds are not left waiting in a queue. You can get
started quickly by using prepackaged build environments, or you can create
custom build environments that use your own build tools. With CodeBuild, you are
charged by the minute for the compute resources you use.

## Source

Build projects are usually associated with a *source*, which is specified via
the `source` property which accepts a class that extends the `Source`
abstract base class.
The default is to have no source associated with the build project;
the `buildSpec` option is required in that case.

Here's a CodeBuild project with no source which simply prints `Hello, CodeBuild!`:

```go
codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("echo \"Hello, CodeBuild!\""),
				},
			},
		},
	}),
})
```

### `CodeCommitSource`

Use an AWS CodeCommit repository as the source of this build:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"


repository := codecommit.NewRepository(this, jsii.String("MyRepo"), &RepositoryProps{
	RepositoryName: jsii.String("foo"),
})
codebuild.NewProject(this, jsii.String("MyFirstCodeCommitProject"), &ProjectProps{
	Source: codebuild.Source_CodeCommit(&CodeCommitSourceProps{
		Repository: *Repository,
	}),
})
```

### `S3Source`

Create a CodeBuild project with an S3 bucket as the source:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	Source: codebuild.Source_S3(&S3SourceProps{
		Bucket: bucket,
		Path: jsii.String("path/to/file.zip"),
	}),
})
```

The CodeBuild role will be granted to read just the given path from the given `bucket`.

### `GitHubSource` and `GitHubEnterpriseSource`

These source types can be used to build code from a GitHub repository.
Example:

```go
gitHubSource := codebuild.Source_GitHub(&GitHubSourceProps{
	Owner: jsii.String("awslabs"),
	Repo: jsii.String("aws-cdk"),
	Webhook: jsii.Boolean(true),
	 // optional, default: true if `webhookFilters` were provided, false otherwise
	WebhookTriggersBatchBuild: jsii.Boolean(true),
	 // optional, default is false
	WebhookFilters: []filterGroup{
		codebuild.*filterGroup_InEventOf(codebuild.EventAction_PUSH).AndBranchIs(jsii.String("main")).AndCommitMessageIs(jsii.String("the commit message")),
	},
})
```

To provide GitHub credentials, please either go to AWS CodeBuild Console to connect
or call `ImportSourceCredentials` to persist your personal access token.
Example:

```console
aws codebuild import-source-credentials --server-type GITHUB --auth-type PERSONAL_ACCESS_TOKEN --token <token_value>
```

### `BitBucketSource`

This source type can be used to build code from a BitBucket repository.

```go
bbSource := codebuild.Source_BitBucket(&BitBucketSourceProps{
	Owner: jsii.String("owner"),
	Repo: jsii.String("repo"),
})
```

### For all Git sources

For all Git sources, you can fetch submodules while cloning git repo.

```go
gitHubSource := codebuild.Source_GitHub(&GitHubSourceProps{
	Owner: jsii.String("awslabs"),
	Repo: jsii.String("aws-cdk"),
	FetchSubmodules: jsii.Boolean(true),
})
```

## BuildSpec

The build spec can be provided from a number of different sources

### File path relative to the root of the source

You can specify a specific filename that exists within the project's source artifact to use as the buildspec.

```go
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromSourceFilename(jsii.String("my-buildspec.yml")),
	Source: codebuild.Source_GitHub(&GitHubSourceProps{
		Owner: jsii.String("awslabs"),
		Repo: jsii.String("aws-cdk"),
	}),
})
```

This will use `my-buildspec.yml` file within the `awslabs/aws-cdk` repository as the build spec.

### File within the CDK project codebuild

You can also specify a file within your cdk project directory to use as the buildspec.

```go
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromAsset(jsii.String("my-buildspec.yml")),
})
```

This file will be uploaded to S3 and referenced from the codebuild project.

### Inline object

```go
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
})
```

This will result in the buildspec being rendered as JSON within the codebuild project, if you prefer it to be rendered as YAML, use `fromObjectToYaml`.

```go
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObjectToYaml(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
})
```

## Artifacts

CodeBuild Projects can produce Artifacts and upload them to S3. For example:

```go
var bucket bucket


project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
	Artifacts: codebuild.Artifacts_S3(&S3ArtifactsProps{
		Bucket: *Bucket,
		IncludeBuildId: jsii.Boolean(false),
		PackageZip: jsii.Boolean(true),
		Path: jsii.String("another/path"),
		Identifier: jsii.String("AddArtifact1"),
	}),
})
```

Because we've not set the `name` property, this example will set the
`overrideArtifactName` parameter, and produce an artifact named as defined in
the Buildspec file, uploaded to an S3 bucket (`bucket`). The path will be
`another/path` and the artifact will be a zipfile.

## CodePipeline

To add a CodeBuild Project as an Action to CodePipeline,
use the `PipelineProject` class instead of `Project`.
It's a simple class that doesn't allow you to specify `sources`,
`secondarySources`, `artifacts` or `secondaryArtifacts`,
as these are handled by setting input and output CodePipeline `Artifact` instances on the Action,
instead of setting them on the Project.

```go
project := codebuild.NewPipelineProject(this, jsii.String("Project"), &PipelineProjectProps{
})
```

For more details, see the readme of the `@aws-cdk/aws-codepipeline-actions` package.

## Caching

You can save time when your project builds by using a cache. A cache can store reusable pieces of your build environment and use them across multiple builds. Your build project can use one of two types of caching: Amazon S3 or local. In general, S3 caching is a good option for small and intermediate build artifacts that are more expensive to build than to download. Local caching is a good option for large intermediate build artifacts because the cache is immediately available on the build host.

### S3 Caching

With S3 caching, the cache is stored in an S3 bucket which is available
regardless from what CodeBuild instance gets selected to run your CodeBuild job
on. When using S3 caching, you must also add in a `cache` section to your
buildspec which indicates the files to be cached:

```go
var myCachingBucket bucket


codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Source: codebuild.Source_BitBucket(&BitBucketSourceProps{
		Owner: jsii.String("awslabs"),
		Repo: jsii.String("aws-cdk"),
	}),

	Cache: codebuild.Cache_Bucket(myCachingBucket),

	// BuildSpec with a 'cache' section necessary for S3 caching. This can
	// also come from 'buildspec.yml' in your source.
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("..."),
				},
			},
		},
		"cache": map[string][]*string{
			"paths": []*string{
				jsii.String("/root/cachedir/**/*"),
			},
		},
	}),
})
```

Note that two different CodeBuild Projects using the same S3 bucket will *not*
share their cache: each Project will get a unique file in the S3 bucket to store
the cache in.

### Local Caching

With local caching, the cache is stored on the codebuild instance itself. This
is simple, cheap and fast, but CodeBuild cannot guarantee a reuse of instance
and hence cannot guarantee cache hits. For example, when a build starts and
caches files locally, if two subsequent builds start at the same time afterwards
only one of those builds would get the cache. Three different cache modes are
supported, which can be turned on individually.

* `LocalCacheMode.SOURCE` caches Git metadata for primary and secondary sources.
* `LocalCacheMode.DOCKER_LAYER` caches existing Docker layers.
* `LocalCacheMode.CUSTOM` caches directories you specify in the buildspec file.

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Source: codebuild.Source_GitHubEnterprise(&GitHubEnterpriseSourceProps{
		HttpsCloneUrl: jsii.String("https://my-github-enterprise.com/owner/repo"),
	}),

	// Enable Docker AND custom caching
	Cache: codebuild.Cache_Local(codebuild.LocalCacheMode_DOCKER_LAYER, codebuild.LocalCacheMode_CUSTOM),

	// BuildSpec with a 'cache' section necessary for 'CUSTOM' caching. This can
	// also come from 'buildspec.yml' in your source.
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("..."),
				},
			},
		},
		"cache": map[string][]*string{
			"paths": []*string{
				jsii.String("/root/cachedir/**/*"),
			},
		},
	}),
})
```

## Environment

By default, projects use a small instance with an Ubuntu 18.04 image. You
can use the `environment` property to customize the build environment:

* `buildImage` defines the Docker image used. See [Images](#images) below for
  details on how to define build images.
* `certificate` defines the location of a PEM encoded certificate to import.
* `computeType` defines the instance type used for the build.
* `privileged` can be set to `true` to allow privileged access.
* `environmentVariables` can be set at this level (and also at the project
  level).

## Images

The CodeBuild library supports both Linux and Windows images via the
`LinuxBuildImage` (or `LinuxArmBuildImage`), and `WindowsBuildImage` classes, respectively.
With the introduction of Lambda compute support, the `LinuxLambdaBuildImage ` (or `LinuxArmLambdaBuildImage`) class
is available for specifying Lambda-compatible images.

You can specify one of the predefined Windows/Linux images by using one
of the constants such as `WindowsBuildImage.WIN_SERVER_CORE_2019_BASE`,
`WindowsBuildImage.WINDOWS_BASE_2_0`, `LinuxBuildImage.STANDARD_2_0`,
`LinuxBuildImage.AMAZON_LINUX_2_5`, `LinuxArmBuildImage.AMAZON_LINUX_2_ARM`,
`LinuxLambdaBuildImage.AMAZON_LINUX_2_NODE_18` or `LinuxArmLambdaBuildImage.AMAZON_LINUX_2_NODE_18`.

Alternatively, you can specify a custom image using one of the static methods on
`LinuxBuildImage`:

* `LinuxBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])` to reference an image in any public or private Docker registry.
* `LinuxBuildImage.fromEcrRepository(repo[, tag])` to reference an image available in an
  ECR repository.
* `LinuxBuildImage.fromAsset(parent, id, props)` to use an image created from a
  local asset.
* `LinuxBuildImage.fromCodeBuildImageId(id)` to reference a pre-defined, CodeBuild-provided Docker image.

or one of the corresponding methods on `WindowsBuildImage`:

* `WindowsBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }, imageType])`
* `WindowsBuildImage.fromEcrRepository(repo[, tag, imageType])`
* `WindowsBuildImage.fromAsset(parent, id, props, [, imageType])`

or one of the corresponding methods on `LinuxArmBuildImage`:

* `LinuxArmBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])`
* `LinuxArmBuildImage.fromEcrRepository(repo[, tag])`

**Note:** You cannot specify custom images on `LinuxLambdaBuildImage` or `LinuxArmLambdaBuildImage` images.

Note that the `WindowsBuildImage` version of the static methods accepts an optional parameter of type `WindowsImageType`,
which can be either `WindowsImageType.STANDARD`, the default, or `WindowsImageType.SERVER_2019`:

```go
var ecrRepository repository


codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.WindowsBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.WindowsImageType_SERVER_2019),
		// optional certificate to include in the build image
		Certificate: &BuildEnvironmentCertificate{
			Bucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
			ObjectKey: jsii.String("path/to/cert.pem"),
		},
	},
})
```

The following example shows how to define an image from a Docker asset:

```go
Environment: &BuildEnvironment{
	BuildImage: codebuild.LinuxBuildImage_FromAsset(this, jsii.String("MyImage"), &DockerImageAssetProps{
		Directory: path.join(__dirname, jsii.String("demo-image")),
	}),
},
```

The following example shows how to define an image from an ECR repository:

```go
Environment: &BuildEnvironment{
	BuildImage: codebuild.LinuxBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0")),
},
```

The following example shows how to define an image from a private docker registry:

```go
Environment: &BuildEnvironment{
	BuildImage: codebuild.LinuxBuildImage_FromDockerRegistry(jsii.String("my-registry/my-repo"), &DockerImageOptions{
		SecretsManagerCredentials: secrets,
	}),
},
```

### GPU images

The class `LinuxGpuBuildImage` contains constants for working with
[AWS Deep Learning Container images](https://aws.amazon.com/releasenotes/available-deep-learning-containers-images):

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_INFERENCE(),
	},
})
```

One complication is that the repositories for the DLC images are in
different accounts in different AWS regions.
In most cases, the CDK will handle providing the correct account for you;
in rare cases (for example, deploying to new regions)
where our information might be out of date,
you can always specify the account
(along with the repository name and tag)
explicitly using the `awsDeepLearningContainersImage` method:

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxGpuBuildImage_AwsDeepLearningContainersImage(jsii.String("tensorflow-inference"), jsii.String("2.1.0-gpu-py36-cu101-ubuntu18.04"), jsii.String("123456789012")),
	},
})
```

Alternatively, you can reference an image available in an ECR repository using the `LinuxGpuBuildImage.fromEcrRepository(repo[, tag])` method.

### Lambda images

The `LinuxLambdaBuildImage` (or `LinuxArmLambdaBuildImage`) class contains constants for working with
Lambda compute:

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxLambdaBuildImage_AMAZON_LINUX_2_NODE_18(),
	},
})
```

> Visit [AWS Lambda compute in AWS CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/lambda.html) for more details.

## Logs

CodeBuild lets you specify an S3 Bucket, CloudWatch Log Group or both to receive logs from your projects.

By default, logs will go to cloudwatch.

### CloudWatch Logs Example

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Logging: &LoggingOptions{
		CloudWatch: &CloudWatchLoggingOptions{
			LogGroup: logs.NewLogGroup(this, jsii.String("MyLogGroup")),
		},
	},
})
```

### S3 Logs Example

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Logging: &LoggingOptions{
		S3: &S3LoggingOptions{
			Bucket: s3.NewBucket(this, jsii.String("LogBucket")),
		},
	},
})
```

## Debugging builds interactively using SSM Session Manager

Integration with SSM Session Manager makes it possible to add breakpoints to your
build commands, pause the build there and log into the container to interactively
debug the environment.

To do so, you need to:

* Create the build with `ssmSessionPermissions: true`.
* Use a build image with SSM agent installed and configured (default CodeBuild images come with the image preinstalled).
* Start the build with [debugSessionEnabled](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuild.html#CodeBuild-StartBuild-request-debugSessionEnabled) set to true.

If these conditions are met, execution of the command `codebuild-breakpoint`
will suspend your build and allow you to attach a Session Manager session from
the CodeBuild console.

For more information, see [View a running build in Session
Manager](https://docs.aws.amazon.com/codebuild/latest/userguide/session-manager.html)
in the CodeBuild documentation.

Example:

```go
codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
	},
	SsmSessionPermissions: jsii.Boolean(true),
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("codebuild-breakpoint"),
					jsii.String("./my-build.sh"),
				},
			},
		},
	}),
})
```

## Credentials

CodeBuild allows you to store credentials used when communicating with various sources,
like GitHub:

```go
codebuild.NewGitHubSourceCredentials(this, jsii.String("CodeBuildGitHubCreds"), &GitHubSourceCredentialsProps{
	AccessToken: awscdk.SecretValue_SecretsManager(jsii.String("my-token")),
})
```

and BitBucket:

```go
codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &BitBucketSourceCredentialsProps{
	Username: awscdk.SecretValue_SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
		JsonField: jsii.String("username"),
	}),
	Password: awscdk.SecretValue_*SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
		JsonField: jsii.String("password"),
	}),
})
```

**Note**: the credentials are global to a given account in a given region -
they are not defined per CodeBuild project.
CodeBuild only allows storing a single credential of a given type
(GitHub, GitHub Enterprise or BitBucket)
in a given account in a given region -
any attempt to save more than one will result in an error.
You can use the [`list-source-credentials` AWS CLI operation](https://docs.aws.amazon.com/cli/latest/reference/codebuild/list-source-credentials.html)
to inspect what credentials are stored in your account.

## Test reports

You can specify a test report in your buildspec:

```go
project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		// ...
		"reports": map[string]map[string]*string{
			"myReport": map[string]*string{
				"files": jsii.String("**/*"),
				"base-directory": jsii.String("build/test-results"),
			},
		},
	}),
})
```

This will create a new test report group,
with the name `<ProjectName>-myReport`.

The project's role in the CDK will always be granted permissions to create and use report groups
with names starting with the project's name;
if you'd rather not have those permissions added,
you can opt out of it when creating the project:

```go
var source source


project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Source: Source,
	GrantReportGroupPermissions: jsii.Boolean(false),
})
```

Alternatively, you can specify an ARN of an existing resource group,
instead of a simple name, in your buildspec:

```go
var source source


// create a new ReportGroup
reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"))

project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Source: Source,
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		// ...
		"reports": map[string]map[string]*string{
			reportGroup.reportGroupArn: map[string]*string{
				"files": jsii.String("**/*"),
				"base-directory": jsii.String("build/test-results"),
			},
		},
	}),
})
```

For a code coverage report, you can specify a report group with the code coverage report group type.

```go
var source source


// create a new ReportGroup
reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"), &ReportGroupProps{
	Type: codebuild.ReportGroupType_CODE_COVERAGE,
})

project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
	Source: Source,
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		// ...
		"reports": map[string]map[string]*string{
			reportGroup.reportGroupArn: map[string]*string{
				"files": jsii.String("**/*"),
				"base-directory": jsii.String("build/coverage-report.xml"),
				"file-format": jsii.String("JACOCOXML"),
			},
		},
	}),
})
```

If you specify a report group, you need to grant the project's role permissions to write reports to that report group:

```go
var project project
var reportGroup reportGroup


reportGroup.grantWrite(project)
```

The created policy will adjust to the report group type. If no type is specified when creating the report group the created policy will contain the action for the test report group type.

For more information on the test reports feature,
see the [AWS CodeBuild documentation](https://docs.aws.amazon.com/codebuild/latest/userguide/test-reporting.html).

## Events

CodeBuild projects can be used either as a source for events or be triggered
by events via an event rule.

### Using Project as an event target

The `aws-cdk-lib/aws-events-targets.CodeBuildProject` allows using an AWS CodeBuild
project as a AWS CloudWatch event rule target:

```go
// start build when a commit is pushed
import codecommit "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var codeCommitRepository repository
var project project


codeCommitRepository.onCommit(jsii.String("OnCommit"), &OnCommitOptions{
	Target: targets.NewCodeBuildProject(project),
})
```

### Using Project as an event source

To define Amazon CloudWatch event rules for build projects, use one of the `onXxx`
methods:

```go
import targets "github.com/aws/aws-cdk-go/awscdk"
var fn function
var project project


rule := project.onStateChange(jsii.String("BuildStateChange"), &OnEventOptions{
	Target: targets.NewLambdaFunction(fn),
})
```

## CodeStar Notifications

To define CodeStar Notification rules for Projects, use one of the `notifyOnXxx()` methods.
They are very similar to `onXxx()` methods for CloudWatch events:

```go
import chatbot "github.com/aws/aws-cdk-go/awscdk"

var project project


target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
})

rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
```

## Secondary sources and artifacts

CodeBuild Projects can get their sources from multiple places, and produce
multiple outputs. For example:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"
var repo repository
var bucket bucket


project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	SecondarySources: []iSource{
		codebuild.Source_CodeCommit(&CodeCommitSourceProps{
			Identifier: jsii.String("source2"),
			Repository: repo,
		}),
	},
	SecondaryArtifacts: []iArtifacts{
		codebuild.Artifacts_S3(&S3ArtifactsProps{
			Identifier: jsii.String("artifact2"),
			Bucket: bucket,
			Path: jsii.String("some/path"),
			Name: jsii.String("file.zip"),
		}),
	},
})
```

Note that the `identifier` property is required for both secondary sources and
artifacts.

The contents of the secondary source is available to the build under the
directory specified by the `CODEBUILD_SRC_DIR_<identifier>` environment variable
(so, `CODEBUILD_SRC_DIR_source2` in the above case).

The secondary artifacts have their own section in the buildspec, under the
regular `artifacts` one. Each secondary artifact has its own section, beginning
with their identifier.

So, a buildspec for the above Project could look something like this:

```go
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	// secondary sources and artifacts as above...
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("cd $CODEBUILD_SRC_DIR_source2"),
					jsii.String("touch output2.txt"),
				},
			},
		},
		"artifacts": map[string]map[string]map[string]interface{}{
			"secondary-artifacts": map[string]map[string]interface{}{
				"artifact2": map[string]interface{}{
					"base-directory": jsii.String("$CODEBUILD_SRC_DIR_source2"),
					"files": []*string{
						jsii.String("output2.txt"),
					},
				},
			},
		},
	}),
})
```

### Definition of VPC configuration in CodeBuild Project

Typically, resources in an VPC are not accessible by AWS CodeBuild. To enable
access, you must provide additional VPC-specific configuration information as
part of your CodeBuild project configuration. This includes the VPC ID, the
VPC subnet IDs, and the VPC security group IDs. VPC-enabled builds are then
able to access resources inside your VPC.

For further Information see https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html

**Use Cases**
VPC connectivity from AWS CodeBuild builds makes it possible to:

* Run integration tests from your build against data in an Amazon RDS database that's isolated on a private subnet.
* Query data in an Amazon ElastiCache cluster directly from tests.
* Interact with internal web services hosted on Amazon EC2, Amazon ECS, or services that use internal Elastic Load Balancing.
* Retrieve dependencies from self-hosted, internal artifact repositories, such as PyPI for Python, Maven for Java, and npm for Node.js.
* Access objects in an Amazon S3 bucket configured to allow access through an Amazon VPC endpoint only.
* Query external web services that require fixed IP addresses through the Elastic IP address of the NAT gateway or NAT instance associated with your subnet(s).

Your builds can access any resource that's hosted in your VPC.

**Enable Amazon VPC Access in your CodeBuild Projects**

Pass the VPC when defining your Project, then make sure to
give the CodeBuild's security group the right permissions
to access the resources that it needs by using the
`connections` object.

For example:

```go
var loadBalancer applicationLoadBalancer


vpc := ec2.NewVpc(this, jsii.String("MyVPC"))
project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	Vpc: vpc,
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
	}),
})

project.connections.AllowTo(loadBalancer, ec2.Port_Tcp(jsii.Number(443)))
```

## Project File System Location EFS

Add support for CodeBuild to build on AWS EFS file system mounts using
the new ProjectFileSystemLocation.
The `fileSystemLocations` property which accepts a list `ProjectFileSystemLocation`
as represented by the interface `IFileSystemLocations`.
The only supported file system type is `EFS`.

For example:

```go
codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
	FileSystemLocations: []iFileSystemLocation{
		codebuild.FileSystemLocation_Efs(&EfsFileSystemLocationProps{
			Identifier: jsii.String("myidentifier2"),
			Location: jsii.String("myclodation.mydnsroot.com:/loc"),
			MountPoint: jsii.String("/media"),
			MountOptions: jsii.String("opts"),
		}),
	},
})
```

Here's a CodeBuild project with a simple example that creates a project mounted on AWS EFS:

[Minimal Example](./test/integ.project-file-system-location.ts)

## Batch builds

To enable batch builds you should call `enableBatchBuilds()` on the project instance.

It returns an object containing the batch service role that was created,
or `undefined` if batch builds could not be enabled, for example if the project was imported.

```go
var source source


project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	Source: Source,
})

if project.EnableBatchBuilds() {
	fmt.Println("Batch builds were enabled")
}
```

## Timeouts

There are two types of timeouts that can be set when creating your Project.
The `timeout` property can be used to set an upper limit on how long your Project is able to run without being marked as completed.
The default is 60 minutes.
An example of overriding the default follows.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	Timeout: awscdk.Duration_Minutes(jsii.Number(90)),
})
```

The `queuedTimeout` property can be used to set an upper limit on how your Project remains queued to run.
There is no default value for this property.
As an example, to allow your Project to queue for up to thirty (30) minutes before the build fails,
use the following code.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	QueuedTimeout: awscdk.Duration_Minutes(jsii.Number(30)),
})
```

## Limiting concurrency

By default if a new build is triggered it will be run even if there is a previous build already in progress.
It is possible to limit the maximum concurrent builds to value between 1 and the account specific maximum limit.
By default there is no explicit limit.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	ConcurrentBuildLimit: jsii.Number(1),
})
```

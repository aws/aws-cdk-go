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
codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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


repository := codecommit.NewRepository(this, jsii.String("MyRepo"), &repositoryProps{
	repositoryName: jsii.String("foo"),
})
codebuild.NewProject(this, jsii.String("MyFirstCodeCommitProject"), &projectProps{
	source: codebuild.source.codeCommit(&codeCommitSourceProps{
		repository: repository,
	}),
})
```

### `S3Source`

Create a CodeBuild project with an S3 bucket as the source:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	source: codebuild.source.s3(&s3SourceProps{
		bucket: bucket,
		path: jsii.String("path/to/file.zip"),
	}),
})
```

The CodeBuild role will be granted to read just the given path from the given `bucket`.

### `GitHubSource` and `GitHubEnterpriseSource`

These source types can be used to build code from a GitHub repository.
Example:

```go
gitHubSource := codebuild.source.gitHub(&gitHubSourceProps{
	owner: jsii.String("awslabs"),
	repo: jsii.String("aws-cdk"),
	webhook: jsii.Boolean(true),
	 // optional, default: true if `webhookFilters` were provided, false otherwise
	webhookTriggersBatchBuild: jsii.Boolean(true),
	 // optional, default is false
	webhookFilters: []filterGroup{
		codebuild.*filterGroup.inEventOf(codebuild.eventAction_PUSH).andBranchIs(jsii.String("main")).andCommitMessageIs(jsii.String("the commit message")),
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
bbSource := codebuild.source.bitBucket(&bitBucketSourceProps{
	owner: jsii.String("owner"),
	repo: jsii.String("repo"),
})
```

### For all Git sources

For all Git sources, you can fetch submodules while cloning git repo.

```go
gitHubSource := codebuild.source.gitHub(&gitHubSourceProps{
	owner: jsii.String("awslabs"),
	repo: jsii.String("aws-cdk"),
	fetchSubmodules: jsii.Boolean(true),
})
```

## Artifacts

CodeBuild Projects can produce Artifacts and upload them to S3. For example:

```go
var bucket bucket


project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
	artifacts: codebuild.artifacts.s3(&s3ArtifactsProps{
		bucket: bucket,
		includeBuildId: jsii.Boolean(false),
		packageZip: jsii.Boolean(true),
		path: jsii.String("another/path"),
		identifier: jsii.String("AddArtifact1"),
	}),
})
```

If you'd prefer your buildspec to be rendered as YAML in the template,
use the `fromObjectToYaml()` method instead of `fromObject()`.

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
project := codebuild.NewPipelineProject(this, jsii.String("Project"), &pipelineProjectProps{
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


codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	source: codebuild.source.bitBucket(&bitBucketSourceProps{
		owner: jsii.String("awslabs"),
		repo: jsii.String("aws-cdk"),
	}),

	cache: codebuild.cache.bucket(myCachingBucket),

	// BuildSpec with a 'cache' section necessary for S3 caching. This can
	// also come from 'buildspec.yml' in your source.
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	source: codebuild.source.gitHubEnterprise(&gitHubEnterpriseSourceProps{
		httpsCloneUrl: jsii.String("https://my-github-enterprise.com/owner/repo"),
	}),

	// Enable Docker AND custom caching
	cache: codebuild.cache.local(codebuild.localCacheMode_DOCKER_LAYER, codebuild.*localCacheMode_CUSTOM),

	// BuildSpec with a 'cache' section necessary for 'CUSTOM' caching. This can
	// also come from 'buildspec.yml' in your source.
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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

You can specify one of the predefined Windows/Linux images by using one
of the constants such as `WindowsBuildImage.WIN_SERVER_CORE_2019_BASE`,
`WindowsBuildImage.WINDOWS_BASE_2_0`, `LinuxBuildImage.STANDARD_2_0`, or
`LinuxArmBuildImage.AMAZON_LINUX_2_ARM`.

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

* `LinuxArmBuildImage.fromEcrRepository(repo[, tag])`

Note that the `WindowsBuildImage` version of the static methods accepts an optional parameter of type `WindowsImageType`,
which can be either `WindowsImageType.STANDARD`, the default, or `WindowsImageType.SERVER_2019`:

```go
var ecrRepository repository


codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.windowsBuildImage.fromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.windowsImageType_SERVER_2019),
		// optional certificate to include in the build image
		certificate: &buildEnvironmentCertificate{
			bucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
			objectKey: jsii.String("path/to/cert.pem"),
		},
	},
})
```

The following example shows how to define an image from a Docker asset:

```go
environment: &buildEnvironment{
	buildImage: codebuild.linuxBuildImage.fromAsset(this, jsii.String("MyImage"), &dockerImageAssetProps{
		directory: path.join(__dirname, jsii.String("demo-image")),
	}),
},
```

The following example shows how to define an image from an ECR repository:

```go
environment: &buildEnvironment{
	buildImage: codebuild.linuxBuildImage.fromEcrRepository(ecrRepository, jsii.String("v1.0")),
},
```

The following example shows how to define an image from a private docker registry:

```go
environment: &buildEnvironment{
	buildImage: codebuild.linuxBuildImage.fromDockerRegistry(jsii.String("my-registry/my-repo"), &dockerImageOptions{
		secretsManagerCredentials: secrets,
	}),
},
```

### GPU images

The class `LinuxGpuBuildImage` contains constants for working with
[AWS Deep Learning Container images](https://aws.amazon.com/releasenotes/available-deep-learning-containers-images):

```go
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.linuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_INFERENCE(),
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
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.linuxGpuBuildImage.awsDeepLearningContainersImage(jsii.String("tensorflow-inference"), jsii.String("2.1.0-gpu-py36-cu101-ubuntu18.04"), jsii.String("123456789012")),
	},
})
```

Alternatively, you can reference an image available in an ECR repository using the `LinuxGpuBuildImage.fromEcrRepository(repo[, tag])` method.

## Logs

CodeBuild lets you specify an S3 Bucket, CloudWatch Log Group or both to receive logs from your projects.

By default, logs will go to cloudwatch.

### CloudWatch Logs Example

```go
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	logging: &loggingOptions{
		cloudWatch: &cloudWatchLoggingOptions{
			logGroup: logs.NewLogGroup(this, jsii.String("MyLogGroup")),
		},
	},
})
```

### S3 Logs Example

```go
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	logging: &loggingOptions{
		s3: &s3LoggingOptions{
			bucket: s3.NewBucket(this, jsii.String("LogBucket")),
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
codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.linuxBuildImage_STANDARD_6_0(),
	},
	ssmSessionPermissions: jsii.Boolean(true),
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
codebuild.NewGitHubSourceCredentials(this, jsii.String("CodeBuildGitHubCreds"), &gitHubSourceCredentialsProps{
	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-token")),
})
```

and BitBucket:

```go
codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &bitBucketSourceCredentialsProps{
	username: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
		jsonField: jsii.String("username"),
	}),
	password: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
		jsonField: jsii.String("password"),
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
project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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


project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	source: source,
	grantReportGroupPermissions: jsii.Boolean(false),
})
```

Alternatively, you can specify an ARN of an existing resource group,
instead of a simple name, in your buildspec:

```go
var source source


// create a new ReportGroup
reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"))

project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	source: source,
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"), &reportGroupProps{
	type: codebuild.reportGroupType_CODE_COVERAGE,
})

project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
	source: source,
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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

The `@aws-cdk/aws-events-targets.CodeBuildProject` allows using an AWS CodeBuild
project as a AWS CloudWatch event rule target:

```go
// start build when a commit is pushed
import codecommit "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var codeCommitRepository repository
var project project


codeCommitRepository.onCommit(jsii.String("OnCommit"), &onCommitOptions{
	target: targets.NewCodeBuildProject(project),
})
```

### Using Project as an event source

To define Amazon CloudWatch event rules for build projects, use one of the `onXxx`
methods:

```go
import targets "github.com/aws/aws-cdk-go/awscdk"
var fn function
var project project


rule := project.onStateChange(jsii.String("BuildStateChange"), &onEventOptions{
	target: targets.NewLambdaFunction(fn),
})
```

## CodeStar Notifications

To define CodeStar Notification rules for Projects, use one of the `notifyOnXxx()` methods.
They are very similar to `onXxx()` methods for CloudWatch events:

```go
import chatbot "github.com/aws/aws-cdk-go/awscdk"

var project project


target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
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


project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	secondarySources: []iSource{
		codebuild.source.codeCommit(&codeCommitSourceProps{
			identifier: jsii.String("source2"),
			repository: repo,
		}),
	},
	secondaryArtifacts: []iArtifacts{
		codebuild.artifacts.s3(&s3ArtifactsProps{
			identifier: jsii.String("artifact2"),
			bucket: bucket,
			path: jsii.String("some/path"),
			name: jsii.String("file.zip"),
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
project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	// secondary sources and artifacts as above...
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	vpc: vpc,
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
	}),
})

project.connections.allowTo(loadBalancer, ec2.port.tcp(jsii.Number(443)))
```

## Project File System Location EFS

Add support for CodeBuild to build on AWS EFS file system mounts using
the new ProjectFileSystemLocation.
The `fileSystemLocations` property which accepts a list `ProjectFileSystemLocation`
as represented by the interface `IFileSystemLocations`.
The only supported file system type is `EFS`.

For example:

```go
codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),
	fileSystemLocations: []iFileSystemLocation{
		codebuild.fileSystemLocation.efs(&efsFileSystemLocationProps{
			identifier: jsii.String("myidentifier2"),
			location: jsii.String("myclodation.mydnsroot.com:/loc"),
			mountPoint: jsii.String("/media"),
			mountOptions: jsii.String("opts"),
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


project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	source: source,
})

if project.enableBatchBuilds() {
	fmt.Println("Batch builds were enabled")
}
```

## Timeouts

There are two types of timeouts that can be set when creating your Project.
The `timeout` property can be used to set an upper limit on how long your Project is able to run without being marked as completed.
The default is 60 minutes.
An example of overriding the default follows.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	timeout: awscdk.Duration.minutes(jsii.Number(90)),
})
```

The `queuedTimeout` property can be used to set an upper limit on how your Project remains queued to run.
There is no default value for this property.
As an example, to allow your Project to queue for up to thirty (30) minutes before the build fails,
use the following code.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	queuedTimeout: awscdk.Duration.minutes(jsii.Number(30)),
})
```

## Limiting concurrency

By default if a new build is triggered it will be run even if there is a previous build already in progress.
It is possible to limit the maximum concurrent builds to value between 1 and the account specific maximum limit.
By default there is no explicit limit.

```go
codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	concurrentBuildLimit: jsii.Number(1),
})
```

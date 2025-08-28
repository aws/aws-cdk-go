# CDK Pipelines

A construct library for painless Continuous Delivery of CDK applications.

CDK Pipelines is an *opinionated construct library*. It is purpose-built to
deploy one or more copies of your CDK applications using CloudFormation with a
minimal amount of effort on your part. It is *not* intended to support arbitrary
deployment pipelines, and very specifically it is not built to use CodeDeploy to
deploy applications to instances, or deploy your custom-built ECR images to an ECS
cluster directly: use CDK file assets with CloudFormation Init for instances, or
CDK container assets for ECS clusters instead.

Give the CDK Pipelines way of doing things a shot first: you might find it does
everything you need. If you need more control, we recommend you drop down to using the `aws-codepipeline`
construct library directly.

> This module contains two sets of APIs: an **original** and a **modern** version of
> CDK Pipelines. The *modern* API has been updated to be easier to work with and
> customize, and will be the preferred API going forward. The *original* version
> of the API is still available for backwards compatibility, but we recommend migrating
> to the new version if possible.
>
> Compared to the original API, the modern API: has more sensible defaults; is
> more flexible; supports parallel deployments; supports multiple synth inputs;
> allows more control of CodeBuild project generation; supports deployment
> engines other than CodePipeline.
>
> The README for the original API, as well as a migration guide, can be found in
> [our GitHub repository](https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/pipelines/ORIGINAL_API.md).

## At a glance

Deploying your application continuously starts by defining a
`MyApplicationStage`, a subclass of `Stage` that contains the stacks that make
up a single copy of your application.

You then define a `Pipeline`, instantiate as many instances of
`MyApplicationStage` as you want for your test and production environments, with
different parameters for each, and calling `pipeline.addStage()` for each of
them. You can deploy to the same account and Region, or to a different one,
with the same amount of code. The *CDK Pipelines* library takes care of the
details.

CDK Pipelines supports multiple *deployment engines* (see
[Using a different deployment engine](#using-a-different-deployment-engine)),
and comes with a deployment engine that deploys CDK apps using AWS CodePipeline.
To use the CodePipeline engine, define a `CodePipeline` construct. The following
example creates a CodePipeline that deploys an application from GitHub:

```go
/** The stacks for our app are minimally defined here.  The internals of these
 * stacks aren't important, except that DatabaseStack exposes an attribute
 * "table" for a database table it defines, and ComputeStack accepts a reference
 * to this table in its properties.
 */
type databaseStack struct {
	stack
	table tableV2
}

func newDatabaseStack(scope construct, id *string) *databaseStack {
	this := &databaseStack{}
	newStack_Override(this, scope, id)
	this.table = dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
		PartitionKey: &Attribute{
			Name: jsii.String("id"),
			Type: dynamodb.AttributeType_STRING,
		},
	})
	return this
}

type computeProps struct {
	table *tableV2
}

type computeStack struct {
	stack
}

func newComputeStack(scope construct, id *string, props computeProps) *computeStack {
	this := &computeStack{}
	newStack_Override(this, scope, id)
	return this
}

/**
 * Stack to hold the pipeline
 */
type myPipelineStack struct {
	stack
}

func newMyPipelineStack(scope construct, id *string, props stackProps) *myPipelineStack {
	this := &myPipelineStack{}
	newStack_Override(this, scope, id, props)

	pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
			// Use a connection created using the AWS console to authenticate to GitHub
			// Other sources are available.
			Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
				ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
			}),
			Commands: []*string{
				jsii.String("npm ci"),
				jsii.String("npm run build"),
				jsii.String("npx cdk synth"),
			},
		}),
	})

	// 'MyApplication' is defined below. Call `addStage` as many times as
	// necessary with any account and region (may be different from the
	// pipeline's).
	pipeline.AddStage(
	NewMyApplication(this, jsii.String("Prod"), &stageProps{
		Env: &Environment{
			Account: jsii.String("123456789012"),
			Region: jsii.String("eu-west-1"),
		},
	}))
	return this
}

/**
 * Your application
 *
 * May consist of one or more Stacks (here, two)
 *
 * By declaring our DatabaseStack and our ComputeStack inside a Stage,
 * we make sure they are deployed together, or not at all.
 */
type myApplication struct {
	stage
}

func newMyApplication(scope construct, id *string, props stageProps) *myApplication {
	this := &myApplication{}
	newStage_Override(this, scope, id, props)

	dbStack := NewDatabaseStack(this, jsii.String("Database"))
	NewComputeStack(this, jsii.String("Compute"), &computeProps{
		table: dbStack.table,
	})
	return this
}

// In your main file
// In your main file
NewMyPipelineStack(this, jsii.String("PipelineStack"), &stackProps{
	Env: &Environment{
		Account: jsii.String("123456789012"),
		Region: jsii.String("eu-west-1"),
	},
})
```

The pipeline is **self-mutating**, which means that if you add new
application stages in the source code, or new stacks to `MyApplication`, the
pipeline will automatically reconfigure itself to deploy those new stages and
stacks.

(Note that you have to *bootstrap* all environments before the above code
will work, and switch on "Modern synthesis" if you are using
CDKv1. See the section **CDK Environment Bootstrapping** below for
more information).

## Provisioning the pipeline

To provision the pipeline you have defined, make sure the target environment
has been bootstrapped (see below), and then execute deploying the
`PipelineStack` *once*. Afterwards, the pipeline will keep itself up-to-date.

> **Important**: be sure to `git commit` and `git push` before deploying the
> Pipeline stack using `cdk deploy`!
>
> The reason is that the pipeline will start deploying and self-mutating
> right away based on the sources in the repository, so the sources it finds
> in there should be the ones you want it to find.

Run the following commands to get the pipeline going:

```console
$ git commit -a
$ git push
$ cdk deploy PipelineStack
```

Administrative permissions to the account are only necessary up until
this point. We recommend you remove access to these credentials after doing this.

### Working on the pipeline

The self-mutation feature of the Pipeline might at times get in the way
of the pipeline development workflow. Each change to the pipeline must be pushed
to git, otherwise, after the pipeline was updated using `cdk deploy`, it will
automatically revert to the state found in git.

To make the development more convenient, the self-mutation feature can be turned
off temporarily, by passing `selfMutation: false` property, example:

```go
pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	SelfMutation: jsii.Boolean(false),
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

## Defining the pipeline

This section of the documentation describes the AWS CodePipeline engine,
which comes with this library. If you want to use a different deployment
engine, read the section
[Using a different deployment engine](#using-a-different-deployment-engine) below.

### Synth and sources

To define a pipeline, instantiate a `CodePipeline` construct from the
`aws-cdk-lib/pipelines` module. It takes one argument, a `synth` step, which is
expected to produce the CDK Cloud Assembly as its single output (the contents of
the `cdk.out` directory after running `cdk synth`). "Steps" are arbitrary
actions in the pipeline, typically used to run scripts or commands.

For the synth, use a `ShellStep` and specify the commands necessary to install
dependencies, the CDK CLI, build your project and run `cdk synth`; the specific
commands required will depend on the programming language you are using. For a
typical NPM-based project, the synth will look like this:

```go
var source iFileSetProducer
// the repository source

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

The pipeline assumes that your `ShellStep` will produce a `cdk.out`
directory in the root, containing the CDK cloud assembly. If your
CDK project lives in a subdirectory, be sure to adjust the
`primaryOutputDirectory` to match:

```go
var source iFileSetProducer
// the repository source

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("cd mysubdir"),
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
		PrimaryOutputDirectory: jsii.String("mysubdir/cdk.out"),
	}),
})
```

The underlying `aws-cdk-lib/aws-codepipeline.Pipeline` construct will be produced
when `app.synth()` is called. You can also force it to be produced
earlier by calling `pipeline.buildPipeline()`. After you've called
that method, you can inspect the constructs that were produced by
accessing the properties of the `pipeline` object.

#### Commands for other languages and package managers

The commands you pass to `new ShellStep` will be very similar to the commands
you run on your own workstation to install dependencies and synth your CDK
project. Here are some (non-exhaustive) examples for what those commands might
look like in a number of different situations.

For Yarn, the install commands are different:

```go
var source iFileSetProducer
// the repository source

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("yarn install --frozen-lockfile"),
			jsii.String("yarn build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

For Python projects, remember to install the CDK CLI globally (as
there is no `package.json` to automatically install it for you):

```go
var source iFileSetProducer
// the repository source

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("pip install -r requirements.txt"),
			jsii.String("npm install -g aws-cdk"),
			jsii.String("cdk synth"),
		},
	}),
})
```

For Java projects, remember to install the CDK CLI globally (as
there is no `package.json` to automatically install it for you),
and the Maven compilation step is automatically executed for you
as you run `cdk synth`:

```go
var source iFileSetProducer
// the repository source

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("npm install -g aws-cdk"),
			jsii.String("cdk synth"),
		},
	}),
})
```

You can adapt these examples to your own situation.

#### Migrating from buildspec.yml files

You may currently have the build instructions for your CodeBuild Projects in a
`buildspec.yml` file in your source repository. In addition to your build
commands, the CodeBuild Project's buildspec also controls some information that
CDK Pipelines manages for you, like artifact identifiers, input artifact
locations, Docker authorization, and exported variables.

Since there is no way in general for CDK Pipelines to modify the file in your
resource repository, CDK Pipelines configures the BuildSpec directly on the
CodeBuild Project, instead of loading it from the `buildspec.yml` file.
This requires a pipeline self-mutation to update.

To avoid this, put your build instructions in a separate script, for example
`build.sh`, and call that script from the build `commands` array:

```go
var source iFileSetProducer


pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: []*string{
			jsii.String("./build.sh"),
		},
	}),
})
```

Doing so keeps your exact build instructions in sync with your source code in
the source repository where it belongs, and provides a convenient build script
for developers at the same time.

#### CodePipeline Sources

In CodePipeline, *Sources* define where the source of your application lives.
When a change to the source is detected, the pipeline will start executing.
Source objects can be created by factory methods on the `CodePipelineSource` class:

##### GitHub, GitHub Enterprise, BitBucket using a connection

The recommended way of connecting to GitHub or BitBucket is by using a *connection*.
You will first use the AWS Console to authenticate to the source control
provider, and then use the connection ARN in your pipeline definition:

```go
pipelines.CodePipelineSource_Connection(jsii.String("org/repo"), jsii.String("branch"), &ConnectionSourceOptions{
	ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
})
```

##### GitHub using OAuth

You can also authenticate to GitHub using a personal access token. This expects
that you've created a personal access token and stored it in Secrets Manager.
By default, the source object will look for a secret named **github-token**, but
you can change the name. The token should have the **repo** and **admin:repo_hook**
scopes.

```go
pipelines.CodePipelineSource_GitHub(jsii.String("org/repo"), jsii.String("branch"), &GitHubSourceOptions{
	// This is optional
	Authentication: cdk.SecretValue_SecretsManager(jsii.String("my-token")),
})
```

##### CodeCommit

You can use a CodeCommit repository as the source. Either create or import
that the CodeCommit repository and then use `CodePipelineSource.codeCommit`
to reference it:

```go
repository := codecommit.Repository_FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-repository"))
pipelines.CodePipelineSource_CodeCommit(repository, jsii.String("main"))
```

##### S3

You can use a zip file in S3 as the source of the pipeline. The pipeline will be
triggered every time the file in S3 is changed:

```go
bucket := s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket"))
pipelines.CodePipelineSource_S3(bucket, jsii.String("my/source.zip"))
```

##### ECR

You can use a Docker image in ECR as the source of the pipeline. The pipeline will be
triggered every time an image is pushed to ECR:

```go
repository := ecr.NewRepository(this, jsii.String("Repository"))
pipelines.CodePipelineSource_Ecr(repository)
```

#### Additional inputs

`ShellStep` allows passing in more than one input: additional
inputs will be placed in the directories you specify. Any step that produces an
output file set can be used as an input, such as a `CodePipelineSource`, but
also other `ShellStep`:

```go
prebuild := pipelines.NewShellStep(jsii.String("Prebuild"), &ShellStepProps{
	Input: pipelines.CodePipelineSource_GitHub(jsii.String("myorg/repo1"), jsii.String("main")),
	PrimaryOutputDirectory: jsii.String("./build"),
	Commands: []*string{
		jsii.String("./build.sh"),
	},
})

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_*GitHub(jsii.String("myorg/repo2"), jsii.String("main")),
		AdditionalInputs: map[string]iFileSetProducer{
			"subdir": pipelines.CodePipelineSource_*GitHub(jsii.String("myorg/repo3"), jsii.String("main")),
			"../siblingdir": prebuild,
		},

		Commands: []*string{
			jsii.String("./build.sh"),
		},
	}),
})
```

### CDK application deployments

After you have defined the pipeline and the `synth` step, you can add one or
more CDK `Stages` which will be deployed to their target environments. To do
so, call `pipeline.addStage()` on the Stage object:

```go
var pipeline codePipeline

// Do this as many times as necessary with any account and region
// Account and region may different from the pipeline's.
pipeline.AddStage(
NewMyApplicationStage(this, jsii.String("Prod"), &stageProps{
	Env: &Environment{
		Account: jsii.String("123456789012"),
		Region: jsii.String("eu-west-1"),
	},
}))
```

CDK Pipelines will automatically discover all `Stacks` in the given `Stage`
object, determine their dependency order, and add appropriate actions to the
pipeline to publish the assets referenced in those stacks and deploy the stacks
in the right order.

If the `Stacks` are targeted at an environment in a different AWS account or
Region and that environment has been
[bootstrapped](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)
, CDK Pipelines will transparently make sure the IAM roles are set up
correctly and any requisite replication Buckets are created.

#### Deploying in parallel

By default, all applications added to CDK Pipelines by calling `addStage()` will
be deployed in sequence, one after the other. If you have a lot of stages, you can
speed up the pipeline by choosing to deploy some stages in parallel. You do this
by calling `addWave()` instead of `addStage()`: a *wave* is a set of stages that
are all deployed in parallel instead of sequentially. Waves themselves are still
deployed in sequence. For example, the following will deploy two copies of your
application to `eu-west-1` and `eu-central-1` in parallel:

```go
var pipeline codePipeline

europeWave := pipeline.AddWave(jsii.String("Europe"))
europeWave.AddStage(
NewMyApplicationStage(this, jsii.String("Ireland"), &stageProps{
	Env: &Environment{
		Region: jsii.String("eu-west-1"),
	},
}))
europeWave.AddStage(
NewMyApplicationStage(this, jsii.String("Germany"), &stageProps{
	Env: &Environment{
		Region: jsii.String("eu-central-1"),
	},
}))
```

#### Deploying to other accounts / encrypting the Artifact Bucket

CDK Pipelines can transparently deploy to other Regions and other accounts
(provided those target environments have been
[*bootstrapped*](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)).
However, deploying to another account requires one additional piece of
configuration: you need to enable `crossAccountKeys: true` when creating the
pipeline.

This will encrypt the artifact bucket(s), but incurs a cost for maintaining the
KMS key.

You may also wish to enable automatic key rotation for the created KMS key.

Example:

```go
pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	// Encrypt artifacts, required for cross-account deployments
	CrossAccountKeys: jsii.Boolean(true),
	EnableKeyRotation: jsii.Boolean(true),
	 // optional
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

#### Deploying without change sets

Deployment is done by default with `CodePipeline` engine using change sets,
i.e. to first create a change set and then execute it. This allows you to inject
steps that inspect the change set and approve or reject it, but failed deployments
are not retryable and creation of the change set costs time.

The creation of change sets can be switched off by setting `useChangeSets: false`:

```go
var synth shellStep


type pipelineStack struct {
	stack
}

func newPipelineStack(scope construct, id *string, props stackProps) *pipelineStack {
	this := &pipelineStack{}
	newStack_Override(this, scope, id, props)

	pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
		Synth: Synth,

		// Disable change set creation and make deployments in pipeline as single step
		UseChangeSets: jsii.Boolean(false),
	})
	return this
}
```

### Validation

Every `addStage()` and `addWave()` command takes additional options. As part of these options,
you can specify `pre` and `post` steps, which are arbitrary steps that run before or after
the contents of the stage or wave, respectively. You can use these to add validations like
manual or automated gates to your pipeline. We recommend putting manual approval gates in the set of `pre` steps, and automated approval gates in
the set of `post` steps.

The following example shows both an automated approval in the form of a `ShellStep`, and
a manual approval in the form of a `ManualApprovalStep` added to the pipeline. Both must
pass in order to promote from the `PreProd` to the `Prod` environment:

```go
var pipeline codePipeline

preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
prod := NewMyApplicationStage(this, jsii.String("Prod"))
topic := sns.NewTopic(this, jsii.String("ChangeApprovalTopic"))

pipeline.AddStage(preprod, &AddStageOpts{
	Post: []step{
		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
			Commands: []*string{
				jsii.String("curl -Ssf https://my.webservice.com/"),
			},
		}),
	},
})
pipeline.AddStage(prod, &AddStageOpts{
	Pre: []*step{
		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd"), &ManualApprovalStepProps{
			//All options below are optional
			Comment: jsii.String("Please validate changes"),
			ReviewUrl: jsii.String("https://my.webservice.com/"),
			NotificationTopic: topic,
		}),
	},
})
```

You can also specify steps to be executed at the stack level. To achieve this, you can specify the stack and step via the `stackSteps` property:

```go
var pipeline codePipeline
type myStacksStage struct {
	stage
	stack1 *stack
	stack2 *stack
}

func newMyStacksStage(scope construct, id *string, props stageProps) *myStacksStage {
	this := &myStacksStage{}
	newStage_Override(this, scope, id, props)
	this.stack1 = awscdk.Newstack(this, jsii.String("stack1"))
	this.stack2 = awscdk.Newstack(this, jsii.String("stack2"))
	return this
}
prod := NewMyStacksStage(this, jsii.String("Prod"))

pipeline.AddStage(prod, &AddStageOpts{
	StackSteps: []stackSteps{
		&stackSteps{
			Stack: prod.stack1,
			Pre: []step{
				pipelines.NewManualApprovalStep(jsii.String("Pre-Stack Check")),
			},
			 // Executed before stack is prepared
			ChangeSet: []*step{
				pipelines.NewManualApprovalStep(jsii.String("ChangeSet Approval")),
			},
			 // Executed after stack is prepared but before the stack is deployed
			Post: []*step{
				pipelines.NewManualApprovalStep(jsii.String("Post-Deploy Check")),
			},
		},
		&stackSteps{
			Stack: prod.stack2,
			Post: []*step{
				pipelines.NewManualApprovalStep(jsii.String("Post-Deploy Check")),
			},
		},
	},
})
```

If you specify multiple steps, they will execute in parallel by default. You can add dependencies between them
to if you wish to specify an order. To add a dependency, call `step.addStepDependency()`:

```go
firstStep := pipelines.NewManualApprovalStep(jsii.String("A"))
secondStep := pipelines.NewManualApprovalStep(jsii.String("B"))
secondStep.AddStepDependency(firstStep)
```

For convenience, `Step.sequence()` will take an array of steps and dependencies between adjacent steps,
so that the whole list executes in order:

```go
// Step A will depend on step B and step B will depend on step C
orderedSteps := pipelines.Step_Sequence([]step{
	pipelines.NewManualApprovalStep(jsii.String("A")),
	pipelines.NewManualApprovalStep(jsii.String("B")),
	pipelines.NewManualApprovalStep(jsii.String("C")),
})
```

#### Using CloudFormation Stack Outputs in approvals

Because many CloudFormation deployments result in the generation of resources with unpredictable
names, validations have support for reading back CloudFormation Outputs after a deployment. This
makes it possible to pass (for example) the generated URL of a load balancer to the test set.

To use Stack Outputs, expose the `CfnOutput` object you're interested in, and
pass it to `envFromCfnOutputs` of the `ShellStep`:

```go
var pipeline codePipeline
type myOutputStage struct {
	stage
	loadBalancerAddress cfnOutput
}

func newMyOutputStage(scope construct, id *string, props stageProps) *myOutputStage {
	this := &myOutputStage{}
	newStage_Override(this, scope, id, props)
	this.loadBalancerAddress = awscdk.NewCfnOutput(this, jsii.String("Output"), &CfnOutputProps{
		Value: jsii.String("value"),
	})
	return this
}

lbApp := NewMyOutputStage(this, jsii.String("MyApp"))
pipeline.AddStage(lbApp, &AddStageOpts{
	Post: []step{
		pipelines.NewShellStep(jsii.String("HitEndpoint"), &ShellStepProps{
			EnvFromCfnOutputs: map[string]*cfnOutput{
				// Make the load balancer address available as $URL inside the commands
				"URL": lbApp.loadBalancerAddress,
			},
			Commands: []*string{
				jsii.String("curl -Ssf $URL"),
			},
		}),
	},
})
```

#### Running scripts compiled during the synth step

As part of a validation, you probably want to run a test suite that's more
elaborate than what can be expressed in a couple of lines of shell script.
You can bring additional files into the shell script validation by supplying
the `input` or `additionalInputs` property of `ShellStep`. The input can
be produced by the `Synth` step, or come from a source or any other build
step.

Here's an example that captures an additional output directory in the synth
step and runs tests from there:

```go
var synth shellStep

stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: Synth,
})

pipeline.AddStage(stage, &AddStageOpts{
	Post: []step{
		pipelines.NewShellStep(jsii.String("Approve"), &ShellStepProps{
			// Use the contents of the 'integ' directory from the synth step as the input
			Input: synth.AddOutputDirectory(jsii.String("integ")),
			Commands: []*string{
				jsii.String("cd integ && ./run.sh"),
			},
		}),
	},
})
```

### Customizing CodeBuild Projects

CDK pipelines will generate CodeBuild projects for each `ShellStep` you use, and it
will also generate CodeBuild projects to publish assets and perform the self-mutation
of the pipeline. To control the various aspects of the CodeBuild projects that get
generated, use a `CodeBuildStep` instead of a `ShellStep`. This class has a number
of properties that allow you to customize various aspects of the projects:

```go
var vpc vpc
var mySecurityGroup securityGroup

pipelines.NewCodeBuildStep(jsii.String("Synth"), &CodeBuildStepProps{
	// ...standard ShellStep props...
	Commands: []*string{
	},
	Env: map[string]interface{}{
	},

	// If you are using a CodeBuildStep explicitly, set the 'cdk.out' directory
	// to be the synth step's output.
	PrimaryOutputDirectory: jsii.String("cdk.out"),

	// Control the name of the project
	ProjectName: jsii.String("MyProject"),

	// Control parts of the BuildSpec other than the regular 'build' and 'install' commands
	PartialBuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
	}),

	// Control the build environment
	BuildEnvironment: &BuildEnvironment{
		ComputeType: codebuild.ComputeType_LARGE,
		Privileged: jsii.Boolean(true),
	},
	Timeout: awscdk.Duration_Minutes(jsii.Number(90)),
	FileSystemLocations: []iFileSystemLocation{
		codebuild.FileSystemLocation_Efs(&EfsFileSystemLocationProps{
			Identifier: jsii.String("myidentifier2"),
			Location: jsii.String("myclodation.mydnsroot.com:/loc"),
			MountPoint: jsii.String("/media"),
			MountOptions: jsii.String("opts"),
		}),
	},

	// Control Elastic Network Interface creation
	Vpc: vpc,
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	SecurityGroups: []iSecurityGroup{
		mySecurityGroup,
	},

	// Control caching
	Cache: codebuild.Cache_Bucket(s3.NewBucket(this, jsii.String("Cache"))),

	// Additional policy statements for the execution role
	RolePolicyStatements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
		}),
	},
})
```

You can also configure defaults for *all* CodeBuild projects by passing `codeBuildDefaults`,
or just for the synth, asset publishing, and self-mutation projects by passing `synthCodeBuildDefaults`,
`assetPublishingCodeBuildDefaults`, or `selfMutationCodeBuildDefaults`:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc
var mySecurityGroup securityGroup


pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	// Standard CodePipeline properties
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),

	// Defaults for all CodeBuild projects
	CodeBuildDefaults: &CodeBuildOptions{
		// Prepend commands and configuration to all projects
		PartialBuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
			"version": jsii.String("0.2"),
		}),

		// Control the build environment
		BuildEnvironment: &BuildEnvironment{
			ComputeType: codebuild.ComputeType_LARGE,
		},

		// Control Elastic Network Interface creation
		Vpc: vpc,
		SubnetSelection: &SubnetSelection{
			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
		SecurityGroups: []iSecurityGroup{
			mySecurityGroup,
		},

		// Additional policy statements for the execution role
		RolePolicy: []policyStatement{
			iam.NewPolicyStatement(&PolicyStatementProps{
			}),
		},

		// Information about logs
		Logging: &LoggingOptions{
			CloudWatch: &CloudWatchLoggingOptions{
				LogGroup: awscdk.Aws_logs.NewLogGroup(this, jsii.String("MyLogGroup")),
			},
			S3: &S3LoggingOptions{
				Bucket: s3.NewBucket(this, jsii.String("LogBucket")),
			},
		},
	},

	SynthCodeBuildDefaults: &CodeBuildOptions{
	},
	AssetPublishingCodeBuildDefaults: &CodeBuildOptions{
	},
	SelfMutationCodeBuildDefaults: &CodeBuildOptions{
	},
})
```

### Arbitrary CodePipeline actions

If you want to add a type of CodePipeline action to the CDK Pipeline that
doesn't have a matching class yet, you can define your own step class that extends
`Step` and implements `ICodePipelineActionFactory`.

Here's an example that adds a Jenkins step:

```go
type myJenkinsStep struct {
	step
}

func newMyJenkinsStep(provider jenkinsProvider, input fileSet) *myJenkinsStep {
	this := &myJenkinsStep{}
	pipelines.NewStep_Override(this, jsii.String("MyJenkinsStep"))

	// This is necessary if your step accepts parameters, like environment variables,
	// that may contain outputs from other steps. It doesn't matter what the
	// structure is, as long as it contains the values that may contain outputs.
	this.DiscoverReferencedOutputs(map[string]map[string]interface{}{
		"env": map[string]interface{}{
		},
	})
	return this
}

func (this *myJenkinsStep) produceAction(stage iStage, options produceActionOptions) codePipelineActionFactoryResult {
	// This is where you control what type of Action gets added to the
	// CodePipeline
	*stage.AddAction(
	cpactions.NewJenkinsAction(&JenkinsActionProps{
		// Copy 'actionName' and 'runOrder' from the options
		ActionName: options.ActionName,
		RunOrder: options.RunOrder,

		// Jenkins-specific configuration
		Type: cpactions.JenkinsActionType_TEST,
		JenkinsProvider: this.provider,
		ProjectName: jsii.String("MyJenkinsProject"),

		// Translate the FileSet into a codepipeline.Artifact
		Inputs: []artifact{
			options.Artifacts.ToCodePipeline(this.input),
		},
	}))

	return &codePipelineActionFactoryResult{
		RunOrdersConsumed: jsii.Number(1),
	}
}
```

Another example, adding a lambda step referencing outputs from a stack:

```go
type myLambdaStep struct {
	step
	stackOutputReference stackOutputReference
}

func newMyLambdaStep(fn function, stackOutput cfnOutput) *myLambdaStep {
	this := &myLambdaStep{}
	pipelines.NewStep_Override(this, jsii.String("MyLambdaStep"))
	this.stackOutputReference = pipelines.stackOutputReference_FromCfnOutput(stackOutput)
	return this
}

func (this *myLambdaStep) produceAction(stage iStage, options produceActionOptions) codePipelineActionFactoryResult {
	*stage.AddAction(
	cpactions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
		ActionName: options.ActionName,
		RunOrder: options.RunOrder,
		// Map the reference to the variable name the CDK has generated for you.
		UserParameters: map[string]interface{}{
			"stackOutput": options.stackOutputsMap.toCodePipeline(this.stackOutputReference),
		},
		Lambda: this.fn,
	}))

	return &codePipelineActionFactoryResult{
		RunOrdersConsumed: jsii.Number(1),
	}
}public get consumedStackOutputs(): pipelines.StackOutputReference[] {
    return [this.stackOutputReference];
  }
```

### Using an existing AWS Codepipeline

If you wish to use an existing `CodePipeline.Pipeline` while using the modern API's
methods and classes, you can pass in the existing `CodePipeline.Pipeline` to be built upon
instead of having the `pipelines.CodePipeline` construct create a new `CodePipeline.Pipeline`.
This also gives you more direct control over the underlying `CodePipeline.Pipeline` construct
if the way the modern API creates it doesn't allow for desired configurations. Use `CodePipelineFileset` to convert CodePipeline **artifacts** into CDK Pipelines **file sets**,
that can be used everywhere a file set or file set producer is expected.

Here's an example of passing in an existing pipeline and using a *source* that's already
in the pipeline:

```go
var codePipeline pipeline


sourceArtifact := codepipeline.NewArtifact(jsii.String("MySourceArtifact"))

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	CodePipeline: codePipeline,
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineFileSet_FromArtifact(sourceArtifact),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

If your existing pipeline already provides a synth step, pass the existing
artifact in place of the `synth` step:

```go
var codePipeline pipeline


buildArtifact := codepipeline.NewArtifact(jsii.String("MyBuildArtifact"))

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	CodePipeline: codePipeline,
	Synth: pipelines.CodePipelineFileSet_FromArtifact(buildArtifact),
})
```

Note that if you provide an existing pipeline, you cannot provide values for
`pipelineName`, `crossAccountKeys`, `reuseCrossRegionSupportStacks`, or `role`
because those values are passed in directly to the underlying `codepipeline.Pipeline`.

### Use pipeline service role as default action role in pipeline

By default CDK automatically creates roles for each action (`CodeBuildStep`, etc).
If you prefer to use the pipeline service role as default instead, set the `usePipelineRoleForActions` property.
This will tell CDK to default to the pipeline service role in AWS CodePipeline if no action role is provided.

## Using Docker in the pipeline

Docker can be used in 3 different places in the pipeline:

* If you are using Docker image assets in your application stages: Docker will
  run in the asset publishing projects.
* If you are using Docker image assets in your stack (for example as
  images for your CodeBuild projects): Docker will run in the self-mutate project.
* If you are using Docker to bundle file assets anywhere in your project (for
  example, if you are using such construct libraries as
  `aws-cdk-lib/aws-lambda-nodejs`): Docker will run in the
  *synth* project.

For the first case, you don't need to do anything special. For the other two cases,
you need to make sure that **privileged mode** is enabled on the correct CodeBuild
projects, so that Docker can run correctly. The follow sections describe how to do
that.

You may also need to authenticate to Docker registries to avoid being throttled.
See the section **Authenticating to Docker registries** below for information on how to do
that.

### Using Docker image assets in the pipeline

If your `PipelineStack` is using Docker image assets (as opposed to the application
stacks the pipeline is deploying), for example by the use of `LinuxBuildImage.fromAsset()`,
you need to pass `dockerEnabledForSelfMutation: true` to the pipeline. For example:

```go
pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),

	// Turn this on because the pipeline uses Docker image assets
	DockerEnabledForSelfMutation: jsii.Boolean(true),
})

pipeline.AddWave(jsii.String("MyWave"), &WaveOptions{
	Post: []step{
		pipelines.NewCodeBuildStep(jsii.String("RunApproval"), &CodeBuildStepProps{
			Commands: []*string{
				jsii.String("command-from-image"),
			},
			BuildEnvironment: &BuildEnvironment{
				// The user of a Docker image asset in the pipeline requires turning on
				// 'dockerEnabledForSelfMutation'.
				BuildImage: codebuild.LinuxBuildImage_FromAsset(this, jsii.String("Image"), &DockerImageAssetProps{
					Directory: jsii.String("./docker-image"),
				}),
			},
		}),
	},
})
```

> **Important**: You must turn on the `dockerEnabledForSelfMutation` flag,
> commit and allow the pipeline to self-update *before* adding the actual
> Docker asset.

### Using bundled file assets

If you are using asset bundling anywhere (such as automatically done for you
if you add a construct like `aws-cdk-lib/aws-lambda-nodejs`), you need to pass
`dockerEnabledForSynth: true` to the pipeline. For example:

```go
pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),

	// Turn this on because the application uses bundled file assets
	DockerEnabledForSynth: jsii.Boolean(true),
})
```

> **Important**: You must turn on the `dockerEnabledForSynth` flag,
> commit and allow the pipeline to self-update *before* adding the actual
> Docker asset.

### Authenticating to Docker registries

You can specify credentials to use for authenticating to Docker registries as part of the
pipeline definition. This can be useful if any Docker image assets — in the pipeline or
any of the application stages — require authentication, either due to being in a
different environment (e.g., ECR repo) or to avoid throttling (e.g., DockerHub).

```go
dockerHubSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
customRegSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("CRSecret"), jsii.String("arn:aws:..."))
repo1 := ecr.Repository_FromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo1"))
repo2 := ecr.Repository_FromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo2"))

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	DockerCredentials: []dockerCredential{
		pipelines.*dockerCredential_DockerHub(dockerHubSecret),
		pipelines.*dockerCredential_CustomRegistry(jsii.String("dockerregistry.example.com"), customRegSecret),
		pipelines.*dockerCredential_Ecr([]iRepository{
			repo1,
			repo2,
		}),
	},
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
})
```

For authenticating to Docker registries that require a username and password combination
(like DockerHub), create a Secrets Manager Secret with fields named `username`
and `secret`, and import it (the field names change be customized).

Authentication to ECR repositories is done using the execution role of the
relevant CodeBuild job. Both types of credentials can be provided with an
optional role to assume before requesting the credentials.

By default, the Docker credentials provided to the pipeline will be available to
the **Synth**, **Self-Update**, and **Asset Publishing** actions within the
*pipeline. The scope of the credentials can be limited via the `DockerCredentialUsage` option.

```go
dockerHubSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
// Only the image asset publishing actions will be granted read access to the secret.
creds := pipelines.DockerCredential_DockerHub(dockerHubSecret, &ExternalDockerCredentialOptions{
	Usages: []dockerCredentialUsage{
		pipelines.*dockerCredentialUsage_ASSET_PUBLISHING,
	},
})
```

## CDK Environment Bootstrapping

An *environment* is an *(account, region)* pair where you want to deploy a
CDK stack (see
[Environments](https://docs.aws.amazon.com/cdk/latest/guide/environments.html)
in the CDK Developer Guide). In a Continuous Deployment pipeline, there are
at least two environments involved: the environment where the pipeline is
provisioned, and the environment where you want to deploy the application (or
different stages of the application). These can be the same, though best
practices recommend you isolate your different application stages from each
other in different AWS accounts or regions.

Before you can provision the pipeline, you have to *bootstrap* the environment you want
to create it in. If you are deploying your application to different environments, you
also have to bootstrap those and be sure to add a *trust* relationship.

After you have bootstrapped an environment and created a pipeline that deploys
to it, it's important that you don't delete the stack or change its *Qualifier*,
or future deployments to this environment will fail. If you want to upgrade
the bootstrap stack to a newer version, do that by updating it in-place.

> This library requires the *modern* bootstrapping stack which has
> been updated specifically to support cross-account continuous delivery.
>
> If you are using CDKv2, you do not need to do anything else. Modern
> bootstrapping and modern stack synthesis (also known as "default stack
> synthesis") is the default.
>
> If you are using CDKv1, you need to opt in to modern bootstrapping and
> modern stack synthesis using a feature flag. Make sure `cdk.json` includes:
>
> ```json
> {
>   "context": {
>     "@aws-cdk/core:newStyleStackSynthesis": true
>   }
> }
> ```
>
> And be sure to run `cdk bootstrap` in the same directory as the `cdk.json`
> file.

To bootstrap an environment for provisioning the pipeline:

```console
$ npx cdk bootstrap \
    [--profile admin-profile-1] \
    --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess \
    aws://111111111111/us-east-1
```

To bootstrap a different environment for deploying CDK applications into using
a pipeline in account `111111111111`:

```console
$ npx cdk bootstrap \
    [--profile admin-profile-2] \
    --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess \
    --trust 11111111111 \
    aws://222222222222/us-east-2
```

If you only want to trust an account to do lookups (e.g, when your CDK application has a
`Vpc.fromLookup()` call), use the option `--trust-for-lookup`:

```console
$ npx cdk bootstrap \
    [--profile admin-profile-2] \
    --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess \
    --trust-for-lookup 11111111111 \
    aws://222222222222/us-east-2
```

These command lines explained:

* `npx`: means to use the CDK CLI from the current NPM install. If you are using
  a global install of the CDK CLI, leave this out.
* `--profile`: should indicate a profile with administrator privileges that has
  permissions to provision a pipeline in the indicated account. You can leave this
  flag out if either the AWS default credentials or the `AWS_*` environment
  variables confer these permissions.
* `--cloudformation-execution-policies`: ARN of the managed policy that future CDK
  deployments should execute with. By default this is `AdministratorAccess`, but
  if you also specify the `--trust` flag to give another Account permissions to
  deploy into the current account, you must specify a value here.
* `--trust`: indicates which other account(s) should have permissions to deploy
  CDK applications into this account. In this case we indicate the Pipeline's account,
  but you could also use this for developer accounts (don't do that for production
  application accounts though!).
* `--trust-for-lookup`: gives a more limited set of permissions to the
  trusted account, only allowing it to look up values such as availability zones, EC2 images and
  VPCs. `--trust-for-lookup` does not give permissions to modify anything in the account.
  Note that `--trust` implies `--trust-for-lookup`, so you don't need to specify
  the same account twice.
* `aws://222222222222/us-east-2`: the account and region we're bootstrapping.

> Be aware that anyone who has access to the trusted Accounts **effectively has all
> permissions conferred by the configured CloudFormation execution policies**,
> allowing them to do things like read arbitrary S3 buckets and create arbitrary
> infrastructure in the bootstrapped account. Restrict the list of `--trust`ed Accounts,
> or restrict the policies configured by `--cloudformation-execution-policies`.

<br>

> **Security tip**: we recommend that you use administrative credentials to an
> account only to bootstrap it and provision the initial pipeline. Otherwise,
> access to administrative credentials should be dropped as soon as possible.

<br>

> **On the use of AdministratorAccess**: The use of the `AdministratorAccess` policy
> ensures that your pipeline can deploy every type of AWS resource to your account.
> Make sure you trust all the code and dependencies that make up your CDK app.
> Check with the appropriate department within your organization to decide on the
> proper policy to use.
>
> If your policy includes permissions to create on attach permission to a role,
> developers can escalate their privilege with more permissive permission.
> Thus, we recommend implementing [permissions boundary](https://aws.amazon.com/premiumsupport/knowledge-center/iam-permission-boundaries/)
> in the CDK Execution role. To do this, you can bootstrap with the `--template` option with
> [a customized template](https://github.com/aws-samples/aws-bootstrap-kit-examples/blob/ba28a97d289128281bc9483bcba12c1793f2c27a/source/1-SDLC-organization/lib/cdk-bootstrap-template.yml#L395) that contains a permission boundary.

### Migrating from old bootstrap stack

The bootstrap stack is a CloudFormation stack in your account named
**CDKToolkit** that provisions a set of resources required for the CDK
to deploy into that environment.

The "new" bootstrap stack (obtained by running `cdk bootstrap` with
`CDK_NEW_BOOTSTRAP=1`) is slightly more elaborate than the "old" stack. It
contains:

* An S3 bucket and ECR repository with predictable names, so that we can reference
  assets in these storage locations *without* the use of CloudFormation template
  parameters.
* A set of roles with permissions to access these asset locations and to execute
  CloudFormation, assumable from whatever accounts you specify under `--trust`.

It is possible and safe to migrate from the old bootstrap stack to the new
bootstrap stack. This will create a new S3 file asset bucket in your account
and orphan the old bucket. You should manually delete the orphaned bucket
after you are sure you have redeployed all CDK applications and there are no
more references to the old asset bucket.

## Considerations around Running at Scale

If you are planning to run pipelines for more than a hundred repos
deploying across multiple regions, then you will want to consider reusing
both artifacts buckets and cross-region replication buckets.

In a situation like this, you will want to have a separate CDK app / dedicated repo which creates
and managed the buckets which will be shared by the pipelines of all your other apps.
Note that this app must NOT be using the shared buckets because of chicken & egg issues.

The following code assumes you have created and are managing your buckets in the aforementioned
separate cdk repo and are just importing them for use in one of your (many) pipelines.

```go
var sharedXRegionUsWest1BucketArn string
var sharedXRegionUsWest1KeyArn string

var sharedXRegionUsWest2BucketArn string
var sharedXRegionUsWest2KeyArn string


usWest1Bucket := s3.Bucket_FromBucketAttributes(*scope, jsii.String("UsWest1Bucket"), &BucketAttributes{
	BucketArn: sharedXRegionUsWest1BucketArn,
	EncryptionKey: kms.Key_FromKeyArn(scope, jsii.String("UsWest1BucketKeyArn"), sharedXRegionUsWest1BucketArn),
})

usWest2Bucket := s3.Bucket_FromBucketAttributes(*scope, jsii.String("UsWest2Bucket"), &BucketAttributes{
	BucketArn: sharedXRegionUsWest2BucketArn,
	EncryptionKey: kms.Key_*FromKeyArn(scope, jsii.String("UsWest2BucketKeyArn"), sharedXRegionUsWest2KeyArn),
})

crossRegionReplicationBuckets := map[string]iBucket{
	"us-west-1": usWest1Bucket,
	"us-west-2": usWest2Bucket,
}

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
		},
	}),
	 // Use shared buckets.
	CrossRegionReplicationBuckets: CrossRegionReplicationBuckets,
})
```

## Context Lookups

You might be using CDK constructs that need to look up [runtime
context](https://docs.aws.amazon.com/cdk/latest/guide/context.html#context_methods),
which is information from the target AWS Account and Region the CDK needs to
synthesize CloudFormation templates appropriate for that environment. Examples
of this kind of context lookups are the number of Availability Zones available
to you, a Route53 Hosted Zone ID, or the ID of an AMI in a given region. This
information is automatically looked up when you run `cdk synth`.

By default, a `cdk synth` performed in a pipeline will not have permissions
to perform these lookups, and the lookups will fail. This is by design.

**Our recommended way of using lookups** is by running `cdk synth` on the
developer workstation and checking in the `cdk.context.json` file, which
contains the results of the context lookups. This will make sure your
synthesized infrastructure is consistent and repeatable. If you do not commit
`cdk.context.json`, the results of the lookups may suddenly be different in
unexpected ways, and even produce results that cannot be deployed or will cause
data loss. To give an account permissions to perform lookups against an
environment, without being able to deploy to it and make changes, run
`cdk bootstrap --trust-for-lookup=<account>`.

If you want to use lookups directly from the pipeline, you either need to accept
the risk of nondeterminism, or make sure you save and load the
`cdk.context.json` file somewhere between synth runs. Finally, you should
give the synth CodeBuild execution role permissions to assume the bootstrapped
lookup roles. As an example, doing so would look like this:

```go
pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	Synth: pipelines.NewCodeBuildStep(jsii.String("Synth"), &CodeBuildStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
		}),
		Commands: []*string{
			jsii.String("..."),
			jsii.String("npm ci"),
			jsii.String("npm run build"),
			jsii.String("npx cdk synth"),
			jsii.String("..."),
		},
		RolePolicyStatements: []policyStatement{
			iam.NewPolicyStatement(&PolicyStatementProps{
				Actions: []*string{
					jsii.String("sts:AssumeRole"),
				},
				Resources: []*string{
					jsii.String("*"),
				},
				Conditions: map[string]interface{}{
					"StringEquals": map[string]*string{
						"iam:ResourceTag/aws-cdk:bootstrap-role": jsii.String("lookup"),
					},
				},
			}),
		},
	}),
})
```

The above example requires that the target environments have all
been bootstrapped with bootstrap stack version `8`, released with
CDK CLI `1.114.0`.

## Security Considerations

It's important to stay safe while employing Continuous Delivery. The CDK Pipelines
library comes with secure defaults to the best of our ability, but by its
very nature the library cannot take care of everything.

We therefore expect you to mind the following:

* Maintain dependency hygiene and vet 3rd-party software you use. Any software you
  run on your build machine has the ability to change the infrastructure that gets
  deployed. Be careful with the software you depend on.
* Use dependency locking to prevent accidental upgrades! The default `CdkSynths` that
  come with CDK Pipelines will expect `package-lock.json` and `yarn.lock` to
  ensure your dependencies are the ones you expect.
* CDK Pipelines runs on resources created in your own account, and the configuration
  of those resources is controlled by developers submitting code through the pipeline.
  Therefore, CDK Pipelines by itself cannot protect against malicious
  developers trying to bypass compliance checks. If your threat model includes
  developers writing CDK code, you should have external compliance mechanisms in place like
  [AWS CloudFormation Hooks](https://aws.amazon.com/blogs/mt/proactively-keep-resources-secure-and-compliant-with-aws-cloudformation-hooks/)
  (preventive) or [AWS Config](https://aws.amazon.com/config/) (reactive) that
  the CloudFormation Execution Role does not have permissions to disable.
* Credentials to production environments should be short-lived. After
  bootstrapping and the initial pipeline provisioning, there is no more need for
  developers to have access to any of the account credentials; all further
  changes can be deployed through git. Avoid the chances of credentials leaking
  by not having them in the first place!

### Confirm permissions broadening

To keep tabs on the security impact of changes going out through your pipeline,
you can insert a security check before any stage deployment. This security check
will check if the upcoming deployment would add any new IAM permissions or
security group rules, and if so pause the pipeline and require you to confirm
the changes.

The security check will appear as two distinct actions in your pipeline: first
a CodeBuild project that runs `cdk diff` on the stage that's about to be deployed,
followed by a Manual Approval action that pauses the pipeline. If it so happens
that there no new IAM permissions or security group rules will be added by the deployment,
the manual approval step is automatically satisfied. The pipeline will look like this:

```txt
Pipeline
├── ...
├── MyApplicationStage
│    ├── MyApplicationSecurityCheck       // Security Diff Action
│    ├── MyApplicationManualApproval      // Manual Approval Action
│    ├── Stack.Prepare
│    └── Stack.Deploy
└── ...
```

You can insert the security check by using a `ConfirmPermissionsBroadening` step:

```go
var pipeline codePipeline

stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
pipeline.AddStage(stage, &AddStageOpts{
	Pre: []step{
		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &PermissionsBroadeningCheckProps{
			Stage: *Stage,
		}),
	},
})
```

To get notified when there is a change that needs your manual approval,
create an SNS Topic, subscribe your own email address, and pass it in as
as the `notificationTopic` property:

```go
var pipeline codePipeline

topic := sns.NewTopic(this, jsii.String("SecurityChangesTopic"))
topic.AddSubscription(subscriptions.NewEmailSubscription(jsii.String("test@email.com")))

stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
pipeline.AddStage(stage, &AddStageOpts{
	Pre: []step{
		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &PermissionsBroadeningCheckProps{
			Stage: *Stage,
			NotificationTopic: topic,
		}),
	},
})
```

**Note**: Manual Approvals notifications only apply when an application has security
check enabled.

## Using a different deployment engine

CDK Pipelines supports multiple *deployment engines*, but this module vends a
construct for only one such engine: AWS CodePipeline. It is also possible to
use CDK Pipelines to build pipelines backed by other deployment engines.

Here is a list of CDK Libraries that integrate CDK Pipelines with
alternative deployment engines:

* GitHub Workflows: [`cdk-pipelines-github`](https://github.com/cdklabs/cdk-pipelines-github)

## Troubleshooting

Here are some common errors you may encounter while using this library.

### Pipeline: Internal Failure

If you see the following error during deployment of your pipeline:

```plaintext
CREATE_FAILED  | AWS::CodePipeline::Pipeline | Pipeline/Pipeline
Internal Failure
```

There's something wrong with your GitHub access token. It might be missing, or not have the
right permissions to access the repository you're trying to access.

### Key: Policy contains a statement with one or more invalid principals

If you see the following error during deployment of your pipeline:

```plaintext
CREATE_FAILED | AWS::KMS::Key | Pipeline/Pipeline/ArtifactsBucketEncryptionKey
Policy contains a statement with one or more invalid principals.
```

One of the target (account, region) environments has not been bootstrapped
with the new bootstrap stack. Check your target environments and make sure
they are all bootstrapped.

### Message: no matching base directory path found for cdk.out

If you see this error during the **Synth** step, it means that CodeBuild
is expecting to find a `cdk.out` directory in the root of your CodeBuild project,
but the directory wasn't there. There are two common causes for this:

* `cdk synth` is not being executed: `cdk synth` used to be run
  implicitly for you, but you now have to explicitly include the command.
  For NPM-based projects, add `npx cdk synth` to the end of the `commands`
  property, for other languages add `npm install -g aws-cdk` and `cdk synth`.
* Your CDK project lives in a subdirectory: you added a `cd <somedirectory>` command
  to the list of commands; don't forget to tell the `ScriptStep` about the
  different location of `cdk.out`, by passing `primaryOutputDirectory: '<somedirectory>/cdk.out'`.

### <Stack> is in ROLLBACK_COMPLETE state and can not be updated

If you see the following error during execution of your pipeline:

```plaintext
Stack ... is in ROLLBACK_COMPLETE state and can not be updated. (Service:
AmazonCloudFormation; Status Code: 400; Error Code: ValidationError; Request
ID: ...)
```

The stack failed its previous deployment, and is in a non-retryable state.
Go into the CloudFormation console, delete the stack, and retry the deployment.

### Cannot find module 'xxxx' or its corresponding type declarations

You may see this if you are using TypeScript or other NPM-based languages,
when using NPM 7 on your workstation (where you generate `package-lock.json`)
and NPM 6 on the CodeBuild image used for synthesizing.

It looks like NPM 7 has started writing less information to `package-lock.json`,
leading NPM 6 reading that same file to not install all required packages anymore.

Make sure you are using the same NPM version everywhere, either downgrade your
workstation's version or upgrade the CodeBuild version.

### Cannot find module '.../check-node-version.js' (MODULE_NOT_FOUND)

The above error may be produced by `npx` when executing the CDK CLI, or any
project that uses the AWS SDK for JavaScript, without the target application
having been installed yet. For example, it can be triggered by `npx cdk synth`
if `aws-cdk` is not in your `package.json`.

Work around this by either installing the target application using NPM *before*
running `npx`, or set the environment variable `NPM_CONFIG_UNSAFE_PERM=true`.

### Cannot connect to the Docker daemon at unix:///var/run/docker.sock

If, in the 'Synth' action (inside the 'Build' stage) of your pipeline, you get an error like this:

```console
stderr: docker: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?.
See 'docker run --help'.
```

It means that the AWS CodeBuild project for 'Synth' is not configured to run in privileged mode,
which prevents Docker builds from happening. This typically happens if you use a CDK construct
that bundles asset using tools run via Docker, like `aws-lambda-nodejs`, `aws-lambda-python`,
`aws-lambda-go` and others.

Make sure you set the `privileged` environment variable to `true` in the synth definition:

After turning on `privilegedMode: true`, you will need to do a one-time manual cdk deploy of your
pipeline to get it going again (as with a broken 'synth' the pipeline will not be able to self
update to the right state).

### Not authorized to perform sts:AssumeRole on arn:aws:iam::*:role/*-lookup-role-*

You may get an error like the following in the **Synth** step:

```text
Could not assume role in target account using current credentials (which are for account 111111111111). User:
arn:aws:sts::111111111111:assumed-role/PipelineStack-PipelineBuildSynthCdkBuildProje-..../AWSCodeBuild-....
is not authorized to perform: sts:AssumeRole on resource:
arn:aws:iam::222222222222:role/cdk-hnb659fds-lookup-role-222222222222-us-east-1.
Please make sure that this role exists in the account. If it doesn't exist, (re)-bootstrap the environment with
the right '--trust', using the latest version of the CDK CLI.
```

This is a sign that the CLI is trying to do Context Lookups during the **Synth** step, which are failing
because it cannot assume the right role. We recommend you don't rely on Context Lookups in the pipeline at
all, and commit a file called `cdk.context.json` with the right lookup values in it to source control.

If you do want to do lookups in the pipeline, the cause is one of the following:

* The target environment has not been bootstrapped; OR
* The target environment has been bootstrapped without the right `--trust` relationship; OR
* The CodeBuild execution role does not have permissions to call `sts:AssumeRole`.

See the section called **Context Lookups** for more information on using this feature.

### IAM policies: Cannot exceed quota for PoliciesPerRole / Maximum policy size exceeded

This happens as a result of having a lot of targets in the Pipeline: the IAM policies that
get generated enumerate all required roles and grow too large.

Make sure you are on version `2.26.0` or higher, and that your `cdk.json` contains the
following:

```json
{
  "context": {
    "aws-cdk-lib/aws-iam:minimizePolicies": true
  }
}
```

### S3 error: Access Denied

An "S3 Access Denied" error can have two causes:

* Asset hashes have changed, but self-mutation has been disabled in the pipeline.
* You have deleted and recreated the bootstrap stack, or changed its qualifier.

#### Bootstrap roles have been renamed or recreated

While attempting to deploy an application stage, the "Prepare" or "Deploy" stage may fail with a cryptic error like:

`Action execution failed Access Denied (Service: Amazon S3; Status Code: 403; Error Code: AccessDenied; Request ID: 0123456ABCDEFGH; S3 Extended Request ID: 3hWcrVkhFGxfiMb/rTJO0Bk7Qn95x5ll4gyHiFsX6Pmk/NT+uX9+Z1moEcfkL7H3cjH7sWZfeD0=; Proxy: null)`

This generally indicates that the roles necessary to deploy have been deleted (or deleted and re-created);
for example, if the bootstrap stack has been deleted and re-created, this scenario will happen. Under the hood,
the resources that rely on these roles (e.g., `cdk-$qualifier-deploy-role-$account-$region`) point to different
canonical IDs than the recreated versions of these roles, which causes the errors. There are no simple solutions
to this issue, and for that reason we **strongly recommend** that bootstrap stacks not be deleted and re-created
once created.

The most automated way to solve the issue is to introduce a secondary bootstrap stack. By changing the qualifier
that the pipeline stack looks for, a change will be detected and the impacted policies and resources will be updated.
A hypothetical recovery workflow would look something like this:

* First, for all impacted environments, create a secondary bootstrap stack:

```sh
$ env CDK_NEW_BOOTSTRAP=1 npx cdk bootstrap \
    --qualifier random1234 \
    --toolkit-stack-name CDKToolkitTemp \
    aws://111111111111/us-east-1
```

* Update all impacted stacks in the pipeline to use this new qualifier.
  See https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html for more info.

```go
awscdk.Newstack(this, jsii.String("MyStack"), &StackProps{
	// Update this qualifier to match the one used above.
	Synthesizer: cdk.NewDefaultStackSynthesizer(&DefaultStackSynthesizerProps{
		Qualifier: jsii.String("randchars1234"),
	}),
})
```

* Deploy the updated stacks. This will update the stacks to use the roles created in the new bootstrap stack.
* (Optional) Restore back to the original state:

  * Revert the change made in step #2 above
  * Re-deploy the pipeline to use the original qualifier.
  * Delete the temporary bootstrap stack(s)

##### Manual Alternative

Alternatively, the errors can be resolved by finding each impacted resource and policy, and correcting the policies
by replacing the canonical IDs (e.g., `AROAYBRETNYCYV6ZF2R93`) with the appropriate ARNs. As an example, the KMS
encryption key policy for the artifacts bucket may have a statement that looks like the following:

```json
{
  "Effect": "Allow",
  "Principal": {
    // "AWS" : "AROAYBRETNYCYV6ZF2R93"  // Indicates this issue; replace this value
    "AWS": "arn:aws:iam::0123456789012:role/cdk-hnb659fds-deploy-role-0123456789012-eu-west-1" // Correct value
  },
  "Action": ["kms:Decrypt", "kms:DescribeKey"],
  "Resource": "*"
}
```

Any resource or policy that references the qualifier (`hnb659fds` by default) will need to be updated.

### This CDK CLI is not compatible with the CDK library used by your application

The CDK CLI version used in your pipeline is too old to read the Cloud Assembly
produced by your CDK app.

Most likely this happens in the `SelfMutate` action, you are passing the `cliVersion`
parameter to control the version of the CDK CLI, and you just updated the CDK
framework version that your application uses. You either forgot to change the
`cliVersion` parameter, or changed the `cliVersion` in the same commit in which
you changed the framework version. Because a change to the pipeline settings needs
a successful run of the `SelfMutate` step to be applied, the next iteration of the
`SelfMutate` step still executes with the *old* CLI version, and that old CLI version
is not able to read the cloud assembly produced by the new framework version.

Solution: change the `cliVersion` first, commit, push and deploy, and only then
change the framework version.

We recommend you avoid specifying the `cliVersion` parameter at all. By default
the pipeline will use the latest CLI version, which will support all cloud assembly
versions.

## Using Drop-in Docker Replacements

By default, the AWS CDK will build and publish Docker image assets using the
`docker` command. However, by specifying the `CDK_DOCKER` environment variable,
you can override the command that will be used to build and publish your
assets. To learn more, see [How to replace Docker with another container management tool](https://docs.aws.amazon.com/cdk/v2/guide/build-containers.html#build-container-replace)
in the *AWS CDK Developer Guide*.

In CDK Pipelines, the drop-in replacement for the `docker` command must be
included in the CodeBuild environment and configured for your pipeline.

### Adding to the default CodeBuild image

You can add a drop-in Docker replacement command to the default CodeBuild
environment by adding install-phase commands that encode how to install
your tooling and by adding the `CDK_DOCKER` environment variable to your
build environment.

```go
var source iFileSetProducer // the repository source
var synthCommands []*string // Commands to synthesize your app
var installCommands []*string
// Commands to install your toolchain

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	// Standard CodePipeline properties...
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: synthCommands,
	}),

	// Configure CodeBuild to use a drop-in Docker replacement.
	CodeBuildDefaults: &CodeBuildOptions{
		PartialBuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
			"phases": map[string]map[string][]*string{
				"install": map[string][]*string{
					// Add the shell commands to install your drop-in Docker
					// replacement to the CodeBuild enviromment.
					"commands": installCommands,
				},
			},
		}),
		BuildEnvironment: &BuildEnvironment{
			EnvironmentVariables: map[string]buildEnvironmentVariable{
				// Instruct the AWS CDK to use `drop-in-replacement` instead of
				// `docker` when building / publishing docker images.
				// e.g., `drop-in-replacement build . -f path/to/Dockerfile`
				"CDK_DOCKER": &buildEnvironmentVariable{
					"value": jsii.String("drop-in-replacement"),
				},
			},
		},
	},
})
```

### Using a custom build image

If you're using a custom build image in CodeBuild, you can override the
command the AWS CDK uses to build Docker images by providing `CDK_DOCKER` as
an `ENV` in your `Dockerfile` or by providing the environment variable in the
pipeline as shown below.

```go
var source iFileSetProducer // the repository source
var synthCommands []*string
// Commands to synthesize your app

pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
	// Standard CodePipeline properties...
	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
		Input: source,
		Commands: synthCommands,
	}),

	// Configure CodeBuild to use a drop-in Docker replacement.
	CodeBuildDefaults: &CodeBuildOptions{
		BuildEnvironment: &BuildEnvironment{
			// Provide a custom build image containing your toolchain and the
			// pre-installed replacement for the `docker` command.
			BuildImage: codebuild.LinuxBuildImage_FromDockerRegistry(jsii.String("your-docker-registry")),
			EnvironmentVariables: map[string]buildEnvironmentVariable{
				// If you haven't provided an `ENV` in your Dockerfile that overrides
				// `CDK_DOCKER`, then you must provide the name of the command that
				// the AWS CDK should run instead of `docker` here.
				"CDK_DOCKER": &buildEnvironmentVariable{
					"value": jsii.String("drop-in-replacement"),
				},
			},
		},
	},
})
```

## Migrating a pipeline type from V1 to V2

To migrate your pipeline type from V1 to V2, you just need to update the `pipelineType` property to `PipelineType.V2`.
This migration does not cause replacement of your pipeline.

When the `@aws-cdk/aws-codepipeline:defaultPipelineTypeToV2` feature flag is set to `true` (default for new projects),
the V2 type is selected by default if you do not specify a value for `pipelineType` property. Otherwise, the V1 type is selected.

See the [CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipeline-types-planning.html)
for more details on the differences between each type.

## Known Issues

There are some usability issues that are caused by underlying technology, and
cannot be remedied by CDK at this point. They are reproduced here for completeness.

* **Console links to other accounts will not work**: the AWS CodePipeline
  console will assume all links are relative to the current account. You will
  not be able to use the pipeline console to click through to a CloudFormation
  stack in a different account.
* **If a change set failed to apply the pipeline must be restarted**: if a change
  set failed to apply, it cannot be retried. The pipeline must be restarted from
  the top by clicking **Release Change**.
* **A stack that failed to create must be deleted manually**: if a stack
  failed to create on the first attempt, you must delete it using the
  CloudFormation console before starting the pipeline again by clicking
  **Release Change**.

# AWS CodePipeline Actions

This package contains Actions that can be used in a CodePipeline.

```go
import codepipeline "github.com/aws/aws-cdk-go/awscdk"
import codepipeline_actions "github.com/aws/aws-cdk-go/awscdk"
```

## Sources

### AWS CodeCommit

To use a CodeCommit Repository in a CodePipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
repo := codecommit.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	repositoryName: jsii.String("MyRepo"),
})

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
	pipelineName: jsii.String("MyPipeline"),
})
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("CodeCommit"),
	repository: repo,
	output: sourceOutput,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
		sourceAction,
	},
})
```

If you want to use existing role which can be used by on commit event rule.
You can specify the role object in eventRole property.

```go
// Example automatically generated from non-compiling source. May contain errors.
var repo repository
eventRole := iam.role.fromRoleArn(this, jsii.String("Event-role"), jsii.String("roleArn"))
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("CodeCommit"),
	repository: repo,
	output: codepipeline.NewArtifact(),
	eventRole: eventRole,
})
```

If you want to clone the entire CodeCommit repository (only available for CodeBuild actions),
you can set the `codeBuildCloneOutput` property to `true`:

```go
// Example automatically generated from non-compiling source. May contain errors.
var project pipelineProject
var repo repository

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("CodeCommit"),
	repository: repo,
	output: sourceOutput,
	codeBuildCloneOutput: jsii.Boolean(true),
})

buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	 // The build action must use the CodeCommitSourceAction output as input.
	outputs: []artifact{
		codepipeline.NewArtifact(),
	},
})
```

The CodeCommit source action emits variables:

```go
// Example automatically generated from non-compiling source. May contain errors.
var project pipelineProject
var repo repository

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("CodeCommit"),
	repository: repo,
	output: sourceOutput,
	variablesNamespace: jsii.String("MyNamespace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"COMMIT_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.commitId,
		},
	},
})
```

### GitHub

If you want to use a GitHub repository as the source, you must create:

* A [GitHub Access Token](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line),
  with scopes **repo** and **admin:repo_hook**.
* A [Secrets Manager Secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html)
  with the value of the **GitHub Access Token**. Pick whatever name you want (for example `my-github-token`).
  This token can be stored either as Plaintext or as a Secret key/value.
  If you stored the token as Plaintext,
  set `SecretValue.secretsManager('my-github-token')` as the value of `oauthToken`.
  If you stored it as a Secret key/value,
  you must set `SecretValue.secretsManager('my-github-token', { jsonField : 'my-github-token' })` as the value of `oauthToken`.

To use GitHub as the source of a CodePipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Read the secret from Secrets Manager
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
	actionName: jsii.String("GitHub_Source"),
	owner: jsii.String("awslabs"),
	repo: jsii.String("aws-cdk"),
	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	output: sourceOutput,
	branch: jsii.String("develop"),
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
		sourceAction,
	},
})
```

The GitHub source action emits variables:

```go
// Example automatically generated from non-compiling source. May contain errors.
var sourceOutput artifact
var project pipelineProject


sourceAction := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
	actionName: jsii.String("Github_Source"),
	output: sourceOutput,
	owner: jsii.String("my-owner"),
	repo: jsii.String("my-repo"),
	oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	variablesNamespace: jsii.String("MyNamespace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"COMMIT_URL": &buildEnvironmentVariable{
			"value": sourceAction.variables.commitUrl,
		},
	},
})
```

### BitBucket

CodePipeline can use a BitBucket Git repository as a source:

**Note**: you have to manually connect CodePipeline through the AWS Console with your BitBucket account.
This is a one-time operation for a given AWS account in a given region.
The simplest way to do that is to either start creating a new CodePipeline,
or edit an existing one, while being logged in to BitBucket.
Choose BitBucket as the source,
and grant CodePipeline permissions to your BitBucket account.
Copy & paste the Connection ARN that you get in the console,
or use the [`codestar-connections list-connections` AWS CLI operation](https://docs.aws.amazon.com/cli/latest/reference/codestar-connections/list-connections.html)
to find it.
After that, you can safely abort creating or editing the pipeline -
the connection has already been created.

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&codeStarConnectionsSourceActionProps{
	actionName: jsii.String("BitBucket_Source"),
	owner: jsii.String("aws"),
	repo: jsii.String("aws-cdk"),
	output: sourceOutput,
	connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
})
```

You can also use the `CodeStarConnectionsSourceAction` to connect to GitHub, in the same way
(you just have to select GitHub as the source when creating the connection in the console).

Similarly to `GitHubSourceAction`, `CodeStarConnectionsSourceAction` also emits the variables:

```go
// Example automatically generated from non-compiling source. May contain errors.
var project project


sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&codeStarConnectionsSourceActionProps{
	actionName: jsii.String("BitBucket_Source"),
	owner: jsii.String("aws"),
	repo: jsii.String("aws-cdk"),
	output: sourceOutput,
	connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
	variablesNamespace: jsii.String("SomeSpace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"COMMIT_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.commitId,
		},
	},
})
```

### AWS S3 Source

To use an S3 Bucket as a source in CodePipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceBucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	versioned: jsii.Boolean(true),
})

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
	actionName: jsii.String("S3Source"),
	bucket: sourceBucket,
	bucketKey: jsii.String("path/to/file.zip"),
	output: sourceOutput,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
		sourceAction,
	},
})
```

The region of the action will be determined by the region the bucket itself is in.
When using a newly created bucket,
that region will be taken from the stack the bucket belongs to;
for an imported bucket,
you can specify the region explicitly:

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceBucket := s3.bucket.fromBucketAttributes(this, jsii.String("SourceBucket"), &bucketAttributes{
	bucketName: jsii.String("my-bucket"),
	region: jsii.String("ap-southeast-1"),
})
```

By default, the Pipeline will poll the Bucket to detect changes.
You can change that behavior to use CloudWatch Events by setting the `trigger`
property to `S3Trigger.EVENTS` (it's `S3Trigger.POLL` by default).
If you do that, make sure the source Bucket is part of an AWS CloudTrail Trail -
otherwise, the CloudWatch Events will not be emitted,
and your Pipeline will not react to changes in the Bucket.
You can do it through the CDK:

```go
// Example automatically generated from non-compiling source. May contain errors.
import cloudtrail "github.com/aws/aws-cdk-go/awscdk"

var sourceBucket bucket

sourceOutput := codepipeline.NewArtifact()
key := "some/key.zip"
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
trail.addS3EventSelector([]s3EventSelector{
	&s3EventSelector{
		bucket: sourceBucket,
		objectPrefix: key,
	},
}, &addEventSelectorOptions{
	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
})
sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
	actionName: jsii.String("S3Source"),
	bucketKey: key,
	bucket: sourceBucket,
	output: sourceOutput,
	trigger: codepipeline_actions.s3Trigger_EVENTS,
})
```

The S3 source action emits variables:

```go
// Example automatically generated from non-compiling source. May contain errors.
var sourceBucket bucket

// later:
var project pipelineProject
key := "some/key.zip"
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
	actionName: jsii.String("S3Source"),
	bucketKey: key,
	bucket: sourceBucket,
	output: sourceOutput,
	variablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"VERSION_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.versionId,
		},
	},
})
```

### AWS ECR

To use an ECR Repository as a source in a Pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
import ecr "github.com/aws/aws-cdk-go/awscdk"

var ecrRepository repository

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
	actionName: jsii.String("ECR"),
	repository: ecrRepository,
	imageTag: jsii.String("some-tag"),
	 // optional, default: 'latest'
	output: sourceOutput,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
		sourceAction,
	},
})
```

The ECR source action emits variables:

```go
// Example automatically generated from non-compiling source. May contain errors.
import ecr "github.com/aws/aws-cdk-go/awscdk"
var ecrRepository repository

// later:
var project pipelineProject


sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
	actionName: jsii.String("Source"),
	output: sourceOutput,
	repository: ecrRepository,
	variablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"IMAGE_URI": &buildEnvironmentVariable{
			"value": sourceAction.variables.imageUri,
		},
	},
})
```

## Build & test

### AWS CodeBuild

Example of a CodeBuild Project used in a Pipeline, alongside CodeCommit:

```go
// Example automatically generated from non-compiling source. May contain errors.
var project pipelineProject

repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &repositoryProps{
	repositoryName: jsii.String("MyRepository"),
})
project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("CodeCommit"),
	repository: repository,
	output: sourceOutput,
})
buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	outputs: []artifact{
		codepipeline.NewArtifact(),
	},
	 // optional
	executeBatchBuild: jsii.Boolean(true),
	 // optional, defaults to false
	combineBatchBuildArtifacts: jsii.Boolean(true),
})

codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
	stages: []stageProps{
		&stageProps{
			stageName: jsii.String("Source"),
			actions: []iAction{
				sourceAction,
			},
		},
		&stageProps{
			stageName: jsii.String("Build"),
			actions: []*iAction{
				buildAction,
			},
		},
	},
})
```

The default category of the CodeBuild Action is `Build`;
if you want a `Test` Action instead,
override the `type` property:

```go
// Example automatically generated from non-compiling source. May contain errors.
var project pipelineProject

sourceOutput := codepipeline.NewArtifact()
testAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("IntegrationTest"),
	project: project,
	input: sourceOutput,
	type: codepipeline_actions.codeBuildActionType_TEST,
})
```

#### Multiple inputs and outputs

When you want to have multiple inputs and/or outputs for a Project used in a
Pipeline, instead of using the `secondarySources` and `secondaryArtifacts`
properties of the `Project` class, you need to use the `extraInputs` and
`outputs` properties of the CodeBuild CodePipeline
Actions. Example:

```go
// Example automatically generated from non-compiling source. May contain errors.
var repository1 repository
var repository2 repository

var project pipelineProject

sourceOutput1 := codepipeline.NewArtifact()
sourceAction1 := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("Source1"),
	repository: repository1,
	output: sourceOutput1,
})
sourceOutput2 := codepipeline.NewArtifact(jsii.String("source2"))
sourceAction2 := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("Source2"),
	repository: repository2,
	output: sourceOutput2,
})
buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("Build"),
	project: project,
	input: sourceOutput1,
	extraInputs: []artifact{
		sourceOutput2,
	},
	outputs: []*artifact{
		codepipeline.NewArtifact(jsii.String("artifact1")),
		 // for better buildspec readability - see below
		codepipeline.NewArtifact(jsii.String("artifact2")),
	},
})
```

**Note**: when a CodeBuild Action in a Pipeline has more than one output, it
only uses the `secondary-artifacts` field of the buildspec, never the
primary output specification directly under `artifacts`. Because of that, it
pays to explicitly name all output artifacts of that Action, like we did
above, so that you know what name to use in the buildspec.

Example buildspec for the above project:

```go
// Example automatically generated from non-compiling source. May contain errors.
project := codebuild.NewPipelineProject(this, jsii.String("MyProject"), &pipelineProjectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]interface{}{
			"build": map[string][]interface{}{
				"commands": []interface{}{
				},
			},
		},
		"artifacts": map[string]map[string]map[string]interface{}{
			"secondary-artifacts": map[string]map[string]interface{}{
				"artifact1": map[string]interface{}{
				},
				"artifact2": map[string]interface{}{
				},
			},
		},
	}),
})
```

#### Variables

The CodeBuild action emits variables.
Unlike many other actions, the variables are not static,
but dynamic, defined in the buildspec,
in the 'exported-variables' subsection of the 'env' section.
Example:

```go
// Example automatically generated from non-compiling source. May contain errors.
// later:
var project pipelineProject
sourceOutput := codepipeline.NewArtifact()
buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("Build1"),
	input: sourceOutput,
	project: codebuild.NewPipelineProject(this, jsii.String("Project"), &pipelineProjectProps{
		buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
			"version": jsii.String("0.2"),
			"env": map[string][]*string{
				"exported-variables": []*string{
					jsii.String("MY_VAR"),
				},
			},
			"phases": map[string]map[string]*string{
				"build": map[string]*string{
					"commands": jsii.String("export MY_VAR=\"some value\""),
				},
			},
		}),
	}),
	variablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"MyVar": &buildEnvironmentVariable{
			"value": buildAction.variable(jsii.String("MY_VAR")),
		},
	},
})
```

### Jenkins

In order to use Jenkins Actions in the Pipeline,
you first need to create a `JenkinsProvider`:

```go
// Example automatically generated from non-compiling source. May contain errors.
jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
	providerName: jsii.String("MyJenkinsProvider"),
	serverUrl: jsii.String("http://my-jenkins.com:8080"),
	version: jsii.String("2"),
})
```

If you've registered a Jenkins provider in a different CDK app,
or outside the CDK (in the CodePipeline AWS Console, for example),
you can import it:

```go
// Example automatically generated from non-compiling source. May contain errors.
jenkinsProvider := codepipeline_actions.jenkinsProvider.fromJenkinsProviderAttributes(this, jsii.String("JenkinsProvider"), &jenkinsProviderAttributes{
	providerName: jsii.String("MyJenkinsProvider"),
	serverUrl: jsii.String("http://my-jenkins.com:8080"),
	version: jsii.String("2"),
})
```

Note that a Jenkins provider
(identified by the provider name-category(build/test)-version tuple)
must always be registered in the given account, in the given AWS region,
before it can be used in CodePipeline.

With a `JenkinsProvider`,
we can create a Jenkins Action:

```go
// Example automatically generated from non-compiling source. May contain errors.
var jenkinsProvider jenkinsProvider

buildAction := codepipeline_actions.NewJenkinsAction(&jenkinsActionProps{
	actionName: jsii.String("JenkinsBuild"),
	jenkinsProvider: jenkinsProvider,
	projectName: jsii.String("MyProject"),
	type: codepipeline_actions.jenkinsActionType_BUILD,
})
```

## Deploy

### AWS CloudFormation

This module contains Actions that allows you to deploy to CloudFormation from AWS CodePipeline.

For example, the following code fragment defines a pipeline that automatically deploys a CloudFormation template
directly from a CodeCommit repository, with a manual approval step in between to confirm the changes:

```go
// Source stage: read from repository
repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
	repositoryName: jsii.String("template-repo"),
})
sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	actionName: jsii.String("Source"),
	repository: repo,
	output: sourceOutput,
	trigger: cpactions.codeCommitTrigger_POLL,
})
sourceStage := map[string]interface{}{
	"stageName": jsii.String("Source"),
	"actions": []CodeCommitSourceAction{
		source,
	},
}

// Deployment stage: create and deploy changeset with manual approval
stackName := "OurStack"
changeSetName := "StagedChangeSet"

prodStage := map[string]interface{}{
	"stageName": jsii.String("Deploy"),
	"actions": []interface{}{
		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
			"actionName": jsii.String("PrepareChanges"),
			"stackName": jsii.String(stackName),
			"changeSetName": jsii.String(changeSetName),
			"adminPermissions": jsii.Boolean(true),
			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
			"runOrder": jsii.Number(1),
		}),
		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
			"actionName": jsii.String("ApproveChanges"),
			"runOrder": jsii.Number(2),
		}),
		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
			"actionName": jsii.String("ExecuteChanges"),
			"stackName": jsii.String(stackName),
			"changeSetName": jsii.String(changeSetName),
			"runOrder": jsii.Number(3),
		}),
	},
}

codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
	stages: []stageProps{
		sourceStage,
		prodStage,
	},
})
```

See [the AWS documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline.html)
for more details about using CloudFormation in CodePipeline.

#### Actions for updating individual CloudFormation Stacks

This package contains the following CloudFormation actions:

* **CloudFormationCreateUpdateStackAction** - Deploy a CloudFormation template directly from the pipeline. The indicated stack is created,
  or updated if it already exists. If the stack is in a failure state, deployment will fail (unless `replaceOnFailure`
  is set to `true`, in which case it will be destroyed and recreated).
* **CloudFormationDeleteStackAction** - Delete the stack with the given name.
* **CloudFormationCreateReplaceChangeSetAction** - Prepare a change set to be applied later. You will typically use change sets if you want
  to manually verify the changes that are being staged, or if you want to separate the people (or system) preparing the
  changes from the people (or system) applying the changes.
* **CloudFormationExecuteChangeSetAction** - Execute a change set prepared previously.

#### Actions for deploying CloudFormation StackSets to multiple accounts

You can use CloudFormation StackSets to deploy the same CloudFormation template to multiple
accounts in a managed way. If you use AWS Organizations, StackSets can be deployed to
all accounts in a particular Organizational Unit (OU), and even automatically to new
accounts as soon as they are added to a particular OU. For more information, see
the [Working with StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html)
section of the CloudFormation developer guide.

The actions available for updating StackSets are:

* **CloudFormationDeployStackSetAction** - Create or update a CloudFormation StackSet directly from the pipeline, optionally
  immediately create and update Stack Instances as well.
* **CloudFormationDeployStackInstancesAction** - Update outdated Stack Instances using the current version of the StackSet.

Here's an example of using both of these actions:

```go
// Example automatically generated from non-compiling source. May contain errors.
var pipeline pipeline
var sourceOutput artifact


pipeline.addStage(&stageOptions{
	stageName: jsii.String("DeployStackSets"),
	actions: []iAction{
		// First, update the StackSet itself with the newest template
		codepipeline_actions.NewCloudFormationDeployStackSetAction(&cloudFormationDeployStackSetActionProps{
			actionName: jsii.String("UpdateStackSet"),
			runOrder: jsii.Number(1),
			stackSetName: jsii.String("MyStackSet"),
			template: codepipeline_actions.stackSetTemplate.fromArtifactPath(sourceOutput.atPath(jsii.String("template.yaml"))),

			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
			deploymentModel: codepipeline_actions.stackSetDeploymentModel.selfManaged(),
			// This deploys to a set of accounts
			stackInstances: codepipeline_actions.stackInstances.inAccounts([]*string{
				jsii.String("111111111111"),
			}, []*string{
				jsii.String("us-east-1"),
				jsii.String("eu-west-1"),
			}),
		}),

		// Afterwards, update/create additional instances in other accounts
		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&cloudFormationDeployStackInstancesActionProps{
			actionName: jsii.String("AddMoreInstances"),
			runOrder: jsii.Number(2),
			stackSetName: jsii.String("MyStackSet"),
			stackInstances: codepipeline_actions.*stackInstances.inAccounts([]*string{
				jsii.String("222222222222"),
				jsii.String("333333333333"),
			}, []*string{
				jsii.String("us-east-1"),
				jsii.String("eu-west-1"),
			}),
		}),
	},
})
```

#### Lambda deployed through CodePipeline

If you want to deploy your Lambda through CodePipeline,
and you don't use assets (for example, because your CDK code and Lambda code are separate),
you can use a special Lambda `Code` class, `CfnParametersCode`.
Note that your Lambda must be in a different Stack than your Pipeline.
The Lambda itself will be deployed, alongside the entire Stack it belongs to,
using a CloudFormation CodePipeline Action. Example:

```go
lambdaStack := cdk.NewStack(app, jsii.String("LambdaStack"))
lambdaCode := lambda.code.fromCfnParameters()
lambda.NewFunction(lambdaStack, jsii.String("Lambda"), &functionProps{
	code: lambdaCode,
	handler: jsii.String("index.handler"),
	runtime: lambda.runtime_NODEJS_14_X(),
})
// other resources that your Lambda needs, added to the lambdaStack...

pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("Pipeline"))

// add the source code repository containing this code to your Pipeline,
// and the source code of the Lambda Function, if they're separate
cdkSourceOutput := codepipeline.NewArtifact()
cdkSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	repository: codecommit.NewRepository(pipelineStack, jsii.String("CdkCodeRepo"), &repositoryProps{
		repositoryName: jsii.String("CdkCodeRepo"),
	}),
	actionName: jsii.String("CdkCode_Source"),
	output: cdkSourceOutput,
})
lambdaSourceOutput := codepipeline.NewArtifact()
lambdaSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
	repository: codecommit.NewRepository(pipelineStack, jsii.String("LambdaCodeRepo"), &repositoryProps{
		repositoryName: jsii.String("LambdaCodeRepo"),
	}),
	actionName: jsii.String("LambdaCode_Source"),
	output: lambdaSourceOutput,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
		cdkSourceAction,
		lambdaSourceAction,
	},
})

// synthesize the Lambda CDK template, using CodeBuild
// the below values are just examples, assuming your CDK code is in TypeScript/JavaScript -
// adjust the build environment and/or commands accordingly
cdkBuildProject := codebuild.NewProject(pipelineStack, jsii.String("CdkBuildProject"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.linuxBuildImage_UBUNTU_14_04_NODEJS_10_1_0(),
	},
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string]*string{
			"install": map[string]*string{
				"commands": jsii.String("npm install"),
			},
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("npm run build"),
					jsii.String("npm run cdk synth LambdaStack -- -o ."),
				},
			},
		},
		"artifacts": map[string]*string{
			"files": jsii.String("LambdaStack.template.yaml"),
		},
	}),
})
cdkBuildOutput := codepipeline.NewArtifact()
cdkBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CDK_Build"),
	project: cdkBuildProject,
	input: cdkSourceOutput,
	outputs: []artifact{
		cdkBuildOutput,
	},
})

// build your Lambda code, using CodeBuild
// again, this example assumes your Lambda is written in TypeScript/JavaScript -
// make sure to adjust the build environment and/or commands if they don't match your specific situation
lambdaBuildProject := codebuild.NewProject(pipelineStack, jsii.String("LambdaBuildProject"), &projectProps{
	environment: &buildEnvironment{
		buildImage: codebuild.*linuxBuildImage_UBUNTU_14_04_NODEJS_10_1_0(),
	},
	buildSpec: codebuild.*buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string]*string{
			"install": map[string]*string{
				"commands": jsii.String("npm install"),
			},
			"build": map[string]*string{
				"commands": jsii.String("npm run build"),
			},
		},
		"artifacts": map[string][]*string{
			"files": []*string{
				jsii.String("index.js"),
				jsii.String("node_modules/**/*"),
			},
		},
	}),
})
lambdaBuildOutput := codepipeline.NewArtifact()
lambdaBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("Lambda_Build"),
	project: lambdaBuildProject,
	input: lambdaSourceOutput,
	outputs: []*artifact{
		lambdaBuildOutput,
	},
})

pipeline.addStage(&stageOptions{
	stageName: jsii.String("Build"),
	actions: []*iAction{
		cdkBuildAction,
		lambdaBuildAction,
	},
})

// finally, deploy your Lambda Stack
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []*iAction{
		codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
			actionName: jsii.String("Lambda_CFN_Deploy"),
			templatePath: cdkBuildOutput.atPath(jsii.String("LambdaStack.template.yaml")),
			stackName: jsii.String("LambdaStackDeployedName"),
			adminPermissions: jsii.Boolean(true),
			parameterOverrides: lambdaCode.assign(lambdaBuildOutput.s3Location),
			extraInputs: []*artifact{
				lambdaBuildOutput,
			},
		}),
	},
})
```

#### Cross-account actions

If you want to update stacks in a different account,
pass the `account` property when creating the action:

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
	actionName: jsii.String("CloudFormationCreateUpdate"),
	stackName: jsii.String("MyStackName"),
	adminPermissions: jsii.Boolean(true),
	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
	account: jsii.String("123456789012"),
})
```

This will create a new stack, called `<PipelineStackName>-support-123456789012`, in your `App`,
that will contain the role that the pipeline will assume in account 123456789012 before executing this action.
This support stack will automatically be deployed before the stack containing the pipeline.

You can also pass a role explicitly when creating the action -
in that case, the `account` property is ignored,
and the action will operate in the same account the role belongs to:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"

// in stack for account 123456789012...
var otherAccountStack stack

actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &roleProps{
	assumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
	// the role has to have a physical name set
	roleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
})

// in the pipeline stack...
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
	actionName: jsii.String("CloudFormationCreateUpdate"),
	stackName: jsii.String("MyStackName"),
	adminPermissions: jsii.Boolean(true),
	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
	role: actionRole,
})
```

### AWS CodeDeploy

#### Server deployments

To use CodeDeploy for EC2/on-premise deployments in a Pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
var deploymentGroup serverDeploymentGroup
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
	pipelineName: jsii.String("MyPipeline"),
})

// add the source and build Stages to the Pipeline...
buildOutput := codepipeline.NewArtifact()
deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&codeDeployServerDeployActionProps{
	actionName: jsii.String("CodeDeploy"),
	input: buildOutput,
	deploymentGroup: deploymentGroup,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []iAction{
		deployAction,
	},
})
```

##### Lambda deployments

To use CodeDeploy for blue-green Lambda deployments in a Pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
lambdaCode := lambda.code.fromCfnParameters()
func := lambda.NewFunction(this, jsii.String("Lambda"), &functionProps{
	code: lambdaCode,
	handler: jsii.String("index.handler"),
	runtime: lambda.runtime_NODEJS_14_X(),
})
// used to make sure each CDK synthesis produces a different Version
version := func.currentVersion
alias := lambda.NewAlias(this, jsii.String("LambdaAlias"), &aliasProps{
	aliasName: jsii.String("Prod"),
	version: version,
})

codedeploy.NewLambdaDeploymentGroup(this, jsii.String("DeploymentGroup"), &lambdaDeploymentGroupProps{
	alias: alias,
	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
})
```

Then, you need to create your Pipeline Stack,
where you will define your Pipeline,
and deploy the `lambdaStack` using a CloudFormation CodePipeline Action
(see above for a complete example).

### ECS

CodePipeline can deploy an ECS service.
The deploy Action receives one input Artifact which contains the [image definition file](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions):

```go
// Example automatically generated from non-compiling source. May contain errors.
import ecs "github.com/aws/aws-cdk-go/awscdk"

var service fargateService

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
buildOutput := codepipeline.NewArtifact()
deployStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []iAction{
		codepipeline_actions.NewEcsDeployAction(&ecsDeployActionProps{
			actionName: jsii.String("DeployAction"),
			service: service,
			// if your file is called imagedefinitions.json,
			// use the `input` property,
			// and leave out the `imageFile` property
			input: buildOutput,
			// if your file name is _not_ imagedefinitions.json,
			// use the `imageFile` property,
			// and leave out the `input` property
			imageFile: buildOutput.atPath(jsii.String("imageDef.json")),
			deploymentTimeout: awscdk.Duration.minutes(jsii.Number(60)),
		}),
	},
})
```

#### Deploying ECS applications to existing services

CodePipeline can deploy to an existing ECS service which uses the
[ECS service ARN format that contains the Cluster name](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids).
This also works if the service is in a different account and/or region than the pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
import ecs "github.com/aws/aws-cdk-go/awscdk"


service := ecs.baseService.fromServiceArnWithCluster(this, jsii.String("EcsService"), jsii.String("arn:aws:ecs:us-east-1:123456789012:service/myClusterName/myServiceName"))
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
buildOutput := codepipeline.NewArtifact()
// add source and build stages to the pipeline as usual...
deployStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []iAction{
		codepipeline_actions.NewEcsDeployAction(&ecsDeployActionProps{
			actionName: jsii.String("DeployAction"),
			service: service,
			input: buildOutput,
		}),
	},
})
```

When deploying across accounts, especially in a CDK Pipelines self-mutating pipeline,
it is recommended to provide the `role` property to the `EcsDeployAction`.
The Role will need to have permissions assigned to it for ECS deployment.
See [the CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-custom-role.html#how-to-update-role-new-services)
for the permissions needed.

#### Deploying ECS applications stored in a separate source code repository

The idiomatic CDK way of deploying an ECS application is to have your Dockerfiles and your CDK code in the same source code repository,
leveraging [Docker Assets](https://docs.aws.amazon.com/cdk/latest/guide/assets.html#assets_types_docker),
and use the [CDK Pipelines module](https://docs.aws.amazon.com/cdk/api/latest/docs/pipelines-readme.html).

However, if you want to deploy a Docker application whose source code is kept in a separate version control repository than the CDK code,
you can use the `TagParameterContainerImage` class from the ECS module.
Here's an example:

```go
/**
 * These are the construction properties for {@link EcsAppStack}.
 * They extend the standard Stack properties,
 * but also require providing the ContainerImage that the service will use.
 * That Image will be provided from the Stack containing the CodePipeline.
 */
type ecsAppStackProps struct {
	stackProps
	image containerImage
}

/**
 * This is the Stack containing a simple ECS Service that uses the provided ContainerImage.
 */
type EcsAppStack struct {
	stack
}

func NewEcsAppStack(scope construct, id *string, props ecsAppStackProps) *EcsAppStack {
	this := &EcsAppStack{}
	cdk.NewStack_Override(this, scope, id, props)

	taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDefinition"), &taskDefinitionProps{
		compatibility: ecs.compatibility_FARGATE,
		cpu: jsii.String("1024"),
		memoryMiB: jsii.String("2048"),
	})
	taskDefinition.addContainer(jsii.String("AppContainer"), &containerDefinitionOptions{
		image: props.image,
	})
	ecs.NewFargateService(this, jsii.String("EcsService"), &fargateServiceProps{
		taskDefinition: taskDefinition,
		cluster: ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
			vpc: ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
				maxAzs: jsii.Number(1),
			}),
		}),
	})
	return this
}

/**
 * This is the Stack containing the CodePipeline definition that deploys an ECS Service.
 */
type PipelineStack struct {
	stack
	tagParameterContainerImage tagParameterContainerImage
}tagParameterContainerImage tagParameterContainerImage

func NewPipelineStack(scope construct, id *string, props stackProps) *PipelineStack {
	this := &PipelineStack{}
	cdk.NewStack_Override(this, scope, id, props)

	/* ********** ECS part **************** */

	// this is the ECR repository where the built Docker image will be pushed
	appEcrRepo := ecr.NewRepository(this, jsii.String("EcsDeployRepository"))
	// the build that creates the Docker image, and pushes it to the ECR repo
	appCodeDockerBuild := codebuild.NewPipelineProject(this, jsii.String("AppCodeDockerImageBuildAndPushProject"), &pipelineProjectProps{
		environment: &buildEnvironment{
			// we need to run Docker
			privileged: jsii.Boolean(true),
		},
		buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
			"version": jsii.String("0.2"),
			"phases": map[string]map[string][]*string{
				"build": map[string][]*string{
					"commands": []*string{
						jsii.String("$(aws ecr get-login --region $AWS_DEFAULT_REGION --no-include-email)"),
						jsii.String("docker build -t $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION ."),
					},
				},
				"post_build": map[string][]*string{
					"commands": []*string{
						jsii.String("docker push $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION"),
						jsii.String("export imageTag=$CODEBUILD_RESOLVED_SOURCE_VERSION"),
					},
				},
			},
			"env": map[string][]*string{
				// save the imageTag environment variable as a CodePipeline Variable
				"exported-variables": []*string{
					jsii.String("imageTag"),
				},
			},
		}),
		environmentVariables: map[string]buildEnvironmentVariable{
			"REPOSITORY_URI": &buildEnvironmentVariable{
				"value": appEcrRepo.repositoryUri,
			},
		},
	})
	// needed for `docker push`
	appEcrRepo.grantPullPush(appCodeDockerBuild)
	// create the ContainerImage used for the ECS application Stack
	this.tagParameterContainerImage = ecs.NewTagParameterContainerImage(appEcrRepo)

	cdkCodeBuild := codebuild.NewPipelineProject(this, jsii.String("CdkCodeBuildProject"), &pipelineProjectProps{
		buildSpec: codebuild.*buildSpec.fromObject(map[string]interface{}{
			"version": jsii.String("0.2"),
			"phases": map[string]map[string][]*string{
				"install": map[string][]*string{
					"commands": []*string{
						jsii.String("npm install"),
					},
				},
				"build": map[string][]*string{
					"commands": []*string{
						jsii.String("npx cdk synth --verbose"),
					},
				},
			},
			"artifacts": map[string]*string{
				// store the entire Cloud Assembly as the output artifact
				"base-directory": jsii.String("cdk.out"),
				"files": jsii.String("**/*"),
			},
		}),
	})

	/* ********** Pipeline part **************** */

	appCodeSourceOutput := codepipeline.NewArtifact()
	cdkCodeSourceOutput := codepipeline.NewArtifact()
	cdkCodeBuildOutput := codepipeline.NewArtifact()
	appCodeBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
		actionName: jsii.String("AppCodeDockerImageBuildAndPush"),
		project: appCodeDockerBuild,
		input: appCodeSourceOutput,
	})
	codepipeline.NewPipeline(this, jsii.String("CodePipelineDeployingEcsApplication"), &pipelineProps{
		artifactBucket: s3.NewBucket(this, jsii.String("ArtifactBucket"), &bucketProps{
			removalPolicy: cdk.removalPolicy_DESTROY,
		}),
		stages: []stageProps{
			&stageProps{
				stageName: jsii.String("Source"),
				actions: []iAction{
					// this is the Action that takes the source of your application code
					codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
						actionName: jsii.String("AppCodeSource"),
						repository: codecommit.NewRepository(this, jsii.String("AppCodeSourceRepository"), &repositoryProps{
							repositoryName: jsii.String("AppCodeSourceRepository"),
						}),
						output: appCodeSourceOutput,
					}),
					// this is the Action that takes the source of your CDK code
					// (which would probably include this Pipeline code as well)
					codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
						actionName: jsii.String("CdkCodeSource"),
						repository: codecommit.NewRepository(this, jsii.String("CdkCodeSourceRepository"), &repositoryProps{
							repositoryName: jsii.String("CdkCodeSourceRepository"),
						}),
						output: cdkCodeSourceOutput,
					}),
				},
			},
			&stageProps{
				stageName: jsii.String("Build"),
				actions: []*iAction{
					appCodeBuildAction,
					codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
						actionName: jsii.String("CdkCodeBuildAndSynth"),
						project: cdkCodeBuild,
						input: cdkCodeSourceOutput,
						outputs: []artifact{
							cdkCodeBuildOutput,
						},
					}),
				},
			},
			&stageProps{
				stageName: jsii.String("Deploy"),
				actions: []*iAction{
					codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
						actionName: jsii.String("CFN_Deploy"),
						stackName: jsii.String("SampleEcsStackDeployedFromCodePipeline"),
						// this name has to be the same name as used below in the CDK code for the application Stack
						templatePath: cdkCodeBuildOutput.atPath(jsii.String("EcsStackDeployedInPipeline.template.json")),
						adminPermissions: jsii.Boolean(true),
						parameterOverrides: map[string]interface{}{
							// read the tag pushed to the ECR repository from the CodePipeline Variable saved by the application build step,
							// and pass it as the CloudFormation Parameter for the tag
							this.tagParameterContainerImage.tagParameterName: appCodeBuildAction.variable(jsii.String("imageTag")),
						},
					}),
				},
			},
		},
	})
	return this
}

app := cdk.NewApp()

// the CodePipeline Stack needs to be created first
pipelineStack := NewPipelineStack(app, jsii.String("aws-cdk-pipeline-ecs-separate-sources"))
// we supply the image to the ECS application Stack from the CodePipeline Stack
// we supply the image to the ECS application Stack from the CodePipeline Stack
NewEcsAppStack(app, jsii.String("EcsStackDeployedInPipeline"), &ecsAppStackProps{
	image: pipelineStack.tagParameterContainerImage,
})
```

### AWS S3 Deployment

To use an S3 Bucket as a deployment target in CodePipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceOutput := codepipeline.NewArtifact()
targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
deployAction := codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
	actionName: jsii.String("S3Deploy"),
	bucket: targetBucket,
	input: sourceOutput,
})
deployStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []iAction{
		deployAction,
	},
})
```

#### Invalidating the CloudFront cache when deploying to S3

There is currently no native support in CodePipeline for invalidating a CloudFront cache after deployment.
One workaround is to add another build step after the deploy step,
and use the AWS CLI to invalidate the cache:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Create a Cloudfront Web Distribution
import cloudfront "github.com/aws/aws-cdk-go/awscdk"
var distribution distribution


// Create the build project that will invalidate the cache
invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &pipelineProjectProps{
	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
				},
			},
		},
	}),
	environmentVariables: map[string]buildEnvironmentVariable{
		"CLOUDFRONT_ID": &buildEnvironmentVariable{
			"value": distribution.distributionId,
		},
	},
})

// Add Cloudfront invalidation permissions to the project
distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.account, distribution.distributionId)
invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
	resources: []*string{
		distributionArn,
	},
	actions: []*string{
		jsii.String("cloudfront:CreateInvalidation"),
	},
}))

// Create the pipeline (here only the S3 deploy and Invalidate cache build)
deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
deployInput := codepipeline.NewArtifact()
codepipeline.NewPipeline(this, jsii.String("Pipeline"), &pipelineProps{
	stages: []stageProps{
		&stageProps{
			stageName: jsii.String("Deploy"),
			actions: []iAction{
				codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
					actionName: jsii.String("S3Deploy"),
					bucket: deployBucket,
					input: deployInput,
					runOrder: jsii.Number(1),
				}),
				codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
					actionName: jsii.String("InvalidateCache"),
					project: invalidateBuildProject,
					input: deployInput,
					runOrder: jsii.Number(2),
				}),
			},
		},
	},
})
```

### Elastic Beanstalk Deployment

To deploy an Elastic Beanstalk Application in CodePipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
sourceOutput := codepipeline.NewArtifact()
targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
deployAction := codepipeline_actions.NewElasticBeanstalkDeployAction(&elasticBeanstalkDeployActionProps{
	actionName: jsii.String("ElasticBeanstalkDeploy"),
	input: sourceOutput,
	environmentName: jsii.String("envName"),
	applicationName: jsii.String("appName"),
})

deployStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
	actions: []iAction{
		deployAction,
	},
})
```

### Alexa Skill

You can deploy to Alexa using CodePipeline with the following Action:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Read the secrets from ParameterStore
clientId := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientId"))
clientSecret := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientSecret"))
refreshToken := awscdk.SecretValue.secretsManager(jsii.String("AlexaRefreshToken"))

// Add deploy action
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewAlexaSkillDeployAction(&alexaSkillDeployActionProps{
	actionName: jsii.String("DeploySkill"),
	runOrder: jsii.Number(1),
	input: sourceOutput,
	clientId: clientId.toString(),
	clientSecret: clientSecret,
	refreshToken: refreshToken,
	skillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
})
```

If you need manifest overrides you can specify them as `parameterOverridesArtifact` in the action:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Deploy some CFN change set and store output
executeOutput := codepipeline.NewArtifact(jsii.String("CloudFormation"))
executeChangeSetAction := codepipeline_actions.NewCloudFormationExecuteChangeSetAction(&cloudFormationExecuteChangeSetActionProps{
	actionName: jsii.String("ExecuteChangesTest"),
	runOrder: jsii.Number(2),
	stackName: jsii.String("MyStack"),
	changeSetName: jsii.String("MyChangeSet"),
	outputFileName: jsii.String("overrides.json"),
	output: executeOutput,
})

// Provide CFN output as manifest overrides
clientId := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientId"))
clientSecret := awscdk.SecretValue.secretsManager(jsii.String("AlexaClientSecret"))
refreshToken := awscdk.SecretValue.secretsManager(jsii.String("AlexaRefreshToken"))
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewAlexaSkillDeployAction(&alexaSkillDeployActionProps{
	actionName: jsii.String("DeploySkill"),
	runOrder: jsii.Number(1),
	input: sourceOutput,
	parameterOverridesArtifact: executeOutput,
	clientId: clientId.toString(),
	clientSecret: clientSecret,
	refreshToken: refreshToken,
	skillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
})
```

### AWS Service Catalog

You can deploy a CloudFormation template to an existing Service Catalog product with the following Action:

```go
// Example automatically generated from non-compiling source. May contain errors.
cdkBuildOutput := codepipeline.NewArtifact()
serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&serviceCatalogDeployActionBeta1Props{
	actionName: jsii.String("ServiceCatalogDeploy"),
	templatePath: cdkBuildOutput.atPath(jsii.String("Sample.template.json")),
	productVersionName: jsii.String("Version - " + date.now.toString),
	productVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
	productId: jsii.String("prod-XXXXXXXX"),
})
```

## Approve & invoke

### Manual approval Action

This package contains an Action that stops the Pipeline until someone manually clicks the approve button:

```go
// Example automatically generated from non-compiling source. May contain errors.
import sns "github.com/aws/aws-cdk-go/awscdk"


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
approveStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Approve"),
})
manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
	actionName: jsii.String("Approve"),
	notificationTopic: sns.NewTopic(this, jsii.String("Topic")),
	 // optional
	notifyEmails: []*string{
		jsii.String("some_email@example.com"),
	},
	 // optional
	additionalInformation: jsii.String("additional info"),
})
approveStage.addAction(manualApprovalAction)
```

If the `notificationTopic` has not been provided,
but `notifyEmails` were,
a new SNS Topic will be created
(and accessible through the `notificationTopic` property of the Action).

If you want to grant a principal permissions to approve the changes,
you can invoke the method `grantManualApproval` passing it a `IGrantable`:

```go
// Example automatically generated from non-compiling source. May contain errors.
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
approveStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Approve"),
})
manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
	actionName: jsii.String("Approve"),
})
approveStage.addAction(manualApprovalAction)

role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
	service: jsii.String("iam"),
	resource: jsii.String("role"),
	resourceName: jsii.String("Admin"),
}, this))
manualApprovalAction.grantManualApproval(role)
```

### AWS Lambda

This module contains an Action that allows you to invoke a Lambda function in a Pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
var fn function

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
	actionName: jsii.String("Lambda"),
	lambda: fn,
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("Lambda"),
	actions: []iAction{
		lambdaAction,
	},
})
```

The Lambda Action can have up to 5 inputs,
and up to 5 outputs:

```go
// Example automatically generated from non-compiling source. May contain errors.
var fn function

sourceOutput := codepipeline.NewArtifact()
buildOutput := codepipeline.NewArtifact()
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
	actionName: jsii.String("Lambda"),
	inputs: []artifact{
		sourceOutput,
		buildOutput,
	},
	outputs: []*artifact{
		codepipeline.NewArtifact(jsii.String("Out1")),
		codepipeline.NewArtifact(jsii.String("Out2")),
	},
	lambda: fn,
})
```

The Lambda Action supports custom user parameters that pipeline
will pass to the Lambda function:

```go
// Example automatically generated from non-compiling source. May contain errors.
var fn function


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
	actionName: jsii.String("Lambda"),
	lambda: fn,
	userParameters: map[string]interface{}{
		"foo": jsii.String("bar"),
		"baz": jsii.String("qux"),
	},
	// OR
	userParametersString: jsii.String("my-parameter-string"),
})
```

The Lambda invoke action emits variables.
Unlike many other actions, the variables are not static,
but dynamic, defined by the function calling the `PutJobSuccessResult`
API with the `outputVariables` property filled with the map of variables
Example:

```go
// Example automatically generated from non-compiling source. May contain errors.
// later:
var project pipelineProject
lambdaInvokeAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
	actionName: jsii.String("Lambda"),
	lambda: lambda.NewFunction(this, jsii.String("Func"), &functionProps{
		runtime: lambda.runtime_NODEJS_14_X(),
		handler: jsii.String("index.handler"),
		code: lambda.code.fromInline(jsii.String("\n        const AWS = require('aws-sdk');\n\n        exports.handler = async function(event, context) {\n            const codepipeline = new AWS.CodePipeline();\n            await codepipeline.putJobSuccessResult({\n                jobId: event['CodePipeline.job'].id,\n                outputVariables: {\n                    MY_VAR: \"some value\",\n                },\n            }).promise();\n        }\n    ")),
	}),
	variablesNamespace: jsii.String("MyNamespace"),
})
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	environmentVariables: map[string]buildEnvironmentVariable{
		"MyVar": &buildEnvironmentVariable{
			"value": lambdaInvokeAction.variable(jsii.String("MY_VAR")),
		},
	},
})
```

See [the AWS documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html)
on how to write a Lambda function invoked from CodePipeline.

### AWS Step Functions

This module contains an Action that allows you to invoke a Step Function in a Pipeline:

```go
// Example automatically generated from non-compiling source. May contain errors.
import stepfunctions "github.com/aws/aws-cdk-go/awscdk"

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
startState := stepfunctions.NewPass(this, jsii.String("StartState"))
simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
	definition: startState,
})
stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
	actionName: jsii.String("Invoke"),
	stateMachine: simpleStateMachine,
	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
		"IsHelloWorldExample": jsii.Boolean(true),
	}),
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("StepFunctions"),
	actions: []iAction{
		stepFunctionAction,
	},
})
```

The `StateMachineInput` can be created with one of 2 static factory methods:
`literal`, which takes an arbitrary map as its only argument, or `filePath`:

```go
// Example automatically generated from non-compiling source. May contain errors.
import stepfunctions "github.com/aws/aws-cdk-go/awscdk"


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
inputArtifact := codepipeline.NewArtifact()
startState := stepfunctions.NewPass(this, jsii.String("StartState"))
simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
	definition: startState,
})
stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
	actionName: jsii.String("Invoke"),
	stateMachine: simpleStateMachine,
	stateMachineInput: codepipeline_actions.stateMachineInput.filePath(inputArtifact.atPath(jsii.String("assets/input.json"))),
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("StepFunctions"),
	actions: []iAction{
		stepFunctionAction,
	},
})
```

See [the AWS documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-StepFunctions.html)
for information on Action structure reference.

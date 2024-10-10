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
repo := codecommit.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	RepositoryName: jsii.String("MyRepo"),
})

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
	PipelineName: jsii.String("MyPipeline"),
})
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: repo,
	Output: sourceOutput,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
		sourceAction,
	},
})
```

If you want to use existing role which can be used by on commit event rule.
You can specify the role object in eventRole property.

```go
var repo repository
eventRole := iam.Role_FromRoleArn(this, jsii.String("Event-role"), jsii.String("roleArn"))
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: repo,
	Output: codepipeline.NewArtifact(),
	EventRole: EventRole,
})
```

If you want to clone the entire CodeCommit repository (only available for CodeBuild actions),
you can set the `codeBuildCloneOutput` property to `true`:

```go
var project pipelineProject
var repo repository

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: repo,
	Output: sourceOutput,
	CodeBuildCloneOutput: jsii.Boolean(true),
})

buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	 // The build action must use the CodeCommitSourceAction output as input.
	Outputs: []artifact{
		codepipeline.NewArtifact(),
	},
})
```

The CodeCommit source action emits variables:

```go
var project pipelineProject
var repo repository

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: repo,
	Output: sourceOutput,
	VariablesNamespace: jsii.String("MyNamespace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
		"COMMIT_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.commitId,
		},
	},
})
```

If you want to use a custom event for your `CodeCommitSourceAction`, you can pass in
a `customEventRule` which needs an event pattern (see [here](https://docs.aws.amazon.com/codecommit/latest/userguide/monitoring-events.html)) and an `IRuleTarget` (see [here](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_events_targets-readme.html))

```go
var repo repository
var lambdaFuntion function
eventPattern := map[string]interface{}{
	"detail-type": []*string{
		jsii.String("CodeCommit Repository State Change"),
	},
	"resources": []*string{
		jsii.String("foo"),
	},
	"source": []*string{
		jsii.String("aws.codecommit"),
	},
	"detail": map[string][]*string{
		"referenceType": []*string{
			jsii.String("branch"),
		},
		"event": []*string{
			jsii.String("referenceCreated"),
			jsii.String("referenceUpdated"),
		},
		"referenceName": []*string{
			jsii.String("master"),
		},
	},
}
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: repo,
	Output: sourceOutput,
	CustomEventRule: map[string]interface{}{
		"eventPattern": eventPattern,
		"target": targets.NewLambdaFunction(lambdaFuntion),
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
// Read the secret from Secrets Manager
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewGitHubSourceAction(&GitHubSourceActionProps{
	ActionName: jsii.String("GitHub_Source"),
	Owner: jsii.String("awslabs"),
	Repo: jsii.String("aws-cdk"),
	OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	Output: sourceOutput,
	Branch: jsii.String("develop"),
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
		sourceAction,
	},
})
```

The GitHub source action emits variables:

```go
var sourceOutput artifact
var project pipelineProject


sourceAction := codepipeline_actions.NewGitHubSourceAction(&GitHubSourceActionProps{
	ActionName: jsii.String("Github_Source"),
	Output: sourceOutput,
	Owner: jsii.String("my-owner"),
	Repo: jsii.String("my-repo"),
	OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	VariablesNamespace: jsii.String("MyNamespace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
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
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
	ActionName: jsii.String("BitBucket_Source"),
	Owner: jsii.String("aws"),
	Repo: jsii.String("aws-cdk"),
	Output: sourceOutput,
	ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
})
```

You can also use the `CodeStarConnectionsSourceAction` to connect to GitHub, in the same way
(you just have to select GitHub as the source when creating the connection in the console).

Similarly to `GitHubSourceAction`, `CodeStarConnectionsSourceAction` also emits the variables:

```go
var project project


sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeStarConnectionsSourceAction(&CodeStarConnectionsSourceActionProps{
	ActionName: jsii.String("BitBucket_Source"),
	Owner: jsii.String("aws"),
	Repo: jsii.String("aws-cdk"),
	Output: sourceOutput,
	ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"),
	VariablesNamespace: jsii.String("SomeSpace"),
})

// later:

// later:
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
		"COMMIT_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.commitId,
		},
	},
})
```

### AWS S3 Source

To use an S3 Bucket as a source in CodePipeline:

```go
sourceBucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	Versioned: jsii.Boolean(true),
})

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
	ActionName: jsii.String("S3Source"),
	Bucket: sourceBucket,
	BucketKey: jsii.String("path/to/file.zip"),
	Output: sourceOutput,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
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
sourceBucket := s3.Bucket_FromBucketAttributes(this, jsii.String("SourceBucket"), &BucketAttributes{
	BucketName: jsii.String("amzn-s3-demo-bucket"),
	Region: jsii.String("ap-southeast-1"),
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
import "github.com/aws/aws-cdk-go/awscdk"

var sourceBucket bucket

sourceOutput := codepipeline.NewArtifact()
key := "some/key.zip"
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
trail.AddS3EventSelector([]s3EventSelector{
	&s3EventSelector{
		Bucket: sourceBucket,
		ObjectPrefix: key,
	},
}, &AddEventSelectorOptions{
	ReadWriteType: cloudtrail.ReadWriteType_WRITE_ONLY,
})
sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
	ActionName: jsii.String("S3Source"),
	BucketKey: key,
	Bucket: sourceBucket,
	Output: sourceOutput,
	Trigger: codepipeline_actions.S3Trigger_EVENTS,
})
```

The S3 source action emits variables:

```go
var sourceBucket bucket

// later:
var project pipelineProject
key := "some/key.zip"
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
	ActionName: jsii.String("S3Source"),
	BucketKey: key,
	Bucket: sourceBucket,
	Output: sourceOutput,
	VariablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
		"VERSION_ID": &buildEnvironmentVariable{
			"value": sourceAction.variables.versionId,
		},
	},
})
```

### AWS ECR

To use an ECR Repository as a source in a Pipeline:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"

var ecrRepository repository

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewEcrSourceAction(&EcrSourceActionProps{
	ActionName: jsii.String("ECR"),
	Repository: ecrRepository,
	ImageTag: jsii.String("some-tag"),
	 // optional, default: 'latest'
	Output: sourceOutput,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
		sourceAction,
	},
})
```

The ECR source action emits variables:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"
var ecrRepository repository

// later:
var project pipelineProject


sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewEcrSourceAction(&EcrSourceActionProps{
	ActionName: jsii.String("Source"),
	Output: sourceOutput,
	Repository: ecrRepository,
	VariablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
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
var project pipelineProject

repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &RepositoryProps{
	RepositoryName: jsii.String("MyRepository"),
})
project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))

sourceOutput := codepipeline.NewArtifact()
sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("CodeCommit"),
	Repository: Repository,
	Output: sourceOutput,
})
buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	Outputs: []artifact{
		codepipeline.NewArtifact(),
	},
	 // optional
	ExecuteBatchBuild: jsii.Boolean(true),
	 // optional, defaults to false
	CombineBatchBuildArtifacts: jsii.Boolean(true),
})

codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
	Stages: []stageProps{
		&stageProps{
			StageName: jsii.String("Source"),
			Actions: []iAction{
				sourceAction,
			},
		},
		&stageProps{
			StageName: jsii.String("Build"),
			Actions: []*iAction{
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
var project pipelineProject

sourceOutput := codepipeline.NewArtifact()
testAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("IntegrationTest"),
	Project: Project,
	Input: sourceOutput,
	Type: codepipeline_actions.CodeBuildActionType_TEST,
})
```

#### Multiple inputs and outputs

When you want to have multiple inputs and/or outputs for a Project used in a
Pipeline, instead of using the `secondarySources` and `secondaryArtifacts`
properties of the `Project` class, you need to use the `extraInputs` and
`outputs` properties of the CodeBuild CodePipeline
Actions. Example:

```go
var repository1 repository
var repository2 repository

var project pipelineProject

sourceOutput1 := codepipeline.NewArtifact()
sourceAction1 := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("Source1"),
	Repository: repository1,
	Output: sourceOutput1,
})
sourceOutput2 := codepipeline.NewArtifact(jsii.String("source2"))
sourceAction2 := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("Source2"),
	Repository: repository2,
	Output: sourceOutput2,
})
buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("Build"),
	Project: Project,
	Input: sourceOutput1,
	ExtraInputs: []artifact{
		sourceOutput2,
	},
	Outputs: []*artifact{
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
project := codebuild.NewPipelineProject(this, jsii.String("MyProject"), &PipelineProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
// later:
var project pipelineProject
sourceOutput := codepipeline.NewArtifact()
buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("Build1"),
	Input: sourceOutput,
	Project: codebuild.NewPipelineProject(this, jsii.String("Project"), &PipelineProjectProps{
		BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
	VariablesNamespace: jsii.String("MyNamespace"),
})
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
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
jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &JenkinsProviderProps{
	ProviderName: jsii.String("MyJenkinsProvider"),
	ServerUrl: jsii.String("http://my-jenkins.com:8080"),
	Version: jsii.String("2"),
})
```

If you've registered a Jenkins provider in a different CDK app,
or outside the CDK (in the CodePipeline AWS Console, for example),
you can import it:

```go
jenkinsProvider := codepipeline_actions.JenkinsProvider_FromJenkinsProviderAttributes(this, jsii.String("JenkinsProvider"), &JenkinsProviderAttributes{
	ProviderName: jsii.String("MyJenkinsProvider"),
	ServerUrl: jsii.String("http://my-jenkins.com:8080"),
	Version: jsii.String("2"),
})
```

Note that a Jenkins provider
(identified by the provider name-category(build/test)-version tuple)
must always be registered in the given account, in the given AWS region,
before it can be used in CodePipeline.

With a `JenkinsProvider`,
we can create a Jenkins Action:

```go
var jenkinsProvider jenkinsProvider

buildAction := codepipeline_actions.NewJenkinsAction(&JenkinsActionProps{
	ActionName: jsii.String("JenkinsBuild"),
	JenkinsProvider: jenkinsProvider,
	ProjectName: jsii.String("MyProject"),
	Type: codepipeline_actions.JenkinsActionType_BUILD,
})
```

## Deploy

### AWS CloudFormation

This module contains Actions that allows you to deploy to CloudFormation from AWS CodePipeline.

For example, the following code fragment defines a pipeline that automatically deploys a CloudFormation template
directly from a CodeCommit repository, with a manual approval step in between to confirm the changes:

```go
// Source stage: read from repository
repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &RepositoryProps{
	RepositoryName: jsii.String("template-repo"),
})
sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
source := cpactions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	ActionName: jsii.String("Source"),
	Repository: repo,
	Output: sourceOutput,
	Trigger: cpactions.CodeCommitTrigger_POLL,
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

codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &PipelineProps{
	CrossAccountKeys: jsii.Boolean(true),
	Stages: []stageProps{
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
var pipeline pipeline
var sourceOutput artifact


pipeline.AddStage(&StageOptions{
	StageName: jsii.String("DeployStackSets"),
	Actions: []iAction{
		// First, update the StackSet itself with the newest template
		codepipeline_actions.NewCloudFormationDeployStackSetAction(&CloudFormationDeployStackSetActionProps{
			ActionName: jsii.String("UpdateStackSet"),
			RunOrder: jsii.Number(1),
			StackSetName: jsii.String("MyStackSet"),
			Template: codepipeline_actions.StackSetTemplate_FromArtifactPath(sourceOutput.AtPath(jsii.String("template.yaml"))),

			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
			DeploymentModel: codepipeline_actions.StackSetDeploymentModel_SelfManaged(),
			// This deploys to a set of accounts
			StackInstances: codepipeline_actions.StackInstances_InAccounts([]*string{
				jsii.String("111111111111"),
			}, []*string{
				jsii.String("us-east-1"),
				jsii.String("eu-west-1"),
			}),
		}),

		// Afterwards, update/create additional instances in other accounts
		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&CloudFormationDeployStackInstancesActionProps{
			ActionName: jsii.String("AddMoreInstances"),
			RunOrder: jsii.Number(2),
			StackSetName: jsii.String("MyStackSet"),
			StackInstances: codepipeline_actions.StackInstances_*InAccounts([]*string{
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
lambdaCode := lambda.Code_FromCfnParameters()
lambda.NewFunction(lambdaStack, jsii.String("Lambda"), &FunctionProps{
	Code: lambdaCode,
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
})
// other resources that your Lambda needs, added to the lambdaStack...

pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("Pipeline"), &PipelineProps{
	CrossAccountKeys: jsii.Boolean(true),
})

// add the source code repository containing this code to your Pipeline,
// and the source code of the Lambda Function, if they're separate
cdkSourceOutput := codepipeline.NewArtifact()
cdkSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	Repository: codecommit.NewRepository(pipelineStack, jsii.String("CdkCodeRepo"), &RepositoryProps{
		RepositoryName: jsii.String("CdkCodeRepo"),
	}),
	ActionName: jsii.String("CdkCode_Source"),
	Output: cdkSourceOutput,
})
lambdaSourceOutput := codepipeline.NewArtifact()
lambdaSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
	Repository: codecommit.NewRepository(pipelineStack, jsii.String("LambdaCodeRepo"), &RepositoryProps{
		RepositoryName: jsii.String("LambdaCodeRepo"),
	}),
	ActionName: jsii.String("LambdaCode_Source"),
	Output: lambdaSourceOutput,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
		cdkSourceAction,
		lambdaSourceAction,
	},
})

// synthesize the Lambda CDK template, using CodeBuild
// the below values are just examples, assuming your CDK code is in TypeScript/JavaScript -
// adjust the build environment and/or commands accordingly
cdkBuildProject := codebuild.NewProject(pipelineStack, jsii.String("CdkBuildProject"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
	},
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
cdkBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CDK_Build"),
	Project: cdkBuildProject,
	Input: cdkSourceOutput,
	Outputs: []artifact{
		cdkBuildOutput,
	},
})

// build your Lambda code, using CodeBuild
// again, this example assumes your Lambda is written in TypeScript/JavaScript -
// make sure to adjust the build environment and/or commands if they don't match your specific situation
lambdaBuildProject := codebuild.NewProject(pipelineStack, jsii.String("LambdaBuildProject"), &ProjectProps{
	Environment: &BuildEnvironment{
		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
	},
	BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
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
lambdaBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("Lambda_Build"),
	Project: lambdaBuildProject,
	Input: lambdaSourceOutput,
	Outputs: []*artifact{
		lambdaBuildOutput,
	},
})

pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Build"),
	Actions: []*iAction{
		cdkBuildAction,
		lambdaBuildAction,
	},
})

// finally, deploy your Lambda Stack
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []*iAction{
		codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
			ActionName: jsii.String("Lambda_CFN_Deploy"),
			TemplatePath: cdkBuildOutput.AtPath(jsii.String("LambdaStack.template.yaml")),
			StackName: jsii.String("LambdaStackDeployedName"),
			AdminPermissions: jsii.Boolean(true),
			ParameterOverrides: lambdaCode.Assign(lambdaBuildOutput.s3Location),
			ExtraInputs: []*artifact{
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
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
	ActionName: jsii.String("CloudFormationCreateUpdate"),
	StackName: jsii.String("MyStackName"),
	AdminPermissions: jsii.Boolean(true),
	TemplatePath: sourceOutput.AtPath(jsii.String("template.yaml")),
	Account: jsii.String("123456789012"),
})
```

This will create a new stack, called `<PipelineStackName>-support-123456789012`, in your `App`,
that will contain the role that the pipeline will assume in account 123456789012 before executing this action.
This support stack will automatically be deployed before the stack containing the pipeline.

You can also pass a role explicitly when creating the action -
in that case, the `account` property is ignored,
and the action will operate in the same account the role belongs to:

```go
import "github.com/aws/aws-cdk-go/awscdk"

// in stack for account 123456789012...
var otherAccountStack stack

actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
	// the role has to have a physical name set
	RoleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
})

// in the pipeline stack...
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
	ActionName: jsii.String("CloudFormationCreateUpdate"),
	StackName: jsii.String("MyStackName"),
	AdminPermissions: jsii.Boolean(true),
	TemplatePath: sourceOutput.AtPath(jsii.String("template.yaml")),
	Role: actionRole,
})
```

### AWS CodeDeploy

#### Server deployments

To use CodeDeploy for EC2/on-premise deployments in a Pipeline:

```go
var deploymentGroup serverDeploymentGroup
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
	PipelineName: jsii.String("MyPipeline"),
})

// add the source and build Stages to the Pipeline...
buildOutput := codepipeline.NewArtifact()
deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&CodeDeployServerDeployActionProps{
	ActionName: jsii.String("CodeDeploy"),
	Input: buildOutput,
	DeploymentGroup: DeploymentGroup,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []iAction{
		deployAction,
	},
})
```

##### Lambda deployments

To use CodeDeploy for blue-green Lambda deployments in a Pipeline:

```go
lambdaCode := lambda.Code_FromCfnParameters()
func := lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
	Code: lambdaCode,
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
})
// used to make sure each CDK synthesis produces a different Version
version := func.currentVersion
alias := lambda.NewAlias(this, jsii.String("LambdaAlias"), &AliasProps{
	AliasName: jsii.String("Prod"),
	Version: Version,
})

codedeploy.NewLambdaDeploymentGroup(this, jsii.String("DeploymentGroup"), &LambdaDeploymentGroupProps{
	Alias: Alias,
	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
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
import ecs "github.com/aws/aws-cdk-go/awscdk"

var service fargateService

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
buildOutput := codepipeline.NewArtifact()
deployStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []iAction{
		codepipeline_actions.NewEcsDeployAction(&EcsDeployActionProps{
			ActionName: jsii.String("DeployAction"),
			Service: *Service,
			// if your file is called imagedefinitions.json,
			// use the `input` property,
			// and leave out the `imageFile` property
			Input: buildOutput,
			// if your file name is _not_ imagedefinitions.json,
			// use the `imageFile` property,
			// and leave out the `input` property
			ImageFile: buildOutput.AtPath(jsii.String("imageDef.json")),
			DeploymentTimeout: awscdk.Duration_Minutes(jsii.Number(60)),
		}),
	},
})
```

#### Deploying ECS applications to existing services

CodePipeline can deploy to an existing ECS service which uses the
[ECS service ARN format that contains the Cluster name](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids).
This also works if the service is in a different account and/or region than the pipeline:

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"


service := ecs.BaseService_FromServiceArnWithCluster(this, jsii.String("EcsService"), jsii.String("arn:aws:ecs:us-east-1:123456789012:service/myClusterName/myServiceName"))
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
buildOutput := codepipeline.NewArtifact()
// add source and build stages to the pipeline as usual...
deployStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []iAction{
		codepipeline_actions.NewEcsDeployAction(&EcsDeployActionProps{
			ActionName: jsii.String("DeployAction"),
			Service: service,
			Input: buildOutput,
		}),
	},
})
```

When deploying across accounts, especially in a CDK Pipelines self-mutating pipeline,
it is recommended to provide the `role` property to the `EcsDeployAction`.
The Role will need to have permissions assigned to it for ECS deployment.
See [the CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/security-iam.html#how-to-custom-role)
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
 * These are the construction properties for `EcsAppStack`.
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

	taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDefinition"), &TaskDefinitionProps{
		Compatibility: ecs.Compatibility_FARGATE,
		Cpu: jsii.String("1024"),
		MemoryMiB: jsii.String("2048"),
	})
	taskDefinition.AddContainer(jsii.String("AppContainer"), &ContainerDefinitionOptions{
		Image: props.image,
	})
	ecs.NewFargateService(this, jsii.String("EcsService"), &FargateServiceProps{
		TaskDefinition: TaskDefinition,
		Cluster: ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
			Vpc: ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
				MaxAzs: jsii.Number(1),
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
	appCodeDockerBuild := codebuild.NewPipelineProject(this, jsii.String("AppCodeDockerImageBuildAndPushProject"), &PipelineProjectProps{
		Environment: &BuildEnvironment{
			// we need to run Docker
			Privileged: jsii.Boolean(true),
		},
		BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
		EnvironmentVariables: map[string]buildEnvironmentVariable{
			"REPOSITORY_URI": &buildEnvironmentVariable{
				"value": appEcrRepo.repositoryUri,
			},
		},
	})
	// needed for `docker push`
	appEcrRepo.GrantPullPush(appCodeDockerBuild)
	// create the ContainerImage used for the ECS application Stack
	this.tagParameterContainerImage = ecs.NewTagParameterContainerImage(appEcrRepo)

	cdkCodeBuild := codebuild.NewPipelineProject(this, jsii.String("CdkCodeBuildProject"), &PipelineProjectProps{
		BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
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
	appCodeBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
		ActionName: jsii.String("AppCodeDockerImageBuildAndPush"),
		Project: appCodeDockerBuild,
		Input: appCodeSourceOutput,
	})
	codepipeline.NewPipeline(this, jsii.String("CodePipelineDeployingEcsApplication"), &PipelineProps{
		ArtifactBucket: s3.NewBucket(this, jsii.String("ArtifactBucket"), &BucketProps{
			RemovalPolicy: cdk.RemovalPolicy_DESTROY,
		}),
		Stages: []stageProps{
			&stageProps{
				StageName: jsii.String("Source"),
				Actions: []iAction{
					// this is the Action that takes the source of your application code
					codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
						ActionName: jsii.String("AppCodeSource"),
						Repository: codecommit.NewRepository(this, jsii.String("AppCodeSourceRepository"), &RepositoryProps{
							RepositoryName: jsii.String("AppCodeSourceRepository"),
						}),
						Output: appCodeSourceOutput,
					}),
					// this is the Action that takes the source of your CDK code
					// (which would probably include this Pipeline code as well)
					codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
						ActionName: jsii.String("CdkCodeSource"),
						Repository: codecommit.NewRepository(this, jsii.String("CdkCodeSourceRepository"), &RepositoryProps{
							RepositoryName: jsii.String("CdkCodeSourceRepository"),
						}),
						Output: cdkCodeSourceOutput,
					}),
				},
			},
			&stageProps{
				StageName: jsii.String("Build"),
				Actions: []*iAction{
					appCodeBuildAction,
					codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
						ActionName: jsii.String("CdkCodeBuildAndSynth"),
						Project: cdkCodeBuild,
						Input: cdkCodeSourceOutput,
						Outputs: []artifact{
							cdkCodeBuildOutput,
						},
					}),
				},
			},
			&stageProps{
				StageName: jsii.String("Deploy"),
				Actions: []*iAction{
					codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
						ActionName: jsii.String("CFN_Deploy"),
						StackName: jsii.String("SampleEcsStackDeployedFromCodePipeline"),
						// this name has to be the same name as used below in the CDK code for the application Stack
						TemplatePath: cdkCodeBuildOutput.AtPath(jsii.String("EcsStackDeployedInPipeline.template.json")),
						AdminPermissions: jsii.Boolean(true),
						ParameterOverrides: map[string]interface{}{
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
import kms "github.com/aws/aws-cdk-go/awscdk"


sourceOutput := codepipeline.NewArtifact()
targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
key := kms.NewKey(this, jsii.String("EnvVarEncryptKey"), &KeyProps{
	Description: jsii.String("sample key"),
})

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
deployAction := codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
	ActionName: jsii.String("S3Deploy"),
	Bucket: targetBucket,
	Input: sourceOutput,
	EncryptionKey: key,
})
deployStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []iAction{
		deployAction,
	},
})
```

#### Invalidating the CloudFront cache when deploying to S3

There is currently no native support in CodePipeline for invalidating a CloudFront cache after deployment.
One workaround is to add another build step after the deploy step,
and use the AWS CLI to invalidate the cache:

```go
// Create a Cloudfront Web Distribution
import cloudfront "github.com/aws/aws-cdk-go/awscdk"
var distribution distribution


// Create the build project that will invalidate the cache
invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &PipelineProjectProps{
	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
		"version": jsii.String("0.2"),
		"phases": map[string]map[string][]*string{
			"build": map[string][]*string{
				"commands": []*string{
					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
				},
			},
		},
	}),
	EnvironmentVariables: map[string]buildEnvironmentVariable{
		"CLOUDFRONT_ID": &buildEnvironmentVariable{
			"value": distribution.distributionId,
		},
	},
})

// Add Cloudfront invalidation permissions to the project
distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.Account, distribution.DistributionId)
invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Resources: []*string{
		distributionArn,
	},
	Actions: []*string{
		jsii.String("cloudfront:CreateInvalidation"),
	},
}))

// Create the pipeline (here only the S3 deploy and Invalidate cache build)
deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
deployInput := codepipeline.NewArtifact()
codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	Stages: []stageProps{
		&stageProps{
			StageName: jsii.String("Deploy"),
			Actions: []iAction{
				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
					ActionName: jsii.String("S3Deploy"),
					Bucket: deployBucket,
					Input: deployInput,
					RunOrder: jsii.Number(1),
				}),
				codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
					ActionName: jsii.String("InvalidateCache"),
					Project: invalidateBuildProject,
					Input: deployInput,
					RunOrder: jsii.Number(2),
				}),
			},
		},
	},
})
```

### Elastic Beanstalk Deployment

To deploy an Elastic Beanstalk Application in CodePipeline:

```go
sourceOutput := codepipeline.NewArtifact()
targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
deployAction := codepipeline_actions.NewElasticBeanstalkDeployAction(&ElasticBeanstalkDeployActionProps{
	ActionName: jsii.String("ElasticBeanstalkDeploy"),
	Input: sourceOutput,
	EnvironmentName: jsii.String("envName"),
	ApplicationName: jsii.String("appName"),
})

deployStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Deploy"),
	Actions: []iAction{
		deployAction,
	},
})
```

### Alexa Skill

You can deploy to Alexa using CodePipeline with the following Action:

```go
// Read the secrets from ParameterStore
clientId := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientId"))
clientSecret := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientSecret"))
refreshToken := awscdk.SecretValue_SecretsManager(jsii.String("AlexaRefreshToken"))

// Add deploy action
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewAlexaSkillDeployAction(&AlexaSkillDeployActionProps{
	ActionName: jsii.String("DeploySkill"),
	RunOrder: jsii.Number(1),
	Input: sourceOutput,
	ClientId: clientId.ToString(),
	ClientSecret: clientSecret,
	RefreshToken: refreshToken,
	SkillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
})
```

If you need manifest overrides you can specify them as `parameterOverridesArtifact` in the action:

```go
// Deploy some CFN change set and store output
executeOutput := codepipeline.NewArtifact(jsii.String("CloudFormation"))
executeChangeSetAction := codepipeline_actions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
	ActionName: jsii.String("ExecuteChangesTest"),
	RunOrder: jsii.Number(2),
	StackName: jsii.String("MyStack"),
	ChangeSetName: jsii.String("MyChangeSet"),
	OutputFileName: jsii.String("overrides.json"),
	Output: executeOutput,
})

// Provide CFN output as manifest overrides
clientId := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientId"))
clientSecret := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientSecret"))
refreshToken := awscdk.SecretValue_SecretsManager(jsii.String("AlexaRefreshToken"))
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewAlexaSkillDeployAction(&AlexaSkillDeployActionProps{
	ActionName: jsii.String("DeploySkill"),
	RunOrder: jsii.Number(1),
	Input: sourceOutput,
	ParameterOverridesArtifact: executeOutput,
	ClientId: clientId.ToString(),
	ClientSecret: clientSecret,
	RefreshToken: refreshToken,
	SkillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
})
```

### AWS Service Catalog

You can deploy a CloudFormation template to an existing Service Catalog product with the following Action:

```go
cdkBuildOutput := codepipeline.NewArtifact()
serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&ServiceCatalogDeployActionBeta1Props{
	ActionName: jsii.String("ServiceCatalogDeploy"),
	TemplatePath: cdkBuildOutput.AtPath(jsii.String("Sample.template.json")),
	ProductVersionName: jsii.String("Version - " + date.now.toString),
	ProductVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
	ProductId: jsii.String("prod-XXXXXXXX"),
})
```

## Approve & invoke

### Manual approval Action

This package contains an Action that stops the Pipeline until someone manually clicks the approve button:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
approveStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Approve"),
})
manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&ManualApprovalActionProps{
	ActionName: jsii.String("Approve"),
	NotificationTopic: sns.NewTopic(this, jsii.String("Topic")),
	 // optional
	NotifyEmails: []*string{
		jsii.String("some_email@example.com"),
	},
	 // optional
	AdditionalInformation: jsii.String("additional info"),
})
approveStage.AddAction(manualApprovalAction)
```

If the `notificationTopic` has not been provided,
but `notifyEmails` were,
a new SNS Topic will be created
(and accessible through the `notificationTopic` property of the Action).

If you want to grant a principal permissions to approve the changes,
you can invoke the method `grantManualApproval` passing it a `IGrantable`:

```go
pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
approveStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Approve"),
})
manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&ManualApprovalActionProps{
	ActionName: jsii.String("Approve"),
})
approveStage.AddAction(manualApprovalAction)

role := iam.Role_FromRoleArn(this, jsii.String("Admin"), awscdk.Arn_Format(&ArnComponents{
	Service: jsii.String("iam"),
	Resource: jsii.String("role"),
	ResourceName: jsii.String("Admin"),
}, this))
manualApprovalAction.GrantManualApproval(role)
```

### AWS Lambda

This module contains an Action that allows you to invoke a Lambda function in a Pipeline:

```go
var fn function

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
	ActionName: jsii.String("Lambda"),
	Lambda: fn,
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Lambda"),
	Actions: []iAction{
		lambdaAction,
	},
})
```

The Lambda Action can have up to 5 inputs,
and up to 5 outputs:

```go
var fn function

sourceOutput := codepipeline.NewArtifact()
buildOutput := codepipeline.NewArtifact()
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
	ActionName: jsii.String("Lambda"),
	Inputs: []artifact{
		sourceOutput,
		buildOutput,
	},
	Outputs: []*artifact{
		codepipeline.NewArtifact(jsii.String("Out1")),
		codepipeline.NewArtifact(jsii.String("Out2")),
	},
	Lambda: fn,
})
```

The Lambda Action supports custom user parameters that pipeline
will pass to the Lambda function:

```go
var fn function


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
	ActionName: jsii.String("Lambda"),
	Lambda: fn,
	UserParameters: map[string]interface{}{
		"foo": jsii.String("bar"),
		"baz": jsii.String("qux"),
	},
	// OR
	UserParametersString: jsii.String("my-parameter-string"),
})
```

The Lambda invoke action emits variables.
Unlike many other actions, the variables are not static,
but dynamic, defined by the function calling the `PutJobSuccessResult`
API with the `outputVariables` property filled with the map of variables
Example:

```go
// later:
var project pipelineProject
lambdaInvokeAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
	ActionName: jsii.String("Lambda"),
	Lambda: lambda.NewFunction(this, jsii.String("Func"), &FunctionProps{
		Runtime: lambda.Runtime_NODEJS_LATEST(),
		Handler: jsii.String("index.handler"),
		Code: lambda.Code_FromInline(jsii.String(`
		        const { CodePipeline } = require('@aws-sdk/client-codepipeline');

		        exports.handler = async function(event, context) {
		            const codepipeline = new AWS.CodePipeline();
		            await codepipeline.putJobSuccessResult({
		                jobId: event['CodePipeline.job'].id,
		                outputVariables: {
		                    MY_VAR: "some value",
		                },
		            });
		        }
		    `)),
	}),
	VariablesNamespace: jsii.String("MyNamespace"),
})
sourceOutput := codepipeline.NewArtifact()
codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
	ActionName: jsii.String("CodeBuild"),
	Project: Project,
	Input: sourceOutput,
	EnvironmentVariables: map[string]buildEnvironmentVariable{
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
import "github.com/aws/aws-cdk-go/awscdk"

pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
startState := stepfunctions.NewPass(this, jsii.String("StartState"))
simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &StateMachineProps{
	Definition: startState,
})
stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&StepFunctionsInvokeActionProps{
	ActionName: jsii.String("Invoke"),
	StateMachine: simpleStateMachine,
	StateMachineInput: codepipeline_actions.StateMachineInput_Literal(map[string]*bool{
		"IsHelloWorldExample": jsii.Boolean(true),
	}),
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("StepFunctions"),
	Actions: []iAction{
		stepFunctionAction,
	},
})
```

The `StateMachineInput` can be created with one of 2 static factory methods:
`literal`, which takes an arbitrary map as its only argument, or `filePath`:

```go
import "github.com/aws/aws-cdk-go/awscdk"


pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
inputArtifact := codepipeline.NewArtifact()
startState := stepfunctions.NewPass(this, jsii.String("StartState"))
simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &StateMachineProps{
	Definition: startState,
})
stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&StepFunctionsInvokeActionProps{
	ActionName: jsii.String("Invoke"),
	StateMachine: simpleStateMachine,
	StateMachineInput: codepipeline_actions.StateMachineInput_FilePath(inputArtifact.AtPath(jsii.String("assets/input.json"))),
})
pipeline.AddStage(&StageOptions{
	StageName: jsii.String("StepFunctions"),
	Actions: []iAction{
		stepFunctionAction,
	},
})
```

See [the AWS documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-StepFunctions.html)
for information on Action structure reference.

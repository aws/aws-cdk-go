# AWS CodePipeline Construct Library

## Pipeline

To construct an empty Pipeline:

```go
// Construct an empty Pipeline
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"))
```

To give the Pipeline a nice, human-readable name:

```go
// Give the Pipeline a nice, human-readable name
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &PipelineProps{
	PipelineName: jsii.String("MyPipeline"),
})
```

Be aware that in the default configuration, the `Pipeline` construct creates
an AWS Key Management Service (AWS KMS) Customer Master Key (CMK) for you to
encrypt the artifacts in the artifact bucket, which incurs a cost of
**$1/month**. This default configuration is necessary to allow cross-account
actions.

If you do not intend to perform cross-account deployments, you can disable
the creation of the Customer Master Keys by passing `crossAccountKeys: false`
when defining the Pipeline:

```go
// Don't create Customer Master Keys
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &PipelineProps{
	CrossAccountKeys: jsii.Boolean(false),
})
```

If you want to enable key rotation for the generated KMS keys,
you can configure it by passing `enableKeyRotation: true` when creating the pipeline.
Note that key rotation will incur an additional cost of **$1/month**.

```go
// Enable key rotation for the generated KMS key
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &PipelineProps{
	// ...
	EnableKeyRotation: jsii.Boolean(true),
})
```

## Stages

You can provide Stages when creating the Pipeline:

```go
// Provide a Stage when creating a pipeline
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &PipelineProps{
	Stages: []stageProps{
		&stageProps{
			StageName: jsii.String("Source"),
			Actions: []iAction{
			},
		},
	},
})
```

Or append a Stage to an existing Pipeline:

```go
// Append a Stage to an existing Pipeline
var pipeline pipeline

sourceStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("Source"),
	Actions: []iAction{
	},
})
```

You can insert the new Stage at an arbitrary point in the Pipeline:

```go
// Insert a new Stage at an arbitrary point
var pipeline pipeline
var anotherStage iStage
var yetAnotherStage iStage


someStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("SomeStage"),
	Placement: &StagePlacement{
		// note: you can only specify one of the below properties
		RightBefore: anotherStage,
		JustAfter: yetAnotherStage,
	},
})
```

You can disable transition to a Stage:

```go
// Disable transition to a stage
var pipeline pipeline


someStage := pipeline.AddStage(&StageOptions{
	StageName: jsii.String("SomeStage"),
	TransitionToEnabled: jsii.Boolean(false),
	TransitionDisabledReason: jsii.String("Manual transition only"),
})
```

This is useful if you don't want every executions of the pipeline to flow into
this stage automatically. The transition can then be "manually" enabled later on.

## Actions

Actions live in a separate package, `aws-cdk-lib/aws-codepipeline-actions`.

To add an Action to a Stage, you can provide it when creating the Stage,
in the `actions` property,
or you can use the `IStage.addAction()` method to mutate an existing Stage:

```go
// Use the `IStage.addAction()` method to mutate an existing Stage.
var sourceStage iStage
var someAction action

sourceStage.AddAction(someAction)
```

## Custom Action Registration

To make your own custom CodePipeline Action requires registering the action provider. Look to the `JenkinsProvider` in `aws-cdk-lib/aws-codepipeline-actions` for an implementation example.

```go
// Make a custom CodePipeline Action
// Make a custom CodePipeline Action
codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &CustomActionRegistrationProps{
	Category: codepipeline.ActionCategory_SOURCE,
	ArtifactBounds: &ActionArtifactBounds{
		MinInputs: jsii.Number(0),
		MaxInputs: jsii.Number(0),
		MinOutputs: jsii.Number(1),
		MaxOutputs: jsii.Number(1),
	},
	Provider: jsii.String("GenericGitSource"),
	Version: jsii.String("1"),
	EntityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
	ExecutionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
	ActionProperties: []customActionProperty{
		&customActionProperty{
			Name: jsii.String("Branch"),
			Required: jsii.Boolean(true),
			Key: jsii.Boolean(false),
			Secret: jsii.Boolean(false),
			Queryable: jsii.Boolean(false),
			Description: jsii.String("Git branch to pull"),
			Type: jsii.String("String"),
		},
		&customActionProperty{
			Name: jsii.String("GitUrl"),
			Required: jsii.Boolean(true),
			Key: jsii.Boolean(false),
			Secret: jsii.Boolean(false),
			Queryable: jsii.Boolean(false),
			Description: jsii.String("SSH git clone URL"),
			Type: jsii.String("String"),
		},
	},
})
```

## Cross-account CodePipelines

> Cross-account Pipeline actions require that the Pipeline has *not* been
> created with `crossAccountKeys: false`.

Most pipeline Actions accept an AWS resource object to operate on. For example:

* `S3DeployAction` accepts an `s3.IBucket`.
* `CodeBuildAction` accepts a `codebuild.IProject`.
* etc.

These resources can be either newly defined (`new s3.Bucket(...)`) or imported
(`s3.Bucket.fromBucketAttributes(...)`) and identify the resource that should
be changed.

These resources can be in different accounts than the pipeline itself. For
example, the following action deploys to an imported S3 bucket from a
different account:

```go
// Deploy an imported S3 bucket from a different account
var stage iStage
var input artifact

stage.AddAction(codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
	Bucket: s3.Bucket_FromBucketAttributes(this, jsii.String("Bucket"), &BucketAttributes{
		Account: jsii.String("123456789012"),
	}),
	Input: input,
	ActionName: jsii.String("s3-deploy-action"),
}))
```

Actions that don't accept a resource object accept an explicit `account` parameter:

```go
// Actions that don't accept a resource objet accept an explicit `account` parameter
var stage iStage
var templatePath artifactPath

stage.AddAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
	Account: jsii.String("123456789012"),
	TemplatePath: TemplatePath,
	AdminPermissions: jsii.Boolean(false),
	StackName: awscdk.*stack_Of(this).stackName,
	ActionName: jsii.String("cloudformation-create-update"),
}))
```

The `Pipeline` construct automatically defines an **IAM Role** for you in the
target account which the pipeline will assume to perform that action. This
Role will be defined in a **support stack** named
`<PipelineStackName>-support-<account>`, that will automatically be deployed
before the stack containing the pipeline.

If you do not want to use the generated role, you can also explicitly pass a
`role` when creating the action. In that case, the action will operate in the
account the role belongs to:

```go
// Explicitly pass in a `role` when creating an action.
var stage iStage
var templatePath artifactPath

stage.AddAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
	TemplatePath: TemplatePath,
	AdminPermissions: jsii.Boolean(false),
	StackName: awscdk.*stack_Of(this).stackName,
	ActionName: jsii.String("cloudformation-create-update"),
	// ...
	Role: iam.Role_FromRoleArn(this, jsii.String("ActionRole"), jsii.String("...")),
}))
```

## Cross-region CodePipelines

Similar to how you set up a cross-account Action, the AWS resource object you
pass to actions can also be in different *Regions*. For example, the
following Action deploys to an imported S3 bucket from a different Region:

```go
// Deploy to an imported S3 bucket from a different Region.
var stage iStage
var input artifact

stage.AddAction(codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
	Bucket: s3.Bucket_FromBucketAttributes(this, jsii.String("Bucket"), &BucketAttributes{
		Region: jsii.String("us-west-1"),
	}),
	Input: input,
	ActionName: jsii.String("s3-deploy-action"),
}))
```

Actions that don't take an AWS resource will accept an explicit `region`
parameter:

```go
// Actions that don't take an AWS resource will accept an explicit `region` parameter.
var stage iStage
var templatePath artifactPath

stage.AddAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
	TemplatePath: TemplatePath,
	AdminPermissions: jsii.Boolean(false),
	StackName: awscdk.*stack_Of(this).stackName,
	ActionName: jsii.String("cloudformation-create-update"),
	// ...
	Region: jsii.String("us-west-1"),
}))
```

The `Pipeline` construct automatically defines a **replication bucket** for
you in the target region, which the pipeline will replicate artifacts to and
from. This Bucket will be defined in a **support stack** named
`<PipelineStackName>-support-<region>`, that will automatically be deployed
before the stack containing the pipeline.

If you don't want to use these support stacks, and already have buckets in
place to serve as replication buckets, you can supply these at Pipeline definition
time using the `crossRegionReplicationBuckets` parameter. Example:

```go
// Supply replication buckets for the Pipeline instead of using the generated support stack
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &PipelineProps{
	// ...

	CrossRegionReplicationBuckets: map[string]iBucket{
		// note that a physical name of the replication Bucket must be known at synthesis time
		"us-west-1": s3.Bucket_fromBucketAttributes(this, jsii.String("UsWest1ReplicationBucket"), &BucketAttributes{
			"bucketName": jsii.String("my-us-west-1-replication-bucket"),
			// optional KMS key
			"encryptionKey": kms.Key_fromKeyArn(this, jsii.String("UsWest1ReplicationKey"), jsii.String("arn:aws:kms:us-west-1:123456789012:key/1234-5678-9012")),
		}),
	},
})
```

See [the AWS docs here](https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-cross-region.html)
for more information on cross-region CodePipelines.

### Creating an encrypted replication bucket

If you're passing a replication bucket created in a different stack,
like this:

```go
// Passing a replication bucket created in a different stack.
app := awscdk.NewApp()
replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-1"),
	},
})
key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &BucketProps{
	// like was said above - replication buckets need a set physical name
	BucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	EncryptionKey: key,
})

// later...
// later...
codepipeline.NewPipeline(replicationStack, jsii.String("Pipeline"), &PipelineProps{
	CrossRegionReplicationBuckets: map[string]iBucket{
		"us-west-1": replicationBucket,
	},
})
```

When trying to encrypt it
(and note that if any of the cross-region actions happen to be cross-account as well,
the bucket *has to* be encrypted - otherwise the pipeline will fail at runtime),
you cannot use a key directly - KMS keys don't have physical names,
and so you can't reference them across environments.

In this case, you need to use an alias in place of the key when creating the bucket:

```go
// Passing an encrypted replication bucket created in a different stack.
app := awscdk.NewApp()
replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-1"),
	},
})
key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
alias := kms.NewAlias(replicationStack, jsii.String("ReplicationAlias"), &AliasProps{
	// aliasName is required
	AliasName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	TargetKey: key,
})
replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &BucketProps{
	BucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	EncryptionKey: alias,
})
```

## Variables

Variables are key-value pairs that can be used to dynamically configure actions in your pipeline.

There are two types of variables, Action-level variables and Pipeline-level variables. Action-level
variables are produced when an action is executed. Pipeline-level variables are defined when the
pipeline is created and resolved at pipeline run time. You specify the Pipeline-level variables
when the pipeline is created, and you can provide values at the time of the pipeline execution.

### Action-level variables

The library supports action-level variables.
Each action class that emits variables has a separate variables interface,
accessed as a property of the action instance called `variables`.
You instantiate the action class and assign it to a local variable;
when you want to use a variable in the configuration of a different action,
you access the appropriate property of the interface returned from `variables`,
which represents a single variable.
Example:

```go
// MyAction is some action type that produces variables, like EcrSourceAction
myAction := NewMyAction(&myActionProps{
	// ...
	actionName: jsii.String("myAction"),
})
NewOtherAction(&otherActionProps{
	// ...
	config: myAction.variables.myVariable,
	actionName: jsii.String("otherAction"),
})
```

The namespace name that will be used will be automatically generated by the pipeline construct,
based on the stage and action name;
you can pass a custom name when creating the action instance:

```go
// MyAction is some action type that produces variables, like EcrSourceAction
myAction := NewMyAction(&myActionProps{
	// ...
	variablesNamespace: jsii.String("MyNamespace"),
	actionName: jsii.String("myAction"),
})
```

There are also global variables available,
not tied to any action;
these are accessed through static properties of the `GlobalVariables` class:

```go
// OtherAction is some action type that produces variables, like EcrSourceAction
// OtherAction is some action type that produces variables, like EcrSourceAction
NewOtherAction(&otherActionProps{
	// ...
	config: codepipeline.GlobalVariables_ExecutionId(),
	actionName: jsii.String("otherAction"),
})
```

The following is an actual code example.

```go
var sourceAction s3SourceAction
var sourceOutput artifact
var deployBucket bucket


codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	Stages: []stageProps{
		&stageProps{
			StageName: jsii.String("Source"),
			Actions: []iAction{
				sourceAction,
			},
		},
		&stageProps{
			StageName: jsii.String("Deploy"),
			Actions: []*iAction{
				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
					ActionName: jsii.String("DeployAction"),
					// can reference the variables
					ObjectKey: fmt.Sprintf("%v.txt", sourceAction.variables.VersionId),
					Input: sourceOutput,
					Bucket: deployBucket,
				}),
			},
		},
	},
})
```

Check the documentation of the `aws-cdk-lib/aws-codepipeline-actions`
for details on how to use the variables for each action class.

See the [CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-variables.html)
for more details on how to use the variables feature.

### Pipeline-level variables

You can add one or more variables at the pipeline level. You can reference
this value in the configuration of CodePipeline actions. You can add the
variable names, default values, and descriptions when you create the pipeline.
Variables are resolved at the time of execution.

Note that using pipeline-level variables in any kind of Source action is not supported.
Also, the variables can only be used with pipeline type V2.

```go
var sourceAction s3SourceAction
var sourceOutput artifact
var deployBucket bucket


// Pipeline-level variable
variable := codepipeline.NewVariable(&VariableProps{
	VariableName: jsii.String("bucket-var"),
	Description: jsii.String("description"),
	DefaultValue: jsii.String("sample"),
})

codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	PipelineType: codepipeline.PipelineType_V2,
	Variables: []variable{
		variable,
	},
	Stages: []stageProps{
		&stageProps{
			StageName: jsii.String("Source"),
			Actions: []iAction{
				sourceAction,
			},
		},
		&stageProps{
			StageName: jsii.String("Deploy"),
			Actions: []*iAction{
				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
					ActionName: jsii.String("DeployAction"),
					// can reference the variables
					ObjectKey: fmt.Sprintf("%v.txt", variable.Reference()),
					Input: sourceOutput,
					Bucket: deployBucket,
				}),
			},
		},
	},
})
```

Or append a variable to an existing pipeline:

```go
var pipeline pipeline


variable := codepipeline.NewVariable(&VariableProps{
	VariableName: jsii.String("bucket-var"),
	Description: jsii.String("description"),
	DefaultValue: jsii.String("sample"),
})
pipeline.AddVariable(variable)
```

## Events

### Using a pipeline as an event target

A pipeline can be used as a target for a CloudWatch event rule:

```go
// A pipeline being used as a target for a CloudWatch event rule.
import targets "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline


// kick off the pipeline every day
rule := events.NewRule(this, jsii.String("Daily"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(1))),
})
rule.AddTarget(targets.NewCodePipeline(pipeline))
```

When a pipeline is used as an event target, the
"codepipeline:StartPipelineExecution" permission is granted to the AWS
CloudWatch Events service.

### Event sources

Pipelines emit CloudWatch events. To define event rules for events emitted by
the pipeline, stages or action, use the `onXxx` methods on the respective
construct:

```go
// Define event rules for events emitted by the pipeline
import events "github.com/aws/aws-cdk-go/awscdk"

var myPipeline pipeline
var myStage iStage
var myAction action
var target iRuleTarget

myPipeline.onStateChange(jsii.String("MyPipelineStateChange"), &OnEventOptions{
	Target: target,
})
myStage.OnStateChange(jsii.String("MyStageStateChange"), target)
myAction.OnStateChange(jsii.String("MyActionStateChange"), target)
```

## CodeStar Notifications

To define CodeStar Notification rules for Pipelines, use one of the `notifyOnXxx()` methods.
They are very similar to `onXxx()` methods for CloudWatch events:

```go
// Define CodeStar Notification rules for Pipelines
import chatbot "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline

target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
})
rule := pipeline.notifyOnExecutionStateChange(jsii.String("NotifyOnExecutionStateChange"), target)
```

## Trigger

To trigger a pipeline with Git tags, specify the `triggers` property. When a Git tag is pushed,
your pipeline starts. You can filter with glob patterns. The `tagsExcludes` takes priority over
the `tagsIncludes`.

The triggers can only be used with pipeline type V2.

```go
var sourceAction codeStarConnectionsSourceAction
var buildAction codeBuildAction


codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	PipelineType: codepipeline.PipelineType_V2,
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
	Triggers: []triggerProps{
		&triggerProps{
			ProviderType: codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
			GitConfiguration: &GitConfiguration{
				SourceAction: *SourceAction,
				PushFilter: []gitPushFilter{
					&gitPushFilter{
						TagsExcludes: []*string{
							jsii.String("exclude1"),
							jsii.String("exclude2"),
						},
						TagsIncludes: []*string{
							jsii.String("include*"),
						},
					},
				},
			},
		},
	},
})
```

Or append a trigger to an existing pipeline:

```go
var pipeline pipeline
var sourceAction codeStarConnectionsSourceAction


pipeline.AddTrigger(&TriggerProps{
	ProviderType: codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
	GitConfiguration: &GitConfiguration{
		SourceAction: *SourceAction,
		PushFilter: []gitPushFilter{
			&gitPushFilter{
				TagsExcludes: []*string{
					jsii.String("exclude1"),
					jsii.String("exclude2"),
				},
				TagsIncludes: []*string{
					jsii.String("include*"),
				},
			},
		},
	},
})
```

## Execution mode

To control the concurrency behavior when multiple executions of a pipeline are started, you can use the `executionMode` property.

The execution mode can only be used with pipeline type V2.

```go
codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	PipelineType: codepipeline.PipelineType_V2,
	ExecutionMode: codepipeline.ExecutionMode_PARALLEL,
})
```

## Migrating a pipeline type from V1 to V2

To migrate your pipeline type from V1 to V2, you just need to update the `pipelineType` property to `PipelineType.V2`.
This migration does not cause replacement of your pipeline.

When the `@aws-cdk/aws-codepipeline:defaultPipelineTypeToV2` feature flag is set to `true` (default for new projects),
the V2 type is selected by default if you do not specify a value for `pipelineType` property. Otherwise, the V1 type is selected.

```go
codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	PipelineType: codepipeline.PipelineType_V2,
})
```

See the [CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipeline-types-planning.html)
for more details on the differences between each type.

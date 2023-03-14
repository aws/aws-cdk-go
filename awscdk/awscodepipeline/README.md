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
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &pipelineProps{
	pipelineName: jsii.String("MyPipeline"),
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
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &pipelineProps{
	crossAccountKeys: jsii.Boolean(false),
})
```

If you want to enable key rotation for the generated KMS keys,
you can configure it by passing `enableKeyRotation: true` when creating the pipeline.
Note that key rotation will incur an additional cost of **$1/month**.

```go
// Enable key rotation for the generated KMS key
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &pipelineProps{
	// ...
	enableKeyRotation: jsii.Boolean(true),
})
```

## Stages

You can provide Stages when creating the Pipeline:

```go
// Provide a Stage when creating a pipeline
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &pipelineProps{
	stages: []stageProps{
		&stageProps{
			stageName: jsii.String("Source"),
			actions: []iAction{
			},
		},
	},
})
```

Or append a Stage to an existing Pipeline:

```go
// Append a Stage to an existing Pipeline
var pipeline pipeline

sourceStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Source"),
	actions: []iAction{
	},
})
```

You can insert the new Stage at an arbitrary point in the Pipeline:

```go
// Insert a new Stage at an arbitrary point
var pipeline pipeline
var anotherStage iStage
var yetAnotherStage iStage


someStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("SomeStage"),
	placement: &stagePlacement{
		// note: you can only specify one of the below properties
		rightBefore: anotherStage,
		justAfter: yetAnotherStage,
	},
})
```

You can disable transition to a Stage:

```go
// Disable transition to a stage
var pipeline pipeline


someStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("SomeStage"),
	transitionToEnabled: jsii.Boolean(false),
	transitionDisabledReason: jsii.String("Manual transition only"),
})
```

This is useful if you don't want every executions of the pipeline to flow into
this stage automatically. The transition can then be "manually" enabled later on.

## Actions

Actions live in a separate package, `@aws-cdk/aws-codepipeline-actions`.

To add an Action to a Stage, you can provide it when creating the Stage,
in the `actions` property,
or you can use the `IStage.addAction()` method to mutate an existing Stage:

```go
// Use the `IStage.addAction()` method to mutate an existing Stage.
var sourceStage iStage
var someAction action

sourceStage.addAction(someAction)
```

## Custom Action Registration

To make your own custom CodePipeline Action requires registering the action provider. Look to the `JenkinsProvider` in `@aws-cdk/aws-codepipeline-actions` for an implementation example.

```go
// Make a custom CodePipeline Action
// Make a custom CodePipeline Action
codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &customActionRegistrationProps{
	category: codepipeline.actionCategory_SOURCE,
	artifactBounds: &actionArtifactBounds{
		minInputs: jsii.Number(0),
		maxInputs: jsii.Number(0),
		minOutputs: jsii.Number(1),
		maxOutputs: jsii.Number(1),
	},
	provider: jsii.String("GenericGitSource"),
	version: jsii.String("1"),
	entityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
	executionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
	actionProperties: []customActionProperty{
		&customActionProperty{
			name: jsii.String("Branch"),
			required: jsii.Boolean(true),
			key: jsii.Boolean(false),
			secret: jsii.Boolean(false),
			queryable: jsii.Boolean(false),
			description: jsii.String("Git branch to pull"),
			type: jsii.String("String"),
		},
		&customActionProperty{
			name: jsii.String("GitUrl"),
			required: jsii.Boolean(true),
			key: jsii.Boolean(false),
			secret: jsii.Boolean(false),
			queryable: jsii.Boolean(false),
			description: jsii.String("SSH git clone URL"),
			type: jsii.String("String"),
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

stage.addAction(codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
	bucket: s3.bucket.fromBucketAttributes(this, jsii.String("Bucket"), &bucketAttributes{
		account: jsii.String("123456789012"),
	}),
	input: input,
	actionName: jsii.String("s3-deploy-action"),
}))
```

Actions that don't accept a resource object accept an explicit `account` parameter:

```go
// Actions that don't accept a resource objet accept an explicit `account` parameter
var stage iStage
var templatePath artifactPath

stage.addAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
	account: jsii.String("123456789012"),
	templatePath: templatePath,
	adminPermissions: jsii.Boolean(false),
	stackName: awscdk.*stack.of(this).stackName,
	actionName: jsii.String("cloudformation-create-update"),
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

stage.addAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
	templatePath: templatePath,
	adminPermissions: jsii.Boolean(false),
	stackName: awscdk.*stack.of(this).stackName,
	actionName: jsii.String("cloudformation-create-update"),
	// ...
	role: iam.role.fromRoleArn(this, jsii.String("ActionRole"), jsii.String("...")),
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

stage.addAction(codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
	bucket: s3.bucket.fromBucketAttributes(this, jsii.String("Bucket"), &bucketAttributes{
		region: jsii.String("us-west-1"),
	}),
	input: input,
	actionName: jsii.String("s3-deploy-action"),
}))
```

Actions that don't take an AWS resource will accept an explicit `region`
parameter:

```go
// Actions that don't take an AWS resource will accept an explicit `region` parameter.
var stage iStage
var templatePath artifactPath

stage.addAction(codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
	templatePath: templatePath,
	adminPermissions: jsii.Boolean(false),
	stackName: awscdk.*stack.of(this).stackName,
	actionName: jsii.String("cloudformation-create-update"),
	// ...
	region: jsii.String("us-west-1"),
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
pipeline := codepipeline.NewPipeline(this, jsii.String("MyFirstPipeline"), &pipelineProps{
	// ...

	crossRegionReplicationBuckets: map[string]iBucket{
		// note that a physical name of the replication Bucket must be known at synthesis time
		"us-west-1": s3.Bucket.fromBucketAttributes(this, jsii.String("UsWest1ReplicationBucket"), &BucketAttributes{
			"bucketName": jsii.String("my-us-west-1-replication-bucket"),
			// optional KMS key
			"encryptionKey": kms.Key.fromKeyArn(this, jsii.String("UsWest1ReplicationKey"), jsii.String("arn:aws:kms:us-west-1:123456789012:key/1234-5678-9012")),
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
replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &stackProps{
	env: &environment{
		region: jsii.String("us-west-1"),
	},
})
key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &bucketProps{
	// like was said above - replication buckets need a set physical name
	bucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	encryptionKey: key,
})

// later...
// later...
codepipeline.NewPipeline(replicationStack, jsii.String("Pipeline"), &pipelineProps{
	crossRegionReplicationBuckets: map[string]iBucket{
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
replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &stackProps{
	env: &environment{
		region: jsii.String("us-west-1"),
	},
})
key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
alias := kms.NewAlias(replicationStack, jsii.String("ReplicationAlias"), &aliasProps{
	// aliasName is required
	aliasName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	targetKey: key,
})
replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &bucketProps{
	bucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
	encryptionKey: alias,
})
```

## Variables

The library supports the CodePipeline Variables feature.
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
	config: codepipeline.globalVariables_ExecutionId(),
	actionName: jsii.String("otherAction"),
})
```

Check the documentation of the `@aws-cdk/aws-codepipeline-actions`
for details on how to use the variables for each action class.

See the [CodePipeline documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-variables.html)
for more details on how to use the variables feature.

## Events

### Using a pipeline as an event target

A pipeline can be used as a target for a CloudWatch event rule:

```go
// A pipeline being used as a target for a CloudWatch event rule.
import targets "github.com/aws/aws-cdk-go/awscdk"
import events "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline


// kick off the pipeline every day
rule := events.NewRule(this, jsii.String("Daily"), &ruleProps{
	schedule: events.schedule.rate(awscdk.Duration.days(jsii.Number(1))),
})
rule.addTarget(targets.NewCodePipeline(pipeline))
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

myPipeline.onStateChange(jsii.String("MyPipelineStateChange"), &onEventOptions{
	target: target,
})
myStage.onStateChange(jsii.String("MyStageStateChange"), target)
myAction.onStateChange(jsii.String("MyActionStateChange"), target)
```

## CodeStar Notifications

To define CodeStar Notification rules for Pipelines, use one of the `notifyOnXxx()` methods.
They are very similar to `onXxx()` methods for CloudWatch events:

```go
// Define CodeStar Notification rules for Pipelines
import chatbot "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline

target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
})
rule := pipeline.notifyOnExecutionStateChange(jsii.String("NotifyOnExecutionStateChange"), target)
```

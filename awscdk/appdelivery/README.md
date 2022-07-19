# Continuous Integration / Continuous Delivery for CDK Applications

This library includes a *CodePipeline* composite Action for deploying AWS CDK Applications.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Replacement recommended

This library has been deprecated. We recommend you use the
[@aws-cdk/pipelines](https://docs.aws.amazon.com/cdk/api/latest/docs/pipelines-readme.html) module instead.

## Limitations

The construct library in it's current form has the following limitations:

1. It can only deploy stacks that are hosted in the same AWS account and region as the *CodePipeline* is.
2. Stacks that make use of `Asset`s cannot be deployed successfully.

## Getting Started

In order to add the `PipelineDeployStackAction` to your *CodePipeline*, you need to have a *CodePipeline* artifact that
contains the result of invoking `cdk synth -o <dir>` on your *CDK App*. You can for example achieve this using a
*CodeBuild* project.

The example below defines a *CDK App* that contains 3 stacks:

* `CodePipelineStack` manages the *CodePipeline* resources, and self-updates before deploying any other stack
* `ServiceStackA` and `ServiceStackB` are service infrastructure stacks, and need to be deployed in this order

```plaintext
  ┏━━━━━━━━━━━━━━━━┓  ┏━━━━━━━━━━━━━━━━┓  ┏━━━━━━━━━━━━━━━━━┓  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
  ┃     Source     ┃  ┃     Build      ┃  ┃  Self-Update    ┃  ┃             Deploy              ┃
  ┃                ┃  ┃                ┃  ┃                 ┃  ┃                                 ┃
  ┃ ┌────────────┐ ┃  ┃ ┌────────────┐ ┃  ┃ ┌─────────────┐ ┃  ┃ ┌─────────────┐ ┌─────────────┐ ┃
  ┃ │   GitHub   ┣━╋━━╋━▶ CodeBuild  ┣━╋━━╋━▶Deploy Stack ┣━╋━━╋━▶Deploy Stack ┣━▶Deploy Stack │ ┃
  ┃ │            │ ┃  ┃ │            │ ┃  ┃ │PipelineStack│ ┃  ┃ │ServiceStackA│ │ServiceStackB│ ┃
  ┃ └────────────┘ ┃  ┃ └────────────┘ ┃  ┃ └─────────────┘ ┃  ┃ └─────────────┘ └─────────────┘ ┃
  ┗━━━━━━━━━━━━━━━━┛  ┗━━━━━━━━━━━━━━━━┛  ┗━━━━━━━━━━━━━━━━━┛  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### `index.ts`

```go
// Example automatically generated from non-compiling source. May contain errors.
import codebuild "github.com/aws/aws-cdk-go/awscdk"
import codepipeline "github.com/aws/aws-cdk-go/awscdk"
import codepipeline_actions "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
import cicd "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"


type myServiceStackA struct {
	stack
}
type myServiceStackB struct {
	stack
}

app := cdk.NewApp()

// We define a stack that contains the CodePipeline
pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("CodePipeline"), &pipelineProps{
	// Mutating a CodePipeline can cause the currently propagating state to be
	// "lost". Ensure we re-run the latest change through the pipeline after it's
	// been mutated so we're sure the latest state is fully deployed through.
	restartExecutionOnUpdate: jsii.Boolean(true),
})

// Configure the CodePipeline source - where your CDK App's source code is hosted
sourceOutput := codepipeline.NewArtifact()
source := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
	actionName: jsii.String("GitHub"),
	output: sourceOutput,
	owner: jsii.String("myName"),
	repo: jsii.String("myRepo"),
	oauthToken: cdk.secretValue.unsafePlainText(jsii.String("secret")),
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("source"),
	actions: []iAction{
		source,
	},
})

project := codebuild.NewPipelineProject(pipelineStack, jsii.String("CodeBuild"), &pipelineProjectProps{
})
synthesizedApp := codepipeline.NewArtifact()
buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
	actionName: jsii.String("CodeBuild"),
	project: project,
	input: sourceOutput,
	outputs: []artifact{
		synthesizedApp,
	},
})
pipeline.addStage(&stageOptions{
	stageName: jsii.String("build"),
	actions: []*iAction{
		buildAction,
	},
})

// Optionally, self-update the pipeline stack
selfUpdateStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("SelfUpdate"),
})
selfUpdateStage.addAction(cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
	stack: pipelineStack,
	input: synthesizedApp,
	adminPermissions: jsii.Boolean(true),
}))

// Now add our service stacks
deployStage := pipeline.addStage(&stageOptions{
	stageName: jsii.String("Deploy"),
})
serviceStackA := NewMyServiceStackA(app, jsii.String("ServiceStackA"), &stackProps{
})
// Add actions to deploy the stacks in the deploy stage:
deployServiceAAction := cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
	stack: serviceStackA,
	input: synthesizedApp,
	// See the note below for details about this option.
	adminPermissions: jsii.Boolean(false),
})
deployStage.addAction(deployServiceAAction)
// Add the necessary permissions for you service deploy action. This role is
// is passed to CloudFormation and needs the permissions necessary to deploy
// stack. Alternatively you can enable [Administrator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_job-functions.html#jf_administrator) permissions above,
// users should understand the privileged nature of this role.
myResourceArn := "arn:partition:service:region:account-id:resource-id"
deployServiceAAction.addToDeploymentRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("service:SomeAction"),
	},
	resources: []*string{
		myResourceArn,
	},
}))

serviceStackB := NewMyServiceStackB(app, jsii.String("ServiceStackB"), &stackProps{
})
deployStage.addAction(cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
	stack: serviceStackB,
	input: synthesizedApp,
	createChangeSetRunOrder: jsii.Number(998),
	adminPermissions: jsii.Boolean(true),
}))
```

### `buildspec.yml`

The repository can contain a file at the root level named `buildspec.yml`, or
you can in-line the buildspec. Note that `buildspec.yaml` is not compatible.

For example, a *TypeScript* or *Javascript* CDK App can add the following `buildspec.yml`
at the root of the repository:

```yml
version: 0.2
phases:
  install:
    commands:
      # Installs the npm dependencies as defined by the `package.json` file
      # present in the root directory of the package
      # (`cdk init app --language=typescript` would have created one for you)
      - npm install
  build:
    commands:
      # Builds the CDK App so it can be synthesized
      - npm run build
      # Synthesizes the CDK App and puts the resulting artifacts into `dist`
      - npm run cdk synth -- -o dist
artifacts:
  # The output artifact is all the files in the `dist` directory
  base-directory: dist
  files: '**/*'
```

The `PipelineDeployStackAction` expects it's `input` to contain the result of
synthesizing a CDK App using the `cdk synth -o <directory>`.

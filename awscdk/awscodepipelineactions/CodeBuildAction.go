package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline build action that uses AWS CodeBuild.
//
// Example:
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution Distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &PipelineProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	EnvironmentVariables: map[string]BuildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &BuildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.Account, distribution.DistributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Resources: []*string{
//   		distributionArn,
//   	},
//   	Actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	Stages: []StageProps{
//   		&StageProps{
//   			StageName: jsii.String("Deploy"),
//   			Actions: []IAction{
//   				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   					ActionName: jsii.String("S3Deploy"),
//   					Bucket: deployBucket,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   					ActionName: jsii.String("InvalidateCache"),
//   					Project: invalidateBuildProject,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
type CodeBuildAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Reference a CodePipeline variable defined by the CodeBuild project this action points to.
	//
	// Variables in CodeBuild actions are defined using the 'exported-variables' subsection of the 'env'
	// section of the buildspec.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-syntax
	//
	Variable(variableName *string) *string
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CodeBuildAction
type jsiiProxy_CodeBuildAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeBuildAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeBuildAction(props *CodeBuildActionProps) CodeBuildAction {
	_init_.Initialize()

	if err := validateNewCodeBuildActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeBuildAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeBuildAction_Override(c CodeBuildAction, props *CodeBuildActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeBuildAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) Bound(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBoundParameters(scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, _stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := c.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) Variable(variableName *string) *string {
	if err := c.validateVariableParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"variable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildAction) VariableExpression(variableName *string) *string {
	if err := c.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}


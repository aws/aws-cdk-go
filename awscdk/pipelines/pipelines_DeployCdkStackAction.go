package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Action to deploy a CDK Stack.
//
// Adds two CodePipeline Actions to the pipeline: one to create a ChangeSet
// and one to execute it.
//
// You do not need to instantiate this action yourself -- it will automatically
// be added by the pipeline when you add stack artifacts or entire stages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var role role
//
//   deployCdkStackAction := awscdk.Pipelines.NewDeployCdkStackAction(&deployCdkStackActionProps{
//   	actionRole: role,
//   	cloudAssemblyInput: artifact,
//   	stackName: jsii.String("stackName"),
//   	templatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	cloudFormationExecutionRole: role,
//   	dependencyStackArtifactIds: []*string{
//   		jsii.String("dependencyStackArtifactIds"),
//   	},
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   	region: jsii.String("region"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	templateConfigurationPath: jsii.String("templateConfigurationPath"),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackAction interface {
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// Artifact ids of the artifact this stack artifact depends on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DependencyStackArtifactIds() *[]*string
	// The runorder for the execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder() *float64
	// The runorder for the prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder() *float64
	// Artifact id of the artifact this action was based on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackArtifactId() *string
	// Name of the deployed stack.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackName() *string
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for DeployCdkStackAction
type jsiiProxy_DeployCdkStackAction struct {
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_DeployCdkStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) DependencyStackArtifactIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyStackArtifactIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) ExecuteRunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeRunOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) PrepareRunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prepareRunOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) StackArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewDeployCdkStackAction(props *DeployCdkStackActionProps) DeployCdkStackAction {
	_init_.Initialize()

	j := jsiiProxy_DeployCdkStackAction{}

	_jsii_.Create(
		"monocdk.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewDeployCdkStackAction_Override(d DeployCdkStackAction, props *DeployCdkStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		d,
	)
}

// Construct a DeployCdkStackAction from a Stack artifact.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func DeployCdkStackAction_FromStackArtifact(scope constructs.Construct, artifact cxapi.CloudFormationStackArtifact, options *CdkStackActionFromArtifactOptions) DeployCdkStackAction {
	_init_.Initialize()

	var returns DeployCdkStackAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.DeployCdkStackAction",
		"fromStackArtifact",
		[]interface{}{scope, artifact, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeployCdkStackAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeployCdkStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}


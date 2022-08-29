package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to connect to an existing ServiceCatalog product.
//
// **Note**: this class is still experimental, and may have breaking changes in the future!
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   cdkBuildOutput := codepipeline.NewArtifact()
//   serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&serviceCatalogDeployActionBeta1Props{
//   	actionName: jsii.String("ServiceCatalogDeploy"),
//   	templatePath: cdkBuildOutput.atPath(jsii.String("Sample.template.json")),
//   	productVersionName: jsii.String("Version - " + date.now.toString),
//   	productVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
//   	productId: jsii.String("prod-XXXXXXXX"),
//   })
//
type ServiceCatalogDeployActionBeta1 interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for ServiceCatalogDeployActionBeta1
type jsiiProxy_ServiceCatalogDeployActionBeta1 struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_ServiceCatalogDeployActionBeta1) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceCatalogDeployActionBeta1) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewServiceCatalogDeployActionBeta1(props *ServiceCatalogDeployActionBeta1Props) ServiceCatalogDeployActionBeta1 {
	_init_.Initialize()

	j := jsiiProxy_ServiceCatalogDeployActionBeta1{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewServiceCatalogDeployActionBeta1_Override(s ServiceCatalogDeployActionBeta1, props *ServiceCatalogDeployActionBeta1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}


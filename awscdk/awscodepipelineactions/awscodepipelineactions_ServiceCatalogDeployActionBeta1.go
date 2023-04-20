package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// CodePipeline action to connect to an existing ServiceCatalog product.
//
// **Note**: this class is still experimental, and may have breaking changes in the future!
//
// Example:
//   cdkBuildOutput := codepipeline.NewArtifact()
//   serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&ServiceCatalogDeployActionBeta1Props{
//   	ActionName: jsii.String("ServiceCatalogDeploy"),
//   	TemplatePath: cdkBuildOutput.AtPath(jsii.String("Sample.template.json")),
//   	ProductVersionName: jsii.String("Version - " + date.now.toString),
//   	ProductVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
//   	ProductId: jsii.String("prod-XXXXXXXX"),
//   })
//
// Experimental.
type ServiceCatalogDeployActionBeta1 interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	// Experimental.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	// Experimental.
	Bound(_scope awscdk.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Experimental.
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


// Experimental.
func NewServiceCatalogDeployActionBeta1(props *ServiceCatalogDeployActionBeta1Props) ServiceCatalogDeployActionBeta1 {
	_init_.Initialize()

	if err := validateNewServiceCatalogDeployActionBeta1Parameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceCatalogDeployActionBeta1{}

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceCatalogDeployActionBeta1_Override(s ServiceCatalogDeployActionBeta1, props *ServiceCatalogDeployActionBeta1Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.ServiceCatalogDeployActionBeta1",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) Bound(_scope awscdk.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBoundParameters(_scope, _stage, options); err != nil {
		panic(err)
	}
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
	if err := s.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
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
	if err := s.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}


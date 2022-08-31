package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Action to publish an asset in the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to prepare and publish the asset.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline when you add stacks that use assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dependable iDependable
//   var role role
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   publishAssetsAction := awscdk.Pipelines.NewPublishAssetsAction(this, jsii.String("MyPublishAssetsAction"), &publishAssetsActionProps{
//   	actionName: jsii.String("actionName"),
//   	assetType: awscdk.*Pipelines.assetType_FILE,
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	createBuildspecFile: jsii.Boolean(false),
//   	dependable: dependable,
//   	preInstallCommands: []*string{
//   		jsii.String("preInstallCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	role: role,
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   	vpc: vpc,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type PublishAssetsAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add a single publishing command.
	//
	// Manifest path should be relative to the root Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddPublishCommand(relativeManifestPath *string, assetSelector *string)
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Validate() *[]*string
}

// The jsii proxy struct for PublishAssetsAction
type jsiiProxy_PublishAssetsAction struct {
	internal.Type__awscdkConstruct
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_PublishAssetsAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublishAssetsAction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewPublishAssetsAction(scope constructs.Construct, id *string, props *PublishAssetsActionProps) PublishAssetsAction {
	_init_.Initialize()

	if err := validateNewPublishAssetsActionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PublishAssetsAction{}

	_jsii_.Create(
		"monocdk.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewPublishAssetsAction_Override(p PublishAssetsAction, scope constructs.Construct, id *string, props *PublishAssetsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func PublishAssetsAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublishAssetsAction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.PublishAssetsAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublishAssetsAction) AddPublishCommand(relativeManifestPath *string, assetSelector *string) {
	if err := p.validateAddPublishCommandParameters(relativeManifestPath, assetSelector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPublishCommand",
		[]interface{}{relativeManifestPath, assetSelector},
	)
}

func (p *jsiiProxy_PublishAssetsAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := p.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublishAssetsAction) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublishAssetsAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := p.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublishAssetsAction) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublishAssetsAction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublishAssetsAction) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublishAssetsAction) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublishAssetsAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublishAssetsAction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


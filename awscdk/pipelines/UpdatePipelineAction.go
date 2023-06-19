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

// Action to self-mutate the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to deploy the pipeline stack.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dockerCredential dockerCredential
//
//   updatePipelineAction := awscdk.Pipelines.NewUpdatePipelineAction(this, jsii.String("MyUpdatePipelineAction"), &UpdatePipelineActionProps{
//   	CloudAssemblyInput: artifact,
//   	PipelineStackHierarchicalId: jsii.String("pipelineStackHierarchicalId"),
//
//   	// the properties below are optional
//   	BuildSpec: buildSpec,
//   	CdkCliVersion: jsii.String("cdkCliVersion"),
//   	DockerCredentials: []*dockerCredential{
//   		dockerCredential,
//   	},
//   	PipelineStackName: jsii.String("pipelineStackName"),
//   	Privileged: jsii.Boolean(false),
//   	ProjectName: jsii.String("projectName"),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type UpdatePipelineAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for UpdatePipelineAction
type jsiiProxy_UpdatePipelineAction struct {
	internal.Type__awscdkConstruct
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_UpdatePipelineAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdatePipelineAction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewUpdatePipelineAction(scope constructs.Construct, id *string, props *UpdatePipelineActionProps) UpdatePipelineAction {
	_init_.Initialize()

	if err := validateNewUpdatePipelineActionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UpdatePipelineAction{}

	_jsii_.Create(
		"monocdk.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewUpdatePipelineAction_Override(u UpdatePipelineAction, scope constructs.Construct, id *string, props *UpdatePipelineActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func UpdatePipelineAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpdatePipelineAction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.UpdatePipelineAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdatePipelineAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := u.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdatePipelineAction) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdatePipelineAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := u.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		u,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdatePipelineAction) OnSynthesize(session constructs.ISynthesisSession) {
	if err := u.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UpdatePipelineAction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdatePipelineAction) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdatePipelineAction) Synthesize(session awscdk.ISynthesisSession) {
	if err := u.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_UpdatePipelineAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpdatePipelineAction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


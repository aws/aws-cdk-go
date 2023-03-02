package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Stage in a CdkPipeline.
//
// You don't need to instantiate this class directly. Use
// `cdkPipeline.addStage()` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var stage iStage
//   var stageHost iStageHost
//   var topic topic
//
//   cdkStage := awscdk.Pipelines.NewCdkStage(this, jsii.String("MyCdkStage"), &cdkStageProps{
//   	cloudAssemblyArtifact: artifact,
//   	host: stageHost,
//   	pipelineStage: stage,
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkStage interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add one or more CodePipeline Actions.
	//
	// You need to make sure it is created with the right runOrder. Call `nextSequentialRunOrder()`
	// for every action to get actions to execute in sequence.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddActions(actions ...awscodepipeline.IAction)
	// Add all stacks in the application Stage to this stage.
	//
	// The application construct should subclass `Stage` and can contain any
	// number of `Stacks` inside it that may have dependency relationships
	// on one another.
	//
	// All stacks in the application will be deployed in the appropriate order,
	// and all assets found in the application will be added to the asset
	// publishing stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddApplication(appStage awscdk.Stage, options *AddStageOptions)
	// Add a manual approval action.
	//
	// If you need more flexibility than what this method offers,
	// use `addAction` with a `ManualApprovalAction`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddManualApprovalAction(options *AddManualApprovalOptions)
	// Add a deployment action based on a stack artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions)
	// Whether this Stage contains an action to deploy the given stack, identified by its artifact ID.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DeploysStack(artifactId *string) *bool
	// Return the runOrder number necessary to run the next Action in sequence with the rest.
	//
	// FIXME: This is here because Actions are immutable and can't be reordered
	// after creation, nor is there a way to specify relative priorities, which
	// is a limitation that we should take away in the base library.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	NextSequentialRunOrder(count *float64) *float64
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

// The jsii proxy struct for CdkStage
type jsiiProxy_CdkStage struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CdkStage) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkStage(scope constructs.Construct, id *string, props *CdkStageProps) CdkStage {
	_init_.Initialize()

	if err := validateNewCdkStageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdkStage{}

	_jsii_.Create(
		"monocdk.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkStage_Override(c CdkStage, scope constructs.Construct, id *string, props *CdkStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func CdkStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkStage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CdkStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkStage) AddActions(actions ...awscodepipeline.IAction) {
	args := []interface{}{}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addActions",
		args,
	)
}

func (c *jsiiProxy_CdkStage) AddApplication(appStage awscdk.Stage, options *AddStageOptions) {
	if err := c.validateAddApplicationParameters(appStage, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addApplication",
		[]interface{}{appStage, options},
	)
}

func (c *jsiiProxy_CdkStage) AddManualApprovalAction(options *AddManualApprovalOptions) {
	if err := c.validateAddManualApprovalActionParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addManualApprovalAction",
		[]interface{}{options},
	)
}

func (c *jsiiProxy_CdkStage) AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions) {
	if err := c.validateAddStackArtifactDeploymentParameters(stackArtifact, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addStackArtifactDeployment",
		[]interface{}{stackArtifact, options},
	)
}

func (c *jsiiProxy_CdkStage) DeploysStack(artifactId *string) *bool {
	if err := c.validateDeploysStackParameters(artifactId); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"deploysStack",
		[]interface{}{artifactId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkStage) NextSequentialRunOrder(count *float64) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"nextSequentialRunOrder",
		[]interface{}{count},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkStage) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkStage) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CdkStage) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkStage) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkStage) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CdkStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkStage) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


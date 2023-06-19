package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The resource representing registering a custom Action with CodePipeline.
//
// For the Action to be usable, it has to be registered for every region and every account it's used in.
// In addition to this class, you should most likely also provide your clients a class
// representing your custom Action, extending the Action class,
// and taking the `actionProperties` as properly typed, construction properties.
//
// Example:
//   // Make a custom CodePipeline Action
//   // Make a custom CodePipeline Action
//   codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &CustomActionRegistrationProps{
//   	Category: codepipeline.ActionCategory_SOURCE,
//   	ArtifactBounds: &ActionArtifactBounds{
//   		MinInputs: jsii.Number(0),
//   		MaxInputs: jsii.Number(0),
//   		MinOutputs: jsii.Number(1),
//   		MaxOutputs: jsii.Number(1),
//   	},
//   	Provider: jsii.String("GenericGitSource"),
//   	Version: jsii.String("1"),
//   	EntityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	ExecutionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	ActionProperties: []customActionProperty{
//   		&customActionProperty{
//   			Name: jsii.String("Branch"),
//   			Required: jsii.Boolean(true),
//   			Key: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Queryable: jsii.Boolean(false),
//   			Description: jsii.String("Git branch to pull"),
//   			Type: jsii.String("String"),
//   		},
//   		&customActionProperty{
//   			Name: jsii.String("GitUrl"),
//   			Required: jsii.Boolean(true),
//   			Key: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Queryable: jsii.Boolean(false),
//   			Description: jsii.String("SSH git clone URL"),
//   			Type: jsii.String("String"),
//   		},
//   	},
//   })
//
// Experimental.
type CustomActionRegistration interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for CustomActionRegistration
type jsiiProxy_CustomActionRegistration struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CustomActionRegistration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomActionRegistration(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) CustomActionRegistration {
	_init_.Initialize()

	if err := validateNewCustomActionRegistrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomActionRegistration{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomActionRegistration_Override(c CustomActionRegistration, scope constructs.Construct, id *string, props *CustomActionRegistrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CustomActionRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomActionRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomActionRegistration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomActionRegistration) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomActionRegistration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomActionRegistration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomActionRegistration) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomActionRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomActionRegistration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


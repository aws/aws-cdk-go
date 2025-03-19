package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
type CustomActionRegistration interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomActionRegistration
type jsiiProxy_CustomActionRegistration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomActionRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCustomActionRegistration(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) CustomActionRegistration {
	_init_.Initialize()

	if err := validateNewCustomActionRegistrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomActionRegistration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomActionRegistration_Override(c CustomActionRegistration, scope constructs.Construct, id *string, props *CustomActionRegistrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CustomActionRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomActionRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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


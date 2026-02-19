package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to create a Prompt Version with CDK.
//
// Creates a version of the prompt. Use this to create a static snapshot of your
// prompt that can be deployed to production. Versions allow you to easily switch
// between different configurations for your prompt and update your application
// with the most appropriate version for your use-case.
//
// Example:
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//   claudeModel := bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0()
//
//   variant1 := bedrock.PromptVariant_Text(&TextPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	Model: claudeModel,
//   	PromptVariables: []*string{
//   		jsii.String("topic"),
//   	},
//   	PromptText: jsii.String("This is my first text prompt. Please summarize our conversation on: {{topic}}."),
//   	InferenceConfiguration: bedrock.PromptInferenceConfiguration_Text(&PromptInferenceConfigurationProps{
//   		Temperature: jsii.Number(1),
//   		TopP: jsii.Number(0.999),
//   		MaxTokens: jsii.Number(2000),
//   	}),
//   })
//
//   prompt1 := bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
//   	PromptName: jsii.String("prompt1"),
//   	Description: jsii.String("my first prompt"),
//   	DefaultVariant: variant1,
//   	Variants: []IPromptVariant{
//   		variant1,
//   	},
//   	KmsKey: cmk,
//   })
//
//   promptVersion := bedrock.NewPromptVersion(this, jsii.String("MyPromptVersion"), &PromptVersionProps{
//   	Prompt: prompt1,
//   	Description: jsii.String("my first version"),
//   })
//   //or alternatively:
//   // const promptVersion = prompt1.createVersion('my first version');
//   versionString := promptVersion.Version
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-deploy.html
//
// Experimental.
type PromptVersion interface {
	constructs.Construct
	// The description of the prompt version.
	// Experimental.
	Description() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The prompt used by this version.
	// Experimental.
	Prompt() IPrompt
	// The version of the prompt that was created.
	// Experimental.
	Version() *string
	// The Amazon Resource Name (ARN) of the prompt version.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345:1"
	//
	// Experimental.
	VersionArn() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for PromptVersion
type jsiiProxy_PromptVersion struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PromptVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) Prompt() IPrompt {
	var returns IPrompt
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) VersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionArn",
		&returns,
	)
	return returns
}


// ****************************************************************************                           CONSTRUCTOR ***************************************************************************.
// Experimental.
func NewPromptVersion(scope constructs.Construct, id *string, props *PromptVersionProps) PromptVersion {
	_init_.Initialize()

	if err := validateNewPromptVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PromptVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// ****************************************************************************                           CONSTRUCTOR ***************************************************************************.
// Experimental.
func NewPromptVersion_Override(p PromptVersion, scope constructs.Construct, id *string, props *PromptVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptVersion",
		[]interface{}{scope, id, props},
		p,
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
// Experimental.
func PromptVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePromptVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PromptVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PromptVersion) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}


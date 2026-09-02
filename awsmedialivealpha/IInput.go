package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Input.
// Experimental.
type IInput interface {
	interfacesawsmedialive.IInputRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the input.
	// Experimental.
	InputArn() *string
	// The input class (STANDARD or SINGLE_PIPELINE).
	//
	// Only available for input types where the pipeline count is known at construct time
	// (e.g. mediaConnectRouter). Undefined for imported inputs and other types.
	// Experimental.
	InputClass() *string
	// For push inputs, the destination URLs where the upstream system sends content.
	// Experimental.
	InputDestinations() *[]*string
	// The ID of the input.
	// Experimental.
	InputId() *string
	// For pull inputs, the source URLs where MediaLive pulls content from.
	// Experimental.
	InputSources() *[]*string
	// The input type (e.g. SRT_CALLER, MP4_FILE, URL_PULL). Undefined for imported inputs.
	// Experimental.
	InputType() *string
}

// The jsii proxy for IInput
type jsiiProxy_IInput struct {
	internal.Type__interfacesawsmedialiveIInputRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IInput) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInput) InputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputRef() *interfacesawsmedialive.InputReference {
	var returns *interfacesawsmedialive.InputReference
	_jsii_.Get(
		j,
		"inputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


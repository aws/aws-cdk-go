package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents predefined topics that can be used to filter content.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   	// Configure tier for topic filters (optional)
//   	TopicsTierConfig: bedrock.TierConfig_STANDARD,
//   })
//
//   // Use a predefined topic
//   guardrail.AddDeniedTopicFilter(bedrock.Topic_FINANCIAL_ADVICE())
//
//   // Create a custom topic with input/output actions
//   guardrail.AddDeniedTopicFilter(bedrock.Topic_Custom(&CustomTopicProps{
//   	Name: jsii.String("Legal_Advice"),
//   	Definition: jsii.String("Offering guidance or suggestions on legal matters, legal actions, interpretation of laws, or legal rights and responsibilities."),
//   	Examples: []*string{
//   		jsii.String("Can I sue someone for this?"),
//   		jsii.String("What are my legal rights in this situation?"),
//   		jsii.String("Is this action against the law?"),
//   		jsii.String("What should I do to file a legal complaint?"),
//   		jsii.String("Can you explain this law to me?"),
//   	},
//   	// props below are optional
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_NONE,
//   	OutputEnabled: jsii.Boolean(true),
//   }))
//
// Experimental.
type Topic interface {
	// Definition of the topic.
	// Experimental.
	Definition() *string
	// Representative phrases that refer to the topic.
	// Experimental.
	Examples() *[]*string
	// The action to take when a topic is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction() GuardrailAction
	// Whether the topic filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled() *bool
	// The name of the topic to deny.
	// Experimental.
	Name() *string
	// The action to take when a topic is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction() GuardrailAction
	// Whether the topic filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled() *bool
}

// The jsii proxy struct for Topic
type jsiiProxy_Topic struct {
	_ byte // padding
}

func (j *jsiiProxy_Topic) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Examples() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"examples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) InputAction() GuardrailAction {
	var returns GuardrailAction
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) InputEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) OutputAction() GuardrailAction {
	var returns GuardrailAction
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) OutputEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopic(props *CustomTopicProps) Topic {
	_init_.Initialize()

	if err := validateNewTopicParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Topic{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopic_Override(t Topic, props *CustomTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		[]interface{}{props},
		t,
	)
}

// Create a custom topic filter.
// Experimental.
func Topic_Custom(props *CustomTopicProps) Topic {
	_init_.Initialize()

	if err := validateTopic_CustomParameters(props); err != nil {
		panic(err)
	}
	var returns Topic

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"custom",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Topic_FINANCIAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"FINANCIAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_INAPPROPRIATE_CONTENT() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"INAPPROPRIATE_CONTENT",
		&returns,
	)
	return returns
}

func Topic_LEGAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"LEGAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_MEDICAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"MEDICAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_POLITICAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		"POLITICAL_ADVICE",
		&returns,
	)
	return returns
}


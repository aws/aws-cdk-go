package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Memory class for managing Bedrock Agent memory configurations.
//
// Enables conversational context retention
// across multiple sessions through session identifiers. Memory context is stored with unique
// memory IDs per user, allowing access to conversation history and summaries. Supports viewing
// stored sessions and clearing memory.
//
// Example:
//   agent := bedrock.NewAgent(this, jsii.String("MyAgent"), &AgentProps{
//   	AgentName: jsii.String("MyAgent"),
//   	Instruction: jsii.String("Your instruction here"),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Memory: awsbedrockalpha.Memory_SessionSummary(&SessionSummaryMemoryProps{
//   		MaxRecentSessions: jsii.Number(10),
//   		 // Keep the last 10 session summaries
//   		MemoryDuration: awscdk.Duration_Days(jsii.Number(20)),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-memory.html
//
// Experimental.
type Memory interface {
}

// The jsii proxy struct for Memory
type jsiiProxy_Memory struct {
	_ byte // padding
}

// Experimental.
func NewMemory(props *SessionSummaryMemoryProps) Memory {
	_init_.Initialize()

	if err := validateNewMemoryParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Memory{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Memory",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMemory_Override(m Memory, props *SessionSummaryMemoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Memory",
		[]interface{}{props},
		m,
	)
}

// Creates a session summary memory with custom configuration.
//
// Returns: Memory instance.
// Experimental.
func Memory_SessionSummary(props *SessionSummaryMemoryProps) Memory {
	_init_.Initialize()

	if err := validateMemory_SessionSummaryParameters(props); err != nil {
		panic(err)
	}
	var returns Memory

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.Memory",
		"sessionSummary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Memory_SESSION_SUMMARY() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.Memory",
		"SESSION_SUMMARY",
		&returns,
	)
	return returns
}


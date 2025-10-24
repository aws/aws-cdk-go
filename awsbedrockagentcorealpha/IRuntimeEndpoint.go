package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Interface for Runtime Endpoint resources.
// Experimental.
type IRuntimeEndpoint interface {
	awscdk.IResource
	// The ARN of the parent agent runtime.
	// Experimental.
	AgentRuntimeArn() *string
	// The ARN of the runtime endpoint resource.
	//
	// Example:
	//   "arn:aws:bedrock-agentcore:us-west-2:123456789012:agent-runtime-endpoint/endpoint-abc123"
	//
	// Experimental.
	AgentRuntimeEndpointArn() *string
	// When the endpoint was created.
	// Experimental.
	CreatedAt() *string
	// The description of the runtime endpoint.
	// Experimental.
	Description() *string
	// The name of the runtime endpoint.
	// Experimental.
	EndpointName() *string
	// The live version of the agent runtime that is currently serving requests.
	// Experimental.
	LiveVersion() *string
	// The current status of the runtime endpoint.
	// Experimental.
	Status() *string
	// The target version the endpoint is transitioning to (during updates).
	// Experimental.
	TargetVersion() *string
}

// The jsii proxy for IRuntimeEndpoint
type jsiiProxy_IRuntimeEndpoint struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRuntimeEndpoint) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) AgentRuntimeEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) LiveVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeEndpoint) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}


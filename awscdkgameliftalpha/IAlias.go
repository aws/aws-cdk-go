package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Gamelift Alias for a Gamelift fleet destination.
// Experimental.
type IAlias interface {
	IGameSessionQueueDestination
	awscdk.IResource
	// The ARN of the alias.
	// Experimental.
	AliasArn() *string
	// The Identifier of the alias.
	// Experimental.
	AliasId() *string
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	jsiiProxy_IGameSessionQueueDestination
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) ResourceArnForDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnForDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


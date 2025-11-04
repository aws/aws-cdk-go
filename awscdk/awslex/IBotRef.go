package awslex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bot.
// Experimental.
type IBotRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Bot resource.
	// Experimental.
	BotRef() *BotReference
}

// The jsii proxy for IBotRef
type jsiiProxy_IBotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IBotRef) BotRef() *BotReference {
	var returns *BotReference
	_jsii_.Get(
		j,
		"botRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBotRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Runtime.
// Experimental.
type IRuntimeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Runtime resource.
	// Experimental.
	RuntimeRef() *RuntimeReference
}

// The jsii proxy for IRuntimeRef
type jsiiProxy_IRuntimeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRuntimeRef) RuntimeRef() *RuntimeReference {
	var returns *RuntimeReference
	_jsii_.Get(
		j,
		"runtimeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


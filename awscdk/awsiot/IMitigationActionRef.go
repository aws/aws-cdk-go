package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MitigationAction.
// Experimental.
type IMitigationActionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MitigationAction resource.
	// Experimental.
	MitigationActionRef() *MitigationActionReference
}

// The jsii proxy for IMitigationActionRef
type jsiiProxy_IMitigationActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMitigationActionRef) MitigationActionRef() *MitigationActionReference {
	var returns *MitigationActionReference
	_jsii_.Get(
		j,
		"mitigationActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMitigationActionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMitigationActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


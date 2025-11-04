package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WALWorkspace.
// Experimental.
type IWALWorkspaceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WALWorkspace resource.
	// Experimental.
	WalWorkspaceRef() *WALWorkspaceReference
}

// The jsii proxy for IWALWorkspaceRef
type jsiiProxy_IWALWorkspaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IWALWorkspaceRef) WalWorkspaceRef() *WALWorkspaceReference {
	var returns *WALWorkspaceReference
	_jsii_.Get(
		j,
		"walWorkspaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWALWorkspaceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWALWorkspaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


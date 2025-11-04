package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DelegatedAdmin.
// Experimental.
type IDelegatedAdminRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DelegatedAdmin resource.
	// Experimental.
	DelegatedAdminRef() *DelegatedAdminReference
}

// The jsii proxy for IDelegatedAdminRef
type jsiiProxy_IDelegatedAdminRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDelegatedAdminRef) DelegatedAdminRef() *DelegatedAdminReference {
	var returns *DelegatedAdminReference
	_jsii_.Get(
		j,
		"delegatedAdminRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegatedAdminRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegatedAdminRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


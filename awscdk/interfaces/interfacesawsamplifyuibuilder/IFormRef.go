package interfacesawsamplifyuibuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Form.
// Experimental.
type IFormRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Form resource.
	// Experimental.
	FormRef() *FormReference
}

// The jsii proxy for IFormRef
type jsiiProxy_IFormRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFormRef) FormRef() *FormReference {
	var returns *FormReference
	_jsii_.Get(
		j,
		"formRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


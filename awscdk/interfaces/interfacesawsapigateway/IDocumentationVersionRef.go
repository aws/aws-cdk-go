package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentationVersion.
// Experimental.
type IDocumentationVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DocumentationVersion resource.
	// Experimental.
	DocumentationVersionRef() *DocumentationVersionReference
}

// The jsii proxy for IDocumentationVersionRef
type jsiiProxy_IDocumentationVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDocumentationVersionRef) DocumentationVersionRef() *DocumentationVersionReference {
	var returns *DocumentationVersionReference
	_jsii_.Get(
		j,
		"documentationVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentationVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentationVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


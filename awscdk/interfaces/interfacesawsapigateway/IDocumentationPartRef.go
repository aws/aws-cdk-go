package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentationPart.
// Experimental.
type IDocumentationPartRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DocumentationPart resource.
	// Experimental.
	DocumentationPartRef() *DocumentationPartReference
}

// The jsii proxy for IDocumentationPartRef
type jsiiProxy_IDocumentationPartRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDocumentationPartRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDocumentationPartRef) DocumentationPartRef() *DocumentationPartReference {
	var returns *DocumentationPartReference
	_jsii_.Get(
		j,
		"documentationPartRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentationPartRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentationPartRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


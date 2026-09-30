package interfacesawscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentClassifierEndpoint.
// Experimental.
type IDocumentClassifierEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DocumentClassifierEndpoint resource.
	// Experimental.
	DocumentClassifierEndpointRef() *DocumentClassifierEndpointReference
}

// The jsii proxy for IDocumentClassifierEndpointRef
type jsiiProxy_IDocumentClassifierEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDocumentClassifierEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDocumentClassifierEndpointRef) DocumentClassifierEndpointRef() *DocumentClassifierEndpointReference {
	var returns *DocumentClassifierEndpointReference
	_jsii_.Get(
		j,
		"documentClassifierEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentClassifierEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentClassifierEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


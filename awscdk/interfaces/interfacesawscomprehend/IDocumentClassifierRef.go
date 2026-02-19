package interfacesawscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentClassifier.
// Experimental.
type IDocumentClassifierRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DocumentClassifier resource.
	// Experimental.
	DocumentClassifierRef() *DocumentClassifierReference
}

// The jsii proxy for IDocumentClassifierRef
type jsiiProxy_IDocumentClassifierRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDocumentClassifierRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDocumentClassifierRef) DocumentClassifierRef() *DocumentClassifierReference {
	var returns *DocumentClassifierReference
	_jsii_.Get(
		j,
		"documentClassifierRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentClassifierRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDocumentClassifierRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


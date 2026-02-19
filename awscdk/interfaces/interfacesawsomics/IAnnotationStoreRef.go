package interfacesawsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnnotationStore.
// Experimental.
type IAnnotationStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AnnotationStore resource.
	// Experimental.
	AnnotationStoreRef() *AnnotationStoreReference
}

// The jsii proxy for IAnnotationStoreRef
type jsiiProxy_IAnnotationStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAnnotationStoreRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAnnotationStoreRef) AnnotationStoreRef() *AnnotationStoreReference {
	var returns *AnnotationStoreReference
	_jsii_.Get(
		j,
		"annotationStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnnotationStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnnotationStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Reference.
// Experimental.
type IReferenceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Reference resource.
	// Experimental.
	ReferenceRef() *ReferenceReference
}

// The jsii proxy for IReferenceRef
type jsiiProxy_IReferenceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReferenceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReferenceRef) ReferenceRef() *ReferenceReference {
	var returns *ReferenceReference
	_jsii_.Get(
		j,
		"referenceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReferenceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReferenceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


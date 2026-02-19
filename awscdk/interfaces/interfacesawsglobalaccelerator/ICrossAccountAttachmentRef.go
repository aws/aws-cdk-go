package interfacesawsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CrossAccountAttachment.
// Experimental.
type ICrossAccountAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CrossAccountAttachment resource.
	// Experimental.
	CrossAccountAttachmentRef() *CrossAccountAttachmentReference
}

// The jsii proxy for ICrossAccountAttachmentRef
type jsiiProxy_ICrossAccountAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICrossAccountAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICrossAccountAttachmentRef) CrossAccountAttachmentRef() *CrossAccountAttachmentReference {
	var returns *CrossAccountAttachmentReference
	_jsii_.Get(
		j,
		"crossAccountAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICrossAccountAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICrossAccountAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


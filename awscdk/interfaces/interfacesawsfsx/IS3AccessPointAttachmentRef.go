package interfacesawsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a S3AccessPointAttachment.
// Experimental.
type IS3AccessPointAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a S3AccessPointAttachment resource.
	// Experimental.
	S3AccessPointAttachmentRef() *S3AccessPointAttachmentReference
}

// The jsii proxy for IS3AccessPointAttachmentRef
type jsiiProxy_IS3AccessPointAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IS3AccessPointAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IS3AccessPointAttachmentRef) S3AccessPointAttachmentRef() *S3AccessPointAttachmentReference {
	var returns *S3AccessPointAttachmentReference
	_jsii_.Get(
		j,
		"s3AccessPointAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IS3AccessPointAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IS3AccessPointAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


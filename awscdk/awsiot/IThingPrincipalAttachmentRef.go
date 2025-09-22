package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingPrincipalAttachment.
// Experimental.
type IThingPrincipalAttachmentRef interface {
	constructs.IConstruct
	// A reference to a ThingPrincipalAttachment resource.
	// Experimental.
	ThingPrincipalAttachmentRef() *ThingPrincipalAttachmentReference
}

// The jsii proxy for IThingPrincipalAttachmentRef
type jsiiProxy_IThingPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IThingPrincipalAttachmentRef) ThingPrincipalAttachmentRef() *ThingPrincipalAttachmentReference {
	var returns *ThingPrincipalAttachmentReference
	_jsii_.Get(
		j,
		"thingPrincipalAttachmentRef",
		&returns,
	)
	return returns
}


package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyPrincipalAttachment.
// Experimental.
type IPolicyPrincipalAttachmentRef interface {
	constructs.IConstruct
	// A reference to a PolicyPrincipalAttachment resource.
	// Experimental.
	PolicyPrincipalAttachmentRef() *PolicyPrincipalAttachmentReference
}

// The jsii proxy for IPolicyPrincipalAttachmentRef
type jsiiProxy_IPolicyPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyPrincipalAttachmentRef) PolicyPrincipalAttachmentRef() *PolicyPrincipalAttachmentReference {
	var returns *PolicyPrincipalAttachmentReference
	_jsii_.Get(
		j,
		"policyPrincipalAttachmentRef",
		&returns,
	)
	return returns
}


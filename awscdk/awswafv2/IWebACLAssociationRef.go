package awswafv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACLAssociation.
// Experimental.
type IWebACLAssociationRef interface {
	constructs.IConstruct
	// A reference to a WebACLAssociation resource.
	// Experimental.
	WebAclAssociationRef() *WebACLAssociationReference
}

// The jsii proxy for IWebACLAssociationRef
type jsiiProxy_IWebACLAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWebACLAssociationRef) WebAclAssociationRef() *WebACLAssociationReference {
	var returns *WebACLAssociationReference
	_jsii_.Get(
		j,
		"webAclAssociationRef",
		&returns,
	)
	return returns
}


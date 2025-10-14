package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACL.
// Experimental.
type IWebACLRef interface {
	constructs.IConstruct
	// A reference to a WebACL resource.
	// Experimental.
	WebAclRef() *WebACLReference
}

// The jsii proxy for IWebACLRef
type jsiiProxy_IWebACLRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWebACLRef) WebAclRef() *WebACLReference {
	var returns *WebACLReference
	_jsii_.Get(
		j,
		"webAclRef",
		&returns,
	)
	return returns
}


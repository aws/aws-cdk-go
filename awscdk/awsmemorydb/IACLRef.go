package awsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ACL.
// Experimental.
type IACLRef interface {
	constructs.IConstruct
	// A reference to a ACL resource.
	// Experimental.
	AclRef() *ACLReference
}

// The jsii proxy for IACLRef
type jsiiProxy_IACLRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IACLRef) AclRef() *ACLReference {
	var returns *ACLReference
	_jsii_.Get(
		j,
		"aclRef",
		&returns,
	)
	return returns
}


package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityControl.
// Experimental.
type ISecurityControlRef interface {
	constructs.IConstruct
	// A reference to a SecurityControl resource.
	// Experimental.
	SecurityControlRef() *SecurityControlReference
}

// The jsii proxy for ISecurityControlRef
type jsiiProxy_ISecurityControlRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityControlRef) SecurityControlRef() *SecurityControlReference {
	var returns *SecurityControlReference
	_jsii_.Get(
		j,
		"securityControlRef",
		&returns,
	)
	return returns
}


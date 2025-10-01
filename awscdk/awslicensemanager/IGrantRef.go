package awslicensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Grant.
// Experimental.
type IGrantRef interface {
	constructs.IConstruct
	// A reference to a Grant resource.
	// Experimental.
	GrantRef() *GrantReference
}

// The jsii proxy for IGrantRef
type jsiiProxy_IGrantRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGrantRef) GrantRef() *GrantReference {
	var returns *GrantReference
	_jsii_.Get(
		j,
		"grantRef",
		&returns,
	)
	return returns
}


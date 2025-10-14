package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GuardHook.
// Experimental.
type IGuardHookRef interface {
	constructs.IConstruct
	// A reference to a GuardHook resource.
	// Experimental.
	GuardHookRef() *GuardHookReference
}

// The jsii proxy for IGuardHookRef
type jsiiProxy_IGuardHookRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGuardHookRef) GuardHookRef() *GuardHookReference {
	var returns *GuardHookReference
	_jsii_.Get(
		j,
		"guardHookRef",
		&returns,
	)
	return returns
}


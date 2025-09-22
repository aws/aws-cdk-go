package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookDefaultVersion.
// Experimental.
type IHookDefaultVersionRef interface {
	constructs.IConstruct
	// A reference to a HookDefaultVersion resource.
	// Experimental.
	HookDefaultVersionRef() *HookDefaultVersionReference
}

// The jsii proxy for IHookDefaultVersionRef
type jsiiProxy_IHookDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHookDefaultVersionRef) HookDefaultVersionRef() *HookDefaultVersionReference {
	var returns *HookDefaultVersionReference
	_jsii_.Get(
		j,
		"hookDefaultVersionRef",
		&returns,
	)
	return returns
}


package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookTypeConfig.
// Experimental.
type IHookTypeConfigRef interface {
	constructs.IConstruct
	// A reference to a HookTypeConfig resource.
	// Experimental.
	HookTypeConfigRef() *HookTypeConfigReference
}

// The jsii proxy for IHookTypeConfigRef
type jsiiProxy_IHookTypeConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHookTypeConfigRef) HookTypeConfigRef() *HookTypeConfigReference {
	var returns *HookTypeConfigReference
	_jsii_.Get(
		j,
		"hookTypeConfigRef",
		&returns,
	)
	return returns
}


package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitConditionHandle.
// Experimental.
type IWaitConditionHandleRef interface {
	constructs.IConstruct
	// A reference to a WaitConditionHandle resource.
	// Experimental.
	WaitConditionHandleRef() *WaitConditionHandleReference
}

// The jsii proxy for IWaitConditionHandleRef
type jsiiProxy_IWaitConditionHandleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWaitConditionHandleRef) WaitConditionHandleRef() *WaitConditionHandleReference {
	var returns *WaitConditionHandleReference
	_jsii_.Get(
		j,
		"waitConditionHandleRef",
		&returns,
	)
	return returns
}


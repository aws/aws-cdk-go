package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WarmPool.
// Experimental.
type IWarmPoolRef interface {
	constructs.IConstruct
	// A reference to a WarmPool resource.
	// Experimental.
	WarmPoolRef() *WarmPoolReference
}

// The jsii proxy for IWarmPoolRef
type jsiiProxy_IWarmPoolRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWarmPoolRef) WarmPoolRef() *WarmPoolReference {
	var returns *WarmPoolReference
	_jsii_.Get(
		j,
		"warmPoolRef",
		&returns,
	)
	return returns
}


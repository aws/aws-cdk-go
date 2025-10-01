package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LiveSource.
// Experimental.
type ILiveSourceRef interface {
	constructs.IConstruct
	// A reference to a LiveSource resource.
	// Experimental.
	LiveSourceRef() *LiveSourceReference
}

// The jsii proxy for ILiveSourceRef
type jsiiProxy_ILiveSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILiveSourceRef) LiveSourceRef() *LiveSourceReference {
	var returns *LiveSourceReference
	_jsii_.Get(
		j,
		"liveSourceRef",
		&returns,
	)
	return returns
}


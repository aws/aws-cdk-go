package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickResponse.
// Experimental.
type IQuickResponseRef interface {
	constructs.IConstruct
	// A reference to a QuickResponse resource.
	// Experimental.
	QuickResponseRef() *QuickResponseReference
}

// The jsii proxy for IQuickResponseRef
type jsiiProxy_IQuickResponseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQuickResponseRef) QuickResponseRef() *QuickResponseReference {
	var returns *QuickResponseReference
	_jsii_.Get(
		j,
		"quickResponseRef",
		&returns,
	)
	return returns
}


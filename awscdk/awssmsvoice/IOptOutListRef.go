package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptOutList.
// Experimental.
type IOptOutListRef interface {
	constructs.IConstruct
	// A reference to a OptOutList resource.
	// Experimental.
	OptOutListRef() *OptOutListReference
}

// The jsii proxy for IOptOutListRef
type jsiiProxy_IOptOutListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOptOutListRef) OptOutListRef() *OptOutListReference {
	var returns *OptOutListReference
	_jsii_.Get(
		j,
		"optOutListRef",
		&returns,
	)
	return returns
}


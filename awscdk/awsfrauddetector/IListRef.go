package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a List.
// Experimental.
type IListRef interface {
	constructs.IConstruct
	// A reference to a List resource.
	// Experimental.
	ListRef() *ListReference
}

// The jsii proxy for IListRef
type jsiiProxy_IListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IListRef) ListRef() *ListReference {
	var returns *ListReference
	_jsii_.Get(
		j,
		"listRef",
		&returns,
	)
	return returns
}


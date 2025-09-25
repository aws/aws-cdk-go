package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Index.
// Experimental.
type IIndexRef interface {
	constructs.IConstruct
	// A reference to a Index resource.
	// Experimental.
	IndexRef() *IndexReference
}

// The jsii proxy for IIndexRef
type jsiiProxy_IIndexRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIndexRef) IndexRef() *IndexReference {
	var returns *IndexReference
	_jsii_.Get(
		j,
		"indexRef",
		&returns,
	)
	return returns
}


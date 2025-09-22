package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrimaryTaskSet.
// Experimental.
type IPrimaryTaskSetRef interface {
	constructs.IConstruct
	// A reference to a PrimaryTaskSet resource.
	// Experimental.
	PrimaryTaskSetRef() *PrimaryTaskSetReference
}

// The jsii proxy for IPrimaryTaskSetRef
type jsiiProxy_IPrimaryTaskSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPrimaryTaskSetRef) PrimaryTaskSetRef() *PrimaryTaskSetReference {
	var returns *PrimaryTaskSetReference
	_jsii_.Get(
		j,
		"primaryTaskSetRef",
		&returns,
	)
	return returns
}


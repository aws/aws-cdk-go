package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pool.
// Experimental.
type IPoolRef interface {
	constructs.IConstruct
	// A reference to a Pool resource.
	// Experimental.
	PoolRef() *PoolReference
}

// The jsii proxy for IPoolRef
type jsiiProxy_IPoolRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPoolRef) PoolRef() *PoolReference {
	var returns *PoolReference
	_jsii_.Get(
		j,
		"poolRef",
		&returns,
	)
	return returns
}


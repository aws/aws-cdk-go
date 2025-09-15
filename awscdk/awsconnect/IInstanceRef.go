package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Instance.
// Experimental.
type IInstanceRef interface {
	constructs.IConstruct
	// A reference to a Instance resource.
	// Experimental.
	InstanceRef() *InstanceReference
}

// The jsii proxy for IInstanceRef
type jsiiProxy_IInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceRef) InstanceRef() *InstanceReference {
	var returns *InstanceReference
	_jsii_.Get(
		j,
		"instanceRef",
		&returns,
	)
	return returns
}


package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SdiSource.
// Experimental.
type ISdiSourceRef interface {
	constructs.IConstruct
	// A reference to a SdiSource resource.
	// Experimental.
	SdiSourceRef() *SdiSourceReference
}

// The jsii proxy for ISdiSourceRef
type jsiiProxy_ISdiSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISdiSourceRef) SdiSourceRef() *SdiSourceReference {
	var returns *SdiSourceReference
	_jsii_.Get(
		j,
		"sdiSourceRef",
		&returns,
	)
	return returns
}


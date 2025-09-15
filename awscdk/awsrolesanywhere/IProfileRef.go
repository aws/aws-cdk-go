package awsrolesanywhere

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Profile.
// Experimental.
type IProfileRef interface {
	constructs.IConstruct
	// A reference to a Profile resource.
	// Experimental.
	ProfileRef() *ProfileReference
}

// The jsii proxy for IProfileRef
type jsiiProxy_IProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProfileRef) ProfileRef() *ProfileReference {
	var returns *ProfileReference
	_jsii_.Get(
		j,
		"profileRef",
		&returns,
	)
	return returns
}


package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MitigationAction.
// Experimental.
type IMitigationActionRef interface {
	constructs.IConstruct
	// A reference to a MitigationAction resource.
	// Experimental.
	MitigationActionRef() *MitigationActionReference
}

// The jsii proxy for IMitigationActionRef
type jsiiProxy_IMitigationActionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMitigationActionRef) MitigationActionRef() *MitigationActionReference {
	var returns *MitigationActionReference
	_jsii_.Get(
		j,
		"mitigationActionRef",
		&returns,
	)
	return returns
}


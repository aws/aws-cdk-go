package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PatchBaseline.
// Experimental.
type IPatchBaselineRef interface {
	constructs.IConstruct
	// A reference to a PatchBaseline resource.
	// Experimental.
	PatchBaselineRef() *PatchBaselineReference
}

// The jsii proxy for IPatchBaselineRef
type jsiiProxy_IPatchBaselineRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPatchBaselineRef) PatchBaselineRef() *PatchBaselineReference {
	var returns *PatchBaselineReference
	_jsii_.Get(
		j,
		"patchBaselineRef",
		&returns,
	)
	return returns
}


package awspersonalize

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Solution.
// Experimental.
type ISolutionRef interface {
	constructs.IConstruct
	// A reference to a Solution resource.
	// Experimental.
	SolutionRef() *SolutionReference
}

// The jsii proxy for ISolutionRef
type jsiiProxy_ISolutionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISolutionRef) SolutionRef() *SolutionReference {
	var returns *SolutionReference
	_jsii_.Get(
		j,
		"solutionRef",
		&returns,
	)
	return returns
}


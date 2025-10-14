package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTarget.
// Experimental.
type IAssessmentTargetRef interface {
	constructs.IConstruct
	// A reference to a AssessmentTarget resource.
	// Experimental.
	AssessmentTargetRef() *AssessmentTargetReference
}

// The jsii proxy for IAssessmentTargetRef
type jsiiProxy_IAssessmentTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssessmentTargetRef) AssessmentTargetRef() *AssessmentTargetReference {
	var returns *AssessmentTargetReference
	_jsii_.Get(
		j,
		"assessmentTargetRef",
		&returns,
	)
	return returns
}


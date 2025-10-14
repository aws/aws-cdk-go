package awsauditmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsauditmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assessment.
// Experimental.
type IAssessmentRef interface {
	constructs.IConstruct
	// A reference to a Assessment resource.
	// Experimental.
	AssessmentRef() *AssessmentReference
}

// The jsii proxy for IAssessmentRef
type jsiiProxy_IAssessmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssessmentRef) AssessmentRef() *AssessmentReference {
	var returns *AssessmentReference
	_jsii_.Get(
		j,
		"assessmentRef",
		&returns,
	)
	return returns
}


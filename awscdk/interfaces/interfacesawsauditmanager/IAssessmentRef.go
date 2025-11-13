package interfacesawsauditmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsauditmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assessment.
// Experimental.
type IAssessmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Assessment resource.
	// Experimental.
	AssessmentRef() *AssessmentReference
}

// The jsii proxy for IAssessmentRef
type jsiiProxy_IAssessmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IAssessmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


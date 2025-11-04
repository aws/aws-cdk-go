package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTarget.
// Experimental.
type IAssessmentTargetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AssessmentTarget resource.
	// Experimental.
	AssessmentTargetRef() *AssessmentTargetReference
}

// The jsii proxy for IAssessmentTargetRef
type jsiiProxy_IAssessmentTargetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAssessmentTargetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentTargetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


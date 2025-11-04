package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExperimentTemplate.
// Experimental.
type IExperimentTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ExperimentTemplate resource.
	// Experimental.
	ExperimentTemplateRef() *ExperimentTemplateReference
}

// The jsii proxy for IExperimentTemplateRef
type jsiiProxy_IExperimentTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IExperimentTemplateRef) ExperimentTemplateRef() *ExperimentTemplateReference {
	var returns *ExperimentTemplateReference
	_jsii_.Get(
		j,
		"experimentTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExperimentTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


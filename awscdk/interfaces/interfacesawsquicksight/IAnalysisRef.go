package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Analysis.
// Experimental.
type IAnalysisRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Analysis resource.
	// Experimental.
	AnalysisRef() *AnalysisReference
}

// The jsii proxy for IAnalysisRef
type jsiiProxy_IAnalysisRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAnalysisRef) AnalysisRef() *AnalysisReference {
	var returns *AnalysisReference
	_jsii_.Get(
		j,
		"analysisRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnalysisRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnalysisRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


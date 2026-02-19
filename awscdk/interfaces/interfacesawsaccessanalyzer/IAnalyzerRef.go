package interfacesawsaccessanalyzer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaccessanalyzer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Analyzer.
// Experimental.
type IAnalyzerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Analyzer resource.
	// Experimental.
	AnalyzerRef() *AnalyzerReference
}

// The jsii proxy for IAnalyzerRef
type jsiiProxy_IAnalyzerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAnalyzerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAnalyzerRef) AnalyzerRef() *AnalyzerReference {
	var returns *AnalyzerReference
	_jsii_.Get(
		j,
		"analyzerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnalyzerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnalyzerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


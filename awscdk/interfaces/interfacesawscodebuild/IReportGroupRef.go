package interfacesawscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportGroup.
// Experimental.
type IReportGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReportGroup resource.
	// Experimental.
	ReportGroupRef() *ReportGroupReference
}

// The jsii proxy for IReportGroupRef
type jsiiProxy_IReportGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReportGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReportGroupRef) ReportGroupRef() *ReportGroupReference {
	var returns *ReportGroupReference
	_jsii_.Get(
		j,
		"reportGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


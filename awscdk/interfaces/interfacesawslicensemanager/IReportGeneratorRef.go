package interfacesawslicensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportGenerator.
// Experimental.
type IReportGeneratorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReportGenerator resource.
	// Experimental.
	ReportGeneratorRef() *ReportGeneratorReference
}

// The jsii proxy for IReportGeneratorRef
type jsiiProxy_IReportGeneratorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReportGeneratorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReportGeneratorRef) ReportGeneratorRef() *ReportGeneratorReference {
	var returns *ReportGeneratorReference
	_jsii_.Get(
		j,
		"reportGeneratorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGeneratorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGeneratorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


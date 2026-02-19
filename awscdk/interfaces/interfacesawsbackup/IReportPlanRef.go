package interfacesawsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportPlan.
// Experimental.
type IReportPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReportPlan resource.
	// Experimental.
	ReportPlanRef() *ReportPlanReference
}

// The jsii proxy for IReportPlanRef
type jsiiProxy_IReportPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReportPlanRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReportPlanRef) ReportPlanRef() *ReportPlanReference {
	var returns *ReportPlanReference
	_jsii_.Get(
		j,
		"reportPlanRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


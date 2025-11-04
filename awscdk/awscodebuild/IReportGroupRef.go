package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportGroup.
// Experimental.
type IReportGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReportGroup resource.
	// Experimental.
	ReportGroupRef() *ReportGroupReference
}

// The jsii proxy for IReportGroupRef
type jsiiProxy_IReportGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IReportGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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


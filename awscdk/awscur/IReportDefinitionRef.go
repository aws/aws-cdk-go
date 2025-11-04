package awscur

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscur/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportDefinition.
// Experimental.
type IReportDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReportDefinition resource.
	// Experimental.
	ReportDefinitionRef() *ReportDefinitionReference
}

// The jsii proxy for IReportDefinitionRef
type jsiiProxy_IReportDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IReportDefinitionRef) ReportDefinitionRef() *ReportDefinitionReference {
	var returns *ReportDefinitionReference
	_jsii_.Get(
		j,
		"reportDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


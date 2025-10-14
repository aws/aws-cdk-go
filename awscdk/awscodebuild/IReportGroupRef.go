package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportGroup.
// Experimental.
type IReportGroupRef interface {
	constructs.IConstruct
	// A reference to a ReportGroup resource.
	// Experimental.
	ReportGroupRef() *ReportGroupReference
}

// The jsii proxy for IReportGroupRef
type jsiiProxy_IReportGroupRef struct {
	internal.Type__constructsIConstruct
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


package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportPlan.
// Experimental.
type IReportPlanRef interface {
	constructs.IConstruct
	// A reference to a ReportPlan resource.
	// Experimental.
	ReportPlanRef() *ReportPlanReference
}

// The jsii proxy for IReportPlanRef
type jsiiProxy_IReportPlanRef struct {
	internal.Type__constructsIConstruct
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


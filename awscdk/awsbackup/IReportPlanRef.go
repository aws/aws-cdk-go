package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReportPlan.
// Experimental.
type IReportPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReportPlanRef
type jsiiProxy_IReportPlanRef struct {
	internal.Type__constructsIConstruct
}


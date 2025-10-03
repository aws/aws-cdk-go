package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RefreshSchedule.
// Experimental.
type IRefreshScheduleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRefreshScheduleRef
type jsiiProxy_IRefreshScheduleRef struct {
	internal.Type__constructsIConstruct
}


package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RefreshSchedule.
// Experimental.
type IRefreshScheduleRef interface {
	constructs.IConstruct
	// A reference to a RefreshSchedule resource.
	// Experimental.
	RefreshScheduleRef() *RefreshScheduleReference
}

// The jsii proxy for IRefreshScheduleRef
type jsiiProxy_IRefreshScheduleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRefreshScheduleRef) RefreshScheduleRef() *RefreshScheduleReference {
	var returns *RefreshScheduleReference
	_jsii_.Get(
		j,
		"refreshScheduleRef",
		&returns,
	)
	return returns
}


package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// A FlowLog.
// Experimental.
type IFlowLog interface {
	awscdk.IResource
	// The Id of the VPC Flow Log.
	// Experimental.
	FlowLogId() *string
}

// The jsii proxy for IFlowLog
type jsiiProxy_IFlowLog struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFlowLog) FlowLogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogId",
		&returns,
	)
	return returns
}


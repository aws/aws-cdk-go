package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A FlowLog.
type IFlowLog interface {
	awscdk.IResource
	// The Id of the VPC Flow Log.
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

